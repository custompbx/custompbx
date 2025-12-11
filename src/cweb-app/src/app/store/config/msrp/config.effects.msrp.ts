
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchMsrpParameter,
  GetMsrp,
  StoreDelMsrpParameter,
  StoreSwitchMsrpParameter,
  UpdateMsrpParameter,
  StoreGetMsrp,
  StoreAddMsrpParameter,
  DelMsrpParameter,
  StoreUpdateMsrpParameter,
  StoreGotMsrpError,
  AddMsrpParameter,
} from './config.actions.msrp';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsMsrp {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetMsrp: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetMsrp),
      map((action: GetMsrp) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMsrpError({error: response.error});
              }
              return new StoreGetMsrp({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMsrpError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateMsrpParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateMsrpParameter),
      map((action: UpdateMsrpParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMsrpError({error: response.error});
              }
              return new StoreUpdateMsrpParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMsrpError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchMsrpParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchMsrpParameter),
      map((action: SwitchMsrpParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMsrpError({error: response.error});
              }
              return new StoreSwitchMsrpParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMsrpError({error: error}));
            }),
          );
        }
      ));
  });

  AddMsrpParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddMsrpParameter),
      map((action: AddMsrpParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMsrpError({error: response.error});
              }
              return new StoreAddMsrpParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMsrpError({error: error}));
            }),
          );
        }
      ));
  });

  DelMsrpParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelMsrpParameter),
      map((action: DelMsrpParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotMsrpError({error: response.error});
              }
              return new StoreDelMsrpParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotMsrpError({error: error}));
            }),
          );
        }
      ));
  });
}

