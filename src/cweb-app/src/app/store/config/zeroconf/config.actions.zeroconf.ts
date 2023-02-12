
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetZeroconf = 'GetZeroconf',
  StoreGetZeroconf = 'StoreGetZeroconf',
  UpdateZeroconfParameter = 'UpdateZeroconfParameter',
  StoreUpdateZeroconfParameter = 'StoreUpdateZeroconfParameter',
  SwitchZeroconfParameter = 'SwitchZeroconfParameter',
  StoreSwitchZeroconfParameter = 'StoreSwitchZeroconfParameter',
  AddZeroconfParameter = 'AddZeroconfParameter',
  StoreAddZeroconfParameter = 'StoreAddZeroconfParameter',
  DelZeroconfParameter = 'DelZeroconfParameter',
  StoreDelZeroconfParameter = 'StoreDelZeroconfParameter',
  StoreNewZeroconfParameter = 'StoreNewZeroconfParameter',
  StoreDropNewZeroconfParameter = 'StoreDropNewZeroconfParameter',
  StoreGotZeroconfError = 'StoreGotZeroconfError',
}

export class GetZeroconf implements Action {
  readonly type = ConfigActionTypes.GetZeroconf;
  constructor(public payload: any) {}
}

export class StoreGetZeroconf implements Action {
  readonly type = ConfigActionTypes.StoreGetZeroconf;
  constructor(public payload: any) {}
}

export class UpdateZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.UpdateZeroconfParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateZeroconfParameter;
  constructor(public payload: any) {}
}

export class SwitchZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.SwitchZeroconfParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchZeroconfParameter;
  constructor(public payload: any) {}
}

export class AddZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.AddZeroconfParameter;
  constructor(public payload: any) {}
}

export class StoreAddZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddZeroconfParameter;
  constructor(public payload: any) {}
}

export class DelZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.DelZeroconfParameter;
  constructor(public payload: any) {}
}

export class StoreDelZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelZeroconfParameter;
  constructor(public payload: any) {}
}

export class StoreGotZeroconfError implements Action {
  readonly type = ConfigActionTypes.StoreGotZeroconfError;
  constructor(public payload: any) {}
}

export class StoreDropNewZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewZeroconfParameter;
  constructor(public payload: any) {}
}

export class StoreNewZeroconfParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewZeroconfParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetZeroconf
  | StoreGetZeroconf
  | UpdateZeroconfParameter
  | StoreUpdateZeroconfParameter
  | SwitchZeroconfParameter
  | StoreSwitchZeroconfParameter
  | AddZeroconfParameter
  | StoreAddZeroconfParameter
  | DelZeroconfParameter
  | StoreDelZeroconfParameter
  | StoreGotZeroconfError
  | StoreDropNewZeroconfParameter
  | StoreNewZeroconfParameter
;

