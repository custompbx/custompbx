import {initialState} from '../config.state.struct';
import {StoreGetSofiaProfileGateways} from './config.actions.sofia';
import {reducer} from './config.reducer.sofia';

describe('Sofia reducer', () => {
  it('completes and clears an empty gateway response for its requested profile', () => {
    const state = {
      ...initialState,
      loadCounter: 1,
      sofia: {
        ...initialState.sofia,
        profiles: {
          1: {id: 1, gateways: {9: {id: 9, name: 'stale'}}},
        },
      },
    } as any;

    const result = reducer(state, new StoreGetSofiaProfileGateways({
      response: {data: {}},
      profileId: 1,
    }));

    expect(result.loadCounter).toBe(0);
    expect(Object.keys(result.sofia.profiles[1].gateways)).toEqual([]);
  });
});
