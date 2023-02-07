import {Injectable} from '@angular/core';
import {Router, CanActivate} from '@angular/router';
import {Store} from '@ngrx/store';
import { AppState } from '../store/app.states';
import {WsDataService} from './ws-data.service';
import { RouterStateSnapshot, ActivatedRouteSnapshot } from '@angular/router';
import {CookiesStorageService} from './cookies-storage.service';

@Injectable()
export class AuthGuardService implements CanActivate {

  isAuthenticated: false;
  noConnect = false;

  constructor(
    public router: Router,
    private store: Store<AppState>,
    private ws: WsDataService,
    private cookie: CookiesStorageService,
  ) {
    this.ws.websocketService.status
      .subscribe((isConnected) => {
        if (typeof(isConnected) !== 'boolean') {
          return;
        }
        this.noConnect = !isConnected;
      });
  }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    const notLogin = state.url !== 'login' && state.url !== '' && state.url !== '/login' && state.url !== '/';
    if (notLogin && this.noConnect) {
      this.router.navigateByUrl('/login');
      return false;
    }

    if (!this.cookie.getToken() && notLogin) {
      this.router.navigateByUrl('/login');
      return false;
    }

    return true;
  }

}
