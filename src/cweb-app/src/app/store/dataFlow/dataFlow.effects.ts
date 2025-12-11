import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {WsDataService} from '../../services/ws-data.service';
import {catchError, concatMap, map, switchMap} from 'rxjs/operators';
import {
  DataFlowActionTypes,
  GetDashboard,
  StoreGetDashboard,
  ReduceLoadCounter,
  Failure,
  UnSubscribe, SubscriptionList, PersistentSubscription
} from './dataFlow.actions';


@Injectable({
  providedIn: 'root'
})
export class DataFlowEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetDashboard: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DataFlowActionTypes.GET_DASHBOARD),
      map((action: GetDashboard) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            concatMap((response) => [
              new StoreGetDashboard({response}),
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

  UnSubscribe: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DataFlowActionTypes.UnSubscribe),
      map((action: UnSubscribe) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  }, { dispatch: false });

  SubscriptionList: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DataFlowActionTypes.SubscriptionList),
      map((action: SubscriptionList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  }, { dispatch: false });

  PersistentSubscription: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(DataFlowActionTypes.PersistentSubscription),
      map((action: SubscriptionList) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  }, { dispatch: false });

}
