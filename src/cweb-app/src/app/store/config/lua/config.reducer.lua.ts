
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.lua';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetLua:
    case ConfigActionTypes.UpdateLuaParameter:
    case ConfigActionTypes.SwitchLuaParameter:
    case ConfigActionTypes.AddLuaParameter:
    case ConfigActionTypes.DelLuaParameter: {
      return {...state,
        lua: {
          ...state.lua,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotLuaError: {
      return {
        ...state,
        lua: {
          ...state.lua,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetLua: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          lua: {...state.lua, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.lua) {
        state.lua = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        lua: {
          ...state.lua,
          settings: {...state.lua.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelLuaParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.lua.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.lua.settings;

      return {
        ...state,
        lua: {
          ...state.lua, settings: {...rest, new: state.lua.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchLuaParameter:
    case ConfigActionTypes.StoreUpdateLuaParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        lua: <IsimpleModule>{
          ...state.lua, settings: {...state.lua.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewLuaParameter: {
      const rest = [
        ...state.lua.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        lua: {
          ...state.lua, settings: {...state.lua.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewLuaParameter: {
      const rest = [
        ...state.lua.settings.new.slice(0, action.payload.index),
        null,
        ...state.lua.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        lua: {
          ...state.lua, settings: {...state.lua.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddLuaParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.lua.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.lua.settings.new.slice(0, action.payload.index),
          null,
          ...state.lua.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        lua: <IsimpleModule>{
          ...state.lua, settings: {...state.lua.settings, [data.id]: data, new: rest},
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

