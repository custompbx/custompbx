import {Iitem, Iunicall, initialState, State} from '../config.state.struct';
import {All, ConfigActionTypes} from './config.actions.unicall';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case ConfigActionTypes.GetUnicall:
    case ConfigActionTypes.GetUnicallSpanParameters:
    case ConfigActionTypes.UpdateUnicallParameter:
    case ConfigActionTypes.SwitchUnicallParameter:
    case ConfigActionTypes.AddUnicallParameter:
    case ConfigActionTypes.DelUnicallParameter:
    case ConfigActionTypes.AddUnicallSpanParameter:
    case ConfigActionTypes.UpdateUnicallSpanParameter:
    case ConfigActionTypes.SwitchUnicallSpanParameter:
    case ConfigActionTypes.DelUnicallSpanParameter:
    case ConfigActionTypes.AddUnicallSpan:
    case ConfigActionTypes.DelUnicallSpan:
    case ConfigActionTypes.UpdateUnicallSpan: {
      return {...state,
        unicall: {
          ...state.unicall,
          errorMessage: null,
        }, loadCounter: state.loadCounter + 1};
    }

    case ConfigActionTypes.StoreGotUnicallError: {
      return {
        ...state,
        unicall: {
          ...state.unicall,
          errorMessage: action.payload.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    // case ConfigActionTypes.StoreSwitchUnicallSpan:
    case ConfigActionTypes.StoreGetUnicall: {
      const settings = action.payload.response.data['settings'] || {};
      const spans = action.payload.response.data['profiles'] || {};
      if (action.payload.response.exists === false) {
        return {
          ...state,
          unicall: {...state.unicall, exists: action.payload.response.exists},
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0
        };
      }

      if (!state.unicall) {
        state.unicall = <Iunicall>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        unicall: {
          ...state.unicall,
          settings: {...state.unicall.settings, ...settings},
          spans: {...state.unicall.spans, ...spans},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddUnicallSpan:
    case ConfigActionTypes.StoreUpdateUnicallSpan: {
      const data = action.payload.response.data || {};
      if (!data.id) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      if (!state.unicall) {
        state.unicall = <Iunicall>{};
        state.loadCounter = 0;
      }

      return {
        ...state,
        unicall: {
          ...state.unicall,
          spans: {...state.unicall.spans, [data.id]: data},
          exists: action.payload.response.exists,
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelUnicallParameter: {
      const id = action.payload.response.data?.id || 0;
      if (!state.unicall.settings[id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[id]: toDel, ...rest} = state.unicall.settings;

      return {
        ...state,
        unicall: {
          ...state.unicall, settings: {...rest, new: state.unicall.settings?.new || []},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreSwitchUnicallParameter:
    case ConfigActionTypes.StoreUpdateUnicallParameter: {
      const data = action.payload.response.data || {};

      return {
        ...state,
        unicall: <Iunicall>{
          ...state.unicall, settings: {...state.unicall.settings, [data.id]: data},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewUnicallParameter: {
      const rest = [
        ...state.unicall.settings?.new || [],
        <Iitem>{}
      ];

      return {
        ...state,
        unicall: {
          ...state.unicall, settings: {...state.unicall.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewUnicallParameter: {
      const rest = [
        ...state.unicall.settings.new.slice(0, action.payload.index),
        null,
        ...state.unicall.settings.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        unicall: {
          ...state.unicall, settings: {...state.unicall.settings, new: rest},
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddUnicallParameter: {
      const data = action.payload.response.data || {};
      let rest = [...state.unicall.settings?.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...state.unicall.settings.new.slice(0, action.payload.index),
          null,
          ...state.unicall.settings.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        unicall: <Iunicall>{
          ...state.unicall, settings: {...state.unicall.settings, [data.id]: data, new: rest},
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreAddUnicallSpanParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const span = state.unicall.spans[parentId];
      if (!span) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      let rest = [...span.parameters.new || []];

      if (action.payload.index !== undefined) {
        rest = [
          ...span.parameters.new.slice(0, action.payload.index),
          null,
          ...span.parameters.new.slice(action.payload.index + 1)
        ];
      }

      return {
        ...state,
        unicall: <Iunicall>{
          ...state.unicall, spans: {
            ...state.unicall.spans, [parentId]:
              {...span, parameters: {...span.parameters, [data.id]: data, new: rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StorePasteUnicallSpanParameters: {
      const fromId = action.payload.from_id;
      const toId = action.payload.to_id;
      if (
        !fromId || !toId || !state.unicall.spans[fromId] || !state.unicall.spans[toId]
      ) {
        return {
          ...state
        };
      }

      let new_items = state.unicall.spans[toId].parameters ? state.unicall.spans[toId].parameters?.new || [] : [];

      const newArray = Object.keys(state.unicall.spans[fromId].parameters).map(i => {
        if (i === 'new') {
          return;
        }
        return state.unicall.spans[fromId].parameters[i];
      });

      new_items = [...new_items, ...newArray];
      return {
        ...state,
        unicall: {
          ...state.unicall,
          spans: {
            ...state.unicall.spans,
            [toId]: {
              ...state.unicall.spans[toId],
              parameters: {
                ...state.unicall.spans[toId].parameters,
                new: [...new_items],
              },
            },
          }
        }
      };
    }

    case ConfigActionTypes.StoreDelUnicallSpanParameter: {
      const data = action.payload.response.data || {};
      const parentId = data?.parent?.id || 0;
      const span = state.unicall.spans[parentId];
      if (!span || !span.parameters[data.id]) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = span.parameters;

      return {
        ...state,
        unicall: <Iunicall>{
          ...state.unicall, spans: {
            ...state.unicall.spans, [parentId]:
              {...span, parameters: {...rest}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreGetUnicallSpanParameters: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const span = state.unicall.spans[parentId];
      if (!span) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        unicall: <Iunicall>{
          ...state.unicall, spans: {
            ...state.unicall.spans, [parentId]:
              {...span, parameters: {...span.parameters, ...data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }
    case ConfigActionTypes.StoreSwitchUnicallSpanParameter:
    case ConfigActionTypes.StoreUpdateUnicallSpanParameter: {
      const data = action.payload.response.data || {};
      const parentId = getParentId(data);
      if (parentId === 0) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const span = state.unicall.spans[parentId];
      if (!span) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      return {
        ...state,
        unicall: <Iunicall>{
          ...state.unicall, spans: {
            ...state.unicall.spans, [parentId]:
              {...span, parameters: {...span.parameters, [data.id]: data}}
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreNewUnicallSpanParameter: {
      const span = state.unicall.spans[action.payload.id];
      if (!span) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...span.parameters?.new || [],
        <Iitem>{}
      ];
      return {
        ...state,
        unicall: {
          ...state.unicall, spans: {
            ...state.unicall.spans, [action.payload.id]:
              {...span, parameters: {...span.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDropNewUnicallSpanParameter: {
      const span = state.unicall.spans[action.payload.id];
      if (!span) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      const rest = [
        ...span.parameters.new.slice(0, action.payload.index),
        null,
        ...span.parameters.new.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        unicall: {
          ...state.unicall, spans: {
            ...state.unicall.spans, [action.payload.id]:
              {...span, parameters: {...span.parameters, new: rest}}
          },
          errorMessage: null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreDelUnicallSpan: {
      const data = action.payload.response.data || {};
      const span = state.unicall.spans[data.id];
      if (!span) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const {[data.id]: toDel, ...rest} = state.unicall.spans;

      return {
        ...state,
        unicall: {
          ...state.unicall, spans: {...rest},
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

function getParentId (data): number {
  let id = 0;
  if (data.id) {
    id = data?.parent?.id || 0;
  } else {
    const ids = Object.keys(data);
    if (ids.length === 0) {
      return id;
    }
    id = data[ids[0]]?.parent?.id || 0;
  }
  return id;
}
