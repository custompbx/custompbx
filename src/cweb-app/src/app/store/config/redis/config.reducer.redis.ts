
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.redis';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetRedis:
    case ConfigActionTypes.UpdateRedisParameter:
    case ConfigActionTypes.SwitchRedisParameter:
    case ConfigActionTypes.AddRedisParameter:
    case ConfigActionTypes.DelRedisParameter: {
      return {...state,
        redis: {
          ...state.redis,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotRedisError: {
      return {
        ...state,
        redis: {
          ...state.redis,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetRedis: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          redis: {...state.redis, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.redis) {
        state.redis = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        redis: {
          ...state.redis,
          settings: {...state.redis.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelRedisParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.redis.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.redis.settings;

      return {
        ...state,
        redis: {
          ...state.redis, settings: {...rest, new: state.redis.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchRedisParameter:
    case ConfigActionTypes.StoreUpdateRedisParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        redis: <IsimpleModule>{
          ...state.redis, settings: {...state.redis.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewRedisParameter: {
      const rest = [
        ...state.redis.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        redis: {
          ...state.redis, settings: {...state.redis.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewRedisParameter: {
      const rest = [
        ...state.redis.settings.new.slice(0, action.payload.index),
        null,
        ...state.redis.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        redis: {
          ...state.redis, settings: {...state.redis.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddRedisParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.redis.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.redis.settings.new.slice(0, action.payload.index),
          null,
          ...state.redis.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        redis: <IsimpleModule>{
          ...state.redis, settings: {...state.redis.settings, [data.id]: data, new: rest},
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

