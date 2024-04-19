export interface State {
  errorMessage: string;
  loadCounter: number;
  acl: Iacl;
  sofia: Isofia;
  cdr_pg_csv: IcdrPgCsv;
  verto: Iverto;
  callcenter: Icallcenter;
  odbc_cdr: IodbcCdr;
  lcr: Ilcr;
  shout: IsimpleModule;
  redis: IsimpleModule;
  nibblebill: IsimpleModule;
  avmd: IsimpleModule;
  cdr_mongodb: IsimpleModule;
  db: IsimpleModule;
  memcache: IsimpleModule;
  opus: IsimpleModule;
  python: IsimpleModule;
  tts_commandline: IsimpleModule;
  http_cache: Ihttpcache;
  alsa: IsimpleModule;
  amr: IsimpleModule;
  amrwb: IsimpleModule;
  cepstral: IsimpleModule;
  cidlookup: IsimpleModule;
  curl: IsimpleModule;
  dialplan_directory: IsimpleModule;
  easyroute: IsimpleModule;
  erlang_event: IsimpleModule;
  event_multicast: IsimpleModule;
  fax: IsimpleModule;
  lua: IsimpleModule;
  mongo: IsimpleModule;
  msrp: IsimpleModule;
  oreka: IsimpleModule;
  perl: IsimpleModule;
  pocketsphinx: IsimpleModule;
  sangoma_codec: IsimpleModule;
  sndfile: IsimpleModule;
  xml_cdr: IsimpleModule;
  xml_rpc: IsimpleModule;
  zeroconf: IsimpleModule;
  post_load_switch: IpostSwitcheModule;
  post_load_modules: IpostLoadModules;
  distributor: Idistributor;
  directory: Idirectory;
  fifo: Ififo;
  opal: Iopal;
  osp: Iosp;
  unicall: Iunicall;
  conference: Iconference;
  voicemail: Ivoicemail;
}

export interface Iconference {
  advertise: Irooms;
  caller_controls: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      controls: Icontrols;
    }
  };
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  chat_profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      users: IpermissionUsers;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Ivoicemail {
  parameters: Iparameter;
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface IpermissionUsers {
  [index: number]: IpermissionUser;

  new: Array<object>;
}

export interface IpermissionUser {
  id: number;
  name: string;
  commands: string;
  enabled: boolean;
}

export interface Irooms {
  [index: number]: Iroom;

  new: Array<object>;
}

export interface Iroom {
  id: number;
  name: string;
  state: string;
  enabled: boolean;
}

export interface Icontrols {
  [index: number]: Icontrol;

  new: Array<object>;
}

export interface Icontrol {
  id: number;
  action: string;
  digits: string;
  enabled?: boolean;
}

export interface Isofia {
  global_settings: Iparameter;
  profiles: Iprofiles;
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Iverto {
  settings: Iparameter;
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: IvertoParameter;
      started?: boolean;
      state?: string;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Ihttpcache {
  settings: Iparameter;
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      aws_s3: IhttpcacheS3Item;
      azure: IvertoAzureItem;
      domains: IhttpcacheDomain;
      started?: boolean;
      state?: string;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface IhttpcacheDomain {
  [index: number]: IhttpcacheDomainItem;

  new: Array<object>;
}

export interface IhttpcacheDomainItem {
  id: number;
  name: string;
  enabled: boolean;
}

export interface IhttpcacheS3Item {
  id: number;
  AccessKeyId: string;
  SecretAccessKey: string;
  BaseDomain: string;
  Region: string;
  Expires: number;
  enabled: boolean;
}

export interface IvertoAzureItem {
  id: number;
  SecretAccessKey: string;
  enabled: boolean;
}

export interface Ilcr {
  settings: Iparameter;
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Idirectory {
  settings: Iparameter;
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Ififo {
  settings: Iparameter;
  fifos: {
    [index: number]: {
      id: number;
      importance: string;
      enabled: boolean;
      name: string;
      members: IfifoMembers;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface IfifoMembers {
  [index: number]: IfifoMember;

  new: Array<object>;
}

export interface IfifoMember {
  id: number;
  timeout: string;
  simo: string;
  lag: string;
  body: string;
  enabled: boolean;
}

export interface Iopal {
  settings: Iparameter;
  listeners: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Iosp {
  settings: Iparameter;
  profiles: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Iunicall {
  settings: Iparameter;
  spans: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface IsimpleModule {
  settings: Iparameter;
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface IpostSwitcheModule {
  settings: Iparameter;
  cli_keybindings: Iparameter;
  default_ptimes: IdefaultPtimes;
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface IpostLoadModules {
  modules: Imodules;
  newModules: Array<object>;
  id?: number;
  enabled?: boolean;
  unloadable?: boolean;
  exists?: boolean;
}

export interface Icallcenter {
  settings: Iparameter;
  queues: {
    [index: number]: {
      id: number;
      enabled: boolean;
      name: string;
      parameters: Iparameter;
    }
  };
  agents: {
    table: Array<{ id: number }>;
    list: { [index: number]: object };
    total: number;
  };
  tiers: {
    table: Array<{ id: number }>;
    list: { [index: number]: object };
    total: number;
  };
  members: {
    table: Array<{ uuid: string }>;
    list: { [index: number]: object };
    total: number;
  };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
  changed?: any;
}

export interface Iprofiles {
  [index: number]: Iprofile;
}

export interface Iprofile {
  id: number;
  name: string;
  parameters: Iparameter;
  aliases: Ialiases;
  domains: Idomains;
  gateways: Igateways;
  enabled?: boolean;
  started?: boolean;
  state?: string;
}

export interface Igateways {
  [index: number]: Igateway;

  new: Array<object>;
}

export interface Igateway {
  id: number;
  name: string;
  started?: boolean;
  state?: string;
  parameters: Iparameter;
  variables: Ivariable;
  new: Array<object>;
}

export interface Ialiases {
  [index: number]: Ialias;

  new: Array<object>;
}

export interface Ialias {
  name: string;
  enabled: boolean;
}

export interface Idomains {
  [index: number]: Idomain;

  new: Array<object>;
}

export interface Idomain {
  id: number;
  name: string;
  alias: boolean;
  parse: boolean;
  enabled: boolean;
}

export interface Iparameter {
  [index: number]: Iitem;

  new: Array<object>;
}

export interface IdefaultPtimes {
  [index: number]: IdefaultPtime;

  new: Array<object>;
}

export interface IvertoParameter {
  [index: number]: IvertoParameterItem;

  new: Array<object>;
}

export interface IvertoParameterItem {
  id: number;
  name: string;
  value: string;
  secure: string;
  enabled: boolean;
}

export interface Ivariable {
  [index: number]: IdirectionItem;

  new: Array<object>;
}

export interface IdirectionItem {
  id: number;
  name: string;
  value: string;
  direction: string;
  enabled: boolean;
}

export interface Iitem {
  id: number;
  name: string;
  value: string;
  enabled: boolean;
}

export interface IdefaultPtime {
  id: number;
  name: string;
  ptime: string;
  enabled: boolean;
}

export interface Iacl {
  lists: Ilists;
  id?: number;
  enabled?: boolean;
  unloadable?: boolean;
  exists?: boolean;
  errorMessage: string | null;
}

export interface Ilists {
  [index: number]: {
    id: number;
    nodes: Inodes;
    name: string;
    default: string;
  };
}

export interface Imodules {
  [index: number]: {
    id: number;
    name: string;
    default: string;
  };
}

export interface Inodes {
  [index: number]: Inode;

  new: Array<object>;
}

export interface Inode {
  id: number;
  type: string;
  cidr?: string;
  domain?: string;
  enabled?: boolean;
}

export interface Idistributor {
  lists: IdistributorLists;
  id?: number;
  enabled?: boolean;
  unloadable?: boolean;
  exists?: boolean;
  errorMessage: string | null;
}

export interface IdistributorLists {
  [index: number]: {
    id: number;
    nodes: IdistributorNodes;
    name: string;
    default: string;
  };
}

export interface IdistributorNodes {
  [index: number]: IdistributorNode;

  new: Array<object>;
}

export interface IdistributorNode {
  id: number;
  name: string;
  weight: string;
  enabled?: boolean;
}

export interface IcdrPgCsv {
  settings: { [index: number]: Iitem };
  schema: { [index: number]: Ifield };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  newSettingParams: Array<Iitem>;
  newSchemaFields: Array<Ifield>;
  loaded?: boolean;
}

export interface IodbcCdr {
  settings: Iparameter;
  tables: { [index: number]: Itable };
  id?: number;
  enabled?: boolean;
  exists?: boolean;
  errorMessage: string | null;
  loaded?: boolean;
}

export interface Ifield {
  id: number;
  var: string;
  column: string;
  quote: string;
  enabled: boolean;
}

export interface Itable {
  id: number;
  name: string;
  log_leg: string;
  fields: IODBCFields;
  newField: Array<IodbcField>;
  enabled: boolean;
}

export interface IodbcField {
  id: number;
  name: string;
  chan_var_name: string;
  enabled: boolean;
}

export interface IODBCFields {
  [index: number]: IodbcField;
  new: Array<object>;
}

export const initialState: State = {
  errorMessage: '',
  loadCounter: 0,
  acl: <Iacl>null,
  sofia: <Isofia>null,
  cdr_pg_csv: <IcdrPgCsv>null,
  verto: <Iverto>null,
  callcenter: <Icallcenter>null,
  odbc_cdr: <IodbcCdr>null,
  lcr: <Ilcr>null,
  shout: <IsimpleModule>null,
  redis: <IsimpleModule>null,
  nibblebill: <IsimpleModule>null,
  avmd: <IsimpleModule>null,
  cdr_mongodb: <IsimpleModule>null,
  db: <IsimpleModule>null,
  memcache: <IsimpleModule>null,
  opus: <IsimpleModule>null,
  python: <IsimpleModule>null,
  tts_commandline: <IsimpleModule>null,
  http_cache: <Ihttpcache>null,
  alsa: <IsimpleModule>null,
  amr: <IsimpleModule>null,
  amrwb: <IsimpleModule>null,
  cepstral: <IsimpleModule>null,
  cidlookup: <IsimpleModule>null,
  curl: <IsimpleModule>null,
  dialplan_directory: <IsimpleModule>null,
  easyroute: <IsimpleModule>null,
  erlang_event: <IsimpleModule>null,
  event_multicast: <IsimpleModule>null,
  fax: <IsimpleModule>null,
  lua: <IsimpleModule>null,
  mongo: <IsimpleModule>null,
  msrp: <IsimpleModule>null,
  oreka: <IsimpleModule>null,
  perl: <IsimpleModule>null,
  pocketsphinx: <IsimpleModule>null,
  sangoma_codec: <IsimpleModule>null,
  sndfile: <IsimpleModule>null,
  xml_cdr: <IsimpleModule>null,
  xml_rpc: <IsimpleModule>null,
  zeroconf: <IsimpleModule>null,
  post_load_switch: <IpostSwitcheModule>null,
  distributor: <Idistributor>null,
  directory: <Idirectory>null,
  fifo: <Ififo>null,
  opal: <Iopal>null,
  osp: <Iosp>null,
  unicall: <Iunicall>null,
  conference: <Iconference>null,
  post_load_modules: <IpostLoadModules>null,
  voicemail: <Ivoicemail>null,
};
