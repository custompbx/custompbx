import {Iitem, Ilcr, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.lcr';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetLcr:
    case ConfigActionTypes.GetLcrProfileParameters:
    case ConfigActionTypes.UpdateLcrParameter:
    case ConfigActionTypes.SwitchLcrParameter:
    case ConfigActionTypes.AddLcrParameter:
    case ConfigActionTypes.DelLcrParameter:
    case ConfigActionTypes.AddLcrProfileParameter:
    case ConfigActionTypes.UpdateLcrProfileParameter:
    case ConfigActionTypes.SwitchLcrProfileParameter:
    case ConfigActionTypes.DelLcrProfileParameter:
    case ConfigActionTypes.AddLcrProfile:
    case ConfigActionTypes.DelLcrProfile:
    case ConfigActionTypes.UpdateLcrProfile: {
      return {...state,
        lcr: {
          ...state.lcr,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotLcrError: {
      return {
        ...state,
        lcr: {
          ...state.lcr,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchLcrProfile:
    case ConfigActionTypes.StoreGetLcr: {
      const settings = action.payload.response.data['settings'] || {};
      const  profiles = action.payload.response.data['profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          lcr: {...state.lcr, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.lcr) {
        state.lcr = <Ilcr>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        lcr: {
          ...state.lcr,
          settings: {...settings},
          profiles: {...profiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreAddLcrProfile:
    case ConfigActionTypes.StoreUpdateLcrProfile: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.lcr) {
        state.lcr = <Ilcr>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        lcr: {
          ...state.lcr,
          profiles: {...state.lcr.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelLcrParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.lcr.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.lcr.settings;

      return {
        ...state,
        lcr: {
          ...state.lcr, settings: {...rest, new: state.lcr.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchLcrParameter:
    case ConfigActionTypes.StoreUpdateLcrParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        lcr: <Ilcr>{
          ...state.lcr, settings: {...state.lcr.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewLcrParameter: {
      const rest = [
        ...state.lcr.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        lcr: {
          ...state.lcr, settings: {...state.lcr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewLcrParameter: {
      const rest = [
        ...state.lcr.settings.new.slice(0, action.payload.index),
        null,
        ...state.lcr.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        lcr: {
          ...state.lcr, settings: {...state.lcr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddLcrParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.lcr.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.lcr.settings.new.slice(0, action.payload.index),
          null,
          ...state.lcr.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        lcr: <Ilcr>{
          ...state.lcr, settings: {...state.lcr.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddLcrProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const profile = state.lcr.profiles[parentId];
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
        lcr: <Ilcr>{
          ...state.lcr, profiles: {
            ...state.lcr.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteLcrProfileParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.lcr.profiles[fromId] || !state.lcr.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.lcr.profiles[toId].parameters ? state.lcr.profiles[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.lcr.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.lcr.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        lcr: {
          ...state.lcr,
          profiles: {
            ...state.lcr.profiles,
            [toId]: {
              ...state.lcr.profiles[toId],
              parameters: {
                ...state.lcr.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelLcrProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const profile = state.lcr.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        lcr: <Ilcr>{
          ...state.lcr, profiles: {
            ...state.lcr.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetLcrProfileParameters: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.lcr.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        lcr: <Ilcr>{
          ...state.lcr, profiles: {
            ...state.lcr.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, ...data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchLcrProfileParameter:
    case ConfigActionTypes.StoreUpdateLcrProfileParameter: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.lcr.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        lcr: <Ilcr>{
          ...state.lcr, profiles: {
            ...state.lcr.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewLcrProfileParameter: {
      const id = action.payload.id;
      const profile = state.lcr.profiles[id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        lcr: {
          ...state.lcr, profiles: {
            ...state.lcr.profiles, [profile.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewLcrProfileParameter: {
      const profile = state.lcr.profiles[action.payload.id];
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
        lcr: {
          ...state.lcr, profiles: {
            ...state.lcr.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelLcrProfile: {
      const data = action.payload.response.data;
      const profile = state.lcr.profiles[data.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.lcr.profiles;

      return {
        ...state,
        lcr: {
          ...state.lcr, profiles: {...rest},
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

function getParentId (data): number {
  let id = 0;
  if (data.id) {
    id = data?.parent?.id || 0;
  } else {
    const ids = Object.keys(data);
    if (ids.length === 0) {
      return id;
    }
    id = data[ids[0]]?.parent?.id || 0;
  }
  return id;
}
