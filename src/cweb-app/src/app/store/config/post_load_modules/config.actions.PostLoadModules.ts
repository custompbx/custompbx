
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetPostLoadModules = 'GetPostLoadModules',
  StoreGetPostLoadModules = 'StoreGetPostLoadModules',
  UpdatePostLoadModule = 'UpdatePostLoadModule',
  StoreUpdatePostLoadModule = 'StoreUpdatePostLoadModule',
  SwitchPostLoadModule = 'SwitchPostLoadModule',
  StoreSwitchPostLoadModule = 'StoreSwitchPostLoadModule',
  AddPostLoadModule = 'AddPostLoadModule',
  StoreAddPostLoadModule = 'StoreAddPostLoadModule',
  DelPostLoadModule = 'DelPostLoadModule',
  StoreDelPostLoadModule = 'StoreDelPostLoadModule',
  StoreNewPostLoadModule = 'StoreNewPostLoadModule',
  StoreDropNewPostLoadModule = 'StoreDropNewPostLoadModule',

  StoreGotPostLoadModuleError = 'StoreGotPostLoadModuleError',
}

export class GetPostLoadModules implements Action {
  readonly type = ConfigActionTypes.GetPostLoadModules;
  constructor(public payload: any) {}
}

export class StoreGetPostLoadModules implements Action {
  readonly type = ConfigActionTypes.StoreGetPostLoadModules;
  constructor(public payload: any) {}
}

export class UpdatePostLoadModule implements Action {
  readonly type = ConfigActionTypes.UpdatePostLoadModule;
  constructor(public payload: any) {}
}

export class StoreUpdatePostLoadModule implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePostLoadModule;
  constructor(public payload: any) {}
}

export class SwitchPostLoadModule implements Action {
  readonly type = ConfigActionTypes.SwitchPostLoadModule;
  constructor(public payload: any) {}
}

export class StoreSwitchPostLoadModule implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPostLoadModule;
  constructor(public payload: any) {}
}

export class AddPostLoadModule implements Action {
  readonly type = ConfigActionTypes.AddPostLoadModule;
  constructor(public payload: any) {}
}

export class StoreAddPostLoadModule implements Action {
  readonly type = ConfigActionTypes.StoreAddPostLoadModule;
  constructor(public payload: any) {}
}

export class DelPostLoadModule implements Action {
  readonly type = ConfigActionTypes.DelPostLoadModule;
  constructor(public payload: any) {}
}

export class StoreDelPostLoadModule implements Action {
  readonly type = ConfigActionTypes.StoreDelPostLoadModule;
  constructor(public payload: any) {}
}

export class StoreGotPostLoadModuleError implements Action {
  readonly type = ConfigActionTypes.StoreGotPostLoadModuleError;
  constructor(public payload: any) {}
}

export class StoreNewPostLoadModule implements Action {
  readonly type = ConfigActionTypes.StoreNewPostLoadModule;
  constructor(public payload: any) {}
}

export class StoreDropNewPostLoadModule implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPostLoadModule;
  constructor(public payload: any) {}
}

export type All =
  | GetPostLoadModules
  | StoreGetPostLoadModules
  | UpdatePostLoadModule
  | StoreUpdatePostLoadModule
  | SwitchPostLoadModule
  | StoreSwitchPostLoadModule
  | AddPostLoadModule
  | StoreAddPostLoadModule
  | DelPostLoadModule
  | StoreDelPostLoadModule
  | StoreGotPostLoadModuleError
  | StoreNewPostLoadModule
  | StoreDropNewPostLoadModule
;

