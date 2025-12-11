import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddLcrParameter, AddLcrProfile,
  AddLcrProfileParameter,
  DelLcrParameter, DelLcrProfile, DelLcrProfileParameter,
  GetLcr,
  GetLcrProfileParameters,
  StoreAddLcrParameter, StoreAddLcrProfile,
  StoreAddLcrProfileParameter,
  StoreDelLcrParameter, StoreDelLcrProfile, StoreDelLcrProfileParameter,
  StoreGetLcr,
  StoreGetLcrProfileParameters,
  StoreGotLcrError,
  StoreSwitchLcrParameter, StoreSwitchLcrProfileParameter,
  StoreUpdateLcrParameter, StoreUpdateLcrProfile,
  StoreUpdateLcrProfileParameter,
  SwitchLcrParameter,
  SwitchLcrProfileParameter,
  UpdateLcrParameter, UpdateLcrProfile,
  UpdateLcrProfileParameter
} from './config.actions.lcr';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsLcr {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetLcr: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetLcr),
      map((action: GetLcr) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreGetLcr({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  GetLcrProfileParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetLcrProfileParameters),
      map((action: GetLcrProfileParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreGetLcrProfileParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateLcrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateLcrParameter),
      map((action: UpdateLcrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreUpdateLcrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchLcrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchLcrParameter),
      map((action: SwitchLcrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreSwitchLcrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  AddLcrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddLcrParameter),
      map((action: AddLcrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreAddLcrParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  DelLcrParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelLcrParameter),
      map((action: DelLcrParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreDelLcrParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  AddLcrProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddLcrProfileParameter),
      map((action: AddLcrProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreAddLcrProfileParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateLcrProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateLcrProfileParameter),
      map((action: UpdateLcrProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreUpdateLcrProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchLcrProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchLcrProfileParameter),
      map((action: SwitchLcrProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreSwitchLcrProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  DelLcrProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelLcrProfileParameter),
      map((action: DelLcrProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreDelLcrProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  AddLcrProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddLcrProfile),
      map((action: AddLcrProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreAddLcrProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  DelLcrProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelLcrProfile),
      map((action: DelLcrProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreDelLcrProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateLcrProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateLcrProfile),
      map((action: UpdateLcrProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotLcrError({error: response.error});
              }
              return new StoreUpdateLcrProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotLcrError({error: error}));
            }),
          );
        }
      ));
  });

}

