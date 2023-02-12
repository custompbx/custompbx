
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetPocketsphinx = 'GetPocketsphinx',
  StoreGetPocketsphinx = 'StoreGetPocketsphinx',
  UpdatePocketsphinxParameter = 'UpdatePocketsphinxParameter',
  StoreUpdatePocketsphinxParameter = 'StoreUpdatePocketsphinxParameter',
  SwitchPocketsphinxParameter = 'SwitchPocketsphinxParameter',
  StoreSwitchPocketsphinxParameter = 'StoreSwitchPocketsphinxParameter',
  AddPocketsphinxParameter = 'AddPocketsphinxParameter',
  StoreAddPocketsphinxParameter = 'StoreAddPocketsphinxParameter',
  DelPocketsphinxParameter = 'DelPocketsphinxParameter',
  StoreDelPocketsphinxParameter = 'StoreDelPocketsphinxParameter',
  StoreNewPocketsphinxParameter = 'StoreNewPocketsphinxParameter',
  StoreDropNewPocketsphinxParameter = 'StoreDropNewPocketsphinxParameter',
  StoreGotPocketsphinxError = 'StoreGotPocketsphinxError',
}

export class GetPocketsphinx implements Action {
  readonly type = ConfigActionTypes.GetPocketsphinx;
  constructor(public payload: any) {}
}

export class StoreGetPocketsphinx implements Action {
  readonly type = ConfigActionTypes.StoreGetPocketsphinx;
  constructor(public payload: any) {}
}

export class UpdatePocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.UpdatePocketsphinxParameter;
  constructor(public payload: any) {}
}

export class StoreUpdatePocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePocketsphinxParameter;
  constructor(public payload: any) {}
}

export class SwitchPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.SwitchPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class AddPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.AddPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class StoreAddPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class DelPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.DelPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class StoreDelPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class StoreGotPocketsphinxError implements Action {
  readonly type = ConfigActionTypes.StoreGotPocketsphinxError;
  constructor(public payload: any) {}
}

export class StoreDropNewPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPocketsphinxParameter;
  constructor(public payload: any) {}
}

export class StoreNewPocketsphinxParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewPocketsphinxParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetPocketsphinx
  | StoreGetPocketsphinx
  | UpdatePocketsphinxParameter
  | StoreUpdatePocketsphinxParameter
  | SwitchPocketsphinxParameter
  | StoreSwitchPocketsphinxParameter
  | AddPocketsphinxParameter
  | StoreAddPocketsphinxParameter
  | DelPocketsphinxParameter
  | StoreDelPocketsphinxParameter
  | StoreGotPocketsphinxError
  | StoreDropNewPocketsphinxParameter
  | StoreNewPocketsphinxParameter
;

