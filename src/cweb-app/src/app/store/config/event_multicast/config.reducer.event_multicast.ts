
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.event_multicast';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetEventMulticast:
    case ConfigActionTypes.UpdateEventMulticastParameter:
    case ConfigActionTypes.SwitchEventMulticastParameter:
    case ConfigActionTypes.AddEventMulticastParameter:
    case ConfigActionTypes.DelEventMulticastParameter: {
      return {...state,
        event_multicast: {
          ...state.event_multicast,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotEventMulticastError: {
      return {
        ...state,
        event_multicast: {
          ...state.event_multicast,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetEventMulticast: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          event_multicast: {...state.event_multicast, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.event_multicast) {
        state.event_multicast = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        event_multicast: {
          ...state.event_multicast,
          settings: {...state.event_multicast.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelEventMulticastParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.event_multicast.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.event_multicast.settings;

      return {
        ...state,
        event_multicast: {
          ...state.event_multicast, settings: {...rest, new: state.event_multicast.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchEventMulticastParameter:
    case ConfigActionTypes.StoreUpdateEventMulticastParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        event_multicast: <IsimpleModule>{
          ...state.event_multicast, settings: {...state.event_multicast.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewEventMulticastParameter: {
      const rest = [
        ...state.event_multicast.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        event_multicast: {
          ...state.event_multicast, settings: {...state.event_multicast.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewEventMulticastParameter: {
      const rest = [
        ...state.event_multicast.settings.new.slice(0, action.payload.index),
        null,
        ...state.event_multicast.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        event_multicast: {
          ...state.event_multicast, settings: {...state.event_multicast.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddEventMulticastParameter: {
      const data = action.payload.response.data;
      let rest = [...state.event_multicast.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.event_multicast.settings.new.slice(0, action.payload.index),
          null,
          ...state.event_multicast.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        event_multicast: <IsimpleModule>{
          ...state.event_multicast, settings: {...state.event_multicast.settings, [data.id]: data, new: rest},
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

