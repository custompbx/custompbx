import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotVoicemailError = 'StoreGotVoicemailError',
  GetVoicemailSettings = 'GetVoicemailSettings',
  StoreGetVoicemailSettings = 'StoreGetVoicemailSettings',
  UpdateVoicemailSetting = 'UpdateVoicemailSetting',
  StoreUpdateVoicemailSetting = 'StoreUpdateVoicemailSetting',
  SwitchVoicemailSetting = 'SwitchVoicemailSetting',
  StoreSwitchVoicemailSetting = 'StoreSwitchVoicemailSetting',
  AddVoicemailSetting = 'AddVoicemailSetting',
  StoreAddVoicemailSetting = 'StoreAddVoicemailSetting',
  DelVoicemailSetting = 'DelVoicemailSetting',
  StoreDelVoicemailSetting = 'StoreDelVoicemailSetting',
  StoreNewVoicemailSetting = 'StoreNewVoicemailSetting',
  StoreDropNewVoicemailSetting = 'StoreDropNewVoicemailSetting',

  GetVoicemailProfileParameters = 'GetVoicemailProfileParameters',
  StoreGetVoicemailProfileParameters = 'StoreGetVoicemailProfileParameters',
  AddVoicemailProfileParameter = 'AddVoicemailProfileParameter',
  StoreAddVoicemailProfileParameter = 'StoreAddVoicemailProfileParameter',
  UpdateVoicemailProfileParameter = 'UpdateVoicemailProfileParameter',
  StoreUpdateVoicemailProfileParameter = 'StoreUpdateVoicemailProfileParameter',
  SwitchVoicemailProfileParameter = 'SwitchVoicemailProfileParameter',
  StoreSwitchVoicemailProfileParameter = 'StoreSwitchVoicemailProfileParameter',
  DelVoicemailProfileParameter = 'DelVoicemailProfileParameter',
  StoreDelVoicemailProfileParameter = 'StoreDelVoicemailProfileParameter',
  StoreNewVoicemailProfileParameter = 'StoreNewVoicemailProfileParameter',
  StoreDropNewVoicemailProfileParameter = 'StoreDropNewVoicemailProfileParameter',
  StorePasteVoicemailProfileParameters = 'StorePasteVoicemailProfileParameters',

  GetVoicemailProfiles = 'GetVoicemailProfiles',
  StoreGetVoicemailProfiles = 'StoreGetVoicemailProfiles',
  AddVoicemailProfile = 'AddVoicemailProfile',
  StoreAddVoicemailProfile = 'StoreAddVoicemailProfile',
  DelVoicemailProfile = 'DelVoicemailProfile',
  StoreDelVoicemailProfile = 'StoreDelVoicemailProfile',
  UpdateVoicemailProfile = 'UpdateVoicemailProfile',
  StoreUpdateVoicemailProfile = 'StoreUpdateVoicemailProfile',

}

export class GetVoicemailSettings implements Action {
  readonly type = ConfigActionTypes.GetVoicemailSettings;
  constructor(public payload: any) {}
}

export class StoreGetVoicemailSettings implements Action {
  readonly type = ConfigActionTypes.StoreGetVoicemailSettings;
  constructor(public payload: any) {}
}

export class UpdateVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.UpdateVoicemailSetting;
  constructor(public payload: any) {}
}

export class StoreUpdateVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.StoreUpdateVoicemailSetting;
  constructor(public payload: any) {}
}

export class SwitchVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.SwitchVoicemailSetting;
  constructor(public payload: any) {}
}

export class StoreSwitchVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.StoreSwitchVoicemailSetting;
  constructor(public payload: any) {}
}

export class AddVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.AddVoicemailSetting;
  constructor(public payload: any) {}
}

export class StoreAddVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.StoreAddVoicemailSetting;
  constructor(public payload: any) {}
}

export class DelVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.DelVoicemailSetting;
  constructor(public payload: any) {}
}

export class StoreDelVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.StoreDelVoicemailSetting;
  constructor(public payload: any) {}
}

export class GetVoicemailProfileParameters implements Action {
  readonly type = ConfigActionTypes.GetVoicemailProfileParameters;
  constructor(public payload: any) {}
}

export class StoreGetVoicemailProfileParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetVoicemailProfileParameters;
  constructor(public payload: any) {}
}

export class AddVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.AddVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StoreAddVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class UpdateVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.UpdateVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class SwitchVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.SwitchVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class DelVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.DelVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDelVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StoreNewVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewVoicemailProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewVoicemailProfileParameter;
  constructor(public payload: any) {}
}

export class StorePasteVoicemailProfileParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteVoicemailProfileParameters;
  constructor(public payload: any) {}
}

export class AddVoicemailProfile implements Action {
  readonly type = ConfigActionTypes.AddVoicemailProfile;
  constructor(public payload: any) {}
}

export class StoreAddVoicemailProfile implements Action {
  readonly type = ConfigActionTypes.StoreAddVoicemailProfile;
  constructor(public payload: any) {}
}

export class DelVoicemailProfile implements Action {
  readonly type = ConfigActionTypes.DelVoicemailProfile;
  constructor(public payload: any) {}
}

export class StoreDelVoicemailProfile implements Action {
  readonly type = ConfigActionTypes.StoreDelVoicemailProfile;
  constructor(public payload: any) {}
}

export class UpdateVoicemailProfile implements Action {
  readonly type = ConfigActionTypes.UpdateVoicemailProfile;
  constructor(public payload: any) {}
}

export class StoreUpdateVoicemailProfile implements Action {
  readonly type = ConfigActionTypes.StoreUpdateVoicemailProfile;
  constructor(public payload: any) {}
}

export class StoreGotVoicemailError implements Action {
  readonly type = ConfigActionTypes.StoreGotVoicemailError;
  constructor(public payload: any) {}
}

export class StoreDropNewVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.StoreDropNewVoicemailSetting;
  constructor(public payload: any) {}
}

export class StoreNewVoicemailSetting implements Action {
  readonly type = ConfigActionTypes.StoreNewVoicemailSetting;
  constructor(public payload: any) {}
}

export class GetVoicemailProfiles implements Action {
  readonly type = ConfigActionTypes.GetVoicemailProfiles;
  constructor(public payload: any) {}
}

export class StoreGetVoicemailProfiles implements Action {
  readonly type = ConfigActionTypes.StoreGetVoicemailProfiles;
  constructor(public payload: any) {}
}

export type All =
  | GetVoicemailSettings
  | StoreGetVoicemailSettings
  | UpdateVoicemailSetting
  | StoreUpdateVoicemailSetting
  | SwitchVoicemailSetting
  | StoreSwitchVoicemailSetting
  | AddVoicemailSetting
  | StoreAddVoicemailSetting
  | DelVoicemailSetting
  | StoreDelVoicemailSetting
  | StoreGotVoicemailError
  | StoreDropNewVoicemailSetting
  | StoreNewVoicemailSetting
  | GetVoicemailProfileParameters
  | StoreGetVoicemailProfileParameters
  | AddVoicemailProfileParameter
  | StoreAddVoicemailProfileParameter
  | UpdateVoicemailProfileParameter
  | StoreUpdateVoicemailProfileParameter
  | SwitchVoicemailProfileParameter
  | StoreSwitchVoicemailProfileParameter
  | DelVoicemailProfileParameter
  | StoreDelVoicemailProfileParameter
  | StoreNewVoicemailProfileParameter
  | StoreDropNewVoicemailProfileParameter
  | StorePasteVoicemailProfileParameters
  | AddVoicemailProfile
  | StoreAddVoicemailProfile
  | DelVoicemailProfile
  | StoreDelVoicemailProfile
  | UpdateVoicemailProfile
  | StoreUpdateVoicemailProfile
  | GetVoicemailProfiles
  | StoreGetVoicemailProfiles
;

