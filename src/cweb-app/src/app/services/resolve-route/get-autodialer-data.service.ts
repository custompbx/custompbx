import { Injectable } from '@angular/core';
import {ActivatedRouteSnapshot, Resolve, RouterStateSnapshot} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {WsDataService} from '../ws-data.service';
import {UnSubscribe} from '../../store/dataFlow/dataFlow.actions';
import {GetAutoDialerCompanies} from '../../store/apps/autodialer/autodialer.actions';
import {GetDirectoryDomains} from '../../store/directory/directory.actions';

@Injectable({
  providedIn: 'root'
})
export class GetAutodialerDataService implements Resolve<void> {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new UnSubscribe(null));
      this.store.dispatch(new GetDirectoryDomains(null));
      this.store.dispatch(new GetAutoDialerCompanies(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new UnSubscribe(null));
        this.store.dispatch(new GetDirectoryDomains(null));
        this.store.dispatch(new GetAutoDialerCompanies(null));
      }
    });
  }
}
