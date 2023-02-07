
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchEasyrouteParameter,
  GetEasyroute,
  StoreDelEasyrouteParameter,
  StoreSwitchEasyrouteParameter,
  UpdateEasyrouteParameter,
  StoreGetEasyroute,
  StoreAddEasyrouteParameter,
  DelEasyrouteParameter,
  StoreUpdateEasyrouteParameter,
  StoreGotEasyrouteError,
  AddEasyrouteParameter,
} from './config.actions.easyroute';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsEasyroute {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetEasyroute: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetEasyroute),
      map((action: GetEasyroute) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEasyrouteError({error: response.error});
              }
              return new StoreGetEasyroute({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEasyrouteError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateEasyrouteParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateEasyrouteParameter),
      map((action: UpdateEasyrouteParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEasyrouteError({error: response.error});
              }
              return new StoreUpdateEasyrouteParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEasyrouteError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchEasyrouteParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchEasyrouteParameter),
      map((action: SwitchEasyrouteParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEasyrouteError({error: response.error});
              }
              return new StoreSwitchEasyrouteParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEasyrouteError({error: error}));
            }),
          );
        }
      ));
  });

  AddEasyrouteParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddEasyrouteParameter),
      map((action: AddEasyrouteParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEasyrouteError({error: response.error});
              }
              return new StoreAddEasyrouteParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEasyrouteError({error: error}));
            }),
          );
        }
      ));
  });

  DelEasyrouteParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelEasyrouteParameter),
      map((action: DelEasyrouteParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEasyrouteError({error: response.error});
              }
              return new StoreDelEasyrouteParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEasyrouteError({error: error}));
            }),
          );
        }
      ));
  });
}

