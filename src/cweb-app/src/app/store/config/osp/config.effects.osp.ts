import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddOspParameter, AddOspProfile,
  AddOspProfileParameter,
  DelOspParameter, DelOspProfile, DelOspProfileParameter,
  GetOsp,
  GetOspProfileParameters,
  StoreAddOspParameter, StoreAddOspProfile,
  StoreAddOspProfileParameter,
  StoreDelOspParameter, StoreDelOspProfile, StoreDelOspProfileParameter,
  StoreGetOsp,
  StoreGetOspProfileParameters,
  StoreGotOspError,
  StoreSwitchOspParameter, StoreSwitchOspProfileParameter,
  StoreUpdateOspParameter, StoreUpdateOspProfile,
  StoreUpdateOspProfileParameter,
  SwitchOspParameter,
  SwitchOspProfileParameter,
  UpdateOspParameter, UpdateOspProfile,
  UpdateOspProfileParameter
} from './config.actions.osp';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsOsp {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetOsp: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOsp),
      map((action: GetOsp) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreGetOsp({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  GetOspProfileParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetOspProfileParameters),
      map((action: GetOspProfileParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreGetOspProfileParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOspParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOspParameter),
      map((action: UpdateOspParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreUpdateOspParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOspParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOspParameter),
      map((action: SwitchOspParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreSwitchOspParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  AddOspParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOspParameter),
      map((action: AddOspParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreAddOspParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  DelOspParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOspParameter),
      map((action: DelOspParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreDelOspParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  AddOspProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOspProfileParameter),
      map((action: AddOspProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreAddOspProfileParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOspProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOspProfileParameter),
      map((action: UpdateOspProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreUpdateOspProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchOspProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchOspProfileParameter),
      map((action: SwitchOspProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreSwitchOspProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  DelOspProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOspProfileParameter),
      map((action: DelOspProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreDelOspProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  AddOspProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddOspProfile),
      map((action: AddOspProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreAddOspProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  DelOspProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelOspProfile),
      map((action: DelOspProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreDelOspProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateOspProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateOspProfile),
      map((action: UpdateOspProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotOspError({error: response.error});
              }
              return new StoreUpdateOspProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotOspError({error: error}));
            }),
          );
        }
      ));
  });

}

