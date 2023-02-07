import { Action } from '@ngrx/store';
import {StoreImportAllModules} from '../config.actions';

export enum ConfigActionTypes {
  GetOdbcCdr = 'GetOdbcCdr',
  StoreGetOdbcCdr = 'StoreGetOdbcCdr',
  StoreNewOdbcCdrParameter = 'StoreNewOdbcCdrParameter',
  StoreDropNewOdbcCdrParameter = 'StoreDropNewOdbcCdrParameter',
  UpdateOdbcCdrParameter = 'UpdateOdbcCdrParameter',
  StoreUpdateOdbcCdrParameter = 'StoreUpdateOdbcCdrParameter',
  SwitchOdbcCdrParameter = 'SwitchOdbcCdrParameter',
  StoreSwitchOdbcCdrParameter = 'StoreSwitchOdbcCdrParameter',
  DeleteOdbcCdrParameter = 'DeleteOdbcCdrParameter',
  StoreDeleteOdbcCdrParameter = 'StoreDeleteOdbcCdrParameter',
  StoreAddOdbcCdrParameter = 'StoreAddOdbcCdrParameter',
  AddOdbcCdrParameter = 'AddOdbcCdrParameter',

  StoreNewOdbcCdrTable = 'StoreNewOdbcCdrTable',
  StoreDropNewOdbcCdrTable = 'StoreDropNewOdbcCdrTable',
  AddOdbcCdrTable = 'AddOdbcCdrTable',
  StoreAddOdbcCdrTable = 'StoreAddOdbcCdrTable',
  UpdateOdbcCdrTable = 'UpdateOdbcCdrTable',
  StoreUpdateOdbcCdrTable = 'StoreUpdateOdbcCdrTable',
  SwitchOdbcCdrTable = 'SwitchOdbcCdrTable',
  StoreSwitchOdbcCdrTable = 'StoreSwitchOdbcCdrTable',
  DeleteOdbcCdrTable = 'DeleteOdbcCdrTable',
  StoreDeleteOdbcCdrTable = 'StoreDeleteOdbcCdrTable',

  StoreNewOdbcCdrField = 'StoreNewOdbcCdrField',
  StoreDropNewOdbcCdrField = 'StoreDropNewOdbcCdrField',
  GetOdbcCdrField = 'GetOdbcCdrField',
  StoreGetOdbcCdrField = 'StoreGetOdbcCdrField',
  AddOdbcCdrField = 'AddOdbcCdrField',
  StoreAddOdbcCdrField = 'StoreAddOdbcCdrField',
  UpdateOdbcCdrField = 'UpdateOdbcCdrField',
  StoreUpdateOdbcCdrField = 'StoreUpdateOdbcCdrField',
  SwitchOdbcCdrField = 'SwitchOdbcCdrField',
  StoreSwitchOdbcCdrField = 'StoreSwitchOdbcCdrField',
  DeleteOdbcCdrField = 'DeleteOdbcCdrField',
  StoreDeleteOdbcCdrField = 'StoreDeleteOdbcCdrField',

  StoreGotOdbcCdrError = 'StoreGotOdbcCdrError',
}

export class StoreGotOdbcCdrError implements Action {
  readonly type = ConfigActionTypes.StoreGotOdbcCdrError;
  constructor(public payload: any) {}
}

export class GetOdbcCdr implements Action {
  readonly type = ConfigActionTypes.GetOdbcCdr;
  constructor(public payload: any) {}
}

export class StoreGetOdbcCdr implements Action {
  readonly type = ConfigActionTypes.StoreGetOdbcCdr;
  constructor(public payload: any) {}
}

export class StoreNewOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class UpdateOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class SwitchOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class DeleteOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.DeleteOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class StoreDeleteOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreDeleteOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class StoreAddOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class AddOdbcCdrParameter implements Action {
  readonly type = ConfigActionTypes.AddOdbcCdrParameter;
  constructor(public payload: any) {}
}

export class StoreNewOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.StoreNewOdbcCdrTable;
  constructor(public payload: any) {}
}

export class StoreDropNewOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOdbcCdrTable;
  constructor(public payload: any) {}
}

export class AddOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.AddOdbcCdrTable;
  constructor(public payload: any) {}
}

export class StoreAddOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.StoreAddOdbcCdrTable;
  constructor(public payload: any) {}
}

export class UpdateOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.UpdateOdbcCdrTable;
  constructor(public payload: any) {}
}

export class StoreUpdateOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOdbcCdrTable;
  constructor(public payload: any) {}
}

export class SwitchOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.SwitchOdbcCdrTable;
  constructor(public payload: any) {}
}

export class StoreSwitchOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOdbcCdrTable;
  constructor(public payload: any) {}
}

export class DeleteOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.DeleteOdbcCdrTable;
  constructor(public payload: any) {}
}

export class StoreDeleteOdbcCdrTable implements Action {
  readonly type = ConfigActionTypes.StoreDeleteOdbcCdrTable;
  constructor(public payload: any) {}
}

export class StoreNewOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreNewOdbcCdrField;
  constructor(public payload: any) {}
}

export class StoreDropNewOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOdbcCdrField;
  constructor(public payload: any) {}
}

export class AddOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.AddOdbcCdrField;
  constructor(public payload: any) {}
}

export class StoreAddOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreAddOdbcCdrField;
  constructor(public payload: any) {}
}

export class UpdateOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.UpdateOdbcCdrField;
  constructor(public payload: any) {}
}

export class StoreUpdateOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOdbcCdrField;
  constructor(public payload: any) {}
}

export class SwitchOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.SwitchOdbcCdrField;
  constructor(public payload: any) {}
}

export class StoreSwitchOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOdbcCdrField;
  constructor(public payload: any) {}
}

export class DeleteOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.DeleteOdbcCdrField;
  constructor(public payload: any) {}
}

export class StoreDeleteOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreDeleteOdbcCdrField;
  constructor(public payload: any) {}
}

export class GetOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.GetOdbcCdrField;
  constructor(public payload: any) {}
}

export class StoreGetOdbcCdrField implements Action {
  readonly type = ConfigActionTypes.StoreGetOdbcCdrField;
  constructor(public payload: any) {}
}

export type All =
  | GetOdbcCdr
  | StoreGetOdbcCdr
  | StoreNewOdbcCdrParameter
  | StoreDropNewOdbcCdrParameter
  | UpdateOdbcCdrParameter
  | StoreUpdateOdbcCdrParameter
  | SwitchOdbcCdrParameter
  | StoreSwitchOdbcCdrParameter
  | DeleteOdbcCdrParameter
  | StoreDeleteOdbcCdrParameter
  | StoreAddOdbcCdrParameter
  | AddOdbcCdrParameter
  | StoreNewOdbcCdrTable
  | StoreDropNewOdbcCdrTable
  | AddOdbcCdrTable
  | StoreAddOdbcCdrTable
  | UpdateOdbcCdrTable
  | StoreUpdateOdbcCdrTable
  | SwitchOdbcCdrTable
  | StoreSwitchOdbcCdrTable
  | DeleteOdbcCdrTable
  | StoreDeleteOdbcCdrTable
  | StoreNewOdbcCdrField
  | StoreDropNewOdbcCdrField
  | AddOdbcCdrField
  | StoreAddOdbcCdrField
  | UpdateOdbcCdrField
  | StoreUpdateOdbcCdrField
  | SwitchOdbcCdrField
  | StoreSwitchOdbcCdrField
  | DeleteOdbcCdrField
  | StoreDeleteOdbcCdrField
  | GetOdbcCdrField
  | StoreGetOdbcCdrField
  | StoreGotOdbcCdrError
;

