import { Action } from '@ngrx/store';
import {StoreImportXMLModuleConfig} from '../config.actions';

export enum ConfigActionTypes {
  GetPostSwitch = 'GetPostSwitch',
  StoreGetPostSwitch = 'StoreGetPostSwitch',
  StoreGotPostSwitchError = 'StoreGotPostSwitchError',

  UpdatePostSwitchParameter = 'UpdatePostSwitchParameter',
  StoreUpdatePostSwitchParameter = 'StoreUpdatePostSwitchParameter',
  SwitchPostSwitchParameter = 'SwitchPostSwitchParameter',
  StoreSwitchPostSwitchParameter = 'StoreSwitchPostSwitchParameter',
  AddPostSwitchParameter = 'AddPostSwitchParameter',
  StoreAddPostSwitchParameter = 'StoreAddPostSwitchParameter',
  DelPostSwitchParameter = 'DelPostSwitchParameter',
  StoreDelPostSwitchParameter = 'StoreDelPostSwitchParameter',
  StoreNewPostSwitchParameter = 'StoreNewPostSwitchParameter',
  StoreDropNewPostSwitchParameter = 'StoreDropNewPostSwitchParameter',

  UpdatePostSwitchDefaultPtime = 'UpdatePostSwitchDefaultPtime',
  StoreUpdatePostSwitchDefaultPtime = 'StoreUpdatePostSwitchDefaultPtime',
  SwitchPostSwitchDefaultPtime = 'SwitchPostSwitchDefaultPtime',
  StoreSwitchPostSwitchDefaultPtime = 'StoreSwitchPostSwitchDefaultPtime',
  AddPostSwitchDefaultPtime = 'AddPostSwitchDefaultPtime',
  StoreAddPostSwitchDefaultPtime = 'StoreAddPostSwitchDefaultPtime',
  DelPostSwitchDefaultPtime = 'DelPostSwitchDefaultPtime',
  StoreDelPostSwitchDefaultPtime = 'StoreDelPostSwitchDefaultPtime',
  StoreNewPostSwitchDefaultPtime = 'StoreNewPostSwitchDefaultPtime',
  StoreDropNewPostSwitchDefaultPtime = 'StoreDropNewPostSwitchDefaultPtime',

  UpdatePostSwitchCliKeybinding = 'UpdatePostSwitchCliKeybinding',
  StoreUpdatePostSwitchCliKeybinding = 'StoreUpdatePostSwitchCliKeybinding',
  SwitchPostSwitchCliKeybinding = 'SwitchPostSwitchCliKeybinding',
  StoreSwitchPostSwitchCliKeybinding = 'StoreSwitchPostSwitchCliKeybinding',
  AddPostSwitchCliKeybinding = 'AddPostSwitchCliKeybinding',
  StoreAddPostSwitchCliKeybinding = 'StoreAddPostSwitchCliKeybinding',
  DelPostSwitchCliKeybinding = 'DelPostSwitchCliKeybinding',
  StoreDelPostSwitchCliKeybinding = 'StoreDelPostSwitchCliKeybinding',
  StoreNewPostSwitchCliKeybinding = 'StoreNewPostSwitchCliKeybinding',
  StoreDropNewPostSwitchCliKeybinding = 'StoreDropNewPostSwitchCliKeybinding',

}

export class GetPostSwitch implements Action {
  readonly type = ConfigActionTypes.GetPostSwitch;
  constructor(public payload: any) {}
}

export class StoreGetPostSwitch implements Action {
  readonly type = ConfigActionTypes.StoreGetPostSwitch;
  constructor(public payload: any) {}
}

export class UpdatePostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.UpdatePostSwitchParameter;
  constructor(public payload: any) {}
}

export class StoreUpdatePostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePostSwitchParameter;
  constructor(public payload: any) {}
}

export class SwitchPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.SwitchPostSwitchParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPostSwitchParameter;
  constructor(public payload: any) {}
}

export class AddPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.AddPostSwitchParameter;
  constructor(public payload: any) {}
}

export class StoreAddPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddPostSwitchParameter;
  constructor(public payload: any) {}
}

export class DelPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.DelPostSwitchParameter;
  constructor(public payload: any) {}
}

export class StoreDelPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelPostSwitchParameter;
  constructor(public payload: any) {}
}

export class StoreGotPostSwitchError implements Action {
  readonly type = ConfigActionTypes.StoreGotPostSwitchError;
  constructor(public payload: any) {}
}

export class StoreDropNewPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPostSwitchParameter;
  constructor(public payload: any) {}
}

export class StoreNewPostSwitchParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewPostSwitchParameter;
  constructor(public payload: any) {}
}

export class UpdatePostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.UpdatePostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class StoreUpdatePostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class SwitchPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.SwitchPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class StoreSwitchPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class AddPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.AddPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class StoreAddPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.StoreAddPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class DelPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.DelPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class StoreDelPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.StoreDelPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class StoreDropNewPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class StoreNewPostSwitchDefaultPtime implements Action {
  readonly type = ConfigActionTypes.StoreNewPostSwitchDefaultPtime;
  constructor(public payload: any) {}
}

export class UpdatePostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.UpdatePostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class StoreUpdatePostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.StoreUpdatePostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class SwitchPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.SwitchPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class StoreSwitchPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.StoreSwitchPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class AddPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.AddPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class StoreAddPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.StoreAddPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class DelPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.DelPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class StoreDelPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.StoreDelPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class StoreDropNewPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.StoreDropNewPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export class StoreNewPostSwitchCliKeybinding implements Action {
  readonly type = ConfigActionTypes.StoreNewPostSwitchCliKeybinding;
  constructor(public payload: any) {}
}

export type All =
  | GetPostSwitch
  | StoreGetPostSwitch
  | UpdatePostSwitchParameter
  | StoreUpdatePostSwitchParameter
  | SwitchPostSwitchParameter
  | StoreSwitchPostSwitchParameter
  | AddPostSwitchParameter
  | StoreAddPostSwitchParameter
  | DelPostSwitchParameter
  | StoreDelPostSwitchParameter
  | StoreGotPostSwitchError
  | StoreDropNewPostSwitchParameter
  | StoreNewPostSwitchParameter
  | UpdatePostSwitchCliKeybinding
  | StoreUpdatePostSwitchCliKeybinding
  | SwitchPostSwitchCliKeybinding
  | StoreSwitchPostSwitchCliKeybinding
  | AddPostSwitchCliKeybinding
  | StoreAddPostSwitchCliKeybinding
  | DelPostSwitchCliKeybinding
  | StoreDelPostSwitchCliKeybinding
  | StoreDropNewPostSwitchCliKeybinding
  | StoreNewPostSwitchCliKeybinding
  | UpdatePostSwitchDefaultPtime
  | StoreUpdatePostSwitchDefaultPtime
  | SwitchPostSwitchDefaultPtime
  | StoreSwitchPostSwitchDefaultPtime
  | AddPostSwitchDefaultPtime
  | StoreAddPostSwitchDefaultPtime
  | DelPostSwitchDefaultPtime
  | StoreDelPostSwitchDefaultPtime
  | StoreDropNewPostSwitchDefaultPtime
  | StoreNewPostSwitchDefaultPtime
;

