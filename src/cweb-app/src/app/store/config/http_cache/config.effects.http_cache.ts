
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchHttpCacheParameter,
  GetHttpCache,
  StoreDelHttpCacheParameter,
  StoreSwitchHttpCacheParameter,
  UpdateHttpCacheParameter,
  StoreGetHttpCache,
  StoreAddHttpCacheParameter,
  DelHttpCacheParameter,
  StoreUpdateHttpCacheParameter,
  StoreGotHttpCacheError,
  AddHttpCacheParameter,
  UpdateHttpCacheProfileParam,
  StoreUpdateHttpCacheProfileParam,
  SwitchHttpCacheProfileParam,
  StoreSwitchHttpCacheProfileParam,
  DelHttpCacheProfileParam,
  AddHttpCacheProfileParam,
  StoreAddHttpCacheProfileParam,
  StoreDelHttpCacheProfileParam,
  AddHttpCacheProfile,
  DelHttpCacheProfile,
  StoreAddHttpCacheProfile,
  StoreDelHttpCacheProfile,
  RenameHttpCacheProfile,
  StoreRenameHttpCacheProfile,
  StoreUpdateHttpCacheProfileDomain,
  StoreGetHttpCacheProfileParameters,
  StoreSwitchHttpCacheProfileDomain,
  StoreAddHttpCacheProfileDomain,
  StoreDelHttpCacheProfileDomain,
  UpdateHttpCacheProfileAws,
  UpdateHttpCacheProfileAzure,
  StoreUpdateHttpCacheProfileAws, StoreUpdateHttpCacheProfileAzure,
} from './config.actions.http_cache';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {Failure} from '../config.actions';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsHttpCache {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetHttpCache: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetHttpCache),
      map((action: GetHttpCache) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreGetHttpCache({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateHttpCacheParameter),
      map((action: UpdateHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreUpdateHttpCacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchHttpCacheParameter),
      map((action: SwitchHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreSwitchHttpCacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  AddHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddHttpCacheParameter),
      map((action: AddHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreAddHttpCacheParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  DelHttpCacheParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelHttpCacheParameter),
      map((action: DelHttpCacheParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreDelHttpCacheParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotHttpCacheError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateHttpCacheProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateHttpCacheProfileParam),
      map((action: UpdateHttpCacheProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreUpdateHttpCacheProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchHttpCacheProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchHttpCacheProfileParam),
      map((action: SwitchHttpCacheProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreSwitchHttpCacheProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddHttpCacheProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddHttpCacheProfileParam),
      map((action: AddHttpCacheProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreAddHttpCacheProfileParam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelHttpCacheProfileParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelHttpCacheProfileParam),
      map((action: DelHttpCacheProfileParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreDelHttpCacheProfileParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddHttpCacheProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddHttpCacheProfile),
      map((action: AddHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreAddHttpCacheProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelHttpCacheProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelHttpCacheProfile),
      map((action: DelHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreDelHttpCacheProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameHttpCacheProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.RenameHttpCacheProfile),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreRenameHttpCacheProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetHttpCacheProfileParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetHttpCacheProfileParameters),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreGetHttpCacheProfileParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
  AddHttpCacheProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddHttpCacheProfileDomain),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreAddHttpCacheProfileDomain({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
  DelHttpCacheProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelHttpCacheProfileDomain),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreDelHttpCacheProfileDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
  SwitchHttpCacheProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchHttpCacheProfileDomain),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreSwitchHttpCacheProfileDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
  UpdateHttpCacheProfileDomain: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateHttpCacheProfileDomain),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreUpdateHttpCacheProfileDomain({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
  UpdateHttpCacheProfileAws: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateHttpCacheProfileAws),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreUpdateHttpCacheProfileAws({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });
  UpdateHttpCacheProfileAzure: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateHttpCacheProfileAzure),
      map((action: RenameHttpCacheProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHttpCacheError({error: response.error});
              }
              return new StoreUpdateHttpCacheProfileAzure({response});
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

