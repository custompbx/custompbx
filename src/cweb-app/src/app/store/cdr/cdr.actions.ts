import { Action } from '@ngrx/store';

export enum CDRActionTypes {
  UPDATE_FAILURE = '[CDR] Failure',
  GET_CDR = '[CDR] Get',
  STORE_GET_CDR = '[CDR][Store] Get',
  StoreGotCdrError = 'StoreGotCdrError',
}

export class Failure implements Action {
  readonly type = CDRActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}

export class GetCDR implements Action {
  readonly type = CDRActionTypes.GET_CDR;
  constructor(public payload: any) {}
}

export class StoreGetCDR implements Action {
  readonly type = CDRActionTypes.STORE_GET_CDR;
  constructor(public payload: any) {}
}

export class StoreGotCdrError implements Action {
  readonly type = CDRActionTypes.StoreGotCdrError;
  constructor(public payload: any) {}
}

export type All =
  | Failure
  | GetCDR
  | StoreGetCDR
  | StoreGotCdrError
;
