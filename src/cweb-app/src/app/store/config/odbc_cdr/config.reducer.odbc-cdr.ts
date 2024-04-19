import { ConfigActionTypes, All} from './config.actions.odbc-cdr';
import {Iitem, initialState, IodbcCdr, State} from '../config.state.struct';
import {getParentId} from "../config.reducers";

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetOdbcCdrField:
    case ConfigActionTypes.GetOdbcCdr:
    case ConfigActionTypes.UpdateOdbcCdrParameter:
    case ConfigActionTypes.SwitchOdbcCdrParameter:
    case ConfigActionTypes.DeleteOdbcCdrParameter:
    case ConfigActionTypes.AddOdbcCdrParameter:
    case ConfigActionTypes.AddOdbcCdrTable:
    case ConfigActionTypes.UpdateOdbcCdrTable:
    case ConfigActionTypes.SwitchOdbcCdrTable:
    case ConfigActionTypes.DeleteOdbcCdrTable:
    case ConfigActionTypes.AddOdbcCdrField:
    case ConfigActionTypes.UpdateOdbcCdrField:
    case ConfigActionTypes.SwitchOdbcCdrField:
    case ConfigActionTypes.DeleteOdbcCdrField: {
      return {...state,
        odbc_cdr: {
          ...state.odbc_cdr,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotOdbcCdrError: {
      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetOdbcCdr: {
      const parameters = action.payload.response.data['settings'];
      const tables = action.payload.response.data['tables'];
      if (action.payload.response.exists === false) {
        return {
          ...state,
          odbc_cdr: {...state.odbc_cdr, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }
      const odbcCdr = state.odbc_cdr ? state.odbc_cdr : <IodbcCdr>{};

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr,
          settings: {...odbcCdr.settings, ...parameters},
          tables: {...odbcCdr.tables, ...tables},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOdbcCdrTable: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0 };
      }
      const odbcCdr = state.odbc_cdr ? state.odbc_cdr : <IodbcCdr>{};

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr,
          tables: {...odbcCdr.tables, [data.id]: data},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDeleteOdbcCdrParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.odbc_cdr.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.odbc_cdr.settings;

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr, settings: {...rest, new: state.odbc_cdr.settings?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateOdbcCdrParameter:
    case ConfigActionTypes.StoreSwitchOdbcCdrParameter: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0 };
      }

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, settings: {...state.odbc_cdr.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateOdbcCdrTable: {
      const data = action.payload.response.data || {};
      if (!data.id || !state.odbc_cdr.tables[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, tables: {...state.odbc_cdr.tables, [data.id]: {...state.odbc_cdr.tables[data.id], ...data}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOdbcCdrParameter: {
      const rest = [
        ...state.odbc_cdr.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr, settings: {...state.odbc_cdr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOdbcCdrParameter: {
      const rest = [
        ...state.odbc_cdr.settings.new.slice(0, action.payload.index),
        null,
        ...state.odbc_cdr.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr, settings: {...state.odbc_cdr.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOdbcCdrParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.odbc_cdr.settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.odbc_cdr.settings.new.slice(0, action.payload.index),
          null,
          ...state.odbc_cdr.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, settings: {...state.odbc_cdr.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddOdbcCdrField: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const table = state.odbc_cdr.tables[parentId];
      if (!table) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...table.fields?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...table.fields.new.slice(0, action.payload.index),
          null,
          ...table.fields.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, tables: {
            ...state.odbc_cdr.tables, [parentId]:
              {...table, fields: {...table.fields, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
/////// paste

    case ConfigActionTypes.StoreDeleteOdbcCdrField: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const table = state.odbc_cdr.tables[parentId];
      if (!table || !table.fields[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = table.fields;

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, tables: {
            ...state.odbc_cdr.tables, [parentId]:
              {...table, fields: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetOdbcCdrField: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const table = state.odbc_cdr.tables[parentId];
      if (!table) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, tables: {
            ...state.odbc_cdr.tables, [parentId]:
              {...table, fields: {...table.fields, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateOdbcCdrField:
    case ConfigActionTypes.StoreSwitchOdbcCdrField: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const table = state.odbc_cdr.tables[parentId];
      if (!table) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        odbc_cdr: <IodbcCdr>{
          ...state.odbc_cdr, tables: {
            ...state.odbc_cdr.tables, [parentId]:
              {...table, fields: {...table.fields, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewOdbcCdrField: {
      const table = state.odbc_cdr.tables[action.payload.id];
      console.log(action.payload);
      if (!table) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...table.fields?.new || [],
        <Iitem>{}
      ];

      console.log(rest);
      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr, tables: {
            ...state.odbc_cdr.tables, [action.payload.id]:
              {...table, fields: {...table.fields, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewOdbcCdrField: {
      const table = state.odbc_cdr.tables[action.payload.id];
      if (!table) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...table.fields.new.slice(0, action.payload.index),
        null,
        ...table.fields.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr, tables: {
            ...state.odbc_cdr.tables, [action.payload.id]:
              {...table, fields: {...table.fields, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDeleteOdbcCdrTable: {
      const id = action.payload.response.data?.id || 0;
      if (!state.odbc_cdr.tables[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.odbc_cdr.tables;

      return {
        ...state,
        odbc_cdr: {
          ...state.odbc_cdr, tables: {...rest},
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
