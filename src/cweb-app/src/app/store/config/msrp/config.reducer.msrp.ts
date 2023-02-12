
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.msrp';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetMsrp:
    case ConfigActionTypes.UpdateMsrpParameter:
    case ConfigActionTypes.SwitchMsrpParameter:
    case ConfigActionTypes.AddMsrpParameter:
    case ConfigActionTypes.DelMsrpParameter: {
      return {...state,
        msrp: {
          ...state.msrp,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotMsrpError: {
      return {
        ...state,
        msrp: {
          ...state.msrp,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetMsrp: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          msrp: {...state.msrp, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.msrp) {
        state.msrp = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        msrp: {
          ...state.msrp,
          settings: {...state.msrp.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelMsrpParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.msrp.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.msrp.settings;

      return {
        ...state,
        msrp: {
          ...state.msrp, settings: {...rest, new: state.msrp.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchMsrpParameter:
    case ConfigActionTypes.StoreUpdateMsrpParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        msrp: <IsimpleModule>{
          ...state.msrp, settings: {...state.msrp.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewMsrpParameter: {
      const rest = [
        ...state.msrp.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        msrp: {
          ...state.msrp, settings: {...state.msrp.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewMsrpParameter: {
      const rest = [
        ...state.msrp.settings.new.slice(0, action.payload.index),
        null,
        ...state.msrp.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        msrp: {
          ...state.msrp, settings: {...state.msrp.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddMsrpParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.msrp.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.msrp.settings.new.slice(0, action.payload.index),
          null,
          ...state.msrp.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        msrp: <IsimpleModule>{
          ...state.msrp, settings: {...state.msrp.settings, [data.id]: data, new: rest},
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

