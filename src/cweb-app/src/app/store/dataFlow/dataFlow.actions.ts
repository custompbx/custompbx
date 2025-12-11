import { Action } from '@ngrx/store';

export enum DataFlowActionTypes {
  UPDATE_FAILURE = '[Dataflow] Failure',
  REDUCE_LOAD_COUNTER = '[Dataflow] Reduce load counter',
  GET_DASHBOARD = '[Dashboard]',
  STORE_GET_DASHBOARD = '[Dashboard][Store]',
  UnSubscribe = '[UnSubscribe]',
  SubscriptionList = 'SubscriptionList',
  PersistentSubscription = 'PersistentSubscription',
}

export class Failure implements Action {
  readonly type = DataFlowActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}
export class ReduceLoadCounter implements Action {
  readonly type = DataFlowActionTypes.REDUCE_LOAD_COUNTER;
}

export class GetDashboard implements Action {
  readonly type = DataFlowActionTypes.GET_DASHBOARD;
  constructor(public payload: any) {}
}

export class StoreGetDashboard implements Action {
  readonly type = DataFlowActionTypes.STORE_GET_DASHBOARD;
  constructor(public payload: any) {}
}

export class UnSubscribe implements Action {
  readonly type = DataFlowActionTypes.UnSubscribe;
  constructor(public payload: any) {}
}

export class SubscriptionList implements Action {
  readonly type = DataFlowActionTypes.SubscriptionList;
  constructor(public payload: any) {}
}

export class PersistentSubscription implements Action {
  readonly type = DataFlowActionTypes.PersistentSubscription;
  constructor(public payload: any) {}
}

export type All =
  | Failure
  | ReduceLoadCounter
  | GetDashboard
  | StoreGetDashboard
  | UnSubscribe
  | SubscriptionList
  | PersistentSubscription
;
