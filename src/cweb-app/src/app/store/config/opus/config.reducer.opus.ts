
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.opus';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetOpus:
    case ConfigActionTypes.UpdateOpusParameter:
    case ConfigActionTypes.SwitchOpusParameter:
    case ConfigActionTypes.AddOpusParameter:
    case ConfigActionTypes.DelOpusParameter: {
      return {...state,
        opus: {
          ...state.opus,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotOpusError: {
      return {
        ...state,
        opus: {
          ...state.opus,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetOpus: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          opus: {...state.opus, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.opus) {
        state.opus = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        opus: {
          ...state.opus,
          settings: {...state.opus.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelOpusParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.opus.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.opus.settings;

      return {
        ...state,
        opus: {
          ...state.opus, settings: {...rest, new: state.opus.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchOpusParameter:
    case ConfigActionTypes.StoreUpdateOpusParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        opus: <IsimpleModule>{
          ...state.opus, settings: {...state.opus.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOpusParameter: {
      const rest = [
        ...state.opus.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        opus: {
          ...state.opus, settings: {...state.opus.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOpusParameter: {
      const rest = [
        ...state.opus.settings.new.slice(0, action.payload.index),
        null,
        ...state.opus.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        opus: {
          ...state.opus, settings: {...state.opus.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOpusParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.opus.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.opus.settings.new.slice(0, action.payload.index),
          null,
          ...state.opus.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        opus: <IsimpleModule>{
          ...state.opus, settings: {...state.opus.settings, [data.id]: data, new: rest},
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

