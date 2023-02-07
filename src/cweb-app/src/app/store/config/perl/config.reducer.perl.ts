
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.perl';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetPerl:
    case ConfigActionTypes.UpdatePerlParameter:
    case ConfigActionTypes.SwitchPerlParameter:
    case ConfigActionTypes.AddPerlParameter:
    case ConfigActionTypes.DelPerlParameter: {
      return {...state,
        perl: {
          ...state.perl,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotPerlError: {
      return {
        ...state,
        perl: {
          ...state.perl,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetPerl: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          perl: {...state.perl, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.perl) {
        state.perl = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        perl: {
          ...state.perl,
          settings: {...state.perl.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPerlParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.perl.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.perl.settings;

      return {
        ...state,
        perl: {
          ...state.perl, settings: {...rest, new: state.perl.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchPerlParameter:
    case ConfigActionTypes.StoreUpdatePerlParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        perl: <IsimpleModule>{
          ...state.perl, settings: {...state.perl.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPerlParameter: {
      const rest = [
        ...state.perl.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        perl: {
          ...state.perl, settings: {...state.perl.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPerlParameter: {
      const rest = [
        ...state.perl.settings.new.slice(0, action.payload.index),
        null,
        ...state.perl.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        perl: {
          ...state.perl, settings: {...state.perl.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPerlParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.perl.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.perl.settings.new.slice(0, action.payload.index),
          null,
          ...state.perl.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        perl: <IsimpleModule>{
          ...state.perl, settings: {...state.perl.settings, [data.id]: data, new: rest},
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

