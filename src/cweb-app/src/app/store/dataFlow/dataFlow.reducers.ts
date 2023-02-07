import {All, DataFlowActionTypes} from './dataFlow.actions';

export interface State {
  dashboardData: Idashboard;
  loadCounter: number;
  errorMessage: string | null;
}

export interface Idashboard {
  timestamp: string;
  hostname: string;
  os: string;
  platform: string;
  cpu_model: string;
  cpu_frequency: number;
  dynamic_metrics: IdynamicMetrics;
  domain_sip_regs: { [domain: string]: number };
  domain_verto_regs: { [domain: string]: number };
  sofia_profiles: Array<Iprofile>;
  sofia_gateways: Array<Igateway>;
}

export interface Iprofiles {
  [index: number]: Iprofile;
}

export interface Iprofile {
  enabled: boolean;
  id: number;
  name: string;
  started: boolean;
  state: string;
  uri: string;
}

export interface Igateways {
  [index: number]: Igateway;
}

export interface Igateway {
  enabled: boolean;
  id: number;
  name: string;
  started: boolean;
  state: string;
}

export interface IdynamicMetrics {
  total_memory: number;
  free_memory: number;
  percentage_used_memory: number;
  total_disc_space: number;
  free_disk_space: number;
  percentage_disk_usage: number;
  core_utilization: Array<number>;
}

export const initialState: State = {
  dashboardData: <Idashboard>{},
  loadCounter: 0,
  errorMessage: '',
};

export function reducer(state = initialState, action: All): State {
  switch (action.type) {
    case DataFlowActionTypes.UPDATE_FAILURE: {
      return {
        ...state,
        loadCounter: state.loadCounter >= 0 ? --state.loadCounter : 0,
        // errorMessage: 'Cant get data from server',
      };
    }
    case DataFlowActionTypes.REDUCE_LOAD_COUNTER: {
      return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
    }

    case DataFlowActionTypes.GET_DASHBOARD: {
      return {
        ...state,
        errorMessage: null, loadCounter: state.loadCounter + 1
      };
    }

    case DataFlowActionTypes.STORE_GET_DASHBOARD: {
      const data = <Idashboard>action.payload.response['dashboard_data'];
      const profiles = state.dashboardData.sofia_profiles || [];
      const gateways = state.dashboardData.sofia_gateways || [];
      if (data && data.sofia_profiles && Array.isArray(data.sofia_profiles) && data.sofia_profiles.length > 0) {
        let found = false;
        const newProfiles = profiles.map(profile => {
            if (profile.id === data.sofia_profiles[0].id) {
              profile = data.sofia_profiles[0];
              found = true;
            }
            return profile;
          }
        );
        if (found) {
          data.sofia_profiles = newProfiles;
        } else {
          // data.sofia_profiles = [...profiles, ...data.sofia_profiles];
        }
      }
      if (data && data.sofia_gateways && Array.isArray(data.sofia_gateways) && data.sofia_gateways.length > 0) {
        let found = false;
        const newGateways = gateways.map(gateway => {
            if (gateway.id === data.sofia_gateways[0].id) {
              gateway = data.sofia_gateways[0];
              found = true;
            }
            return gateway;
          }
        );
        if (found) {
          data.sofia_gateways = newGateways || [];
        } else {
          data.sofia_gateways = [...gateways, ...data.sofia_gateways];
        }
      }

      if (data && data.domain_sip_regs) {
        data.domain_sip_regs = {...state.dashboardData.domain_sip_regs, ...data.domain_sip_regs};
      }

      return {
        ...state,
        dashboardData: {...state.dashboardData, ...data},
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    default: {
      return state;
    }
  }
}
