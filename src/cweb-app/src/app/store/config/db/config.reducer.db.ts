
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.db';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetDb:
    case ConfigActionTypes.UpdateDbParameter:
    case ConfigActionTypes.SwitchDbParameter:
    case ConfigActionTypes.AddDbParameter:
    case ConfigActionTypes.DelDbParameter: {
      return {...state,
        db: {
          ...state.db,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotDbError: {
      return {
        ...state,
        db: {
          ...state.db,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetDb: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          db: {...state.db, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.db) {
        state.db = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        db: {
          ...state.db,
          settings: {...state.db.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelDbParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.db.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.db.settings;

      return {
        ...state,
        db: {
          ...state.db, settings: {...rest, new: state.db.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchDbParameter:
    case ConfigActionTypes.StoreUpdateDbParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        db: <IsimpleModule>{
          ...state.db, settings: {...state.db.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewDbParameter: {
      const rest = [
        ...state.db.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        db: {
          ...state.db, settings: {...state.db.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewDbParameter: {
      const rest = [
        ...state.db.settings.new.slice(0, action.payload.index),
        null,
        ...state.db.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        db: {
          ...state.db, settings: {...state.db.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddDbParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.db.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.db.settings.new.slice(0, action.payload.index),
          null,
          ...state.db.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        db: <IsimpleModule>{
          ...state.db, settings: {...state.db.settings, [data.id]: data, new: rest},
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

