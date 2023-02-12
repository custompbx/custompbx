
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetAvmd = 'GetAvmd',
  StoreGetAvmd = 'StoreGetAvmd',
  UpdateAvmdParameter = 'UpdateAvmdParameter',
  StoreUpdateAvmdParameter = 'StoreUpdateAvmdParameter',
  SwitchAvmdParameter = 'SwitchAvmdParameter',
  StoreSwitchAvmdParameter = 'StoreSwitchAvmdParameter',
  AddAvmdParameter = 'AddAvmdParameter',
  StoreAddAvmdParameter = 'StoreAddAvmdParameter',
  DelAvmdParameter = 'DelAvmdParameter',
  StoreDelAvmdParameter = 'StoreDelAvmdParameter',
  StoreNewAvmdParameter = 'StoreNewAvmdParameter',
  StoreDropNewAvmdParameter = 'StoreDropNewAvmdParameter',
  StoreGotAvmdError = 'StoreGotAvmdError',
}

export class GetAvmd implements Action {
  readonly type = ConfigActionTypes.GetAvmd;
  constructor(public payload: any) {}
}

export class StoreGetAvmd implements Action {
  readonly type = ConfigActionTypes.StoreGetAvmd;
  constructor(public payload: any) {}
}

export class UpdateAvmdParameter implements Action {
  readonly type = ConfigActionTypes.UpdateAvmdParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateAvmdParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateAvmdParameter;
  constructor(public payload: any) {}
}

export class SwitchAvmdParameter implements Action {
  readonly type = ConfigActionTypes.SwitchAvmdParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchAvmdParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchAvmdParameter;
  constructor(public payload: any) {}
}

export class AddAvmdParameter implements Action {
  readonly type = ConfigActionTypes.AddAvmdParameter;
  constructor(public payload: any) {}
}

export class StoreAddAvmdParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddAvmdParameter;
  constructor(public payload: any) {}
}

export class DelAvmdParameter implements Action {
  readonly type = ConfigActionTypes.DelAvmdParameter;
  constructor(public payload: any) {}
}

export class StoreDelAvmdParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelAvmdParameter;
  constructor(public payload: any) {}
}

export class StoreGotAvmdError implements Action {
  readonly type = ConfigActionTypes.StoreGotAvmdError;
  constructor(public payload: any) {}
}

export class StoreDropNewAvmdParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewAvmdParameter;
  constructor(public payload: any) {}
}

export class StoreNewAvmdParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewAvmdParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetAvmd
  | StoreGetAvmd
  | UpdateAvmdParameter
  | StoreUpdateAvmdParameter
  | SwitchAvmdParameter
  | StoreSwitchAvmdParameter
  | AddAvmdParameter
  | StoreAddAvmdParameter
  | DelAvmdParameter
  | StoreDelAvmdParameter
  | StoreGotAvmdError
  | StoreDropNewAvmdParameter
  | StoreNewAvmdParameter
;

