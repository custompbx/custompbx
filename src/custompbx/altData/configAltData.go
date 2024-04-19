package altData

import (
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"encoding/xml"
	"errors"
)

func GetConfigSection(confName string, post map[string]string) *mainStruct.Configuration {
	switch confName {
	case mainStruct.ConfPostLoadSwitch:
		return XMLPostSwitch(confName)
	case mainStruct.ConfAcl:
		return XMLAcl(confName)
	case mainStruct.ConfCallcenter:
		return XMLCallcenter(confName, post["CC-Queue"])
	case mainStruct.ConfCdrPgCsv:
		return XMLCdrPgCsv(confName)
	case mainStruct.ConfOdbcCdr:
		return XMLOdbcCdr(confName)
	case mainStruct.ConfSofia:
		return XMLSofia(confName, post["profile"])
	case mainStruct.ConfVerto:
		return XMLVerto(confName)
	case mainStruct.ConfLcr:
		return XMLLcr(confName)
	case mainStruct.ConfShout:
		return XMLShout(confName)
	case mainStruct.ConfRedis:
		return XMLRedis(confName)
	case mainStruct.ConfNibblebill:
		return XMLNibblebill(confName)
	case mainStruct.ConfDb:
		return XMLDb(confName)
	case mainStruct.ConfMemcache:
		return XMLMemcache(confName)
	case mainStruct.ConfAvmd:
		return XMLAvmd(confName)
	case mainStruct.ConfTtsCommandline:
		return XMLTtsCommandline(confName)
	case mainStruct.ConfCdrMongodb:
		return XMLCdrMongodb(confName)
	case mainStruct.ConfHttpCache:
		return XMLHttpCache(confName)
	case mainStruct.ConfOpus:
		return XMLOpus(confName)
	case mainStruct.ConfPython:
		return XMLPython(confName)
	case mainStruct.ConfAlsa:
		return XMLAlsa(confName)
	case mainStruct.ConfAmr:
		return XMLAmr(confName)
	case mainStruct.ConfAmrwb:
		return XMLAmrwb(confName)
	case mainStruct.ConfCepstral:
		return XMLCepstral(confName)
	case mainStruct.ConfCidlookup:
		return XMLCidlookup(confName)
	case mainStruct.ConfCurl:
		return XMLCurl(confName)
	case mainStruct.ConfDialplanDirectory:
		return XMLDialplanDirectory(confName)
	case mainStruct.ConfEasyroute:
		return XMLEasyroute(confName)
	case mainStruct.ConfErlangEvent:
		return XMLErlangEvent(confName)
	case mainStruct.ConfEventMulticast:
		return XMLEventMulticast(confName)
	case mainStruct.ConfFax:
		return XMLFax(confName)
	case mainStruct.ConfLua:
		return XMLLua(confName)
	case mainStruct.ConfMongo:
		return XMLMongo(confName)
	case mainStruct.ConfMsrp:
		return XMLMsrp(confName)
	case mainStruct.ConfOreka:
		return XMLOreka(confName)
	case mainStruct.ConfPerl:
		return XMLPerl(confName)
	case mainStruct.ConfPocketsphinx:
		return XMLPocketsphinx(confName)
	case mainStruct.ConfSangomaCodec:
		return XMLSangomaCodec(confName)
	case mainStruct.ConfSndfile:
		return XMLSndfile(confName)
	case mainStruct.ConfXmlCdr:
		return XMLXmlCdr(confName)
	case mainStruct.ConfXmlRpc:
		return XMLXmlRpc(confName)
	case mainStruct.ConfZeroconf:
		return XMLZeroconf(confName)
	case mainStruct.ConfDirectory:
		return XMLDirectory(confName)
	case mainStruct.ConfFifo:
		return XMLFifo(confName)
	case mainStruct.ConfOpal:
		return XMLOpal(confName)
	case mainStruct.ConfOsp:
		return XMLOsp(confName)
	case mainStruct.ConfUnicall:
		return XMLUnicall(confName)
	case mainStruct.ConfConference:
		return XMLConference(confName)
	case mainStruct.ConfConferenceLayouts:
		return XMLConferenceLayouts(confName)
	case mainStruct.ConfPostLoadModules:
		return XMLPostLoadModules(confName)
	case mainStruct.ConfVoicemail:
		return XMLVoicemail(confName)
	case mainStruct.ConfDistributor:
		return XMLDistributor(confName)
	default:
		return nil
	}
}

func GetModuleByName(name string) (*altStruct.ConfigurationsList, error) {
	switch name {
	case mainStruct.ModCdrPgCsv:
		return getCurrentModule(mainStruct.ConfCdrPgCsv, name)
	case mainStruct.ModSofiaAlias:
		return getCurrentModule(mainStruct.ConfSofia, name)
	case mainStruct.ModSofia:
		return getCurrentModule(mainStruct.ConfSofia, name)
	case mainStruct.ModAcl:
		return getCurrentModule(mainStruct.ConfAcl, name)
	case mainStruct.ModVerto:
		return getCurrentModule(mainStruct.ConfVerto, name)
	case mainStruct.ModVertoAlias:
		return getCurrentModule(mainStruct.ConfVerto, name)
	case mainStruct.ModCallcenter:
		return getCurrentModule(mainStruct.ConfCallcenter, name)
	case mainStruct.ModCallcenterAlias:
		return getCurrentModule(mainStruct.ConfCallcenter, name)
	case mainStruct.ModOdbcCdr:
		return getCurrentModule(mainStruct.ConfOdbcCdr, name)
	case mainStruct.ModLcrAlias:
		return getCurrentModule(mainStruct.ConfLcr, name)
	case mainStruct.ModLcr:
		return getCurrentModule(mainStruct.ConfLcr, name)
	case mainStruct.ModShout:
		return getCurrentModule(mainStruct.ConfShout, name)
	case mainStruct.ModRedis:
		return getCurrentModule(mainStruct.ConfRedis, name)
	case mainStruct.ModNibblebill:
		return getCurrentModule(mainStruct.ConfNibblebill, name)
	case mainStruct.ModDb:
		return getCurrentModule(mainStruct.ConfDb, name)
	case mainStruct.ModMemcache:
		return getCurrentModule(mainStruct.ConfMemcache, name)
	case mainStruct.ModAvmd:
		return getCurrentModule(mainStruct.ConfAvmd, name)
	case mainStruct.ModTtsCommandline:
		return getCurrentModule(mainStruct.ConfTtsCommandline, name)
	case mainStruct.ModCdrMongodb:
		return getCurrentModule(mainStruct.ConfCdrMongodb, name)
	case mainStruct.ModHttpCache:
		return getCurrentModule(mainStruct.ConfHttpCache, name)
	case mainStruct.ModOpus:
		return getCurrentModule(mainStruct.ConfOpus, name)
	case mainStruct.ModPython:
		return getCurrentModule(mainStruct.ConfPython, name)
	case mainStruct.ModAlsa:
		return getCurrentModule(mainStruct.ConfAlsa, name)
	case mainStruct.ModAmr:
		return getCurrentModule(mainStruct.ConfAmr, name)
	case mainStruct.ModAmrwb:
		return getCurrentModule(mainStruct.ConfAmrwb, name)
	case mainStruct.ModCepstral:
		return getCurrentModule(mainStruct.ConfCepstral, name)
	case mainStruct.ModCidlookup:
		return getCurrentModule(mainStruct.ConfCidlookup, name)
	case mainStruct.ModCurl:
		return getCurrentModule(mainStruct.ConfCurl, name)
	case mainStruct.ModDialplanDirectory:
		return getCurrentModule(mainStruct.ConfDialplanDirectory, name)
	case mainStruct.ModEasyroute:
		return getCurrentModule(mainStruct.ConfEasyroute, name)
	case mainStruct.ModErlangEvent:
		return getCurrentModule(mainStruct.ConfErlangEvent, name)
	case mainStruct.ModEventMulticast:
		return getCurrentModule(mainStruct.ConfEventMulticast, name)
	case mainStruct.ModFax:
		return getCurrentModule(mainStruct.ConfFax, name)
	case mainStruct.ModLua:
		return getCurrentModule(mainStruct.ConfLua, name)
	case mainStruct.ModMongo:
		return getCurrentModule(mainStruct.ConfMongo, name)
	case mainStruct.ModMsrp:
		return getCurrentModule(mainStruct.ConfMsrp, name)
	case mainStruct.ModOreka:
		return getCurrentModule(mainStruct.ConfOreka, name)
	case mainStruct.ModPerl:
		return getCurrentModule(mainStruct.ConfPerl, name)
	case mainStruct.ModPocketsphinx:
		return getCurrentModule(mainStruct.ConfPocketsphinx, name)
	case mainStruct.ModSangomaCodec:
		return getCurrentModule(mainStruct.ConfSangomaCodec, name)
	case mainStruct.ModSndfile:
		return getCurrentModule(mainStruct.ConfSndfile, name)
	case mainStruct.ModXmlCdr:
		return getCurrentModule(mainStruct.ConfXmlCdr, name)
	case mainStruct.ModXmlRpc:
		return getCurrentModule(mainStruct.ConfXmlRpc, name)
	case mainStruct.ModZeroconf:
		return getCurrentModule(mainStruct.ConfZeroconf, name)
	case mainStruct.ModDistributor:
		return getCurrentModule(mainStruct.ConfDistributor, name)
	case mainStruct.ModOpal:
		return getCurrentModule(mainStruct.ConfOpal, name)
	case mainStruct.ModUnicall:
		return getCurrentModule(mainStruct.ConfUnicall, name)
	case mainStruct.ModDirectory:
		return getCurrentModule(mainStruct.ConfDirectory, name)
	case mainStruct.ModFifo:
		return getCurrentModule(mainStruct.ConfFifo, name)
	case mainStruct.ModOsp:
		return getCurrentModule(mainStruct.ConfOsp, name)
	case mainStruct.ModConference:
		return getCurrentModule(mainStruct.ConfConference, name)
	case mainStruct.ModConferenceLayouts:
		return getCurrentModule(mainStruct.ConfConferenceLayouts, name)
	case mainStruct.ModPostLoadModules:
		return getCurrentModule(mainStruct.ConfPostLoadModules, name)
	case mainStruct.ModPostLoadSwitch:
		return getCurrentModule(mainStruct.ConfPostLoadSwitch, name)
	case mainStruct.ModVoicemail:
		return getCurrentModule(mainStruct.ConfVoicemail, name)
	default:
		return nil, errors.New("no config")
	}
}

/*
acl

alsa
amr
amrwb
avmd
callcenter
cdr_mongodb
cdr_pg_csv
cepstral
cidlookup
conference
curl
db
dialplan_directory
directory
distributor
easyroute
erlang_event
event_multicast
fax
fifo
http_cache
lcr
lua
memcache
mongo
msrp
nibblebill
odbc_cdr
opal
opus
oreka
osp
perl
pocketsphinx
post-switch
post_load_modules
python
redis
sangoma_codec
shout
sndfile
sofia
tts_commandline
unicall
verto
voicemail
xml_cdr
xml_rpc
zeroconf
*/
func GetConfNameByStruct(structure interface{}) string {
	str, _ := GetConfNameAndInstanceByStruct(structure, nil)

	return str
}

func GetConfInstanceByStruct(structure interface{}, parent interface{}) interface{} {
	_, res := GetConfNameAndInstanceByStruct(structure, parent)
	return res
}

func GetConfNameAndInstanceByStruct(structure interface{}, par interface{}) (string, interface{}) {
	switch structure.(type) {
	case altStruct.ConfigAclList, *altStruct.ConfigAclList:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfAcl, &altStruct.ConfigAclList{Parent: parent}
	case altStruct.ConfigAclNode, *altStruct.ConfigAclNode:
		parent, ok := par.(*altStruct.ConfigAclList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfAcl, &altStruct.ConfigAclNode{Parent: parent}
	case altStruct.ConfigAlsaSetting, *altStruct.ConfigAlsaSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfAlsa, &altStruct.ConfigAlsaSetting{Parent: parent}
	case altStruct.ConfigAmrSetting, *altStruct.ConfigAmrSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfAmr, &altStruct.ConfigAmrSetting{Parent: parent}
	case altStruct.ConfigAmrwbSetting, *altStruct.ConfigAmrwbSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfAmrwb, &altStruct.ConfigAmrwbSetting{Parent: parent}
	case altStruct.ConfigAvmdSetting, *altStruct.ConfigAvmdSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfAvmd, &altStruct.ConfigAvmdSetting{Parent: parent}
	case altStruct.ConfigCallcenterSetting, *altStruct.ConfigCallcenterSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCallcenter, &altStruct.ConfigCallcenterSetting{Parent: parent}
	case altStruct.ConfigCallcenterQueue, *altStruct.ConfigCallcenterQueue:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCallcenter, &altStruct.ConfigCallcenterQueue{Parent: parent}
	case altStruct.ConfigCallcenterQueueParameter, *altStruct.ConfigCallcenterQueueParameter:
		parent, ok := par.(*altStruct.ConfigCallcenterQueue)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCallcenter, &altStruct.ConfigCallcenterQueueParameter{Parent: parent}
	case altStruct.Agent:
		/*		parent, ok := par.(string)
				if !ok {
					parent = nil
				}
				return mainStruct.ConfCallcenter, &altStruct.Agent{InstanceId: parent}*/
		return mainStruct.ConfCallcenter, &altStruct.Agent{}
	case *altStruct.Agent:
		return mainStruct.ConfCallcenter, &altStruct.Agent{}
	case altStruct.Tier:
		return mainStruct.ConfCallcenter, &altStruct.Tier{}
	case *altStruct.Tier:
		return mainStruct.ConfCallcenter, &altStruct.Tier{}
	case altStruct.Member:
		return mainStruct.ConfCallcenter, &altStruct.Member{}
	case *altStruct.Member:
		return mainStruct.ConfCallcenter, &altStruct.Member{}
	case altStruct.ConfigCdrMongodbSetting, *altStruct.ConfigCdrMongodbSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCdrMongodb, &altStruct.ConfigCdrMongodbSetting{Parent: parent}
	case altStruct.ConfigMongoSetting, *altStruct.ConfigMongoSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfMongo, &altStruct.ConfigMongoSetting{Parent: parent}
	case altStruct.ConfigCdrPgCsvSetting, *altStruct.ConfigCdrPgCsvSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCdrPgCsv, &altStruct.ConfigCdrPgCsvSetting{Parent: parent}
	case altStruct.ConfigCdrPgCsvSchema, *altStruct.ConfigCdrPgCsvSchema:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCdrPgCsv, &altStruct.ConfigCdrPgCsvSchema{Parent: parent}
	case altStruct.ConfigCespalSetting, *altStruct.ConfigCespalSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCepstral, &altStruct.ConfigCespalSetting{Parent: parent}
	case altStruct.ConfigCidlookupSetting, *altStruct.ConfigCidlookupSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCidlookup, &altStruct.ConfigCidlookupSetting{Parent: parent}
	case altStruct.ConfigConferenceAdvertiseRoom, *altStruct.ConfigConferenceAdvertiseRoom:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceAdvertiseRoom{Parent: parent}
	case altStruct.ConfigConferenceCallerControlGroup, *altStruct.ConfigConferenceCallerControlGroup:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceCallerControlGroup{Parent: parent}
	case altStruct.ConfigConferenceCallerControlGroupControl, *altStruct.ConfigConferenceCallerControlGroupControl:
		parent, ok := par.(*altStruct.ConfigConferenceCallerControlGroup)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceCallerControlGroupControl{Parent: parent}
	case altStruct.ConfigConferenceChatPermissionProfile, *altStruct.ConfigConferenceChatPermissionProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceChatPermissionProfile{Parent: parent}
	case altStruct.ConfigConferenceChatPermissionProfileUser, *altStruct.ConfigConferenceChatPermissionProfileUser:
		parent, ok := par.(*altStruct.ConfigConferenceChatPermissionProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceChatPermissionProfileUser{Parent: parent}
	case altStruct.ConfigConferenceProfile, *altStruct.ConfigConferenceProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceProfile{Parent: parent}
	case altStruct.ConfigConferenceProfileParameter, *altStruct.ConfigConferenceProfileParameter:
		parent, ok := par.(*altStruct.ConfigConferenceProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConference, &altStruct.ConfigConferenceProfileParameter{Parent: parent}
	case altStruct.ConfigConferenceLayout, *altStruct.ConfigConferenceLayout:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConferenceLayouts, &altStruct.ConfigConferenceLayout{Parent: parent}
	case altStruct.ConfigConferenceLayoutImage, *altStruct.ConfigConferenceLayoutImage:
		parent, ok := par.(*altStruct.ConfigConferenceLayout)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConferenceLayouts, &altStruct.ConfigConferenceLayoutImage{Parent: parent}
	case altStruct.ConfigConferenceLayoutGroup, *altStruct.ConfigConferenceLayoutGroup:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConferenceLayouts, &altStruct.ConfigConferenceLayoutGroup{Parent: parent}
	case altStruct.ConfigConferenceLayoutGroupLayout, *altStruct.ConfigConferenceLayoutGroupLayout:
		parent, ok := par.(*altStruct.ConfigConferenceLayoutGroup)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfConferenceLayouts, &altStruct.ConfigConferenceLayoutGroupLayout{Parent: parent}
	case altStruct.ConfigCurlSetting, *altStruct.ConfigCurlSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfCurl, &altStruct.ConfigCurlSetting{Parent: parent}
	case altStruct.ConfigDbSetting, *altStruct.ConfigDbSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDb, &altStruct.ConfigDbSetting{Parent: parent}
	case altStruct.ConfigDialplanDirectorySetting, *altStruct.ConfigDialplanDirectorySetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDialplanDirectory, &altStruct.ConfigDialplanDirectorySetting{Parent: parent}
	case altStruct.ConfigDirectorySetting, *altStruct.ConfigDirectorySetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDialplanDirectory, &altStruct.ConfigDirectorySetting{Parent: parent}
	case altStruct.ConfigDirectoryProfile, *altStruct.ConfigDirectoryProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDialplanDirectory, &altStruct.ConfigDirectoryProfile{Parent: parent}
	case altStruct.ConfigDirectoryProfileParameter, *altStruct.ConfigDirectoryProfileParameter:
		parent, ok := par.(*altStruct.ConfigDirectoryProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDialplanDirectory, &altStruct.ConfigDirectoryProfileParameter{Parent: parent}
	case altStruct.ConfigDistributorList, *altStruct.ConfigDistributorList:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDistributor, &altStruct.ConfigDistributorList{Parent: parent}
	case altStruct.ConfigDistributorListNode, *altStruct.ConfigDistributorListNode:
		parent, ok := par.(*altStruct.ConfigDistributorList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfDistributor, &altStruct.ConfigDistributorListNode{Parent: parent}
	case altStruct.ConfigEasyrouteSetting, *altStruct.ConfigEasyrouteSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfEasyroute, &altStruct.ConfigEasyrouteSetting{Parent: parent}
	case altStruct.ConfigErlangEventSetting, *altStruct.ConfigErlangEventSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfErlangEvent, &altStruct.ConfigErlangEventSetting{Parent: parent}
	case altStruct.ConfigEventMulticastSetting, *altStruct.ConfigEventMulticastSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfEventMulticast, &altStruct.ConfigEventMulticastSetting{Parent: parent}
	case altStruct.ConfigFaxSetting, *altStruct.ConfigFaxSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfFax, &altStruct.ConfigFaxSetting{Parent: parent}
	case altStruct.ConfigFifoSetting, *altStruct.ConfigFifoSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfFifo, &altStruct.ConfigFifoSetting{Parent: parent}
	case altStruct.ConfigFifoFifo, *altStruct.ConfigFifoFifo:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfFifo, &altStruct.ConfigFifoFifo{Parent: parent}
	case altStruct.ConfigFifoFifoMember, *altStruct.ConfigFifoFifoMember:
		parent, ok := par.(*altStruct.ConfigFifoFifo)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfFifo, &altStruct.ConfigFifoFifoMember{Parent: parent}
	case altStruct.ConfigHttpCacheSetting, *altStruct.ConfigHttpCacheSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfHttpCache, &altStruct.ConfigHttpCacheSetting{Parent: parent}
	case altStruct.ConfigHttpCacheProfile, *altStruct.ConfigHttpCacheProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfHttpCache, &altStruct.ConfigHttpCacheProfile{Parent: parent}
	case altStruct.ConfigHttpCacheProfileDomain, *altStruct.ConfigHttpCacheProfileDomain:
		parent, ok := par.(*altStruct.ConfigHttpCacheProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfHttpCache, &altStruct.ConfigHttpCacheProfileDomain{Parent: parent}
	case altStruct.ConfigHttpCacheProfileAWSS3, *altStruct.ConfigHttpCacheProfileAWSS3:
		parent, ok := par.(*altStruct.ConfigHttpCacheProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfHttpCache, &altStruct.ConfigHttpCacheProfileAWSS3{Parent: parent}
	case altStruct.ConfigHttpCacheProfileAzureBlob, *altStruct.ConfigHttpCacheProfileAzureBlob:
		parent, ok := par.(*altStruct.ConfigHttpCacheProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfHttpCache, &altStruct.ConfigHttpCacheProfileAzureBlob{Parent: parent}
	case altStruct.ConfigLcrSetting, *altStruct.ConfigLcrSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfLcr, &altStruct.ConfigLcrSetting{Parent: parent}
	case altStruct.ConfigLcrProfile, *altStruct.ConfigLcrProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfLcr, &altStruct.ConfigLcrProfile{Parent: parent}
	case altStruct.ConfigLcrProfileParameter, *altStruct.ConfigLcrProfileParameter:
		parent, ok := par.(*altStruct.ConfigLcrProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfLcr, &altStruct.ConfigLcrProfileParameter{Parent: parent}
	case altStruct.ConfigLuaSetting, *altStruct.ConfigLuaSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfLua, &altStruct.ConfigLuaSetting{Parent: parent}
	case altStruct.ConfigMemcacheSetting, *altStruct.ConfigMemcacheSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfMemcache, &altStruct.ConfigMemcacheSetting{Parent: parent}
	case altStruct.ConfigMsrpSetting, *altStruct.ConfigMsrpSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfMsrp, &altStruct.ConfigMsrpSetting{Parent: parent}
	case altStruct.ConfigNibblebillSetting, *altStruct.ConfigNibblebillSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfNibblebill, &altStruct.ConfigNibblebillSetting{Parent: parent}
	case altStruct.ConfigOdbcCdrSetting, *altStruct.ConfigOdbcCdrSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOdbcCdr, &altStruct.ConfigOdbcCdrSetting{Parent: parent}
	case altStruct.ConfigOdbcCdrTable, *altStruct.ConfigOdbcCdrTable:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOdbcCdr, &altStruct.ConfigOdbcCdrTable{Parent: parent}
	case altStruct.ConfigOdbcCdrTableField, *altStruct.ConfigOdbcCdrTableField:
		parent, ok := par.(*altStruct.ConfigOdbcCdrTable)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOdbcCdr, &altStruct.ConfigOdbcCdrTableField{Parent: parent}
	case altStruct.ConfigOpalSetting, *altStruct.ConfigOpalSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOpal, &altStruct.ConfigOpalSetting{Parent: parent}
	case altStruct.ConfigOpalListener, *altStruct.ConfigOpalListener:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOpal, &altStruct.ConfigOpalListener{Parent: parent}
	case altStruct.ConfigOpalListenerParameter, *altStruct.ConfigOpalListenerParameter:
		parent, ok := par.(*altStruct.ConfigOpalListener)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOpal, &altStruct.ConfigOpalListenerParameter{Parent: parent}
	case altStruct.ConfigOpusSetting, *altStruct.ConfigOpusSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOpus, &altStruct.ConfigOpusSetting{Parent: parent}
	case altStruct.ConfigOrekaSetting, *altStruct.ConfigOrekaSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOreka, &altStruct.ConfigOrekaSetting{Parent: parent}
	case altStruct.ConfigOspSetting, *altStruct.ConfigOspSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOsp, &altStruct.ConfigOspSetting{Parent: parent}
	case altStruct.ConfigOspProfile, *altStruct.ConfigOspProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOsp, &altStruct.ConfigOspProfile{Parent: parent}
	case altStruct.ConfigOspProfileParameter, *altStruct.ConfigOspProfileParameter:
		parent, ok := par.(*altStruct.ConfigOspProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfOsp, &altStruct.ConfigOspProfileParameter{Parent: parent}
	case altStruct.ConfigPerlSetting, *altStruct.ConfigPerlSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPerl, &altStruct.ConfigPerlSetting{Parent: parent}
	case altStruct.ConfigPocketsphinxSetting, *altStruct.ConfigPocketsphinxSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPocketsphinx, &altStruct.ConfigPocketsphinxSetting{Parent: parent}
	case altStruct.ConfigPostLoadSwitchSetting, *altStruct.ConfigPostLoadSwitchSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPostLoadSwitch, &altStruct.ConfigPostLoadSwitchSetting{Parent: parent}
	case altStruct.ConfigPostLoadSwitchCliKeybinding, *altStruct.ConfigPostLoadSwitchCliKeybinding:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPostLoadSwitch, &altStruct.ConfigPostLoadSwitchCliKeybinding{Parent: parent}
	case altStruct.ConfigPostLoadSwitchDefaultPtime, *altStruct.ConfigPostLoadSwitchDefaultPtime:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPostLoadSwitch, &altStruct.ConfigPostLoadSwitchDefaultPtime{Parent: parent}
	case altStruct.ConfigPostLoadModule, *altStruct.ConfigPostLoadModule:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPostLoadModules, &altStruct.ConfigPostLoadModule{Parent: parent}
	case altStruct.ConfigPythonSetting, *altStruct.ConfigPythonSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfPython, &altStruct.ConfigPythonSetting{Parent: parent}
	case altStruct.ConfigRedisSetting, *altStruct.ConfigRedisSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfRedis, &altStruct.ConfigRedisSetting{Parent: parent}
	case altStruct.ConfigSangomaCodecSetting, *altStruct.ConfigSangomaCodecSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfRedis, &altStruct.ConfigSangomaCodecSetting{Parent: parent}
	case altStruct.ConfigShoutSetting, *altStruct.ConfigShoutSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfShout, &altStruct.ConfigShoutSetting{Parent: parent}
	case altStruct.ConfigSndfileSetting, *altStruct.ConfigSndfileSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSndfile, &altStruct.ConfigSndfileSetting{Parent: parent}
	case altStruct.ConfigSofiaGlobalSetting, *altStruct.ConfigSofiaGlobalSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaGlobalSetting{Parent: parent}
	case altStruct.ConfigSofiaProfile, *altStruct.ConfigSofiaProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfile{Parent: parent}
	case altStruct.ConfigSofiaProfileAlias, *altStruct.ConfigSofiaProfileAlias:
		parent, ok := par.(*altStruct.ConfigSofiaProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfileAlias{Parent: parent}
	case altStruct.ConfigSofiaProfileDomain, *altStruct.ConfigSofiaProfileDomain:
		parent, ok := par.(*altStruct.ConfigSofiaProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfileDomain{Parent: parent}
	case altStruct.ConfigSofiaProfileParameter, *altStruct.ConfigSofiaProfileParameter:
		parent, ok := par.(*altStruct.ConfigSofiaProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfileParameter{Parent: parent}
	case altStruct.ConfigSofiaProfileGateway, *altStruct.ConfigSofiaProfileGateway:
		parent, ok := par.(*altStruct.ConfigSofiaProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfileGateway{Parent: parent}
	case altStruct.ConfigSofiaProfileGatewayParameter, *altStruct.ConfigSofiaProfileGatewayParameter:
		parent, ok := par.(*altStruct.ConfigSofiaProfileGateway)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfileGatewayParameter{Parent: parent}
	case altStruct.ConfigSofiaProfileGatewayVariable, *altStruct.ConfigSofiaProfileGatewayVariable:
		parent, ok := par.(*altStruct.ConfigSofiaProfileGateway)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfSofia, &altStruct.ConfigSofiaProfileGatewayVariable{Parent: parent}
	case altStruct.ConfigTtsCommandlineSetting, *altStruct.ConfigTtsCommandlineSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfTtsCommandline, &altStruct.ConfigTtsCommandlineSetting{Parent: parent}
	case altStruct.ConfigUnicallSetting, *altStruct.ConfigUnicallSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfUnicall, &altStruct.ConfigUnicallSetting{Parent: parent}
	case altStruct.ConfigUnicallSpan, *altStruct.ConfigUnicallSpan:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfUnicall, &altStruct.ConfigUnicallSpan{Parent: parent}
	case altStruct.ConfigUnicallSpanParameter, *altStruct.ConfigUnicallSpanParameter:
		parent, ok := par.(*altStruct.ConfigUnicallSpan)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfUnicall, &altStruct.ConfigUnicallSpanParameter{Parent: parent}
	case altStruct.ConfigVertoSetting, *altStruct.ConfigVertoSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVerto, &altStruct.ConfigVertoSetting{Parent: parent}
	case altStruct.ConfigVertoProfile, *altStruct.ConfigVertoProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVerto, &altStruct.ConfigVertoProfile{Parent: parent}
	case altStruct.ConfigVertoProfileParameter, *altStruct.ConfigVertoProfileParameter:
		parent, ok := par.(*altStruct.ConfigVertoProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVerto, &altStruct.ConfigVertoProfileParameter{Parent: parent}
	case altStruct.ConfigVoicemailSetting, *altStruct.ConfigVoicemailSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVoicemail, &altStruct.ConfigVoicemailSetting{Parent: parent}
	case altStruct.ConfigVoicemailProfile, *altStruct.ConfigVoicemailProfile:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVoicemail, &altStruct.ConfigVoicemailProfile{Parent: parent}
	case altStruct.ConfigVoicemailProfileParameter, *altStruct.ConfigVoicemailProfileParameter:
		parent, ok := par.(*altStruct.ConfigVoicemailProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVoicemail, &altStruct.ConfigVoicemailProfileParameter{Parent: parent}
	case altStruct.ConfigVoicemailProfileEmailParameter, *altStruct.ConfigVoicemailProfileEmailParameter:
		parent, ok := par.(*altStruct.ConfigVoicemailProfile)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfVoicemail, &altStruct.ConfigVoicemailProfileEmailParameter{Parent: parent}
	case altStruct.ConfigXmlRpcSetting, *altStruct.ConfigXmlRpcSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfXmlRpc, &altStruct.ConfigXmlRpcSetting{Parent: parent}
	case altStruct.ConfigXmlCdrSetting, *altStruct.ConfigXmlCdrSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfXmlCdr, &altStruct.ConfigXmlCdrSetting{Parent: parent}
	case altStruct.ConfigZeroconfSetting, *altStruct.ConfigZeroconfSetting:
		parent, ok := par.(*altStruct.ConfigurationsList)
		if !ok {
			parent = nil
		}
		return mainStruct.ConfZeroconf, &altStruct.ConfigZeroconfSetting{Parent: parent}
	default:
		return "", nil
	}
}

func castConfigurationsList(par interface{}) *altStruct.ConfigurationsList {
	parent, ok := par.(*altStruct.ConfigurationsList)
	if ok {
		return parent
	}
	return nil
}

func getCurrentModule(confName, modName string) (*altStruct.ConfigurationsList, error) {
	conf, err := getCurrentConfigByName(confName)
	if err != nil {
		return nil, err
	}
	if conf == nil {
		return nil, errors.New("no config")
	}
	conf.Module = modName
	return conf, err
}

func getCurrentConfigByName(name string) (*altStruct.ConfigurationsList, error) {
	if name == "" {
		return nil, errors.New("no name")
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}, Name: name},
		map[string]bool{"Parent": true, "Name": true},
	)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("no config")
	}

	item, ok := res[0].(altStruct.ConfigurationsList)
	if !ok {
		return nil, errors.New("no config")
	}

	return &item, nil
}

func GetCurrentConfPostSwitch() (*altStruct.ConfigurationsList, error) {
	return getCurrentConfigByName(mainStruct.ConfPostLoadSwitch)
}

func XMLPostSwitch(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPostLoadSwitchSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	keyBindings, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPostLoadSwitchSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPostLoadSwitchSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Post load switch Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"cli-keybindings,omitempty"`
				Inner   interface{} `xml:"key"`
			}{Inner: &keyBindings},
			struct {
				XMLName xml.Name    `xml:"default-ptimes,omitempty"`
				Inner   interface{} `xml:"codec"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLAcl(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	lists, _ := intermediateDB.GetByValue(
		&altStruct.ConfigAclList{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	var listLists []interface{}
	for _, ilist := range lists {
		list, ok := ilist.(altStruct.ConfigAclList)
		if !ok {
			continue
		}
		nodes, _ := intermediateDB.GetByValue(
			&altStruct.ConfigAclNode{Parent: &altStruct.ConfigAclList{Id: list.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigAclList
			Nodes interface{} `xml:"node"`
		}{
			&list,
			nodes,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "ACL Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"network-lists,omitempty"`
				Inner   interface{} `xml:"list"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLCallcenter(name string, arg string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCallcenterSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	filter := map[string]bool{"Parent": true, "Enabled": true}
	if arg != "" {
		filter["Name"] = true
	}
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCallcenterQueue{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Name: arg},
		filter,
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigCallcenterQueue)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigCallcenterQueueParameter{Parent: &altStruct.ConfigCallcenterQueue{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigCallcenterQueue
			Params interface{} `xml:"param"`
		}{
			&profile,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Callcenter Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"queues,omitempty"`
				Inner   interface{} `xml:"queue"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLCdrPgCsv(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCdrPgCsvSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	schema, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCdrPgCsvSchema{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "CdrPgCsv Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"schema,omitempty"`
				Inner   interface{} `xml:"field"`
			}{Inner: &schema},
		},
	}
	return &currentConfig
}

func XMLOdbcCdr(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOdbcCdrSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	tables, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOdbcCdrTable{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, table := range tables {
		table, ok := table.(altStruct.ConfigOdbcCdrTable)
		if !ok {
			continue
		}
		fields, _ := intermediateDB.GetByValue(
			&altStruct.ConfigOdbcCdrTableField{Parent: &altStruct.ConfigOdbcCdrTable{Id: table.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigOdbcCdrTable
			Fields interface{} `xml:"field"`
		}{
			&table,
			fields,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "OdbcCdr Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"tables,omitempty"`
				Inner   interface{} `xml:"table"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLSofia(name string, arg string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSofiaGlobalSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	filter := map[string]bool{"Parent": true, "Enabled": true}
	if arg != "" {
		filter["Name"] = true
	}
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSofiaProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Name: arg, Enabled: true},
		filter,
	)
	var listLists []interface{}
	for _, profileI := range profiles {
		profile, ok := profileI.(altStruct.ConfigSofiaProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigSofiaProfileParameter{Parent: &altStruct.ConfigSofiaProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		profileAliases, _ := intermediateDB.GetByValue(
			&altStruct.ConfigSofiaProfileAlias{Parent: &altStruct.ConfigSofiaProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		profileDomains, _ := intermediateDB.GetByValue(
			&altStruct.ConfigSofiaProfileDomain{Parent: &altStruct.ConfigSofiaProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		profileGateways, _ := intermediateDB.GetByValue(
			&altStruct.ConfigSofiaProfileGateway{Parent: &altStruct.ConfigSofiaProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)

		var gatewayLists []interface{}
		for _, gateway := range profileGateways {
			gateway, ok := gateway.(altStruct.ConfigSofiaProfileGateway)
			if !ok {
				continue
			}
			gatewayParams, _ := intermediateDB.GetByValue(
				&altStruct.ConfigSofiaProfileGatewayParameter{Parent: &altStruct.ConfigSofiaProfileGateway{Id: gateway.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
			gatewayVars, _ := intermediateDB.GetByValue(
				&altStruct.ConfigSofiaProfileGatewayVariable{Parent: &altStruct.ConfigSofiaProfileGateway{Id: gateway.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
			gatewayLists = append(listLists, struct {
				*altStruct.ConfigSofiaProfileGateway
				Params interface{} `xml:"param"`
				Vars   interface{} `xml:"variable"`
			}{
				&gateway,
				gatewayParams,
				gatewayVars,
			})
		}
		listLists = append(listLists, struct {
			*altStruct.ConfigSofiaProfile
			Params   interface{} `xml:"settings>param"`
			Aliases  interface{} `xml:"aliases>alias"`
			Domains  interface{} `xml:"domains>domain"`
			Gateways interface{} `xml:"gateways>gateway"`
		}{
			&profile,
			profileParams,
			profileAliases,
			profileDomains,
			gatewayLists,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Sofia Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"global_settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLVerto(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigVertoSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigVertoProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigVertoProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigVertoProfileParameter{Parent: &altStruct.ConfigVertoProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigVertoProfile
			Params interface{} `xml:"param"`
		}{
			&profile,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Verto Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLLcr(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigLcrSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigLcrProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigLcrProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigLcrProfileParameter{Parent: &altStruct.ConfigLcrProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigLcrProfile
			Params interface{} `xml:"param"`
		}{
			&profile,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Lcr Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLShout(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigShoutSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Shout Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLRedis(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigRedisSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Redis Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLNibblebill(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigNibblebillSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Nibblebill Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLDb(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigDbSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Db Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLMemcache(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigMemcacheSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Memcache Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLAvmd(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigAvmdSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Avmd Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLTtsCommandline(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigTtsCommandlineSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "TtsCommandline Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLCdrMongodb(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCdrMongodbSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "CdrMongodb Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLHttpCache(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigHttpCacheSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigHttpCacheProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigHttpCacheProfile)
		if !ok {
			continue
		}
		profileAWS, _ := intermediateDB.GetByValue(
			&altStruct.ConfigHttpCacheProfileAWSS3{Parent: &altStruct.ConfigHttpCacheProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		profileAZURE, _ := intermediateDB.GetByValue(
			&altStruct.ConfigHttpCacheProfileAzureBlob{Parent: &altStruct.ConfigHttpCacheProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		profileDomains, _ := intermediateDB.GetByValue(
			&altStruct.ConfigHttpCacheProfileDomain{Parent: &altStruct.ConfigHttpCacheProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigHttpCacheProfile
			AWS     interface{} `xml:"aws-s3"`
			AZURE   interface{} `xml:"azure-blob"`
			Domains interface{} `xml:"domains>domain"`
		}{
			&profile,
			profileAWS,
			profileAZURE,
			profileDomains,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "HttpCache Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLOpus(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOpusSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Opus Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLPython(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPythonSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Python Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLAlsa(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigAlsaSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Alsa Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLAmr(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigAmrSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Amr Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLAmrwb(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigAmrwbSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Amrwb Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLCepstral(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCespalSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Cespal Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLCidlookup(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCidlookupSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Cidlookup Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLCurl(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigCurlSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Curl Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLDialplanDirectory(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigDialplanDirectorySetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "DialplanDirectory Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLEasyroute(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigEasyrouteSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Easyroute Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLErlangEvent(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigErlangEventSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "ErlangEvent Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLEventMulticast(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigEventMulticastSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "EventMulticast Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLFax(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigFaxSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Fax Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLLua(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigLuaSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Lua Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLMongo(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigMongoSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Mongo Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLMsrp(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigMsrpSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Msrp Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLOreka(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOrekaSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Oreka Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLPerl(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPerlSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Perl Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLPocketsphinx(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPocketsphinxSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Pocketsphinx Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLSangomaCodec(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSangomaCodecSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "SangomaCodec Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLSndfile(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSndfileSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Sndfile Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLXmlCdr(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigXmlCdrSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "XmlCdr Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLXmlRpc(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigXmlRpcSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "XmlRpc Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLZeroconf(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	ptimes, _ := intermediateDB.GetByValue(
		&altStruct.ConfigZeroconfSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Zeroconf Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &ptimes},
		},
	}
	return &currentConfig
}

func XMLDirectory(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigDirectorySetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigDirectoryProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigDirectoryProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigDirectoryProfileParameter{Parent: &altStruct.ConfigDirectoryProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigDirectoryProfile
			Params interface{} `xml:"param"`
		}{
			&profile,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Directory Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLFifo(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigFifoSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	fifos, _ := intermediateDB.GetByValue(
		&altStruct.ConfigFifoFifo{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, fifo := range fifos {
		fifo, ok := fifo.(altStruct.ConfigFifoFifo)
		if !ok {
			continue
		}
		fifoMembers, _ := intermediateDB.GetByValue(
			&altStruct.ConfigFifoFifoMember{Parent: &altStruct.ConfigFifoFifo{Id: fifo.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigFifoFifo
			Params interface{} `xml:"member"`
		}{
			&fifo,
			fifoMembers,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Fifo Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"fifos,omitempty"`
				Inner   interface{} `xml:"fifo"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLOpal(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOpalSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	listeners, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOpalListener{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, listener := range listeners {
		listener, ok := listener.(altStruct.ConfigOpalListener)
		if !ok {
			continue
		}
		listenerParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigOpalListenerParameter{Parent: &altStruct.ConfigOpalListener{Id: listener.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigOpalListener
			Params interface{} `xml:"param"`
		}{
			&listener,
			listenerParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Opal Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"listeners,omitempty"`
				Inner   interface{} `xml:"listener"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLOsp(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOspSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigOspProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigOspProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigOspProfileParameter{Parent: &altStruct.ConfigOspProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigOspProfile
			Params interface{} `xml:"param"`
		}{
			&profile,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Osp Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLUnicall(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigUnicallSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	spans, _ := intermediateDB.GetByValue(
		&altStruct.ConfigUnicallSpan{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, span := range spans {
		span, ok := span.(altStruct.ConfigUnicallSpan)
		if !ok {
			continue
		}
		spanParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigUnicallSpanParameter{Parent: &altStruct.ConfigUnicallSpan{Id: span.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigUnicallSpan
			Params interface{} `xml:"param"`
		}{
			&span,
			spanParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Unicall Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"spans,omitempty"`
				Inner   interface{} `xml:"span"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLConference(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	status, _ := intermediateDB.GetByValue(
		&altStruct.ConfigConferenceAdvertiseRoom{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	ccgroups, _ := intermediateDB.GetByValue(
		&altStruct.ConfigConferenceCallerControlGroup{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listControlGroups []interface{}
	for _, group := range ccgroups {
		group, ok := group.(altStruct.ConfigConferenceCallerControlGroup)
		if !ok {
			continue
		}
		groupControl, _ := intermediateDB.GetByValue(
			&altStruct.ConfigConferenceCallerControlGroupControl{Parent: &altStruct.ConfigConferenceCallerControlGroup{Id: group.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listControlGroups = append(listControlGroups, struct {
			*altStruct.ConfigConferenceCallerControlGroup
			Controls interface{} `xml:"control"`
		}{
			&group,
			groupControl,
		})
	}
	ccpprofiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigConferenceChatPermissionProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listPermissionProfiles []interface{}
	for _, profile := range ccpprofiles {
		profile, ok := profile.(altStruct.ConfigConferenceChatPermissionProfile)
		if !ok {
			continue
		}
		users, _ := intermediateDB.GetByValue(
			&altStruct.ConfigConferenceChatPermissionProfileUser{Parent: &altStruct.ConfigConferenceChatPermissionProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listPermissionProfiles = append(listPermissionProfiles, struct {
			*altStruct.ConfigConferenceChatPermissionProfile
			Params interface{} `xml:"user"`
		}{
			&profile,
			users,
		})
	}
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigConferenceProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigConferenceProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigConferenceProfileParameter{Parent: &altStruct.ConfigConferenceProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigConferenceProfile
			Params interface{} `xml:"param"`
		}{
			&profile,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Conference Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"advertise,omitempty"`
				Inner   interface{} `xml:"room"`
			}{Inner: &status},
			struct {
				XMLName xml.Name    `xml:"caller-controls,omitempty"`
				Inner   interface{} `xml:"group"`
			}{Inner: &listControlGroups},
			struct {
				XMLName xml.Name    `xml:"chat-permissions,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{Inner: &listPermissionProfiles},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{Inner: &listLists},
		},
	}
	return &currentConfig
}

func XMLConferenceLayouts(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	llayouts, _ := intermediateDB.GetByValue(
		&altStruct.ConfigConferenceLayout{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLayouts []interface{}
	for _, llayout := range llayouts {
		llayout, ok := llayout.(altStruct.ConfigConferenceLayout)
		if !ok {
			continue
		}
		layoutlayout, _ := intermediateDB.GetByValue(
			&altStruct.ConfigConferenceLayoutImage{Parent: &altStruct.ConfigConferenceLayout{Id: llayout.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLayouts = append(listLayouts, struct {
			*altStruct.ConfigConferenceLayout
			Controls interface{} `xml:"image"`
		}{
			&llayout,
			layoutlayout,
		})
	}
	lgroups, _ := intermediateDB.GetByValue(
		&altStruct.ConfigConferenceLayoutGroup{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, group := range lgroups {
		group, ok := group.(altStruct.ConfigConferenceLayoutGroup)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigConferenceLayoutGroupLayout{Parent: &altStruct.ConfigConferenceLayoutGroup{Id: group.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigConferenceLayoutGroup
			Params interface{} `xml:"layout"`
		}{
			&group,
			profileParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "ConferenceLayouts Config",
		AnyXML: struct {
			T interface{} `xml:"layout-settings"`
		}{
			[]interface{}{
				struct {
					XMLName xml.Name    `xml:"layouts,omitempty"`
					Inner   interface{} `xml:"layout"`
				}{Inner: &listLayouts},
				struct {
					XMLName xml.Name    `xml:"groups,omitempty"`
					Inner   interface{} `xml:"group"`
				}{Inner: &listLists},
			},
		},
	}
	return &currentConfig
}

func XMLPostLoadModules(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	modules, _ := intermediateDB.GetByValue(
		&altStruct.ConfigPostLoadModule{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "PostLoadModule Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"modules,omitempty"`
				Inner   interface{} `xml:"load"`
			}{Inner: &modules},
		},
	}
	return &currentConfig
}

func XMLVoicemail(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigVoicemailSetting{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigVoicemailProfile{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, profile := range profiles {
		profile, ok := profile.(altStruct.ConfigVoicemailProfile)
		if !ok {
			continue
		}
		profileParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigVoicemailProfileParameter{Parent: &altStruct.ConfigVoicemailProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		profileEmailParams, _ := intermediateDB.GetByValue(
			&altStruct.ConfigVoicemailProfileEmailParameter{Parent: &altStruct.ConfigVoicemailProfile{Id: profile.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigVoicemailProfile
			Params      interface{} `xml:"param"`
			EmailParams interface{} `xml:"email>param"`
		}{
			&profile,
			profileParams,
			profileEmailParams,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Voicemail Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"settings,omitempty"`
				Inner   interface{} `xml:"param"`
			}{Inner: &params},
			struct {
				XMLName xml.Name    `xml:"profiles,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func XMLDistributor(name string) *mainStruct.Configuration {
	c, err := getCurrentConfigByName(name)
	if err != nil || c == nil || !c.Enabled {
		return nil
	}
	lists, _ := intermediateDB.GetByValue(
		&altStruct.ConfigDistributorList{Parent: &altStruct.ConfigurationsList{Id: c.Id}, Enabled: true},
		map[string]bool{"Parent": true, "Enabled": true},
	)

	var listLists []interface{}
	for _, list := range lists {
		list, ok := list.(altStruct.ConfigDistributorList)
		if !ok {
			return nil
		}
		listNodes, _ := intermediateDB.GetByValue(
			&altStruct.ConfigDistributorListNode{Parent: &altStruct.ConfigDistributorList{Id: list.Id}, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		listLists = append(listLists, struct {
			*altStruct.ConfigDistributorList
			Params interface{} `xml:"node"`
		}{
			&list,
			listNodes,
		})
	}
	currentConfig := mainStruct.Configuration{
		Name:        name,
		Description: "Distributor Config",
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"lists,omitempty"`
				Inner   interface{} `xml:"list"`
			}{
				Inner: &listLists,
			},
		},
	}
	return &currentConfig
}

func setConf(name string, unloadable bool) (*altStruct.ConfigurationsList, error) {
	c, err := getCurrentConfigByName(name)
	if c != nil {
		return nil, errors.New("already exists")
	}
	_, err = intermediateDB.InsertItem(&altStruct.ConfigurationsList{
		Enabled:    true,
		Unloadable: unloadable,
		Name:       name,
		Parent:     &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()},
	})
	if err != nil {
		return nil, err
	}

	return getCurrentConfigByName(name)
}

func SetConfAcl() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfAcl, true)
}

func SetConfCdrPgCsv() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfCdrPgCsv, false)
}

func SetConfSofia() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfSofia, false)
}

func SetConfVerto() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfVerto, false)
}

func SetConfCallcenter() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfCallcenter, false)
}

func SetConfOdbcCdr() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfOdbcCdr, false)
}

func SetConfLcr() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfLcr, false)
}

func SetConfShout() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfShout, false)
}

func SetConfRedis() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfRedis, false)
}

func SetConfNibblebill() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfNibblebill, false)
}

func SetConfDb() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfDb, false)
}

func SetConfMemcache() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfMemcache, false)
}

func SetConfAvmd() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfAvmd, false)
}

func SetConfTtsCommandline() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfTtsCommandline, false)
}

func SetConfCdrMongodb() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfCdrMongodb, false)
}

func SetConfHttpCache() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfHttpCache, false)
}

func SetConfOpus() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfOpus, false)
}

func SetConfPython() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfPython, false)
}

func SetConfAlsa() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfAlsa, false)
}

func SetConfAmr() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfAmr, false)
}

func SetConfAmrwb() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfAmrwb, false)
}

func SetConfCepstral() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfCepstral, false)
}

func SetConfCidlookup() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfCidlookup, false)
}

func SetConfCurl() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfCurl, false)
}

func SetConfDialplanDirectory() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfDialplanDirectory, false)
}

func SetConfEasyroute() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfEasyroute, false)
}

func SetConfErlangEvent() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfErlangEvent, false)
}

func SetConfEventMulticast() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfEventMulticast, false)
}

func SetConfFax() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfFax, false)
}

func SetConfLua() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfLua, false)
}

func SetConfMongo() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfMongo, false)
}

func SetConfMsrp() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfMsrp, false)
}

func SetConfOreka() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfOreka, false)
}

func SetConfPerl() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfPerl, false)
}

func SetConfPocketsphinx() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfPocketsphinx, false)
}

func SetConfSangomaCodec() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfSangomaCodec, false)
}

func SetConfSndfile() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfSndfile, false)
}

func SetConfXmlCdr() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfXmlCdr, false)
}

func SetConfXmlRpc() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfXmlRpc, false)
}

func SetConfZeroconf() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfZeroconf, false)
}

func SetConfDistributor() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfDistributor, false)
}

func SetConfOpal() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfOpal, false)
}

func SetConfUnicall() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfUnicall, false)
}

func SetConfDirectory() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfDirectory, false)
}

func SetConfFifo() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfFifo, false)
}

func SetConfOsp() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfOsp, false)
}

func SetConfConference() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfConference, false)
}

func SetConfConferenceLayouts() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfConferenceLayouts, false)
}

func SetConfPostLoadModules() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfPostLoadModules, true)
}

func SetConfVoicemail() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfVoicemail, false)
}

func SetConfPostSwitch() (*altStruct.ConfigurationsList, error) {
	return setConf(mainStruct.ConfPostLoadSwitch, false)
}

func SetConfAclList(c *altStruct.ConfigurationsList, listNname, listDefault string) (int64, error) {
	if c.Name != mainStruct.ConfAcl {
		return 0, errors.New("wrong config")
	}
	return intermediateDB.InsertItem(&altStruct.ConfigAclList{
		Name:    listNname,
		Default: listDefault,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfAclNode(parentId int64, nodeType, cidr, domain string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigAclNode{
		Type:    nodeType,
		Cidr:    cidr,
		Domain:  domain,
		Parent:  &altStruct.ConfigAclList{Id: parentId},
		Enabled: true,
	})
}

func SetConfigSofiaGlobalSetting(c *altStruct.ConfigurationsList, paramName, paramValue string) (int64, error) {
	if c.Name != mainStruct.ConfSofia {
		return 0, errors.New("wrong config")
	}
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaGlobalSetting{
		Name:    paramName,
		Value:   paramValue,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigSofiaProfile(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	if c.Name != mainStruct.ConfSofia {
		return 0, errors.New("wrong config")
	}
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfile{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigSofiaProfileAliases(parentId int64, aliasName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfileAlias{
		Name:    aliasName,
		Parent:  &altStruct.ConfigSofiaProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigSofiaGateway(parentId int64, gatewayName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfileGateway{
		Name:    gatewayName,
		Parent:  &altStruct.ConfigSofiaProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigSofiaGatewayParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfileGatewayParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigSofiaProfileGateway{Id: parentId},
		Enabled: true,
	})
}

func SetConfigSofiaGatewayVar(parentId int64, name, value, direction string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfileGatewayParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigSofiaProfileGateway{Id: parentId},
		Enabled: true,
	})
}

func SetConfigSofiaProfileDomain(parentId int64, domainName, alias, parse string) (int64, error) {
	var aliasV bool
	var parseV bool
	if alias == "true" {
		aliasV = true
	}
	if parse == "true" {
		parseV = true
	}

	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfileDomain{
		Name:    domainName,
		Alias:   aliasV,
		Parse:   parseV,
		Parent:  &altStruct.ConfigSofiaProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigSofiaProfileParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSofiaProfileParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigSofiaProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfCdrPgCsvSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCdrPgCsvSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfCdrPgCsvSchemaField(c *altStruct.ConfigurationsList, variable, colunm, quote string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCdrPgCsvSchema{
		Var:     variable,
		Column:  colunm,
		Quote:   quote,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigVertoSetting(c *altStruct.ConfigurationsList, paramName, paramValue string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVertoSetting{
		Name:    paramName,
		Value:   paramValue,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigVertoProfile(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVertoProfile{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigVertoProfileParam(parentId int64, name, value, secure string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVertoProfileParameter{
		Name:    name,
		Value:   value,
		Secure:  secure,
		Parent:  &altStruct.ConfigVertoProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfOdbcCdrSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOdbcCdrSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfOdbcCdrTable(c *altStruct.ConfigurationsList, name, logLeg string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOdbcCdrTable{
		Name:    name,
		LogLeg:  logLeg,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfOdbcCdrTableField(parentId int64, name, chanVarName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOdbcCdrTableField{
		Name:        name,
		ChanVarName: chanVarName,
		Parent:      &altStruct.ConfigOdbcCdrTable{Id: parentId},
		Enabled:     true,
	})
}

func SetConfigLcrSetting(c *altStruct.ConfigurationsList, paramName, paramValue string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigLcrSetting{
		Name:    paramName,
		Value:   paramValue,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigLcrProfile(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigLcrProfile{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigLcrProfileParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigLcrProfileParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigLcrProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigShoutSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigShoutSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigRedisSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigRedisSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigNibblebillSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigNibblebillSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigDbSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDbSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigMemcacheSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigMemcacheSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigAvmdSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigAvmdSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigTtsCommandlineSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigTtsCommandlineSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigCdrMongodbSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCdrMongodbSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigHttpCacheSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigHttpCacheSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigHttpCacheProfile(c *altStruct.ConfigurationsList, name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigHttpCacheProfile{
		Name:    name,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigHttpCacheProfileAWSS3(parentId int64, AccessKeyId, SecretAccessKey, BaseDomain, Region string, Expires int64) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigHttpCacheProfileAWSS3{
		AccessKeyId:     AccessKeyId,
		SecretAccessKey: SecretAccessKey,
		BaseDomain:      BaseDomain,
		Region:          Region,
		Expires:         Expires,
		Parent:          &altStruct.ConfigHttpCacheProfile{Id: parentId},
		Enabled:         true,
	})
}

func SetConfigHttpCacheProfileAzureBlob(parentId int64, SecretAccessKey string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigHttpCacheProfileAzureBlob{
		SecretAccessKey: SecretAccessKey,
		Parent:          &altStruct.ConfigHttpCacheProfile{Id: parentId},
		Enabled:         true,
	})
}

func SetConfigHttpCacheProfileDomain(parentId int64, name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigHttpCacheProfileDomain{
		Name:    name,
		Parent:  &altStruct.ConfigHttpCacheProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigOpusSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOpusSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigPythonSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPythonSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigAlsaSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigAlsaSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigAmrSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigAmrSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigAmrwbSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigAmrwbSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigCepstralSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCespalSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigCidlookupSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCidlookupSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigCurlSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCurlSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigDialplanDirectorySetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDialplanDirectorySetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigEasyrouteSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigEasyrouteSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigErlangEventSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigErlangEventSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigEventMulticastSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigEventMulticastSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigFaxSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigFaxSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigLuaSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigLuaSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigMongoSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigMongoSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigMsrpSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigMsrpSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigOrekaSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOrekaSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigPerlSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPerlSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigPocketsphinxSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPocketsphinxSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigSangomaCodecSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSangomaCodecSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigSndfileSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigSndfileSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigXmlCdrSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigXmlCdrSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigXmlRpcSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigXmlRpcSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigZeroconfSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigZeroconfSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigPostSwitchSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPostLoadSwitchSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigPostSwitchCliKeybinding(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPostLoadSwitchCliKeybinding{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigPostSwitchDefaultPtime(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPostLoadSwitchDefaultPtime{
		CodecName:  name,
		CodecPtime: value,
		Parent:     c,
		Enabled:    true,
	})
}

func SetConfDistributorList(c *altStruct.ConfigurationsList, listName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDistributorList{
		Name:    listName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfDistributorNode(parentId int64, name, weight string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDistributorListNode{
		Name:    name,
		Weight:  weight,
		Parent:  &altStruct.ConfigDistributorList{Id: parentId},
		Enabled: true,
	})
}

func SetConfigDirectorySetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDirectorySetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigDirectoryProfile(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDirectoryProfile{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigDirectoryProfileParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigDirectoryProfileParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigDirectoryProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigFifoSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigFifoSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigFifoFifo(c *altStruct.ConfigurationsList, profileName, importance string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigFifoFifo{
		Name:       profileName,
		Importance: importance,
		Parent:     c,
		Enabled:    true,
	})
}

func SetConfigFifoFifoParam(parentId int64, timeout, simo, lag, body string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigFifoFifoMember{
		Timeout: timeout,
		Simo:    simo,
		Lag:     lag,
		Body:    body,
		Parent:  &altStruct.ConfigFifoFifo{Id: parentId},
		Enabled: true,
	})
}

func SetConfigOpalSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOpalSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigOpalListener(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOpalListener{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigOpalListenerParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOpalListenerParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigOpalListener{Id: parentId},
		Enabled: true,
	})
}

func SetConfigOspSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOspSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigOspProfile(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOspProfile{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigOspProfileParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigOspProfileParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigOspProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigUnicallSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigUnicallSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigUnicallSpan(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigUnicallSpan{
		SpanId:  profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfigUnicallSpanParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigUnicallSpanParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigUnicallSpan{Id: parentId},
		Enabled: true,
	})
}

func SetConfConferenceAdvertise(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceAdvertiseRoom{
		Name:    name,
		Status:  value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfConferenceCallerControlsGroup(c *altStruct.ConfigurationsList, groupName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceCallerControlGroup{
		Name:    groupName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfConferenceCallerControlsGroupControl(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceCallerControlGroupControl{
		Action:  name,
		Digits:  value,
		Parent:  &altStruct.ConfigConferenceCallerControlGroup{Id: parentId},
		Enabled: true,
	})
}

func SetConfConferenceProfile(c *altStruct.ConfigurationsList, itemName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceProfile{
		Name:    itemName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfConferenceProfileParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceProfileParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigConferenceProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfConferenceChatPermissionsProfile(c *altStruct.ConfigurationsList, profileName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceChatPermissionProfile{
		Name:    profileName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfConferenceChatPermissionsUser(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceChatPermissionProfileUser{
		Name:     name,
		Commands: value,
		Parent:   &altStruct.ConfigConferenceChatPermissionProfile{Id: parentId},
		Enabled:  true,
	})
}

func SetConfConferenceLayoutsGroups(c *altStruct.ConfigurationsList, itemName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceLayoutGroup{
		Name:    itemName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfConferenceLayoutsGroupLayout(parentId int64, body string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceLayoutGroupLayout{
		Body:    body,
		Parent:  &altStruct.ConfigConferenceLayoutGroup{Id: parentId},
		Enabled: true,
	})
}

func SetConfConferenceLayoutLayouts(c *altStruct.ConfigurationsList, itemName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceLayout{
		Name:    itemName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfConferenceLayoutLayoutsImage(parentId int64, x, y, scale, floor, floorOnly, hScale, overlap, reservationId, zoom string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigConferenceLayoutImage{
		X:             x,
		Y:             y,
		Scale:         scale,
		Floor:         floor,
		FloorOnly:     floorOnly,
		Hscale:        hScale,
		Overlap:       overlap,
		ReservationId: reservationId,
		Zoom:          zoom,
		Parent:        &altStruct.ConfigConferenceLayout{Id: parentId},
		Enabled:       true,
	})
}

func SetPostLoadModule(c *altStruct.ConfigurationsList, name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigPostLoadModule{
		Name:    name,
		Parent:  c,
		Enabled: true,
	})
}

func SetVoicemailSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVoicemailSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetVoicemailProfile(c *altStruct.ConfigurationsList, name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVoicemailProfile{
		Name:    name,
		Parent:  c,
		Enabled: true,
	})
}

func SetVoicemailProfileParam(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVoicemailProfileParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigVoicemailProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfigVoicemailProfileEmail(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigVoicemailProfileEmailParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.ConfigVoicemailProfile{Id: parentId},
		Enabled: true,
	})
}

func SetConfCallcenterSetting(c *altStruct.ConfigurationsList, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCallcenterSetting{
		Name:    name,
		Value:   value,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfCallcenterQueue(c *altStruct.ConfigurationsList, queueName string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCallcenterQueue{
		Name:    queueName,
		Parent:  c,
		Enabled: true,
	})
}

func SetConfCallcenterQueueParam(queueId int64, paramName, paramValue string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.ConfigCallcenterQueueParameter{
		Name:    paramName,
		Value:   paramValue,
		Parent:  &altStruct.ConfigCallcenterQueue{Id: queueId},
		Enabled: true,
	})
}

func SetConfCallcenterAgent(
	name,
	agentType,
	system,
	instanceId,
	uuid,
	contact,
	status,
	state string,
	maxNoAnswer,
	wrapUpTime,
	rejectDelayTime,
	busyDelayTime,
	noAnswerDelayTime,
	lastBridgeStart,
	lastBridgeEnd,
	lastOfferedCall,
	lastStatusChange,
	noAnswerCount,
	callsAnswered,
	talkTime,
	readyTime int64) (int64, error) {

	return intermediateDB.InsertItem(&altStruct.Agent{
		Name:              name,
		Status:            status,
		Contact:           contact,
		BusyDelayTime:     busyDelayTime,
		MaxNoAnswer:       maxNoAnswer,
		RejectDelayTime:   rejectDelayTime,
		WrapUpTime:        wrapUpTime,
		Type:              agentType,
		System:            system,
		InstanceId:        instanceId,
		Uuid:              uuid,
		State:             state,
		NoAnswerDelayTime: noAnswerDelayTime,
		LastBridgeStart:   lastBridgeStart,
		LastBridgeEnd:     lastBridgeEnd,
		LastOfferedCall:   lastOfferedCall,
		LastStatusChange:  lastStatusChange,
		NoAnswerCount:     noAnswerCount,
		CallsAnswered:     callsAnswered,
		TalkTime:          talkTime,
		ReadyTime:         readyTime})

	//TODO: correlate with directory user after i redo directory in the new way

	/*	if directory != nil {
		r := regexp.MustCompile(`^(.+)@(.+)$`)
		res := r.FindStringSubmatch(agent.Name)

		if len(res) != 3 || res[1] == "" || res[2] == "" {
			return agent, nil
		}
		domain := directory.Domains.GetByName(res[2])
		if domain == nil {
			return agent, nil
		}
		user := domain.Users.GetByName(res[1])
		if user == nil {
			return agent, nil
		}
		user.CCAgent = agent.Id
	}*/

}

func SetConfCallcenterTier(queue, agent, state string, position, level int64) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.Tier{Queue: queue, Agent: agent, Position: position, Level: level, State: state})
}

func SetConfCallcenterMember(
	uuid, state, queue, instanceId string, abandonedEpoch, baseScore,
	bridgeEpoch int64, cidName, cidNumber string, joinedEpoch, rejoinedEpoch int64,
	servingAgent, servingSystem, sessionUuid string, skillScore, systemEpoch int64,
) (int64, error) {

	return intermediateDB.InsertItem(&altStruct.Member{Uuid: uuid, State: state, Queue: queue, InstanceId: instanceId, AbandonedEpoch: abandonedEpoch, BaseScore: baseScore,
		BridgeEpoch: bridgeEpoch, CidName: cidName, CidNumber: cidNumber, JoinedEpoch: joinedEpoch, RejoinedEpoch: rejoinedEpoch,
		ServingAgent: servingAgent, ServingSystem: servingSystem, SessionUuid: sessionUuid, SkillScore: skillScore, SystemEpoch: systemEpoch})
}

func IsSofiaExists() bool {
	conf, err := getCurrentConfigByName(mainStruct.ConfSofia)
	if err != nil {
		return false
	}
	return conf != nil
}

func GetSofiaProfileByName(name string) *altStruct.ConfigSofiaProfile {
	conf, err := getCurrentConfigByName(mainStruct.ConfSofia)
	if err != nil {
		return nil
	}
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSofiaProfile{Parent: &altStruct.ConfigurationsList{Id: conf.Id}, Name: name},
		map[string]bool{"Parent": true, "Name": true},
	)
	if len(profiles) == 0 {
		return nil
	}
	profile, ok := profiles[0].(altStruct.ConfigSofiaProfile)
	if !ok {
		return nil
	}
	return &profile
}

func GetSofiaProfileGateways(id int64) []altStruct.ConfigSofiaProfileGateway {
	igateways, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSofiaProfileGateway{Parent: &altStruct.ConfigSofiaProfile{Id: id}},
		map[string]bool{"Parent": true},
	)
	var gateways []altStruct.ConfigSofiaProfileGateway
	for _, gateway := range igateways {
		gateway, ok := gateway.(altStruct.ConfigSofiaProfileGateway)
		if !ok {
			continue
		}
		gateways = append(gateways, gateway)
	}
	return gateways
}

func GetSofiaProfileGateway(name string) *altStruct.ConfigSofiaProfileGateway {
	gateways, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSofiaProfileGateway{Name: name},
		map[string]bool{"Name": true},
	)
	if len(gateways) == 0 {
		return nil
	}

	gateway, ok := gateways[0].(altStruct.ConfigSofiaProfileGateway)
	if !ok {
		return nil
	}
	return &gateway
	/*conf, err := getCurrentConfigByName(mainStruct.ConfSofia)
		if err != nil {
			return nil
		}

	profiles, _ := intermediateDB.GetByValueAsMap(
			&altStruct.ConfigSofiaProfile{Parent: &altStruct.ConfigurationsList{Id: conf.Id}},
			map[string]bool{"Parent": true},
		)
		if len(profiles) == 0 {
			return nil
		}

		for _, profile := range profiles {
			profileStr, ok := profile.(*altStruct.ConfigSofiaProfile)
			if !ok || profileStr == nil {
				continue
			}
			if gateway.Parent.Id != profileStr.Id {
				continue
			}
			return &gateway
		}

		return nil*/
}

func IsVertoExists() bool {
	conf, err := getCurrentConfigByName(mainStruct.ConfVerto)
	if err != nil {
		return false
	}
	return conf != nil
}

func GetVertoProfileByName(name string) *altStruct.ConfigVertoProfile {
	conf, err := getCurrentConfigByName(mainStruct.ConfVerto)
	if err != nil {
		return nil
	}
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigVertoProfile{Parent: &altStruct.ConfigurationsList{Id: conf.Id}, Name: name},
		map[string]bool{"Parent": true, "Name": true},
	)
	if len(profiles) == 0 {
		return nil
	}
	profile, ok := profiles[0].(altStruct.ConfigVertoProfile)
	if !ok {
		return nil
	}
	return &profile
}

func GetVertoProfileParamByName(id int64, name string) *altStruct.ConfigVertoProfileParameter {
	params, _ := intermediateDB.GetByValue(
		&altStruct.ConfigVertoProfileParameter{Parent: &altStruct.ConfigVertoProfile{Id: id}, Name: name},
		map[string]bool{"Parent": true, "Name": true},
	)
	if len(params) == 0 {
		return nil
	}
	param, ok := params[0].(altStruct.ConfigVertoProfileParameter)
	if !ok {
		return nil
	}
	return &param
}

func IsCallcenterEnabled() bool {
	conf, err := getCurrentConfigByName(mainStruct.ConfCallcenter)
	if err != nil {
		return false
	}
	if conf == nil {
		return false
	}
	return conf.Enabled
}

func GetCallcenterAgentByName(name string) *altStruct.Agent {
	agents, _ := intermediateDB.GetByValue(
		&altStruct.Agent{Name: name},
		map[string]bool{"Name": true},
	)
	if len(agents) == 0 {
		return nil
	}
	agent, ok := agents[0].(altStruct.Agent)
	if !ok {
		return nil
	}
	return &agent
}

func GetSofiaProfilesAndGateways() ([]interface{}, []interface{}) {
	conf, err := GetModuleByName(mainStruct.ModSofia)
	if err != nil {
		conf = &altStruct.ConfigurationsList{}
	}
	profiles, _ := intermediateDB.GetByValue(
		&altStruct.ConfigSofiaProfile{Parent: conf},
		map[string]bool{"Parent": true},
	)
	gateways, _ := intermediateDB.GetAllFromDBAsSlice(
		&altStruct.ConfigSofiaProfileGateway{},
	)
	profileIds := map[int64]bool{}

	for _, p := range profiles {
		pp, ok := p.(altStruct.ConfigSofiaProfile)
		if !ok {
			continue
		}
		profileIds[pp.Id] = true
	}

	for i, g := range gateways {
		gg, ok := g.(altStruct.ConfigSofiaProfileGateway)
		if !ok {
			continue
		}
		if profileIds[gg.Parent.Id] {
			continue
		}
		gateways[i] = gateways[len(gateways)-1]
		gateways[len(gateways)-1] = nil
		gateways = gateways[:len(gateways)-1]
	}

	return profiles, gateways
}
