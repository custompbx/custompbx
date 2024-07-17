import { Action } from '@ngrx/store';
import {createActionHelper} from '../../../services/rxjs-helper/actions-helper';

export enum ConfigActionTypes {
  StoreGotConferenceError = 'StoreGotConferenceError',
  GetConference = 'GetConference',
  StoreGetConference = 'StoreGetConference',
  UpdateConferenceRoom = 'UpdateConferenceRoom',
  StoreUpdateConferenceRoom = 'StoreUpdateConferenceRoom',
  SwitchConferenceRoom = 'SwitchConferenceRoom',
  StoreSwitchConferenceRoom = 'StoreSwitchConferenceRoom',
  AddConferenceRoom = 'AddConferenceRoom',
  StoreAddConferenceRoom = 'StoreAddConferenceRoom',
  DelConferenceRoom = 'DelConferenceRoom',
  StoreDelConferenceRoom = 'StoreDelConferenceRoom',
  StoreNewConferenceRoom = 'StoreNewConferenceRoom',
  StoreDropNewConferenceRoom = 'StoreDropNewConferenceRoom',

  GetConferenceCallerControls = 'GetConferenceCallerControls',
  StoreGetConferenceCallerControls = 'StoreGetConferenceCallerControls',
  AddConferenceCallerControl = 'AddConferenceCallerControl',
  StoreAddConferenceCallerControl = 'StoreAddConferenceCallerControl',
  UpdateConferenceCallerControl = 'UpdateConferenceCallerControl',
  StoreUpdateConferenceCallerControl = 'StoreUpdateConferenceCallerControl',
  SwitchConferenceCallerControl = 'SwitchConferenceCallerControl',
  StoreSwitchConferenceCallerControl = 'StoreSwitchConferenceCallerControl',
  DelConferenceCallerControl = 'DelConferenceCallerControl',
  StoreDelConferenceCallerControl = 'StoreDelConferenceCallerControl',
  StoreNewConferenceCallerControl = 'StoreNewConferenceCallerControl',
  StoreDropNewConferenceCallerControl = 'StoreDropNewConferenceCallerControl',
  StorePasteConferenceCallerControls = 'StorePasteConferenceCallerControls',
  AddConferenceCallerControlGroup = 'AddConferenceCallerControlGroup',
  StoreAddConferenceCallerControlGroup = 'StoreAddConferenceCallerControlGroup',
  DelConferenceCallerControlGroup = 'DelConferenceCallerControlGroup',
  StoreDelConferenceCallerControlGroup = 'StoreDelConferenceCallerControlGroup',
  UpdateConferenceCallerControlGroup = 'UpdateConferenceCallerControlGroup',
  StoreUpdateConferenceCallerControlGroup = 'StoreUpdateConferenceCallerControlGroup',

  GetConferenceProfileParameters = 'GetConferenceProfileParameters',
  StoreGetConferenceProfileParameters = 'StoreGetConferenceProfileParameters',
  AddConferenceProfileParameter = 'AddConferenceProfileParameter',
  StoreAddConferenceProfileParameter = 'StoreAddConferenceProfileParameter',
  UpdateConferenceProfileParameter = 'UpdateConferenceProfileParameter',
  StoreUpdateConferenceProfileParameter = 'StoreUpdateConferenceProfileParameter',
  SwitchConferenceProfileParameter = 'SwitchConferenceProfileParameter',
  StoreSwitchConferenceProfileParameter = 'StoreSwitchConferenceProfileParameter',
  DelConferenceProfileParameter = 'DelConferenceProfileParameter',
  StoreDelConferenceProfileParameter = 'StoreDelConferenceProfileParameter',
  StoreNewConferenceProfileParameter = 'StoreNewConferenceProfileParameter',
  StoreDropNewConferenceProfileParameter = 'StoreDropNewConferenceProfileParameter',
  StorePasteConferenceProfileParameters = 'StorePasteConferenceProfileParameters',
  AddConferenceProfile = 'AddConferenceProfile',
  StoreAddConferenceProfile = 'StoreAddConferenceProfile',
  DelConferenceProfile = 'DelConferenceProfile',
  StoreDelConferenceProfile = 'StoreDelConferenceProfile',
  UpdateConferenceProfile = 'UpdateConferenceProfile',
  StoreUpdateConferenceProfile = 'StoreUpdateConferenceProfile',

  GetConferenceChatPermissionUsers = 'GetConferenceChatPermissionUsers',
  StoreGetConferenceChatPermissionUsers = 'StoreGetConferenceChatPermissionUsers',
  AddConferenceChatPermissionUser = 'AddConferenceChatPermissionUser',
  StoreAddConferenceChatPermissionUser = 'StoreAddConferenceChatPermissionUser',
  UpdateConferenceChatPermissionUser = 'UpdateConferenceChatPermissionUser',
  StoreUpdateConferenceChatPermissionUser = 'StoreUpdateConferenceChatPermissionUser',
  SwitchConferenceChatPermissionUser = 'SwitchConferenceChatPermissionUser',
  StoreSwitchConferenceChatPermissionUser = 'StoreSwitchConferenceChatPermissionUser',
  DelConferenceChatPermissionUser = 'DelConferenceChatPermissionUser',
  StoreDelConferenceChatPermissionUser = 'StoreDelConferenceChatPermissionUser',
  StoreNewConferenceChatPermissionUser = 'StoreNewConferenceChatPermissionUser',
  StoreDropNewConferenceChatPermissionUser = 'StoreDropNewConferenceChatPermissionUser',
  StorePasteConferenceChatPermissionUsers = 'StorePasteConferenceChatPermissionUsers',
  AddConferenceChatPermission = 'AddConferenceChatPermission',
  StoreAddConferenceChatPermission = 'StoreAddConferenceChatPermission',
  DelConferenceChatPermission = 'DelConferenceChatPermission',
  StoreDelConferenceChatPermission = 'StoreDelConferenceChatPermission',
  UpdateConferenceChatPermission = 'UpdateConferenceChatPermission',
  StoreUpdateConferenceChatPermission = 'StoreUpdateConferenceChatPermission',
}

export class GetConference implements Action {
  readonly type = ConfigActionTypes.GetConference;
  constructor(public payload: any) {}
}

export class StoreGetConference implements Action {
  readonly type = ConfigActionTypes.StoreGetConference;
  constructor(public payload: any) {}
}

export class UpdateConferenceRoom implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceRoom;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceRoom implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceRoom;
  constructor(public payload: any) {}
}

export class SwitchConferenceRoom implements Action {
  readonly type = ConfigActionTypes.SwitchConferenceRoom;
  constructor(public payload: any) {}
}

export class StoreSwitchConferenceRoom implements Action {
  readonly type = ConfigActionTypes.StoreSwitchConferenceRoom;
  constructor(public payload: any) {}
}

export class AddConferenceRoom implements Action {
  readonly type = ConfigActionTypes.AddConferenceRoom;
  constructor(public payload: any) {}
}

export class StoreAddConferenceRoom implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceRoom;
  constructor(public payload: any) {}
}

export class DelConferenceRoom implements Action {
  readonly type = ConfigActionTypes.DelConferenceRoom;
  constructor(public payload: any) {}
}

export class StoreDelConferenceRoom implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceRoom;
  constructor(public payload: any) {}
}

export class GetConferenceCallerControls implements Action {
  readonly type = ConfigActionTypes.GetConferenceCallerControls;
  constructor(public payload: any) {}
}

export class StoreGetConferenceCallerControls implements Action {
  readonly type = ConfigActionTypes.StoreGetConferenceCallerControls;
  constructor(public payload: any) {}
}

export class AddConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.AddConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StoreAddConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceCallerControl;
  constructor(public payload: any) {}
}

export class UpdateConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceCallerControl;
  constructor(public payload: any) {}
}

export class SwitchConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.SwitchConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StoreSwitchConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.StoreSwitchConferenceCallerControl;
  constructor(public payload: any) {}
}

export class DelConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.DelConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StoreDelConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StoreNewConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.StoreNewConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StoreDropNewConferenceCallerControl implements Action {
  readonly type = ConfigActionTypes.StoreDropNewConferenceCallerControl;
  constructor(public payload: any) {}
}

export class StorePasteConferenceCallerControls implements Action {
  readonly type = ConfigActionTypes.StorePasteConferenceCallerControls;
  constructor(public payload: any) {}
}

export class AddConferenceCallerControlGroup implements Action {
  readonly type = ConfigActionTypes.AddConferenceCallerControlGroup;
  constructor(public payload: any) {}
}

export class StoreAddConferenceCallerControlGroup implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceCallerControlGroup;
  constructor(public payload: any) {}
}

export class DelConferenceCallerControlGroup implements Action {
  readonly type = ConfigActionTypes.DelConferenceCallerControlGroup;
  constructor(public payload: any) {}
}

export class StoreDelConferenceCallerControlGroup implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceCallerControlGroup;
  constructor(public payload: any) {}
}

export class UpdateConferenceCallerControlGroup implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceCallerControlGroup;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceCallerControlGroup implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceCallerControlGroup;
  constructor(public payload: any) {}
}

export class GetConferenceProfileParameters implements Action {
  readonly type = ConfigActionTypes.GetConferenceProfileParameters;
  constructor(public payload: any) {}
}

export class StoreGetConferenceProfileParameters implements Action {
  readonly type = ConfigActionTypes.StoreGetConferenceProfileParameters;
  constructor(public payload: any) {}
}

export class AddConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.AddConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StoreAddConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class UpdateConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class SwitchConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.SwitchConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreSwitchConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class DelConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.DelConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDelConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StoreNewConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreNewConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StoreDropNewConferenceProfileParameter implements Action {
  readonly type = ConfigActionTypes.StoreDropNewConferenceProfileParameter;
  constructor(public payload: any) {}
}

export class StorePasteConferenceProfileParameters implements Action {
  readonly type = ConfigActionTypes.StorePasteConferenceProfileParameters;
  constructor(public payload: any) {}
}

export class AddConferenceProfile implements Action {
  readonly type = ConfigActionTypes.AddConferenceProfile;
  constructor(public payload: any) {}
}

export class StoreAddConferenceProfile implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceProfile;
  constructor(public payload: any) {}
}

export class DelConferenceProfile implements Action {
  readonly type = ConfigActionTypes.DelConferenceProfile;
  constructor(public payload: any) {}
}

export class StoreDelConferenceProfile implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceProfile;
  constructor(public payload: any) {}
}

export class UpdateConferenceProfile implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceProfile;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceProfile implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceProfile;
  constructor(public payload: any) {}
}

export class StoreGotConferenceError implements Action {
  readonly type = ConfigActionTypes.StoreGotConferenceError;
  constructor(public payload: any) {}
}

export class StoreDropNewConferenceRoom implements Action {
  readonly type = ConfigActionTypes.StoreDropNewConferenceRoom;
  constructor(public payload: any) {}
}

export class StoreNewConferenceRoom implements Action {
  readonly type = ConfigActionTypes.StoreNewConferenceRoom;
  constructor(public payload: any) {}
}

export class GetConferenceChatPermissionUsers implements Action {
  readonly type = ConfigActionTypes.GetConferenceChatPermissionUsers;
  constructor(public payload: any) {}
}

export class StoreGetConferenceChatPermissionUsers implements Action {
  readonly type = ConfigActionTypes.StoreGetConferenceChatPermissionUsers;
  constructor(public payload: any) {}
}

export class AddConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.AddConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StoreAddConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class UpdateConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class SwitchConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.SwitchConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StoreSwitchConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.StoreSwitchConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class DelConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.DelConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StoreDelConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StoreNewConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.StoreNewConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StoreDropNewConferenceChatPermissionUser implements Action {
  readonly type = ConfigActionTypes.StoreDropNewConferenceChatPermissionUser;
  constructor(public payload: any) {}
}

export class StorePasteConferenceChatPermissionUsers implements Action {
  readonly type = ConfigActionTypes.StorePasteConferenceChatPermissionUsers;
  constructor(public payload: any) {}
}

export class AddConferenceChatPermission implements Action {
  readonly type = ConfigActionTypes.AddConferenceChatPermission;
  constructor(public payload: any) {}
}

export class StoreAddConferenceChatPermission implements Action {
  readonly type = ConfigActionTypes.StoreAddConferenceChatPermission;
  constructor(public payload: any) {}
}

export class DelConferenceChatPermission implements Action {
  readonly type = ConfigActionTypes.DelConferenceChatPermission;
  constructor(public payload: any) {}
}

export class StoreDelConferenceChatPermission implements Action {
  readonly type = ConfigActionTypes.StoreDelConferenceChatPermission;
  constructor(public payload: any) {}
}

export class UpdateConferenceChatPermission implements Action {
  readonly type = ConfigActionTypes.UpdateConferenceChatPermission;
  constructor(public payload: any) {}
}

export class StoreUpdateConferenceChatPermission implements Action {
  readonly type = ConfigActionTypes.StoreUpdateConferenceChatPermission;
  constructor(public payload: any) {}
}

export const GetConferenceLayouts = createActionHelper('GetConferenceLayouts');
export const StoreGetConferenceLayouts = createActionHelper('StoreGetConferenceLayouts');
export const GetConferenceLayoutImages = createActionHelper('GetConferenceLayoutImages');
export const StoreGetConferenceLayoutImages = createActionHelper('StoreGetConferenceLayoutImages');
export const GetConferenceLayoutGroupLayouts = createActionHelper('GetConferenceLayoutGroupLayouts');
export const StoreGetConferenceLayoutGroupLayouts = createActionHelper('StoreGetConferenceLayoutGroupLayouts');

export const UpdateConferenceLayout = createActionHelper('UpdateConferenceLayout');
export const StoreUpdateConferenceLayout = createActionHelper('StoreUpdateConferenceLayout');
export const UpdateConferenceLayout3D = createActionHelper('UpdateConferenceLayout3D');
export const StoreUpdateConferenceLayout3D = createActionHelper('StoreUpdateConferenceLayout3D');
export const UpdateConferenceLayoutGroup = createActionHelper('UpdateConferenceLayoutGroup');
export const StoreUpdateConferenceLayoutGroup = createActionHelper('StoreUpdateConferenceLayoutGroup');
export const SwitchConferenceLayout = createActionHelper('SwitchConferenceLayout');
export const StoreSwitchConferenceLayout = createActionHelper('StoreSwitchConferenceLayout');
export const AddConferenceLayout = createActionHelper('AddConferenceLayout');
export const StoreAddConferenceLayout = createActionHelper('StoreAddConferenceLayout');
export const AddConferenceLayoutGroup = createActionHelper('AddConferenceLayoutGroup');
export const StoreAddConferenceLayoutGroup = createActionHelper('StoreAddConferenceLayoutGroup');
export const DelConferenceLayout = createActionHelper('DelConferenceLayout');
export const StoreDelConferenceLayout = createActionHelper('StoreDelConferenceLayout');
export const DelConferenceLayoutGroup = createActionHelper('DelConferenceLayoutGroup');
export const StoreDelConferenceLayoutGroup = createActionHelper('StoreDelConferenceLayoutGroup');
export const UpdateConferenceLayoutGroupLayout = createActionHelper('UpdateConferenceLayoutGroupLayout');
export const StoreUpdateConferenceLayoutGroupLayout = createActionHelper('StoreUpdateConferenceLayoutGroupLayout');
export const SwitchConferenceLayoutGroupLayout = createActionHelper('SwitchConferenceLayoutGroupLayout');
export const StoreSwitchConferenceLayoutGroupLayout = createActionHelper('StoreSwitchConferenceLayoutGroupLayout');
export const AddConferenceLayoutGroupLayout = createActionHelper('AddConferenceLayoutGroupLayout');
export const StoreAddConferenceLayoutGroupLayout = createActionHelper('StoreAddConferenceLayoutGroupLayout');
export const DelConferenceLayoutGroupLayout = createActionHelper('DelConferenceLayoutGroupLayout');
export const StoreDelConferenceLayoutGroupLayout = createActionHelper('StoreDelConferenceLayoutGroupLayout');
export const AddConferenceLayoutImage = createActionHelper('AddConferenceLayoutImage');
export const StoreAddConferenceLayoutImage = createActionHelper('StoreAddConferenceLayoutImage');
export const DelConferenceLayoutImage = createActionHelper('DelConferenceLayoutImage');
export const StoreDelConferenceLayoutImage = createActionHelper('StoreDelConferenceLayoutImage');
export const SwitchConferenceLayoutImage = createActionHelper('SwitchConferenceLayoutImage');
export const StoreSwitchConferenceLayoutImage = createActionHelper('StoreSwitchConferenceLayoutImage');
export const UpdateConferenceLayoutImage = createActionHelper('UpdateConferenceLayoutImage');
export const StoreUpdateConferenceLayoutImage = createActionHelper('StoreUpdateConferenceLayoutImage');
export const StoreNewConferenceLayoutImage = createActionHelper('StoreNewConferenceLayoutImage');
export const StoreNewConferenceLayoutGroupLayout = createActionHelper('StoreNewConferenceLayoutGroupLayout');
export const StoreDropConferenceLayoutImage = createActionHelper('StoreDropConferenceLayoutImage');
export const StoreDropConferenceLayoutGroupLayout = createActionHelper('StoreDropConferenceLayoutGroupLayout');
export const StorePasteConferenceLayoutImage = createActionHelper('StorePasteConferenceLayoutImage');
export const StorePasteConferenceLayoutGroupLayout = createActionHelper('StorePasteConferenceLayoutGroupLayout');

export const StoreConferenceError = createActionHelper('StoreConferenceError');

export type All =
  | GetConference
  | StoreGetConference
  | UpdateConferenceRoom
  | StoreUpdateConferenceRoom
  | SwitchConferenceRoom
  | StoreSwitchConferenceRoom
  | AddConferenceRoom
  | StoreAddConferenceRoom
  | DelConferenceRoom
  | StoreDelConferenceRoom
  | GetConferenceCallerControls
  | StoreGetConferenceCallerControls
  | AddConferenceCallerControl
  | StoreAddConferenceCallerControl
  | UpdateConferenceCallerControl
  | StoreUpdateConferenceCallerControl
  | SwitchConferenceCallerControl
  | StoreSwitchConferenceCallerControl
  | DelConferenceCallerControl
  | StoreDelConferenceCallerControl
  | StoreNewConferenceCallerControl
  | StoreDropNewConferenceCallerControl
  | StorePasteConferenceCallerControls
  | AddConferenceCallerControlGroup
  | StoreAddConferenceCallerControlGroup
  | DelConferenceCallerControlGroup
  | StoreDelConferenceCallerControlGroup
  | UpdateConferenceCallerControlGroup
  | StoreUpdateConferenceCallerControlGroup
  | StoreGotConferenceError
  | StoreDropNewConferenceRoom
  | StoreNewConferenceRoom
  | GetConferenceProfileParameters
  | StoreGetConferenceProfileParameters
  | AddConferenceProfileParameter
  | StoreAddConferenceProfileParameter
  | UpdateConferenceProfileParameter
  | StoreUpdateConferenceProfileParameter
  | SwitchConferenceProfileParameter
  | StoreSwitchConferenceProfileParameter
  | DelConferenceProfileParameter
  | StoreDelConferenceProfileParameter
  | StoreNewConferenceProfileParameter
  | StoreDropNewConferenceProfileParameter
  | StorePasteConferenceProfileParameters
  | AddConferenceProfile
  | StoreAddConferenceProfile
  | DelConferenceProfile
  | StoreDelConferenceProfile
  | UpdateConferenceProfile
  | StoreUpdateConferenceProfile
  | GetConferenceChatPermissionUsers
  | StoreGetConferenceChatPermissionUsers
  | AddConferenceChatPermissionUser
  | StoreAddConferenceChatPermissionUser
  | UpdateConferenceChatPermissionUser
  | StoreUpdateConferenceChatPermissionUser
  | SwitchConferenceChatPermissionUser
  | StoreSwitchConferenceChatPermissionUser
  | DelConferenceChatPermissionUser
  | StoreDelConferenceChatPermissionUser
  | StoreNewConferenceChatPermissionUser
  | StoreDropNewConferenceChatPermissionUser
  | StorePasteConferenceChatPermissionUsers
  | AddConferenceChatPermission
  | StoreAddConferenceChatPermission
  | DelConferenceChatPermission
  | StoreDelConferenceChatPermission
  | UpdateConferenceChatPermission
  | StoreUpdateConferenceChatPermission
;

