
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetMongo = 'GetMongo',
  StoreGetMongo = 'StoreGetMongo',
  UpdateMongoParameter = 'UpdateMongoParameter',
  StoreUpdateMongoParameter = 'StoreUpdateMongoParameter',
  SwitchMongoParameter = 'SwitchMongoParameter',
  StoreSwitchMongoParameter = 'StoreSwitchMongoParameter',
  AddMongoParameter = 'AddMongoParameter',
  StoreAddMongoParameter = 'StoreAddMongoParameter',
  DelMongoParameter = 'DelMongoParameter',
  StoreDelMongoParameter = 'StoreDelMongoParameter',
  StoreNewMongoParameter = 'StoreNewMongoParameter',
  StoreDropNewMongoParameter = 'StoreDropNewMongoParameter',
  StoreGotMongoError = 'StoreGotMongoError',
}

export class GetMongo implements Action {
  readonly type = ConfigActionTypes.GetMongo;
  constructor(public payload: any) {}
}

export class StoreGetMongo implements Action {
  readonly type = ConfigActionTypes.StoreGetMongo;
  constructor(public payload: any) {}
}

export class UpdateMongoParameter implements Action {
  readonly type = ConfigActionTypes.UpdateMongoParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateMongoParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateMongoParameter;
  constructor(public payload: any) {}
}

export class SwitchMongoParameter implements Action {
  readonly type = ConfigActionTypes.SwitchMongoParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchMongoParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchMongoParameter;
  constructor(public payload: any) {}
}

export class AddMongoParameter implements Action {
  readonly type = ConfigActionTypes.AddMongoParameter;
  constructor(public payload: any) {}
}

export class StoreAddMongoParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddMongoParameter;
  constructor(public payload: any) {}
}

export class DelMongoParameter implements Action {
  readonly type = ConfigActionTypes.DelMongoParameter;
  constructor(public payload: any) {}
}

export class StoreDelMongoParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelMongoParameter;
  constructor(public payload: any) {}
}

export class StoreGotMongoError implements Action {
  readonly type = ConfigActionTypes.StoreGotMongoError;
  constructor(public payload: any) {}
}

export class StoreDropNewMongoParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewMongoParameter;
  constructor(public payload: any) {}
}

export class StoreNewMongoParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewMongoParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetMongo
  | StoreGetMongo
  | UpdateMongoParameter
  | StoreUpdateMongoParameter
  | SwitchMongoParameter
  | StoreSwitchMongoParameter
  | AddMongoParameter
  | StoreAddMongoParameter
  | DelMongoParameter
  | StoreDelMongoParameter
  | StoreGotMongoError
  | StoreDropNewMongoParameter
  | StoreNewMongoParameter
;

