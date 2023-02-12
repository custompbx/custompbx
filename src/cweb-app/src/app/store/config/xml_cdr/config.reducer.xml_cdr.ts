
import {Iitem, initialState, State, IsimpleModule} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.xml_cdr';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetXmlCdr:
    case ConfigActionTypes.UpdateXmlCdrParameter:
    case ConfigActionTypes.SwitchXmlCdrParameter:
    case ConfigActionTypes.AddXmlCdrParameter:
    case ConfigActionTypes.DelXmlCdrParameter: {
      return {...state,
        xml_cdr: {
          ...state.xml_cdr,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotXmlCdrError: {
      return {
        ...state,
        xml_cdr: {
          ...state.xml_cdr,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetXmlCdr: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          xml_cdr: {...state.xml_cdr, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.xml_cdr) {
        state.xml_cdr = <IsimpleModule>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        xml_cdr: {
          ...state.xml_cdr,
          settings: {...state.xml_cdr.settings, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelXmlCdrParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.xml_cdr.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.xml_cdr.settings;

      return {
        ...state,
        xml_cdr: {
          ...state.xml_cdr, settings: {...rest, new: state.xml_cdr.settings.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchXmlCdrParameter:
    case ConfigActionTypes.StoreUpdateXmlCdrParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        xml_cdr: <IsimpleModule>{
          ...state.xml_cdr, settings: {...state.xml_cdr.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewXmlCdrParameter: {
      const rest = [
        ...state.xml_cdr.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        xml_cdr: {
          ...state.xml_cdr, settings: {...state.xml_cdr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewXmlCdrParameter: {
      const rest = [
        ...state.xml_cdr.settings.new.slice(0, action.payload.index),
        null,
        ...state.xml_cdr.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        xml_cdr: {
          ...state.xml_cdr, settings: {...state.xml_cdr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddXmlCdrParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.xml_cdr.settings.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.xml_cdr.settings.new.slice(0, action.payload.index),
          null,
          ...state.xml_cdr.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        xml_cdr: <IsimpleModule>{
          ...state.xml_cdr, settings: {...state.xml_cdr.settings, [data.id]: data, new: rest},
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

