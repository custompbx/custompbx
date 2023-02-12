import {Action} from '@ngrx/store';

export enum AutoDialerActionTypes {
  StoreAutoDialerError = 'StoreAutoDialerError',
  GetAutoDialerCompanies = 'GetAutoDialerCompanies',
  StoreGetAutoDialerCompanies = 'StoreGetAutoDialerCompanies',
  AddAutoDialerCompany = 'AddAutoDialerCompany',
  StoreAddAutoDialerCompany = 'StoreAddAutoDialerCompany',
  DelAutoDialerCompany = 'DelAutoDialerCompany',
  StoreDelAutoDialerCompany = 'StoreDelAutoDialerCompany',
  UpdateAutoDialerCompany = 'UpdateAutoDialerCompany',
  StoreUpdateAutoDialerCompany = 'StoreUpdateAutoDialerCompany',
  GetAutoDialerTeams = 'GetAutoDialerTeams',
  StoreGetAutoDialerTeams = 'StoreGetAutoDialerTeams',
  AddAutoDialerTeam = 'AddAutoDialerTeam',
  StoreAddAutoDialerTeam = 'StoreAddAutoDialerTeam',
  DelAutoDialerTeam = 'DelAutoDialerTeam',
  StoreDelAutoDialerTeam = 'StoreDelAutoDialerTeam',
  UpdateAutoDialerTeam = 'UpdateAutoDialerTeam',
  StoreUpdateAutoDialerTeam = 'StoreUpdateAutoDialerTeam',
  GetAutoDialerTeamMembers = 'GetAutoDialerTeamMembers',
  StoreGetAutoDialerTeamMembers = 'StoreGetAutoDialerTeamMembers',
  AddAutoDialerTeamMember = 'AddAutoDialerTeamMember',
  StoreAddAutoDialerTeamMember = 'StoreAddAutoDialerTeamMember',
  AddAutoDialerTeamMembers = 'AddAutoDialerTeamMembers',
  StoreAddAutoDialerTeamMembers = 'StoreAddAutoDialerTeamMembers',
  DelAutoDialerTeamMember = 'DelAutoDialerTeamMember',
  StoreDelAutoDialerTeamMember = 'StoreDelAutoDialerTeamMember',
  UpdateAutoDialerTeamMember = 'UpdateAutoDialerTeamMember',
  StoreUpdateAutoDialerTeamMember = 'StoreUpdateAutoDialerTeamMember',
  GetAutoDialerLists = 'GetAutoDialerLists',
  StoreGetAutoDialerLists = 'StoreGetAutoDialerLists',
  AddAutoDialerList = 'AddAutoDialerList',
  StoreAddAutoDialerList = 'StoreAddAutoDialerList',
  DelAutoDialerList = 'DelAutoDialerList',
  StoreDelAutoDialerList = 'StoreDelAutoDialerList',
  UpdateAutoDialerList = 'UpdateAutoDialerList',
  StoreUpdateAutoDialerList = 'StoreUpdateAutoDialerList',
  GetAutoDialerListMembers = 'GetAutoDialerListMembers',
  StoreGetAutoDialerListMembers = 'StoreGetAutoDialerListMembers',
  AddAutoDialerListMember = 'AddAutoDialerListMember',
  StoreAddAutoDialerListMember = 'StoreAddAutoDialerListMember',
  AddAutoDialerListMembers = 'AddAutoDialerListMembers',
  StoreAddAutoDialerListMembers = 'StoreAddAutoDialerListMembers',
  DelAutoDialerListMember = 'DelAutoDialerListMember',
  StoreDelAutoDialerListMember = 'StoreDelAutoDialerListMember',
  UpdateAutoDialerListMember = 'UpdateAutoDialerListMember',
  StoreUpdateAutoDialerListMember = 'StoreUpdateAutoDialerListMember',
  GetAutoDialerReducers = 'GetAutoDialerReducers',
  StoreGetAutoDialerReducers = 'StoreGetAutoDialerReducers',
  AddAutoDialerReducer = 'AddAutoDialerReducer',
  StoreAddAutoDialerReducer = 'StoreAddAutoDialerReducer',
  DelAutoDialerReducer = 'DelAutoDialerReducer',
  StoreDelAutoDialerReducer = 'StoreDelAutoDialerReducer',
  UpdateAutoDialerReducer = 'UpdateAutoDialerReducer',
  StoreUpdateAutoDialerReducer = 'StoreUpdateAutoDialerReducer',
  GetAutoDialerReducerMembers = 'GetAutoDialerReducerMembers',
  StoreGetAutoDialerReducerMembers = 'StoreGetAutoDialerReducerMembers',
  AddAutoDialerReducerMember = 'AddAutoDialerReducerMember',
  StoreAddAutoDialerReducerMember = 'StoreAddAutoDialerReducerMember',
  DelAutoDialerReducerMember = 'DelAutoDialerReducerMember',
  StoreDelAutoDialerReducerMember = 'StoreDelAutoDialerReducerMember',
  UpdateAutoDialerReducerMember = 'UpdateAutoDialerReducerMember',
  StoreUpdateAutoDialerReducerMember = 'StoreUpdateAutoDialerReducerMember',
  StoreNewAutoDialerReducerMembers = 'StoreNewAutoDialerReducerMembers',
  StoreDropNewAutoDialerReducerMembers = 'StoreDropNewAutoDialerReducerMembers',
  StoreNewAutoDialerTeamMembers = 'StoreNewAutoDialerTeamMembers',
  StoreDropNewAutoDialerTeamMembers = 'StoreDropNewAutoDialerTeamMembers',
  StoreSetChangedAutodialerListMemberField = 'StoreSetChangedAutodialerListMemberField',
}


export class StoreAutoDialerError implements Action {
  readonly type = AutoDialerActionTypes.StoreAutoDialerError;
  constructor(public payload: any) {
  }
}

export class StoreSetChangedAutodialerListMemberField implements Action {
  readonly type = AutoDialerActionTypes.StoreSetChangedAutodialerListMemberField;
  constructor(public payload: any) {
  }
}

export class GetAutoDialerCompanies implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerCompanies;
  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerCompanies implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerCompanies;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerCompany implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerCompany;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerCompany implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerCompany;

  constructor(public payload: any) {
  }
}

export class DelAutoDialerCompany implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerCompany;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerCompany implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerCompany;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerCompany implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerCompany;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerCompany implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerCompany;

  constructor(public payload: any) {
  }
}

export class GetAutoDialerTeams implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerTeams;

  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerTeams implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerTeams;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerTeam implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerTeam;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerTeam implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerTeam;

  constructor(public payload: any) {
  }
}

export class DelAutoDialerTeam implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerTeam;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerTeam implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerTeam;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerTeam implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerTeam;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerTeam implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerTeam;

  constructor(public payload: any) {
  }
}

export class GetAutoDialerTeamMembers implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerTeamMembers;

  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerTeamMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerTeamMembers;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerTeamMember implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerTeamMember;
  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerTeamMember implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerTeamMember;
  constructor(public payload: any) {
  }
}

export class AddAutoDialerTeamMembers implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerTeamMembers;
  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerTeamMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerTeamMembers;
  constructor(public payload: any) {
  }
}

export class DelAutoDialerTeamMember implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerTeamMember;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerTeamMember implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerTeamMember;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerTeamMember implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerTeamMember;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerTeamMember implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerTeamMember;

  constructor(public payload: any) {
  }
}

export class GetAutoDialerLists implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerLists;

  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerLists implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerLists;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerList implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerList;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerList implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerList;

  constructor(public payload: any) {
  }
}

export class DelAutoDialerList implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerList;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerList implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerList;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerList implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerList;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerList implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerList;

  constructor(public payload: any) {
  }
}

export class GetAutoDialerListMembers implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerListMembers;

  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerListMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerListMembers;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerListMember implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerListMember;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerListMember implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerListMember;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerListMembers implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerListMembers;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerListMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerListMembers;

  constructor(public payload: any) {
  }
}

export class DelAutoDialerListMember implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerListMember;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerListMember implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerListMember;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerListMember implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerListMember;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerListMember implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerListMember;

  constructor(public payload: any) {
  }
}

export class GetAutoDialerReducers implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerReducers;

  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerReducers implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerReducers;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerReducer implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerReducer;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerReducer implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerReducer;

  constructor(public payload: any) {
  }
}

export class DelAutoDialerReducer implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerReducer;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerReducer implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerReducer;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerReducer implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerReducer;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerReducer implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerReducer;

  constructor(public payload: any) {
  }
}

export class GetAutoDialerReducerMembers implements Action {
  readonly type = AutoDialerActionTypes.GetAutoDialerReducerMembers;

  constructor(public payload: any) {
  }
}

export class StoreGetAutoDialerReducerMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreGetAutoDialerReducerMembers;

  constructor(public payload: any) {
  }
}

export class AddAutoDialerReducerMember implements Action {
  readonly type = AutoDialerActionTypes.AddAutoDialerReducerMember;

  constructor(public payload: any) {
  }
}

export class StoreAddAutoDialerReducerMember implements Action {
  readonly type = AutoDialerActionTypes.StoreAddAutoDialerReducerMember;

  constructor(public payload: any) {
  }
}

export class DelAutoDialerReducerMember implements Action {
  readonly type = AutoDialerActionTypes.DelAutoDialerReducerMember;

  constructor(public payload: any) {
  }
}

export class StoreDelAutoDialerReducerMember implements Action {
  readonly type = AutoDialerActionTypes.StoreDelAutoDialerReducerMember;

  constructor(public payload: any) {
  }
}

export class UpdateAutoDialerReducerMember implements Action {
  readonly type = AutoDialerActionTypes.UpdateAutoDialerReducerMember;

  constructor(public payload: any) {
  }
}

export class StoreUpdateAutoDialerReducerMember implements Action {
  readonly type = AutoDialerActionTypes.StoreUpdateAutoDialerReducerMember;

  constructor(public payload: any) {
  }
}

export class StoreNewAutoDialerReducerMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreNewAutoDialerReducerMembers;

  constructor(public payload: any) {
  }
}

export class StoreDropNewAutoDialerReducerMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreDropNewAutoDialerReducerMembers;

  constructor(public payload: any) {
  }
}

export class StoreNewAutoDialerTeamMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreNewAutoDialerTeamMembers;

  constructor(public payload: any) {
  }
}

export class StoreDropNewAutoDialerTeamMembers implements Action {
  readonly type = AutoDialerActionTypes.StoreDropNewAutoDialerTeamMembers;

  constructor(public payload: any) {
  }
}

export type All =
  | StoreAutoDialerError
  | GetAutoDialerCompanies
  | StoreGetAutoDialerCompanies
  | AddAutoDialerCompany
  | StoreAddAutoDialerCompany
  | DelAutoDialerCompany
  | StoreDelAutoDialerCompany
  | UpdateAutoDialerCompany
  | StoreUpdateAutoDialerCompany
  | GetAutoDialerTeams
  | StoreGetAutoDialerTeams
  | AddAutoDialerTeam
  | StoreAddAutoDialerTeam
  | DelAutoDialerTeam
  | StoreDelAutoDialerTeam
  | UpdateAutoDialerTeam
  | StoreUpdateAutoDialerTeam
  | GetAutoDialerTeamMembers
  | StoreGetAutoDialerTeamMembers
  | AddAutoDialerTeamMember
  | StoreAddAutoDialerTeamMember
  | AddAutoDialerTeamMembers
  | StoreAddAutoDialerTeamMembers
  | DelAutoDialerTeamMember
  | StoreDelAutoDialerTeamMember
  | UpdateAutoDialerTeamMember
  | StoreUpdateAutoDialerTeamMember
  | GetAutoDialerLists
  | StoreGetAutoDialerLists
  | AddAutoDialerList
  | StoreAddAutoDialerList
  | DelAutoDialerList
  | StoreDelAutoDialerList
  | UpdateAutoDialerList
  | StoreUpdateAutoDialerList
  | GetAutoDialerListMembers
  | StoreGetAutoDialerListMembers
  | AddAutoDialerListMember
  | StoreAddAutoDialerListMember
  | AddAutoDialerListMembers
  | StoreAddAutoDialerListMembers
  | DelAutoDialerListMember
  | StoreDelAutoDialerListMember
  | UpdateAutoDialerListMember
  | StoreUpdateAutoDialerListMember
  | GetAutoDialerReducers
  | StoreGetAutoDialerReducers
  | AddAutoDialerReducer
  | StoreAddAutoDialerReducer
  | DelAutoDialerReducer
  | StoreDelAutoDialerReducer
  | UpdateAutoDialerReducer
  | StoreUpdateAutoDialerReducer
  | GetAutoDialerReducerMembers
  | StoreGetAutoDialerReducerMembers
  | AddAutoDialerReducerMember
  | StoreAddAutoDialerReducerMember
  | DelAutoDialerReducerMember
  | StoreDelAutoDialerReducerMember
  | UpdateAutoDialerReducerMember
  | StoreUpdateAutoDialerReducerMember
  | StoreNewAutoDialerReducerMembers
  | StoreDropNewAutoDialerReducerMembers
  | StoreNewAutoDialerTeamMembers
  | StoreDropNewAutoDialerTeamMembers
  | StoreSetChangedAutodialerListMemberField
  ;
