import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GET_SOFIA_GLOBAL_SETTINGS = '[Config] Get_sofia_global_settings',
  STORE_GET_SOFIA_GLOBAL_SETTINGS = '[Config] Store_get_sofia_global_settings',
  GET_SOFIA_PROFILES = '[Config] Get_sofia_profiles',
  STORE_GET_SOFIA_PROFILES = '[Config] Store_get_sofia_profiles',
  GET_SOFIA_PROFILES_PARAMS = '[Config] Get_sofia_profiles_params',
  STORE_GET_SOFIA_PROFILES_PARAMS = '[Config] Store_get_sofia_profiles_params',
  UPDATE_SOFIA_GLOBAL_SETTING = '[Config] Update_sofia_global_setting',
  STORE_UPDATE_SOFIA_GLOBAL_SETTING = '[Config] Store_update_sofia_global_setting',
  SWITCH_SOFIA_GLOBAL_SETTING = '[Config] Switch_sofia_global_setting',
  STORE_SWITCH_SOFIA_GLOBAL_SETTING = '[Config] Store_switch_sofia_global_setting',
  ADD_SOFIA_GLOBAL_SETTING = '[Config] Add_sofia_global_setting',
  STORE_ADD_SOFIA_GLOBAL_SETTING = '[Config] Store_add_sofia_global_setting',
  DEL_SOFIA_GLOBAL_SETTING = '[Config] Del_sofia_global_setting',
  STORE_DEL_SOFIA_GLOBAL_SETTING = '[Config] Store_del_sofia_global_setting',
  STORE_NEW_SOFIA_GLOBAL_SETTING = '[Config] Store_new_sofia_global_setting',
  STORE_DROP_NEW_SOFIA_GLOBAL_SETTING = '[Config] Store_drop_new_sofia_global_setting',

  UPDATE_SOFIA_PROFILE_PARAM = '[Config] Update_sofia_profile_param',
  STORE_UPDATE_SOFIA_PROFILE_PARAM = '[Config] Store_update_sofia_profile_param',
  SWITCH_SOFIA_PROFILE_PARAM = '[Config] Switch_sofia_profile_param',
  STORE_SWITCH_SOFIA_PROFILE_PARAM = '[Config] Store_switch_sofia_profile_param',
  ADD_SOFIA_PROFILE_PARAM = '[Config] Add_sofia_profile_param',
  STORE_ADD_SOFIA_PROFILE_PARAM = '[Config] Store_add_sofia_profile_param',
  DEL_SOFIA_PROFILE_PARAM = '[Config] Del_sofia_profile_param',
  STORE_DEL_SOFIA_PROFILE_PARAM = '[Config] Store_del_sofia_profile_param',
  STORE_NEW_SOFIA_PROFILE_PARAM = '[Config] Store_new_sofia_profile_param',
  STORE_DROP_NEW_SOFIA_PROFILE_PARAM = '[Config] Store_drop_new_sofia_profile_param',
  STORE_PASTE_SOFIA_PROFILE_PARAMS = '[Config][Store][Parameters] Sofia Profile',

  GET_SOFIA_PROFILE_GATEWAYS = '[Config] Get_sofia_profile_gateways',
  STORE_GET_SOFIA_PROFILE_GATEWAYS = '[Config] Store_get_sofia_profile_gateways',
  STORE_PASTE_SOFIA_GATEWAY_PARAMS = '[Gateways][Store][Parameters] Sofia Gateway',
  STORE_PASTE_SOFIA_GATEWAY_VARS = '[Gateways][Store][Variables] Sofia Gateway',
  GetSofiaProfileGatewayVariables = 'GetSofiaProfileGatewayVariables',
  StoreGetSofiaProfileGatewayVariables = 'GetSofiaProfileGatewayVariables',
  GetSofiaProfileGatewayParameters = 'GetSofiaProfileGatewayParameters',
  StoreGetSofiaProfileGatewayParameters = 'StoreGetSofiaProfileGatewayParameters',

  UPDATE_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Update_sofia_profile_gateway_param',
  STORE_UPDATE_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Store_update_sofia_profile_gateway_param',
  SWITCH_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Switch_sofia_profile_gateway_param',
  STORE_SWITCH_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Store_switch_sofia_profile_gateway_param',
  ADD_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Add_sofia_profile_gateway_param',
  STORE_ADD_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Store_add_sofia_profile_gateway_param',
  DEL_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Del_sofia_profile_gateway_param',
  STORE_DEL_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Store_del_sofia_profile_gateway_param',
  STORE_NEW_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Store_new_sofia_profile_gateway_param',
  STORE_DROP_NEW_SOFIA_PROFILE_GATEWAY_PARAM = '[Config] Store_drop_new_sofia_profile_gateway_param',

  UPDATE_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Update_sofia_profile_gateway_var',
  STORE_UPDATE_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Store_update_sofia_profile_gateway_var',
  SWITCH_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Switch_sofia_profile_gateway_var',
  STORE_SWITCH_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Store_switch_sofia_profile_gateway_var',
  ADD_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Add_sofia_profile_gateway_var',
  STORE_ADD_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Store_add_sofia_profile_gateway_var',
  DEL_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Del_sofia_profile_gateway_var',
  STORE_DEL_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Store_del_sofia_profile_gateway_var',
  STORE_NEW_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Store_new_sofia_profile_gateway_var',
  STORE_DROP_NEW_SOFIA_PROFILE_GATEWAY_VAR = '[Config] Store_drop_new_sofia_profile_gateway_var',

  GET_SOFIA_PROFILE_DOMAINS = '[Config] Get_sofia_profile_domains',
  STORE_GET_SOFIA_PROFILE_DOMAINS = '[Config] Store_get_sofia_profile_domains',

  UPDATE_SOFIA_PROFILE_DOMAIN = '[Config] Update_sofia_profile_domain',
  STORE_UPDATE_SOFIA_PROFILE_DOMAIN = '[Config] Store_update_sofia_profile_domain',
  SWITCH_SOFIA_PROFILE_DOMAIN = '[Config] Switch_sofia_profile_domain',
  STORE_SWITCH_SOFIA_PROFILE_DOMAIN = '[Config] Store_switch_sofia_profile_domain',
  ADD_SOFIA_PROFILE_DOMAIN = '[Config] Add_sofia_profile_domain',
  STORE_ADD_SOFIA_PROFILE_DOMAIN = '[Config] Store_add_sofia_profile_domain',
  DEL_SOFIA_PROFILE_DOMAIN = '[Config] Del_sofia_profile_domain',
  STORE_DEL_SOFIA_PROFILE_DOMAIN = '[Config] Store_del_sofia_profile_domain',
  STORE_NEW_SOFIA_PROFILE_DOMAIN = '[Config] Store_new_sofia_profile_domain',
  STORE_DROP_NEW_SOFIA_PROFILE_DOMAIN = '[Config] Store_drop_new_sofia_profile_domain',

  GET_SOFIA_PROFILE_ALIASES = '[Config] Get_sofia_profile_aliases',
  STORE_GET_SOFIA_PROFILE_ALIASES = '[Config] Store_get_sofia_profile_aliases',

  UPDATE_SOFIA_PROFILE_ALIAS = '[Config] Update_sofia_profile_alias',
  STORE_UPDATE_SOFIA_PROFILE_ALIAS = '[Config] Store_update_sofia_profile_alias',
  SWITCH_SOFIA_PROFILE_ALIAS = '[Config] Switch_sofia_profile_alias',
  STORE_SWITCH_SOFIA_PROFILE_ALIAS = '[Config] Store_switch_sofia_profile_alias',
  ADD_SOFIA_PROFILE_ALIAS = '[Config] Add_sofia_profile_alias',
  STORE_ADD_SOFIA_PROFILE_ALIAS = '[Config] Store_add_sofia_profile_alias',
  DEL_SOFIA_PROFILE_ALIAS = '[Config] Del_sofia_profile_alias',
  STORE_DEL_SOFIA_PROFILE_ALIAS = '[Config] Store_del_sofia_profile_alias',
  STORE_NEW_SOFIA_PROFILE_ALIAS = '[Config] Store_new_sofia_profile_alias',
  STORE_DROP_NEW_SOFIA_PROFILE_ALIAS = '[Config] Store_drop_new_sofia_profile_alias',

  ADD_SOFIA_PROFILE = '[Config] Add_sofia_profile',
  STORE_ADD_SOFIA_PROFILE = '[Config] Store_add_sofia_profile',
  ADD_SOFIA_PROFILE_GATEWAY = '[Config] Add_sofia_profile_gateway',
  STORE_ADD_SOFIA_PROFILE_GATEWAY = '[Config] Store_add_sofia_profile_gateway',

  DEL_SOFIA_PROFILE = '[Config] Del_sofia_profile',
  STORE_DEL_SOFIA_PROFILE = '[Config] Store_del_sofia_profile',
  RENAME_SOFIA_PROFILE = '[Config] Rename_sofia_profile',
  STORE_RENAME_SOFIA_PROFILE = '[Config] Store_rename_sofia_profile',
  DEL_SOFIA_PROFILE_GATEWAY = '[Config] Del_sofia_profile_gateway',
  STORE_DEL_SOFIA_PROFILE_GATEWAY = '[Config] Store_del_sofia_profile_gateway',
  RENAME_SOFIA_PROFILE_GATEWAY = '[Config] Rename_sofia_profile_gateway',
  STORE_RENAME_SOFIA_PROFILE_GATEWAY = '[Config] Store_rename_sofia_profile_gateway',
  StoreGotSofiaError = 'StoreGotSofiaError',
  SOFIA_PROFILE_COMMAND = '[API] Sofia profile command',
  STORE_SOFIA_PROFILE_COMMAND = '[API][Store] Sofia profile command',
  SWITCH_SOFIA_PROFILE = '[Config] Switch_sofia_profile',
  STORE_SWITCH_SOFIA_PROFILE = '[Config] Store_switch_sofia_profile',
}

export class GetSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.GET_SOFIA_GLOBAL_SETTINGS;
  constructor(public payload: any) {}
}

export class StoreGetSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_GET_SOFIA_GLOBAL_SETTINGS;
  constructor(public payload: any) {}
}

export class UpdateSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.UPDATE_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class StoreUpdateSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class SwitchSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class AddSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class StoreAddSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class DelSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class StoreDelSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class StoreNewSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class StoreDropNewSofiaGlobalSettings implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_SOFIA_GLOBAL_SETTING;
  constructor(public payload: any) {}
}

export class GetSofiaProfiles implements Action {
  readonly type = ConfigActionTypes.GET_SOFIA_PROFILES;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfiles implements Action {
  readonly type = ConfigActionTypes.STORE_GET_SOFIA_PROFILES;
  constructor(public payload: any) {}
}

export class GetSofiaProfilesParams implements Action {
  readonly type = ConfigActionTypes.GET_SOFIA_PROFILES_PARAMS;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfilesParams implements Action {
  readonly type = ConfigActionTypes.STORE_GET_SOFIA_PROFILES_PARAMS;
  constructor(public payload: any) {}
}

export class UpdateSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.UPDATE_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreUpdateSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class SwitchSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class AddSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class DelSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreNewSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreDropNewSofiaProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StorePasteSofiaProfileParams implements Action {
  readonly type = ConfigActionTypes.STORE_PASTE_SOFIA_PROFILE_PARAMS;
  constructor(public payload: any) {}
}

export class GetSofiaProfileGateways implements Action {
  readonly type = ConfigActionTypes.GET_SOFIA_PROFILE_GATEWAYS;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfileGateways implements Action {
  readonly type = ConfigActionTypes.STORE_GET_SOFIA_PROFILE_GATEWAYS;
  constructor(public payload: any) {}
}

export class StorePasteSofiaProfileGatewayParams implements Action {
  readonly type = ConfigActionTypes.STORE_PASTE_SOFIA_GATEWAY_PARAMS;
  constructor(public payload: any) {}
}

export class GetSofiaProfileGatewayParameters implements Action {
  readonly type = ConfigActionTypes.GetSofiaProfileGatewayParameters;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfileGatewayParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetSofiaProfileGatewayParameters;
  constructor(public payload: any) {}
}

export class GetSofiaProfileGatewayVariables implements Action {
  readonly type = ConfigActionTypes.GetSofiaProfileGatewayVariables;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfileGatewayVariables implements Action {
  readonly type = ConfigActionTypes.StoreGetSofiaProfileGatewayVariables;
  constructor(public payload: any) {}
}

export class StorePasteSofiaProfileGatewayVars implements Action {
  readonly type = ConfigActionTypes.STORE_PASTE_SOFIA_GATEWAY_VARS;
  constructor(public payload: any) {}
}

export class UpdateSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.UPDATE_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class StoreUpdateSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class SwitchSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class AddSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class DelSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class StoreNewSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class StoreDropNewSofiaProfileGatewayParam implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_GATEWAY_PARAM;
  constructor(public payload: any) {}
}

export class UpdateSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.UPDATE_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class StoreUpdateSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class SwitchSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class AddSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class DelSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class StoreNewSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class StoreDropNewSofiaProfileGatewayVar implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_GATEWAY_VAR;
  constructor(public payload: any) {}
}

export class GetSofiaProfileDomains implements Action {
  readonly type = ConfigActionTypes.GET_SOFIA_PROFILE_DOMAINS;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfileDomains implements Action {
  readonly type = ConfigActionTypes.STORE_GET_SOFIA_PROFILE_DOMAINS;
  constructor(public payload: any) {}
}

export class UpdateSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.UPDATE_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class StoreUpdateSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class SwitchSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class AddSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class DelSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class StoreNewSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class StoreDropNewSofiaProfileDomain implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_DOMAIN;
  constructor(public payload: any) {}
}

export class GetSofiaProfileAliases implements Action {
  readonly type = ConfigActionTypes.GET_SOFIA_PROFILE_ALIASES;
  constructor(public payload: any) {}
}

export class StoreGetSofiaProfileAliases implements Action {
  readonly type = ConfigActionTypes.STORE_GET_SOFIA_PROFILE_ALIASES;
  constructor(public payload: any) {}
}

export class UpdateSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.UPDATE_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class StoreUpdateSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class SwitchSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class AddSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class DelSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class StoreNewSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class StoreDropNewSofiaProfileAlias implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_ALIAS;
  constructor(public payload: any) {}
}

export class AddSofiaProfile implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfile implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class AddSofiaProfileGateway implements Action {
  readonly type = ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY;
  constructor(public payload: any) {}
}

export class StoreAddSofiaProfileGateway implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_GATEWAY;
  constructor(public payload: any) {}
}

export class DelSofiaProfile implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfile implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class RenameSofiaProfile implements Action {
  readonly type = ConfigActionTypes.RENAME_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class StoreRenameSofiaProfile implements Action {
  readonly type = ConfigActionTypes.STORE_RENAME_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class DelSofiaProfileGateway implements Action {
  readonly type = ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY;
  constructor(public payload: any) {}
}

export class StoreDelSofiaProfileGateway implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_GATEWAY;
  constructor(public payload: any) {}
}

export class RenameSofiaProfileGateway implements Action {
  readonly type = ConfigActionTypes.RENAME_SOFIA_PROFILE_GATEWAY;
  constructor(public payload: any) {}
}

export class StoreRenameSofiaProfileGateway implements Action {
  readonly type = ConfigActionTypes.STORE_RENAME_SOFIA_PROFILE_GATEWAY;
  constructor(public payload: any) {}
}

export class SofiaProfileCommand implements Action {
  readonly type = ConfigActionTypes.SOFIA_PROFILE_COMMAND;
  constructor(public payload: any) {}
}

export class StoreSofiaProfileCommand implements Action {
  readonly type = ConfigActionTypes.STORE_SOFIA_PROFILE_COMMAND;
  constructor(public payload: any) {}
}

export class SwitchSofiaProfile implements Action {
  readonly type = ConfigActionTypes.SWITCH_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class StoreSwitchSofiaProfile implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE;
  constructor(public payload: any) {}
}

export class StoreGotSofiaError implements Action {
  readonly type = ConfigActionTypes.StoreGotSofiaError;
  constructor(public payload: any) {}
}

export type All =
  | GetSofiaGlobalSettings
  | StoreGetSofiaGlobalSettings
  | GetSofiaProfiles
  | StoreGetSofiaProfiles
  | GetSofiaProfilesParams
  | StoreGetSofiaProfilesParams
  | UpdateSofiaGlobalSettings
  | StoreUpdateSofiaGlobalSettings
  | SwitchSofiaGlobalSettings
  | StoreSwitchSofiaGlobalSettings
  | StoreNewSofiaGlobalSettings
  | StoreDropNewSofiaGlobalSettings
  | AddSofiaGlobalSettings
  | StoreAddSofiaGlobalSettings
  | DelSofiaGlobalSettings
  | StoreDelSofiaGlobalSettings
  | UpdateSofiaProfileParam
  | StoreUpdateSofiaProfileParam
  | SwitchSofiaProfileParam
  | StoreSwitchSofiaProfileParam
  | AddSofiaProfileParam
  | StoreAddSofiaProfileParam
  | DelSofiaProfileParam
  | StoreDelSofiaProfileParam
  | StoreNewSofiaProfileParam
  | StoreDropNewSofiaProfileParam
  | StorePasteSofiaProfileParams
  | GetSofiaProfileGateways
  | StoreGetSofiaProfileGateways
  | StorePasteSofiaProfileGatewayParams
  | StorePasteSofiaProfileGatewayVars
  | UpdateSofiaProfileGatewayParam
  | StoreUpdateSofiaProfileGatewayParam
  | SwitchSofiaProfileGatewayParam
  | StoreSwitchSofiaProfileGatewayParam
  | AddSofiaProfileGatewayParam
  | StoreAddSofiaProfileGatewayParam
  | DelSofiaProfileGatewayParam
  | StoreDelSofiaProfileGatewayParam
  | StoreNewSofiaProfileGatewayParam
  | StoreDropNewSofiaProfileGatewayParam
  | UpdateSofiaProfileGatewayVar
  | StoreUpdateSofiaProfileGatewayVar
  | SwitchSofiaProfileGatewayVar
  | StoreSwitchSofiaProfileGatewayVar
  | AddSofiaProfileGatewayVar
  | StoreAddSofiaProfileGatewayVar
  | DelSofiaProfileGatewayVar
  | StoreDelSofiaProfileGatewayVar
  | StoreNewSofiaProfileGatewayVar
  | StoreDropNewSofiaProfileGatewayVar
  | GetSofiaProfileDomains
  | StoreGetSofiaProfileDomains
  | UpdateSofiaProfileDomain
  | StoreUpdateSofiaProfileDomain
  | SwitchSofiaProfileDomain
  | StoreSwitchSofiaProfileDomain
  | AddSofiaProfileDomain
  | StoreAddSofiaProfileDomain
  | DelSofiaProfileDomain
  | StoreDelSofiaProfileDomain
  | StoreNewSofiaProfileDomain
  | StoreDropNewSofiaProfileDomain
  | GetSofiaProfileAliases
  | StoreGetSofiaProfileAliases
  | UpdateSofiaProfileAlias
  | StoreUpdateSofiaProfileAlias
  | SwitchSofiaProfileAlias
  | StoreSwitchSofiaProfileAlias
  | AddSofiaProfileAlias
  | StoreAddSofiaProfileAlias
  | DelSofiaProfileAlias
  | StoreDelSofiaProfileAlias
  | StoreNewSofiaProfileAlias
  | StoreDropNewSofiaProfileAlias
  | AddSofiaProfile
  | StoreAddSofiaProfile
  | AddSofiaProfileGateway
  | StoreAddSofiaProfileGateway
  | DelSofiaProfileGateway
  | StoreDelSofiaProfileGateway
  | RenameSofiaProfileGateway
  | StoreRenameSofiaProfileGateway
  | DelSofiaProfile
  | StoreDelSofiaProfile
  | RenameSofiaProfile
  | StoreRenameSofiaProfile
  | SofiaProfileCommand
  | StoreSofiaProfileCommand
  | SwitchSofiaProfile
  | StoreSwitchSofiaProfile
  | StoreGotSofiaError
  | GetSofiaProfileGatewayParameters
  | StoreGetSofiaProfileGatewayParameters
  | GetSofiaProfileGatewayVariables
  | StoreGetSofiaProfileGatewayVariables
;

