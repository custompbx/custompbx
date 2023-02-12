import { Action } from '@ngrx/store';

export enum SettingsActionTypes {
  STORE_SETTINGS = '[Settings] Update',
  UPDATE_FAILURE = '[Settings] Failure',
  GET_SETTINGS = '[Settings] Get',
  SET_SETTINGS = '[Settings] Set',

  GET_WEB_USERS = '[Settings][Users] Get',
  STORE_GET_WEB_USERS = '[Settings][Store][Users] Get',
  ADD_WEB_USER = '[Settings][Users] Add',
  STORE_ADD_WEB_USER = '[Settings][Store][Users] Add',
  RENAME_WEB_USER = '[Settings][Users] Rename',
  STORE_RENAME_WEB_USER = '[Settings][Store][Users] Rename',
  DELETE_WEB_USER = '[Settings][Users] Delete',
  STORE_DELETE_WEB_USER = '[Settings][Store][Users] Delete',

  SWITCH_WEB_USER = '[Settings][Users][Switch] Web user',
  STORE_SWITCH_WEB_USER = '[Settings][Store][Users][Switch] Web user',
  UPDATE_WEB_USER_PASSWORD = '[Settings][Users][Update] Password',
  STORE_UPDATE_WEB_USER_PASSWORD = '[Settings][Store][Users][Update] Password',
  UPDATE_WEB_USER_LANG = '[Settings][Users][Update] Lang',
  STORE_UPDATE_WEB_USER_LANG = '[Settings][Store][Users][Update] Lang',
  UPDATE_WEB_USER_SIP_USER = '[Settings][Users][Update] Sip user',
  STORE_UPDATE_WEB_USER_SIP_USER = '[Settings][Store][Users][Update] Sip user',
  UPDATE_WEB_USER_WS = '[Settings][Users][Update] Ws',
  STORE_UPDATE_WEB_USER_WS = '[Settings][Store][Users][Update] Ws',
  UPDATE_WEB_USER_VERTO_WS = '[Settings][Users][Update] Verto Ws',
  STORE_UPDATE_WEB_USER_VERTO_WS = '[Settings][Store][Users][Update] Verto Ws',
  UPDATE_WEB_USER_WEBRTC_LIB = '[Settings][Users][Update] WebRTC Lib',
  STORE_UPDATE_WEB_USER_WEBRTC_LIB = '[Settings][Store][Users][Update]  WebRTC Lib',
  UPDATE_WEB_USER_STUN = '[Settings][Users][Update] Stun',
  STORE_UPDATE_WEB_USER_STUN = '[Settings][Store][Users][Update] Stun',
  UPDATE_WEB_USER_AVATAR = '[Settings][Users][Update] Avatar',
  STORE_UPDATE_WEB_USER_AVATAR = '[Settings][Store][Users][Update] Avatar',
  CLEAR_WEB_USER_AVATAR = '[Settings][Users][Clear] Avatar',
  STORE_CLEAR_WEB_USER_AVATAR = '[Settings][Store][Users][Clear] Avatar',
  UpdateWebUserGroup = 'UpdateWebUserGroup',
  StoreUpdateWebUserGroup = 'StoreUpdateWebUserGroup',

  GetWebSettings = 'GetWebSettings',
  StoreGetWebSettings = 'StoreGetWebSettings',
  SaveWebSettings = 'SaveWebSettings',
  StoreSaveWebSettings = 'StoreSaveWebSettings',
  StoreGotWebError = 'StoreGotWebError',

  GetUserTokens = 'GetUserTokens',
  StoreGetUserTokens = 'StoreGetUserTokens',
  RemoveUserToken = 'RemoveUserToken',
  StoreRemoveUserToken = 'StoreRemoveUserToken',
  AddUserToken = 'AddUserToken',
  StoreAddUserToken = 'StoreAddUserToken',

  GetWebDirectoryUsersTemplates = 'GetWebDirectoryUsersTemplates',
  StoreGetWebDirectoryUsersTemplates = 'StoreGetWebDirectoryUsersTemplates',
  AddWebDirectoryUsersTemplate = 'AddWebDirectoryUsersTemplate',
  StoreAddWebDirectoryUsersTemplate = 'StoreAddWebDirectoryUsersTemplate',
  DelWebDirectoryUsersTemplate = 'DelWebDirectoryUsersTemplate',
  StoreDelWebDirectoryUsersTemplate = 'StoreDelWebDirectoryUsersTemplate',
  SwitchWebDirectoryUsersTemplate = 'SwitchWebDirectoryUsersTemplate',
  StoreSwitchWebDirectoryUsersTemplate = 'StoreSwitchWebDirectoryUsersTemplate',
  UpdateWebDirectoryUsersTemplate = 'UpdateWebDirectoryUsersTemplate',
  StoreUpdateWebDirectoryUsersTemplate = 'StoreUpdateWebDirectoryUsersTemplate',

  GetWebDirectoryUsersTemplateParameters = 'GetWebDirectoryUsersTemplateParameters',
  StoreGetWebDirectoryUsersTemplateParameters = 'StoreGetWebDirectoryUsersTemplateParameters',
  AddWebDirectoryUsersTemplateParameter = 'AddWebDirectoryUsersTemplateParameter',
  StoreAddWebDirectoryUsersTemplateParameter = 'StoreAddWebDirectoryUsersTemplateParameter',
  StoreNewWebDirectoryUsersTemplateParameter = 'StoreNewWebDirectoryUsersTemplateParameter',
  StoreDelNewWebDirectoryUsersTemplateParameter = 'StoreDelNewWebDirectoryUsersTemplateParameter',
  DelWebDirectoryUsersTemplateParameter = 'DelWebDirectoryUsersTemplateParameter',
  StoreDelWebDirectoryUsersTemplateParameter = 'StoreDelWebDirectoryUsersTemplateParameter',
  SwitchWebDirectoryUsersTemplateParameter = 'SwitchWebDirectoryUsersTemplateParameter',
  StoreSwitchWebDirectoryUsersTemplateParameter = 'StoreSwitchWebDirectoryUsersTemplateParameter',
  UpdateWebDirectoryUsersTemplateParameter = 'UpdateWebDirectoryUsersTemplateParameter',
  StoreUpdateWebDirectoryUsersTemplateParameter = 'StoreUpdateWebDirectoryUsersTemplateParameter',

  GetWebDirectoryUsersTemplateVariables = 'GetWebDirectoryUsersTemplateVariables',
  StoreGetWebDirectoryUsersTemplateVariables = 'StoreGetWebDirectoryUsersTemplateVariables',
  AddWebDirectoryUsersTemplateVariable = 'AddWebDirectoryUsersTemplateVariable',
  StoreAddWebDirectoryUsersTemplateVariable = 'StoreAddWebDirectoryUsersTemplateVariable',
  StoreNewWebDirectoryUsersTemplateVariable = 'StoreNewWebDirectoryUsersTemplateVariable',
  StoreDelNewWebDirectoryUsersTemplateVariable = 'StoreDelNewWebDirectoryUsersTemplateVariable',
  DelWebDirectoryUsersTemplateVariable = 'DelWebDirectoryUsersTemplateVariable',
  StoreDelWebDirectoryUsersTemplateVariable = 'StoreDelWebDirectoryUsersTemplateVariable',
  SwitchWebDirectoryUsersTemplateVariable = 'SwitchWebDirectoryUsersTemplateVariable',
  StoreSwitchWebDirectoryUsersTemplateVariable = 'StoreSwitchWebDirectoryUsersTemplateVariable',
  UpdateWebDirectoryUsersTemplateVariable = 'UpdateWebDirectoryUsersTemplateVariable',
  StoreUpdateWebDirectoryUsersTemplateVariable = 'StoreUpdateWebDirectoryUsersTemplateVariable',
}

export class AddUserToken implements Action {
  readonly type = SettingsActionTypes.AddUserToken;
  constructor(public payload: any) {}
}

export class StoreAddUserToken implements Action {
  readonly type = SettingsActionTypes.StoreAddUserToken;
  constructor(public payload: any) {}
}

export class RemoveUserToken implements Action {
  readonly type = SettingsActionTypes.RemoveUserToken;
  constructor(public payload: any) {}
}

export class StoreRemoveUserToken implements Action {
  readonly type = SettingsActionTypes.StoreRemoveUserToken;
  constructor(public payload: any) {}
}

export class GetUserTokens implements Action {
  readonly type = SettingsActionTypes.GetUserTokens;
  constructor(public payload: any) {}
}

export class StoreGetUserTokens implements Action {
  readonly type = SettingsActionTypes.StoreGetUserTokens;
  constructor(public payload: any) {}
}

export class StoreGotWebError implements Action {
  readonly type = SettingsActionTypes.StoreGotWebError;
  constructor(public payload: any) {}
}

export class GetWebSettings implements Action {
  readonly type = SettingsActionTypes.GetWebSettings;
  constructor(public payload: any) {}
}

export class StoreGetWebSettings implements Action {
  readonly type = SettingsActionTypes.StoreGetWebSettings;
  constructor(public payload: any) {}
}

export class SaveWebSettings implements Action {
  readonly type = SettingsActionTypes.SaveWebSettings;
  constructor(public payload: any) {}
}

export class StoreSaveWebSettings implements Action {
  readonly type = SettingsActionTypes.StoreSaveWebSettings;
  constructor(public payload: any) {}
}

export class UpdateSettings implements Action {
  readonly type = SettingsActionTypes.STORE_SETTINGS;
  constructor(public payload: any) {}
}

export class Failure implements Action {
  readonly type = SettingsActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}

export class GetSettings implements Action {
  readonly type = SettingsActionTypes.GET_SETTINGS;
  constructor(public payload: any) {}
}

export class SetSettings implements Action {
  readonly type = SettingsActionTypes.SET_SETTINGS;
  constructor(public payload: any) {}
}

export class GetWebUsers implements Action {
  readonly type = SettingsActionTypes.GET_WEB_USERS;
  constructor(public payload: any) {}
}

export class StoreGetWebUsers implements Action {
  readonly type = SettingsActionTypes.STORE_GET_WEB_USERS;
  constructor(public payload: any) {}
}

export class AddWebUser implements Action {
  readonly type = SettingsActionTypes.ADD_WEB_USER;
  constructor(public payload: any) {}
}

export class StoreAddWebUser implements Action {
  readonly type = SettingsActionTypes.STORE_ADD_WEB_USER;
  constructor(public payload: any) {}
}

export class RenameWebUser implements Action {
  readonly type = SettingsActionTypes.RENAME_WEB_USER;
  constructor(public payload: any) {}
}

export class StoreRenameWebUser implements Action {
  readonly type = SettingsActionTypes.STORE_RENAME_WEB_USER;
  constructor(public payload: any) {}
}

export class DeleteWebUser implements Action {
  readonly type = SettingsActionTypes.DELETE_WEB_USER;
  constructor(public payload: any) {}
}

export class StoreDeleteWebUser implements Action {
  readonly type = SettingsActionTypes.STORE_DELETE_WEB_USER;
  constructor(public payload: any) {}
}

export class SwitchWebUser implements Action {
  readonly type = SettingsActionTypes.SWITCH_WEB_USER;
  constructor(public payload: any) {}
}

export class StoreSwitchWebUser implements Action {
  readonly type = SettingsActionTypes.STORE_SWITCH_WEB_USER;
  constructor(public payload: any) {}
}

export class UpdateWebUserPassword implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_PASSWORD;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserPassword implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_PASSWORD;
  constructor(public payload: any) {}
}

export class UpdateWebUserLang implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_LANG;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserLang implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_LANG;
  constructor(public payload: any) {}
}

export class UpdateWebUserSipUser implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_SIP_USER;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserSipUser implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_SIP_USER;
  constructor(public payload: any) {}
}

export class UpdateWebUserWs implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_WS;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserWs implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_WS;
  constructor(public payload: any) {}
}

export class UpdateWebUserVertoWs implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_VERTO_WS;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserVertoWs implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_VERTO_WS;
  constructor(public payload: any) {}
}

export class UpdateWebUserWebRTCLib implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_WEBRTC_LIB;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserWebRTCLib implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_WEBRTC_LIB;
  constructor(public payload: any) {}
}

export class UpdateWebUserStun implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_STUN;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserStun implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_STUN;
  constructor(public payload: any) {}
}

export class UpdateWebUserAvatar implements Action {
  readonly type = SettingsActionTypes.UPDATE_WEB_USER_AVATAR;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserAvatar implements Action {
  readonly type = SettingsActionTypes.STORE_UPDATE_WEB_USER_AVATAR;
  constructor(public payload: any) {}
}

export class ClearWebUserAvatar implements Action {
  readonly type = SettingsActionTypes.CLEAR_WEB_USER_AVATAR;
  constructor(public payload: any) {}
}

export class StoreClearWebUserAvatar implements Action {
  readonly type = SettingsActionTypes.STORE_CLEAR_WEB_USER_AVATAR;
  constructor(public payload: any) {}
}

export class UpdateWebUserGroup implements Action {
  readonly type = SettingsActionTypes.UpdateWebUserGroup;
  constructor(public payload: any) {}
}

export class StoreUpdateWebUserGroup implements Action {
  readonly type = SettingsActionTypes.StoreUpdateWebUserGroup;
  constructor(public payload: any) {}
}

export class GetWebDirectoryUsersTemplates implements Action {
  readonly type = SettingsActionTypes.GetWebDirectoryUsersTemplates;
  constructor(public payload: any) {}
}

export class StoreGetWebDirectoryUsersTemplates implements Action {
  readonly type = SettingsActionTypes.StoreGetWebDirectoryUsersTemplates;
  constructor(public payload: any) {}
}

export class AddWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.AddWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class StoreAddWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.StoreAddWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class DelWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.DelWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class StoreDelWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.StoreDelWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class SwitchWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.SwitchWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class StoreSwitchWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.StoreSwitchWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class UpdateWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.UpdateWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class StoreUpdateWebDirectoryUsersTemplate implements Action {
  readonly type = SettingsActionTypes.StoreUpdateWebDirectoryUsersTemplate;
  constructor(public payload: any) {}
}

export class GetWebDirectoryUsersTemplateParameters implements Action {
  readonly type = SettingsActionTypes.GetWebDirectoryUsersTemplateParameters;
  constructor(public payload: any) {}
}

export class StoreGetWebDirectoryUsersTemplateParameters implements Action {
  readonly type = SettingsActionTypes.StoreGetWebDirectoryUsersTemplateParameters;
  constructor(public payload: any) {}
}

export class AddWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.AddWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class StoreAddWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.StoreAddWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class DelWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.DelWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class StoreDelWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.StoreDelWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class SwitchWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.SwitchWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.StoreSwitchWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class UpdateWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.UpdateWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.StoreUpdateWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class GetWebDirectoryUsersTemplateVariables implements Action {
  readonly type = SettingsActionTypes.GetWebDirectoryUsersTemplateVariables;
  constructor(public payload: any) {}
}

export class StoreGetWebDirectoryUsersTemplateVariables implements Action {
  readonly type = SettingsActionTypes.StoreGetWebDirectoryUsersTemplateVariables;
  constructor(public payload: any) {}
}

export class AddWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.AddWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class StoreAddWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.StoreAddWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class DelWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.DelWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class StoreDelWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.StoreDelWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class SwitchWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.SwitchWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class StoreSwitchWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.StoreSwitchWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class UpdateWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.UpdateWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class StoreUpdateWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.StoreUpdateWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class StoreNewWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.StoreNewWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class StoreNewWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.StoreNewWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export class StoreDelNewWebDirectoryUsersTemplateParameter implements Action {
  readonly type = SettingsActionTypes.StoreDelNewWebDirectoryUsersTemplateParameter;
  constructor(public payload: any) {}
}

export class StoreDelNewWebDirectoryUsersTemplateVariable implements Action {
  readonly type = SettingsActionTypes.StoreDelNewWebDirectoryUsersTemplateVariable;
  constructor(public payload: any) {}
}

export type All =
  | UpdateSettings
  | Failure
  | GetSettings
  | SetSettings
  | GetWebUsers
  | StoreGetWebUsers
  | AddWebUser
  | StoreAddWebUser
  | RenameWebUser
  | StoreRenameWebUser
  | DeleteWebUser
  | StoreDeleteWebUser
  | SwitchWebUser
  | StoreSwitchWebUser
  | UpdateWebUserPassword
  | StoreUpdateWebUserPassword
  | UpdateWebUserLang
  | StoreUpdateWebUserLang
  | UpdateWebUserSipUser
  | StoreUpdateWebUserSipUser
  | UpdateWebUserWs
  | StoreUpdateWebUserWs
  | UpdateWebUserVertoWs
  | StoreUpdateWebUserVertoWs
  | UpdateWebUserWebRTCLib
  | StoreUpdateWebUserWebRTCLib
  | UpdateWebUserStun
  | StoreUpdateWebUserStun
  | UpdateWebUserAvatar
  | StoreUpdateWebUserAvatar
  | ClearWebUserAvatar
  | StoreClearWebUserAvatar
  | GetWebSettings
  | StoreGetWebSettings
  | SaveWebSettings
  | StoreSaveWebSettings
  | StoreGotWebError
  | GetUserTokens
  | StoreGetUserTokens
  | RemoveUserToken
  | StoreRemoveUserToken
  | AddUserToken
  | StoreAddUserToken
  | UpdateWebUserGroup
  | StoreUpdateWebUserGroup
  | GetWebDirectoryUsersTemplates
  | StoreGetWebDirectoryUsersTemplates
  | AddWebDirectoryUsersTemplate
  | StoreAddWebDirectoryUsersTemplate
  | DelWebDirectoryUsersTemplate
  | StoreDelWebDirectoryUsersTemplate
  | SwitchWebDirectoryUsersTemplate
  | StoreSwitchWebDirectoryUsersTemplate
  | UpdateWebDirectoryUsersTemplate
  | StoreUpdateWebDirectoryUsersTemplate
  | GetWebDirectoryUsersTemplateParameters
  | StoreGetWebDirectoryUsersTemplateParameters
  | AddWebDirectoryUsersTemplateParameter
  | StoreAddWebDirectoryUsersTemplateParameter
  | DelWebDirectoryUsersTemplateParameter
  | StoreDelWebDirectoryUsersTemplateParameter
  | SwitchWebDirectoryUsersTemplateParameter
  | StoreSwitchWebDirectoryUsersTemplateParameter
  | UpdateWebDirectoryUsersTemplateParameter
  | StoreUpdateWebDirectoryUsersTemplateParameter
  | GetWebDirectoryUsersTemplateVariables
  | StoreGetWebDirectoryUsersTemplateVariables
  | AddWebDirectoryUsersTemplateVariable
  | StoreAddWebDirectoryUsersTemplateVariable
  | DelWebDirectoryUsersTemplateVariable
  | StoreDelWebDirectoryUsersTemplateVariable
  | SwitchWebDirectoryUsersTemplateVariable
  | StoreSwitchWebDirectoryUsersTemplateVariable
  | UpdateWebDirectoryUsersTemplateVariable
  | StoreUpdateWebDirectoryUsersTemplateVariable
  | StoreNewWebDirectoryUsersTemplateParameter
  | StoreNewWebDirectoryUsersTemplateVariable
  | StoreDelNewWebDirectoryUsersTemplateParameter
  | StoreDelNewWebDirectoryUsersTemplateVariable
  ;
