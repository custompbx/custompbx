import {
  DirectoryActionTypes,
  All, StoreDirectoryError,
} from './directory.actions';
import {All as SettingsAll, SettingsActionTypes} from '../settings/settings.actions';

export interface State {
  domains: Idomains;
  users: Iusers;
  groupNames: {[index: number]: object};
  userGateways: {[index: number]: object};
  domainDetails: Idetails;
  gatewayDetails: IGatewaydetails;
  groupUsers: {[index: number]: object};
  errorMessage: string | null;
  loadCounter: number;
  additionalData: object;
  webUsersTemplates: { [id: number]: object };
  templatesItems: {
    [id: number]:
      {
        id: number;
        name: string;
        parameters: Array<{ id: number; name: string; value: string; description: string; disabled: boolean; editable: boolean; }>;
        variables: Array<{ id: number; name: string; value: string; description: string; disabled: boolean; editable: boolean; }>;
      };
  };
}

export const initialState: State = {
  domains: <Idomains>{},
  users: <Iusers>{},
  groupNames: {},
  userGateways: {},
  domainDetails: <Idetails>{},
  groupUsers: {},
  gatewayDetails: {},
  errorMessage: null,
  loadCounter: 0,
  additionalData: {},
  webUsersTemplates: {},
  templatesItems: {},
};

export interface Idomains {
  [index: number]: {
    id: number,
    name: string,
    enabled: boolean,
  };
}

export interface Iusers {
  [index: number]: {
    id: number,
    name: string,
    enabled: boolean,
    cache: number,
    number_alias: number,
    cidr: string,
    parameters: Iparameters,
    variables: Ivariables,
  };
}

export interface Idetails {
  [index: number]: {
    parameters: Iparameters,
    variables: Ivariables,
  };
}

export interface IGatewaydetails {
  [index: number]: {
    parameters: Iparameters,
    variables: IgatewayVariables,
  };
}

export interface Iparameters {
  [index: number]: Iitem;
  new: Array<object>;
}

export interface Ivariables {
  [index: number]: Iitem;
  new: Array<object>;
}

export interface Iitem {
  id?: number;
  name: string;
  value: string;
  enabled?: boolean;
}

export interface IgatewayVariables {
  [index: number]: IdirectionItem;
  new: Array<object>;
}

export interface IdirectionItem {
  id?: number;
  name: string;
  value: string;
  direction: string;
  enabled?: boolean;
}

const isNumeric = (val: string): boolean => {
  return !isNaN(Number(val));
};

export function reducer(state = initialState, action: All | SettingsAll): State {
  switch (action.type) {
    case DirectoryActionTypes.ReduceLoadCounter: {
      return {
        ...state,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
      };
    }
    case DirectoryActionTypes.CreateWebDirectoryUsersByTemplate:
    case DirectoryActionTypes.GetWebDirectoryUsersTemplateForm:
    case DirectoryActionTypes.GetWebDirectoryUsersTemplatesList:
    case DirectoryActionTypes.ImportXMLDomainUser:
    case DirectoryActionTypes.ImportXMLDomain:
    case DirectoryActionTypes.GetWebUsersByDirectory:
    case DirectoryActionTypes.ImportDirectory:
    case DirectoryActionTypes.SwitchDirectoryUser:
    case DirectoryActionTypes.SwitchDirectoryUserParameter:
    case DirectoryActionTypes.SwitchDirectoryUserVariable:
    case DirectoryActionTypes.SwitchDirectoryDomainParameter:
    case DirectoryActionTypes.SwitchDirectoryDomainVariable:
    case DirectoryActionTypes.GetDirectoryDomains:
    case DirectoryActionTypes.AddDirectoryDomain:
    case DirectoryActionTypes.SwitchDirectoryDomain:
    case DirectoryActionTypes.GetDirectoryDomainDetails:
    case DirectoryActionTypes.RenameDirectoryDomain:
    // case DirectoryActionTypes.DOMAIN_NEW_DOMAIN:
    // case DirectoryActionTypes.NEW_DOMAIN_PARAM:
    // case DirectoryActionTypes.NEW_DOMAIN_VAR:
    case DirectoryActionTypes.DeleteDirectoryDomain:
    case DirectoryActionTypes.DeleteDirectoryDomainParameter:
    case DirectoryActionTypes.DeleteDirectoryDomainVariable:
    case DirectoryActionTypes.UpdateDirectoryDomainParameter:
    case DirectoryActionTypes.UpdateDirectoryDomainVariable:
    case DirectoryActionTypes.GetDirectoryUsers:
    case DirectoryActionTypes.GetDirectoryUserDetails:
    case DirectoryActionTypes.AddDirectoryUserParameter:
    case DirectoryActionTypes.AddDirectoryUserVariable:
    case DirectoryActionTypes.DeleteDirectoryUserParameter:
    case DirectoryActionTypes.DeleteDirectoryUserVariable:
    case DirectoryActionTypes.UpdateDirectoryUserParameter:
    case DirectoryActionTypes.UpdateDirectoryUserVariable:
    case DirectoryActionTypes.UpdateDirectoryUserCache:
    case DirectoryActionTypes.UpdateDirectoryUserCidr:
    case DirectoryActionTypes.UpdateDirectoryUserNumberAlias:
    case DirectoryActionTypes.AddDirectoryUser:
    case DirectoryActionTypes.DeleteDirectoryUser:
    case DirectoryActionTypes.UpdateDirectoryUserName:
    case DirectoryActionTypes.GetDirectoryGroups:
    case DirectoryActionTypes.GetDirectoryGroupUsers:
    // case DirectoryActionTypes.ADD_NEW_GROUP:
    case DirectoryActionTypes.DeleteDirectoryGroup:
    case DirectoryActionTypes.UpdateDirectoryGroupName:
    case DirectoryActionTypes.AddDirectoryGroupUser:
    case DirectoryActionTypes.DeleteDirectoryGroupUser:
    case DirectoryActionTypes.GetDirectoryUserGateways:
    case DirectoryActionTypes.GetDirectoryUserGatewayDetails:
    case DirectoryActionTypes.AddDirectoryUserGatewayParameter:
    case DirectoryActionTypes.DeleteDirectoryUserGatewayParameter:
    case DirectoryActionTypes.UpdateDirectoryUserGatewayParameter:
    case DirectoryActionTypes.SwitchDirectoryUserGatewayParameter:
    // case DirectoryActionTypes.ADD_NEW_USER_GATEWAY:
    case DirectoryActionTypes.DeleteDirectoryUserGateway:
    case DirectoryActionTypes.UpdateDirectoryUserGatewayName:
    case DirectoryActionTypes.UpdateDirectoryUserGatewayVariable:
    case DirectoryActionTypes.SwitchDirectoryUserGatewayVariable:
    case DirectoryActionTypes.AddDirectoryUserGatewayVariable:
    case DirectoryActionTypes.DeleteDirectoryUserGatewayVariable: {
      return {...state,
        errorMessage: null, loadCounter: state.loadCounter + 1};
    }

    case DirectoryActionTypes.StoreDirectoryError: {
      console.log(action);
      return {
        ...state,
        errorMessage: action.payload.error,
        loadCounter: 0,
      };
    }

    case DirectoryActionTypes.StoreImportXMLDomain:
    case DirectoryActionTypes.StoreGetDirectoryDomains: {
      const data = action.payload.response.data || {};
      return {
        ...state,
        domains: data,
        domainDetails: {...state.domainDetails},
        errorMessage: action.payload.response.error ||  '',
        loadCounter: 0,
      };
    }
    case DirectoryActionTypes.StoreAddDirectoryDomain:
    case DirectoryActionTypes.StoreSwitchDirectoryDomain:
    case DirectoryActionTypes.StoreRenameDirectoryDomain: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        domains: {...state.domains, [data.id]: data},
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreDeleteDirectoryDomain: {
      const data = action.payload.response.data || {};
      if (data.id === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.domains;
      return {
        ...state,
        domains: {...rest},
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }


    case DirectoryActionTypes.StoreGetDirectoryDomainDetails: {
      const variables = action.payload.response.data.variables || {};
      const parameters = action.payload.response.data.parameters || {};
      const parentId = getParentId(variables) || getParentId(parameters);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        domainDetails: {...state.domainDetails, [parentId]: {parameters: parameters, variables: variables}},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreUpdateDirectoryDomainParameter:
    case DirectoryActionTypes.StoreSwitchDirectoryDomainParameter:
    case DirectoryActionTypes.StoreAddDirectoryDomainParameter: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      let new_params = !!state.domainDetails[parentId]?.parameters ? state.domainDetails[parentId]?.parameters?.new || [] : [];
      if (isNumeric(action.payload.param_index)) {
        new_params = [
          ...new_params.slice(0, action.payload.param_index),
          null,
          ...new_params.slice(action.payload.param_index + 1)
        ];
      }

      return {
        ...state,
        domainDetails: <Idetails>{
          ...state.domainDetails, [parentId]: {
            parameters: {
              ...state.domainDetails[parentId]?.parameters, [data.id]: {...data}, new: [...new_params],
            },
            variables: state.domainDetails[parentId]?.variables,
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryDomainVariable:
    case DirectoryActionTypes.StoreSwitchDirectoryDomainVariable:
    case DirectoryActionTypes.StoreAddDirectoryDomainVariable: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let new_vars = !!state.domainDetails[parentId]?.variables ? state.domainDetails[parentId]?.variables?.new || [] : [];
      if (isNumeric(action.payload.var_index)) {
        new_vars = [
          ...new_vars.slice(0, action.payload.var_index),
          null,
          ...new_vars.slice(action.payload.var_index + 1)
        ];
      }

      return {
        ...state,
        domainDetails: <Idetails>{
          ...state.domainDetails, [parentId]: {
            variables: {
              ...state.domainDetails[parentId]?.variables, [data.id]: {...data}, new: [...new_vars],
            },
            parameters: state.domainDetails[parentId]?.parameters,
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.DropDirectoryDomainVariable: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      if (!state.domainDetails[parentId]?.variables) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.domainDetails[parentId]?.variables;

      return {
        ...state,
        domainDetails: {
          ...state.domainDetails,
          [parentId]: {...state.domainDetails[parentId], variables: <Ivariables>{...rest}}},
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.DropDirectoryDomainParameter: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      if (!state.domainDetails[parentId]?.parameters) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.domainDetails[parentId]?.parameters;

      return {
        ...state,
        domainDetails: {
          ...state.domainDetails,
          [parentId]: {...state.domainDetails[parentId], parameters: <Iparameters>{...rest}}
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreAddNewDirectoryDomainParameter: {
      let oldArray = [];
      if (state.domainDetails[action.payload.id]?.parameters && state.domainDetails[action.payload.id]?.parameters.new) {
        oldArray = state.domainDetails[action.payload.id]?.parameters?.new || [];
      }
      return {
        ...state,
        domainDetails: {
          ...state.domainDetails, [action.payload.id]: {
            ...state.domainDetails[action.payload.id],
            parameters: {
              ...state.domainDetails[action.payload.id]?.parameters || {},
              new: [...oldArray, {name: null, value: null}]
            }
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreAddNewDirectoryDomainVariable: {
      let oldArray = [];
      if (state.domainDetails[action.payload.id]?.variables && state.domainDetails[action.payload.id]?.variables.new) {
        oldArray = state.domainDetails[action.payload.id]?.variables?.new || [];
      }
      return {
        ...state,
        domainDetails: {
          ...state.domainDetails, [action.payload.id]: {
            ...state.domainDetails[action.payload.id],
            variables: {
              ...state.domainDetails[action.payload.id]?.variables || {},
              new: [...oldArray, {name: null, value: null}]
            }
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteNewDirectoryDomainParameter: {
      let oldArray = [];
      if (state.domainDetails[action.payload.id]?.parameters && state.domainDetails[action.payload.id]?.parameters.new) {
        oldArray = state.domainDetails[action.payload.id]?.parameters?.new || [];
      }
      const newArray = [
        ...oldArray.slice(0, action.payload.index),
        null,
        ...oldArray.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        domainDetails: {
          ...state.domainDetails, [action.payload.id]: {
            ...state.domainDetails[action.payload.id],
            parameters: {
              ...state.domainDetails[action.payload.id]?.parameters || {},
              new: [...newArray]
            }
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteNewDirectoryDomainVariable: {
      let oldArray = [];
      if (state.domainDetails[action.payload.id]?.variables && state.domainDetails[action.payload.id]?.variables.new) {
        oldArray = state.domainDetails[action.payload.id]?.variables?.new || [];
      }
      const newArray = [
        ...oldArray.slice(0, action.payload.index),
        null,
        ...oldArray.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        domainDetails: {
          ...state.domainDetails, [action.payload.id]: {
            ...state.domainDetails[action.payload.id],
            variables: {
              ...state.domainDetails[action.payload.id]?.variables || {},
              new: [...newArray]
            }
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.ClearDetails: {
      // state.domainDetails[action.payload] = {};
      return {
        ...state,
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.UpdateFailure: {
      return {
        ...state,
        errorMessage: 'Cant get data from server',
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreGetDirectoryUsers: {
      const domains = action.payload.response.data?.domains || null;
      let users = action.payload.response.data['directory_users'] || {};
      if (!users) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      if (users.id) {
        users = {[users.id]: users};
      }

      return {
        ...state,
        domains: domains || state.domains,
        users: {
          ...state.users,
          ...users,
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: 0,
      };
    }
    case DirectoryActionTypes.StoreGetDirectoryUserDetails: {
      const data = action.payload.response.data || {};
      if (!data?.user?.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        users: {
          ...state.users,
          [data?.user?.id]: {
            ...data.user, variables: data.variables, parameters: data.parameters,
          }
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StorePasteDirectoryDomainVariables: {
      const fromDomainId = action.payload.from_id;
      const toDomainId = action.payload.to_id;
      if (
        !fromDomainId || !toDomainId || !state.domains[fromDomainId] || !state.domains[toDomainId] || !state.domainDetails[fromDomainId]
      ) {
        return {
          ...state
        };
      }

      let newItems = state.domainDetails[toDomainId]?.variables ? state.domainDetails[toDomainId]?.variables.new || [] : [];

      const newArray = Object.keys(state.domainDetails[fromDomainId]?.variables).map(i => {
        if (i === 'new') {
          return;
        }
        return state.domainDetails[fromDomainId]?.variables[i];
      });

      newItems = [...newItems, ...newArray];
      const details = state.domainDetails[toDomainId] || {variables: {}};
      return {
        ...state,
        domainDetails: {...state.domainDetails,
          [toDomainId]: {
            ...state.domainDetails[toDomainId],
            variables: {...details.variables, new: [...newItems]},
          }
        },
      };
    }

    case DirectoryActionTypes.StorePasteDirectoryDomainParameters: {
      const fromDomainId = action.payload.from_id;
      const toDomainId = action.payload.to_id;
      if (
        !fromDomainId || !toDomainId || !state.domains[fromDomainId] || !state.domains[toDomainId] || !state.domainDetails[fromDomainId]
      ) {
        return {
          ...state
        };
      }

      let newItems = state.domainDetails[toDomainId]?.parameters ? state.domainDetails[toDomainId]?.parameters.new || [] : [];

      const newArray = Object.keys(state.domainDetails[fromDomainId]?.parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.domainDetails[fromDomainId]?.parameters[i];
      });

      newItems = [...newItems, ...newArray];
      const details = state.domainDetails[toDomainId] || {parameters: {}};
      return {
        ...state,
        domainDetails: {...state.domainDetails,
          [toDomainId]: {
            ...state.domainDetails[toDomainId],
            parameters: {...details.parameters, new: [...newItems]},
          }
        },
      };
    }

    case DirectoryActionTypes.StorePasteDirectoryUserVariables: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.users[fromId]
      ) {
        return {
          ...state
        };
      }

      let newItems = state.users[toId]?.variables ? state.users[toId]?.variables.new || [] : [];

      const newArray = Object.keys(state.users[fromId]?.variables).map(i => {
        if (i === 'new') {
          return;
        }
        return state.users[fromId]?.variables[i];
      });

      newItems = [...newItems, ...newArray];
      const details = state.users[toId] || {variables: {}};
      return {
        ...state,
        users: {...state.users,
          [toId]: {
            ...state.users[toId],
            variables: {...details.variables, new: [...newItems]},
          }
        },
      };
    }

    case DirectoryActionTypes.StorePasteDirectoryUserParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId  || !state.users[fromId]
      ) {
        return {
          ...state
        };
      }

      let newItems = state.users[toId]?.parameters ? state.users[toId]?.parameters.new || [] : [];

      const newArray = Object.keys(state.users[fromId]?.parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.users[fromId]?.parameters[i];
      });

      newItems = [...newItems, ...newArray];
      const details = state.users[toId] || {parameters: {}};
      return {
        ...state,
        users: {...state.users,
          [toId]: {
            ...state.users[toId],
            parameters: {...details.parameters, new: [...newItems]},
          }
        },
      };
    }

    case DirectoryActionTypes.StorePasteDirectoryUserGatewayVariables: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.gatewayDetails[fromId] || !state.gatewayDetails[toId]
      ) {
        return {
          ...state
        };
      }

      let newItems = state.gatewayDetails[toId]?.variables ? state.gatewayDetails[toId]?.variables.new || [] : [];

      const newArray = Object.keys(state.gatewayDetails[fromId]?.variables).map(i => {
        if (i === 'new') {
          return;
        }
        return state.gatewayDetails[fromId]?.variables[i];
      });

      newItems = [...newItems, ...newArray];
      const details = state.gatewayDetails[toId] || {variables: {}};
      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails,
          [toId]: {
            ...state.gatewayDetails[toId],
            variables: {...details.variables, new: [...newItems]},
          }
        },
      };
    }

    case DirectoryActionTypes.StorePasteDirectoryUserGatewayParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.gatewayDetails[fromId]
      ) {
        return {
          ...state
        };
      }

      let newItems = state.gatewayDetails[toId]?.parameters ? state.gatewayDetails[toId]?.parameters.new || [] : [];

      const newArray = Object.keys(state.gatewayDetails[fromId]?.parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.gatewayDetails[fromId]?.parameters[i];
      });

      newItems = [...newItems, ...newArray];
      const details = state.gatewayDetails[toId] || {parameters: {}};
      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails,
          [toId]: {
            ...state.gatewayDetails[toId],
            parameters: {...details.parameters, new: [...newItems]},
          }
        },
      };
    }

    case DirectoryActionTypes.StoreUpdateDirectoryUserParameter:
    case DirectoryActionTypes.StoreAddDirectoryUserParameter:
    case DirectoryActionTypes.StoreSwitchDirectoryUserParameter: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      let new_params = !!state.users[parentId]?.parameters ? state.users[parentId]?.parameters.new || [] : [];
      if (isNumeric(action.payload.param_index)) {
        new_params = [
          ...new_params.slice(0, action.payload.param_index),
          null,
          ...new_params.slice(action.payload.param_index + 1)
        ];
      }

      return {
        ...state,
        users: {...state.users,
          [parentId]: {
            ...state.users[parentId],
            parameters: {...state.users[parentId]?.parameters, [data.id]: data, new: [...new_params]}
          }
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryUserVariable:
    case DirectoryActionTypes.StoreAddDirectoryUserVariable:
    case DirectoryActionTypes.StoreSwitchDirectoryUserVariable: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      let new_vars = !!state.users[parentId]?.variables ? state.users[parentId]?.variables.new || [] : [];
      if (isNumeric(action.payload.var_index)) {
        new_vars = [
          ...new_vars.slice(0, action.payload.var_index),
          null,
          ...new_vars.slice(action.payload.var_index + 1)
        ];
      }

      return {
        ...state,
        users: {...state.users,
          [parentId]: {
            ...state.users[parentId],
            variables: {...state.users[parentId]?.variables, [data.id]: data, new: [...new_vars]},
          }
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: 0,
      };
    }
    case DirectoryActionTypes.StoreAddNewDirectoryUserParameter: {
      let oldArray = [];
      if (state.users[action.payload.id]?.parameters && state.users[action.payload.id]?.parameters.new) {
        oldArray = state.users[action.payload.id]?.parameters.new || [];
      }
      return {
        ...state,
        users: {
          ...state.users, [action.payload.id]: {
            ...state.users[action.payload.id],
            parameters: {
              ...state.users[action.payload.id]?.parameters || {},
              new: [...oldArray, {name: null, value: null}]
            }
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreAddNewDirectoryUserVariable: {
      let oldArray = [];
      if (state.users[action.payload.id]?.variables && state.users[action.payload.id]?.variables.new) {
        oldArray = state.users[action.payload.id]?.variables.new || [];
      }
      return {
        ...state,
        users: {
          ...state.users, [action.payload.id]: {
            ...state.users[action.payload.id],
            variables: {
              ...state.users[action.payload.id]?.variables || {},
              new: [...oldArray, {name: null, value: null}]
            }
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteNewDirectoryUserVariable: {
      let oldArray = [];
      if (state.users[action.payload.id]?.variables && state.users[action.payload.id]?.variables.new) {
        oldArray = state.users[action.payload.id]?.variables.new || [];
      }
      const newArray = [
        ...oldArray.slice(0, action.payload.index),
        null,
        ...oldArray.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        users: {
          ...state.users, [action.payload.id]: {
            ...state.users[action.payload.id],
            variables: {
              ...state.users[action.payload.id]?.variables || {},
              new: [...newArray]
            }
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteNewDirectoryUserParameter: {
      let oldArray = [];
      if (state.users[action.payload.id]?.parameters && state.users[action.payload.id]?.parameters.new) {
        oldArray = state.users[action.payload.id]?.parameters.new || [];
      }
      const newArray = [
        ...oldArray.slice(0, action.payload.index),
        null,
        ...oldArray.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        users: {
          ...state.users, [action.payload.id]: {
            ...state.users[action.payload.id],
            parameters: {
              ...state.users[action.payload.id]?.parameters || {},
              new: [...newArray]
            }
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryUserParameter: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.users[parentId]?.parameters;
      return {
        ...state,
        users: <Iusers>{
          ...state.users,
          [parentId]: {
            ...state.users[parentId],
            parameters: {
              ...rest
            },
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryUserVariable: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.users[parentId]?.variables;

      return {
        ...state,
        users: <Iusers>{
          ...state.users,
          [parentId]: {
            ...state.users[parentId],
            variables: {
              ...rest
            },
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryUserCache: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        users: {
          ...state.users, [data.id]: {
            ...state.users[data.id],
            cache: data.cache,
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryUserNumberAlias: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        users: {
          ...state.users, [data.id]: {
            ...state.users[data.id],
            number_alias: data.number_alias,
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryUserCidr: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        users: {
          ...state.users, [data.id]: {
            ...state.users[data.id],
            cidr: data.cidr,
          }
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreCreateWebDirectoryUsersByTemplate:
    case DirectoryActionTypes.StoreImportXMLDomainUser:
    case DirectoryActionTypes.StoreAddDirectoryUser: {
      const data = action.payload.response.data || {};
      if (!data) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        users: {
          ...state.users,
          ...data,
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreSwitchDirectoryUser:
    case DirectoryActionTypes.StoreUpdateDirectoryUserName: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        users: {
          ...state.users,
          [data.id]: {...state.users[data.id], ...data},
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryUser: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.users;
      return {
        ...state,
        users: {...rest},
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreGetDirectoryGroups: {
      const domains = action.payload.response.data['domains'] || {};
      const groups = action.payload.response.data['list'] || {};
      if (!domains) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        domains: domains || state.domains,
        groupNames: groups,
        errorMessage: action.payload.response.error || null,
        loadCounter: 0,
      };
    }
    case DirectoryActionTypes.StoreGetDirectoryGroupUsers: {
      const groupUsers = action.payload.response.data['group_users'] || {};
      const users = action.payload.response.data['users'] || {};
      if (!groupUsers) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        groupUsers: {...state.groupUsers, ...groupUsers},
        users: {...state.users, ...users},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryGroupName:
    case DirectoryActionTypes.StoreAddNewDirectoryGroup: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        groupNames: {
          ...state.groupNames,
          [data.id]: { ...state.groupNames[data.id], ...data},
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryGroup: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.groupNames;

      return {
        ...state,
        groupNames: rest,
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreAddDirectoryGroupUser: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        groupUsers: {
          ...state.groupUsers,
          [data.id]: { ...state.groupUsers[data.id], ...data},
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryGroupUser: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.groupUsers;

      return {
        ...state,
        groupUsers: rest,
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreGetDirectoryUserGateways: {
      const data = action.payload.response.data || {};
      return {
        ...state,
        domains: data['domains'],
        users: data['directory_users'],
        userGateways: data['user_gateways'],
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreGetDirectoryUserGatewayDetails: {
      const variables = action.payload.response.data.variables || {};
      const parameters = action.payload.response.data.parameters || {};
      const parentId = getParentId(variables) || getParentId(parameters);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails, [parentId]: {parameters: parameters, variables: variables}},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryUserGatewayParameter:
    case DirectoryActionTypes.StoreSwitchDirectoryUserGatewayParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        gatewayDetails: {
          ...state.gatewayDetails,
          [parentId]: {...state.gatewayDetails[parentId], parameters: {...state.gatewayDetails[parentId].parameters, [data.id]: data}}},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreSwitchDirectoryUserGatewayVariable:
    case DirectoryActionTypes.StoreUpdateDirectoryUserGatewayVariable: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        gatewayDetails: {
          ...state.gatewayDetails,
          [parentId]: {...state.gatewayDetails[parentId], variables: {...state.gatewayDetails[parentId].variables, [data.id]: data}}
          },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreDeleteDirectoryUserGatewayVariable: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.gatewayDetails[parentId].variables;

      return {
        ...state,
        gatewayDetails: <IGatewaydetails>{...state.gatewayDetails,
          [parentId]: {
          ...state.gatewayDetails[parentId],
            variables: rest}
          },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryUserGatewayParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.gatewayDetails[parentId].parameters;

      return {
        ...state,
        gatewayDetails: <IGatewaydetails>{...state.gatewayDetails,
          [parentId]: {
            ...state.gatewayDetails[parentId],
            parameters: rest}
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreAddDirectoryUserGatewayVariable: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.gatewayDetails[parentId].variables.new || []];
      if (action.payload.index !== undefined) {
        rest = [
          ...rest.slice(0, action.payload.index),
          null,
          ...rest.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails, [parentId]: {...state.gatewayDetails[parentId],
            variables: {...state.gatewayDetails[parentId].variables, [data.id]: data, new: rest}}},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreNewDirectoryUserGatewayParameter: {
      let details = state.gatewayDetails[action.payload.id];
      if (!details) {
        // return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
        details = {parameters: <Iparameters>{}, variables: <IgatewayVariables>{}};
      }
      const rest = [
        ...details.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        gatewayDetails: {
          ...state.gatewayDetails, [action.payload.id]: {
            ...details, parameters: {...details.parameters, new: rest},
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.DropNewDirectoryUserGatewayParameter: {
      const details = state.gatewayDetails[action.payload.id];
      if (!details) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...details.parameters.new.slice(0, action.payload.index),
        null,
        ...details.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        gatewayDetails: {
          ...state.gatewayDetails, [action.payload.id]: {
            ...details, parameters: {...details.parameters, new: rest},
          },
        },
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDirectoryUserGatewayParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.gatewayDetails[parentId].parameters.new || []];
      if (action.payload.index !== undefined) {
        rest = [
          ...rest.slice(0, action.payload.index),
          null,
          ...rest.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails, [parentId]: {...state.gatewayDetails[parentId],
            parameters: {...state.gatewayDetails[parentId].parameters, [data.id]: data, new: rest}}},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreUpdateDirectoryUserGatewayName:
    case DirectoryActionTypes.StoreAddDirectoryUserGateway: {
      const data = action.payload.response.data || {};
      const parentId: number = data.parent?.id || 0;
      if (!data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        userGateways: {
          ...state.userGateways,
          [data.id]: {...state.userGateways[data.id], ...data},
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case DirectoryActionTypes.StoreDeleteDirectoryUserGateway: {
      const data = action.payload.response.data || {};
      if (data.id === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = state.userGateways;
      return {
        ...state,
        userGateways: rest,
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreNewDirectoryUserGatewayVariable: {
      let details = state.gatewayDetails[action.payload.id];
      if (!details) {
        // return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
        details = {parameters: <Iparameters>{}, variables: <IgatewayVariables>{}};
      }
      if (!details.variables) {
        details.variables = {new: []};
      }
      const rest = [
        ...details.variables?.new || [],
        <IdirectionItem>{}
      ];
      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails, [action.payload.id]:
            {...details, variables: {...details.variables,  new: rest}}},
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.DropNewDirectoryUserGatewayVariable: {
      const details = state.gatewayDetails[action.payload.id];
      if (!details) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...details.variables.new.slice(0, action.payload.index),
        null,
        ...details.variables.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        gatewayDetails: {...state.gatewayDetails, [action.payload.id]:
            {...details, variables: {...details.variables,  new: rest}}},
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DirectoryActionTypes.StoreGetWebUsersByDirectory: {
      return {
        ...state,
        additionalData: action.payload.response['additional_data'],
        errorMessage: action.payload.response.error || null,
        loadCounter: 0,
      };
    }

    case SettingsActionTypes.STORE_CLEAR_WEB_USER_AVATAR:
    case SettingsActionTypes.STORE_UPDATE_WEB_USER_AVATAR: {
      const data = action.payload.response['web_users'];
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
      let format = data[id]?.avatar_format;
      if (format.length) {
        format = format + '?' + (+new Date());
      }
      return {
        ...state,
        additionalData: {...state.additionalData, [id]: {
            ...state.additionalData[id],
            avatar_format: format
          }},
        errorMessage: action.payload.response.error || null,
      };
    }

    case DirectoryActionTypes.StoreGetWebDirectoryUsersTemplatesList: {
      const data = action.payload.response['data'] || {};
      return {
        ...state,
        webUsersTemplates: data,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case DirectoryActionTypes.StoreGetWebDirectoryUsersTemplateForm: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        templatesItems: {...state.templatesItems, ...data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    default: {
      return state;
    }
  }
}

function getParentId(data: any): number {
  if (data.id) {
    return data?.parent?.id || 0;
  } else {
    const firstKey = Object.keys(data)[0];
    return data[firstKey]?.parent?.id || 0;
  }
}
