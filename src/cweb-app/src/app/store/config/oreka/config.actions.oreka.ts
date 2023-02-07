
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetOreka = 'GetOreka',
  StoreGetOreka = 'StoreGetOreka',
  UpdateOrekaParameter = 'UpdateOrekaParameter',
  StoreUpdateOrekaParameter = 'StoreUpdateOrekaParameter',
  SwitchOrekaParameter = 'SwitchOrekaParameter',
  StoreSwitchOrekaParameter = 'StoreSwitchOrekaParameter',
  AddOrekaParameter = 'AddOrekaParameter',
  StoreAddOrekaParameter = 'StoreAddOrekaParameter',
  DelOrekaParameter = 'DelOrekaParameter',
  StoreDelOrekaParameter = 'StoreDelOrekaParameter',
  StoreNewOrekaParameter = 'StoreNewOrekaParameter',
  StoreDropNewOrekaParameter = 'StoreDropNewOrekaParameter',
  StoreGotOrekaError = 'StoreGotOrekaError',
}

export class GetOreka implements Action {
  readonly type = ConfigActionTypes.GetOreka;
  constructor(public payload: any) {}
}

export class StoreGetOreka implements Action {
  readonly type = ConfigActionTypes.StoreGetOreka;
  constructor(public payload: any) {}
}

export class UpdateOrekaParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOrekaParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOrekaParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOrekaParameter;
  constructor(public payload: any) {}
}

export class SwitchOrekaParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOrekaParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOrekaParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOrekaParameter;
  constructor(public payload: any) {}
}

export class AddOrekaParameter implements Action {
  readonly type = ConfigActionTypes.AddOrekaParameter;
  constructor(public payload: any) {}
}

export class StoreAddOrekaParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOrekaParameter;
  constructor(public payload: any) {}
}

export class DelOrekaParameter implements Action {
  readonly type = ConfigActionTypes.DelOrekaParameter;
  constructor(public payload: any) {}
}

export class StoreDelOrekaParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelOrekaParameter;
  constructor(public payload: any) {}
}

export class StoreGotOrekaError implements Action {
  readonly type = ConfigActionTypes.StoreGotOrekaError;
  constructor(public payload: any) {}
}

export class StoreDropNewOrekaParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOrekaParameter;
  constructor(public payload: any) {}
}

export class StoreNewOrekaParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOrekaParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetOreka
  | StoreGetOreka
  | UpdateOrekaParameter
  | StoreUpdateOrekaParameter
  | SwitchOrekaParameter
  | StoreSwitchOrekaParameter
  | AddOrekaParameter
  | StoreAddOrekaParameter
  | DelOrekaParameter
  | StoreDelOrekaParameter
  | StoreGotOrekaError
  | StoreDropNewOrekaParameter
  | StoreNewOrekaParameter
;

