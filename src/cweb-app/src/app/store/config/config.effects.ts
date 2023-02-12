import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  Failure,
  ConfigActionTypes,
  GetModules,
  StoreGetModules,
  ReloadModule,
  StoreReloadModule,
  StoreUnloadModule,
  UnloadModule,
  SwitchModule,
  StoreSwitchModule,
  ImportConfModule,
  StoreImportConfModule,
  AutoloadModule,
  StoreAutoloadModule,
  ImportAllModules,
  StoreImportAllModules,
  FromScratchConfModule,
  StoreFromScratchConfModule,
  LoadModule,
  StoreLoadModule,
  TruncateModuleConfig,
  StoreTruncateModuleConfig,
  ImportXMLModuleConfig,
  StoreImportXMLModuleConfig,
  StoreGotModuleError,
} from './config.actions';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../services/ws-data.service';

@Injectable()
export class ConfigEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetModules: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_MODULES),
      map((action: GetModules) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreGetModules({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  ReloadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.RELOAD_MODULE),
      map((action: ReloadModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreReloadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  ImportAllModules: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.IMPORT_ALL_MODULES),
      map((action: ImportAllModules) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreImportAllModules({response}) && new GetModules(null);
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UnloadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UNLOAD_MODULE),
      map((action: UnloadModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreUnloadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  LoadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.LOAD_MODULE),
      map((action: LoadModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreLoadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_MODULE),
      map((action: SwitchModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreSwitchModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  ImportConfModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.IMPORT_MODULE),
      map((action: ImportConfModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreImportConfModule({response}) && new GetModules(null);
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  FromScratchConfModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.FROM_SCRATCH_MODULE),
      map((action: FromScratchConfModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreFromScratchConfModule({response}) && new GetModules(null);
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AutoloadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AUTOLOAD_MODULE),
      map((action: AutoloadModule) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreAutoloadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  TruncateModuleConfig: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.TruncateModuleConfig),
      map((action: TruncateModuleConfig) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreTruncateModuleConfig({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  ImportXMLModuleConfig: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ImportXMLModuleConfig),
      map((action: ImportXMLModuleConfig) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotModuleError({error: response.error});
              }
              return new StoreImportXMLModuleConfig({response});
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
