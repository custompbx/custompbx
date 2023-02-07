import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  AutoDialerActionTypes,
  StoreUpdateAutoDialerList,
  StoreAddAutoDialerListMember,
  StoreGetAutoDialerListMembers,
  StoreUpdateAutoDialerTeam,
  StoreUpdateAutoDialerReducer,
  StoreDelAutoDialerListMember,
  StoreGetAutoDialerReducers,
  StoreAddAutoDialerCompany,
  StoreAddAutoDialerList,
  StoreDelAutoDialerTeamMember,
  StoreAddAutoDialerTeam,
  StoreUpdateAutoDialerCompany,
  StoreGetAutoDialerTeamMembers,
  StoreDelAutoDialerReducer,
  StoreGetAutoDialerTeams,
  StoreDelAutoDialerCompany,
  StoreGetAutoDialerCompanies,
  StoreAddAutoDialerTeamMember,
  StoreGetAutoDialerLists,
  StoreUpdateAutoDialerListMember,
  StoreAddAutoDialerReducer,
  StoreDelAutoDialerList,
  StoreUpdateAutoDialerTeamMember,
  StoreAutoDialerError,
  StoreDelAutoDialerTeam,
  StoreDelAutoDialerReducerMember,
  StoreGetAutoDialerReducerMembers,
  StoreAddAutoDialerReducerMember,
  StoreUpdateAutoDialerReducerMember, StoreDropNewAutoDialerReducerMembers, StoreAddAutoDialerTeamMembers, StoreAddAutoDialerListMembers,
} from './autodialer.actions';
import {catchError, concatMap, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {AddUserToken, StoreGotWebError} from '../../settings/settings.actions';

@Injectable()
export class AutodialerEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAutoDialerCompanies: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerCompanies),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerCompanies({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  AddAutoDialerCompany: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerCompany),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerCompany({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerCompany: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerCompany),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerCompany({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerCompany: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerCompany),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerCompany({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  GetAutoDialerTeams: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerTeams),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerTeams({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerTeam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerTeam),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerTeam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerTeam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerTeam),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerTeam({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerTeam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerTeam),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerTeam({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  GetAutoDialerTeamMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerTeamMembers),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerTeamMembers({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerTeamMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerTeamMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerTeamMember({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerTeamMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerTeamMembers),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerTeamMembers({response: response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerTeamMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerTeamMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerTeamMember({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerTeamMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerTeamMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerTeamMember({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  GetAutoDialerLists: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerLists),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerLists({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerList),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerList({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerList),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerList({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerList),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerList({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  GetAutoDialerListMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerListMembers),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerListMembers({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerListMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerListMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerListMember({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerListMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerListMembers),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerListMembers({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerListMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerListMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerListMember({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerListMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerListMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerListMember({response, payload: action.payload});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  GetAutoDialerReducers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerReducers),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerReducers({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerReducer: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerReducer),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreAddAutoDialerReducer({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerReducer: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerReducer),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerReducer({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerReducer: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerReducer),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerReducer({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  GetAutoDialerReducerMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.GetAutoDialerReducerMembers),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreGetAutoDialerReducerMembers({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  AddAutoDialerReducerMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.AddAutoDialerReducerMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            concatMap((response) => {
              if (response.error) {
                return [new StoreAutoDialerError({error: response.error})];
              }
              return [
                new StoreAddAutoDialerReducerMember({response}),
                new StoreDropNewAutoDialerReducerMembers({id: action.payload.id, index: action.payload.index}),
              ];
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
  DelAutoDialerReducerMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.DelAutoDialerReducerMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreDelAutoDialerReducerMember({response, id: action.payload.parent});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAutoDialerReducerMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(AutoDialerActionTypes.UpdateAutoDialerReducerMember),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreAutoDialerError({error: response.error});
              }
              return new StoreUpdateAutoDialerReducerMember({response});
            }),
            catchError((error) => {
              return of(new StoreAutoDialerError({error: error}));
            }),
          );
        }
      ));
  });
}
