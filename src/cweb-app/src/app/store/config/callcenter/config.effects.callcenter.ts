import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddCallcenterAgent,
  AddCallcenterQueue,
  AddCallcenterQueueParam,
  AddCallcenterSettings, AddCallcenterTier,
  DelCallcenterAgent, DelCallcenterMember,
  DelCallcenterQueue,
  DelCallcenterQueueParam,
  DelCallcenterSettings, DelCallcenterTier,
  GetCallcenterAgents, GetCallcenterMembers,
  GetCallcenterQueues,
  GetCallcenterQueuesParams,
  GetCallcenterSettings, GetCallcenterTiers,
  ImportCallcenterAgentsAndTiers,
  RenameCallcenterQueue, SendCallcenterCommand,
  StoreAddCallcenterAgent,
  StoreAddCallcenterQueue,
  StoreAddCallcenterQueueParam,
  StoreAddCallcenterSettings, StoreAddCallcenterTier,
  StoreDelCallcenterAgent, StoreDelCallcenterMember,
  StoreDelCallcenterQueue,
  StoreDelCallcenterQueueParam,
  StoreDelCallcenterSettings, StoreDelCallcenterTier,
  StoreGetCallcenterAgents, StoreGetCallcenterMembers,
  StoreGetCallcenterQueues,
  StoreGetCallcenterQueuesParams,
  StoreGetCallcenterSettings, StoreGetCallcenterTiers,
  StoreGotCallcenterError, StoreImportCallcenterAgentsAndTiers,
  StoreRenameCallcenterQueue, StoreSendCallcenterCommand, StoreSubscribeCallcenterAgents, StoreSubscribeCallcenterTiers,
  StoreSwitchCallcenterQueueParam,
  StoreSwitchCallcenterSettings,
  StoreUpdateCallcenterAgent,
  StoreUpdateCallcenterQueueParam,
  StoreUpdateCallcenterSettings, StoreUpdateCallcenterTier, SubscribeCallcenterAgents, SubscribeCallcenterTiers,
  SwitchCallcenterQueueParam,
  SwitchCallcenterSettings,
  UpdateCallcenterAgent,
  UpdateCallcenterQueueParam,
  UpdateCallcenterSettings, UpdateCallcenterTier
} from './config.actions.callcenter';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsCallcenter {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCallcenterSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCallcenterSettings),
      map((action: GetCallcenterSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreGetCallcenterSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCallcenterSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCallcenterSettings),
      map((action: UpdateCallcenterSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreUpdateCallcenterSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCallcenterSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchCallcenterSettings),
      map((action: SwitchCallcenterSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreSwitchCallcenterSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  AddCallcenterSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCallcenterSettings),
      map((action: AddCallcenterSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreAddCallcenterSettings({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  DelCallcenterSettings: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCallcenterSettings),
      map((action: DelCallcenterSettings) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreDelCallcenterSettings({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  GetCallcenterQueues: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCallcenterQueues),
      map((action: GetCallcenterQueues) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreGetCallcenterQueues({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  GetCallcenterQueuesParams: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCallcenterQueuesParams),
      map((action: GetCallcenterQueuesParams) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreGetCallcenterQueuesParams({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCallcenterQueueParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCallcenterQueueParam),
      map((action: UpdateCallcenterQueueParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreUpdateCallcenterQueueParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchCallcenterQueueParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchCallcenterQueueParam),
      map((action: SwitchCallcenterQueueParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreSwitchCallcenterQueueParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  AddCallcenterQueueParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCallcenterQueueParam),
      map((action: AddCallcenterQueueParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreAddCallcenterQueueParam({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  DelCallcenterQueueParam: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCallcenterQueueParam),
      map((action: DelCallcenterQueueParam) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreDelCallcenterQueueParam({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  AddCallcenterQueue: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCallcenterQueue),
      map((action: AddCallcenterQueue) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreAddCallcenterQueue({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  DelCallcenterQueue: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCallcenterQueue),
      map((action: DelCallcenterQueue) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreDelCallcenterQueue({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  RenameCallcenterQueue: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.RenameCallcenterQueue),
      map((action: RenameCallcenterQueue) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreRenameCallcenterQueue({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  GetCallcenterAgents: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCallcenterAgents),
      map((action: GetCallcenterAgents) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreGetCallcenterAgents({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCallcenterAgent: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCallcenterAgent),
      map((action: UpdateCallcenterAgent) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreUpdateCallcenterAgent({response, payload: action.payload});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  AddCallcenterAgent: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCallcenterAgent),
      map((action: AddCallcenterAgent) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreAddCallcenterAgent({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  DelCallcenterAgent: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCallcenterAgent),
      map((action: DelCallcenterAgent) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreDelCallcenterAgent({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  ImportCallcenterAgentsAndTiers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.ImportCallcenterAgentsAndTiers),
      map((action: ImportCallcenterAgentsAndTiers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreImportCallcenterAgentsAndTiers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  GetCallcenterTiers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCallcenterTiers),
      map((action: GetCallcenterTiers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreGetCallcenterTiers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateCallcenterTier: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateCallcenterTier),
      map((action: UpdateCallcenterTier) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreUpdateCallcenterTier({response, payload: action.payload});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  AddCallcenterTier: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddCallcenterTier),
      map((action: AddCallcenterTier) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreAddCallcenterTier({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  DelCallcenterTier: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCallcenterTier),
      map((action: DelCallcenterTier) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreDelCallcenterTier({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  SendCallcenterCommand: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SendCallcenterCommand),
      map((action: SendCallcenterCommand) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreSendCallcenterCommand({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  SubscribeCallcenterAgents: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SubscribeCallcenterAgents),
      map((action: SubscribeCallcenterAgents) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreSubscribeCallcenterAgents({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  SubscribeCallcenterTiers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SubscribeCallcenterTiers),
      map((action: SubscribeCallcenterTiers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreSubscribeCallcenterTiers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  GetCallcenterMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetCallcenterMembers),
      map((action: GetCallcenterMembers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreGetCallcenterMembers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

  DelCallcenterMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelCallcenterMember),
      map((action: DelCallcenterMember) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCallcenterError({error: response.error});
              }
              return new StoreDelCallcenterMember({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotCallcenterError({error: error}));
            }),
          );
        }
      ));
  });

}

