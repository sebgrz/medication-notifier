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

  private static getRequiredHeaders = (): { [key: string]: string; } => {
    const clientId = CookieManager.get("client-id") ?? crypto.randomUUID();
    Headers
    return {
      [Header.CLIENT_ID]: clientId,
      [Header.USER_AGENT]: window.navigator.userAgent
    };
  }
}

export default ApiManager;
