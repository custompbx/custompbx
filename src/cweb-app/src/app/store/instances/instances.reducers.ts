import {All, ConfigActionTypes} from './instances.actions';

export interface State {
  instances: object;
  currentInstanceId: number;
  loadCounter: number;
  errorMessage: string | null;
}

export const initialState: State = {
  instances: <Iinstances>{},
  currentInstanceId: 0,
  loadCounter: 0,
  errorMessage: '',
};

export interface Iinstances {
  [index: number]: {
    id: number,
    name: string,
    host: string,
    port: number,
    auth: string,
    token: string,
    description: string,
    enabled: boolean,
  };
}

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.UpdateInstanceDescription:
    case ConfigActionTypes.GetInstances: {
      return {...state, loadCounter: state.loadCounter + 1, errorMessage: null};
    }

    case ConfigActionTypes.StoreGotInstancesError: {
      return {
        ...state,
        errorMessage: action.payload.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetInstances: {
      return {
        ...state,
        instances: action.payload.response.fs_instances || {},
        currentInstanceId: action.payload.response.id || 0,
        errorMessage: action.payload.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreUpdateInstanceDescription: {
      return {
        ...state,
        instances: {...state.instances, ...action.payload.response.fs_instances},
        errorMessage: action.payload.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return state;
    }
  }
}
