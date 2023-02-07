
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.dialplan_directory';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetDialplanDirectory:
    case ConfigActionTypes.UpdateDialplanDirectoryParameter:
    case ConfigActionTypes.SwitchDialplanDirectoryParameter:
    case ConfigActionTypes.AddDialplanDirectoryParameter:
    case ConfigActionTypes.DelDialplanDirectoryParameter: {
      return {...state,
        dialplan_directory: {
          ...state.dialplan_directory,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotDialplanDirectoryError: {
      return {
        ...state,
        dialplan_directory: {
          ...state.dialplan_directory,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetDialplanDirectory: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          dialplan_directory: {...state.dialplan_directory, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.dialplan_directory) {
        state.dialplan_directory = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        dialplan_directory: {
          ...state.dialplan_directory,
          settings: {...state.dialplan_directory.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelDialplanDirectoryParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.dialplan_directory.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.dialplan_directory.settings;

      return {
        ...state,
        dialplan_directory: {
          ...state.dialplan_directory, settings: {...rest, new: state.dialplan_directory.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchDialplanDirectoryParameter:
    case ConfigActionTypes.StoreUpdateDialplanDirectoryParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        dialplan_directory: <IsimpleModule>{
          ...state.dialplan_directory, settings: {...state.dialplan_directory.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewDialplanDirectoryParameter: {
      const rest = [
        ...state.dialplan_directory.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        dialplan_directory: {
          ...state.dialplan_directory, settings: {...state.dialplan_directory.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewDialplanDirectoryParameter: {
      const rest = [
        ...state.dialplan_directory.settings.new.slice(0, action.payload.index),
        null,
        ...state.dialplan_directory.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        dialplan_directory: {
          ...state.dialplan_directory, settings: {...state.dialplan_directory.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddDialplanDirectoryParameter: {
      const data = action.payload.response.data;
      let rest = [...state.dialplan_directory.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.dialplan_directory.settings.new.slice(0, action.payload.index),
          null,
          ...state.dialplan_directory.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        dialplan_directory: <IsimpleModule>{
          ...state.dialplan_directory, settings: {...state.dialplan_directory.settings, [data.id]: data, new: rest},
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

