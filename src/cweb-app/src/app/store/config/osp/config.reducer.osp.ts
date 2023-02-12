import {Iitem, Iosp, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.osp';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetOsp:
    case ConfigActionTypes.GetOspProfileParameters:
    case ConfigActionTypes.UpdateOspParameter:
    case ConfigActionTypes.SwitchOspParameter:
    case ConfigActionTypes.AddOspParameter:
    case ConfigActionTypes.DelOspParameter:
    case ConfigActionTypes.AddOspProfileParameter:
    case ConfigActionTypes.UpdateOspProfileParameter:
    case ConfigActionTypes.SwitchOspProfileParameter:
    case ConfigActionTypes.DelOspProfileParameter:
    case ConfigActionTypes.AddOspProfile:
    case ConfigActionTypes.DelOspProfile:
    case ConfigActionTypes.UpdateOspProfile: {
      return {...state,
        osp: {
          ...state.osp,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotOspError: {
      return {
        ...state,
        osp: {
          ...state.osp,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchOspProfile:
    case ConfigActionTypes.StoreGetOsp: {
      const settings = action.payload.response.data['settings'] || {};
      const  profiles = action.payload.response.data['profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          osp: {...state.osp, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.osp) {
        state.osp = <Iosp>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        osp: {
          ...state.osp,
          settings: {...settings},
          profiles: {...profiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreAddOspProfile:
    case ConfigActionTypes.StoreUpdateOspProfile: {
      const data = action.payload.response.data;
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.osp) {
        state.osp = <Iosp>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        osp: {
          ...state.osp,
          profiles: {...state.osp.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelOspParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.osp.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.osp.settings;

      return {
        ...state,
        osp: {
          ...state.osp, settings: {...rest, new: state.osp.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchOspParameter:
    case ConfigActionTypes.StoreUpdateOspParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        osp: <Iosp>{
          ...state.osp, settings: {...state.osp.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOspParameter: {
      const rest = [
        ...state.osp.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        osp: {
          ...state.osp, settings: {...state.osp.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOspParameter: {
      const rest = [
        ...state.osp.settings.new.slice(0, action.payload.index),
        null,
        ...state.osp.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        osp: {
          ...state.osp, settings: {...state.osp.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOspParameter: {
      const data = action.payload.response.data;
      let rest = [...state.osp.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.osp.settings.new.slice(0, action.payload.index),
          null,
          ...state.osp.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        osp: <Iosp>{
          ...state.osp, settings: {...state.osp.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOspProfileParameter: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.osp.profiles[parentId];
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
        osp: <Iosp>{
          ...state.osp, profiles: {
            ...state.osp.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteOspProfileParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.osp.profiles[fromId] || !state.osp.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.osp.profiles[toId].parameters ? state.osp.profiles[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.osp.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.osp.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        osp: {
          ...state.osp,
          profiles: {
            ...state.osp.profiles,
            [toId]: {
              ...state.osp.profiles[toId],
              parameters: {
                ...state.osp.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelOspProfileParameter: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.osp.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        osp: <Iosp>{
          ...state.osp, profiles: {
            ...state.osp.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetOspProfileParameters: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.osp.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        osp: <Iosp>{
          ...state.osp, profiles: {
            ...state.osp.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, ...data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchOspProfileParameter:
    case ConfigActionTypes.StoreUpdateOspProfileParameter: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.osp.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        osp: <Iosp>{
          ...state.osp, profiles: {
            ...state.osp.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOspProfileParameter: {
      const id = action.payload.id;
      const profile = state.osp.profiles[id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        osp: {
          ...state.osp, profiles: {
            ...state.osp.profiles, [profile.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOspProfileParameter: {
      const profile = state.osp.profiles[action.payload.id];
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
        osp: {
          ...state.osp, profiles: {
            ...state.osp.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelOspProfile: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.osp.profiles[data.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.osp.profiles;

      return {
        ...state,
        osp: {
          ...state.osp, profiles: {...rest},
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
