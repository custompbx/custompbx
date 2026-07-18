import {Actions} from '@ngrx/effects';
import {Subject, of, throwError} from 'rxjs';

import {createActionHelper} from './actions-helper';
import {createEffectForActions} from './effects-helper';

describe('createEffectForActions', () => {
  const request = createActionHelper('AddAclList');
  const completed = createActionHelper('StoreAclList');
  const failed = createActionHelper('StoreGotAclError');

  it('adds explicit write metadata to ambiguous completion actions', () => {
    const source = new Subject<any>();
    const ws = {universalSender: () => of({id: 9})} as any;
    let result: any;

    createEffectForActions(new Actions(source), ws, request, completed, failed).subscribe(action => result = action);
    source.next(request({name: 'trusted'}));

    expect(result.type).toBe('StoreAclList');
    expect(result.operationFeedback).toEqual({kind: 'add', sourceType: 'AddAclList'});
  });

  it('always emits the failure action for server and transport errors', () => {
    const serverSource = new Subject<any>();
    const transportSource = new Subject<any>();
    const failures: any[] = [];

    createEffectForActions(
      new Actions(serverSource),
      {universalSender: () => of({error: 'rejected'})} as any,
      request,
      completed,
      failed,
    ).subscribe(action => failures.push(action));
    createEffectForActions(
      new Actions(transportSource),
      {universalSender: () => throwError(() => new Error('offline'))} as any,
      request,
      completed,
      failed,
    ).subscribe(action => failures.push(action));

    serverSource.next(request({}));
    transportSource.next(request({}));

    expect(failures.length).toBe(2);
    expect(failures.every(action => action.type === 'StoreGotAclError')).toBeTrue();
  });

  it('can explicitly suppress operation feedback for quiet operations', () => {
    const source = new Subject<any>();
    const ws = {universalSender: () => of({id: 9})} as any;
    let result: any;

    createEffectForActions(
      new Actions(source),
      ws,
      request,
      completed,
      failed,
      undefined,
      false,
    ).subscribe(action => result = action);
    source.next(request({name: 'trusted'}));

    expect(result.type).toBe('StoreAclList');
    expect(result.operationFeedback).toBeFalse();
  });
});
