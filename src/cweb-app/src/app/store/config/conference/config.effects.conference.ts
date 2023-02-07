import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddConferenceRoom,
  AddConferenceProfile,
  AddConferenceProfileParameter,
  DelConferenceRoom,
  DelConferenceProfile,
  DelConferenceProfileParameter,
  GetConference,
  GetConferenceProfileParameters,
  StoreAddConferenceRoom,
  StoreAddConferenceProfile,
  StoreAddConferenceProfileParameter,
  StoreDelConferenceRoom,
  StoreDelConferenceProfile,
  StoreDelConferenceProfileParameter,
  StoreGetConference,
  StoreGetConferenceProfileParameters,
  StoreGotConferenceError,
  StoreSwitchConferenceRoom,
  StoreSwitchConferenceProfileParameter,
  StoreUpdateConferenceRoom,
  StoreUpdateConferenceProfile,
  StoreUpdateConferenceProfileParameter,
  SwitchConferenceRoom,
  SwitchConferenceProfileParameter,
  UpdateConferenceRoom,
  UpdateConferenceProfile,
  UpdateConferenceProfileParameter,
  UpdateConferenceCallerControlGroup,
  UpdateConferenceCallerControl,
  AddConferenceCallerControl,
  StoreUpdateConferenceCallerControl,
  StoreGetConferenceCallerControls,
  StoreAddConferenceCallerControlGroup,
  StoreDelConferenceCallerControlGroup,
  StoreDelConferenceCallerControl,
  DelConferenceCallerControlGroup,
  SwitchConferenceCallerControl,
  StoreSwitchConferenceCallerControl,
  StoreUpdateConferenceCallerControlGroup,
  StoreAddConferenceCallerControl,
  DelConferenceCallerControl,
  AddConferenceCallerControlGroup,
  GetConferenceCallerControls,
  StoreGetConferenceChatPermissionUsers,
  AddConferenceChatPermissionUser,
  StoreDelConferenceChatPermission,
  UpdateConferenceChatPermission,
  StoreAddConferenceChatPermission,
  GetConferenceChatPermissionUsers,
  StoreAddConferenceChatPermissionUser,
  DelConferenceChatPermission,
  StoreSwitchConferenceChatPermissionUser,
  DelConferenceChatPermissionUser,
  StoreUpdateConferenceChatPermissionUser,
  AddConferenceChatPermission,
  StoreDelConferenceChatPermissionUser,
  StoreUpdateConferenceChatPermission,
  SwitchConferenceChatPermissionUser,
  UpdateConferenceChatPermissionUser
} from './config.actions.conference';

@Injectable()
export class ConfigEffectsConference {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetConference: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetConference),
      map((action: GetConference) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreGetConference({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceRoom: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceRoom),
      map((action: UpdateConferenceRoom) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceRoom({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchConferenceRoom: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchConferenceRoom),
      map((action: SwitchConferenceRoom) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreSwitchConferenceRoom({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceRoom: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceRoom),
      map((action: AddConferenceRoom) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceRoom({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceRoom: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceRoom),
      map((action: DelConferenceRoom) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceRoom({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  GetConferenceCallerControls: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetConferenceCallerControls),
      map((action: GetConferenceCallerControls) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreGetConferenceCallerControls({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceCallerControl: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceCallerControl),
      map((action: AddConferenceCallerControl) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceCallerControl({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceCallerControl: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceCallerControl),
      map((action: UpdateConferenceCallerControl) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceCallerControl({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchConferenceCallerControl: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchConferenceCallerControl),
      map((action: SwitchConferenceCallerControl) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreSwitchConferenceCallerControl({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceCallerControl: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceCallerControl),
      map((action: DelConferenceCallerControl) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceCallerControl({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceCallerControlGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceCallerControlGroup),
      map((action: AddConferenceCallerControlGroup) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceCallerControlGroup({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceCallerControlGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceCallerControlGroup),
      map((action: DelConferenceCallerControlGroup) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceCallerControlGroup({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceCallerControlGroup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceCallerControlGroup),
      map((action: UpdateConferenceCallerControlGroup) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceCallerControlGroup({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  GetConferenceProfileParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetConferenceProfileParameters),
      map((action: GetConferenceProfileParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreGetConferenceProfileParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceProfileParameter),
      map((action: AddConferenceProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceProfileParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceProfileParameter),
      map((action: UpdateConferenceProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchConferenceProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchConferenceProfileParameter),
      map((action: SwitchConferenceProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreSwitchConferenceProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceProfileParameter),
      map((action: DelConferenceProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceProfile),
      map((action: AddConferenceProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceProfile),
      map((action: DelConferenceProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceProfile),
      map((action: UpdateConferenceProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  GetConferenceChatPermissionUsers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetConferenceChatPermissionUsers),
      map((action: GetConferenceChatPermissionUsers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreGetConferenceChatPermissionUsers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceChatPermissionUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceChatPermissionUser),
      map((action: AddConferenceChatPermissionUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceChatPermissionUser({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceChatPermissionUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceChatPermissionUser),
      map((action: UpdateConferenceChatPermissionUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceChatPermissionUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchConferenceChatPermissionUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchConferenceChatPermissionUser),
      map((action: SwitchConferenceChatPermissionUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreSwitchConferenceChatPermissionUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceChatPermissionUser: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceChatPermissionUser),
      map((action: DelConferenceChatPermissionUser) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceChatPermissionUser({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  AddConferenceChatPermission: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddConferenceChatPermission),
      map((action: AddConferenceChatPermission) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreAddConferenceChatPermission({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  DelConferenceChatPermission: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelConferenceChatPermission),
      map((action: DelConferenceChatPermission) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreDelConferenceChatPermission({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateConferenceChatPermission: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateConferenceChatPermission),
      map((action: UpdateConferenceChatPermission) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotConferenceError({error: response.error});
              }
              return new StoreUpdateConferenceChatPermission({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotConferenceError({error: error}));
            }),
          );
        }
      ));
  });

}

