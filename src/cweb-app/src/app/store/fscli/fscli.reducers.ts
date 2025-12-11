import {All, FSCLIActionTypes} from './fscli.actions';

export interface State {
  fsCliData: string;
  loadCounter: number;
  errorMessage: string | null;
}


export const initialState: State = {
  fsCliData: '',
  loadCounter: 0,
  errorMessage: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case FSCLIActionTypes.UPDATE_FAILURE: {
      return {
        ...state,
        loadCounter: state.loadCounter >= 0 ? --state.loadCounter : 0,
      };
    }
    case FSCLIActionTypes.SendFSCLICommand: {
      return {...state, loadCounter: state.loadCounter + 1, errorMessage: null};
    }

    case FSCLIActionTypes.StoreSendFSCLICommand: {
      if (action.payload.response.error) {
        return {
          ...state,
          errorMessage: action.payload.response.error,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        };
      }
      return {
        ...state,
        fsCliData: action.payload.response.response || '',
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
    return state;
  }
  }
}
