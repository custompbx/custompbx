import {Iitem, Ivoicemail, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.voicemail';
import {getParentId} from "../config.reducers";

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetVoicemailSettings:
    case ConfigActionTypes.GetVoicemailProfiles:
    case ConfigActionTypes.UpdateVoicemailSetting:
    case ConfigActionTypes.SwitchVoicemailSetting:
    case ConfigActionTypes.AddVoicemailSetting:
    case ConfigActionTypes.DelVoicemailSetting:
    case ConfigActionTypes.GetVoicemailProfileParameters:
    case ConfigActionTypes.AddVoicemailProfileParameter:
    case ConfigActionTypes.UpdateVoicemailProfileParameter:
    case ConfigActionTypes.SwitchVoicemailProfileParameter:
    case ConfigActionTypes.DelVoicemailProfileParameter:
    case ConfigActionTypes.AddVoicemailProfile:
    case ConfigActionTypes.DelVoicemailProfile:
    case ConfigActionTypes.UpdateVoicemailProfile: {
      return {...state,
        voicemail: {
          ...state.voicemail,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotVoicemailError: {
      return {
        ...state,
        voicemail: {
          ...state.voicemail,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetVoicemailProfiles: {
      const profiles = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          voicemail: {...state.voicemail, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.voicemail) {
        state.voicemail = <Ivoicemail>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        voicemail: {
          ...state.voicemail,
          profiles: {...state.voicemail.profiles, ...profiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddVoicemailProfile:
    case ConfigActionTypes.StoreUpdateVoicemailProfile: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.voicemail) {
        state.voicemail = <Ivoicemail>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        voicemail: {
          ...state.voicemail,
          profiles: {...state.voicemail.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetVoicemailSettings: {
      const settings = action.payload.response.data || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          voicemail: {...state.voicemail, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.voicemail) {
        state.voicemail = <Ivoicemail>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        voicemail: {
          ...state.voicemail,
          parameters: {...state.voicemail.parameters, ...settings},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelVoicemailSetting: {
      const id = action.payload.response.data?.id || 0;
      if (!state.voicemail.parameters[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.voicemail.parameters;

      return {
        ...state,
        voicemail: {
          ...state.voicemail, parameters: {...rest, new: state.voicemail.parameters.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchVoicemailSetting:
    case ConfigActionTypes.StoreUpdateVoicemailSetting: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        voicemail: <Ivoicemail>{
          ...state.voicemail, parameters: {...state.voicemail.parameters, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewVoicemailSetting: {
      const rest = [
        ...state.voicemail.parameters?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        voicemail: {
          ...state.voicemail, parameters: {...state.voicemail.parameters, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewVoicemailSetting: {
      const rest = [
        ...state.voicemail.parameters.new.slice(0, action.payload.index),
        null,
        ...state.voicemail.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        voicemail: {
          ...state.voicemail, parameters: {...state.voicemail.parameters, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddVoicemailSetting: {
      const data = action.payload.response.data || {};
      let rest = [...state.voicemail.parameters.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.voicemail.parameters.new.slice(0, action.payload.index),
          null,
          ...state.voicemail.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        voicemail: <Ivoicemail>{
          ...state.voicemail, parameters: {...state.voicemail.parameters, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddVoicemailProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const profile = state.voicemail.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.parameters.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.parameters.new.slice(0, action.payload.index),
          null,
          ...profile.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        voicemail: <Ivoicemail>{
          ...state.voicemail, profiles: {
            ...state.voicemail.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteVoicemailProfileParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.voicemail.profiles[fromId] || !state.voicemail.profiles[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.voicemail.profiles[toId].parameters ? state.voicemail.profiles[toId].parameters.new || [] : [];

      const newArray = Object.keys(state.voicemail.profiles[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.voicemail.profiles[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        voicemail: {
          ...state.voicemail,
          profiles: {
            ...state.voicemail.profiles,
            [toId]: {
              ...state.voicemail.profiles[toId],
              parameters: {
                ...state.voicemail.profiles[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelVoicemailProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const profile = state.voicemail.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.parameters;

      return {
        ...state,
        voicemail: <Ivoicemail>{
          ...state.voicemail, profiles: {
            ...state.voicemail.profiles, [parentId]:
              {...profile, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetVoicemailProfileParameters: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.voicemail.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        voicemail: <Ivoicemail>{
          ...state.voicemail, profiles: {
            ...state.voicemail.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchVoicemailProfileParameter:
    case ConfigActionTypes.StoreUpdateVoicemailProfileParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.voicemail.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        voicemail: <Ivoicemail>{
          ...state.voicemail, profiles: {
            ...state.voicemail.profiles, [parentId]:
              {...profile, parameters: {...profile.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewVoicemailProfileParameter: {
      const profile = state.voicemail.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        voicemail: {
          ...state.voicemail, profiles: {
            ...state.voicemail.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewVoicemailProfileParameter: {
      const profile = state.voicemail.profiles[action.payload.id];
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
        voicemail: {
          ...state.voicemail, profiles: {
            ...state.voicemail.profiles, [action.payload.id]:
              {...profile, parameters: {...profile.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelVoicemailProfile: {
      const data = action.payload.response.data || {};
      const profile = state.voicemail.profiles[data.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.voicemail.profiles;

      return {
        ...state,
        voicemail: {
          ...state.voicemail, profiles: {...rest},
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
