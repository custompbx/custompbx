
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.easyroute';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetEasyroute:
    case ConfigActionTypes.UpdateEasyrouteParameter:
    case ConfigActionTypes.SwitchEasyrouteParameter:
    case ConfigActionTypes.AddEasyrouteParameter:
    case ConfigActionTypes.DelEasyrouteParameter: {
      return {...state,
        easyroute: {
          ...state.easyroute,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotEasyrouteError: {
      return {
        ...state,
        easyroute: {
          ...state.easyroute,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetEasyroute: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          easyroute: {...state.easyroute, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.easyroute) {
        state.easyroute = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        easyroute: {
          ...state.easyroute,
          settings: {...state.easyroute.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelEasyrouteParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.easyroute.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.easyroute.settings;

      return {
        ...state,
        easyroute: {
          ...state.easyroute, settings: {...rest, new: state.easyroute.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchEasyrouteParameter:
    case ConfigActionTypes.StoreUpdateEasyrouteParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        easyroute: <IsimpleModule>{
          ...state.easyroute, settings: {...state.easyroute.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewEasyrouteParameter: {
      const rest = [
        ...state.easyroute.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        easyroute: {
          ...state.easyroute, settings: {...state.easyroute.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewEasyrouteParameter: {
      const rest = [
        ...state.easyroute.settings.new.slice(0, action.payload.index),
        null,
        ...state.easyroute.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        easyroute: {
          ...state.easyroute, settings: {...state.easyroute.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddEasyrouteParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.easyroute.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.easyroute.settings.new.slice(0, action.payload.index),
          null,
          ...state.easyroute.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        easyroute: <IsimpleModule>{
          ...state.easyroute, settings: {...state.easyroute.settings, [data.id]: data, new: rest},
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

