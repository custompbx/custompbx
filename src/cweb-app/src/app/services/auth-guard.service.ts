import {Injectable} from '@angular/core';
import {Router, UrlTree} from '@angular/router';
import {WsDataService} from './ws-data.service';
import { RouterStateSnapshot, ActivatedRouteSnapshot } from '@angular/router';
import {CookiesStorageService} from './cookies-storage.service';
import {Observable} from 'rxjs';
import {filter, map, take} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService  {

  constructor(
    public router: Router,
    private ws: WsDataService,
    private cookie: CookiesStorageService,
  ) {}

  canActivate(_route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean | UrlTree | Observable<boolean | UrlTree> {
    const isPublicRoute = state.url === '' || state.url === '/' || state.url === '/login';
    if (isPublicRoute) return true;
    if (!this.cookie.getToken()) return this.router.createUrlTree(['/login']);

    return this.ws.websocketService.status.pipe(
      filter((connected): connected is boolean => typeof connected === 'boolean'),
      take(1),
      map(connected => connected || this.router.createUrlTree(['/login']))
    );
  }

}
