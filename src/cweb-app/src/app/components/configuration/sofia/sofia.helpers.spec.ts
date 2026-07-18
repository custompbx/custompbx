import {profilesNeedingGatewaySubscription} from './sofia.helpers';

describe('profilesNeedingGatewaySubscription', () => {
  it('returns only profiles that are neither loaded nor already requested', () => {
    const profiles = {
      1: {gateways: {}},
      2: {},
      3: {},
    };

    expect(profilesNeedingGatewaySubscription(profiles, new Set([2]))).toEqual([3]);
  });

  it('handles an absent profile collection', () => {
    expect(profilesNeedingGatewaySubscription(undefined, new Set())).toEqual([]);
  });
});
