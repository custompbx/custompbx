import {All, CDRActionTypes} from './cdr.actions';
import {All as SettingsAll, SettingsActionTypes} from '../settings/settings.actions';

export interface State {
  cdrData: Array<object>;
  settings: object;
  loadCounter: number;
  errorMessage: string | null;
}

export const initialState: State = {
  cdrData: [],
  settings: {},
  loadCounter: 0,
  errorMessage: '',
};

export function reducer(state = initialState, action: All | SettingsAll): State {
  switch (action.type) {
    case CDRActionTypes.UPDATE_FAILURE: {
      return {
        ...state,
        loadCounter: state.loadCounter >= 0 ? --state.loadCounter : 0,
      };
    }

    case SettingsActionTypes.SaveWebSettings:
    case SettingsActionTypes.GetWebSettings:
    case CDRActionTypes.GET_CDR: {
      return {...state, loadCounter: state.loadCounter + 1, errorMessage: null};
    }

    case SettingsActionTypes.StoreGotWebError:
    case CDRActionTypes.StoreGotCdrError: {
        return {
          ...state,
          errorMessage: action.payload.error,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        };
    }

    case SettingsActionTypes.StoreSaveWebSettings:
    case SettingsActionTypes.StoreGetWebSettings: {
      return {
        ...state,
        settings: action.payload.response['web_settings'] || {},
        loadCounter: --state.loadCounter,
        errorMessage: null};
    }


    case CDRActionTypes.STORE_GET_CDR: {
      return {
        ...state,
        cdrData: action.payload.response.cdr || [],
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
    return state;
  }
  }
}
