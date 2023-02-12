import { NgModule } from '@angular/core';
import { Routes, RouterModule  } from '@angular/router';
import {LoginComponent} from './components/login/login.component';
import {DashboardComponent} from './components/dashboard/dashboard.component';
import { AuthGuardService as AuthGuard } from './services/auth-guard.service';
import {NotFoundComponent} from './components/not-found/not-found.component';
import {GetDirectoryDomainsDataService} from './services/resolve-route/get-directory-domains-data.service';
import {GetDirectoryUsersDataService} from './services/resolve-route/get-directory-users-data.service';
import {GetDirectoryGroupsDataService} from './services/resolve-route/get-directory-groups-data.service';
import {GetDirectoryGatewayDataService} from './services/resolve-route/get-directory-gateways-data.service';
import {GetConfigAclDataService} from './services/resolve-route/get-config-acl-data.service';
import {GetConfigSofiaDataService} from './services/resolve-route/get-config-sofia-data.service';
import {GetConfigModulesDataService} from './services/resolve-route/get-config-modules-data.service';
import {GetDialplanContextsDataService} from './services/resolve-route/get-dialplan-contexts-data.service';
import {GetDataFlowDataService} from './services/resolve-route/get-data-flow-data.service';
import {GetConfigCdrPgCsvDataService} from './services/resolve-route/get-config-cdr-pg-csv-data.service';
import {GetDirectoryUsersDataWithSubscriptionService} from './services/resolve-route/get-directory-users-data-with-subscription.service';
import {UnsubscribeService} from './services/resolve-route/unsubscribe.service';
import {GetConfigVertoDataService} from './services/resolve-route/get-config-verto-data.service';
import {GetConfigCallcenterDataService} from './services/resolve-route/get-config-callcenter-data.service';
import {GetConfigOdbcCdrDataService} from './services/resolve-route/get-config-odbc-cdr-data.service';
import {GetConfigLcrDataService} from './services/resolve-route/get-config-lcr-data.service';
import {GetConfigShoutDataService} from './services/resolve-route/get-config-shout-data.service';
import {GetConfigRedisDataService} from './services/resolve-route/get-config-redis-data.service';
import {GetConfigNibblebillDataService} from './services/resolve-route/get-config-nibblebill-data.service';
import {GetConfigAvmdDataService} from './services/resolve-route/get-config-avmd-data.service';
import {GetConfigCdrMongodbDataService} from './services/resolve-route/get-config-cdr-mongodb-data.service';
import {GetConfigTtsCommandlineDataService} from './services/resolve-route/get-config-tts-commandline-data.service';
import {GetConfigPythonDataService} from './services/resolve-route/get-config-python-data.service';
import {GetConfigOpusDataService} from './services/resolve-route/get-config-opus-data.service';
import {GetConfigMemcacheDataService} from './services/resolve-route/get-config-memcache-data.service';
import {GetConfigHttpCacheDataService} from './services/resolve-route/get-config-http-cache-data.service';
import {GetConfigDbDataService} from './services/resolve-route/get-config-db-data.service';
import {GetConfigSndfileDataService} from './services/resolve-route/get-config-sndfile-data.service';
import {GetConfigXmlCdrDataService} from './services/resolve-route/get-config-xml-cdr-data.service';
import {GetConfigXmlRpcDataService} from './services/resolve-route/get-config-xml-rpc-data.service';
import {GetConfigZeroconfDataService} from './services/resolve-route/get-config-zeroconf-data.service';
import {GetConfigErlangEventDataService} from './services/resolve-route/get-config-erlang-event-data.service';
import {GetConfigCurlDataService} from './services/resolve-route/get-config-curl-data.service';
import {GetConfigOrekaDataService} from './services/resolve-route/get-config-oreka-data.service';
import {GetConfigEasyrouteDataService} from './services/resolve-route/get-config-easyroute-data.service';
import {GetConfigMongoDataService} from './services/resolve-route/get-config-mongo-data.service';
import {GetConfigAmrDataService} from './services/resolve-route/get-config-amr-data.service';
import {GetConfigCepstralDataService} from './services/resolve-route/get-config-cepstral-data.service';
import {GetConfigDialplanDirectoryDataService} from './services/resolve-route/get-config-dialplan-directory-data.service';
import {GetConfigAmrwbDataService} from './services/resolve-route/get-config-amrwb-data.service';
import {GetConfigEventMulticastDataService} from './services/resolve-route/get-config-event-multicast-data.service';
import {GetConfigAlsaDataService} from './services/resolve-route/get-config-alsa-data.service';
import {GetConfigSangomaCodecDataService} from './services/resolve-route/get-config-sangoma-codec-data.service';
import {GetConfigFaxDataService} from './services/resolve-route/get-config-fax-data.service';
import {GetConfigMsrpDataService} from './services/resolve-route/get-config-msrp-data.service';
import {GetConfigPerlDataService} from './services/resolve-route/get-config-perl-data.service';
import {GetConfigCidlookupDataService} from './services/resolve-route/get-config-cidlookup-data.service';
import {GetConfigPocketsphinxDataService} from './services/resolve-route/get-config-pocketsphinx-data.service';
import {GetConfigLuaDataService} from './services/resolve-route/get-config-lua-data.service';
import {GetConfigPostSwitchDataService} from './services/resolve-route/get-config-post-switch-data.service';
import {GetConfigDistributorDataService} from './services/resolve-route/get-config-distributor-data.service';
import {GetConfigDirectoryDataService} from './services/resolve-route/get-config-directory-data.service';
import {GetConfigFifoDataService} from './services/resolve-route/get-config-fifo-data.service';
import {GetConfigOpalDataService} from './services/resolve-route/get-config-opal-data.service';
import {GetConfigOspDataService} from './services/resolve-route/get-config-osp-data.service';
import {GetConfigUnicallDataService} from './services/resolve-route/get-config-unicall-data.service';
import {GetConfigConferenceDataService} from './services/resolve-route/get-config-conference-data.service';
import {GetInstancesDataService} from './services/resolve-route/get-instances-data.service';
import {GetGlobalVariablesDataService} from './services/resolve-route/get-global-variables-data.service';
import {GetPostLoadModulesDataService} from './services/resolve-route/get-post-load-modules-data.service';
import {LazyWrapperComponent} from './components/lazy-wrapper/lazy-wrapper.component';
import {GetVoicemailDataService} from './services/resolve-route/get-voicemail-data.service';
import {GetAutodialerDataService} from './services/resolve-route/get-autodialer-data.service';

const routes: Routes = [
  {path: '', component: LoginComponent, canActivate: [AuthGuard]},
  {path: 'login',
    title: 'login', component: LoginComponent, canActivate: [AuthGuard]},
  {path: 'dashboard',
    title: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard], resolve: {reconnectUpdater: GetDataFlowDataService}},
  {path: 'settings',
    title: 'settings', component: LazyWrapperComponent, canActivate: [AuthGuard], resolve: {reconnectUpdater: UnsubscribeService}},
  {
    path: 'directory/domains',
    title: 'directory/domains',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDirectoryDomainsDataService}
  },
  {
    path: 'directory/users',
    title: 'directory/users',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDirectoryUsersDataService}
  },
  {
    path: 'directory/groups',
    title: 'directory/groups',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDirectoryGroupsDataService}
  },
  {
    path: 'directory/gateways',
    title: 'directory/gateways',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDirectoryGatewayDataService}
  },
  {
    path: 'configuration/modules',
    title: 'configuration/modules',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigModulesDataService}
  },
  {
    path: 'configuration/acl',
    title: 'configuration/acl',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigAclDataService}
  },
  {
    path: 'configuration/callcenter',
    title: 'configuration/callcenter',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigCallcenterDataService}
  },
  {
    path: 'configuration/sofia',
    title: 'configuration/sofia',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigSofiaDataService}
  },
  {
    path: 'configuration/verto',
    title: 'configuration/verto',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigVertoDataService}
  },
  {
    path: 'configuration/cdr-pg-csv',
    title: 'configuration/cdr-pg-csv',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigCdrPgCsvDataService}
  },
  {
    path: 'configuration/odbc-cdr',
    title: 'configuration/odbc-cdr',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigOdbcCdrDataService}
  },
  {
    path: 'configuration/lcr',
    title: 'configuration/lcr',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigLcrDataService}
  },
  {
    path: 'configuration/shout',
    title: 'configuration/shout',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigShoutDataService}
  },
  {
    path: 'configuration/redis',
    title: 'configuration/redis',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigRedisDataService}
  },
  {
    path: 'configuration/nibblebill',
    title: 'configuration/nibblebill',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigNibblebillDataService}
  },
  {
    path: 'configuration/avmd',
    title: 'configuration/avmd',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigAvmdDataService}
  },
  {
    path: 'configuration/cdr-mongodb',
    title: 'configuration/cdr-mongodb',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigCdrMongodbDataService}
  },
  {
    path: 'configuration/db',
    title: 'configuration/db',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigDbDataService}
  },
  {
    path: 'configuration/http-cache',
    title: 'configuration/http-cache',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigHttpCacheDataService}
  },
  {
    path: 'configuration/memcache',
    title: 'configuration/memcache',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigMemcacheDataService}
  },
  {
    path: 'configuration/opus',
    title: 'configuration/opus',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigOpusDataService}
  },
  {
    path: 'configuration/python',
    title: 'configuration/python',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigPythonDataService}
  },
  {
    path: 'configuration/tts-commandline',
    title: 'configuration/tts-commandline',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigTtsCommandlineDataService}
  },
  {
    path: 'configuration/alsa',
    title: 'configuration/alsa',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigAlsaDataService}
  },
  {
    path: 'configuration/amr',
    title: 'configuration/amr',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigAmrDataService}
  },
  {
    path: 'configuration/amrwb',
    title: 'configuration/amrwb',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigAmrwbDataService}
  },
  {
    path: 'configuration/cepstral',
    title: 'configuration/cepstral',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigCepstralDataService}
  },
  {
    path: 'configuration/cidlookup',
    title: 'configuration/cidlookup',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigCidlookupDataService}
  },
  {
    path: 'configuration/curl',
    title: 'configuration/curl',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigCurlDataService}
  },
  {
    path: 'configuration/dialplan-directory',
    title: 'configuration/dialplan-directory',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigDialplanDirectoryDataService}
  },
  {
    path: 'configuration/easyroute',
    title: 'configuration/easyroute',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigEasyrouteDataService}
  },
  {
    path: 'configuration/erlang-event',
    title: 'configuration/erlang-event',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigErlangEventDataService}
  },
  {
    path: 'configuration/event-multicast',
    title: 'configuration/event-multicast',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigEventMulticastDataService}
  },
  {
    path: 'configuration/fax',
    title: 'configuration/fax',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigFaxDataService}
  },
  {
    path: 'configuration/lua',
    title: 'configuration/lua',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigLuaDataService}
  },
  {
    path: 'configuration/mongo',
    title: 'configuration/mongo',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigMongoDataService}
  },
  {
    path: 'configuration/msrp',
    title: 'configuration/msrp',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigMsrpDataService}
  },
  {
    path: 'configuration/oreka',
    title: 'configuration/oreka',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigOrekaDataService}
  },
  {
    path: 'configuration/perl',
    title: 'configuration/perl',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigPerlDataService}
  },
  {
    path: 'configuration/pocketsphinx',
    title: 'configuration/pocketsphinx',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigPocketsphinxDataService}
  },
  {
    path: 'configuration/sangoma-codec',
    title: 'configuration/sangoma-codec',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigSangomaCodecDataService}
  },
  {
    path: 'configuration/sndfile',
    title: 'configuration/sndfile',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigSndfileDataService}
  },
  {
    path: 'configuration/xml-cdr',
    title: 'configuration/xml-cdr',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigXmlCdrDataService}
  },
  {
    path: 'configuration/xml-rpc',
    title: 'configuration/xml-rpc',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigXmlRpcDataService}
  },
  {
    path: 'configuration/zeroconf',
    title: 'configuration/zeroconf',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigZeroconfDataService}
  },
  {
    path: 'configuration/post-load-switch',
    title: 'configuration/post-load-switch',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigPostSwitchDataService}
  },
  {
    path: 'configuration/distributor',
    title: 'configuration/distributor',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigDistributorDataService}
  },
  {
    path: 'configuration/directory',
    title: 'configuration/directory',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigDirectoryDataService}
  },
  {
    path: 'configuration/fifo',
    title: 'configuration/fifo',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigFifoDataService}
  },
  {
    path: 'configuration/opal',
    title: 'configuration/opal',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigOpalDataService}
  },
  {
    path: 'configuration/osp',
    title: 'configuration/osp',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigOspDataService}
  },
  {
    path: 'configuration/unicall',
    title: 'configuration/unicall',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigUnicallDataService}
  },

  {
    path: 'dialplan/contexts',
    title: 'dialplan/contexts',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDialplanContextsDataService}
  },
  {
    path: 'dashboard/users-panel',
    title: 'dashboard/users-panel',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDirectoryUsersDataWithSubscriptionService}
  },
  {
    path: 'configuration/conference',
    title: 'configuration/conference',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetConfigConferenceDataService}
  },
  {
    path: 'instances',
    title: 'instances',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetInstancesDataService}
  },
  {
    path: 'global-variables',
    title: 'global-variables',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetGlobalVariablesDataService}
  },
  {
    path: 'configuration/post-load-modules',
    title: 'configuration/post-load-modules',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetPostLoadModulesDataService}
  },
  {
    path: 'configuration/voicemail',
    title: 'configuration/voicemail',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetVoicemailDataService}
  },
  {path: 'cdr',
    title: 'cdr', component: LazyWrapperComponent, canActivate: [AuthGuard], resolve: {reconnectUpdater: UnsubscribeService}},
  {path: 'logs',
    title: 'logs', component: LazyWrapperComponent, canActivate: [AuthGuard], resolve: {reconnectUpdater: UnsubscribeService}},
  {path: 'fs-cli',
    title: 'fs-cli', component: LazyWrapperComponent, canActivate: [AuthGuard], resolve: {reconnectUpdater: UnsubscribeService}},
  {path: 'hep',
    title: 'hep', component: LazyWrapperComponent, canActivate: [AuthGuard], resolve: {reconnectUpdater: UnsubscribeService}},
  {
    path: 'apps/autodialer',
    title: 'apps/autodialer',
    component: LazyWrapperComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetAutodialerDataService}
  },
  {path: '**',
    title: '**', component: NotFoundComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {})],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
