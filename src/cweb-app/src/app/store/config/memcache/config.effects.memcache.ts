
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchMemcacheParameter,
  GetMemcache,
  StoreDelMemcacheParameter,
  StoreSwitchMemcacheParameter,
  UpdateMemcacheParameter,
  StoreGetMemcache,
  StoreAddMemcacheParameter,
  DelMemcacheParameter,
  StoreUpdateMemcacheParameter,
  StoreGotMemcacheError,
  AddMemcacheParameter,
} from './config.actions.memcache';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsMemcache {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetMemcache: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetMemcache),
      map((action: GetMemcache) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMemcacheError({error: response.error});
              }
              return new StoreGetMemcache({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMemcacheError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateMemcacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateMemcacheParameter),
      map((action: UpdateMemcacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMemcacheError({error: response.error});
              }
              return new StoreUpdateMemcacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMemcacheError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchMemcacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchMemcacheParameter),
      map((action: SwitchMemcacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMemcacheError({error: response.error});
              }
              return new StoreSwitchMemcacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMemcacheError({error: error}));
            }),
          );
        }
      ));
  });

  AddMemcacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddMemcacheParameter),
      map((action: AddMemcacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMemcacheError({error: response.error});
              }
              return new StoreAddMemcacheParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMemcacheError({error: error}));
            }),
          );
        }
      ));
  });

  DelMemcacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelMemcacheParameter),
      map((action: DelMemcacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMemcacheError({error: response.error});
              }
              return new StoreDelMemcacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMemcacheError({error: error}));
            }),
          );
        }
      ));
  });
}

