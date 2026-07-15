import {
  StoreConversationError,
  StoreGetNewConversationMessage,
  GetConversationPrivateMessages,
  StoreGetConversationPrivateMessages,
  StoreCurrentUser,
  StoreGetConversationPrivateCalls,
  GetConversationPrivateCalls,
} from './conversations.actions';

import {Iuser} from '../auth/auth.reducers';

export interface State {
  conversations: { [index: number]: Array<Messages> };
  calls: { [index: number]: Array<Messages> };
  loadCounter: number;
  errorMessage: string | null;
  user: Iuser;
  scrollDown: boolean;
  event: {
    type: 'new-call' | null;
    data: any;
  } | null;
}

export interface Messages {
  id: number;
  created_at: string;
  sender_id: {id: number};
  receiver_id: {id: number};
  text: string;
  timestamp: number;
}

export interface Calls {
  id: number;
  created_at: string;
  sender_id: {id: number};
  receiver_id: {id: number};
  duration: number;
}

export const initialState: State = {
  conversations: {},
  calls: {},
  loadCounter: 0,
  errorMessage: '',
  user: {},
  scrollDown: false,
  event: null,
};

const nullEvent = {type: null, data: null};

export function reducer(state: State = initialState, action): State {
  switch (action.type) {
    case GetConversationPrivateMessages.type:
    case GetConversationPrivateCalls.type: {
      return {
        ...state,
        errorMessage: null,
        loadCounter: state.loadCounter + 1,
        scrollDown: false,
        event: nullEvent,
      };
    }

    case StoreConversationError.type: {
      return {
        ...state,
        errorMessage: action.payload.error || null,
        loadCounter: Math.max(0, state.loadCounter - 1),
        scrollDown: false,
        event: null,
      };
    }
    case StoreCurrentUser.type: {
      const {user} = action.payload;
      return {
        ...state,
        user: user,
        scrollDown: true,
        event: nullEvent,
      };
    }

    case StoreGetConversationPrivateMessages.type: {
      const {response, payload} = action.payload;
      const {id} = payload;
      let {data, error} = response;
      if (!Array.isArray(data)) {
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
        event: nullEvent,
      };
    }
    case StoreGetConversationPrivateCalls.type: {
      const {response, payload} = action.payload;
      const {id} = payload;
      let {data, error} = response;
      if (!Array.isArray(data)) {
        data = [];
      }
      const lastMes = state.calls[id] || [];
      const isFirst = lastMes.length;
      let messages = [...data, ...lastMes];
      if (!payload.up_to_time) {
        messages = [...data];
      }
      return {
        ...state,
        calls: {
          ...state.calls,
          [id]: messages,
        },
        errorMessage: error || null,
        loadCounter: 0,
        scrollDown: isFirst === 0,
        event: nullEvent,
      };
    }

    case StoreGetNewConversationMessage.type: {
      const {response} = action.payload;
      const {data, error} = response;
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
      if (!id) return {...state, errorMessage: null, loadCounter: 0};

      let conversations = state.conversations;
      let calls = state.calls;
      let event = nullEvent;
      if (data.duration === 0 || data.duration) {
        calls = {...calls, [id]: [...(calls[id] || []), data]};
        event = {type: 'new-call', data: {sid, rid}};
      } else {
        conversations = {...conversations, [id]: [...(conversations[id] || []), data]};
      }

      return {
        ...state,
        conversations: conversations,
        calls: calls,
        errorMessage: error || null,
        loadCounter: 0,
        scrollDown: true,
        event: event,
      };
    }

    default: {
      return state;
    }
  }
}
