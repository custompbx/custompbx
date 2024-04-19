import {
  ConfigActionTypes,
} from './config.actions';
import {reducer as sofiaReducer} from './sofia/config.reducer.sofia';
import {reducer as vertoReducer} from './verto/config.reducer.verto';
import {reducer as cdrPgCsvReducer} from './cdr_pg_csv/config.reducer.cdr-pg-csv';
import {reducer as aclReducer} from './acl/config.reducer.acl';
import {reducer as callcenterReducer} from './callcenter/config.reducer.callcenter';
import {reducer as odbcCdrReducer} from './odbc_cdr/config.reducer.odbc-cdr';
import {reducer as lcrReducer} from './lcr/config.reducer.lcr';
import {reducer as httpCacheReducer} from './http_cache/config.reducer.http_cache';
import {reducer as postSwitchReducer} from './post-switch/config.reducer.post-switch';
import {reducer as distributorReducer} from './distributor/config.reducer.distributor';
import {reducer as alsaReducer} from './alsa/config.reducer.alsa';
import {reducer as shoutReducer} from './shout/config.reducer.shout';
import {reducer as redisReducer} from './redis/config.reducer.redis';
import {reducer as nibblebillReducer} from './nibblebill/config.reducer.nibblebill';
import {reducer as avmdReducer} from './avmd/config.reducer.avmd';
import {reducer as cdrMongodbReducer} from './cdr_mongodb/config.reducer.cdr_mongodb';
import {reducer as dbReducer} from './db/config.reducer.db';
import {reducer as memcacheReducer} from './memcache/config.reducer.memcache';
import {reducer as opusReducer} from './opus/config.reducer.opus';
import {reducer as pythonReducer} from './python/config.reducer.python';
import {reducer as ttsCommandlineReducer} from './tts_commandline/config.reducer.tts_commandline';
import {reducer as amrReducer} from './amr/config.reducer.amr';
import {reducer as amrwbReducer} from './amrwb/config.reducer.amrwb';
import {reducer as cepstralReducer} from './cepstral/config.reducer.cepstral';
import {reducer as cidlookupReducer} from './cidlookup/config.reducer.cidlookup';
import {reducer as curlReducer} from './curl/config.reducer.curl';
import {reducer as dialplanDirectoryReducer} from './dialplan_directory/config.reducer.dialplan_directory';
import {reducer as easyrouteReducer} from './easyroute/config.reducer.easyroute';
import {reducer as erlangEventReducer} from './erlang_event/config.reducer.erlang_event';
import {reducer as eventMulticastReducer} from './event_multicast/config.reducer.event_multicast';
import {reducer as faxReducer} from './fax/config.reducer.fax';
import {reducer as luaReducer} from './lua/config.reducer.lua';
import {reducer as mongoReducer} from './mongo/config.reducer.mongo';
import {reducer as msrpReducer} from './msrp/config.reducer.msrp';
import {reducer as orekaReducer} from './oreka/config.reducer.oreka';
import {reducer as perlReducer} from './perl/config.reducer.perl';
import {reducer as pocketsphinxReducer} from './pocketsphinx/config.reducer.pocketsphinx';
import {reducer as sangomaCodecReducer} from './sangoma_codec/config.reducer.sangoma_codec';
import {reducer as sndfileReducer} from './sndfile/config.reducer.sndfile';
import {reducer as xmlCdrReducer} from './xml_cdr/config.reducer.xml_cdr';
import {reducer as xmlRpcReducer} from './xml_rpc/config.reducer.xml_rpc';
import {reducer as zeroconfReducer} from './zeroconf/config.reducer.zeroconf';
import {reducer as directoryReducer} from './directory/config.reducer.directory';
import {reducer as fifoReducer} from './fifo/config.reducer.fifo';
import {reducer as opalReducer} from './opal/config.reducer.opal';
import {reducer as ospReducer} from './osp/config.reducer.osp';
import {reducer as unicallReducer} from './unicall/config.reducer.unicall';
import {reducer as conferenceReducer} from './conference/config.reducer.conference';
import {reducer as postLoadReducer} from './post_load_modules/config.reducer.PostLoadModules';
import {reducer as voicemail} from './voicemail/config.reducer.voicemail';

import {
  initialState, Iverto,
  State
} from './config.state.struct';

export function reducer(state = initialState, action: any): State {
  const sofiaState = sofiaReducer(state, action);
  if (sofiaState) {
    return sofiaState;
  }
  const vertoState = vertoReducer(state, action);
  if (vertoState) {
    return vertoState;
  }
  const cdrPgCsvState = cdrPgCsvReducer(state, action);
  if (cdrPgCsvState) {
    return cdrPgCsvState;
  }
  const aclState = aclReducer(state, action);
  if (aclState) {
    return aclState;
  }
  const callcenterState = callcenterReducer(state, action);
  if (callcenterState) {
    return callcenterState;
  }
  const odbcCdrState = odbcCdrReducer(state, action);
  if (odbcCdrState) {
    return odbcCdrState;
  }
  const lcrState = lcrReducer(state, action);
  if (lcrState) {
    return lcrState;
  }
  const shoutState = shoutReducer(state, action);
  if (shoutState) {
    return shoutState;
  }
  const redisState = redisReducer(state, action);
  if (redisState) {
    return redisState;
  }
  const nibblebillState = nibblebillReducer(state, action);
  if (nibblebillState) {
    return nibblebillState;
  }
  const avmdState = avmdReducer(state, action);
  if (avmdState) {
    return avmdState;
  }
  const cdrMongodbState = cdrMongodbReducer(state, action);
  if (cdrMongodbState) {
    return cdrMongodbState;
  }
  const dbState = dbReducer(state, action);
  if (dbState) {
    return dbState;
  }
  const httpCacheState = httpCacheReducer(state, action);
  if (httpCacheState) {
    return httpCacheState;
  }
  const memcacheState = memcacheReducer(state, action);
  if (memcacheState) {
    return memcacheState;
  }
  const opusState = opusReducer(state, action);
  if (opusState) {
    return opusState;
  }
  const pythonState = pythonReducer(state, action);
  if (pythonState) {
    return pythonState;
  }
  const ttsCommandlineState = ttsCommandlineReducer(state, action);
  if (ttsCommandlineState) {
    return ttsCommandlineState;
  }
  const alsaState = alsaReducer(state, action);
  if (alsaState) {
    return alsaState;
  }
  const amrState = amrReducer(state, action);
  if (amrState) {
    return amrState;
  }
  const amrwbState = amrwbReducer(state, action);
  if (amrwbState) {
    return amrwbState;
  }
  const cepstralState = cepstralReducer(state, action);
  if (cepstralState) {
    return cepstralState;
  }
  const cidlookupState = cidlookupReducer(state, action);
  if (cidlookupState) {
    return cidlookupState;
  }
  const curlState = curlReducer(state, action);
  if (curlState) {
    return curlState;
  }
  const dialplanDirectoryState = dialplanDirectoryReducer(state, action);
  if (dialplanDirectoryState) {
    return dialplanDirectoryState;
  }
  const easyrouteState = easyrouteReducer(state, action);
  if (easyrouteState) {
    return easyrouteState;
  }
  const erlangEventState = erlangEventReducer(state, action);
  if (erlangEventState) {
    return erlangEventState;
  }
  const eventMulticastState = eventMulticastReducer(state, action);
  if (eventMulticastState) {
    return eventMulticastState;
  }
  const faxState = faxReducer(state, action);
  if (faxState) {
    return faxState;
  }
  const luaState = luaReducer(state, action);
  if (luaState) {
    return luaState;
  }
  const mongoState = mongoReducer(state, action);
  if (mongoState) {
    return mongoState;
  }
  const msrpState = msrpReducer(state, action);
  if (msrpState) {
    return msrpState;
  }
  const orekaState = orekaReducer(state, action);
  if (orekaState) {
    return orekaState;
  }
  const perlState = perlReducer(state, action);
  if (perlState) {
    return perlState;
  }
  const pocketsphinxState = pocketsphinxReducer(state, action);
  if (pocketsphinxState) {
    return pocketsphinxState;
  }
  const sangomaCodecState = sangomaCodecReducer(state, action);
  if (sangomaCodecState) {
    return sangomaCodecState;
  }
  const sndfileState = sndfileReducer(state, action);
  if (sndfileState) {
    return sndfileState;
  }
  const xmlCdrState = xmlCdrReducer(state, action);
  if (xmlCdrState) {
    return xmlCdrState;
  }
  const xmlRpcState = xmlRpcReducer(state, action);
  if (xmlRpcState) {
    return xmlRpcState;
  }
  const zeroconfState = zeroconfReducer(state, action);
  if (zeroconfState) {
    return zeroconfState;
  }
  const postSwitchState = postSwitchReducer(state, action);
  if (postSwitchState) {
    return postSwitchState;
  }
  const distributorState = distributorReducer(state, action);
  if (distributorState) {
    return distributorState;
  }
  const directoryState = directoryReducer(state, action);
  if (directoryState) {
    return directoryState;
  }
  const fifoState = fifoReducer(state, action);
  if (fifoState) {
    return fifoState;
  }
  const opalState = opalReducer(state, action);
  if (opalState) {
    return opalState;
  }
  const opspState = ospReducer(state, action);
  if (opspState) {
    return opspState;
  }
  const unicallState = unicallReducer(state, action);
  if (unicallState) {
    return unicallState;
  }
  const conferenceState = conferenceReducer(state, action);
  if (conferenceState) {
    return conferenceState;
  }

  const postLoadModulesState = postLoadReducer(state, action);
  if (postLoadModulesState) {
    return postLoadModulesState;
  }

  const voicemailState = voicemail(state, action);
  if (voicemailState) {
    return voicemailState;
  }

  switch (action.type) {
    case ConfigActionTypes.UPDATE_FAILURE: {
      return {
        ...state,
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
        // errorMessage: 'Cant get data from server',
      };
    }

    case ConfigActionTypes.ImportXMLModuleConfig:
    case ConfigActionTypes.TruncateModuleConfig:
    case ConfigActionTypes.GET_MODULES:
    case ConfigActionTypes.RELOAD_MODULE:
    case ConfigActionTypes.UNLOAD_MODULE:
    case ConfigActionTypes.LOAD_MODULE:
    case ConfigActionTypes.SWITCH_MODULE:
    case ConfigActionTypes.IMPORT_MODULE:
    case ConfigActionTypes.FROM_SCRATCH_MODULE:
    case ConfigActionTypes.AUTOLOAD_MODULE:
    case ConfigActionTypes.IMPORT_ALL_MODULES: {
      return {
        ...state,
        errorMessage: '', loadCounter: state.loadCounter + 1
      };
    }

    case ConfigActionTypes.StoreGotModuleError: {
      return {
        ...state,
        errorMessage: action.payload.error || '',
        loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0,
      };
    }

    case ConfigActionTypes.StoreImportXMLModuleConfig:
    case ConfigActionTypes.StoreTruncateModuleConfig:
    case ConfigActionTypes.STORE_SWITCH_MODULE:
    case ConfigActionTypes.STORE_GET_MODULES: {
      if (action.payload.response.modules) {
        return {
          ...state,
          acl: action.payload.response.modules.acl !== null ?
            {...state.acl, ...action.payload.response.modules.acl} : null,
          sofia: action.payload.response.modules.sofia !== null ?
            {...state.sofia, ...action.payload.response.modules.sofia} : null,
          cdr_pg_csv: action.payload.response.modules.cdr_pg_csv !== null ?
            {...state.cdr_pg_csv, ...action.payload.response.modules.cdr_pg_csv} : null,
          odbc_cdr: action.payload.response.modules.odbc_cdr !== null ?
            {...state.odbc_cdr, ...action.payload.response.modules.odbc_cdr} : null,
          callcenter: action.payload.response.modules.callcenter !== null ?
            {...state.callcenter, ...action.payload.response.modules.callcenter} : null,
          verto: action.payload.response.modules.verto !== null ?
            {...state.verto, ...action.payload.response.modules.verto} : null,
          lcr: action.payload.response.modules.lcr !== null ?
            {...state.lcr, ...action.payload.response.modules.lcr} : null,
          shout: action.payload.response.modules.shout !== null ?
            {...state.shout, ...action.payload.response.modules.shout} : null,
          redis: action.payload.response.modules.redis !== null ?
            {...state.redis, ...action.payload.response.modules.redis} : null,
          nibblebill: action.payload.response.modules.nibblebill !== null ?
            {...state.nibblebill, ...action.payload.response.modules.nibblebill} : null,
          avmd: action.payload.response.modules.avmd !== null ?
            {...state.avmd, ...action.payload.response.modules.avmd} : null,
          cdr_mongodb: action.payload.response.modules.cdr_mongodb !== null ?
            {...state.cdr_mongodb, ...action.payload.response.modules.cdr_mongodb} : null,
          db: action.payload.response.modules.db !== null ?
            {...state.db, ...action.payload.response.modules.db} : null,
          memcache: action.payload.response.modules.memcache !== null ?
            {...state.memcache, ...action.payload.response.modules.memcache} : null,
          opus: action.payload.response.modules.opus !== null ?
            {...state.opus, ...action.payload.response.modules.opus} : null,
          python: action.payload.response.modules.python !== null ?
            {...state.python, ...action.payload.response.modules.python} : null,
          tts_commandline: action.payload.response.modules.tts_commandline !== null ?
            {...state.tts_commandline, ...action.payload.response.modules.tts_commandline} : null,
          http_cache: action.payload.response.modules.http_cache !== null ?
            {...state.http_cache, ...action.payload.response.modules.http_cache} : null,
          alsa: action.payload.response.modules.alsa !== null ?
            {...state.alsa, ...action.payload.response.modules.alsa} : null,
          amr: action.payload.response.modules.amr !== null ?
            {...state.amr, ...action.payload.response.modules.amr} : null,
          amrwb: action.payload.response.modules.amrwb !== null ?
            {...state.amrwb, ...action.payload.response.modules.amrwb} : null,
          cepstral: action.payload.response.modules.cepstral !== null ?
            {...state.cepstral, ...action.payload.response.modules.cepstral} : null,
          cidlookup: action.payload.response.modules.cidlookup !== null ?
            {...state.cidlookup, ...action.payload.response.modules.cidlookup} : null,
          curl: action.payload.response.modules.curl !== null ?
            {...state.curl, ...action.payload.response.modules.curl} : null,
          dialplan_directory: action.payload.response.modules.dialplan_directory !== null ?
            {...state.dialplan_directory, ...action.payload.response.modules.dialplan_directory} : null,
          easyroute: action.payload.response.modules.easyroute !== null ?
            {...state.easyroute, ...action.payload.response.modules.easyroute} : null,
          erlang_event: action.payload.response.modules.erlang_event !== null ?
            {...state.erlang_event, ...action.payload.response.modules.erlang_event} : null,
          event_multicast: action.payload.response.modules.event_multicast !== null ?
            {...state.event_multicast, ...action.payload.response.modules.event_multicast} : null,
          fax: action.payload.response.modules.fax !== null ?
            {...state.fax, ...action.payload.response.modules.fax} : null,
          lua: action.payload.response.modules.lua !== null ?
            {...state.lua, ...action.payload.response.modules.lua} : null,
          mongo: action.payload.response.modules.mongo !== null ?
            {...state.mongo, ...action.payload.response.modules.mongo} : null,
          msrp: action.payload.response.modules.msrp !== null ?
            {...state.msrp, ...action.payload.response.modules.msrp} : null,
          oreka: action.payload.response.modules.oreka !== null ?
            {...state.oreka, ...action.payload.response.modules.oreka} : null,
          perl: action.payload.response.modules.perl !== null ?
            {...state.perl, ...action.payload.response.modules.perl} : null,
          pocketsphinx: action.payload.response.modules.pocketsphinx !== null ?
            {...state.pocketsphinx, ...action.payload.response.modules.pocketsphinx} : null,
          sangoma_codec: action.payload.response.modules.sangoma_codec !== null ?
            {...state.sangoma_codec, ...action.payload.response.modules.sangoma_codec} : null,
          sndfile: action.payload.response.modules.sndfile !== null ?
            {...state.sndfile, ...action.payload.response.modules.sndfile} : null,
          xml_cdr: action.payload.response.modules.xml_cdr !== null ?
            {...state.xml_cdr, ...action.payload.response.modules.xml_cdr} : null,
          xml_rpc: action.payload.response.modules.xml_rpc !== null ?
            {...state.xml_rpc, ...action.payload.response.modules.xml_rpc} : null,
          zeroconf: action.payload.response.modules.zeroconf !== null ?
            {...state.zeroconf, ...action.payload.response.modules.zeroconf} : null,
          distributor: action.payload.response.modules.distributor !== null ?
            {...state.distributor, ...action.payload.response.modules.distributor} : null,
          post_load_switch: action.payload.response.modules.post_load_switch !== null ?
            {...state.post_load_switch, ...action.payload.response.modules.post_load_switch} : null,
          directory: action.payload.response.modules.directory !== null ?
            {...state.directory, ...action.payload.response.modules.directory} : null,
          fifo: action.payload.response.modules.fifo !== null ?
            {...state.fifo, ...action.payload.response.modules.fifo} : null,
          opal: action.payload.response.modules.opal !== null ?
            {...state.opal, ...action.payload.response.modules.opal} : null,
          osp: action.payload.response.modules.osp !== null ?
            {...state.osp, ...action.payload.response.modules.osp} : null,
          unicall: action.payload.response.modules.unicall !== null ?
            {...state.unicall, ...action.payload.response.modules.unicall} : null,
          conference: action.payload.response.modules.conference !== null ?
            {...state.conference, ...action.payload.response.modules.conference} : null,
          post_load_modules: action.payload.response.modules.post_load_modules !== null ?
            {...state.post_load_modules, ...action.payload.response.modules.post_load_modules} : null,
          voicemail: action.payload.response.modules.voicemail !== null ?
            {...state.voicemail, ...action.payload.response.modules.voicemail} : null,
          errorMessage: '',
          loadCounter: 0,
        };
      }
      if (action.payload.response.module) {
        return {
          ...state,
          acl: {...state.acl, ...action.payload.response.module.acl},
          sofia: {...state.sofia, ...action.payload.response.module.sofia},
          cdr_pg_csv: {...state.cdr_pg_csv, ...action.payload.response.module.cdr_pg_csv},
          odbc_cdr: {...state.odbc_cdr, ...action.payload.response.module.odbc_cdr},
          callcenter: {...state.callcenter, ...action.payload.response.module.callcenter},
          verto: {...state.verto, ...action.payload.response.module.verto},
          lcr: {...state.lcr, ...action.payload.response.module.lcr},
          shout: {...state.shout, ...action.payload.response.module.shout},
          redis: {...state.redis, ...action.payload.response.module.redis},
          nibblebill: {...state.nibblebill, ...action.payload.response.module.nibblebill},
          avmd: {...state.avmd, ...action.payload.response.module.avmd},
          cdr_mongodb: {...state.cdr_mongodb, ...action.payload.response.module.cdr_mongodb},
          db: {...state.db, ...action.payload.response.module.db},
          memcache: {...state.memcache, ...action.payload.response.module.memcache},
          opus: {...state.opus, ...action.payload.response.module.opus},
          python: {...state.python, ...action.payload.response.module.python},
          tts_commandline: {...state.tts_commandline, ...action.payload.response.module.tts_commandline},
          http_cache: {...state.http_cache, ...action.payload.response.module.http_cache},
          alsa: {...state.alsa, ...action.payload.response.module.alsa},
          amr: {...state.amr, ...action.payload.response.module.amr},
          amrwb: {...state.amrwb, ...action.payload.response.module.amrwb},
          cepstral: {...state.cepstral, ...action.payload.response.module.cepstral},
          cidlookup: {...state.cidlookup, ...action.payload.response.module.cidlookup},
          curl: {...state.curl, ...action.payload.response.module.curl},
          dialplan_directory: {...state.dialplan_directory, ...action.payload.response.module.dialplan_directory},
          easyroute: {...state.easyroute, ...action.payload.response.module.easyroute},
          erlang_event: {...state.erlang_event, ...action.payload.response.module.erlang_event},
          event_multicast: {...state.event_multicast, ...action.payload.response.module.event_multicast},
          fax: {...state.fax, ...action.payload.response.module.fax},
          lua: {...state.lua, ...action.payload.response.module.lua},
          mongo: {...state.mongo, ...action.payload.response.module.mongo},
          msrp: {...state.msrp, ...action.payload.response.module.msrp},
          oreka: {...state.oreka, ...action.payload.response.module.oreka},
          perl: {...state.perl, ...action.payload.response.module.perl},
          pocketsphinx: {...state.pocketsphinx, ...action.payload.response.module.pocketsphinx},
          sangoma_codec: {...state.sangoma_codec, ...action.payload.response.module.sangoma_codec},
          sndfile: {...state.sndfile, ...action.payload.response.module.sndfile},
          xml_cdr: {...state.xml_cdr, ...action.payload.response.module.xml_cdr},
          xml_rpc: {...state.xml_rpc, ...action.payload.response.module.xml_rpc},
          zeroconf: {...state.zeroconf, ...action.payload.response.module.zeroconf},
          distributor: {...state.distributor, ...action.payload.response.module.distributor},
          post_load_switch: {...state.post_load_switch, ...action.payload.response.module.post_load_switch},
          directory: {...state.directory, ...action.payload.response.module.directory},
          fifo: {...state.fifo, ...action.payload.response.module.fifo},
          opal: {...state.opal, ...action.payload.response.module.opal},
          osp: {...state.osp, ...action.payload.response.module.osp},
          unicall: {...state.unicall, ...action.payload.response.module.unicall},
          conference: {...state.conference, ...action.payload.response.module.conference},
          post_load_modules: {...state.post_load_modules, ...action.payload.response.module.post_load_modules},
          voicemail: {...state.voicemail, ...action.payload.response.module.voicemail},
          errorMessage: '',
          loadCounter: 0,
        };
      }
      return {...state, loadCounter: state.loadCounter > 0 ? --state.loadCounter : 0};
    }

    default: {
      return state;
    }
  }
}

export function increaseStateLoadField(state: State): State {
  return {
    ...state,
    loadCounter: state.loadCounter++,
  };
}
export function decreaseStateLoadField(state: State): State {
  return {
    ...state,
    loadCounter: Math.max(state.loadCounter - 1, 0),
  };
}
export function updateStateItem<State>(fieldName: keyof State, fieldValue: State[keyof State], state: State): State {
  return {
    ...state,
    [fieldName]: {
      ...state[fieldName],
      ...fieldValue,
    },
  };
}
export function updateStateWithFieldValue(objectKey: string, fieldName: string, fieldValue: any, state: State): State {
  if (state[objectKey] && typeof state[objectKey] === 'object') {
    return {
      ...state,
      [objectKey]: {
        ...state[objectKey],
        [fieldName]: fieldValue,
      },
    };
  }
  return state;
}
export function removeFromObject<T extends object>(obj: T, propertyName: keyof T) {
  const { [propertyName]: toDel, ...rest } = obj;
  return rest;
}
export function updateNestedState<State>(
  state: State,
  updates: { path: string[], value: any }[]
): State {
  let newState = { ...state };

  updates.forEach(update => {
    const { path, value } = update;
    let current = newState;

    for (let i = 0; i < path.length - 1; i++) {
      const key = path[i];
      if (current.hasOwnProperty(key) && typeof current[key] === 'object' && current[key] !== null) {
        current = current[key];
      } else {
        // Create intermediate objects if they don't exist
        current[key] = {};
        current = current[key];
      }
    }

    current[path[path.length - 1]] = value;
  });

  return newState;
}

export function getParentId(data: any): number {
  if (data.id) {
    return data?.parent?.id || 0;
  } else {
    const firstKey = Object.keys(data)[0];
    return data[firstKey]?.parent?.id || 0;
  }
}
