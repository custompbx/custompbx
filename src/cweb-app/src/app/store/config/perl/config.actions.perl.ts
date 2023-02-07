
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetPerl = 'GetPerl',
  StoreGetPerl = 'StoreGetPerl',
  UpdatePerlParameter = 'UpdatePerlParameter',
  StoreUpdatePerlParameter = 'StoreUpdatePerlParameter',
  SwitchPerlParameter = 'SwitchPerlParameter',
  StoreSwitchPerlParameter = 'StoreSwitchPerlParameter',
  AddPerlParameter = 'AddPerlParameter',
  StoreAddPerlParameter = 'StoreAddPerlParameter',
  DelPerlParameter = 'DelPerlParameter',
  StoreDelPerlParameter = 'StoreDelPerlParameter',
  StoreNewPerlParameter = 'StoreNewPerlParameter',
  StoreDropNewPerlParameter = 'StoreDropNewPerlParameter',
  StoreGotPerlError = 'StoreGotPerlError',
}

export class GetPerl implements Action {
  readonly type = ConfigActionTypes.GetPerl;
  constructor(public payload: any) {}
}

export class StoreGetPerl implements Action {
  readonly type = ConfigActionTypes.StoreGetPerl;
  constructor(public payload: any) {}
}

export class UpdatePerlParameter implements Action {
  readonly type = ConfigActionTypes.UpdatePerlParameter;
  constructor(public payload: any) {}
}

export class StoreUpdatePerlParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePerlParameter;
  constructor(public payload: any) {}
}

export class SwitchPerlParameter implements Action {
  readonly type = ConfigActionTypes.SwitchPerlParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchPerlParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPerlParameter;
  constructor(public payload: any) {}
}

export class AddPerlParameter implements Action {
  readonly type = ConfigActionTypes.AddPerlParameter;
  constructor(public payload: any) {}
}

export class StoreAddPerlParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddPerlParameter;
  constructor(public payload: any) {}
}

export class DelPerlParameter implements Action {
  readonly type = ConfigActionTypes.DelPerlParameter;
  constructor(public payload: any) {}
}

export class StoreDelPerlParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelPerlParameter;
  constructor(public payload: any) {}
}

export class StoreGotPerlError implements Action {
  readonly type = ConfigActionTypes.StoreGotPerlError;
  constructor(public payload: any) {}
}

export class StoreDropNewPerlParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPerlParameter;
  constructor(public payload: any) {}
}

export class StoreNewPerlParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewPerlParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetPerl
  | StoreGetPerl
  | UpdatePerlParameter
  | StoreUpdatePerlParameter
  | SwitchPerlParameter
  | StoreSwitchPerlParameter
  | AddPerlParameter
  | StoreAddPerlParameter
  | DelPerlParameter
  | StoreDelPerlParameter
  | StoreGotPerlError
  | StoreDropNewPerlParameter
  | StoreNewPerlParameter
;

