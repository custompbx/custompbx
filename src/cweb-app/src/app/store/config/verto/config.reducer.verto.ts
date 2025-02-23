import {Iitem, initialState, Iverto, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.verto';
import {getParentId} from '../config.reducers';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.MoveVertoProfileParameter:
    case ConfigActionTypes.GET_VERTO_CONFIG:
    case ConfigActionTypes.GET_VERTO_PROFILES_PARAMS:
    case ConfigActionTypes.UPDATE_VERTO_SETTING:
    case ConfigActionTypes.SWITCH_VERTO_SETTING:
    case ConfigActionTypes.ADD_VERTO_SETTING:
    case ConfigActionTypes.DEL_VERTO_SETTING:
    case ConfigActionTypes.UPDATE_VERTO_PROFILE_PARAM:
    case ConfigActionTypes.SWITCH_VERTO_PROFILE_PARAM:
    case ConfigActionTypes.ADD_VERTO_PROFILE_PARAM:
    case ConfigActionTypes.DEL_VERTO_PROFILE_PARAM:
    case ConfigActionTypes.ADD_VERTO_PROFILE:
    case ConfigActionTypes.DEL_VERTO_PROFILE:
    case ConfigActionTypes.RENAME_VERTO_PROFILE: {
      return {...state,
        verto: {
          ...state.verto,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotVertoError: {
      return {
        ...state,
        verto: {
          ...state.verto,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.STORE_SWITCH_VERTO_PROFILE:
    case ConfigActionTypes.STORE_GET_VERTO_CONFIG: {
      const settings = action.payload.response.data['settings'] || {};
      const profiles = action.payload.response.data['profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          verto: {...state.verto, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.verto) {
        state.verto = <Iverto>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        verto: {
          ...state.verto,
          settings: {...state.verto.settings, ...settings},
          profiles: {...state.verto.profiles, ...profiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_VERTO_PROFILE:
    case ConfigActionTypes.STORE_RENAME_VERTO_PROFILE: {
      const data = action.payload.response.data;
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.verto) {
        state.verto = <Iverto>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        verto: {
          ...state.verto,
          profiles: {...state.verto.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_VERTO_SETTING: {
      const id = action.payload.response.data?.id || 0;
      if (!state.verto.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.verto.settings;

      return {
        ...state,
        verto: {
          ...state.verto, settings: {...rest, new: state.verto.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_VERTO_SETTING:
    case ConfigActionTypes.STORE_UPDATE_VERTO_SETTING: {
      const data = action.payload.response.data;

      return {
        ...state,
        verto: <Iverto>{
          ...state.verto, settings: {...state.verto.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_VERTO_SETTING: {
      const rest = [
        ...state.verto.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        verto: {
          ...state.verto, settings: {...state.verto.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_VERTO_SETTING: {
      const rest = [
        ...state.verto.settings.new.slice(0, action.payload.index),
        null,
        ...state.verto.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        verto: {
          ...state.verto, settings: {...state.verto.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_VERTO_SETTING: {
      const data = action.payload.response.data;
      let rest = [...state.verto.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.verto.settings.new.slice(0, action.payload.index),
          null,
          ...state.verto.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        verto: <Iverto>{
          ...state.verto, settings: {...state.verto.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_VERTO_PROFILE_PARAM: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.verto.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.parameters.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.parameters.new.slice(0, action.payload.index),
          null,
          ...profile.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        verto: <Iverto>{
          ...state.verto, profiles: {
            ...state.verto.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_PASTE_VERTO_PROFILE_PARAMS: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.verto.profiles[fromId] || !state.verto.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.verto.profiles[toId].parameters ? state.verto.profiles[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.verto.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.verto.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        verto: {
          ...state.verto,
          profiles: {
            ...state.verto.profiles,
            [toId]: {
              ...state.verto.profiles[toId],
              parameters: {
                ...state.verto.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.STORE_DEL_VERTO_PROFILE_PARAM: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.verto.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        verto: <Iverto>{
          ...state.verto, profiles: {
            ...state.verto.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreMoveVertoProfileParameter:
    case ConfigActionTypes.STORE_GET_VERTO_PROFILES_PARAMS: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.verto.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        verto: <Iverto>{
          ...state.verto, profiles: {
            ...state.verto.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, ...data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_VERTO_PROFILE_PARAM:
    case ConfigActionTypes.STORE_UPDATE_VERTO_PROFILE_PARAM: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.verto.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }


      return {
        ...state,
        verto: <Iverto>{
          ...state.verto, profiles: {
            ...state.verto.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_VERTO_PROFILE_PARAM: {
      const profile = state.verto.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        verto: {
          ...state.verto, profiles: {
            ...state.verto.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_VERTO_PROFILE_PARAM: {
      const profile = state.verto.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters.new.slice(0, action.payload.index),
        null,
        ...profile.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        verto: {
          ...state.verto, profiles: {
            ...state.verto.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_VERTO_PROFILE: {
      const data = action.payload.response.data;
      const profile = state.verto.profiles[data.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.verto.profiles;

      return {
        ...state,
        verto: {
          ...state.verto, profiles: {...rest},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return null;
    }
  }
}
