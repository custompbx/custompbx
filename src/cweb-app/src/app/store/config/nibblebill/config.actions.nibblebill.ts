
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetNibblebill = 'GetNibblebill',
  StoreGetNibblebill = 'StoreGetNibblebill',
  UpdateNibblebillParameter = 'UpdateNibblebillParameter',
  StoreUpdateNibblebillParameter = 'StoreUpdateNibblebillParameter',
  SwitchNibblebillParameter = 'SwitchNibblebillParameter',
  StoreSwitchNibblebillParameter = 'StoreSwitchNibblebillParameter',
  AddNibblebillParameter = 'AddNibblebillParameter',
  StoreAddNibblebillParameter = 'StoreAddNibblebillParameter',
  DelNibblebillParameter = 'DelNibblebillParameter',
  StoreDelNibblebillParameter = 'StoreDelNibblebillParameter',
  StoreNewNibblebillParameter = 'StoreNewNibblebillParameter',
  StoreDropNewNibblebillParameter = 'StoreDropNewNibblebillParameter',
  StoreGotNibblebillError = 'StoreGotNibblebillError',
}

export class GetNibblebill implements Action {
  readonly type = ConfigActionTypes.GetNibblebill;
  constructor(public payload: any) {}
}

export class StoreGetNibblebill implements Action {
  readonly type = ConfigActionTypes.StoreGetNibblebill;
  constructor(public payload: any) {}
}

export class UpdateNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.UpdateNibblebillParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateNibblebillParameter;
  constructor(public payload: any) {}
}

export class SwitchNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.SwitchNibblebillParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchNibblebillParameter;
  constructor(public payload: any) {}
}

export class AddNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.AddNibblebillParameter;
  constructor(public payload: any) {}
}

export class StoreAddNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddNibblebillParameter;
  constructor(public payload: any) {}
}

export class DelNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.DelNibblebillParameter;
  constructor(public payload: any) {}
}

export class StoreDelNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelNibblebillParameter;
  constructor(public payload: any) {}
}

export class StoreGotNibblebillError implements Action {
  readonly type = ConfigActionTypes.StoreGotNibblebillError;
  constructor(public payload: any) {}
}

export class StoreDropNewNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewNibblebillParameter;
  constructor(public payload: any) {}
}

export class StoreNewNibblebillParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewNibblebillParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetNibblebill
  | StoreGetNibblebill
  | UpdateNibblebillParameter
  | StoreUpdateNibblebillParameter
  | SwitchNibblebillParameter
  | StoreSwitchNibblebillParameter
  | AddNibblebillParameter
  | StoreAddNibblebillParameter
  | DelNibblebillParameter
  | StoreDelNibblebillParameter
  | StoreGotNibblebillError
  | StoreDropNewNibblebillParameter
  | StoreNewNibblebillParameter
;

