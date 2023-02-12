import {reducer as autodialerReducer} from './autodialer/autodialer.reducers';

import {
  initialState,
  State
} from './apps.state.struct';

export function reducer(state = initialState, action: any): State {
  const autodialerState = autodialerReducer(state, action);
  if (autodialerState) {
    return autodialerState;
  }

  return state;

}
