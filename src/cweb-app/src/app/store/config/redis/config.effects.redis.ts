
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchRedisParameter,
  GetRedis,
  StoreDelRedisParameter,
  StoreSwitchRedisParameter,
  UpdateRedisParameter,
  StoreGetRedis,
  StoreAddRedisParameter,
  DelRedisParameter,
  StoreUpdateRedisParameter,
  StoreGotRedisError,
  AddRedisParameter,
} from './config.actions.redis';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsRedis {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetRedis: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetRedis),
      map((action: GetRedis) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotRedisError({error: response.error});
              }
              return new StoreGetRedis({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotRedisError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateRedisParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateRedisParameter),
      map((action: UpdateRedisParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotRedisError({error: response.error});
              }
              return new StoreUpdateRedisParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotRedisError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchRedisParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchRedisParameter),
      map((action: SwitchRedisParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotRedisError({error: response.error});
              }
              return new StoreSwitchRedisParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotRedisError({error: error}));
            }),
          );
        }
      ));
  });

  AddRedisParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddRedisParameter),
      map((action: AddRedisParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotRedisError({error: response.error});
              }
              return new StoreAddRedisParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotRedisError({error: error}));
            }),
          );
        }
      ));
  });

  DelRedisParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelRedisParameter),
      map((action: DelRedisParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotRedisError({error: response.error});
              }
              return new StoreDelRedisParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotRedisError({error: error}));
            }),
          );
        }
      ));
  });
}

