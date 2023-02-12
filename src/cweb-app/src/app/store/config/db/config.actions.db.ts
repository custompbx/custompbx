
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetDb = 'GetDb',
  StoreGetDb = 'StoreGetDb',
  UpdateDbParameter = 'UpdateDbParameter',
  StoreUpdateDbParameter = 'StoreUpdateDbParameter',
  SwitchDbParameter = 'SwitchDbParameter',
  StoreSwitchDbParameter = 'StoreSwitchDbParameter',
  AddDbParameter = 'AddDbParameter',
  StoreAddDbParameter = 'StoreAddDbParameter',
  DelDbParameter = 'DelDbParameter',
  StoreDelDbParameter = 'StoreDelDbParameter',
  StoreNewDbParameter = 'StoreNewDbParameter',
  StoreDropNewDbParameter = 'StoreDropNewDbParameter',
  StoreGotDbError = 'StoreGotDbError',
}

export class GetDb implements Action {
  readonly type = ConfigActionTypes.GetDb;
  constructor(public payload: any) {}
}

export class StoreGetDb implements Action {
  readonly type = ConfigActionTypes.StoreGetDb;
  constructor(public payload: any) {}
}

export class UpdateDbParameter implements Action {
  readonly type = ConfigActionTypes.UpdateDbParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDbParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDbParameter;
  constructor(public payload: any) {}
}

export class SwitchDbParameter implements Action {
  readonly type = ConfigActionTypes.SwitchDbParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDbParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchDbParameter;
  constructor(public payload: any) {}
}

export class AddDbParameter implements Action {
  readonly type = ConfigActionTypes.AddDbParameter;
  constructor(public payload: any) {}
}

export class StoreAddDbParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddDbParameter;
  constructor(public payload: any) {}
}

export class DelDbParameter implements Action {
  readonly type = ConfigActionTypes.DelDbParameter;
  constructor(public payload: any) {}
}

export class StoreDelDbParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelDbParameter;
  constructor(public payload: any) {}
}

export class StoreGotDbError implements Action {
  readonly type = ConfigActionTypes.StoreGotDbError;
  constructor(public payload: any) {}
}

export class StoreDropNewDbParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewDbParameter;
  constructor(public payload: any) {}
}

export class StoreNewDbParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewDbParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetDb
  | StoreGetDb
  | UpdateDbParameter
  | StoreUpdateDbParameter
  | SwitchDbParameter
  | StoreSwitchDbParameter
  | AddDbParameter
  | StoreAddDbParameter
  | DelDbParameter
  | StoreDelDbParameter
  | StoreGotDbError
  | StoreDropNewDbParameter
  | StoreNewDbParameter
;

