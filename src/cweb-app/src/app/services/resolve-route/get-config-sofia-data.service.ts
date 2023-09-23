import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetSofiaProfiles} from '../../store/config/sofia/config.actions.sofia';
import {WsDataService} from '../ws-data.service';
import {SubscriptionList, UnSubscribe} from '../../store/dataFlow/dataFlow.actions';

@Injectable({
  providedIn: 'root'
})
export class GetConfigSofiaDataService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new SubscriptionList({values: [new GetSofiaProfiles(null).type]}));
      this.store.dispatch(new GetSofiaProfiles(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new SubscriptionList({values: [new GetSofiaProfiles(null).type]}));
        this.store.dispatch(new GetSofiaProfiles(null));
      }
    });
  }
}
