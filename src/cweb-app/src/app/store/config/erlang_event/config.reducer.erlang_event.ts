
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.erlang_event';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetErlangEvent:
    case ConfigActionTypes.UpdateErlangEventParameter:
    case ConfigActionTypes.SwitchErlangEventParameter:
    case ConfigActionTypes.AddErlangEventParameter:
    case ConfigActionTypes.DelErlangEventParameter: {
      return {...state,
        erlang_event: {
          ...state.erlang_event,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotErlangEventError: {
      return {
        ...state,
        erlang_event: {
          ...state.erlang_event,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetErlangEvent: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          erlang_event: {...state.erlang_event, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.erlang_event) {
        state.erlang_event = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        erlang_event: {
          ...state.erlang_event,
          settings: {...state.erlang_event.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelErlangEventParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.erlang_event.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.erlang_event.settings;

      return {
        ...state,
        erlang_event: {
          ...state.erlang_event, settings: {...rest, new: state.erlang_event.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchErlangEventParameter:
    case ConfigActionTypes.StoreUpdateErlangEventParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        erlang_event: <IsimpleModule>{
          ...state.erlang_event, settings: {...state.erlang_event.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewErlangEventParameter: {
      const rest = [
        ...state.erlang_event.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        erlang_event: {
          ...state.erlang_event, settings: {...state.erlang_event.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewErlangEventParameter: {
      const rest = [
        ...state.erlang_event.settings.new.slice(0, action.payload.index),
        null,
        ...state.erlang_event.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        erlang_event: {
          ...state.erlang_event, settings: {...state.erlang_event.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddErlangEventParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.erlang_event.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.erlang_event.settings.new.slice(0, action.payload.index),
          null,
          ...state.erlang_event.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        erlang_event: <IsimpleModule>{
          ...state.erlang_event, settings: {...state.erlang_event.settings, [data.id]: data, new: rest},
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

