
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchLuaParameter,
  GetLua,
  StoreDelLuaParameter,
  StoreSwitchLuaParameter,
  UpdateLuaParameter,
  StoreGetLua,
  StoreAddLuaParameter,
  DelLuaParameter,
  StoreUpdateLuaParameter,
  StoreGotLuaError,
  AddLuaParameter,
} from './config.actions.lua';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsLua {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetLua: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetLua),
      map((action: GetLua) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLuaError({error: response.error});
              }
              return new StoreGetLua({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLuaError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateLuaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateLuaParameter),
      map((action: UpdateLuaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLuaError({error: response.error});
              }
              return new StoreUpdateLuaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLuaError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchLuaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchLuaParameter),
      map((action: SwitchLuaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLuaError({error: response.error});
              }
              return new StoreSwitchLuaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLuaError({error: error}));
            }),
          );
        }
      ));
  });

  AddLuaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddLuaParameter),
      map((action: AddLuaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLuaError({error: response.error});
              }
              return new StoreAddLuaParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLuaError({error: error}));
            }),
          );
        }
      ));
  });

  DelLuaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelLuaParameter),
      map((action: DelLuaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLuaError({error: response.error});
              }
              return new StoreDelLuaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLuaError({error: error}));
            }),
          );
        }
      ));
  });
}

