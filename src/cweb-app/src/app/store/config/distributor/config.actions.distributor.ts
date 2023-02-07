import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotDistributorError = 'StoreGotDistributorError',

  GetDistributorConfig = 'GetDistributorConfig',
  StoreGetDistributorConfig = 'StoreGetDistributorConfig',
  AddDistributorList = 'AddDistributorList',
  StoreAddDistributorList = 'StoreAddDistributorList',
  UpdateDistributorList = 'UpdateDistributorList',
  StoreUpdateDistributorList = 'StoreUpdateDistributorList',
  DelDistributorList = 'DelDistributorList',
  StoreDelDistributorList = 'StoreDelDistributorList',

  GetDistributorNodes = 'GetDistributorNodes',
  StoreGetDistributorNodes = 'StoreGetDistributorNodes',
  AddDistributorNode = 'AddDistributorNode',
  StoreAddDistributorNode = 'StoreAddDistributorNode',
  UpdateDistributorNode = 'UpdateDistributorNode',
  StoreUpdateDistributorNode = 'StoreUpdateDistributorNode',
  DelDistributorNode = 'DelDistributorNode',
  StoreDelDistributorNode = 'StoreDelDistributorNode',
  StoreNewDistributorNode = 'StoreNewDistributorNode',
  StoreDelNewDistributorNode = 'StoreDelNewDistributorNode',
  SwitchDistributorNode = 'SwitchDistributorNode',
  StoreSwitchDistributorNode = 'StoreSwitchDistributorNode',
}

export class StoreGotDistributorError implements Action {
  readonly type = ConfigActionTypes.StoreGotDistributorError;
  constructor(public payload: any) {}
}

export class GetDistributorConfig implements Action {
  readonly type = ConfigActionTypes.GetDistributorConfig;
  constructor(public payload: any) {}
}

export class StoreGetDistributorConfig implements Action {
  readonly type = ConfigActionTypes.StoreGetDistributorConfig;
  constructor(public payload: any) {}
}

export class AddDistributorList implements Action {
  readonly type = ConfigActionTypes.AddDistributorList;
  constructor(public payload: any) {}
}

export class StoreAddDistributorList implements Action {
  readonly type = ConfigActionTypes.StoreAddDistributorList;
  constructor(public payload: any) {}
}

export class UpdateDistributorList implements Action {
  readonly type = ConfigActionTypes.UpdateDistributorList;
  constructor(public payload: any) {}
}

export class StoreUpdateDistributorList implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDistributorList;
  constructor(public payload: any) {}
}

export class DelDistributorList implements Action {
  readonly type = ConfigActionTypes.DelDistributorList;
  constructor(public payload: any) {}
}

export class StoreDelDistributorList implements Action {
  readonly type = ConfigActionTypes.StoreDelDistributorList;
  constructor(public payload: any) {}
}

export class GetDistributorNodes implements Action {
  readonly type = ConfigActionTypes.GetDistributorNodes;
  constructor(public payload: any) {}
}

export class StoreGetDistributorNodes implements Action {
  readonly type = ConfigActionTypes.StoreGetDistributorNodes;
  constructor(public payload: any) {}
}

export class AddDistributorNode implements Action {
  readonly type = ConfigActionTypes.AddDistributorNode;
  constructor(public payload: any) {}
}

export class StoreAddDistributorNode implements Action {
  readonly type = ConfigActionTypes.StoreAddDistributorNode;
  constructor(public payload: any) {}
}

export class UpdateDistributorNode implements Action {
  readonly type = ConfigActionTypes.UpdateDistributorNode;
  constructor(public payload: any) {}
}

export class StoreUpdateDistributorNode implements Action {
  readonly type = ConfigActionTypes.StoreUpdateDistributorNode;
  constructor(public payload: any) {}
}

export class DelDistributorNode implements Action {
  readonly type = ConfigActionTypes.DelDistributorNode;
  constructor(public payload: any) {}
}

export class StoreDelDistributorNode implements Action {
  readonly type = ConfigActionTypes.StoreDelDistributorNode;
  constructor(public payload: any) {}
}

export class StoreNewDistributorNode implements Action {
  readonly type = ConfigActionTypes.StoreNewDistributorNode;
  constructor(public payload: any) {}
}

export class StoreDelNewDistributorNode implements Action {
  readonly type = ConfigActionTypes.StoreDelNewDistributorNode;
  constructor(public payload: any) {}
}

export class SwitchDistributorNode implements Action {
  readonly type = ConfigActionTypes.SwitchDistributorNode;
  constructor(public payload: any) {}
}

export class StoreSwitchDistributorNode implements Action {
  readonly type = ConfigActionTypes.StoreSwitchDistributorNode;
  constructor(public payload: any) {}
}

export type All =
  | StoreGotDistributorError
  | GetDistributorConfig
  | StoreGetDistributorConfig
  | AddDistributorList
  | StoreAddDistributorList
  | UpdateDistributorList
  | StoreUpdateDistributorList
  | DelDistributorList
  | StoreDelDistributorList
  | GetDistributorNodes
  | StoreGetDistributorNodes
  | AddDistributorNode
  | StoreAddDistributorNode
  | UpdateDistributorNode
  | StoreUpdateDistributorNode
  | DelDistributorNode
  | StoreDelDistributorNode
  | StoreNewDistributorNode
  | StoreDelNewDistributorNode
  | SwitchDistributorNode
  | StoreSwitchDistributorNode
;
