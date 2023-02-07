
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.cidlookup';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetCidlookup:
    case ConfigActionTypes.UpdateCidlookupParameter:
    case ConfigActionTypes.SwitchCidlookupParameter:
    case ConfigActionTypes.AddCidlookupParameter:
    case ConfigActionTypes.DelCidlookupParameter: {
      return {...state,
        cidlookup: {
          ...state.cidlookup,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotCidlookupError: {
      return {
        ...state,
        cidlookup: {
          ...state.cidlookup,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCidlookup: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          cidlookup: {...state.cidlookup, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.cidlookup) {
        state.cidlookup = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        cidlookup: {
          ...state.cidlookup,
          settings: {...state.cidlookup.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCidlookupParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.cidlookup.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.cidlookup.settings;

      return {
        ...state,
        cidlookup: {
          ...state.cidlookup, settings: {...rest, new: state.cidlookup.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchCidlookupParameter:
    case ConfigActionTypes.StoreUpdateCidlookupParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        cidlookup: <IsimpleModule>{
          ...state.cidlookup, settings: {...state.cidlookup.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewCidlookupParameter: {
      const rest = [
        ...state.cidlookup.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        cidlookup: {
          ...state.cidlookup, settings: {...state.cidlookup.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewCidlookupParameter: {
      const rest = [
        ...state.cidlookup.settings.new.slice(0, action.payload.index),
        null,
        ...state.cidlookup.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        cidlookup: {
          ...state.cidlookup, settings: {...state.cidlookup.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCidlookupParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.cidlookup.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.cidlookup.settings.new.slice(0, action.payload.index),
          null,
          ...state.cidlookup.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        cidlookup: <IsimpleModule>{
          ...state.cidlookup, settings: {...state.cidlookup.settings, [data.id]: data, new: rest},
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

