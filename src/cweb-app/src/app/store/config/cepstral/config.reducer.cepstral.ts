
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.cepstral';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetCepstral:
    case ConfigActionTypes.UpdateCepstralParameter:
    case ConfigActionTypes.SwitchCepstralParameter:
    case ConfigActionTypes.AddCepstralParameter:
    case ConfigActionTypes.DelCepstralParameter: {
      return {...state,
        cepstral: {
          ...state.cepstral,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotCepstralError: {
      return {
        ...state,
        cepstral: {
          ...state.cepstral,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetCepstral: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          cepstral: {...state.cepstral, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.cepstral) {
        state.cepstral = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        cepstral: {
          ...state.cepstral,
          settings: {...state.cepstral.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelCepstralParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.cepstral.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.cepstral.settings;

      return {
        ...state,
        cepstral: {
          ...state.cepstral, settings: {...rest, new: state.cepstral.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchCepstralParameter:
    case ConfigActionTypes.StoreUpdateCepstralParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        cepstral: <IsimpleModule>{
          ...state.cepstral, settings: {...state.cepstral.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewCepstralParameter: {
      const rest = [
        ...state.cepstral.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        cepstral: {
          ...state.cepstral, settings: {...state.cepstral.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewCepstralParameter: {
      const rest = [
        ...state.cepstral.settings.new.slice(0, action.payload.index),
        null,
        ...state.cepstral.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        cepstral: {
          ...state.cepstral, settings: {...state.cepstral.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddCepstralParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.cepstral.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.cepstral.settings.new.slice(0, action.payload.index),
          null,
          ...state.cepstral.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        cepstral: <IsimpleModule>{
          ...state.cepstral, settings: {...state.cepstral.settings, [data.id]: data, new: rest},
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

