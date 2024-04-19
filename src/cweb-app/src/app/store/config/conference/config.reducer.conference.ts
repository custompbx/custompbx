import {Iitem, Iconference, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.conference';
import {getParentId} from "../config.reducers";

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetConference:
    case ConfigActionTypes.UpdateConferenceRoom:
    case ConfigActionTypes.SwitchConferenceRoom:
    case ConfigActionTypes.AddConferenceRoom:
    case ConfigActionTypes.DelConferenceRoom:
    case ConfigActionTypes.GetConferenceCallerControls:
    case ConfigActionTypes.AddConferenceCallerControl:
    case ConfigActionTypes.UpdateConferenceCallerControl:
    case ConfigActionTypes.SwitchConferenceCallerControl:
    case ConfigActionTypes.DelConferenceCallerControl:
    case ConfigActionTypes.AddConferenceCallerControlGroup:
    case ConfigActionTypes.DelConferenceCallerControlGroup:
    case ConfigActionTypes.UpdateConferenceCallerControlGroup:
    case ConfigActionTypes.GetConferenceProfileParameters:
    case ConfigActionTypes.AddConferenceProfileParameter:
    case ConfigActionTypes.UpdateConferenceProfileParameter:
    case ConfigActionTypes.SwitchConferenceProfileParameter:
    case ConfigActionTypes.DelConferenceProfileParameter:
    case ConfigActionTypes.AddConferenceProfile:
    case ConfigActionTypes.DelConferenceProfile:
    case ConfigActionTypes.UpdateConferenceProfile:
    case ConfigActionTypes.GetConferenceChatPermissionUsers:
    case ConfigActionTypes.AddConferenceChatPermissionUser:
    case ConfigActionTypes.UpdateConferenceChatPermissionUser:
    case ConfigActionTypes.SwitchConferenceChatPermissionUser:
    case ConfigActionTypes.DelConferenceChatPermissionUser:
    case ConfigActionTypes.AddConferenceChatPermission:
    case ConfigActionTypes.DelConferenceChatPermission:
    case ConfigActionTypes.UpdateConferenceChatPermission: {
      return {...state,
        conference: {
          ...state.conference,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotConferenceError: {
      return {
        ...state,
        conference: {
          ...state.conference,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchConferenceProfile:
    case ConfigActionTypes.StoreGetConference: {
      const advertise = action.payload.response.data['conference_rooms'] || {};
      const ccGroups = action.payload.response.data['conference_caller_control_groups'] || {};
      const profiles = action.payload.response.data['conference_profiles'] || {};
      const chatProfiles = action.payload.response.data['conference_chat_permissions_profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          conference: {...state.conference, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          advertise: {...state.conference.advertise, ...advertise},
          caller_controls: {...state.conference.caller_controls, ...ccGroups},
          profiles: {...state.conference.profiles, ...profiles},
          chat_profiles: {...state.conference.chat_profiles, ...chatProfiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddConferenceProfile:
    case ConfigActionTypes.StoreUpdateConferenceProfile: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          profiles: {...state.conference.profiles, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddConferenceChatPermission:
    case ConfigActionTypes.StoreUpdateConferenceChatPermission: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          chat_profiles: {...state.conference.chat_profiles, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreAddConferenceCallerControlGroup:
    case ConfigActionTypes.StoreUpdateConferenceCallerControlGroup: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          caller_controls: {...state.conference.caller_controls, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelConferenceRoom: {
      const id = action.payload.response.data?.id || 0;
      if (!state.conference.advertise[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.conference.advertise;

      return {
        ...state,
        conference: {
          ...state.conference, advertise: {...rest, new: state.conference.advertise?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchConferenceRoom:
    case ConfigActionTypes.StoreUpdateConferenceRoom: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, advertise: {...state.conference.advertise, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewConferenceRoom: {
      const rest = [
        ...state.conference.advertise?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        conference: {
          ...state.conference, advertise: {...state.conference.advertise, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewConferenceRoom: {
      const rest = [
        ...state.conference.advertise.new.slice(0, action.payload.index),
        null,
        ...state.conference.advertise.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        conference: {
          ...state.conference, advertise: {...state.conference.advertise, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddConferenceRoom: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...state.conference.advertise?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.conference.advertise.new.slice(0, action.payload.index),
          null,
          ...state.conference.advertise.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, advertise: {...state.conference.advertise, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddConferenceCallerControl: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const group = state.conference.caller_controls[parentId];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...group.controls?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...group.controls.new.slice(0, action.payload.index),
          null,
          ...group.controls.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, caller_controls: {
            ...state.conference.caller_controls, [parentId]:
              {...group, controls: {...group.controls, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteConferenceCallerControls: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.conference.caller_controls[fromId] || !state.conference.caller_controls[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.conference.caller_controls[toId].controls ? state.conference.caller_controls[toId].controls?.new || [] : [];

      const newArray = Object.keys(state.conference.caller_controls[fromId].controls).map(i => {
        if (i === 'new') {
          return;
        }
        return state.conference.caller_controls[fromId].controls[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        conference: {
          ...state.conference,
          caller_controls: {
            ...state.conference.caller_controls,
            [toId]: {
              ...state.conference.caller_controls[toId],
              controls: {
                ...state.conference.caller_controls[toId].controls,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelConferenceCallerControl: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const group = state.conference.caller_controls[parentId];
      if (!group || !group.controls[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = group.controls;

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, caller_controls: {
            ...state.conference.caller_controls, [parentId]:
              {...group, controls: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetConferenceCallerControls: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const group = state.conference.caller_controls[parentId];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, caller_controls: {
            ...state.conference.caller_controls, [parentId]:
              {...group, controls: {...group.controls, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchConferenceCallerControl:
    case ConfigActionTypes.StoreUpdateConferenceCallerControl: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const group = state.conference.caller_controls[parentId];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, caller_controls: {
            ...state.conference.caller_controls, [parentId]:
              {...group, controls: {...group.controls, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewConferenceCallerControl: {
      const group = state.conference.caller_controls[action.payload.id];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...group.controls?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        conference: {
          ...state.conference, caller_controls: {
            ...state.conference.caller_controls, [action.payload.id]:
              {...group, controls: {...group.controls, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewConferenceCallerControl: {
      const group = state.conference.caller_controls[action.payload.id];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...group.controls.new.slice(0, action.payload.index),
        null,
        ...group.controls.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        conference: {
          ...state.conference, caller_controls: {
            ...state.conference.caller_controls, [action.payload.id]:
              {...group, controls: {...group.controls, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelConferenceCallerControlGroup: {
      const id = action.payload.response.data?.id || 0;
      const group = state.conference.caller_controls[id];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.conference.caller_controls;

      return {
        ...state,
        conference: {
          ...state.conference, caller_controls: {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddConferenceProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.profiles[parentId];
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
        conference: <Iconference>{
          ...state.conference, profiles: {
            ...state.conference.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteConferenceProfileParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.conference.profiles[fromId] || !state.conference.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.conference.profiles[toId].parameters ? state.conference.profiles[toId].parameters?.new || [] : [];

      const newArray = Object.keys(state.conference.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.conference.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        conference: {
          ...state.conference,
          profiles: {
            ...state.conference.profiles,
            [toId]: {
              ...state.conference.profiles[toId],
              parameters: {
                ...state.conference.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelConferenceProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.profiles[parentId];
      if (!profile || !profile.parameters[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, profiles: {
            ...state.conference.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetConferenceProfileParameters: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, profiles: {
            ...state.conference.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchConferenceProfileParameter:
    case ConfigActionTypes.StoreUpdateConferenceProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, profiles: {
            ...state.conference.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewConferenceProfileParameter: {
      const profile = state.conference.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        conference: {
          ...state.conference, profiles: {
            ...state.conference.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewConferenceProfileParameter: {
      const profile = state.conference.profiles[action.payload.id];
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
        conference: {
          ...state.conference, profiles: {
            ...state.conference.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelConferenceProfile: {
      const id = action.payload.response.data?.id || 0;
      const profile = state.conference.profiles[id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.conference.profiles;

      return {
        ...state,
        conference: {
          ...state.conference, profiles: {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddConferenceChatPermissionUser: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.chat_profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.users?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.users.new.slice(0, action.payload.index),
          null,
          ...profile.users.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, chat_profiles: {
            ...state.conference.chat_profiles, [parentId]:
              {...profile, users: {...profile.users, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteConferenceChatPermissionUsers: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.conference.chat_profiles[fromId] || !state.conference.chat_profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.conference.chat_profiles[toId].users ? state.conference.chat_profiles[toId].users?.new || [] : [];

      const newArray = Object.keys(state.conference.chat_profiles[fromId].users).map(i => {
        if (i === 'new') {
          return;
        }
        return state.conference.chat_profiles[fromId].users[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        conference: {
          ...state.conference,
          chat_profiles: {
            ...state.conference.chat_profiles,
            [toId]: {
              ...state.conference.chat_profiles[toId],
              users: {
                ...state.conference.chat_profiles[toId].users,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelConferenceChatPermissionUser: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.chat_profiles[parentId];
      if (!profile || !profile.users[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.users;

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, chat_profiles: {
            ...state.conference.chat_profiles, [parentId]:
              {...profile, users: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetConferenceChatPermissionUsers: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.chat_profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, chat_profiles: {
            ...state.conference.chat_profiles, [parentId]:
              {...profile, users: {...profile.users, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchConferenceChatPermissionUser:
    case ConfigActionTypes.StoreUpdateConferenceChatPermissionUser: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.chat_profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, chat_profiles: {
            ...state.conference.chat_profiles, [parentId]:
              {...profile, users: {...profile.users, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewConferenceChatPermissionUser: {
      const profile = state.conference.chat_profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.users?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        conference: {
          ...state.conference, chat_profiles: {
            ...state.conference.chat_profiles, [action.payload.id]:
              {...profile, users: {...profile.users, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewConferenceChatPermissionUser: {
      const profile = state.conference.chat_profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.users.new.slice(0, action.payload.index),
        null,
        ...profile.users.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        conference: {
          ...state.conference, chat_profiles: {
            ...state.conference.chat_profiles, [action.payload.id]:
              {...profile, users: {...profile.users, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelConferenceChatPermission: {
      const id = action.payload.response.data?.id || 0;
      const profile = state.conference.chat_profiles[id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.conference.chat_profiles;

      return {
        ...state,
        conference: {
          ...state.conference, chat_profiles: {...rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return null;
    }
  }
}
