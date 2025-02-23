import {Iitem, Ififo, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.fifo';
import {getParentId} from '../config.reducers';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.UpdateFifoFifoImportance:
    case ConfigActionTypes.GetFifo:
    case ConfigActionTypes.GetFifoFifoMembers:
    case ConfigActionTypes.UpdateFifoParameter:
    case ConfigActionTypes.SwitchFifoParameter:
    case ConfigActionTypes.AddFifoParameter:
    case ConfigActionTypes.DelFifoParameter:
    case ConfigActionTypes.AddFifoFifoMember:
    case ConfigActionTypes.UpdateFifoFifoMember:
    case ConfigActionTypes.SwitchFifoFifoMember:
    case ConfigActionTypes.DelFifoFifoMember:
    case ConfigActionTypes.AddFifoFifo:
    case ConfigActionTypes.DelFifoFifo:
    case ConfigActionTypes.UpdateFifoFifo: {
      return {...state,
        fifo: {
          ...state.fifo,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotFifoError: {
      return {
        ...state,
        fifo: {
          ...state.fifo,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchFifoFifo:
    case ConfigActionTypes.StoreGetFifo: {
      const settings = action.payload.response.data['settings'] || {};
      const fifos = action.payload.response.data['profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          fifo: {...state.fifo, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.fifo) {
        state.fifo = <Ififo>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        fifo: {
          ...state.fifo,
          settings: {...state.fifo.settings, ...settings},
          fifos: {...state.fifo.fifos, ...fifos},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchFifoFifo:
    case ConfigActionTypes.StoreAddFifoFifo:
    case ConfigActionTypes.StoreUpdateFifoFifo: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0 };
      }

      if (!state.fifo) {
        state.fifo = <Ififo>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        fifo: {
          ...state.fifo,
          fifos: {...state.fifo.fifos, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelFifoParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.fifo.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.fifo.settings;

      return {
        ...state,
        fifo: {
          ...state.fifo, settings: {...rest, new: state.fifo.settings?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchFifoParameter:
    case ConfigActionTypes.StoreUpdateFifoParameter: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0 };
      }

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, settings: {...state.fifo.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateFifoFifoImportance: {
      const data = action.payload.response.data || {};
      if (!data.id || !state.fifo.fifos[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, fifos: {...state.fifo.fifos, [data.id]: {...state.fifo.fifos[data.id], importance: data.importance}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewFifoParameter: {
      const rest = [
        ...state.fifo.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        fifo: {
          ...state.fifo, settings: {...state.fifo.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewFifoParameter: {
      const rest = [
        ...state.fifo.settings.new.slice(0, action.payload.index),
        null,
        ...state.fifo.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        fifo: {
          ...state.fifo, settings: {...state.fifo.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddFifoParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.fifo.settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.fifo.settings.new.slice(0, action.payload.index),
          null,
          ...state.fifo.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, settings: {...state.fifo.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddFifoFifoMember: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const fifo = state.fifo.fifos[parentId];
      if (!fifo) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...fifo.members?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...fifo.members.new.slice(0, action.payload.index),
          null,
          ...fifo.members.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, fifos: {
            ...state.fifo.fifos, [parentId]:
              {...fifo, members: {...fifo.members, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteFifoFifoMembers: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.fifo.fifos[fromId] || !state.fifo.fifos[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.fifo.fifos[toId].members ? state.fifo.fifos[toId].members?.new || [] : [];

      const newArray = Object.keys(state.fifo.fifos[fromId].members).map(i => {
        if (i === 'new') {
          return;
        }
        return state.fifo.fifos[fromId].members[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        fifo: {
          ...state.fifo,
          fifos: {
            ...state.fifo.fifos,
            [toId]: {
              ...state.fifo.fifos[toId],
              members: {
                ...state.fifo.fifos[toId].members,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelFifoFifoMember: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const fifo = state.fifo.fifos[parentId];
      if (!fifo || !fifo.members[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = fifo.members;

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, fifos: {
            ...state.fifo.fifos, [parentId]:
              {...fifo, members: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetFifoFifoMembers: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const fifo = state.fifo.fifos[parentId];
      if (!fifo) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, fifos: {
            ...state.fifo.fifos, [parentId]:
              {...fifo, members: {...fifo.members, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchFifoFifoMember:
    case ConfigActionTypes.StoreUpdateFifoFifoMember: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const fifo = state.fifo.fifos[parentId];
      if (!fifo) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        fifo: <Ififo>{
          ...state.fifo, fifos: {
            ...state.fifo.fifos, [parentId]:
              {...fifo, members: {...fifo.members, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewFifoFifoMember: {
      const fifo = state.fifo.fifos[action.payload.id];
      if (!fifo) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...fifo.members?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        fifo: {
          ...state.fifo, fifos: {
            ...state.fifo.fifos, [action.payload.id]:
              {...fifo, members: {...fifo.members, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewFifoFifoMember: {
      const fifo = state.fifo.fifos[action.payload.id];
      if (!fifo) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...fifo.members.new.slice(0, action.payload.index),
        null,
        ...fifo.members.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        fifo: {
          ...state.fifo, fifos: {
            ...state.fifo.fifos, [action.payload.id]:
              {...fifo, members: {...fifo.members, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelFifoFifo: {
      const data = action.payload.response.data || {};
      const fifo = state.fifo.fifos[data.id];
      if (!fifo) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.fifo.fifos;

      return {
        ...state,
        fifo: {
          ...state.fifo, fifos: {...rest},
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
