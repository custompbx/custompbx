
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchXmlRpcParameter,
  GetXmlRpc,
  StoreDelXmlRpcParameter,
  StoreSwitchXmlRpcParameter,
  UpdateXmlRpcParameter,
  StoreGetXmlRpc,
  StoreAddXmlRpcParameter,
  DelXmlRpcParameter,
  StoreUpdateXmlRpcParameter,
  StoreGotXmlRpcError,
  AddXmlRpcParameter,
} from './config.actions.xml_rpc';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsXmlRpc {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetXmlRpc: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetXmlRpc),
      map((action: GetXmlRpc) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlRpcError({error: response.error});
              }
              return new StoreGetXmlRpc({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlRpcError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateXmlRpcParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateXmlRpcParameter),
      map((action: UpdateXmlRpcParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlRpcError({error: response.error});
              }
              return new StoreUpdateXmlRpcParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlRpcError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchXmlRpcParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchXmlRpcParameter),
      map((action: SwitchXmlRpcParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlRpcError({error: response.error});
              }
              return new StoreSwitchXmlRpcParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlRpcError({error: error}));
            }),
          );
        }
      ));
  });

  AddXmlRpcParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddXmlRpcParameter),
      map((action: AddXmlRpcParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlRpcError({error: response.error});
              }
              return new StoreAddXmlRpcParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlRpcError({error: error}));
            }),
          );
        }
      ));
  });

  DelXmlRpcParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelXmlRpcParameter),
      map((action: DelXmlRpcParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotXmlRpcError({error: response.error});
              }
              return new StoreDelXmlRpcParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotXmlRpcError({error: error}));
            }),
          );
        }
      ));
  });
}

