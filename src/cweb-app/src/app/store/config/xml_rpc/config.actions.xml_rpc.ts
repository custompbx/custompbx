
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetXmlRpc = 'GetXmlRpc',
  StoreGetXmlRpc = 'StoreGetXmlRpc',
  UpdateXmlRpcParameter = 'UpdateXmlRpcParameter',
  StoreUpdateXmlRpcParameter = 'StoreUpdateXmlRpcParameter',
  SwitchXmlRpcParameter = 'SwitchXmlRpcParameter',
  StoreSwitchXmlRpcParameter = 'StoreSwitchXmlRpcParameter',
  AddXmlRpcParameter = 'AddXmlRpcParameter',
  StoreAddXmlRpcParameter = 'StoreAddXmlRpcParameter',
  DelXmlRpcParameter = 'DelXmlRpcParameter',
  StoreDelXmlRpcParameter = 'StoreDelXmlRpcParameter',
  StoreNewXmlRpcParameter = 'StoreNewXmlRpcParameter',
  StoreDropNewXmlRpcParameter = 'StoreDropNewXmlRpcParameter',
  StoreGotXmlRpcError = 'StoreGotXmlRpcError',
}

export class GetXmlRpc implements Action {
  readonly type = ConfigActionTypes.GetXmlRpc;
  constructor(public payload: any) {}
}

export class StoreGetXmlRpc implements Action {
  readonly type = ConfigActionTypes.StoreGetXmlRpc;
  constructor(public payload: any) {}
}

export class UpdateXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.UpdateXmlRpcParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateXmlRpcParameter;
  constructor(public payload: any) {}
}

export class SwitchXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.SwitchXmlRpcParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchXmlRpcParameter;
  constructor(public payload: any) {}
}

export class AddXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.AddXmlRpcParameter;
  constructor(public payload: any) {}
}

export class StoreAddXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddXmlRpcParameter;
  constructor(public payload: any) {}
}

export class DelXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.DelXmlRpcParameter;
  constructor(public payload: any) {}
}

export class StoreDelXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelXmlRpcParameter;
  constructor(public payload: any) {}
}

export class StoreGotXmlRpcError implements Action {
  readonly type = ConfigActionTypes.StoreGotXmlRpcError;
  constructor(public payload: any) {}
}

export class StoreDropNewXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewXmlRpcParameter;
  constructor(public payload: any) {}
}

export class StoreNewXmlRpcParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewXmlRpcParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetXmlRpc
  | StoreGetXmlRpc
  | UpdateXmlRpcParameter
  | StoreUpdateXmlRpcParameter
  | SwitchXmlRpcParameter
  | StoreSwitchXmlRpcParameter
  | AddXmlRpcParameter
  | StoreAddXmlRpcParameter
  | DelXmlRpcParameter
  | StoreDelXmlRpcParameter
  | StoreGotXmlRpcError
  | StoreDropNewXmlRpcParameter
  | StoreNewXmlRpcParameter
;

