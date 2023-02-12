import { Action } from '@ngrx/store';

export enum DaemonActionTypes {
  CONNECTION = '[Daemon] Connection',
  STORE_FLUSH_TOKEN_STATE = '[Daemon][Store][Token] State',
}

export class Status implements Action {
  readonly type = DaemonActionTypes.CONNECTION;
  constructor(public payload: any) {}
}

export class StoreFlushTokenState implements Action {
  readonly type = DaemonActionTypes.STORE_FLUSH_TOKEN_STATE;
  constructor(public payload: any) {}
}

export type All = Status | StoreFlushTokenState;
