import {createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {Action} from '@ngrx/store';
import {WsDataService} from "../ws-data.service";

export function createEffectForActions(
  actions,
  ws: WsDataService,
  action1: Action,
  action2: (payload: any) => Action,
  action3: (payload: any) => Action,
  payloadType?: string
): Observable<any> {
  return createEffect(() =>
    actions.pipe(
      ofType(action1.type),
      switchMap((action: any) => {
        console.log(action.payload)
        return ws.universalSender(action.type, action.payload).pipe(
          map((response: any) => {
            if (response.error) {
              return action3({error: response.error});
            }
            switch (payloadType) {
              case 'index':
                return action2({response: response, index: action.payload.index});
              case 'id':
                return action2({id: action.payload.id, response});
              default:
                return action2({response});
            }
          }),
          catchError((error) => {
            console.log(error);
            return of(action3({error}));
          })
        );
      })
    )
  );
}
