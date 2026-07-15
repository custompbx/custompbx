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
        loadCounter: state.loadCounter >= 0 ? state.loadCounter - 1 : 0,
        // errorMessage: 'Cant get data from server',
      };
    }
    case DataFlowActionTypes.REDUCE_LOAD_COUNTER: {
      return {...state, loadCounter: state.loadCounter > 0 ? state.loadCounter - 1 : 0};
    }

    case DataFlowActionTypes.GET_DASHBOARD: {
      return {
        ...state,
        errorMessage: null, loadCounter: state.loadCounter + 1
      };
    }

    case DataFlowActionTypes.STORE_GET_DASHBOARD: {
      const responseData = <Idashboard>action.payload.response['dashboard_data'];
      const profiles = state.dashboardData.sofia_profiles || [];
      const gateways = state.dashboardData.sofia_gateways || [];
      let sofiaProfiles = responseData?.sofia_profiles;
      let sofiaGateways = responseData?.sofia_gateways;
      if (Array.isArray(sofiaProfiles) && sofiaProfiles.length > 0) {
        let found = false;
        const newProfiles = profiles.map(profile => {
            if (profile.id === sofiaProfiles[0].id) {
              profile = sofiaProfiles[0];
              found = true;
            }
            return profile;
          }
        );
        if (found) sofiaProfiles = newProfiles;
      }
      if (Array.isArray(sofiaGateways) && sofiaGateways.length > 0) {
        let found = false;
        const newGateways = gateways.map(gateway => {
            if (gateway.id === sofiaGateways[0].id) {
              gateway = sofiaGateways[0];
              found = true;
            }
            return gateway;
          }
        );
        sofiaGateways = found ? newGateways : [...gateways, ...sofiaGateways];
      }

      const data = responseData ? {
        ...responseData,
        ...(sofiaProfiles ? {sofia_profiles: sofiaProfiles} : {}),
        ...(sofiaGateways ? {sofia_gateways: sofiaGateways} : {}),
        ...(responseData.domain_sip_regs ? {
          domain_sip_regs: {...state.dashboardData.domain_sip_regs, ...responseData.domain_sip_regs},
        } : {}),
      } : {};

      return {
        ...state,
        dashboardData: {...state.dashboardData, ...data},
        loadCounter: state.loadCounter > 0 ? state.loadCounter - 1 : 0,
      };
    }

    default: {
      return state;
    }
  }
}
