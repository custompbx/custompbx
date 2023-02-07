import { Action } from '@ngrx/store';

export enum HEPActionTypes {
  GetHEP = 'GetHEP',
  StoreGetHEP = 'StoreGetHEP',
  StoreGotHEPError = 'StoreGotHEPError',
  GetHEPDetails = 'GetHEPDetails',
  StoreGetHEPDetails = 'StoreGetHEPDetails',
}

export class GetHEP implements Action {
  readonly type = HEPActionTypes.GetHEP;
  constructor(public payload: any) {}
}

export class StoreGetHEP implements Action {
  readonly type = HEPActionTypes.StoreGetHEP;
  constructor(public payload: any) {}
}

export class StoreGotHEPError implements Action {
  readonly type = HEPActionTypes.StoreGotHEPError;
  constructor(public payload: any) {}
}

export class GetHEPDetails implements Action {
  readonly type = HEPActionTypes.GetHEPDetails;
  constructor(public payload: any) {}
}

export class StoreGetHEPDetails implements Action {
  readonly type = HEPActionTypes.StoreGetHEPDetails;
  constructor(public payload: any) {}
}

export type All =
  | GetHEP
  | StoreGetHEP
  | StoreGotHEPError
  | GetHEPDetails
  | StoreGetHEPDetails
;
