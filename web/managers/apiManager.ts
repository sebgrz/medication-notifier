import CookieManager from "./cookieManager";

enum Header {
  CLIENT_ID = "X-Client-Id",
  USER_AGENT = "User-Agent"
}

class ApiManager {
  private static BASE_URL = "http://localhost:8080"

  static authRegister = async (username: string, password: string): Promise<boolean> => {
    const body = JSON.stringify({ username: username, password: password });
    const headers = ApiManager.getRequiredHeaders();

    try {
      await fetch(ApiManager.BASE_URL + "/api/auth/register",
        { method: "POST", body: body, headers: headers }
      );
    } catch {
      return false;
    }
    CookieManager.set("client-id", headers[Header.CLIENT_ID], 5 * 365)
    return true;
  }

  static authLogin = async (username: string, password: string): Promise<boolean> => {
    const body = JSON.stringify({ username: username, password: password });
    const headers = ApiManager.getRequiredHeaders();

    try {
      const resp = await fetch(ApiManager.BASE_URL + "/api/auth/login",
        { method: "POST", body: body, headers: headers }
      );
      const tokens = JSON.stringify(await resp.json());
      ApiManager.storeTokens(tokens);
    } catch {
      return false;
    }
    return true;
  }

  static authRefresh = async (): Promise<{ auth_token: string, refresh_token: string } | undefined> => {
    const tokens = ApiManager.getLocalTokens();
    if (!tokens) {
      return undefined;
    }
    const body = JSON.stringify({ refresh_token: tokens.refresh_token });
    const headers = ApiManager.getRequiredHeaders();

    try {
      const resp = await fetch(ApiManager.BASE_URL + "/api/auth/refresh_token",
        { method: "POST", body: body, headers: headers }
      );
      const tokensObj = await resp.json();
      const tokens = JSON.stringify(tokensObj);
      ApiManager.storeTokens(tokens);

      return tokensObj;
    } catch {
      return undefined;
    }
  }

  private static getRequiredHeaders = (): { [key: string]: string; } => {
    const clientId = CookieManager.get("client-id") ?? crypto.randomUUID();
    Headers
    return {
      [Header.CLIENT_ID]: clientId,
      [Header.USER_AGENT]: window.navigator.userAgent
    };
  }

  private static storeTokens = (data: string) =>
    CookieManager.set("token", data, 7);

  private static getLocalTokens = (): { auth_token: string, refresh_token: string } | undefined => {
    const cookieToken = CookieManager.get("token");
    if (!cookieToken) {
      return undefined;
    }
    return JSON.parse(cookieToken);
  }
}

export default ApiManager;
