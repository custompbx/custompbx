
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetCdrMongodb = 'GetCdrMongodb',
  StoreGetCdrMongodb = 'StoreGetCdrMongodb',
  UpdateCdrMongodbParameter = 'UpdateCdrMongodbParameter',
  StoreUpdateCdrMongodbParameter = 'StoreUpdateCdrMongodbParameter',
  SwitchCdrMongodbParameter = 'SwitchCdrMongodbParameter',
  StoreSwitchCdrMongodbParameter = 'StoreSwitchCdrMongodbParameter',
  AddCdrMongodbParameter = 'AddCdrMongodbParameter',
  StoreAddCdrMongodbParameter = 'StoreAddCdrMongodbParameter',
  DelCdrMongodbParameter = 'DelCdrMongodbParameter',
  StoreDelCdrMongodbParameter = 'StoreDelCdrMongodbParameter',
  StoreNewCdrMongodbParameter = 'StoreNewCdrMongodbParameter',
  StoreDropNewCdrMongodbParameter = 'StoreDropNewCdrMongodbParameter',
  StoreGotCdrMongodbError = 'StoreGotCdrMongodbError',
}

export class GetCdrMongodb implements Action {
  readonly type = ConfigActionTypes.GetCdrMongodb;
  constructor(public payload: any) {}
}

export class StoreGetCdrMongodb implements Action {
  readonly type = ConfigActionTypes.StoreGetCdrMongodb;
  constructor(public payload: any) {}
}

export class UpdateCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.UpdateCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class SwitchCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.SwitchCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class AddCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.AddCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class StoreAddCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class DelCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.DelCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class StoreDelCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class StoreGotCdrMongodbError implements Action {
  readonly type = ConfigActionTypes.StoreGotCdrMongodbError;
  constructor(public payload: any) {}
}

export class StoreDropNewCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewCdrMongodbParameter;
  constructor(public payload: any) {}
}

export class StoreNewCdrMongodbParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewCdrMongodbParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetCdrMongodb
  | StoreGetCdrMongodb
  | UpdateCdrMongodbParameter
  | StoreUpdateCdrMongodbParameter
  | SwitchCdrMongodbParameter
  | StoreSwitchCdrMongodbParameter
  | AddCdrMongodbParameter
  | StoreAddCdrMongodbParameter
  | DelCdrMongodbParameter
  | StoreDelCdrMongodbParameter
  | StoreGotCdrMongodbError
  | StoreDropNewCdrMongodbParameter
  | StoreNewCdrMongodbParameter
;

