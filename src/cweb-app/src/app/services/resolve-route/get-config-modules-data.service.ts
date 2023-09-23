import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetModules} from '../../store/config/config.actions';
import {WsDataService} from '../ws-data.service';
import {SubscriptionList, UnSubscribe} from '../../store/dataFlow/dataFlow.actions';
import {GetPostLoadModules} from '../../store/config/post_load_modules/config.actions.PostLoadModules';

@Injectable({
  providedIn: 'root'
})
export class GetConfigModulesDataService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new SubscriptionList({values: [new GetModules(null).type]}));
      this.store.dispatch(new GetModules(null));
      this.store.dispatch(new GetPostLoadModules(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new SubscriptionList({values: [new GetModules(null).type]}));
        this.store.dispatch(new GetModules(null));
        this.store.dispatch(new GetPostLoadModules(null));
      }
    });
  }
}
