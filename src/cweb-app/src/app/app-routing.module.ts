import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// --- Static Components (Keep imports for non-lazy components) ---
import { LoginComponent } from './components/login/login.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { NotFoundComponent } from './components/not-found/not-found.component';
import { AuthGuardService as AuthGuard } from './services/auth-guard.service';

// --- Data Resolvers (Keep imports for resolvers) ---
// NOTE: All resolver imports are kept to maintain the existing data loading logic.
import { GetDataFlowDataService } from './services/resolve-route/get-data-flow-data.service';
import { UnsubscribeService } from './services/resolve-route/unsubscribe.service';
// ... (All your resolver imports are assumed to be here but are omitted for brevity) ...
import { GetDirectoryDomainsDataService } from './services/resolve-route/get-directory-domains-data.service';
import { GetDirectoryUsersDataService } from './services/resolve-route/get-directory-users-data.service';
import { GetDirectoryGroupsDataService } from './services/resolve-route/get-directory-groups-data.service';
import { GetDirectoryGatewayDataService } from './services/resolve-route/get-directory-gateways-data.service';
import { GetConfigAclDataService } from './services/resolve-route/get-config-acl-data.service';
import { GetConfigSofiaDataService } from './services/resolve-route/get-config-sofia-data.service';
import { GetConfigModulesDataService } from './services/resolve-route/get-config-modules-data.service';
import { GetDialplanContextsDataService } from './services/resolve-route/get-dialplan-contexts-data.service';
import { GetConfigCdrPgCsvDataService } from './services/resolve-route/get-config-cdr-pg-csv-data.service';
import { GetDirectoryUsersDataWithSubscriptionService } from './services/resolve-route/get-directory-users-data-with-subscription.service';
import { GetConfigVertoDataService } from './services/resolve-route/get-config-verto-data.service';
import { GetConfigCallcenterDataService } from './services/resolve-route/get-config-callcenter-data.service';
import { GetConfigOdbcCdrDataService } from './services/resolve-route/get-config-odbc-cdr-data.service';
import { GetConfigLcrDataService } from './services/resolve-route/get-config-lcr-data.service';
import { GetConfigShoutDataService } from './services/resolve-route/get-config-shout-data.service';
import { GetConfigRedisDataService } from './services/resolve-route/get-config-redis-data.service';
import { GetConfigNibblebillDataService } from './services/resolve-route/get-config-nibblebill-data.service';
import { GetConfigAvmdDataService } from './services/resolve-route/get-config-avmd-data.service';
import { GetConfigCdrMongodbDataService } from './services/resolve-route/get-config-cdr-mongodb-data.service';
import { GetConfigTtsCommandlineDataService } from './services/resolve-route/get-config-tts-commandline-data.service';
import { GetConfigPythonDataService } from './services/resolve-route/get-config-python-data.service';
import { GetConfigOpusDataService } from './services/resolve-route/get-config-opus-data.service';
import { GetConfigMemcacheDataService } from './services/resolve-route/get-config-memcache-data.service';
import { GetConfigHttpCacheDataService } from './services/resolve-route/get-config-http-cache-data.service';
import { GetConfigDbDataService } from './services/resolve-route/get-config-db-data.service';
import { GetConfigSndfileDataService } from './services/resolve-route/get-config-sndfile-data.service';
import { GetConfigXmlCdrDataService } from './services/resolve-route/get-config-xml-cdr-data.service';
import { GetConfigXmlRpcDataService } from './services/resolve-route/get-config-xml-rpc-data.service';
import { GetConfigZeroconfDataService } from './services/resolve-route/get-config-zeroconf-data.service';
import { GetConfigErlangEventDataService } from './services/resolve-route/get-config-erlang-event-data.service';
import { GetConfigCurlDataService } from './services/resolve-route/get-config-curl-data.service';
import { GetConfigOrekaDataService } from './services/resolve-route/get-config-oreka-data.service';
import { GetConfigEasyrouteDataService } from './services/resolve-route/get-config-easyroute-data.service';
import { GetConfigMongoDataService } from './services/resolve-route/get-config-mongo-data.service';
import { GetConfigAmrDataService } from './services/resolve-route/get-config-amr-data.service';
import { GetConfigCepstralDataService } from './services/resolve-route/get-config-cepstral-data.service';
import { GetConfigDialplanDirectoryDataService } from './services/resolve-route/get-config-dialplan-directory-data.service';
import { GetConfigAmrwbDataService } from './services/resolve-route/get-config-amrwb-data.service';
import { GetConfigEventMulticastDataService } from './services/resolve-route/get-config-event-multicast-data.service';
import { GetConfigAlsaDataService } from './services/resolve-route/get-config-alsa-data.service';
import { GetConfigSangomaCodecDataService } from './services/resolve-route/get-config-sangoma-codec-data.service';
import { GetConfigFaxDataService } from './services/resolve-route/get-config-fax-data.service';
import { GetConfigMsrpDataService } from './services/resolve-route/get-config-msrp-data.service';
import { GetConfigPerlDataService } from './services/resolve-route/get-config-perl-data.service';
import { GetConfigCidlookupDataService } from './services/resolve-route/get-config-cidlookup-data.service';
import { GetConfigPocketsphinxDataService } from './services/resolve-route/get-config-pocketsphinx-data.service';
import { GetConfigLuaDataService } from './services/resolve-route/get-config-lua-data.service';
import { GetConfigPostSwitchDataService } from './services/resolve-route/get-config-post-switch-data.service';
import { GetConfigDistributorDataService } from './services/resolve-route/get-config-distributor-data.service';
import { GetConfigDirectoryDataService } from './services/resolve-route/get-config-directory-data.service';
import { GetConfigFifoDataService } from './services/resolve-route/get-config-fifo-data.service';
import { GetConfigOpalDataService } from './services/resolve-route/get-config-opal-data.service';
import { GetConfigOspDataService } from './services/resolve-route/get-config-osp-data.service';
import { GetConfigUnicallDataService } from './services/resolve-route/get-config-unicall-data.service';
import { GetConfigConferenceDataService } from './services/resolve-route/get-config-conference-data.service';
import { GetInstancesDataService } from './services/resolve-route/get-instances-data.service';
import { GetGlobalVariablesDataService } from './services/resolve-route/get-global-variables-data.service';
import { GetPostLoadModulesDataService } from './services/resolve-route/get-post-load-modules-data.service';
import { GetVoicemailDataService } from './services/resolve-route/get-voicemail-data.service';
import { GetAutodialerDataService } from './services/resolve-route/get-autodialer-data.service';

export const routes: Routes = [
  // --- Static Routes ---
  {path: '', component: LoginComponent, canActivate: [AuthGuard]},
  {path: 'login', title: 'login', component: LoginComponent, canActivate: [AuthGuard]},
  {
    path: 'dashboard',
    title: 'dashboard',
    component: DashboardComponent,
    canActivate: [AuthGuard],
    resolve: {reconnectUpdater: GetDataFlowDataService}
  },

  // --- Dynamic/Lazy Routes using loadComponent ---

  // NOTE: Assuming settings component is at: ../settings/settings.component
  {
    path: 'settings',
    title: 'settings',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/settings/settings.component').then(m => m.SettingsComponent),
    resolve: {reconnectUpdater: UnsubscribeService}
  },
  // NOTE: All components are assumed to be in a path matching the route segment.

  // --- Directory Routes ---
  {
    path: 'directory/domains',
    title: 'directory/domains',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/directory/domains/domains.component').then(m => m.DomainsComponent),
    resolve: {reconnectUpdater: GetDirectoryDomainsDataService}
  },
  {
    path: 'directory/users',
    title: 'directory/users',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/directory/users/users.component').then(m => m.UsersComponent),
    resolve: {reconnectUpdater: GetDirectoryUsersDataService}
  },
  {
    path: 'directory/groups',
    title: 'directory/groups',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/directory/groups/groups.component').then(m => m.GroupsComponent),
    resolve: {reconnectUpdater: GetDirectoryGroupsDataService}
  },
  {
    path: 'directory/gateways',
    title: 'directory/gateways',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/directory/gateways/gateways.component').then(m => m.GatewaysComponent),
    resolve: {reconnectUpdater: GetDirectoryGatewayDataService}
  },

  // --- Configuration Routes ---

  // NOTE: For brevity, components are assumed to be in a shared path like ./components/config/modules/modules.component
  // You may need to adjust the path based on your actual file structure.
  {
    path: 'configuration/modules',
    title: 'configuration/modules',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/modules/modules.component').then(m => m.ModulesComponent),
    resolve: {reconnectUpdater: GetConfigModulesDataService}
  },
  {
    path: 'configuration/acl',
    title: 'configuration/acl',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/acl/acl.component').then(m => m.AclComponent),
    resolve: {reconnectUpdater: GetConfigAclDataService}
  },
  {
    path: 'configuration/callcenter',
    title: 'configuration/callcenter',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/callcenter/callcenter.component').then(m => m.CallcenterComponent),
    resolve: {reconnectUpdater: GetConfigCallcenterDataService}
  },
  {
    path: 'configuration/sofia',
    title: 'configuration/sofia',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/sofia/sofia.component').then(m => m.SofiaComponent),
    resolve: {reconnectUpdater: GetConfigSofiaDataService}
  },
  {
    path: 'configuration/verto',
    title: 'configuration/verto',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/verto/verto.component').then(m => m.VertoComponent),
    resolve: {reconnectUpdater: GetConfigVertoDataService}
  },
  {
    path: 'configuration/cdr-pg-csv',
    title: 'configuration/cdr-pg-csv',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/cdr-pg-csv/cdr-pg-csv.component').then(m => m.CdrPgCsvComponent),
    resolve: {reconnectUpdater: GetConfigCdrPgCsvDataService}
  },
  {
    path: 'configuration/odbc-cdr',
    title: 'configuration/odbc-cdr',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/odbc-cdr/odbc-cdr.component').then(m => m.OdbcCdrComponent),
    resolve: {reconnectUpdater: GetConfigOdbcCdrDataService}
  },
  {
    path: 'configuration/lcr',
    title: 'configuration/lcr',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/lcr/lcr.component').then(m => m.LcrComponent),
    resolve: {reconnectUpdater: GetConfigLcrDataService}
  },
  {
    path: 'configuration/shout',
    title: 'configuration/shout',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/shout/shout.component').then(m => m.ShoutComponent),
    resolve: {reconnectUpdater: GetConfigShoutDataService}
  },
  {
    path: 'configuration/redis',
    title: 'configuration/redis',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/redis/redis.component').then(m => m.RedisComponent),
    resolve: {reconnectUpdater: GetConfigRedisDataService}
  },
  {
    path: 'configuration/nibblebill',
    title: 'configuration/nibblebill',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/nibblebill/nibblebill.component').then(m => m.NibblebillComponent),
    resolve: {reconnectUpdater: GetConfigNibblebillDataService}
  },
  {
    path: 'configuration/avmd',
    title: 'configuration/avmd',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/avmd/avmd.component').then(m => m.AvmdComponent),
    resolve: {reconnectUpdater: GetConfigAvmdDataService}
  },
  {
    path: 'configuration/cdr-mongodb',
    title: 'configuration/cdr-mongodb',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/cdr-mongodb/cdr-mongodb.component').then(m => m.CdrMongodbComponent),
    resolve: {reconnectUpdater: GetConfigCdrMongodbDataService}
  },
  {
    path: 'configuration/db',
    title: 'configuration/db',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/db/db.component').then(m => m.DbComponent),
    resolve: {reconnectUpdater: GetConfigDbDataService}
  },
  {
    path: 'configuration/http-cache',
    title: 'configuration/http-cache',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/http-cache/http-cache.component').then(m => m.HttpCacheComponent),
    resolve: {reconnectUpdater: GetConfigHttpCacheDataService}
  },
  {
    path: 'configuration/memcache',
    title: 'configuration/memcache',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/memcache/memcache.component').then(m => m.MemcacheComponent),
    resolve: {reconnectUpdater: GetConfigMemcacheDataService}
  },
  {
    path: 'configuration/opus',
    title: 'configuration/opus',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/opus/opus.component').then(m => m.OpusComponent),
    resolve: {reconnectUpdater: GetConfigOpusDataService}
  },
  {
    path: 'configuration/python',
    title: 'configuration/python',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/python/python.component').then(m => m.PythonComponent),
    resolve: {reconnectUpdater: GetConfigPythonDataService}
  },
  {
    path: 'configuration/tts-commandline',
    title: 'configuration/tts-commandline',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/tts-commandline/tts-commandline.component').then(m => m.TtsCommandlineComponent),
    resolve: {reconnectUpdater: GetConfigTtsCommandlineDataService}
  },
  {
    path: 'configuration/alsa',
    title: 'configuration/alsa',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/alsa/alsa.component').then(m => m.AlsaComponent),
    resolve: {reconnectUpdater: GetConfigAlsaDataService}
  },
  {
    path: 'configuration/amr',
    title: 'configuration/amr',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/amr/amr.component').then(m => m.AmrComponent),
    resolve: {reconnectUpdater: GetConfigAmrDataService}
  },
  {
    path: 'configuration/amrwb',
    title: 'configuration/amrwb',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/amrwb/amrwb.component').then(m => m.AmrwbComponent),
    resolve: {reconnectUpdater: GetConfigAmrwbDataService}
  },
  {
    path: 'configuration/cepstral',
    title: 'configuration/cepstral',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/cepstral/cepstral.component').then(m => m.CepstralComponent),
    resolve: {reconnectUpdater: GetConfigCepstralDataService}
  },
  {
    path: 'configuration/cidlookup',
    title: 'configuration/cidlookup',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/cidlookup/cidlookup.component').then(m => m.CidlookupComponent),
    resolve: {reconnectUpdater: GetConfigCidlookupDataService}
  },
  {
    path: 'configuration/curl',
    title: 'configuration/curl',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/curl/curl.component').then(m => m.CurlComponent),
    resolve: {reconnectUpdater: GetConfigCurlDataService}
  },
  {
    path: 'configuration/dialplan-directory',
    title: 'configuration/dialplan-directory',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/dialplan-directory/dialplan-directory.component').then(m => m.DialplanDirectoryComponent),
    resolve: {reconnectUpdater: GetConfigDialplanDirectoryDataService}
  },
  {
    path: 'configuration/easyroute',
    title: 'configuration/easyroute',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/easyroute/easyroute.component').then(m => m.EasyrouteComponent),
    resolve: {reconnectUpdater: GetConfigEasyrouteDataService}
  },
  {
    path: 'configuration/erlang-event',
    title: 'configuration/erlang-event',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/erlang-event/erlang-event.component').then(m => m.ErlangEventComponent),
    resolve: {reconnectUpdater: GetConfigErlangEventDataService}
  },
  {
    path: 'configuration/event-multicast',
    title: 'configuration/event-multicast',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/event-multicast/event-multicast.component').then(m => m.EventMulticastComponent),
    resolve: {reconnectUpdater: GetConfigEventMulticastDataService}
  },
  {
    path: 'configuration/fax',
    title: 'configuration/fax',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/fax/fax.component').then(m => m.FaxComponent),
    resolve: {reconnectUpdater: GetConfigFaxDataService}
  },
  {
    path: 'configuration/lua',
    title: 'configuration/lua',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/lua/lua.component').then(m => m.LuaComponent),
    resolve: {reconnectUpdater: GetConfigLuaDataService}
  },
  {
    path: 'configuration/mongo',
    title: 'configuration/mongo',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/mongo/mongo.component').then(m => m.MongoComponent),
    resolve: {reconnectUpdater: GetConfigMongoDataService}
  },
  {
    path: 'configuration/msrp',
    title: 'configuration/msrp',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/msrp/msrp.component').then(m => m.MsrpComponent),
    resolve: {reconnectUpdater: GetConfigMsrpDataService}
  },
  {
    path: 'configuration/oreka',
    title: 'configuration/oreka',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/oreka/oreka.component').then(m => m.OrekaComponent),
    resolve: {reconnectUpdater: GetConfigOrekaDataService}
  },
  {
    path: 'configuration/perl',
    title: 'configuration/perl',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/perl/perl.component').then(m => m.PerlComponent),
    resolve: {reconnectUpdater: GetConfigPerlDataService}
  },
  {
    path: 'configuration/pocketsphinx',
    title: 'configuration/pocketsphinx',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/pocketsphinx/pocketsphinx.component').then(m => m.PocketsphinxComponent),
    resolve: {reconnectUpdater: GetConfigPocketsphinxDataService}
  },
  {
    path: 'configuration/sangoma-codec',
    title: 'configuration/sangoma-codec',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/sangoma-codec/sangoma-codec.component').then(m => m.SangomaCodecComponent),
    resolve: {reconnectUpdater: GetConfigSangomaCodecDataService}
  },
  {
    path: 'configuration/sndfile',
    title: 'configuration/sndfile',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/sndfile/sndfile.component').then(m => m.SndfileComponent),
    resolve: {reconnectUpdater: GetConfigSndfileDataService}
  },
  {
    path: 'configuration/xml-cdr',
    title: 'configuration/xml-cdr',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/xml-cdr/xml-cdr.component').then(m => m.XmlCdrComponent),
    resolve: {reconnectUpdater: GetConfigXmlCdrDataService}
  },
  {
    path: 'configuration/xml-rpc',
    title: 'configuration/xml-rpc',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/xml-rpc/xml-rpc.component').then(m => m.XmlRpcComponent),
    resolve: {reconnectUpdater: GetConfigXmlRpcDataService}
  },
  {
    path: 'configuration/zeroconf',
    title: 'configuration/zeroconf',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/zeroconf/zeroconf.component').then(m => m.ZeroconfComponent),
    resolve: {reconnectUpdater: GetConfigZeroconfDataService}
  },
  {
    path: 'configuration/post-load-switch',
    title: 'configuration/post-load-switch',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/post-load-switch/post-load-switch.component').then(m => m.PostLoadSwitchComponent),
    resolve: {reconnectUpdater: GetConfigPostSwitchDataService}
  },
  {
    path: 'configuration/distributor',
    title: 'configuration/distributor',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/distributor/distributor.component').then(m => m.DistributorComponent),
    resolve: {reconnectUpdater: GetConfigDistributorDataService}
  },
  {
    path: 'configuration/directory',
    title: 'configuration/directory',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/directory/directory.component').then(m => m.DirectoryComponent),
    resolve: {reconnectUpdater: GetConfigDirectoryDataService}
  },
  {
    path: 'configuration/fifo',
    title: 'configuration/fifo',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/fifo/fifo.component').then(m => m.FifoComponent),
    resolve: {reconnectUpdater: GetConfigFifoDataService}
  },
  {
    path: 'configuration/opal',
    title: 'configuration/opal',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/opal/opal.component').then(m => m.OpalComponent),
    resolve: {reconnectUpdater: GetConfigOpalDataService}
  },
  {
    path: 'configuration/osp',
    title: 'configuration/osp',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/osp/osp.component').then(m => m.OspComponent),
    resolve: {reconnectUpdater: GetConfigOspDataService}
  },
  {
    path: 'configuration/unicall',
    title: 'configuration/unicall',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/unicall/unicall.component').then(m => m.UnicallComponent),
    resolve: {reconnectUpdater: GetConfigUnicallDataService}
  },
  {
    path: 'configuration/conference',
    title: 'configuration/conference',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/conference/conference.component').then(m => m.ConferenceComponent),
    resolve: {reconnectUpdater: GetConfigConferenceDataService}
  },
  {
    path: 'configuration/post-load-modules',
    title: 'configuration/post-load-modules',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/post-load-modules/post-load-modules.component').then(m => m.PostLoadModulesComponent),
    resolve: {reconnectUpdater: GetPostLoadModulesDataService}
  },
  {
    path: 'configuration/voicemail',
    title: 'configuration/voicemail',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/configuration/voicemail/voicemail.component').then(m => m.VoicemailComponent),
    resolve: {reconnectUpdater: GetVoicemailDataService}
  },

  // --- Dialplan, Monitoring, and Admin Routes ---
  {
    path: 'dialplan/contexts',
    title: 'dialplan/contexts',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/dialplan/contexts/contexts.component').then(m => m.ContextsComponent),
    resolve: {reconnectUpdater: GetDialplanContextsDataService}
  },
  {
    path: 'monitoring/users-panel',
    title: 'monitoring/users-panel',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/monitoring/users-panel/users-panel.component').then(m => m.UsersPanelComponent),
    resolve: {reconnectUpdater: GetDirectoryUsersDataWithSubscriptionService}
  },
  {
    path: 'instances',
    title: 'instances',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/instances/instances.component').then(m => m.InstancesComponent),
    resolve: {reconnectUpdater: GetInstancesDataService}
  },
  {
    path: 'global-variables',
    title: 'global-variables',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/global-variables/global-variables.component').then(m => m.GlobalVariablesComponent),
    resolve: {reconnectUpdater: GetGlobalVariablesDataService}
  },

  // --- Utility/Standalone App Routes ---
  {
    path: 'cdr',
    title: 'cdr',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/cdr/cdr.component').then(m => m.CdrComponent),
    resolve: {reconnectUpdater: UnsubscribeService}
  },
  {
    path: 'logs',
    title: 'logs',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/logs/logs.component').then(m => m.LogsComponent),
    resolve: {reconnectUpdater: UnsubscribeService}
  },
  {
    path: 'fs-cli',
    title: 'fs-cli',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/fs-cli/fs-cli.component').then(m => m.FsCliComponent),
    resolve: {reconnectUpdater: UnsubscribeService}
  },
  {
    path: 'hep',
    title: 'hep',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/hep/hep.component').then(m => m.HepComponent),
    resolve: {reconnectUpdater: UnsubscribeService}
  },
  {
    path: 'apps/autodialer',
    title: 'apps/autodialer',
    canActivate: [AuthGuard],
    loadComponent: () => import('./components/apps/autodialer/autodialer.component').then(m => m.AutodialerComponent),
    resolve: {reconnectUpdater: GetAutodialerDataService}
  },

  // --- Catch-All Route ---
  {
    path: '**',
    title: '**',
    component: NotFoundComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
