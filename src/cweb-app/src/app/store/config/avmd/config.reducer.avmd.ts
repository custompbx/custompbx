
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.avmd';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetAvmd:
    case ConfigActionTypes.UpdateAvmdParameter:
    case ConfigActionTypes.SwitchAvmdParameter:
    case ConfigActionTypes.AddAvmdParameter:
    case ConfigActionTypes.DelAvmdParameter: {
      return {...state,
        avmd: {
          ...state.avmd,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotAvmdError: {
      return {
        ...state,
        avmd: {
          ...state.avmd,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetAvmd: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          avmd: {...state.avmd, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.avmd) {
        state.avmd = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        avmd: {
          ...state.avmd,
          settings: {...state.avmd.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelAvmdParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.avmd.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.avmd.settings;

      return {
        ...state,
        avmd: {
          ...state.avmd, settings: {...rest, new: state.avmd.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchAvmdParameter:
    case ConfigActionTypes.StoreUpdateAvmdParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        avmd: <IsimpleModule>{
          ...state.avmd, settings: {...state.avmd.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewAvmdParameter: {
      const rest = [
        ...state.avmd.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        avmd: {
          ...state.avmd, settings: {...state.avmd.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewAvmdParameter: {
      const rest = [
        ...state.avmd.settings.new.slice(0, action.payload.index),
        null,
        ...state.avmd.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        avmd: {
          ...state.avmd, settings: {...state.avmd.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddAvmdParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.avmd.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.avmd.settings.new.slice(0, action.payload.index),
          null,
          ...state.avmd.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        avmd: <IsimpleModule>{
          ...state.avmd, settings: {...state.avmd.settings, [data.id]: data, new: rest},
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

