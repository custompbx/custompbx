
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetPython = 'GetPython',
  StoreGetPython = 'StoreGetPython',
  UpdatePythonParameter = 'UpdatePythonParameter',
  StoreUpdatePythonParameter = 'StoreUpdatePythonParameter',
  SwitchPythonParameter = 'SwitchPythonParameter',
  StoreSwitchPythonParameter = 'StoreSwitchPythonParameter',
  AddPythonParameter = 'AddPythonParameter',
  StoreAddPythonParameter = 'StoreAddPythonParameter',
  DelPythonParameter = 'DelPythonParameter',
  StoreDelPythonParameter = 'StoreDelPythonParameter',
  StoreNewPythonParameter = 'StoreNewPythonParameter',
  StoreDropNewPythonParameter = 'StoreDropNewPythonParameter',
  StoreGotPythonError = 'StoreGotPythonError',
}

export class GetPython implements Action {
  readonly type = ConfigActionTypes.GetPython;
  constructor(public payload: any) {}
}

export class StoreGetPython implements Action {
  readonly type = ConfigActionTypes.StoreGetPython;
  constructor(public payload: any) {}
}

export class UpdatePythonParameter implements Action {
  readonly type = ConfigActionTypes.UpdatePythonParameter;
  constructor(public payload: any) {}
}

export class StoreUpdatePythonParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePythonParameter;
  constructor(public payload: any) {}
}

export class SwitchPythonParameter implements Action {
  readonly type = ConfigActionTypes.SwitchPythonParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchPythonParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPythonParameter;
  constructor(public payload: any) {}
}

export class AddPythonParameter implements Action {
  readonly type = ConfigActionTypes.AddPythonParameter;
  constructor(public payload: any) {}
}

export class StoreAddPythonParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddPythonParameter;
  constructor(public payload: any) {}
}

export class DelPythonParameter implements Action {
  readonly type = ConfigActionTypes.DelPythonParameter;
  constructor(public payload: any) {}
}

export class StoreDelPythonParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelPythonParameter;
  constructor(public payload: any) {}
}

export class StoreGotPythonError implements Action {
  readonly type = ConfigActionTypes.StoreGotPythonError;
  constructor(public payload: any) {}
}

export class StoreDropNewPythonParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPythonParameter;
  constructor(public payload: any) {}
}

export class StoreNewPythonParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewPythonParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetPython
  | StoreGetPython
  | UpdatePythonParameter
  | StoreUpdatePythonParameter
  | SwitchPythonParameter
  | StoreSwitchPythonParameter
  | AddPythonParameter
  | StoreAddPythonParameter
  | DelPythonParameter
  | StoreDelPythonParameter
  | StoreGotPythonError
  | StoreDropNewPythonParameter
  | StoreNewPythonParameter
;

