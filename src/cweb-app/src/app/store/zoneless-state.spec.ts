import {reducer as conversationsReducer, initialState as conversationsInitial} from './conversations/conversations.reducer';
import {StoreGetNewConversationMessage} from './conversations/conversations.actions';
import {reducer as headerReducer, initialState as headerInitial} from './header/header.reducer';
import {ToggleShowPhone} from './header/header.actions';
import {updateNestedState} from './config/config.reducers';
import {initialState as configInitial} from './config/config.state.struct';
import {reducer as directoryReducer} from './config/directory/config.reducer.directory';
import {ConfigActionTypes as DirectoryActions} from './config/directory/config.actions.directory';
import {reducer as fifoReducer} from './config/fifo/config.reducer.fifo';
import {ConfigActionTypes as FifoActions} from './config/fifo/config.actions.fifo';
import {reducer as ospReducer} from './config/osp/config.reducer.osp';
import {ConfigActionTypes as OspActions} from './config/osp/config.actions.osp';
import {reducer as lcrReducer} from './config/lcr/config.reducer.lcr';
import {ConfigActionTypes as LcrActions} from './config/lcr/config.actions.lcr';
import {reducer as httpCacheReducer} from './config/http_cache/config.reducer.http_cache';
import {ConfigActionTypes as HttpCacheActions} from './config/http_cache/config.actions.http_cache';
import {reducer as distributorReducer} from './config/distributor/config.reducer.distributor';
import {ConfigActionTypes as DistributorActions} from './config/distributor/config.actions.distributor';
import {reducer as dialplanReducer} from './dialplan/dialplan.reducers';
import {DialplanActionTypes} from './dialplan/dialplan.actions';

describe('zoneless reducer reference guarantees', () => {
  it('appends conversation messages without mutating the previous peer array', () => {
    const oldMessage = {id: 1, text: 'old'} as any;
    const oldMessages = Object.freeze([oldMessage]);
    const state = {
      ...conversationsInitial,
      user: {id: 10},
      conversations: {20: oldMessages},
    } as any;
    const message = {id: 2, sender_id: {id: 20}, receiver_id: {id: 10}, text: 'new'};

    const next = conversationsReducer(state, StoreGetNewConversationMessage({response: {data: message}}));

    expect(next).not.toBe(state);
    expect(next.conversations).not.toBe(state.conversations);
    expect(next.conversations[20]).not.toBe(oldMessages);
    expect(next.conversations[20]).toEqual([oldMessage, message]);
    expect(oldMessages).toEqual([oldMessage]);
  });

  it('appends call events to a fresh call collection', () => {
    const oldCalls = Object.freeze([{id: 1}] as any[]);
    const state = {...conversationsInitial, user: {id: 10}, calls: {20: oldCalls}} as any;
    const call = {id: 2, sender_id: {id: 20}, receiver_id: {id: 10}, duration: 0};

    const next = conversationsReducer(state, StoreGetNewConversationMessage({response: {data: call}}));

    expect(next.calls).not.toBe(state.calls);
    expect(next.calls[20]).not.toBe(oldCalls);
    expect(next.calls[20] as any).toEqual([{id: 1}, call]);
    expect(next.event).toEqual({type: 'new-call', data: {sid: 20, rid: 10}});
  });

  it('honors an explicit false phone visibility request', () => {
    const shown = {...headerInitial, phone: {...headerInitial.phone, shown: true}};
    const hidden = headerReducer(shown, ToggleShowPhone({show: false}));

    expect(hidden.phone).not.toBe(shown.phone);
    expect(hidden.phone.shown).toBeFalse();
    expect(headerReducer(hidden, ToggleShowPhone({})).phone.shown).toBeTrue();
  });

  it('clones every level updated by updateNestedState', () => {
    const state = {module: {profiles: {1: {parameters: {old: true}}}}};
    const next = updateNestedState(state, [{path: ['module', 'profiles', '1', 'parameters'], value: {fresh: true}}]);

    expect(next).not.toBe(state);
    expect(next.module).not.toBe(state.module);
    expect(next.module.profiles).not.toBe(state.module.profiles);
    expect(next.module.profiles[1]).not.toBe(state.module.profiles[1]);
    expect(state.module.profiles[1].parameters).toEqual({old: true});
    expect((next as any).module.profiles[1].parameters).toEqual({fresh: true});
  });

  const nestedParameterCases = [
    {
      name: 'Directory',
      reducer: directoryReducer,
      type: DirectoryActions.StoreGetDirectoryProfileParameters,
      module: 'directory',
      child: 'profiles',
      collection: 'parameters',
    },
    {
      name: 'OSP',
      reducer: ospReducer,
      type: OspActions.StoreGetOspProfileParameters,
      module: 'osp',
      child: 'profiles',
      collection: 'parameters',
    },
    {
      name: 'LCR',
      reducer: lcrReducer,
      type: LcrActions.StoreGetLcrProfileParameters,
      module: 'lcr',
      child: 'profiles',
      collection: 'parameters',
    },
    {
      name: 'FIFO',
      reducer: fifoReducer,
      type: FifoActions.StoreGetFifoFifoMembers,
      module: 'fifo',
      child: 'fifos',
      collection: 'members',
    },
  ];

  for (const testCase of nestedParameterCases) {
    it(`publishes fresh nested references for ${testCase.name} websocket data`, () => {
      const oldCollection = {9: {id: 9, name: 'old'}};
      const oldChild = {id: 1, [testCase.collection]: oldCollection};
      const moduleState = {[testCase.child]: {1: oldChild}};
      const state = {...configInitial, [testCase.module]: moduleState, loadCounter: 1} as any;
      const data = {10: {id: 10, name: 'new', parent: {id: 1}}};

      const next = (testCase.reducer as any)(state, {
        type: testCase.type,
        payload: {response: {data}},
      } as any) as any;

      expect(next).not.toBe(state);
      expect(next[testCase.module]).not.toBe(moduleState);
      expect(next[testCase.module][testCase.child]).not.toBe(moduleState[testCase.child]);
      expect(next[testCase.module][testCase.child][1]).not.toBe(oldChild);
      expect(next[testCase.module][testCase.child][1][testCase.collection]).not.toBe(oldCollection);
      expect(oldCollection[10]).toBeUndefined();
      expect(next[testCase.module][testCase.child][1][testCase.collection][10]).toEqual(data[10]);
    });
  }

  it('publishes fresh HTTP Cache profile and provider references', () => {
    const oldProfile = {id: 1, domains: {9: {id: 9}}, azure: {}, aws_s3: {}};
    const httpCache = {settings: {}, profiles: {1: oldProfile}};
    const state = {...configInitial, http_cache: httpCache, loadCounter: 1} as any;
    const domains = {10: {id: 10, parent: {id: 1}}};

    const next = httpCacheReducer(state, {
      type: HttpCacheActions.StoreGetHttpCacheProfileParameters,
      payload: {response: {data: {domains, azure: {}, aws_s3: {}}}},
    } as any) as any;

    expect(next.http_cache).not.toBe(httpCache);
    expect(next.http_cache.profiles).not.toBe(httpCache.profiles);
    expect(next.http_cache.profiles[1]).not.toBe(oldProfile);
    expect(next.http_cache.profiles[1].domains).not.toBe(oldProfile.domains);
    expect(next.http_cache.profiles[1].domains).toEqual(domains);
  });

  it('publishes fresh Distributor list and node references', () => {
    const oldNodes = {9: {id: 9}};
    const oldList = {id: 1, nodes: oldNodes};
    const distributor = {lists: {1: oldList}};
    const state = {...configInitial, distributor, loadCounter: 1} as any;
    const nodes = {10: {id: 10, parent: {id: 1}}};

    const next = distributorReducer(state, {
      type: DistributorActions.StoreGetDistributorNodes,
      payload: {id: 1, response: {data: nodes}},
    } as any) as any;

    expect(next.distributor).not.toBe(distributor);
    expect(next.distributor.lists).not.toBe(distributor.lists);
    expect(next.distributor.lists[1]).not.toBe(oldList);
    expect(next.distributor.lists[1].nodes).toBe(nodes);
  });

  it('updates dialplan condition data without mutating extensions or conditions', () => {
    const oldCondition = {
      id: 4,
      position: 1,
      enabled: true,
      regexes: [],
      actions: [],
      antiactions: [],
      newRegexes: [],
      newActions: [],
      newAntiactions: [],
      new: [],
    };
    const oldExtension = {id: 3, position: 1, name: 'ext', continue: '', conditions: [oldCondition]};
    const context = {id: 2, name: 'ctx', extensions: [oldExtension]};
    const state = {contexts: {2: context}, debug: {log: [], enabled: false}, staticDialplan: false, loadCounter: 1, errorMessage: null};

    const next = dialplanReducer(state as any, {
      type: DialplanActionTypes.STORE_UPDATE_CONDITION,
      payload: {response: {id: 2, dialplan_conditions: {3: [{...oldCondition, enabled: false}]}}},
    } as any);

    expect(next.contexts).not.toBe(state.contexts);
    expect(next.contexts[2]).not.toBe(context);
    expect(next.contexts[2].extensions).not.toBe(context.extensions);
    expect(next.contexts[2].extensions[0]).not.toBe(oldExtension);
    expect(next.contexts[2].extensions[0].conditions[0]).not.toBe(oldCondition);
    expect(oldCondition.enabled).toBeTrue();
    expect(next.contexts[2].extensions[0].conditions[0].enabled).toBeFalse();
  });
});
