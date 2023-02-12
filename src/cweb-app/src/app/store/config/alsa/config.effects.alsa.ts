import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchAlsaParameter,
  GetAlsa,
  StoreDelAlsaParameter,
  StoreSwitchAlsaParameter,
  UpdateAlsaParameter,
  StoreGetAlsa,
  StoreAddAlsaParameter,
  DelAlsaParameter,
  StoreUpdateAlsaParameter,
  StoreGotAlsaError,
  AddAlsaParameter,
} from './config.actions.alsa';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsAlsa {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAlsa: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetAlsa),
      map((action: GetAlsa) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAlsaError({error: response.error});
              }
              return new StoreGetAlsa({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAlsaError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAlsaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateAlsaParameter),
      map((action: UpdateAlsaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAlsaError({error: response.error});
              }
              return new StoreUpdateAlsaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAlsaError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAlsaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchAlsaParameter),
      map((action: SwitchAlsaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAlsaError({error: response.error});
              }
              return new StoreSwitchAlsaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAlsaError({error: error}));
            }),
          );
        }
      ));
  });

  AddAlsaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddAlsaParameter),
      map((action: AddAlsaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAlsaError({error: response.error});
              }
              return new StoreAddAlsaParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAlsaError({error: error}));
            }),
          );
        }
      ));
  });

  DelAlsaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelAlsaParameter),
      map((action: DelAlsaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAlsaError({error: response.error});
              }
              return new StoreDelAlsaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAlsaError({error: error}));
            }),
          );
        }
      ));
  });
}
