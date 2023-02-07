import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../services/ws-data.service';
import {
  GetInstances, ConfigActionTypes, StoreGetInstances, StoreGotInstancesError, UpdateInstanceDescription, StoreUpdateInstanceDescription,
} from './instances.actions';

@Injectable()
export class EffectsInstances {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetInstances: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetInstances),
      map((action: GetInstances) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotInstancesError({error: response.error});
              }
              return new StoreGetInstances({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotInstancesError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateInstanceDescription: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateInstanceDescription),
      map((action: UpdateInstanceDescription) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotInstancesError({error: response.error});
              }
              return new StoreUpdateInstanceDescription({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotInstancesError({error: error}));
            }),
          );
        }
      ));
  });

}

