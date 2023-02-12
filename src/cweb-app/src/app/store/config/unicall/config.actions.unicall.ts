import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotUnicallError = 'StoreGotUnicallError',
  GetUnicall = 'GetUnicall',
  StoreGetUnicall = 'StoreGetUnicall',
  GetUnicallSpanParameters = 'GetUnicallSpanParameters',
  StoreGetUnicallSpanParameters = 'StoreGetUnicallSpanParameters',
  UpdateUnicallParameter = 'UpdateUnicallParameter',
  StoreUpdateUnicallParameter = 'StoreUpdateUnicallParameter',
  SwitchUnicallParameter = 'SwitchUnicallParameter',
  StoreSwitchUnicallParameter = 'StoreSwitchUnicallParameter',
  AddUnicallParameter = 'AddUnicallParameter',
  StoreAddUnicallParameter = 'StoreAddUnicallParameter',
  DelUnicallParameter = 'DelUnicallParameter',
  StoreDelUnicallParameter = 'StoreDelUnicallParameter',
  StoreNewUnicallParameter = 'StoreNewUnicallParameter',
  StoreDropNewUnicallParameter = 'StoreDropNewUnicallParameter',
  AddUnicallSpanParameter = 'AddUnicallSpanParameter',
  StoreAddUnicallSpanParameter = 'StoreAddUnicallSpanParameter',
  UpdateUnicallSpanParameter = 'UpdateUnicallSpanParameter',
  StoreUpdateUnicallSpanParameter = 'StoreUpdateUnicallSpanParameter',
  SwitchUnicallSpanParameter = 'SwitchUnicallSpanParameter',
  StoreSwitchUnicallSpanParameter = 'StoreSwitchUnicallSpanParameter',
  DelUnicallSpanParameter = 'DelUnicallSpanParameter',
  StoreDelUnicallSpanParameter = 'StoreDelUnicallSpanParameter',
  StoreNewUnicallSpanParameter = 'StoreNewUnicallSpanParameter',
  StoreDropNewUnicallSpanParameter = 'StoreDropNewUnicallSpanParameter',
  StorePasteUnicallSpanParameters = 'StorePasteUnicallSpanParameters',
  AddUnicallSpan = 'AddUnicallSpan',
  StoreAddUnicallSpan = 'StoreAddUnicallSpan',
  DelUnicallSpan = 'DelUnicallSpan',
  StoreDelUnicallSpan = 'StoreDelUnicallSpan',
  UpdateUnicallSpan = 'UpdateUnicallSpan',
  StoreUpdateUnicallSpan = 'StoreUpdateUnicallSpan',
}

export class GetUnicall implements Action {
  readonly type = ConfigActionTypes.GetUnicall;
  constructor(public payload: any) {}
}

export class StoreGetUnicall implements Action {
  readonly type = ConfigActionTypes.StoreGetUnicall;
  constructor(public payload: any) {}
}

export class GetUnicallSpanParameters implements Action {
  readonly type = ConfigActionTypes.GetUnicallSpanParameters;
  constructor(public payload: any) {}
}

export class StoreGetUnicallSpanParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetUnicallSpanParameters;
  constructor(public payload: any) {}
}

export class UpdateUnicallParameter implements Action {
  readonly type = ConfigActionTypes.UpdateUnicallParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateUnicallParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateUnicallParameter;
  constructor(public payload: any) {}
}

export class SwitchUnicallParameter implements Action {
  readonly type = ConfigActionTypes.SwitchUnicallParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchUnicallParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchUnicallParameter;
  constructor(public payload: any) {}
}

export class AddUnicallParameter implements Action {
  readonly type = ConfigActionTypes.AddUnicallParameter;
  constructor(public payload: any) {}
}

export class StoreAddUnicallParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddUnicallParameter;
  constructor(public payload: any) {}
}

export class DelUnicallParameter implements Action {
  readonly type = ConfigActionTypes.DelUnicallParameter;
  constructor(public payload: any) {}
}

export class StoreDelUnicallParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelUnicallParameter;
  constructor(public payload: any) {}
}

export class AddUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.AddUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StoreAddUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class UpdateUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.UpdateUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class SwitchUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.SwitchUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class DelUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.DelUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StoreDelUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StoreNewUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewUnicallSpanParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewUnicallSpanParameter;
  constructor(public payload: any) {}
}

export class StorePasteUnicallSpanParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteUnicallSpanParameters;
  constructor(public payload: any) {}
}

export class AddUnicallSpan implements Action {
  readonly type = ConfigActionTypes.AddUnicallSpan;
  constructor(public payload: any) {}
}

export class StoreAddUnicallSpan implements Action {
  readonly type = ConfigActionTypes.StoreAddUnicallSpan;
  constructor(public payload: any) {}
}

export class DelUnicallSpan implements Action {
  readonly type = ConfigActionTypes.DelUnicallSpan;
  constructor(public payload: any) {}
}

export class StoreDelUnicallSpan implements Action {
  readonly type = ConfigActionTypes.StoreDelUnicallSpan;
  constructor(public payload: any) {}
}

export class UpdateUnicallSpan implements Action {
  readonly type = ConfigActionTypes.UpdateUnicallSpan;
  constructor(public payload: any) {}
}

export class StoreUpdateUnicallSpan implements Action {
  readonly type = ConfigActionTypes.StoreUpdateUnicallSpan;
  constructor(public payload: any) {}
}

export class StoreGotUnicallError implements Action {
  readonly type = ConfigActionTypes.StoreGotUnicallError;
  constructor(public payload: any) {}
}

export class StoreDropNewUnicallParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewUnicallParameter;
  constructor(public payload: any) {}
}

export class StoreNewUnicallParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewUnicallParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetUnicall
  | StoreGetUnicall
  | GetUnicallSpanParameters
  | StoreGetUnicallSpanParameters
  | UpdateUnicallParameter
  | StoreUpdateUnicallParameter
  | SwitchUnicallParameter
  | StoreSwitchUnicallParameter
  | AddUnicallParameter
  | StoreAddUnicallParameter
  | DelUnicallParameter
  | StoreDelUnicallParameter
  | AddUnicallSpanParameter
  | StoreAddUnicallSpanParameter
  | UpdateUnicallSpanParameter
  | StoreUpdateUnicallSpanParameter
  | SwitchUnicallSpanParameter
  | StoreSwitchUnicallSpanParameter
  | DelUnicallSpanParameter
  | StoreDelUnicallSpanParameter
  | StoreNewUnicallSpanParameter
  | StoreDropNewUnicallSpanParameter
  | StorePasteUnicallSpanParameters
  | AddUnicallSpan
  | StoreAddUnicallSpan
  | DelUnicallSpan
  | StoreDelUnicallSpan
  | UpdateUnicallSpan
  | StoreUpdateUnicallSpan
  | StoreGotUnicallError
  | StoreDropNewUnicallParameter
  | StoreNewUnicallParameter
;

