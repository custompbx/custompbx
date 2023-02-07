
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.shout';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetShout:
    case ConfigActionTypes.UpdateShoutParameter:
    case ConfigActionTypes.SwitchShoutParameter:
    case ConfigActionTypes.AddShoutParameter:
    case ConfigActionTypes.DelShoutParameter: {
      return {...state,
        shout: {
          ...state.shout,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotShoutError: {
      return {
        ...state,
        shout: {
          ...state.shout,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetShout: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          shout: {...state.shout, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.shout) {
        state.shout = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        shout: {
          ...state.shout,
          settings: {...state.shout.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelShoutParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.shout.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.shout.settings;

      return {
        ...state,
        shout: {
          ...state.shout, settings: {...rest, new: state.shout.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchShoutParameter:
    case ConfigActionTypes.StoreUpdateShoutParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        shout: <IsimpleModule>{
          ...state.shout, settings: {...state.shout.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewShoutParameter: {
      const rest = [
        ...state.shout.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        shout: {
          ...state.shout, settings: {...state.shout.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewShoutParameter: {
      const rest = [
        ...state.shout.settings.new.slice(0, action.payload.index),
        null,
        ...state.shout.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        shout: {
          ...state.shout, settings: {...state.shout.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddShoutParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.shout.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.shout.settings.new.slice(0, action.payload.index),
          null,
          ...state.shout.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        shout: <IsimpleModule>{
          ...state.shout, settings: {...state.shout.settings, [data.id]: data, new: rest},
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

