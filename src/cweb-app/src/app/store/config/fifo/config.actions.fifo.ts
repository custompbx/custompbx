import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotFifoError = 'StoreGotFifoError',
  GetFifo = 'GetFifo',
  StoreGetFifo = 'StoreGetFifo',
  GetFifoFifoMembers = 'GetFifoFifoMembers',
  StoreGetFifoFifoMembers = 'StoreGetFifoFifoMembers',
  UpdateFifoParameter = 'UpdateFifoParameter',
  StoreUpdateFifoParameter = 'StoreUpdateFifoParameter',
  SwitchFifoParameter = 'SwitchFifoParameter',
  StoreSwitchFifoParameter = 'StoreSwitchFifoParameter',
  AddFifoParameter = 'AddFifoParameter',
  StoreAddFifoParameter = 'StoreAddFifoParameter',
  DelFifoParameter = 'DelFifoParameter',
  StoreDelFifoParameter = 'StoreDelFifoParameter',
  StoreNewFifoParameter = 'StoreNewFifoParameter',
  StoreDropNewFifoParameter = 'StoreDropNewFifoParameter',
  AddFifoFifoMember = 'AddFifoFifoMember',
  StoreAddFifoFifoMember = 'StoreAddFifoFifoMember',
  UpdateFifoFifoMember = 'UpdateFifoFifoMember',
  StoreUpdateFifoFifoMember = 'StoreUpdateFifoFifoMember',
  SwitchFifoFifoMember = 'SwitchFifoFifoMember',
  StoreSwitchFifoFifoMember = 'StoreSwitchFifoFifoMember',
  DelFifoFifoMember = 'DelFifoFifoMember',
  StoreDelFifoFifoMember = 'StoreDelFifoFifoMember',
  StoreNewFifoFifoMember = 'StoreNewFifoFifoMember',
  StoreDropNewFifoFifoMember = 'StoreDropNewFifoFifoMember',
  StorePasteFifoFifoMembers = 'StorePasteFifoFifoMembers',
  AddFifoFifo = 'AddFifoFifo',
  StoreAddFifoFifo = 'StoreAddFifoFifo',
  DelFifoFifo = 'DelFifoFifo',
  StoreDelFifoFifo = 'StoreDelFifoFifo',
  UpdateFifoFifo = 'UpdateFifoFifo',
  StoreUpdateFifoFifo = 'StoreUpdateFifoFifo',
  UpdateFifoFifoImportance = 'UpdateFifoFifoImportance',
  StoreUpdateFifoFifoImportance = 'StoreUpdateFifoFifoImportance',
}

export class UpdateFifoFifoImportance implements Action {
  readonly type = ConfigActionTypes.UpdateFifoFifoImportance;
  constructor(public payload: any) {}
}

export class StoreUpdateFifoFifoImportance implements Action {
  readonly type = ConfigActionTypes.StoreUpdateFifoFifoImportance;
  constructor(public payload: any) {}
}

export class GetFifo implements Action {
  readonly type = ConfigActionTypes.GetFifo;
  constructor(public payload: any) {}
}

export class StoreGetFifo implements Action {
  readonly type = ConfigActionTypes.StoreGetFifo;
  constructor(public payload: any) {}
}

export class GetFifoFifoMembers implements Action {
  readonly type = ConfigActionTypes.GetFifoFifoMembers;
  constructor(public payload: any) {}
}

export class StoreGetFifoFifoMembers implements Action {
  readonly type = ConfigActionTypes.StoreGetFifoFifoMembers;
  constructor(public payload: any) {}
}

export class UpdateFifoParameter implements Action {
  readonly type = ConfigActionTypes.UpdateFifoParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateFifoParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateFifoParameter;
  constructor(public payload: any) {}
}

export class SwitchFifoParameter implements Action {
  readonly type = ConfigActionTypes.SwitchFifoParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchFifoParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchFifoParameter;
  constructor(public payload: any) {}
}

export class AddFifoParameter implements Action {
  readonly type = ConfigActionTypes.AddFifoParameter;
  constructor(public payload: any) {}
}

export class StoreAddFifoParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddFifoParameter;
  constructor(public payload: any) {}
}

export class DelFifoParameter implements Action {
  readonly type = ConfigActionTypes.DelFifoParameter;
  constructor(public payload: any) {}
}

export class StoreDelFifoParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelFifoParameter;
  constructor(public payload: any) {}
}

export class AddFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.AddFifoFifoMember;
  constructor(public payload: any) {}
}

export class StoreAddFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.StoreAddFifoFifoMember;
  constructor(public payload: any) {}
}

export class UpdateFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.UpdateFifoFifoMember;
  constructor(public payload: any) {}
}

export class StoreUpdateFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.StoreUpdateFifoFifoMember;
  constructor(public payload: any) {}
}

export class SwitchFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.SwitchFifoFifoMember;
  constructor(public payload: any) {}
}

export class StoreSwitchFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.StoreSwitchFifoFifoMember;
  constructor(public payload: any) {}
}

export class DelFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.DelFifoFifoMember;
  constructor(public payload: any) {}
}

export class StoreDelFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.StoreDelFifoFifoMember;
  constructor(public payload: any) {}
}

export class StoreNewFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.StoreNewFifoFifoMember;
  constructor(public payload: any) {}
}

export class StoreDropNewFifoFifoMember implements Action {
  readonly type = ConfigActionTypes.StoreDropNewFifoFifoMember;
  constructor(public payload: any) {}
}

export class StorePasteFifoFifoMembers implements Action {
  readonly type = ConfigActionTypes.StorePasteFifoFifoMembers;
  constructor(public payload: any) {}
}

export class AddFifoFifo implements Action {
  readonly type = ConfigActionTypes.AddFifoFifo;
  constructor(public payload: any) {}
}

export class StoreAddFifoFifo implements Action {
  readonly type = ConfigActionTypes.StoreAddFifoFifo;
  constructor(public payload: any) {}
}

export class DelFifoFifo implements Action {
  readonly type = ConfigActionTypes.DelFifoFifo;
  constructor(public payload: any) {}
}

export class StoreDelFifoFifo implements Action {
  readonly type = ConfigActionTypes.StoreDelFifoFifo;
  constructor(public payload: any) {}
}

export class UpdateFifoFifo implements Action {
  readonly type = ConfigActionTypes.UpdateFifoFifo;
  constructor(public payload: any) {}
}

export class StoreUpdateFifoFifo implements Action {
  readonly type = ConfigActionTypes.StoreUpdateFifoFifo;
  constructor(public payload: any) {}
}

export class StoreGotFifoError implements Action {
  readonly type = ConfigActionTypes.StoreGotFifoError;
  constructor(public payload: any) {}
}

export class StoreDropNewFifoParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewFifoParameter;
  constructor(public payload: any) {}
}

export class StoreNewFifoParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewFifoParameter;
  constructor(public payload: any) {}
}

export type All =
  | GetFifo
  | StoreGetFifo
  | GetFifoFifoMembers
  | StoreGetFifoFifoMembers
  | UpdateFifoParameter
  | StoreUpdateFifoParameter
  | SwitchFifoParameter
  | StoreSwitchFifoParameter
  | AddFifoParameter
  | StoreAddFifoParameter
  | DelFifoParameter
  | StoreDelFifoParameter
  | AddFifoFifoMember
  | StoreAddFifoFifoMember
  | UpdateFifoFifoMember
  | StoreUpdateFifoFifoMember
  | SwitchFifoFifoMember
  | StoreSwitchFifoFifoMember
  | DelFifoFifoMember
  | StoreDelFifoFifoMember
  | StoreNewFifoFifoMember
  | StoreDropNewFifoFifoMember
  | StorePasteFifoFifoMembers
  | AddFifoFifo
  | StoreAddFifoFifo
  | DelFifoFifo
  | StoreDelFifoFifo
  | UpdateFifoFifo
  | StoreUpdateFifoFifo
  | StoreGotFifoError
  | StoreDropNewFifoParameter
  | StoreNewFifoParameter
  | UpdateFifoFifoImportance
  | StoreUpdateFifoFifoImportance
;

