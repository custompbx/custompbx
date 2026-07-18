import {hasOperationError, isConfirmedOperation, operationMessage} from './operation-feedback.effects';
import {operationKindFromType, withOperationFeedback} from '../services/operation-feedback';

describe('operation feedback', () => {
  it('recognizes confirmed bracketed Dialplan mutations', () => {
    expect(isConfirmedOperation({
      type: '[Dialplan]{Store}[Add] Condition',
      payload: {response: {condition: {id: 1}}}
    })).toBeTrue();
  });

  it('recognizes token removal confirmations', () => {
    expect(isConfirmedOperation({
      type: 'StoreRemoveUserToken',
      payload: {response: {id: 3}}
    })).toBeTrue();
    expect(operationMessage('StoreRemoveUserToken')).toBe('feedback.itemRemoved');
  });

  it('recognizes legacy persisted delete action names', () => {
    expect(isConfirmedOperation({
      type: 'StoreDelGlobalVariable',
      payload: {response: {id: 4}}
    })).toBeTrue();
  });

  it('ignores local draft actions and reset actions', () => {
    expect(isConfirmedOperation({
      type: '[Dialplan]{Store}[Add] Delete new action',
      payload: {index: 0}
    })).toBeFalse();
    expect(isConfirmedOperation({
      type: 'StoreResetItem',
      payload: {response: {}}
    })).toBeFalse();
  });

  it('ignores failed server confirmations', () => {
    const response = {error: 'unable to save'};
    expect(hasOperationError(response)).toBeTrue();
    expect(isConfirmedOperation({
      type: '[Dialplan]{Store}[Update] Condition',
      payload: {response}
    })).toBeFalse();
  });

  it('does not treat entity names containing Switch as mutations', () => {
    expect(isConfirmedOperation({
      type: 'StoreGetPostSwitch',
      payload: {response: {settings: {}}}
    })).toBeFalse();
    expect(isConfirmedOperation({
      type: '[Config][Store][Get] Post switch',
      payload: {response: {settings: {}}}
    })).toBeFalse();
  });

  it('parses only explicit operation opcodes', () => {
    expect(operationKindFromType('[Config][Store][Get] Post switch')).toBeNull();
    expect(operationKindFromType('[Dialplan][Switch][Store] Parameter')).toBe('switch');
    expect(operationKindFromType('StoreSwitchPostSwitchParameter')).toBe('switch');
    expect(operationKindFromType('StoreGetPostSwitch')).toBeNull();
  });

  it('uses request metadata for ambiguous completion action names', () => {
    const action = withOperationFeedback({
      type: 'StoreAclList',
      payload: {response: {id: 1}},
    }, 'AddAclList');

    expect(action.operationFeedback).toEqual({kind: 'add', sourceType: 'AddAclList'});
    expect(isConfirmedOperation(action)).toBeTrue();
  });
});
