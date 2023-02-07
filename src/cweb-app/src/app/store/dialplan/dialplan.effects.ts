import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  AddAction,
  AddAntiaction,
  AddCondition,
  AddContext,
  AddExtension,
  DeleteAction,
  DeleteAntiaction,
  DeleteCondition,
  DeleteContext,
  DeleteExtension,
  DialplanActionTypes,
  Failure,
  GetExtensionDetails,
  GetConditions,
  GetContexts,
  GetExtensions,
  MoveAction,
  MoveAntiaction,
  MoveCondition,
  MoveExtension,
  RenameContext,
  RenameExtension,
  StoreAddAction,
  StoreAddAntiaction,
  StoreAddCondition,
  StoreAddContext,
  StoreAddExtension,
  StoreDeleteAction,
  StoreDeleteAntiaction,
  StoreDeleteCondition,
  StoreDeleteContext,
  StoreDeleteExtension,
  StoreGetExtensionDetails,
  StoreGetConditions,
  StoreGetContexts,
  StoreGetExtensions,
  StoreMoveAction,
  StoreMoveAntiaction,
  StoreMoveCondition,
  StoreMoveExtension,
  StoreRenameContext,
  StoreRenameExtension,
  StoreSwitchAction,
  StoreSwitchAntiaction,
  StoreSwitchCondition,
  StoreSwitchContext,
  StoreSwitchExtension,
  StoreUpdateAction,
  StoreUpdateAntiaction,
  StoreUpdateCondition,
  StoreUpdateExtension,
  SwitchAction,
  SwitchAntiaction,
  SwitchCondition,
  SwitchContext,
  SwitchExtension,
  UpdateAction,
  UpdateAntiaction,
  UpdateCondition,
  UpdateExtension,
  AddRegex,
  StoreAddRegex,
  DeleteRegex,
  StoreDeleteRegex,
  SwitchRegex,
  StoreSwitchRegex,
  DeleteNewAction,
  DeleteNewAntiaction,
  DeleteNewRegex,
  SwitchExtensionContinue,
  StoreSwitchExtensionContinue,
  ReduceLoadCounter,
  SwitchDialplanDebug,
  StoreSwitchDialplanDebug,
  DialplanDebug,
  StoreDialplanDebug,
  DialplanSettings,
  StoreDialplanSettings,
  StoreSwitchDialplanStatic, SwitchDialplanStatic,
} from './dialplan.actions';
import {catchError, concatMap, map, mergeMap, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../services/ws-data.service';

@Injectable()
export class DialplanEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  ImportDialplan: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.IMPORT_DIALPLAN),
      map((action: GetContexts) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            concatMap((response) => [
              new GetContexts(null),
              new ReduceLoadCounter(),
            ]),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetContexts: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.GET_CONTEXTS),
      map((action: GetContexts) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreGetContexts({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddContext: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.ADD_CONTEXT),
      map((action: AddContext) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreAddContext({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameContext: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.RENAME_CONTEXT),
      map((action: RenameContext) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreRenameContext({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteContext: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DELETE_CONTEXT),
      map((action: DeleteContext) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDeleteContext({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchContext: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_CONTEXT),
      map((action: SwitchContext) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchContext({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetConditions: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.GET_CONDITIONS),
      map((action: GetConditions) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreGetConditions({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddCondition: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.ADD_CONDITION),
      map((action: AddCondition) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreAddCondition({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteCondition: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DELETE_CONDITION),
      map((action: DeleteCondition) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDeleteCondition({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCondition: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_CONDITION),
      map((action: SwitchCondition) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchCondition({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetExtensions: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.GET_EXTENSIONS),
      map((action: GetExtensions) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreGetExtensions({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddExtension: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.ADD_EXTENSION),
      map((action: AddExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreAddExtension({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  RenameExtension: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.RENAME_EXTENSION),
      map((action: RenameExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreRenameExtension({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteExtension: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DELETE_EXTENSION),
      map((action: DeleteExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDeleteExtension({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchExtension: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_EXTENSION),
      map((action: SwitchExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchExtension({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchExtensionContinue: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_EXTENSION_CONTINUE),
      map((action: SwitchExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchExtensionContinue({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  GetActions: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.GET_EXTENSION_DETAILS),
      map((action: GetExtensionDetails) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreGetExtensionDetails({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddAction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.ADD_ACTION),
      map((action: AddAction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, {id: action.payload.id, action: action.payload.action}).pipe(
            mergeMap(response => [
              new StoreAddAction({response: response}),
              new DeleteNewAction({
                contextId: action.payload.contextId,
                extensionId: action.payload.extensionId,
                conditionId: action.payload.id,
                index: action.payload.index,
              }),
            ]),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteAction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DELETE_ACTION),
      map((action: DeleteAction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDeleteAction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_ACTION),
      map((action: SwitchAction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchAction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddAntiaction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.ADD_ANTIACTION),
      map((action: AddAntiaction) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, {id: action.payload.id, antiaction: action.payload.antiaction}).pipe(
          mergeMap(response => [
            new StoreAddAntiaction({response: response}),
            new DeleteNewAntiaction({
              contextId: action.payload.contextId,
              extensionId: action.payload.extensionId,
              conditionId: action.payload.id,
              index: action.payload.index,
            }),
          ]),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteAntiaction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DELETE_ANTIACTION),
      map((action: DeleteAntiaction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDeleteAntiaction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchAntiaction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_ANTIACTION),
      map((action: SwitchAntiaction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchAntiaction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.UPDATE_ACTION),
      map((action: UpdateAction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreUpdateAction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateAntiaction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.UPDATE_ANTIACTION),
      map((action: UpdateAntiaction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreUpdateAntiaction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateExtension: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.UPDATE_EXTENSION),
      map((action: UpdateExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreUpdateExtension({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  MoveExtension: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.MOVE_EXTENSION),
      map((action: MoveExtension) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreMoveExtension({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  MoveCondition: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.MOVE_CONDITION),
      map((action: MoveCondition) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreMoveCondition({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  MoveAction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.MOVE_ACTION),
      map((action: MoveAction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreMoveAction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  MoveAntiaction: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.MOVE_ANTIACTION),
      map((action: MoveAntiaction) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreMoveAntiaction({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCondition: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.UPDATE_CONDITION),
      map((action: UpdateCondition) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreUpdateCondition({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddRegex: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.ADD_REGEX),
      map((action: AddRegex) => action),
      switchMap(action => {
        return this.ws.universalSender(action.type, {id: action.payload.id, action: action.payload.action}).pipe(
          mergeMap(response => [
            new StoreAddRegex({response: response}),
            new DeleteNewRegex({
              contextId: action.payload.contextId,
              extensionId: action.payload.extensionId,
              conditionId: action.payload.id,
              index: action.payload.index,
            }),
          ]),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DeleteRegex: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DELETE_REGEX),
      map((action: DeleteRegex) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDeleteRegex({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchRegex: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_EXTENSION),
      map((action: SwitchRegex) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchRegex({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DialplanDebug: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DIALPLAN_DEBUG),
      map((action: DialplanDebug) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDialplanDebug({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDialplanDebug: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_DIALPLAN_DEBUG),
      map((action: SwitchDialplanDebug) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchDialplanDebug({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DialplanSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.DIALPLAN_SETTINGS),
      map((action: DialplanSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreDialplanSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDialplanStatic: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DialplanActionTypes.SWITCH_DIALPLAN_STATIC),
      map((action: SwitchDialplanStatic) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSwitchDialplanStatic({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

}
