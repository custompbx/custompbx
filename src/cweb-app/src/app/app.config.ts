import {APP_INITIALIZER, ApplicationConfig, importProvidersFrom} from '@angular/core';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideRouter } from '@angular/router';
import { routes } from './app-routing.module';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';

import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';

import { reducers } from './store/app.states';
import { AuthEffects } from './store/auth/auth.effects';
// ... all your effects here

import { MaterialModule } from '../material-module';
import { WebsocketModule } from './services/websocket';
import { environment } from '../environments/environment';
import { MAT_FORM_FIELD_DEFAULT_OPTIONS } from '@angular/material/form-field';
import {APP_BASE_HREF} from "@angular/common";
import {SettingsEffects} from "./store/settings/settings.effects";
import {DirectoryEffects} from "./store/directory/directory.effects";
import {DaemonEffects} from "./store/daemon/daemon.effects";
import {ConfigEffects} from "./store/config/config.effects";
import {ConfigEffectsAlsa} from "./store/config/alsa/config.effects.alsa";
import {ConfigEffectsShout} from "./store/config/shout/config.effects.shout";
import {ConfigEffectsRedis} from "./store/config/redis/config.effects.redis";
import {ConfigEffectsNibblebill} from "./store/config/nibblebill/config.effects.nibblebill";
import {ConfigEffectsAvmd} from "./store/config/avmd/config.effects.avmd";
import {ConfigEffectsCdrMongodb} from "./store/config/cdr_mongodb/config.effects.cdr_mongodb";
import {ConfigEffectsDb} from "./store/config/db/config.effects.db";
import {ConfigEffectsMemcache} from "./store/config/memcache/config.effects.memcache";
import {ConfigEffectsOpus} from "./store/config/opus/config.effects.opus";
import {ConfigEffectsPython} from "./store/config/python/config.effects.python";
import {ConfigEffectsTtsCommandline} from "./store/config/tts_commandline/config.effects.tts_commandline";
import {ConfigEffectsHttpCache} from "./store/config/http_cache/config.effects.http_cache";
import {ConfigEffectsAmr} from "./store/config/amr/config.effects.amr";
import {ConfigEffectsAmrwb} from "./store/config/amrwb/config.effects.amrwb";
import {ConfigEffectsCidlookup} from "./store/config/cidlookup/config.effects.cidlookup";
import {ConfigEffectsCepstral} from "./store/config/cepstral/config.effects.cepstral";
import {ConfigEffectsCurl} from "./store/config/curl/config.effects.curl";
import {ConfigEffectsDialplanDirectory} from "./store/config/dialplan_directory/config.effects.dialplan_directory";
import {ConfigEffectsEasyroute} from "./store/config/easyroute/config.effects.easyroute";
import {ConfigEffectsErlangEvent} from "./store/config/erlang_event/config.effects.erlang_event";
import {ConfigEffectsEventMulticast} from "./store/config/event_multicast/config.effects.event_multicast";
import {ConfigEffectsFax} from "./store/config/fax/config.effects.fax";
import {ConfigEffectsLua} from "./store/config/lua/config.effects.lua";
import {ConfigEffectsMongo} from "./store/config/mongo/config.effects.mongo";
import {ConfigEffectsMsrp} from "./store/config/msrp/config.effects.msrp";
import {ConfigEffectsOreka} from "./store/config/oreka/config.effects.oreka";
import {ConfigEffectsPerl} from "./store/config/perl/config.effects.perl";
import {ConfigEffectsPocketsphinx} from "./store/config/pocketsphinx/config.effects.pocketsphinx";
import {ConfigEffectsSangomaCodec} from "./store/config/sangoma_codec/config.effects.sangoma_codec";
import {ConfigEffectsSndfile} from "./store/config/sndfile/config.effects.sndfile";
import {ConfigEffectsXmlCdr} from "./store/config/xml_cdr/config.effects.xml_cdr";
import {ConfigEffectsXmlRpc} from "./store/config/xml_rpc/config.effects.xml_rpc";
import {ConfigEffectsZeroconf} from "./store/config/zeroconf/config.effects.zeroconf";
import {DialplanEffects} from "./store/dialplan/dialplan.effects";
import {DataFlowEffects} from "./store/dataFlow/dataFlow.effects";
import {CdrEffects} from "./store/cdr/cdr.effects";
import {LogsEffects} from "./store/logs/logs.effects";
import {PhoneEffects} from "./store/phone/phone.effects";
import {FscliEffects} from "./store/fscli/fscli.effects";
import {ConfigEffectsLcr} from "./store/config/lcr/config.effects.lcr";
import {ConfigEffectsAcl} from "./store/config/acl/config.effects.acl";
import {ConfigEffectsCallcenter} from "./store/config/callcenter/config.effects.callcenter";
import {EffectsInstances} from "./store/instances/instances.effects";
import {ConfigEffectsSofia} from "./store/config/sofia/config.effects.sofia";
import {ConfigEffectsVerto} from "./store/config/verto/config.effects.verto";
import {ConfigEffectsOdbcCdr} from "./store/config/odbc_cdr/config.effects.odbc-cdr";
import {ConfigEffectsPostSwitch} from "./store/config/post-switch/config.effects.post-switch";
import {ConfigEffectsCdrPgCsv} from "./store/config/cdr_pg_csv/config.effects.cdr-pg-csv";
import {ConfigEffectsDistributor} from "./store/config/distributor/config.effects.distributor";
import {ConfigEffectsDirectory} from "./store/config/directory/config.effects.directory";
import {ConfigEffectsFifo} from "./store/config/fifo/config.effects.fifo";
import {ConfigEffectsOpal} from "./store/config/opal/config.effects.opal";
import {ConfigEffectsOsp} from "./store/config/osp/config.effects.osp";
import {ConfigEffectsUnicall} from "./store/config/unicall/config.effects.unicall";
import {HepEffects} from "./store/hep/hep.effects";
import {ConfigEffectsConference} from "./store/config/conference/config.effects.conference";
import {GlobalVariablesEffects} from "./store/global-variables/global-variables.effects";
import {ConfigEffectsPostLoadModules} from "./store/config/post_load_modules/config.effects.PostLoadModules";
import {ConfigEffectsVoicemail} from "./store/config/voicemail/config.effects.voicemail";
import {AutodialerEffects} from "./store/apps/autodialer/autodialer.effects";
import {ConversationsEffects} from "./store/conversations/conversations.effects";
import {provideCharts, withDefaultRegisterables} from "ng2-charts";
import {MatIconRegistry} from "@angular/material/icon";

export const appConfig: ApplicationConfig = {
  providers: [

    provideRouter(routes),
    provideAnimations(),

    importProvidersFrom(
      MaterialModule,
      WebsocketModule.config({ url: environment.WSServ }),

      StoreModule.forRoot({}, {
        runtimeChecks: {
          strictStateImmutability: false,
          strictActionImmutability: false,
          strictStateSerializability: true,
          strictActionSerializability: true,
        }

      }),

      StoreModule.forFeature('app', reducers),

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
    ),

    provideHttpClient(withInterceptorsFromDi()),

    {
      provide: MAT_FORM_FIELD_DEFAULT_OPTIONS,
      useValue: { appearance: 'outline' }
    },

    { provide: APP_BASE_HREF, useValue: '/cweb/' },
    {
      provide: APP_INITIALIZER,
      multi: true,
      useFactory: (iconRegistry: MatIconRegistry) => {
        return () => {
          iconRegistry.setDefaultFontSetClass('material-icons-outlined');
        };
      },
      deps: [MatIconRegistry]
    }
  ]
};
