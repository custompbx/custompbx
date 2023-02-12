import {
  ConfigActionTypes,
  All,
} from './config.actions.cdr-pg-csv';
import {
  IcdrPgCsv, Ifield,
  Iitem,
  initialState,
  State
} from '../config.state.struct';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.UPDATE_CDR_PG_CSV_PARAMETER:
    case ConfigActionTypes.UPDATE_CDR_PG_CSV_FIELD:
    case ConfigActionTypes.SWITCH_CDR_PG_CSV_PARAMETER:
    case ConfigActionTypes.SWITCH_CDR_PG_CSV_FIELD:
    case ConfigActionTypes.DELETE_CDR_PG_CSV_PARAMETER:
    case ConfigActionTypes.DELETE_CDR_PG_CSV_FIELD:
    case ConfigActionTypes.ADD_CDR_PG_CSV_PARAMETER:
    case ConfigActionTypes.ADD_CDR_PG_CSV_FIELD:
    case ConfigActionTypes.GET_CDR_PG_CSV: {
      return {...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          errorMessage: null,
        },  loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotCdrPgCsvError: {
      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_GET_CDR_PG_CSV: {
      if (action.payload.response.exists === false) {
        return {
          ...state,
          cdr_pg_csv: {...state.cdr_pg_csv, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }
      const parameters = action.payload.response.data['settings'];
      const cdrSchema = action.payload.response.data['schemas'];
      const cdrPgCsv = state.cdr_pg_csv ? state.cdr_pg_csv : <IcdrPgCsv>{};

      let restParams = cdrPgCsv.newSettingParams || [];
      let restFields = cdrPgCsv.newSchemaFields || [];
      if (action.payload.index !== undefined) {
        if (parameters) {
          restParams = [
            ...restParams.slice(0, action.payload.index),
            null,
            ...restParams.slice(action.payload.index + 1)
          ];
        }
        if (cdrSchema) {
          restFields = [
            ...restFields.slice(0, action.payload.index),
            null,
            ...restFields.slice(action.payload.index + 1)
          ];
        }
      }

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          settings: {...cdrPgCsv.settings, ...parameters},
          schema: {...cdrPgCsv.schema, ...cdrSchema},
          newSettingParams: [...restParams],
          newSchemaFields: [...restFields],
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_CDR_PG_CSV_PARAMETER:
    case ConfigActionTypes.STORE_UPDATE_CDR_PG_CSV_PARAMETER:
    case ConfigActionTypes.STORE_ADD_CDR_PG_CSV_PARAMETER: {
      if (action.payload.response.exists === false) {
        return {
          ...state,
          cdr_pg_csv: {...state.cdr_pg_csv, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }
      const data = action.payload.response.data || {};
      const cdrPgCsv = state.cdr_pg_csv ? state.cdr_pg_csv : <IcdrPgCsv>{};

      let restParams = cdrPgCsv.newSettingParams || [];
      if (action.payload.index !== undefined) {
        restParams = [
          ...restParams.slice(0, action.payload.index),
          null,
          ...restParams.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          settings: {...cdrPgCsv.settings, [data.id]: data},
          newSettingParams: [...restParams],
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_CDR_PG_CSV_FIELD:
    case ConfigActionTypes.STORE_UPDATE_CDR_PG_CSV_FIELD:
    case ConfigActionTypes.STORE_ADD_CDR_PG_CSV_FIELD: {
      if (action.payload.response.exists === false) {
        return {
          ...state,
          cdr_pg_csv: {...state.cdr_pg_csv, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }
      const data = action.payload.response.data || {};
      const cdrPgCsv = state.cdr_pg_csv ? state.cdr_pg_csv : <IcdrPgCsv>{};

      let restFields = cdrPgCsv.newSchemaFields || [];
      if (action.payload.index !== undefined) {
        restFields = [
          ...restFields.slice(0, action.payload.index),
          null,
          ...restFields.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          schema: {...cdrPgCsv.schema, [data.id]: data},
          newSchemaFields: [...restFields],
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_CDR_PG_CSV_PARAMETER: {
      const rest = [
        ...state.cdr_pg_csv.newSettingParams || [],
        <Iitem>{}
      ];
      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          newSettingParams: [...rest],
          errorMessage: null
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_CDR_PG_CSV_FIELD: {
      const rest = [
        ...state.cdr_pg_csv.newSchemaFields || [],
        <Ifield>{}
      ];
      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          newSchemaFields: [...rest],
          errorMessage: null
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_CDR_PG_CSV_PARAMETER: {
      const rest = [
        ...state.cdr_pg_csv.newSettingParams.slice(0, action.payload.index),
        null,
        ...state.cdr_pg_csv.newSettingParams.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          newSettingParams: [...rest],
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_CDR_PG_CSV_FIELD: {
      const rest = [
        ...state.cdr_pg_csv.newSchemaFields.slice(0, action.payload.index),
        null,
        ...state.cdr_pg_csv.newSchemaFields.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv,
          newSchemaFields: [...rest],
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DELETE_CDR_PG_CSV_PARAMETER: {
      const id = action.payload.response.data?.id || 0;
      if (!state.cdr_pg_csv.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.cdr_pg_csv.settings;

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv, settings: {...rest},
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DELETE_CDR_PG_CSV_FIELD: {
      const id = action.payload.response.data?.id || 0;
      if (!state.cdr_pg_csv.schema[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.cdr_pg_csv.schema;

      return {
        ...state,
        cdr_pg_csv: {
          ...state.cdr_pg_csv, schema: {...rest},
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
