
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchNibblebillParameter,
  GetNibblebill,
  StoreDelNibblebillParameter,
  StoreSwitchNibblebillParameter,
  UpdateNibblebillParameter,
  StoreGetNibblebill,
  StoreAddNibblebillParameter,
  DelNibblebillParameter,
  StoreUpdateNibblebillParameter,
  StoreGotNibblebillError,
  AddNibblebillParameter,
} from './config.actions.nibblebill';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsNibblebill {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetNibblebill: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetNibblebill),
      map((action: GetNibblebill) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotNibblebillError({error: response.error});
              }
              return new StoreGetNibblebill({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotNibblebillError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateNibblebillParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateNibblebillParameter),
      map((action: UpdateNibblebillParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotNibblebillError({error: response.error});
              }
              return new StoreUpdateNibblebillParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotNibblebillError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchNibblebillParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchNibblebillParameter),
      map((action: SwitchNibblebillParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotNibblebillError({error: response.error});
              }
              return new StoreSwitchNibblebillParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotNibblebillError({error: error}));
            }),
          );
        }
      ));
  });

  AddNibblebillParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddNibblebillParameter),
      map((action: AddNibblebillParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotNibblebillError({error: response.error});
              }
              return new StoreAddNibblebillParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotNibblebillError({error: error}));
            }),
          );
        }
      ));
  });

  DelNibblebillParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelNibblebillParameter),
      map((action: DelNibblebillParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotNibblebillError({error: response.error});
              }
              return new StoreDelNibblebillParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotNibblebillError({error: error}));
            }),
          );
        }
      ));
  });
}

