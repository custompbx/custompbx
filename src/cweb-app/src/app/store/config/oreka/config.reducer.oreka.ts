
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.oreka';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetOreka:
    case ConfigActionTypes.UpdateOrekaParameter:
    case ConfigActionTypes.SwitchOrekaParameter:
    case ConfigActionTypes.AddOrekaParameter:
    case ConfigActionTypes.DelOrekaParameter: {
      return {...state,
        oreka: {
          ...state.oreka,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotOrekaError: {
      return {
        ...state,
        oreka: {
          ...state.oreka,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetOreka: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          oreka: {...state.oreka, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.oreka) {
        state.oreka = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        oreka: {
          ...state.oreka,
          settings: {...state.oreka.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelOrekaParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.oreka.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.oreka.settings;

      return {
        ...state,
        oreka: {
          ...state.oreka, settings: {...rest, new: state.oreka.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchOrekaParameter:
    case ConfigActionTypes.StoreUpdateOrekaParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        oreka: <IsimpleModule>{
          ...state.oreka, settings: {...state.oreka.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOrekaParameter: {
      const rest = [
        ...state.oreka.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        oreka: {
          ...state.oreka, settings: {...state.oreka.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOrekaParameter: {
      const rest = [
        ...state.oreka.settings.new.slice(0, action.payload.index),
        null,
        ...state.oreka.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        oreka: {
          ...state.oreka, settings: {...state.oreka.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOrekaParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.oreka.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.oreka.settings.new.slice(0, action.payload.index),
          null,
          ...state.oreka.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        oreka: <IsimpleModule>{
          ...state.oreka, settings: {...state.oreka.settings, [data.id]: data, new: rest},
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

