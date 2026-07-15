import {
  StartPhone, ToggleShowPhone,
} from './header.actions';

export interface State {
  phone: {
    started: boolean;
    shown: boolean;
  };
}

export interface Messages {
  id: number;
  sender_id: number;
  receiver_id: number;
  text: string;
  timestamp: number;
  new: boolean;
}

export const initialState: State = {
  phone: {
    started: false,
    shown: false,
  },
};

export function reducer(state: State = initialState, action): State {
  switch (action.type) {
    case StartPhone.type: {
      return {
        ...state,
        phone: {
          ...state.phone,
          started: true,
        },
      };
    }

    case ToggleShowPhone.type: {
      const requested = action.payload?.show;
      const show = typeof requested === 'boolean' ? requested : !state.phone.shown;
      return {
        ...state,
        phone: {
          ...state.phone,
          shown: show,
        },
      };
    }


    default: {
      return state;
    }
  }
}
