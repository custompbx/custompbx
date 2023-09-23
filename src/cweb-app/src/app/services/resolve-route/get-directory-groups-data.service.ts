import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetDirectoryGroups} from '../../store/directory/directory.actions';
import {WsDataService} from '../ws-data.service';
import {UnSubscribe} from '../../store/dataFlow/dataFlow.actions';
@Injectable({
  providedIn: 'root'
})
export class GetDirectoryGroupsDataService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new UnSubscribe(null));
      this.store.dispatch(new GetDirectoryGroups(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new UnSubscribe(null));
        this.store.dispatch(new GetDirectoryGroups(null));
      }
    });
  }
}
