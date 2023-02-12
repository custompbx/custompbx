
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.pocketsphinx';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetPocketsphinx:
    case ConfigActionTypes.UpdatePocketsphinxParameter:
    case ConfigActionTypes.SwitchPocketsphinxParameter:
    case ConfigActionTypes.AddPocketsphinxParameter:
    case ConfigActionTypes.DelPocketsphinxParameter: {
      return {...state,
        pocketsphinx: {
          ...state.pocketsphinx,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotPocketsphinxError: {
      return {
        ...state,
        pocketsphinx: {
          ...state.pocketsphinx,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetPocketsphinx: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          pocketsphinx: {...state.pocketsphinx, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.pocketsphinx) {
        state.pocketsphinx = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        pocketsphinx: {
          ...state.pocketsphinx,
          settings: {...state.pocketsphinx.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPocketsphinxParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.pocketsphinx.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.pocketsphinx.settings;

      return {
        ...state,
        pocketsphinx: {
          ...state.pocketsphinx, settings: {...rest, new: state.pocketsphinx.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchPocketsphinxParameter:
    case ConfigActionTypes.StoreUpdatePocketsphinxParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        pocketsphinx: <IsimpleModule>{
          ...state.pocketsphinx, settings: {...state.pocketsphinx.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPocketsphinxParameter: {
      const rest = [
        ...state.pocketsphinx.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        pocketsphinx: {
          ...state.pocketsphinx, settings: {...state.pocketsphinx.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPocketsphinxParameter: {
      const rest = [
        ...state.pocketsphinx.settings.new.slice(0, action.payload.index),
        null,
        ...state.pocketsphinx.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        pocketsphinx: {
          ...state.pocketsphinx, settings: {...state.pocketsphinx.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPocketsphinxParameter: {
      const data = action.payload.response.data;
      let rest = [...state.pocketsphinx.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.pocketsphinx.settings.new.slice(0, action.payload.index),
          null,
          ...state.pocketsphinx.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        pocketsphinx: <IsimpleModule>{
          ...state.pocketsphinx, settings: {...state.pocketsphinx.settings, [data.id]: data, new: rest},
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

