import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {WsDataService} from '../../services/ws-data.service';
import {catchError, map, switchMap} from 'rxjs/operators';
import {
  HEPActionTypes,
  GetHEP,
  StoreGetHEP, StoreGotHEPError, GetHEPDetails, StoreGetHEPDetails,
} from './hep.actions';


@Injectable()
export class HepEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetHEP: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(HEPActionTypes.GetHEP),
      map((action: GetHEP) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHEPError({error: response.error});
              }
              return new StoreGetHEP({response});
            }),
            catchError((error) => {
              return of(new StoreGotHEPError({error: error}));
            }),
          );
        }
      ));
  });

  GetHEPDetails: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(HEPActionTypes.GetHEPDetails),
      map((action: GetHEPDetails) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotHEPError({error: response.error});
              }
              return new StoreGetHEPDetails({response});
            }),
            catchError((error) => {
              return of(new StoreGotHEPError({error: error}));
            }),
          );
        }
      ));
  });

}
