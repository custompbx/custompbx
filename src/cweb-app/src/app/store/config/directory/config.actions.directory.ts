import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotDirectoryError = 'StoreGotDirectoryError',
  GetDirectory = 'GetDirectory',
  StoreGetDirectory = 'StoreGetDirectory',
  GetDirectoryProfileParameters = 'GetDirectoryProfileParameters',
  StoreGetDirectoryProfileParameters = 'StoreGetDirectoryProfileParameters',
  UpdateDirectoryParameter = 'UpdateDirectoryParameter',
  StoreUpdateDirectoryParameter = 'StoreUpdateDirectoryParameter',
  SwitchDirectoryParameter = 'SwitchDirectoryParameter',
  StoreSwitchDirectoryParameter = 'StoreSwitchDirectoryParameter',
  AddDirectoryParameter = 'AddDirectoryParameter',
  StoreAddDirectoryParameter = 'StoreAddDirectoryParameter',
  DelDirectoryParameter = 'DelDirectoryParameter',
  StoreDelDirectoryParameter = 'StoreDelDirectoryParameter',
  StoreNewDirectoryParameter = 'StoreNewDirectoryParameter',
  StoreDropNewDirectoryParameter = 'StoreDropNewDirectoryParameter',
  AddDirectoryProfileParameter = 'AddDirectoryProfileParameter',
  StoreAddDirectoryProfileParameter = 'StoreAddDirectoryProfileParameter',
  UpdateDirectoryProfileParameter = 'UpdateDirectoryProfileParameter',
  StoreUpdateDirectoryProfileParameter = 'StoreUpdateDirectoryProfileParameter',
  SwitchDirectoryProfileParameter = 'SwitchDirectoryProfileParameter',
  StoreSwitchDirectoryProfileParameter = 'StoreSwitchDirectoryProfileParameter',
  DelDirectoryProfileParameter = 'DelDirectoryProfileParameter',
  StoreDelDirectoryProfileParameter = 'StoreDelDirectoryProfileParameter',
  StoreNewDirectoryProfileParameter = 'StoreNewDirectoryProfileParameter',
  StoreDropNewDirectoryProfileParameter = 'StoreDropNewDirectoryProfileParameter',
  StorePasteDirectoryProfileParameters = 'StorePasteDirectoryProfileParameters',
  AddDirectoryProfile = 'AddDirectoryProfile',
  StoreAddDirectoryProfile = 'StoreAddDirectoryProfile',
  DelDirectoryProfile = 'DelDirectoryProfile',
  StoreDelDirectoryProfile = 'StoreDelDirectoryProfile',
  UpdateDirectoryProfile = 'UpdateDirectoryProfile',
  StoreUpdateDirectoryProfile = 'StoreUpdateDirectoryProfile',
}

export class GetDirectory implements Action {
  readonly type = ConfigActionTypes.GetDirectory;
  constructor(public payload: any) {}
}

export class StoreGetDirectory implements Action {
  readonly type = ConfigActionTypes.StoreGetDirectory;
  constructor(public payload: any) {}
}

export class GetDirectoryProfileParameters implements Action {
  readonly type = ConfigActionTypes.GetDirectoryProfileParameters;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryProfileParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetDirectoryProfileParameters;
  constructor(public payload: any) {}
}

export class UpdateDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.UpdateDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDirectoryParameter;
  constructor(public payload: any) {}
}

export class SwitchDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.SwitchDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchDirectoryParameter;
  constructor(public payload: any) {}
}

export class AddDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.AddDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddDirectoryParameter;
  constructor(public payload: any) {}
}

export class DelDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.DelDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreDelDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelDirectoryParameter;
  constructor(public payload: any) {}
}

export class AddDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.AddDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class UpdateDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.UpdateDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class SwitchDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.SwitchDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class DelDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.DelDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDelDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StoreNewDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewDirectoryProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewDirectoryProfileParameter;
  constructor(public payload: any) {}
}

export class StorePasteDirectoryProfileParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteDirectoryProfileParameters;
  constructor(public payload: any) {}
}

export class AddDirectoryProfile implements Action {
  readonly type = ConfigActionTypes.AddDirectoryProfile;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryProfile implements Action {
  readonly type = ConfigActionTypes.StoreAddDirectoryProfile;
  constructor(public payload: any) {}
}

export class DelDirectoryProfile implements Action {
  readonly type = ConfigActionTypes.DelDirectoryProfile;
  constructor(public payload: any) {}
}

export class StoreDelDirectoryProfile implements Action {
  readonly type = ConfigActionTypes.StoreDelDirectoryProfile;
  constructor(public payload: any) {}
}

export class UpdateDirectoryProfile implements Action {
  readonly type = ConfigActionTypes.UpdateDirectoryProfile;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryProfile implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDirectoryProfile;
  constructor(public payload: any) {}
}

export class StoreGotDirectoryError implements Action {
  readonly type = ConfigActionTypes.StoreGotDirectoryError;
  constructor(public payload: any) {}
}

export class StoreDropNewDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewDirectoryParameter;
  constructor(public payload: any) {}
}

export class StoreNewDirectoryParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewDirectoryParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetDirectory
  | StoreGetDirectory
  | GetDirectoryProfileParameters
  | StoreGetDirectoryProfileParameters
  | UpdateDirectoryParameter
  | StoreUpdateDirectoryParameter
  | SwitchDirectoryParameter
  | StoreSwitchDirectoryParameter
  | AddDirectoryParameter
  | StoreAddDirectoryParameter
  | DelDirectoryParameter
  | StoreDelDirectoryParameter
  | AddDirectoryProfileParameter
  | StoreAddDirectoryProfileParameter
  | UpdateDirectoryProfileParameter
  | StoreUpdateDirectoryProfileParameter
  | SwitchDirectoryProfileParameter
  | StoreSwitchDirectoryProfileParameter
  | DelDirectoryProfileParameter
  | StoreDelDirectoryProfileParameter
  | StoreNewDirectoryProfileParameter
  | StoreDropNewDirectoryProfileParameter
  | StorePasteDirectoryProfileParameters
  | AddDirectoryProfile
  | StoreAddDirectoryProfile
  | DelDirectoryProfile
  | StoreDelDirectoryProfile
  | UpdateDirectoryProfile
  | StoreUpdateDirectoryProfile
  | StoreGotDirectoryError
  | StoreDropNewDirectoryParameter
  | StoreNewDirectoryParameter
;

