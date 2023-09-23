import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetDirectoryUsers, GetWebUsersByDirectory} from '../../store/directory/directory.actions';
import {WsDataService} from '../ws-data.service';
import {GetDashboard, SubscriptionList, UnSubscribe} from '../../store/dataFlow/dataFlow.actions';

@Injectable({
  providedIn: 'root'
})
export class GetDirectoryUsersDataWithSubscriptionService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new SubscriptionList({values: [new GetWebUsersByDirectory(null).type, new GetDirectoryUsers(null).type]}));
      this.store.dispatch(new GetDirectoryUsers(null));
      this.store.dispatch(new GetWebUsersByDirectory(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new SubscriptionList({values: [new GetWebUsersByDirectory(null).type, new GetDirectoryUsers(null).type]}));
        this.store.dispatch(new GetDirectoryUsers(null));
        this.store.dispatch(new GetWebUsersByDirectory(null));
      }
    });
  }
}
