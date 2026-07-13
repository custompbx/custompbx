import {Inject, Injectable, OnDestroy} from '@angular/core';
import {EMPTY, Observable, ReplaySubject, Subject, Subscription, interval} from 'rxjs';
import {distinctUntilChanged, filter, map, shareReplay, takeWhile} from 'rxjs/operators';
import {WebSocketSubject, WebSocketSubjectConfig} from 'rxjs/webSocket';
import {IWebsocketService, IWsMessage, WebSocketConfig} from './websocket.interfaces';
import {config} from './websocket.config';

@Injectable({
  providedIn: 'root',
})
export class WebsocketService implements IWebsocketService, OnDestroy {

  readonly config: WebSocketSubjectConfig<IWsMessage<any>>;

  private websocketSub = new Subscription();
  private socketSub = new Subscription();
  private reconnectSub = new Subscription();

  private reconnection$: Observable<number>;
  private websocket$: WebSocketSubject<IWsMessage<any>>;
  private wsMessages$: Subject<IWsMessage<any>>;
  private readonly connectionStatus = new ReplaySubject<boolean>(1);
  private reconnectInterval: number;
  private reconnectAttempts: number;
  private isConnected = false;
  private destroyed = false;

  public readonly status = this.connectionStatus.asObservable().pipe(
    distinctUntilChanged(),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(@Inject(config) private wsConfig: WebSocketConfig) {
    this.wsMessages$ = new Subject<IWsMessage<any>>();
    this.reconnectInterval = wsConfig.reconnectInterval || 5000; // pause between connections
    this.reconnectAttempts = wsConfig.reconnectAttempts || 10; // number of connection attempts

    this.config = {
      url: wsConfig.url,
      closeObserver: {
        next: (event: CloseEvent) => {
          this.websocket$ = null;
          this.emitConnectionStatus(false);
        }
      },
      openObserver: {
        next: (event: Event) => {
          this.emitConnectionStatus(true);
        }
      }
    };

    this.websocketSub = this.wsMessages$.subscribe({
      error: (error: ErrorEvent) => console.error('WebSocket error!', error)
    });

    this.connect();
  }

  ngOnDestroy() {
    this.close();
  }

  /*
  * connect to WebSocked
  * */
  private connect(): void {
    if (this.destroyed || this.websocket$) return;

    this.websocket$ = new WebSocketSubject(this.config);
    this.socketSub.unsubscribe();
    this.socketSub = this.websocket$.subscribe({
        next: (message) => this.emitMessage(message),
        error: () => this.emitConnectionStatus(false)
      }
    );
  }

  /*
  * reconnect if not connecting or errors
  * */
  private reconnect(): void {
    if (this.destroyed || this.reconnection$) return;

    this.reconnection$ = interval(this.reconnectInterval)
      .pipe(takeWhile((v, index) => index < this.reconnectAttempts && !this.websocket$));

    this.reconnectSub.unsubscribe();
    this.reconnectSub = this.reconnection$.subscribe({
      next: () => this.connect(),
      complete: () => {
        this.reconnection$ = null;

        if (!this.websocket$ && !this.destroyed) {
          this.wsMessages$.complete();
          this.connectionStatus.complete();
        }
      }
    });
  }

  /*
  * on message event
  * */
  public on<T>(event: string): Observable<IWsMessage<T>> {
    if (!event) return EMPTY;
    return this.wsMessages$.pipe(
      filter((message: IWsMessage<T>) => message.MessageType === event),
      map((message: IWsMessage<T>) => message)
    );
  }

  public send(event: string, data: any = {}): void {
    if (event && this.isConnected) {
      this.websocket$.next({event, data} as any);
    } else {
      console.error('Send error!');
    }
  }

  public close(): void {
    if (this.destroyed) return;
    this.destroyed = true;
    this.isConnected = false;
    this.reconnectSub.unsubscribe();
    this.socketSub.unsubscribe();
    this.websocketSub.unsubscribe();
    this.websocket$?.complete();
    this.wsMessages$.complete();
    this.connectionStatus.complete();
  }

  private emitConnectionStatus(isConnected: boolean): void {
    this.isConnected = isConnected;
    this.connectionStatus.next(isConnected);
    if (!isConnected) {
      this.websocket$ = null;
      this.reconnect();
    }
  }

  private emitMessage(message: IWsMessage<any>): void {
    this.wsMessages$.next(message);
  }
}
