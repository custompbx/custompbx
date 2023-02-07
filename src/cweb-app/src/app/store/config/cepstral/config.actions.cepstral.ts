
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetCepstral = 'GetCepstral',
  StoreGetCepstral = 'StoreGetCepstral',
  UpdateCepstralParameter = 'UpdateCepstralParameter',
  StoreUpdateCepstralParameter = 'StoreUpdateCepstralParameter',
  SwitchCepstralParameter = 'SwitchCepstralParameter',
  StoreSwitchCepstralParameter = 'StoreSwitchCepstralParameter',
  AddCepstralParameter = 'AddCepstralParameter',
  StoreAddCepstralParameter = 'StoreAddCepstralParameter',
  DelCepstralParameter = 'DelCepstralParameter',
  StoreDelCepstralParameter = 'StoreDelCepstralParameter',
  StoreNewCepstralParameter = 'StoreNewCepstralParameter',
  StoreDropNewCepstralParameter = 'StoreDropNewCepstralParameter',
  StoreGotCepstralError = 'StoreGotCepstralError',
}

export class GetCepstral implements Action {
  readonly type = ConfigActionTypes.GetCepstral;
  constructor(public payload: any) {}
}

export class StoreGetCepstral implements Action {
  readonly type = ConfigActionTypes.StoreGetCepstral;
  constructor(public payload: any) {}
}

export class UpdateCepstralParameter implements Action {
  readonly type = ConfigActionTypes.UpdateCepstralParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateCepstralParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCepstralParameter;
  constructor(public payload: any) {}
}

export class SwitchCepstralParameter implements Action {
  readonly type = ConfigActionTypes.SwitchCepstralParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchCepstralParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchCepstralParameter;
  constructor(public payload: any) {}
}

export class AddCepstralParameter implements Action {
  readonly type = ConfigActionTypes.AddCepstralParameter;
  constructor(public payload: any) {}
}

export class StoreAddCepstralParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddCepstralParameter;
  constructor(public payload: any) {}
}

export class DelCepstralParameter implements Action {
  readonly type = ConfigActionTypes.DelCepstralParameter;
  constructor(public payload: any) {}
}

export class StoreDelCepstralParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelCepstralParameter;
  constructor(public payload: any) {}
}

export class StoreGotCepstralError implements Action {
  readonly type = ConfigActionTypes.StoreGotCepstralError;
  constructor(public payload: any) {}
}

export class StoreDropNewCepstralParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewCepstralParameter;
  constructor(public payload: any) {}
}

export class StoreNewCepstralParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewCepstralParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetCepstral
  | StoreGetCepstral
  | UpdateCepstralParameter
  | StoreUpdateCepstralParameter
  | SwitchCepstralParameter
  | StoreSwitchCepstralParameter
  | AddCepstralParameter
  | StoreAddCepstralParameter
  | DelCepstralParameter
  | StoreDelCepstralParameter
  | StoreGotCepstralError
  | StoreDropNewCepstralParameter
  | StoreNewCepstralParameter
;

