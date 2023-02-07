
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.sangoma_codec';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetSangomaCodec:
    case ConfigActionTypes.UpdateSangomaCodecParameter:
    case ConfigActionTypes.SwitchSangomaCodecParameter:
    case ConfigActionTypes.AddSangomaCodecParameter:
    case ConfigActionTypes.DelSangomaCodecParameter: {
      return {...state,
        sangoma_codec: {
          ...state.sangoma_codec,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotSangomaCodecError: {
      return {
        ...state,
        sangoma_codec: {
          ...state.sangoma_codec,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetSangomaCodec: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          sangoma_codec: {...state.sangoma_codec, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.sangoma_codec) {
        state.sangoma_codec = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        sangoma_codec: {
          ...state.sangoma_codec,
          settings: {...state.sangoma_codec.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelSangomaCodecParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.sangoma_codec.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.sangoma_codec.settings;

      return {
        ...state,
        sangoma_codec: {
          ...state.sangoma_codec, settings: {...rest, new: state.sangoma_codec.settings?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchSangomaCodecParameter:
    case ConfigActionTypes.StoreUpdateSangomaCodecParameter: {
      const data = action.payload.response.data;

      return {
        ...state,
        sangoma_codec: <IsimpleModule>{
          ...state.sangoma_codec, settings: {...state.sangoma_codec.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewSangomaCodecParameter: {
      const rest = [
        ...state.sangoma_codec.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        sangoma_codec: {
          ...state.sangoma_codec, settings: {...state.sangoma_codec.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewSangomaCodecParameter: {
      const rest = [
        ...state.sangoma_codec.settings.new.slice(0, action.payload.index),
        null,
        ...state.sangoma_codec.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        sangoma_codec: {
          ...state.sangoma_codec, settings: {...state.sangoma_codec.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddSangomaCodecParameter: {
      const data = action.payload.response.data;
      let rest = [...state.sangoma_codec.settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.sangoma_codec.settings.new.slice(0, action.payload.index),
          null,
          ...state.sangoma_codec.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        sangoma_codec: <IsimpleModule>{
          ...state.sangoma_codec, settings: {...state.sangoma_codec.settings, [data.id]: data, new: rest},
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

