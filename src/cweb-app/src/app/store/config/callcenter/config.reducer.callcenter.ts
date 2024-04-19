import {
  ConfigActionTypes,
  All,
} from './config.actions.callcenter';
import {
  Icallcenter, Iitem,
  initialState,
  State
} from '../config.state.struct';
import {getParentId} from "../config.reducers";

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetCallcenterTiers:
    case ConfigActionTypes.UpdateCallcenterTier:
    case ConfigActionTypes.AddCallcenterTier:
    case ConfigActionTypes.DelCallcenterTier:
    case ConfigActionTypes.SendCallcenterCommand:
    case ConfigActionTypes.GetCallcenterMembers:
    case ConfigActionTypes.DelCallcenterMember:
    case ConfigActionTypes.AddCallcenterQueue:
    case ConfigActionTypes.RenameCallcenterQueue:
    case ConfigActionTypes.DelCallcenterQueue:
    case ConfigActionTypes.AddCallcenterSettings:
    case ConfigActionTypes.DelCallcenterSettings:
    case ConfigActionTypes.SwitchCallcenterSettings:
    case ConfigActionTypes.UpdateCallcenterSettings:
    case ConfigActionTypes.GetCallcenterSettings:
    case ConfigActionTypes.AddCallcenterQueueParam:
    case ConfigActionTypes.DelCallcenterQueueParam:
    case ConfigActionTypes.SwitchCallcenterQueueParam:
    case ConfigActionTypes.UpdateCallcenterQueueParam:
    case ConfigActionTypes.GetCallcenterQueuesParams:
    case ConfigActionTypes.ImportCallcenterAgentsAndTiers:
    case ConfigActionTypes.GetCallcenterAgents:
    case ConfigActionTypes.AddCallcenterAgent:
    case ConfigActionTypes.DelCallcenterAgent:
    case ConfigActionTypes.UpdateCallcenterAgent:
    case ConfigActionTypes.GetCallcenterQueues: {
      return {...state,
        callcenter: {
          ...state.callcenter,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotCallcenterError: {
      return {
        ...state,
        callcenter: {
          ...state.callcenter,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCallcenterQueues: {
      const data = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        callcenter: {
          ...state.callcenter, queues: {...state.callcenter?.queues, ...data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreRenameCallcenterQueue:
    case ConfigActionTypes.StoreAddCallcenterQueue: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        callcenter: {
          ...state.callcenter, queues: {...state.callcenter?.queues, [data.id]: data},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCallcenterQueue: {
      const id = action.payload.response.data?.id || 0;
      const queue = state.callcenter?.queues[id];
      if (!queue) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.callcenter?.queues;

      return {
        ...state,
        callcenter: {
          ...state.callcenter, queues: {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCallcenterSettings: {
      const data = action.payload.response.data || {};
      let rest = [...state.callcenter?.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.callcenter?.settings.new.slice(0, action.payload.index),
          null,
          ...state.callcenter?.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, settings: {...state.callcenter?.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCallcenterSettings: {
      const id = action.payload.response.data?.id || 0;
      if (!state.callcenter?.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.callcenter?.settings;

      return {
        ...state,
        callcenter: {
          ...state.callcenter, settings: {...rest, new: state.callcenter?.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCallcenterSettings: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, settings: {...state.callcenter?.settings, ...data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchCallcenterSettings:
    case ConfigActionTypes.StoreUpdateCallcenterSettings: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, settings: {...state.callcenter?.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewCallcenterSettings: {
      const rest = [
        ...state.callcenter?.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        callcenter: {
          ...state.callcenter, settings: {...state.callcenter?.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewCallcenterSettings: {
      const rest = [
        ...state.callcenter?.settings.new.slice(0, action.payload.index),
        null,
        ...state.callcenter?.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        callcenter: {
          ...state.callcenter, settings: {...state.callcenter?.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCallcenterQueueParam: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const queue = state.callcenter?.queues[parentId];
      if (!queue) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...queue.parameters.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...queue.parameters.new.slice(0, action.payload.index),
          null,
          ...queue.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, queues: {
            ...state.callcenter?.queues, [parentId]:
              {...queue, parameters: {...queue.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteCallcenterQueueParams: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.callcenter?.queues[fromId] || !state.callcenter?.queues[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.callcenter?.queues[toId].parameters ? state.callcenter?.queues[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.callcenter?.queues[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.callcenter?.queues[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        callcenter: {
          ...state.callcenter,
          queues: {
            ...state.callcenter?.queues,
            [toId]: {
              ...state.callcenter?.queues[toId],
              parameters: {
                ...state.callcenter?.queues[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelCallcenterQueueParam: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const queue = state.callcenter?.queues[parentId];
      if (!queue || !queue.parameters[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = queue.parameters;

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, queues: {
            ...state.callcenter?.queues, [parentId]:
              {...queue, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCallcenterQueuesParams: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const queue = state.callcenter?.queues[parentId];
      if (!queue) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, queues: {
            ...state.callcenter?.queues, [parentId]:
              {...queue, parameters: {...queue.parameters, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchCallcenterQueueParam:
    case ConfigActionTypes.StoreUpdateCallcenterQueueParam: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const queue = state.callcenter?.queues[parentId];
      if (!queue) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, queues: {
            ...state.callcenter?.queues, [parentId]:
              {...queue, parameters: {...queue.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewCallcenterQueueParam: {
      const queue = state.callcenter?.queues[action.payload.id];
      if (!queue) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...queue.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        callcenter: {
          ...state.callcenter, queues: {
            ...state.callcenter?.queues, [action.payload.id]:
              {...queue, parameters: {...queue.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewCallcenterQueueParam: {
      const queue = state.callcenter?.queues[action.payload.id];
      if (!queue) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...queue.parameters.new.slice(0, action.payload.index),
        null,
        ...queue.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        callcenter: {
          ...state.callcenter, queues: {
            ...state.callcenter?.queues, [action.payload.id]:
              {...queue, parameters: {...queue.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreImportCallcenterAgentsAndTiers: {
      const dataA = action.payload.response.data['callcenter_agents']?.items || [];
      const dataT = action.payload.response.data['callcenter_tiers']?.items || [];
      const totalA = action.payload.response.data['callcenter_agents']?.total || 0;
      const totalT = action.payload.response.data['callcenter_tiers']?.total || 0;
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const agents = state.callcenter?.agents || {table: [], list: {}, total: 0};
      const tiers = state.callcenter?.tiers || {table: [], list: {}, total: 0};

      return {
        ...state,
        callcenter: {
          ...state.callcenter,
          agents: {...agents, table: [ ...dataA], total: totalA},
          tiers: {...tiers, table: [ ...dataT], total: totalT},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCallcenterAgents: {
      const data = action.payload.response.data.items || [];
      let total = action.payload.response.data.total || 0;
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const agents = state.callcenter?.agents || {table: [], list: {}, total: 0};

      if (total < data.length) {
        total = data.length;
      }

      return {
        ...state,
        callcenter: {
          ...state.callcenter, agents: {...agents, table: [ ...data], total: total},
          exists: action.payload.response.exists,
          changed: {
            ...state.callcenter?.changed,
            agents: {},
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCallcenterAgent: {
      const data = action.payload.response.data || {};

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const agents = state.callcenter?.agents || {table: [], list: {}, total: 0};
      let total = agents.total || 0;

      return {
        ...state,
        callcenter: {
          ...state.callcenter, agents: {...agents, table: [data, ...agents.table], total: ++total},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCallcenterAgent: {
      const id: number = action.payload.response?.data?.id || 0;
      const agent = state.callcenter?.agents?.table || [];
      if (!agent) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = state.callcenter?.agents?.table || [];
      if (id) {
        rest = [
          ...rest.filter((item) => item && item.id && item.id !== id)
        ];
      }
      return {
        ...state,
        callcenter: {
          ...state.callcenter, agents: {...state.callcenter?.agents, table: [...rest]},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateCallcenterAgent: {
      const data = action.payload.response.data || {};
      if (!data || !data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const field = action.payload?.payload?.param?.name;

      const rest = state.callcenter?.agents?.table.map(item => {
          if (item.id && item.id === data.id) {
            Object.keys(data).forEach(k => {
              if (!state.callcenter?.changed.agents[String(data.id)] || !state.callcenter?.changed.agents[String(data.id)][k]) {
                return;
              }
              delete data[k];
            });
            return {...item, ...data};
          }
          return item;
        });

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter, agents: {table: [...rest], total: state.callcenter?.agents?.total},
          changed: {
            ...state.callcenter?.changed,
            agents: {
              ...state.callcenter?.changed.agents,
              [String(data.id)]: {
                ...state.callcenter?.changed.agents[String(data.id)],
                [field]: false,
              },
            },
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCallcenterTiers: {
      const data = action.payload.response.data.items || [];
      const total = action.payload.response.data.total || 0;
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const tiers = state.callcenter?.tiers || {table: [], list: {}, total: 0};

      return {
        ...state,
        callcenter: {
          ...state.callcenter, tiers: {...tiers, table: [ ...data], total: total},
          exists: action.payload.response.exists,
          changed: {
            ...state.callcenter?.changed,
            tiers: {},
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCallcenterTier: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const tiers = state.callcenter?.tiers || {table: [], list: {}, total: 0};
      let total = tiers.total || 0;

      return {
        ...state,
        callcenter: {
          ...state.callcenter, tiers: {...tiers, table: [ data, ...tiers.table], total: ++total},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCallcenterTier: {
      const id: number = action.payload.response?.data?.id || 0;
      const tier = state.callcenter?.tiers?.table || [];
      if (!tier) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = state.callcenter?.tiers?.table || [];
      if (id) {
        rest = [
          ...rest.filter((item) => item && item.id && item.id !== id)
        ];
      }
      return {
        ...state,
        callcenter: {
          ...state.callcenter, tiers: {...state.callcenter?.tiers, table: [...rest]},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateCallcenterTier: {
      const data = action.payload.response.data;
      if (!data || !data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const field = action.payload?.payload?.param?.name;
      const rest = state.callcenter?.tiers?.table.map(item => {
          if (item.id && item.id === data.id) {
            Object.keys(data).forEach(k => {
              if (!state.callcenter?.changed?.tiers[String(data.id)] || !state.callcenter?.changed?.tiers[String(data.id)][k]) {
                return;
              }
              delete data[k];
            });
            return {...item, ...data};
          }
          return item;
        });

      return {
        ...state,
        callcenter: <Icallcenter>{
          ...state.callcenter,
          tiers: {table: [...rest], total: state.callcenter?.tiers?.total},
          changed: {
            ...state.callcenter?.changed,
            tiers: {
              ...state.callcenter?.changed.tiers,
              [String(data.id)]: {
                ...state.callcenter?.changed.tiers[String(data.id)],
                [field]: false,
              },
            },
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSubscribeCallcenterAgents: {
      const data = action.payload.response.data['callcenter_agents_list'] || [];
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const agents = state.callcenter?.agents || {list: {}, table: [], total: 0};

      return {
        ...state,
        callcenter: {
          ...state.callcenter,
          agents: {...agents, list: {...agents.list, ...data}},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSubscribeCallcenterTiers: {
      const data = action.payload.response.data['callcenter_tiers_list'] || [];
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const tiers = state.callcenter?.tiers || {list: {}, table: [], total: 0};

      return {
        ...state,
        callcenter: {
          ...state.callcenter,
          tiers: {...tiers, list: {...tiers.list, ...data}},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCallcenterMembers: {
      const data = action.payload.response.data.items || [];
      const total = action.payload.response.total || 0;
      if (action.payload.response.exists === false) {
        return {
          ...state,
          callcenter: {...state.callcenter, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.callcenter) {
        state.callcenter = <Icallcenter>{};
        state.loadCounter = 0;
      }
      const tiers = state.callcenter?.members || {table: [], list: {}, total: 0};

      return {
        ...state,
        callcenter: {
          ...state.callcenter, members: {...tiers, table: [ ...data], total: total},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCallcenterMember: {
      const uuid: string = action.payload.response.uuid;
      const tier = state.callcenter?.members.table || [];
      if (!tier) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = state.callcenter?.members.table || [];
      if (uuid) {
        rest = [
          ...rest.filter((item) => item && item.uuid && item.uuid !== uuid)
        ];
      }
      return {
        ...state,
        callcenter: {
          ...state.callcenter, members: {...state.callcenter?.members, table: [...rest]},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSetChangedCallcenterTableField: {
      const tableName = action.payload.tableName;
      const fieldName = action.payload.fieldName;
      const rowId = action.payload.rowId;
      if (!tableName || !fieldName || !rowId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        callcenter: {
          ...state.callcenter,
          changed: {
            ...state.callcenter?.changed || {},
            [tableName]: {
              ...state.callcenter?.changed[tableName] || {},
              [rowId]: {
                ...state.callcenter?.changed[tableName][rowId] || {},
                [fieldName]: true,
              },
            }
          },
        },
      };
    }

    default: {
      return null;
    }
  }
}
