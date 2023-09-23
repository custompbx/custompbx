import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CookiesStorageService {

  constructor() {
  }


  public getToken(): string {
    return this.getCookie('token');
  }

  public setToken(token: string): void {
    this.setCookie('token', token, null, '/', null, false, null);
  }

  public delToken(): void {
    this.deleteCookie('token', '/');
  }

  // Set a cookie with a specific path
  setCookie(
    name: string,
    value: string,
    expires?: number | Date,
    path?: string,
    domain?: string,
    secure?: boolean,
    sameSite: 'Lax' | 'None' | 'Strict' = 'None'
  ): void {
    let cookieValue = `${encodeURIComponent(name)}=${encodeURIComponent(value)}`;

    if (expires) {
      if (expires instanceof Date) {
        cookieValue += `;expires=${expires.toUTCString()}`;
      } else {
        const expirationDate = new Date();
        expirationDate.setDate(expirationDate.getDate() + expires);
        cookieValue += `;expires=${expirationDate.toUTCString()}`;
      }
    }

    if (path) {
      cookieValue += `;path=${path}`;
    }

    if (domain) {
      cookieValue += `;domain=${domain}`;
    }

    if (secure) {
      cookieValue += `;secure`;
    }

    cookieValue += `;samesite=${sameSite}`;

    document.cookie = cookieValue;
  }

  // Get a cookie by name
  getCookie(name: string): string | null {
    const cookies = document.cookie.split(';').map(cookie => cookie.trim());
    for (const cookie of cookies) {
      const [cookieName, cookieValue] = cookie.split('=');
      if (cookieName === encodeURIComponent(name)) {
        return decodeURIComponent(cookieValue);
      }
    }
    return null;
  }

  // Delete a cookie by name
  deleteCookie(name: string, path?: string, domain?: string) {
    let cookieValue = `${encodeURIComponent(name)}=;`;

    if (path) {
      cookieValue += `path=${path};`;
    }

    if (domain) {
      cookieValue += `domain=${domain};`;
    }

    // Set an expired date to delete the cookie
    cookieValue += 'expires=Thu, 01 Jan 1970 00:00:00 GMT;';

    document.cookie = cookieValue;
  }
}
