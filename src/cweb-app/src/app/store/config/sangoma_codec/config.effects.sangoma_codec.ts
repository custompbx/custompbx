import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchSangomaCodecParameter,
  GetSangomaCodec,
  StoreDelSangomaCodecParameter,
  StoreSwitchSangomaCodecParameter,
  UpdateSangomaCodecParameter,
  StoreGetSangomaCodec,
  StoreAddSangomaCodecParameter,
  DelSangomaCodecParameter,
  StoreUpdateSangomaCodecParameter,
  StoreGotSangomaCodecError,
  AddSangomaCodecParameter,
} from './config.actions.sangoma_codec';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsSangomaCodec {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetSangomaCodec: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetSangomaCodec),
      map((action: GetSangomaCodec) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSangomaCodecError({error: response.error});
              }
              return new StoreGetSangomaCodec({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSangomaCodecError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateSangomaCodecParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateSangomaCodecParameter),
      map((action: UpdateSangomaCodecParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSangomaCodecError({error: response.error});
              }
              return new StoreUpdateSangomaCodecParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSangomaCodecError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchSangomaCodecParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchSangomaCodecParameter),
      map((action: SwitchSangomaCodecParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSangomaCodecError({error: response.error});
              }
              return new StoreSwitchSangomaCodecParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSangomaCodecError({error: error}));
            }),
          );
        }
      ));
  });

  AddSangomaCodecParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddSangomaCodecParameter),
      map((action: AddSangomaCodecParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSangomaCodecError({error: response.error});
              }
              return new StoreAddSangomaCodecParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSangomaCodecError({error: error}));
            }),
          );
        }
      ));
  });

  DelSangomaCodecParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelSangomaCodecParameter),
      map((action: DelSangomaCodecParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotSangomaCodecError({error: response.error});
              }
              return new StoreDelSangomaCodecParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotSangomaCodecError({error: error}));
            }),
          );
        }
      ));
  });
}

