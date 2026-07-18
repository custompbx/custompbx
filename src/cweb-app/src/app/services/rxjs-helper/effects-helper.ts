import {createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {Action} from '@ngrx/store';
import {WsDataService} from '../ws-data.service';
import {withOperationFeedback} from '../operation-feedback';

export function createEffectForActions(
  actions,
  ws: WsDataService,
  action1: Action,
  action2: (payload: any) => Action,
  action3: (payload: any) => Action,
  payloadType?: string,
  operationFeedback = true,
): Observable<any> {
  return createEffect(() =>
    actions.pipe(
      ofType(action1.type),
      switchMap((action: any) => {
        return ws.universalSender(action.type, action.payload).pipe(
          map((response: any) => {
            if (response.error) {
              return action3({error: response.error});
            }
            let completedAction: Action;
            switch (payloadType) {
              case 'index':
                completedAction = action2({response: response, index: action.payload.index});
                break;
              case 'id':
                completedAction = action2({id: action.payload.id, response});
                break;
              default:
                completedAction = action2({response, payload: action.payload});
            }
            return operationFeedback
              ? withOperationFeedback(completedAction, action.type)
              : {...completedAction, operationFeedback: false};
          }),
          catchError((error) => {
            const errorMessage = error instanceof Error ? error.message : error;
            return of(action3({error: errorMessage}));
          })
        );
      })
    )
  );
}
