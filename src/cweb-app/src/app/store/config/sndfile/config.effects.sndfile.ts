
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchSndfileParameter,
  GetSndfile,
  StoreDelSndfileParameter,
  StoreSwitchSndfileParameter,
  UpdateSndfileParameter,
  StoreGetSndfile,
  StoreAddSndfileParameter,
  DelSndfileParameter,
  StoreUpdateSndfileParameter,
  StoreGotSndfileError,
  AddSndfileParameter,
} from './config.actions.sndfile';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsSndfile {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetSndfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetSndfile),
      map((action: GetSndfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSndfileError({error: response.error});
              }
              return new StoreGetSndfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSndfileError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSndfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateSndfileParameter),
      map((action: UpdateSndfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSndfileError({error: response.error});
              }
              return new StoreUpdateSndfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSndfileError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSndfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchSndfileParameter),
      map((action: SwitchSndfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSndfileError({error: response.error});
              }
              return new StoreSwitchSndfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSndfileError({error: error}));
            }),
          );
        }
      ));
  });

  AddSndfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddSndfileParameter),
      map((action: AddSndfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSndfileError({error: response.error});
              }
              return new StoreAddSndfileParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSndfileError({error: error}));
            }),
          );
        }
      ));
  });

  DelSndfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelSndfileParameter),
      map((action: DelSndfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSndfileError({error: response.error});
              }
              return new StoreDelSndfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSndfileError({error: error}));
            }),
          );
        }
      ));
  });
}

