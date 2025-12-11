
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchCidlookupParameter,
  GetCidlookup,
  StoreDelCidlookupParameter,
  StoreSwitchCidlookupParameter,
  UpdateCidlookupParameter,
  StoreGetCidlookup,
  StoreAddCidlookupParameter,
  DelCidlookupParameter,
  StoreUpdateCidlookupParameter,
  StoreGotCidlookupError,
  AddCidlookupParameter,
} from './config.actions.cidlookup';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsCidlookup {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCidlookup: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCidlookup),
      map((action: GetCidlookup) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCidlookupError({error: response.error});
              }
              return new StoreGetCidlookup({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCidlookupError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCidlookupParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCidlookupParameter),
      map((action: UpdateCidlookupParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCidlookupError({error: response.error});
              }
              return new StoreUpdateCidlookupParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCidlookupError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCidlookupParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchCidlookupParameter),
      map((action: SwitchCidlookupParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCidlookupError({error: response.error});
              }
              return new StoreSwitchCidlookupParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCidlookupError({error: error}));
            }),
          );
        }
      ));
  });

  AddCidlookupParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCidlookupParameter),
      map((action: AddCidlookupParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCidlookupError({error: response.error});
              }
              return new StoreAddCidlookupParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCidlookupError({error: error}));
            }),
          );
        }
      ));
  });

  DelCidlookupParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCidlookupParameter),
      map((action: DelCidlookupParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCidlookupError({error: response.error});
              }
              return new StoreDelCidlookupParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCidlookupError({error: error}));
            }),
          );
        }
      ));
  });
}

