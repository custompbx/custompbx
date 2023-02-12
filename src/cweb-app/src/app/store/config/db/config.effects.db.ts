
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchDbParameter,
  GetDb,
  StoreDelDbParameter,
  StoreSwitchDbParameter,
  UpdateDbParameter,
  StoreGetDb,
  StoreAddDbParameter,
  DelDbParameter,
  StoreUpdateDbParameter,
  StoreGotDbError,
  AddDbParameter,
} from './config.actions.db';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsDb {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetDb: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetDb),
      map((action: GetDb) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDbError({error: response.error});
              }
              return new StoreGetDb({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDbError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDbParameter),
      map((action: UpdateDbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDbError({error: response.error});
              }
              return new StoreUpdateDbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDbError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchDbParameter),
      map((action: SwitchDbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDbError({error: response.error});
              }
              return new StoreSwitchDbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDbError({error: error}));
            }),
          );
        }
      ));
  });

  AddDbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDbParameter),
      map((action: AddDbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDbError({error: response.error});
              }
              return new StoreAddDbParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDbError({error: error}));
            }),
          );
        }
      ));
  });

  DelDbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDbParameter),
      map((action: DelDbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDbError({error: response.error});
              }
              return new StoreDelDbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDbError({error: error}));
            }),
          );
        }
      ));
  });
}

