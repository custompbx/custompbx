import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  UPDATE_FAILURE = '[Config] Failure',

  GET_MODULES = '[Config][Get] Modules',
  STORE_GET_MODULES = '[Config][Store][Get] Modules',
  RELOAD_MODULE = '[Config][Reload] Module',
  STORE_RELOAD_MODULE = '[Config][Store][Reload] Module',
  UNLOAD_MODULE = '[Config][Unload] Module',
  STORE_UNLOAD_MODULE = '[Config][Store][Unload] Module',
  LOAD_MODULE = '[Config][Load] Module',
  STORE_LOAD_MODULE = '[Config][Store][Load] Module',
  SWITCH_MODULE = '[Config][Switch] Module',
  STORE_SWITCH_MODULE = '[Config][Store][Switch] Module',
  IMPORT_MODULE = '[Config][Import] Module',
  STORE_IMPORT_MODULE = '[Config][Store][Import] Module',
  FROM_SCRATCH_MODULE = '[Config][From scratch] Module',
  STORE_FROM_SCRATCH_MODULE = '[Config][Store][From scratch] Module',
  AUTOLOAD_MODULE = '[Config][Autoload] Module',
  STORE_AUTOLOAD_MODULE = '[Config][Store][Autoload] Module',
  IMPORT_ALL_MODULES = '[Config][Import] All Modules',
  STORE_IMPORT_ALL_MODULES = '[Config][Store][Import] All Modules',
  TruncateModuleConfig = 'TruncateModuleConfig',
  StoreTruncateModuleConfig = 'StoreTruncateModuleConfig',
  ImportXMLModuleConfig = 'ImportXMLModuleConfig',
  StoreImportXMLModuleConfig = 'StoreImportXMLModuleConfig',

  StoreGotModuleError = 'StoreGotModuleError',
}

export class Failure implements Action {
  readonly type = ConfigActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}

export class StoreGotModuleError implements Action {
  readonly type = ConfigActionTypes.StoreGotModuleError;
  constructor(public payload: any) {}
}

export class ImportAllModules implements Action {
  readonly type = ConfigActionTypes.IMPORT_ALL_MODULES;
  constructor(public payload: any) {}
}

export class StoreImportAllModules implements Action {
  readonly type = ConfigActionTypes.STORE_IMPORT_ALL_MODULES;
  constructor(public payload: any) {}
}

export class TruncateModuleConfig implements Action {
  readonly type = ConfigActionTypes.TruncateModuleConfig;
  constructor(public payload: any) {}
}

export class StoreTruncateModuleConfig implements Action {
  readonly type = ConfigActionTypes.StoreTruncateModuleConfig;
  constructor(public payload: any) {}
}

export class ImportXMLModuleConfig implements Action {
  readonly type = ConfigActionTypes.ImportXMLModuleConfig;
  constructor(public payload: any) {}
}

export class StoreImportXMLModuleConfig implements Action {
  readonly type = ConfigActionTypes.StoreImportXMLModuleConfig;
  constructor(public payload: any) {}
}

export class GetModules implements Action {
  readonly type = ConfigActionTypes.GET_MODULES;
  constructor(public payload: any) {}
}

export class StoreGetModules implements Action {
  readonly type = ConfigActionTypes.STORE_GET_MODULES;
  constructor(public payload: any) {}
}

export class ReloadModule implements Action {
  readonly type = ConfigActionTypes.RELOAD_MODULE;
  constructor(public payload: any) {}
}

export class StoreReloadModule implements Action {
  readonly type = ConfigActionTypes.STORE_RELOAD_MODULE;
  constructor(public payload: any) {}
}

export class UnloadModule implements Action {
  readonly type = ConfigActionTypes.UNLOAD_MODULE;
  constructor(public payload: any) {}
}

export class StoreUnloadModule implements Action {
  readonly type = ConfigActionTypes.STORE_UNLOAD_MODULE;
  constructor(public payload: any) {}
}

export class LoadModule implements Action {
  readonly type = ConfigActionTypes.LOAD_MODULE;
  constructor(public payload: any) {}
}

export class StoreLoadModule implements Action {
  readonly type = ConfigActionTypes.STORE_LOAD_MODULE;
  constructor(public payload: any) {}
}

export class SwitchModule implements Action {
  readonly type = ConfigActionTypes.SWITCH_MODULE;
  constructor(public payload: any) {}
}

export class StoreSwitchModule implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_MODULE;
  constructor(public payload: any) {}
}

export class ImportConfModule implements Action {
  readonly type = ConfigActionTypes.IMPORT_MODULE;
  constructor(public payload: any) {}
}

export class StoreImportConfModule implements Action {
  readonly type = ConfigActionTypes.STORE_IMPORT_MODULE;
  constructor(public payload: any) {}
}

export class FromScratchConfModule implements Action {
  readonly type = ConfigActionTypes.FROM_SCRATCH_MODULE;
  constructor(public payload: any) {}
}

export class StoreFromScratchConfModule implements Action {
  readonly type = ConfigActionTypes.STORE_FROM_SCRATCH_MODULE;
  constructor(public payload: any) {}
}

export class AutoloadModule implements Action {
  readonly type = ConfigActionTypes.AUTOLOAD_MODULE;
  constructor(public payload: any) {}
}

export class StoreAutoloadModule implements Action {
  readonly type = ConfigActionTypes.STORE_AUTOLOAD_MODULE;
  constructor(public payload: any) {}
}

export type All =
  | Failure
  | GetModules
  | StoreGetModules
  | ReloadModule
  | StoreReloadModule
  | UnloadModule
  | StoreUnloadModule
  | LoadModule
  | StoreLoadModule
  | SwitchModule
  | StoreSwitchModule
  | ImportConfModule
  | StoreImportConfModule
  | FromScratchConfModule
  | StoreFromScratchConfModule
  | AutoloadModule
  | StoreAutoloadModule
  | ImportAllModules
  | StoreImportAllModules
  | StoreGotModuleError
  | TruncateModuleConfig
  | StoreTruncateModuleConfig
  | ImportXMLModuleConfig
  | StoreImportXMLModuleConfig
;
