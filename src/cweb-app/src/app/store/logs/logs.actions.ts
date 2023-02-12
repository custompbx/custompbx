import { Action } from '@ngrx/store';

export enum LogsActionTypes {
  GetLogs = 'GetLogs',
  StoreGetLogs = 'StoreGetLogs',
  StoreGotLogsError = 'StoreGotLogsError',
}


export class GetLogs implements Action {
  readonly type = LogsActionTypes.GetLogs;
  constructor(public payload: any) {}
}

export class StoreGetLogs implements Action {
  readonly type = LogsActionTypes.StoreGetLogs;
  constructor(public payload: any) {}
}

export class StoreGotLogsError implements Action {
  readonly type = LogsActionTypes.StoreGotLogsError;
  constructor(public payload: any) {}
}

export type All =
  | GetLogs
  | StoreGetLogs
  | StoreGotLogsError
;
