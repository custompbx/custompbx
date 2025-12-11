import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  AddDistributorList,
  AddDistributorNode,
  ConfigActionTypes,
  DelDistributorList,
  DelDistributorNode,
  GetDistributorConfig,
  GetDistributorNodes,
  StoreAddDistributorList,
  StoreAddDistributorNode,
  StoreDelDistributorList,
  StoreDelDistributorNode,
  StoreGetDistributorConfig,
  StoreGetDistributorNodes,
  StoreGotDistributorError,
  StoreSwitchDistributorNode,
  StoreUpdateDistributorList,
  StoreUpdateDistributorNode,
  SwitchDistributorNode, UpdateDistributorList,
  UpdateDistributorNode
} from './config.actions.distributor';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsDistributor {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetDistributorConfig: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetDistributorConfig),
      map((action: GetDistributorConfig) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreGetDistributorConfig({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  AddDistributorList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDistributorList),
      map((action: AddDistributorList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreAddDistributorList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  DelDistributorList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDistributorList),
      map((action: DelDistributorList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreDelDistributorList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDistributorList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDistributorList),
      map((action: UpdateDistributorList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreUpdateDistributorList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  GetDistributorNodes: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetDistributorNodes),
      map((action: GetDistributorNodes) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreGetDistributorNodes({response: response, id: action.payload.id});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  AddDistributorNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDistributorNode),
      map((action: AddDistributorNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreAddDistributorNode({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDistributorNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDistributorNode),
      map((action: UpdateDistributorNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreUpdateDistributorNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  DelDistributorNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDistributorNode),
      map((action: DelDistributorNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreDelDistributorNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDistributorNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchDistributorNode),
      map((action: SwitchDistributorNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDistributorError({error: response.error});
              }
              return new StoreSwitchDistributorNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDistributorError({error: error}));
            }),
          );
        }
      ));
  });

}
