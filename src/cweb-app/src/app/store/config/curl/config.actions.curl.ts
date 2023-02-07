
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetCurl = 'GetCurl',
  StoreGetCurl = 'StoreGetCurl',
  UpdateCurlParameter = 'UpdateCurlParameter',
  StoreUpdateCurlParameter = 'StoreUpdateCurlParameter',
  SwitchCurlParameter = 'SwitchCurlParameter',
  StoreSwitchCurlParameter = 'StoreSwitchCurlParameter',
  AddCurlParameter = 'AddCurlParameter',
  StoreAddCurlParameter = 'StoreAddCurlParameter',
  DelCurlParameter = 'DelCurlParameter',
  StoreDelCurlParameter = 'StoreDelCurlParameter',
  StoreNewCurlParameter = 'StoreNewCurlParameter',
  StoreDropNewCurlParameter = 'StoreDropNewCurlParameter',
  StoreGotCurlError = 'StoreGotCurlError',
}

export class GetCurl implements Action {
  readonly type = ConfigActionTypes.GetCurl;
  constructor(public payload: any) {}
}

export class StoreGetCurl implements Action {
  readonly type = ConfigActionTypes.StoreGetCurl;
  constructor(public payload: any) {}
}

export class UpdateCurlParameter implements Action {
  readonly type = ConfigActionTypes.UpdateCurlParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateCurlParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCurlParameter;
  constructor(public payload: any) {}
}

export class SwitchCurlParameter implements Action {
  readonly type = ConfigActionTypes.SwitchCurlParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchCurlParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchCurlParameter;
  constructor(public payload: any) {}
}

export class AddCurlParameter implements Action {
  readonly type = ConfigActionTypes.AddCurlParameter;
  constructor(public payload: any) {}
}

export class StoreAddCurlParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddCurlParameter;
  constructor(public payload: any) {}
}

export class DelCurlParameter implements Action {
  readonly type = ConfigActionTypes.DelCurlParameter;
  constructor(public payload: any) {}
}

export class StoreDelCurlParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelCurlParameter;
  constructor(public payload: any) {}
}

export class StoreGotCurlError implements Action {
  readonly type = ConfigActionTypes.StoreGotCurlError;
  constructor(public payload: any) {}
}

export class StoreDropNewCurlParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewCurlParameter;
  constructor(public payload: any) {}
}

export class StoreNewCurlParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewCurlParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetCurl
  | StoreGetCurl
  | UpdateCurlParameter
  | StoreUpdateCurlParameter
  | SwitchCurlParameter
  | StoreSwitchCurlParameter
  | AddCurlParameter
  | StoreAddCurlParameter
  | DelCurlParameter
  | StoreDelCurlParameter
  | StoreGotCurlError
  | StoreDropNewCurlParameter
  | StoreNewCurlParameter
;

