
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetTtsCommandline = 'GetTtsCommandline',
  StoreGetTtsCommandline = 'StoreGetTtsCommandline',
  UpdateTtsCommandlineParameter = 'UpdateTtsCommandlineParameter',
  StoreUpdateTtsCommandlineParameter = 'StoreUpdateTtsCommandlineParameter',
  SwitchTtsCommandlineParameter = 'SwitchTtsCommandlineParameter',
  StoreSwitchTtsCommandlineParameter = 'StoreSwitchTtsCommandlineParameter',
  AddTtsCommandlineParameter = 'AddTtsCommandlineParameter',
  StoreAddTtsCommandlineParameter = 'StoreAddTtsCommandlineParameter',
  DelTtsCommandlineParameter = 'DelTtsCommandlineParameter',
  StoreDelTtsCommandlineParameter = 'StoreDelTtsCommandlineParameter',
  StoreNewTtsCommandlineParameter = 'StoreNewTtsCommandlineParameter',
  StoreDropNewTtsCommandlineParameter = 'StoreDropNewTtsCommandlineParameter',
  StoreGotTtsCommandlineError = 'StoreGotTtsCommandlineError',
}

export class GetTtsCommandline implements Action {
  readonly type = ConfigActionTypes.GetTtsCommandline;
  constructor(public payload: any) {}
}

export class StoreGetTtsCommandline implements Action {
  readonly type = ConfigActionTypes.StoreGetTtsCommandline;
  constructor(public payload: any) {}
}

export class UpdateTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.UpdateTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class SwitchTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.SwitchTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class AddTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.AddTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class StoreAddTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class DelTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.DelTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class StoreDelTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class StoreGotTtsCommandlineError implements Action {
  readonly type = ConfigActionTypes.StoreGotTtsCommandlineError;
  constructor(public payload: any) {}
}

export class StoreDropNewTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export class StoreNewTtsCommandlineParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewTtsCommandlineParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetTtsCommandline
  | StoreGetTtsCommandline
  | UpdateTtsCommandlineParameter
  | StoreUpdateTtsCommandlineParameter
  | SwitchTtsCommandlineParameter
  | StoreSwitchTtsCommandlineParameter
  | AddTtsCommandlineParameter
  | StoreAddTtsCommandlineParameter
  | DelTtsCommandlineParameter
  | StoreDelTtsCommandlineParameter
  | StoreGotTtsCommandlineError
  | StoreDropNewTtsCommandlineParameter
  | StoreNewTtsCommandlineParameter
;

