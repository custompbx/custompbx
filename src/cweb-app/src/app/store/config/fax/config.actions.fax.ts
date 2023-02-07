
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetFax = 'GetFax',
  StoreGetFax = 'StoreGetFax',
  UpdateFaxParameter = 'UpdateFaxParameter',
  StoreUpdateFaxParameter = 'StoreUpdateFaxParameter',
  SwitchFaxParameter = 'SwitchFaxParameter',
  StoreSwitchFaxParameter = 'StoreSwitchFaxParameter',
  AddFaxParameter = 'AddFaxParameter',
  StoreAddFaxParameter = 'StoreAddFaxParameter',
  DelFaxParameter = 'DelFaxParameter',
  StoreDelFaxParameter = 'StoreDelFaxParameter',
  StoreNewFaxParameter = 'StoreNewFaxParameter',
  StoreDropNewFaxParameter = 'StoreDropNewFaxParameter',
  StoreGotFaxError = 'StoreGotFaxError',
}

export class GetFax implements Action {
  readonly type = ConfigActionTypes.GetFax;
  constructor(public payload: any) {}
}

export class StoreGetFax implements Action {
  readonly type = ConfigActionTypes.StoreGetFax;
  constructor(public payload: any) {}
}

export class UpdateFaxParameter implements Action {
  readonly type = ConfigActionTypes.UpdateFaxParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateFaxParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateFaxParameter;
  constructor(public payload: any) {}
}

export class SwitchFaxParameter implements Action {
  readonly type = ConfigActionTypes.SwitchFaxParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchFaxParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchFaxParameter;
  constructor(public payload: any) {}
}

export class AddFaxParameter implements Action {
  readonly type = ConfigActionTypes.AddFaxParameter;
  constructor(public payload: any) {}
}

export class StoreAddFaxParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddFaxParameter;
  constructor(public payload: any) {}
}

export class DelFaxParameter implements Action {
  readonly type = ConfigActionTypes.DelFaxParameter;
  constructor(public payload: any) {}
}

export class StoreDelFaxParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelFaxParameter;
  constructor(public payload: any) {}
}

export class StoreGotFaxError implements Action {
  readonly type = ConfigActionTypes.StoreGotFaxError;
  constructor(public payload: any) {}
}

export class StoreDropNewFaxParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewFaxParameter;
  constructor(public payload: any) {}
}

export class StoreNewFaxParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewFaxParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetFax
  | StoreGetFax
  | UpdateFaxParameter
  | StoreUpdateFaxParameter
  | SwitchFaxParameter
  | StoreSwitchFaxParameter
  | AddFaxParameter
  | StoreAddFaxParameter
  | DelFaxParameter
  | StoreDelFaxParameter
  | StoreGotFaxError
  | StoreDropNewFaxParameter
  | StoreNewFaxParameter
;

