import {
  ConfigActionTypes,
  All,
} from './config.actions.acl';
import {
  initialState, Inode, Inodes,
  State
} from '../config.state.struct';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.MoveAclListNode:
    case ConfigActionTypes.ADD_ACL_LIST:
    case ConfigActionTypes.UPDATE_ACL_LIST:
    case ConfigActionTypes.UPDATE_ACL_LIST_DEFAULT:
    case ConfigActionTypes.DEL_ACL_LIST:
    case ConfigActionTypes.GET_ACL_NODES:
    case ConfigActionTypes.ADD_ACL_NODE:
    case ConfigActionTypes.UPDATE_ACL_NODE:
    case ConfigActionTypes.DEL_ACL_NODE:
    case ConfigActionTypes.SWITCH_ACL_NODE:
    case ConfigActionTypes.GET_ACL_LISTS: {
      return {
        ...state,
        acl: {
          ...state.acl,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1
      };
    }

    case ConfigActionTypes.StoreGotAclError: {
      return {
        ...state,
        acl: {
          ...state.acl,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ACL_LISTS: {
      if (action.payload.response.exists === false) {
        return {
          ...state,
          acl: {...state.acl, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }
      let data = action.payload.response.data;
      if (data && data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        acl: {
          ...state.acl,
          lists: {
            ...data
          },
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: 0,
      };
    }

    case ConfigActionTypes.STORE_UPDATED_ACL_LIST_DEFAULT:
    case ConfigActionTypes.STORE_UPDATED_ACL_LIST:
    case ConfigActionTypes.STORE_ACL_LIST: {
      let data = action.payload.response.data || {};
      if (!data.id) {
        const ids = Object.keys(data);
        if (ids.length === 0) {
          return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
        }
        data = data[ids[0]];
      }
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      if (state.acl.lists[data.id]) {
        data.nodes = state.acl.lists[data.id].nodes;
      }

      return {
        ...state,
        acl: {
          ...state.acl, lists: {...state.acl.lists, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.DROP_ACL_LIST: {
      const id = action.payload.response.data?.id;
      if (!state.acl.lists[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.acl.lists;

      return {
        ...state,
        acl: {
          ...state.acl,
          lists:
            {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreMoveAclListNode:
    case ConfigActionTypes.STORE_ACL_NODES: {
      let parent_id = action.payload.id;
      const data = action.payload.response.data || {};

      if (!parent_id && Object.keys(data).length > 0) {
        parent_id = data[Object.keys(data)[0]].parent?.id || 0;
      }

      if (!state.acl.lists[parent_id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        acl: {
          ...state.acl, lists: {...state.acl.lists, [parent_id]: {...state.acl.lists[parent_id], nodes: data}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_ACL_NODE:
    case ConfigActionTypes.STORE_UPDATED_ACL_NODE: {
      const data = action.payload.response.data || {};
      const parent_id = data.parent?.id || 0;
      if (!parent_id || !state.acl.lists[parent_id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        acl: {
          ...state.acl, lists: {
            ...state.acl.lists, [parent_id]: {
              ...state.acl.lists[parent_id], nodes: {
                ...state.acl.lists[parent_id].nodes,
                [data.id]: {...data}
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.DROP_ACL_NODE: {
      const data = action.payload.response.data || {};
      const parent_id = data.parent?.id || 0;
      if (!parent_id || !state.acl.lists[parent_id] || !state.acl.lists[parent_id].nodes[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.acl.lists[parent_id].nodes;

      return {
        ...state,
        acl: {
          ...state.acl, lists: {...state.acl.lists, [parent_id]: {...state.acl.lists[parent_id], nodes: <Inodes>{...rest}}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_ACL_NODE: {
      const id = action.payload;
      if (!state.acl.lists[id].nodes) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const rest = [
        ...state.acl.lists[id].nodes?.new || [],
        <Inode>{}
      ];

      return {
        ...state,
        acl: {
          ...state.acl, lists: {...state.acl.lists, [id]: {...state.acl.lists[id], nodes: {...state.acl.lists[id].nodes, new: rest}}},
          errorMessage: null
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.DROP_NEW_ACL_NODE: {
      const id = action.payload.id;
      if (!state.acl.lists[id].nodes || !state.acl.lists[id].nodes.new) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const rest = [
        ...state.acl.lists[id].nodes.new.slice(0, action.payload.index),
        null,
        ...state.acl.lists[id].nodes.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        acl: {
          ...state.acl, lists: {
            ...state.acl.lists, [id]: {
              ...state.acl.lists[id], nodes: {...state.acl.lists[id].nodes, new: rest}
            }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ACL_NODE: {
      let data = action.payload.response.data || {};
      let id: any = 0;
      if (!data.id) {
        const ids = Object.keys(data);
        if (ids.length === 0) {
          return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
        }
        data = data[ids[0]];
      }
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      if (!data.parent?.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      id = data.parent.id;
      data = {[id]: data};
      const rest = [
        ...state.acl.lists[id].nodes.new.slice(0, action.payload.index),
        null,
        ...state.acl.lists[id].nodes.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        acl: {
          ...state.acl,
          lists: {
            ...state.acl.lists,
            [id]: {...state.acl.lists[id], nodes: {...state.acl.lists[id].nodes, [data[id].id]: data[id], new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return null;
    }
  }
}
