
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchCdrMongodbParameter,
  GetCdrMongodb,
  StoreDelCdrMongodbParameter,
  StoreSwitchCdrMongodbParameter,
  UpdateCdrMongodbParameter,
  StoreGetCdrMongodb,
  StoreAddCdrMongodbParameter,
  DelCdrMongodbParameter,
  StoreUpdateCdrMongodbParameter,
  StoreGotCdrMongodbError,
  AddCdrMongodbParameter,
} from './config.actions.cdr_mongodb';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsCdrMongodb {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCdrMongodb: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCdrMongodb),
      map((action: GetCdrMongodb) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrMongodbError({error: response.error});
              }
              return new StoreGetCdrMongodb({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCdrMongodbError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCdrMongodbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCdrMongodbParameter),
      map((action: UpdateCdrMongodbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrMongodbError({error: response.error});
              }
              return new StoreUpdateCdrMongodbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCdrMongodbError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCdrMongodbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchCdrMongodbParameter),
      map((action: SwitchCdrMongodbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrMongodbError({error: response.error});
              }
              return new StoreSwitchCdrMongodbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCdrMongodbError({error: error}));
            }),
          );
        }
      ));
  });

  AddCdrMongodbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCdrMongodbParameter),
      map((action: AddCdrMongodbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrMongodbError({error: response.error});
              }
              return new StoreAddCdrMongodbParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCdrMongodbError({error: error}));
            }),
          );
        }
      ));
  });

  DelCdrMongodbParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCdrMongodbParameter),
      map((action: DelCdrMongodbParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrMongodbError({error: response.error});
              }
              return new StoreDelCdrMongodbParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCdrMongodbError({error: error}));
            }),
          );
        }
      ));
  });
}

