import { DaemonActionTypes, All } from './daemon.actions';

export interface State {
  eslConnection: boolean;
  dbConnection: boolean;
  errorMessage: string | null;
  tokenFailed: boolean;
}

export const initialState: State = {
  eslConnection: true,
  dbConnection: true,
  errorMessage: null,
  tokenFailed: false,
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case DaemonActionTypes.CONNECTION: {
      return {
        ...state,
        dbConnection: action.payload.daemon['database_connection'],
        eslConnection: action.payload.daemon['esl_connection'],
        errorMessage: action.payload.daemon['esl_error'] + '. ' + action.payload.daemon['data_base_error'],
        tokenFailed: action.payload['no_token'] || false,
      };
    }
    case DaemonActionTypes.STORE_FLUSH_TOKEN_STATE: {
      return {
        ...state,
        tokenFailed: false,
      };
    }
    default: {
      return state;
    }
  }
}
