
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetSndfile = 'GetSndfile',
  StoreGetSndfile = 'StoreGetSndfile',
  UpdateSndfileParameter = 'UpdateSndfileParameter',
  StoreUpdateSndfileParameter = 'StoreUpdateSndfileParameter',
  SwitchSndfileParameter = 'SwitchSndfileParameter',
  StoreSwitchSndfileParameter = 'StoreSwitchSndfileParameter',
  AddSndfileParameter = 'AddSndfileParameter',
  StoreAddSndfileParameter = 'StoreAddSndfileParameter',
  DelSndfileParameter = 'DelSndfileParameter',
  StoreDelSndfileParameter = 'StoreDelSndfileParameter',
  StoreNewSndfileParameter = 'StoreNewSndfileParameter',
  StoreDropNewSndfileParameter = 'StoreDropNewSndfileParameter',
  StoreGotSndfileError = 'StoreGotSndfileError',
}

export class GetSndfile implements Action {
  readonly type = ConfigActionTypes.GetSndfile;
  constructor(public payload: any) {}
}

export class StoreGetSndfile implements Action {
  readonly type = ConfigActionTypes.StoreGetSndfile;
  constructor(public payload: any) {}
}

export class UpdateSndfileParameter implements Action {
  readonly type = ConfigActionTypes.UpdateSndfileParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateSndfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateSndfileParameter;
  constructor(public payload: any) {}
}

export class SwitchSndfileParameter implements Action {
  readonly type = ConfigActionTypes.SwitchSndfileParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchSndfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchSndfileParameter;
  constructor(public payload: any) {}
}

export class AddSndfileParameter implements Action {
  readonly type = ConfigActionTypes.AddSndfileParameter;
  constructor(public payload: any) {}
}

export class StoreAddSndfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddSndfileParameter;
  constructor(public payload: any) {}
}

export class DelSndfileParameter implements Action {
  readonly type = ConfigActionTypes.DelSndfileParameter;
  constructor(public payload: any) {}
}

export class StoreDelSndfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelSndfileParameter;
  constructor(public payload: any) {}
}

export class StoreGotSndfileError implements Action {
  readonly type = ConfigActionTypes.StoreGotSndfileError;
  constructor(public payload: any) {}
}

export class StoreDropNewSndfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewSndfileParameter;
  constructor(public payload: any) {}
}

export class StoreNewSndfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewSndfileParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetSndfile
  | StoreGetSndfile
  | UpdateSndfileParameter
  | StoreUpdateSndfileParameter
  | SwitchSndfileParameter
  | StoreSwitchSndfileParameter
  | AddSndfileParameter
  | StoreAddSndfileParameter
  | DelSndfileParameter
  | StoreDelSndfileParameter
  | StoreGotSndfileError
  | StoreDropNewSndfileParameter
  | StoreNewSndfileParameter
;

