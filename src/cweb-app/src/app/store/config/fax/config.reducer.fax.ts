
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.fax';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetFax:
    case ConfigActionTypes.UpdateFaxParameter:
    case ConfigActionTypes.SwitchFaxParameter:
    case ConfigActionTypes.AddFaxParameter:
    case ConfigActionTypes.DelFaxParameter: {
      return {...state,
        fax: {
          ...state.fax,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotFaxError: {
      return {
        ...state,
        fax: {
          ...state.fax,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetFax: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          fax: {...state.fax, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.fax) {
        state.fax = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        fax: {
          ...state.fax,
          settings: {...state.fax.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelFaxParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.fax.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.fax.settings;

      return {
        ...state,
        fax: {
          ...state.fax, settings: {...rest, new: state.fax.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchFaxParameter:
    case ConfigActionTypes.StoreUpdateFaxParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        fax: <IsimpleModule>{
          ...state.fax, settings: {...state.fax.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewFaxParameter: {
      const rest = [
        ...state.fax.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        fax: {
          ...state.fax, settings: {...state.fax.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewFaxParameter: {
      const rest = [
        ...state.fax.settings.new.slice(0, action.payload.index),
        null,
        ...state.fax.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        fax: {
          ...state.fax, settings: {...state.fax.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddFaxParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.fax.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.fax.settings.new.slice(0, action.payload.index),
          null,
          ...state.fax.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        fax: <IsimpleModule>{
          ...state.fax, settings: {...state.fax.settings, [data.id]: data, new: rest},
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

