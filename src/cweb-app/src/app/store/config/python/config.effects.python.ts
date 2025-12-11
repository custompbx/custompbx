
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchPythonParameter,
  GetPython,
  StoreDelPythonParameter,
  StoreSwitchPythonParameter,
  UpdatePythonParameter,
  StoreGetPython,
  StoreAddPythonParameter,
  DelPythonParameter,
  StoreUpdatePythonParameter,
  StoreGotPythonError,
  AddPythonParameter,
} from './config.actions.python';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsPython {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetPython: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetPython),
      map((action: GetPython) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPythonError({error: response.error});
              }
              return new StoreGetPython({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPythonError({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePythonParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePythonParameter),
      map((action: UpdatePythonParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPythonError({error: response.error});
              }
              return new StoreUpdatePythonParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPythonError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPythonParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPythonParameter),
      map((action: SwitchPythonParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPythonError({error: response.error});
              }
              return new StoreSwitchPythonParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPythonError({error: error}));
            }),
          );
        }
      ));
  });

  AddPythonParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPythonParameter),
      map((action: AddPythonParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPythonError({error: response.error});
              }
              return new StoreAddPythonParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPythonError({error: error}));
            }),
          );
        }
      ));
  });

  DelPythonParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPythonParameter),
      map((action: DelPythonParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPythonError({error: response.error});
              }
              return new StoreDelPythonParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPythonError({error: error}));
            }),
          );
        }
      ));
  });
}

