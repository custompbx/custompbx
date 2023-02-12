
import { Action } from '@ngrx/store';
import {StoreImportAllModules} from '../config.actions';

export enum ConfigActionTypes {
  GET_CDR_PG_CSV = '[Config][Get] Cdr_Pg_Csv',
  STORE_GET_CDR_PG_CSV = '[Config][Store][Get] Cdr_Pg_Csv',
  STORE_NEW_CDR_PG_CSV_PARAMETER = '[Config][Store][New] Cdr_Pg_Csv Parameter',
  STORE_DROP_NEW_CDR_PG_CSV_PARAMETER = '[Config][Store][New][Drop] Cdr_Pg_Csv Parameter',
  STORE_NEW_CDR_PG_CSV_FIELD = '[Config][Store][New] Cdr_Pg_Csv Field',
  STORE_DROP_NEW_CDR_PG_CSV_FIELD = '[Config][Store][New][Drop] Cdr_Pg_Csv Field',
  ADD_CDR_PG_CSV_PARAMETER = '[Config][Add] Cdr_Pg_Csv Parameter',
  STORE_ADD_CDR_PG_CSV_PARAMETER = '[Config][Store][Add] Cdr_Pg_Csv Parameter',
  ADD_CDR_PG_CSV_FIELD = '[Config][Add] Cdr_Pg_Csv Field',
  STORE_ADD_CDR_PG_CSV_FIELD = '[Config][Store][Add] Cdr_Pg_Csv Field',
  UPDATE_CDR_PG_CSV_PARAMETER = '[Config][Update] Cdr_Pg_Csv Parameter',
  STORE_UPDATE_CDR_PG_CSV_PARAMETER = '[Config][Store][Update] Cdr_Pg_Csv Parameter',
  SWITCH_CDR_PG_CSV_PARAMETER = '[Config][Switch] Cdr_Pg_Csv Parameter',
  STORE_SWITCH_CDR_PG_CSV_PARAMETER = '[Config][Store][Switch] Cdr_Pg_Csv Parameter',
  DELETE_CDR_PG_CSV_PARAMETER = '[Config][Delete] Cdr_Pg_Csv Parameter',
  STORE_DELETE_CDR_PG_CSV_PARAMETER = '[Config][Store][Delete] Cdr_Pg_Csv Parameter',
  UPDATE_CDR_PG_CSV_FIELD = '[Config][Update] Cdr_Pg_Csv Field',
  STORE_UPDATE_CDR_PG_CSV_FIELD = '[Config][Store][Update] Cdr_Pg_Csv Field',
  SWITCH_CDR_PG_CSV_FIELD = '[Config][Switch] Cdr_Pg_Csv Field',
  STORE_SWITCH_CDR_PG_CSV_FIELD = '[Config][Store][Switch] Cdr_Pg_Csv Field',
  DELETE_CDR_PG_CSV_FIELD = '[Config][Delete] Cdr_Pg_Csv Field',
  STORE_DELETE_CDR_PG_CSV_FIELD = '[Config][Store][Delete] Cdr_Pg_Csv Field',
  StoreGotCdrPgCsvError = 'StoreGotCdrPgCsvError',
}

export class StoreGotCdrPgCsvError implements Action {
  readonly type = ConfigActionTypes.StoreGotCdrPgCsvError;
  constructor(public payload: any) {}
}

export class GetCdrPgCsv implements Action {
  readonly type = ConfigActionTypes.GET_CDR_PG_CSV;
  constructor(public payload: any) {}
}

export class StoreGetCdrPgCsv implements Action {
  readonly type = ConfigActionTypes.STORE_GET_CDR_PG_CSV;
  constructor(public payload: any) {}
}

export class StoreNewCdrPgCsvParam implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class StoreDropNewCdrPgCsvParam implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class StoreNewCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class StoreDropNewCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class AddCdrPgCsvParam implements Action {
  readonly type = ConfigActionTypes.ADD_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class StoreAddCdrPgCsvParam implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class AddCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.ADD_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class StoreAddCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class UpdateCdrPgCsvParameter implements Action {
  readonly type = ConfigActionTypes.UPDATE_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class StoreUpdateCdrPgCsvParameter implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class SwitchCdrPgCsvParameter implements Action {
  readonly type = ConfigActionTypes.SWITCH_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class StoreSwitchCdrPgCsvParameter implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class DeleteCdrPgCsvParameter implements Action {
  readonly type = ConfigActionTypes.DELETE_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class StoreDeleteCdrPgCsvParameter implements Action {
  readonly type = ConfigActionTypes.STORE_DELETE_CDR_PG_CSV_PARAMETER;
  constructor(public payload: any) {}
}

export class UpdateCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.UPDATE_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class StoreUpdateCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class SwitchCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.SWITCH_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class StoreSwitchCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class DeleteCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.DELETE_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export class StoreDeleteCdrPgCsvField implements Action {
  readonly type = ConfigActionTypes.STORE_DELETE_CDR_PG_CSV_FIELD;
  constructor(public payload: any) {}
}

export type All =
  | GetCdrPgCsv
  | StoreGetCdrPgCsv
  | StoreNewCdrPgCsvParam
  | StoreDropNewCdrPgCsvParam
  | StoreNewCdrPgCsvField
  | StoreDropNewCdrPgCsvField
  | AddCdrPgCsvParam
  | StoreAddCdrPgCsvParam
  | AddCdrPgCsvField
  | StoreAddCdrPgCsvField
  | UpdateCdrPgCsvParameter
  | StoreUpdateCdrPgCsvParameter
  | SwitchCdrPgCsvParameter
  | StoreSwitchCdrPgCsvParameter
  | DeleteCdrPgCsvParameter
  | StoreDeleteCdrPgCsvParameter
  | UpdateCdrPgCsvField
  | StoreUpdateCdrPgCsvField
  | SwitchCdrPgCsvField
  | StoreSwitchCdrPgCsvField
  | DeleteCdrPgCsvField
  | StoreDeleteCdrPgCsvField
  | StoreGotCdrPgCsvError
;

