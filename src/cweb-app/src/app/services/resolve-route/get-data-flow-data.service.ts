import { Injectable } from '@angular/core';
import {ActivatedRouteSnapshot, Resolve, RouterStateSnapshot} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {WsDataService} from '../ws-data.service';
import {GetDashboard, SubscriptionList} from '../../store/dataFlow/dataFlow.actions';
import {GetModules} from '../../store/config/config.actions';

@Injectable({
  providedIn: 'root'
})
export class GetDataFlowDataService implements Resolve<void> {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new SubscriptionList({values: [new GetDashboard(null).type, new GetModules(null).type]}));
      this.store.dispatch(new GetModules(null));
      this.store.dispatch(new GetDashboard(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new SubscriptionList({values: [new GetDashboard(null).type, new GetModules(null).type]}));
        this.store.dispatch(new GetModules(null));
        this.store.dispatch(new GetDashboard(null));
      }
    });
  }
}
