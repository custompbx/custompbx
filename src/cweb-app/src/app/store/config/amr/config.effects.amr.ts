
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchAmrParameter,
  GetAmr,
  StoreDelAmrParameter,
  StoreSwitchAmrParameter,
  UpdateAmrParameter,
  StoreGetAmr,
  StoreAddAmrParameter,
  DelAmrParameter,
  StoreUpdateAmrParameter,
  StoreGotAmrError,
  AddAmrParameter,
} from './config.actions.amr';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsAmr {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAmr: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetAmr),
      map((action: GetAmr) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrError({error: response.error});
              }
              return new StoreGetAmr({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAmrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateAmrParameter),
      map((action: UpdateAmrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrError({error: response.error});
              }
              return new StoreUpdateAmrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAmrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchAmrParameter),
      map((action: SwitchAmrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrError({error: response.error});
              }
              return new StoreSwitchAmrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrError({error: error}));
            }),
          );
        }
      ));
  });

  AddAmrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddAmrParameter),
      map((action: AddAmrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrError({error: response.error});
              }
              return new StoreAddAmrParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrError({error: error}));
            }),
          );
        }
      ));
  });

  DelAmrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelAmrParameter),
      map((action: DelAmrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAmrError({error: response.error});
              }
              return new StoreDelAmrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAmrError({error: error}));
            }),
          );
        }
      ));
  });
}

