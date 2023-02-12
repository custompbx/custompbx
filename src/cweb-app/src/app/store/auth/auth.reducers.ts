import {All, AuthActionTypes} from './auth.actions';
import {All as SettingsAll, SettingsActionTypes} from '../settings/settings.actions';

export interface State {
  isAuthenticated: boolean;
  user: Iuser | null;
  token: string;
  errorMessage: string | null;
}
export interface Iuser {
  id?: string;
  login?: string;
  password?: string;
  group_id?: number;
  token?: string;
  avatar_format?: string;
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

    case SettingsActionTypes.STORE_CLEAR_WEB_USER_AVATAR:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_AVATAR: {
      const data = action.payload.response['web_users'] || {};
      if (!data) {
        return {
          ...state,
          errorMessage: action.payload.response.error,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state};
      }
      const id = ids[0];
      let format = data[id].avatar_format;
      if (format.length) {
        format = format + '?' + (+new Date());
      }
      return {
        ...state,
        user: {...state.user, avatar_format: format},
      };
    }

    case SettingsActionTypes.STORE_RENAME_WEB_USER: {
      const data = action.payload.response['web_users'] || {};
      if (!data) {
        return {
          ...state,
          errorMessage: action.payload.response.error,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state};
      }
      const id = ids[0];
      return {
        ...state,
        user: {...state.user, login: data[id].login},
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
