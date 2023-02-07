
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetOpus = 'GetOpus',
  StoreGetOpus = 'StoreGetOpus',
  UpdateOpusParameter = 'UpdateOpusParameter',
  StoreUpdateOpusParameter = 'StoreUpdateOpusParameter',
  SwitchOpusParameter = 'SwitchOpusParameter',
  StoreSwitchOpusParameter = 'StoreSwitchOpusParameter',
  AddOpusParameter = 'AddOpusParameter',
  StoreAddOpusParameter = 'StoreAddOpusParameter',
  DelOpusParameter = 'DelOpusParameter',
  StoreDelOpusParameter = 'StoreDelOpusParameter',
  StoreNewOpusParameter = 'StoreNewOpusParameter',
  StoreDropNewOpusParameter = 'StoreDropNewOpusParameter',
  StoreGotOpusError = 'StoreGotOpusError',
}

export class GetOpus implements Action {
  readonly type = ConfigActionTypes.GetOpus;
  constructor(public payload: any) {}
}

export class StoreGetOpus implements Action {
  readonly type = ConfigActionTypes.StoreGetOpus;
  constructor(public payload: any) {}
}

export class UpdateOpusParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOpusParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOpusParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOpusParameter;
  constructor(public payload: any) {}
}

export class SwitchOpusParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOpusParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOpusParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOpusParameter;
  constructor(public payload: any) {}
}

export class AddOpusParameter implements Action {
  readonly type = ConfigActionTypes.AddOpusParameter;
  constructor(public payload: any) {}
}

export class StoreAddOpusParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOpusParameter;
  constructor(public payload: any) {}
}

export class DelOpusParameter implements Action {
  readonly type = ConfigActionTypes.DelOpusParameter;
  constructor(public payload: any) {}
}

export class StoreDelOpusParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelOpusParameter;
  constructor(public payload: any) {}
}

export class StoreGotOpusError implements Action {
  readonly type = ConfigActionTypes.StoreGotOpusError;
  constructor(public payload: any) {}
}

export class StoreDropNewOpusParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOpusParameter;
  constructor(public payload: any) {}
}

export class StoreNewOpusParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOpusParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetOpus
  | StoreGetOpus
  | UpdateOpusParameter
  | StoreUpdateOpusParameter
  | SwitchOpusParameter
  | StoreSwitchOpusParameter
  | AddOpusParameter
  | StoreAddOpusParameter
  | DelOpusParameter
  | StoreDelOpusParameter
  | StoreGotOpusError
  | StoreDropNewOpusParameter
  | StoreNewOpusParameter
;

