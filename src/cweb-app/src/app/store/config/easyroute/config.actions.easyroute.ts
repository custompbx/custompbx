
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetEasyroute = 'GetEasyroute',
  StoreGetEasyroute = 'StoreGetEasyroute',
  UpdateEasyrouteParameter = 'UpdateEasyrouteParameter',
  StoreUpdateEasyrouteParameter = 'StoreUpdateEasyrouteParameter',
  SwitchEasyrouteParameter = 'SwitchEasyrouteParameter',
  StoreSwitchEasyrouteParameter = 'StoreSwitchEasyrouteParameter',
  AddEasyrouteParameter = 'AddEasyrouteParameter',
  StoreAddEasyrouteParameter = 'StoreAddEasyrouteParameter',
  DelEasyrouteParameter = 'DelEasyrouteParameter',
  StoreDelEasyrouteParameter = 'StoreDelEasyrouteParameter',
  StoreNewEasyrouteParameter = 'StoreNewEasyrouteParameter',
  StoreDropNewEasyrouteParameter = 'StoreDropNewEasyrouteParameter',
  StoreGotEasyrouteError = 'StoreGotEasyrouteError',
}

export class GetEasyroute implements Action {
  readonly type = ConfigActionTypes.GetEasyroute;
  constructor(public payload: any) {}
}

export class StoreGetEasyroute implements Action {
  readonly type = ConfigActionTypes.StoreGetEasyroute;
  constructor(public payload: any) {}
}

export class UpdateEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.UpdateEasyrouteParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateEasyrouteParameter;
  constructor(public payload: any) {}
}

export class SwitchEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.SwitchEasyrouteParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchEasyrouteParameter;
  constructor(public payload: any) {}
}

export class AddEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.AddEasyrouteParameter;
  constructor(public payload: any) {}
}

export class StoreAddEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddEasyrouteParameter;
  constructor(public payload: any) {}
}

export class DelEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.DelEasyrouteParameter;
  constructor(public payload: any) {}
}

export class StoreDelEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelEasyrouteParameter;
  constructor(public payload: any) {}
}

export class StoreGotEasyrouteError implements Action {
  readonly type = ConfigActionTypes.StoreGotEasyrouteError;
  constructor(public payload: any) {}
}

export class StoreDropNewEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewEasyrouteParameter;
  constructor(public payload: any) {}
}

export class StoreNewEasyrouteParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewEasyrouteParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetEasyroute
  | StoreGetEasyroute
  | UpdateEasyrouteParameter
  | StoreUpdateEasyrouteParameter
  | SwitchEasyrouteParameter
  | StoreSwitchEasyrouteParameter
  | AddEasyrouteParameter
  | StoreAddEasyrouteParameter
  | DelEasyrouteParameter
  | StoreDelEasyrouteParameter
  | StoreGotEasyrouteError
  | StoreDropNewEasyrouteParameter
  | StoreNewEasyrouteParameter
;

