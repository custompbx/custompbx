import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {WsDataService} from '../../services/ws-data.service';
import {catchError, concatMap, map, switchMap} from 'rxjs/operators';
import {
  Failure, FSCLIActionTypes, SendFSCLICommand, StoreSendFSCLICommand,
} from './fscli.actions';


@Injectable({
  providedIn: 'root'
})
export class FscliEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  SendFSCLICommand: Observable<any> = createEffect(() => {
    return this.actions.pipe(
      ofType(FSCLIActionTypes.SendFSCLICommand),
      map((action: SendFSCLICommand) => action),
      switchMap(action => {
          return this.ws.universalSender(action.type, action.payload).pipe(
            map((response) => {
              return new StoreSendFSCLICommand({response});
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
