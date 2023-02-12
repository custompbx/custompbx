import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddVoicemailSetting,
  AddVoicemailProfile,
  AddVoicemailProfileParameter,
  DelVoicemailSetting,
  DelVoicemailProfile,
  DelVoicemailProfileParameter,
  GetVoicemailSettings,
  GetVoicemailProfileParameters,
  StoreAddVoicemailSetting,
  StoreAddVoicemailProfile,
  StoreAddVoicemailProfileParameter,
  StoreDelVoicemailSetting,
  StoreDelVoicemailProfile,
  StoreDelVoicemailProfileParameter,
  StoreGetVoicemailSettings,
  StoreGetVoicemailProfileParameters,
  StoreGotVoicemailError,
  StoreSwitchVoicemailSetting,
  StoreSwitchVoicemailProfileParameter,
  StoreUpdateVoicemailSetting,
  StoreUpdateVoicemailProfile,
  StoreUpdateVoicemailProfileParameter,
  SwitchVoicemailSetting,
  SwitchVoicemailProfileParameter,
  UpdateVoicemailSetting,
  UpdateVoicemailProfile,
  UpdateVoicemailProfileParameter, StoreGetVoicemailProfiles, GetVoicemailProfiles,
} from './config.actions.voicemail';

@Injectable()
export class ConfigEffectsVoicemail {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetVoicemailSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetVoicemailSettings),
      map((action: GetVoicemailSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreGetVoicemailSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateVoicemailSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateVoicemailSetting),
      map((action: UpdateVoicemailSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreUpdateVoicemailSetting({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchVoicemailSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchVoicemailSetting),
      map((action: SwitchVoicemailSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreSwitchVoicemailSetting({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  AddVoicemailSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddVoicemailSetting),
      map((action: AddVoicemailSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreAddVoicemailSetting({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  DelVoicemailSetting: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelVoicemailSetting),
      map((action: DelVoicemailSetting) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreDelVoicemailSetting({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  GetVoicemailProfileParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetVoicemailProfileParameters),
      map((action: GetVoicemailProfileParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreGetVoicemailProfileParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  AddVoicemailProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddVoicemailProfileParameter),
      map((action: AddVoicemailProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreAddVoicemailProfileParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateVoicemailProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateVoicemailProfileParameter),
      map((action: UpdateVoicemailProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreUpdateVoicemailProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchVoicemailProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchVoicemailProfileParameter),
      map((action: SwitchVoicemailProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreSwitchVoicemailProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  GetVoicemailProfiles: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetVoicemailProfiles),
      map((action: GetVoicemailProfiles) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreGetVoicemailProfiles({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  DelVoicemailProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelVoicemailProfileParameter),
      map((action: DelVoicemailProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreDelVoicemailProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  AddVoicemailProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddVoicemailProfile),
      map((action: AddVoicemailProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreAddVoicemailProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  DelVoicemailProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelVoicemailProfile),
      map((action: DelVoicemailProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreDelVoicemailProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateVoicemailProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateVoicemailProfile),
      map((action: UpdateVoicemailProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotVoicemailError({error: response.error});
              }
              return new StoreUpdateVoicemailProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotVoicemailError({error: error}));
            }),
          );
        }
      ));
  });

}

