import { Observable } from 'rxjs';

export interface IWebsocketService {
    status: Observable<boolean>;
    on<T>(message: string): Observable<IWsMessage<T>>;
    send(MessageType: string, data: any): void;
}

export interface WebSocketConfig {
    url: string;
    reconnectInterval?: number;
    reconnectAttempts?: number;
}

export interface IWsMessage<T> {
    MessageType: string;
}
