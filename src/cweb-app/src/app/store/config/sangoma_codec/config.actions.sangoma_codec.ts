
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetSangomaCodec = 'GetSangomaCodec',
  StoreGetSangomaCodec = 'StoreGetSangomaCodec',
  UpdateSangomaCodecParameter = 'UpdateSangomaCodecParameter',
  StoreUpdateSangomaCodecParameter = 'StoreUpdateSangomaCodecParameter',
  SwitchSangomaCodecParameter = 'SwitchSangomaCodecParameter',
  StoreSwitchSangomaCodecParameter = 'StoreSwitchSangomaCodecParameter',
  AddSangomaCodecParameter = 'AddSangomaCodecParameter',
  StoreAddSangomaCodecParameter = 'StoreAddSangomaCodecParameter',
  DelSangomaCodecParameter = 'DelSangomaCodecParameter',
  StoreDelSangomaCodecParameter = 'StoreDelSangomaCodecParameter',
  StoreNewSangomaCodecParameter = 'StoreNewSangomaCodecParameter',
  StoreDropNewSangomaCodecParameter = 'StoreDropNewSangomaCodecParameter',
  StoreGotSangomaCodecError = 'StoreGotSangomaCodecError',
}

export class GetSangomaCodec implements Action {
  readonly type = ConfigActionTypes.GetSangomaCodec;
  constructor(public payload: any) {}
}

export class StoreGetSangomaCodec implements Action {
  readonly type = ConfigActionTypes.StoreGetSangomaCodec;
  constructor(public payload: any) {}
}

export class UpdateSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.UpdateSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class SwitchSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.SwitchSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class AddSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.AddSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class StoreAddSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class DelSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.DelSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class StoreDelSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class StoreGotSangomaCodecError implements Action {
  readonly type = ConfigActionTypes.StoreGotSangomaCodecError;
  constructor(public payload: any) {}
}

export class StoreDropNewSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewSangomaCodecParameter;
  constructor(public payload: any) {}
}

export class StoreNewSangomaCodecParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewSangomaCodecParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetSangomaCodec
  | StoreGetSangomaCodec
  | UpdateSangomaCodecParameter
  | StoreUpdateSangomaCodecParameter
  | SwitchSangomaCodecParameter
  | StoreSwitchSangomaCodecParameter
  | AddSangomaCodecParameter
  | StoreAddSangomaCodecParameter
  | DelSangomaCodecParameter
  | StoreDelSangomaCodecParameter
  | StoreGotSangomaCodecError
  | StoreDropNewSangomaCodecParameter
  | StoreNewSangomaCodecParameter
;

