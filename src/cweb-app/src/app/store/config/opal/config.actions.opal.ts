import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotOpalError = 'StoreGotOpalError',
  GetOpal = 'GetOpal',
  StoreGetOpal = 'StoreGetOpal',
  GetOpalListenerParameters = 'GetOpalListenerParameters',
  StoreGetOpalListenerParameters = 'StoreGetOpalListenerParameters',
  UpdateOpalParameter = 'UpdateOpalParameter',
  StoreUpdateOpalParameter = 'StoreUpdateOpalParameter',
  SwitchOpalParameter = 'SwitchOpalParameter',
  StoreSwitchOpalParameter = 'StoreSwitchOpalParameter',
  AddOpalParameter = 'AddOpalParameter',
  StoreAddOpalParameter = 'StoreAddOpalParameter',
  DelOpalParameter = 'DelOpalParameter',
  StoreDelOpalParameter = 'StoreDelOpalParameter',
  StoreNewOpalParameter = 'StoreNewOpalParameter',
  StoreDropNewOpalParameter = 'StoreDropNewOpalParameter',
  AddOpalListenerParameter = 'AddOpalListenerParameter',
  StoreAddOpalListenerParameter = 'StoreAddOpalListenerParameter',
  UpdateOpalListenerParameter = 'UpdateOpalListenerParameter',
  StoreUpdateOpalListenerParameter = 'StoreUpdateOpalListenerParameter',
  SwitchOpalListenerParameter = 'SwitchOpalListenerParameter',
  StoreSwitchOpalListenerParameter = 'StoreSwitchOpalListenerParameter',
  DelOpalListenerParameter = 'DelOpalListenerParameter',
  StoreDelOpalListenerParameter = 'StoreDelOpalListenerParameter',
  StoreNewOpalListenerParameter = 'StoreNewOpalListenerParameter',
  StoreDropNewOpalListenerParameter = 'StoreDropNewOpalListenerParameter',
  StorePasteOpalListenerParameters = 'StorePasteOpalListenerParameters',
  AddOpalListener = 'AddOpalListener',
  StoreAddOpalListener = 'StoreAddOpalListener',
  DelOpalListener = 'DelOpalListener',
  StoreDelOpalListener = 'StoreDelOpalListener',
  UpdateOpalListener = 'UpdateOpalListener',
  StoreUpdateOpalListener = 'StoreUpdateOpalListener',
}

export class GetOpal implements Action {
  readonly type = ConfigActionTypes.GetOpal;
  constructor(public payload: any) {}
}

export class StoreGetOpal implements Action {
  readonly type = ConfigActionTypes.StoreGetOpal;
  constructor(public payload: any) {}
}

export class GetOpalListenerParameters implements Action {
  readonly type = ConfigActionTypes.GetOpalListenerParameters;
  constructor(public payload: any) {}
}

export class StoreGetOpalListenerParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetOpalListenerParameters;
  constructor(public payload: any) {}
}

export class UpdateOpalParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOpalParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOpalParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOpalParameter;
  constructor(public payload: any) {}
}

export class SwitchOpalParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOpalParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOpalParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOpalParameter;
  constructor(public payload: any) {}
}

export class AddOpalParameter implements Action {
  readonly type = ConfigActionTypes.AddOpalParameter;
  constructor(public payload: any) {}
}

export class StoreAddOpalParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOpalParameter;
  constructor(public payload: any) {}
}

export class DelOpalParameter implements Action {
  readonly type = ConfigActionTypes.DelOpalParameter;
  constructor(public payload: any) {}
}

export class StoreDelOpalParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelOpalParameter;
  constructor(public payload: any) {}
}

export class AddOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.AddOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StoreAddOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddOpalListenerParameter;
  constructor(public payload: any) {}
}

export class UpdateOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.UpdateOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOpalListenerParameter;
  constructor(public payload: any) {}
}

export class SwitchOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.SwitchOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchOpalListenerParameter;
  constructor(public payload: any) {}
}

export class DelOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.DelOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StoreDelOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StoreNewOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewOpalListenerParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOpalListenerParameter;
  constructor(public payload: any) {}
}

export class StorePasteOpalListenerParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteOpalListenerParameters;
  constructor(public payload: any) {}
}

export class AddOpalListener implements Action {
  readonly type = ConfigActionTypes.AddOpalListener;
  constructor(public payload: any) {}
}

export class StoreAddOpalListener implements Action {
  readonly type = ConfigActionTypes.StoreAddOpalListener;
  constructor(public payload: any) {}
}

export class DelOpalListener implements Action {
  readonly type = ConfigActionTypes.DelOpalListener;
  constructor(public payload: any) {}
}

export class StoreDelOpalListener implements Action {
  readonly type = ConfigActionTypes.StoreDelOpalListener;
  constructor(public payload: any) {}
}

export class UpdateOpalListener implements Action {
  readonly type = ConfigActionTypes.UpdateOpalListener;
  constructor(public payload: any) {}
}

export class StoreUpdateOpalListener implements Action {
  readonly type = ConfigActionTypes.StoreUpdateOpalListener;
  constructor(public payload: any) {}
}

export class StoreGotOpalError implements Action {
  readonly type = ConfigActionTypes.StoreGotOpalError;
  constructor(public payload: any) {}
}

export class StoreDropNewOpalParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewOpalParameter;
  constructor(public payload: any) {}
}

export class StoreNewOpalParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewOpalParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetOpal
  | StoreGetOpal
  | GetOpalListenerParameters
  | StoreGetOpalListenerParameters
  | UpdateOpalParameter
  | StoreUpdateOpalParameter
  | SwitchOpalParameter
  | StoreSwitchOpalParameter
  | AddOpalParameter
  | StoreAddOpalParameter
  | DelOpalParameter
  | StoreDelOpalParameter
  | AddOpalListenerParameter
  | StoreAddOpalListenerParameter
  | UpdateOpalListenerParameter
  | StoreUpdateOpalListenerParameter
  | SwitchOpalListenerParameter
  | StoreSwitchOpalListenerParameter
  | DelOpalListenerParameter
  | StoreDelOpalListenerParameter
  | StoreNewOpalListenerParameter
  | StoreDropNewOpalListenerParameter
  | StorePasteOpalListenerParameters
  | AddOpalListener
  | StoreAddOpalListener
  | DelOpalListener
  | StoreDelOpalListener
  | UpdateOpalListener
  | StoreUpdateOpalListener
  | StoreGotOpalError
  | StoreDropNewOpalParameter
  | StoreNewOpalParameter
;

