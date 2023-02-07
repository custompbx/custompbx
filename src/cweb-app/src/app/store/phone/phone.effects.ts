import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {map, switchMap, catchError} from 'rxjs/operators';
import {
  AuthActionTypes,
  GetPhoneCreds,
  StoreGetPhoneCreds,
  Failure,
} from './phone.actions';
import {WsDataService} from '../../services/ws-data.service';

@Injectable()
export class PhoneEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {}

  GetPhoneCreds: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AuthActionTypes.GET_PHONE_CREDS),
      map((action: GetPhoneCreds) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreGetPhoneCreds({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
}
