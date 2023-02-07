
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetShout = 'GetShout',
  StoreGetShout = 'StoreGetShout',
  UpdateShoutParameter = 'UpdateShoutParameter',
  StoreUpdateShoutParameter = 'StoreUpdateShoutParameter',
  SwitchShoutParameter = 'SwitchShoutParameter',
  StoreSwitchShoutParameter = 'StoreSwitchShoutParameter',
  AddShoutParameter = 'AddShoutParameter',
  StoreAddShoutParameter = 'StoreAddShoutParameter',
  DelShoutParameter = 'DelShoutParameter',
  StoreDelShoutParameter = 'StoreDelShoutParameter',
  StoreNewShoutParameter = 'StoreNewShoutParameter',
  StoreDropNewShoutParameter = 'StoreDropNewShoutParameter',
  StoreGotShoutError = 'StoreGotShoutError',
}

export class GetShout implements Action {
  readonly type = ConfigActionTypes.GetShout;
  constructor(public payload: any) {}
}

export class StoreGetShout implements Action {
  readonly type = ConfigActionTypes.StoreGetShout;
  constructor(public payload: any) {}
}

export class UpdateShoutParameter implements Action {
  readonly type = ConfigActionTypes.UpdateShoutParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateShoutParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateShoutParameter;
  constructor(public payload: any) {}
}

export class SwitchShoutParameter implements Action {
  readonly type = ConfigActionTypes.SwitchShoutParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchShoutParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchShoutParameter;
  constructor(public payload: any) {}
}

export class AddShoutParameter implements Action {
  readonly type = ConfigActionTypes.AddShoutParameter;
  constructor(public payload: any) {}
}

export class StoreAddShoutParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddShoutParameter;
  constructor(public payload: any) {}
}

export class DelShoutParameter implements Action {
  readonly type = ConfigActionTypes.DelShoutParameter;
  constructor(public payload: any) {}
}

export class StoreDelShoutParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelShoutParameter;
  constructor(public payload: any) {}
}

export class StoreGotShoutError implements Action {
  readonly type = ConfigActionTypes.StoreGotShoutError;
  constructor(public payload: any) {}
}

export class StoreDropNewShoutParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewShoutParameter;
  constructor(public payload: any) {}
}

export class StoreNewShoutParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewShoutParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetShout
  | StoreGetShout
  | UpdateShoutParameter
  | StoreUpdateShoutParameter
  | SwitchShoutParameter
  | StoreSwitchShoutParameter
  | AddShoutParameter
  | StoreAddShoutParameter
  | DelShoutParameter
  | StoreDelShoutParameter
  | StoreGotShoutError
  | StoreDropNewShoutParameter
  | StoreNewShoutParameter
;

