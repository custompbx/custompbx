import {Iitem, Iopal, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.opal';
import {getParentId} from '../config.reducers';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetOpal:
    case ConfigActionTypes.GetOpalListenerParameters:
    case ConfigActionTypes.UpdateOpalParameter:
    case ConfigActionTypes.SwitchOpalParameter:
    case ConfigActionTypes.AddOpalParameter:
    case ConfigActionTypes.DelOpalParameter:
    case ConfigActionTypes.AddOpalListenerParameter:
    case ConfigActionTypes.UpdateOpalListenerParameter:
    case ConfigActionTypes.SwitchOpalListenerParameter:
    case ConfigActionTypes.DelOpalListenerParameter:
    case ConfigActionTypes.AddOpalListener:
    case ConfigActionTypes.DelOpalListener:
    case ConfigActionTypes.UpdateOpalListener: {
      return {...state,
        opal: {
          ...state.opal,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotOpalError: {
      return {
        ...state,
        opal: {
          ...state.opal,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchOpalListener:
    case ConfigActionTypes.StoreGetOpal: {
      const settings = action.payload.response.data['settings'] || {};
      const listeners = action.payload.response.data['listeners'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          opal: {...state.opal, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.opal) {
        state.opal = <Iopal>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        opal: {
          ...state.opal,
          settings: {...state.opal.settings, ...settings},
          listeners: {...state.opal.listeners, ...listeners},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreAddOpalListener:
    case ConfigActionTypes.StoreUpdateOpalListener: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.opal) {
        state.opal = <Iopal>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        opal: {
          ...state.opal,
          listeners: {...state.opal.listeners, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelOpalParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.opal.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.opal.settings;

      return {
        ...state,
        opal: {
          ...state.opal, settings: {...rest, new: state.opal.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchOpalParameter:
    case ConfigActionTypes.StoreUpdateOpalParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        opal: <Iopal>{
          ...state.opal, settings: {...state.opal.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOpalParameter: {
      const rest = [
        ...state.opal.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        opal: {
          ...state.opal, settings: {...state.opal.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOpalParameter: {
      const rest = [
        ...state.opal.settings.new.slice(0, action.payload.index),
        null,
        ...state.opal.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        opal: {
          ...state.opal, settings: {...state.opal.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOpalParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.opal.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.opal.settings.new.slice(0, action.payload.index),
          null,
          ...state.opal.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        opal: <Iopal>{
          ...state.opal, settings: {...state.opal.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOpalListenerParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const listener = state.opal.listeners[parentId];
      if (!listener) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...listener.parameters.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...listener.parameters.new.slice(0, action.payload.index),
          null,
          ...listener.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        opal: <Iopal>{
          ...state.opal, listeners: {
            ...state.opal.listeners, [parentId]:
              {...listener, parameters: {...listener.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteOpalListenerParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.opal.listeners[fromId] || !state.opal.listeners[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.opal.listeners[toId].parameters ? state.opal.listeners[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.opal.listeners[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.opal.listeners[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        opal: {
          ...state.opal,
          listeners: {
            ...state.opal.listeners,
            [toId]: {
              ...state.opal.listeners[toId],
              parameters: {
                ...state.opal.listeners[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelOpalListenerParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const listener = state.opal.listeners[parentId];
      if (!listener || !listener.parameters[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = listener.parameters;

      return {
        ...state,
        opal: <Iopal>{
          ...state.opal, listeners: {
            ...state.opal.listeners, [parentId]:
              {...listener, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetOpalListenerParameters: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const listener = state.opal.listeners[parentId];
      if (!listener) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        opal: <Iopal>{
          ...state.opal, listeners: {
            ...state.opal.listeners, [parentId]:
              {...listener, parameters: {...listener.parameters, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchOpalListenerParameter:
    case ConfigActionTypes.StoreUpdateOpalListenerParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const listener = state.opal.listeners[parentId];
      if (!listener) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        opal: <Iopal>{
          ...state.opal, listeners: {
            ...state.opal.listeners, [parentId]:
              {...listener, parameters: {...listener.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOpalListenerParameter: {
      const listener = state.opal.listeners[action.payload.id];
      if (!listener) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...listener.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        opal: {
          ...state.opal, listeners: {
            ...state.opal.listeners, [action.payload.id]:
              {...listener, parameters: {...listener.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOpalListenerParameter: {
      const listener = state.opal.listeners[action.payload.id];
      if (!listener) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...listener.parameters.new.slice(0, action.payload.index),
        null,
        ...listener.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        opal: {
          ...state.opal, listeners: {
            ...state.opal.listeners, [action.payload.id]:
              {...listener, parameters: {...listener.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelOpalListener: {
      const data = action.payload.response.data || {};
      const listener = state.opal.listeners[data.id];
      if (!listener) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.opal.listeners;

      return {
        ...state,
        opal: {
          ...state.opal, listeners: {...rest},
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
