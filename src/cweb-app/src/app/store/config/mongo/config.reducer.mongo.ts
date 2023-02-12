
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.mongo';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetMongo:
    case ConfigActionTypes.UpdateMongoParameter:
    case ConfigActionTypes.SwitchMongoParameter:
    case ConfigActionTypes.AddMongoParameter:
    case ConfigActionTypes.DelMongoParameter: {
      return {...state,
        mongo: {
          ...state.mongo,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotMongoError: {
      return {
        ...state,
        mongo: {
          ...state.mongo,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetMongo: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          mongo: {...state.mongo, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.mongo) {
        state.mongo = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        mongo: {
          ...state.mongo,
          settings: {...state.mongo.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelMongoParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.mongo.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.mongo.settings;

      return {
        ...state,
        mongo: {
          ...state.mongo, settings: {...rest, new: state.mongo.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchMongoParameter:
    case ConfigActionTypes.StoreUpdateMongoParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        mongo: <IsimpleModule>{
          ...state.mongo, settings: {...state.mongo.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewMongoParameter: {
      const rest = [
        ...state.mongo.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        mongo: {
          ...state.mongo, settings: {...state.mongo.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewMongoParameter: {
      const rest = [
        ...state.mongo.settings.new.slice(0, action.payload.index),
        null,
        ...state.mongo.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        mongo: {
          ...state.mongo, settings: {...state.mongo.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddMongoParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.mongo.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.mongo.settings.new.slice(0, action.payload.index),
          null,
          ...state.mongo.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        mongo: <IsimpleModule>{
          ...state.mongo, settings: {...state.mongo.settings, [data.id]: data, new: rest},
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

