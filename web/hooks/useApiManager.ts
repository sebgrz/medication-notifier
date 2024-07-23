import { Medication } from "@/components/medicationsPanel";
import CookieManager from "@/managers/cookieManager";
import { useRouter } from "next/navigation";

enum Header {
  CLIENT_ID = "X-Client-Id",
  USER_AGENT = "User-Agent",
  AUTHORIZATION = "Authorization"
}

export const useApiManager = () => {
  const BASE_URL = "http://localhost:8080"

  const router = useRouter();

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
    return await tryCallWithReauthorize("POST", "/api/add", body);
  }

  const tryCallWithReauthorize = async <T,>(method: string, url: string, body?: string): Promise<T | undefined> => {
    const headers = getRequiredHeaders();
    let tokens = getLocalTokens();
    if (!tokens) {
      moveToLoginPage();
      return;
    }
    headers[Header.AUTHORIZATION] = `Bearer ${tokens.auth_token}`;

    try {
      let resp = await fetch(BASE_URL + url,
        { method: method, body: body, headers: headers }
      );
      if (resp.status === 401) {
        let tokens = await authRefresh();
        if (!tokens) {
          moveToLoginPage();
          return;
        }

        // second shot with fresh auth_token
        headers[Header.AUTHORIZATION] = `Bearer ${tokens.auth_token}`;
        resp = await fetch(BASE_URL + url,
          { method: method, body: body, headers: headers }
        );
      }

      return await resp.json();
    } catch {
      return undefined;
    }
  }

  const moveToLoginPage = () => router.push("/auth/login");

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
