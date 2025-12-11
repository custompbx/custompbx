import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  DirectoryActionTypes,
  StoreGetDirectoryDomains,
  UpdateFailure,
  GetDirectoryDomains,
  GetDirectoryDomainDetails,
  StoreGetDirectoryDomainDetails,
  StoreAddNewDirectoryDomainParameter,
  AddDirectoryDomain,
  StoreAddDirectoryDomain,
  DeleteDirectoryDomain,
  StoreDeleteDirectoryDomain,
  AddDirectoryDomainVariable,
  StoreAddDirectoryDomainVariable,
  AddDirectoryDomainParameter,
  StoreAddDirectoryDomainParameter,
  DropDirectoryDomainParameter,
  DeleteDirectoryDomainParameter,
  DropDirectoryDomainVariable,
  DeleteDirectoryDomainVariable,
  GetDirectoryUsers,
  StoreGetDirectoryUsers,
  StoreGetDirectoryUserDetails,
  UpdateDirectoryDomainParameter,
  StoreUpdateDirectoryDomainParameter,
  UpdateDirectoryDomainVariable,
  StoreUpdateDirectoryDomainVariable,
  RenameDirectoryDomain,
  StoreRenameDirectoryDomain,
  StoreAddNewDirectoryUserVariable,
  StoreAddNewDirectoryUserParameter,
  StoreDeleteNewDirectoryUserParameter,
  StoreDeleteNewDirectoryUserVariable,
  AddDirectoryUserParameter,
  AddDirectoryUserVariable,
  StoreAddDirectoryUserParameter,
  StoreAddDirectoryUserVariable,
  StoreDeleteDirectoryUserParameter,
  StoreDeleteDirectoryUserVariable,
  StoreUpdateDirectoryUserParameter,
  StoreUpdateDirectoryUserVariable,
  StoreUpdateDirectoryUserCache,
  StoreUpdateDirectoryUserCidr,
  AddDirectoryUser,
  StoreAddDirectoryUser,
  DeleteDirectoryUser,
  UpdateDirectoryUserName,
  StoreUpdateDirectoryUserName,
  StoreDeleteDirectoryUser,
  GetDirectoryGroups,
  StoreGetDirectoryGroups,
  GetDirectoryGroupUsers,
  StoreGetDirectoryGroupUsers,
  AddNewDirectoryGroup,
  StoreAddNewDirectoryGroup,
  DeleteDirectoryGroup,
  StoreDeleteDirectoryGroup,
  UpdateDirectoryGroupName,
  StoreUpdateDirectoryGroupName,
  StoreAddDirectoryGroupUser,
  AddDirectoryGroupUser,
  DeleteDirectoryGroupUser,
  StoreDeleteDirectoryGroupUser,
  GetDirectoryUserGateways,
  StoreGetDirectoryUserGateways,
  GetDirectoryUserGatewayDetails,
  StoreGetDirectoryUserGatewayDetails,
  GetDirectoryUserDetails,
  AddDirectoryUserGatewayParameter,
  DeleteDirectoryUserParameter,
  DeleteDirectoryUserVariable,
  UpdateDirectoryUserParameter,
  UpdateDirectoryUserVariable,
  UpdateDirectoryUserCache,
  UpdateDirectoryUserCidr,
  DeleteDirectoryUserGatewayParameter,
  StoreDirectoryUserGatewayParameter,
  StoreDeleteDirectoryUserGatewayParameter,
  UpdateDirectoryUserGatewayParameter,
  StoreUpdateDirectoryUserGatewayParameter,
  AddDirectoryUserGateway,
  StoreAddDirectoryUserGateway,
  DeleteDirectoryUserGateway,
  UpdateDirectoryUserGatewayName,
  StoreDeleteDirectoryUserGateway,
  StoreUpdateDirectoryUserGatewayName,
  UpdateDirectoryUserGatewayVariable,
  SwitchDirectoryUserGatewayVariable,
  StoreUpdateDirectoryUserGatewayVariable,
  StoreSwitchDirectoryUserGatewayVariable,
  AddDirectoryUserGatewayVariable,
  StoreAddDirectoryUserGatewayVariable,
  DeleteDirectoryUserGatewayVariable,
  StoreDeleteDirectoryUserGatewayVariable,
  SwitchDirectoryDomainParameter,
  SwitchDirectoryDomainVariable,
  StoreSwitchDirectoryDomainVariable,
  StoreSwitchDirectoryDomainParameter,
  SwitchDirectoryUser,
  StoreSwitchDirectoryUser,
  SwitchDirectoryUserParameter,
  StoreSwitchDirectoryUserParameter,
  SwitchDirectoryUserVariable,
  StoreSwitchDirectoryUserVariable,
  ImportDirectory,
  ReduceLoadCounter,
  SwitchDirectoryUserGatewayParameter,
  StoreSwitchDirectoryUserGatewayParameter,
  GetWebUsersByDirectory,
  StoreGetWebUsersByDirectory,
  ImportXMLDomain,
  StoreImportXMLDomain,
  ImportXMLDomainUser,
  StoreImportXMLDomainUser,
  UpdateDirectoryUserNumberAlias,
  StoreUpdateDirectoryUserNumberAlias,
  GetWebDirectoryUsersTemplatesList,
  StoreGetWebDirectoryUsersTemplatesList,
  GetWebDirectoryUsersTemplateForm,
  StoreGetWebDirectoryUsersTemplateForm,
  CreateWebDirectoryUsersByTemplate,
  StoreCreateWebDirectoryUsersByTemplate,
  StoreDirectoryError,
  SwitchDirectoryDomain,
  StoreSwitchDirectoryDomain
} from './directory.actions';
import {catchError, concatMap, map, switchMap, tap} from 'rxjs/operators';
import {WsDataService} from '../../services/ws-data.service';
import {StoreGotAlsaError} from '../config/alsa/config.actions.alsa';


@Injectable({
  providedIn: 'root'
})
export class DirectoryEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  ImportDirectory: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.ImportDirectory),
      map((action: ImportDirectory) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            concatMap((response) => [
              new GetDirectoryDomains(null),
              new ReduceLoadCounter(),
            ]),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetDomains: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryDomains),
      map((action: GetDirectoryDomains) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryDomains({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SetDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryDomain),
      map((action: AddDirectoryDomain) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryDomain({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetDetails: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryDomainDetails),
      map((action: GetDirectoryDomainDetails) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryDomainDetails({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryDomain),
      map((action: DeleteDirectoryDomain) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryDomain({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  NewDomainParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryDomainParameter),
      map((action: AddDirectoryDomainParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryDomainParameter({response: response, param_index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  NewDomainVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryDomainVariable),
      map((action: AddDirectoryDomainVariable) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryDomainVariable({response: response, var_index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelDomainParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryDomainParameter),
      map((action: DeleteDirectoryDomainParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, {id: action.payload.index}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new DropDirectoryDomainParameter({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelDomainVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryDomainVariable),
      map((action: DeleteDirectoryDomainVariable) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, {id: action.payload.index}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new DropDirectoryDomainVariable({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDomainParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryDomainParameter),
      map((action: UpdateDirectoryDomainParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryDomainParameter({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDomainVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryDomainVariable),
      map((action: UpdateDirectoryDomainVariable) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryDomainVariable({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  RenameDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.RenameDirectoryDomain),
      map((action: RenameDirectoryDomain) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreRenameDirectoryDomain({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryDomain),
      map((action: SwitchDirectoryDomain) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryDomain({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  // users
  GetUsers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryUsers),
      map((action: GetDirectoryUsers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryUsers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetUserDetails: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryUserDetails),
      map((action: GetDirectoryUserDetails) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryUserDetails({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  NewUserParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryUserParameter),
      map((action: AddDirectoryUserParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryUserParameter({response: response, param_index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  NewUserVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryUserVariable),
      map((action: AddDirectoryUserVariable) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryUserVariable({response: response, var_index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelUserParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryUserParameter),
      map((action: DeleteDirectoryUserParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, {id: action.payload.index}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryUserParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelUserVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryUserVariable),
      map((action: DeleteDirectoryUserVariable) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, {id: action.payload.index}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryUserVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserParameter),
      map((action: UpdateDirectoryUserParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserVariable),
      map((action: UpdateDirectoryUserVariable) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserCache: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserCache),
      map((action: UpdateDirectoryUserCache) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserCache({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserNumberAlias: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserNumberAlias),
      map((action: UpdateDirectoryUserNumberAlias) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserNumberAlias({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserCidr: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserCidr),
      map((action: UpdateDirectoryUserCidr) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserCidr({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  AddNewUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryUser),
      map((action: AddDirectoryUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryUser),
      map((action: DeleteDirectoryUser) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  RenameUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserName),
      map((action: UpdateDirectoryUserName) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserName({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  // groups
  GetGroups: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryGroups),
      map((action: GetDirectoryGroups) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryGroups({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetGroupUsers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryGroupUsers),
      map((action: GetDirectoryGroupUsers) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryGroupUsers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  AddNewGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddNewDirectoryGroup),
      map((action: AddNewDirectoryGroup) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, {id: action.payload.domainId, name: action.payload.userName}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddNewDirectoryGroup({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryGroup),
      map((action: DeleteDirectoryGroup) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryGroup({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  RenameGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryGroupName),
      map((action: UpdateDirectoryGroupName) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryGroupName({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  AddGroupUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryGroupUser),
      map((action: AddDirectoryGroupUser) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryGroupUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelGroupUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryGroupUser),
      map((action: DeleteDirectoryGroupUser) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryGroupUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  // gateways
  GetUserGateways: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes. GetDirectoryUserGateways),
      map((action: GetDirectoryUserGateways) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryUserGateways({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetUserGatewaysDetails: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetDirectoryUserGatewayDetails),
      map((action: GetDirectoryUserGatewayDetails) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryUserGatewayDetails({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  NewUserGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryUserGatewayParameter),
      map((action: AddDirectoryUserGatewayParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDirectoryUserGatewayParameter({response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelUserGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryUserGatewayParameter),
      map((action: DeleteDirectoryUserGatewayParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, {id: action.payload.index}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryUserGatewayParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserGatewayParameter),
      map((action: UpdateDirectoryUserGatewayParameter) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserGatewayParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  AddNewUserGateway: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryUserGateway),
      map((action: AddDirectoryUserGateway) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryUserGateway({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelUserGateway: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryUserGateway),
      map((action: DeleteDirectoryUserGateway) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryUserGateway({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  RenameUserGateway: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserGatewayName),
      map((action: UpdateDirectoryUserGatewayName) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserGatewayName({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUserGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.UpdateDirectoryUserGatewayVariable),
      map((action: UpdateDirectoryUserGatewayVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryUserGatewayVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUserGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryUserGatewayVariable),
      map((action: SwitchDirectoryUserGatewayVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryUserGatewayVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUserGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryUserGatewayParameter),
      map((action: SwitchDirectoryUserGatewayParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryUserGatewayParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  AddUserGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.AddDirectoryUserGatewayVariable),
      map((action: AddDirectoryUserGatewayVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryUserGatewayVariable({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  DelUserGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.DeleteDirectoryUserGatewayVariable),
      map((action: DeleteDirectoryUserGatewayVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreDeleteDirectoryUserGatewayVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDomainParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryDomainParameter),
      map((action: SwitchDirectoryDomainParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryDomainParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDomainVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryDomainVariable),
      map((action: SwitchDirectoryDomainVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryDomainVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryUser),
      map((action: SwitchDirectoryUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUserParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryUserParameter),
      map((action: SwitchDirectoryUserParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryUserParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUserVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.SwitchDirectoryUserVariable),
      map((action: SwitchDirectoryUserVariable) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryUserVariable({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebUsersByDirectory: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetWebUsersByDirectory),
      map((action: GetWebUsersByDirectory) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetWebUsersByDirectory({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  ImportXMLDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.ImportXMLDomain),
      map((action: ImportXMLDomain) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreImportXMLDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  ImportXMLDomainUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.ImportXMLDomainUser),
      map((action: ImportXMLDomainUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreImportXMLDomainUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebDirectoryUsersTemplatesList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetWebDirectoryUsersTemplatesList),
      map((action: GetWebDirectoryUsersTemplatesList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetWebDirectoryUsersTemplatesList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebDirectoryUsersTemplateForm: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.GetWebDirectoryUsersTemplateForm),
      map((action: GetWebDirectoryUsersTemplateForm) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreGetWebDirectoryUsersTemplateForm({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

  CreateWebDirectoryUsersByTemplate: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DirectoryActionTypes.CreateWebDirectoryUsersByTemplate),
      map((action: CreateWebDirectoryUsersByTemplate) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreDirectoryError({error: response.error});
              }
              return new StoreCreateWebDirectoryUsersByTemplate({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new UpdateFailure({error: error}));
            }),
          );
        }
      ));
  });

}
