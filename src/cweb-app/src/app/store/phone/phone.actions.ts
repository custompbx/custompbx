import { Action } from '@ngrx/store';
import {CDRActionTypes} from '../cdr/cdr.actions';
import {createActionHelper} from '../../services/rxjs-helper/actions-helper';

export enum AuthActionTypes {
  UPDATE_FAILURE = '[Phone] Failure',
  GET_PHONE_CREDS = '[Phone][Get] Creds',
  StoreGetPhoneCreds = 'StoreGetPhoneCreds',
  StorePhoneStatus = 'StorePhoneStatus',
}

export class Failure implements Action {
  readonly type = CDRActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}

export class GetPhoneCreds implements Action {
  readonly type = AuthActionTypes.GET_PHONE_CREDS;
  constructor(public payload: any) {}
}

export class StoreGetPhoneCreds implements Action {
  readonly type = AuthActionTypes.StoreGetPhoneCreds;
  constructor(public payload: any) {}
}

export class StorePhoneStatus implements Action {
  readonly type = AuthActionTypes.StorePhoneStatus;
  constructor(public payload: any) {}
}

export const StoreTicker = createActionHelper('StoreTicker');
export const StoreCommand = createActionHelper('StoreCommand');

export type All =
  | Failure
  | GetPhoneCreds
  | StoreGetPhoneCreds
  | StorePhoneStatus
  ;
