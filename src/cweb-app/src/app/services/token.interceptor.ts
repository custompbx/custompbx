import { Injectable } from '@angular/core';
import { HttpEvent, HttpInterceptor, HttpHandler, HttpRequest, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { Router } from '@angular/router';

import {catchError} from 'rxjs/operators';
import {CookiesStorageService} from './cookies-storage.service';

@Injectable({ providedIn: 'root' })
export class ErrorInterceptor implements HttpInterceptor {
  constructor(private router: Router, private cookiesStorageService: CookiesStorageService) {}
  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {

    return next.handle(request).pipe(
      catchError((response: unknown) => {
        if (response instanceof HttpErrorResponse && response.status === 401) {
          this.cookiesStorageService.delToken();
          void this.router.navigateByUrl('/login');
        }
        return throwError(() => response);
      })
  );
  }
}
