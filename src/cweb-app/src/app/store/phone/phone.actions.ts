import { Action } from '@ngrx/store';
import {CDRActionTypes} from '../cdr/cdr.actions';

export enum AuthActionTypes {
  UPDATE_FAILURE = '[Phone] Failure',
  GET_PHONE_CREDS = '[Phone][Get] Creds',
  STORE_GET_PHONE_CREDS = '[Phone][Store][Get] Creds',
  STORE_PHONE_STATUS = '[Phone][Store] Status',
  STORE_MAKE_PHONE_CALL = '[Phone][Store][Make] Call',
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
  readonly type = AuthActionTypes.STORE_GET_PHONE_CREDS;
  constructor(public payload: any) {}
}

export class StorePhoneStatus implements Action {
  readonly type = AuthActionTypes.STORE_PHONE_STATUS;
  constructor(public payload: any) {}
}

export class StoreMakePhoneCall implements Action {
  readonly type = AuthActionTypes.STORE_MAKE_PHONE_CALL;
  constructor(public payload: any) {}
}

export type All =
  | Failure
  | GetPhoneCreds
  | StoreGetPhoneCreds
  | StorePhoneStatus
  | StoreMakePhoneCall
  ;
