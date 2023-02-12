
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetHttpCache = 'GetHttpCache',
  StoreGetHttpCache = 'StoreGetHttpCache',
  UpdateHttpCacheParameter = 'UpdateHttpCacheParameter',
  StoreUpdateHttpCacheParameter = 'StoreUpdateHttpCacheParameter',
  SwitchHttpCacheParameter = 'SwitchHttpCacheParameter',
  StoreSwitchHttpCacheParameter = 'StoreSwitchHttpCacheParameter',
  AddHttpCacheParameter = 'AddHttpCacheParameter',
  StoreAddHttpCacheParameter = 'StoreAddHttpCacheParameter',
  DelHttpCacheParameter = 'DelHttpCacheParameter',
  StoreDelHttpCacheParameter = 'StoreDelHttpCacheParameter',
  StoreNewHttpCacheParameter = 'StoreNewHttpCacheParameter',
  StoreDropNewHttpCacheParameter = 'StoreDropNewHttpCacheParameter',
  StoreGotHttpCacheError = 'StoreGotHttpCacheError',
}

export class GetHttpCache implements Action {
  readonly type = ConfigActionTypes.GetHttpCache;
  constructor(public payload: any) {}
}

export class StoreGetHttpCache implements Action {
  readonly type = ConfigActionTypes.StoreGetHttpCache;
  constructor(public payload: any) {}
}

export class UpdateHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.UpdateHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateHttpCacheParameter;
  constructor(public payload: any) {}
}

export class SwitchHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.SwitchHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchHttpCacheParameter;
  constructor(public payload: any) {}
}

export class AddHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.AddHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreAddHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddHttpCacheParameter;
  constructor(public payload: any) {}
}

export class DelHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.DelHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreDelHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreGotHttpCacheError implements Action {
  readonly type = ConfigActionTypes.StoreGotHttpCacheError;
  constructor(public payload: any) {}
}

export class StoreDropNewHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreNewHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewHttpCacheParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetHttpCache
  | StoreGetHttpCache
  | UpdateHttpCacheParameter
  | StoreUpdateHttpCacheParameter
  | SwitchHttpCacheParameter
  | StoreSwitchHttpCacheParameter
  | AddHttpCacheParameter
  | StoreAddHttpCacheParameter
  | DelHttpCacheParameter
  | StoreDelHttpCacheParameter
  | StoreGotHttpCacheError
  | StoreDropNewHttpCacheParameter
  | StoreNewHttpCacheParameter
;

