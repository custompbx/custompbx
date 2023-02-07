import { Injectable } from '@angular/core';
import {CookieService} from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class CookiesStorageService {

  constructor(
    private cookieService: CookieService
  ) { }


  public getToken(): string {
    return  this.cookieService.get('token');
  }

  public setToken(token: string): void {
    this.cookieService.set('token', token, null, '/', null, false, null);
  }

  public delToken(): void {
    this.cookieService.delete('token', '/');
  }
}
