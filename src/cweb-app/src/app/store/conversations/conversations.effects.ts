import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {WsDataService} from '../../services/ws-data.service';
import {createEffectForActions} from '../../services/rxjs-helper/effects-helper';

import {
  GetConversationPrivateMessages,
  GetNewConversationMessage,
  SendConversationPrivateMessage,
  StoreConversationError,
  StoreGetConversationPrivateMessages,
  StoreGetNewConversationMessage,
  StoreSendConversationPrivateMessage,
  SendConversationPrivateCall,
  StoreSendConversationPrivateCall, StoreGetConversationPrivateCalls, GetConversationPrivateCalls,
} from './conversations.actions';
import {catchError, map, switchMap} from 'rxjs/operators';

@Injectable()
export class ConversationsEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetConversationMessages: Observable<any> = createEffectForActions(this.actions, this.ws,
    GetConversationPrivateMessages, StoreGetConversationPrivateMessages, StoreConversationError);
  SendConversationMessage: Observable<any> = createEffectForActions(this.actions, this.ws,
    SendConversationPrivateMessage, StoreSendConversationPrivateMessage, StoreConversationError);
  SendConversationPrivateCall: Observable<any> = createEffectForActions(this.actions, this.ws,
    SendConversationPrivateCall, StoreSendConversationPrivateCall, StoreConversationError);
  GetConversationPrivateCalls: Observable<any> = createEffectForActions(this.actions, this.ws,
    GetConversationPrivateCalls, StoreGetConversationPrivateCalls, StoreConversationError);

  GetNewConversationMessage: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(GetNewConversationMessage.type),
      map((action) => action),
      switchMap(action => {
          return this.ws.proceedMessageType(action.type).pipe(
            map((response) => {
              return StoreGetNewConversationMessage({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(StoreConversationError({error: error}));
            }),
          );
        }
      ));
  });
}
