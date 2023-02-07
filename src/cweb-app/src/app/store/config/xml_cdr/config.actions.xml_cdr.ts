
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetXmlCdr = 'GetXmlCdr',
  StoreGetXmlCdr = 'StoreGetXmlCdr',
  UpdateXmlCdrParameter = 'UpdateXmlCdrParameter',
  StoreUpdateXmlCdrParameter = 'StoreUpdateXmlCdrParameter',
  SwitchXmlCdrParameter = 'SwitchXmlCdrParameter',
  StoreSwitchXmlCdrParameter = 'StoreSwitchXmlCdrParameter',
  AddXmlCdrParameter = 'AddXmlCdrParameter',
  StoreAddXmlCdrParameter = 'StoreAddXmlCdrParameter',
  DelXmlCdrParameter = 'DelXmlCdrParameter',
  StoreDelXmlCdrParameter = 'StoreDelXmlCdrParameter',
  StoreNewXmlCdrParameter = 'StoreNewXmlCdrParameter',
  StoreDropNewXmlCdrParameter = 'StoreDropNewXmlCdrParameter',
  StoreGotXmlCdrError = 'StoreGotXmlCdrError',
}

export class GetXmlCdr implements Action {
  readonly type = ConfigActionTypes.GetXmlCdr;
  constructor(public payload: any) {}
}

export class StoreGetXmlCdr implements Action {
  readonly type = ConfigActionTypes.StoreGetXmlCdr;
  constructor(public payload: any) {}
}

export class UpdateXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.UpdateXmlCdrParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateXmlCdrParameter;
  constructor(public payload: any) {}
}

export class SwitchXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.SwitchXmlCdrParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchXmlCdrParameter;
  constructor(public payload: any) {}
}

export class AddXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.AddXmlCdrParameter;
  constructor(public payload: any) {}
}

export class StoreAddXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddXmlCdrParameter;
  constructor(public payload: any) {}
}

export class DelXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.DelXmlCdrParameter;
  constructor(public payload: any) {}
}

export class StoreDelXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelXmlCdrParameter;
  constructor(public payload: any) {}
}

export class StoreGotXmlCdrError implements Action {
  readonly type = ConfigActionTypes.StoreGotXmlCdrError;
  constructor(public payload: any) {}
}

export class StoreDropNewXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewXmlCdrParameter;
  constructor(public payload: any) {}
}

export class StoreNewXmlCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewXmlCdrParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetXmlCdr
  | StoreGetXmlCdr
  | UpdateXmlCdrParameter
  | StoreUpdateXmlCdrParameter
  | SwitchXmlCdrParameter
  | StoreSwitchXmlCdrParameter
  | AddXmlCdrParameter
  | StoreAddXmlCdrParameter
  | DelXmlCdrParameter
  | StoreDelXmlCdrParameter
  | StoreGotXmlCdrError
  | StoreDropNewXmlCdrParameter
  | StoreNewXmlCdrParameter
;

