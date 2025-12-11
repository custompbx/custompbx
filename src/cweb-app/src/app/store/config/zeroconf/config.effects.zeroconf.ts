
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchZeroconfParameter,
  GetZeroconf,
  StoreDelZeroconfParameter,
  StoreSwitchZeroconfParameter,
  UpdateZeroconfParameter,
  StoreGetZeroconf,
  StoreAddZeroconfParameter,
  DelZeroconfParameter,
  StoreUpdateZeroconfParameter,
  StoreGotZeroconfError,
  AddZeroconfParameter,
} from './config.actions.zeroconf';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsZeroconf {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetZeroconf: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetZeroconf),
      map((action: GetZeroconf) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotZeroconfError({error: response.error});
              }
              return new StoreGetZeroconf({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotZeroconfError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateZeroconfParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateZeroconfParameter),
      map((action: UpdateZeroconfParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotZeroconfError({error: response.error});
              }
              return new StoreUpdateZeroconfParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotZeroconfError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchZeroconfParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchZeroconfParameter),
      map((action: SwitchZeroconfParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotZeroconfError({error: response.error});
              }
              return new StoreSwitchZeroconfParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotZeroconfError({error: error}));
            }),
          );
        }
      ));
  });

  AddZeroconfParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddZeroconfParameter),
      map((action: AddZeroconfParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotZeroconfError({error: response.error});
              }
              return new StoreAddZeroconfParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotZeroconfError({error: error}));
            }),
          );
        }
      ));
  });

  DelZeroconfParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelZeroconfParameter),
      map((action: DelZeroconfParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotZeroconfError({error: response.error});
              }
              return new StoreDelZeroconfParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotZeroconfError({error: error}));
            }),
          );
        }
      ));
  });
}

