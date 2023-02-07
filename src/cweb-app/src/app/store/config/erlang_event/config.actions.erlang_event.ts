
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetErlangEvent = 'GetErlangEvent',
  StoreGetErlangEvent = 'StoreGetErlangEvent',
  UpdateErlangEventParameter = 'UpdateErlangEventParameter',
  StoreUpdateErlangEventParameter = 'StoreUpdateErlangEventParameter',
  SwitchErlangEventParameter = 'SwitchErlangEventParameter',
  StoreSwitchErlangEventParameter = 'StoreSwitchErlangEventParameter',
  AddErlangEventParameter = 'AddErlangEventParameter',
  StoreAddErlangEventParameter = 'StoreAddErlangEventParameter',
  DelErlangEventParameter = 'DelErlangEventParameter',
  StoreDelErlangEventParameter = 'StoreDelErlangEventParameter',
  StoreNewErlangEventParameter = 'StoreNewErlangEventParameter',
  StoreDropNewErlangEventParameter = 'StoreDropNewErlangEventParameter',
  StoreGotErlangEventError = 'StoreGotErlangEventError',
}

export class GetErlangEvent implements Action {
  readonly type = ConfigActionTypes.GetErlangEvent;
  constructor(public payload: any) {}
}

export class StoreGetErlangEvent implements Action {
  readonly type = ConfigActionTypes.StoreGetErlangEvent;
  constructor(public payload: any) {}
}

export class UpdateErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.UpdateErlangEventParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateErlangEventParameter;
  constructor(public payload: any) {}
}

export class SwitchErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.SwitchErlangEventParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchErlangEventParameter;
  constructor(public payload: any) {}
}

export class AddErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.AddErlangEventParameter;
  constructor(public payload: any) {}
}

export class StoreAddErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddErlangEventParameter;
  constructor(public payload: any) {}
}

export class DelErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.DelErlangEventParameter;
  constructor(public payload: any) {}
}

export class StoreDelErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelErlangEventParameter;
  constructor(public payload: any) {}
}

export class StoreGotErlangEventError implements Action {
  readonly type = ConfigActionTypes.StoreGotErlangEventError;
  constructor(public payload: any) {}
}

export class StoreDropNewErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewErlangEventParameter;
  constructor(public payload: any) {}
}

export class StoreNewErlangEventParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewErlangEventParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetErlangEvent
  | StoreGetErlangEvent
  | UpdateErlangEventParameter
  | StoreUpdateErlangEventParameter
  | SwitchErlangEventParameter
  | StoreSwitchErlangEventParameter
  | AddErlangEventParameter
  | StoreAddErlangEventParameter
  | DelErlangEventParameter
  | StoreDelErlangEventParameter
  | StoreGotErlangEventError
  | StoreDropNewErlangEventParameter
  | StoreNewErlangEventParameter
;

