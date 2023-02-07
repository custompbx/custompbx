
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchPostLoadModule,
  StoreDelPostLoadModule,
  StoreSwitchPostLoadModule,
  UpdatePostLoadModule,
  StoreAddPostLoadModule,
  DelPostLoadModule,
  StoreUpdatePostLoadModule,
  StoreGotPostLoadModuleError,
  AddPostLoadModule, GetPostLoadModules, StoreGetPostLoadModules,
} from './config.actions.PostLoadModules';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsPostLoadModules {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetPostLoadModules: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetPostLoadModules),
      map((action: GetPostLoadModules) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostLoadModuleError({error: response.error});
              }
              return new StoreGetPostLoadModules({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPostLoadModuleError({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePostLoadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePostLoadModule),
      map((action: UpdatePostLoadModule) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostLoadModuleError({error: response.error});
              }
              return new StoreUpdatePostLoadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPostLoadModuleError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPostLoadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPostLoadModule),
      map((action: SwitchPostLoadModule) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostLoadModuleError({error: response.error});
              }
              return new StoreSwitchPostLoadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPostLoadModuleError({error: error}));
            }),
          );
        }
      ));
  });

  AddPostLoadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPostLoadModule),
      map((action: AddPostLoadModule) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostLoadModuleError({error: response.error});
              }
              return new StoreAddPostLoadModule({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPostLoadModuleError({error: error}));
            }),
          );
        }
      ));
  });

  DelPostLoadModule: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPostLoadModule),
      map((action: DelPostLoadModule) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostLoadModuleError({error: response.error});
              }
              return new StoreDelPostLoadModule({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPostLoadModuleError({error: error}));
            }),
          );
        }
      ));
  });
}

