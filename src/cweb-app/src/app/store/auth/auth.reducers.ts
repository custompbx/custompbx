import {All, AuthActionTypes} from './auth.actions';
import {All as SettingsAll, SettingsActionTypes} from '../settings/settings.actions';

export interface State {
  isAuthenticated: boolean;
  user: Iuser | null;
  token: string;
  errorMessage: string | null;
}
export interface Iuser {
  id?: any;
  login?: string;
  password?: string;
  group_id?: number;
  token?: string;
  avatar_format?: string;
  sip_id?: object;
}

export const initialState: State = {
  isAuthenticated: false,
  user: null,
  token: '',
  errorMessage: null,
};

export function reducer(state = initialState, action: (All | SettingsAll)): State {
  switch (action.type) {
    case AuthActionTypes.LOGIN_SUCCESS: {
      return {
        ...state,
        isAuthenticated: true,
        token: action.payload.response.token,
        user: action.payload.response.user,
        errorMessage: null
      };
    }

    case SettingsActionTypes.StoreUpdateWebUserGroup:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_STUN:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_WEBRTC_LIB:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_VERTO_WS:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_WS:
    case SettingsActionTypes.STORE_RENAME_WEB_USER:
    case SettingsActionTypes.STORE_CLEAR_WEB_USER_AVATAR:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_AVATAR: {
      const data = action.payload.response['web_users'] || {};
      if (!data) {
        return {
          ...state,
          errorMessage: action.payload.response.error,
        };
      }

      const updatedUser = data[state.user.id]
      if (!updatedUser) {
        return {...state}
      }

      if (updatedUser?.avatar_format?.length) {
        updatedUser.avatar_format += '?' + (+new Date());
      }
      return {
        ...state,
        user: {...state.user, ...updatedUser},
      };
    }

    case AuthActionTypes.LOGIN_FAILURE: {
      return {
        ...state,
        errorMessage: 'Incorrect login and/or password.'
      };
    }

    case AuthActionTypes.LOGOUT: {
      return initialState;
    }

    default: {
      return state;
    }
  }
}
