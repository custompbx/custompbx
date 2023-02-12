import {All, HEPActionTypes} from './hep.actions';

export interface State {
  hepData: Array<object>;
  hepDetails: Array<object>;
  settings: object;
  loadCounter: number;
  errorMessage: string | null;
}

export const initialState: State = {
  hepData: [],
  hepDetails: [],
  settings: {},
  loadCounter: 0,
  errorMessage: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case HEPActionTypes.GetHEPDetails:
    case HEPActionTypes.GetHEP: {
      return {...state, loadCounter: state.loadCounter + 1, errorMessage: null};
    }

    case HEPActionTypes.StoreGotHEPError: {
        return {
          ...state,
          errorMessage: action.payload.error,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        };
    }

    case HEPActionTypes.StoreGetHEP: {
      return {
        ...state,
        hepData: action.payload.response.heps || [],
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case HEPActionTypes.StoreGetHEPDetails: {
      return {
        ...state,
        hepDetails: action.payload.response['hep_details'] || [],
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
    return state;
  }
  }
}
