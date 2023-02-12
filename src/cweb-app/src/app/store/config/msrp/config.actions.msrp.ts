
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetMsrp = 'GetMsrp',
  StoreGetMsrp = 'StoreGetMsrp',
  UpdateMsrpParameter = 'UpdateMsrpParameter',
  StoreUpdateMsrpParameter = 'StoreUpdateMsrpParameter',
  SwitchMsrpParameter = 'SwitchMsrpParameter',
  StoreSwitchMsrpParameter = 'StoreSwitchMsrpParameter',
  AddMsrpParameter = 'AddMsrpParameter',
  StoreAddMsrpParameter = 'StoreAddMsrpParameter',
  DelMsrpParameter = 'DelMsrpParameter',
  StoreDelMsrpParameter = 'StoreDelMsrpParameter',
  StoreNewMsrpParameter = 'StoreNewMsrpParameter',
  StoreDropNewMsrpParameter = 'StoreDropNewMsrpParameter',
  StoreGotMsrpError = 'StoreGotMsrpError',
}

export class GetMsrp implements Action {
  readonly type = ConfigActionTypes.GetMsrp;
  constructor(public payload: any) {}
}

export class StoreGetMsrp implements Action {
  readonly type = ConfigActionTypes.StoreGetMsrp;
  constructor(public payload: any) {}
}

export class UpdateMsrpParameter implements Action {
  readonly type = ConfigActionTypes.UpdateMsrpParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateMsrpParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateMsrpParameter;
  constructor(public payload: any) {}
}

export class SwitchMsrpParameter implements Action {
  readonly type = ConfigActionTypes.SwitchMsrpParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchMsrpParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchMsrpParameter;
  constructor(public payload: any) {}
}

export class AddMsrpParameter implements Action {
  readonly type = ConfigActionTypes.AddMsrpParameter;
  constructor(public payload: any) {}
}

export class StoreAddMsrpParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddMsrpParameter;
  constructor(public payload: any) {}
}

export class DelMsrpParameter implements Action {
  readonly type = ConfigActionTypes.DelMsrpParameter;
  constructor(public payload: any) {}
}

export class StoreDelMsrpParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelMsrpParameter;
  constructor(public payload: any) {}
}

export class StoreGotMsrpError implements Action {
  readonly type = ConfigActionTypes.StoreGotMsrpError;
  constructor(public payload: any) {}
}

export class StoreDropNewMsrpParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewMsrpParameter;
  constructor(public payload: any) {}
}

export class StoreNewMsrpParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewMsrpParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetMsrp
  | StoreGetMsrp
  | UpdateMsrpParameter
  | StoreUpdateMsrpParameter
  | SwitchMsrpParameter
  | StoreSwitchMsrpParameter
  | AddMsrpParameter
  | StoreAddMsrpParameter
  | DelMsrpParameter
  | StoreDelMsrpParameter
  | StoreGotMsrpError
  | StoreDropNewMsrpParameter
  | StoreNewMsrpParameter
;

