
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchHttpCacheParameter,
  GetHttpCache,
  StoreDelHttpCacheParameter,
  StoreSwitchHttpCacheParameter,
  UpdateHttpCacheParameter,
  StoreGetHttpCache,
  StoreAddHttpCacheParameter,
  DelHttpCacheParameter,
  StoreUpdateHttpCacheParameter,
  StoreGotHttpCacheError,
  AddHttpCacheParameter,
} from './config.actions.http_cache';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsHttpCache {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetHttpCache: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetHttpCache),
      map((action: GetHttpCache) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreGetHttpCache({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateHttpCacheParameter),
      map((action: UpdateHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreUpdateHttpCacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchHttpCacheParameter),
      map((action: SwitchHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreSwitchHttpCacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  AddHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddHttpCacheParameter),
      map((action: AddHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreAddHttpCacheParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  DelHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelHttpCacheParameter),
      map((action: DelHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreDelHttpCacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });
}

