import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddOdbcCdrField,
  AddOdbcCdrParameter,
  AddOdbcCdrTable,
  DeleteOdbcCdrField,
  DeleteOdbcCdrParameter,
  DeleteOdbcCdrTable,
  GetOdbcCdr,
  StoreAddOdbcCdrField,
  StoreAddOdbcCdrParameter,
  StoreAddOdbcCdrTable,
  StoreDeleteOdbcCdrField,
  StoreDeleteOdbcCdrParameter,
  StoreDeleteOdbcCdrTable,
  StoreGetOdbcCdr,
  StoreGotOdbcCdrError,
  StoreSwitchOdbcCdrField,
  StoreSwitchOdbcCdrParameter,
  StoreSwitchOdbcCdrTable,
  StoreUpdateOdbcCdrField,
  StoreUpdateOdbcCdrParameter,
  StoreUpdateOdbcCdrTable,
  SwitchOdbcCdrField,
  SwitchOdbcCdrParameter,
  SwitchOdbcCdrTable,
  UpdateOdbcCdrField,
  UpdateOdbcCdrParameter,
  UpdateOdbcCdrTable, StoreGetOdbcCdrField, GetOdbcCdrField
} from './config.actions.odbc-cdr';
import {Failure} from '../config.actions';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsOdbcCdr {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetOdbcCdr: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOdbcCdr),
      map((action: GetOdbcCdr) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreGetOdbcCdr({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOdbcCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOdbcCdrParameter),
      map((action: UpdateOdbcCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreUpdateOdbcCdrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOdbcCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOdbcCdrParameter),
      map((action: SwitchOdbcCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreSwitchOdbcCdrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteOdbcCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DeleteOdbcCdrParameter),
      map((action: DeleteOdbcCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreDeleteOdbcCdrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddOdbcCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOdbcCdrParameter),
      map((action: AddOdbcCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreAddOdbcCdrParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddOdbcCdrTable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOdbcCdrTable),
      map((action: AddOdbcCdrTable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreAddOdbcCdrTable({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOdbcCdrTable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOdbcCdrTable),
      map((action: UpdateOdbcCdrTable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreUpdateOdbcCdrTable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOdbcCdrTable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOdbcCdrTable),
      map((action: SwitchOdbcCdrTable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreSwitchOdbcCdrTable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteOdbcCdrTable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DeleteOdbcCdrTable),
      map((action: DeleteOdbcCdrTable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreDeleteOdbcCdrTable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddOdbcCdrField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOdbcCdrField),
      map((action: AddOdbcCdrField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreAddOdbcCdrField({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOdbcCdrField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOdbcCdrField),
      map((action: UpdateOdbcCdrField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreUpdateOdbcCdrField({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOdbcCdrField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOdbcCdrField),
      map((action: SwitchOdbcCdrField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreSwitchOdbcCdrField({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteOdbcCdrField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DeleteOdbcCdrField),
      map((action: DeleteOdbcCdrField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreDeleteOdbcCdrField({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetOdbcCdrField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOdbcCdrField),
      map((action: GetOdbcCdrField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOdbcCdrError({error: response.error});
              }
              return new StoreGetOdbcCdrField({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

}

