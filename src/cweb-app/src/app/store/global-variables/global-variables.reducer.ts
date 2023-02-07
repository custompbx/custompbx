import {All, ConfigActionTypes} from './global-variables.actions';

export interface State {
  globalVariables: IglobalVariables;
  newGlobalVariables: Array<IglobalVariable>;
  loadCounter: number;
  errorMessage: string | null;
}

export const initialState: State = {
  globalVariables: <IglobalVariables>{},
  newGlobalVariables: [],
  loadCounter: 0,
  errorMessage: '',
};

export interface IglobalVariables {
  [index: number]: IglobalVariable;
}

export interface IglobalVariable {
    id: number;
    name: string;
    value: string;
    type: string;
    enabled: boolean;
    position: number;
    dynamic: boolean;
}

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.MoveGlobalVariable:
    case ConfigActionTypes.ImportGlobalVariables:
    case ConfigActionTypes.GetGlobalVariables:
    case ConfigActionTypes.UpdateGlobalVariable:
    case ConfigActionTypes.SwitchGlobalVariable:
    case ConfigActionTypes.AddGlobalVariable:
    case ConfigActionTypes.DelGlobalVariable: {
      return {
        ...state,
        errorMessage: null,
        loadCounter: state.loadCounter + 1
      };
    }

    case ConfigActionTypes.StoreGotGlobalVariableError: {
      return {
        ...state,
        errorMessage: action.payload.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreMoveGlobalVariable:
    case ConfigActionTypes.StoreImportGlobalVariables:
    case ConfigActionTypes.StoreGetGlobalVariables: {
      return {
        ...state,
        globalVariables: action.payload.response.global_variables || {},
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelGlobalVariable: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.globalVariables;

      return {
        ...state,
        globalVariables: {
          ...rest,
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchGlobalVariable:
    case ConfigActionTypes.StoreUpdateGlobalVariable: {
      const data = action.payload.response.global_variables;

      return {
        ...state,
        globalVariables: {
          ...state.globalVariables, ...data,
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewGlobalVariable: {
      const rest = [
        ...state.newGlobalVariables || [],
        <IglobalVariable>{}
      ];

      return {
        ...state,
        newGlobalVariables: [
          ...rest
        ],
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewGlobalVariable: {
      const rest = [
        ...state.newGlobalVariables.slice(0, action.payload.index),
        null,
        ...state.newGlobalVariables.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        newGlobalVariables: [
          ...rest
        ],
        errorMessage: null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddGlobalVariable: {
      const data = action.payload.response.global_variables;
      let rest = [...state.newGlobalVariables || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.newGlobalVariables.slice(0, action.payload.index),
          null,
          ...state.newGlobalVariables.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        globalVariables: <IglobalVariables>{
          ...state.globalVariables,
          ...data,
        },
        newGlobalVariables: [
          ...rest
        ],
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return state;
    }
  }
}

