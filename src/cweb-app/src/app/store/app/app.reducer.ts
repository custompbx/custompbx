import {ToggleShowConversations} from "./app.actions";

export interface State {
  showConversations: boolean;
}

export const initialState: State = {
  showConversations: false,
};

export function reducer(state: State = initialState, action): State {
  // TODO: fix this
  if (state === null) {
    state = initialState
  }
  switch (action.type) {
    case ToggleShowConversations.type: {
      let showConversations = action.payload?.showConversations;
      if (typeof showConversations !== "boolean") {
        showConversations = !state.showConversations;
      }
      return {
        ...state,
        showConversations: showConversations,
      };
    }

    default: {
      return state;
    }
  }
}
