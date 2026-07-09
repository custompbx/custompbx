package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventShoutGet                     = "GetShout"
	eventShoutParamUpdate             = "UpdateShoutParameter"
	eventShoutParamSwitch             = "SwitchShoutParameter"
	eventShoutParamAdd                = "AddShoutParameter"
	eventShoutParamDelete             = "DelShoutParameter"
	eventRedisGet                     = "GetRedis"
	eventRedisParamUpdate             = "UpdateRedisParameter"
	eventRedisParamSwitch             = "SwitchRedisParameter"
	eventRedisParamAdd                = "AddRedisParameter"
	eventRedisParamDelete             = "DelRedisParameter"
	eventNibblebillGet                = "GetNibblebill"
	eventNibblebillParamUpdate        = "UpdateNibblebillParameter"
	eventNibblebillParamSwitch        = "SwitchNibblebillParameter"
	eventNibblebillParamAdd           = "AddNibblebillParameter"
	eventNibblebillParamDelete        = "DelNibblebillParameter"
	eventDBGet                        = "GetDb"
	eventDBParamUpdate                = "UpdateDbParameter"
	eventDBParamSwitch                = "SwitchDbParameter"
	eventDBParamAdd                   = "AddDbParameter"
	eventDBParamDelete                = "DelDbParameter"
	eventMemcacheGet                  = "GetMemcache"
	eventMemcacheParamUpdate          = "UpdateMemcacheParameter"
	eventMemcacheParamSwitch          = "SwitchMemcacheParameter"
	eventMemcacheParamAdd             = "AddMemcacheParameter"
	eventMemcacheParamDelete          = "DelMemcacheParameter"
	eventAvmdGet                      = "GetAvmd"
	eventAvmdParamUpdate              = "UpdateAvmdParameter"
	eventAvmdParamSwitch              = "SwitchAvmdParameter"
	eventAvmdParamAdd                 = "AddAvmdParameter"
	eventAvmdParamDelete              = "DelAvmdParameter"
	eventTtsCommandlineGet            = "GetTtsCommandline"
	eventTtsCommandlineParamUpdate    = "UpdateTtsCommandlineParameter"
	eventTtsCommandlineParamSwitch    = "SwitchTtsCommandlineParameter"
	eventTtsCommandlineParamAdd       = "AddTtsCommandlineParameter"
	eventTtsCommandlineParamDelete    = "DelTtsCommandlineParameter"
	eventCdrMongodbGet                = "GetCdrMongodb"
	eventCdrMongodbParamUpdate        = "UpdateCdrMongodbParameter"
	eventCdrMongodbParamSwitch        = "SwitchCdrMongodbParameter"
	eventCdrMongodbParamAdd           = "AddCdrMongodbParameter"
	eventCdrMongodbParamDelete        = "DelCdrMongodbParameter"
	eventOpusGet                      = "GetOpus"
	eventOpusParamUpdate              = "UpdateOpusParameter"
	eventOpusParamSwitch              = "SwitchOpusParameter"
	eventOpusParamAdd                 = "AddOpusParameter"
	eventOpusParamDelete              = "DelOpusParameter"
	eventPythonGet                    = "GetPython"
	eventPythonParamUpdate            = "UpdatePythonParameter"
	eventPythonParamSwitch            = "SwitchPythonParameter"
	eventPythonParamAdd               = "AddPythonParameter"
	eventPythonParamDelete            = "DelPythonParameter"
	eventAlsaGet                      = "GetAlsa"
	eventAlsaParamUpdate              = "UpdateAlsaParameter"
	eventAlsaParamSwitch              = "SwitchAlsaParameter"
	eventAlsaParamAdd                 = "AddAlsaParameter"
	eventAlsaParamDelete              = "DelAlsaParameter"
	eventAmrGet                       = "GetAmr"
	eventAmrParamUpdate               = "UpdateAmrParameter"
	eventAmrParamSwitch               = "SwitchAmrParameter"
	eventAmrParamAdd                  = "AddAmrParameter"
	eventAmrParamDelete               = "DelAmrParameter"
	eventAmrwbGet                     = "GetAmrwb"
	eventAmrwbParamUpdate             = "UpdateAmrwbParameter"
	eventAmrwbParamSwitch             = "SwitchAmrwbParameter"
	eventAmrwbParamAdd                = "AddAmrwbParameter"
	eventAmrwbParamDelete             = "DelAmrwbParameter"
	eventCepstralGet                  = "GetCepstral"
	eventCepstralParamUpdate          = "UpdateCepstralParameter"
	eventCepstralParamSwitch          = "SwitchCepstralParameter"
	eventCepstralParamAdd             = "AddCepstralParameter"
	eventCepstralParamDelete          = "DelCepstralParameter"
	eventCidlookupGet                 = "GetCidlookup"
	eventCidlookupParamUpdate         = "UpdateCidlookupParameter"
	eventCidlookupParamSwitch         = "SwitchCidlookupParameter"
	eventCidlookupParamAdd            = "AddCidlookupParameter"
	eventCidlookupParamDelete         = "DelCidlookupParameter"
	eventCurlGet                      = "GetCurl"
	eventCurlParamUpdate              = "UpdateCurlParameter"
	eventCurlParamSwitch              = "SwitchCurlParameter"
	eventCurlParamAdd                 = "AddCurlParameter"
	eventCurlParamDelete              = "DelCurlParameter"
	eventDialplanDirectoryGet         = "GetDialplanDirectory"
	eventDialplanDirectoryParamUpdate = "UpdateDialplanDirectoryParameter"
	eventDialplanDirectoryParamSwitch = "SwitchDialplanDirectoryParameter"
	eventDialplanDirectoryParamAdd    = "AddDialplanDirectoryParameter"
	eventDialplanDirectoryParamDelete = "DelDialplanDirectoryParameter"
	eventEasyrouteGet                 = "GetEasyroute"
	eventEasyrouteParamUpdate         = "UpdateEasyrouteParameter"
	eventEasyrouteParamSwitch         = "SwitchEasyrouteParameter"
	eventEasyrouteParamAdd            = "AddEasyrouteParameter"
	eventEasyrouteParamDelete         = "DelEasyrouteParameter"
	eventErlangEventGet               = "GetErlangEvent"
	eventErlangEventParamUpdate       = "UpdateErlangEventParameter"
	eventErlangEventParamSwitch       = "SwitchErlangEventParameter"
	eventErlangEventParamAdd          = "AddErlangEventParameter"
	eventErlangEventParamDelete       = "DelErlangEventParameter"
	eventEventMulticastGet            = "GetEventMulticast"
	eventEventMulticastParamUpdate    = "UpdateEventMulticastParameter"
	eventEventMulticastParamSwitch    = "SwitchEventMulticastParameter"
	eventEventMulticastParamAdd       = "AddEventMulticastParameter"
	eventEventMulticastParamDelete    = "DelEventMulticastParameter"
	eventFaxGet                       = "GetFax"
	eventFaxParamUpdate               = "UpdateFaxParameter"
	eventFaxParamSwitch               = "SwitchFaxParameter"
	eventFaxParamAdd                  = "AddFaxParameter"
	eventFaxParamDelete               = "DelFaxParameter"
	eventLuaGet                       = "GetLua"
	eventLuaParamUpdate               = "UpdateLuaParameter"
	eventLuaParamSwitch               = "SwitchLuaParameter"
	eventLuaParamAdd                  = "AddLuaParameter"
	eventLuaParamDelete               = "DelLuaParameter"
	eventMongoGet                     = "GetMongo"
	eventMongoParamUpdate             = "UpdateMongoParameter"
	eventMongoParamSwitch             = "SwitchMongoParameter"
	eventMongoParamAdd                = "AddMongoParameter"
	eventMongoParamDelete             = "DelMongoParameter"
	eventMsrpGet                      = "GetMsrp"
	eventMsrpParamUpdate              = "UpdateMsrpParameter"
	eventMsrpParamSwitch              = "SwitchMsrpParameter"
	eventMsrpParamAdd                 = "AddMsrpParameter"
	eventMsrpParamDelete              = "DelMsrpParameter"
	eventOrekaGet                     = "GetOreka"
	eventOrekaParamUpdate             = "UpdateOrekaParameter"
	eventOrekaParamSwitch             = "SwitchOrekaParameter"
	eventOrekaParamAdd                = "AddOrekaParameter"
	eventOrekaParamDelete             = "DelOrekaParameter"
	eventPerlGet                      = "GetPerl"
	eventPerlParamUpdate              = "UpdatePerlParameter"
	eventPerlParamSwitch              = "SwitchPerlParameter"
	eventPerlParamAdd                 = "AddPerlParameter"
	eventPerlParamDelete              = "DelPerlParameter"
	eventPocketsphinxGet              = "GetPocketsphinx"
	eventPocketsphinxParamUpdate      = "UpdatePocketsphinxParameter"
	eventPocketsphinxParamSwitch      = "SwitchPocketsphinxParameter"
	eventPocketsphinxParamAdd         = "AddPocketsphinxParameter"
	eventPocketsphinxParamDelete      = "DelPocketsphinxParameter"
	eventSangomaCodecGet              = "GetSangomaCodec"
	eventSangomaCodecParamUpdate      = "UpdateSangomaCodecParameter"
	eventSangomaCodecParamSwitch      = "SwitchSangomaCodecParameter"
	eventSangomaCodecParamAdd         = "AddSangomaCodecParameter"
	eventSangomaCodecParamDelete      = "DelSangomaCodecParameter"
	eventSndfileGet                   = "GetSndfile"
	eventSndfileParamUpdate           = "UpdateSndfileParameter"
	eventSndfileParamSwitch           = "SwitchSndfileParameter"
	eventSndfileParamAdd              = "AddSndfileParameter"
	eventSndfileParamDelete           = "DelSndfileParameter"
	eventXmlCdrGet                    = "GetXmlCdr"
	eventXmlCdrParamUpdate            = "UpdateXmlCdrParameter"
	eventXmlCdrParamSwitch            = "SwitchXmlCdrParameter"
	eventXmlCdrParamAdd               = "AddXmlCdrParameter"
	eventXmlCdrParamDelete            = "DelXmlCdrParameter"
	eventXmlRpcGet                    = "GetXmlRpc"
	eventXmlRpcParamUpdate            = "UpdateXmlRpcParameter"
	eventXmlRpcParamSwitch            = "SwitchXmlRpcParameter"
	eventXmlRpcParamAdd               = "AddXmlRpcParameter"
	eventXmlRpcParamDelete            = "DelXmlRpcParameter"
	eventZeroconfGet                  = "GetZeroconf"
	eventZeroconfParamUpdate          = "UpdateZeroconfParameter"
	eventZeroconfParamSwitch          = "SwitchZeroconfParameter"
	eventZeroconfParamAdd             = "AddZeroconfParameter"
	eventZeroconfParamDelete          = "DelZeroconfParameter"
)

func registerCoreSimpleModuleSettingEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventShoutGet, Update: eventShoutParamUpdate, Switch: eventShoutParamSwitch, Add: eventShoutParamAdd, Delete: eventShoutParamDelete},
		&altStruct.ConfigShoutSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigShoutSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigShoutSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigShoutSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigShoutSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigShoutSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventRedisGet, Update: eventRedisParamUpdate, Switch: eventRedisParamSwitch, Add: eventRedisParamAdd, Delete: eventRedisParamDelete},
		&altStruct.ConfigRedisSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigRedisSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigRedisSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigRedisSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigRedisSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigRedisSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventNibblebillGet, Update: eventNibblebillParamUpdate, Switch: eventNibblebillParamSwitch, Add: eventNibblebillParamAdd, Delete: eventNibblebillParamDelete},
		&altStruct.ConfigNibblebillSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigNibblebillSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigNibblebillSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigNibblebillSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigNibblebillSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigNibblebillSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventDBGet, Update: eventDBParamUpdate, Switch: eventDBParamSwitch, Add: eventDBParamAdd, Delete: eventDBParamDelete},
		&altStruct.ConfigDbSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigDbSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigDbSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigDbSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigDbSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigDbSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventMemcacheGet, Update: eventMemcacheParamUpdate, Switch: eventMemcacheParamSwitch, Add: eventMemcacheParamAdd, Delete: eventMemcacheParamDelete},
		&altStruct.ConfigMemcacheSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigMemcacheSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigMemcacheSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigMemcacheSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigMemcacheSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigMemcacheSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventAvmdGet, Update: eventAvmdParamUpdate, Switch: eventAvmdParamSwitch, Add: eventAvmdParamAdd, Delete: eventAvmdParamDelete},
		&altStruct.ConfigAvmdSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigAvmdSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigAvmdSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigAvmdSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigAvmdSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigAvmdSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventTtsCommandlineGet, Update: eventTtsCommandlineParamUpdate, Switch: eventTtsCommandlineParamSwitch, Add: eventTtsCommandlineParamAdd, Delete: eventTtsCommandlineParamDelete},
		&altStruct.ConfigTtsCommandlineSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigTtsCommandlineSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigTtsCommandlineSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigTtsCommandlineSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigTtsCommandlineSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigTtsCommandlineSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfig(r, overrides, simpleParamConfigEvents{Get: eventCdrMongodbGet, Update: eventCdrMongodbParamUpdate, Switch: eventCdrMongodbParamSwitch, Add: eventCdrMongodbParamAdd, Delete: eventCdrMongodbParamDelete},
		&altStruct.ConfigCdrMongodbSetting{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigCdrMongodbSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigCdrMongodbSetting{})}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigCdrMongodbSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigCdrMongodbSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
		},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigCdrMongodbSetting{Id: data.Param.Id}
		},
	)
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventOpusGet, Update: eventOpusParamUpdate, Switch: eventOpusParamSwitch, Add: eventOpusParamAdd, Delete: eventOpusParamDelete}, &altStruct.ConfigOpusSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventPythonGet, Update: eventPythonParamUpdate, Switch: eventPythonParamSwitch, Add: eventPythonParamAdd, Delete: eventPythonParamDelete}, &altStruct.ConfigPythonSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventAlsaGet, Update: eventAlsaParamUpdate, Switch: eventAlsaParamSwitch, Add: eventAlsaParamAdd, Delete: eventAlsaParamDelete}, &altStruct.ConfigAlsaSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventAmrGet, Update: eventAmrParamUpdate, Switch: eventAmrParamSwitch, Add: eventAmrParamAdd, Delete: eventAmrParamDelete}, &altStruct.ConfigAmrSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventAmrwbGet, Update: eventAmrwbParamUpdate, Switch: eventAmrwbParamSwitch, Add: eventAmrwbParamAdd, Delete: eventAmrwbParamDelete}, &altStruct.ConfigAmrwbSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventCepstralGet, Update: eventCepstralParamUpdate, Switch: eventCepstralParamSwitch, Add: eventCepstralParamAdd, Delete: eventCepstralParamDelete}, &altStruct.ConfigCespalSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventCidlookupGet, Update: eventCidlookupParamUpdate, Switch: eventCidlookupParamSwitch, Add: eventCidlookupParamAdd, Delete: eventCidlookupParamDelete}, &altStruct.ConfigCidlookupSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventCurlGet, Update: eventCurlParamUpdate, Switch: eventCurlParamSwitch, Add: eventCurlParamAdd, Delete: eventCurlParamDelete}, &altStruct.ConfigCurlSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventDialplanDirectoryGet, Update: eventDialplanDirectoryParamUpdate, Switch: eventDialplanDirectoryParamSwitch, Add: eventDialplanDirectoryParamAdd, Delete: eventDialplanDirectoryParamDelete}, &altStruct.ConfigDialplanDirectorySetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventEasyrouteGet, Update: eventEasyrouteParamUpdate, Switch: eventEasyrouteParamSwitch, Add: eventEasyrouteParamAdd, Delete: eventEasyrouteParamDelete}, &altStruct.ConfigEasyrouteSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventErlangEventGet, Update: eventErlangEventParamUpdate, Switch: eventErlangEventParamSwitch, Add: eventErlangEventParamAdd, Delete: eventErlangEventParamDelete}, &altStruct.ConfigErlangEventSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventEventMulticastGet, Update: eventEventMulticastParamUpdate, Switch: eventEventMulticastParamSwitch, Add: eventEventMulticastParamAdd, Delete: eventEventMulticastParamDelete}, &altStruct.ConfigEventMulticastSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventFaxGet, Update: eventFaxParamUpdate, Switch: eventFaxParamSwitch, Add: eventFaxParamAdd, Delete: eventFaxParamDelete}, &altStruct.ConfigFaxSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventLuaGet, Update: eventLuaParamUpdate, Switch: eventLuaParamSwitch, Add: eventLuaParamAdd, Delete: eventLuaParamDelete}, &altStruct.ConfigLuaSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventMongoGet, Update: eventMongoParamUpdate, Switch: eventMongoParamSwitch, Add: eventMongoParamAdd, Delete: eventMongoParamDelete}, &altStruct.ConfigMongoSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventMsrpGet, Update: eventMsrpParamUpdate, Switch: eventMsrpParamSwitch, Add: eventMsrpParamAdd, Delete: eventMsrpParamDelete}, &altStruct.ConfigMsrpSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventOrekaGet, Update: eventOrekaParamUpdate, Switch: eventOrekaParamSwitch, Add: eventOrekaParamAdd, Delete: eventOrekaParamDelete}, &altStruct.ConfigOrekaSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventPerlGet, Update: eventPerlParamUpdate, Switch: eventPerlParamSwitch, Add: eventPerlParamAdd, Delete: eventPerlParamDelete}, &altStruct.ConfigPerlSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventPocketsphinxGet, Update: eventPocketsphinxParamUpdate, Switch: eventPocketsphinxParamSwitch, Add: eventPocketsphinxParamAdd, Delete: eventPocketsphinxParamDelete}, &altStruct.ConfigPocketsphinxSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventSangomaCodecGet, Update: eventSangomaCodecParamUpdate, Switch: eventSangomaCodecParamSwitch, Add: eventSangomaCodecParamAdd, Delete: eventSangomaCodecParamDelete}, &altStruct.ConfigSangomaCodecSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventSndfileGet, Update: eventSndfileParamUpdate, Switch: eventSndfileParamSwitch, Add: eventSndfileParamAdd, Delete: eventSndfileParamDelete}, &altStruct.ConfigSndfileSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventXmlCdrGet, Update: eventXmlCdrParamUpdate, Switch: eventXmlCdrParamSwitch, Add: eventXmlCdrParamAdd, Delete: eventXmlCdrParamDelete}, &altStruct.ConfigXmlCdrSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventXmlRpcGet, Update: eventXmlRpcParamUpdate, Switch: eventXmlRpcParamSwitch, Add: eventXmlRpcParamAdd, Delete: eventXmlRpcParamDelete}, &altStruct.ConfigXmlRpcSetting{})
	registerSimpleParamConfigForSample(r, overrides, simpleParamConfigEvents{Get: eventZeroconfGet, Update: eventZeroconfParamUpdate, Switch: eventZeroconfParamSwitch, Add: eventZeroconfParamAdd, Delete: eventZeroconfParamDelete}, &altStruct.ConfigZeroconfSetting{})
}

func simpleModuleSettingEvents() []string {
	return []string{
		eventShoutGet, eventShoutParamUpdate, eventShoutParamSwitch, eventShoutParamAdd, eventShoutParamDelete,
		eventRedisGet, eventRedisParamUpdate, eventRedisParamSwitch, eventRedisParamAdd, eventRedisParamDelete,
		eventNibblebillGet, eventNibblebillParamUpdate, eventNibblebillParamSwitch, eventNibblebillParamAdd, eventNibblebillParamDelete,
		eventDBGet, eventDBParamUpdate, eventDBParamSwitch, eventDBParamAdd, eventDBParamDelete,
		eventMemcacheGet, eventMemcacheParamUpdate, eventMemcacheParamSwitch, eventMemcacheParamAdd, eventMemcacheParamDelete,
		eventAvmdGet, eventAvmdParamUpdate, eventAvmdParamSwitch, eventAvmdParamAdd, eventAvmdParamDelete,
		eventTtsCommandlineGet, eventTtsCommandlineParamUpdate, eventTtsCommandlineParamSwitch, eventTtsCommandlineParamAdd, eventTtsCommandlineParamDelete,
		eventCdrMongodbGet, eventCdrMongodbParamUpdate, eventCdrMongodbParamSwitch, eventCdrMongodbParamAdd, eventCdrMongodbParamDelete,
		eventOpusGet, eventOpusParamUpdate, eventOpusParamSwitch, eventOpusParamAdd, eventOpusParamDelete,
		eventPythonGet, eventPythonParamUpdate, eventPythonParamSwitch, eventPythonParamAdd, eventPythonParamDelete,
		eventAlsaGet, eventAlsaParamUpdate, eventAlsaParamSwitch, eventAlsaParamAdd, eventAlsaParamDelete,
		eventAmrGet, eventAmrParamUpdate, eventAmrParamSwitch, eventAmrParamAdd, eventAmrParamDelete,
		eventAmrwbGet, eventAmrwbParamUpdate, eventAmrwbParamSwitch, eventAmrwbParamAdd, eventAmrwbParamDelete,
		eventCepstralGet, eventCepstralParamUpdate, eventCepstralParamSwitch, eventCepstralParamAdd, eventCepstralParamDelete,
		eventCidlookupGet, eventCidlookupParamUpdate, eventCidlookupParamSwitch, eventCidlookupParamAdd, eventCidlookupParamDelete,
		eventCurlGet, eventCurlParamUpdate, eventCurlParamSwitch, eventCurlParamAdd, eventCurlParamDelete,
		eventDialplanDirectoryGet, eventDialplanDirectoryParamUpdate, eventDialplanDirectoryParamSwitch, eventDialplanDirectoryParamAdd, eventDialplanDirectoryParamDelete,
		eventEasyrouteGet, eventEasyrouteParamUpdate, eventEasyrouteParamSwitch, eventEasyrouteParamAdd, eventEasyrouteParamDelete,
		eventErlangEventGet, eventErlangEventParamUpdate, eventErlangEventParamSwitch, eventErlangEventParamAdd, eventErlangEventParamDelete,
		eventEventMulticastGet, eventEventMulticastParamUpdate, eventEventMulticastParamSwitch, eventEventMulticastParamAdd, eventEventMulticastParamDelete,
		eventFaxGet, eventFaxParamUpdate, eventFaxParamSwitch, eventFaxParamAdd, eventFaxParamDelete,
		eventLuaGet, eventLuaParamUpdate, eventLuaParamSwitch, eventLuaParamAdd, eventLuaParamDelete,
		eventMongoGet, eventMongoParamUpdate, eventMongoParamSwitch, eventMongoParamAdd, eventMongoParamDelete,
		eventMsrpGet, eventMsrpParamUpdate, eventMsrpParamSwitch, eventMsrpParamAdd, eventMsrpParamDelete,
		eventOrekaGet, eventOrekaParamUpdate, eventOrekaParamSwitch, eventOrekaParamAdd, eventOrekaParamDelete,
		eventPerlGet, eventPerlParamUpdate, eventPerlParamSwitch, eventPerlParamAdd, eventPerlParamDelete,
		eventPocketsphinxGet, eventPocketsphinxParamUpdate, eventPocketsphinxParamSwitch, eventPocketsphinxParamAdd, eventPocketsphinxParamDelete,
		eventSangomaCodecGet, eventSangomaCodecParamUpdate, eventSangomaCodecParamSwitch, eventSangomaCodecParamAdd, eventSangomaCodecParamDelete,
		eventSndfileGet, eventSndfileParamUpdate, eventSndfileParamSwitch, eventSndfileParamAdd, eventSndfileParamDelete,
		eventXmlCdrGet, eventXmlCdrParamUpdate, eventXmlCdrParamSwitch, eventXmlCdrParamAdd, eventXmlCdrParamDelete,
		eventXmlRpcGet, eventXmlRpcParamUpdate, eventXmlRpcParamSwitch, eventXmlRpcParamAdd, eventXmlRpcParamDelete,
		eventZeroconfGet, eventZeroconfParamUpdate, eventZeroconfParamSwitch, eventZeroconfParamAdd, eventZeroconfParamDelete,
	}
}
