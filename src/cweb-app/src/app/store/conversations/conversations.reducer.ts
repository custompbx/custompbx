import {
  StoreConversationError,
  StoreGetNewConversationMessage,
  GetConversationPrivateMessages, StoreGetConversationPrivateMessages, StoreCurrentUser,
} from './conversations.actions';

import {isArray} from "chart.js/helpers";
import {Iuser} from "../auth/auth.reducers";

export interface State {
  conversations: { [index: number]: Array<Messages> };
  loadCounter: number;
  errorMessage: string | null;
  user: Iuser;
  scrollDown: boolean;
}

export interface Messages {
  id: number;
  sender_id: {id: number};
  receiver_id: {id: number};
  text: string;
  timestamp: number;
  new: boolean;
}

export const initialState: State = {
  conversations: {},
  loadCounter: 0,
  errorMessage: '',
  user: {},
  scrollDown: false,
};

export function reducer(state: State = initialState, action): State {
  // TODO: fix this
  if (state === null) {
    state = initialState
  }
  switch (action.type) {
    case GetConversationPrivateMessages.type: {
      return {
        ...state,
        errorMessage: null,
        loadCounter: state.loadCounter + 1,
        scrollDown: false,
      };
    }

    case StoreConversationError.type: {
      return {
        ...state,
        errorMessage: action.payload.error || null,
        loadCounter: Math.max(0, state.loadCounter - 1),
        scrollDown: false,
      };
    }
    case StoreCurrentUser.type: {
      const {user} = action.payload;
      return {
        ...state,
        user: user,
        scrollDown: true,
      };
    }

    case StoreGetConversationPrivateMessages.type: {
      const {response, payload} = action.payload;
      const {id} = payload;
      let {data, error} = response;
      if (isArray(data)) {
        data.sort((a: any, b: any) => b.created_at - a.created_at).reverse();
      } else {
        data = [];
      }
      const lastMes = state.conversations[id] || [];
      const isFirst = lastMes.length;
      let messages = [...data, ...lastMes];
      if (!payload.up_to_time) {
        messages = [...data];
      }
      return {
        ...state,
        conversations: {
          ...state.conversations,
          [id]: messages,
        },
        errorMessage: error || null,
        loadCounter: 0,
        scrollDown: isFirst === 0,
      };
    }

    case StoreGetNewConversationMessage.type: {
      const {response} = action.payload;
      let {data, error} = response;
      if (error) {
        return {
          ...state,
          errorMessage: action.payload.error || null,
          loadCounter: Math.max(0, state.loadCounter - 1),
        };
      }
      if (!state.user?.id) {
        return {
          ...state,
          loadCounter: Math.max(0, state.loadCounter - 1),
        };
      }
      const sid = data.sender_id?.id;
      const rid = data.receiver_id?.id;
      let id = sid;
      if (state.user?.id === sid) {
          id = rid;
      }
      const conversations = {...state.conversations};
      if (!conversations[id]) {
        conversations[id] = [];
      }
      data.new = true;
      conversations[id].push(data);
      return {
        ...state,
        conversations: conversations,
        errorMessage: error || null,
        loadCounter: 0,
        scrollDown: true,
      };
    }

    default: {
      return state;
    }
  }
}
