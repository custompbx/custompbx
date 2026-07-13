import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetXmlRpc} from '../../store/config/xml_rpc/config.actions.xml_rpc';
import {WsDataService} from '../ws-data.service';
import {dispatchWhenConnected} from './dispatch-when-connected';
import {UnSubscribe} from '../../store/dataFlow/dataFlow.actions';

@Injectable({
  providedIn: 'root'
})
export class GetConfigXmlRpcDataService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    return dispatchWhenConnected(this.ws, () => {
      this.store.dispatch(new UnSubscribe(null));
      this.store.dispatch(new GetXmlRpc(null));
    });
  }
}


