import {Iitem, initialState, State, IpostSwitcheModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.post-switch';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetPostSwitch:
    case ConfigActionTypes.UpdatePostSwitchParameter:
    case ConfigActionTypes.SwitchPostSwitchParameter:
    case ConfigActionTypes.AddPostSwitchParameter:
    case ConfigActionTypes.DelPostSwitchParameter:
    case ConfigActionTypes.UpdatePostSwitchCliKeybinding:
    case ConfigActionTypes.SwitchPostSwitchCliKeybinding:
    case ConfigActionTypes.AddPostSwitchCliKeybinding:
    case ConfigActionTypes.DelPostSwitchCliKeybinding:
    case ConfigActionTypes.UpdatePostSwitchDefaultPtime:
    case ConfigActionTypes.SwitchPostSwitchDefaultPtime:
    case ConfigActionTypes.AddPostSwitchDefaultPtime:
    case ConfigActionTypes.DelPostSwitchDefaultPtime: {
      return {...state,
        post_load_switch: {
          ...state.post_load_switch,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotPostSwitchError: {
      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetPostSwitch: {
      const settings = action.payload.response.data['settings'] || {};
      const cliKeybindings = action.payload.response.data['cli_keybinding'] || {};
      const defaultPtime = action.payload.response.data['default_ptime'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          post_load_switch: {
            ...state.post_load_switch, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.post_load_switch) {
        state.post_load_switch = <IpostSwitcheModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch,
          settings: {...state.post_load_switch.settings, ...settings},
          cli_keybindings: {...state.post_load_switch.cli_keybindings, ...cliKeybindings},
          default_ptimes: {...state.post_load_switch.default_ptimes, ...defaultPtime},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPostSwitchParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.post_load_switch.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.post_load_switch.settings;

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, settings: {...rest, new: state.post_load_switch.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchPostSwitchParameter:
    case ConfigActionTypes.StoreUpdatePostSwitchParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        post_load_switch: <IpostSwitcheModule>{
          ...state.post_load_switch, settings: {...state.post_load_switch.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPostSwitchParameter: {
      const rest = [
        ...state.post_load_switch.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, settings: {...state.post_load_switch.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPostSwitchParameter: {
      const rest = [
        ...state.post_load_switch.settings.new.slice(0, action.payload.index),
        null,
        ...state.post_load_switch.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, settings: {...state.post_load_switch.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPostSwitchParameter: {
      const data = action.payload.response.data;
      let rest = [...state.post_load_switch.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.post_load_switch.settings.new.slice(0, action.payload.index),
          null,
          ...state.post_load_switch.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        post_load_switch: <IpostSwitcheModule>{
          ...state.post_load_switch, settings: {...state.post_load_switch.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPostSwitchCliKeybinding: {
      const id = action.payload.response.data?.id || 0;
      if (!state.post_load_switch.cli_keybindings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.post_load_switch.cli_keybindings;

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, cli_keybindings: {...rest, new: state.post_load_switch.cli_keybindings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchPostSwitchCliKeybinding:
    case ConfigActionTypes.StoreUpdatePostSwitchCliKeybinding: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        post_load_switch: <IpostSwitcheModule>{
          ...state.post_load_switch, cli_keybindings: {...state.post_load_switch.cli_keybindings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPostSwitchCliKeybinding: {
      const rest = [
        ...state.post_load_switch.cli_keybindings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, cli_keybindings: {...state.post_load_switch.cli_keybindings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPostSwitchCliKeybinding: {
      const rest = [
        ...state.post_load_switch.cli_keybindings.new.slice(0, action.payload.index),
        null,
        ...state.post_load_switch.cli_keybindings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, cli_keybindings: {...state.post_load_switch.cli_keybindings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPostSwitchCliKeybinding: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.post_load_switch.cli_keybindings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.post_load_switch.cli_keybindings.new.slice(0, action.payload.index),
          null,
          ...state.post_load_switch.cli_keybindings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        post_load_switch: <IpostSwitcheModule>{
          ...state.post_load_switch, cli_keybindings: {...state.post_load_switch.cli_keybindings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPostSwitchDefaultPtime: {
      const id = action.payload.response.data?.id || 0;
      if (!state.post_load_switch.default_ptimes[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.post_load_switch.default_ptimes;

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, default_ptimes: {...rest, new: state.post_load_switch.default_ptimes.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchPostSwitchDefaultPtime:
    case ConfigActionTypes.StoreUpdatePostSwitchDefaultPtime: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        post_load_switch: <IpostSwitcheModule>{
          ...state.post_load_switch, default_ptimes: {...state.post_load_switch.default_ptimes, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPostSwitchDefaultPtime: {
      const rest = [
        ...state.post_load_switch.default_ptimes?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, default_ptimes: {...state.post_load_switch.default_ptimes, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPostSwitchDefaultPtime: {
      const rest = [
        ...state.post_load_switch.default_ptimes.new.slice(0, action.payload.index),
        null,
        ...state.post_load_switch.default_ptimes.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        post_load_switch: {
          ...state.post_load_switch, default_ptimes: {...state.post_load_switch.default_ptimes, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPostSwitchDefaultPtime: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.post_load_switch.default_ptimes.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.post_load_switch.default_ptimes.new.slice(0, action.payload.index),
          null,
          ...state.post_load_switch.default_ptimes.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        post_load_switch: <IpostSwitcheModule>{
          ...state.post_load_switch, default_ptimes: {...state.post_load_switch.default_ptimes, [data.id]: data, new: rest},
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

