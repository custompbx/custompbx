
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetCidlookup = 'GetCidlookup',
  StoreGetCidlookup = 'StoreGetCidlookup',
  UpdateCidlookupParameter = 'UpdateCidlookupParameter',
  StoreUpdateCidlookupParameter = 'StoreUpdateCidlookupParameter',
  SwitchCidlookupParameter = 'SwitchCidlookupParameter',
  StoreSwitchCidlookupParameter = 'StoreSwitchCidlookupParameter',
  AddCidlookupParameter = 'AddCidlookupParameter',
  StoreAddCidlookupParameter = 'StoreAddCidlookupParameter',
  DelCidlookupParameter = 'DelCidlookupParameter',
  StoreDelCidlookupParameter = 'StoreDelCidlookupParameter',
  StoreNewCidlookupParameter = 'StoreNewCidlookupParameter',
  StoreDropNewCidlookupParameter = 'StoreDropNewCidlookupParameter',
  StoreGotCidlookupError = 'StoreGotCidlookupError',
}

export class GetCidlookup implements Action {
  readonly type = ConfigActionTypes.GetCidlookup;
  constructor(public payload: any) {}
}

export class StoreGetCidlookup implements Action {
  readonly type = ConfigActionTypes.StoreGetCidlookup;
  constructor(public payload: any) {}
}

export class UpdateCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.UpdateCidlookupParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCidlookupParameter;
  constructor(public payload: any) {}
}

export class SwitchCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.SwitchCidlookupParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchCidlookupParameter;
  constructor(public payload: any) {}
}

export class AddCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.AddCidlookupParameter;
  constructor(public payload: any) {}
}

export class StoreAddCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddCidlookupParameter;
  constructor(public payload: any) {}
}

export class DelCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.DelCidlookupParameter;
  constructor(public payload: any) {}
}

export class StoreDelCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelCidlookupParameter;
  constructor(public payload: any) {}
}

export class StoreGotCidlookupError implements Action {
  readonly type = ConfigActionTypes.StoreGotCidlookupError;
  constructor(public payload: any) {}
}

export class StoreDropNewCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewCidlookupParameter;
  constructor(public payload: any) {}
}

export class StoreNewCidlookupParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewCidlookupParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetCidlookup
  | StoreGetCidlookup
  | UpdateCidlookupParameter
  | StoreUpdateCidlookupParameter
  | SwitchCidlookupParameter
  | StoreSwitchCidlookupParameter
  | AddCidlookupParameter
  | StoreAddCidlookupParameter
  | DelCidlookupParameter
  | StoreDelCidlookupParameter
  | StoreGotCidlookupError
  | StoreDropNewCidlookupParameter
  | StoreNewCidlookupParameter
;

