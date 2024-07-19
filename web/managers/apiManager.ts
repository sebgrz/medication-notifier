import CookieManager from "./cookieManager";

class ApiManager {
  private static BASE_URL = "http://localhost:8080"

  static authRegister = async (username: string, password: string): Promise<boolean> => {
    const body = JSON.stringify({ username: username, password: password });
    const clientId = crypto.randomUUID();

    try {
      await fetch(ApiManager.BASE_URL + "/api/auth/register",
        { method: "POST", body: body, headers: { 'X-Client-Id': clientId, "User-Agent": window.navigator.userAgent } }
      );
    } catch {
      return false;
    }
    CookieManager.setCookie("client-id", clientId, 5)
    return true;
  }
}

export default ApiManager;
