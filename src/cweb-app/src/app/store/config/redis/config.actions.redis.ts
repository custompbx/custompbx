
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetRedis = 'GetRedis',
  StoreGetRedis = 'StoreGetRedis',
  UpdateRedisParameter = 'UpdateRedisParameter',
  StoreUpdateRedisParameter = 'StoreUpdateRedisParameter',
  SwitchRedisParameter = 'SwitchRedisParameter',
  StoreSwitchRedisParameter = 'StoreSwitchRedisParameter',
  AddRedisParameter = 'AddRedisParameter',
  StoreAddRedisParameter = 'StoreAddRedisParameter',
  DelRedisParameter = 'DelRedisParameter',
  StoreDelRedisParameter = 'StoreDelRedisParameter',
  StoreNewRedisParameter = 'StoreNewRedisParameter',
  StoreDropNewRedisParameter = 'StoreDropNewRedisParameter',
  StoreGotRedisError = 'StoreGotRedisError',
}

export class GetRedis implements Action {
  readonly type = ConfigActionTypes.GetRedis;
  constructor(public payload: any) {}
}

export class StoreGetRedis implements Action {
  readonly type = ConfigActionTypes.StoreGetRedis;
  constructor(public payload: any) {}
}

export class UpdateRedisParameter implements Action {
  readonly type = ConfigActionTypes.UpdateRedisParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateRedisParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateRedisParameter;
  constructor(public payload: any) {}
}

export class SwitchRedisParameter implements Action {
  readonly type = ConfigActionTypes.SwitchRedisParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchRedisParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchRedisParameter;
  constructor(public payload: any) {}
}

export class AddRedisParameter implements Action {
  readonly type = ConfigActionTypes.AddRedisParameter;
  constructor(public payload: any) {}
}

export class StoreAddRedisParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddRedisParameter;
  constructor(public payload: any) {}
}

export class DelRedisParameter implements Action {
  readonly type = ConfigActionTypes.DelRedisParameter;
  constructor(public payload: any) {}
}

export class StoreDelRedisParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelRedisParameter;
  constructor(public payload: any) {}
}

export class StoreGotRedisError implements Action {
  readonly type = ConfigActionTypes.StoreGotRedisError;
  constructor(public payload: any) {}
}

export class StoreDropNewRedisParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewRedisParameter;
  constructor(public payload: any) {}
}

export class StoreNewRedisParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewRedisParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetRedis
  | StoreGetRedis
  | UpdateRedisParameter
  | StoreUpdateRedisParameter
  | SwitchRedisParameter
  | StoreSwitchRedisParameter
  | AddRedisParameter
  | StoreAddRedisParameter
  | DelRedisParameter
  | StoreDelRedisParameter
  | StoreGotRedisError
  | StoreDropNewRedisParameter
  | StoreNewRedisParameter
;

