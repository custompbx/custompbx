
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchErlangEventParameter,
  GetErlangEvent,
  StoreDelErlangEventParameter,
  StoreSwitchErlangEventParameter,
  UpdateErlangEventParameter,
  StoreGetErlangEvent,
  StoreAddErlangEventParameter,
  DelErlangEventParameter,
  StoreUpdateErlangEventParameter,
  StoreGotErlangEventError,
  AddErlangEventParameter,
} from './config.actions.erlang_event';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsErlangEvent {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetErlangEvent: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetErlangEvent),
      map((action: GetErlangEvent) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotErlangEventError({error: response.error});
              }
              return new StoreGetErlangEvent({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotErlangEventError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateErlangEventParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateErlangEventParameter),
      map((action: UpdateErlangEventParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotErlangEventError({error: response.error});
              }
              return new StoreUpdateErlangEventParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotErlangEventError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchErlangEventParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchErlangEventParameter),
      map((action: SwitchErlangEventParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotErlangEventError({error: response.error});
              }
              return new StoreSwitchErlangEventParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotErlangEventError({error: error}));
            }),
          );
        }
      ));
  });

  AddErlangEventParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddErlangEventParameter),
      map((action: AddErlangEventParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotErlangEventError({error: response.error});
              }
              return new StoreAddErlangEventParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotErlangEventError({error: error}));
            }),
          );
        }
      ));
  });

  DelErlangEventParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelErlangEventParameter),
      map((action: DelErlangEventParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotErlangEventError({error: response.error});
              }
              return new StoreDelErlangEventParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotErlangEventError({error: error}));
            }),
          );
        }
      ));
  });
}

