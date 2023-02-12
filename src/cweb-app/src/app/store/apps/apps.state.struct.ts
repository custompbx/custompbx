import {IfilterField, IsortField} from '../../components/cdr/cdr.component';
import {PageEvent} from '@angular/material/paginator';

export interface State {
  autodialer: Iautodialer;
  errorMessage: string;
  loadCounter: number;
}

export interface IlistMemberMeta {
  csvData: string;
  filters: Array<IfilterField>;
  pageEvent: PageEvent;
  sortObject: IsortField;
  filter: IfilterField;
  columns: Array<string>;
  toEditAgentFilter: number;
  sortColumns: string;
  toEdit: {};
  showDel: {};
}

export interface Iautodialer {
  AutoDialerCompanies: { [id: number]: object };
  AutoDialerTeams: { [id: number]: object };
  AutoDialerTeamMembers: { [id: number]: object };
  AutoDialerLists: { [id: number]: object };
  AutoDialerReducers: { [id: number]: object };
  AutoDialerReducerMembers: { [id: number]: object };
  AutoDialerListMembers: {
    [id: number]: {
      table: Array<object>;
      list: { [index: number]: object };
      total: number;

      tableMeta: IlistMemberMeta;
    }
    changed?: any;
    lastAdded?: number;
  };
  NewAutoDialerReducerMembers: { [id: number]: Array<object> };
  NewAutoDialerTeamMembers: { [id: number]: Array<object> };
}

export const autodialerInitialState: Iautodialer = {
  AutoDialerCompanies: {},
  AutoDialerTeams: {},
  AutoDialerTeamMembers: {},
  AutoDialerLists: {},
  AutoDialerListMembers: {},
  AutoDialerReducers: {},
  AutoDialerReducerMembers: {},
  NewAutoDialerReducerMembers: {},
  NewAutoDialerTeamMembers: {},
};

export const initialState: State = {
  autodialer: autodialerInitialState,
  errorMessage: '',
  loadCounter: 0,
};
