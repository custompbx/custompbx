import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetAlsa = 'GetAlsa',
  StoreGetAlsa = 'StoreGetAlsa',
  UpdateAlsaParameter = 'UpdateAlsaParameter',
  StoreUpdateAlsaParameter = 'StoreUpdateAlsaParameter',
  SwitchAlsaParameter = 'SwitchAlsaParameter',
  StoreSwitchAlsaParameter = 'StoreSwitchAlsaParameter',
  AddAlsaParameter = 'AddAlsaParameter',
  StoreAddAlsaParameter = 'StoreAddAlsaParameter',
  DelAlsaParameter = 'DelAlsaParameter',
  StoreDelAlsaParameter = 'StoreDelAlsaParameter',
  StoreNewAlsaParameter = 'StoreNewAlsaParameter',
  StoreDropNewAlsaParameter = 'StoreDropNewAlsaParameter',
  StoreGotAlsaError = 'StoreGotAlsaError',
}

export class GetAlsa implements Action {
  readonly type = ConfigActionTypes.GetAlsa;
  constructor(public payload: any) {}
}

export class StoreGetAlsa implements Action {
  readonly type = ConfigActionTypes.StoreGetAlsa;
  constructor(public payload: any) {}
}

export class UpdateAlsaParameter implements Action {
  readonly type = ConfigActionTypes.UpdateAlsaParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateAlsaParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateAlsaParameter;
  constructor(public payload: any) {}
}

export class SwitchAlsaParameter implements Action {
  readonly type = ConfigActionTypes.SwitchAlsaParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchAlsaParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchAlsaParameter;
  constructor(public payload: any) {}
}

export class AddAlsaParameter implements Action {
  readonly type = ConfigActionTypes.AddAlsaParameter;
  constructor(public payload: any) {}
}

export class StoreAddAlsaParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddAlsaParameter;
  constructor(public payload: any) {}
}

export class DelAlsaParameter implements Action {
  readonly type = ConfigActionTypes.DelAlsaParameter;
  constructor(public payload: any) {}
}

export class StoreDelAlsaParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelAlsaParameter;
  constructor(public payload: any) {}
}

export class StoreGotAlsaError implements Action {
  readonly type = ConfigActionTypes.StoreGotAlsaError;
  constructor(public payload: any) {}
}

export class StoreDropNewAlsaParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewAlsaParameter;
  constructor(public payload: any) {}
}

export class StoreNewAlsaParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewAlsaParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetAlsa
  | StoreGetAlsa
  | UpdateAlsaParameter
  | StoreUpdateAlsaParameter
  | SwitchAlsaParameter
  | StoreSwitchAlsaParameter
  | AddAlsaParameter
  | StoreAddAlsaParameter
  | DelAlsaParameter
  | StoreDelAlsaParameter
  | StoreGotAlsaError
  | StoreDropNewAlsaParameter
  | StoreNewAlsaParameter
;
