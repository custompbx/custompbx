import {Actions} from '@ngrx/effects';
import {Subject} from 'rxjs';

import {
  GetDirectoryUsers,
  StoreGetDirectoryUsers,
} from './directory.actions';
import {DirectoryEffects} from './directory.effects';

describe('DirectoryEffects', () => {
  it('keeps applying subscribed directory-user updates after the initial response', () => {
    const actions = new Subject<any>();
    const responses = new Subject<any>();
    const ws = {
      subscriptionSender: jasmine.createSpy('subscriptionSender').and.returnValue(responses),
    };
    const effects = new DirectoryEffects(new Actions(actions), ws as any);
    const results: StoreGetDirectoryUsers[] = [];

    effects.GetUsers.subscribe((action) => results.push(action));
    actions.next(new GetDirectoryUsers(null));

    responses.next({
      MessageType: 'GetDirectoryUsers',
      data: {directory_users: {id: 5, in_call: false, talking: false}},
    });
    responses.next({
      MessageType: 'GetDirectoryUsers',
      data: {
        directory_users: {
          id: 5,
          in_call: true,
          talking: true,
          last_uuid: 'af3eaa26-c07e-41f6-b58a-d6786201e82c',
        },
      },
    });

    expect(ws.subscriptionSender).toHaveBeenCalledOnceWith(
      new GetDirectoryUsers(null).type,
      null,
    );
    expect(results.length).toBe(2);
    expect(results[1].payload.response.data.directory_users).toEqual(
      jasmine.objectContaining({
        id: 5,
        in_call: true,
        talking: true,
        last_uuid: 'af3eaa26-c07e-41f6-b58a-d6786201e82c',
      }),
    );
  });
});
