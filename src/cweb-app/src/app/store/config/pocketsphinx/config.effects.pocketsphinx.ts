
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchPocketsphinxParameter,
  GetPocketsphinx,
  StoreDelPocketsphinxParameter,
  StoreSwitchPocketsphinxParameter,
  UpdatePocketsphinxParameter,
  StoreGetPocketsphinx,
  StoreAddPocketsphinxParameter,
  DelPocketsphinxParameter,
  StoreUpdatePocketsphinxParameter,
  StoreGotPocketsphinxError,
  AddPocketsphinxParameter,
} from './config.actions.pocketsphinx';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsPocketsphinx {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetPocketsphinx: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetPocketsphinx),
      map((action: GetPocketsphinx) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPocketsphinxError({error: response.error});
              }
              return new StoreGetPocketsphinx({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPocketsphinxError({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePocketsphinxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePocketsphinxParameter),
      map((action: UpdatePocketsphinxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPocketsphinxError({error: response.error});
              }
              return new StoreUpdatePocketsphinxParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPocketsphinxError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPocketsphinxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPocketsphinxParameter),
      map((action: SwitchPocketsphinxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPocketsphinxError({error: response.error});
              }
              return new StoreSwitchPocketsphinxParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPocketsphinxError({error: error}));
            }),
          );
        }
      ));
  });

  AddPocketsphinxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPocketsphinxParameter),
      map((action: AddPocketsphinxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPocketsphinxError({error: response.error});
              }
              return new StoreAddPocketsphinxParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPocketsphinxError({error: error}));
            }),
          );
        }
      ));
  });

  DelPocketsphinxParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPocketsphinxParameter),
      map((action: DelPocketsphinxParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPocketsphinxError({error: response.error});
              }
              return new StoreDelPocketsphinxParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPocketsphinxError({error: error}));
            }),
          );
        }
      ));
  });
}

