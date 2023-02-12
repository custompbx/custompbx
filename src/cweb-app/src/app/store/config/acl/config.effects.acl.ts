import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddAclList,
  AddAclNode,
  DelAclList,
  DelAclNode,
  DropAclList,
  DropAclNode,
  GetAclLists,
  GetAclNodes,
  StoreAclList,
  StoreAclLists,
  StoreAclNode,
  StoreAclNodes,
  StoreGotAclError,
  StoreSwitchAclNode,
  StoreUpdatadAclList,
  StoreUpdatadAclListDefault,
  StoreUpdatadAclNode,
  UpdateAclList,
  UpdateAclNode, StoreMoveAclListNode, MoveAclListNode,
} from './config.actions.acl';

@Injectable()
export class ConfigEffectsAcl {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAclLists: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_ACL_LISTS),
      map((action: GetAclLists) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreAclLists({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  AddAclList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_ACL_LIST),
      map((action: AddAclList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreAclList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAclList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_ACL_LIST),
      map((action: UpdateAclList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreUpdatadAclList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAclListDefault: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_ACL_LIST_DEFAULT),
      map((action: UpdateAclList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreUpdatadAclListDefault({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  DelAclList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_ACL_LIST),
      map((action: DelAclList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new DropAclList({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  GetAclNodes: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GET_ACL_NODES),
      map((action: GetAclNodes) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, {id: action.payload}).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreAclNodes({id: action.payload, response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  AddAclNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ADD_ACL_NODE),
      map((action: AddAclNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreAclNode({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAclNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UPDATE_ACL_NODE),
      map((action: UpdateAclNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreUpdatadAclNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  DelAclNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DEL_ACL_NODE),
      map((action: DelAclNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new DropAclNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAclNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SWITCH_ACL_NODE),
      map((action: DelAclNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreSwitchAclNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });

  MoveAclListNode: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.MoveAclListNode),
      map((action: MoveAclListNode) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotAclError({error: response.error});
              }
              return new StoreMoveAclListNode({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotAclError({error: error}));
            }),
          );
        }
      ));
  });
}
