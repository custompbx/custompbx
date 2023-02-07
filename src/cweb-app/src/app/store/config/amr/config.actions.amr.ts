
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetAmr = 'GetAmr',
  StoreGetAmr = 'StoreGetAmr',
  UpdateAmrParameter = 'UpdateAmrParameter',
  StoreUpdateAmrParameter = 'StoreUpdateAmrParameter',
  SwitchAmrParameter = 'SwitchAmrParameter',
  StoreSwitchAmrParameter = 'StoreSwitchAmrParameter',
  AddAmrParameter = 'AddAmrParameter',
  StoreAddAmrParameter = 'StoreAddAmrParameter',
  DelAmrParameter = 'DelAmrParameter',
  StoreDelAmrParameter = 'StoreDelAmrParameter',
  StoreNewAmrParameter = 'StoreNewAmrParameter',
  StoreDropNewAmrParameter = 'StoreDropNewAmrParameter',
  StoreGotAmrError = 'StoreGotAmrError',
}

export class GetAmr implements Action {
  readonly type = ConfigActionTypes.GetAmr;
  constructor(public payload: any) {}
}

export class StoreGetAmr implements Action {
  readonly type = ConfigActionTypes.StoreGetAmr;
  constructor(public payload: any) {}
}

export class UpdateAmrParameter implements Action {
  readonly type = ConfigActionTypes.UpdateAmrParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateAmrParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateAmrParameter;
  constructor(public payload: any) {}
}

export class SwitchAmrParameter implements Action {
  readonly type = ConfigActionTypes.SwitchAmrParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchAmrParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchAmrParameter;
  constructor(public payload: any) {}
}

export class AddAmrParameter implements Action {
  readonly type = ConfigActionTypes.AddAmrParameter;
  constructor(public payload: any) {}
}

export class StoreAddAmrParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddAmrParameter;
  constructor(public payload: any) {}
}

export class DelAmrParameter implements Action {
  readonly type = ConfigActionTypes.DelAmrParameter;
  constructor(public payload: any) {}
}

export class StoreDelAmrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelAmrParameter;
  constructor(public payload: any) {}
}

export class StoreGotAmrError implements Action {
  readonly type = ConfigActionTypes.StoreGotAmrError;
  constructor(public payload: any) {}
}

export class StoreDropNewAmrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewAmrParameter;
  constructor(public payload: any) {}
}

export class StoreNewAmrParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewAmrParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetAmr
  | StoreGetAmr
  | UpdateAmrParameter
  | StoreUpdateAmrParameter
  | SwitchAmrParameter
  | StoreSwitchAmrParameter
  | AddAmrParameter
  | StoreAddAmrParameter
  | DelAmrParameter
  | StoreDelAmrParameter
  | StoreGotAmrError
  | StoreDropNewAmrParameter
  | StoreNewAmrParameter
;

