import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  SettingsActionTypes,
  UpdateSettings,
  Failure,
  GetSettings,
  GetWebUsers,
  StoreGetWebUsers,
  AddWebUser,
  StoreAddWebUser,
  RenameWebUser,
  StoreRenameWebUser,
  DeleteWebUser,
  StoreDeleteWebUser,
  SwitchWebUser,
  StoreSwitchWebUser,
  UpdateWebUserPassword,
  StoreUpdateWebUserPassword,
  UpdateWebUserLang,
  StoreUpdateWebUserLang,
  UpdateWebUserSipUser,
  StoreUpdateWebUserSipUser,
  UpdateWebUserWs,
  StoreUpdateWebUserWs,
  UpdateWebUserAvatar,
  StoreUpdateWebUserAvatar,
  ClearWebUserAvatar,
  StoreClearWebUserAvatar,
  StoreUpdateWebUserStun,
  UpdateWebUserVertoWs,
  StoreUpdateWebUserVertoWs,
  UpdateWebUserWebRTCLib,
  StoreUpdateWebUserWebRTCLib,
  SaveWebSettings,
  GetWebSettings,
  StoreSaveWebSettings,
  StoreGetWebSettings,
  StoreGotWebError,
  StoreGetUserTokens,
  GetUserTokens,
  RemoveUserToken,
  StoreRemoveUserToken,
  AddUserToken,
  StoreAddUserToken,
  StoreUpdateWebUserGroup,
  GetWebDirectoryUsersTemplates,
  StoreGetWebDirectoryUsersTemplates,
  AddWebDirectoryUsersTemplate,
  StoreAddWebDirectoryUsersTemplate,
  StoreDelWebDirectoryUsersTemplate,
  StoreGetWebDirectoryUsersTemplateParameters,
  StoreAddWebDirectoryUsersTemplateVariable,
  StoreDelWebDirectoryUsersTemplateVariable,
  StoreUpdateWebDirectoryUsersTemplate,
  StoreSwitchWebDirectoryUsersTemplateParameter,
  StoreUpdateWebDirectoryUsersTemplateParameter,
  StoreSwitchWebDirectoryUsersTemplateVariable,
  StoreAddWebDirectoryUsersTemplateParameter,
  StoreGetWebDirectoryUsersTemplateVariables,
  StoreSwitchWebDirectoryUsersTemplate,
  StoreDelWebDirectoryUsersTemplateParameter,
  StoreUpdateWebDirectoryUsersTemplateVariable,
} from './settings.actions';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class SettingsEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  Get: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GET_SETTINGS),
      map((action: GetSettings) => action),
      switchMap(payload => {
          return this.ws.getSettings().pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new UpdateSettings({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  Set: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.SET_SETTINGS),
      map((action: GetSettings) => action),
      switchMap(payload => {
          return this.ws.setSettings(payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new UpdateSettings({response});
            }),
            catchError((error) => {
              console.log(error);

              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebUsers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GET_WEB_USERS),
      map((action: GetWebUsers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreGetWebUsers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddWebUsers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.ADD_WEB_USER),
      map((action: AddWebUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreAddWebUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameWebUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.RENAME_WEB_USER),
      map((action: RenameWebUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreRenameWebUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteWebUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.DELETE_WEB_USER),
      map((action: DeleteWebUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreDeleteWebUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchWebUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.SWITCH_WEB_USER),
      map((action: SwitchWebUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreSwitchWebUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserPassword: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_PASSWORD),
      map((action: UpdateWebUserPassword) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserPassword({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserLang: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_LANG),
      map((action: UpdateWebUserLang) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserLang({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserSipUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_SIP_USER),
      map((action: UpdateWebUserSipUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserSipUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserWs: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_WS),
      map((action: UpdateWebUserWs) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserWs({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserVertoWs: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_VERTO_WS),
      map((action: UpdateWebUserVertoWs) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserVertoWs({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserWebRTCLib: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_WEBRTC_LIB),
      map((action: UpdateWebUserWebRTCLib) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserWebRTCLib({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserStun: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_STUN),
      map((action: UpdateWebUserWs) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserStun({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserAvatar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UPDATE_WEB_USER_AVATAR),
      map((action: UpdateWebUserAvatar) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserAvatar({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  ClearWebUserAvatar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.CLEAR_WEB_USER_AVATAR),
      map((action: ClearWebUserAvatar) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreClearWebUserAvatar({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SaveWebSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.SaveWebSettings),
      map((action: SaveWebSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreSaveWebSettings({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GetWebSettings),
      map((action: GetWebSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreGetWebSettings({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetUserTokens: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GetUserTokens),
      map((action: GetUserTokens) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreGetUserTokens({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RemoveUserToken: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.RemoveUserToken),
      map((action: RemoveUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreRemoveUserToken({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddUserToken: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.AddUserToken),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreAddUserToken({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebUserGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UpdateWebUserGroup),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebUserGroup({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebDirectoryUsersTemplates: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GetWebDirectoryUsersTemplates),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreGetWebDirectoryUsersTemplates({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddWebDirectoryUsersTemplate: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.AddWebDirectoryUsersTemplate),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreAddWebDirectoryUsersTemplate({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelWebDirectoryUsersTemplate: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.DelWebDirectoryUsersTemplate),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreDelWebDirectoryUsersTemplate({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchWebDirectoryUsersTemplate: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.SwitchWebDirectoryUsersTemplate),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreSwitchWebDirectoryUsersTemplate({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebDirectoryUsersTemplate: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UpdateWebDirectoryUsersTemplate),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebDirectoryUsersTemplate({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebDirectoryUsersTemplateParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GetWebDirectoryUsersTemplateParameters),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreGetWebDirectoryUsersTemplateParameters({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddWebDirectoryUsersTemplateParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.AddWebDirectoryUsersTemplateParameter),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreAddWebDirectoryUsersTemplateParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelWebDirectoryUsersTemplateParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.DelWebDirectoryUsersTemplateParameter),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreDelWebDirectoryUsersTemplateParameter({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchWebDirectoryUsersTemplateParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.SwitchWebDirectoryUsersTemplateParameter),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreSwitchWebDirectoryUsersTemplateParameter({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebDirectoryUsersTemplateParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UpdateWebDirectoryUsersTemplateParameter),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebDirectoryUsersTemplateParameter({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetWebDirectoryUsersTemplateVariables: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.GetWebDirectoryUsersTemplateVariables),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreGetWebDirectoryUsersTemplateVariables({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddWebDirectoryUsersTemplateVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.AddWebDirectoryUsersTemplateVariable),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreAddWebDirectoryUsersTemplateVariable({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelWebDirectoryUsersTemplateVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.DelWebDirectoryUsersTemplateVariable),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreDelWebDirectoryUsersTemplateVariable({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchWebDirectoryUsersTemplateVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.SwitchWebDirectoryUsersTemplateVariable),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreSwitchWebDirectoryUsersTemplateVariable({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateWebDirectoryUsersTemplateVariable: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(SettingsActionTypes.UpdateWebDirectoryUsersTemplateVariable),
      map((action: AddUserToken) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotWebError({error: response.error});
              }
              return new StoreUpdateWebDirectoryUsersTemplateVariable({response});
            }),
            catchError((error) => {
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

}
