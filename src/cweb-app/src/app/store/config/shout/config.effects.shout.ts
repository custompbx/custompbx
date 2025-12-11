
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchShoutParameter,
  GetShout,
  StoreDelShoutParameter,
  StoreSwitchShoutParameter,
  UpdateShoutParameter,
  StoreGetShout,
  StoreAddShoutParameter,
  DelShoutParameter,
  StoreUpdateShoutParameter,
  StoreGotShoutError,
  AddShoutParameter,
} from './config.actions.shout';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsShout {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetShout: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetShout),
      map((action: GetShout) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotShoutError({error: response.error});
              }
              return new StoreGetShout({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotShoutError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateShoutParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateShoutParameter),
      map((action: UpdateShoutParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotShoutError({error: response.error});
              }
              return new StoreUpdateShoutParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotShoutError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchShoutParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchShoutParameter),
      map((action: SwitchShoutParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotShoutError({error: response.error});
              }
              return new StoreSwitchShoutParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotShoutError({error: error}));
            }),
          );
        }
      ));
  });

  AddShoutParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddShoutParameter),
      map((action: AddShoutParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotShoutError({error: response.error});
              }
              return new StoreAddShoutParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotShoutError({error: error}));
            }),
          );
        }
      ));
  });

  DelShoutParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelShoutParameter),
      map((action: DelShoutParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotShoutError({error: response.error});
              }
              return new StoreDelShoutParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotShoutError({error: error}));
            }),
          );
        }
      ));
  });
}

