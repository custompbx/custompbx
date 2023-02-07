import { Action } from '@ngrx/store';

export enum ConfigActionTypes {
  StoreGotAclError = 'StoreGotAclError',
  GET_ACL_LISTS = '[Config] Get_acl_lists',
  STORE_ACL_LISTS = '[Config] Store_acl_lists',
  ADD_ACL_LIST = '[Config] Add_acl_list',
  STORE_ACL_LIST = '[Config] Store_acl_list',
  UPDATE_ACL_LIST = '[Config] Update_acl_list',
  STORE_UPDATED_ACL_LIST = '[Config] Store_updated_acl_list',
  UPDATE_ACL_LIST_DEFAULT = '[Config] Update_acl_list_default',
  STORE_UPDATED_ACL_LIST_DEFAULT = '[Config] Store_updated_acl_list_default',
  DEL_ACL_LIST = '[Config] Del_acl_list',
  DROP_ACL_LIST = '[Config] Drop_acl_list',
  STORE_NEW_ACL_LIST = '[Config] Store_new_acl_list',
  DROP_NEW_ACL_LIST = '[Config] Drop_new_acl_list',
  GET_ACL_NODES = '[Config] Get_acl_nodes',
  STORE_ACL_NODES = '[Config] Store_acl_nodes',
  ADD_ACL_NODE = '[Config] Add_acl_node',
  STORE_ACL_NODE = '[Config] Store_acl_node',
  UPDATE_ACL_NODE = '[Config] Update_acl_node',
  STORE_UPDATED_ACL_NODE = '[Config] Store_updated_acl_node',
  DEL_ACL_NODE = '[Config] Del_acl_node',
  DROP_ACL_NODE = '[Config] Drop_acl_node',
  STORE_NEW_ACL_NODE = '[Config] Store_new_acl_node',
  DROP_NEW_ACL_NODE = '[Config] Drop_new_acl_node',
  SWITCH_ACL_NODE = '[Config] Switch_acl_node',
  STORE_SWITCH_ACL_NODE = '[Config] Store_switch_acl_node',
  MoveAclListNode = 'MoveAclListNode',
  StoreMoveAclListNode = 'StoreMoveAclListNode',

}

export class StoreMoveAclListNode implements Action {
  readonly type = ConfigActionTypes.StoreMoveAclListNode;
  constructor(public payload: any) {}
}

export class MoveAclListNode implements Action {
  readonly type = ConfigActionTypes.MoveAclListNode;
  constructor(public payload: any) {}
}

export class StoreGotAclError implements Action {
  readonly type = ConfigActionTypes.StoreGotAclError;
  constructor(public payload: any) {}
}

export class GetAclLists implements Action {
  readonly type = ConfigActionTypes.GET_ACL_LISTS;
  constructor(public payload: any) {}
}

export class StoreAclLists implements Action {
  readonly type = ConfigActionTypes.STORE_ACL_LISTS;
  constructor(public payload: any) {}
}

export class AddAclList implements Action {
  readonly type = ConfigActionTypes.ADD_ACL_LIST;
  constructor(public payload: any) {}
}

export class StoreAclList implements Action {
  readonly type = ConfigActionTypes.STORE_ACL_LIST;
  constructor(public payload: any) {}
}

export class UpdateAclList implements Action {
  readonly type = ConfigActionTypes.UPDATE_ACL_LIST;
  constructor(public payload: any) {}
}

export class StoreUpdatadAclList implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATED_ACL_LIST;
  constructor(public payload: any) {}
}

export class UpdateAclListDefault implements Action {
  readonly type = ConfigActionTypes.UPDATE_ACL_LIST_DEFAULT;
  constructor(public payload: any) {}
}

export class StoreUpdatadAclListDefault implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATED_ACL_LIST_DEFAULT;
  constructor(public payload: any) {}
}

export class DropAclList implements Action {
  readonly type = ConfigActionTypes.DROP_ACL_LIST;
  constructor(public payload: any) {}
}

export class DelAclList implements Action {
  readonly type = ConfigActionTypes.DEL_ACL_LIST;
  constructor(public payload: any) {}
}

export class DropNewAclList implements Action {
  readonly type = ConfigActionTypes.DROP_NEW_ACL_LIST;
  constructor(public payload: any) {}
}

export class StoreNewAclList implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_ACL_LIST;
  constructor(public payload: any) {}
}

// nodes
export class GetAclNodes implements Action {
  readonly type = ConfigActionTypes.GET_ACL_NODES;
  constructor(public payload: any) {}
}

export class StoreAclNodes implements Action {
  readonly type = ConfigActionTypes.STORE_ACL_NODES;
  constructor(public payload: any) {}
}

export class AddAclNode implements Action {
  readonly type = ConfigActionTypes.ADD_ACL_NODE;
  constructor(public payload: any) {}
}

export class StoreAclNode implements Action {
  readonly type = ConfigActionTypes.STORE_ACL_NODE;
  constructor(public payload: any) {}
}

export class UpdateAclNode implements Action {
  readonly type = ConfigActionTypes.UPDATE_ACL_NODE;
  constructor(public payload: any) {}
}

export class StoreUpdatadAclNode implements Action {
  readonly type = ConfigActionTypes.STORE_UPDATED_ACL_NODE;
  constructor(public payload: any) {}
}

export class DropAclNode implements Action {
  readonly type = ConfigActionTypes.DROP_ACL_NODE;
  constructor(public payload: any) {}
}

export class DelAclNode implements Action {
  readonly type = ConfigActionTypes.DEL_ACL_NODE;
  constructor(public payload: any) {}
}

export class DropNewAclNode implements Action {
  readonly type = ConfigActionTypes.DROP_NEW_ACL_NODE;
  constructor(public payload: any) {}
}

export class StoreNewAclNode implements Action {
  readonly type = ConfigActionTypes.STORE_NEW_ACL_NODE;
  constructor(public payload: any) {}
}

export class SwitchAclNode implements Action {
  readonly type = ConfigActionTypes.SWITCH_ACL_NODE;
  constructor(public payload: any) {}
}

export class StoreSwitchAclNode implements Action {
  readonly type = ConfigActionTypes.STORE_SWITCH_ACL_NODE;
  constructor(public payload: any) {}
}

export type All =
  | StoreGotAclError
  | GetAclLists
  | StoreAclLists
  | AddAclList
  | StoreAclList
  | UpdateAclList
  | StoreUpdatadAclList
  | UpdateAclListDefault
  | StoreUpdatadAclListDefault
  | DropAclList
  | DelAclList
  | StoreNewAclList
  | DropNewAclList
  | GetAclNodes
  | StoreAclNodes
  | AddAclNode
  | StoreAclNode
  | UpdateAclNode
  | StoreUpdatadAclNode
  | DropAclNode
  | DelAclNode
  | StoreNewAclNode
  | DropNewAclNode
  | SwitchAclNode
  | StoreSwitchAclNode
  | MoveAclListNode
  | StoreMoveAclListNode
;
