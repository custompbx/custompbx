import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {StoreModule} from '@ngrx/store';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginComponent} from './components/login/login.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MaterialModule} from '../material-module';
import {DashboardComponent} from './components/dashboard/dashboard.component';
import {HTTP_INTERCEPTORS, HttpClientModule} from '@angular/common/http';
import {EffectsModule} from '@ngrx/effects';
import {AuthEffects} from './store/auth/auth.effects';
import {reducers} from './store/app.states';
import {ErrorInterceptor} from './services/token.interceptor';
import {AuthGuardService as AuthGuard} from './services/auth-guard.service';
import {NotFoundComponent} from './components/not-found/not-found.component';
import {HeaderComponent} from './components/header/header.component';
import {SidenavComponent} from './components/sidenav/sidenav.component';
import {WebsocketModule} from './services/websocket';
import {LayoutModule} from '@angular/cdk/layout';
import {UsersComponent} from './components/directory/users/users.component';
import {ServiceStatusComponent} from './components/service-status/service-status.component';
import {DomainsComponent} from './components/directory/domains/domains.component';
import {GroupsBottomSheetComponent, GroupsComponent} from './components/directory/groups/groups.component';
import {GatewaysComponent} from './components/directory/gateways/gateways.component';
import {SettingsComponent} from './components/settings/settings.component';
import {SettingsEffects} from './store/settings/settings.effects';
import {GetSettingsDataService} from './services/resolve-route/get-settings-data.service';
import {DirectoryEffects} from './store/directory/directory.effects';
import {GetDirectoryDomainsDataService} from './services/resolve-route/get-directory-domains-data.service';
import {GetDirectoryUsersDataService} from './services/resolve-route/get-directory-users-data.service';
import {LogPipe} from './pipes/log.pipe';
import {GetDirectoryGatewayDataService} from './services/resolve-route/get-directory-gateways-data.service';
import {ContextsComponent, ObjectToNamePipe} from './components/dialplan/contexts/contexts.component';
import {AclComponent} from './components/configuration/acl/acl.component';
import {CallcenterComponent} from './components/configuration/callcenter/callcenter.component';
import {GetConfigAclDataService} from './services/resolve-route/get-config-acl-data.service';
import {ConfigEffects} from './store/config/config.effects';
import {SofiaComponent} from './components/configuration/sofia/sofia.component';
import {ConfirmBottomSheetComponent} from './components/confirm-bottom-sheet/confirm-bottom-sheet.component';
import {ModulesComponent} from './components/configuration/modules/modules.component';
import {GetConfigModulesDataService} from './services/resolve-route/get-config-modules-data.service';
import {GetDialplanContextsDataService} from './services/resolve-route/get-dialplan-contexts-data.service';
import {DialplanEffects} from './store/dialplan/dialplan.effects';
import {GetDataFlowDataService} from './services/resolve-route/get-data-flow-data.service';
import {DataFlowEffects} from './store/dataFlow/dataFlow.effects';
import {NgChartsModule} from 'ng2-charts';
import {CdrPgCsvComponent} from './components/configuration/cdr-pg-csv/cdr-pg-csv.component';
import {GetConfigCdrPgCsvDataService} from './services/resolve-route/get-config-cdr-pg-csv-data.service';
import {CdrComponent} from './components/cdr/cdr.component';
import {CdrEffects} from './store/cdr/cdr.effects';
import {PhoneComponent} from './components/phone/phone.component';
import {PhoneEffects} from './store/phone/phone.effects';
import {environment} from '../environments/environment';
import {ResizeInputDirective} from './directives/resize-input.directive';
import {KeyValuePadComponent} from './components/key-value-pad/key-value-pad.component';
import {UsersPanelComponent} from './components/monitoring/users-panel/users-panel.component';
import {FormatTimerPipe} from './pipes/format-timer.pipe';
import {VertoComponent} from './components/configuration/verto/verto.component';
import {GetConfigCallcenterDataService} from './services/resolve-route/get-config-callcenter-data.service';
import {AppAutoFocusDirective} from './directives/auto-focus.directive';
import {FsCliComponent} from './components/fs-cli/fs-cli.component';
import {FscliEffects} from './store/fscli/fscli.effects';
import {OdbcCdrComponent} from './components/configuration/odbc-cdr/odbc-cdr.component';
import {KeyValuePad2Component} from './components/key-value-pad-2/key-value-pad-2.component';
import {LcrComponent} from './components/configuration/lcr/lcr.component';
import {ShoutComponent} from './components/configuration/shout/shout.component';
import {RedisComponent} from './components/configuration/redis/redis.component';
import {NibblebillComponent} from './components/configuration/nibblebill/nibblebill.component';
import {AvmdComponent} from './components/configuration/avmd/avmd.component';
import {CdrMongodbComponent} from './components/configuration/cdr-mongodb/cdr-mongodb.component';
import {DbComponent} from './components/configuration/db/db.component';
import {HttpCacheComponent} from './components/configuration/http-cache/http-cache.component';
import {MemcacheComponent} from './components/configuration/memcache/memcache.component';
import {OpusComponent} from './components/configuration/opus/opus.component';
import {PythonComponent} from './components/configuration/python/python.component';
import {TtsCommandlineComponent} from './components/configuration/tts-commandline/tts-commandline.component';
import {AlsaComponent} from './components/configuration/alsa/alsa.component';
import {DialplanDirectoryComponent} from './components/configuration/dialplan-directory/dialplan-directory.component';
import {SndfileComponent} from './components/configuration/sndfile/sndfile.component';
import {AmrComponent} from './components/configuration/amr/amr.component';
import {XmlCdrComponent} from './components/configuration/xml-cdr/xml-cdr.component';
import {FaxComponent} from './components/configuration/fax/fax.component';
import {MsrpComponent} from './components/configuration/msrp/msrp.component';
import {OrekaComponent} from './components/configuration/oreka/oreka.component';
import {PocketsphinxComponent} from './components/configuration/pocketsphinx/pocketsphinx.component';
import {CidlookupComponent} from './components/configuration/cidlookup/cidlookup.component';
import {ZeroconfComponent} from './components/configuration/zeroconf/zeroconf.component';
import {AmrwbComponent} from './components/configuration/amrwb/amrwb.component';
import {CepstralComponent} from './components/configuration/cepstral/cepstral.component';
import {XmlRpcComponent} from './components/configuration/xml-rpc/xml-rpc.component';
import {LuaComponent} from './components/configuration/lua/lua.component';
import {PerlComponent} from './components/configuration/perl/perl.component';
import {ErlangEventComponent} from './components/configuration/erlang-event/erlang-event.component';
import {EventMulticastComponent} from './components/configuration/event-multicast/event-multicast.component';
import {SangomaCodecComponent} from './components/configuration/sangoma-codec/sangoma-codec.component';
import {EasyrouteComponent} from './components/configuration/easyroute/easyroute.component';
import {CurlComponent} from './components/configuration/curl/curl.component';
import {MongoComponent} from './components/configuration/mongo/mongo.component';
import {InnerHeaderComponent} from './components/inner-header/inner-header.component';
import {ModuleNotExistsBannerComponent} from './components/configuration/module-not-exists-banner/module-not-exists-banner.component';
import {LogsComponent} from './components/logs/logs.component';
import {LogsEffects} from './store/logs/logs.effects';
import {PostLoadSwitchComponent} from './components/configuration/post-load-switch/post-load-switch.component';
import {ConfigEffectsErlangEvent} from './store/config/erlang_event/config.effects.erlang_event';
import {ConfigEffectsCidlookup} from './store/config/cidlookup/config.effects.cidlookup';
import {ConfigEffectsPerl} from './store/config/perl/config.effects.perl';
import {ConfigEffectsXmlRpc} from './store/config/xml_rpc/config.effects.xml_rpc';
import {ConfigEffectsSndfile} from './store/config/sndfile/config.effects.sndfile';
import {ConfigEffectsMsrp} from './store/config/msrp/config.effects.msrp';
import {ConfigEffectsOreka} from './store/config/oreka/config.effects.oreka';
import {ConfigEffectsEventMulticast} from './store/config/event_multicast/config.effects.event_multicast';
import {ConfigEffectsLua} from './store/config/lua/config.effects.lua';
import {ConfigEffectsCurl} from './store/config/curl/config.effects.curl';
import {ConfigEffectsNibblebill} from './store/config/nibblebill/config.effects.nibblebill';
import {ConfigEffectsZeroconf} from './store/config/zeroconf/config.effects.zeroconf';
import {ConfigEffectsPython} from './store/config/python/config.effects.python';
import {ConfigEffectsXmlCdr} from './store/config/xml_cdr/config.effects.xml_cdr';
import {ConfigEffectsEasyroute} from './store/config/easyroute/config.effects.easyroute';
import {ConfigEffectsOpus} from './store/config/opus/config.effects.opus';
import {ConfigEffectsHttpCache} from './store/config/http_cache/config.effects.http_cache';
import {ConfigEffectsAlsa} from './store/config/alsa/config.effects.alsa';
import {ConfigEffectsRedis} from './store/config/redis/config.effects.redis';
import {ConfigEffectsMongo} from './store/config/mongo/config.effects.mongo';
import {ConfigEffectsTtsCommandline} from './store/config/tts_commandline/config.effects.tts_commandline';
import {ConfigEffectsFax} from './store/config/fax/config.effects.fax';
import {ConfigEffectsAvmd} from './store/config/avmd/config.effects.avmd';
import {ConfigEffectsPocketsphinx} from './store/config/pocketsphinx/config.effects.pocketsphinx';
import {ConfigEffectsDialplanDirectory} from './store/config/dialplan_directory/config.effects.dialplan_directory';
import {ConfigEffectsShout} from './store/config/shout/config.effects.shout';
import {ConfigEffectsDb} from './store/config/db/config.effects.db';
import {ConfigEffectsAmr} from './store/config/amr/config.effects.amr';
import {ConfigEffectsMemcache} from './store/config/memcache/config.effects.memcache';
import {ConfigEffectsCepstral} from './store/config/cepstral/config.effects.cepstral';
import {ConfigEffectsSangomaCodec} from './store/config/sangoma_codec/config.effects.sangoma_codec';
import {ConfigEffectsCdrMongodb} from './store/config/cdr_mongodb/config.effects.cdr_mongodb';
import {ConfigEffectsAmrwb} from './store/config/amrwb/config.effects.amrwb';
import {ConfigEffectsLcr} from './store/config/lcr/config.effects.lcr';
import {ConfigEffectsAcl} from './store/config/acl/config.effects.acl';
import {ConfigEffectsCallcenter} from './store/config/callcenter/config.effects.callcenter';
import {ConfigEffectsSofia} from './store/config/sofia/config.effects.sofia';
import {ConfigEffectsVerto} from './store/config/verto/config.effects.verto';
import {ConfigEffectsOdbcCdr} from './store/config/odbc_cdr/config.effects.odbc-cdr';
import {ConfigEffectsPostSwitch} from './store/config/post-switch/config.effects.post-switch';
import {ConfigEffectsCdrPgCsv} from './store/config/cdr_pg_csv/config.effects.cdr-pg-csv';
import {DaemonEffects} from './store/daemon/daemon.effects';
import {ConfigEffectsDistributor} from './store/config/distributor/config.effects.distributor';
import {DistributorComponent} from './components/configuration/distributor/distributor.component';
import {DirectoryComponent} from './components/configuration/directory/directory.component';
import {FifoComponent} from './components/configuration/fifo/fifo.component';
import {OpalComponent} from './components/configuration/opal/opal.component';
import {OspComponent} from './components/configuration/osp/osp.component';
import {UnicallComponent} from './components/configuration/unicall/unicall.component';
import {ConfigEffectsDirectory} from './store/config/directory/config.effects.directory';
import {ConfigEffectsFifo} from './store/config/fifo/config.effects.fifo';
import {ConfigEffectsOpal} from './store/config/opal/config.effects.opal';
import {ConfigEffectsOsp} from './store/config/osp/config.effects.osp';
import {ConfigEffectsUnicall} from './store/config/unicall/config.effects.unicall';
import {HepEffects} from './store/hep/hep.effects';
import {BottomSheetExportComponent, HepComponent} from './components/hep/hep.component';
import {SvgSeqDiagramComponent} from './components/svg-seq-diagram/svg-seq-diagram.component';
import {ConferenceComponent} from './components/configuration/conference/conference.component';
import {ConfigEffectsConference} from './store/config/conference/config.effects.conference';
import {CodeEditorComponent} from './components/code-editor/code-editor.component';
import {InstancesComponent} from './components/instances/instances.component';
import {EffectsInstances} from './store/instances/instances.effects';
import {GlobalVariablesEffects} from './store/global-variables/global-variables.effects';
import {GlobalVariablesComponent} from './components/global-variables/global-variables.component';
import {PostLoadModulesComponent} from './components/configuration/post-load-modules/post-load-modules.component';
import {ConfigEffectsPostLoadModules} from './store/config/post_load_modules/config.effects.PostLoadModules';
import {LazyWrapperComponent} from './components/lazy-wrapper/lazy-wrapper.component';
import {ConfigEffectsVoicemail} from './store/config/voicemail/config.effects.voicemail';
import {VoicemailComponent} from './components/configuration/voicemail/voicemail.component';
import {AutodialerEffects} from './store/apps/autodialer/autodialer.effects';
import {AutodialerComponent} from './components/apps/autodialer/autodialer.component';
import {KeyValuePadPositionComponent} from './components/key-value-pad-position/key-value-pad-position.component';
import {MAT_FORM_FIELD_DEFAULT_OPTIONS} from '@angular/material/form-field';
import {APP_BASE_HREF} from '@angular/common';
import {MatIconRegistry} from '@angular/material/icon';
import {ConversationsComponent} from './components/conversations/conversations.component';
import {ConversationsEffects} from "./store/conversations/conversations.effects";

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DashboardComponent,
    NotFoundComponent,
    HeaderComponent,
    SidenavComponent,
    UsersComponent,
    GroupsBottomSheetComponent,
    ServiceStatusComponent,
    DomainsComponent,
    GroupsComponent,
    GatewaysComponent,
    SettingsComponent,
    LogPipe,
    ContextsComponent,
    AclComponent,
    CallcenterComponent,
    SofiaComponent,
    VertoComponent,
    ConfirmBottomSheetComponent,
    ModulesComponent,
    ObjectToNamePipe,
    CdrPgCsvComponent,
    CdrComponent,
    PhoneComponent,
    ResizeInputDirective,
    KeyValuePadComponent,
    KeyValuePad2Component,
    KeyValuePadPositionComponent,
    UsersPanelComponent,
    FormatTimerPipe,
    AppAutoFocusDirective,
    FsCliComponent,
    OdbcCdrComponent,
    LcrComponent,
    ShoutComponent,
    RedisComponent,
    NibblebillComponent,
    AvmdComponent,
    CdrMongodbComponent,
    DbComponent,
    HttpCacheComponent,
    MemcacheComponent,
    OpusComponent,
    PythonComponent,
    TtsCommandlineComponent,
    AlsaComponent,
    AmrComponent,
    AmrwbComponent,
    CepstralComponent,
    CidlookupComponent,
    CurlComponent,
    DialplanDirectoryComponent,
    EasyrouteComponent,
    ErlangEventComponent,
    EventMulticastComponent,
    FaxComponent,
    LuaComponent,
    MongoComponent,
    MsrpComponent,
    OrekaComponent,
    PerlComponent,
    PocketsphinxComponent,
    SangomaCodecComponent,
    SndfileComponent,
    XmlCdrComponent,
    XmlRpcComponent,
    ZeroconfComponent,
    InnerHeaderComponent,
    ModuleNotExistsBannerComponent,
    LogsComponent,
    PostLoadSwitchComponent,
    DistributorComponent,
    DirectoryComponent,
    FifoComponent,
    OpalComponent,
    OspComponent,
    UnicallComponent,
    HepComponent,
    SvgSeqDiagramComponent,
    BottomSheetExportComponent,
    ConferenceComponent,
    CodeEditorComponent,
    InstancesComponent,
    GlobalVariablesComponent,
    PostLoadModulesComponent,
    LazyWrapperComponent,
    VoicemailComponent,
    AutodialerComponent,
    ConversationsComponent,
  ],
  imports: [
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    EffectsModule.forRoot([
      AuthEffects,
      SettingsEffects,
      DirectoryEffects,
      DaemonEffects,
      ConfigEffects,
      ConfigEffectsAlsa,
      ConfigEffectsShout,
      ConfigEffectsRedis,
      ConfigEffectsNibblebill,
      ConfigEffectsAvmd,
      ConfigEffectsCdrMongodb,
      ConfigEffectsDb,
      ConfigEffectsMemcache,
      ConfigEffectsOpus,
      ConfigEffectsPython,
      ConfigEffectsTtsCommandline,
      ConfigEffectsHttpCache,
      ConfigEffectsAmr,
      ConfigEffectsAmrwb,
      ConfigEffectsCepstral,
      ConfigEffectsCidlookup,
      ConfigEffectsCurl,
      ConfigEffectsDialplanDirectory,
      ConfigEffectsEasyroute,
      ConfigEffectsErlangEvent,
      ConfigEffectsEventMulticast,
      ConfigEffectsFax,
      ConfigEffectsLua,
      ConfigEffectsMongo,
      ConfigEffectsMsrp,
      ConfigEffectsOreka,
      ConfigEffectsPerl,
      ConfigEffectsPocketsphinx,
      ConfigEffectsSangomaCodec,
      ConfigEffectsSndfile,
      ConfigEffectsXmlCdr,
      ConfigEffectsXmlRpc,
      ConfigEffectsZeroconf,
      DialplanEffects,
      DataFlowEffects,
      CdrEffects,
      LogsEffects,
      PhoneEffects,
      FscliEffects,
      ConfigEffectsLcr,
      ConfigEffectsAcl,
      ConfigEffectsCallcenter,
      ConfigEffectsSofia,
      ConfigEffectsVerto,
      ConfigEffectsOdbcCdr,
      ConfigEffectsPostSwitch,
      ConfigEffectsCdrPgCsv,
      ConfigEffectsDistributor,
      ConfigEffectsDirectory,
      ConfigEffectsFifo,
      ConfigEffectsOpal,
      ConfigEffectsOsp,
      ConfigEffectsUnicall,
      HepEffects,
      ConfigEffectsConference,
      EffectsInstances,
      GlobalVariablesEffects,
      ConfigEffectsPostLoadModules,
      ConfigEffectsVoicemail,
      AutodialerEffects,
      ConversationsEffects,
    ]),
    StoreModule.forRoot({}, {
      runtimeChecks: {
        strictStateImmutability: false,
        strictActionImmutability: false,
        strictStateSerializability: true,
        strictActionSerializability: true,
      }
    }),
    StoreModule.forFeature('app', reducers),
    WebsocketModule.config({
      url: environment.WSServ,
    }),
    LayoutModule,
    NgChartsModule,
  ],
  providers: [
    AuthGuard,
    GetSettingsDataService,
    GetDirectoryDomainsDataService,
    GetDirectoryUsersDataService,
    GetDirectoryGatewayDataService,
    GetConfigAclDataService,
    GetConfigModulesDataService,
    GetDialplanContextsDataService,
    GetDataFlowDataService,
    GetConfigCdrPgCsvDataService,
    GetConfigCallcenterDataService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: ErrorInterceptor,
      multi: true
    },
    {
      provide: MAT_FORM_FIELD_DEFAULT_OPTIONS,
      useValue: {appearance: 'outline'}
    },
    {provide: APP_BASE_HREF, useValue: '/cweb/'}
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(iconRegistry: MatIconRegistry) {
    iconRegistry.setDefaultFontSetClass('material-icons-outlined');
  }
}
