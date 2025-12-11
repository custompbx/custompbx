
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchTtsCommandlineParameter,
  GetTtsCommandline,
  StoreDelTtsCommandlineParameter,
  StoreSwitchTtsCommandlineParameter,
  UpdateTtsCommandlineParameter,
  StoreGetTtsCommandline,
  StoreAddTtsCommandlineParameter,
  DelTtsCommandlineParameter,
  StoreUpdateTtsCommandlineParameter,
  StoreGotTtsCommandlineError,
  AddTtsCommandlineParameter,
} from './config.actions.tts_commandline';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsTtsCommandline {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetTtsCommandline: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetTtsCommandline),
      map((action: GetTtsCommandline) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotTtsCommandlineError({error: response.error});
              }
              return new StoreGetTtsCommandline({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotTtsCommandlineError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateTtsCommandlineParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateTtsCommandlineParameter),
      map((action: UpdateTtsCommandlineParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotTtsCommandlineError({error: response.error});
              }
              return new StoreUpdateTtsCommandlineParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotTtsCommandlineError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchTtsCommandlineParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchTtsCommandlineParameter),
      map((action: SwitchTtsCommandlineParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotTtsCommandlineError({error: response.error});
              }
              return new StoreSwitchTtsCommandlineParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotTtsCommandlineError({error: error}));
            }),
          );
        }
      ));
  });

  AddTtsCommandlineParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddTtsCommandlineParameter),
      map((action: AddTtsCommandlineParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotTtsCommandlineError({error: response.error});
              }
              return new StoreAddTtsCommandlineParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotTtsCommandlineError({error: error}));
            }),
          );
        }
      ));
  });

  DelTtsCommandlineParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelTtsCommandlineParameter),
      map((action: DelTtsCommandlineParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotTtsCommandlineError({error: response.error});
              }
              return new StoreDelTtsCommandlineParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotTtsCommandlineError({error: error}));
            }),
          );
        }
      ));
  });
}

