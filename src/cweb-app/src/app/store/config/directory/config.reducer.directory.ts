import {Iitem, Idirectory, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.directory';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetDirectory:
    case ConfigActionTypes.GetDirectoryProfileParameters:
    case ConfigActionTypes.UpdateDirectoryParameter:
    case ConfigActionTypes.SwitchDirectoryParameter:
    case ConfigActionTypes.AddDirectoryParameter:
    case ConfigActionTypes.DelDirectoryParameter:
    case ConfigActionTypes.AddDirectoryProfileParameter:
    case ConfigActionTypes.UpdateDirectoryProfileParameter:
    case ConfigActionTypes.SwitchDirectoryProfileParameter:
    case ConfigActionTypes.DelDirectoryProfileParameter:
    case ConfigActionTypes.AddDirectoryProfile:
    case ConfigActionTypes.DelDirectoryProfile:
    case ConfigActionTypes.UpdateDirectoryProfile: {
      return {...state,
        directory: {
          ...state.directory,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotDirectoryError: {
      return {
        ...state,
        directory: {
          ...state.directory,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchDirectoryProfile:
    case ConfigActionTypes.StoreGetDirectory: {
      const settings = action.payload.response.data['settings'] || {};
      const profiles = action.payload.response.data['profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          directory: {...state.directory, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.directory) {
        state.directory = <Idirectory>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        directory: {
          ...state.directory,
          settings: {...state.directory.settings, ...settings},
          profiles: {...state.directory.profiles, ...profiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddDirectoryProfile:
    case ConfigActionTypes.StoreUpdateDirectoryProfile: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.directory) {
        state.directory = <Idirectory>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        directory: {
          ...state.directory,
          profiles: {...state.directory.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelDirectoryParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.directory.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.directory.settings;

      return {
        ...state,
        directory: {
          ...state.directory, settings: {...rest, new: state.directory.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchDirectoryParameter:
    case ConfigActionTypes.StoreUpdateDirectoryParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        directory: <Idirectory>{
          ...state.directory, settings: {...state.directory.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewDirectoryParameter: {
      const rest = [
        ...state.directory.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        directory: {
          ...state.directory, settings: {...state.directory.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewDirectoryParameter: {
      const rest = [
        ...state.directory.settings.new.slice(0, action.payload.index),
        null,
        ...state.directory.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        directory: {
          ...state.directory, settings: {...state.directory.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddDirectoryParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.directory.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.directory.settings.new.slice(0, action.payload.index),
          null,
          ...state.directory.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        directory: <Idirectory>{
          ...state.directory, settings: {...state.directory.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddDirectoryProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const profile = state.directory.profiles[parentId];
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
        directory: <Idirectory>{
          ...state.directory, profiles: {
            ...state.directory.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteDirectoryProfileParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.directory.profiles[fromId] || !state.directory.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.directory.profiles[toId].parameters ? state.directory.profiles[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.directory.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.directory.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        directory: {
          ...state.directory,
          profiles: {
            ...state.directory.profiles,
            [toId]: {
              ...state.directory.profiles[toId],
              parameters: {
                ...state.directory.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelDirectoryProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const profile = state.directory.profiles[parentId];
      if (!profile || !profile.parameters[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        directory: <Idirectory>{
          ...state.directory, profiles: {
            ...state.directory.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetDirectoryProfileParameters: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.directory.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        directory: <Idirectory>{
          ...state.directory, profiles: {
            ...state.directory.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, ...data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchDirectoryProfileParameter:
    case ConfigActionTypes.StoreUpdateDirectoryProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.directory.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        directory: <Idirectory>{
          ...state.directory, profiles: {
            ...state.directory.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewDirectoryProfileParameter: {
      const profile = state.directory.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        directory: {
          ...state.directory, profiles: {
            ...state.directory.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewDirectoryProfileParameter: {
      const profile = state.directory.profiles[action.payload.id];
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
        directory: {
          ...state.directory, profiles: {
            ...state.directory.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelDirectoryProfile: {
      const id = action.payload.response.data?.id || 0;
      const profile = state.directory.profiles[id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.directory.profiles;

      return {
        ...state,
        directory: {
          ...state.directory, profiles: {...rest},
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
