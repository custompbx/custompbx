import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotLcrError = 'StoreGotLcrError',
  GetLcr = 'GetLcr',
  StoreGetLcr = 'StoreGetLcr',
  GetLcrProfileParameters = 'GetLcrProfileParameters',
  StoreGetLcrProfileParameters = 'StoreGetLcrProfileParameters',
  UpdateLcrParameter = 'UpdateLcrParameter',
  StoreUpdateLcrParameter = 'StoreUpdateLcrParameter',
  SwitchLcrParameter = 'SwitchLcrParameter',
  StoreSwitchLcrParameter = 'StoreSwitchLcrParameter',
  AddLcrParameter = 'AddLcrParameter',
  StoreAddLcrParameter = 'StoreAddLcrParameter',
  DelLcrParameter = 'DelLcrParameter',
  StoreDelLcrParameter = 'StoreDelLcrParameter',
  StoreNewLcrParameter = 'StoreNewLcrParameter',
  StoreDropNewLcrParameter = 'StoreDropNewLcrParameter',
  AddLcrProfileParameter = 'AddLcrProfileParameter',
  StoreAddLcrProfileParameter = 'StoreAddLcrProfileParameter',
  UpdateLcrProfileParameter = 'UpdateLcrProfileParameter',
  StoreUpdateLcrProfileParameter = 'StoreUpdateLcrProfileParameter',
  SwitchLcrProfileParameter = 'SwitchLcrProfileParameter',
  StoreSwitchLcrProfileParameter = 'StoreSwitchLcrProfileParameter',
  DelLcrProfileParameter = 'DelLcrProfileParameter',
  StoreDelLcrProfileParameter = 'StoreDelLcrProfileParameter',
  StoreNewLcrProfileParameter = 'StoreNewLcrProfileParameter',
  StoreDropNewLcrProfileParameter = 'StoreDropNewLcrProfileParameter',
  StorePasteLcrProfileParameters = 'StorePasteLcrProfileParameters',
  AddLcrProfile = 'AddLcrProfile',
  StoreAddLcrProfile = 'StoreAddLcrProfile',
  DelLcrProfile = 'DelLcrProfile',
  StoreDelLcrProfile = 'StoreDelLcrProfile',
  UpdateLcrProfile = 'UpdateLcrProfile',
  StoreUpdateLcrProfile = 'StoreUpdateLcrProfile',
}

export class GetLcr implements Action {
  readonly type = ConfigActionTypes.GetLcr;
  constructor(public payload: any) {}
}

export class StoreGetLcr implements Action {
  readonly type = ConfigActionTypes.StoreGetLcr;
  constructor(public payload: any) {}
}

export class GetLcrProfileParameters implements Action {
  readonly type = ConfigActionTypes.GetLcrProfileParameters;
  constructor(public payload: any) {}
}

export class StoreGetLcrProfileParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetLcrProfileParameters;
  constructor(public payload: any) {}
}

export class UpdateLcrParameter implements Action {
  readonly type = ConfigActionTypes.UpdateLcrParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateLcrParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateLcrParameter;
  constructor(public payload: any) {}
}

export class SwitchLcrParameter implements Action {
  readonly type = ConfigActionTypes.SwitchLcrParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchLcrParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchLcrParameter;
  constructor(public payload: any) {}
}

export class AddLcrParameter implements Action {
  readonly type = ConfigActionTypes.AddLcrParameter;
  constructor(public payload: any) {}
}

export class StoreAddLcrParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddLcrParameter;
  constructor(public payload: any) {}
}

export class DelLcrParameter implements Action {
  readonly type = ConfigActionTypes.DelLcrParameter;
  constructor(public payload: any) {}
}

export class StoreDelLcrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelLcrParameter;
  constructor(public payload: any) {}
}

export class AddLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.AddLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StoreAddLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddLcrProfileParameter;
  constructor(public payload: any) {}
}

export class UpdateLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.UpdateLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateLcrProfileParameter;
  constructor(public payload: any) {}
}

export class SwitchLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.SwitchLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchLcrProfileParameter;
  constructor(public payload: any) {}
}

export class DelLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.DelLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDelLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StoreNewLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewLcrProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewLcrProfileParameter;
  constructor(public payload: any) {}
}

export class StorePasteLcrProfileParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteLcrProfileParameters;
  constructor(public payload: any) {}
}

export class AddLcrProfile implements Action {
  readonly type = ConfigActionTypes.AddLcrProfile;
  constructor(public payload: any) {}
}

export class StoreAddLcrProfile implements Action {
  readonly type = ConfigActionTypes.StoreAddLcrProfile;
  constructor(public payload: any) {}
}

export class DelLcrProfile implements Action {
  readonly type = ConfigActionTypes.DelLcrProfile;
  constructor(public payload: any) {}
}

export class StoreDelLcrProfile implements Action {
  readonly type = ConfigActionTypes.StoreDelLcrProfile;
  constructor(public payload: any) {}
}

export class UpdateLcrProfile implements Action {
  readonly type = ConfigActionTypes.UpdateLcrProfile;
  constructor(public payload: any) {}
}

export class StoreUpdateLcrProfile implements Action {
  readonly type = ConfigActionTypes.StoreUpdateLcrProfile;
  constructor(public payload: any) {}
}

export class StoreGotLcrError implements Action {
  readonly type = ConfigActionTypes.StoreGotLcrError;
  constructor(public payload: any) {}
}

export class StoreDropNewLcrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewLcrParameter;
  constructor(public payload: any) {}
}

export class StoreNewLcrParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewLcrParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetLcr
  | StoreGetLcr
  | GetLcrProfileParameters
  | StoreGetLcrProfileParameters
  | UpdateLcrParameter
  | StoreUpdateLcrParameter
  | SwitchLcrParameter
  | StoreSwitchLcrParameter
  | AddLcrParameter
  | StoreAddLcrParameter
  | DelLcrParameter
  | StoreDelLcrParameter
  | AddLcrProfileParameter
  | StoreAddLcrProfileParameter
  | UpdateLcrProfileParameter
  | StoreUpdateLcrProfileParameter
  | SwitchLcrProfileParameter
  | StoreSwitchLcrProfileParameter
  | DelLcrProfileParameter
  | StoreDelLcrProfileParameter
  | StoreNewLcrProfileParameter
  | StoreDropNewLcrProfileParameter
  | StorePasteLcrProfileParameters
  | AddLcrProfile
  | StoreAddLcrProfile
  | DelLcrProfile
  | StoreDelLcrProfile
  | UpdateLcrProfile
  | StoreUpdateLcrProfile
  | StoreGotLcrError
  | StoreDropNewLcrParameter
  | StoreNewLcrParameter
;

