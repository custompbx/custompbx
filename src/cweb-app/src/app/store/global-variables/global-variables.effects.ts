import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchGlobalVariable,
  GetGlobalVariables,
  StoreDelGlobalVariable,
  StoreSwitchGlobalVariable,
  UpdateGlobalVariable,
  StoreGetGlobalVariables,
  StoreAddGlobalVariable,
  DelGlobalVariable,
  StoreUpdateGlobalVariable,
  StoreGotGlobalVariableError,
  AddGlobalVariable, ImportGlobalVariables, StoreImportGlobalVariables, MoveGlobalVariable, StoreMoveGlobalVariable,
} from './global-variables.actions';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../services/ws-data.service';

@Injectable()
export class GlobalVariablesEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetGlobalVariables: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetGlobalVariables),
      map((action: GetGlobalVariables) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreGetGlobalVariables({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateGlobalVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateGlobalVariable),
      map((action: UpdateGlobalVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreUpdateGlobalVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchGlobalVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchGlobalVariable),
      map((action: SwitchGlobalVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreSwitchGlobalVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });

  AddGlobalVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddGlobalVariable),
      map((action: AddGlobalVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreAddGlobalVariable({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });

  DelGlobalVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelGlobalVariable),
      map((action: DelGlobalVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreDelGlobalVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });

  ImportGlobalVariables: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ImportGlobalVariables),
      map((action: ImportGlobalVariables) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreImportGlobalVariables({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });

  MoveGlobalVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.MoveGlobalVariable),
      map((action: MoveGlobalVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotGlobalVariableError({error: response.error});
              }
              return new StoreMoveGlobalVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotGlobalVariableError({error: error}));
            }),
          );
        }
      ));
  });
}

