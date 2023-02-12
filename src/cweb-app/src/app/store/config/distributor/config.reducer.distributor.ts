import {
  ConfigActionTypes,
  All,
} from './config.actions.distributor';
import {
  IdistributorNode,
  IdistributorNodes,
  initialState,
  State
} from '../config.state.struct';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetDistributorConfig:
    case ConfigActionTypes.AddDistributorList:
    case ConfigActionTypes.UpdateDistributorList:
    case ConfigActionTypes.DelDistributorList:
    case ConfigActionTypes.GetDistributorNodes:
    case ConfigActionTypes.AddDistributorNode:
    case ConfigActionTypes.UpdateDistributorNode:
    case ConfigActionTypes.DelDistributorNode:
    case ConfigActionTypes.SwitchDistributorNode: {
      return {...state,
        distributor: {
          ...state.distributor,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotDistributorError: {
      return {
        ...state,
        distributor: {
          ...state.distributor,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetDistributorConfig: {
      return {
        ...state,
        distributor: {
          ...state.distributor,
          lists: {
            ...action.payload.response.data
          },
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: 0,
      };
    }
    case ConfigActionTypes.StoreAddDistributorList:
    case ConfigActionTypes.StoreUpdateDistributorList: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        distributor: {
          ...state.distributor, lists: {...state.distributor.lists, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreDelDistributorList: {
      const id = action.payload.response.data?.id || 0;
      if (!state.distributor.lists[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.distributor.lists;

      return {
        ...state,
        distributor: {
          ...state.distributor,
          lists:
            {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreGetDistributorNodes: {
      const id = action.payload.id;
      if (!state.distributor.lists[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        distributor: {
          ...state.distributor, lists: {...state.distributor.lists,
            [id]: {...state.distributor.lists[id], nodes: action.payload.response.data}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreSwitchDistributorNode:
    case ConfigActionTypes.StoreUpdateDistributorNode: {
      const data = action.payload.response.data;
      const id = data?.id || 0;
      const parentId = data?.parent?.id || 0;
      if (!state.distributor.lists[parentId] || !state.distributor.lists[parentId].nodes[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        distributor: {
          ...state.distributor,
          lists: {
            ...state.distributor.lists,
            [parentId]: {...state.distributor.lists[parentId], nodes: {...state.distributor.lists[parentId].nodes, [id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelDistributorNode: {
      const id = action.payload.response.data?.id || 0;
      const parentId = action.payload.response.data?.parent?.id || 0;
      if (!state.distributor.lists[parentId] || !state.distributor.lists[parentId].nodes[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.distributor.lists[parentId].nodes;

      return {
        ...state,
        distributor: {
          ...state.distributor, lists: {...state.distributor.lists,
            [parentId]: {...state.distributor.lists[parentId], nodes: <IdistributorNodes>{...rest}}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreNewDistributorNode: {
      const id = action.payload;
      if (!state.distributor.lists[id].nodes) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const rest = [
        ...state.distributor.lists[id].nodes?.new || [],
        <IdistributorNode>{}
      ];

      return {
        ...state,
        distributor: {
          ...state.distributor, lists: {...state.distributor.lists,
            [id]: {...state.distributor.lists[id], nodes: {...state.distributor.lists[id].nodes, new: rest}}},
          errorMessage: null
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreDelNewDistributorNode: {
      const id = action.payload.id;
      if (!state.distributor.lists[id].nodes || !state.distributor.lists[id].nodes.new) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const rest = [
        ...state.distributor.lists[id].nodes.new.slice(0, action.payload.index),
        null,
        ...state.distributor.lists[id].nodes.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        distributor: {
          ...state.distributor, lists: {
            ...state.distributor.lists, [id]: {
              ...state.distributor.lists[id], nodes: {...state.distributor.lists[id].nodes, new: rest}
            }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreAddDistributorNode: {
      const data = action.payload.response.data || {};
      const parentId = data.parent?.id || 0;
      if (!parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const rest = [
        ...state.distributor.lists[parentId].nodes.new.slice(0, action.payload.index),
        null,
        ...state.distributor.lists[parentId].nodes.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        distributor: {
          ...state.distributor,
          lists: {
            ...state.distributor.lists,
            [parentId]: {...state.distributor.lists[parentId],
              nodes: {...state.distributor.lists[parentId].nodes, [data.id]: data, new: rest}
            }
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
