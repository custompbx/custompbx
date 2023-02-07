
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.nibblebill';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetNibblebill:
    case ConfigActionTypes.UpdateNibblebillParameter:
    case ConfigActionTypes.SwitchNibblebillParameter:
    case ConfigActionTypes.AddNibblebillParameter:
    case ConfigActionTypes.DelNibblebillParameter: {
      return {...state,
        nibblebill: {
          ...state.nibblebill,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotNibblebillError: {
      return {
        ...state,
        nibblebill: {
          ...state.nibblebill,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetNibblebill: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          nibblebill: {...state.nibblebill, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.nibblebill) {
        state.nibblebill = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        nibblebill: {
          ...state.nibblebill,
          settings: {...state.nibblebill.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelNibblebillParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.nibblebill.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.nibblebill.settings;

      return {
        ...state,
        nibblebill: {
          ...state.nibblebill, settings: {...rest, new: state.nibblebill.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchNibblebillParameter:
    case ConfigActionTypes.StoreUpdateNibblebillParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        nibblebill: <IsimpleModule>{
          ...state.nibblebill, settings: {...state.nibblebill.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewNibblebillParameter: {
      const rest = [
        ...state.nibblebill.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        nibblebill: {
          ...state.nibblebill, settings: {...state.nibblebill.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewNibblebillParameter: {
      const rest = [
        ...state.nibblebill.settings.new.slice(0, action.payload.index),
        null,
        ...state.nibblebill.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        nibblebill: {
          ...state.nibblebill, settings: {...state.nibblebill.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddNibblebillParameter: {
      const data = action.payload.response.data;
      let rest = [...state.nibblebill.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.nibblebill.settings.new.slice(0, action.payload.index),
          null,
          ...state.nibblebill.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        nibblebill: <IsimpleModule>{
          ...state.nibblebill, settings: {...state.nibblebill.settings, [data.id]: data, new: rest},
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

