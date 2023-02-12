package pbxcache

import (
	"custompbx/cache"
	"custompbx/db"
	"custompbx/mainStruct"
	"errors"
)

var directory *mainStruct.DirectoryItems
var configs *mainStruct.Configurations
var dialplan *mainStruct.Dialplans
var channels *mainStruct.Channels
var globalVariables *mainStruct.GlobalVariables

func InitCacheObjects() {
	//directory = mainStruct.NewDirectoryItems()
	//configs = &mainStruct.Configurations{}
	dialplan = mainStruct.NewDialplanItems()
	channels = mainStruct.NewChannelsCache()
	globalVariables = mainStruct.NewGlobalVariables()
}

func InitRootDB() {
	db.InitRootDB()
}

func InitPBXCache() {
	//InitDirectoryCache()
	//InitConfigurationCache()
	InitDialplanCache()
	InitGlobalVariablesCache()
}

func GetGlobalVariables() *mainStruct.GlobalVariables {
	return globalVariables
}

func GetGlobalVariableByName(name string) *mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetByName(name)
}

func GetGlobalVariableById(id int64) *mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetById(id)
}

func GetGlobalVariableNamedList() map[string]*mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetNamedList()
}

func GetGlobalVariableList() map[int64]*mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetList()
}

func GetGlobalVariableNotDynamicsProps() []*mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.NotDynamicsProps()
}

func UpdateFSInstanceDescription(instance *mainStruct.FsInstance, description string) error {
	if instance == nil {
		return errors.New("no id")
	}

	err := db.UpdateFSInstanceDescription(instance.Id, description)
	if err != nil {
		return err
	}
	instance.Description = description
	return nil
}

func InitGlobalVariablesCache() {
	db.GetGlobalVariables(globalVariables, cache.GetCurrentInstanceId())

}
func InitDirectoryCache() {
	db.GetDomains(directory, cache.GetCurrentInstanceId())
	for _, domain := range directory.Domains.Props() {
		domain.NewParams()
		domain.NewVars()
		domain.NewUsers()
		domain.NewGroups()
		db.GetDomainParams(domain, directory)
		db.GetDomainVars(domain, directory)
		db.GetUser(domain, directory)
		for _, user := range domain.Users.Props() {
			user.NewUserParams()
			user.NewUserVars()
			user.NewUserGateways()
			db.GetUserParams(user, directory)
			db.GetUserVars(user, directory)
			db.GetUserGateways(user, directory)
			for _, gateway := range user.Gateways.Props() {
				gateway.NewGatewayParams()
				gateway.NewGatewayVars()
				db.GetUserGatewaysParams(gateway, directory)
				db.GetUserGatewaysVars(gateway, directory)
			}
		}
		db.GetDomainGroups(domain, directory)
		for _, group := range domain.Groups.Props() {
			group.NewGroupUsers()
			db.GetGroupMembers(group, directory)
		}
	}
}

func InitConfigurationCache() {
	db.GetConfigs(configs, cache.GetCurrentInstanceId()) //get conf from db
	if configs.Acl != nil {
		db.GetConfigAclLists(configs)
		for _, list := range configs.Acl.Lists.Props() {
			db.GetConfigAclListNodes(list, configs)
		}
	}
	if configs.Callcenter != nil {
		db.GetConfigCallcenterSettings(configs)
		db.GetConfigCallcenterQueues(configs)
		for _, queue := range configs.Callcenter.Queues.Props() {
			db.GetConfigCallcenterQueuesParams(queue, configs)
		}
		// agents tires dynamic statuses
		db.GetConfigCallcenterAgents(configs, directory)
		db.GetConfigCallcenterTiers(configs)
	}
	if configs.Sofia != nil {
		db.GetConfigSofiaSettings(configs)
		db.GetConfigSofiaProfiles(configs)
		for _, profile := range configs.Sofia.Profiles.Props() {
			db.GetConfigSofiaProfileAliases(profile, configs)
			db.GetConfigSofiaProfileDomains(profile, configs)
			db.GetConfigSofiaProfileParams(profile, configs)
			db.GetConfigSofiaProfileGateways(profile, configs)
			for _, gateway := range profile.Gateways.Props() {
				db.GetConfigSofiaProfileGatewayParams(gateway, configs)
				db.GetConfigSofiaProfileGatewayVars(gateway, configs)
			}
		}
	}
	if configs.CdrPgCsv != nil {
		db.GetConfigCdrPgCsvSettings(configs)
		db.GetConfigCdrPgCsvSchema(configs)

	}
	if configs.Verto != nil {
		db.GetConfigVertoSettings(configs)
		db.GetConfigVertoProfiles(configs)
		for _, profile := range configs.Verto.Profiles.Props() {
			db.GetConfigVertoProfileParams(profile, configs)
		}
	}
	if configs.OdbcCdr != nil {
		db.GetConfigOdbcCdrSettings(configs)
		db.GetConfigOdbcCdrTables(configs)
		for _, table := range configs.OdbcCdr.Tables.Props() {
			db.GetConfigOdbcCdrTableFields(table, configs)
		}
	}
	if configs.Lcr != nil {
		db.GetConfigLcrSettings(configs)
		db.GetConfigLcrProfiles(configs)
		for _, profile := range configs.Lcr.Profiles.Props() {
			db.GetConfigLcrProfileParams(profile, configs)
		}
	}
	if configs.Shout != nil {
		db.GetConfigShoutSettings(configs)
	}
	if configs.Redis != nil {
		db.GetConfigRedisSettings(configs)
	}
	if configs.Nibblebill != nil {
		db.GetConfigNibblebillSettings(configs)
	}
	if configs.Db != nil {
		db.GetConfigDbSettings(configs)
	}
	if configs.Memcache != nil {
		db.GetConfigMemcacheSettings(configs)
	}
	if configs.Avmd != nil {
		db.GetConfigAvmdSettings(configs)
	}
	if configs.TtsCommandline != nil {
		db.GetConfigTtsCommandlineSettings(configs)
	}
	if configs.CdrMongodb != nil {
		db.GetConfigCdrMongodbSettings(configs)
	}
	if configs.HttpCache != nil {
		db.GetConfigHttpCacheSettings(configs)
	}
	if configs.Opus != nil {
		db.GetConfigOpusSettings(configs)
	}
	if configs.Python != nil {
		db.GetConfigPythonSettings(configs)
	}
	if configs.Alsa != nil {
		db.GetConfigAlsaSettings(configs)
	}
	if configs.Amr != nil {
		db.GetConfigAmrSettings(configs)
	}
	if configs.Amrwb != nil {
		db.GetConfigAmrwbSettings(configs)
	}
	if configs.Cepstral != nil {
		db.GetConfigCepstralSettings(configs)
	}
	if configs.Cidlookup != nil {
		db.GetConfigCidlookupSettings(configs)
	}
	if configs.Curl != nil {
		db.GetConfigCurlSettings(configs)
	}
	if configs.DialplanDirectory != nil {
		db.GetConfigDialplanDirectorySettings(configs)
	}
	if configs.Easyroute != nil {
		db.GetConfigEasyrouteSettings(configs)
	}
	if configs.ErlangEvent != nil {
		db.GetConfigErlangEventSettings(configs)
	}
	if configs.EventMulticast != nil {
		db.GetConfigEventMulticastSettings(configs)
	}
	if configs.Fax != nil {
		db.GetConfigFaxSettings(configs)
	}
	if configs.Lua != nil {
		db.GetConfigLuaSettings(configs)
	}
	if configs.Mongo != nil {
		db.GetConfigMongoSettings(configs)
	}
	if configs.Msrp != nil {
		db.GetConfigMsrpSettings(configs)
	}
	if configs.Oreka != nil {
		db.GetConfigOrekaSettings(configs)
	}
	if configs.Perl != nil {
		db.GetConfigPerlSettings(configs)
	}
	if configs.Pocketsphinx != nil {
		db.GetConfigPocketsphinxSettings(configs)
	}
	if configs.SangomaCodec != nil {
		db.GetConfigSangomaCodecSettings(configs)
	}
	if configs.Sndfile != nil {
		db.GetConfigSndfileSettings(configs)
	}
	if configs.XmlCdr != nil {
		db.GetConfigXmlCdrSettings(configs)
	}
	if configs.XmlRpc != nil {
		db.GetConfigXmlRpcSettings(configs)
	}
	if configs.Zeroconf != nil {
		db.GetConfigZeroconfSettings(configs)
	}
	if configs.PostSwitch != nil {
		db.GetConfigPostSwitchSettings(configs)
		db.GetConfigPostSwitchDefaultPtimes(configs)
		db.GetConfigPostSwitchCliKeybindings(configs)
	}
	if configs.Distributor != nil {
		db.GetConfigDistributorLists(configs)
		for _, list := range configs.Distributor.Lists.Props() {
			db.GetConfigDistributorListNodes(list, configs)
		}
	}
	if configs.Directory != nil {
		db.GetConfigDirectorySettings(configs)
		db.GetConfigDirectoryProfiles(configs)
		for _, profile := range configs.Directory.Profiles.Props() {
			db.GetConfigDirectoryProfileParams(profile, configs)
		}
	}
	if configs.Fifo != nil {
		db.GetConfigFifoSettings(configs)
		db.GetConfigFifoFifos(configs)
		for _, fifo := range configs.Fifo.Fifos.Props() {
			db.GetConfigFifoFifoParams(fifo, configs)
		}
	}
	if configs.Opal != nil {
		db.GetConfigOpalSettings(configs)
		db.GetConfigOpalListeners(configs)
		for _, listener := range configs.Opal.Listeners.Props() {
			db.GetConfigOpalListenerParams(listener, configs)
		}
	}
	if configs.Osp != nil {
		db.GetConfigOspSettings(configs)
		db.GetConfigOspProfiles(configs)
		for _, profile := range configs.Osp.Profiles.Props() {
			db.GetConfigOspProfileParams(profile, configs)
		}
	}
	if configs.Unicall != nil {
		db.GetConfigUnicallSettings(configs)
		db.GetConfigUnicallSpans(configs)
		for _, span := range configs.Unicall.Spans.Props() {
			db.GetConfigUnicallSpanParams(span, configs)
		}
	}
	if configs.Conference != nil {
		GetConfCache(configs.Conference.Id, configs.Conference.Advertise)
		GetConfCache(configs.Conference.Id, configs.Conference.CallerControlsGroups)
		for _, item := range configs.Conference.CallerControlsGroups.Props() {
			GetConfCache(item.Id, item.Controls)
		}
		GetConfCache(configs.Conference.Id, configs.Conference.Profiles)
		for _, item := range configs.Conference.Profiles.Props() {
			GetConfCache(item.Id, item.Params)
		}
		GetConfCache(configs.Conference.Id, configs.Conference.ChatPermissions)
		for _, item := range configs.Conference.ChatPermissions.Props() {
			GetConfCache(item.Id, item.Users)
		}
	}
	if configs.ConferenceLayouts != nil {
		GetConfCache(configs.ConferenceLayouts.Id, configs.ConferenceLayouts.ConferenceLayoutsGroups)
		for _, item := range configs.ConferenceLayouts.ConferenceLayoutsGroups.Props() {
			GetConfCache(item.Id, item.Layouts)
		}
		GetConfCache(configs.ConferenceLayouts.Id, configs.ConferenceLayouts.Layouts)
		for _, item := range configs.ConferenceLayouts.Layouts.Props() {
			GetConfCache(item.Id, item.LayoutsImages)
		}
	}
	if configs.PostLoadModules != nil {
		GetConfCache(configs.PostLoadModules.Id, configs.PostLoadModules.Modules)
	} else {
		//always exist no need to import from fs(for now)
		SetConfPostLoadModules()
	}
	if configs.Voicemail != nil {
		GetConfCache(configs.Voicemail.Id, configs.Voicemail.Settings)
		GetConfCache(configs.Voicemail.Id, configs.Voicemail.Profiles)
		for _, item := range configs.Voicemail.Profiles.Props() {
			GetConfCache(item.Id, item.Params)
		}
	}
}

func InitDialplanCache() {
	db.GetDialplanSettings(dialplan, cache.GetCurrentInstanceId())
	db.GetContexts(dialplan, cache.GetCurrentInstanceId())
	for _, context := range dialplan.Contexts.Props() {
		context.Extensions = mainStruct.NewExtensions()
		db.GetContextExtensions(context, dialplan)
		for _, extension := range context.Extensions.Props() {
			extension.Conditions = mainStruct.NewConditions()
			db.GetExtensionConditions(extension, dialplan)
			for _, condition := range extension.Conditions.Props() {
				condition.Regexes = mainStruct.NewRegexes()
				condition.Actions = mainStruct.NewActions()
				condition.AntiActions = mainStruct.NewAntiActions()
				db.GetConditionRegexes(condition, dialplan)
				db.GetConditionActions(condition, dialplan)
				db.GetConditionAntiActions(condition, dialplan)
			}
		}
		context.CacheFullXML()
	}
}

func GetChannelsCache() *mainStruct.Channels {
	return channels
}

func GetChannelsCounter() (int, int) {
	return channels.Total, channels.Answered
}

func ReloadCallcenter() {
	if configs.Callcenter == nil {
		return
	}
	//configs.Callcenter.Members = mainStruct.NewMembers()

	configs.Callcenter.Settings = mainStruct.NewParams()
	db.GetConfigCallcenterSettings(configs)

	configs.Callcenter.Queues = mainStruct.NewQueues()
	configs.Callcenter.QueueParams = mainStruct.NewQueueParams()
	db.GetConfigCallcenterQueues(configs)
	for _, queue := range configs.Callcenter.Queues.Props() {
		db.GetConfigCallcenterQueuesParams(queue, configs)
	}
	// agents tires dynamic statuses
	configs.Callcenter.Agents = mainStruct.NewAgents()
	db.GetConfigCallcenterAgents(configs, directory)
	configs.Callcenter.Tiers = mainStruct.NewTiers()
	db.GetConfigCallcenterTiers(configs)
}
