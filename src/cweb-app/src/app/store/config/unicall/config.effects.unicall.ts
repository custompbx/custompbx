import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddUnicallParameter, AddUnicallSpan,
  AddUnicallSpanParameter,
  DelUnicallParameter, DelUnicallSpan, DelUnicallSpanParameter,
  GetUnicall,
  GetUnicallSpanParameters,
  StoreAddUnicallParameter, StoreAddUnicallSpan,
  StoreAddUnicallSpanParameter,
  StoreDelUnicallParameter, StoreDelUnicallSpan, StoreDelUnicallSpanParameter,
  StoreGetUnicall,
  StoreGetUnicallSpanParameters,
  StoreGotUnicallError,
  StoreSwitchUnicallParameter, StoreSwitchUnicallSpanParameter,
  StoreUpdateUnicallParameter, StoreUpdateUnicallSpan,
  StoreUpdateUnicallSpanParameter,
  SwitchUnicallParameter,
  SwitchUnicallSpanParameter,
  UpdateUnicallParameter, UpdateUnicallSpan,
  UpdateUnicallSpanParameter
} from './config.actions.unicall';

@Injectable()
export class ConfigEffectsUnicall {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetUnicall: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetUnicall),
      map((action: GetUnicall) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreGetUnicall({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  GetUnicallSpanParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetUnicallSpanParameters),
      map((action: GetUnicallSpanParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreGetUnicallSpanParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUnicallParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateUnicallParameter),
      map((action: UpdateUnicallParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreUpdateUnicallParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUnicallParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchUnicallParameter),
      map((action: SwitchUnicallParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreSwitchUnicallParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  AddUnicallParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddUnicallParameter),
      map((action: AddUnicallParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreAddUnicallParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  DelUnicallParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelUnicallParameter),
      map((action: DelUnicallParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreDelUnicallParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  AddUnicallSpanParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddUnicallSpanParameter),
      map((action: AddUnicallSpanParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreAddUnicallSpanParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUnicallSpanParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateUnicallSpanParameter),
      map((action: UpdateUnicallSpanParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreUpdateUnicallSpanParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchUnicallSpanParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchUnicallSpanParameter),
      map((action: SwitchUnicallSpanParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreSwitchUnicallSpanParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  DelUnicallSpanParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelUnicallSpanParameter),
      map((action: DelUnicallSpanParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreDelUnicallSpanParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  AddUnicallSpan: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddUnicallSpan),
      map((action: AddUnicallSpan) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreAddUnicallSpan({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  DelUnicallSpan: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelUnicallSpan),
      map((action: DelUnicallSpan) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreDelUnicallSpan({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateUnicallSpan: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateUnicallSpan),
      map((action: UpdateUnicallSpan) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotUnicallError({error: response.error});
              }
              return new StoreUpdateUnicallSpan({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotUnicallError({error: error}));
            }),
          );
        }
      ));
  });

}

