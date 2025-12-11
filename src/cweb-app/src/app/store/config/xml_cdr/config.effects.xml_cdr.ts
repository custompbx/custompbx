
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchXmlCdrParameter,
  GetXmlCdr,
  StoreDelXmlCdrParameter,
  StoreSwitchXmlCdrParameter,
  UpdateXmlCdrParameter,
  StoreGetXmlCdr,
  StoreAddXmlCdrParameter,
  DelXmlCdrParameter,
  StoreUpdateXmlCdrParameter,
  StoreGotXmlCdrError,
  AddXmlCdrParameter,
} from './config.actions.xml_cdr';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsXmlCdr {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetXmlCdr: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetXmlCdr),
      map((action: GetXmlCdr) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlCdrError({error: response.error});
              }
              return new StoreGetXmlCdr({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlCdrError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateXmlCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateXmlCdrParameter),
      map((action: UpdateXmlCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlCdrError({error: response.error});
              }
              return new StoreUpdateXmlCdrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlCdrError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchXmlCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchXmlCdrParameter),
      map((action: SwitchXmlCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlCdrError({error: response.error});
              }
              return new StoreSwitchXmlCdrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlCdrError({error: error}));
            }),
          );
        }
      ));
  });

  AddXmlCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddXmlCdrParameter),
      map((action: AddXmlCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlCdrError({error: response.error});
              }
              return new StoreAddXmlCdrParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlCdrError({error: error}));
            }),
          );
        }
      ));
  });

  DelXmlCdrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelXmlCdrParameter),
      map((action: DelXmlCdrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlCdrError({error: response.error});
              }
              return new StoreDelXmlCdrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlCdrError({error: error}));
            }),
          );
        }
      ));
  });
}

