import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  GetCallcenterSettings = 'GetCallcenterSettings',
  StoreGetCallcenterSettings = 'StoreGetCallcenterSettings',
  UpdateCallcenterSettings = 'UpdateCallcenterSettings',
  StoreUpdateCallcenterSettings = 'StoreUpdateCallcenterSettings',
  SwitchCallcenterSettings = 'SwitchCallcenterSettings',
  StoreSwitchCallcenterSettings = 'StoreSwitchCallcenterSettings',
  AddCallcenterSettings = 'AddCallcenterSettings',
  StoreAddCallcenterSettings = 'StoreAddCallcenterSettings',
  DelCallcenterSettings = 'DelCallcenterSettings',
  StoreDelCallcenterSettings = 'StoreDelCallcenterSettings',
  StoreNewCallcenterSettings = 'StoreNewCallcenterSettings',
  StoreDropNewCallcenterSettings = 'StoreDropNewCallcenterSettings',

  GetCallcenterQueues = 'GetCallcenterQueues',
  StoreGetCallcenterQueues = 'StoreGetCallcenterQueues',
  GetCallcenterQueuesParams = 'GetCallcenterQueuesParams',
  StoreGetCallcenterQueuesParams = 'StoreGetCallcenterQueuesParams',
  UpdateCallcenterQueueParam = 'UpdateCallcenterQueueParam',
  StoreUpdateCallcenterQueueParam = 'StoreUpdateCallcenterQueueParam',
  SwitchCallcenterQueueParam = 'SwitchCallcenterQueueParam',
  StoreSwitchCallcenterQueueParam = 'StoreSwitchCallcenterQueueParam',
  AddCallcenterQueueParam = 'AddCallcenterQueueParam',
  StoreAddCallcenterQueueParam = 'StoreAddCallcenterQueueParam',
  DelCallcenterQueueParam = 'DelCallcenterQueueParam',
  StoreDelCallcenterQueueParam = 'StoreDelCallcenterQueueParam',
  StoreNewCallcenterQueueParam = 'StoreNewCallcenterQueueParam',
  StoreDropNewCallcenterQueueParam = 'StoreDropNewCallcenterQueueParam',
  StorePasteCallcenterQueueParams = 'StorePasteCallcenterQueueParams',
  AddCallcenterQueue = 'AddCallcenterQueue',
  StoreAddCallcenterQueue = 'StoreAddCallcenterQueue',
  DelCallcenterQueue = 'DelCallcenterQueue',
  StoreDelCallcenterQueue = 'StoreDelCallcenterQueue',
  RenameCallcenterQueue = 'RenameCallcenterQueue',
  StoreRenameCallcenterQueue = 'StoreRenameCallcenterQueue',

  GetCallcenterAgents = 'GetCallcenterAgents',
  StoreGetCallcenterAgents = 'StoreGetCallcenterAgents',
  UpdateCallcenterAgent = 'UpdateCallcenterAgent',
  StoreUpdateCallcenterAgent = 'StoreUpdateCallcenterAgent',
  AddCallcenterAgent = 'AddCallcenterAgent',
  StoreAddCallcenterAgent = 'StoreAddCallcenterAgent',
  DelCallcenterAgent = 'DelCallcenterAgent',
  StoreDelCallcenterAgent = 'StoreDelCallcenterAgent',
  ImportCallcenterAgentsAndTiers = 'ImportCallcenterAgentsAndTiers',
  StoreImportCallcenterAgentsAndTiers = 'StoreImportCallcenterAgentsAndTiers',

  GetCallcenterTiers = 'GetCallcenterTiers',
  StoreGetCallcenterTiers = 'StoreGetCallcenterTiers',
  UpdateCallcenterTier = 'UpdateCallcenterTier',
  StoreUpdateCallcenterTier = 'StoreUpdateCallcenterTier',
  AddCallcenterTier = 'AddCallcenterTier',
  StoreAddCallcenterTier = 'StoreAddCallcenterTier',
  DelCallcenterTier = 'DelCallcenterTier',
  StoreDelCallcenterTier = 'StoreDelCallcenterTier',

  SendCallcenterCommand = 'SendCallcenterCommand',
  StoreSendCallcenterCommand = 'StoreSendCallcenterCommand',
  SubscribeCallcenterAgents = 'SubscribeCallcenterAgents',
  StoreSubscribeCallcenterAgents = 'StoreSubscribeCallcenterAgents',
  SubscribeCallcenterTiers = 'SubscribeCallcenterTiers',
  StoreSubscribeCallcenterTiers = 'StoreSubscribeCallcenterTiers',

  GetCallcenterMembers = 'GetCallcenterMembers',
  StoreGetCallcenterMembers = 'StoreGetCallcenterMembers',
  DelCallcenterMember = 'DelCallcenterMember',
  StoreDelCallcenterMember = 'StoreDelCallcenterMember',

  StoreSetChangedCallcenterTableField = 'StoreSetChangedCallcenterTableField',
  StoreGotCallcenterError = 'StoreGotCallcenterError',

}

export class StoreSetChangedCallcenterTableField implements Action {
  readonly type = ConfigActionTypes.StoreSetChangedCallcenterTableField;
  constructor(public payload: any) {}
}

export class StoreGotCallcenterError implements Action {
  readonly type = ConfigActionTypes.StoreGotCallcenterError;
  constructor(public payload: any) {}
}

export class GetCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.GetCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreGetCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreGetCallcenterSettings;
  constructor(public payload: any) {}
}

export class UpdateCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.UpdateCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreUpdateCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCallcenterSettings;
  constructor(public payload: any) {}
}

export class SwitchCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.SwitchCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreSwitchCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreSwitchCallcenterSettings;
  constructor(public payload: any) {}
}

export class AddCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.AddCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreAddCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreAddCallcenterSettings;
  constructor(public payload: any) {}
}

export class DelCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.DelCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreDelCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreDelCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreNewCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreNewCallcenterSettings;
  constructor(public payload: any) {}
}

export class StoreDropNewCallcenterSettings implements Action {
  readonly type = ConfigActionTypes.StoreDropNewCallcenterSettings;
  constructor(public payload: any) {}
}

export class GetCallcenterQueues implements Action {
  readonly type = ConfigActionTypes.GetCallcenterQueues;
  constructor(public payload: any) {}
}

export class StoreGetCallcenterQueues implements Action {
  readonly type = ConfigActionTypes.StoreGetCallcenterQueues;
  constructor(public payload: any) {}
}

export class GetCallcenterQueuesParams implements Action {
  readonly type = ConfigActionTypes.GetCallcenterQueuesParams;
  constructor(public payload: any) {}
}

export class StoreGetCallcenterQueuesParams implements Action {
  readonly type = ConfigActionTypes.StoreGetCallcenterQueuesParams;
  constructor(public payload: any) {}
}

export class UpdateCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.UpdateCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StoreUpdateCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class SwitchCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.SwitchCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StoreSwitchCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.StoreSwitchCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class AddCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.AddCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StoreAddCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.StoreAddCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class DelCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.DelCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StoreDelCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.StoreDelCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StoreNewCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.StoreNewCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StoreDropNewCallcenterQueueParam implements Action {
  readonly type = ConfigActionTypes.StoreDropNewCallcenterQueueParam;
  constructor(public payload: any) {}
}

export class StorePasteCallcenterQueueParams implements Action {
  readonly type = ConfigActionTypes.StorePasteCallcenterQueueParams;
  constructor(public payload: any) {}
}

export class AddCallcenterQueue implements Action {
  readonly type = ConfigActionTypes.AddCallcenterQueue;
  constructor(public payload: any) {}
}

export class StoreAddCallcenterQueue implements Action {
  readonly type = ConfigActionTypes.StoreAddCallcenterQueue;
  constructor(public payload: any) {}
}

export class DelCallcenterQueue implements Action {
  readonly type = ConfigActionTypes.DelCallcenterQueue;
  constructor(public payload: any) {}
}

export class StoreDelCallcenterQueue implements Action {
  readonly type = ConfigActionTypes.StoreDelCallcenterQueue;
  constructor(public payload: any) {}
}

export class RenameCallcenterQueue implements Action {
  readonly type = ConfigActionTypes.RenameCallcenterQueue;
  constructor(public payload: any) {}
}

export class StoreRenameCallcenterQueue implements Action {
  readonly type = ConfigActionTypes.StoreRenameCallcenterQueue;
  constructor(public payload: any) {}
}

export class GetCallcenterAgents implements Action {
  readonly type = ConfigActionTypes.GetCallcenterAgents;
  constructor(public payload: any) {}
}

export class StoreGetCallcenterAgents implements Action {
  readonly type = ConfigActionTypes.StoreGetCallcenterAgents;
  constructor(public payload: any) {}
}

export class UpdateCallcenterAgent implements Action {
  readonly type = ConfigActionTypes.UpdateCallcenterAgent;
  constructor(public payload: any) {}
}

export class StoreUpdateCallcenterAgent implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCallcenterAgent;
  constructor(public payload: any) {}
}

export class AddCallcenterAgent implements Action {
  readonly type = ConfigActionTypes.AddCallcenterAgent;
  constructor(public payload: any) {}
}

export class StoreAddCallcenterAgent implements Action {
  readonly type = ConfigActionTypes.StoreAddCallcenterAgent;
  constructor(public payload: any) {}
}

export class DelCallcenterAgent implements Action {
  readonly type = ConfigActionTypes.DelCallcenterAgent;
  constructor(public payload: any) {}
}

export class StoreDelCallcenterAgent implements Action {
  readonly type = ConfigActionTypes.StoreDelCallcenterAgent;
  constructor(public payload: any) {}
}

export class ImportCallcenterAgentsAndTiers implements Action {
  readonly type = ConfigActionTypes.ImportCallcenterAgentsAndTiers;
  constructor(public payload: any) {}
}

export class StoreImportCallcenterAgentsAndTiers implements Action {
  readonly type = ConfigActionTypes.StoreImportCallcenterAgentsAndTiers;
  constructor(public payload: any) {}
}

export class GetCallcenterTiers implements Action {
  readonly type = ConfigActionTypes.GetCallcenterTiers;
  constructor(public payload: any) {}
}

export class StoreGetCallcenterTiers implements Action {
  readonly type = ConfigActionTypes.StoreGetCallcenterTiers;
  constructor(public payload: any) {}
}

export class UpdateCallcenterTier implements Action {
  readonly type = ConfigActionTypes.UpdateCallcenterTier;
  constructor(public payload: any) {}
}

export class StoreUpdateCallcenterTier implements Action {
  readonly type = ConfigActionTypes.StoreUpdateCallcenterTier;
  constructor(public payload: any) {}
}

export class AddCallcenterTier implements Action {
  readonly type = ConfigActionTypes.AddCallcenterTier;
  constructor(public payload: any) {}
}

export class StoreAddCallcenterTier implements Action {
  readonly type = ConfigActionTypes.StoreAddCallcenterTier;
  constructor(public payload: any) {}
}

export class DelCallcenterTier implements Action {
  readonly type = ConfigActionTypes.DelCallcenterTier;
  constructor(public payload: any) {}
}

export class StoreDelCallcenterTier implements Action {
  readonly type = ConfigActionTypes.StoreDelCallcenterTier;
  constructor(public payload: any) {}
}

export class SendCallcenterCommand implements Action {
  readonly type = ConfigActionTypes.SendCallcenterCommand;
  constructor(public payload: any) {}
}

export class StoreSendCallcenterCommand implements Action {
  readonly type = ConfigActionTypes.StoreSendCallcenterCommand;
  constructor(public payload: any) {}
}

export class SubscribeCallcenterAgents implements Action {
  readonly type = ConfigActionTypes.SubscribeCallcenterAgents;
  constructor(public payload: any) {}
}

export class StoreSubscribeCallcenterAgents implements Action {
  readonly type = ConfigActionTypes.StoreSubscribeCallcenterAgents;
  constructor(public payload: any) {}
}

export class SubscribeCallcenterTiers implements Action {
  readonly type = ConfigActionTypes.SubscribeCallcenterTiers;
  constructor(public payload: any) {}
}

export class StoreSubscribeCallcenterTiers implements Action {
  readonly type = ConfigActionTypes.StoreSubscribeCallcenterTiers;
  constructor(public payload: any) {}
}

export class GetCallcenterMembers implements Action {
  readonly type = ConfigActionTypes.GetCallcenterMembers;
  constructor(public payload: any) {}
}

export class StoreGetCallcenterMembers implements Action {
  readonly type = ConfigActionTypes.StoreGetCallcenterMembers;
  constructor(public payload: any) {}
}

export class DelCallcenterMember implements Action {
  readonly type = ConfigActionTypes.DelCallcenterMember;
  constructor(public payload: any) {}
}

export class StoreDelCallcenterMember implements Action {
  readonly type = ConfigActionTypes.StoreDelCallcenterMember;
  constructor(public payload: any) {}
}

export type All =
  | GetCallcenterSettings
  | StoreGetCallcenterSettings
  | UpdateCallcenterSettings
  | StoreUpdateCallcenterSettings
  | SwitchCallcenterSettings
  | StoreSwitchCallcenterSettings
  | AddCallcenterSettings
  | StoreAddCallcenterSettings
  | DelCallcenterSettings
  | StoreDelCallcenterSettings
  | StoreNewCallcenterSettings
  | StoreDropNewCallcenterSettings
  | GetCallcenterQueues
  | StoreGetCallcenterQueues
  | GetCallcenterQueuesParams
  | StoreGetCallcenterQueuesParams
  | UpdateCallcenterQueueParam
  | StoreUpdateCallcenterQueueParam
  | SwitchCallcenterQueueParam
  | StoreSwitchCallcenterQueueParam
  | AddCallcenterQueueParam
  | StoreAddCallcenterQueueParam
  | DelCallcenterQueueParam
  | StoreDelCallcenterQueueParam
  | StoreNewCallcenterQueueParam
  | StoreDropNewCallcenterQueueParam
  | StorePasteCallcenterQueueParams
  | AddCallcenterQueue
  | StoreAddCallcenterQueue
  | DelCallcenterQueue
  | StoreDelCallcenterQueue
  | RenameCallcenterQueue
  | StoreRenameCallcenterQueue
  | GetCallcenterAgents
  | StoreGetCallcenterAgents
  | UpdateCallcenterAgent
  | StoreUpdateCallcenterAgent
  | AddCallcenterAgent
  | StoreAddCallcenterAgent
  | DelCallcenterAgent
  | StoreDelCallcenterAgent
  | ImportCallcenterAgentsAndTiers
  | StoreImportCallcenterAgentsAndTiers
  | GetCallcenterTiers
  | StoreGetCallcenterTiers
  | UpdateCallcenterTier
  | StoreUpdateCallcenterTier
  | AddCallcenterTier
  | StoreAddCallcenterTier
  | DelCallcenterTier
  | StoreDelCallcenterTier
  | SendCallcenterCommand
  | StoreSendCallcenterCommand
  | SubscribeCallcenterAgents
  | StoreSubscribeCallcenterAgents
  | SubscribeCallcenterTiers
  | StoreSubscribeCallcenterTiers
  | GetCallcenterMembers
  | StoreGetCallcenterMembers
  | DelCallcenterMember
  | StoreDelCallcenterMember
  | StoreGotCallcenterError
  | StoreSetChangedCallcenterTableField
;

