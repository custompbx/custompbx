import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '../../store/app.states';
import {GetDirectory} from '../../store/config/directory/config.actions.directory';
import {WsDataService} from '../ws-data.service';
import {UnSubscribe} from '../../store/dataFlow/dataFlow.actions';

@Injectable({
  providedIn: 'root'
})
export class GetConfigDirectoryDataService  {

  constructor(
    private store: Store<AppState>,
    private ws: WsDataService,
    ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    if (this.ws.isConnected) {
      this.store.dispatch(new UnSubscribe(null));
      this.store.dispatch(new GetDirectory(null));
    }

    return this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new UnSubscribe(null));
        this.store.dispatch(new GetDirectory(null));
      }
    });
  }
}

