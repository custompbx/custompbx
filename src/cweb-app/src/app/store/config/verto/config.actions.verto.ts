import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GET_VERTO_CONFIG = '[Config][Verto][Get]',
  STORE_GET_VERTO_CONFIG = '[Config][Verto][Store][Get]',
  GET_VERTO_PROFILES_PARAMS = '[Config][Verto][Profile][Parameters][Get]',
  STORE_GET_VERTO_PROFILES_PARAMS = 'Config][Verto][Profile][Parameters][Store][Get]',
  UPDATE_VERTO_SETTING = '[Config][Verto][Settings][Update]',
  STORE_UPDATE_VERTO_SETTING = '[Config][Verto][Settings][Store][Update]',
  SWITCH_VERTO_SETTING = '[Config][Verto][Setting][Switch]',
  STORE_SWITCH_VERTO_SETTING = '[Config][Verto][Setting][Store][Switch]',
  ADD_VERTO_SETTING = '[Config][Verto][Setting][Add]',
  STORE_ADD_VERTO_SETTING = '[Config][Verto][Setting][Store][Add]',
  DEL_VERTO_SETTING = '[Config][Verto][Setting][Del]',
  STORE_DEL_VERTO_SETTING = '[Config][Verto][Setting][Store][Del]',
  STORE_NEW_VERTO_SETTING = '[Config][Verto][Setting][New]',
  STORE_DROP_NEW_VERTO_SETTING = '[Config][Verto][Setting][Store][New]',
  UPDATE_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Update]',
  STORE_UPDATE_VERTO_PROFILE_PARAM = '[Config][Verto][Store][Profile][Store][Update]',
  SWITCH_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Switch]',
  STORE_SWITCH_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Store][Switch]',
  ADD_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Add]',
  STORE_ADD_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Store][Add]',
  DEL_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Del]',
  STORE_DEL_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Store][Del]',
  STORE_NEW_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Store][New]',
  STORE_DROP_NEW_VERTO_PROFILE_PARAM = '[Config][Verto][Profile][Param][Store][Drop New]',
  STORE_PASTE_VERTO_PROFILE_PARAMS = '[Config][Verto][Profile][Param][Store][Paste]',
  ADD_VERTO_PROFILE = '[Config][Verto][Profile][Add]',
  STORE_ADD_VERTO_PROFILE = '[Config][Verto][Profile][Store][Add]',
  DEL_VERTO_PROFILE = '[Config][Verto][Profile][Del]',
  STORE_DEL_VERTO_PROFILE = '[Config][Verto][Profile][Store][Del]',
  RENAME_VERTO_PROFILE = '[Config][Verto][Profile][Rename]',
  STORE_RENAME_VERTO_PROFILE = '[Config][Verto][Profile][Store][Rename]',
  StoreGotVertoError = 'StoreGotVertoError',
  MoveVertoProfileParameter = 'MoveVertoProfileParameter',
  StoreMoveVertoProfileParameter = 'StoreMoveVertoProfileParameter',
}

export class StoreGotVertoError implements Action {
  readonly type = ConfigActionTypes.StoreGotVertoError;
  constructor(public payload: any) {}
}

export class GetVertoConfig implements Action {
  readonly type = ConfigActionTypes.GET_VERTO_CONFIG;
  constructor(public payload: any) {}
}

export class StoreGetVertoConfig implements Action {
  readonly type = ConfigActionTypes.STORE_GET_VERTO_CONFIG;
  constructor(public payload: any) {}
}

export class GetVertoProfileParams implements Action {
  readonly type = ConfigActionTypes.GET_VERTO_PROFILES_PARAMS;
  constructor(public payload: any) {}
}

export class StoreGetVertoProfileParams implements Action {
  readonly type = ConfigActionTypes.STORE_GET_VERTO_PROFILES_PARAMS;
  constructor(public payload: any) {}
}

export class UpdateVertoSetting implements Action {
  readonly type = ConfigActionTypes.UPDATE_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class StoreUpdateVertoSetting implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class SwitchVertoSetting implements Action {
  readonly type = ConfigActionTypes.SWITCH_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class StoreSwitchVertoSetting implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class AddVertoSetting implements Action {
  readonly type = ConfigActionTypes.ADD_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class StoreAddVertoSetting implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class DelVertoSetting implements Action {
  readonly type = ConfigActionTypes.DEL_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class StoreDelVertoSetting implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class StoreNewVertoSetting implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class StoreDropNewVertoSetting implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_VERTO_SETTING;
  constructor(public payload: any) {}
}

export class UpdateVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.UPDATE_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreUpdateVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATE_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class SwitchVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.SWITCH_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreSwitchVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class AddVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.ADD_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreAddVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class DelVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.DEL_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreDelVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreNewVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StoreDropNewVertoProfileParam implements Action {
  readonly type = ConfigActionTypes.STORE_DROP_NEW_VERTO_PROFILE_PARAM;
  constructor(public payload: any) {}
}

export class StorePasteVertoProfileParams implements Action {
  readonly type = ConfigActionTypes.STORE_PASTE_VERTO_PROFILE_PARAMS;
  constructor(public payload: any) {}
}

export class AddVertoProfile implements Action {
  readonly type = ConfigActionTypes.ADD_VERTO_PROFILE;
  constructor(public payload: any) {}
}

export class StoreAddVertoProfile implements Action {
  readonly type = ConfigActionTypes.STORE_ADD_VERTO_PROFILE;
  constructor(public payload: any) {}
}

export class DelVertoProfile implements Action {
  readonly type = ConfigActionTypes.DEL_VERTO_PROFILE;
  constructor(public payload: any) {}
}

export class StoreDelVertoProfile implements Action {
  readonly type = ConfigActionTypes.STORE_DEL_VERTO_PROFILE;
  constructor(public payload: any) {}
}

export class RenameVertoProfile implements Action {
  readonly type = ConfigActionTypes.RENAME_VERTO_PROFILE;
  constructor(public payload: any) {}
}

export class StoreRenameVertoProfile implements Action {
  readonly type = ConfigActionTypes.STORE_RENAME_VERTO_PROFILE;
  constructor(public payload: any) {}
}

export class MoveVertoProfileParameter implements Action {
  readonly type = ConfigActionTypes.MoveVertoProfileParameter;
  constructor(public payload: any) {}
}

export class StoreMoveVertoProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreMoveVertoProfileParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetVertoConfig
  | StoreGetVertoConfig
  | GetVertoProfileParams
  | StoreGetVertoProfileParams
  | UpdateVertoSetting
  | StoreUpdateVertoSetting
  | SwitchVertoSetting
  | StoreSwitchVertoSetting
  | AddVertoSetting
  | StoreAddVertoSetting
  | DelVertoSetting
  | StoreDelVertoSetting
  | StoreNewVertoSetting
  | StoreDropNewVertoSetting
  | UpdateVertoProfileParam
  | StoreUpdateVertoProfileParam
  | SwitchVertoProfileParam
  | StoreSwitchVertoProfileParam
  | AddVertoProfileParam
  | StoreAddVertoProfileParam
  | DelVertoProfileParam
  | StoreDelVertoProfileParam
  | StoreNewVertoProfileParam
  | StoreDropNewVertoProfileParam
  | StorePasteVertoProfileParams
  | AddVertoProfile
  | StoreAddVertoProfile
  | DelVertoProfile
  | StoreDelVertoProfile
  | RenameVertoProfile
  | StoreRenameVertoProfile
  | StoreGotVertoError
  | MoveVertoProfileParameter
  | StoreMoveVertoProfileParameter
;

