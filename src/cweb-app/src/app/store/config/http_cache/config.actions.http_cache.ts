
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetHttpCache = 'GetHttpCache',
  StoreGetHttpCache = 'StoreGetHttpCache',
  UpdateHttpCacheParameter = 'UpdateHttpCacheParameter',
  StoreUpdateHttpCacheParameter = 'StoreUpdateHttpCacheParameter',
  SwitchHttpCacheParameter = 'SwitchHttpCacheParameter',
  StoreSwitchHttpCacheParameter = 'StoreSwitchHttpCacheParameter',
  AddHttpCacheParameter = 'AddHttpCacheParameter',
  StoreAddHttpCacheParameter = 'StoreAddHttpCacheParameter',
  DelHttpCacheParameter = 'DelHttpCacheParameter',
  StoreDelHttpCacheParameter = 'StoreDelHttpCacheParameter',
  StoreNewHttpCacheParameter = 'StoreNewHttpCacheParameter',
  StoreDropNewHttpCacheParameter = 'StoreDropNewHttpCacheParameter',
  StoreGotHttpCacheError = 'StoreGotHttpCacheError',
  UpdateHttpCacheProfileParam = 'UpdateHttpCacheProfileParam',
  StoreUpdateHttpCacheProfileParam = 'StoreUpdateHttpCacheProfileParam',
  SwitchHttpCacheProfileParam = 'SwitchHttpCacheProfileParam',
  StoreSwitchHttpCacheProfileParam = 'StoreSwitchHttpCacheProfileParam',
  AddHttpCacheProfileParam = 'AddHttpCacheProfileParam',
  StoreAddHttpCacheProfileParam = 'StoreAddHttpCacheProfileParam',
  DelHttpCacheProfileParam = 'DelHttpCacheProfileParam',
  StoreDelHttpCacheProfileParam = 'StoreDelHttpCacheProfileParam',
  StoreNewHttpCacheProfileParam = 'StoreNewHttpCacheProfileParam',
  StoreDropNewHttpCacheProfileParam = 'StoreDropNewHttpCacheProfileParam',
  AddHttpCacheProfile = 'AddHttpCacheProfile',
  StoreAddHttpCacheProfile = 'StoreAddHttpCacheProfile',
  DelHttpCacheProfile = 'DelHttpCacheProfile',
  StoreDelHttpCacheProfile = 'StoreDelHttpCacheProfile',
  RenameHttpCacheProfile = 'RenameHttpCacheProfile',
  StoreRenameHttpCacheProfile = 'StoreRenameHttpCacheProfile',

  GetHttpCacheProfileParameters = 'GetHttpCacheProfileParameters',
  StoreGetHttpCacheProfileParameters = 'StoreGetHttpCacheProfileParameters',
  AddHttpCacheProfileDomain = 'AddHttpCacheProfileDomain',
  StoreAddHttpCacheProfileDomain = 'StoreAddHttpCacheProfileDomain',
  DelHttpCacheProfileDomain = 'DelHttpCacheProfileDomain',
  StoreDelHttpCacheProfileDomain = 'StoreDelHttpCacheProfileDomain',
  SwitchHttpCacheProfileDomain = 'SwitchHttpCacheProfileDomain',
  StoreSwitchHttpCacheProfileDomain = 'StoreSwitchHttpCacheProfileDomain',
  UpdateHttpCacheProfileDomain = 'UpdateHttpCacheProfileDomain',
  StoreUpdateHttpCacheProfileDomain = 'StoreUpdateHttpCacheProfileDomain',
  StoreDropNewHttpCacheProfileDomain = 'StoreDropNewHttpCacheProfileDomain',
  StoreNewHttpCacheProfileDomain = 'StoreNewHttpCacheProfileDomain',

  UpdateHttpCacheProfileAws = 'UpdateHttpCacheProfileAws',
  StoreUpdateHttpCacheProfileAws = 'StoreUpdateHttpCacheProfileAws',
  UpdateHttpCacheProfileAzure = 'UpdateHttpCacheProfileAzure',
  StoreUpdateHttpCacheProfileAzure = 'StoreUpdateHttpCacheProfileAzure',
}

export class GetHttpCache implements Action {
  readonly type = ConfigActionTypes.GetHttpCache;
  constructor(public payload: any) {}
}

export class StoreGetHttpCache implements Action {
  readonly type = ConfigActionTypes.StoreGetHttpCache;
  constructor(public payload: any) {}
}

export class UpdateHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.UpdateHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateHttpCacheParameter;
  constructor(public payload: any) {}
}

export class SwitchHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.SwitchHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchHttpCacheParameter;
  constructor(public payload: any) {}
}

export class AddHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.AddHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreAddHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddHttpCacheParameter;
  constructor(public payload: any) {}
}

export class DelHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.DelHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreDelHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreGotHttpCacheError implements Action {
  readonly type = ConfigActionTypes.StoreGotHttpCacheError;
  constructor(public payload: any) {}
}

export class StoreDropNewHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewHttpCacheParameter;
  constructor(public payload: any) {}
}

export class StoreNewHttpCacheParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewHttpCacheParameter;
  constructor(public payload: any) {}
}

export class UpdateHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.UpdateHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class SwitchHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.SwitchHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class AddHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.AddHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class DelHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.DelHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class StoreNewHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.StoreNewHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class StoreDropNewHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.StoreDropNewHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class AddHttpCacheProfile implements Action {
  readonly type = ConfigActionTypes.AddHttpCacheProfile;
  constructor(public payload: any) {}
}

export class StoreAddHttpCacheProfile implements Action {
  readonly type = ConfigActionTypes.StoreAddHttpCacheProfile;
  constructor(public payload: any) {}
}

export class DelHttpCacheProfile implements Action {
  readonly type = ConfigActionTypes.DelHttpCacheProfile;
  constructor(public payload: any) {}
}

export class StoreDelHttpCacheProfile implements Action {
  readonly type = ConfigActionTypes.StoreDelHttpCacheProfile;
  constructor(public payload: any) {}
}

export class RenameHttpCacheProfile implements Action {
  readonly type = ConfigActionTypes.RenameHttpCacheProfile;
  constructor(public payload: any) {}
}

export class StoreRenameHttpCacheProfile implements Action {
  readonly type = ConfigActionTypes.StoreRenameHttpCacheProfile;
  constructor(public payload: any) {}
}

export class StoreUpdateHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.StoreUpdateHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class StoreSwitchHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.StoreSwitchHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class StoreAddHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.StoreAddHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class StoreDelHttpCacheProfileParam implements Action {
  readonly type = ConfigActionTypes.StoreDelHttpCacheProfileParam;
  constructor(public payload: any) {}
}

export class GetHttpCacheProfileParameters implements Action {
  readonly type = ConfigActionTypes.GetHttpCacheProfileParameters;
  constructor(public payload: any) {}
}

export class StoreGetHttpCacheProfileParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetHttpCacheProfileParameters;
  constructor(public payload: any) {}
}

export class AddHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.AddHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class StoreAddHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.StoreAddHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class DelHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.DelHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class StoreDelHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.StoreDelHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class SwitchHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.SwitchHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class StoreSwitchHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.StoreSwitchHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class UpdateHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.UpdateHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class StoreUpdateHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.StoreUpdateHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class StoreDropNewHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.StoreDropNewHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class StoreNewHttpCacheProfileDomain implements Action {
  readonly type = ConfigActionTypes.StoreNewHttpCacheProfileDomain;
  constructor(public payload: any) {}
}

export class UpdateHttpCacheProfileAws implements Action {
  readonly type = ConfigActionTypes.UpdateHttpCacheProfileAws;
  constructor(public payload: any) {}
}

export class StoreUpdateHttpCacheProfileAws implements Action {
  readonly type = ConfigActionTypes.StoreUpdateHttpCacheProfileAws;
  constructor(public payload: any) {}
}

export class UpdateHttpCacheProfileAzure implements Action {
  readonly type = ConfigActionTypes.UpdateHttpCacheProfileAzure;
  constructor(public payload: any) {}
}

export class StoreUpdateHttpCacheProfileAzure implements Action {
  readonly type = ConfigActionTypes.StoreUpdateHttpCacheProfileAzure;
  constructor(public payload: any) {}
}

export type All =
  | GetHttpCache
  | StoreGetHttpCache
  | UpdateHttpCacheParameter
  | StoreUpdateHttpCacheParameter
  | SwitchHttpCacheParameter
  | StoreSwitchHttpCacheParameter
  | AddHttpCacheParameter
  | StoreAddHttpCacheParameter
  | DelHttpCacheParameter
  | StoreDelHttpCacheParameter
  | StoreGotHttpCacheError
  | StoreDropNewHttpCacheParameter
  | StoreNewHttpCacheParameter
  | UpdateHttpCacheProfileParam
  | SwitchHttpCacheProfileParam
  | AddHttpCacheProfileParam
  | DelHttpCacheProfileParam
  | StoreNewHttpCacheProfileParam
  | StoreDropNewHttpCacheProfileParam
  | AddHttpCacheProfile
  | StoreAddHttpCacheProfile
  | DelHttpCacheProfile
  | StoreDelHttpCacheProfile
  | RenameHttpCacheProfile
  | StoreRenameHttpCacheProfile
  | StoreUpdateHttpCacheProfileParam
  | StoreSwitchHttpCacheProfileParam
  | StoreAddHttpCacheProfileParam
  | StoreDelHttpCacheProfileParam
  | GetHttpCacheProfileParameters
  | StoreGetHttpCacheProfileParameters
  | AddHttpCacheProfileDomain
  | StoreAddHttpCacheProfileDomain
  | DelHttpCacheProfileDomain
  | StoreDelHttpCacheProfileDomain
  | SwitchHttpCacheProfileDomain
  | StoreSwitchHttpCacheProfileDomain
  | UpdateHttpCacheProfileDomain
  | StoreUpdateHttpCacheProfileDomain
  | StoreDropNewHttpCacheProfileDomain
  | StoreNewHttpCacheProfileDomain
  | UpdateHttpCacheProfileAws
  | StoreUpdateHttpCacheProfileAws
  | UpdateHttpCacheProfileAzure
  | StoreUpdateHttpCacheProfileAzure
;

