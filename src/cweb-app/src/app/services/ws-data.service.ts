import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { filter, map } from 'rxjs/operators';
import { WebsocketService } from './websocket';
import { WS } from './websocket.events';
import {CookiesStorageService} from './cookies-storage.service';
import {Iuser} from '../store/auth/auth.reducers';

export interface IMessage {
  MessageType: string;
}

@Injectable({
  providedIn: 'root'
})

export class WsDataService {
  public websocketService: WebsocketService;
  public isConnected: boolean;

  constructor(
    private wsService: WebsocketService,
    private cookie: CookiesStorageService
  ) {
    this.websocketService = this.wsService;
    this.websocketService.status.subscribe(connected => {
      this.isConnected = connected;
    });
  }

  connectToAnotherHost(url: string): Observable<any> {
    // this.websocketService.reconnectToAnotherHost(url);
    this.websocketService.close();
    this.websocketService = new WebsocketService({url: url});
    return this.websocketService.status;
  }

/*
  connectToAnotherHost2(url: string): Observable<any> {
      const newWebsocketService = new WebsocketService({url: url});
      newWebsocketService.status.subscribe(
        connected => {
          if (!connected) {
            newWebsocketService.close();
          } else {
            this.websocketService.close();
            this.websocketService = newWebsocketService;
          }

          return this.websocketService.status;
        }
      );
    }
*/

  logIn(login: string, password: string): Observable<any> {
    this.websocketService.send(WS.ON.LOGIN, {login, password});
    return this.websocketService.on<IMessage>(WS.ON.LOGIN);
  }

  getUserByToken(): Observable<any> {
    this.websocketService.send(WS.ON.RELOGIN, {token: this.cookie.getToken()});
    return this.websocketService.on<IMessage>(WS.ON.RELOGIN);
  }

  getSettings(): Observable<any> {
    this.websocketService.send('get_settings', {token: this.cookie.getToken()});
    return this.websocketService.on<IMessage>(WS.ON.SETTINGS);
  }

  setSettings(data: {}): Observable<any> {
    this.websocketService.send('set_settings', {token: this.cookie.getToken(), ...data});
    return this.websocketService.on<IMessage>(WS.ON.SETTINGS);
  }

  universalSender(method: string, data: any): Observable<any> {
    this.websocketService.send(method, {token: this.cookie.getToken(), ...data});
    return this.websocketService.on<IMessage>(method);
  }

  waitDaemonData(): Observable<any> {
    this.websocketService.send('get_status', {});
    return this.websocketService.on<IMessage>(WS.ON.CONNECTION);
  }

  getStatus(): Observable<Iuser> {
    this.websocketService.send('login');
    return this.websocketService.on<IMessage>(WS.ON.STATUS).pipe(
      filter((action) => 'status' === action.MessageType),
      map(message => <Iuser>{})
    );
  }
}
