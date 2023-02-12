
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetGlobalVariables = 'GetGlobalVariables',
  StoreGetGlobalVariables = 'StoreGetGlobalVariables',
  UpdateGlobalVariable = 'UpdateGlobalVariable',
  StoreUpdateGlobalVariable = 'StoreUpdateGlobalVariable',
  SwitchGlobalVariable = 'SwitchGlobalVariable',
  StoreSwitchGlobalVariable = 'StoreSwitchGlobalVariable',
  AddGlobalVariable = 'AddGlobalVariable',
  StoreAddGlobalVariable = 'StoreAddGlobalVariable',
  DelGlobalVariable = 'DelGlobalVariable',
  StoreDelGlobalVariable = 'StoreDelGlobalVariable',
  StoreNewGlobalVariable = 'StoreNewGlobalVariable',
  StoreDropNewGlobalVariable = 'StoreDropNewGlobalVariable',
  ImportGlobalVariables = 'ImportGlobalVariables',
  StoreImportGlobalVariables = 'StoreImportGlobalVariables',
  StoreGotGlobalVariableError = 'StoreGotGlobalVariableError',
  MoveGlobalVariable = 'MoveGlobalVariable',
  StoreMoveGlobalVariable = 'StoreMoveGlobalVariable',
}

export class GetGlobalVariables implements Action {
  readonly type = ConfigActionTypes.GetGlobalVariables;
  constructor(public payload: any) {}
}

export class StoreGetGlobalVariables implements Action {
  readonly type = ConfigActionTypes.StoreGetGlobalVariables;
  constructor(public payload: any) {}
}

export class UpdateGlobalVariable implements Action {
  readonly type = ConfigActionTypes.UpdateGlobalVariable;
  constructor(public payload: any) {}
}

export class StoreUpdateGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreUpdateGlobalVariable;
  constructor(public payload: any) {}
}

export class SwitchGlobalVariable implements Action {
  readonly type = ConfigActionTypes.SwitchGlobalVariable;
  constructor(public payload: any) {}
}

export class StoreSwitchGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreSwitchGlobalVariable;
  constructor(public payload: any) {}
}

export class AddGlobalVariable implements Action {
  readonly type = ConfigActionTypes.AddGlobalVariable;
  constructor(public payload: any) {}
}

export class StoreAddGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreAddGlobalVariable;
  constructor(public payload: any) {}
}

export class DelGlobalVariable implements Action {
  readonly type = ConfigActionTypes.DelGlobalVariable;
  constructor(public payload: any) {}
}

export class StoreDelGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreDelGlobalVariable;
  constructor(public payload: any) {}
}

export class StoreGotGlobalVariableError implements Action {
  readonly type = ConfigActionTypes.StoreGotGlobalVariableError;
  constructor(public payload: any) {}
}

export class StoreDropNewGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreDropNewGlobalVariable;
  constructor(public payload: any) {}
}

export class StoreNewGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreNewGlobalVariable;
  constructor(public payload: any) {}
}

export class ImportGlobalVariables implements Action {
  readonly type = ConfigActionTypes.ImportGlobalVariables;
  constructor(public payload: any) {}
}

export class StoreImportGlobalVariables implements Action {
  readonly type = ConfigActionTypes.StoreImportGlobalVariables;
  constructor(public payload: any) {}
}
export class MoveGlobalVariable implements Action {
  readonly type = ConfigActionTypes.MoveGlobalVariable;
  constructor(public payload: any) {}
}
export class StoreMoveGlobalVariable implements Action {
  readonly type = ConfigActionTypes.StoreMoveGlobalVariable;
  constructor(public payload: any) {}
}

export type All =
  | GetGlobalVariables
  | StoreGetGlobalVariables
  | UpdateGlobalVariable
  | StoreUpdateGlobalVariable
  | SwitchGlobalVariable
  | StoreSwitchGlobalVariable
  | AddGlobalVariable
  | StoreAddGlobalVariable
  | DelGlobalVariable
  | StoreDelGlobalVariable
  | StoreGotGlobalVariableError
  | StoreDropNewGlobalVariable
  | StoreNewGlobalVariable
  | ImportGlobalVariables
  | StoreImportGlobalVariables
  | MoveGlobalVariable
  | StoreMoveGlobalVariable
;

