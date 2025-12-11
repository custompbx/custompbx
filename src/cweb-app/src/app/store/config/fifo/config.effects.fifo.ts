import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';
import {
  ConfigActionTypes,
  AddFifoParameter, AddFifoFifo,
  AddFifoFifoMember,
  DelFifoParameter, DelFifoFifo, DelFifoFifoMember,
  GetFifo,
  GetFifoFifoMembers,
  StoreAddFifoParameter, StoreAddFifoFifo,
  StoreAddFifoFifoMember,
  StoreDelFifoParameter, StoreDelFifoFifo, StoreDelFifoFifoMember,
  StoreGetFifo,
  StoreGetFifoFifoMembers,
  StoreGotFifoError,
  StoreSwitchFifoParameter, StoreSwitchFifoFifoMember,
  StoreUpdateFifoParameter, StoreUpdateFifoFifo,
  StoreUpdateFifoFifoMember,
  SwitchFifoParameter,
  SwitchFifoFifoMember,
  UpdateFifoParameter, UpdateFifoFifo,
  UpdateFifoFifoMember, UpdateFifoFifoImportance, StoreUpdateFifoFifoImportance
} from './config.actions.fifo';

@Injectable({
  providedIn: 'root'
})
export class ConfigEffectsFifo {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  UpdateFifoFifoImportance: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateFifoFifoImportance),
      map((action: UpdateFifoFifoImportance) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreUpdateFifoFifoImportance({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  GetFifo: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetFifo),
      map((action: GetFifo) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreGetFifo({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  GetFifoFifoMembers: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetFifoFifoMembers),
      map((action: GetFifoFifoMembers) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreGetFifoFifoMembers({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateFifoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateFifoParameter),
      map((action: UpdateFifoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreUpdateFifoParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchFifoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchFifoParameter),
      map((action: SwitchFifoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreSwitchFifoParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  AddFifoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddFifoParameter),
      map((action: AddFifoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreAddFifoParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  DelFifoParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelFifoParameter),
      map((action: DelFifoParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreDelFifoParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  AddFifoFifoMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddFifoFifoMember),
      map((action: AddFifoFifoMember) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreAddFifoFifoMember({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateFifoFifoMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateFifoFifoMember),
      map((action: UpdateFifoFifoMember) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreUpdateFifoFifoMember({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchFifoFifoMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchFifoFifoMember),
      map((action: SwitchFifoFifoMember) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreSwitchFifoFifoMember({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  DelFifoFifoMember: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelFifoFifoMember),
      map((action: DelFifoFifoMember) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreDelFifoFifoMember({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  AddFifoFifo: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddFifoFifo),
      map((action: AddFifoFifo) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreAddFifoFifo({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  DelFifoFifo: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelFifoFifo),
      map((action: DelFifoFifo) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreDelFifoFifo({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

  UpdateFifoFifo: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdateFifoFifo),
      map((action: UpdateFifoFifo) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotFifoError({error: response.error});
              }
              return new StoreUpdateFifoFifo({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotFifoError({error: error}));
            }),
          );
        }
      ));
  });

}

