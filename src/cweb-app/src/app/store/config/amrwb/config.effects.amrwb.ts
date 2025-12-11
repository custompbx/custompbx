
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchAmrwbParameter,
  GetAmrwb,
  StoreDelAmrwbParameter,
  StoreSwitchAmrwbParameter,
  UpdateAmrwbParameter,
  StoreGetAmrwb,
  StoreAddAmrwbParameter,
  DelAmrwbParameter,
  StoreUpdateAmrwbParameter,
  StoreGotAmrwbError,
  AddAmrwbParameter,
} from './config.actions.amrwb';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsAmrwb {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAmrwb: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetAmrwb),
      map((action: GetAmrwb) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrwbError({error: response.error});
              }
              return new StoreGetAmrwb({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrwbError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAmrwbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateAmrwbParameter),
      map((action: UpdateAmrwbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrwbError({error: response.error});
              }
              return new StoreUpdateAmrwbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrwbError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAmrwbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchAmrwbParameter),
      map((action: SwitchAmrwbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrwbError({error: response.error});
              }
              return new StoreSwitchAmrwbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrwbError({error: error}));
            }),
          );
        }
      ));
  });

  AddAmrwbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddAmrwbParameter),
      map((action: AddAmrwbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrwbError({error: response.error});
              }
              return new StoreAddAmrwbParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrwbError({error: error}));
            }),
          );
        }
      ));
  });

  DelAmrwbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelAmrwbParameter),
      map((action: DelAmrwbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrwbError({error: response.error});
              }
              return new StoreDelAmrwbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrwbError({error: error}));
            }),
          );
        }
      ));
  });
}

