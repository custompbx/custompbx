import {All, ConfigActionTypes} from './config.actions.sofia';
import {
  IdirectionItem,
  Igateways,
  Iitem,
  initialState,
  Isofia,
  State
} from '../config.state.struct';
import {getParentId} from '../config.reducers';


export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GET_SOFIA_GLOBAL_SETTINGS:
    case ConfigActionTypes.GET_SOFIA_PROFILES:
    case ConfigActionTypes.GET_SOFIA_PROFILES_PARAMS:
    case ConfigActionTypes.UPDATE_SOFIA_GLOBAL_SETTING:
    case ConfigActionTypes.SWITCH_SOFIA_GLOBAL_SETTING:
    case ConfigActionTypes.ADD_SOFIA_GLOBAL_SETTING:
    case ConfigActionTypes.DEL_SOFIA_GLOBAL_SETTING:
    case ConfigActionTypes.UPDATE_SOFIA_PROFILE_PARAM:
    case ConfigActionTypes.SWITCH_SOFIA_PROFILE_PARAM:
    case ConfigActionTypes.ADD_SOFIA_PROFILE_PARAM:
    case ConfigActionTypes.DEL_SOFIA_PROFILE_PARAM:
    case ConfigActionTypes.GET_SOFIA_PROFILE_GATEWAYS:
    case ConfigActionTypes.UPDATE_SOFIA_PROFILE_GATEWAY_PARAM:
    case ConfigActionTypes.SWITCH_SOFIA_PROFILE_GATEWAY_PARAM:
    case ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY_PARAM:
    case ConfigActionTypes.UPDATE_SOFIA_PROFILE_GATEWAY_VAR:
    case ConfigActionTypes.SWITCH_SOFIA_PROFILE_GATEWAY_VAR:
    case ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY_VAR:
    case ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY_VAR:
    case ConfigActionTypes.GET_SOFIA_PROFILE_DOMAINS:
    case ConfigActionTypes.UPDATE_SOFIA_PROFILE_DOMAIN:
    case ConfigActionTypes.SWITCH_SOFIA_PROFILE_DOMAIN:
    case ConfigActionTypes.ADD_SOFIA_PROFILE_DOMAIN:
    case ConfigActionTypes.DEL_SOFIA_PROFILE_DOMAIN:
    case ConfigActionTypes.GET_SOFIA_PROFILE_ALIASES:
    case ConfigActionTypes.UPDATE_SOFIA_PROFILE_ALIAS:
    case ConfigActionTypes.SWITCH_SOFIA_PROFILE_ALIAS:
    case ConfigActionTypes.ADD_SOFIA_PROFILE_ALIAS:
    case ConfigActionTypes.DEL_SOFIA_PROFILE_ALIAS:
    case ConfigActionTypes.ADD_SOFIA_PROFILE:
    case ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY:
    case ConfigActionTypes.DEL_SOFIA_PROFILE:
    case ConfigActionTypes.RENAME_SOFIA_PROFILE:
    case ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY:
    case ConfigActionTypes.RENAME_SOFIA_PROFILE_GATEWAY:
    // case ConfigActionTypes.SOFIA_PROFILE_COMMAND:
    case ConfigActionTypes.SWITCH_SOFIA_PROFILE:
    case ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY_PARAM: {
      return {...state,
        sofia: {
          ...state.sofia,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotSofiaError: {
      return {
        ...state,
        sofia: {
          ...state.sofia,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_GET_SOFIA_PROFILES: {
      const data = action.payload.response.data;
      if (action.payload.response.exists === false) {
        return {
          ...state,
          sofia: {...state.sofia, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.sofia) {
        state.sofia = <Isofia>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {...state.sofia.profiles, ...data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE:
    case ConfigActionTypes.STORE_RENAME_SOFIA_PROFILE: {
      const data = action.payload.response.data;
      if (action.payload.response.exists === false) {
        return {
          ...state,
          sofia: {...state.sofia, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.sofia) {
        state.sofia = <Isofia>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {...state.sofia.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_GLOBAL_SETTING: {
      let data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      data = {[data.id]: data};
      let rest = [...state.sofia.global_settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.sofia.global_settings.new.slice(0, action.payload.index),
          null,
          ...state.sofia.global_settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, global_settings: {...state.sofia.global_settings, ...data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE: {
      const data = action.payload.response.data;
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.sofia.profiles[data.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.sofia.profiles;

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_GLOBAL_SETTING: {
      const id = action.payload.response.data?.id || 0;
      if (!state.sofia.global_settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.sofia.global_settings;

      return {
        ...state,
        sofia: {
          ...state.sofia, global_settings: {...rest, new: state.sofia.global_settings?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_GET_SOFIA_GLOBAL_SETTINGS: {
      const data = action.payload.response.data || {};
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, global_settings: {...state.sofia.global_settings, ...data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_SOFIA_GLOBAL_SETTING:
    case ConfigActionTypes.STORE_UPDATE_SOFIA_GLOBAL_SETTING: {
      const data = action.payload.response.data;
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, global_settings: {...state.sofia.global_settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_SOFIA_GLOBAL_SETTING: {
      const rest = [
        ...state.sofia.global_settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        sofia: {
          ...state.sofia, global_settings: {...state.sofia.global_settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_SOFIA_GLOBAL_SETTING: {
      const rest = [
        ...state.sofia.global_settings.new.slice(0, action.payload.index),
        null,
        ...state.sofia.global_settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        sofia: {
          ...state.sofia, global_settings: {...state.sofia.global_settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_PARAM: {
      const data = action.payload.response.data;
      const profileId = getParentId(data);
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.parameters?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.parameters.new.slice(0, action.payload.index),
          null,
          ...profile.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [profileId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_PASTE_SOFIA_PROFILE_PARAMS: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.sofia.profiles[fromId] || !state.sofia.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.sofia.profiles[toId].parameters ? state.sofia.profiles[toId].parameters?.new || [] : [];

      const newArray = Object.keys(state.sofia.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.sofia.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        sofia: {
          ...state.sofia,
          profiles: {
            ...state.sofia.profiles,
            [toId]: {
              ...state.sofia.profiles[toId],
              parameters: {
                ...state.sofia.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_PARAM: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.sofia.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_GET_SOFIA_PROFILES_PARAMS: {
      const data = action.payload.response.data;
      const profileId = getParentId(data);
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [profileId]:
              {...state.sofia.profiles[profileId], parameters: {...state.sofia.profiles[profileId].parameters, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_PARAM:
    case ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_PARAM: {
      const data = action.payload.response.data;
      const profileId = getParentId(data);
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [profileId]:
              {...state.sofia.profiles[profileId], parameters: {...state.sofia.profiles[profileId].parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_PARAM: {
      const profile = state.sofia.profiles[action.payload.id];
      console.log(profile);
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_PARAM: {
      const profile = state.sofia.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters.new.slice(0, action.payload.index),
        null,
        ...profile.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_RENAME_SOFIA_PROFILE_GATEWAY:
    case ConfigActionTypes.STORE_GET_SOFIA_PROFILE_GATEWAYS: {
      const data = action.payload.response.data || {};
      const ids = Object.keys(data);

      if (ids.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profileId = data[ids[0]].parent?.id || 0;
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {...state.sofia.profiles, [profileId]: {...state.sofia.profiles[profileId], gateways: data} },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
/*    case ConfigActionTypes.STORE_SOFIA_GATEWAY_UPDATES: {
      const profileId = action.payload.response.id;
      if (!state.sofia.profiles || !state.sofia.profiles[profileId]) {
        return {...state};
      }
      const data = action.payload.response.data['sofia_gateway'];
      if (!data.id || !state.sofia.profiles[profileId].gateways || !state.sofia.profiles[profileId].gateways[data.id]) {
        return {...state};
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia,
          profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...state.sofia.profiles[profileId],
              gateways: {
                ...state.sofia.profiles[profileId].gateways,
                [data.id]: {...state.sofia.profiles[profileId].gateways[data.id], ...data}
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }*/
    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_GATEWAY: {
      const data = action.payload.response.data || {};
      const profile = state.sofia.profiles[data?.parent?.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const id = action.payload.response.data.id;
      const gateway = profile.gateways[id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = profile.gateways;

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {...state.sofia.profiles, [profile.id]: {...profile, gateways: {...<Igateways>rest}}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreGetSofiaProfileGatewayVariables:
    case ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_GATEWAY_VAR:
    case ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_GATEWAY_VAR: {
      let data = action.payload.response.data || {};
      if (data && data.id) {
        data = {[data.id]: data};
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      let profileId = 0;
      const gatewayId = data[ids[0]].parent?.id || 0;
      Object.keys(state.sofia.profiles).forEach(
        (key) => {
          if (state.sofia.profiles[key]?.gateways[gatewayId]) {
            profileId = Number(key);
          }
        }
      );
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...state.sofia.profiles[profileId], gateways: {
                ...state.sofia.profiles[profileId].gateways, [gatewayId]:
                  {
                    ...state.sofia.profiles[profileId].gateways[gatewayId],
                    variables: {...state.sofia.profiles[profileId].gateways[gatewayId].variables, ...data}
                  }
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetSofiaProfileGatewayParameters:
    case ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_GATEWAY_PARAM:
    case ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_GATEWAY_PARAM: {
      let data = action.payload.response.data || {};
      if (data && data.id) {
        data = {[data.id]: data};
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      let profileId = 0;
      const gatewayId = data[ids[0]].parent?.id || 0;
      Object.keys(state.sofia.profiles).forEach(
        (key) => {
          if (state.sofia.profiles[key] && state.sofia.profiles[key]?.gateways && state.sofia.profiles[key]?.gateways[gatewayId]) {
            profileId = Number(key);
          }
        }
      );
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...state.sofia.profiles[profileId], gateways: {
                ...state.sofia.profiles[profileId].gateways, [gatewayId]:
                  {
                    ...state.sofia.profiles[profileId].gateways[gatewayId],
                    parameters: {...state.sofia.profiles[profileId].gateways[gatewayId].parameters, ...data},
                  }
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_PASTE_SOFIA_GATEWAY_PARAMS: {
      const profileId = action.payload.id;
      const toProfile = state.sofia.profiles[profileId];
      if (!toProfile) {
        return {...state};
      }
      const fromProfileId = action.payload.from_profile;
      const fromProfile = state.sofia.profiles[fromProfileId];
      if (!fromProfile) {
        return {...state};
      }
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !fromProfile.gateways[fromId] || !toProfile.gateways[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = toProfile.gateways[toId].parameters ? toProfile.gateways[toId].parameters?.new || [] : [];

      const newArray = Object.keys(fromProfile.gateways[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return fromProfile.gateways[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        sofia: {
          ...state.sofia,
          profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...toProfile,
              gateways: {
                ...toProfile.gateways,
                [toId]: {
                  ...toProfile.gateways[toId],
                  parameters: {
                    ...toProfile.gateways[toId].parameters,
                    new: [...new_items],
                  }
                }
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.STORE_PASTE_SOFIA_GATEWAY_VARS: {
      const profileId = action.payload.id;
      const toProfile = state.sofia.profiles[profileId];
      if (!toProfile) {
        return {...state};
      }
      const fromProfileId = action.payload.from_profile;
      const fromProfile = state.sofia.profiles[fromProfileId];
      if (!fromProfile) {
        return {...state};
      }
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !fromProfile.gateways[fromId] || !toProfile.gateways[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = toProfile.gateways[toId].variables ? toProfile.gateways[toId].variables?.new || [] : [];

      const newArray = Object.keys(fromProfile.gateways[fromId].variables).map(i => {
        if (i === 'new') {
          return;
        }
        return fromProfile.gateways[fromId].variables[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        sofia: {
          ...state.sofia,
          profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...toProfile,
              gateways: {
                ...toProfile.gateways,
                [toId]: {
                  ...toProfile.gateways[toId],
                  variables: {
                    ...toProfile.gateways[toId].variables,
                    new: [...new_items],
                  }
                }
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_GATEWAY_PARAM: {
      const data = action.payload.response.data;
      const gatewayId = getParentId(data);
      if (gatewayId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let profileId = 0;
      Object.keys(state.sofia.profiles).forEach(
        (key) => {
          if (state.sofia.profiles[key] && state.sofia.profiles[key]?.gateways && state.sofia.profiles[key]?.gateways[gatewayId]) {
            profileId = Number(key);
          }
        }
      );
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.sofia.profiles[profileId];
      const gateway = profile.gateways[gatewayId];
      let rest = [...gateway.parameters?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...gateway.parameters.new.slice(0, action.payload.index),
          null,
          ...gateway.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...profile, gateways: {
                ...profile.gateways, [gatewayId]:
                  {
                    ...profile.gateways[gatewayId],
                    parameters: {...profile.gateways[gatewayId].parameters, [data.id]: data, new: rest}
                  }
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_GATEWAY_PARAM: {
      const data = action.payload.response.data || {};
      const gatewayId = data?.parent?.id || 0;
      if (gatewayId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let profile = null;
      Object.values(state.sofia.profiles).forEach(
        (profileItem) => {
          if (profileItem.gateways && profileItem.gateways[gatewayId]) {
            profile = profileItem;
          }
        });
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const gateway = profile.gateways[gatewayId];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = gateway.parameters;

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profile.id]: {
              ...profile, gateways: {
                ...profile.gateways, [gatewayId]:
                  {
                    ...gateway,
                    parameters: {...rest}
                  }
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_GATEWAY_PARAM: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const gateway = profile.gateways[action.payload.id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...gateway.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {
                ...profile, gateways: {
                  ...profile.gateways, [action.payload.id]:
                    {...gateway, parameters: {...gateway.parameters, new: rest}}
                }
              }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_GATEWAY_PARAM: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const gateway = profile.gateways[action.payload.id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...gateway.parameters.new.slice(0, action.payload.index),
        null,
        ...gateway.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {
                ...profile, gateways: {
                  ...profile.gateways, [action.payload.id]:
                    {...gateway, parameters: {...gateway.parameters, new: rest}}
                }
              }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_GATEWAY_VAR: {
      const profileId = action.payload.response.id;
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const data = action.payload.response.data || {};
      const ids = Object.keys(data);
      if (ids.length > 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const id = ids[0];
      const gateway = profile.gateways[id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...gateway.variables?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...gateway.variables.new.slice(0, action.payload.index),
          null,
          ...gateway.variables.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...profile, gateways: {
                ...profile.gateways, [id]:
                  {
                    ...profile.gateways[id],
                    variables: {...profile.gateways[id].variables, ...data[id].variables, new: rest}
                  }
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_GATEWAY_VAR: {
      const profileId = action.payload.response.id;
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const data = action.payload.response.data || {};
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const id = ids[0];
      const gateway = profile.gateways[id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const paramIds = Object.keys(data[id].variables);
      if (paramIds.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const paramId = paramIds[0];
      const {[paramId]: toDel, ...rest} = gateway.variables;

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {
              ...profile, gateways: {
                ...profile.gateways, [id]:
                  {
                    ...profile.gateways[id],
                    variables: {...rest}
                  }
              }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_GATEWAY_VAR: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const gateway = profile.gateways[action.payload.id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...gateway.variables?.new || [],
        <IdirectionItem>{}
      ];
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {
                ...profile, gateways: {
                  ...profile.gateways, [action.payload.id]:
                    {...gateway, variables: {...gateway.variables, new: rest}}
                }
              }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_GATEWAY_VAR: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const gateway = profile.gateways[action.payload.id];
      if (!gateway) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...gateway.variables.new.slice(0, action.payload.index),
        null,
        ...gateway.variables.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {
                ...profile, gateways: {
                  ...profile.gateways, [action.payload.id]:
                    {...gateway, variables: {...gateway.variables, new: rest}}
                }
              }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_GET_SOFIA_PROFILE_DOMAINS: {
      const data = action.payload.response.data;
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profileId = data[ids[0]]?.parent.id || 0;
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [profileId]:
              {...profile, domains: {...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_DOMAIN:
    case ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_DOMAIN: {
      const data = action.payload.response.data || {};
      const profile = state.sofia.profiles[data.parent?.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profile.id]: {...profile, domains: {...profile.domains, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_DOMAIN: {
      const data = action.payload.response.data;
      const profileId = getParentId(data);
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.domains?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.domains.new.slice(0, action.payload.index),
          null,
          ...profile.domains.new.slice(action.payload.index + 1)
        ];
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {...profile, domains: {...profile.domains, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_DOMAIN: {
      const data = action.payload.response.data || {};
      const profile = state.sofia.profiles[data.parent?.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.domains;
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profile.id]: {...profile, domains: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_DOMAIN: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.domains?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {...profile, domains: {...profile.domains, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_DOMAIN: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.domains?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.domains.new.slice(0, action.payload.index),
          null,
          ...profile.domains.new.slice(action.payload.index + 1)
        ];
      }
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {...profile, domains: {...profile.domains, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }


    case ConfigActionTypes.STORE_GET_SOFIA_PROFILE_ALIASES: {
      const data = action.payload.response.data || {};
      const profile = state.sofia.profiles[data.parent?.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [profile.id]:
              {...profile, aliases: {...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_SWITCH_SOFIA_PROFILE_ALIAS:
    case ConfigActionTypes.STORE_UPDATE_SOFIA_PROFILE_ALIAS: {
      const data = action.payload.response.data || {};
      const profile = state.sofia.profiles[data.parent?.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profile.id]: {...profile, aliases: {...profile.aliases, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_ALIAS: {
      const data = action.payload.response.data;
      const profileId = getParentId(data);
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.sofia.profiles[profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.aliases?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.aliases.new.slice(0, action.payload.index),
          null,
          ...profile.aliases.new.slice(action.payload.index + 1)
        ];
      }
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profileId]: {...profile, aliases: {...profile.aliases, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DEL_SOFIA_PROFILE_ALIAS: {
      const data = action.payload.response.data || {};
      const profile = state.sofia.profiles[data.parent?.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const {[data.id]: toDel, ...rest} = profile.aliases;
      return {
        ...state,
        sofia: <Isofia>{
          ...state.sofia, profiles: {
            ...state.sofia.profiles,
            [profile.id]: {...profile, aliases: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_NEW_SOFIA_PROFILE_ALIAS: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.aliases?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {...profile, aliases: {...profile.aliases, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_DROP_NEW_SOFIA_PROFILE_ALIAS: {
      const profile = state.sofia.profiles[action.payload.profileId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.aliases?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.aliases.new.slice(0, action.payload.index),
          null,
          ...profile.aliases.new.slice(action.payload.index + 1)
        ];
      }
      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [action.payload.profileId]:
              {...profile, aliases: {...profile.aliases, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {...state.sofia.profiles, [data.id]: data},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.STORE_ADD_SOFIA_PROFILE_GATEWAY: {
      const data = action.payload.response.data;
      const profileId = getParentId(data);
      if (profileId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        sofia: {
          ...state.sofia, profiles: {
            ...state.sofia.profiles, [profileId]: {
              ...state.sofia.profiles[profileId], gateways:
                {...state.sofia.profiles[profileId].gateways, [data.id]: data}
            }
          }
          ,
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return null;
    }
  }
}
