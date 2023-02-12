
import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {
  ConfigActionTypes,
  SwitchPerlParameter,
  GetPerl,
  StoreDelPerlParameter,
  StoreSwitchPerlParameter,
  UpdatePerlParameter,
  StoreGetPerl,
  StoreAddPerlParameter,
  DelPerlParameter,
  StoreUpdatePerlParameter,
  StoreGotPerlError,
  AddPerlParameter,
} from './config.actions.perl';
import {catchError, map, switchMap} from 'rxjs/operators';
import {WsDataService} from '../../../services/ws-data.service';

@Injectable()
export class ConfigEffectsPerl {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetPerl: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.GetPerl),
      map((action: GetPerl) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPerlError({error: response.error});
              }
              return new StoreGetPerl({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPerlError({error: error}));
            }),
          );
        }
      ));
  });

  UpdatePerlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.UpdatePerlParameter),
      map((action: UpdatePerlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPerlError({error: response.error});
              }
              return new StoreUpdatePerlParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPerlError({error: error}));
            }),
          );
        }
      ));
  });

  SwitchPerlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.SwitchPerlParameter),
      map((action: SwitchPerlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPerlError({error: response.error});
              }
              return new StoreSwitchPerlParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPerlError({error: error}));
            }),
          );
        }
      ));
  });

  AddPerlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.AddPerlParameter),
      map((action: AddPerlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPerlError({error: response.error});
              }
              return new StoreAddPerlParameter({response: response, index: action.payload.index});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPerlError({error: error}));
            }),
          );
        }
      ));
  });

  DelPerlParameter: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(ConfigActionTypes.DelPerlParameter),
      map((action: DelPerlParameter) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              if (response.error) {
                return new StoreGotPerlError({error: response.error});
              }
              return new StoreDelPerlParameter({response});
            }),
            catchError((error) => {
              console.log(error);
              return of(new StoreGotPerlError({error: error}));
            }),
          );
        }
      ));
  });
}

