class CookieManager {
  static set = (cookieName: string, cookieValue: string, ageInDays: number): string | undefined => {
    const year = 60 * 60 * 24 * ageInDays;
    document.cookie = `${cookieName}=${cookieValue}; max-age=${year}; path=/; SameSite=Lax; Secure;`;

    const cookie = document.cookie
      .split("; ")
      .find((row) => row.startsWith(cookieName))
      ?.split("=")[1];
    return cookie;
  }

  static get = (cookieName: string): string | undefined => {
    const cookies = document.cookie?.split(';');
    if (!cookies) {
      return undefined;
    }

    for (let cookie in cookies) {
      const parts = cookie.split('=');

      if (parts[0] === cookieName) {
        return parts[1];
      }
    }

    return undefined;
  }

  static delete = (cookieName: string) =>
    document.cookie = `${cookieName}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`
}

export default CookieManager;
