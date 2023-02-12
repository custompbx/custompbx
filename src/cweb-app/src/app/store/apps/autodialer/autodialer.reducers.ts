import {
  AutoDialerActionTypes,
  All,
} from './autodialer.actions';
import {State, initialState, Iautodialer} from '../apps.state.struct';
import {IfilterField, IsortField} from '../../../components/cdr/cdr.component';
import {PageEvent} from '@angular/material/paginator';

export function reducer(state = initialState, action: All): State {
  switch (action.type) {

    case AutoDialerActionTypes.GetAutoDialerReducerMembers:
    case AutoDialerActionTypes.AddAutoDialerReducerMember:
    case AutoDialerActionTypes.DelAutoDialerReducerMember:
    case AutoDialerActionTypes.UpdateAutoDialerReducerMember:
    case AutoDialerActionTypes.GetAutoDialerCompanies:
    case AutoDialerActionTypes.AddAutoDialerCompany:
    case AutoDialerActionTypes.DelAutoDialerCompany:
    case AutoDialerActionTypes.UpdateAutoDialerCompany:
    case AutoDialerActionTypes.GetAutoDialerTeams:
    case AutoDialerActionTypes.AddAutoDialerTeam:
    case AutoDialerActionTypes.DelAutoDialerTeam:
    case AutoDialerActionTypes.UpdateAutoDialerTeam:
    case AutoDialerActionTypes.GetAutoDialerTeamMembers:
    case AutoDialerActionTypes.AddAutoDialerTeamMember:
    case AutoDialerActionTypes.DelAutoDialerTeamMember:
    case AutoDialerActionTypes.UpdateAutoDialerTeamMember:
    case AutoDialerActionTypes.GetAutoDialerLists:
    case AutoDialerActionTypes.AddAutoDialerList:
    case AutoDialerActionTypes.DelAutoDialerList:
    case AutoDialerActionTypes.UpdateAutoDialerList:
    case AutoDialerActionTypes.GetAutoDialerListMembers:
    case AutoDialerActionTypes.AddAutoDialerListMember:
    case AutoDialerActionTypes.AddAutoDialerListMembers:
    case AutoDialerActionTypes.DelAutoDialerListMember:
    case AutoDialerActionTypes.UpdateAutoDialerListMember:
    case AutoDialerActionTypes.GetAutoDialerReducers:
    case AutoDialerActionTypes.AddAutoDialerReducer:
    case AutoDialerActionTypes.DelAutoDialerReducer:
    case AutoDialerActionTypes.UpdateAutoDialerReducer: {
      return {
        ...state,
        errorMessage: null, loadCounter: state.loadCounter + 1
      };
    }

    case AutoDialerActionTypes.StoreAutoDialerError: {
      return {
        ...state,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.error,
      };
    }

    case AutoDialerActionTypes.StoreGetAutoDialerCompanies: {
      const data = action.payload.response['data'] || {};
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerCompanies: data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreAddAutoDialerCompany: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        autodialer: {
          ...state.autodialer,
          AutoDialerCompanies: {...state.autodialer.AutoDialerCompanies, ...data},
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    /*
    {
        "MessageType": "UpdateAutoDialerCompany",
        "data": {
            "id": 1,
            "position": 1,
            "name": "ddddd",
            "enabled": false,
            "predictive": false,
            "domain": {
                "id": 1,
                "position": 0,
                "enabled": false,
                "name": "",
                "parent": null,
                "sip_regs_counter": 0
            },
            "reducer": {
                "id": 0,
                "position": 0,
                "name": "",
                "enabled": false
            },
            "team": {
                "id": 1,
                "position": 0,
                "name": "",
                "enabled": false
            },
            "list": {
                "id": 3,
                "position": 0,
                "name": "",
                "enabled": false
            }
        }
    }
        */
    case AutoDialerActionTypes.StoreUpdateAutoDialerCompany: {
      const data = action.payload.response['data'] || {};
      const listId = data.list?.id || 0;
      if (!data.id) {
        return {
          ...state,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        };
      }
      const listMembers = state.autodialer.AutoDialerListMembers[listId] || <any>{table: [], tableMeta: newTablemeta(), list: {}, total: 0};

      return {
        ...state,
        autodialer: {
          ...state.autodialer,
          AutoDialerCompanies: {
            ...state.autodialer.AutoDialerCompanies,
            [data.id]: {...state.autodialer.AutoDialerCompanies[data.id], ...data}
          },
          AutoDialerListMembers: {...state.autodialer.AutoDialerListMembers, [listId]: listMembers},
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error || null,
      };
    }

    case AutoDialerActionTypes.StoreDelAutoDialerCompany: {
      const id = action.payload.response['id'] || 0;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerCompanies;

      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerCompanies: {...rest}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreUpdateAutoDialerTeam:
    case AutoDialerActionTypes.StoreAddAutoDialerTeam:
    case AutoDialerActionTypes.StoreGetAutoDialerTeams: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerTeams: {...state.autodialer.AutoDialerTeams, ...data}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreDelAutoDialerTeam: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerTeams;
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerTeams: {...rest}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreUpdateAutoDialerTeamMember:
    case AutoDialerActionTypes.StoreAddAutoDialerTeamMember:
    case AutoDialerActionTypes.StoreGetAutoDialerTeamMembers: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerTeamMembers: {...state.autodialer.AutoDialerTeamMembers, ...data}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreDelAutoDialerTeamMember: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerTeamMembers;
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerTeamMembers: {...rest}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreUpdateAutoDialerList:
    case AutoDialerActionTypes.StoreAddAutoDialerList: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerLists: {...state.autodialer.AutoDialerLists, ...data}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreGetAutoDialerLists: {
      const data = action.payload.response['data'] || {};
      const ids = Object.keys(data);

      const listMembers = {};
      ids.forEach(id => {
        listMembers[id] = state.autodialer.AutoDialerListMembers[id] || {table: [], tableMeta: newTablemeta(), list: {}, total: 0};
      });

      return {
        ...state,
        autodialer: {
          ...state.autodialer, AutoDialerLists: {...state.autodialer.AutoDialerLists, ...data},
          AutoDialerListMembers: <any>{
            ...state.autodialer.AutoDialerListMembers, ...listMembers,
          },
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreDelAutoDialerList: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerLists;
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerLists: {...rest}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreAddAutoDialerListMember: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerListMembers: {...state.autodialer.AutoDialerListMembers, ...data}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreUpdateAutoDialerListMember: {
      const data = action.payload.response['data'] || {};
      const parentId = data?.parent?.id || 0;
      if (!data || !data.id || !parentId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }

      const field = action.payload?.payload?.param?.name;
      const rest = state.autodialer.AutoDialerListMembers[parentId].table.map(item => {
        console.log(item);
        if (item['id'] && item['id'] === data.id) {
          Object.keys(data).forEach(k => {
            if (!state.autodialer.AutoDialerListMembers.changed[String(data.id)][k]) {
              return;
            }
            delete data[k];
          });
          return {...item, ...data};
        }
        return item;
      });
      return {
        ...state,
        autodialer: {
          ...state.autodialer, AutoDialerListMembers: {
            ...state.autodialer.AutoDialerListMembers,
            [parentId]: {
              ...state.autodialer.AutoDialerListMembers[parentId],
              table: [...rest],
            },
            changed: {
              ...state.autodialer.AutoDialerListMembers.changed || {},
              [data.id]: {
                ...state.autodialer.AutoDialerListMembers.changed[data.id] || {},
                [field]: false,
              }
            },
          }
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreDelAutoDialerListMember: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerListMembers;
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerListMembers: {...rest}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreUpdateAutoDialerReducer:
    case AutoDialerActionTypes.StoreAddAutoDialerReducer:
    case AutoDialerActionTypes.StoreGetAutoDialerReducers: {
      let data = action.payload.response['data'] || {};
      if (data.id) {
        data = {[data.id]: data};
      }
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerReducers: {...state.autodialer.AutoDialerReducers, ...data}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreDelAutoDialerReducer: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerReducers;
      return {
        ...state,
        autodialer: {...state.autodialer, AutoDialerReducers: {...rest}},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreUpdateAutoDialerReducerMember:
    case AutoDialerActionTypes.StoreAddAutoDialerReducerMember:
    case AutoDialerActionTypes.StoreGetAutoDialerReducerMembers: {
      let id = 0;
      let data = action.payload.response['data'] || {};
      const first = data[Object.keys(data)[0]];
      if (first && first.parent && first.parent.id !== 0) {
        id = first.parent.id;
      }
      if (data.parent && data.parent.id !== 0) {
        id = data.parent.id;
      }
      if (data.id) {
        data = {[data.id]: data};
      }
      if (id === 0) {
        return {
          ...state,
          loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
          // errorMessage: 'no id'
        };
      }
      return {
        ...state,
        autodialer: {
          ...state.autodialer, AutoDialerReducerMembers: {
            ...state.autodialer.AutoDialerReducerMembers,
            [id]: {...state.autodialer.AutoDialerReducerMembers[id], ...data},
          }
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }
////////////////
    case AutoDialerActionTypes.StoreDelAutoDialerReducerMember: {
      const id = action.payload.response.id;
      const {[id]: toDel, ...rest} = state.autodialer.AutoDialerReducerMembers;
      return {
        ...state,
        autodialer: {
          ...state.autodialer, AutoDialerReducerMembers: {
            ...state.autodialer.AutoDialerReducerMembers,
            [id]: {...rest},
          }
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.response.error,
      };
    }

    case AutoDialerActionTypes.StoreNewAutoDialerReducerMembers: {
      const id = action.payload.id;
      const members = state.autodialer.NewAutoDialerReducerMembers[id] || [];
      return {
        ...state,
        autodialer: {
          ...state.autodialer, NewAutoDialerReducerMembers: {
            ...state.autodialer.NewAutoDialerReducerMembers, [id]: [
              ...members,
              {},
            ]
          }
        },
      };
    }

    case AutoDialerActionTypes.StoreDropNewAutoDialerReducerMembers: {
      const id = action.payload.id;
      const members = state.autodialer.NewAutoDialerReducerMembers[id] || [];
      const rest = [
        ...members.slice(0, action.payload.index),
        null,
        ...members.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        autodialer: {
          ...state.autodialer, NewAutoDialerReducerMembers: {
            ...state.autodialer.NewAutoDialerReducerMembers, [id]: [
              ...rest
            ]
          }
        }
      };
    }

    case AutoDialerActionTypes.StoreNewAutoDialerTeamMembers: {
      const id = action.payload.id;
      const members = state.autodialer.NewAutoDialerTeamMembers[id] || [];
      return {
        ...state,
        autodialer: {
          ...state.autodialer, NewAutoDialerTeamMembers: {
            ...state.autodialer.NewAutoDialerTeamMembers, [id]: [
              ...members,
              {},
            ]
          }
        },
      };
    }

    case AutoDialerActionTypes.StoreDropNewAutoDialerTeamMembers: {
      const id = action.payload.id;
      const members = state.autodialer.NewAutoDialerTeamMembers[id] || [];
      const rest = [
        ...members.slice(0, action.payload.index),
        null,
        ...members.slice(action.payload.index + 1)
      ];

      return {
        ...state,
        autodialer: {
          ...state.autodialer, NewAutoDialerTeamMembers: {
            ...state.autodialer.NewAutoDialerTeamMembers, [id]: [
              ...rest
            ]
          }
        }
      };
    }

    case AutoDialerActionTypes.StoreGetAutoDialerListMembers: {
      const data = action.payload.response.data.items || [];
      let total = action.payload.response.data.total || 0;
      const parentId = (data.length > 0) ? data[0].parent.id : 0;

      const members = state.autodialer.AutoDialerListMembers[parentId] || {table: [], list: {}, total: 0, tableMeta: newTablemeta()};

      if (total < data.length) {
        total = data.length;
      }
      const columns = [];
      if (data && data.length > 0) {
        Object.keys(data[0]).forEach(key => {
          if (key === 'parent') {
            return;
          }
          columns.push(key);
        });
      }

      return {
        ...state,
        autodialer: <Iautodialer>{
          ...state.autodialer,
          AutoDialerListMembers: {
            ...state.autodialer.AutoDialerListMembers,
            [parentId]: {
              ...members,
              table: [...data],
              total: total,
              tableMeta: {
                ...members.tableMeta,
                columns: columns,
              },
            },
            // later put into tableMeta
            changed: {},
          },
          errorMessage: action.payload.response.error || null,
        },
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case AutoDialerActionTypes.StoreAddAutoDialerListMembers: {
      return {
        ...state,
        /*        autodialer: <Iautodialer>{
                  ...state.autodialer,
                  AutoDialerListMembers: {
                    ...state.autodialer.AutoDialerTeamMembers,
                    lastAdded: action.payload.total,
                  },
                },*/
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        errorMessage: action.payload.error || null,
      };
    }


    case AutoDialerActionTypes.StoreSetChangedAutodialerListMemberField: {
      const fieldName = action.payload.fieldName;
      const rowId = action.payload.rowId;
      if (!fieldName || !rowId) {
        return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
      }
      return {
        ...state,
        autodialer: {
          ...state.autodialer,
          AutoDialerListMembers: {
            ...state.autodialer.AutoDialerListMembers || {},
            changed: {
              ...state.autodialer.AutoDialerListMembers.changed || {},
              [rowId]: {
                ...state.autodialer.AutoDialerListMembers.changed[rowId] || {},
                [fieldName]: true,
              }
            }
          },
        },
      };
    }
    default: {
      return state;
    }
  }

}

function newTablemeta() {
  return {
    filters: [],
    pageEvent: {},
    sortObject: <IsortField>{fields: [], desc: false},
    filter: <IfilterField>{},
    columns: [],
    toEdit: {},
    showDel: {},
    toEditAgentFilter: '',
    sortColumns: '',
    csvData: '',
  };
}
