
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchCurlParameter,
  GetCurl,
  StoreDelCurlParameter,
  StoreSwitchCurlParameter,
  UpdateCurlParameter,
  StoreGetCurl,
  StoreAddCurlParameter,
  DelCurlParameter,
  StoreUpdateCurlParameter,
  StoreGotCurlError,
  AddCurlParameter,
} from './config.actions.curl';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsCurl {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCurl: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCurl),
      map((action: GetCurl) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCurlError({error: response.error});
              }
              return new StoreGetCurl({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCurlError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCurlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCurlParameter),
      map((action: UpdateCurlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCurlError({error: response.error});
              }
              return new StoreUpdateCurlParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCurlError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCurlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchCurlParameter),
      map((action: SwitchCurlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCurlError({error: response.error});
              }
              return new StoreSwitchCurlParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCurlError({error: error}));
            }),
          );
        }
      ));
  });

  AddCurlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCurlParameter),
      map((action: AddCurlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCurlError({error: response.error});
              }
              return new StoreAddCurlParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCurlError({error: error}));
            }),
          );
        }
      ));
  });

  DelCurlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCurlParameter),
      map((action: DelCurlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCurlError({error: response.error});
              }
              return new StoreDelCurlParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCurlError({error: error}));
            }),
          );
        }
      ));
  });
}

