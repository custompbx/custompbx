import {Injectable} from '@angular/core';
import {Actions, createEffect, ofType} from '@ngrx/effects';
import {WsDataService} from '../../services/ws-data.service';

import {

} from './header.actions';

@Injectable()
export class HeaderEffects {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

}
