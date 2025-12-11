import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddOpalParameter, AddOpalListener,
  AddOpalListenerParameter,
  DelOpalParameter, DelOpalListener, DelOpalListenerParameter,
  GetOpal,
  GetOpalListenerParameters,
  StoreAddOpalParameter, StoreAddOpalListener,
  StoreAddOpalListenerParameter,
  StoreDelOpalParameter, StoreDelOpalListener, StoreDelOpalListenerParameter,
  StoreGetOpal,
  StoreGetOpalListenerParameters,
  StoreGotOpalError,
  StoreSwitchOpalParameter, StoreSwitchOpalListenerParameter,
  StoreUpdateOpalParameter, StoreUpdateOpalListener,
  StoreUpdateOpalListenerParameter,
  SwitchOpalParameter,
  SwitchOpalListenerParameter,
  UpdateOpalParameter, UpdateOpalListener,
  UpdateOpalListenerParameter
} from './config.actions.opal';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsOpal {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetOpal: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOpal),
      map((action: GetOpal) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreGetOpal({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  GetOpalListenerParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOpalListenerParameters),
      map((action: GetOpalListenerParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreGetOpalListenerParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOpalParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOpalParameter),
      map((action: UpdateOpalParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreUpdateOpalParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOpalParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOpalParameter),
      map((action: SwitchOpalParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreSwitchOpalParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  AddOpalParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOpalParameter),
      map((action: AddOpalParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreAddOpalParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  DelOpalParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOpalParameter),
      map((action: DelOpalParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreDelOpalParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  AddOpalListenerParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOpalListenerParameter),
      map((action: AddOpalListenerParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreAddOpalListenerParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOpalListenerParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOpalListenerParameter),
      map((action: UpdateOpalListenerParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreUpdateOpalListenerParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOpalListenerParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOpalListenerParameter),
      map((action: SwitchOpalListenerParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreSwitchOpalListenerParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  DelOpalListenerParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOpalListenerParameter),
      map((action: DelOpalListenerParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreDelOpalListenerParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  AddOpalListener: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOpalListener),
      map((action: AddOpalListener) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreAddOpalListener({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  DelOpalListener: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOpalListener),
      map((action: DelOpalListener) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreDelOpalListener({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOpalListener: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOpalListener),
      map((action: UpdateOpalListener) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOpalError({error: response.error});
              }
              return new StoreUpdateOpalListener({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOpalError({error: error}));
            }),
          );
        }
      ));
  });

}

