
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchOpusParameter,
  GetOpus,
  StoreDelOpusParameter,
  StoreSwitchOpusParameter,
  UpdateOpusParameter,
  StoreGetOpus,
  StoreAddOpusParameter,
  DelOpusParameter,
  StoreUpdateOpusParameter,
  StoreGotOpusError,
  AddOpusParameter,
} from './config.actions.opus';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsOpus {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetOpus: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOpus),
      map((action: GetOpus) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpusError({error: response.error});
              }
              return new StoreGetOpus({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpusError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOpusParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOpusParameter),
      map((action: UpdateOpusParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpusError({error: response.error});
              }
              return new StoreUpdateOpusParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpusError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOpusParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOpusParameter),
      map((action: SwitchOpusParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpusError({error: response.error});
              }
              return new StoreSwitchOpusParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpusError({error: error}));
            }),
          );
        }
      ));
  });

  AddOpusParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOpusParameter),
      map((action: AddOpusParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpusError({error: response.error});
              }
              return new StoreAddOpusParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpusError({error: error}));
            }),
          );
        }
      ));
  });

  DelOpusParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOpusParameter),
      map((action: DelOpusParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpusError({error: response.error});
              }
              return new StoreDelOpusParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpusError({error: error}));
            }),
          );
        }
      ));
  });
}

