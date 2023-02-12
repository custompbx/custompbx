
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchFaxParameter,
  GetFax,
  StoreDelFaxParameter,
  StoreSwitchFaxParameter,
  UpdateFaxParameter,
  StoreGetFax,
  StoreAddFaxParameter,
  DelFaxParameter,
  StoreUpdateFaxParameter,
  StoreGotFaxError,
  AddFaxParameter,
} from './config.actions.fax';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsFax {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetFax: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetFax),
      map((action: GetFax) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFaxError({error: response.error});
              }
              return new StoreGetFax({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFaxError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateFaxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateFaxParameter),
      map((action: UpdateFaxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFaxError({error: response.error});
              }
              return new StoreUpdateFaxParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFaxError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchFaxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchFaxParameter),
      map((action: SwitchFaxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFaxError({error: response.error});
              }
              return new StoreSwitchFaxParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFaxError({error: error}));
            }),
          );
        }
      ));
  });

  AddFaxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddFaxParameter),
      map((action: AddFaxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFaxError({error: response.error});
              }
              return new StoreAddFaxParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFaxError({error: error}));
            }),
          );
        }
      ));
  });

  DelFaxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelFaxParameter),
      map((action: DelFaxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFaxError({error: response.error});
              }
              return new StoreDelFaxParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFaxError({error: error}));
            }),
          );
        }
      ));
  });
}

