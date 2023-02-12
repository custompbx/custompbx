import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotOspError = 'StoreGotOspError',
  GetOsp = 'GetOsp',
  StoreGetOsp = 'StoreGetOsp',
  GetOspProfileParameters = 'GetOspProfileParameters',
  StoreGetOspProfileParameters = 'StoreGetOspProfileParameters',
  UpdateOspParameter = 'UpdateOspParameter',
  StoreUpdateOspParameter = 'StoreUpdateOspParameter',
  SwitchOspParameter = 'SwitchOspParameter',
  StoreSwitchOspParameter = 'StoreSwitchOspParameter',
  AddOspParameter = 'AddOspParameter',
  StoreAddOspParameter = 'StoreAddOspParameter',
  DelOspParameter = 'DelOspParameter',
  StoreDelOspParameter = 'StoreDelOspParameter',
  StoreNewOspParameter = 'StoreNewOspParameter',
  StoreDropNewOspParameter = 'StoreDropNewOspParameter',
  AddOspProfileParameter = 'AddOspProfileParameter',
  StoreAddOspProfileParameter = 'StoreAddOspProfileParameter',
  UpdateOspProfileParameter = 'UpdateOspProfileParameter',
  StoreUpdateOspProfileParameter = 'StoreUpdateOspProfileParameter',
  SwitchOspProfileParameter = 'SwitchOspProfileParameter',
  StoreSwitchOspProfileParameter = 'StoreSwitchOspProfileParameter',
  DelOspProfileParameter = 'DelOspProfileParameter',
  StoreDelOspProfileParameter = 'StoreDelOspProfileParameter',
  StoreNewOspProfileParameter = 'StoreNewOspProfileParameter',
  StoreDropNewOspProfileParameter = 'StoreDropNewOspProfileParameter',
  StorePasteOspProfileParameters = 'StorePasteOspProfileParameters',
  AddOspProfile = 'AddOspProfile',
  StoreAddOspProfile = 'StoreAddOspProfile',
  DelOspProfile = 'DelOspProfile',
  StoreDelOspProfile = 'StoreDelOspProfile',
  UpdateOspProfile = 'UpdateOspProfile',
  StoreUpdateOspProfile = 'StoreUpdateOspProfile',
}

export class GetOsp implements Action {
  readonly type = ConfigActionTypes.GetOsp;
  constructor(public payload: any) {}
}

export class StoreGetOsp implements Action {
  readonly type = ConfigActionTypes.StoreGetOsp;
  constructor(public payload: any) {}
}

export class GetOspProfileParameters implements Action {
  readonly type = ConfigActionTypes.GetOspProfileParameters;
  constructor(public payload: any) {}
}

export class StoreGetOspProfileParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetOspProfileParameters;
  constructor(public payload: any) {}
}

export class UpdateOspParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOspParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOspParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOspParameter;
  constructor(public payload: any) {}
}

export class SwitchOspParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOspParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOspParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOspParameter;
  constructor(public payload: any) {}
}

export class AddOspParameter implements Action {
  readonly type = ConfigActionTypes.AddOspParameter;
  constructor(public payload: any) {}
}

export class StoreAddOspParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOspParameter;
  constructor(public payload: any) {}
}

export class DelOspParameter implements Action {
  readonly type = ConfigActionTypes.DelOspParameter;
  constructor(public payload: any) {}
}

export class StoreDelOspParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelOspParameter;
  constructor(public payload: any) {}
}

export class AddOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.AddOspProfileParameter;
  constructor(public payload: any) {}
}

export class StoreAddOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOspProfileParameter;
  constructor(public payload: any) {}
}

export class UpdateOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOspProfileParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOspProfileParameter;
  constructor(public payload: any) {}
}

export class SwitchOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOspProfileParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOspProfileParameter;
  constructor(public payload: any) {}
}

export class DelOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.DelOspProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDelOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelOspProfileParameter;
  constructor(public payload: any) {}
}

export class StoreNewOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOspProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewOspProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOspProfileParameter;
  constructor(public payload: any) {}
}

export class StorePasteOspProfileParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteOspProfileParameters;
  constructor(public payload: any) {}
}

export class AddOspProfile implements Action {
  readonly type = ConfigActionTypes.AddOspProfile;
  constructor(public payload: any) {}
}

export class StoreAddOspProfile implements Action {
  readonly type = ConfigActionTypes.StoreAddOspProfile;
  constructor(public payload: any) {}
}

export class DelOspProfile implements Action {
  readonly type = ConfigActionTypes.DelOspProfile;
  constructor(public payload: any) {}
}

export class StoreDelOspProfile implements Action {
  readonly type = ConfigActionTypes.StoreDelOspProfile;
  constructor(public payload: any) {}
}

export class UpdateOspProfile implements Action {
  readonly type = ConfigActionTypes.UpdateOspProfile;
  constructor(public payload: any) {}
}

export class StoreUpdateOspProfile implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOspProfile;
  constructor(public payload: any) {}
}

export class StoreGotOspError implements Action {
  readonly type = ConfigActionTypes.StoreGotOspError;
  constructor(public payload: any) {}
}

export class StoreDropNewOspParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOspParameter;
  constructor(public payload: any) {}
}

export class StoreNewOspParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOspParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetOsp
  | StoreGetOsp
  | GetOspProfileParameters
  | StoreGetOspProfileParameters
  | UpdateOspParameter
  | StoreUpdateOspParameter
  | SwitchOspParameter
  | StoreSwitchOspParameter
  | AddOspParameter
  | StoreAddOspParameter
  | DelOspParameter
  | StoreDelOspParameter
  | AddOspProfileParameter
  | StoreAddOspProfileParameter
  | UpdateOspProfileParameter
  | StoreUpdateOspProfileParameter
  | SwitchOspProfileParameter
  | StoreSwitchOspProfileParameter
  | DelOspProfileParameter
  | StoreDelOspProfileParameter
  | StoreNewOspProfileParameter
  | StoreDropNewOspProfileParameter
  | StorePasteOspProfileParameters
  | AddOspProfile
  | StoreAddOspProfile
  | DelOspProfile
  | StoreDelOspProfile
  | UpdateOspProfile
  | StoreUpdateOspProfile
  | StoreGotOspError
  | StoreDropNewOspParameter
  | StoreNewOspParameter
;

