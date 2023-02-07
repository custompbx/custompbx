
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetMemcache = 'GetMemcache',
  StoreGetMemcache = 'StoreGetMemcache',
  UpdateMemcacheParameter = 'UpdateMemcacheParameter',
  StoreUpdateMemcacheParameter = 'StoreUpdateMemcacheParameter',
  SwitchMemcacheParameter = 'SwitchMemcacheParameter',
  StoreSwitchMemcacheParameter = 'StoreSwitchMemcacheParameter',
  AddMemcacheParameter = 'AddMemcacheParameter',
  StoreAddMemcacheParameter = 'StoreAddMemcacheParameter',
  DelMemcacheParameter = 'DelMemcacheParameter',
  StoreDelMemcacheParameter = 'StoreDelMemcacheParameter',
  StoreNewMemcacheParameter = 'StoreNewMemcacheParameter',
  StoreDropNewMemcacheParameter = 'StoreDropNewMemcacheParameter',
  StoreGotMemcacheError = 'StoreGotMemcacheError',
}

export class GetMemcache implements Action {
  readonly type = ConfigActionTypes.GetMemcache;
  constructor(public payload: any) {}
}

export class StoreGetMemcache implements Action {
  readonly type = ConfigActionTypes.StoreGetMemcache;
  constructor(public payload: any) {}
}

export class UpdateMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.UpdateMemcacheParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateMemcacheParameter;
  constructor(public payload: any) {}
}

export class SwitchMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.SwitchMemcacheParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchMemcacheParameter;
  constructor(public payload: any) {}
}

export class AddMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.AddMemcacheParameter;
  constructor(public payload: any) {}
}

export class StoreAddMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddMemcacheParameter;
  constructor(public payload: any) {}
}

export class DelMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.DelMemcacheParameter;
  constructor(public payload: any) {}
}

export class StoreDelMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelMemcacheParameter;
  constructor(public payload: any) {}
}

export class StoreGotMemcacheError implements Action {
  readonly type = ConfigActionTypes.StoreGotMemcacheError;
  constructor(public payload: any) {}
}

export class StoreDropNewMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewMemcacheParameter;
  constructor(public payload: any) {}
}

export class StoreNewMemcacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewMemcacheParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetMemcache
  | StoreGetMemcache
  | UpdateMemcacheParameter
  | StoreUpdateMemcacheParameter
  | SwitchMemcacheParameter
  | StoreSwitchMemcacheParameter
  | AddMemcacheParameter
  | StoreAddMemcacheParameter
  | DelMemcacheParameter
  | StoreDelMemcacheParameter
  | StoreGotMemcacheError
  | StoreDropNewMemcacheParameter
  | StoreNewMemcacheParameter
;

