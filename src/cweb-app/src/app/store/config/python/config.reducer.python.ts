
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.python';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetPython:
    case ConfigActionTypes.UpdatePythonParameter:
    case ConfigActionTypes.SwitchPythonParameter:
    case ConfigActionTypes.AddPythonParameter:
    case ConfigActionTypes.DelPythonParameter: {
      return {...state,
        python: {
          ...state.python,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotPythonError: {
      return {
        ...state,
        python: {
          ...state.python,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetPython: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          python: {...state.python, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.python) {
        state.python = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        python: {
          ...state.python,
          settings: {...state.python.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPythonParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.python.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.python.settings;

      return {
        ...state,
        python: {
          ...state.python, settings: {...rest, new: state.python.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchPythonParameter:
    case ConfigActionTypes.StoreUpdatePythonParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        python: <IsimpleModule>{
          ...state.python, settings: {...state.python.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPythonParameter: {
      const rest = [
        ...state.python.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        python: {
          ...state.python, settings: {...state.python.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPythonParameter: {
      const rest = [
        ...state.python.settings.new.slice(0, action.payload.index),
        null,
        ...state.python.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        python: {
          ...state.python, settings: {...state.python.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPythonParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.python.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.python.settings.new.slice(0, action.payload.index),
          null,
          ...state.python.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        python: <IsimpleModule>{
          ...state.python, settings: {...state.python.settings, [data.id]: data, new: rest},
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

