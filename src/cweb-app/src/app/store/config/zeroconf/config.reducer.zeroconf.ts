
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.zeroconf';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetZeroconf:
    case ConfigActionTypes.UpdateZeroconfParameter:
    case ConfigActionTypes.SwitchZeroconfParameter:
    case ConfigActionTypes.AddZeroconfParameter:
    case ConfigActionTypes.DelZeroconfParameter: {
      return {...state,
        zeroconf: {
          ...state.zeroconf,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotZeroconfError: {
      return {
        ...state,
        zeroconf: {
          ...state.zeroconf,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetZeroconf: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          zeroconf: {...state.zeroconf, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.zeroconf) {
        state.zeroconf = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        zeroconf: {
          ...state.zeroconf,
          settings: {...state.zeroconf.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelZeroconfParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.zeroconf.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.zeroconf.settings;

      return {
        ...state,
        zeroconf: {
          ...state.zeroconf, settings: {...rest, new: state.zeroconf.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchZeroconfParameter:
    case ConfigActionTypes.StoreUpdateZeroconfParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        zeroconf: <IsimpleModule>{
          ...state.zeroconf, settings: {...state.zeroconf.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewZeroconfParameter: {
      const rest = [
        ...state.zeroconf.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        zeroconf: {
          ...state.zeroconf, settings: {...state.zeroconf.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewZeroconfParameter: {
      const rest = [
        ...state.zeroconf.settings.new.slice(0, action.payload.index),
        null,
        ...state.zeroconf.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        zeroconf: {
          ...state.zeroconf, settings: {...state.zeroconf.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddZeroconfParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.zeroconf.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.zeroconf.settings.new.slice(0, action.payload.index),
          null,
          ...state.zeroconf.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        zeroconf: <IsimpleModule>{
          ...state.zeroconf, settings: {...state.zeroconf.settings, [data.id]: data, new: rest},
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

