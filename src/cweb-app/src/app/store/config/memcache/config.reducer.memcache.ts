
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.memcache';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetMemcache:
    case ConfigActionTypes.UpdateMemcacheParameter:
    case ConfigActionTypes.SwitchMemcacheParameter:
    case ConfigActionTypes.AddMemcacheParameter:
    case ConfigActionTypes.DelMemcacheParameter: {
      return {...state,
        memcache: {
          ...state.memcache,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotMemcacheError: {
      return {
        ...state,
        memcache: {
          ...state.memcache,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetMemcache: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          memcache: {...state.memcache, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.memcache) {
        state.memcache = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        memcache: {
          ...state.memcache,
          settings: {...state.memcache.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelMemcacheParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.memcache.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.memcache.settings;

      return {
        ...state,
        memcache: {
          ...state.memcache, settings: {...rest, new: state.memcache.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchMemcacheParameter:
    case ConfigActionTypes.StoreUpdateMemcacheParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        memcache: <IsimpleModule>{
          ...state.memcache, settings: {...state.memcache.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewMemcacheParameter: {
      const rest = [
        ...state.memcache.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        memcache: {
          ...state.memcache, settings: {...state.memcache.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewMemcacheParameter: {
      const rest = [
        ...state.memcache.settings.new.slice(0, action.payload.index),
        null,
        ...state.memcache.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        memcache: {
          ...state.memcache, settings: {...state.memcache.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddMemcacheParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.memcache.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.memcache.settings.new.slice(0, action.payload.index),
          null,
          ...state.memcache.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        memcache: <IsimpleModule>{
          ...state.memcache, settings: {...state.memcache.settings, [data.id]: data, new: rest},
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

