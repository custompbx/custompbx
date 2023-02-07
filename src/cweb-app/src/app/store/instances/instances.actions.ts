import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotInstancesError = 'StoreGotInstancesError',
  GetInstances = 'GetInstances',
  StoreGetInstances = 'StoreGetInstances',
  UpdateInstanceDescription = 'UpdateInstanceDescription',
  StoreUpdateInstanceDescription = 'StoreUpdateInstanceDescription',

}

export class StoreGotInstancesError implements Action {
  readonly type = ConfigActionTypes.StoreGotInstancesError;
  constructor(public payload: any) {}
}

export class GetInstances implements Action {
  readonly type = ConfigActionTypes.GetInstances;
  constructor(public payload: any) {}
}

export class StoreGetInstances implements Action {
  readonly type = ConfigActionTypes.StoreGetInstances;
  constructor(public payload: any) {}
}

export class UpdateInstanceDescription implements Action {
  readonly type = ConfigActionTypes.UpdateInstanceDescription;
  constructor(public payload: any) {}
}

export class StoreUpdateInstanceDescription implements Action {
  readonly type = ConfigActionTypes.StoreUpdateInstanceDescription;
  constructor(public payload: any) {}
}

export type All =
  | StoreGotInstancesError
  | GetInstances
  | StoreGetInstances
  | UpdateInstanceDescription
  | StoreUpdateInstanceDescription
  ;

