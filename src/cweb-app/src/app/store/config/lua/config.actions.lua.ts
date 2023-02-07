
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetLua = 'GetLua',
  StoreGetLua = 'StoreGetLua',
  UpdateLuaParameter = 'UpdateLuaParameter',
  StoreUpdateLuaParameter = 'StoreUpdateLuaParameter',
  SwitchLuaParameter = 'SwitchLuaParameter',
  StoreSwitchLuaParameter = 'StoreSwitchLuaParameter',
  AddLuaParameter = 'AddLuaParameter',
  StoreAddLuaParameter = 'StoreAddLuaParameter',
  DelLuaParameter = 'DelLuaParameter',
  StoreDelLuaParameter = 'StoreDelLuaParameter',
  StoreNewLuaParameter = 'StoreNewLuaParameter',
  StoreDropNewLuaParameter = 'StoreDropNewLuaParameter',
  StoreGotLuaError = 'StoreGotLuaError',
}

export class GetLua implements Action {
  readonly type = ConfigActionTypes.GetLua;
  constructor(public payload: any) {}
}

export class StoreGetLua implements Action {
  readonly type = ConfigActionTypes.StoreGetLua;
  constructor(public payload: any) {}
}

export class UpdateLuaParameter implements Action {
  readonly type = ConfigActionTypes.UpdateLuaParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateLuaParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateLuaParameter;
  constructor(public payload: any) {}
}

export class SwitchLuaParameter implements Action {
  readonly type = ConfigActionTypes.SwitchLuaParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchLuaParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchLuaParameter;
  constructor(public payload: any) {}
}

export class AddLuaParameter implements Action {
  readonly type = ConfigActionTypes.AddLuaParameter;
  constructor(public payload: any) {}
}

export class StoreAddLuaParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddLuaParameter;
  constructor(public payload: any) {}
}

export class DelLuaParameter implements Action {
  readonly type = ConfigActionTypes.DelLuaParameter;
  constructor(public payload: any) {}
}

export class StoreDelLuaParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelLuaParameter;
  constructor(public payload: any) {}
}

export class StoreGotLuaError implements Action {
  readonly type = ConfigActionTypes.StoreGotLuaError;
  constructor(public payload: any) {}
}

export class StoreDropNewLuaParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewLuaParameter;
  constructor(public payload: any) {}
}

export class StoreNewLuaParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewLuaParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetLua
  | StoreGetLua
  | UpdateLuaParameter
  | StoreUpdateLuaParameter
  | SwitchLuaParameter
  | StoreSwitchLuaParameter
  | AddLuaParameter
  | StoreAddLuaParameter
  | DelLuaParameter
  | StoreDelLuaParameter
  | StoreGotLuaError
  | StoreDropNewLuaParameter
  | StoreNewLuaParameter
;

