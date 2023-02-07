
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetAmrwb = 'GetAmrwb',
  StoreGetAmrwb = 'StoreGetAmrwb',
  UpdateAmrwbParameter = 'UpdateAmrwbParameter',
  StoreUpdateAmrwbParameter = 'StoreUpdateAmrwbParameter',
  SwitchAmrwbParameter = 'SwitchAmrwbParameter',
  StoreSwitchAmrwbParameter = 'StoreSwitchAmrwbParameter',
  AddAmrwbParameter = 'AddAmrwbParameter',
  StoreAddAmrwbParameter = 'StoreAddAmrwbParameter',
  DelAmrwbParameter = 'DelAmrwbParameter',
  StoreDelAmrwbParameter = 'StoreDelAmrwbParameter',
  StoreNewAmrwbParameter = 'StoreNewAmrwbParameter',
  StoreDropNewAmrwbParameter = 'StoreDropNewAmrwbParameter',
  StoreGotAmrwbError = 'StoreGotAmrwbError',
}

export class GetAmrwb implements Action {
  readonly type = ConfigActionTypes.GetAmrwb;
  constructor(public payload: any) {}
}

export class StoreGetAmrwb implements Action {
  readonly type = ConfigActionTypes.StoreGetAmrwb;
  constructor(public payload: any) {}
}

export class UpdateAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.UpdateAmrwbParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateAmrwbParameter;
  constructor(public payload: any) {}
}

export class SwitchAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.SwitchAmrwbParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchAmrwbParameter;
  constructor(public payload: any) {}
}

export class AddAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.AddAmrwbParameter;
  constructor(public payload: any) {}
}

export class StoreAddAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddAmrwbParameter;
  constructor(public payload: any) {}
}

export class DelAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.DelAmrwbParameter;
  constructor(public payload: any) {}
}

export class StoreDelAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelAmrwbParameter;
  constructor(public payload: any) {}
}

export class StoreGotAmrwbError implements Action {
  readonly type = ConfigActionTypes.StoreGotAmrwbError;
  constructor(public payload: any) {}
}

export class StoreDropNewAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewAmrwbParameter;
  constructor(public payload: any) {}
}

export class StoreNewAmrwbParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewAmrwbParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetAmrwb
  | StoreGetAmrwb
  | UpdateAmrwbParameter
  | StoreUpdateAmrwbParameter
  | SwitchAmrwbParameter
  | StoreSwitchAmrwbParameter
  | AddAmrwbParameter
  | StoreAddAmrwbParameter
  | DelAmrwbParameter
  | StoreDelAmrwbParameter
  | StoreGotAmrwbError
  | StoreDropNewAmrwbParameter
  | StoreNewAmrwbParameter
;

