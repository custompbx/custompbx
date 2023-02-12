
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.http_cache';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetHttpCache:
    case ConfigActionTypes.UpdateHttpCacheParameter:
    case ConfigActionTypes.SwitchHttpCacheParameter:
    case ConfigActionTypes.AddHttpCacheParameter:
    case ConfigActionTypes.DelHttpCacheParameter: {
      return {...state,
        http_cache: {
          ...state.http_cache,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotHttpCacheError: {
      return {
        ...state,
        http_cache: {
          ...state.http_cache,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetHttpCache: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          http_cache: {...state.http_cache, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.http_cache) {
        state.http_cache = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        http_cache: {
          ...state.http_cache,
          settings: {...state.http_cache.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelHttpCacheParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.http_cache.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.http_cache.settings;

      return {
        ...state,
        http_cache: {
          ...state.http_cache, settings: {...rest, new: state.http_cache.settings?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchHttpCacheParameter:
    case ConfigActionTypes.StoreUpdateHttpCacheParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        http_cache: <IsimpleModule>{
          ...state.http_cache, settings: {...state.http_cache.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewHttpCacheParameter: {
      const rest = [
        ...state.http_cache.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        http_cache: {
          ...state.http_cache, settings: {...state.http_cache.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewHttpCacheParameter: {
      const rest = [
        ...state.http_cache.settings.new.slice(0, action.payload.index),
        null,
        ...state.http_cache.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        http_cache: {
          ...state.http_cache, settings: {...state.http_cache.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddHttpCacheParameter: {
      const data = action.payload.response.data;
      let rest = [...state.http_cache.settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.http_cache.settings.new.slice(0, action.payload.index),
          null,
          ...state.http_cache.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        http_cache: <IsimpleModule>{
          ...state.http_cache, settings: {...state.http_cache.settings, [data.id]: data, new: rest},
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

