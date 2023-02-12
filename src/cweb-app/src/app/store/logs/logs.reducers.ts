import {All, LogsActionTypes } from './logs.actions';

export interface State {
  logsData: Array<object>;
  loadCounter: number;
  errorMessage: string | null;
}

export const initialState: State = {
  logsData: [],
  loadCounter: 0,
  errorMessage: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case LogsActionTypes.GetLogs: {
      return {...state, loadCounter: state.loadCounter + 1, errorMessage: null};
    }

    case LogsActionTypes.StoreGotLogsError: {
        return {
          ...state,
          errorMessage: action.payload.error,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        };
    }

    case LogsActionTypes.StoreGetLogs: {
      return {
        ...state,
        logsData: action.payload.response.logs || [],
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
    return state;
  }
  }
}
