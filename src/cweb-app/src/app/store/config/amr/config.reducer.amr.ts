
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.amr';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetAmr:
    case ConfigActionTypes.UpdateAmrParameter:
    case ConfigActionTypes.SwitchAmrParameter:
    case ConfigActionTypes.AddAmrParameter:
    case ConfigActionTypes.DelAmrParameter: {
      return {...state,
        amr: {
          ...state.amr,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotAmrError: {
      return {
        ...state,
        amr: {
          ...state.amr,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetAmr: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          amr: {...state.amr, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.amr) {
        state.amr = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        amr: {
          ...state.amr,
          settings: {...state.amr.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelAmrParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.amr.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.amr.settings;

      return {
        ...state,
        amr: {
          ...state.amr, settings: {...rest, new: state.amr.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchAmrParameter:
    case ConfigActionTypes.StoreUpdateAmrParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        amr: <IsimpleModule>{
          ...state.amr, settings: {...state.amr.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewAmrParameter: {
      const rest = [
        ...state.amr.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        amr: {
          ...state.amr, settings: {...state.amr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewAmrParameter: {
      const rest = [
        ...state.amr.settings.new.slice(0, action.payload.index),
        null,
        ...state.amr.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        amr: {
          ...state.amr, settings: {...state.amr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddAmrParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.amr.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.amr.settings.new.slice(0, action.payload.index),
          null,
          ...state.amr.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        amr: <IsimpleModule>{
          ...state.amr, settings: {...state.amr.settings, [data.id]: data, new: rest},
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

