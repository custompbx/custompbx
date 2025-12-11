import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddDirectoryParameter, AddDirectoryProfile,
  AddDirectoryProfileParameter,
  DelDirectoryParameter, DelDirectoryProfile, DelDirectoryProfileParameter,
  GetDirectory,
  GetDirectoryProfileParameters,
  StoreAddDirectoryParameter, StoreAddDirectoryProfile,
  StoreAddDirectoryProfileParameter,
  StoreDelDirectoryParameter, StoreDelDirectoryProfile, StoreDelDirectoryProfileParameter,
  StoreGetDirectory,
  StoreGetDirectoryProfileParameters,
  StoreGotDirectoryError,
  StoreSwitchDirectoryParameter, StoreSwitchDirectoryProfileParameter,
  StoreUpdateDirectoryParameter, StoreUpdateDirectoryProfile,
  StoreUpdateDirectoryProfileParameter,
  SwitchDirectoryParameter,
  SwitchDirectoryProfileParameter,
  UpdateDirectoryParameter, UpdateDirectoryProfile,
  UpdateDirectoryProfileParameter
} from './config.actions.directory';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsDirectory {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetDirectory: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetDirectory),
      map((action: GetDirectory) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreGetDirectory({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  GetDirectoryProfileParameters: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetDirectoryProfileParameters),
      map((action: GetDirectoryProfileParameters) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreGetDirectoryProfileParameters({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDirectoryParameter),
      map((action: UpdateDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchDirectoryParameter),
      map((action: SwitchDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  AddDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDirectoryParameter),
      map((action: AddDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  DelDirectoryParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDirectoryParameter),
      map((action: DelDirectoryParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreDelDirectoryParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  AddDirectoryProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDirectoryProfileParameter),
      map((action: AddDirectoryProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryProfileParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDirectoryProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDirectoryProfileParameter),
      map((action: UpdateDirectoryProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchDirectoryProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchDirectoryProfileParameter),
      map((action: SwitchDirectoryProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreSwitchDirectoryProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  DelDirectoryProfileParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDirectoryProfileParameter),
      map((action: DelDirectoryProfileParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreDelDirectoryProfileParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  AddDirectoryProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddDirectoryProfile),
      map((action: AddDirectoryProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreAddDirectoryProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  DelDirectoryProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelDirectoryProfile),
      map((action: DelDirectoryProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreDelDirectoryProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateDirectoryProfile: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateDirectoryProfile),
      map((action: UpdateDirectoryProfile) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotDirectoryError({error: response.error});
              }
              return new StoreUpdateDirectoryProfile({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotDirectoryError({error: error}));
            }),
          );
        }
      ));
  });

}

