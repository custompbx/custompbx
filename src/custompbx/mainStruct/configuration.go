package mainStruct

const (
	ConfSwitch                   = "switch.conf"
	ConfPostLoadSwitch           = "post_load_switch.conf"
	ConfAcl                      = "acl.conf"
	ConfCallcenter               = "callcenter.conf"
	ConfCdrPgCsv                 = "cdr_pg_csv.conf"
	ConfOdbcCdr                  = "odbc_cdr.conf"
	ConfConference               = "conference.conf"
	ConfConferenceLayouts        = "conference_layouts.conf"
	ConfEventSocket              = "event_socket.conf"
	ConfFormatCdr                = "format_cdr.conf"
	ConfHttapi                   = "httapi.conf"
	ConfHttapiCache              = "http_cache.conf"
	ConfIvr                      = "ivr.conf"
	ConfLcr                      = "lcr.conf"
	ConfDirectory                = "directory.conf"
	ConfFifo                     = "fifo.conf"
	ConfOpal                     = "opal.conf"
	ConfOsp                      = "osp.conf"
	ConfUnicall                  = "unicall.conf"
	ConfShout                    = "shout.conf"
	ConfRedis                    = "redis.conf"
	ConfDb                       = "db.conf"
	ConfDistributor              = "distributor.conf"
	ConfMemcache                 = "memcache.conf"
	ConfAvmd                     = "avmd.conf"
	ConfTtsCommandline           = "tts_commandline.conf"
	ConfCdrMongodb               = "cdr_mongodb.conf"
	ConfHttpCache                = "http_cache.conf"
	ConfOpus                     = "opus.conf"
	ConfPython                   = "python.conf"
	ConfAlsa                     = "alsa.conf"
	ConfAmr                      = "amr.conf"
	ConfAmrwb                    = "amrwb.conf"
	ConfCepstral                 = "cepstral.conf"
	ConfCidlookup                = "cidlookup.conf"
	ConfCurl                     = "curl.conf"
	ConfDialplanDirectory        = "dialplan_directory.conf"
	ConfEasyroute                = "easyroute.conf"
	ConfErlangEvent              = "erlang_event.conf"
	ConfEventMulticast           = "event_multicast.conf"
	ConfFax                      = "fax.conf"
	ConfLua                      = "lua.conf"
	ConfMongo                    = "mongo.conf"
	ConfMsrp                     = "msrp.conf"
	ConfOreka                    = "oreka.conf"
	ConfPerl                     = "perl.conf"
	ConfPocketsphinx             = "pocketsphinx.conf"
	ConfSangomaCodec             = "sangoma_codec.conf"
	ConfSndfile                  = "sndfile.conf"
	ConfXmlCdr                   = "xml_cdr.conf"
	ConfXmlRpc                   = "xml_rpc.conf"
	ConfZeroconf                 = "zeroconf.conf"
	ConfLogfile                  = "logfile.conf"
	ConfModules                  = "modules.conf"
	ConfNibblebill               = "nibblebill.conf"
	ConfPostLoadModules          = "post_load_modules.conf"
	ConfSofia                    = "sofia.conf"
	ConfVerto                    = "verto.conf"
	ConfVoicemail                = "voicemail.conf"
	ConfXmlCurl                  = "xml_curl.conf"
	ModAcl                       = "mod_acl"
	ModCallcenter                = "mod_callcenter"
	ModCallcenterAlias           = "callcenter"
	ModCdrPgCsv                  = "mod_cdr_pg_csv"
	ModOdbcCdr                   = "mod_odbc_cdr"
	ModConference                = "mod_conference"
	ModConferenceLayouts         = "mod_conference_layouts"
	ModEventSocket               = "mod_event_socket"
	ModFormatCdr                 = "mod_format_cdr"
	ModHttapi                    = "mod_httapi"
	ModHttapiCache               = "mod_http_cache"
	ModIvr                       = "mod_ivr"
	ModLcr                       = "mod_lcr"
	ModLcrAlias                  = "lcr"
	ModDirectory                 = "mod_directory"
	ModFifo                      = "mod_fifo"
	ModOpal                      = "mod_opal"
	ModOsp                       = "mod_osp"
	ModUnicall                   = "mod_unicall"
	ModShout                     = "mod_shout"
	ModRedis                     = "mod_redis"
	ModDb                        = "mod_db"
	ModDistributor               = "mod_distributor"
	ModMemcache                  = "mod_memcache"
	ModAvmd                      = "mod_avmd"
	ModTtsCommandline            = "mod_tts_commandline"
	ModCdrMongodb                = "mod_cdr_mongodb"
	ModHttpCache                 = "mod_http_cache"
	ModOpus                      = "mod_opus"
	ModPython                    = "mod_python"
	ModAlsa                      = "mod_alsa"
	ModAmr                       = "mod_amr"
	ModAmrwb                     = "mod_amrwb"
	ModCepstral                  = "mod_cepstral"
	ModCidlookup                 = "mod_cidlookup"
	ModCurl                      = "mod_curl"
	ModDialplanDirectory         = "mod_dialplan_directory"
	ModEasyroute                 = "mod_easyroute"
	ModErlangEvent               = "mod_erlang_event"
	ModEventMulticast            = "mod_event_multicast"
	ModFax                       = "mod_fax"
	ModLua                       = "mod_lua"
	ModMongo                     = "mod_mongo"
	ModMsrp                      = "mod_msrp"
	ModOreka                     = "mod_oreka"
	ModPerl                      = "mod_perl"
	ModPocketsphinx              = "mod_pocketsphinx"
	ModSangomaCodec              = "mod_sangoma_codec"
	ModSndfile                   = "mod_sndfile"
	ModXmlCdr                    = "mod_xml_cdr"
	ModXmlRpc                    = "mod_xml_rpc"
	ModZeroconf                  = "mod_zeroconf"
	ModPostLoadSwitch            = "mod_post_load_switch"
	ModLogfile                   = "mod_logfile"
	ModModules                   = "mod_modules"
	ModNibblebill                = "mod_nibblebill"
	ModPostLoadModules           = "mod_post_load_modules"
	ModSofia                     = "mod_sofia"
	ModSofiaAlias                = "sofia"
	ModSwitch                    = "mod_switch"
	ModVerto                     = "mod_verto"
	ModVertoAlias                = "verto"
	ModVoicemail                 = "mod_voicemail"
	ModXmlCurl                   = "xml_curl.conf"
	CommandSofiaProfileStart     = "start"
	CommandSofiaProfileStop      = "stop"
	CommandSofiaProfileRestart   = "restart"
	CommandSofiaProfileRescan    = "rescan"
	CommandSofiaProfileStartgw   = "startgw"
	CommandSofiaProfileKillgw    = "killgw"
	CommandCallcenterQueueLoad   = "load"
	CommandCallcenterQueueReload = "reload"
	CommandCallcenterQueueUnload = "unload"
)

type Configurations struct {
	PostSwitch        *PostLoadSwitch    `xml:"" json:"post_load_switch"`
	Acl               *Acl               `xml:"" json:"acl"`
	Callcenter        *Callcenter        `xml:"" json:"callcenter"`
	CdrPgCsv          *CdrPgCsv          `xml:"" json:"cdr_pg_csv"`
	OdbcCdr           *OdbcCdr           `xml:"" json:"odbc_cdr"`
	Lcr               *Lcr               `xml:"" json:"lcr"`
	Sofia             *Sofia             `xml:"" json:"sofia"`
	Verto             *Verto             `xml:"" json:"verto"`
	Shout             *Shout             `xml:"" json:"shout"`
	Redis             *Redis             `xml:"" json:"redis"`
	Nibblebill        *Nibblebill        `xml:"" json:"nibblebill"`
	Db                *Db                `xml:"" json:"db"`
	Distributor       *Distributor       `xml:"" json:"distributor"`
	Memcache          *Memcache          `xml:"" json:"memcache"`
	Avmd              *Avmd              `xml:"" json:"avmd"`
	TtsCommandline    *TtsCommandline    `xml:"" json:"tts_commandline"`
	CdrMongodb        *CdrMongodb        `xml:"" json:"cdr_mongodb"`
	HttpCache         *HttpCache         `xml:"" json:"http_cache"`
	Opus              *Opus              `xml:"" json:"opus"`
	Python            *Python            `xml:"" json:"python"`
	Alsa              *Alsa              `xml:"" json:"alsa"`
	Amr               *Amr               `xml:"" json:"amr"`
	Amrwb             *Amrwb             `xml:"" json:"amrwb"`
	Cepstral          *Cepstral          `xml:"" json:"cepstral"`
	Cidlookup         *Cidlookup         `xml:"" json:"cidlookup"`
	Curl              *Curl              `xml:"" json:"curl"`
	DialplanDirectory *DialplanDirectory `xml:"" json:"dialplan_directory"`
	Easyroute         *Easyroute         `xml:"" json:"easyroute"`
	ErlangEvent       *ErlangEvent       `xml:"" json:"erlang_event"`
	EventMulticast    *EventMulticast    `xml:"" json:"event_multicast"`
	Fax               *Fax               `xml:"" json:"fax"`
	Lua               *Lua               `xml:"" json:"lua"`
	Mongo             *Mongo             `xml:"" json:"mongo"`
	Msrp              *Msrp              `xml:"" json:"msrp"`
	Oreka             *Oreka             `xml:"" json:"oreka"`
	Perl              *Perl              `xml:"" json:"perl"`
	Pocketsphinx      *Pocketsphinx      `xml:"" json:"pocketsphinx"`
	SangomaCodec      *SangomaCodec      `xml:"" json:"sangoma_codec"`
	Sndfile           *Sndfile           `xml:"" json:"sndfile"`
	XmlCdr            *XmlCdr            `xml:"" json:"xml_cdr"`
	XmlRpc            *XmlRpc            `xml:"" json:"xml_rpc"`
	Zeroconf          *Zeroconf          `xml:"" json:"zeroconf"`
	Directory         *Directory         `xml:"" json:"directory"`
	Fifo              *Fifo              `xml:"" json:"fifo"`
	Opal              *Opal              `xml:"" json:"opal"`
	Osp               *Osp               `xml:"" json:"osp"`
	Unicall           *Unicall           `xml:"" json:"unicall"`
	Conference        *Conference        `xml:"" json:"conference"`
	ConferenceLayouts *ConferenceLayouts `xml:"" json:"conference_layouts"`
	PostLoadModules   *PostLoadModules   `xml:"" json:"post_load_modules"`
	Voicemail         *Voicemail         `xml:"" json:"voicemail"`
}

func GetModulesNames() []string {
	return []string{
		ModCdrPgCsv,
		ModSofia,
		ModAcl,
		ModVerto,
		ModCallcenter,
		ModOdbcCdr,
		ModLcr,
		ModShout,
		ModRedis,
		ModNibblebill,
		ModDb,
		ModMemcache,
		ModAvmd,
		ModTtsCommandline,
		ModCdrMongodb,
		ModHttpCache,
		ModOpus,
		ModPython,
		ModAlsa,
		ModAmr,
		ModAmrwb,
		ModCepstral,
		ModCidlookup,
		ModCurl,
		ModDialplanDirectory,
		ModEasyroute,
		ModErlangEvent,
		ModEventMulticast,
		ModFax,
		ModLua,
		ModMongo,
		ModMsrp,
		ModOreka,
		ModPerl,
		ModPocketsphinx,
		ModSangomaCodec,
		ModSndfile,
		ModXmlCdr,
		ModXmlRpc,
		ModZeroconf,
		ModDistributor,
		ModPostLoadSwitch,
		ModDirectory,
		ModFifo,
		ModOpal,
		ModOsp,
		ModUnicall,
		ModConference,
		ModPostLoadModules,
		ModVoicemail,
	}
}

func IsConfName(name string) bool {
	switch name {
	case ConfCdrPgCsv:
		return true
	case ConfSofia:
		return true
	case ConfAcl:
		return true
	case ConfVerto:
		return true
	case ConfCallcenter:
		return true
	case ConfOdbcCdr:
		return true
	case ConfLcr:
		return true
	case ConfShout:
		return true
	case ConfRedis:
		return true
	case ConfNibblebill:
		return true
	case ConfDb:
		return true
	case ConfMemcache:
		return true
	case ConfAvmd:
		return true
	case ConfTtsCommandline:
		return true
	case ConfCdrMongodb:
		return true
	case ConfHttpCache:
		return true
	case ConfOpus:
		return true
	case ConfPython:
		return true
	case ConfAlsa:
		return true
	case ConfAmr:
		return true
	case ConfAmrwb:
		return true
	case ConfCepstral:
		return true
	case ConfCidlookup:
		return true
	case ConfCurl:
		return true
	case ConfDialplanDirectory:
		return true
	case ConfEasyroute:
		return true
	case ConfErlangEvent:
		return true
	case ConfEventMulticast:
		return true
	case ConfFax:
		return true
	case ConfLua:
		return true
	case ConfMongo:
		return true
	case ConfMsrp:
		return true
	case ConfOreka:
		return true
	case ConfPerl:
		return true
	case ConfPocketsphinx:
		return true
	case ConfSangomaCodec:
		return true
	case ConfSndfile:
		return true
	case ConfXmlCdr:
		return true
	case ConfXmlRpc:
		return true
	case ConfZeroconf:
		return true
	case ConfDistributor:
		return true
	case ConfPostLoadSwitch:
		return true
	case ConfSwitch:
		return true
	case ConfDirectory:
		return true
	case ConfFifo:
		return true
	case ConfOpal:
		return true
	case ConfOsp:
		return true
	case ConfUnicall:
		return true
	case ConfConference:
		return true
	case ConfPostLoadModules:
		return true
	case ConfVoicemail:
		return true
	}
	return false
}

func GetConfNameByModuleName(name string) string {
	switch name {
	case ModCdrPgCsv:
		return ConfCdrPgCsv
	case ModSofiaAlias:
		return ConfSofia
	case ModSofia:
		return ConfSofia
	case ModAcl:
		return ConfAcl
	case ModVerto:
		return ConfVerto
	case ModVertoAlias:
		return ConfVerto
	case ModCallcenter:
		return ConfCallcenter
	case ModCallcenterAlias:
		return ConfCallcenter
	case ModOdbcCdr:
		return ConfOdbcCdr
	case ModLcrAlias:
		return ConfLcr
	case ModLcr:
		return ConfLcr
	case ModShout:
		return ConfShout
	case ModRedis:
		return ConfRedis
	case ModNibblebill:
		return ConfNibblebill
	case ModDb:
		return ConfDb
	case ModMemcache:
		return ConfMemcache
	case ModAvmd:
		return ConfAvmd
	case ModTtsCommandline:
		return ConfTtsCommandline
	case ModCdrMongodb:
		return ConfCdrMongodb
	case ModHttpCache:
		return ConfHttpCache
	case ModOpus:
		return ConfOpus
	case ModPython:
		return ConfPython
	case ModAlsa:
		return ConfAlsa
	case ModAmr:
		return ConfAmr
	case ModAmrwb:
		return ConfAmrwb
	case ModCepstral:
		return ConfCepstral
	case ModCidlookup:
		return ConfCidlookup
	case ModCurl:
		return ConfCurl
	case ModDialplanDirectory:
		return ConfDialplanDirectory
	case ModEasyroute:
		return ConfEasyroute
	case ModErlangEvent:
		return ConfErlangEvent
	case ModEventMulticast:
		return ConfEventMulticast
	case ModFax:
		return ConfFax
	case ModLua:
		return ConfLua
	case ModMongo:
		return ConfMongo
	case ModMsrp:
		return ConfMsrp
	case ModOreka:
		return ConfOreka
	case ModPerl:
		return ConfPerl
	case ModPocketsphinx:
		return ConfPocketsphinx
	case ModSangomaCodec:
		return ConfSangomaCodec
	case ModSndfile:
		return ConfSndfile
	case ModXmlCdr:
		return ConfXmlCdr
	case ModXmlRpc:
		return ConfXmlRpc
	case ModZeroconf:
		return ConfZeroconf
	case ModDistributor:
		return ConfDistributor
	case ModPostLoadSwitch:
		return ConfPostLoadSwitch
	case ModSwitch:
		return ConfSwitch
	case ModDirectory:
		return ConfDirectory
	case ModFifo:
		return ConfFifo
	case ModOpal:
		return ConfOpal
	case ModOsp:
		return ConfOsp
	case ModUnicall:
		return ConfUnicall
	case ModConference:
		return ConfConference
	case ModPostLoadModules:
		return ConfPostLoadModules
	case ModVoicemail:
		return ConfVoicemail
	default:
		return ""
	}
}

func GetModuleNameByConfName(name string) string {
	switch name {
	case ConfCdrPgCsv:
		return ModCdrPgCsv
	case ConfSofia:
		return ModSofia
	case ConfAcl:
		return ModAcl
	case ConfVerto:
		return ModVerto
	case ConfCallcenter:
		return ModCallcenter
	case ConfOdbcCdr:
		return ModOdbcCdr
	case ConfLcr:
		return ModLcr
	case ConfShout:
		return ModShout
	case ConfRedis:
		return ModRedis
	case ConfNibblebill:
		return ModNibblebill
	case ConfDb:
		return ModDb
	case ConfMemcache:
		return ModMemcache
	case ConfAvmd:
		return ModAvmd
	case ConfTtsCommandline:
		return ModTtsCommandline
	case ConfCdrMongodb:
		return ModCdrMongodb
	case ConfHttpCache:
		return ModHttpCache
	case ConfOpus:
		return ModOpus
	case ConfPython:
		return ModPython
	case ConfAlsa:
		return ModAlsa
	case ConfAmr:
		return ModAmr
	case ConfAmrwb:
		return ModAmrwb
	case ConfCepstral:
		return ModCepstral
	case ConfCidlookup:
		return ModCidlookup
	case ConfCurl:
		return ModCurl
	case ConfDialplanDirectory:
		return ModDialplanDirectory
	case ConfEasyroute:
		return ModEasyroute
	case ConfErlangEvent:
		return ModErlangEvent
	case ConfEventMulticast:
		return ModEventMulticast
	case ConfFax:
		return ModFax
	case ConfLua:
		return ModLua
	case ConfMongo:
		return ModMongo
	case ConfMsrp:
		return ModMsrp
	case ConfOreka:
		return ModOreka
	case ConfPerl:
		return ModPerl
	case ConfPocketsphinx:
		return ModPocketsphinx
	case ConfSangomaCodec:
		return ModSangomaCodec
	case ConfSndfile:
		return ModSndfile
	case ConfXmlCdr:
		return ModXmlCdr
	case ConfXmlRpc:
		return ModXmlRpc
	case ConfZeroconf:
		return ModZeroconf
	case ConfDistributor:
		return ModDistributor
	case ConfPostLoadSwitch:
		return ModPostLoadSwitch
	case ConfSwitch:
		return ModSwitch
	case ConfDirectory:
		return ModDirectory
	case ConfFifo:
		return ModFifo
	case ConfOpal:
		return ModOpal
	case ConfOsp:
		return ModOsp
	case ConfUnicall:
		return ModUnicall
	case ConfConference:
		return ModConference
	case ConfPostLoadModules:
		return ModPostLoadModules
	case ConfVoicemail:
		return ModVoicemail
	default:
		return ""
	}
}

func (c *Configurations) TruncateModuleConfig(name string) {
	switch true {
	case name == ModCdrPgCsv && c.CdrPgCsv != nil:
		c.CdrPgCsv = nil
	case name == ModSofia && c.Sofia != nil:
		c.Sofia = nil
	case name == ModAcl && c.Acl != nil:
		c.Acl = nil
	case name == ModVerto && c.Verto != nil:
		c.Verto = nil
	case name == ModCallcenter && c.Callcenter != nil:
		c.Callcenter = nil
	case name == ModOdbcCdr && c.OdbcCdr != nil:
		c.OdbcCdr = nil
	case name == ModLcr && c.Lcr != nil:
		c.Lcr = nil
	case name == ModShout && c.Shout != nil:
		c.Shout = nil
	case name == ModRedis && c.Redis != nil:
		c.Redis = nil
	case name == ModNibblebill && c.Nibblebill != nil:
		c.Nibblebill = nil
	case name == ModDb && c.Db != nil:
		c.Db = nil
	case name == ModMemcache && c.Memcache != nil:
		c.Memcache = nil
	case name == ModAvmd && c.Avmd != nil:
		c.Avmd = nil
	case name == ModTtsCommandline && c.TtsCommandline != nil:
		c.TtsCommandline = nil
	case name == ModCdrMongodb && c.CdrMongodb != nil:
		c.CdrMongodb = nil
	case name == ModHttpCache && c.HttpCache != nil:
		c.HttpCache = nil
	case name == ModOpus && c.Opus != nil:
		c.Opus = nil
	case name == ModPython && c.Python != nil:
		c.Python = nil
	case name == ModAlsa && c.Alsa != nil:
		c.Alsa = nil
	case name == ModAmr && c.Amr != nil:
		c.Amr = nil
	case name == ModAmrwb && c.Amrwb != nil:
		c.Amrwb = nil
	case name == ModCepstral && c.Cepstral != nil:
		c.Cepstral = nil
	case name == ModCidlookup && c.Cidlookup != nil:
		c.Cidlookup = nil
	case name == ModCurl && c.Curl != nil:
		c.Curl = nil
	case name == ModDialplanDirectory && c.DialplanDirectory != nil:
		c.DialplanDirectory = nil
	case name == ModEasyroute && c.Easyroute != nil:
		c.Easyroute = nil
	case name == ModErlangEvent && c.ErlangEvent != nil:
		c.ErlangEvent = nil
	case name == ModEventMulticast && c.EventMulticast != nil:
		c.EventMulticast = nil
	case name == ModFax && c.Fax != nil:
		c.Fax = nil
	case name == ModLua && c.Lua != nil:
		c.Lua = nil
	case name == ModMongo && c.Mongo != nil:
		c.Mongo = nil
	case name == ModMsrp && c.Msrp != nil:
		c.Msrp = nil
	case name == ModOreka && c.Oreka != nil:
		c.Oreka = nil
	case name == ModPerl && c.Perl != nil:
		c.Perl = nil
	case name == ModPocketsphinx && c.Pocketsphinx != nil:
		c.Pocketsphinx = nil
	case name == ModSangomaCodec && c.SangomaCodec != nil:
		c.SangomaCodec = nil
	case name == ModSndfile && c.Sndfile != nil:
		c.Sndfile = nil
	case name == ModXmlCdr && c.XmlCdr != nil:
		c.XmlCdr = nil
	case name == ModXmlRpc && c.XmlRpc != nil:
		c.XmlRpc = nil
	case name == ModZeroconf && c.Zeroconf != nil:
		c.Zeroconf = nil
	case name == ModPostLoadSwitch && c.PostSwitch != nil:
		c.PostSwitch = nil
	case name == ModDistributor && c.Distributor != nil:
		c.Distributor = nil
	case name == ModDirectory && c.Directory != nil:
		c.Directory = nil
	case name == ModFifo && c.Fifo != nil:
		c.Fifo = nil
	case name == ModOpal && c.Opal != nil:
		c.Opal = nil
	case name == ModOsp && c.Osp != nil:
		c.Osp = nil
	case name == ModUnicall && c.Unicall != nil:
		c.Unicall = nil
	case name == ModConference && c.Conference != nil:
		c.Conference = nil
	case name == ModPostLoadModules && c.PostLoadModules != nil:
		c.PostLoadModules = nil
	case name == ModVoicemail && c.Voicemail != nil:
		c.Voicemail = nil
	}
}
