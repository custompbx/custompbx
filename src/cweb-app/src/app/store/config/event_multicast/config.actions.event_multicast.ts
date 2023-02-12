
import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetEventMulticast = 'GetEventMulticast',
  StoreGetEventMulticast = 'StoreGetEventMulticast',
  UpdateEventMulticastParameter = 'UpdateEventMulticastParameter',
  StoreUpdateEventMulticastParameter = 'StoreUpdateEventMulticastParameter',
  SwitchEventMulticastParameter = 'SwitchEventMulticastParameter',
  StoreSwitchEventMulticastParameter = 'StoreSwitchEventMulticastParameter',
  AddEventMulticastParameter = 'AddEventMulticastParameter',
  StoreAddEventMulticastParameter = 'StoreAddEventMulticastParameter',
  DelEventMulticastParameter = 'DelEventMulticastParameter',
  StoreDelEventMulticastParameter = 'StoreDelEventMulticastParameter',
  StoreNewEventMulticastParameter = 'StoreNewEventMulticastParameter',
  StoreDropNewEventMulticastParameter = 'StoreDropNewEventMulticastParameter',
  StoreGotEventMulticastError = 'StoreGotEventMulticastError',
}

export class GetEventMulticast implements Action {
  readonly type = ConfigActionTypes.GetEventMulticast;
  constructor(public payload: any) {}
}

export class StoreGetEventMulticast implements Action {
  readonly type = ConfigActionTypes.StoreGetEventMulticast;
  constructor(public payload: any) {}
}

export class UpdateEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.UpdateEventMulticastParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateEventMulticastParameter;
  constructor(public payload: any) {}
}

export class SwitchEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.SwitchEventMulticastParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchEventMulticastParameter;
  constructor(public payload: any) {}
}

export class AddEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.AddEventMulticastParameter;
  constructor(public payload: any) {}
}

export class StoreAddEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddEventMulticastParameter;
  constructor(public payload: any) {}
}

export class DelEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.DelEventMulticastParameter;
  constructor(public payload: any) {}
}

export class StoreDelEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelEventMulticastParameter;
  constructor(public payload: any) {}
}

export class StoreGotEventMulticastError implements Action {
  readonly type = ConfigActionTypes.StoreGotEventMulticastError;
  constructor(public payload: any) {}
}

export class StoreDropNewEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewEventMulticastParameter;
  constructor(public payload: any) {}
}

export class StoreNewEventMulticastParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewEventMulticastParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetEventMulticast
  | StoreGetEventMulticast
  | UpdateEventMulticastParameter
  | StoreUpdateEventMulticastParameter
  | SwitchEventMulticastParameter
  | StoreSwitchEventMulticastParameter
  | AddEventMulticastParameter
  | StoreAddEventMulticastParameter
  | DelEventMulticastParameter
  | StoreDelEventMulticastParameter
  | StoreGotEventMulticastError
  | StoreDropNewEventMulticastParameter
  | StoreNewEventMulticastParameter
;

