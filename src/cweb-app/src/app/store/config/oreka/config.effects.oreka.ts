
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchOrekaParameter,
  GetOreka,
  StoreDelOrekaParameter,
  StoreSwitchOrekaParameter,
  UpdateOrekaParameter,
  StoreGetOreka,
  StoreAddOrekaParameter,
  DelOrekaParameter,
  StoreUpdateOrekaParameter,
  StoreGotOrekaError,
  AddOrekaParameter,
} from './config.actions.oreka';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsOreka {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetOreka: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOreka),
      map((action: GetOreka) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOrekaError({error: response.error});
              }
              return new StoreGetOreka({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOrekaError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOrekaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOrekaParameter),
      map((action: UpdateOrekaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOrekaError({error: response.error});
              }
              return new StoreUpdateOrekaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOrekaError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOrekaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOrekaParameter),
      map((action: SwitchOrekaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOrekaError({error: response.error});
              }
              return new StoreSwitchOrekaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOrekaError({error: error}));
            }),
          );
        }
      ));
  });

  AddOrekaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOrekaParameter),
      map((action: AddOrekaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOrekaError({error: response.error});
              }
              return new StoreAddOrekaParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOrekaError({error: error}));
            }),
          );
        }
      ));
  });

  DelOrekaParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOrekaParameter),
      map((action: DelOrekaParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOrekaError({error: response.error});
              }
              return new StoreDelOrekaParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOrekaError({error: error}));
            }),
          );
        }
      ));
  });
}

