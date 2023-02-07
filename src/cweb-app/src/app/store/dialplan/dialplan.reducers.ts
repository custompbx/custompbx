import {DialplanActionTypes, All} from './dialplan.actions';

export interface State {
  contexts: Icontexts;
  debug: Idebug;
  staticDialplan: boolean;
  loadCounter: number;
  errorMessage: string | null;
}

export interface Idebug {
  log: Array<{log: Array<string>, actions: {application: string, data: string, inline: boolean}}>;
  enabled: boolean;
}

export interface Icontexts {
  [index: number]: Icontext;
}

export interface Icontext {
  // extensions: Iextensions;
  extensions: Array<Iextension>;
  id?: number;
  enabled?: boolean;
  name: string;
}

export interface Iextensions {
  [index: number]: Iextension;
  new: Array<object>;
}

export interface Iextension {
  id: number;
  position: number;
  name: string;
  continue: string;
  // conditions: Iconditions;
  conditions: Array<Icondition>;
  enabled?: boolean;
}

export interface Iconditions {
  [index: number]: Icondition;
  new: Array<object>;
}

export interface Icondition {
  id: number;
  position: number;
  enabled: boolean;
  regexes: Array<Iregex>;
  actions: Array<Iaction>;
  antiactions: Array<Iantiaction>;
  newRegexes: Array<Iregex>;
  newActions: Array<Iaction>;
  newAntiactions: Array<Iantiaction>;
  new: Array<object>;
}

export interface Iactions {
  [index: number]: Iaction;
  new: Array<object>;
}

export interface Iaction {
  id: number;
  position: number;
  application: string;
  data: string;
  inline: string;
  enabled: boolean;
}

export interface Iregex {
  id: number;
  position: number;
  field: string;
  expression: string;
  enabled: boolean;
}

export interface Iantiactions {
  [index: number]: Iantiaction;
  new: Array<object>;
}

export interface Iantiaction {
  id: number;
  position: number;
  application: string;
  data: string;
  enabled: boolean;
}

export const initialState: State = {
   contexts: <Icontexts>null,
   debug: <Idebug>{},
   staticDialplan: false,
   loadCounter: 0,
  errorMessage: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case DialplanActionTypes.UPDATE_FAILURE: {
      return {
        ...state,
        loadCounter: state.loadCounter >= 0 ? --state.loadCounter : 0,
        // errorMessage: 'Cant get data from server',
      };
    }

    case DialplanActionTypes.REDUCE_LOAD_COUNTER: {
      return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
    }

    case DialplanActionTypes.SWITCH_DIALPLAN_STATIC:
    case DialplanActionTypes.DIALPLAN_SETTINGS:
    case DialplanActionTypes.DIALPLAN_DEBUG:
    case DialplanActionTypes.SWITCH_DIALPLAN_DEBUG:
    case DialplanActionTypes.IMPORT_DIALPLAN:
    case DialplanActionTypes.UPDATE_CONDITION:
    case DialplanActionTypes.SWITCH_CONDITION:
    case DialplanActionTypes.DELETE_CONDITION:
    case DialplanActionTypes.SWITCH_EXTENSION_CONTINUE:
    case DialplanActionTypes.DELETE_CONTEXT:
    case DialplanActionTypes.DELETE_EXTENSION:
    case DialplanActionTypes.RENAME_CONTEXT:
    case DialplanActionTypes.RENAME_EXTENSION:
    case DialplanActionTypes.ADD_CONTEXT:
    case DialplanActionTypes.ADD_EXTENSION:
    case DialplanActionTypes.ADD_CONDITION:
    case DialplanActionTypes.SWITCH_REGEX:
    case DialplanActionTypes.SWITCH_ACTION:
    case DialplanActionTypes.SWITCH_ANTIACTION:
    case DialplanActionTypes.UPDATE_REGEX:
    case DialplanActionTypes.UPDATE_ACTION:
    case DialplanActionTypes.UPDATE_ANTIACTION:
    case DialplanActionTypes.DELETE_REGEX:
    case DialplanActionTypes.DELETE_ACTION:
    case DialplanActionTypes.DELETE_ANTIACTION:
    case DialplanActionTypes.ADD_REGEX:
    case DialplanActionTypes.ADD_ACTION:
    case DialplanActionTypes.ADD_ANTIACTION:
    case DialplanActionTypes.GET_CONDITIONS:
    case DialplanActionTypes.GET_CONTEXTS:
    case DialplanActionTypes.GET_EXTENSION_DETAILS:
    case DialplanActionTypes.MOVE_EXTENSION:
    case DialplanActionTypes.MOVE_CONDITION:
    case DialplanActionTypes.MOVE_ACTION:
    case DialplanActionTypes.MOVE_ANTIACTION:
    case DialplanActionTypes.GET_EXTENSIONS: {
      return {...state,
        errorMessage: null, loadCounter: state.loadCounter + 1};
    }

    case DialplanActionTypes.STORE_GET_CONTEXTS: {
      const data = action.payload.response['dialplan_contexts'];
      return {
        ...state,
        contexts: { ...data, },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_RENAME_CONTEXT:
    case DialplanActionTypes.STORE_ADD_CONTEXT: {
      const data = action.payload.response['dialplan_contexts'];
      return {
        ...state,
        contexts: {
          ...state.contexts, ...data,
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_DELETE_CONTEXT: {
      const id = action.payload.response.id;
      if (!state.contexts[id]) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      const {[id]: toDel, ...rest} = state.contexts;

      return {
        ...state,
        contexts: {
          ...rest,
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_MOVE_EXTENSION:
    case DialplanActionTypes.STORE_GET_EXTENSIONS: {
      const id = action.payload.response.id;
      const data = action.payload.response['dialplan_extensions'];
      if (!state.contexts[id] || !Array.isArray(data)) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      data.sort(function (a, b) {
        if (a.position > b.position) {
          return 1;
        }
        if (a.position < b.position) {
          return -1;
        }
        return 0;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [id]: {...state.contexts[id], extensions: data},
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_ADD_EXTENSION: {
      const id = action.payload.response.id;
      const data = action.payload.response['dialplan_extensions'];
      if (!state.contexts[id] || !Array.isArray(data)) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const extensions: Array<Iextension> = Array.isArray(state.contexts[id].extensions) ? state.contexts[id].extensions : [];

      return {
        ...state,
        contexts: {
          ...state.contexts, [id]: {...state.contexts[id], extensions: [...extensions, ...data]},
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_SWITCH_EXTENSION_CONTINUE:
    case DialplanActionTypes.STORE_RENAME_EXTENSION: {
      const id = action.payload.response.id;
      const data = action.payload.response['dialplan_extensions'];
      if (!state.contexts[id] || !Array.isArray(data) || data.length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      return {
        ...state,
        contexts: {
          ...state.contexts, [id]: {...state.contexts[id], extensions: state.contexts[id].extensions.map(exten => {
              if (exten.id === data[0].id) {
                exten = {...exten, ...data[0]};
              }
              return exten;
            })
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_DELETE_EXTENSION: {
      const id = action.payload.response.id;
      const extenId = action.payload.response['affected_id'];
      if (!state.contexts[id]) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      return {
        ...state,
        contexts: {
          ...state.contexts, [id]: {...state.contexts[id], extensions: state.contexts[id].extensions.filter(exten =>
            exten.id !== extenId
            )
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_MOVE_CONDITION:
    case DialplanActionTypes.STORE_GET_CONDITIONS: {
      const contextId = action.payload.response.id;
      const data = action.payload.response['dialplan_conditions'];
      if (!data) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const id = ids[0];

      if (!state.contexts[contextId] || !Array.isArray(state.contexts[contextId].extensions) || !Array.isArray(data[id])) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      data[id].sort(function (a, b) {
        if (a.position > b.position) {
          return 1;
        }
        if (a.position < b.position) {
          return -1;
        }
        return 0;
      });

      const extensions = <Array<Iextension>>state.contexts[contextId].extensions.map(exten => {
        if (exten.id === Number(id)) {
          exten.conditions = data[id];
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts,
          [contextId]: {
            ...state.contexts[contextId],
            extensions: extensions
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_ADD_CONDITION: {
      const contextId = action.payload.response.id;
      const data = action.payload.response['dialplan_conditions'];
      if (!data) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const id = ids[0];

      if (!state.contexts[contextId] || !Array.isArray(state.contexts[contextId].extensions) || !Array.isArray(data[id])) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      const extensions = <Array<Iextension>>state.contexts[contextId].extensions.map(exten => {
        if (exten.id === Number(id)) {
          if (!exten.conditions) {
            exten.conditions = [];
          }
          exten.conditions = [...exten.conditions, ...<Array<Icondition>>data[id]];
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts,
          [contextId]: {
            ...state.contexts[contextId],
            extensions: extensions
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_UPDATE_CONDITION:
    case DialplanActionTypes.STORE_SWITCH_CONDITION: {
      const contextId = action.payload.response.id;
      const data = action.payload.response['dialplan_conditions'];
      if (!data) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const id = ids[0];
      if (!state.contexts[contextId] || !Array.isArray(state.contexts[contextId].extensions) || !Array.isArray(data[id]) || data[id].length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      const extensions = <Array<Iextension>>state.contexts[contextId].extensions.map(exten => {
        if (exten.id === Number(id)) {
          exten.conditions = exten.conditions.map(cond => {if (cond.id === data[id][0].id) {
            cond = {...cond, ...data[id][0]};
          }
          return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts,
          [contextId]: {
            ...state.contexts[contextId],
            extensions: extensions
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_DELETE_CONDITION: {
      const contextId = action.payload.response.id;
      const data = action.payload.response['dialplan_conditions'];
      if (!data) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const id = ids[0];
      if (!state.contexts[contextId] || !Array.isArray(state.contexts[contextId].extensions) || !Array.isArray(data[id])) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const extensions = <Array<Iextension>>state.contexts[contextId].extensions.map(exten => {
        if (exten.id === Number(id)) {
          exten.conditions = exten.conditions.filter(cond => cond.id !== data[id][0].id);
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts,
          [contextId]: {
            ...state.contexts[contextId],
            extensions: extensions
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.STORE_SWITCH_ACTION:
    case DialplanActionTypes.STORE_SWITCH_ANTIACTION:
    case DialplanActionTypes.STORE_SWITCH_REGEX:
    case DialplanActionTypes.STORE_UPDATE_ACTION:
    case DialplanActionTypes.STORE_UPDATE_ANTIACTION:
    case DialplanActionTypes.STORE_UPDATE_REGEX:
    case DialplanActionTypes.STORE_DELETE_ACTION:
    case DialplanActionTypes.STORE_DELETE_ANTIACTION:
    case DialplanActionTypes.STORE_DELETE_REGEX:
    case DialplanActionTypes.STORE_ADD_ACTION:
    case DialplanActionTypes.STORE_ADD_ANTIACTION:
    case DialplanActionTypes.STORE_ADD_REGEX:
    case DialplanActionTypes.STORE_MOVE_ACTION:
    case DialplanActionTypes.STORE_MOVE_ANTIACTION:
    case DialplanActionTypes.STORE_EXTENSIONS_DETAILS: {
      if (action.payload.response.error) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error,
        };
      }
      const contextId = action.payload.response.id;
      const extenId = action.payload.response['affected_id'];
      const data = action.payload.response['dialplan_details'];
      if (!data) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const ids = Object.keys(data);
      if (ids.length === 0) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }
      const id = ids[0];
      if (
        !state.contexts[contextId] &&
        !Array.isArray(state.contexts[contextId].extensions)
        // !extensions[0].conditions[id]
      ) {
        return {
          ...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          errorMessage: action.payload.response.error || null,
        };
      }

      if (data[id].actions) {
        data[id].actions.sort(function (a, b) {
          if (a.position > b.position) {
            return 1;
          }
          if (a.position < b.position) {
            return -1;
          }
          return 0;
        });
      }

      if (data[id].antiactions) {
        data[id].antiactions.sort(function (a, b) {
          if (a.position > b.position) {
            return 1;
          }
          if (a.position < b.position) {
            return -1;
          }
          return 0;
        });
      }

      const extensions = <Array<Iextension>>state.contexts[contextId].extensions.map(exten => {
        if (exten.id === Number(extenId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(id)) {
              cond.regexes = data[id].regexes ? data[id].regexes : cond.regexes;
              cond.actions = data[id].actions ? data[id].actions : cond.actions;
              cond.antiactions = data[id].antiactions ? data[id].antiactions : cond.antiactions;
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts,
          [contextId]: {
            ...state.contexts[contextId],
            extensions: extensions
          },
        },
        errorMessage: action.payload.response.error || null,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case DialplanActionTypes.ADD_NEW_REGEX: {
      const extensions = <Array<Iextension>>state.contexts[action.payload.contextId].extensions.map(exten => {
        if (exten.id === Number(action.payload.extensionId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(action.payload.conditionId)) {
              cond.newRegexes = [
                ...cond.newRegexes || [],
                <Iregex>{}
              ];
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [action.payload.contextId]: {...state.contexts[action.payload.contextId], extensions: extensions},
        },
      };
    }

    case DialplanActionTypes.DELETE_NEW_REGEX: {
      const extensions = <Array<Iextension>>state.contexts[action.payload.contextId].extensions.map(exten => {
        if (exten.id === Number(action.payload.extensionId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(action.payload.conditionId)) {
              cond.newRegexes = [
                ...cond.newRegexes.slice(0, action.payload.index),
                null,
                ...cond.newRegexes.slice(action.payload.index + 1)
              ];
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [action.payload.contextId]: {...state.contexts[action.payload.contextId], extensions: extensions},
        },
      };
    }

    case DialplanActionTypes.ADD_NEW_ACTION: {
      const extensions = <Array<Iextension>>state.contexts[action.payload.contextId].extensions.map(exten => {
        if (exten.id === Number(action.payload.extensionId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(action.payload.conditionId)) {
              cond.newActions = [
                ...cond.newActions || [],
                <Iaction>{}
              ];
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [action.payload.contextId]: {...state.contexts[action.payload.contextId], extensions: extensions},
        },
      };
    }

    case DialplanActionTypes.DELETE_NEW_ACTION: {
      const extensions = <Array<Iextension>>state.contexts[action.payload.contextId].extensions.map(exten => {
        if (exten.id === Number(action.payload.extensionId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(action.payload.conditionId)) {
              cond.newActions = [
                ...cond.newActions.slice(0, action.payload.index),
                null,
                ...cond.newActions.slice(action.payload.index + 1)
              ];
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [action.payload.contextId]: {...state.contexts[action.payload.contextId], extensions: extensions},
        },
      };
    }

    case DialplanActionTypes.ADD_NEW_ANTIACTION: {
      const extensions = <Array<Iextension>>state.contexts[action.payload.contextId].extensions.map(exten => {
        if (exten.id === Number(action.payload.extensionId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(action.payload.conditionId)) {
              cond.newAntiactions = [
                ...cond.newAntiactions || [],
                <Iantiaction>{}
              ];
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [action.payload.contextId]: {...state.contexts[action.payload.contextId], extensions: extensions},
        },
      };
    }

    case DialplanActionTypes.DELETE_NEW_ANTIACTION: {
      const extensions = <Array<Iextension>>state.contexts[action.payload.contextId].extensions.map(exten => {
        if (exten.id === Number(action.payload.extensionId)) {
          exten.conditions = <Array<Icondition>>exten.conditions.map(cond => {
            if (cond.id === Number(action.payload.conditionId)) {
              cond.newAntiactions = [
                ...cond.newAntiactions.slice(0, action.payload.index),
                null,
                ...cond.newAntiactions.slice(action.payload.index + 1)
              ];
            }
            return cond;
          });
        }
        return exten;
      });

      return {
        ...state,
        contexts: {
          ...state.contexts, [action.payload.contextId]: {...state.contexts[action.payload.contextId], extensions: extensions},
        },
      };
    }

    case DialplanActionTypes.STORE_SWITCH_DIALPLAN_DEBUG:
    case DialplanActionTypes.STORE_DIALPLAN_DEBUG: {
      const enabled = action.payload.response.enabled;
      const log = action.payload.response['dialplan_debug'];
      const data = state.debug;
      if (typeof enabled === 'boolean') {
        data.enabled = enabled;
      }
      if (log) {
        if (!data.log) {
          data.log = [];
        }
        data.log.push(log);
      }

      return {
        ...state,
        debug: {...data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error || null,
      };
    }

    case DialplanActionTypes.STORE_SWITCH_DIALPLAN_STATIC:
    case DialplanActionTypes.STORE_DIALPLAN_SETTINGS: {
      const dialplanSettings = action.payload.response.dialplan_settings;
      let noProceed = state.staticDialplan;
      if (dialplanSettings && typeof dialplanSettings.no_proceed === 'boolean') {
        noProceed = dialplanSettings.no_proceed;
      }

      return {
        ...state,
        staticDialplan: noProceed,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error || null,
      };
    }

    case DialplanActionTypes.STORE_CLEAR_DIALPLAN_DEBUG: {
      return {
        ...state,
        debug: {log: [], enabled: state.debug.enabled},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: null,
      };
    }

    default: {
      return state;
    }
  }
}
