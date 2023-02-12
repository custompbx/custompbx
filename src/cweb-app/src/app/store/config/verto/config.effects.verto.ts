import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddVertoProfile,
  AddVertoProfileParam,
  AddVertoSetting, DelVertoProfile, DelVertoProfileParam,
  DelVertoSetting,
  GetVertoConfig,
  GetVertoProfileParams, RenameVertoProfile, StoreAddVertoProfile, StoreAddVertoProfileParam,
  StoreAddVertoSetting, StoreDelVertoProfile, StoreDelVertoProfileParam,
  StoreDelVertoSetting,
  StoreGetVertoConfig,
  StoreGetVertoProfileParams,
  StoreGotVertoError, StoreRenameVertoProfile, StoreSwitchVertoProfileParam,
  StoreSwitchVertoSetting,
  StoreUpdateVertoProfileParam,
  StoreUpdateVertoSetting, SwitchVertoProfileParam,
  SwitchVertoSetting,
  UpdateVertoProfileParam,
  UpdateVertoSetting, MoveVertoProfileParameter, StoreMoveVertoProfileParameter
} from './config.actions.verto';
import {Failure} from '../config.actions';

@Injectable()
export class ConfigEffectsVerto {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetVertoConfig: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_VERTO_CONFIG),
      map((action: GetVertoConfig) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreGetVertoConfig({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetVertoProfileParams: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_VERTO_PROFILES_PARAMS),
      map((action: GetVertoProfileParams) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreGetVertoProfileParams({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateVertoSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_VERTO_SETTING),
      map((action: UpdateVertoSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreUpdateVertoSetting({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchVertoSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_VERTO_SETTING),
      map((action: SwitchVertoSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreSwitchVertoSetting({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddVertoSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_VERTO_SETTING),
      map((action: AddVertoSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreAddVertoSetting({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelVertoSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_VERTO_SETTING),
      map((action: DelVertoSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreDelVertoSetting({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateVertoProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_VERTO_PROFILE_PARAM),
      map((action: UpdateVertoProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreUpdateVertoProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchVertoProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_VERTO_PROFILE_PARAM),
      map((action: SwitchVertoProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreSwitchVertoProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddVertoProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_VERTO_PROFILE_PARAM),
      map((action: AddVertoProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreAddVertoProfileParam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelVertoProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_VERTO_PROFILE_PARAM),
      map((action: DelVertoProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreDelVertoProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddVertoProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_VERTO_PROFILE),
      map((action: AddVertoProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreAddVertoProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelVertoProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_VERTO_PROFILE),
      map((action: DelVertoProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreDelVertoProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameVertoProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.RENAME_VERTO_PROFILE),
      map((action: RenameVertoProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreRenameVertoProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  MoveVertoProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.MoveVertoProfileParameter),
      map((action: RenameVertoProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVertoError({error: response.error});
              }
              return new StoreMoveVertoProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

}

