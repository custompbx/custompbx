package pbxcache

import (
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/db"
	"custompbx/mainStruct"
	"errors"
	"log"
	"regexp"
	"time"
)

var ConfigurationCache altStruct.Configurations

func GetConfigs(s []interface{}) *altStruct.Configurations {
	conf := &altStruct.Configurations{}
	for _, c := range s {
		config, ok := c.(altStruct.ConfigurationsList)
		if !ok {
			continue
		}
		subCached := ConfigurationCache.GetConfiguration(config.Name)
		if subCached != nil {
			config.Loaded = subCached.Loaded
		}
		conf.FillConfigurations(&config)
	}
	return conf
}

func GetCachedConfigSection(confName string, post map[string]string) *mainStruct.Configuration {
	switch confName {
	case mainStruct.ConfPostLoadSwitch:
		return configs.XMLPostSwitch()
	case mainStruct.ConfAcl:
		return configs.XMLAcl()
	case mainStruct.ConfCallcenter:
		return configs.XMLCallcenter(post["CC-Queue"])
	case mainStruct.ConfCdrPgCsv:
		return configs.XMLCdrPgCsv()
	case mainStruct.ConfOdbcCdr:
		return configs.XMLOdbcCdr()
	case mainStruct.ConfSofia:
		return configs.XMLSofia(post["profile"])
	case mainStruct.ConfVerto:
		return configs.XMLVerto()
	case mainStruct.ConfLcr:
		return configs.XMLLcr()
	case mainStruct.ConfShout:
		return configs.XMLShout()
	case mainStruct.ConfRedis:
		return configs.XMLRedis()
	case mainStruct.ConfNibblebill:
		return configs.XMLNibblebill()
	case mainStruct.ConfDb:
		return configs.XMLDb()
	case mainStruct.ConfMemcache:
		return configs.XMLMemcache()
	case mainStruct.ConfAvmd:
		return configs.XMLAvmd()
	case mainStruct.ConfTtsCommandline:
		return configs.XMLTtsCommandline()
	case mainStruct.ConfCdrMongodb:
		return configs.XMLCdrMongodb()
	case mainStruct.ConfHttpCache:
		return configs.XMLHttpCache()
	case mainStruct.ConfOpus:
		return configs.XMLOpus()
	case mainStruct.ConfPython:
		return configs.XMLPython()
	case mainStruct.ConfAlsa:
		return configs.XMLAlsa()
	case mainStruct.ConfAmr:
		return configs.XMLAmr()
	case mainStruct.ConfAmrwb:
		return configs.XMLAmrwb()
	case mainStruct.ConfCepstral:
		return configs.XMLCepstral()
	case mainStruct.ConfCidlookup:
		return configs.XMLCidlookup()
	case mainStruct.ConfCurl:
		return configs.XMLCurl()
	case mainStruct.ConfDialplanDirectory:
		return configs.XMLDialplanDirectory()
	case mainStruct.ConfEasyroute:
		return configs.XMLEasyroute()
	case mainStruct.ConfErlangEvent:
		return configs.XMLErlangEvent()
	case mainStruct.ConfEventMulticast:
		return configs.XMLEventMulticast()
	case mainStruct.ConfFax:
		return configs.XMLFax()
	case mainStruct.ConfLua:
		return configs.XMLLua()
	case mainStruct.ConfMongo:
		return configs.XMLMongo()
	case mainStruct.ConfMsrp:
		return configs.XMLMsrp()
	case mainStruct.ConfOreka:
		return configs.XMLOreka()
	case mainStruct.ConfPerl:
		return configs.XMLPerl()
	case mainStruct.ConfPocketsphinx:
		return configs.XMLPocketsphinx()
	case mainStruct.ConfSangomaCodec:
		return configs.XMLSangomaCodec()
	case mainStruct.ConfSndfile:
		return configs.XMLSndfile()
	case mainStruct.ConfXmlCdr:
		return configs.XMLXmlCdr()
	case mainStruct.ConfXmlRpc:
		return configs.XMLXmlRpc()
	case mainStruct.ConfZeroconf:
		return configs.XMLZeroconf()
	case mainStruct.ConfDirectory:
		return configs.XMLDirectory()
	case mainStruct.ConfFifo:
		return configs.XMLFifo()
	case mainStruct.ConfOpal:
		return configs.XMLOpal()
	case mainStruct.ConfOsp:
		return configs.XMLOsp()
	case mainStruct.ConfUnicall:
		return configs.XMLUnicall()
	case mainStruct.ConfConference:
		return configs.XMLConference()
	case mainStruct.ConfConferenceLayouts:
		return configs.XMLConferenceLayouts()
	case mainStruct.ConfPostLoadModules:
		return configs.XMLPostLoadModules()
	case mainStruct.ConfVoicemail:
		return configs.XMLVoicemail()
	default:
		return nil
	}

	return nil
}

func SetConfAcl() (int64, error) {
	if configs.Acl != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfAcl, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewAcl(res, true)
	return res, nil
}

func SetConfAclList(listNname, listDefault string) (*mainStruct.List, error) {
	if configs.Acl == nil {
		return &mainStruct.List{}, errors.New("no config")
	}
	res, err := db.SetConfAclList(configs.Acl.Id, listNname, listDefault)
	if err != nil {
		return &mainStruct.List{}, err
	}

	list := &mainStruct.List{Id: res, Name: listNname, Default: listDefault, Nodes: mainStruct.NewNodes(), Enabled: true}
	configs.Acl.Lists.Set(list)
	configs.Acl.Nodes = mainStruct.NewNodes()
	return list, err
}

func SetConfAclNode(list *mainStruct.List, nodeType, cidr, domain string) (int64, error) {
	if configs.Acl == nil {
		return 0, errors.New("no config")
	}

	if list == nil {
		return 0, errors.New("list name doesn't exists")
	}
	res, position, err := db.SetConfAclListNode(list.Id, nodeType, cidr, domain)
	if err != nil {
		return 0, err
	}

	node := &mainStruct.Node{Id: res, Type: nodeType, Cidr: cidr, Domain: domain, List: list, Enabled: true, Position: position}
	configs.Acl.Nodes.Set(node)
	list.Nodes.Set(node)
	return res, err
}

func MoveAclListNode(node *mainStruct.Node, newPosition int64) error {
	if node == nil || newPosition == 0 {
		return errors.New("node doesn't exists")
	}

	err := db.MoveAclListNode(node, newPosition)
	if err != nil {
		return err
	}

	return err
}

func SetConfCallcenter() (int64, error) {
	if configs.Callcenter != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfCallcenter, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewCallcenter(res, true)
	return res, nil
}

func SetConfCallcenterSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCallcenterSetting(configs.Callcenter.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Callcenter.Settings.Set(param)
	return param, err
}

func SetConfCallcenterQueue(queueName string) (*mainStruct.Queue, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCallcenterQueue(configs.Callcenter.Id, queueName)
	if err != nil {
		return nil, err
	}

	queue := &mainStruct.Queue{Id: res, Name: queueName, Params: mainStruct.NewQueueParams(), Enabled: true}
	configs.Callcenter.Queues.Set(queue)
	return queue, err
}

func SetConfCallcenterQueueParam(queue *mainStruct.Queue, paramName, paramValue string) (*mainStruct.QueueParam, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}
	if queue == nil {
		return nil, errors.New("queue name doesn't exists")
	}
	res, err := db.SetConfCallcenterQueueParam(queue.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}
	param := &mainStruct.QueueParam{Id: res, Name: paramName, Value: paramValue, Enabled: true, Queue: queue}
	queue.Params.Set(param)
	configs.Callcenter.QueueParams.Set(param)
	return param, err
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
	readyTime int64) (*mainStruct.Agent, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}

	res, err := db.SetConfCallcenterAgent(
		name,
		agentType,
		system,
		instanceId,
		uuid,
		contact,
		status,
		state,
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
		readyTime)
	if err != nil {
		return nil, err
	}
	agent := &mainStruct.Agent{Id: res,
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
		ReadyTime:         readyTime}

	configs.Callcenter.Agents.Set(agent)

	if directory != nil {
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
	}

	return agent, nil
}

func SetConfCallcenterTier(queue, agent, state string, position, level int64) (*mainStruct.Tier, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}

	res, err := db.SetConfCallcenterTier(queue, agent, state, position, level)
	if err != nil {
		return nil, err
	}
	tier := &mainStruct.Tier{Id: res, Queue: queue, Agent: agent, Position: position, Level: level, State: state}
	configs.Callcenter.Tiers.Set(tier)
	return tier, nil
}

func SetConfCallcenterMember(
	uuid, state, queue, instanceId string, abandonedEpoch, baseScore,
	bridgeEpoch int64, cidName, cidNumber string, joinedEpoch, rejoinedEpoch int64,
	servingAgent, servingSystem, sessionUuid string, skillScore, systemEpoch int64,
) (*mainStruct.Member, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}

	err := db.SetConfCallcenterMember(uuid, state, queue, instanceId, abandonedEpoch, baseScore,
		bridgeEpoch, cidName, cidNumber, joinedEpoch, rejoinedEpoch,
		servingAgent, servingSystem, sessionUuid, skillScore, systemEpoch)
	if err != nil {
		return nil, err
	}
	member := &mainStruct.Member{Uuid: uuid, State: state, Queue: queue, InstanceId: instanceId, AbandonedEpoch: abandonedEpoch, BaseScore: baseScore,
		BridgeEpoch: bridgeEpoch, CidName: cidName, CidNumber: cidNumber, JoinedEpoch: joinedEpoch, RejoinedEpoch: rejoinedEpoch,
		ServingAgent: servingAgent, ServingSystem: servingSystem, SessionUuid: sessionUuid, SkillScore: skillScore, SystemEpoch: systemEpoch}
	configs.Callcenter.Members.Set(member)
	return member, nil
}

func SetConfCallcenterMemberCache(
	uuid, state, queue, instanceId string, abandonedEpoch, baseScore,
	bridgeEpoch int64, cidName, cidNumber string, joinedEpoch, rejoinedEpoch int64,
	servingAgent, servingSystem, sessionUuid string, skillScore, systemEpoch int64,
) (*mainStruct.Member, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}

	member := &mainStruct.Member{Uuid: uuid, State: state, Queue: queue, InstanceId: instanceId, AbandonedEpoch: abandonedEpoch, BaseScore: baseScore,
		BridgeEpoch: bridgeEpoch, CidName: cidName, CidNumber: cidNumber, JoinedEpoch: joinedEpoch, RejoinedEpoch: rejoinedEpoch,
		ServingAgent: servingAgent, ServingSystem: servingSystem, SessionUuid: sessionUuid, SkillScore: skillScore, SystemEpoch: systemEpoch}
	configs.Callcenter.Members.Set(member)
	return member, nil
}

func SetConfigSofiaGlobalSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfSofiaGlobalSetting(configs.Sofia.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Sofia.GlobalSettings.Set(param)
	return param, err
}

func SetConfSofia() (int64, error) {
	if configs.Sofia != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfSofia, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewSofia(res, true)
	return res, nil
}

func SetConfigSofiaProfile(profileName string) (*mainStruct.SofiaProfile, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfSofiaProfile(configs.Sofia.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.SofiaProfile{
		Id: res, Name: profileName,
		Params:   mainStruct.NewSofiaProfileParams(),
		Aliases:  mainStruct.NewAliases(),
		Gateways: mainStruct.NewSofiaGateways(),
		Domains:  mainStruct.NewSofiaDomains(),
		Enabled:  true,
	}
	configs.Sofia.Profiles.Set(profile)
	return profile, err
}

func SetConfigSofiaProfileAliases(profile *mainStruct.SofiaProfile, aliasName string) (*mainStruct.Alias, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}

	res, err := db.SetConfSofiaProfileAlias(profile.Id, aliasName)
	if err != nil {
		return nil, err
	}

	alias := &mainStruct.Alias{Id: res, Name: aliasName, Enabled: true, Profile: profile}
	profile.Aliases.Set(alias)
	configs.Sofia.ProfileAliases.Set(alias)
	return alias, err
}

func SetConfigSofiaGateway(profile *mainStruct.SofiaProfile, gatewayName string) (*mainStruct.SofiaGateway, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}

	res, err := db.SetConfSofiaProfileGateway(profile.Id, gatewayName)
	if err != nil {
		return nil, err
	}

	gateway := &mainStruct.SofiaGateway{Id: res, Name: gatewayName, Params: mainStruct.NewSofiaGatewayParams(), Vars: mainStruct.NewSofiaGatewayVars(), Enabled: true, Profile: profile}
	profile.Gateways.Set(gateway)
	configs.Sofia.ProfileGateways.Set(gateway)
	return gateway, err
}

func SetConfigSofiaGatewayParam(gateway *mainStruct.SofiaGateway, name, value string) (*mainStruct.SofiaGatewayParam, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	if gateway == nil {
		return nil, errors.New("no gateway")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, err := db.SetConfSofiaProfileGatewayParam(gateway.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.SofiaGatewayParam{Id: res, Name: name, Value: value, Enabled: true, Gateway: gateway}
	gateway.Params.Set(param)
	configs.Sofia.GatewayParams.Set(param)
	return param, err
}

func SetConfigSofiaGatewayVar(gateway *mainStruct.SofiaGateway, name, value, direction string) (*mainStruct.SofiaGatewayVariable, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	if gateway == nil {
		return nil, errors.New("no gateway")
	}
	if name == "" {
		return nil, errors.New("no variable")
	}

	res, err := db.SetConfSofiaProfileGatewayVar(gateway.Id, name, value, direction)
	if err != nil {
		return nil, err
	}

	variable := &mainStruct.SofiaGatewayVariable{Id: res, Name: name, Value: value, Direction: direction, Enabled: true, Gateway: gateway}
	gateway.Vars.Set(variable)
	configs.Sofia.GatewayVars.Set(variable)
	return variable, err
}

func SetConfigSofiaProfileDomain(profile *mainStruct.SofiaProfile, domainName, alias, parse string) (*mainStruct.SofiaDomain, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}

	var aliasValue bool
	var parseValue bool
	if alias != "false" {
		aliasValue = true
	}
	if parse != "false" {
		parseValue = true
	}

	res, err := db.SetConfSofiaProfileDomain(profile.Id, domainName, aliasValue, parseValue)
	if err != nil {
		return nil, err
	}

	domain := &mainStruct.SofiaDomain{Id: res, Name: domainName, Alias: aliasValue, Parse: parseValue, Enabled: true, Profile: profile}
	profile.Domains.Set(domain)
	configs.Sofia.ProfileDomains.Set(domain)
	return domain, err
}

func SetConfigSofiaProfileParam(profile *mainStruct.SofiaProfile, name, value string) (int64, error) {
	if configs.Sofia == nil {
		return 0, errors.New("no config")
	}
	if profile == nil {
		return 0, errors.New("no profile")
	}
	if name == "" {
		return 0, errors.New("no param")
	}

	res, err := db.SetConfSofiaProfileParam(profile.Id, name, value)
	if err != nil {
		return 0, err
	}

	param := &mainStruct.SofiaProfileParam{Id: res, Name: name, Value: value, Enabled: true, Profile: profile}
	profile.Params.Set(param)
	configs.Sofia.ProfileParams.Set(param)
	return res, err
}

func GetAclLists() (map[int64]*mainStruct.List, bool) {
	if configs.Acl == nil {
		return map[int64]*mainStruct.List{}, false
	}

	item := configs.Acl.Lists.GetList()
	return item, true
}

func GetAclList(id int64) *mainStruct.List {
	if configs.Acl == nil {
		return nil
	}
	item := configs.Acl.Lists.GetById(id)
	return item
}

func GetSofiaProfile(id int64) *mainStruct.SofiaProfile {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.Profiles.GetById(id)
	return item
}

func IsSofiaExists() bool {
	return configs.Sofia != nil
}

func GetSofiaProfileByName(name string) *mainStruct.SofiaProfile {
	if configs.Sofia == nil || configs.Sofia.Profiles == nil {
		return nil
	}
	item := configs.Sofia.Profiles.GetByName(name)
	return item
}

func GetSofiaProfileGateway(id int64) *mainStruct.SofiaGateway {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.ProfileGateways.GetById(id)
	return item
}

func GetSofiaProfileGatewayByName(name string) *mainStruct.SofiaGateway {
	if configs.Sofia == nil || configs.Sofia.ProfileGateways == nil {
		return nil
	}
	item := configs.Sofia.ProfileGateways.GetByName(name)
	return item
}

func IsVertoExists() bool {
	return configs.Verto != nil
}

func IsOdbcCdrExists() bool {
	return configs.OdbcCdr != nil
}

func IsCallcenterExists() bool {
	return configs.Callcenter != nil
}

func IsCdrPgCsvExists() bool {
	return configs.CdrPgCsv != nil
}

func IsAclExists() bool {
	return configs.Acl != nil
}

func UpdateAclListDefault(id int64, newValue string) (int64, error) {
	list := configs.Acl.Lists.GetById(id)
	if list == nil {
		return 0, errors.New("list doesn't exists")
	}
	res, err := db.UpdateAclListDefault(list.Id, newValue)
	if err != nil {
		return 0, err
	}
	list.Default = newValue
	return res, err
}

func DelAclNode(node *mainStruct.Node) int64 {
	parentId := node.List.Id
	ok := db.DelAclNode(node.Id)
	if !ok {
		return 0
	}

	node.List.Nodes.Remove(node)
	configs.Acl.Nodes.Remove(node)
	return parentId
}

func GetAclNode(id int64) *mainStruct.Node {
	if configs.Acl == nil {
		return nil
	}
	item := configs.Acl.Nodes.GetById(id)
	return item
}

func UpdateAclNode(node *mainStruct.Node, nodeType, cidr, domain string) error {
	if node == nil {
		return errors.New("node doesn't exists")
	}
	if cidr != "" {
		domain = ""
	}
	_, err := db.UpdateAclNode(node.Id, nodeType, cidr, domain)
	if err != nil {
		return err
	}
	node.Type = nodeType
	node.Cidr = cidr
	node.Domain = domain
	return nil
}

func SwitchAclNode(node *mainStruct.Node, switcher bool) error {
	if node == nil {
		return errors.New("node doesn't exists")
	}
	_, err := db.SwitchAclNode(node.Id, switcher)
	if err != nil {
		return err
	}
	node.Enabled = switcher
	return nil
}

func DelAclList(id int64) bool {
	list := configs.Acl.Lists.GetById(id)
	if list == nil {
		return false
	}
	ok := db.DelAclList(list.Id)
	if !ok {
		return false
	}

	configs.Acl.Lists.Remove(list)
	configs.Acl.Nodes.ClearUp(configs)
	return true
}

func UpdateAclList(id int64, newName string) error {
	domain := configs.Acl.Lists.GetById(id)
	if domain == nil {
		return errors.New("domain name doesn't exists")
	}
	err := db.UpdateAclList(id, newName)
	if err != nil {
		return err
	}
	configs.Acl.Lists.Rename(domain.Name, newName)
	return err
}

func GetSofiaGlobalSettings() map[int64]*mainStruct.Param {
	if configs.Sofia == nil {
		return map[int64]*mainStruct.Param{}
	}

	item := configs.Sofia.GlobalSettings.GetList()
	return item
}

func GetSofiaProfilesLists() (map[int64]*mainStruct.SofiaProfile, bool) {
	if configs.Sofia == nil {
		return map[int64]*mainStruct.SofiaProfile{}, false
	}

	item := configs.Sofia.Profiles.GetList()
	return item, true
}

func GetSofiaProfilesProps() ([]*mainStruct.SofiaProfile, bool) {
	if configs.Sofia == nil {
		return []*mainStruct.SofiaProfile{}, false
	}

	item := configs.Sofia.Profiles.Props()
	return item, true
}

func GetSofiaGlobalSetting(id int64) *mainStruct.Param {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.GlobalSettings.GetById(id)
	return item
}

func UpdateSofiaGlobalSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateSofiaGlobalSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchSofiaGlobalSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchSofiaGlobalSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelSofiaGlobalSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelSofiaGlobalSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Sofia.GlobalSettings.Remove(param)
	return id
}

func SetConfSofiaProfileParam(profile *mainStruct.SofiaProfile, name, value string) (*mainStruct.SofiaProfileParam, error) {
	if configs.Sofia == nil {
		return nil, errors.New("no config")
	}

	if profile == nil {
		return nil, errors.New("profile name doesn't exists")
	}
	res, err := db.SetConfSofiaProfileParam(profile.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.SofiaProfileParam{Id: res, Name: name, Value: value, Enabled: true, Profile: profile}
	configs.Sofia.ProfileParams.Set(param)
	profile.Params.Set(param)
	return param, err
}

func GetSofiaProfileParam(id int64) *mainStruct.SofiaProfileParam {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.ProfileParams.GetById(id)
	return item
}

func GetSofiaProfileGatewayParam(id int64) *mainStruct.SofiaGatewayParam {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.GatewayParams.GetById(id)
	return item
}

func GetSofiaProfileGatewayVariable(id int64) *mainStruct.SofiaGatewayVariable {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.GatewayVars.GetById(id)
	return item
}

func DelProfileParam(param *mainStruct.SofiaProfileParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Profile.Id
	ok := db.DelProfileParam(param.Id)
	if !ok {
		return 0
	}

	param.Profile.Params.Remove(param)
	configs.Sofia.ProfileParams.Remove(param)
	return parentId
}

func SwitchProfileParam(param *mainStruct.SofiaProfileParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchProfileParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Profile.Id, err
}

func UpdateProfileParam(param *mainStruct.SofiaProfileParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateProfileParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Profile.Id, err
}

func GetSofiasGateways() map[int64]*mainStruct.SofiaGateway {
	if configs.Sofia == nil {
		return map[int64]*mainStruct.SofiaGateway{}
	}

	item := configs.Sofia.ProfileGateways.GetList()
	return item
}

func GetSofiasGatewaysProps() []*mainStruct.SofiaGateway {
	if configs.Sofia == nil {
		return []*mainStruct.SofiaGateway{}
	}

	item := configs.Sofia.ProfileGateways.Props()
	return item
}

func GetSofiaParentsGateways() map[int64]map[int64]*mainStruct.SofiaGateway {
	if configs.Sofia == nil {
		return map[int64]map[int64]*mainStruct.SofiaGateway{}
	}

	item := configs.Sofia.ProfileGateways.GetParentList()
	return item
}

func UpdateProfileGatewayParam(param *mainStruct.SofiaGatewayParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateProfileGatewayParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Gateway.Id, err
}

func SwitchProfileGatewayParam(param *mainStruct.SofiaGatewayParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchProfileGatewayParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Gateway.Id, err
}

func DelProfileGatewayParam(param *mainStruct.SofiaGatewayParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Gateway.Id
	ok := db.DelProfileGatewayParam(param.Id)
	if !ok {
		return 0
	}

	param.Gateway.Params.Remove(param)
	configs.Sofia.GatewayParams.Remove(param)
	return parentId
}

func UpdateProfileGatewayVariable(variable *mainStruct.SofiaGatewayVariable, name, value, direction string) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	_, err := db.UpdateProfileGatewayVariable(variable.Id, name, value, direction)
	if err != nil {
		return 0, err
	}
	variable.Name = name
	variable.Value = value
	variable.Direction = direction
	return variable.Gateway.Id, err
}

func SwitchProfileGatewayVariable(variable *mainStruct.SofiaGatewayVariable, switcher bool) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	_, err := db.SwitchProfileGatewayVariable(variable.Id, switcher)
	if err != nil {
		return 0, err
	}
	variable.Enabled = switcher
	return variable.Gateway.Id, err
}

func DelProfileGatewayVariable(variable *mainStruct.SofiaGatewayVariable) int64 {
	if variable == nil {
		return 0
	}
	parentId := variable.Gateway.Id
	ok := db.DelProfileGatewayVariable(variable.Id)
	if !ok {
		return 0
	}

	variable.Gateway.Vars.Remove(variable)
	configs.Sofia.GatewayVars.Remove(variable)
	return parentId
}

func GetSofiaParentsDomains() map[int64]map[int64]*mainStruct.SofiaDomain {
	if configs.Sofia == nil {
		return map[int64]map[int64]*mainStruct.SofiaDomain{}
	}

	item := configs.Sofia.ProfileDomains.GetParentList()
	return item
}

func GetSofiaProfileDomain(id int64) *mainStruct.SofiaDomain {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.ProfileDomains.GetById(id)
	return item
}

func DelProfileDomain(domain *mainStruct.SofiaDomain) int64 {
	if domain == nil {
		return 0
	}
	parentId := domain.Profile.Id
	ok := db.DelProfileDomain(domain.Id)
	if !ok {
		return 0
	}

	domain.Profile.Domains.Remove(domain)
	configs.Sofia.ProfileDomains.Remove(domain)
	return parentId
}

func SwitchProfileDomain(domain *mainStruct.SofiaDomain, switcher bool) (int64, error) {
	if domain == nil {
		return 0, errors.New("domain doesn't exists")
	}
	_, err := db.SwitchProfileDomain(domain.Id, switcher)
	if err != nil {
		return 0, err
	}
	domain.Enabled = switcher
	return domain.Profile.Id, err
}

func UpdateProfileDomain(domain *mainStruct.SofiaDomain, name string, alias, parse bool) (int64, error) {
	if domain == nil {
		return 0, errors.New("domain doesn't exists")
	}
	_, err := db.UpdateProfileDomain(domain.Id, name, alias, parse)
	if err != nil {
		return 0, err
	}
	domain.Name = name
	domain.Alias = alias
	domain.Parse = parse
	return domain.Profile.Id, err
}

func GetSofiaProfileAlias(id int64) *mainStruct.Alias {
	if configs.Sofia == nil {
		return nil
	}
	item := configs.Sofia.ProfileAliases.GetById(id)
	return item
}

func DelProfileAlias(alias *mainStruct.Alias) int64 {
	if alias == nil {
		return 0
	}
	parentId := alias.Profile.Id
	ok := db.DelProfileAlias(alias.Id)
	if !ok {
		return 0
	}

	alias.Profile.Aliases.Remove(alias)
	configs.Sofia.ProfileAliases.Remove(alias)
	return parentId
}

func SwitchProfileAlias(alias *mainStruct.Alias, switcher bool) (int64, error) {
	if alias == nil {
		return 0, errors.New("domain doesn't exists")
	}
	_, err := db.SwitchProfileAlias(alias.Id, switcher)
	if err != nil {
		return 0, err
	}
	alias.Enabled = switcher
	return alias.Profile.Id, err
}

func SwitchProfile(profile *mainStruct.SofiaProfile, switcher bool) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.SwitchProfile(profile.Id, switcher)
	if err != nil {
		return 0, err
	}
	profile.Enabled = switcher
	return profile.Id, err
}

func UpdateProfileAlias(alias *mainStruct.Alias, name string) (int64, error) {
	if alias == nil {
		return 0, errors.New("domain doesn't exists")
	}
	_, err := db.UpdateProfileAlias(alias.Id, name)
	if err != nil {
		return 0, err
	}
	alias.Name = name

	return alias.Profile.Id, err
}

func UpdateSofiaProfile(profile *mainStruct.SofiaProfile, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateSofiaProfile(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelProfile(profile *mainStruct.SofiaProfile) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelProfile(profile.Id)
	if !ok {
		return 0
	}

	configs.Sofia.Profiles.Remove(profile)
	configs.Sofia.ClearSofiaProfile()
	return profile.Id
}

func UpdateSofiaProfileGateway(gateway *mainStruct.SofiaGateway, name string) (int64, error) {
	if gateway == nil {
		return 0, errors.New("gateway doesn't exists")
	}
	_, err := db.UpdateSofiaProfileGateway(gateway.Id, name)
	if err != nil {
		return 0, err
	}
	gateway.Name = name
	return gateway.Id, err
}

func DelProfileGateway(gateway *mainStruct.SofiaGateway) int64 {
	if gateway == nil {
		return 0
	}
	ok := db.DelProfileGateway(gateway.Id)
	if !ok {
		return 0
	}

	gateway.Profile.Gateways.Remove(gateway)
	configs.Sofia.ProfileGateways.Remove(gateway)
	configs.Sofia.ClearSofiaProfileGateways()
	return gateway.Id
}

func GetModules() *mainStruct.Configurations {
	return configs
}

func GetModule(id int64) mainStruct.Module {
	switch true {
	case configs.Acl != nil && configs.Acl.Id == id:
		return configs.Acl
	case configs.Callcenter != nil && configs.Callcenter.Id == id:
		return configs.Callcenter
	case configs.Sofia != nil && configs.Sofia.Id == id:
		return configs.Sofia
	case configs.CdrPgCsv != nil && configs.CdrPgCsv.Id == id:
		return configs.CdrPgCsv
	case configs.Verto != nil && configs.Verto.Id == id:
		return configs.Verto
	case configs.OdbcCdr != nil && configs.OdbcCdr.Id == id:
		return configs.OdbcCdr
	case configs.Lcr != nil && configs.Lcr.Id == id:
		return configs.Lcr
	case configs.Shout != nil && configs.Shout.Id == id:
		return configs.Shout
	case configs.Redis != nil && configs.Redis.Id == id:
		return configs.Redis
	case configs.Nibblebill != nil && configs.Nibblebill.Id == id:
		return configs.Nibblebill
	case configs.Db != nil && configs.Db.Id == id:
		return configs.Db
	case configs.Memcache != nil && configs.Memcache.Id == id:
		return configs.Memcache
	case configs.Avmd != nil && configs.Avmd.Id == id:
		return configs.Avmd
	case configs.TtsCommandline != nil && configs.TtsCommandline.Id == id:
		return configs.TtsCommandline
	case configs.CdrMongodb != nil && configs.CdrMongodb.Id == id:
		return configs.CdrMongodb
	case configs.HttpCache != nil && configs.HttpCache.Id == id:
		return configs.HttpCache
	case configs.Opus != nil && configs.Opus.Id == id:
		return configs.Opus
	case configs.Python != nil && configs.Python.Id == id:
		return configs.Python
	case configs.Alsa != nil && configs.Alsa.Id == id:
		return configs.Alsa
	case configs.Amr != nil && configs.Amr.Id == id:
		return configs.Amr
	case configs.Amrwb != nil && configs.Amrwb.Id == id:
		return configs.Amrwb
	case configs.Cepstral != nil && configs.Cepstral.Id == id:
		return configs.Cepstral
	case configs.Cidlookup != nil && configs.Cidlookup.Id == id:
		return configs.Cidlookup
	case configs.Curl != nil && configs.Curl.Id == id:
		return configs.Curl
	case configs.DialplanDirectory != nil && configs.DialplanDirectory.Id == id:
		return configs.DialplanDirectory
	case configs.Easyroute != nil && configs.Easyroute.Id == id:
		return configs.Easyroute
	case configs.ErlangEvent != nil && configs.ErlangEvent.Id == id:
		return configs.ErlangEvent
	case configs.EventMulticast != nil && configs.EventMulticast.Id == id:
		return configs.EventMulticast
	case configs.Fax != nil && configs.Fax.Id == id:
		return configs.Fax
	case configs.Lua != nil && configs.Lua.Id == id:
		return configs.Lua
	case configs.Mongo != nil && configs.Mongo.Id == id:
		return configs.Mongo
	case configs.Msrp != nil && configs.Msrp.Id == id:
		return configs.Msrp
	case configs.Oreka != nil && configs.Oreka.Id == id:
		return configs.Oreka
	case configs.Perl != nil && configs.Perl.Id == id:
		return configs.Perl
	case configs.Pocketsphinx != nil && configs.Pocketsphinx.Id == id:
		return configs.Pocketsphinx
	case configs.SangomaCodec != nil && configs.SangomaCodec.Id == id:
		return configs.SangomaCodec
	case configs.Sndfile != nil && configs.Sndfile.Id == id:
		return configs.Sndfile
	case configs.XmlCdr != nil && configs.XmlCdr.Id == id:
		return configs.XmlCdr
	case configs.XmlRpc != nil && configs.XmlRpc.Id == id:
		return configs.XmlRpc
	case configs.Zeroconf != nil && configs.Zeroconf.Id == id:
		return configs.Zeroconf
	case configs.PostSwitch != nil && configs.PostSwitch.Id == id:
		return configs.PostSwitch
	case configs.Distributor != nil && configs.Distributor.Id == id:
		return configs.Distributor
	case configs.Opal != nil && configs.Opal.Id == id:
		return configs.Opal
	case configs.Unicall != nil && configs.Unicall.Id == id:
		return configs.Unicall
	case configs.Directory != nil && configs.Directory.Id == id:
		return configs.Directory
	case configs.Fifo != nil && configs.Fifo.Id == id:
		return configs.Fifo
	case configs.Osp != nil && configs.Osp.Id == id:
		return configs.Osp
	case configs.Conference != nil && configs.Conference.Id == id:
		return configs.Conference
	case configs.PostLoadModules != nil && configs.PostLoadModules.Id == id:
		return configs.PostLoadModules
	case configs.Voicemail != nil && configs.Voicemail.Id == id:
		return configs.Voicemail
	default:
		return nil
	}
}

func GetModuleByName(name string) mainStruct.Module {
	switch name {
	case mainStruct.ModCdrPgCsv:
		return configs.CdrPgCsv
	case mainStruct.ModSofiaAlias:
		return configs.Sofia
	case mainStruct.ModSofia:
		return configs.Sofia
	case mainStruct.ModAcl:
		return configs.Acl
	case mainStruct.ModVerto:
		return configs.Verto
	case mainStruct.ModVertoAlias:
		return configs.Verto
	case mainStruct.ModCallcenter:
		return configs.Callcenter
	case mainStruct.ModCallcenterAlias:
		return configs.Callcenter
	case mainStruct.ModOdbcCdr:
		return configs.OdbcCdr
	case mainStruct.ModLcrAlias:
		return configs.Lcr
	case mainStruct.ModLcr:
		return configs.Lcr
	case mainStruct.ModShout:
		return configs.Shout
	case mainStruct.ModRedis:
		return configs.Redis
	case mainStruct.ModNibblebill:
		return configs.Nibblebill
	case mainStruct.ModDb:
		return configs.Db
	case mainStruct.ModMemcache:
		return configs.Memcache
	case mainStruct.ModAvmd:
		return configs.Avmd
	case mainStruct.ModTtsCommandline:
		return configs.TtsCommandline
	case mainStruct.ModCdrMongodb:
		return configs.CdrMongodb
	case mainStruct.ModHttpCache:
		return configs.HttpCache
	case mainStruct.ModOpus:
		return configs.Opus
	case mainStruct.ModPython:
		return configs.Python
	case mainStruct.ModAlsa:
		return configs.Alsa
	case mainStruct.ModAmr:
		return configs.Amr
	case mainStruct.ModAmrwb:
		return configs.Amrwb
	case mainStruct.ModCepstral:
		return configs.Cepstral
	case mainStruct.ModCidlookup:
		return configs.Cidlookup
	case mainStruct.ModCurl:
		return configs.Curl
	case mainStruct.ModDialplanDirectory:
		return configs.DialplanDirectory
	case mainStruct.ModEasyroute:
		return configs.Easyroute
	case mainStruct.ModErlangEvent:
		return configs.ErlangEvent
	case mainStruct.ModEventMulticast:
		return configs.EventMulticast
	case mainStruct.ModFax:
		return configs.Fax
	case mainStruct.ModLua:
		return configs.Lua
	case mainStruct.ModMongo:
		return configs.Mongo
	case mainStruct.ModMsrp:
		return configs.Msrp
	case mainStruct.ModOreka:
		return configs.Oreka
	case mainStruct.ModPerl:
		return configs.Perl
	case mainStruct.ModPocketsphinx:
		return configs.Pocketsphinx
	case mainStruct.ModSangomaCodec:
		return configs.SangomaCodec
	case mainStruct.ModSndfile:
		return configs.Sndfile
	case mainStruct.ModXmlCdr:
		return configs.XmlCdr
	case mainStruct.ModXmlRpc:
		return configs.XmlRpc
	case mainStruct.ModZeroconf:
		return configs.Zeroconf
	case mainStruct.ModDistributor:
		return configs.Distributor
	case mainStruct.ModOpal:
		return configs.Opal
	case mainStruct.ModUnicall:
		return configs.Unicall
	case mainStruct.ModDirectory:
		return configs.Directory
	case mainStruct.ModFifo:
		return configs.Fifo
	case mainStruct.ModOsp:
		return configs.Osp
	case mainStruct.ModConference:
		return configs.Conference
	case mainStruct.ModPostLoadModules:
		return configs.PostLoadModules
	case mainStruct.ModVoicemail:
		return configs.Voicemail
	default:
		return nil
	}
}

func SwitchModule(module mainStruct.Module, switcher bool) (int64, error) {
	if module == nil {
		return 0, errors.New("module doesn't exists")
	}
	_, err := db.SwitchModule(module.GetId(), switcher)
	if err != nil {
		return 0, err
	}
	module.Switch(switcher)
	return module.GetId(), err
}

func TruncateModuleConfig(module mainStruct.Module) error {
	if module == nil {
		return errors.New("module doesn't exists")
	}
	err := db.TruncateModuleConfig(module.GetId())
	if err != nil {
		return err
	}
	configs.TruncateModuleConfig(module.GetModuleName())
	return nil
}

func SetConfCdrPgCsv() (int64, error) {
	if configs.CdrPgCsv != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfCdrPgCsv, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewCdrPgCsv(res, true)
	return res, nil
}

func SetConfCdrPgCsvSetting(name, value string) (*mainStruct.Param, error) {
	if configs.CdrPgCsv == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCdrPgCsvSetting(configs.CdrPgCsv.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: name, Value: value, Enabled: true}
	configs.CdrPgCsv.Settings.Set(param)
	return param, err
}

func SetConfCdrPgCsvSchemaField(variable, colunm string) (*mainStruct.Field, error) {
	if configs.CdrPgCsv == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCdrPgCsvSchema(configs.CdrPgCsv.Id, variable, colunm)
	if err != nil {
		return nil, err
	}

	field := &mainStruct.Field{Id: res, Var: variable, Column: colunm, Enabled: true}
	configs.CdrPgCsv.Schema.Set(field)
	return field, err
}

func GetCdrPgCsvSettings() (map[int64]*mainStruct.Param, bool) {
	if configs.CdrPgCsv == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.CdrPgCsv.Settings.GetList()
	return item, true
}

func GetCdrPgCsvSchema() (map[int64]*mainStruct.Field, bool) {
	if configs.CdrPgCsv == nil {
		return map[int64]*mainStruct.Field{}, false
	}

	item := configs.CdrPgCsv.Schema.GetList()
	return item, true
}

func SetConfCdrPgCsvField(variable, column string) (*mainStruct.Field, error) {
	if configs.CdrPgCsv == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCdrPgCsvSchema(configs.CdrPgCsv.Id, variable, column)
	if err != nil {
		return nil, err
	}

	field := &mainStruct.Field{Id: res, Var: variable, Column: column, Enabled: true}
	configs.CdrPgCsv.Schema.Set(field)
	return field, err
}

func GetCdrPgCsvParam(id int64) *mainStruct.Param {
	if configs.CdrPgCsv == nil {
		return nil
	}
	item := configs.CdrPgCsv.Settings.GetById(id)
	return item
}

func SwitchCdrPgCsvParam(param *mainStruct.Param, switcher bool) error {
	if param == nil {
		return errors.New("param doesn't exists")
	}
	err := db.SwitchCdrPgCsvParam(param.Id, switcher)
	if err != nil {
		return err
	}
	param.Enabled = switcher
	return err
}

func GetCdrPgCsvField(id int64) *mainStruct.Field {
	if configs.CdrPgCsv == nil {
		return nil
	}
	item := configs.CdrPgCsv.Schema.GetById(id)
	return item
}

func SwitchCdrPgCsvField(field *mainStruct.Field, switcher bool) error {
	if field == nil {
		return errors.New("field doesn't exists")
	}
	err := db.SwitchCdrPgCsvField(field.Id, switcher)
	if err != nil {
		return err
	}
	field.Enabled = switcher
	return err
}

func UpdateCdrPgCsvParam(param *mainStruct.Param, name, value string) error {
	if param == nil {
		return errors.New("param doesn't exists")
	}
	err := db.UpdateCdrPgCsvParam(param.Id, name, value)
	if err != nil {
		return err
	}
	param.Name = name
	param.Value = value
	return err
}

func UpdateCdrPgCsvField(field *mainStruct.Field, variable, column string) error {
	if field == nil {
		return errors.New("field doesn't exists")
	}
	err := db.UpdateCdrPgCsvField(field.Id, variable, column)
	if err != nil {
		return err
	}
	field.Var = variable
	field.Column = column
	return err
}

func DelCdrPgCsvParam(param *mainStruct.Param) int64 {
	id := param.Id
	err := db.DelCdrPgCsvParam(param.Id)
	if err != nil {
		return 0
	}

	configs.CdrPgCsv.Settings.Remove(param)
	return id
}

func DelCdrPgCsvField(field *mainStruct.Field) int64 {
	id := field.Id
	err := db.DelCdrPgCsvField(field.Id)
	if err != nil {
		return 0
	}

	configs.CdrPgCsv.Schema.Remove(field)
	return id
}

func GetCdrPgCsvConnectorData(name string) string {
	if configs.CdrPgCsv == nil {
		return ""
	}

	item := configs.CdrPgCsv.Settings.GetByName(name)
	if item == nil || !item.Enabled {
		return ""
	}

	return item.Value
}

func SetConfVerto() (int64, error) {
	if configs.Verto != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfVerto, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewVerto(res, true)
	return res, nil
}

func SetConfigVertoSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Verto == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfVertoSetting(configs.Verto.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Verto.Settings.Set(param)
	return param, err
}

func SetConfigVertoProfile(profileName string) (*mainStruct.VertoProfile, error) {
	if configs.Verto == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfVertoProfile(configs.Verto.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.VertoProfile{
		Id: res, Name: profileName,
		Params:  mainStruct.NewVertoProfileParams(),
		Enabled: true,
	}
	configs.Verto.Profiles.Set(profile)
	return profile, err
}

func SetConfigVertoProfileParam(profile *mainStruct.VertoProfile, name, value, secure string) (*mainStruct.VertoProfileParam, error) {
	if configs.Verto == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, position, err := db.SetConfVertoProfileParam(profile.Id, name, value, secure)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.VertoProfileParam{Id: res, Name: name, Value: value, Secure: secure, Enabled: true, Profile: profile, Position: position}
	profile.Params.Set(param)
	configs.Verto.ProfileParams.Set(param)
	return param, err
}

func MoveVertoProfileParam(param *mainStruct.VertoProfileParam, newPosition int64) error {
	if param == nil || newPosition == 0 {
		return errors.New("node doesn't exists")
	}

	err := db.MoveVertoProfileParam(param, newPosition)
	if err != nil {
		return err
	}

	return err
}

func GetConfigVerto() (map[int64]*mainStruct.Param, bool) {
	if configs.Verto == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Verto.Settings.GetList()
	return item, true
}

func GetProfiles() (map[int64]*mainStruct.VertoProfile, bool) {
	if configs.Verto == nil {
		return map[int64]*mainStruct.VertoProfile{}, false
	}

	item := configs.Verto.Profiles.GetList()
	return item, true
}

func GetVertoProfile(id int64) *mainStruct.VertoProfile {
	if configs.Verto == nil {
		return nil
	}
	item := configs.Verto.Profiles.GetById(id)
	return item
}

func GetVertoProfileByName(name string) *mainStruct.VertoProfile {
	if configs.Verto == nil {
		return nil
	}
	item := configs.Verto.Profiles.GetByName(name)
	return item
}

func GetVertoSetting(id int64) *mainStruct.Param {
	if configs.Verto == nil {
		return nil
	}
	item := configs.Verto.Settings.GetById(id)
	return item
}

func GetVertoSettingByName(name string) *mainStruct.Param {
	if configs.Verto == nil {
		return nil
	}
	item := configs.Verto.Settings.GetByName(name)
	return item
}

func UpdateVertoSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateVertoSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchVertoSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchVertoSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelVertoSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelVertoSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Verto.Settings.Remove(param)
	return id
}

func GetVertoProfileParam(id int64) *mainStruct.VertoProfileParam {
	if configs.Verto == nil {
		return nil
	}
	item := configs.Verto.ProfileParams.GetById(id)
	return item
}

func DelVertoProfileParam(param *mainStruct.VertoProfileParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Profile.Id
	ok := db.DelVertoProfileParam(param.Id)
	if !ok {
		return 0
	}

	param.Profile.Params.Remove(param)
	configs.Verto.ProfileParams.Remove(param)
	return parentId
}

func SwitchVertoProfileParam(param *mainStruct.VertoProfileParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchVertoProfileParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Profile.Id, err
}

func UpdateVertoProfileParam(param *mainStruct.VertoProfileParam, name, value, secure string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateVertoProfileParam(param.Id, name, value, secure)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Profile.Id, err
}

func UpdateVertoProfile(profile *mainStruct.VertoProfile, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateVertoProfile(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelVertoProfile(profile *mainStruct.VertoProfile) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelVertoProfile(profile.Id)
	if !ok {
		return 0
	}

	configs.Verto.Profiles.Remove(profile)
	configs.Verto.ProfileParams.ClearUp(configs.Verto)
	return profile.Id
}

func GetCallcenterQueuesLists() (map[int64]*mainStruct.Queue, bool) {
	if configs.Callcenter == nil {
		return nil, false
	}

	item := configs.Callcenter.Queues.GetList()
	return item, true
}

func GetCallcenterQueue(id int64) *mainStruct.Queue {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Queues.GetById(id)
	return item
}

func GetCallcenterQueueByName(name string) *mainStruct.Queue {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Queues.GetByName(name)
	return item
}

func GetCallcenterSetting(id int64) *mainStruct.Param {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Settings.GetById(id)
	return item
}

func UpdateCallcenterSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateCallcenterSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchCallcenterSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchCallcenterSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelCallcenterSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelCallcenterSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Callcenter.Settings.Remove(param)
	return id
}

func GetCallcenterQueueParam(id int64) *mainStruct.QueueParam {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.QueueParams.GetById(id)
	return item
}

func DelCallcenterQueueParam(param *mainStruct.QueueParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Queue.Id
	ok := db.DelCallcenterQueueParam(param.Id)
	if !ok {
		return 0
	}

	param.Queue.Params.Remove(param)
	configs.Callcenter.QueueParams.Remove(param)
	return parentId
}

func SwitchCallcenterQueueParam(param *mainStruct.QueueParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchCallcenterQueueParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Queue.Id, err
}

func UpdateCallcenterQueueParam(param *mainStruct.QueueParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateCallcenterQueueParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Queue.Id, err
}

func UpdateCallcenterQueue(queue *mainStruct.Queue, name string) (int64, error) {
	if queue == nil {
		return 0, errors.New("queue doesn't exists")
	}
	_, err := db.UpdateCallcenterQueue(queue.Id, name)
	if err != nil {
		return 0, err
	}
	queue.Name = name
	return queue.Id, err
}

func DelCallcenterQueue(queue *mainStruct.Queue) int64 {
	if queue == nil {
		return 0
	}
	ok := db.DelCallcenterQueue(queue.Id)
	if !ok {
		return 0
	}

	configs.Callcenter.Queues.Remove(queue)
	configs.Callcenter.QueueParams.ClearUp(configs.Callcenter)
	return queue.Id
}

func GetCallcenterSettings() map[int64]*mainStruct.Param {
	if configs.Callcenter == nil {
		return map[int64]*mainStruct.Param{}
	}

	item := configs.Callcenter.Settings.GetList()
	return item
}

func GetCallcenterAgentsList() (map[int64]*mainStruct.Agent, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}

	item := configs.Callcenter.Agents.GetList()
	return item, nil
}

func GetCallcenterAgentsListsByForm(limit, offset int, filters []mainStruct.Filter, order mainStruct.Order) ([]*mainStruct.Agent, int, error) {
	if configs.Callcenter == nil {
		return nil, 0, errors.New("no config")
	}

	item, total := configs.Callcenter.Agents.FilteredProps(limit, offset, filters, order)
	return item, total, nil
}

func GetCallcenterAgent(id int64) *mainStruct.Agent {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Agents.GetById(id)
	return item
}

func GetCallcenterAgentByName(name string) *mainStruct.Agent {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Agents.GetByName(name)
	return item
}

func UpdateCallcenterAgent(agent *mainStruct.Agent, name, value string, eventChannel chan interface{}) (bool, error) {
	if name == "id" {
		return false, errors.New("please dont")
	}
	if agent == nil {
		return false, errors.New("no agent")
	}
	realValue, err := db.UpdateCallcenterTableColumn("agents", agent.Id, name, value)
	if err != nil {
		return false, err
	}
	ok := agent.Update(name, realValue)
	switch name {
	case "name":
		configs.Callcenter.Agents.Rename(agent.Name, realValue)
	case "state":
		eventChannel <- &map[int64]*mainStruct.Agent{agent.Id: agent}
	case "status":
		agent.LastStatusChange = time.Now().Unix()
		eventChannel <- &map[int64]*mainStruct.Agent{agent.Id: agent}
	}

	return ok, nil
}

func DelCallcenterAgent(agent *mainStruct.Agent) int64 {
	if agent == nil {
		return 0
	}
	ok := db.DelCallcenterAgent(agent.Id)
	if !ok {
		return 0
	}

	configs.Callcenter.Agents.Remove(agent)
	return agent.Id
}

func GetCallcenterTiersList() (map[int64]*mainStruct.Tier, error) {
	if configs.Callcenter == nil {
		return nil, errors.New("no config")
	}

	item := configs.Callcenter.Tiers.GetList()
	return item, nil
}

func GetCallcenterTiersListsByForm(limit, offset int, filters []mainStruct.Filter, order mainStruct.Order) ([]*mainStruct.Tier, int, error) {
	if configs.Callcenter == nil {
		return nil, 0, errors.New("no config")
	}

	item, total := configs.Callcenter.Tiers.FilteredProps(limit, offset, filters, order)
	return item, total, nil
}

func GetCallcenterTier(id int64) *mainStruct.Tier {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Tiers.GetById(id)
	return item
}

func GetCallcenterTierByName(name string) *mainStruct.Tier {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Tiers.GetByName(name)
	return item
}

func UpdateCallcenterTier(tier *mainStruct.Tier, name, value string) (bool, error) {
	if name == "id" {
		return false, errors.New("please dont")
	}
	if tier == nil {
		return false, errors.New("no tier")
	}
	realValue, err := db.UpdateCallcenterTableColumn("tiers", tier.Id, name, value)
	if err != nil {
		return false, err
	}
	ok := tier.Update(name, realValue)
	if name == "queue" {
		configs.Callcenter.Tiers.Rename(tier.Queue, realValue, tier.Agent, tier.Agent)
	} else if name == "agent" {
		configs.Callcenter.Tiers.Rename(tier.Queue, tier.Queue, realValue, tier.Agent)
	}
	return ok, nil
}

func DelCallcenterTier(tier *mainStruct.Tier) int64 {
	if tier == nil {
		return 0
	}
	ok := db.DelCallcenterTier(tier.Id)
	if !ok {
		return 0
	}

	configs.Callcenter.Tiers.Remove(tier)
	return tier.Id
}

func GetCallcenterMembersListsByForm(limit, offset int, filters []mainStruct.Filter, order mainStruct.Order) ([]*mainStruct.Member, int, error) {
	if configs.Callcenter == nil {
		return nil, 0, errors.New("no config")
	}

	item, total := configs.Callcenter.Members.FilteredProps(limit, offset, filters, order)
	return item, total, nil
}

func GetCallcenterMember(uuid string) *mainStruct.Member {
	if configs.Callcenter == nil {
		return nil
	}
	item := configs.Callcenter.Members.GetByUuid(uuid)
	return item
}

func DelCallcenterMember(member *mainStruct.Member) string {
	if member == nil {
		return ""
	}
	ok := db.DelCallcenterMember(member.Uuid)
	if !ok {
		return ""
	}

	configs.Callcenter.Members.Remove(member)
	return member.Uuid
}

func DelCallcenterMemberCache(member *mainStruct.Member) string {
	if member == nil {
		return ""
	}
	configs.Callcenter.Members.Remove(member)
	return member.Uuid
}

func SetConfOdbcCdr() (int64, error) {
	if configs.OdbcCdr != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfOdbcCdr, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewOdbcCdr(res, true)
	return res, nil
}

func GetOdbcCdrParam(id int64) *mainStruct.Param {
	if configs.OdbcCdr == nil {
		return nil
	}
	item := configs.OdbcCdr.Settings.GetById(id)
	return item
}

func GetOdbcCdrSettings() (map[int64]*mainStruct.Param, bool) {
	if configs.OdbcCdr == nil {
		return nil, false
	}

	item := configs.OdbcCdr.Settings.GetList()
	return item, true
}

func GetOdbcCdrTables() map[int64]*mainStruct.Table {
	if configs.OdbcCdr == nil {
		return nil
	}

	item := configs.OdbcCdr.Tables.GetList()
	return item
}

func GetOdbcCdrTable(id int64) *mainStruct.Table {
	if configs.OdbcCdr == nil {
		return nil
	}

	return configs.OdbcCdr.Tables.GetById(id)
}

func GetOdbcCdrTableByName(name string) *mainStruct.Table {
	if configs.OdbcCdr == nil {
		return nil
	}

	return configs.OdbcCdr.Tables.GetByName(name)
}

func GetOdbcCdrFields() map[int64]map[int64]*mainStruct.ODBCField {
	if configs.OdbcCdr == nil {
		return nil
	}

	item := configs.OdbcCdr.TableFields.GetParentList()
	return item
}

func GetOdbcCdrField(id int64) *mainStruct.ODBCField {
	if configs.OdbcCdr == nil {
		return nil
	}
	item := configs.OdbcCdr.TableFields.GetById(id)
	return item
}

func SetConfOdbcCdrSetting(name, value string) (*mainStruct.Param, error) {
	if configs.OdbcCdr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOdbcCdrSetting(configs.OdbcCdr.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: name, Value: value, Enabled: true}
	configs.OdbcCdr.Settings.Set(param)
	return param, err
}

func UpdateOdbcCdrParam(param *mainStruct.Param, name, value string) error {
	if param == nil {
		return errors.New("param doesn't exists")
	}
	err := db.UpdateOdbcCdrParam(param.Id, name, value)
	if err != nil {
		return err
	}
	param.Name = name
	param.Value = value
	return err
}

func SwitchOdbcCdrParam(param *mainStruct.Param, switcher bool) error {
	if param == nil {
		return errors.New("param doesn't exists")
	}
	err := db.SwitchOdbcCdrParam(param.Id, switcher)
	if err != nil {
		return err
	}
	param.Enabled = switcher
	return err
}

func DelOdbcCdrParam(param *mainStruct.Param) int64 {
	id := param.Id
	err := db.DelOdbcCdrParam(param.Id)
	if err != nil {
		return 0
	}

	configs.OdbcCdr.Settings.Remove(param)
	return id
}

func SetConfOdbcCdrTable(name, logLeg string) (*mainStruct.Table, error) {
	if configs.OdbcCdr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOdbcCdrTable(configs.OdbcCdr.Id, name, logLeg)
	if err != nil {
		return nil, err
	}

	table := &mainStruct.Table{Id: res, Name: name, LogLeg: logLeg, Fields: mainStruct.NewOdbcFields(), Enabled: true}
	configs.OdbcCdr.Tables.Set(table)
	return table, err
}

func UpdateOdbcCdrTable(table *mainStruct.Table, name, logLeg string) error {
	if table == nil {
		return errors.New("table doesn't exists")
	}
	err := db.UpdateOdbcCdrTable(table.Id, name, logLeg)
	if err != nil {
		return err
	}
	table.Name = name
	table.LogLeg = logLeg
	return err
}

func SwitchOdbcCdrTable(table *mainStruct.Table, switcher bool) error {
	if table == nil {
		return errors.New("table doesn't exists")
	}
	err := db.SwitchOdbcCdrTable(table.Id, switcher)
	if err != nil {
		return err
	}
	table.Enabled = switcher
	return err
}

func DelOdbcCdrTable(table *mainStruct.Table) int64 {
	id := table.Id
	err := db.DelOdbcCdrTable(table.Id)
	if err != nil {
		return 0
	}

	configs.OdbcCdr.Tables.Remove(table)
	configs.OdbcCdr.TableFields.ClearUp(configs.OdbcCdr)
	return id
}

func SetConfOdbcCdrTableField(table *mainStruct.Table, name, chanVarName string) (*mainStruct.ODBCField, error) {
	if configs.OdbcCdr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOdbcCdrTableField(table.Id, name, chanVarName)
	if err != nil {
		return nil, err
	}

	field := &mainStruct.ODBCField{Id: res, Name: name, ChanVarName: chanVarName, Table: table, Enabled: true}
	table.Fields.Set(field)
	configs.OdbcCdr.TableFields.Set(field)
	return field, err
}

func SwitchOdbcCdrField(field *mainStruct.ODBCField, switcher bool) error {
	if field == nil {
		return errors.New("field doesn't exists")
	}
	err := db.SwitchOdbcCdrField(field.Id, switcher)
	if err != nil {
		return err
	}
	field.Enabled = switcher
	return err
}

func UpdateOdbcCdrField(field *mainStruct.ODBCField, name, chanVarName string) error {
	if field == nil {
		return errors.New("field doesn't exists")
	}
	err := db.UpdateOdbcCdrField(field.Id, name, chanVarName)
	if err != nil {
		return err
	}
	field.Name = name
	field.ChanVarName = chanVarName
	return err
}

func DelOdbcCdrField(field *mainStruct.ODBCField) int64 {
	id := field.Id
	err := db.DelOdbcCdrField(field.Id)
	if err != nil {
		return 0
	}

	field.Table.Fields.Remove(field)
	configs.OdbcCdr.TableFields.Remove(field)
	return id
}

func GetOdbcCdrConnectorData(name string) string {
	if configs.OdbcCdr == nil {
		return ""
	}

	item := configs.OdbcCdr.Settings.GetByName(name)
	if item == nil || !item.Enabled {
		return ""
	}

	return item.Value
}

func SetConfLcr() (int64, error) {
	if configs.Lcr != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfLcr, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewLcr(res, true)
	return res, nil
}

func SetConfigLcrSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Lcr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfLcrSetting(configs.Lcr.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Lcr.Settings.Set(param)
	return param, err
}

func SetConfigLcrProfile(profileName string) (*mainStruct.LcrProfile, error) {
	if configs.Lcr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfLcrProfile(configs.Lcr.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.LcrProfile{
		Id: res, Name: profileName,
		Params:  mainStruct.NewLcrProfileParams(),
		Enabled: true,
	}
	configs.Lcr.Profiles.Set(profile)
	return profile, err
}

func SetConfigLcrProfileParam(profile *mainStruct.LcrProfile, name, value string) (*mainStruct.LcrProfileParam, error) {
	if configs.Lcr == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, err := db.SetConfLcrProfileParam(profile.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.LcrProfileParam{Id: res, Name: name, Value: value, Enabled: true, Profile: profile}
	profile.Params.Set(param)
	configs.Lcr.ProfileParams.Set(param)
	return param, err
}

func GetConfigLcr() (map[int64]*mainStruct.Param, bool) {
	if configs.Lcr == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Lcr.Settings.GetList()
	return item, true
}

func GetLcrProfiles() (map[int64]*mainStruct.LcrProfile, bool) {
	if configs.Lcr == nil {
		return map[int64]*mainStruct.LcrProfile{}, false
	}

	item := configs.Lcr.Profiles.GetList()
	return item, true
}

func GetLcrProfile(id int64) *mainStruct.LcrProfile {
	if configs.Lcr == nil {
		return nil
	}
	item := configs.Lcr.Profiles.GetById(id)
	return item
}

func GetLcrProfileByName(name string) *mainStruct.LcrProfile {
	if configs.Lcr == nil {
		return nil
	}
	item := configs.Lcr.Profiles.GetByName(name)
	return item
}

func GetLcrSetting(id int64) *mainStruct.Param {
	if configs.Lcr == nil {
		return nil
	}
	item := configs.Lcr.Settings.GetById(id)
	return item
}

func GetLcrSettingByName(name string) *mainStruct.Param {
	if configs.Lcr == nil {
		return nil
	}
	item := configs.Lcr.Settings.GetByName(name)
	return item
}

func UpdateLcrSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateLcrSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchLcrSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchLcrSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelLcrSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelLcrSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Lcr.Settings.Remove(param)
	return id
}

func GetLcrProfileParam(id int64) *mainStruct.LcrProfileParam {
	if configs.Lcr == nil {
		return nil
	}
	item := configs.Lcr.ProfileParams.GetById(id)
	return item
}

func DelLcrProfileParam(param *mainStruct.LcrProfileParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Profile.Id
	ok := db.DelLcrProfileParam(param.Id)
	if !ok {
		return 0
	}

	param.Profile.Params.Remove(param)
	configs.Lcr.ProfileParams.Remove(param)
	return parentId
}

func SwitchLcrProfileParam(param *mainStruct.LcrProfileParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchLcrProfileParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Profile.Id, err
}

func UpdateLcrProfileParam(param *mainStruct.LcrProfileParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateLcrProfileParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Profile.Id, err
}

func UpdateLcrProfile(profile *mainStruct.LcrProfile, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateLcrProfile(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelLcrProfile(profile *mainStruct.LcrProfile) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelLcrProfile(profile.Id)
	if !ok {
		return 0
	}

	configs.Lcr.Profiles.Remove(profile)
	configs.Lcr.ProfileParams.ClearUp(configs.Lcr)
	return profile.Id
}

func SetConfShout() (int64, error) {
	if configs.Shout != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfShout, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewShout(res, true)
	return res, nil
}

func SetConfigShoutSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Shout == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfShoutSetting(configs.Shout.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Shout.Settings.Set(param)
	return param, err
}

func GetConfigShout() (map[int64]*mainStruct.Param, bool) {
	if configs.Shout == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Shout.Settings.GetList()
	return item, true
}

func GetShoutSetting(id int64) *mainStruct.Param {
	if configs.Shout == nil {
		return nil
	}
	item := configs.Shout.Settings.GetById(id)
	return item
}

func GetShoutSettingByName(name string) *mainStruct.Param {
	if configs.Shout == nil {
		return nil
	}
	item := configs.Shout.Settings.GetByName(name)
	return item
}

func UpdateShoutSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateShoutSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchShoutSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchShoutSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelShoutSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelShoutSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Shout.Settings.Remove(param)
	return id
}

func SetConfRedis() (int64, error) {
	if configs.Redis != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfRedis, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewRedis(res, true)
	return res, nil
}

func SetConfigRedisSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Redis == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfRedisSetting(configs.Redis.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Redis.Settings.Set(param)
	return param, err
}

func GetConfigRedis() (map[int64]*mainStruct.Param, bool) {
	if configs.Redis == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Redis.Settings.GetList()
	return item, true
}

func GetRedisSetting(id int64) *mainStruct.Param {
	if configs.Redis == nil {
		return nil
	}
	item := configs.Redis.Settings.GetById(id)
	return item
}

func GetRedisSettingByName(name string) *mainStruct.Param {
	if configs.Redis == nil {
		return nil
	}
	item := configs.Redis.Settings.GetByName(name)
	return item
}

func UpdateRedisSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateRedisSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchRedisSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchRedisSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelRedisSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelRedisSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Redis.Settings.Remove(param)
	return id
}

func SetConfNibblebill() (int64, error) {
	if configs.Nibblebill != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfNibblebill, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewNibblebill(res, true)
	return res, nil
}

func SetConfigNibblebillSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Nibblebill == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfNibblebillSetting(configs.Nibblebill.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Nibblebill.Settings.Set(param)
	return param, err
}

func GetConfigNibblebill() (map[int64]*mainStruct.Param, bool) {
	if configs.Nibblebill == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Nibblebill.Settings.GetList()
	return item, true
}

func GetNibblebillSetting(id int64) *mainStruct.Param {
	if configs.Nibblebill == nil {
		return nil
	}
	item := configs.Nibblebill.Settings.GetById(id)
	return item
}

func GetNibblebillSettingByName(name string) *mainStruct.Param {
	if configs.Nibblebill == nil {
		return nil
	}
	item := configs.Nibblebill.Settings.GetByName(name)
	return item
}

func UpdateNibblebillSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateNibblebillSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchNibblebillSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchNibblebillSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelNibblebillSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelNibblebillSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Nibblebill.Settings.Remove(param)
	return id
}

func SetConfDb() (int64, error) {
	if configs.Db != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfDb, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewDb(res, true)
	return res, nil
}

func SetConfigDbSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Db == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfDbSetting(configs.Db.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Db.Settings.Set(param)
	return param, err
}

func GetConfigDb() (map[int64]*mainStruct.Param, bool) {
	if configs.Db == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Db.Settings.GetList()
	return item, true
}

func GetDbSetting(id int64) *mainStruct.Param {
	if configs.Db == nil {
		return nil
	}
	item := configs.Db.Settings.GetById(id)
	return item
}

func GetDbSettingByName(name string) *mainStruct.Param {
	if configs.Db == nil {
		return nil
	}
	item := configs.Db.Settings.GetByName(name)
	return item
}

func UpdateDbSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateDbSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchDbSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchDbSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelDbSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelDbSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Db.Settings.Remove(param)
	return id
}

func SetConfMemcache() (int64, error) {
	if configs.Memcache != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfMemcache, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewMemcache(res, true)
	return res, nil
}

func SetConfigMemcacheSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Memcache == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfMemcacheSetting(configs.Memcache.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Memcache.Settings.Set(param)
	return param, err
}

func GetConfigMemcache() (map[int64]*mainStruct.Param, bool) {
	if configs.Memcache == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Memcache.Settings.GetList()
	return item, true
}

func GetMemcacheSetting(id int64) *mainStruct.Param {
	if configs.Memcache == nil {
		return nil
	}
	item := configs.Memcache.Settings.GetById(id)
	return item
}

func GetMemcacheSettingByName(name string) *mainStruct.Param {
	if configs.Memcache == nil {
		return nil
	}
	item := configs.Memcache.Settings.GetByName(name)
	return item
}

func UpdateMemcacheSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateMemcacheSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchMemcacheSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchMemcacheSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelMemcacheSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelMemcacheSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Memcache.Settings.Remove(param)
	return id
}

func SetConfAvmd() (int64, error) {
	if configs.Avmd != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfAvmd, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewAvmd(res, true)
	return res, nil
}

func SetConfigAvmdSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Avmd == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfAvmdSetting(configs.Avmd.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Avmd.Settings.Set(param)
	return param, err
}

func GetConfigAvmd() (map[int64]*mainStruct.Param, bool) {
	if configs.Avmd == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Avmd.Settings.GetList()
	return item, true
}

func GetAvmdSetting(id int64) *mainStruct.Param {
	if configs.Avmd == nil {
		return nil
	}
	item := configs.Avmd.Settings.GetById(id)
	return item
}

func GetAvmdSettingByName(name string) *mainStruct.Param {
	if configs.Avmd == nil {
		return nil
	}
	item := configs.Avmd.Settings.GetByName(name)
	return item
}

func UpdateAvmdSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateAvmdSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchAvmdSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchAvmdSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelAvmdSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelAvmdSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Avmd.Settings.Remove(param)
	return id
}

func SetConfTtsCommandline() (int64, error) {
	if configs.TtsCommandline != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfTtsCommandline, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewTtsCommandline(res, true)
	return res, nil
}

func SetConfigTtsCommandlineSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.TtsCommandline == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfTtsCommandlineSetting(configs.TtsCommandline.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.TtsCommandline.Settings.Set(param)
	return param, err
}

func GetConfigTtsCommandline() (map[int64]*mainStruct.Param, bool) {
	if configs.TtsCommandline == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.TtsCommandline.Settings.GetList()
	return item, true
}

func GetTtsCommandlineSetting(id int64) *mainStruct.Param {
	if configs.TtsCommandline == nil {
		return nil
	}
	item := configs.TtsCommandline.Settings.GetById(id)
	return item
}

func GetTtsCommandlineSettingByName(name string) *mainStruct.Param {
	if configs.TtsCommandline == nil {
		return nil
	}
	item := configs.TtsCommandline.Settings.GetByName(name)
	return item
}

func UpdateTtsCommandlineSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateTtsCommandlineSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchTtsCommandlineSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchTtsCommandlineSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelTtsCommandlineSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelTtsCommandlineSetting(param.Id)
	if !ok {
		return 0
	}

	configs.TtsCommandline.Settings.Remove(param)
	return id
}

func SetConfCdrMongodb() (int64, error) {
	if configs.CdrMongodb != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfCdrMongodb, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewCdrMongodb(res, true)
	return res, nil
}

func SetConfigCdrMongodbSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.CdrMongodb == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCdrMongodbSetting(configs.CdrMongodb.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.CdrMongodb.Settings.Set(param)
	return param, err
}

func GetConfigCdrMongodb() (map[int64]*mainStruct.Param, bool) {
	if configs.CdrMongodb == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.CdrMongodb.Settings.GetList()
	return item, true
}

func GetCdrMongodbSetting(id int64) *mainStruct.Param {
	if configs.CdrMongodb == nil {
		return nil
	}
	item := configs.CdrMongodb.Settings.GetById(id)
	return item
}

func GetCdrMongodbSettingByName(name string) *mainStruct.Param {
	if configs.CdrMongodb == nil {
		return nil
	}
	item := configs.CdrMongodb.Settings.GetByName(name)
	return item
}

func UpdateCdrMongodbSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateCdrMongodbSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchCdrMongodbSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchCdrMongodbSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelCdrMongodbSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelCdrMongodbSetting(param.Id)
	if !ok {
		return 0
	}

	configs.CdrMongodb.Settings.Remove(param)
	return id
}

func SetConfHttpCache() (int64, error) {
	if configs.HttpCache != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfHttpCache, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewHttpCache(res, true)
	return res, nil
}

func SetConfigHttpCacheSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.HttpCache == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfHttpCacheSetting(configs.HttpCache.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.HttpCache.Settings.Set(param)
	return param, err
}

func GetConfigHttpCache() (map[int64]*mainStruct.Param, bool) {
	if configs.HttpCache == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.HttpCache.Settings.GetList()
	return item, true
}

func GetHttpCacheSetting(id int64) *mainStruct.Param {
	if configs.HttpCache == nil {
		return nil
	}
	item := configs.HttpCache.Settings.GetById(id)
	return item
}

func GetHttpCacheSettingByName(name string) *mainStruct.Param {
	if configs.HttpCache == nil {
		return nil
	}
	item := configs.HttpCache.Settings.GetByName(name)
	return item
}

func UpdateHttpCacheSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateHttpCacheSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchHttpCacheSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchHttpCacheSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelHttpCacheSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelHttpCacheSetting(param.Id)
	if !ok {
		return 0
	}

	configs.HttpCache.Settings.Remove(param)
	return id
}

func SetConfOpus() (int64, error) {
	if configs.Opus != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfOpus, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewOpus(res, true)
	return res, nil
}

func SetConfigOpusSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Opus == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOpusSetting(configs.Opus.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Opus.Settings.Set(param)
	return param, err
}

func GetConfigOpus() (map[int64]*mainStruct.Param, bool) {
	if configs.Opus == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Opus.Settings.GetList()
	return item, true
}

func GetOpusSetting(id int64) *mainStruct.Param {
	if configs.Opus == nil {
		return nil
	}
	item := configs.Opus.Settings.GetById(id)
	return item
}

func GetOpusSettingByName(name string) *mainStruct.Param {
	if configs.Opus == nil {
		return nil
	}
	item := configs.Opus.Settings.GetByName(name)
	return item
}

func UpdateOpusSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateOpusSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchOpusSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchOpusSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelOpusSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelOpusSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Opus.Settings.Remove(param)
	return id
}

func SetConfPython() (int64, error) {
	if configs.Python != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfPython, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewPython(res, true)
	return res, nil
}

func SetConfigPythonSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Python == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfPythonSetting(configs.Python.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Python.Settings.Set(param)
	return param, err
}

func GetConfigPython() (map[int64]*mainStruct.Param, bool) {
	if configs.Python == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Python.Settings.GetList()
	return item, true
}

func GetPythonSetting(id int64) *mainStruct.Param {
	if configs.Python == nil {
		return nil
	}
	item := configs.Python.Settings.GetById(id)
	return item
}

func GetPythonSettingByName(name string) *mainStruct.Param {
	if configs.Python == nil {
		return nil
	}
	item := configs.Python.Settings.GetByName(name)
	return item
}

func UpdatePythonSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdatePythonSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchPythonSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchPythonSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelPythonSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelPythonSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Python.Settings.Remove(param)
	return id
}

func SetConfAlsa() (int64, error) {
	if configs.Alsa != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfAlsa, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewAlsa(res, true)
	return res, nil
}

func SetConfigAlsaSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Alsa == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfAlsaSetting(configs.Alsa.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Alsa.Settings.Set(param)
	return param, err
}

func GetConfigAlsa() (map[int64]*mainStruct.Param, bool) {
	if configs.Alsa == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Alsa.Settings.GetList()
	return item, true
}

func GetAlsaSetting(id int64) *mainStruct.Param {
	if configs.Alsa == nil {
		return nil
	}
	item := configs.Alsa.Settings.GetById(id)
	return item
}

func GetAlsaSettingByName(name string) *mainStruct.Param {
	if configs.Alsa == nil {
		return nil
	}
	item := configs.Alsa.Settings.GetByName(name)
	return item
}

func UpdateAlsaSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateAlsaSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchAlsaSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchAlsaSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelAlsaSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelAlsaSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Alsa.Settings.Remove(param)
	return id
}

func SetConfigfAmr() (int64, error) {
	if configs.Amr != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfAmr, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewAmr(res, true)
	return res, nil
}

func SetConfigAmrSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Amr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfAmrSetting(configs.Amr.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Amr.Settings.Set(param)
	return param, err
}

func GetConfigAmr() (map[int64]*mainStruct.Param, bool) {
	if configs.Amr == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Amr.Settings.GetList()
	return item, true
}

func GetAmrSetting(id int64) *mainStruct.Param {
	if configs.Amr == nil {
		return nil
	}
	item := configs.Amr.Settings.GetById(id)
	return item
}

func GetAmrSettingByName(name string) *mainStruct.Param {
	if configs.Amr == nil {
		return nil
	}
	item := configs.Amr.Settings.GetByName(name)
	return item
}

func UpdateAmrSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateAmrSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchAmrSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchAmrSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelAmrSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelAmrSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Amr.Settings.Remove(param)
	return id
}

func SetConfAmrwb() (int64, error) {
	if configs.Amrwb != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfAmrwb, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewAmrwb(res, true)
	return res, nil
}

func SetConfigAmrwbSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Amrwb == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfAmrwbSetting(configs.Amrwb.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Amrwb.Settings.Set(param)
	return param, err
}

func GetConfigAmrwb() (map[int64]*mainStruct.Param, bool) {
	if configs.Amrwb == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Amrwb.Settings.GetList()
	return item, true
}

func GetAmrwbSetting(id int64) *mainStruct.Param {
	if configs.Amrwb == nil {
		return nil
	}
	item := configs.Amrwb.Settings.GetById(id)
	return item
}

func GetAmrwbSettingByName(name string) *mainStruct.Param {
	if configs.Amrwb == nil {
		return nil
	}
	item := configs.Amrwb.Settings.GetByName(name)
	return item
}

func UpdateAmrwbSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateAmrwbSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchAmrwbSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchAmrwbSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelAmrwbSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelAmrwbSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Amrwb.Settings.Remove(param)
	return id
}

func SetConfCepstral() (int64, error) {
	if configs.Cepstral != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfCepstral, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewCepstral(res, true)
	return res, nil
}

func SetConfigCepstralSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Cepstral == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCepstralSetting(configs.Cepstral.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Cepstral.Settings.Set(param)
	return param, err
}

func GetConfigCepstral() (map[int64]*mainStruct.Param, bool) {
	if configs.Cepstral == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Cepstral.Settings.GetList()
	return item, true
}

func GetCepstralSetting(id int64) *mainStruct.Param {
	if configs.Cepstral == nil {
		return nil
	}
	item := configs.Cepstral.Settings.GetById(id)
	return item
}

func GetCepstralSettingByName(name string) *mainStruct.Param {
	if configs.Cepstral == nil {
		return nil
	}
	item := configs.Cepstral.Settings.GetByName(name)
	return item
}

func UpdateCepstralSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateCepstralSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchCepstralSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchCepstralSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelCepstralSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelCepstralSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Cepstral.Settings.Remove(param)
	return id
}

func SetConfCidlookup() (int64, error) {
	if configs.Cidlookup != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfCidlookup, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewCidlookup(res, true)
	return res, nil
}

func SetConfigCidlookupSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Cidlookup == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCidlookupSetting(configs.Cidlookup.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Cidlookup.Settings.Set(param)
	return param, err
}

func GetConfigCidlookup() (map[int64]*mainStruct.Param, bool) {
	if configs.Cidlookup == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Cidlookup.Settings.GetList()
	return item, true
}

func GetCidlookupSetting(id int64) *mainStruct.Param {
	if configs.Cidlookup == nil {
		return nil
	}
	item := configs.Cidlookup.Settings.GetById(id)
	return item
}

func GetCidlookupSettingByName(name string) *mainStruct.Param {
	if configs.Cidlookup == nil {
		return nil
	}
	item := configs.Cidlookup.Settings.GetByName(name)
	return item
}

func UpdateCidlookupSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateCidlookupSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchCidlookupSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchCidlookupSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelCidlookupSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelCidlookupSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Cidlookup.Settings.Remove(param)
	return id
}

func SetConfCurl() (int64, error) {
	if configs.Curl != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfCurl, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewCurl(res, true)
	return res, nil
}

func SetConfigCurlSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Curl == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfCurlSetting(configs.Curl.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Curl.Settings.Set(param)
	return param, err
}

func GetConfigCurl() (map[int64]*mainStruct.Param, bool) {
	if configs.Curl == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Curl.Settings.GetList()
	return item, true
}

func GetCurlSetting(id int64) *mainStruct.Param {
	if configs.Curl == nil {
		return nil
	}
	item := configs.Curl.Settings.GetById(id)
	return item
}

func GetCurlSettingByName(name string) *mainStruct.Param {
	if configs.Curl == nil {
		return nil
	}
	item := configs.Curl.Settings.GetByName(name)
	return item
}

func UpdateCurlSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateCurlSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchCurlSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchCurlSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelCurlSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelCurlSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Curl.Settings.Remove(param)
	return id
}

func SetConfDialplanDirectory() (int64, error) {
	if configs.DialplanDirectory != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfDialplanDirectory, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewDialplanDirectory(res, true)
	return res, nil
}

func SetConfigDialplanDirectorySetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.DialplanDirectory == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfDialplanDirectorySetting(configs.DialplanDirectory.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.DialplanDirectory.Settings.Set(param)
	return param, err
}

func GetConfigDialplanDirectory() (map[int64]*mainStruct.Param, bool) {
	if configs.DialplanDirectory == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.DialplanDirectory.Settings.GetList()
	return item, true
}

func GetDialplanDirectorySetting(id int64) *mainStruct.Param {
	if configs.DialplanDirectory == nil {
		return nil
	}
	item := configs.DialplanDirectory.Settings.GetById(id)
	return item
}

func GetDialplanDirectorySettingByName(name string) *mainStruct.Param {
	if configs.DialplanDirectory == nil {
		return nil
	}
	item := configs.DialplanDirectory.Settings.GetByName(name)
	return item
}

func UpdateDialplanDirectorySetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateDialplanDirectorySetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchDialplanDirectorySetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchDialplanDirectorySetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelDialplanDirectorySetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelDialplanDirectorySetting(param.Id)
	if !ok {
		return 0
	}

	configs.DialplanDirectory.Settings.Remove(param)
	return id
}

func SetConfEasyroute() (int64, error) {
	if configs.Easyroute != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfEasyroute, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewEasyroute(res, true)
	return res, nil
}

func SetConfigEasyrouteSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Easyroute == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfEasyrouteSetting(configs.Easyroute.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Easyroute.Settings.Set(param)
	return param, err
}

func GetConfigEasyroute() (map[int64]*mainStruct.Param, bool) {
	if configs.Easyroute == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Easyroute.Settings.GetList()
	return item, true
}

func GetEasyrouteSetting(id int64) *mainStruct.Param {
	if configs.Easyroute == nil {
		return nil
	}
	item := configs.Easyroute.Settings.GetById(id)
	return item
}

func GetEasyrouteSettingByName(name string) *mainStruct.Param {
	if configs.Easyroute == nil {
		return nil
	}
	item := configs.Easyroute.Settings.GetByName(name)
	return item
}

func UpdateEasyrouteSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateEasyrouteSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchEasyrouteSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchEasyrouteSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelEasyrouteSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelEasyrouteSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Easyroute.Settings.Remove(param)
	return id
}

func SetConfErlangEvent() (int64, error) {
	if configs.ErlangEvent != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfErlangEvent, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewErlangEvent(res, true)
	return res, nil
}

func SetConfigErlangEventSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.ErlangEvent == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfErlangEventSetting(configs.ErlangEvent.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.ErlangEvent.Settings.Set(param)
	return param, err
}

func GetConfigErlangEvent() (map[int64]*mainStruct.Param, bool) {
	if configs.ErlangEvent == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.ErlangEvent.Settings.GetList()
	return item, true
}

func GetErlangEventSetting(id int64) *mainStruct.Param {
	if configs.ErlangEvent == nil {
		return nil
	}
	item := configs.ErlangEvent.Settings.GetById(id)
	return item
}

func GetErlangEventSettingByName(name string) *mainStruct.Param {
	if configs.ErlangEvent == nil {
		return nil
	}
	item := configs.ErlangEvent.Settings.GetByName(name)
	return item
}

func UpdateErlangEventSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateErlangEventSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchErlangEventSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchErlangEventSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelErlangEventSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelErlangEventSetting(param.Id)
	if !ok {
		return 0
	}

	configs.ErlangEvent.Settings.Remove(param)
	return id
}

func SetConfEventMulticast() (int64, error) {
	if configs.EventMulticast != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfEventMulticast, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewEventMulticast(res, true)
	return res, nil
}

func SetConfigEventMulticastSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.EventMulticast == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfEventMulticastSetting(configs.EventMulticast.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.EventMulticast.Settings.Set(param)
	return param, err
}

func GetConfigEventMulticast() (map[int64]*mainStruct.Param, bool) {
	if configs.EventMulticast == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.EventMulticast.Settings.GetList()
	return item, true
}

func GetEventMulticastSetting(id int64) *mainStruct.Param {
	if configs.EventMulticast == nil {
		return nil
	}
	item := configs.EventMulticast.Settings.GetById(id)
	return item
}

func GetEventMulticastSettingByName(name string) *mainStruct.Param {
	if configs.EventMulticast == nil {
		return nil
	}
	item := configs.EventMulticast.Settings.GetByName(name)
	return item
}

func UpdateEventMulticastSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateEventMulticastSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchEventMulticastSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchEventMulticastSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelEventMulticastSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelEventMulticastSetting(param.Id)
	if !ok {
		return 0
	}

	configs.EventMulticast.Settings.Remove(param)
	return id
}

func SetConfFax() (int64, error) {
	if configs.Fax != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfFax, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewFax(res, true)
	return res, nil
}

func SetConfigFaxSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Fax == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfFaxSetting(configs.Fax.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Fax.Settings.Set(param)
	return param, err
}

func GetConfigFax() (map[int64]*mainStruct.Param, bool) {
	if configs.Fax == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Fax.Settings.GetList()
	return item, true
}

func GetFaxSetting(id int64) *mainStruct.Param {
	if configs.Fax == nil {
		return nil
	}
	item := configs.Fax.Settings.GetById(id)
	return item
}

func GetFaxSettingByName(name string) *mainStruct.Param {
	if configs.Fax == nil {
		return nil
	}
	item := configs.Fax.Settings.GetByName(name)
	return item
}

func UpdateFaxSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateFaxSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchFaxSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchFaxSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelFaxSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelFaxSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Fax.Settings.Remove(param)
	return id
}

func SetConfLua() (int64, error) {
	if configs.Lua != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfLua, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewLua(res, true)
	return res, nil
}

func SetConfigLuaSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Lua == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfLuaSetting(configs.Lua.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Lua.Settings.Set(param)
	return param, err
}

func GetConfigLua() (map[int64]*mainStruct.Param, bool) {
	if configs.Lua == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Lua.Settings.GetList()
	return item, true
}

func GetLuaSetting(id int64) *mainStruct.Param {
	if configs.Lua == nil {
		return nil
	}
	item := configs.Lua.Settings.GetById(id)
	return item
}

func GetLuaSettingByName(name string) *mainStruct.Param {
	if configs.Lua == nil {
		return nil
	}
	item := configs.Lua.Settings.GetByName(name)
	return item
}

func UpdateLuaSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateLuaSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchLuaSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchLuaSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelLuaSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelLuaSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Lua.Settings.Remove(param)
	return id
}

func SetConfMongo() (int64, error) {
	if configs.Mongo != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfMongo, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewMongo(res, true)
	return res, nil
}

func SetConfigMongoSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Mongo == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfMongoSetting(configs.Mongo.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Mongo.Settings.Set(param)
	return param, err
}

func GetConfigMongo() (map[int64]*mainStruct.Param, bool) {
	if configs.Mongo == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Mongo.Settings.GetList()
	return item, true
}

func GetMongoSetting(id int64) *mainStruct.Param {
	if configs.Mongo == nil {
		return nil
	}
	item := configs.Mongo.Settings.GetById(id)
	return item
}

func GetMongoSettingByName(name string) *mainStruct.Param {
	if configs.Mongo == nil {
		return nil
	}
	item := configs.Mongo.Settings.GetByName(name)
	return item
}

func UpdateMongoSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateMongoSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchMongoSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchMongoSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelMongoSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelMongoSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Mongo.Settings.Remove(param)
	return id
}

func SetConfMsrp() (int64, error) {
	if configs.Msrp != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfMsrp, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewMsrp(res, true)
	return res, nil
}

func SetConfigMsrpSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Msrp == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfMsrpSetting(configs.Msrp.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Msrp.Settings.Set(param)
	return param, err
}

func GetConfigMsrp() (map[int64]*mainStruct.Param, bool) {
	if configs.Msrp == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Msrp.Settings.GetList()
	return item, true
}

func GetMsrpSetting(id int64) *mainStruct.Param {
	if configs.Msrp == nil {
		return nil
	}
	item := configs.Msrp.Settings.GetById(id)
	return item
}

func GetMsrpSettingByName(name string) *mainStruct.Param {
	if configs.Msrp == nil {
		return nil
	}
	item := configs.Msrp.Settings.GetByName(name)
	return item
}

func UpdateMsrpSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateMsrpSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchMsrpSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchMsrpSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelMsrpSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelMsrpSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Msrp.Settings.Remove(param)
	return id
}

func SetConfOreka() (int64, error) {
	if configs.Oreka != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfOreka, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewOreka(res, true)
	return res, nil
}

func SetConfigOrekaSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Oreka == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOrekaSetting(configs.Oreka.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Oreka.Settings.Set(param)
	return param, err
}

func GetConfigOreka() (map[int64]*mainStruct.Param, bool) {
	if configs.Oreka == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Oreka.Settings.GetList()
	return item, true
}

func GetOrekaSetting(id int64) *mainStruct.Param {
	if configs.Oreka == nil {
		return nil
	}
	item := configs.Oreka.Settings.GetById(id)
	return item
}

func GetOrekaSettingByName(name string) *mainStruct.Param {
	if configs.Oreka == nil {
		return nil
	}
	item := configs.Oreka.Settings.GetByName(name)
	return item
}

func UpdateOrekaSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateOrekaSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchOrekaSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchOrekaSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelOrekaSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelOrekaSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Oreka.Settings.Remove(param)
	return id
}

func SetConfPerl() (int64, error) {
	if configs.Perl != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfPerl, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewPerl(res, true)
	return res, nil
}

func SetConfigPerlSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Perl == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfPerlSetting(configs.Perl.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Perl.Settings.Set(param)
	return param, err
}

func GetConfigPerl() (map[int64]*mainStruct.Param, bool) {
	if configs.Perl == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Perl.Settings.GetList()
	return item, true
}

func GetPerlSetting(id int64) *mainStruct.Param {
	if configs.Perl == nil {
		return nil
	}
	item := configs.Perl.Settings.GetById(id)
	return item
}

func GetPerlSettingByName(name string) *mainStruct.Param {
	if configs.Perl == nil {
		return nil
	}
	item := configs.Perl.Settings.GetByName(name)
	return item
}

func UpdatePerlSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdatePerlSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchPerlSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchPerlSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelPerlSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelPerlSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Perl.Settings.Remove(param)
	return id
}

func SetConfPocketsphinx() (int64, error) {
	if configs.Pocketsphinx != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfPocketsphinx, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewPocketsphinx(res, true)
	return res, nil
}

func SetConfigPocketsphinxSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Pocketsphinx == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfPocketsphinxSetting(configs.Pocketsphinx.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Pocketsphinx.Settings.Set(param)
	return param, err
}

func GetConfigPocketsphinx() (map[int64]*mainStruct.Param, bool) {
	if configs.Pocketsphinx == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Pocketsphinx.Settings.GetList()
	return item, true
}

func GetPocketsphinxSetting(id int64) *mainStruct.Param {
	if configs.Pocketsphinx == nil {
		return nil
	}
	item := configs.Pocketsphinx.Settings.GetById(id)
	return item
}

func GetPocketsphinxSettingByName(name string) *mainStruct.Param {
	if configs.Pocketsphinx == nil {
		return nil
	}
	item := configs.Pocketsphinx.Settings.GetByName(name)
	return item
}

func UpdatePocketsphinxSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdatePocketsphinxSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchPocketsphinxSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchPocketsphinxSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelPocketsphinxSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelPocketsphinxSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Pocketsphinx.Settings.Remove(param)
	return id
}

func SetConfSangomaCodec() (int64, error) {
	if configs.SangomaCodec != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfSangomaCodec, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewSangomaCodec(res, true)
	return res, nil
}

func SetConfigSangomaCodecSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.SangomaCodec == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfSangomaCodecSetting(configs.SangomaCodec.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.SangomaCodec.Settings.Set(param)
	return param, err
}

func GetConfigSangomaCodec() (map[int64]*mainStruct.Param, bool) {
	if configs.SangomaCodec == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.SangomaCodec.Settings.GetList()
	return item, true
}

func GetSangomaCodecSetting(id int64) *mainStruct.Param {
	if configs.SangomaCodec == nil {
		return nil
	}
	item := configs.SangomaCodec.Settings.GetById(id)
	return item
}

func GetSangomaCodecSettingByName(name string) *mainStruct.Param {
	if configs.SangomaCodec == nil {
		return nil
	}
	item := configs.SangomaCodec.Settings.GetByName(name)
	return item
}

func UpdateSangomaCodecSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateSangomaCodecSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchSangomaCodecSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchSangomaCodecSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelSangomaCodecSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelSangomaCodecSetting(param.Id)
	if !ok {
		return 0
	}

	configs.SangomaCodec.Settings.Remove(param)
	return id
}

func SetConfSndfile() (int64, error) {
	if configs.Sndfile != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfSndfile, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewSndfile(res, true)
	return res, nil
}

func SetConfigSndfileSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Sndfile == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfSndfileSetting(configs.Sndfile.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Sndfile.Settings.Set(param)
	return param, err
}

func GetConfigSndfile() (map[int64]*mainStruct.Param, bool) {
	if configs.Sndfile == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Sndfile.Settings.GetList()
	return item, true
}

func GetSndfileSetting(id int64) *mainStruct.Param {
	if configs.Sndfile == nil {
		return nil
	}
	item := configs.Sndfile.Settings.GetById(id)
	return item
}

func GetSndfileSettingByName(name string) *mainStruct.Param {
	if configs.Sndfile == nil {
		return nil
	}
	item := configs.Sndfile.Settings.GetByName(name)
	return item
}

func UpdateSndfileSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateSndfileSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchSndfileSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchSndfileSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelSndfileSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelSndfileSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Sndfile.Settings.Remove(param)
	return id
}

func SetConfXmlCdr() (int64, error) {
	if configs.XmlCdr != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfXmlCdr, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewXmlCdr(res, true)
	return res, nil
}

func SetConfigXmlCdrSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.XmlCdr == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfXmlCdrSetting(configs.XmlCdr.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.XmlCdr.Settings.Set(param)
	return param, err
}

func GetConfigXmlCdr() (map[int64]*mainStruct.Param, bool) {
	if configs.XmlCdr == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.XmlCdr.Settings.GetList()
	return item, true
}

func GetXmlCdrSetting(id int64) *mainStruct.Param {
	if configs.XmlCdr == nil {
		return nil
	}
	item := configs.XmlCdr.Settings.GetById(id)
	return item
}

func GetXmlCdrSettingByName(name string) *mainStruct.Param {
	if configs.XmlCdr == nil {
		return nil
	}
	item := configs.XmlCdr.Settings.GetByName(name)
	return item
}

func UpdateXmlCdrSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateXmlCdrSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchXmlCdrSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchXmlCdrSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelXmlCdrSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelXmlCdrSetting(param.Id)
	if !ok {
		return 0
	}

	configs.XmlCdr.Settings.Remove(param)
	return id
}

func SetConfXmlRpc() (int64, error) {
	if configs.XmlRpc != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfXmlRpc, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewXmlRpc(res, true)
	return res, nil
}

func SetConfigXmlRpcSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.XmlRpc == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfXmlRpcSetting(configs.XmlRpc.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.XmlRpc.Settings.Set(param)
	return param, err
}

func GetConfigXmlRpc() (map[int64]*mainStruct.Param, bool) {
	if configs.XmlRpc == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.XmlRpc.Settings.GetList()
	return item, true
}

func GetXmlRpcSetting(id int64) *mainStruct.Param {
	if configs.XmlRpc == nil {
		return nil
	}
	item := configs.XmlRpc.Settings.GetById(id)
	return item
}

func GetXmlRpcSettingByName(name string) *mainStruct.Param {
	if configs.XmlRpc == nil {
		return nil
	}
	item := configs.XmlRpc.Settings.GetByName(name)
	return item
}

func UpdateXmlRpcSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateXmlRpcSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchXmlRpcSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchXmlRpcSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelXmlRpcSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelXmlRpcSetting(param.Id)
	if !ok {
		return 0
	}

	configs.XmlRpc.Settings.Remove(param)
	return id
}

func SetConfZeroconf() (int64, error) {
	if configs.Zeroconf != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfZeroconf, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewZeroconf(res, true)
	return res, nil
}

func SetConfigZeroconfSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Zeroconf == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfZeroconfSetting(configs.Zeroconf.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Zeroconf.Settings.Set(param)
	return param, err
}

func GetConfigZeroconf() (map[int64]*mainStruct.Param, bool) {
	if configs.Zeroconf == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Zeroconf.Settings.GetList()
	return item, true
}

func GetZeroconfSetting(id int64) *mainStruct.Param {
	if configs.Zeroconf == nil {
		return nil
	}
	item := configs.Zeroconf.Settings.GetById(id)
	return item
}

func GetZeroconfSettingByName(name string) *mainStruct.Param {
	if configs.Zeroconf == nil {
		return nil
	}
	item := configs.Zeroconf.Settings.GetByName(name)
	return item
}

func UpdateZeroconfSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateZeroconfSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchZeroconfSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchZeroconfSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelZeroconfSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelZeroconfSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Zeroconf.Settings.Remove(param)
	return id
}

func SetConfPostSwitch() (int64, error) {
	if configs.PostSwitch != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfPostLoadSwitch, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewPostSwitch(res, true)
	return res, nil
}

func SetConfigPostSwitchSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.PostSwitch == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfPostSwitchSetting(configs.PostSwitch.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.PostSwitch.Settings.Set(param)
	return param, err
}

func GetConfigPostSwitch() (map[int64]*mainStruct.Param, bool) {
	if configs.PostSwitch == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.PostSwitch.Settings.GetList()
	return item, true
}

func GetPostSwitchSetting(id int64) *mainStruct.Param {
	if configs.PostSwitch == nil {
		return nil
	}
	item := configs.PostSwitch.Settings.GetById(id)
	return item
}

func GetPostSwitchSettingByName(name string) *mainStruct.Param {
	if configs.PostSwitch == nil {
		return nil
	}
	item := configs.PostSwitch.Settings.GetByName(name)
	return item
}

func UpdatePostSwitchSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdatePostSwitchSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchPostSwitchSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchPostSwitchSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelPostSwitchSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelPostSwitchSetting(param.Id)
	if !ok {
		return 0
	}

	configs.PostSwitch.Settings.Remove(param)
	return id
}

func SetConfigPostSwitchCliKeybinding(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.PostSwitch == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfPostSwitchCliKeybinding(configs.PostSwitch.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.PostSwitch.CliKeybindings.Set(param)
	return param, err
}

func GetConfigPostSwitchCliKeybinding() (map[int64]*mainStruct.Param, bool) {
	if configs.PostSwitch == nil {
		return nil, false
	}

	item := configs.PostSwitch.CliKeybindings.GetList()
	return item, true
}

func GetPostSwitchCliKeybinding(id int64) *mainStruct.Param {
	if configs.PostSwitch == nil {
		return nil
	}
	item := configs.PostSwitch.CliKeybindings.GetById(id)
	return item
}

func GetPostSwitchCliKeybindingByName(name string) *mainStruct.Param {
	if configs.PostSwitch == nil {
		return nil
	}
	item := configs.PostSwitch.CliKeybindings.GetByName(name)
	return item
}

func UpdatePostSwitchCliKeybinding(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdatePostSwitchCliKeybinding(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchPostSwitchCliKeybinding(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchPostSwitchCliKeybinding(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelPostSwitchCliKeybinding(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelPostSwitchCliKeybinding(param.Id)
	if !ok {
		return 0
	}

	configs.PostSwitch.CliKeybindings.Remove(param)
	return id
}

func GetConfigPostSwitchDefaultPtime() (map[int64]*mainStruct.DefaultPtime, bool) {
	if configs.PostSwitch == nil {
		return nil, false
	}

	item := configs.PostSwitch.DefaultPtimes.GetList()
	return item, true
}

func SetConfigPostSwitchDefaultPtime(paramName, paramValue string) (*mainStruct.DefaultPtime, error) {
	if configs.PostSwitch == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfPostSwitchDefaultPtime(configs.PostSwitch.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.DefaultPtime{Id: res, Name: paramName, Ptime: paramValue, Enabled: true}
	configs.PostSwitch.DefaultPtimes.Set(param)
	return param, err
}

func GetPostSwitchDefaultPtime(id int64) *mainStruct.DefaultPtime {
	if configs.PostSwitch == nil {
		return nil
	}
	item := configs.PostSwitch.DefaultPtimes.GetById(id)
	return item
}

func GetPostSwitchDefaultPtimeByName(name string) *mainStruct.DefaultPtime {
	if configs.PostSwitch == nil {
		return nil
	}
	item := configs.PostSwitch.DefaultPtimes.GetByName(name)
	return item
}

func UpdatePostSwitchDefaultPtime(param *mainStruct.DefaultPtime, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdatePostSwitchDefaultPtime(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Ptime = value
	return param.Id, err
}

func SwitchPostSwitchDefaultPtime(param *mainStruct.DefaultPtime, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchPostSwitchDefaultPtime(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelPostSwitchDefaultPtime(param *mainStruct.DefaultPtime) int64 {
	id := param.Id
	ok := db.DelPostSwitchDefaultPtime(param.Id)
	if !ok {
		return 0
	}

	configs.PostSwitch.DefaultPtimes.Remove(param)
	return id
}

func SetConfDistributor() (int64, error) {
	if configs.Distributor != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfDistributor, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewDistributor(res, true)
	return res, nil
}

func SetConfDistributorList(listName string) (*mainStruct.DistributorList, error) {
	if configs.Distributor == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfDistributorList(configs.Distributor.Id, listName)
	if err != nil {
		return nil, err
	}

	list := &mainStruct.DistributorList{Id: res, Name: listName, Nodes: mainStruct.NewDistributorNodes(), Enabled: true}
	configs.Distributor.Lists.Set(list)
	configs.Distributor.Nodes = mainStruct.NewDistributorNodes()
	return list, err
}

func SetConfDistributorNode(list *mainStruct.DistributorList, name, weight string) (int64, error) {
	if configs.Distributor == nil {
		return 0, errors.New("no config")
	}

	if list == nil {
		return 0, errors.New("list name doesn't exists")
	}
	res, err := db.SetConfDistributorListNode(list.Id, name, weight)
	if err != nil {
		return 0, err
	}

	node := &mainStruct.DistributorNode{Id: res, Name: name, Weight: weight, List: list, Enabled: true}
	configs.Distributor.Nodes.Set(node)
	list.Nodes.Set(node)
	return res, err
}

func GetDistributorLists() (map[int64]*mainStruct.DistributorList, bool) {
	if configs.Distributor == nil {
		return map[int64]*mainStruct.DistributorList{}, false
	}

	item := configs.Distributor.Lists.GetList()
	return item, true
}

func GetDistributorList(id int64) *mainStruct.DistributorList {
	if configs.Distributor == nil {
		return nil
	}
	item := configs.Distributor.Lists.GetById(id)
	return item
}

func IsDistributorExists() bool {
	return configs.Distributor != nil
}

func DelDistributorNode(node *mainStruct.DistributorNode) int64 {
	parentId := node.List.Id
	ok := db.DelDistributorNode(node.Id)
	if !ok {
		return 0
	}

	node.List.Nodes.Remove(node)
	configs.Distributor.Nodes.Remove(node)
	return parentId
}

func GetDistributorNode(id int64) *mainStruct.DistributorNode {
	if configs.Distributor == nil {
		return nil
	}
	item := configs.Distributor.Nodes.GetById(id)
	return item
}

func UpdateDistributorNode(node *mainStruct.DistributorNode, name, weight string) (int64, error) {
	if node == nil {
		return 0, errors.New("node doesn't exists")
	}
	_, err := db.UpdateDistributorNode(node.Id, name, weight)
	if err != nil {
		return 0, err
	}
	node.Name = name
	node.Weight = weight
	return node.List.Id, err
}

func SwitchDistributorNode(node *mainStruct.DistributorNode, switcher bool) (int64, error) {
	if node == nil {
		return 0, errors.New("node doesn't exists")
	}
	_, err := db.SwitchDistributorNode(node.Id, switcher)
	if err != nil {
		return 0, err
	}
	node.Enabled = switcher
	return node.List.Id, err
}

func DelDistributorList(id int64) bool {
	list := configs.Distributor.Lists.GetById(id)
	if list == nil {
		return false
	}
	ok := db.DelDistributorList(list.Id)
	if !ok {
		return false
	}

	configs.Distributor.Lists.Remove(list)
	configs.Distributor.Nodes.ClearUp(configs)
	return true
}

func UpdateDistributorList(id int64, newName string) error {
	domain := configs.Distributor.Lists.GetById(id)
	if domain == nil {
		return errors.New("domain name doesn't exists")
	}
	err := db.UpdateDistributorList(id, newName)
	if err != nil {
		return err
	}
	configs.Distributor.Lists.Rename(domain.Name, newName)
	return err
}

func SetConfDirectory() (int64, error) {
	if configs.Directory != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfDirectory, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewDirectory(res, true)
	return res, nil
}

func SetConfigDirectorySetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Directory == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfDirectorySetting(configs.Directory.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Directory.Settings.Set(param)
	return param, err
}

func SetConfigDirectoryProfile(profileName string) (*mainStruct.DirectoryProfile, error) {
	if configs.Directory == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfDirectoryProfile(configs.Directory.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.DirectoryProfile{
		Id: res, Name: profileName,
		Params:  mainStruct.NewDirectoryProfileParams(),
		Enabled: true,
	}
	configs.Directory.Profiles.Set(profile)
	return profile, err
}

func SetConfigDirectoryProfileParam(profile *mainStruct.DirectoryProfile, name, value string) (*mainStruct.DirectoryProfileParam, error) {
	if configs.Directory == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, err := db.SetConfDirectoryProfileParam(profile.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.DirectoryProfileParam{Id: res, Name: name, Value: value, Enabled: true, Profile: profile}
	profile.Params.Set(param)
	configs.Directory.ProfileParams.Set(param)
	return param, err
}

func GetConfigDirectory() (map[int64]*mainStruct.Param, bool) {
	if configs.Directory == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Directory.Settings.GetList()
	return item, true
}

func GetDirectoryProfiles() (map[int64]*mainStruct.DirectoryProfile, bool) {
	if configs.Directory == nil {
		return map[int64]*mainStruct.DirectoryProfile{}, false
	}

	item := configs.Directory.Profiles.GetList()
	return item, true
}

func GetDirectoryProfile(id int64) *mainStruct.DirectoryProfile {
	if configs.Directory == nil {
		return nil
	}
	item := configs.Directory.Profiles.GetById(id)
	return item
}

func GetDirectoryProfileByName(name string) *mainStruct.DirectoryProfile {
	if configs.Directory == nil {
		return nil
	}
	item := configs.Directory.Profiles.GetByName(name)
	return item
}

func GetDirectorySetting(id int64) *mainStruct.Param {
	if configs.Directory == nil {
		return nil
	}
	item := configs.Directory.Settings.GetById(id)
	return item
}

func GetDirectorySettingByName(name string) *mainStruct.Param {
	if configs.Directory == nil {
		return nil
	}
	item := configs.Directory.Settings.GetByName(name)
	return item
}

func UpdateDirectorySetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateDirectorySetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchDirectorySetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchDirectorySetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelDirectorySetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelDirectorySetting(param.Id)
	if !ok {
		return 0
	}

	configs.Directory.Settings.Remove(param)
	return id
}

func GetDirectoryProfileParam(id int64) *mainStruct.DirectoryProfileParam {
	if configs.Directory == nil {
		return nil
	}
	item := configs.Directory.ProfileParams.GetById(id)
	return item
}

func DelDirectoryProfileParam(param *mainStruct.DirectoryProfileParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Profile.Id
	ok := db.DelDirectoryProfileParam(param.Id)
	if !ok {
		return 0
	}

	param.Profile.Params.Remove(param)
	configs.Directory.ProfileParams.Remove(param)
	return parentId
}

func SwitchDirectoryProfileParam(param *mainStruct.DirectoryProfileParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchDirectoryProfileParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Profile.Id, err
}

func UpdateDirectoryProfileParam(param *mainStruct.DirectoryProfileParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateDirectoryProfileParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Profile.Id, err
}

func UpdateDirectoryProfile(profile *mainStruct.DirectoryProfile, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateDirectoryProfile(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelDirectoryProfile(profile *mainStruct.DirectoryProfile) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelDirectoryProfile(profile.Id)
	if !ok {
		return 0
	}

	configs.Directory.Profiles.Remove(profile)
	configs.Directory.ProfileParams.ClearUp(configs.Directory)
	return profile.Id
}

func SetConfFifo() (int64, error) {
	if configs.Fifo != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfFifo, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewFifo(res, true)
	return res, nil
}

func SetConfigFifoSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Fifo == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfFifoSetting(configs.Fifo.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Fifo.Settings.Set(param)
	return param, err
}

func SetConfigFifoFifo(profileName, importance string) (*mainStruct.FifoFifo, error) {
	if configs.Fifo == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfFifoFifo(configs.Fifo.Id, profileName, importance)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.FifoFifo{
		Id:         res,
		Name:       profileName,
		Importance: importance,
		Params:     mainStruct.NewFifoFifoParams(),
		Enabled:    true,
	}
	configs.Fifo.Fifos.Set(profile)
	return profile, err
}

func SetConfigFifoFifoParam(profile *mainStruct.FifoFifo, timeout, simo, lag, body string) (*mainStruct.FifoFifoMember, error) {
	if configs.Fifo == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}

	res, err := db.SetConfFifoFifoParam(profile.Id, timeout, simo, lag, body)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.FifoFifoMember{Id: res, Timeout: timeout, Simo: simo, Lag: lag, Body: body, Enabled: true, Fifo: profile}
	profile.Params.Set(param)
	configs.Fifo.FifoParams.Set(param)
	return param, err
}

func GetConfigFifo() (map[int64]*mainStruct.Param, bool) {
	if configs.Fifo == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Fifo.Settings.GetList()
	return item, true
}

func GetFifoFifos() (map[int64]*mainStruct.FifoFifo, bool) {
	if configs.Fifo == nil {
		return map[int64]*mainStruct.FifoFifo{}, false
	}

	item := configs.Fifo.Fifos.GetList()
	return item, true
}

func GetFifoFifo(id int64) *mainStruct.FifoFifo {
	if configs.Fifo == nil {
		return nil
	}
	item := configs.Fifo.Fifos.GetById(id)
	return item
}

func GetFifoFifoByName(name string) *mainStruct.FifoFifo {
	if configs.Fifo == nil {
		return nil
	}
	item := configs.Fifo.Fifos.GetByName(name)
	return item
}

func GetFifoSetting(id int64) *mainStruct.Param {
	if configs.Fifo == nil {
		return nil
	}
	item := configs.Fifo.Settings.GetById(id)
	return item
}

func GetFifoSettingByName(name string) *mainStruct.Param {
	if configs.Fifo == nil {
		return nil
	}
	item := configs.Fifo.Settings.GetByName(name)
	return item
}

func UpdateFifoFifoImportance(id int64, newValue string) (int64, error) {
	list := configs.Fifo.Fifos.GetById(id)
	if list == nil {
		return 0, errors.New("list doesn't exists")
	}
	res, err := db.UpdateFifoFifoImportance(list.Id, newValue)
	if err != nil {
		return 0, err
	}
	list.Importance = newValue
	return res, err
}

func UpdateFifoSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateFifoSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchFifoSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchFifoSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelFifoSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelFifoSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Fifo.Settings.Remove(param)
	return id
}

func GetFifoFifoParam(id int64) *mainStruct.FifoFifoMember {
	if configs.Fifo == nil {
		return nil
	}
	item := configs.Fifo.FifoParams.GetById(id)
	return item
}

func DelFifoFifoParam(param *mainStruct.FifoFifoMember) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Fifo.Id
	ok := db.DelFifoFifoParam(param.Id)
	if !ok {
		return 0
	}

	param.Fifo.Params.Remove(param)
	configs.Fifo.FifoParams.Remove(param)
	return parentId
}

func SwitchFifoFifoParam(param *mainStruct.FifoFifoMember, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchFifoFifoParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Fifo.Id, err
}

func UpdateFifoFifoParam(param *mainStruct.FifoFifoMember, timeout, simo, lag, body string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateFifoFifoParam(param.Id, timeout, simo, lag, body)
	if err != nil {
		return 0, err
	}
	param.Timeout = timeout
	param.Simo = simo
	param.Lag = lag
	param.Body = body
	return param.Fifo.Id, err
}

func UpdateFifoFifo(profile *mainStruct.FifoFifo, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateFifoFifo(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelFifoFifo(profile *mainStruct.FifoFifo) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelFifoFifo(profile.Id)
	if !ok {
		return 0
	}

	configs.Fifo.Fifos.Remove(profile)
	configs.Fifo.FifoParams.ClearUp(configs.Fifo)
	return profile.Id
}

func SetConfOpal() (int64, error) {
	if configs.Opal != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfOpal, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewOpal(res, true)
	return res, nil
}

func SetConfigOpalSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Opal == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOpalSetting(configs.Opal.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Opal.Settings.Set(param)
	return param, err
}

func SetConfigOpalListener(profileName string) (*mainStruct.OpalListener, error) {
	if configs.Opal == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOpalListener(configs.Opal.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.OpalListener{
		Id: res, Name: profileName,
		Params:  mainStruct.NewOpalListenerParams(),
		Enabled: true,
	}
	configs.Opal.Listeners.Set(profile)
	return profile, err
}

func SetConfigOpalListenerParam(profile *mainStruct.OpalListener, name, value string) (*mainStruct.OpalListenerParam, error) {
	if configs.Opal == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, err := db.SetConfOpalListenerParam(profile.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.OpalListenerParam{Id: res, Name: name, Value: value, Enabled: true, Listener: profile}
	profile.Params.Set(param)
	configs.Opal.ListenerParams.Set(param)
	return param, err
}

func GetConfigOpal() (map[int64]*mainStruct.Param, bool) {
	if configs.Opal == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Opal.Settings.GetList()
	return item, true
}

func GetOpalListeners() (map[int64]*mainStruct.OpalListener, bool) {
	if configs.Opal == nil {
		return map[int64]*mainStruct.OpalListener{}, false
	}

	item := configs.Opal.Listeners.GetList()
	return item, true
}

func GetOpalListener(id int64) *mainStruct.OpalListener {
	if configs.Opal == nil {
		return nil
	}
	item := configs.Opal.Listeners.GetById(id)
	return item
}

func GetOpalListenerByName(name string) *mainStruct.OpalListener {
	if configs.Opal == nil {
		return nil
	}
	item := configs.Opal.Listeners.GetByName(name)
	return item
}

func GetOpalSetting(id int64) *mainStruct.Param {
	if configs.Opal == nil {
		return nil
	}
	item := configs.Opal.Settings.GetById(id)
	return item
}

func GetOpalSettingByName(name string) *mainStruct.Param {
	if configs.Opal == nil {
		return nil
	}
	item := configs.Opal.Settings.GetByName(name)
	return item
}

func UpdateOpalSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateOpalSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchOpalSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchOpalSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelOpalSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelOpalSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Opal.Settings.Remove(param)
	return id
}

func GetOpalListenerParam(id int64) *mainStruct.OpalListenerParam {
	if configs.Opal == nil {
		return nil
	}
	item := configs.Opal.ListenerParams.GetById(id)
	return item
}

func DelOpalListenerParam(param *mainStruct.OpalListenerParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Listener.Id
	ok := db.DelOpalListenerParam(param.Id)
	if !ok {
		return 0
	}

	param.Listener.Params.Remove(param)
	configs.Opal.ListenerParams.Remove(param)
	return parentId
}

func SwitchOpalListenerParam(param *mainStruct.OpalListenerParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchOpalListenerParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Listener.Id, err
}

func UpdateOpalListenerParam(param *mainStruct.OpalListenerParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateOpalListenerParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Listener.Id, err
}

func UpdateOpalListener(profile *mainStruct.OpalListener, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateOpalListener(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelOpalListener(profile *mainStruct.OpalListener) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelOpalListener(profile.Id)
	if !ok {
		return 0
	}

	configs.Opal.Listeners.Remove(profile)
	configs.Opal.ListenerParams.ClearUp(configs.Opal)
	return profile.Id
}

func SetConfOsp() (int64, error) {
	if configs.Osp != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfOsp, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewOsp(res, true)
	return res, nil
}

func SetConfigOspSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Osp == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOspSetting(configs.Osp.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Osp.Settings.Set(param)
	return param, err
}

func SetConfigOspProfile(profileName string) (*mainStruct.OspProfile, error) {
	if configs.Osp == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfOspProfile(configs.Osp.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.OspProfile{
		Id: res, Name: profileName,
		Params:  mainStruct.NewOspProfileParams(),
		Enabled: true,
	}
	configs.Osp.Profiles.Set(profile)
	return profile, err
}

func SetConfigOspProfileParam(profile *mainStruct.OspProfile, name, value string) (*mainStruct.OspProfileParam, error) {
	if configs.Osp == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, err := db.SetConfOspProfileParam(profile.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.OspProfileParam{Id: res, Name: name, Value: value, Enabled: true, Profile: profile}
	profile.Params.Set(param)
	configs.Osp.ProfileParams.Set(param)
	return param, err
}

func GetConfigOsp() (map[int64]*mainStruct.Param, bool) {
	if configs.Osp == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Osp.Settings.GetList()
	return item, true
}

func GetOspProfiles() (map[int64]*mainStruct.OspProfile, bool) {
	if configs.Osp == nil {
		return map[int64]*mainStruct.OspProfile{}, false
	}

	item := configs.Osp.Profiles.GetList()
	return item, true
}

func GetOspProfile(id int64) *mainStruct.OspProfile {
	if configs.Osp == nil {
		return nil
	}
	item := configs.Osp.Profiles.GetById(id)
	return item
}

func GetOspProfileByName(name string) *mainStruct.OspProfile {
	if configs.Osp == nil {
		return nil
	}
	item := configs.Osp.Profiles.GetByName(name)
	return item
}

func GetOspSetting(id int64) *mainStruct.Param {
	if configs.Osp == nil {
		return nil
	}
	item := configs.Osp.Settings.GetById(id)
	return item
}

func GetOspSettingByName(name string) *mainStruct.Param {
	if configs.Osp == nil {
		return nil
	}
	item := configs.Osp.Settings.GetByName(name)
	return item
}

func UpdateOspSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateOspSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchOspSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchOspSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelOspSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelOspSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Osp.Settings.Remove(param)
	return id
}

func GetOspProfileParam(id int64) *mainStruct.OspProfileParam {
	if configs.Osp == nil {
		return nil
	}
	item := configs.Osp.ProfileParams.GetById(id)
	return item
}

func DelOspProfileParam(param *mainStruct.OspProfileParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Profile.Id
	ok := db.DelOspProfileParam(param.Id)
	if !ok {
		return 0
	}

	param.Profile.Params.Remove(param)
	configs.Osp.ProfileParams.Remove(param)
	return parentId
}

func SwitchOspProfileParam(param *mainStruct.OspProfileParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchOspProfileParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Profile.Id, err
}

func UpdateOspProfileParam(param *mainStruct.OspProfileParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateOspProfileParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Profile.Id, err
}

func UpdateOspProfile(profile *mainStruct.OspProfile, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateOspProfile(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.Name = name
	return profile.Id, err
}

func DelOspProfile(profile *mainStruct.OspProfile) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelOspProfile(profile.Id)
	if !ok {
		return 0
	}

	configs.Osp.Profiles.Remove(profile)
	configs.Osp.ProfileParams.ClearUp(configs.Osp)
	return profile.Id
}

func SetConfUnicall() (int64, error) {
	if configs.Unicall != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfUnicall, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewUnicall(res, true)
	return res, nil
}

func SetConfigUnicallSetting(paramName, paramValue string) (*mainStruct.Param, error) {
	if configs.Unicall == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfUnicallSetting(configs.Unicall.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.Param{Id: res, Name: paramName, Value: paramValue, Enabled: true}
	configs.Unicall.Settings.Set(param)
	return param, err
}

func SetConfigUnicallSpan(profileName string) (*mainStruct.UnicallSpan, error) {
	if configs.Unicall == nil {
		return nil, errors.New("no config")
	}
	res, err := db.SetConfUnicallSpan(configs.Unicall.Id, profileName)
	if err != nil {
		return nil, err
	}

	profile := &mainStruct.UnicallSpan{
		Id:      res,
		SpanId:  profileName,
		Params:  mainStruct.NewUnicallSpanParams(),
		Enabled: true,
	}
	configs.Unicall.Spans.Set(profile)
	return profile, err
}

func SetConfigUnicallSpanParam(profile *mainStruct.UnicallSpan, name, value string) (*mainStruct.UnicallSpanParam, error) {
	if configs.Unicall == nil {
		return nil, errors.New("no config")
	}
	if profile == nil {
		return nil, errors.New("no profile")
	}
	if name == "" {
		return nil, errors.New("no param")
	}

	res, err := db.SetConfUnicallSpanParam(profile.Id, name, value)
	if err != nil {
		return nil, err
	}

	param := &mainStruct.UnicallSpanParam{Id: res, Name: name, Value: value, Enabled: true, Span: profile}
	profile.Params.Set(param)
	configs.Unicall.SpanParams.Set(param)
	return param, err
}

func GetConfigUnicall() (map[int64]*mainStruct.Param, bool) {
	if configs.Unicall == nil {
		return map[int64]*mainStruct.Param{}, false
	}

	item := configs.Unicall.Settings.GetList()
	return item, true
}

func GetUnicallSpans() (map[int64]*mainStruct.UnicallSpan, bool) {
	if configs.Unicall == nil {
		return map[int64]*mainStruct.UnicallSpan{}, false
	}

	item := configs.Unicall.Spans.GetList()
	return item, true
}

func GetUnicallSpan(id int64) *mainStruct.UnicallSpan {
	item := configs.Unicall.Spans.GetById(id)
	return item
}

func GetUnicallSpanByName(name string) *mainStruct.UnicallSpan {
	if configs.Unicall == nil {
		return nil
	}
	if configs.Unicall == nil {
		return nil
	}
	item := configs.Unicall.Spans.GetByName(name)
	return item
}

func GetUnicallSetting(id int64) *mainStruct.Param {
	if configs.Unicall == nil {
		return nil
	}
	item := configs.Unicall.Settings.GetById(id)
	return item
}

func GetUnicallSettingByName(name string) *mainStruct.Param {
	if configs.Unicall == nil {
		return nil
	}
	item := configs.Unicall.Settings.GetByName(name)
	return item
}

func UpdateUnicallSetting(param *mainStruct.Param, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateUnicallSetting(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Id, err
}

func SwitchUnicallSetting(param *mainStruct.Param, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchUnicallSetting(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Id, err
}

func DelUnicallSetting(param *mainStruct.Param) int64 {
	id := param.Id
	ok := db.DelUnicallSetting(param.Id)
	if !ok {
		return 0
	}

	configs.Unicall.Settings.Remove(param)
	return id
}

func GetUnicallSpanParam(id int64) *mainStruct.UnicallSpanParam {
	if configs.Unicall == nil {
		return nil
	}
	item := configs.Unicall.SpanParams.GetById(id)
	return item
}

func DelUnicallSpanParam(param *mainStruct.UnicallSpanParam) int64 {
	if param == nil {
		return 0
	}
	parentId := param.Span.Id
	ok := db.DelUnicallSpanParam(param.Id)
	if !ok {
		return 0
	}

	param.Span.Params.Remove(param)
	configs.Unicall.SpanParams.Remove(param)
	return parentId
}

func SwitchUnicallSpanParam(param *mainStruct.UnicallSpanParam, switcher bool) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.SwitchUnicallSpanParam(param.Id, switcher)
	if err != nil {
		return 0, err
	}
	param.Enabled = switcher
	return param.Span.Id, err
}

func UpdateUnicallSpanParam(param *mainStruct.UnicallSpanParam, name, value string) (int64, error) {
	if param == nil {
		return 0, errors.New("param doesn't exists")
	}
	_, err := db.UpdateUnicallSpanParam(param.Id, name, value)
	if err != nil {
		return 0, err
	}
	param.Name = name
	param.Value = value
	return param.Span.Id, err
}

func UpdateUnicallSpan(profile *mainStruct.UnicallSpan, name string) (int64, error) {
	if profile == nil {
		return 0, errors.New("profile doesn't exists")
	}
	_, err := db.UpdateUnicallSpan(profile.Id, name)
	if err != nil {
		return 0, err
	}
	profile.SpanId = name
	return profile.Id, err
}

func DelUnicallSpan(profile *mainStruct.UnicallSpan) int64 {
	if profile == nil {
		return 0
	}
	ok := db.DelUnicallSpan(profile.Id)
	if !ok {
		return 0
	}

	configs.Unicall.Spans.Remove(profile)
	configs.Unicall.SpanParams.ClearUp(configs.Unicall)
	return profile.Id
}

func SetConfConference() (int64, error) {
	if configs.Conference != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfConference, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewConference(res, true)
	return res, nil
}

func SetConfConferenceAdvertise(name, value string) (*mainStruct.ConfigConferenceAdvertiseRooms, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := configs.Conference.Advertise.NewSubItem()
	item.Name = name
	item.Status = value
	item.Enabled = true
	res, err := db.InsertFields(item.Parent.Parent.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	configs.Conference.Advertise.Set(item)
	return item, err
}

func SetConfConferenceCallerControlsGroup(groupName string) (*mainStruct.ConfigConferenceCallerControlsGroups, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := configs.Conference.CallerControlsGroups.NewSubItem()
	item.Name = groupName
	item.Controls = mainStruct.NewControls(item)
	item.Enabled = true
	res, err := db.InsertFields(configs.Conference.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}
	item.Id = res
	configs.Conference.CallerControlsGroups.Set(item)
	return item, err
}

func SetConfConferenceCallerControlsGroupControl(group *mainStruct.ConfigConferenceCallerControlsGroups, name, value string) (*mainStruct.ConfigConferenceCallerControlsControls, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := group.Controls.NewSubItem()
	item.Action = name
	item.Digits = value
	item.Enabled = true
	res, err := db.InsertFields(group.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	group.Controls.Set(item)
	return item, err
}

func SetConfConferenceProfile(itemName string) (*mainStruct.ConfigConferenceProfiles, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := configs.Conference.Profiles.NewSubItem()
	item.Name = itemName
	item.Params = mainStruct.NewConferenceProfileParams(item)
	item.Enabled = true
	res, err := db.InsertFields(configs.Conference.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	configs.Conference.Profiles.Set(item)
	return item, err
}

func SetConfConferenceProfileParam(profile *mainStruct.ConfigConferenceProfiles, name, value string) (*mainStruct.ConfigConferenceProfilesParams, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := profile.Params.NewSubItem()
	item.Name = name
	item.Value = value
	item.Enabled = true
	res, err := db.InsertFields(profile.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	profile.Params.Set(item)
	return item, err
}

func GetConfCache(id int64, table mainStruct.TableMethods) {
	sub := table.NewSubItemInterface()

	res := db.GetDataById(id, sub.GetTableName(), sub)
	for _, val := range res {
		table.SetFromInterface(val)
	}

	return
}

func GetConfigConference() (map[int64]*mainStruct.ConfigConferenceAdvertiseRooms, bool) {
	if configs.Conference == nil {
		return map[int64]*mainStruct.ConfigConferenceAdvertiseRooms{}, false
	}

	item := configs.Conference.Advertise.GetList()
	return item, true
}

func GetConferenceProfiles() (map[int64]*mainStruct.ConfigConferenceProfiles, bool) {
	if configs.Conference == nil {
		return map[int64]*mainStruct.ConfigConferenceProfiles{}, false
	}

	item := configs.Conference.Profiles.GetList()
	return item, true
}

func GetConferenceCallerControls() (map[int64]*mainStruct.ConfigConferenceCallerControlsGroups, bool) {
	if configs.Conference == nil {
		return map[int64]*mainStruct.ConfigConferenceCallerControlsGroups{}, false
	}

	item := configs.Conference.CallerControlsGroups.GetList()
	return item, true
}

func GetConferenceProfile(id int64) *mainStruct.ConfigConferenceProfiles {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.Profiles.GetById(id)
	return item
}

func GetConferenceProfileByName(name string) *mainStruct.ConfigConferenceProfiles {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.Profiles.GetByName(name)
	return item
}

func GetConferenceRoom(id int64) *mainStruct.ConfigConferenceAdvertiseRooms {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.Advertise.GetById(id)
	return item
}

func GetConferenceRoomByName(name string) *mainStruct.ConfigConferenceAdvertiseRooms {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.Advertise.GetByName(name)
	return item
}

func SwitchConfigRow(item mainStruct.RowItem, switcher bool) error {
	if item == nil {
		return errors.New("room doesn't exists")
	}
	err := db.SwitchFields(item.GetTableName(), item.GetId(), switcher)
	if err != nil {
		return err
	}
	item.SetEnabled(switcher)
	return err
}

func DelConfigRow(item mainStruct.RowItem) bool {
	if item == nil {
		return false
	}
	ok := db.DeleteRow(item.GetTableName(), item.GetId())
	if !ok {
		return ok
	}
	item.Remove()
	return ok
}

func GetConferenceProfileParam(id int64) *mainStruct.ConfigConferenceProfilesParams {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.ConferenceProfileParams.GetById(id)
	return item
}

func UpdateConfigRow(item mainStruct.RowItem, values ...string) error {
	if item == nil {
		return errors.New("empty item")
	}
	err := db.UpdateFields(item.GetTableName(), item.ForUpdate(values))
	if err != nil {
		return err
	}

	item.Update(values)
	return err
}

func GetConferenceCallerControl(id int64) *mainStruct.ConfigConferenceCallerControlsControls {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.Controls.GetById(id)
	return item
}

func GetConferenceCallerControlGroup(id int64) *mainStruct.ConfigConferenceCallerControlsGroups {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.CallerControlsGroups.GetById(id)
	return item
}

func GetConferenceCallerControlGroupByName(name string) *mainStruct.ConfigConferenceCallerControlsGroups {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.CallerControlsGroups.GetByName(name)
	return item
}

func SetConfConferenceChatPermissionsProfile(profileName string) (*mainStruct.ConfigConferenceChatPermissions, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := configs.Conference.ChatPermissions.NewSubItem()
	item.Name = profileName
	item.Users = mainStruct.NewConferenceChatPermissionUsers(item)
	item.Enabled = true
	res, err := db.InsertFields(configs.Conference.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}
	item.Id = res
	configs.Conference.ChatPermissions.Set(item)
	return item, err
}

func SetConfConferenceChatPermissionsUser(profile *mainStruct.ConfigConferenceChatPermissions, name, value string) (*mainStruct.ConfigConferenceChatPermissionUsers, error) {
	if configs.Conference == nil {
		return nil, errors.New("no config")
	}

	item := profile.Users.NewSubItem()
	item.Name = name
	item.Commands = value
	item.Enabled = true
	res, err := db.InsertFields(profile.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	profile.Users.Set(item)
	return item, err
}

func GetConferenceChatPermissionsProfiles() (map[int64]*mainStruct.ConfigConferenceChatPermissions, bool) {
	if configs.Conference == nil {
		return map[int64]*mainStruct.ConfigConferenceChatPermissions{}, false
	}

	item := configs.Conference.ChatPermissions.GetList()
	return item, true
}

func GetConferenceChatPermissionsProfile(id int64) *mainStruct.ConfigConferenceChatPermissions {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.ChatPermissions.GetById(id)
	return item
}

func GetConferenceChatPermissionsProfileByName(name string) *mainStruct.ConfigConferenceChatPermissions {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.ChatPermissions.GetByName(name)
	return item
}

func GetConferenceChatPermissionsUser(id int64) *mainStruct.ConfigConferenceChatPermissionUsers {
	if configs.Conference == nil {
		return nil
	}
	item := configs.Conference.Users.GetById(id)
	return item
}

func SetConfConferenceLayouts() (int64, error) {
	if configs.ConferenceLayouts != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfConferenceLayouts, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewConferenceLayouts(res, true)
	return res, nil
}

func SetConfConferenceLayoutsGroups(itemName string) (*mainStruct.ConfigConferenceLayoutsGroups, error) {
	if configs.ConferenceLayouts == nil {
		return nil, errors.New("no config")
	}

	item := configs.ConferenceLayouts.ConferenceLayoutsGroups.NewSubItem()
	item.Name = itemName
	item.Layouts = mainStruct.NewGroupLayouts(item)
	item.Enabled = true
	res, err := db.InsertFields(configs.Conference.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}
	item.Id = res
	configs.ConferenceLayouts.ConferenceLayoutsGroups.Set(item)
	return item, err
}

func SetConfConferenceLayoutsGroupLayout(group *mainStruct.ConfigConferenceLayoutsGroups, body string) (*mainStruct.ConfigConferenceLayoutsGroupLayouts, error) {
	if configs.ConferenceLayouts == nil {
		return nil, errors.New("no config")
	}

	item := group.Layouts.NewSubItem()
	item.Body = body
	item.Enabled = true
	res, err := db.InsertFields(group.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	group.Layouts.Set(item)
	return item, err
}

func SetConfConferenceLayoutLayouts(itemName string) (*mainStruct.ConfigConferenceLayouts, error) {
	if configs.ConferenceLayouts == nil {
		return nil, errors.New("no config")
	}

	item := configs.ConferenceLayouts.Layouts.NewSubItem()
	item.Name = itemName
	item.LayoutsImages = mainStruct.NewLayoutsImages(item)
	item.Enabled = true
	res, err := db.InsertFields(configs.Conference.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}
	item.Id = res
	configs.ConferenceLayouts.Layouts.Set(item)
	return item, err
}

func SetConfConferenceLayoutLayoutsImage(image *mainStruct.ConfigConferenceLayouts, x, y, scale, floor, floorOnly, hScale, overlap, reservationId, zoom string) (*mainStruct.ConfigConferenceLayoutsImages, error) {
	if configs.ConferenceLayouts == nil {
		return nil, errors.New("no config")
	}

	item := image.LayoutsImages.NewSubItem()
	item.X = x
	item.Y = y
	item.Scale = scale
	item.Floor = floor
	item.FloorOnly = floorOnly
	item.Hscale = hScale
	item.Overlap = overlap
	item.ReservationId = reservationId
	item.Zoom = zoom
	item.Enabled = true
	res, err := db.InsertFields(image.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	image.LayoutsImages.Set(item)
	return item, err
}

func SetConfPostLoadModules() (int64, error) {
	if configs.PostLoadModules != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfPostLoadModules, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewPostLoadModules(res, true)
	return res, nil
}

func GetConfigPostLoadModules() (map[int64]*mainStruct.ModuleTag, bool) {
	if configs.PostLoadModules == nil {
		return map[int64]*mainStruct.ModuleTag{}, false
	}

	item := configs.PostLoadModules.Modules.GetList()
	return item, true
}

func GetPostLoadModule(id int64) *mainStruct.ModuleTag {
	if configs.PostLoadModules == nil {
		return nil
	}
	item := configs.PostLoadModules.Modules.GetById(id)
	return item
}

func GetPostLoadModuleByName(name string) *mainStruct.ModuleTag {
	if configs.PostLoadModules == nil {
		return nil
	}
	item := configs.PostLoadModules.Modules.GetByName(name)
	return item
}

func SetPostLoadModule(name string) (*mainStruct.ModuleTag, error) {
	if configs.PostLoadModules == nil {
		return nil, errors.New("no config")
	}

	item := configs.PostLoadModules.Modules.NewSubItem()
	item.Name = name
	item.Enabled = true
	res, err := db.InsertFields(item.Parent.Parent.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	configs.PostLoadModules.Modules.Set(item)
	return item, err
}

func SetConfVoicemail() (int64, error) {
	if configs.Voicemail != nil {
		return 0, errors.New("config already exists")
	}

	res, err := db.SetConf(mainStruct.ConfVoicemail, cache.GetCurrentInstanceId())
	if err != nil {
		return 0, err
	}
	configs.NewVoicemail(res, true)
	return res, nil
}

func GetConfigVoicemailSettings() (map[int64]*mainStruct.VoicemailSettingsParameter, bool) {
	if configs.Voicemail == nil {
		return map[int64]*mainStruct.VoicemailSettingsParameter{}, false
	}

	item := configs.Voicemail.Settings.GetList()
	return item, true
}

func GetVoicemailSettings(id int64) *mainStruct.VoicemailSettingsParameter {
	if configs.Voicemail == nil {
		return nil
	}
	item := configs.Voicemail.Settings.GetById(id)
	return item
}

func GetVoicemailSettingByName(name string) *mainStruct.VoicemailSettingsParameter {
	if configs.Voicemail == nil {
		return nil
	}
	item := configs.Voicemail.Settings.GetByName(name)
	return item
}

func SetVoicemailSetting(name, value string) (*mainStruct.VoicemailSettingsParameter, error) {
	if configs.Voicemail == nil {
		return nil, errors.New("no config")
	}

	item := configs.Voicemail.Settings.NewSubItem()
	item.Name = name
	item.Value = value
	item.Enabled = true
	res, err := db.InsertFields(item.Parent.Parent.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	configs.Voicemail.Settings.Set(item)
	return item, err
}

func GetConfigVoicemailProfiles() (map[int64]*mainStruct.VoicemailProfile, bool) {
	if configs.Voicemail == nil {
		return map[int64]*mainStruct.VoicemailProfile{}, false
	}

	item := configs.Voicemail.Profiles.GetList()
	return item, true
}

func SetVoicemailProfile(name string) (*mainStruct.VoicemailProfile, error) {
	if configs.Voicemail == nil {
		return nil, errors.New("no config")
	}
	if name == "" {
		return nil, errors.New("no name")
	}

	item := configs.Voicemail.Profiles.NewSubItem()
	item.Name = name
	item.Enabled = true
	item.Params = mainStruct.NewVoicemailProfileParams(item)
	res, err := db.InsertFields(item.Parent.Parent.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	configs.Voicemail.Profiles.Set(item)
	return item, err
}

func GetVoicemailProfile(id int64) *mainStruct.VoicemailProfile {
	if configs.Voicemail == nil {
		return nil
	}
	item := configs.Voicemail.Profiles.GetById(id)
	return item
}

func GetVoicemailProfileByName(name string) *mainStruct.VoicemailProfile {
	if configs.Voicemail == nil {
		return nil
	}
	item := configs.Voicemail.Profiles.GetByName(name)
	return item
}

func SetVoicemailProfiles(name string) (*mainStruct.VoicemailProfile, error) {
	if configs.Voicemail == nil {
		return nil, errors.New("no config")
	}

	item := configs.Voicemail.Profiles.NewSubItem()
	item.Name = name
	item.Enabled = true
	res, err := db.InsertFields(item.Parent.Parent.Id, item.GetTableName(), item)
	if err != nil {
		return nil, err
	}

	item.Id = res
	configs.Voicemail.Profiles.Set(item)
	return item, err
}

func SetVoicemailProfileParam(profile *mainStruct.VoicemailProfile, name, value string) (*mainStruct.VoicemailProfilesParameter, error) {
	item := profile.Params.NewSubItem()
	item.Name = name
	item.Value = value
	item.Enabled = true
	res, err := db.InsertFields(profile.Id, item.GetTableName(), item)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	item.Id = res
	profile.Params.Set(item)
	return item, err
}

func CheckVoicemailConfigExists() bool {
	return configs.Voicemail != nil
}

func GetVoicemailProfileParam(id int64) *mainStruct.VoicemailProfilesParameter {
	if configs.Voicemail == nil {
		return nil
	}
	item := configs.Voicemail.ProfileParams.GetById(id)
	return item
}

func GetVoicemailProfileParamByName(name string) *mainStruct.VoicemailProfilesParameter {
	if configs.Voicemail == nil {
		return nil
	}
	item := configs.Voicemail.ProfileParams.GetByName(name)
	return item
}
