import {Iitem, Iconference, initialState, State, Ilayouts} from '../config.state.struct';
import {
  AddConferenceLayout, AddConferenceLayoutGroup, AddConferenceLayoutGroupLayout, AddConferenceLayoutImage,
  All,
  ConfigActionTypes,
  DelConferenceLayout,
  DelConferenceLayoutGroup,
  DelConferenceLayoutGroupLayout,
  DelConferenceLayoutImage,
  GetConferenceLayoutGroupLayouts,
  GetConferenceLayouts,
  StoreAddConferenceLayout,
  StoreAddConferenceLayoutGroup,
  StoreAddConferenceLayoutGroupLayout,
  StoreAddConferenceLayoutImage,
  StoreConferenceError,
  StoreDelConferenceLayout,
  StoreDelConferenceLayoutGroup,
  StoreDelConferenceLayoutGroupLayout,
  StoreDelConferenceLayoutImage,
  StoreDropConferenceLayoutGroupLayout,
  StoreDropConferenceLayoutImage,
  StoreGetConferenceLayoutGroupLayouts,
  StoreGetConferenceLayoutImages,
  StoreGetConferenceLayouts,
  StoreNewConferenceLayoutGroupLayout,
  StoreNewConferenceLayoutImage,
  StoreSwitchConferenceLayoutGroupLayout,
  StoreSwitchConferenceLayoutImage,
  StoreUpdateConferenceLayout,
  StoreUpdateConferenceLayout3D,
  StoreUpdateConferenceLayoutGroup,
  StoreUpdateConferenceLayoutGroupLayout,
  StoreUpdateConferenceLayoutImage,
  SwitchConferenceLayout,
  SwitchConferenceLayoutGroupLayout,
  SwitchConferenceLayoutImage,
  UpdateConferenceLayout,
  UpdateConferenceLayout3D,
  UpdateConferenceLayoutGroup,
  UpdateConferenceLayoutGroupLayout,
  UpdateConferenceLayoutImage
} from './config.actions.conference';
import {getParentId} from '../config.reducers';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case GetConferenceLayouts.type:
    case GetConferenceLayoutGroupLayouts.type:
    case UpdateConferenceLayout.type:
    case UpdateConferenceLayoutGroup.type:
    case UpdateConferenceLayout3D.type:
    case UpdateConferenceLayoutImage.type:
    case UpdateConferenceLayoutGroupLayout.type:
    case DelConferenceLayout.type:
    case DelConferenceLayoutGroup.type:
    case DelConferenceLayoutGroupLayout.type:
    case DelConferenceLayoutImage.type:
    case SwitchConferenceLayoutImage.type:
    case SwitchConferenceLayout.type:
    case SwitchConferenceLayoutGroupLayout.type:
    case AddConferenceLayout.type:
    case AddConferenceLayoutGroup.type:
    case AddConferenceLayoutGroupLayout.type:
    case AddConferenceLayoutImage.type:
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
      return {
        ...state,
        conference: {
          ...state.conference,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1
      };
    }

    case StoreConferenceError.type:
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
          layouts: {...state.conference.layouts},
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

    case StoreGetConferenceLayouts.type: {
      const layouts = action.payload.response.data['conference_layouts'] || {};
      const groups = action.payload.response.data['conference_layouts_groups'] || {};

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          layouts: {
            ...state.conference.layouts,
            conference_layouts: {...layouts},
            conference_layouts_groups: {...groups},
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreGetConferenceLayoutImages.type: {
      const {data} = action.payload.response || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.layouts.conference_layouts[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts: {
              ...state.conference.layouts.conference_layouts,
              [parentId]:
                {...profile, images: {...data}}
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreGetConferenceLayoutGroupLayouts.type: {
      const {data} = action.payload.response || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.layouts.conference_layouts_groups[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts_groups: {
              ...state.conference.layouts.conference_layouts_groups,
              [parentId]:
                {...profile, layouts: {...data}}
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreNewConferenceLayoutImage.type: {
      const profile = state.conference.layouts.conference_layouts[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const rest = profile.new || [];
      return {
        ...state,
        conference: {
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts: {
              ...state.conference.layouts.conference_layouts,
              [action.payload.id]:
                {...profile, new: [...rest, {}]}
            }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreDropConferenceLayoutImage.type: {
      const profile = state.conference.layouts.conference_layouts[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.new.slice(0, action.payload.index),
        null,
        ...profile.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        conference: {
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts: {
              ...state.conference.layouts.conference_layouts,
              [action.payload.id]:
                {...profile, new: rest}
            }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreNewConferenceLayoutGroupLayout.type: {
      const profile = state.conference.layouts.conference_layouts_groups[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = profile.new || [];
      return {
        ...state,
        conference: {
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts_groups: {
              ...state.conference.layouts.conference_layouts_groups,
              [action.payload.id]:
                {...profile, new: [ ...rest, {} ]}
            }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreDropConferenceLayoutGroupLayout.type: {
      const profile = state.conference.layouts.conference_layouts_groups[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.new.slice(0, action.payload.index),
        null,
        ...profile.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        conference: {
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts_groups: {
              ...state.conference.layouts.conference_layouts_groups,
              [action.payload.id]:
                {...profile, new: rest}
            }
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreSwitchConferenceLayoutImage.type:
    case StoreUpdateConferenceLayoutImage.type: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const group = state.conference.layouts.conference_layouts[parentId];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts: {
              ...state.conference.layouts.conference_layouts, [parentId]:
                {...group, images: {...state.conference.layouts.conference_layouts[parentId].images, [data.id]: data}}
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case StoreSwitchConferenceLayoutGroupLayout.type:
    case StoreUpdateConferenceLayoutGroupLayout.type: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const group = state.conference.layouts.conference_layouts_groups[parentId];
      if (!group) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts_groups: {
              ...state.conference.layouts.conference_layouts_groups, [parentId]:
                {
                  ...group,
                  layouts: {...state.conference.layouts.conference_layouts_groups[parentId].layouts, [data.id]: data}
                }
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreAddConferenceLayoutImage.type: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.layouts.conference_layouts[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.new || []];
      if (action.payload.payload.index !== undefined) {
        rest = [
          ...profile.new.slice(0, action.payload.payload.index),
          null,
          ...profile.new.slice(action.payload.payload.index + 1)
        ];
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts: {
              ...state.conference.layouts.conference_layouts, [parentId]:
                {...profile, images: {...profile.images, [data.id]: data}, new: rest}
            },
            errorMessage: action.payload.response.error || null,
          },
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        }
      };
    }
    case StoreAddConferenceLayoutGroupLayout.type: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.layouts.conference_layouts_groups[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.new || []];
      if (action.payload.payload.index !== undefined) {
        rest = [
          ...profile.new.slice(0, action.payload.payload.index),
          null,
          ...profile.new.slice(action.payload.payload.index + 1)
        ];
      }

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts_groups: {
              ...state.conference.layouts.conference_layouts_groups, [parentId]:
                {...profile, layouts: {...profile.layouts, [data.id]: data}, new: rest}
            },
            errorMessage: action.payload.response.error || null,
          },
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        }
      };
    }

    case StoreDelConferenceLayoutImage.type: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.layouts.conference_layouts[parentId];
      if (!profile || !profile.images[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.images;

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts: {
              ...state.conference.layouts.conference_layouts, [parentId]:
                {...profile, images: {...rest}}
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreDelConferenceLayoutGroupLayout.type: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.conference.layouts.conference_layouts_groups[parentId];
      if (!profile || !profile.layouts[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.layouts;

      return {
        ...state,
        conference: <Iconference>{
          ...state.conference, layouts: {
            ...state.conference.layouts, conference_layouts_groups: {
              ...state.conference.layouts.conference_layouts_groups, [parentId]:
                {...profile, layouts: {...rest}}
            }
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreAddConferenceLayout.type:
    case StoreUpdateConferenceLayout.type:
    case StoreUpdateConferenceLayout3D.type: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }
      if (!state.conference.layouts) {
        state.conference.layouts = <Ilayouts>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          layouts: {...state.conference.layouts, conference_layouts: {
            ...state.conference.layouts.conference_layouts,
            [data.id]: {...data, images: {...state.conference.layouts.conference_layouts[data.id]?.images}}}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreAddConferenceLayoutGroup.type:
    case StoreUpdateConferenceLayoutGroup.type: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.conference) {
        state.conference = <Iconference>{};
        state.loadCounter = 0;
      }
      if (!state.conference.layouts) {
        state.conference.layouts = <Ilayouts>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        conference: {
          ...state.conference,
          layouts: {...state.conference.layouts, conference_layouts_groups: {
            ...state.conference.layouts.conference_layouts_groups,
            [data.id]: {...data, layouts: {...state.conference.layouts.conference_layouts_groups[data.id]?.layouts}}}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreDelConferenceLayout.type: {
      const id = action.payload.response.data?.id || 0;
      if (!state.conference.layouts.conference_layouts[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.conference.layouts.conference_layouts;

      return {
        ...state,
        conference: {
          ...state.conference, layouts: {...state.conference.layouts, conference_layouts: {...rest}},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case StoreDelConferenceLayoutGroup.type: {
      const id = action.payload.response.data?.id || 0;
      if (!state.conference.layouts.conference_layouts_groups[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.conference.layouts.conference_layouts_groups;

      return {
        ...state,
        conference: {
          ...state.conference, layouts: {...state.conference.layouts, conference_layouts_groups: {...rest}},
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
