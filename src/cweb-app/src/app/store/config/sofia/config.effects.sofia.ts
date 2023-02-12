import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  AddSofiaGlobalSettings, AddSofiaProfile, AddSofiaProfileAlias,
  AddSofiaProfileDomain, AddSofiaProfileGateway,
  AddSofiaProfileGatewayParam, AddSofiaProfileGatewayVar,
  AddSofiaProfileParam,
  ConfigActionTypes,
  DelSofiaGlobalSettings,
  DelSofiaProfile,
  DelSofiaProfileAlias,
  DelSofiaProfileDomain, DelSofiaProfileGateway,
  DelSofiaProfileGatewayParam,
  DelSofiaProfileGatewayVar,
  DelSofiaProfileParam,
  GetSofiaGlobalSettings,
  GetSofiaProfileAliases,
  GetSofiaProfileDomains,
  GetSofiaProfileGateways,
  GetSofiaProfiles,
  GetSofiaProfilesParams, RenameSofiaProfile,
  RenameSofiaProfileGateway,
  SofiaProfileCommand,
  StoreAddSofiaGlobalSettings,
  StoreAddSofiaProfile, StoreAddSofiaProfileAlias,
  StoreAddSofiaProfileDomain,
  StoreAddSofiaProfileGateway,
  StoreAddSofiaProfileGatewayParam, StoreAddSofiaProfileGatewayVar,
  StoreAddSofiaProfileParam,
  StoreDelSofiaGlobalSettings,
  StoreDelSofiaProfile,
  StoreDelSofiaProfileAlias, StoreDelSofiaProfileDomain,
  StoreDelSofiaProfileGateway,
  StoreDelSofiaProfileGatewayParam,
  StoreDelSofiaProfileGatewayVar, StoreDelSofiaProfileParam,
  StoreGetSofiaGlobalSettings, StoreGetSofiaProfileAliases,
  StoreGetSofiaProfileDomains, StoreGetSofiaProfileGatewayParameters,
  StoreGetSofiaProfileGateways, StoreGetSofiaProfileGatewayVariables,
  StoreGetSofiaProfiles,
  StoreGetSofiaProfilesParams,
  StoreGotSofiaError,
  StoreRenameSofiaProfile,
  StoreRenameSofiaProfileGateway,
  StoreSofiaProfileCommand,
  StoreSwitchSofiaGlobalSettings,
  StoreSwitchSofiaProfile,
  StoreSwitchSofiaProfileAlias, StoreSwitchSofiaProfileDomain, StoreSwitchSofiaProfileGatewayParam,
  StoreSwitchSofiaProfileGatewayVar,
  StoreSwitchSofiaProfileParam, StoreUpdateSofiaGlobalSettings,
  StoreUpdateSofiaProfileAlias,
  StoreUpdateSofiaProfileDomain,
  StoreUpdateSofiaProfileGatewayParam,
  StoreUpdateSofiaProfileGatewayVar,
  StoreUpdateSofiaProfileParam, SwitchSofiaGlobalSettings,
  SwitchSofiaProfile,
  SwitchSofiaProfileAlias,
  SwitchSofiaProfileDomain,
  SwitchSofiaProfileGatewayParam, SwitchSofiaProfileGatewayVar,
  SwitchSofiaProfileParam,
  UpdateSofiaGlobalSettings,
  UpdateSofiaProfileAlias,
  UpdateSofiaProfileDomain,
  UpdateSofiaProfileGatewayParam, UpdateSofiaProfileGatewayVar,
  UpdateSofiaProfileParam,
} from './config.actions.sofia';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {Failure} from '../config.actions';

@Injectable()
export class ConfigEffectsSofia {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetSofiaGlobalSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_SOFIA_GLOBAL_SETTINGS),
      map((action: GetSofiaGlobalSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreGetSofiaGlobalSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetSofiaProfiles: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_SOFIA_PROFILES),
      map((action: GetSofiaProfiles) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreGetSofiaProfiles({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetSofiaProfilesParams: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_SOFIA_PROFILES_PARAMS),
      map((action: GetSofiaProfilesParams) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreGetSofiaProfilesParams({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSofiaGlobalSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_SOFIA_GLOBAL_SETTING),
      map((action: UpdateSofiaGlobalSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreUpdateSofiaGlobalSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaGlobalSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_GLOBAL_SETTING),
      map((action: SwitchSofiaGlobalSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaGlobalSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaGlobalSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_GLOBAL_SETTING),
      map((action: AddSofiaGlobalSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaGlobalSettings({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaGlobalSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_GLOBAL_SETTING),
      map((action: DelSofiaGlobalSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaGlobalSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSofiaProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_SOFIA_PROFILE_PARAM),
      map((action: UpdateSofiaProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreUpdateSofiaProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_PROFILE_PARAM),
      map((action: SwitchSofiaProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE_PARAM),
      map((action: AddSofiaProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfileParam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE_PARAM),
      map((action: DelSofiaProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetSofiaProfileGateways: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_SOFIA_PROFILE_GATEWAYS),
      map((action: GetSofiaProfileGateways) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreGetSofiaProfileGateways({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSofiaProfileGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_SOFIA_PROFILE_GATEWAY_PARAM),
      map((action: UpdateSofiaProfileGatewayParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreUpdateSofiaProfileGatewayParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaProfileGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_PROFILE_GATEWAY_PARAM),
      map((action: SwitchSofiaProfileGatewayParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaProfileGatewayParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfileGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY_PARAM),
      map((action: AddSofiaProfileGatewayParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfileGatewayParam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfileGatewayParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY_PARAM),
      map((action: DelSofiaProfileGatewayParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfileGatewayParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });


  UpdateSofiaProfileGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_SOFIA_PROFILE_GATEWAY_VAR),
      map((action: UpdateSofiaProfileGatewayVar) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreUpdateSofiaProfileGatewayVar({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaProfileGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_PROFILE_GATEWAY_VAR),
      map((action: SwitchSofiaProfileGatewayVar) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaProfileGatewayVar({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfileGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY_VAR),
      map((action: AddSofiaProfileGatewayVar) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfileGatewayVar({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfileGatewayVar: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY_VAR),
      map((action: DelSofiaProfileGatewayVar) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfileGatewayVar({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetSofiaProfileDomains: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_SOFIA_PROFILE_DOMAINS),
      map((action: GetSofiaProfileDomains) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreGetSofiaProfileDomains({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSofiaProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_SOFIA_PROFILE_DOMAIN),
      map((action: UpdateSofiaProfileDomain) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreUpdateSofiaProfileDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_PROFILE_DOMAIN),
      map((action: SwitchSofiaProfileDomain) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaProfileDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE_DOMAIN),
      map((action: AddSofiaProfileDomain) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfileDomain({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE_DOMAIN),
      map((action: DelSofiaProfileDomain) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfileDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetSofiaProfileAliases: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_SOFIA_PROFILE_ALIASES),
      map((action: GetSofiaProfileAliases) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreGetSofiaProfileAliases({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSofiaProfileAlias: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_SOFIA_PROFILE_ALIAS),
      map((action: UpdateSofiaProfileAlias) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreUpdateSofiaProfileAlias({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaProfileAlias: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_PROFILE_ALIAS),
      map((action: SwitchSofiaProfileAlias) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaProfileAlias({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfileAlias: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE_ALIAS),
      map((action: AddSofiaProfileAlias) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfileAlias({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfileAlias: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE_ALIAS),
      map((action: DelSofiaProfileAlias) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfileAlias({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE),
      map((action: AddSofiaProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddSofiaProfileGateway: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_SOFIA_PROFILE_GATEWAY),
      map((action: AddSofiaProfileGateway) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreAddSofiaProfileGateway({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE),
      map((action: DelSofiaProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameSofiaProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.RENAME_SOFIA_PROFILE),
      map((action: RenameSofiaProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreRenameSofiaProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelSofiaProfileGateway: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_SOFIA_PROFILE_GATEWAY),
      map((action: DelSofiaProfileGateway) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreDelSofiaProfileGateway({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameSofiaProfileGateway: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.RENAME_SOFIA_PROFILE_GATEWAY),
      map((action: RenameSofiaProfileGateway) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreRenameSofiaProfileGateway({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SofiaProfileCommand: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SOFIA_PROFILE_COMMAND),
      map((action: SofiaProfileCommand) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSofiaProfileCommand({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSofiaProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_SOFIA_PROFILE),
      map((action: SwitchSofiaProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return new StoreSwitchSofiaProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SofiaEffect: Observable<any> = createEffect(() => {
    return this.actions
    .pipe(
      ofType(...[
        ConfigActionTypes.GetSofiaProfileGatewayVariables,
        ConfigActionTypes.GetSofiaProfileGatewayParameters,
      ]),
      map((action: any) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSofiaError({error: response.error});
              }
              return this.SofiaStoreEffects(action.type, response);
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  private SofiaStoreEffects = (name: string, response) => {
    switch (name) {
      case 'GetSofiaProfileGatewayVariables':
        return new StoreGetSofiaProfileGatewayVariables({response});
      case 'GetSofiaProfileGatewayParameters':
        return new StoreGetSofiaProfileGatewayParameters({response});
    }
  }
}

