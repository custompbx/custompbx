import {
  AddAclList,
  AddAclNode,
  DelAclList,
  DelAclNode,
  DropAclList, StoreDelAclNode,
  DropNewAclNode,
  GetAclLists,
  GetAclNodes,
  MoveAclListNode,
  StoreAclList,
  StoreAclLists,
  StoreAclNode,
  StoreAclNodes,
  StoreGotAclError,
  StoreMoveAclListNode,
  StoreNewAclNode,
  StoreSwitchAclNode,
  StoreUpdatedAclList,
  StoreUpdatedAclListDefault,
  StoreUpdatedAclNode,
  SwitchAclNode,
  UpdateAclList,
  UpdateAclListDefault,
  UpdateAclNode,
} from './config.actions.acl';
import {initialState, Inode, Inodes, State} from '../config.state.struct';

export function reducer(state = initialState, action): State {
  switch (action.type) {
    case MoveAclListNode.type:
    case AddAclList.type:
    case UpdateAclList.type:
    case UpdateAclListDefault.type:
    case DelAclList.type:
    case GetAclNodes.type:
    case AddAclNode.type:
    case UpdateAclNode.type:
    case DelAclNode.type:
    case SwitchAclNode.type:
    case GetAclLists.type: {
      return {
        ...state,
        acl: {
          ...state.acl,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1
      };
    }

    case StoreGotAclError.type: {
      return {
        ...state,
        acl: {
          ...state.acl,
          errorMessage: action.payload.error || null,
        },
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case StoreAclLists.type: {
      const { response } = action.payload;
      let { data, error, exists } = response;
      if (exists === false) {
        return {
          ...state,
          acl: {...state.acl, exists: exists},
          loadCounter: Math.max(0, state.loadCounter - 1)
        };
      }
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
          exists: exists,
          errorMessage: error || null,
        },
        loadCounter: 0,
      };
    }

    case StoreUpdatedAclListDefault.type:
    case StoreUpdatedAclList.type:
    case StoreAclList.type: {
      const { response } = action.payload;
      let { data, error, exists } = response;
      if (!data.id) {
        const ids = Object.keys(data);
        if (ids.length === 0) {
          return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
        }
        data = data[ids[0]];
      }
      if (!data.id) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
      }
      if (state.acl.lists[data.id]) {
        data.nodes = state.acl.lists[data.id].nodes;
      }

      return {
        ...state,
        acl: {
          ...state.acl, lists: {...state.acl.lists, [data.id]: data},
          errorMessage: error || null,
        },
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case DropAclList.type: {
      const id = action.payload.response.data?.id;
      if (!state.acl.lists[id]) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
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
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case StoreMoveAclListNode.type:
    case StoreAclNodes.type: {
      let { parent_id } = action.payload;
      const data = action.payload.response.data || {};

      if (!parent_id && Object.keys(data).length > 0) {
        parent_id = data[Object.keys(data)[0]].parent?.id || 0;
      }

      if (!state.acl.lists[parent_id]) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
      }

      return {
        ...state,
        acl: {
          ...state.acl, lists: {...state.acl.lists, [parent_id]: {...state.acl.lists[parent_id], nodes: data}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case StoreSwitchAclNode.type:
    case StoreUpdatedAclNode.type: {
      const data = action.payload.response.data || {};
      const parent_id = data.parent?.id || 0;
      if (!parent_id || !state.acl.lists[parent_id]) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
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
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case StoreDelAclNode.type: {
      const data = action.payload.response.data || {};
      const parent_id = data.parent?.id || 0;
      if (!parent_id || !state.acl.lists[parent_id] || !state.acl.lists[parent_id].nodes[data.id]) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
      }

      const {[data.id]: toDel, ...rest} = state.acl.lists[parent_id].nodes;

      return {
        ...state,
        acl: {
          ...state.acl,
          lists: {...state.acl.lists, [parent_id]: {...state.acl.lists[parent_id], nodes: <Inodes>{...rest}}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case StoreNewAclNode.type: {
      const { id } = action.payload;
      if (!state.acl.lists[id].nodes) {
        state.acl.lists[id].nodes = {new: []}
      }

      const rest = [
        ...state.acl.lists[id].nodes?.new || [],
        <Inode>{}
      ];
      return {
        ...state,
        acl: {
          ...state.acl,
          lists: {...state.acl.lists, [id]: {...state.acl.lists[id], nodes: {...state.acl.lists[id].nodes, new: rest}}},
          errorMessage: null
        },
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case DropNewAclNode.type: {
      const { id } = action.payload;
      if (!state.acl.lists[id].nodes || !state.acl.lists[id].nodes.new) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
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
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    case StoreAclNode.type: {
      let data = action.payload.response.data || {};
      let id: any = 0;
      if (!data.id) {
        const ids = Object.keys(data);
        if (ids.length === 0) {
          return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
        }
        data = data[ids[0]];
      }
      if (!data.id) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
      }
      if (!data.parent?.id) {
        return {...state, loadCounter: Math.max(0, state.loadCounter - 1)};
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
        loadCounter: Math.max(0, state.loadCounter - 1),
      };
    }

    default: {
      return null;
    }
  }
}
