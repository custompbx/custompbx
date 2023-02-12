import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {WsDataService} from '../../services/ws-data.service';
import {catchError, map, switchMap} from 'rxjs/operators';
import {
  CDRActionTypes,
  GetCDR,
  StoreGetCDR, StoreGotCdrError,
} from './cdr.actions';


@Injectable()
export class CdrEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetCDR: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(CDRActionTypes.GET_CDR),
      map((action: GetCDR) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotCdrError({error: response.error});
              }
              return new StoreGetCDR({response});
            }),
            catchError((error) => {
              return of(new StoreGotCdrError({error: error}));
            }),
          );
        }
      ));
  });

}
