import {Action} from '@ngrx/store';

export type OperationKind =
  | 'add'
  | 'update'
  | 'delete'
  | 'switch'
  | 'paste'
  | 'import'
  | 'truncate'
  | 'rename'
  | 'save'
  | 'move'
  | 'load'
  | 'unload'
  | 'reload'
  | 'autoload'
  | 'clear'
  | 'from-scratch'
  | 'send';

export interface OperationFeedbackMetadata {
  kind: OperationKind;
  sourceType: string;
}

export interface OperationFeedbackAction extends Action {
  operationFeedback?: OperationFeedbackMetadata | false;
}

const operationAliases: Readonly<Record<string, OperationKind>> = {
  add: 'add',
  create: 'add',
  update: 'update',
  updated: 'update',
  delete: 'delete',
  remove: 'delete',
  del: 'delete',
  drop: 'delete',
  switch: 'switch',
  paste: 'paste',
  import: 'import',
  truncate: 'truncate',
  rename: 'rename',
  save: 'save',
  move: 'move',
  load: 'load',
  unload: 'unload',
  reload: 'reload',
  autoload: 'autoload',
  clear: 'clear',
  'from scratch': 'from-scratch',
  send: 'send',
};

const localOnlyTokens = new Set(['new', 'reset']);

/**
 * Reads the operation opcode from the action naming convention. Only complete
 * bracket tokens or the verb immediately following the optional Store prefix
 * are considered, so entity names such as PostSwitch cannot become writes.
 */
export function operationKindFromType(type: string): OperationKind | null {
  const tokens = Array.from(type.matchAll(/[\[{]([^\]}]+)[\]}]/g), match => match[1].trim().toLowerCase());
  if (tokens.length) {
    if (tokens.some(token => token === 'get' || localOnlyTokens.has(token))) return null;
    for (const token of tokens) {
      const operation = operationAliases[token];
      if (operation) return operation;
    }
    return null;
  }

  const legacy = type.match(/^(?:Store)?(FromScratch|Updated|Update|Create|Add|Delete|Remove|Del|Drop|Switch|Paste|Import|Truncate|Rename|Save|Move|Unload|Reload|Load|Autoload|Clear|Send)/);
  if (!legacy) return null;

  const token = legacy[1].replace(/([a-z])([A-Z])/g, '$1 $2').toLowerCase();
  return operationAliases[token] ?? null;
}

export function operationMetadata(action: OperationFeedbackAction): OperationFeedbackMetadata | null {
  if (action.operationFeedback === false) return null;
  if (action.operationFeedback) return action.operationFeedback;

  const kind = operationKindFromType(action.type);
  return kind ? {kind, sourceType: action.type} : null;
}

export function withOperationFeedback<T extends Action>(action: T, sourceType: string): T & OperationFeedbackAction {
  const kind = operationKindFromType(sourceType);
  return kind ? {...action, operationFeedback: {kind, sourceType}} : action;
}

export function operationSuccessKey(kind: OperationKind): string {
  switch (kind) {
    case 'add': return 'feedback.itemAdded';
    case 'delete': return 'feedback.itemRemoved';
    case 'switch': return 'feedback.statusUpdated';
    case 'import': return 'feedback.importCompleted';
    case 'paste': return 'feedback.itemsPasted';
    case 'rename': return 'feedback.itemRenamed';
    case 'move': return 'feedback.orderUpdated';
    case 'reload': return 'feedback.reloadCompleted';
    case 'unload': return 'feedback.moduleUnloaded';
    case 'load': return 'feedback.moduleLoaded';
    case 'autoload': return 'feedback.autoloadUpdated';
    case 'clear': return 'feedback.itemCleared';
    case 'truncate': return 'feedback.configurationCleared';
    case 'from-scratch': return 'feedback.configurationCreated';
    case 'send': return 'feedback.commandCompleted';
    default: return 'feedback.changesSaved';
  }
}
