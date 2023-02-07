
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.cdr_mongodb';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetCdrMongodb:
    case ConfigActionTypes.UpdateCdrMongodbParameter:
    case ConfigActionTypes.SwitchCdrMongodbParameter:
    case ConfigActionTypes.AddCdrMongodbParameter:
    case ConfigActionTypes.DelCdrMongodbParameter: {
      return {...state,
        cdr_mongodb: {
          ...state.cdr_mongodb,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotCdrMongodbError: {
      return {
        ...state,
        cdr_mongodb: {
          ...state.cdr_mongodb,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCdrMongodb: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          cdr_mongodb: {...state.cdr_mongodb, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.cdr_mongodb) {
        state.cdr_mongodb = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        cdr_mongodb: {
          ...state.cdr_mongodb,
          settings: {...state.cdr_mongodb.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCdrMongodbParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.cdr_mongodb.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.cdr_mongodb.settings;

      return {
        ...state,
        cdr_mongodb: {
          ...state.cdr_mongodb, settings: {...rest, new: state.cdr_mongodb.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchCdrMongodbParameter:
    case ConfigActionTypes.StoreUpdateCdrMongodbParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        cdr_mongodb: <IsimpleModule>{
          ...state.cdr_mongodb, settings: {...state.cdr_mongodb.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewCdrMongodbParameter: {
      const rest = [
        ...state.cdr_mongodb.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        cdr_mongodb: {
          ...state.cdr_mongodb, settings: {...state.cdr_mongodb.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewCdrMongodbParameter: {
      const rest = [
        ...state.cdr_mongodb.settings.new.slice(0, action.payload.index),
        null,
        ...state.cdr_mongodb.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        cdr_mongodb: {
          ...state.cdr_mongodb, settings: {...state.cdr_mongodb.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCdrMongodbParameter: {
      const data = action.payload.response.data;
      let rest = [...state.cdr_mongodb.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.cdr_mongodb.settings.new.slice(0, action.payload.index),
          null,
          ...state.cdr_mongodb.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        cdr_mongodb: <IsimpleModule>{
          ...state.cdr_mongodb, settings: {...state.cdr_mongodb.settings, [data.id]: data, new: rest},
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

