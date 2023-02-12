
import {Iitem, initialState, State, IpostLoadModules} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.PostLoadModules';
import {StoreAutoloadModule, ConfigActionTypes as ConfigActionTypesModules} from '../config.actions';

export function reducer(state = initialState, action: (All | StoreAutoloadModule)): State {
  switch (action.type) {
    case ConfigActionTypes.GetPostLoadModules:
    case ConfigActionTypes.UpdatePostLoadModule:
    case ConfigActionTypes.SwitchPostLoadModule:
    case ConfigActionTypes.AddPostLoadModule:
    case ConfigActionTypes.DelPostLoadModule: {
      return {...state,
        post_load_modules: {
          ...state.post_load_modules,
        },
        errorMessage: null,
        loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotPostLoadModuleError: {
      return {
        ...state,
        post_load_modules: {
          ...state.post_load_modules,
        },
        errorMessage: action.payload.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetPostLoadModules: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          post_load_modules: {...state.post_load_modules, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.post_load_modules) {
        state.post_load_modules = <IpostLoadModules>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        post_load_modules: {
          ...state.post_load_modules,
          modules: {...state.post_load_modules.modules, ...settings},
          exists: action.payload.response.exists,
        },
        errorMessage: action.payload.response.error ||  '',
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelPostLoadModule: {
      const id = action.payload.response.data?.id || 0;
      if (!state.post_load_modules.modules[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.post_load_modules.modules;

      return {
        ...state,
        post_load_modules: {
          ...state.post_load_modules, modules: {...rest},
          newModules: state.post_load_modules.newModules || [],
        },
        errorMessage: action.payload.response.error ||  '',
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypesModules.STORE_AUTOLOAD_MODULE:
    case ConfigActionTypes.StoreSwitchPostLoadModule:
    case ConfigActionTypes.StoreUpdatePostLoadModule: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        post_load_modules: <IpostLoadModules>{
          ...state.post_load_modules, modules: {...state.post_load_modules.modules, [data.id]: data},
        },
        errorMessage: action.payload.response.error ||  '',
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewPostLoadModule: {
      const rest = [
        ...state.post_load_modules.newModules || [],
        <Iitem>{}
      ];

      return {
        ...state,
        post_load_modules: {
          ...state.post_load_modules, modules: {...state.post_load_modules.modules},
          newModules: rest,
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewPostLoadModule: {
      const rest = [
        ...state.post_load_modules.newModules.slice(0, action.payload.index),
        null,
        ...state.post_load_modules.newModules.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        post_load_modules: {
          ...state.post_load_modules, modules: {...state.post_load_modules.modules},
          newModules: rest,
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddPostLoadModule: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.post_load_modules.newModules || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.post_load_modules.newModules.slice(0, action.payload.index),
          null,
          ...state.post_load_modules.newModules.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        post_load_modules: <IpostLoadModules>{
          ...state.post_load_modules, modules: {...state.post_load_modules.modules, [data.id]: data},
          newModules: rest,
        },
        errorMessage: action.payload.response.error ||  '',
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return null;
    }
  }
}

