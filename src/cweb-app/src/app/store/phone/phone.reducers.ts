import {All, AuthActionTypes} from './phone.actions';

export interface Icreds {
  user_name: string;
  password: string;
  domain: string;
  ws: string;
  stun: string;
}

export interface State {
  phoneCreds: Icreds;
  phoneStatus: Istatuses;
  callTo: string;
  errorMessage: string | null;
  lastActionName: string;
}

export interface Istatuses {
  isRunning: boolean;
}

export const initialState: State = {
  phoneCreds: null,
  phoneStatus: <Istatuses>{},
  callTo: '',
  errorMessage: null,
  lastActionName: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case AuthActionTypes.GET_PHONE_CREDS: {
      return {
        ...state,
        lastActionName: AuthActionTypes.GET_PHONE_CREDS,
      };
    }

    case AuthActionTypes.STORE_GET_PHONE_CREDS: {
      return {
        ...state,
        phoneCreds: action.payload.response['phone_creds'],
        errorMessage: action.payload.response.error ||  '',
        callTo: '',
        lastActionName: AuthActionTypes.STORE_GET_PHONE_CREDS,
      };
    }

    case AuthActionTypes.STORE_PHONE_STATUS: {
      return {
        ...state,
        phoneStatus: {...state.phoneStatus, ...action.payload.phoneStatus},
        callTo: '',
        lastActionName: AuthActionTypes.STORE_PHONE_STATUS,
      };
    }

    case AuthActionTypes.STORE_MAKE_PHONE_CALL: {
      return {
        ...state,
        callTo: action.payload.user,
        lastActionName: AuthActionTypes.STORE_MAKE_PHONE_CALL,
      };
    }

    default: {
      return {
        ...state,
        callTo: '',
        lastActionName: '',
      };
    }
  }
}
