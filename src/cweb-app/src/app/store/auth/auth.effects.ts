import {Injectable} from '@angular/core';
import {Router} from '@angular/router';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {map, switchMap, catchError} from 'rxjs/operators';
import {
  AuthActionTypes,
  LogIn,
  LogInSuccess,
  LogInFailure,
  LogOut,
  GetStatus,
  ReLogIn,
} from './auth.actions';
import {WsDataService} from '../../services/ws-data.service';
import {CookiesStorageService} from '../../services/cookies-storage.service';

@Injectable({
  providedIn: 'root'
})
export class AuthEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
    private router: Router,
    private cookie: CookiesStorageService
  ) {}

  LogIn: Observable<any> = createEffect(() => {
    return this.actions.pipe(
        ofType(AuthActionTypes.LOGIN),
        map((action: LogIn) => action.payload),
        switchMap(payload => {
          return this.ws.logIn(payload.login, payload.password).pipe(
            map((response) => {
              if (!response.token || response.error) {
                return new LogInFailure({error: response.error});
              }
              return new LogInSuccess({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new LogInFailure({error: error}));
            })
          );
        })
      );
  });

  LogInSuccess: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AuthActionTypes.LOGIN_SUCCESS),
      map((user: any) => {
        this.cookie.setToken(user.payload.response.token);
        if (user.payload.route) {
          this.router.navigateByUrl(user.payload.route);
        } else {
          this.router.navigateByUrl('dashboard');
        }
      })
    );
  }, {dispatch: false});

  LogInFailure: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AuthActionTypes.LOGIN_FAILURE),
      map((error) => {
        this.router.navigateByUrl('login');
      }),
    );
  }, { dispatch: false });

  LogOut: Observable<any> =  createEffect(() => {
    return this.actions.pipe(
      ofType(AuthActionTypes.LOGOUT),
      map((user) => {
        this.ws.universalSender(AuthActionTypes.LOGOUT, null).pipe();
        this.cookie.delToken();
        this.router.navigateByUrl('login');
      })
    );
  }, { dispatch: false });

  GetStatus: Observable<any> =  createEffect(() => {
    return this.actions.pipe(
      ofType(AuthActionTypes.GET_STATUS),
      map((action: GetStatus) => action),
      switchMap(payload => {
        return this.ws.getStatus();
      })
    );
  }, { dispatch: false });

  ReLogIn:  Observable<any> = createEffect(() => {
    return this.actions.pipe(
        ofType(AuthActionTypes.RELOGIN),
        map((action: ReLogIn) => action.payload),
        switchMap(payload => {
          return this.ws.getUserByToken().pipe(
            map((response) => {
              if (!response.token || response.error) {
                return new LogInFailure({error: response.error});
              }
              let route = payload.route;
              if (payload.route === 'login' || payload.route === '' || payload.route === '/login' || payload.route === '/') {
                route = 'dashboard';
              }

              return new LogInSuccess({response: response, route: route});
            }),
            catchError((error) => {
              console.log(error);

              return of(new LogInFailure({error: error}));
            }),
          );
        })
      );
  });
}
