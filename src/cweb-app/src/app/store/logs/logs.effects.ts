import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {WsDataService} from '../../services/ws-data.service';
import {catchError, map, switchMap} from 'rxjs/operators';
import {
  LogsActionTypes,
  GetLogs,
  StoreGetLogs, StoreGotLogsError,
} from './logs.actions';


@Injectable({
  providedIn: 'root'
})
export class LogsEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetLogs: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(LogsActionTypes.GetLogs),
      map((action: GetLogs) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLogsError({error: response.error});
              }
              return new StoreGetLogs({response});
            }),
            catchError((error) => {
              return of(new StoreGotLogsError({error: error}));
            }),
          );
        }
      ));
  });

}
