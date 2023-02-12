import { Action } from '@ngrx/store';

export enum AuthActionTypes {
  LOGIN = '[Auth] Login',
  RELOGIN = '[Auth] Relogin',
  LOGIN_SUCCESS = '[Auth] Login Success',
  RELOGIN_SUCCESS = '[Auth] Relogin Success',
  LOGIN_FAILURE = '[Auth] Login Failure',
  LOGOUT = '[Auth] Logout',
  GET_STATUS = '[Auth] GetStatus',
}

export class LogIn implements Action {
  readonly type = AuthActionTypes.LOGIN;
  constructor(public payload: any) {}
}

export class LogInSuccess implements Action {
  readonly type = AuthActionTypes.LOGIN_SUCCESS;
  constructor(public payload: any) {}
}

export class LogInFailure implements Action {
  readonly type = AuthActionTypes.LOGIN_FAILURE;
  constructor(public payload: any) {}
}

export class LogOut implements Action {
  readonly type = AuthActionTypes.LOGOUT;
}

export class GetStatus implements Action {
  readonly type = AuthActionTypes.GET_STATUS;
}

export class ReLogIn implements Action {
  readonly type = AuthActionTypes.RELOGIN;
  constructor(public payload: any) {}
}

export class ReLogInSuccess implements Action {
  readonly type = AuthActionTypes.RELOGIN_SUCCESS;
  constructor(public payload: any) {}
}

export type All =
  | LogIn
  | LogInSuccess
  | LogInFailure
  | LogOut
  | GetStatus
  | ReLogIn
  | ReLogInSuccess
  ;
