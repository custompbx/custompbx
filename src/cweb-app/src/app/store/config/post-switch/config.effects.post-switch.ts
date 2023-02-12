import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddPostSwitchCliKeybinding, AddPostSwitchDefaultPtime,
  AddPostSwitchParameter,
  DelPostSwitchCliKeybinding, DelPostSwitchDefaultPtime,
  DelPostSwitchParameter,
  GetPostSwitch,
  StoreAddPostSwitchCliKeybinding, StoreAddPostSwitchDefaultPtime,
  StoreAddPostSwitchParameter,
  StoreDelPostSwitchCliKeybinding, StoreDelPostSwitchDefaultPtime,
  StoreDelPostSwitchParameter,
  StoreGetPostSwitch,
  StoreGotPostSwitchError,
  StoreSwitchPostSwitchCliKeybinding, StoreSwitchPostSwitchDefaultPtime,
  StoreSwitchPostSwitchParameter,
  StoreUpdatePostSwitchCliKeybinding, StoreUpdatePostSwitchDefaultPtime,
  StoreUpdatePostSwitchParameter,
  SwitchPostSwitchCliKeybinding, SwitchPostSwitchDefaultPtime,
  SwitchPostSwitchParameter,
  UpdatePostSwitchCliKeybinding,
  UpdatePostSwitchDefaultPtime,
  UpdatePostSwitchParameter
} from './config.actions.post-switch';
import {Failure} from '../config.actions';

@Injectable()
export class ConfigEffectsPostSwitch {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetPostSwitch: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetPostSwitch),
      map((action: GetPostSwitch) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreGetPostSwitch({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePostSwitchParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePostSwitchParameter),
      map((action: UpdatePostSwitchParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreUpdatePostSwitchParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPostSwitchParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPostSwitchParameter),
      map((action: SwitchPostSwitchParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreSwitchPostSwitchParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddPostSwitchParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPostSwitchParameter),
      map((action: AddPostSwitchParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreAddPostSwitchParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelPostSwitchParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPostSwitchParameter),
      map((action: DelPostSwitchParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreDelPostSwitchParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePostSwitchCliKeybinding: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePostSwitchCliKeybinding),
      map((action: UpdatePostSwitchCliKeybinding) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreUpdatePostSwitchCliKeybinding({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPostSwitchCliKeybinding: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPostSwitchCliKeybinding),
      map((action: SwitchPostSwitchCliKeybinding) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreSwitchPostSwitchCliKeybinding({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddPostSwitchCliKeybinding: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPostSwitchCliKeybinding),
      map((action: AddPostSwitchCliKeybinding) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreAddPostSwitchCliKeybinding({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelPostSwitchCliKeybinding: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPostSwitchCliKeybinding),
      map((action: DelPostSwitchCliKeybinding) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreDelPostSwitchCliKeybinding({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePostSwitchDefaultPtime: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePostSwitchDefaultPtime),
      map((action: UpdatePostSwitchDefaultPtime) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreUpdatePostSwitchDefaultPtime({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPostSwitchDefaultPtime: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPostSwitchDefaultPtime),
      map((action: SwitchPostSwitchDefaultPtime) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreSwitchPostSwitchDefaultPtime({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  AddPostSwitchDefaultPtime: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPostSwitchDefaultPtime),
      map((action: AddPostSwitchDefaultPtime) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreAddPostSwitchDefaultPtime({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

  DelPostSwitchDefaultPtime: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPostSwitchDefaultPtime),
      map((action: DelPostSwitchDefaultPtime) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPostSwitchError({error: response.error});
              }
              return new StoreDelPostSwitchDefaultPtime({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new Failure({error: error}));
            }),
          );
        }
      ));
  });

}

