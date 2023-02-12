
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.xml_rpc';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetXmlRpc:
    case ConfigActionTypes.UpdateXmlRpcParameter:
    case ConfigActionTypes.SwitchXmlRpcParameter:
    case ConfigActionTypes.AddXmlRpcParameter:
    case ConfigActionTypes.DelXmlRpcParameter: {
      return {...state,
        xml_rpc: {
          ...state.xml_rpc,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotXmlRpcError: {
      return {
        ...state,
        xml_rpc: {
          ...state.xml_rpc,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetXmlRpc: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          xml_rpc: {...state.xml_rpc, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.xml_rpc) {
        state.xml_rpc = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        xml_rpc: {
          ...state.xml_rpc,
          settings: {...state.xml_rpc.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelXmlRpcParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.xml_rpc.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.xml_rpc.settings;

      return {
        ...state,
        xml_rpc: {
          ...state.xml_rpc, settings: {...rest, new: state.xml_rpc.settings.new || []},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchXmlRpcParameter:
    case ConfigActionTypes.StoreUpdateXmlRpcParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        xml_rpc: <IsimpleModule>{
          ...state.xml_rpc, settings: {...state.xml_rpc.settings, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewXmlRpcParameter: {
      const rest = [
        ...state.xml_rpc.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        xml_rpc: {
          ...state.xml_rpc, settings: {...state.xml_rpc.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewXmlRpcParameter: {
      const rest = [
        ...state.xml_rpc.settings.new.slice(0, action.payload.index),
        null,
        ...state.xml_rpc.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        xml_rpc: {
          ...state.xml_rpc, settings: {...state.xml_rpc.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddXmlRpcParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.xml_rpc.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.xml_rpc.settings.new.slice(0, action.payload.index),
          null,
          ...state.xml_rpc.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        xml_rpc: <IsimpleModule>{
          ...state.xml_rpc, settings: {...state.xml_rpc.settings, [data.id]: data, new: rest},
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

