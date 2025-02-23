import {All, AuthActionTypes, StoreCommand, StoreTicker} from './phone.actions';

export interface Icreds {
  user_name: string;
  password: string;
  domain: string;
  ws: string;
  stun: string;
}

export interface State {
  phoneCreds: Icreds;
  phoneStatus: Istatuses;
  timer: number;
  command: Icommand;
  errorMessage: string | null;
  lastActionName: string;
}

export interface Istatuses {
  isRunning: boolean;
  inCall: boolean;
  registered: boolean;
  status: 'answered' | 'ringing' | '';
}

export interface Icommand {
  callTo: string;
  hangup: boolean;
  register: boolean;
  answer: boolean;
}

export const initialState: State = {
  phoneCreds: null,
  phoneStatus: <Istatuses>{},
  timer: 0,
  command: {
    callTo: '',
    hangup: false,
    register: false,
    answer: false,
  },
  errorMessage: null,
  lastActionName: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case AuthActionTypes.GET_PHONE_CREDS: {
      return {
        ...state,
        lastActionName: AuthActionTypes.GET_PHONE_CREDS,
      };
    }

    case AuthActionTypes.StoreGetPhoneCreds: {
      return {
        ...state,
        phoneCreds: action.payload.response['phone_creds'],
        errorMessage: action.payload.response.error ||  '',
        command: <Icommand>{},
        lastActionName: AuthActionTypes.StoreGetPhoneCreds,
      };
    }

    case AuthActionTypes.StorePhoneStatus: {
      return {
        ...state,
        phoneStatus: {...state.phoneStatus, ...action.payload.phoneStatus},
        command: <Icommand>{},
        lastActionName: AuthActionTypes.StorePhoneStatus,
      };
    }

    case StoreCommand.type: {
      return {
        ...state,
        command: <Icommand>{...state.command, ...action.payload},
        lastActionName: StoreCommand.type,
      };
    }

    case StoreTicker.type: {
      const targetString = action.payload.date;
      if (!targetString) {
        return {
          ...state,
          timer: 0,
          lastActionName: StoreTicker.type,
        };
      }
      const targetDate = new Date(targetString);
      const currentDate = new Date();
      const timeDifference = currentDate.getTime() - targetDate.getTime() || 0;
      return {
        ...state,
        timer:  Math.floor(timeDifference / 1000),
        lastActionName: StoreTicker.type,
      };
    }

    default: {
      return {
        ...state,
        command: <Icommand>{},
        lastActionName: '',
      };
    }
  }
}
