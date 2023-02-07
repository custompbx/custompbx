import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.alsa';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetAlsa:
    case ConfigActionTypes.UpdateAlsaParameter:
    case ConfigActionTypes.SwitchAlsaParameter:
    case ConfigActionTypes.AddAlsaParameter:
    case ConfigActionTypes.DelAlsaParameter: {
      return {...state,
        alsa: {
          ...state.alsa,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotAlsaError: {
      return {
        ...state,
        alsa: {
          ...state.alsa,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetAlsa: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          alsa: {...state.alsa, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.alsa) {
        state.alsa = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        alsa: {
          ...state.alsa,
          settings: {...state.alsa.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelAlsaParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.alsa.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.alsa.settings;

      return {
        ...state,
        alsa: {
          ...state.alsa, settings: {...rest, new: state.alsa.settings?.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchAlsaParameter:
    case ConfigActionTypes.StoreUpdateAlsaParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        alsa: <IsimpleModule>{
          ...state.alsa, settings: {...state.alsa.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewAlsaParameter: {
      const rest = [
        ...state.alsa.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        alsa: {
          ...state.alsa, settings: {...state.alsa.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewAlsaParameter: {
      const rest = [
        ...state.alsa.settings.new.slice(0, action.payload.index),
        null,
        ...state.alsa.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        alsa: {
          ...state.alsa, settings: {...state.alsa.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddAlsaParameter: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.alsa.settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.alsa.settings.new.slice(0, action.payload.index),
          null,
          ...state.alsa.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        alsa: <IsimpleModule>{
          ...state.alsa, settings: {...state.alsa.settings, [data.id]: data, new: rest},
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

