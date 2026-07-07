import {Observable, Subscription} from 'rxjs';
import {filter, take} from 'rxjs/operators';

export interface ConnectionSource {
  readonly isConnected: boolean;
  readonly websocketService: {
    readonly status: Observable<boolean>;
  };
}

/** Runs a resolver dispatch exactly once, now or on the next connection. */
export function dispatchWhenConnected(
  connection: ConnectionSource,
  dispatch: () => void
): Subscription {
  if (connection.isConnected) {
    dispatch();
    return new Subscription();
  }

  return connection.websocketService.status.pipe(
    filter(connected => connected),
    take(1)
  ).subscribe(() => dispatch());
}
