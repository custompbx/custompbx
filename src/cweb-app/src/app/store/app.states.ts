import * as auth from './auth/auth.reducers';
import * as settings from './settings/settings.reducers';
import * as daemon from './daemon/daemon.reducers';
import * as directory from './directory/directory.reducers';
import {reducer as configReducer} from './config/config.reducers';
import * as instances from './instances/instances.reducers';
import * as globalVariables from './global-variables/global-variables.reducer';
import {State as configState} from './config/config.state.struct';
import * as dialplan from './dialplan/dialplan.reducers';
import * as dataFlow from './dataFlow/dataFlow.reducers';
import * as cdr from './cdr/cdr.reducers';
import * as logs from './logs/logs.reducers';
import * as phone from './phone/phone.reducers';
import * as fscli from './fscli/fscli.reducers';
import * as hep from './hep/hep.reducers';
import {State as appsState} from './apps/apps.state.struct';
import {reducer as appsReducer} from './apps/apps.reducers';
import * as conversations from './conversations/conversations.reducer';
import * as header from './header/header.reducer';
import * as app from './app/app.reducer';
import {createFeatureSelector, createSelector} from '@ngrx/store';

export interface AppState {
  auth: auth.State;
  settings: settings.State;
  daemon: daemon.State;
  directory: directory.State;
  config: configState;
  dialplan: dialplan.State;
  dataFlow: dataFlow.State;
  cdr: cdr.State;
  logs: logs.State;
  phone: phone.State;
  fscli: fscli.State;
  hep: hep.State;
  instances: instances.State;
  globalVariables: globalVariables.State;
  apps: appsState;
  conversations: conversations.State;
  header: header.State;
  app: app.State;
}

export const reducers = {
  auth: auth.reducer,
  settings: settings.reducer,
  daemon: daemon.reducer,
  directory: directory.reducer,
  config: configReducer,
  dialplan: dialplan.reducer,
  dataFlow: dataFlow.reducer,
  cdr: cdr.reducer,
  logs: logs.reducer,
  phone: phone.reducer,
  fscli: fscli.reducer,
  hep: hep.reducer,
  instances: instances.reducer,
  globalVariables: globalVariables.reducer,
  apps: appsReducer,
  conversations: conversations.reducer,
  header: header.reducer,
  app: app.reducer,
};

export const selectState = createFeatureSelector<AppState>('app');

export const selectAuthState = createSelector(selectState, (state: AppState) => state.auth);
export const selectSettingsState = createSelector(selectState, (state: AppState) => state.settings);
export const selectDaemonState = createSelector(selectState, (state: AppState) => state.daemon);
export const selectDirectoryState = createSelector(selectState, (state: AppState) => state.directory);
export const selectConfigurationState = createSelector(selectState, (state: AppState) => state.config);
export const selectDialplanState = createSelector(selectState, (state: AppState) => state.dialplan);
export const selectDataFlowState = createSelector(selectState, (state: AppState) => state.dataFlow);
export const selectCDRState = createSelector(selectState, (state: AppState) => state.cdr);
export const selectLogsState = createSelector(selectState, (state: AppState) => state.logs);
export const selectPhoneState = createSelector(selectState, (state: AppState) => state.phone);
export const selectFSCLIState = createSelector(selectState, (state: AppState) => state.fscli);
export const selectHEPState = createSelector(selectState, (state: AppState) => state.hep);
export const selectInstancesState = createSelector(selectState, (state: AppState) => state.instances);
export const selectGlobalVariablesState = createSelector(selectState, (state: AppState) => state.globalVariables);
export const selectApps = createSelector(selectState, (state: AppState) => state.apps);
export const selectConversations = createSelector(selectState, (state: AppState) => state.conversations);
export const selectHeader = createSelector(selectState, (state: AppState) => state.header);
export const selectApp = createSelector(selectState, (state: AppState) => state.app);

