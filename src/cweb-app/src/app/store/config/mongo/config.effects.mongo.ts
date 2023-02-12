
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchMongoParameter,
  GetMongo,
  StoreDelMongoParameter,
  StoreSwitchMongoParameter,
  UpdateMongoParameter,
  StoreGetMongo,
  StoreAddMongoParameter,
  DelMongoParameter,
  StoreUpdateMongoParameter,
  StoreGotMongoError,
  AddMongoParameter,
} from './config.actions.mongo';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsMongo {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetMongo: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetMongo),
      map((action: GetMongo) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMongoError({error: response.error});
              }
              return new StoreGetMongo({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMongoError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateMongoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateMongoParameter),
      map((action: UpdateMongoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMongoError({error: response.error});
              }
              return new StoreUpdateMongoParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMongoError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchMongoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchMongoParameter),
      map((action: SwitchMongoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMongoError({error: response.error});
              }
              return new StoreSwitchMongoParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMongoError({error: error}));
            }),
          );
        }
      ));
  });

  AddMongoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddMongoParameter),
      map((action: AddMongoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMongoError({error: response.error});
              }
              return new StoreAddMongoParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMongoError({error: error}));
            }),
          );
        }
      ));
  });

  DelMongoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelMongoParameter),
      map((action: DelMongoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMongoError({error: response.error});
              }
              return new StoreDelMongoParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMongoError({error: error}));
            }),
          );
        }
      ));
  });
}

