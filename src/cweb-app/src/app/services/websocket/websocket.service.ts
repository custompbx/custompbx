import {Injectable, OnDestroy, Inject} from '@angular/core';
import {Observable, SubscriptionLike, Subject, Observer, interval} from 'rxjs';
import {filter, map, tap} from 'rxjs/operators';
import {WebSocketSubject, WebSocketSubjectConfig} from 'rxjs/webSocket';
import {share, distinctUntilChanged, takeWhile} from 'rxjs/operators';
import {IWebsocketService, IWsMessage, WebSocketConfig} from './websocket.interfaces';
import {config} from './websocket.config';

@Injectable({
  providedIn: 'root',
})
export class WebsocketService implements IWebsocketService, OnDestroy {

  readonly config: WebSocketSubjectConfig<IWsMessage<any>>;

  private websocketSub: SubscriptionLike;
  private statusSub: SubscriptionLike;

  private reconnection$: Observable<number>;
  private websocket$: WebSocketSubject<IWsMessage<any>>;
  private connection$: Observer<boolean>;
  private wsMessages$: Subject<IWsMessage<any>>;
  private reconnectInterval: number;
  readonly reconnectAttempts: number;
  private isConnected: boolean;

  public status: Observable<boolean>;

  constructor(@Inject(config) private wsConfig: WebSocketConfig) {
    this.wsMessages$ = new Subject<IWsMessage<any>>();
    this.reconnectInterval = wsConfig.reconnectInterval || 5000; // pause between connections
    this.reconnectAttempts = wsConfig.reconnectAttempts || 10; // number of connection attempts

    this.config = {
      url: wsConfig.url,
      closeObserver: {
        next: (event: CloseEvent) => {
          this.websocket$ = null;
          this.connection$.next(false);
        }
      },
      openObserver: {
        next: (event: Event) => {
          this.connection$.next(true);
        }
      }
    };

    // connection status
    this.status = new Observable<boolean>((observer) => {
      this.connection$ = observer;
    }).pipe(share(), distinctUntilChanged());

    // run reconnect if not connection
    this.statusSub = this.status
      .subscribe((isConnected) => {
        this.isConnected = isConnected;

        if (!this.reconnection$ && typeof (isConnected) === 'boolean' && !isConnected) {
          this.reconnect();
        }
      });

    this.websocketSub = this.wsMessages$.subscribe({
      next: null,
      error: (error: ErrorEvent) => console.error('WebSocket error!', error)
    });

    this.connect();
  }

  ngOnDestroy() {
    this.websocketSub.unsubscribe();
    this.statusSub.unsubscribe();
  }

  /*
  * connect to WebSocked
  * */
  private connect(): void {
    this.websocket$ = new WebSocketSubject(this.config);

    this.websocket$.subscribe({
        next: (message) => this.wsMessages$.next(message),
        error: (error: Event) => {
          if (!this.websocket$) {
            // run reconnect if errors
            this.reconnect();
          }
        }
      }
    );
  }

  /*
  * reconnect if not connecting or errors
  * */
  private reconnect(): void {
    this.reconnection$ = interval(this.reconnectInterval)
      .pipe(takeWhile((v, index) => index < this.reconnectAttempts && !this.websocket$));

    this.reconnection$.subscribe({
      next: () => this.connect(),
      error: null,
      complete: () => {
        // Subject complete if reconnect attemts ending
        this.reconnection$ = null;

        if (!this.websocket$) {
          this.wsMessages$.complete();
          this.connection$.complete();
        }
      }
    });
  }

  /*
  * on message event
  * */
  public on<T>(event: string): Observable<IWsMessage<T>> {
    if (event) {
      return this.wsMessages$.pipe(
        filter((message: IWsMessage<T>) => message.MessageType === event),
        map((message: IWsMessage<T>) => message)
      );
    }
  }

  public send(event: string, data: any = {}): void {
    if (event && this.isConnected) {
      // remote party waits for a string
      this.websocket$.next(<any>JSON.constructor({event, data}));
    } else {
      console.error('Send error!');
    }
  }

  public close(): void {
    this.wsMessages$.complete();
    this.connection$.complete();
    this.wsMessages$.unsubscribe();
    this.websocket$.complete();
    this.statusSub.unsubscribe();
    this.websocket$.unsubscribe();
  }
}
