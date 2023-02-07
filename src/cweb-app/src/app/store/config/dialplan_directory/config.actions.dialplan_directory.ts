
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetDialplanDirectory = 'GetDialplanDirectory',
  StoreGetDialplanDirectory = 'StoreGetDialplanDirectory',
  UpdateDialplanDirectoryParameter = 'UpdateDialplanDirectoryParameter',
  StoreUpdateDialplanDirectoryParameter = 'StoreUpdateDialplanDirectoryParameter',
  SwitchDialplanDirectoryParameter = 'SwitchDialplanDirectoryParameter',
  StoreSwitchDialplanDirectoryParameter = 'StoreSwitchDialplanDirectoryParameter',
  AddDialplanDirectoryParameter = 'AddDialplanDirectoryParameter',
  StoreAddDialplanDirectoryParameter = 'StoreAddDialplanDirectoryParameter',
  DelDialplanDirectoryParameter = 'DelDialplanDirectoryParameter',
  StoreDelDialplanDirectoryParameter = 'StoreDelDialplanDirectoryParameter',
  StoreNewDialplanDirectoryParameter = 'StoreNewDialplanDirectoryParameter',
  StoreDropNewDialplanDirectoryParameter = 'StoreDropNewDialplanDirectoryParameter',
  StoreGotDialplanDirectoryError = 'StoreGotDialplanDirectoryError',
}

export class GetDialplanDirectory implements Action {
  readonly type = ConfigActionTypes.GetDialplanDirectory;
  constructor(public payload: any) {}
}

export class StoreGetDialplanDirectory implements Action {
  readonly type = ConfigActionTypes.StoreGetDialplanDirectory;
  constructor(public payload: any) {}
}

export class UpdateDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.UpdateDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class SwitchDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.SwitchDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class AddDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.AddDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreAddDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class DelDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.DelDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreDelDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreGotDialplanDirectoryError implements Action {
  readonly type = ConfigActionTypes.StoreGotDialplanDirectoryError;
  constructor(public payload: any) {}
}

export class StoreDropNewDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreNewDialplanDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewDialplanDirectoryParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetDialplanDirectory
  | StoreGetDialplanDirectory
  | UpdateDialplanDirectoryParameter
  | StoreUpdateDialplanDirectoryParameter
  | SwitchDialplanDirectoryParameter
  | StoreSwitchDialplanDirectoryParameter
  | AddDialplanDirectoryParameter
  | StoreAddDialplanDirectoryParameter
  | DelDialplanDirectoryParameter
  | StoreDelDialplanDirectoryParameter
  | StoreGotDialplanDirectoryError
  | StoreDropNewDialplanDirectoryParameter
  | StoreNewDialplanDirectoryParameter
;

