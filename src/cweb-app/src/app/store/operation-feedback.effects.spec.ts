import {hasOperationError, isConfirmedOperation, operationMessage} from './operation-feedback.effects';

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
    expect(operationMessage('StoreRemoveUserToken')).toBe('Item removed successfully.');
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
});
