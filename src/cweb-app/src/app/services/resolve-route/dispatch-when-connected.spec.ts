import {Subject} from 'rxjs';
import {dispatchWhenConnected} from './dispatch-when-connected';

describe('dispatchWhenConnected', () => {
  it('dispatches immediately and once when already connected', () => {
    const status = new Subject<boolean>();
    const dispatch = jasmine.createSpy('dispatch');

    dispatchWhenConnected({isConnected: true, websocketService: {status}}, dispatch);
    status.next(true);

    expect(dispatch).toHaveBeenCalledTimes(1);
  });

  it('waits for the first connected state and unsubscribes', () => {
    const status = new Subject<boolean>();
    const dispatch = jasmine.createSpy('dispatch');

    const subscription = dispatchWhenConnected(
      {isConnected: false, websocketService: {status}},
      dispatch
    );
    status.next(false);
    status.next(true);
    status.next(true);

    expect(dispatch).toHaveBeenCalledTimes(1);
    expect(subscription.closed).toBeTrue();
  });
});
