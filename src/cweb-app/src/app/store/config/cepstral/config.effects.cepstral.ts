
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchCepstralParameter,
  GetCepstral,
  StoreDelCepstralParameter,
  StoreSwitchCepstralParameter,
  UpdateCepstralParameter,
  StoreGetCepstral,
  StoreAddCepstralParameter,
  DelCepstralParameter,
  StoreUpdateCepstralParameter,
  StoreGotCepstralError,
  AddCepstralParameter,
} from './config.actions.cepstral';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsCepstral {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCepstral: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCepstral),
      map((action: GetCepstral) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCepstralError({error: response.error});
              }
              return new StoreGetCepstral({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCepstralError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCepstralParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCepstralParameter),
      map((action: UpdateCepstralParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCepstralError({error: response.error});
              }
              return new StoreUpdateCepstralParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCepstralError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCepstralParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchCepstralParameter),
      map((action: SwitchCepstralParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCepstralError({error: response.error});
              }
              return new StoreSwitchCepstralParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCepstralError({error: error}));
            }),
          );
        }
      ));
  });

  AddCepstralParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCepstralParameter),
      map((action: AddCepstralParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCepstralError({error: response.error});
              }
              return new StoreAddCepstralParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCepstralError({error: error}));
            }),
          );
        }
      ));
  });

  DelCepstralParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCepstralParameter),
      map((action: DelCepstralParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCepstralError({error: response.error});
              }
              return new StoreDelCepstralParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCepstralError({error: error}));
            }),
          );
        }
      ));
  });
}

