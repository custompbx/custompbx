
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.tts_commandline';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetTtsCommandline:
    case ConfigActionTypes.UpdateTtsCommandlineParameter:
    case ConfigActionTypes.SwitchTtsCommandlineParameter:
    case ConfigActionTypes.AddTtsCommandlineParameter:
    case ConfigActionTypes.DelTtsCommandlineParameter: {
      return {...state,
        tts_commandline: {
          ...state.tts_commandline,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotTtsCommandlineError: {
      return {
        ...state,
        tts_commandline: {
          ...state.tts_commandline,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetTtsCommandline: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          tts_commandline: {...state.tts_commandline, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.tts_commandline) {
        state.tts_commandline = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        tts_commandline: {
          ...state.tts_commandline,
          settings: {...state.tts_commandline.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelTtsCommandlineParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.tts_commandline.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.tts_commandline.settings;

      return {
        ...state,
        tts_commandline: {
          ...state.tts_commandline, settings: {...rest, new: state.tts_commandline.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchTtsCommandlineParameter:
    case ConfigActionTypes.StoreUpdateTtsCommandlineParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        tts_commandline: <IsimpleModule>{
          ...state.tts_commandline, settings: {...state.tts_commandline.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewTtsCommandlineParameter: {
      const rest = [
        ...state.tts_commandline.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        tts_commandline: {
          ...state.tts_commandline, settings: {...state.tts_commandline.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewTtsCommandlineParameter: {
      const rest = [
        ...state.tts_commandline.settings.new.slice(0, action.payload.index),
        null,
        ...state.tts_commandline.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        tts_commandline: {
          ...state.tts_commandline, settings: {...state.tts_commandline.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddTtsCommandlineParameter: {
      const data = action.payload.response.data;
      let rest = [...state.tts_commandline.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.tts_commandline.settings.new.slice(0, action.payload.index),
          null,
          ...state.tts_commandline.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        tts_commandline: <IsimpleModule>{
          ...state.tts_commandline, settings: {...state.tts_commandline.settings, [data.id]: data, new: rest},
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

