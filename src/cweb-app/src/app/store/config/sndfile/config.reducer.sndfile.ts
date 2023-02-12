
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.sndfile';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetSndfile:
    case ConfigActionTypes.UpdateSndfileParameter:
    case ConfigActionTypes.SwitchSndfileParameter:
    case ConfigActionTypes.AddSndfileParameter:
    case ConfigActionTypes.DelSndfileParameter: {
      return {...state,
        sndfile: {
          ...state.sndfile,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotSndfileError: {
      return {
        ...state,
        sndfile: {
          ...state.sndfile,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetSndfile: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          sndfile: {...state.sndfile, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.sndfile) {
        state.sndfile = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        sndfile: {
          ...state.sndfile,
          settings: {...state.sndfile.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelSndfileParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.sndfile.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.sndfile.settings;

      return {
        ...state,
        sndfile: {
          ...state.sndfile, settings: {...rest, new: state.sndfile.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchSndfileParameter:
    case ConfigActionTypes.StoreUpdateSndfileParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        sndfile: <IsimpleModule>{
          ...state.sndfile, settings: {...state.sndfile.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewSndfileParameter: {
      const rest = [
        ...state.sndfile.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        sndfile: {
          ...state.sndfile, settings: {...state.sndfile.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewSndfileParameter: {
      const rest = [
        ...state.sndfile.settings.new.slice(0, action.payload.index),
        null,
        ...state.sndfile.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        sndfile: {
          ...state.sndfile, settings: {...state.sndfile.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddSndfileParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.sndfile.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.sndfile.settings.new.slice(0, action.payload.index),
          null,
          ...state.sndfile.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        sndfile: <IsimpleModule>{
          ...state.sndfile, settings: {...state.sndfile.settings, [data.id]: data, new: rest},
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

