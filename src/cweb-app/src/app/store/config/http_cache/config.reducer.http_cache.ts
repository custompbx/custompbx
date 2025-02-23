import {Ihttpcache, Iitem, initialState, State} from '../config.state.struct';
import {
  All,
  ConfigActionTypes,
} from './config.actions.http_cache';
import {
  decreaseStateLoadField,
  increaseStateLoadField,
  removeFromObject,
  updateStateItem,
  updateStateWithFieldValue
} from '../config.reducers';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetHttpCache:
    case ConfigActionTypes.UpdateHttpCacheParameter:
    case ConfigActionTypes.SwitchHttpCacheParameter:
    case ConfigActionTypes.AddHttpCacheParameter:
    case ConfigActionTypes.UpdateHttpCacheProfileParam:
    case ConfigActionTypes.SwitchHttpCacheProfileParam:
    case ConfigActionTypes.DelHttpCacheProfileParam:
    case ConfigActionTypes.AddHttpCacheProfile:
    case ConfigActionTypes.DelHttpCacheProfile:
    case ConfigActionTypes.RenameHttpCacheProfile:
    case ConfigActionTypes.AddHttpCacheProfileParam:
    case ConfigActionTypes.GetHttpCacheProfileParameters:
    case ConfigActionTypes.AddHttpCacheProfileDomain:
    case ConfigActionTypes.DelHttpCacheProfileDomain:
    case ConfigActionTypes.SwitchHttpCacheProfileDomain:
    case ConfigActionTypes.UpdateHttpCacheProfileDomain:
    case ConfigActionTypes.UpdateHttpCacheProfileAws:
    case ConfigActionTypes.UpdateHttpCacheProfileAzure:
    case ConfigActionTypes.DelHttpCacheParameter: {
      return updateStateWithFieldValue(
        'http_cache',
        'errorMessage', null,
        increaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreGotHttpCacheError: {
      return updateStateWithFieldValue(
        'http_cache',
        'errorMessage', action.payload.error || null,
        decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreGetHttpCache: {
      const settings = action.payload.response.data.settings || {};
      const profiles = action.payload.response.data.profiles || {};
      if (action.payload.response.exists === false) {
        return updateStateWithFieldValue(
          'http_cache',
          'exists', action.payload.response.exists,
          decreaseStateLoadField(state));
      }

      if (!state.http_cache) {
        state.http_cache = <Ihttpcache>{};
        state.loadCounter = 0;
      }

      return updateStateItem(
        'http_cache',
        {
          settings: {...state.http_cache.settings, ...settings},
          profiles: {...state.http_cache.profiles, ...profiles},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        decreaseStateLoadField(state)
      );
    }

    case ConfigActionTypes.StoreDelHttpCacheParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.http_cache.settings[id]) {
        return decreaseStateLoadField(state);
      }

      return updateStateItem(
        'http_cache',
        {
          settings: {...removeFromObject(state.http_cache.settings, id), new: state.http_cache.settings?.new || []},
          errorMessage: action.payload.response.error || null
        },
        decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreSwitchHttpCacheParameter:
    case ConfigActionTypes.StoreUpdateHttpCacheParameter: {
      const data = action.payload.response.data;

      return updateStateItem(
        'http_cache', {
          settings: {...state.http_cache.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        }, decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreNewHttpCacheParameter: {
      const rest = [
        ...state.http_cache.settings?.new || [],
        <Iitem>{}
      ];

      return updateStateItem(
        'http_cache', {
          settings: {...state.http_cache.settings, new: rest},
          errorMessage: null,
        },
        decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreDropNewHttpCacheParameter: {
      const {settings} = state.http_cache;
      const {index} = action.payload;
      const newSettings = [...settings.new];

      newSettings[index] = null;

      return updateStateItem(
        'http_cache', {
          settings: {
            ...settings,
            new: newSettings
          },
          errorMessage: null
        },
        decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreAddHttpCacheParameter: {
      const {index} = action.payload;
      const {data} = action.payload.response;
      const {settings} = state.http_cache;
      let newSettings = [...(settings?.new || [])];

      // If index is specified, replace the element at that index with null
      if (index !== undefined) {
        newSettings = [
          ...newSettings.slice(0, index),
          null,
          ...newSettings.slice(index + 1)
        ];
      }

      return updateStateItem(
        'http_cache', {
          settings: {
            ...settings,
            [data.id]: data,
            new: newSettings
          },
          errorMessage: action.payload.error || null
        },
        decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreAddHttpCacheProfile:
    case ConfigActionTypes.StoreRenameHttpCacheProfile: {
      const data = action.payload.response.data;
      if (!data.id) {
        return decreaseStateLoadField(state);
      }

      if (!state.http_cache) {
        state.http_cache = <Ihttpcache>{};
        state.loadCounter = 0;
      }

      return updateStateItem(
        'http_cache',
        {
          settings: state.http_cache.settings,
          profiles: {...state.http_cache.profiles, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        decreaseStateLoadField(state)
      );
    }
    case ConfigActionTypes.StoreDelHttpCacheProfile: {
      const data = action.payload.response.data;
      const profile = state.http_cache.profiles[data.id];
      if (!profile) {
        return decreaseStateLoadField(state);
      }

      return updateStateItem(
        'http_cache',
        {
          settings: state.http_cache.settings,
          profiles: {...removeFromObject(state.http_cache.profiles, data.id)},
          errorMessage: action.payload.response.error || null
        },
        decreaseStateLoadField(state));
    }

    case ConfigActionTypes.StoreGetHttpCacheProfileParameters: {
      const domains = action.payload.response.data.domains || {};
      const azure = action.payload.response.data.azure || {};
      const aws = action.payload.response.data.aws_s3 || {};

      const parentId = getParentId(domains) ||  getParentId(azure) ||  getParentId(aws);
      if (parentId === 0) {
        return decreaseStateLoadField(state);
      }
      const profile = state.http_cache.profiles[parentId];
      if (!profile) {
        return decreaseStateLoadField(state);
      }

      return updateStateItem(
        'http_cache',
        {
          settings: state.http_cache.settings,
          profiles: {
            ...state.http_cache.profiles,
            [parentId]:
              {...profile,
                domains: {...domains},
                azure: {...azure},
                aws_s3: { ...aws}
              }
          },
          errorMessage: action.payload.response.error ||  '',
        },
      decreaseStateLoadField(state));

    }
    case ConfigActionTypes.StoreNewHttpCacheProfileDomain: {
      const profile = state.http_cache.profiles[action.payload.id];
      if (!profile) {
        return decreaseStateLoadField(state);
      }
      const rest = [
        ...profile.domains?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        http_cache: {
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [action.payload.id]:
              {...profile, domains: {...profile.domains, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewHttpCacheProfileDomain: {
      const profile = state.http_cache.profiles[action.payload.id];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...profile.domains.new.slice(0, action.payload.index),
        null,
        ...profile.domains.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        http_cache: {
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [action.payload.id]:
              {...profile, domains: {...profile.domains, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelHttpCacheProfileDomain: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.http_cache.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = profile.domains;

      return {
        ...state,
        http_cache: <Ihttpcache>{
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [parentId]:
              {...profile, domains: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchHttpCacheProfileDomain:
    case ConfigActionTypes.StoreUpdateHttpCacheProfileDomain: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.http_cache.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }


      return {
        ...state,
        http_cache: <Ihttpcache>{
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [parentId]:
              {...profile, domains: {...profile.domains, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddHttpCacheProfileDomain: {
      const data = action.payload.response.data;
      const parentId = data?.parent?.id || 0;
      const profile = state.http_cache.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...profile.domains.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...profile.domains.new.slice(0, action.payload.index),
          null,
          ...profile.domains.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        http_cache: <Ihttpcache>{
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [parentId]:
              {...profile, domains: {...profile.domains, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateHttpCacheProfileAws: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.http_cache.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }


      return {
        ...state,
        http_cache: <Ihttpcache>{
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [parentId]:
              {...profile, aws_s3: {...profile.aws_s3, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateHttpCacheProfileAzure: {
      const data = action.payload.response.data;
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const profile = state.http_cache.profiles[parentId];
      if (!profile) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }


      return {
        ...state,
        http_cache: <Ihttpcache>{
          ...state.http_cache, profiles: {
            ...state.http_cache.profiles, [parentId]:
              {...profile, azure: {...profile.azure, [data.id]: data}}
          },
          errorMessage: action.payload.response.error ||  '',
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return null;
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
