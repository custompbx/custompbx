import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {Failure} from '../config.actions';
import {
  ConfigActionTypes,
  AddCdrPgCsvField,
  AddCdrPgCsvParam,
  DeleteCdrPgCsvField,
  DeleteCdrPgCsvParameter,
  GetCdrPgCsv,
  StoreAddCdrPgCsvField,
  StoreAddCdrPgCsvParam,
  StoreDeleteCdrPgCsvField,
  StoreDeleteCdrPgCsvParameter,
  StoreGetCdrPgCsv,
  StoreGotCdrPgCsvError, StoreSwitchCdrPgCsvField,
  StoreSwitchCdrPgCsvParameter,
  StoreUpdateCdrPgCsvField,
  StoreUpdateCdrPgCsvParameter,
  SwitchCdrPgCsvField,
  SwitchCdrPgCsvParameter,
  UpdateCdrPgCsvField,
  UpdateCdrPgCsvParameter
} from './config.actions.cdr-pg-csv';

@Injectable()
export class ConfigEffectsCdrPgCsv {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCdrPgCsv: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_CDR_PG_CSV),
      map((action: GetCdrPgCsv) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreGetCdrPgCsv({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddCdrPgCsvParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_CDR_PG_CSV_PARAMETER),
      map((action: AddCdrPgCsvParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreAddCdrPgCsvParam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddCdrPgCsvField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_CDR_PG_CSV_FIELD),
      map((action: AddCdrPgCsvField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreAddCdrPgCsvField({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCdrPgCsvParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_CDR_PG_CSV_PARAMETER),
      map((action: UpdateCdrPgCsvParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreUpdateCdrPgCsvParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCdrPgCsvParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_CDR_PG_CSV_PARAMETER),
      map((action: SwitchCdrPgCsvParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreSwitchCdrPgCsvParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteCdrPgCsvParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DELETE_CDR_PG_CSV_PARAMETER),
      map((action: DeleteCdrPgCsvParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreDeleteCdrPgCsvParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCdrPgCsvField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_CDR_PG_CSV_FIELD),
      map((action: UpdateCdrPgCsvField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreUpdateCdrPgCsvField({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCdrPgCsvField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_CDR_PG_CSV_FIELD),
      map((action: SwitchCdrPgCsvField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreSwitchCdrPgCsvField({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteCdrPgCsvField: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DELETE_CDR_PG_CSV_FIELD),
      map((action: DeleteCdrPgCsvField) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrPgCsvError({error: response.error});
              }
              return new StoreDeleteCdrPgCsvField({response});
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

