
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchEventMulticastParameter,
  GetEventMulticast,
  StoreDelEventMulticastParameter,
  StoreSwitchEventMulticastParameter,
  UpdateEventMulticastParameter,
  StoreGetEventMulticast,
  StoreAddEventMulticastParameter,
  DelEventMulticastParameter,
  StoreUpdateEventMulticastParameter,
  StoreGotEventMulticastError,
  AddEventMulticastParameter,
} from './config.actions.event_multicast';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsEventMulticast {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetEventMulticast: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetEventMulticast),
      map((action: GetEventMulticast) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEventMulticastError({error: response.error});
              }
              return new StoreGetEventMulticast({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEventMulticastError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateEventMulticastParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateEventMulticastParameter),
      map((action: UpdateEventMulticastParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEventMulticastError({error: response.error});
              }
              return new StoreUpdateEventMulticastParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEventMulticastError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchEventMulticastParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchEventMulticastParameter),
      map((action: SwitchEventMulticastParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEventMulticastError({error: response.error});
              }
              return new StoreSwitchEventMulticastParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEventMulticastError({error: error}));
            }),
          );
        }
      ));
  });

  AddEventMulticastParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddEventMulticastParameter),
      map((action: AddEventMulticastParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEventMulticastError({error: response.error});
              }
              return new StoreAddEventMulticastParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEventMulticastError({error: error}));
            }),
          );
        }
      ));
  });

  DelEventMulticastParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelEventMulticastParameter),
      map((action: DelEventMulticastParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotEventMulticastError({error: response.error});
              }
              return new StoreDelEventMulticastParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotEventMulticastError({error: error}));
            }),
          );
        }
      ));
  });
}

