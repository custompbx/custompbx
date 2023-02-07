import { Action } from '@ngrx/store';

export enum FSCLIActionTypes {
  UPDATE_FAILURE = '[FSCLI] Failure',
  SendFSCLICommand = 'SendFSCLICommand',
  StoreSendFSCLICommand = 'StoreSendFSCLICommand',

}

export class Failure implements Action {
  readonly type = FSCLIActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}

export class SendFSCLICommand implements Action {
  readonly type = FSCLIActionTypes.SendFSCLICommand;
  constructor(public payload: any) {}
}

export class StoreSendFSCLICommand implements Action {
  readonly type = FSCLIActionTypes.StoreSendFSCLICommand;
  constructor(public payload: any) {}
}

export type All =
  | Failure
  | SendFSCLICommand
  | StoreSendFSCLICommand
;
