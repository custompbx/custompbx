import {Actions} from '@ngrx/effects';
import {Subject} from 'rxjs';

import {
  GetSofiaProfileGateways,
  StoreGetSofiaProfileGateways,
} from './config.actions.sofia';
import {ConfigEffectsSofia} from './config.effects.sofia';

describe('ConfigEffectsSofia', () => {
  it('completes concurrent profile gateway requests in dispatch order', () => {
    const actions = new Subject<any>();
    const responses: Subject<any>[] = [];
    const ws = {
      universalSender: jasmine.createSpy('universalSender').and.callFake(() => {
        const response = new Subject<any>();
        responses.push(response);
        return response;
      }),
    };
    const effects = new ConfigEffectsSofia(new Actions(actions), ws as any);
    const results: StoreGetSofiaProfileGateways[] = [];

    effects.GetSofiaProfileGateways.subscribe((action) => results.push(action));
    actions.next(new GetSofiaProfileGateways({id: 1, keep_subscription: true}));
    actions.next(new GetSofiaProfileGateways({id: 2, keep_subscription: true}));

    expect(ws.universalSender).toHaveBeenCalledTimes(1);

    responses[0].next({data: {}});
    responses[0].complete();
    expect(ws.universalSender).toHaveBeenCalledTimes(2);

    responses[1].next({data: {}});
    responses[1].complete();

    expect(results.map((action) => action.payload.profileId)).toEqual([1, 2]);
  });
});
