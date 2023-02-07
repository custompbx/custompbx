
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchDialplanDirectoryParameter,
  GetDialplanDirectory,
  StoreDelDialplanDirectoryParameter,
  StoreSwitchDialplanDirectoryParameter,
  UpdateDialplanDirectoryParameter,
  StoreGetDialplanDirectory,
  StoreAddDialplanDirectoryParameter,
  DelDialplanDirectoryParameter,
  StoreUpdateDialplanDirectoryParameter,
  StoreGotDialplanDirectoryError,
  AddDialplanDirectoryParameter,
} from './config.actions.dialplan_directory';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsDialplanDirectory {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetDialplanDirectory: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetDialplanDirectory),
      map((action: GetDialplanDirectory) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDialplanDirectoryError({error: response.error});
              }
              return new StoreGetDialplanDirectory({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDialplanDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDialplanDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDialplanDirectoryParameter),
      map((action: UpdateDialplanDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDialplanDirectoryError({error: response.error});
              }
              return new StoreUpdateDialplanDirectoryParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDialplanDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDialplanDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchDialplanDirectoryParameter),
      map((action: SwitchDialplanDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDialplanDirectoryError({error: response.error});
              }
              return new StoreSwitchDialplanDirectoryParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDialplanDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  AddDialplanDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDialplanDirectoryParameter),
      map((action: AddDialplanDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDialplanDirectoryError({error: response.error});
              }
              return new StoreAddDialplanDirectoryParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDialplanDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  DelDialplanDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDialplanDirectoryParameter),
      map((action: DelDialplanDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDialplanDirectoryError({error: response.error});
              }
              return new StoreDelDialplanDirectoryParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDialplanDirectoryError({error: error}));
            }),
          );
        }
      ));
  });
}

