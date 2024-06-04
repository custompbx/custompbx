package altStruct

import (
	"custompbx/mainStruct"
)

type ConfigurationsList struct {
	Id          int64                  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                  `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                   `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                 `xml:"-" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Module      string                 `xml:"-" json:"module"`
	Loaded      bool                   `xml:"-" json:"loaded"`
	Unloadable  bool                   `xml:"-" json:"unloadable" customsql:"unloadable;default=FALSE"`
	Description string                 `xml:"-" json:"description,omitempty" customsql:"description"`
	Parent      *mainStruct.FsInstance `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigurationsList) GetTableName() string {
	return "configurations_list"
}

func (c *Configurations) FillConfigurations(conf *ConfigurationsList) {
	if conf == nil {
		return
	}
	c.GetConfigurationAndUpdate(conf.Name, conf)
}

func (c *Configurations) GetConfiguration(name string) *ConfigurationsList {
	return c.GetConfigurationAndUpdate(name, nil)
}

func (c *Configurations) GetConfigurationAndUpdate(name string, newConf *ConfigurationsList) *ConfigurationsList {
	switch name {
	case mainStruct.ConfPostLoadSwitch:
		if newConf != nil {
			c.PostSwitch = newConf
		}
		return c.PostSwitch
	case mainStruct.ConfAcl:
		if newConf != nil {
			c.Acl = newConf
		}
		return c.Acl
	case mainStruct.ConfCallcenter:
		if newConf != nil {
			c.Callcenter = newConf
		}
		return c.Callcenter
	case mainStruct.ConfCdrPgCsv:
		if newConf != nil {
			c.CdrPgCsv = newConf
		}
		return c.CdrPgCsv
	case mainStruct.ConfOdbcCdr:
		if newConf != nil {
			c.OdbcCdr = newConf
		}
		return c.OdbcCdr
	case mainStruct.ConfSofia:
		if newConf != nil {
			c.Sofia = newConf
		}
		return c.Sofia
	case mainStruct.ConfVerto:
		if newConf != nil {
			c.Verto = newConf
		}
		return c.Verto
	case mainStruct.ConfLcr:
		if newConf != nil {
			c.Lcr = newConf
		}
		return c.Lcr
	case mainStruct.ConfShout:
		if newConf != nil {
			c.Shout = newConf
		}
		return c.Shout
	case mainStruct.ConfRedis:
		if newConf != nil {
			c.Redis = newConf
		}
		return c.Redis
	case mainStruct.ConfNibblebill:
		if newConf != nil {
			c.Nibblebill = newConf
		}
		return c.Nibblebill
	case mainStruct.ConfDb:
		if newConf != nil {
			c.Db = newConf
		}
		return c.Db
	case mainStruct.ConfMemcache:
		if newConf != nil {
			c.Memcache = newConf
		}
		return c.Memcache
	case mainStruct.ConfAvmd:
		if newConf != nil {
			c.Avmd = newConf
		}
		return c.Avmd
	case mainStruct.ConfTtsCommandline:
		if newConf != nil {
			c.TtsCommandline = newConf
		}
		return c.TtsCommandline
	case mainStruct.ConfCdrMongodb:
		if newConf != nil {
			c.CdrMongodb = newConf
		}
		return c.CdrMongodb
	case mainStruct.ConfHttpCache:
		if newConf != nil {
			c.HttpCache = newConf
		}
		return c.HttpCache
	case mainStruct.ConfOpus:
		if newConf != nil {
			c.Opus = newConf
		}
		return c.Opus
	case mainStruct.ConfPython:
		if newConf != nil {
			c.Python = newConf
		}
		return c.Python
	case mainStruct.ConfAlsa:
		if newConf != nil {
			c.Alsa = newConf
		}
		return c.Alsa
	case mainStruct.ConfAmr:
		if newConf != nil {
			c.Amr = newConf
		}
		return c.Amr
	case mainStruct.ConfAmrwb:
		if newConf != nil {
			c.Amrwb = newConf
		}
		return c.Amrwb
	case mainStruct.ConfCepstral:
		if newConf != nil {
			c.Cepstral = newConf
		}
		return c.Cepstral
	case mainStruct.ConfCidlookup:
		if newConf != nil {
			c.Cidlookup = newConf
		}
		return c.Cidlookup
	case mainStruct.ConfCurl:
		if newConf != nil {
			c.Curl = newConf
		}
		return c.Curl
	case mainStruct.ConfDialplanDirectory:
		if newConf != nil {
			c.DialplanDirectory = newConf
		}
		return c.DialplanDirectory
	case mainStruct.ConfEasyroute:
		if newConf != nil {
			c.Easyroute = newConf
		}
		return c.Easyroute
	case mainStruct.ConfErlangEvent:
		if newConf != nil {
			c.ErlangEvent = newConf
		}
		return c.ErlangEvent
	case mainStruct.ConfEventMulticast:
		if newConf != nil {
			c.EventMulticast = newConf
		}
		return c.EventMulticast
	case mainStruct.ConfFax:
		if newConf != nil {
			c.Fax = newConf
		}
		return c.Fax
	case mainStruct.ConfLua:
		if newConf != nil {
			c.Lua = newConf
		}
		return c.Lua
	case mainStruct.ConfMongo:
		if newConf != nil {
			c.Mongo = newConf
		}
		return c.Mongo
	case mainStruct.ConfMsrp:
		if newConf != nil {
			c.Msrp = newConf
		}
		return c.Msrp
	case mainStruct.ConfOreka:
		if newConf != nil {
			c.Oreka = newConf
		}
		return c.Oreka
	case mainStruct.ConfPerl:
		if newConf != nil {
			c.Perl = newConf
		}
		return c.Perl
	case mainStruct.ConfPocketsphinx:
		if newConf != nil {
			c.Pocketsphinx = newConf
		}
		return c.Pocketsphinx
	case mainStruct.ConfSangomaCodec:
		if newConf != nil {
			c.SangomaCodec = newConf
		}
		return c.SangomaCodec
	case mainStruct.ConfSndfile:
		if newConf != nil {
			c.Sndfile = newConf
		}
		return c.Sndfile
	case mainStruct.ConfXmlCdr:
		if newConf != nil {
			c.XmlCdr = newConf
		}
		return c.XmlCdr
	case mainStruct.ConfXmlRpc:
		if newConf != nil {
			c.XmlRpc = newConf
		}
		return c.XmlRpc
	case mainStruct.ConfZeroconf:
		if newConf != nil {
			c.Zeroconf = newConf
		}
		return c.Zeroconf
	case mainStruct.ConfDirectory:
		if newConf != nil {
			c.Directory = newConf
		}
		return c.Directory
	case mainStruct.ConfFifo:
		if newConf != nil {
			c.Fifo = newConf
		}
		return c.Fifo
	case mainStruct.ConfOpal:
		if newConf != nil {
			c.Opal = newConf
		}
		return c.Opal
	case mainStruct.ConfOsp:
		if newConf != nil {
			c.Osp = newConf
		}
		return c.Osp
	case mainStruct.ConfUnicall:
		if newConf != nil {
			c.Unicall = newConf
		}
		return c.Unicall
	case mainStruct.ConfConference:
		if newConf != nil {
			c.Conference = newConf
		}
		return c.Conference
	case mainStruct.ConfConferenceLayouts:
		if newConf != nil {
			c.ConferenceLayouts = newConf
		}
		return c.ConferenceLayouts
	case mainStruct.ConfPostLoadModules:
		if newConf != nil {
			c.PostLoadModules = newConf
		}
		return c.PostLoadModules
	case mainStruct.ConfVoicemail:
		if newConf != nil {
			c.Voicemail = newConf
		}
		return c.Voicemail
	case mainStruct.ConfDistributor:
		if newConf != nil {
			c.Distributor = newConf
		}
		return c.Distributor
	}

	return nil
}

func GetConfigs(s []interface{}, cached Configurations) *Configurations {
	conf := &Configurations{}
	for _, c := range s {
		config, ok := c.(ConfigurationsList)
		if !ok {
			continue
		}
		subCached := cached.GetConfiguration(config.Name)
		if subCached != nil {
			config.Loaded = subCached.Loaded
		}
		conf.FillConfigurations(&config)
	}
	return conf
}

type Configurations struct {
	PostSwitch        *ConfigurationsList `xml:"" json:"post_load_switch"`
	Acl               *ConfigurationsList `xml:"" json:"acl"`
	Callcenter        *ConfigurationsList `xml:"" json:"callcenter"`
	CdrPgCsv          *ConfigurationsList `xml:"" json:"cdr_pg_csv"`
	OdbcCdr           *ConfigurationsList `xml:"" json:"odbc_cdr"`
	Lcr               *ConfigurationsList `xml:"" json:"lcr"`
	Sofia             *ConfigurationsList `xml:"" json:"sofia"`
	Verto             *ConfigurationsList `xml:"" json:"verto"`
	Shout             *ConfigurationsList `xml:"" json:"shout"`
	Redis             *ConfigurationsList `xml:"" json:"redis"`
	Nibblebill        *ConfigurationsList `xml:"" json:"nibblebill"`
	Db                *ConfigurationsList `xml:"" json:"db"`
	Distributor       *ConfigurationsList `xml:"" json:"distributor"`
	Memcache          *ConfigurationsList `xml:"" json:"memcache"`
	Avmd              *ConfigurationsList `xml:"" json:"avmd"`
	TtsCommandline    *ConfigurationsList `xml:"" json:"tts_commandline"`
	CdrMongodb        *ConfigurationsList `xml:"" json:"cdr_mongodb"`
	HttpCache         *ConfigurationsList `xml:"" json:"http_cache"`
	Opus              *ConfigurationsList `xml:"" json:"opus"`
	Python            *ConfigurationsList `xml:"" json:"python"`
	Alsa              *ConfigurationsList `xml:"" json:"alsa"`
	Amr               *ConfigurationsList `xml:"" json:"amr"`
	Amrwb             *ConfigurationsList `xml:"" json:"amrwb"`
	Cepstral          *ConfigurationsList `xml:"" json:"cepstral"`
	Cidlookup         *ConfigurationsList `xml:"" json:"cidlookup"`
	Curl              *ConfigurationsList `xml:"" json:"curl"`
	DialplanDirectory *ConfigurationsList `xml:"" json:"dialplan_directory"`
	Easyroute         *ConfigurationsList `xml:"" json:"easyroute"`
	ErlangEvent       *ConfigurationsList `xml:"" json:"erlang_event"`
	EventMulticast    *ConfigurationsList `xml:"" json:"event_multicast"`
	Fax               *ConfigurationsList `xml:"" json:"fax"`
	Lua               *ConfigurationsList `xml:"" json:"lua"`
	Mongo             *ConfigurationsList `xml:"" json:"mongo"`
	Msrp              *ConfigurationsList `xml:"" json:"msrp"`
	Oreka             *ConfigurationsList `xml:"" json:"oreka"`
	Perl              *ConfigurationsList `xml:"" json:"perl"`
	Pocketsphinx      *ConfigurationsList `xml:"" json:"pocketsphinx"`
	SangomaCodec      *ConfigurationsList `xml:"" json:"sangoma_codec"`
	Sndfile           *ConfigurationsList `xml:"" json:"sndfile"`
	XmlCdr            *ConfigurationsList `xml:"" json:"xml_cdr"`
	XmlRpc            *ConfigurationsList `xml:"" json:"xml_rpc"`
	Zeroconf          *ConfigurationsList `xml:"" json:"zeroconf"`
	Directory         *ConfigurationsList `xml:"" json:"directory"`
	Fifo              *ConfigurationsList `xml:"" json:"fifo"`
	Opal              *ConfigurationsList `xml:"" json:"opal"`
	Osp               *ConfigurationsList `xml:"" json:"osp"`
	Unicall           *ConfigurationsList `xml:"" json:"unicall"`
	Conference        *ConfigurationsList `xml:"" json:"conference"`
	ConferenceLayouts *ConfigurationsList `xml:"" json:"conference_layouts"`
	PostLoadModules   *ConfigurationsList `xml:"" json:"post_load_modules"`
	Voicemail         *ConfigurationsList `xml:"" json:"voicemail"`
}
