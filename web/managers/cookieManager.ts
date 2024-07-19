class CookieManager {
  static setCookie = (cookieName: string, cookieValue: string, ageYears: number): string | undefined => {
    const year = 60 * 60 * 24 * 365 * ageYears;
    document.cookie = `${cookieName}=${cookieValue}; max-age=${year}; path=/; SameSite=Lax; Secure;`;

    const cookie = document.cookie
      .split("; ")
      .find((row) => row.startsWith(cookieName))
      ?.split("=")[1];
    return cookie;
  }
}

export default CookieManager;
