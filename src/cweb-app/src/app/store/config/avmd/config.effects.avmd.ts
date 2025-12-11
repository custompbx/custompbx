
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchAvmdParameter,
  GetAvmd,
  StoreDelAvmdParameter,
  StoreSwitchAvmdParameter,
  UpdateAvmdParameter,
  StoreGetAvmd,
  StoreAddAvmdParameter,
  DelAvmdParameter,
  StoreUpdateAvmdParameter,
  StoreGotAvmdError,
  AddAvmdParameter,
} from './config.actions.avmd';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsAvmd {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAvmd: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetAvmd),
      map((action: GetAvmd) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAvmdError({error: response.error});
              }
              return new StoreGetAvmd({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAvmdError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAvmdParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateAvmdParameter),
      map((action: UpdateAvmdParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAvmdError({error: response.error});
              }
              return new StoreUpdateAvmdParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAvmdError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAvmdParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchAvmdParameter),
      map((action: SwitchAvmdParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAvmdError({error: response.error});
              }
              return new StoreSwitchAvmdParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAvmdError({error: error}));
            }),
          );
        }
      ));
  });

  AddAvmdParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddAvmdParameter),
      map((action: AddAvmdParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAvmdError({error: response.error});
              }
              return new StoreAddAvmdParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAvmdError({error: error}));
            }),
          );
        }
      ));
  });

  DelAvmdParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelAvmdParameter),
      map((action: DelAvmdParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAvmdError({error: response.error});
              }
              return new StoreDelAvmdParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAvmdError({error: error}));
            }),
          );
        }
      ));
  });
}

