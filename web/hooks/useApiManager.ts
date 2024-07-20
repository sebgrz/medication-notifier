import { Medication } from "@/components/medicationsPanel";
import CookieManager from "@/managers/cookieManager";

enum Header {
  CLIENT_ID = "X-Client-Id",
  USER_AGENT = "User-Agent"
}

export const useApiManager = () => {
  const BASE_URL = "http://localhost:8080"

  const authRegister = async (username: string, password: string): Promise<boolean> => {
    const body = JSON.stringify({ username: username, password: password });
    const headers = getRequiredHeaders();

    try {
      await fetch(BASE_URL + "/api/auth/register",
        { method: "POST", body: body, headers: headers }
      );
    } catch {
      return false;
    }
    CookieManager.set("client-id", headers[Header.CLIENT_ID], 5 * 365)
    return true;
  }

  const authLogin = async (username: string, password: string): Promise<boolean> => {
    const body = JSON.stringify({ username: username, password: password });
    const headers = getRequiredHeaders();

    try {
      const resp = await fetch(BASE_URL + "/api/auth/login",
        { method: "POST", body: body, headers: headers }
      );
      const tokens = JSON.stringify(await resp.json());
      storeTokens(tokens);
    } catch {
      return false;
    }
    return true;
  }

  const authRefresh = async (): Promise<{ auth_token: string, refresh_token: string } | undefined> => {
    const tokens = getLocalTokens();
    if (!tokens) {
      return undefined;
    }
    const body = JSON.stringify({ refresh_token: tokens.refresh_token });
    const headers = getRequiredHeaders();

    try {
      const resp = await fetch(BASE_URL + "/api/auth/refresh_token",
        { method: "POST", body: body, headers: headers }
      );
      const tokensObj = await resp.json();
      const tokens = JSON.stringify(tokensObj);
      storeTokens(tokens);

      return tokensObj;
    } catch {
      return undefined;
    }
  }

  const appAddMedication = async (name: string, day: string, timeOfDay: string): Promise<Medication | undefined> => {
    const body = JSON.stringify({ name: name, day: day, time_of_day: timeOfDay });
    const headers = getRequiredHeaders();

    try {
      const resp = await fetch(BASE_URL + "/api/auth/login",
        { method: "POST", body: body, headers: headers }
      );
      const tokens = JSON.stringify(await resp.json());
      storeTokens(tokens);
    } catch {
      return undefined;
    }
    return undefined;
  }

  const getRequiredHeaders = (): { [key: string]: string; } => {
    const clientId = CookieManager.get("client-id") ?? crypto.randomUUID();
    Headers
    return {
      [Header.CLIENT_ID]: clientId,
      [Header.USER_AGENT]: window.navigator.userAgent
    };
  }

  const storeTokens = (data: string) =>
    CookieManager.set("token", data, 7);

  const getLocalTokens = (): { auth_token: string, refresh_token: string } | undefined => {
    const cookieToken = CookieManager.get("token");
    if (!cookieToken) {
      return undefined;
    }
    return JSON.parse(cookieToken);
  }

  return { authLogin, authRegister, authRefresh, appAddMedication }
}
