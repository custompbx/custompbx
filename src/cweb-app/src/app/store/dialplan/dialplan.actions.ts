import { Action } from '@ngrx/store';

export enum DialplanActionTypes {
  UPDATE_FAILURE = '[Dialplan] Failure',

  GET_CONTEXTS = '[Dialplan][Get] Contexts',
  STORE_GET_CONTEXTS = '[Dialplan]{Store}[Get] Contexts',
  ADD_CONTEXT = '[Dialplan][Add] Context',
  STORE_ADD_CONTEXT = '[Dialplan]{Store}[Add] Context',
  RENAME_CONTEXT = '[Dialplan][Rename] Context',
  STORE_RENAME_CONTEXT = '[Dialplan]{Store}[Rename] Context',
  DELETE_CONTEXT = '[Dialplan][Delete] Context',
  STORE_DELETE_CONTEXT = '[Dialplan]{Store}[Delete] Context',
  SWITCH_CONTEXT = '[Dialplan][Switch] Context',
  STORE_SWITCH_CONTEXT = '[Dialplan]{Store}[Switch] Context',

  GET_EXTENSIONS = '[Dialplan][Get] Extensions',
  STORE_GET_EXTENSIONS = '[Dialplan]{Store}[Get] Extensions',
  ADD_EXTENSION = '[Dialplan][Add] Extension',
  STORE_ADD_EXTENSION = '[Dialplan]{Store}[Add] Extension',
  RENAME_EXTENSION = '[Dialplan][Rename] Extension',
  STORE_RENAME_EXTENSION = '[Dialplan]{Store}[Rename] Extension',
  DELETE_EXTENSION = '[Dialplan][Delete] Extension',
  STORE_DELETE_EXTENSION = '[Dialplan]{Store}[Delete] Extension',
  SWITCH_EXTENSION = '[Dialplan][Switch] Extension',
  STORE_SWITCH_EXTENSION = '[Dialplan]{Store}[Switch] Extension',
  SWITCH_EXTENSION_CONTINUE = '[Dialplan][Switch] Extension Continue',
  STORE_SWITCH_EXTENSION_CONTINUE = '[Dialplan]{Store}[Switch] Extension Continue',
  UPDATE_EXTENSION = '[Dialplan][Update] Extension',
  STORE_UPDATE_EXTENSION = '[Dialplan]{Store}[Update] Extension',
  MOVE_EXTENSION = '[Dialplan][Move] Extension',
  STORE_MOVE_EXTENSION = '[Dialplan]{Store}[Move] Extension',

  GET_CONDITIONS = '[Dialplan][Get] Conditions',
  STORE_GET_CONDITIONS = '[Dialplan]{Store}[Get] Conditions',
  ADD_CONDITION = '[Dialplan][Add] Condition',
  STORE_ADD_CONDITION = '[Dialplan]{Store}[Add] Condition',
  DELETE_CONDITION = '[Dialplan][Delete] Condition',
  STORE_DELETE_CONDITION = '[Dialplan]{Store}[Delete] Condition',
  SWITCH_CONDITION = '[Dialplan][Switch] Condition',
  STORE_SWITCH_CONDITION = '[Dialplan]{Store}[Switch] Condition',
  UPDATE_CONDITION = '[Dialplan][Update] Condition',
  STORE_UPDATE_CONDITION = '[Dialplan]{Store}[Update] Condition',
  MOVE_CONDITION = '[Dialplan][Move] Condition',
  STORE_MOVE_CONDITION = '[Dialplan]{Store}[Move] Condition',

  GET_EXTENSION_DETAILS = '[Dialplan][Get] Extension details',
  STORE_EXTENSIONS_DETAILS = '[Dialplan]{Store}[Get] Extension details',

  ADD_ACTION = '[Dialplan][Add] Action',
  STORE_ADD_ACTION = '[Dialplan]{Store}[Add] Action',
  ADD_NEW_ACTION = '[Dialplan][Add] New action',
  DELETE_NEW_ACTION = '[Dialplan]{Store}[Add] Delete new action',
  DELETE_ACTION = '[Dialplan][Delete] Action',
  STORE_DELETE_ACTION = '[Dialplan]{Store}[Delete] Action',
  SWITCH_ACTION = '[Dialplan][Switch] Action',
  STORE_SWITCH_ACTION = '[Dialplan]{Store}[Switch] Action',
  UPDATE_ACTION = '[Dialplan][Update] Action',
  STORE_UPDATE_ACTION = '[Dialplan]{Store}[Update] Action',
  MOVE_ACTION = '[Dialplan][Move] Action',
  STORE_MOVE_ACTION = '[Dialplan]{Store}[Move] Action',

  ADD_REGEX = '[Dialplan][Add] Regex',
  STORE_ADD_REGEX = '[Dialplan]{Store}[Add] Regex',
  ADD_NEW_REGEX = '[Dialplan][Add] New regex',
  DELETE_NEW_REGEX = '[Dialplan]{Store}[Add] Delete new regex',
  DELETE_REGEX = '[Dialplan][Delete] Regex',
  STORE_DELETE_REGEX = '[Dialplan]{Store}[Delete] Regex',
  SWITCH_REGEX = '[Dialplan][Switch] Regex',
  STORE_SWITCH_REGEX = '[Dialplan]{Store}[Switch] Regex',
  UPDATE_REGEX = '[Dialplan][Update] Regex',
  STORE_UPDATE_REGEX = '[Dialplan]{Store}[Update] Regex',

  ADD_ANTIACTION = '[Dialplan][Add] Antiaction',
  STORE_ADD_ANTIACTION = '[Dialplan]{Store}[Add] Antiaction',
  ADD_NEW_ANTIACTION = '[Dialplan][Add] New antiaction',
  DELETE_NEW_ANTIACTION = '[Dialplan]{Store}[Add] Delete new antiaction',
  DELETE_ANTIACTION = '[Dialplan][Delete] Antiaction',
  STORE_DELETE_ANTIACTION = '[Dialplan]{Store}[Delete] Antiaction',
  SWITCH_ANTIACTION = '[Dialplan][Switch] Antiaction',
  STORE_SWITCH_ANTIACTION = '[Dialplan]{Store}[Switch] Antiaction',
  UPDATE_ANTIACTION = '[Dialplan][Update] Antiaction',
  STORE_UPDATE_ANTIACTION = '[Dialplan]{Store}[Update] Antiaction',
  MOVE_ANTIACTION = '[Dialplan][Move] Antiaction',
  STORE_MOVE_ANTIACTION = '[Dialplan]{Store}[Move] Antiaction',

  IMPORT_DIALPLAN = '[Dialplan][Import]',
  REDUCE_LOAD_COUNTER = '[Reduce][Load] Counter',

  DIALPLAN_DEBUG = '[Dialplan] Debug',
  STORE_DIALPLAN_DEBUG = '[Dialplan][Store] Debug',
  SWITCH_DIALPLAN_DEBUG = '[Dialplan][Switch] Debug',
  STORE_SWITCH_DIALPLAN_DEBUG = '[Dialplan][Switch][Store] Debug',
  STORE_CLEAR_DIALPLAN_DEBUG = '[Dialplan][Clear][Store] Debug',

  DIALPLAN_SETTINGS = 'DialplanGetSettings',
  STORE_DIALPLAN_SETTINGS = 'StoreDialplanGetSettings',
  SWITCH_DIALPLAN_STATIC = 'DialplanChangeNotProceed',
  STORE_SWITCH_DIALPLAN_STATIC = 'StoreDialplanChangeNotProceed',
}

export class Failure implements Action {
  readonly type = DialplanActionTypes.UPDATE_FAILURE;
  constructor(public payload: any) {}
}

export class ReduceLoadCounter implements Action {
  readonly type = DialplanActionTypes.REDUCE_LOAD_COUNTER;
  // constructor(public payload: any) {}
}

export class ImportDialplan implements Action {
  readonly type = DialplanActionTypes.IMPORT_DIALPLAN;
  constructor(public payload: any) {}
}

export class GetContexts implements Action {
  readonly type = DialplanActionTypes.GET_CONTEXTS;
  constructor(public payload: any) {}
}

export class StoreGetContexts implements Action {
  readonly type = DialplanActionTypes.STORE_GET_CONTEXTS;
  constructor(public payload: any) {}
}

export class AddContext implements Action {
  readonly type = DialplanActionTypes.ADD_CONTEXT;
  constructor(public payload: any) {}
}

export class StoreAddContext implements Action {
  readonly type = DialplanActionTypes.STORE_ADD_CONTEXT;
  constructor(public payload: any) {}
}

export class RenameContext implements Action {
  readonly type = DialplanActionTypes.RENAME_CONTEXT;
  constructor(public payload: any) {}
}

export class StoreRenameContext implements Action {
  readonly type = DialplanActionTypes.STORE_RENAME_CONTEXT;
  constructor(public payload: any) {}
}

export class DeleteContext implements Action {
  readonly type = DialplanActionTypes.DELETE_CONTEXT;
  constructor(public payload: any) {}
}

export class StoreDeleteContext implements Action {
  readonly type = DialplanActionTypes.STORE_DELETE_CONTEXT;
  constructor(public payload: any) {}
}

export class SwitchContext implements Action {
  readonly type = DialplanActionTypes.SWITCH_CONTEXT;
  constructor(public payload: any) {}
}

export class StoreSwitchContext implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_CONTEXT;
  constructor(public payload: any) {}
}

export class GetExtensions implements Action {
  readonly type = DialplanActionTypes.GET_EXTENSIONS;
  constructor(public payload: any) {}
}

export class StoreGetExtensions implements Action {
  readonly type = DialplanActionTypes.STORE_GET_EXTENSIONS;
  constructor(public payload: any) {}
}

export class AddExtension implements Action {
  readonly type = DialplanActionTypes.ADD_EXTENSION;
  constructor(public payload: any) {}
}

export class StoreAddExtension implements Action {
  readonly type = DialplanActionTypes.STORE_ADD_EXTENSION;
  constructor(public payload: any) {}
}

export class RenameExtension implements Action {
  readonly type = DialplanActionTypes.RENAME_EXTENSION;
  constructor(public payload: any) {}
}

export class StoreRenameExtension implements Action {
  readonly type = DialplanActionTypes.STORE_RENAME_EXTENSION;
  constructor(public payload: any) {}
}

export class DeleteExtension implements Action {
  readonly type = DialplanActionTypes.DELETE_EXTENSION;
  constructor(public payload: any) {}
}

export class StoreDeleteExtension implements Action {
  readonly type = DialplanActionTypes.STORE_DELETE_EXTENSION;
  constructor(public payload: any) {}
}

export class SwitchExtension implements Action {
  readonly type = DialplanActionTypes.SWITCH_EXTENSION;
  constructor(public payload: any) {}
}

export class StoreSwitchExtension implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_EXTENSION;
  constructor(public payload: any) {}
}

export class SwitchExtensionContinue implements Action {
  readonly type = DialplanActionTypes.SWITCH_EXTENSION_CONTINUE;
  constructor(public payload: any) {}
}

export class StoreSwitchExtensionContinue implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_EXTENSION_CONTINUE;
  constructor(public payload: any) {}
}

export class GetConditions implements Action {
  readonly type = DialplanActionTypes.GET_CONDITIONS;
  constructor(public payload: any) {}
}

export class StoreGetConditions implements Action {
  readonly type = DialplanActionTypes.STORE_GET_CONDITIONS;
  constructor(public payload: any) {}
}

export class AddCondition implements Action {
  readonly type = DialplanActionTypes.ADD_CONDITION;
  constructor(public payload: any) {}
}

export class StoreAddCondition implements Action {
  readonly type = DialplanActionTypes.STORE_ADD_CONDITION;
  constructor(public payload: any) {}
}

export class DeleteCondition implements Action {
  readonly type = DialplanActionTypes.DELETE_CONDITION;
  constructor(public payload: any) {}
}

export class StoreDeleteCondition implements Action {
  readonly type = DialplanActionTypes.STORE_DELETE_CONDITION;
  constructor(public payload: any) {}
}

export class SwitchCondition implements Action {
  readonly type = DialplanActionTypes.SWITCH_CONDITION;
  constructor(public payload: any) {}
}

export class StoreSwitchCondition implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_CONDITION;
  constructor(public payload: any) {}
}

export class GetExtensionDetails implements Action {
  readonly type = DialplanActionTypes.GET_EXTENSION_DETAILS;
  constructor(public payload: any) {}
}

export class StoreGetExtensionDetails implements Action {
  readonly type = DialplanActionTypes.STORE_EXTENSIONS_DETAILS;
  constructor(public payload: any) {}
}

export class AddAction implements Action {
  readonly type = DialplanActionTypes.ADD_ACTION;
  constructor(public payload: any) {}
}

export class StoreAddAction implements Action {
  readonly type = DialplanActionTypes.STORE_ADD_ACTION;
  constructor(public payload: any) {}
}

export class DeleteAction implements Action {
  readonly type = DialplanActionTypes.DELETE_ACTION;
  constructor(public payload: any) {}
}

export class StoreDeleteAction implements Action {
  readonly type = DialplanActionTypes.STORE_DELETE_ACTION;
  constructor(public payload: any) {}
}

export class SwitchAction implements Action {
  readonly type = DialplanActionTypes.SWITCH_ACTION;
  constructor(public payload: any) {}
}

export class StoreSwitchAction implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_ACTION;
  constructor(public payload: any) {}
}

export class AddAntiaction implements Action {
  readonly type = DialplanActionTypes.ADD_ANTIACTION;
  constructor(public payload: any) {}
}

export class StoreAddAntiaction implements Action {
  readonly type = DialplanActionTypes.STORE_ADD_ANTIACTION;
  constructor(public payload: any) {}
}

export class DeleteAntiaction implements Action {
  readonly type = DialplanActionTypes.DELETE_ANTIACTION;
  constructor(public payload: any) {}
}

export class StoreDeleteAntiaction implements Action {
  readonly type = DialplanActionTypes.STORE_DELETE_ANTIACTION;
  constructor(public payload: any) {}
}

export class SwitchAntiaction implements Action {
  readonly type = DialplanActionTypes.SWITCH_ANTIACTION;
  constructor(public payload: any) {}
}

export class StoreSwitchAntiaction implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_ANTIACTION;
  constructor(public payload: any) {}
}

export class UpdateAction implements Action {
  readonly type = DialplanActionTypes.UPDATE_ACTION;
  constructor(public payload: any) {}
}

export class StoreUpdateAction implements Action {
  readonly type = DialplanActionTypes.STORE_UPDATE_ACTION;
  constructor(public payload: any) {}
}

export class UpdateAntiaction implements Action {
  readonly type = DialplanActionTypes.UPDATE_ANTIACTION;
  constructor(public payload: any) {}
}

export class StoreUpdateAntiaction implements Action {
  readonly type = DialplanActionTypes.STORE_UPDATE_ANTIACTION;
  constructor(public payload: any) {}
}

export class UpdateExtension implements Action {
  readonly type = DialplanActionTypes.UPDATE_EXTENSION;
  constructor(public payload: any) {}
}

export class StoreUpdateExtension implements Action {
  readonly type = DialplanActionTypes.STORE_UPDATE_EXTENSION;
  constructor(public payload: any) {}
}

export class MoveExtension implements Action {
  readonly type = DialplanActionTypes.MOVE_EXTENSION;
  constructor(public payload: any) {}
}

export class StoreMoveExtension implements Action {
  readonly type = DialplanActionTypes.STORE_MOVE_EXTENSION;
  constructor(public payload: any) {}
}

export class MoveCondition implements Action {
  readonly type = DialplanActionTypes.MOVE_CONDITION;
  constructor(public payload: any) {}
}

export class StoreMoveCondition implements Action {
  readonly type = DialplanActionTypes.STORE_MOVE_CONDITION;
  constructor(public payload: any) {}
}

export class MoveAction implements Action {
  readonly type = DialplanActionTypes.MOVE_ACTION;
  constructor(public payload: any) {}
}

export class StoreMoveAction implements Action {
  readonly type = DialplanActionTypes.STORE_MOVE_ACTION;
  constructor(public payload: any) {}
}

export class MoveAntiaction implements Action {
  readonly type = DialplanActionTypes.MOVE_ANTIACTION;
  constructor(public payload: any) {}
}

export class StoreMoveAntiaction implements Action {
  readonly type = DialplanActionTypes.STORE_MOVE_ANTIACTION;
  constructor(public payload: any) {}
}

export class UpdateCondition implements Action {
  readonly type = DialplanActionTypes.UPDATE_CONDITION;
  constructor(public payload: any) {}
}

export class StoreUpdateCondition implements Action {
  readonly type = DialplanActionTypes.STORE_UPDATE_CONDITION;
  constructor(public payload: any) {}
}

export class AddRegex implements Action {
  readonly type = DialplanActionTypes.ADD_REGEX;
  constructor(public payload: any) {}
}

export class StoreAddRegex implements Action {
  readonly type = DialplanActionTypes.STORE_ADD_REGEX;
  constructor(public payload: any) {}
}

export class DeleteRegex implements Action {
  readonly type = DialplanActionTypes.DELETE_REGEX;
  constructor(public payload: any) {}
}

export class StoreDeleteRegex implements Action {
  readonly type = DialplanActionTypes.STORE_DELETE_REGEX;
  constructor(public payload: any) {}
}

export class UpdateRegex implements Action {
  readonly type = DialplanActionTypes.UPDATE_REGEX;
  constructor(public payload: any) {}
}

export class StoreUpdateRegex implements Action {
  readonly type = DialplanActionTypes.STORE_UPDATE_REGEX;
  constructor(public payload: any) {}
}

export class SwitchRegex implements Action {
  readonly type = DialplanActionTypes.SWITCH_REGEX;
  constructor(public payload: any) {}
}

export class StoreSwitchRegex implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_REGEX;
  constructor(public payload: any) {}
}

export class AddNewAction implements Action {
  readonly type = DialplanActionTypes.ADD_NEW_ACTION;
  constructor(public payload: any) {}
}

export class DeleteNewAction implements Action {
  readonly type = DialplanActionTypes.DELETE_NEW_ACTION;
  constructor(public payload: any) {}
}

export class AddNewAntiaction implements Action {
  readonly type = DialplanActionTypes.ADD_NEW_ANTIACTION;
  constructor(public payload: any) {}
}

export class DeleteNewAntiaction implements Action {
  readonly type = DialplanActionTypes.DELETE_NEW_ANTIACTION;
  constructor(public payload: any) {}
}

export class AddNewRegex implements Action {
  readonly type = DialplanActionTypes.ADD_NEW_REGEX;
  constructor(public payload: any) {}
}

export class DeleteNewRegex implements Action {
  readonly type = DialplanActionTypes.DELETE_NEW_REGEX;
  constructor(public payload: any) {}
}

export class DialplanDebug implements Action {
  readonly type = DialplanActionTypes.DIALPLAN_DEBUG;
  constructor(public payload: any) {}
}

export class StoreDialplanDebug implements Action {
  readonly type = DialplanActionTypes.STORE_DIALPLAN_DEBUG;
  constructor(public payload: any) {}
}

export class DialplanSettings implements Action {
  readonly type = DialplanActionTypes.DIALPLAN_SETTINGS;
  constructor(public payload: any) {}
}

export class StoreDialplanSettings implements Action {
  readonly type = DialplanActionTypes.STORE_DIALPLAN_SETTINGS;
  constructor(public payload: any) {}
}

export class SwitchDialplanStatic implements Action {
  readonly type = DialplanActionTypes.SWITCH_DIALPLAN_STATIC;
  constructor(public payload: any) {}
}

export class StoreSwitchDialplanStatic implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_DIALPLAN_STATIC;
  constructor(public payload: any) {}
}

export class SwitchDialplanDebug implements Action {
  readonly type = DialplanActionTypes.SWITCH_DIALPLAN_DEBUG;
  constructor(public payload: any) {}
}

export class StoreSwitchDialplanDebug implements Action {
  readonly type = DialplanActionTypes.STORE_SWITCH_DIALPLAN_DEBUG;
  constructor(public payload: any) {}
}

export class StoreClearDialplanDebug implements Action {
  readonly type = DialplanActionTypes.STORE_CLEAR_DIALPLAN_DEBUG;
  constructor(public payload: any) {}
}

export type All =
  | Failure
  | ReduceLoadCounter
  | ImportDialplan
  | GetContexts
  | StoreGetContexts
  | AddContext
  | StoreAddContext
  | RenameContext
  | StoreRenameContext
  | DeleteContext
  | StoreDeleteContext
  | SwitchContext
  | StoreSwitchContext
  | GetExtensions
  | StoreGetExtensions
  | AddExtension
  | StoreAddExtension
  | RenameExtension
  | StoreRenameExtension
  | DeleteExtension
  | StoreDeleteExtension
  | SwitchExtension
  | StoreSwitchExtension
  | SwitchExtensionContinue
  | StoreSwitchExtensionContinue
  | GetConditions
  | StoreGetConditions
  | AddCondition
  | StoreAddCondition
  | DeleteCondition
  | StoreDeleteCondition
  | SwitchCondition
  | StoreSwitchCondition
  | GetExtensionDetails
  | StoreGetExtensionDetails
  | AddAction
  | StoreAddAction
  | DeleteAction
  | StoreDeleteAction
  | SwitchAction
  | StoreSwitchAction
  | AddAntiaction
  | StoreAddAntiaction
  | DeleteAntiaction
  | StoreDeleteAntiaction
  | SwitchAntiaction
  | StoreSwitchAntiaction
  | UpdateAction
  | StoreUpdateAction
  | UpdateAntiaction
  | StoreUpdateAntiaction
  | UpdateExtension
  | StoreUpdateExtension
  | UpdateCondition
  | StoreUpdateCondition
  | MoveExtension
  | StoreMoveExtension
  | MoveCondition
  | StoreMoveCondition
  | MoveAction
  | StoreMoveAction
  | MoveAntiaction
  | StoreMoveAntiaction
  | AddRegex
  | StoreAddRegex
  | DeleteRegex
  | StoreDeleteRegex
  | UpdateRegex
  | StoreUpdateRegex
  | SwitchRegex
  | StoreSwitchRegex
  | AddNewAction
  | DeleteNewAction
  | AddNewAntiaction
  | DeleteNewAntiaction
  | AddNewRegex
  | DeleteNewRegex
  | DialplanDebug
  | StoreDialplanDebug
  | SwitchDialplanDebug
  | StoreSwitchDialplanDebug
  | StoreClearDialplanDebug
  | DialplanSettings
  | StoreDialplanSettings
  | SwitchDialplanStatic
  | StoreSwitchDialplanStatic
  ;
