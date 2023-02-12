
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.amrwb';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetAmrwb:
    case ConfigActionTypes.UpdateAmrwbParameter:
    case ConfigActionTypes.SwitchAmrwbParameter:
    case ConfigActionTypes.AddAmrwbParameter:
    case ConfigActionTypes.DelAmrwbParameter: {
      return {...state,
        amrwb: {
          ...state.amrwb,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotAmrwbError: {
      return {
        ...state,
        amrwb: {
          ...state.amrwb,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetAmrwb: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          amrwb: {...state.amrwb, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.amrwb) {
        state.amrwb = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        amrwb: {
          ...state.amrwb,
          settings: {...state.amrwb.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelAmrwbParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.amrwb.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.amrwb.settings;

      return {
        ...state,
        amrwb: {
          ...state.amrwb, settings: {...rest, new: state.amrwb.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchAmrwbParameter:
    case ConfigActionTypes.StoreUpdateAmrwbParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        amrwb: <IsimpleModule>{
          ...state.amrwb, settings: {...state.amrwb.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewAmrwbParameter: {
      const rest = [
        ...state.amrwb.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        amrwb: {
          ...state.amrwb, settings: {...state.amrwb.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewAmrwbParameter: {
      const rest = [
        ...state.amrwb.settings.new.slice(0, action.payload.index),
        null,
        ...state.amrwb.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        amrwb: {
          ...state.amrwb, settings: {...state.amrwb.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddAmrwbParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.amrwb.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.amrwb.settings.new.slice(0, action.payload.index),
          null,
          ...state.amrwb.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        amrwb: <IsimpleModule>{
          ...state.amrwb, settings: {...state.amrwb.settings, [data.id]: data, new: rest},
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

