
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.curl';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetCurl:
    case ConfigActionTypes.UpdateCurlParameter:
    case ConfigActionTypes.SwitchCurlParameter:
    case ConfigActionTypes.AddCurlParameter:
    case ConfigActionTypes.DelCurlParameter: {
      return {...state,
        curl: {
          ...state.curl,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotCurlError: {
      return {
        ...state,
        curl: {
          ...state.curl,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCurl: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          curl: {...state.curl, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.curl) {
        state.curl = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        curl: {
          ...state.curl,
          settings: {...state.curl.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCurlParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.curl.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.curl.settings;

      return {
        ...state,
        curl: {
          ...state.curl, settings: {...rest, new: state.curl.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchCurlParameter:
    case ConfigActionTypes.StoreUpdateCurlParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        curl: <IsimpleModule>{
          ...state.curl, settings: {...state.curl.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewCurlParameter: {
      const rest = [
        ...state.curl.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        curl: {
          ...state.curl, settings: {...state.curl.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewCurlParameter: {
      const rest = [
        ...state.curl.settings.new.slice(0, action.payload.index),
        null,
        ...state.curl.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        curl: {
          ...state.curl, settings: {...state.curl.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCurlParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.curl.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.curl.settings.new.slice(0, action.payload.index),
          null,
          ...state.curl.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        curl: <IsimpleModule>{
          ...state.curl, settings: {...state.curl.settings, [data.id]: data, new: rest},
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

