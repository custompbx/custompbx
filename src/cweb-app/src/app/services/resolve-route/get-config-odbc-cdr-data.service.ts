import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetOdbcCdr} from '../../store/config/odbc_cdr/config.actions.odbc-cdr';
import {WsDataService} from '../ws-data.service';
import {dispatchWhenConnected} from './dispatch-when-connected';
import {UnSubscribe} from '../../store/dataFlow/dataFlow.actions';

@Injectable({
  providedIn: 'root'
})
export class GetConfigOdbcCdrDataService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    return dispatchWhenConnected(this.ws, () => {
      this.store.dispatch(new UnSubscribe(null));
      this.store.dispatch(new GetOdbcCdr(null));
    });
  }
}
