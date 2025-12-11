import {
  SettingsActionTypes,
  All, AddWebDirectoryUsersTemplate, StoreNewWebDirectoryUsersTemplateParameter, StoreDelNewWebDirectoryUsersTemplateParameter,
} from './settings.actions';

export interface State {
  settingsData: object;
  webUsersTemplates: { [id: number]: object };
  webUsersTemplateParameters: { [id: number]: object };
  newWUTPs: { [id: number]: Array<object> };
  newWUTVs: { [id: number]: Array<object> };
  webUsersTemplateVariables: { [id: number]: object };
  webUsers: { [id: number]: IwebUser };
  webGroups: { [id: number]: object };
  wssUris: Array<string>;
  vertoWsUris: Array<string>;
  loadCounter: number;
  errorMessage: string | null;
}

export const initialState: State = {
  settingsData: {},
  webUsersTemplates: {},
  webUsers: {},
  webUsersTemplateParameters: {},
  webUsersTemplateVariables: {},
  newWUTPs: {},
  newWUTVs: {},
  webGroups: {},
  wssUris: [],
  vertoWsUris: [],
  loadCounter: 0,
  errorMessage: null,
};

export interface Isettings {
  freeswitch: {esl: object};
  database: object;
  webserver: object;
  xml_curl_server: object;
  cert_path: object;
  key_path: object;
}

export interface IwebUser {
  id: number;
  login: string;
  verto_ws: string;
  stun: string;
  avatar_format: string;
  tokens: Array<{id: number}>;
  webrtc_lib: string;
  ws: string;
  group_id: number;
  sip_id: any;
}

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case SettingsActionTypes.UpdateWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.SwitchWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.DelWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.AddWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.GetWebDirectoryUsersTemplateVariables:
    case SettingsActionTypes.UpdateWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.SwitchWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.DelWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.AddWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.GetWebDirectoryUsersTemplateParameters:
    case SettingsActionTypes.UpdateWebDirectoryUsersTemplate:
    case SettingsActionTypes.SwitchWebDirectoryUsersTemplate:
    case SettingsActionTypes.DelWebDirectoryUsersTemplate:
    case SettingsActionTypes.AddWebDirectoryUsersTemplate:
    case SettingsActionTypes.GetWebDirectoryUsersTemplates:
    case SettingsActionTypes.UpdateWebUserGroup:
    case SettingsActionTypes.AddUserToken:
    case SettingsActionTypes.RemoveUserToken:
    case SettingsActionTypes.GetUserTokens:
    case SettingsActionTypes.CLEAR_WEB_USER_AVATAR:
    case SettingsActionTypes.UPDATE_WEB_USER_AVATAR:
    case SettingsActionTypes.ADD_WEB_USER:
    case SettingsActionTypes.DELETE_WEB_USER:
    case SettingsActionTypes.RENAME_WEB_USER:
    case SettingsActionTypes.GET_WEB_USERS:
    case SettingsActionTypes.SWITCH_WEB_USER:
    case SettingsActionTypes.UPDATE_WEB_USER_PASSWORD:
    case SettingsActionTypes.UPDATE_WEB_USER_LANG:
    case SettingsActionTypes.UPDATE_WEB_USER_SIP_USER:
    case SettingsActionTypes.UPDATE_WEB_USER_WS:
    case SettingsActionTypes.UPDATE_WEB_USER_VERTO_WS:
    case SettingsActionTypes.UPDATE_WEB_USER_WEBRTC_LIB:
    case SettingsActionTypes.SET_SETTINGS:
    case SettingsActionTypes.GET_SETTINGS: {
      return {...state,
        errorMessage: null, loadCounter: state.loadCounter + 1};
    }

    case SettingsActionTypes.StoreGotWebError: {
      return {
        ...state,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.error,
      };
    }

    case SettingsActionTypes.STORE_SETTINGS: {
      return {
        ...state,
        settingsData: action.payload.response.settings,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.UPDATE_FAILURE: {
      return {
        ...state,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: 'Cant get data from server',
      };
    }

    case SettingsActionTypes.STORE_GET_WEB_USERS: {
      return {
        ...state,
        webUsers: action.payload.response['web_users'],
        webGroups: action.payload.response['web_users_groups'],
        wssUris: action.payload.response.options || [],
        vertoWsUris: action.payload.response['alt_options'] || [],
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreUpdateWebUserGroup:
    case SettingsActionTypes.STORE_CLEAR_WEB_USER_AVATAR:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_AVATAR:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_PASSWORD:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_LANG:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_SIP_USER:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_WS:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_VERTO_WS:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_WEBRTC_LIB:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_STUN:
    case SettingsActionTypes.STORE_SWITCH_WEB_USER:
    case SettingsActionTypes.STORE_RENAME_WEB_USER:
    case SettingsActionTypes.STORE_ADD_WEB_USER: {
      const data = action.payload.response['web_users'];
      const groups = action.payload.response['web_users_groups'];
      if (!data) {
        return {
          ...state,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const id = ids[0];
      let newState = {...state};

      switch (action.type) {
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_PASSWORD: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], password: ''}}};
          break;
        }
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_LANG: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], lang: data[id].lang}}};
          break;
        }
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_SIP_USER: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], sip_user: data[id].sip_user}}};
          break;
        }
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_WS: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], ws: data[id].ws}}};
          break;
        }
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_VERTO_WS: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], verto_ws: data[id].verto_ws}}};
          break;
        }
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_WEBRTC_LIB: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], webrtc_lib: data[id].webrtc_lib}}};
          break;
        }
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_STUN: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], stun: data[id].stun}}};
          break;
        }
        case SettingsActionTypes.STORE_CLEAR_WEB_USER_AVATAR:
        case SettingsActionTypes.STORE_UPDATE_WEB_USER_AVATAR: {
          let format = data[id].avatar_format;
          if (format.length) {
            format = format + '?' + (+new Date());
          }
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], avatar_format: format}}};
          break;
        }
        case SettingsActionTypes.STORE_SWITCH_WEB_USER: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], enabled: data[id].enabled}}};
          break;
        }
        case SettingsActionTypes.STORE_RENAME_WEB_USER: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], login: data[id].login}}};
          break;
        }
        case SettingsActionTypes.STORE_ADD_WEB_USER:  {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], ...data[id]}}};
          break;
        }
        case SettingsActionTypes.StoreUpdateWebUserGroup: {
          newState = {...newState, webUsers: {...newState.webUsers, [id]: {...newState.webUsers[id], group_id: data[id].group_id}}};
          break;
        }
      }

      return {
        ...state,
        webUsers: {...newState.webUsers},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.STORE_DELETE_WEB_USER: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.webUsers;

      return {
        ...state,
        webUsers: {...rest},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreGetUserTokens: {
      const id = action.payload.response.id;
      const data = action.payload.response['tokens_list'] || [];
      const ids = Object.keys(data);
      if (ids.length === 0 || !state.webUsers[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        webUsers: {...state.webUsers, [id]: {...state.webUsers[id], tokens: data}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreAddUserToken: {
      const id = action.payload.response.id;
      const data = action.payload.response['tokens_list'] || [];
      const ids = Object.keys(data);
      if (ids.length === 0 || !state.webUsers[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const tokens = state.webUsers[id]?.tokens || [];
      return {
        ...state,
        webUsers: {...state.webUsers, [id]: {...state.webUsers[id], tokens: [ ...tokens, ...data]}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreRemoveUserToken: {
      const id = action.payload.response.id;
      const afId = action.payload.response.affected_id;
      if (!id || !state.webUsers[id] || !afId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const tokens = state.webUsers[id]?.tokens || [];
      const rest = tokens.filter(i => i.id !== afId);

      return {
        ...state,
        webUsers: {...state.webUsers, [id]: {...state.webUsers[id], tokens: [...rest]}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreGetWebDirectoryUsersTemplates: {
      const data = action.payload.response['data'] || {};
      return {
        ...state,
        webUsersTemplates: data,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreUpdateWebDirectoryUsersTemplate:
    case SettingsActionTypes.StoreSwitchWebDirectoryUsersTemplate:
    case SettingsActionTypes.StoreAddWebDirectoryUsersTemplate: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        webUsersTemplates: {...state.webUsersTemplates, ...data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreDelWebDirectoryUsersTemplate: {
      const id = action.payload.response['id'] || 0;
      const {[id]: toDel, ...rest} = state.webUsersTemplates;

      return {
        ...state,
        webUsersTemplates: {...rest},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreUpdateWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.StoreSwitchWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.StoreAddWebDirectoryUsersTemplateParameter:
    case SettingsActionTypes.StoreGetWebDirectoryUsersTemplateParameters: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        if (action.payload.index || action.payload.index === 0) {
          const params = state.newWUTPs[data.parent.id] || [];
          params.splice(action.payload.index, 1);
        }
        data = {[data.id]: data};
      }
      return {
        ...state,
        webUsersTemplateParameters: {...state.webUsersTemplateParameters, ...data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreNewWebDirectoryUsersTemplateParameter: {
      const id = action.payload.id;
      const params = state.newWUTPs[id] || [];
      return {
        ...state,
        newWUTPs: {...state.newWUTPs, [id]: [...params, {}]},
      };
    }

    case SettingsActionTypes.StoreDelNewWebDirectoryUsersTemplateParameter: {
      const id = action.payload.id;
      const index = action.payload.index;
      const params = state.newWUTPs[id] || [];
      params.splice(index, 1);
      return {
        ...state,
        newWUTPs: {...state.newWUTPs, [id]: [...params]},
      };
    }

    case SettingsActionTypes.StoreDelWebDirectoryUsersTemplateParameter: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.webUsersTemplateParameters;
      return {
        ...state,
        webUsersTemplateParameters: {...rest},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreUpdateWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.StoreSwitchWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.StoreAddWebDirectoryUsersTemplateVariable:
    case SettingsActionTypes.StoreGetWebDirectoryUsersTemplateVariables: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        if (action.payload.index || action.payload.index === 0) {
          const params = state.newWUTVs[data.parent.id] || [];
          params.splice(action.payload.index, 1);
        }
        data = {[data.id]: data};
      }
      return {
        ...state,
        webUsersTemplateVariables: {...state.webUsersTemplateVariables, ...data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case SettingsActionTypes.StoreNewWebDirectoryUsersTemplateVariable: {
      const id = action.payload.id;
      const params = state.newWUTVs[id] || [];
      return {
        ...state,
        newWUTVs: {...state.newWUTVs, [id]: [...params, {}]},
      };
    }

    case SettingsActionTypes.StoreDelNewWebDirectoryUsersTemplateVariable: {
      const id = action.payload.id;
      const index = action.payload.index;
      const params = state.newWUTVs[id] || [];
      params.splice(index, 1);
      return {
        ...state,
        newWUTVs: {...state.newWUTVs, [id]: [...params]},
      };
    }

    case SettingsActionTypes.StoreDelWebDirectoryUsersTemplateVariable: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.webUsersTemplateVariables;
      return {
        ...state,
        webUsersTemplateVariables: {...rest},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    default: {
      return state;
    }
  }
}
