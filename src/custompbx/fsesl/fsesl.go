package fsesl

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/cfg"
	"custompbx/daemonCache"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webcache"
	"custompbx/xmlStruct"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/zusrut/fsock"
	"golang.org/x/net/html/charset"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var esl *fsock.FSock

type regsAsJson struct {
	Registrations []Regs `json:"rows"`
	RowCount      int    `json:"row_count"`
}

type Regs struct {
	RegUser     string `json:"reg_user"`
	Realm       string `json:"realm"`
	Token       string `json:"token"`
	Url         string `json:"url"`
	Expires     string `json:"expires"`
	NetworkIp   string `json:"network_ip"`
	NetworkPort string `json:"network_port"`
	Hostname    string `json:"hostname"`
	Metadata    string `json:"metadata"`
}

type callsAsJson struct {
	Calls    []mainStruct.Call `json:"rows"`
	RowCount int               `json:"row_count"`
}

type channelsAsJson struct {
	Channels []mainStruct.Channel `json:"rows"`
	RowCount int                  `json:"row_count"`
}

type XmlVerto struct {
	XMLName xml.Name        `xml:"profiles"`
	Profile []VertoProfiles `xml:"profile"`
	Clients []VertoClients  `xml:"client"`
}

type VertoClients struct {
	XMLName xml.Name `xml:"client"`
	Profile string   `xml:"profile"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
	Data    string   `xml:"data"`
	State   string   `xml:"state"`
}

type VertoProfiles struct {
	XMLName xml.Name `xml:"profile"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
	Data    string   `xml:"data"`
	State   string   `xml:"state"`
}

type XmlSofia struct {
	XMLName  xml.Name       `xml:"profiles"`
	Profiles []SofiaProfile `xml:"profile"`
	Gateways []SofiaGateway `xml:"gateway"`
}

type SofiaGateway struct {
	XMLName xml.Name `xml:"gateway"`
	Profile string   `xml:"profile"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
	Data    string   `xml:"data"`
	State   string   `xml:"state"`
}

type SofiaProfile struct {
	XMLName xml.Name `xml:"profile"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
	Data    string   `xml:"data"`
	State   string   `xml:"state"`
}

const (
	EventName              = "Event-Name"
	EventSubclass          = "Event-Subclass"
	NameChannelCreate      = "CHANNEL_CREATE"
	NameChannelDestroy     = "CHANNEL_DESTROY"
	NameChannelAnswer      = "CHANNEL_ANSWER"
	NameCustom             = "CUSTOM"
	NamePublish            = "PUBLISH"
	NameUnPublish          = "UNPUBLISH"
	HeaderFromHost         = "from-host"
	HeaderVertoProfileName = "verto_profile_name"
	HeaderVertoLogin       = "verto_login"
	HeaderVertoSuccess     = "verto_success"
	HeaderFromUser         = "from-user"
	HeaderHost             = "host"
	HeaderUser             = "user"
	HeaderStatus           = "status"
	HeaderProfileName      = "profile_name"
	HeaderProfileNameDash  = "profile-name"
	HeaderProfileUri       = "profile_uri"
	HeaderGateway          = "Gateway"
	HeaderState            = "State"
	NameCSGatewayAdd       = "CUSTOM sofia::gateway_add"
	NameCSGatewayDelete    = "CUSTOM sofia::gateway_delete"
	NameCSGatewayState     = "CUSTOM sofia::gateway_state"
	NameSubGatewayAdd      = "sofia::gateway_add"
	NameSubGatewayDelete   = "sofia::gateway_delete"
	NameSubGatewayState    = "sofia::gateway_state"
	NameCSRegister         = "CUSTOM sofia::register"
	NameCSUnregister       = "CUSTOM sofia::unregister"
	NameCSExpire           = "CUSTOM sofia::expire"
	NameCSInfo             = "CUSTOM sofia::info"
	NameCVClientConnect    = "CUSTOM verto::client_connect"
	NameCVClientDisconnect = "CUSTOM verto::client_disconnect"
	NameCVLogin            = "CUSTOM verto::login"
	NameCCallcenter        = "CUSTOM callcenter::info"
	NameModuleLoad         = "MODULE_LOAD"
	NameModuleUnload       = "MODULE_UNLOAD"
	NameName               = "name"
	NameKey                = "key"

	// NameChannelUuid              = "Channel-Call-UUID"
	NameChannelUuid                  = "Unique-ID"
	NameChannelDirection             = "Call-Direction"
	NameChannelDate                  = "Event-Date-Local"
	NameChannelCreatedEpoch          = "Caller-Channel-Created-Time"
	NameChannelName                  = "variable_channel_name"
	NameChannelState                 = "Channel-State"
	NameChannelCallerName            = "Caller-Caller-ID-Name"
	NameChannelCallerNumber          = "Caller-Caller-ID-Number"
	NameChannelCallerIp              = "Caller-Network-Addr"
	NameChannelCallerDest            = "Caller-Destination-Number"
	NameChannelPresenceId            = "Channel-Presence-ID"
	NameChannelPresenceData          = "Channel-Presence-Data"
	NameChannelAccountcode           = "variable_accountcode"
	NameChannelCallState             = "Channel-Call-State"
	NameChannelCallUuid              = "variable_call_uuid"
	NameChannelHostname              = "FreeSWITCH-Hostname"
	NameChannelOtherLegUuid          = "Other-Leg-Unique-ID"
	NameChannelCalleeIdName          = "Caller-Callee-ID-Name"
	NameChannelCalleeIdNumber        = "Caller-Callee-ID-Number"
	NameChannelApplication           = "variable_current_application"
	NameChannelApplicationData       = "variable_current_application_data"
	NameChannelDialplan              = "Caller-Dialplan"
	NameChannelContext               = "Caller-Context"
	NameChannelReadCodec             = "Channel-Read-Codec-Name"
	NameChannelReadCodecRate         = "Channel-Read-Codec-Rate"
	NameChannelReadCodecBitRate      = "Channel-Read-Codec-Bit-Rate"
	NameChannelWriteCodec            = "Channel-Write-Codec-Name"
	NameChannelWriteCodecRate        = "Channel-Write-Codec-Rate"
	NameChannelWriteCodecBitRate     = "Channel-Write-Codec-Bit-Rate"
	NameCallerChannelAnswerTIme      = "Caller-Channel-Answered-Time"
	NameCallcenterAction             = "CC-Action"
	NameCallcenterAgent              = "CC-Agent"
	NameCallcenterAgentStatus        = "CC-Agent-Status"
	NameCallcenterAgentState         = "CC-Agent-State"
	NameCallcenterQueue              = "CC-Queue"
	NameCallcenterMemberUuid         = "CC-Member-UUID"
	NameCallcenterMemberSessionUuid  = "CC-Member-Session-UUID"
	NameCallcenterMemberCIDName      = "CC-Member-CID-Name"
	NameCallcenterMemberCIDNumber    = "CC-Member-CID-Number"
	VariableCCQueueJoinedEpoch       = "variable_cc_queue_joined_epoch"
	NameCallcenterMemberAnsweredTime = "CC-Agent-Answered-Time"
	NameEventDateTimestamp           = "Event-Date-Timestamp"

	CommandModuleExists     = "module_exists"
	CommandGlobalGetVar     = "global_getvar"
	CommandSwitchname       = "switchname"
	ValueChannelStateActive = "ACTIVE"

	CCMemberStateUnknown   = "unknown"
	CCMemberStateWaiting   = "waiting"
	CCMemberStateTrying    = "trying"
	CCMemberStateAnswered  = "answered"
	CCMemberStateAbandoned = "abandoned"
)

func CollectAndSetESLDataESLData() {
	if esl == nil || !esl.Connected() {
		return
	}
	GetDirectorySipRegs()
	GetChannels()
	//GetSofiaStatuses()
	GetVertoStatuses()
	GetLoadedModules()
	SetGlobalVariables()
}

func SetGlobalVariables() {
	vars := pbxcache.GetGlobalVariableNamedList()
	if len(vars) == 0 {
		return
	}
	res, err := SendBgapiCmd("global_getvar")
	if err != nil {
		log.Println("Cant get global vars: " + err.Error())
		return
	}
	s := strings.Split(res, "\n")
	var toDel []string
	var toSet []struct {
		Cmd      string
		Position int64
	}
	for _, str := range s {
		keyVal := strings.Split(str, "=")
		if len(keyVal) < 1 {
			continue
		}
		if vars[keyVal[0]] == nil && !mainStruct.IsDynamicGlobalVar(keyVal[0]) {
			toDel = append(toDel, keyVal[0])
		}
		if len(keyVal) < 2 {
			keyVal = append(keyVal, "")
		}
		if vars[keyVal[0]] != nil && vars[keyVal[0]].Value == keyVal[1] {
			delete(vars, keyVal[0])
		}
		//update dynamic global variable value
		if mainStruct.IsDynamicGlobalVar(keyVal[0]) {
			gvar := pbxcache.GetGlobalVariableByName(keyVal[0])
			if gvar == nil {
				pbxcache.SetGlobalVariable(keyVal[0], keyVal[1], "set", true)
				continue
			}
			if gvar.Value == keyVal[1] {
				continue
			}
			pbxcache.UpdateGlobalVariable(gvar, gvar.Name, "set", keyVal[1])
		}
	}
	for _, td := range toDel {
		SendBgapiCmd("global_setvar " + td + "=")
	}

	for name, value := range vars {
		if mainStruct.IsDynamicGlobalVar(name) {
			continue
		}
		toSet = append(toSet, struct {
			Cmd      string
			Position int64
		}{Cmd: "global_setvar " + name + "=" + value.Value, Position: value.Position})
	}

	sort.SliceStable(toSet, func(i, j int) bool {
		return toSet[i].Position < toSet[j].Position
	})
	for _, cmd := range toSet {
		SendBgapiCmd(cmd.Cmd)
	}

	DropToFileWithEsl(pbxcache.GlobalVariablesDropToFile())
}

func SaveToFile(body, path, fileName string, withEsl bool) {
	if body == "" || path == "" || fileName == "" {
		log.Println("Cant DropToFileWithEsl: body, path, fileName: ", body, path, fileName)
		return
	}
	if withEsl {
		log.Println("ESL!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		DropToFileWithEsl(body, path, fileName)
		return
	}
	log.Println("OS!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	DropToFileWithOs(body, path, fileName)
}

func DropToFileWithEsl(body, path, fileName string) {
	_, err := SendBgapiCmd("system echo -e '" + body + "' > " + path + "/" + fileName + ".pr")
	if err != nil {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
		return
	}
	_, err = SendBgapiCmd("system cp -n " + path + "/" + fileName + " " + path + "/" + fileName + ".dump")
	if err != nil {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
		return
	}
	_, err = SendBgapiCmd("system mv " + path + "/" + fileName + ".pr " + path + "/" + fileName)
	if err != nil {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
	}
}

func copyFile(src, dst string, bufferSize int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s already exists", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, bufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

func DropToFileWithOs(body, path, fileName string) {
	f, err := os.Create(path + "/" + fileName + ".pr")
	if err != nil || f == nil {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
	}
	defer f.Close()

	body = strings.Replace(body, "\\n", "\n", -1)
	_, err2 := f.WriteString(body)
	if err2 != nil {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
	}

	err = copyFile(path+"/"+fileName, path+"/"+fileName+".dump", 4096)
	if err != nil && err.Error() != fmt.Sprintf("file %s already exists", path+"/"+fileName+".dump") {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
		return
	}
	err = os.Rename(path+"/"+fileName+".pr", path+"/"+fileName)
	if err != nil {
		log.Println("Cant DropToFileWithEsl: " + err.Error())
	}
}

func ImportGlobalVariables() {
	vars := pbxcache.GetGlobalVariableNamedList()
	if len(vars) != 0 {
		log.Println("Global vars already exists")
		return
	}
	res, err := SendBgapiCmd("global_getvar")
	if err != nil {
		log.Println("Cant get global vars: " + err.Error())
		return
	}
	s := strings.Split(res, "\n")
	for _, str := range s {
		keyVal := strings.Split(str, "=")
		if len(keyVal) < 1 || keyVal[0] == "" {
			continue
		}
		if len(keyVal) < 2 {
			keyVal = append(keyVal, "")
		}
		pbxcache.SetGlobalVariable(keyVal[0], keyVal[1], "set", mainStruct.IsDynamicGlobalVar(keyVal[0]))
	}
}

func ParseGlobalVars(in string) string {
	r := regexp.MustCompile(`\$\${[^}]*}`)
	matches := r.ReplaceAllStringFunc(in, func(part string) string {
		response, err := esl.SendApiCmd(CommandGlobalGetVar + " " + part[3:len(part)-1])
		if err != nil {
			return ""
		}

		return response
	})

	return matches
}

func FirstConnectData() {
	evHandlers := map[string][]func(string, int){}
	evFilters := make(map[string][]string)
	tmpEsl, err := fsock.NewFSock(cfg.CustomPbx.Fs.Esl.Host+":"+strconv.Itoa(cfg.CustomPbx.Fs.Esl.Port), cfg.CustomPbx.Fs.Esl.Pass, cfg.CustomPbx.Fs.Esl.Timeout, evHandlers, nil, evFilters, nil, 1)
	if err != nil {
		return
	}
	defer tmpEsl.Disconnect()

	out, err := tmpEsl.SendBgapiCmd(CommandSwitchname)
	if err != nil {
		return
	}

	val := <-out
	if strings.Contains(val, "-ERR") {
		return
	}
	cache.SetCurrentInstanceName(val)

}

func ESLConnectKeeper(eventChannel chan interface{}, logsChannel chan mainStruct.LogType) {
	log.Println("Connection to ESL")
	tenSecondsTick := time.Tick(10 * time.Second)
	Connect(eventChannel, logsChannel)
	go StartListenFSEvents()
	CollectAndSetESLDataESLData()

	for {
		select {
		case <-tenSecondsTick:
			if esl != nil && esl.Connected() {
				daemonCache.State.ESLConnection = true
				continue
			}
			log.Println("Reconnecting to ESL...")
			daemonCache.State.ESLConnection = false
			Connect(eventChannel, logsChannel)
			if esl == nil || !esl.Connected() {
				continue
			}
			switchName, _ := SendBgapiCmd(CommandSwitchname)
			if switchName == "" || switchName != cache.GetCurrentInstanceName() {
				log.Println("FS actual switchname is different from config file!!!")
				continue
			}
			go StartListenFSEvents()
			CollectAndSetESLDataESLData()
		}
	}
}

func Connect(eventChannel chan interface{}, logsChannel chan mainStruct.LogType) {
	if esl != nil && esl.Connected() {
		return
	}
	var err error
	evFilters := make(map[string][]string)
	evFilters[EventName] = append(evFilters[EventName], NameChannelCreate)
	evFilters[EventName] = append(evFilters[EventName], NameChannelAnswer)
	evFilters[EventName] = append(evFilters[EventName], NameChannelDestroy)
	evFilters[EventName] = append(evFilters[EventName], NameCustom)
	evFilters[EventName] = append(evFilters[EventName], NamePublish)
	evFilters[EventName] = append(evFilters[EventName], NameUnPublish)
	evFilters[EventName] = append(evFilters[EventName], NameModuleUnload)
	evFilters[EventName] = append(evFilters[EventName], NameModuleLoad)

	// We are interested in heartbeats, channel_answer, channel_hangup define handler for them
	evHandlers := map[string][]func(string, int){
		NameChannelCreate:      {func(event string, id int) { channelCreateHandler(event, id, eventChannel) }},
		NameChannelAnswer:      {func(event string, id int) { channelAnswerHandler(event, id, eventChannel) }},
		NameChannelDestroy:     {func(event string, id int) { channelDestroyHandler(event, id, eventChannel) }},
		NameCSRegister:         {func(event string, id int) { sofiaRegsHandler(event, id, eventChannel) }},
		NameCSUnregister:       {func(event string, id int) { sofiaRegsHandler(event, id, eventChannel) }},
		NameCSExpire:           {func(event string, id int) { sofiaRegsHandler(event, id, eventChannel) }},
		NameCSInfo:             {func(event string, id int) { sofiaRegsHandler(event, id, eventChannel) }},
		NameCSGatewayAdd:       {func(event string, id int) { sofiaGatewayHandler(event, id, eventChannel) }},
		NameCSGatewayDelete:    {func(event string, id int) { sofiaGatewayHandler(event, id, eventChannel) }},
		NameCSGatewayState:     {func(event string, id int) { sofiaGatewayHandler(event, id, eventChannel) }},
		NamePublish:            {func(event string, id int) { sofiaProfileHandler(event, id, eventChannel) }},
		NameUnPublish:          {func(event string, id int) { sofiaProfileHandler(event, id, eventChannel) }},
		NameModuleUnload:       {func(event string, id int) { moduleHandler(event, id, eventChannel) }},
		NameModuleLoad:         {func(event string, id int) { moduleHandler(event, id, eventChannel) }},
		NameCVLogin:            {func(event string, id int) { vertoRegsHandler(event, id, eventChannel) }},
		NameCVClientDisconnect: {func(event string, id int) { vertoRegsHandler(event, id, eventChannel) }},
		NameCCallcenter:        {func(event string, id int) { callcenterHandler(event, id, eventChannel) }},
		// "CUSTOM sofia::profile_start": {func(event string, id int) { sofiaProfileHandler(event, id, eventChannel) }},
		// "CUSTOM sofia::profile_stop": {func(event string, id int) { sofiaProfileHandler(event, id, eventChannel) }},
	}

	logHandler := fsock.LogHandler{}
	if cfg.CustomPbx.Fs.Esl.CollectLogs > 0 && cfg.CustomPbx.Fs.Esl.CollectLogs < 11 {
		logHandler = fsock.LogHandler{Level: cfg.CustomPbx.Fs.Esl.CollectLogs, Handler: func(headers map[string]string, body string) { logHandlerFunc(headers, body, logsChannel) }}
	}

	esl, err = fsock.NewFSock(cfg.CustomPbx.Fs.Esl.Host+":"+strconv.Itoa(cfg.CustomPbx.Fs.Esl.Port), cfg.CustomPbx.Fs.Esl.Pass, cfg.CustomPbx.Fs.Esl.Timeout, evHandlers, &logHandler, evFilters, nil, 1)
	if err != nil || esl == nil || !esl.Connected() {
		daemonCache.State.ESLConnection = false
	} else {
		daemonCache.State.ESLConnection = true
		log.Println("Esl connected.")
	}

	eventChannel <- daemonCache.State
}

func OneTimeConnectCommand(command string) string {
	r := rand.New(rand.NewSource(4564))
	connId := r.Intn(997) + 2
	evFilters := make(map[string][]string)
	evHandlers := map[string][]func(string, int){}

	oneTimeESL, err := fsock.NewFSock(cfg.CustomPbx.Fs.Esl.Host+":"+strconv.Itoa(cfg.CustomPbx.Fs.Esl.Port), cfg.CustomPbx.Fs.Esl.Pass, cfg.CustomPbx.Fs.Esl.Timeout, evHandlers, &fsock.LogHandler{}, evFilters, nil, connId)
	if err != nil {
		return err.Error()
	}
	defer oneTimeESL.Disconnect()

	resp, err := oneTimeESL.SendApiCmd(command)
	if err != nil {
		return err.Error()
	}

	return resp
}

func StartListenFSEvents() {
	if esl == nil || !esl.Connected() {
		return
	}
	err := esl.ReadEvents()
	if err != nil && esl != nil {
		_ = esl.Disconnect()
		log.Println("Esl Disconnected.")
		return
	}
}

func logHandlerFunc(headers map[string]string, body string, logsChannel chan mainStruct.LogType) {
	logLine, err := strconv.Atoi(headers["Log-Line"])
	if err != nil {
		logLine = 0
	}
	logLevel, err := strconv.Atoi(headers["Log-Level"])
	if err != nil {
		logLevel = 0
	}
	textChannel, err := strconv.Atoi(headers["Text-Channel"])
	if err != nil {
		textChannel = 0
	}
	logsChannel <- mainStruct.LogType{
		LogFile:     headers["Log-File"],
		LogFunc:     headers["Log-Func"],
		LogLine:     logLine,
		LogLevel:    logLevel,
		TextChannel: textChannel,
		UserData:    headers["User-Data"],
		Body:        body,
	}
}

func moduleHandler(event string, id int, eventChannel chan interface{}) {
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)
	module, err := altData.GetModuleByName(eventMap[NameKey])
	if module == nil || err != nil {
		module, err = altData.GetModuleByName(eventMap[NameName])
	}
	if module == nil || err != nil {
		// log.Println("no cache for module: " + eventMap[NameName])
		return
	}

	module.Loaded = eventMap[EventName] == NameModuleLoad

	conf := &altStruct.Configurations{}
	conf.GetConfigurationAndUpdate(module.Name, module)
	eventChannel <- conf

	/*
	   switch module.Module {
	   case "mod_sofia":

	   	GetSofiaStatuses()

	   case "mod_verto":

	   		GetVertoStatuses()
	   	}
	*/
}

func getUserByPresence(presense string) *altStruct.DirectoryDomainUser {
	r := regexp.MustCompile(`^.*?([^/]*)@(.*)$`)
	res := r.FindStringSubmatch(presense)

	if len(res) != 3 || res[1] == "" || res[2] == "" {
		return nil
	}
	domain := GetDomainByName(res[2])
	if domain == nil {
		return nil
	}
	return GetDomainUserByName(res[1], domain)
}

func channelCreateHandler(event string, id int, eventChannel chan interface{}) {
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)
	cCache := pbxcache.GetChannelsCache()

	channel := &mainStruct.Channel{
		Uuid:            eventMap[NameChannelUuid],
		Direction:       eventMap[NameChannelDirection],
		Created:         eventMap[NameChannelDate],
		CreatedEpoch:    eventMap[NameChannelCreatedEpoch],
		Name:            eventMap[NameChannelName],
		State:           eventMap[NameChannelState],
		CidName:         eventMap[NameChannelCallerName],
		CidNum:          eventMap[NameChannelCallerNumber],
		IpAddr:          eventMap[NameChannelCallerIp],
		Dest:            eventMap[NameChannelCallerDest],
		CallUuid:        eventMap[NameChannelCallUuid],
		Application:     eventMap[NameChannelApplication],
		ApplicationData: eventMap[NameChannelApplicationData],
		Dialplan:        eventMap[NameChannelDialplan],
		Context:         eventMap[NameChannelContext],
		ReadCodec:       eventMap[NameChannelReadCodec],
		ReadRate:        eventMap[NameChannelReadCodecRate],
		ReadBitRate:     eventMap[NameChannelReadCodecBitRate],
		WriteCodec:      eventMap[NameChannelWriteCodec],
		WriteRate:       eventMap[NameChannelWriteCodecRate],
		WriteBitRate:    eventMap[NameChannelWriteCodecBitRate],
		// Secure: eventMap[],
		Hostname:     eventMap[NameChannelHostname],
		PresenceId:   eventMap[NameChannelPresenceId],
		PresenceData: eventMap[NameChannelPresenceData],
		Accountcode:  eventMap[NameChannelAccountcode],
		Callstate:    eventMap[NameChannelCallState],
		/*		CalleeName: eventMap[],
				CalleeNum: eventMap[],
				CalleeDirection: eventMap[],
				CallUuid: eventMap[],
				SentCalleeName: eventMap[],
				SentCalleeNum: eventMap[],
				InitialCidName: eventMap[],
				InitialCidNum: eventMap[],
				InitialIpAddr: eventMap[],
				InitialDest: eventMap[],
				InitialDialplan: eventMap[],
				InitialContext: eventMap[],*/

	}

	cCache.Set(channel)
	cCache.Total++
	eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{CallsCounter: map[string]int{"total": cCache.Total, "answered": cCache.Answered}}}

	possibleUserName := channel.PresenceId
	if possibleUserName == "" {
		possibleUserName = channel.Name
	}
	user := getUserByPresence(possibleUserName)
	if user == nil {
		return
	}

	directoryCache := cache.GetDirectoryCache()
	cUser := directoryCache.UserCache.GetById(user.Id)
	if cUser == nil {
		cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
	}
	cUser.InCall = true
	cUser.CallDirection = channel.Direction
	intDate, err := strconv.ParseInt(channel.CreatedEpoch, 10, 64)
	if err == nil {
		cUser.CallDate = intDate / 1000000
	}
	if channel.Callstate == ValueChannelStateActive {
		cUser.LastUuid = channel.Uuid
		cUser.Talking = true
	}
	cUser.UpdateUser(user)
	eventChannel <- user
}

func channelAnswerHandler(event string, id int, eventChannel chan interface{}) {
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)

	cCache := pbxcache.GetChannelsCache()
	channel := cCache.GetByUuid(eventMap[NameChannelUuid])
	if channel == nil {
		// channel = cache.GetByUuid(eventMap[NameChannelOtherLegUuid])
	}
	if channel == nil {
		return
	}
	channel.Callstate = ValueChannelStateActive
	cCache.Answered++
	eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{CallsCounter: map[string]int{"total": cCache.Total, "answered": cCache.Answered}}}

	possibleUserName := channel.PresenceId
	if possibleUserName == "" {
		possibleUserName = channel.Name
	}
	user := getUserByPresence(possibleUserName)
	if user == nil {
		return
	}
	directoryCache := cache.GetDirectoryCache()
	cUser := directoryCache.UserCache.GetById(user.Id)
	if cUser == nil {
		cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
	}
	cUser.Talking = true
	cUser.CallDirection = channel.Direction
	cUser.LastUuid = channel.Uuid
	cUser.CallDate = time.Now().Unix()
	cUser.UpdateUser(user)
	eventChannel <- user
}

func channelDestroyHandler(event string, id int, eventChannel chan interface{}) {
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)

	cCache := pbxcache.GetChannelsCache()
	channel := cCache.GetByUuid(eventMap[NameChannelUuid])
	if channel == nil {
		channel = cCache.GetByUuid(eventMap[NameChannelOtherLegUuid])
	}
	if channel == nil {
		return
	}

	// if eventMap[NameCallerChannelAnswerTIme] != "0"{
	if channel.Callstate == ValueChannelStateActive {
		cCache.Answered--
	}
	possibleUserName := channel.PresenceId
	if possibleUserName == "" {
		possibleUserName = channel.Name
	}
	cCache.Total--
	cCache.Remove(channel)

	eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{CallsCounter: map[string]int{"total": cCache.Total, "answered": cCache.Answered}}}

	user := getUserByPresence(possibleUserName)
	if user == nil {
		return
	}
	directoryCache := cache.GetDirectoryCache()
	cUser := directoryCache.UserCache.GetById(user.Id)
	if cUser == nil {
		cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
	}
	cUser.InCall = false
	cUser.Talking = false
	cUser.LastUuid = ""
	cUser.CallDirection = ""
	cUser.CallDate = time.Now().Unix()
	cUser.UpdateUser(user)
	eventChannel <- user
}

func sofiaRegsHandler(event string, id int, eventChannel chan interface{}) {
	if !altData.IsDirectoryEnabled() {
		return
	}
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)

	domainName := eventMap[HeaderFromHost]
	if domainName == "" {
		domainName = eventMap[HeaderHost]
	}

	domain := GetDomainByName(domainName)
	if domain == nil {
		log.Println("no domain: " + domainName)
		return
	}

	userName := eventMap[HeaderFromUser]
	if userName == "" {
		userName = eventMap[HeaderUser]
	}

	user := GetDomainUserByName(userName, domain)
	if user == nil {
		log.Println("no user: " + userName)
		return
	}
	directoryCache := cache.GetDirectoryCache()
	cDomain := directoryCache.DomainCache.GetById(domain.Id)
	if cDomain == nil {
		cDomain = directoryCache.DomainCache.SetByData(domain.Id, domain.Name, 0)
	}
	cUser := directoryCache.UserCache.GetById(user.Id)
	if cUser == nil {
		cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
	}
	if len(eventMap[HeaderStatus]) > 10 && eventMap[HeaderStatus][:10] == "Registered" {
		if cUser.SipRegister {
			return
		}
		cUser.SipRegister = true
		cDomain.SipRegsCounter++
		cUser.UpdateUser(user)
		eventChannel <- user
	} else {
		if !cUser.SipRegister {
			return
		}
		cUser.SipRegister = false
		cDomain.SipRegsCounter--
		cUser.UpdateUser(user)
		eventChannel <- user
	}
	domain.SipRegsCounter = cDomain.SipRegsCounter
	eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{DomainSipRegs: map[string]int{domain.Name: domain.SipRegsCounter}}}
}

func GetDomainsWithNames() map[string]*altStruct.DirectoryDomain {
	domainI, err := intermediateDB.GetByValue(&altStruct.DirectoryDomain{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}}, map[string]bool{"Parent": true})
	if err != nil {
		return nil
	}
	res := map[string]*altStruct.DirectoryDomain{}
	for _, d := range domainI {
		domain, ok := d.(altStruct.DirectoryDomain)
		if !ok {
			continue
		}
		res[domain.Name] = &domain
	}
	return res
}

func GetDomainByName(domainName string) *altStruct.DirectoryDomain {
	domainI, err := intermediateDB.GetByValue(&altStruct.DirectoryDomain{Name: domainName, Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}}, map[string]bool{"Name": true, "Parent": true})
	if err != nil || len(domainI) == 0 {
		return nil
	}
	domain, ok := domainI[0].(altStruct.DirectoryDomain)
	if !ok {
		return nil
	}
	return &domain
}

func GetDomainUserByName(userName string, domain *altStruct.DirectoryDomain) *altStruct.DirectoryDomainUser {
	userI, err := intermediateDB.GetByValue(&altStruct.DirectoryDomainUser{Name: userName, Parent: domain}, map[string]bool{"Name": true, "Parent": true})
	if err != nil || len(userI) == 0 {
		return nil
	}
	user, ok := userI[0].(altStruct.DirectoryDomainUser)
	if !ok {
		return nil
	}
	return &user
}

func sofiaProfileHandler(event string, id int, eventChannel chan interface{}) {
	if !altData.IsSofiaExists() {
		return
	}
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)
	profileName := eventMap[HeaderProfileName]
	profile := altData.GetSofiaProfileByName(profileName)
	if profile == nil {
		log.Println("no profile")
		return
	}
	if eventMap[EventName] == NamePublish {
		profile.Started = true
		profile.Uri = eventMap[HeaderProfileUri]
	} else if eventMap[EventName] == NameUnPublish {
		gateways := altData.GetSofiaProfileGateways(profile.Id)
		eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{SofiaGateways: &gateways}}
		for _, gateway := range gateways {
			eventChannel <- &gateway
		}
	}
	profiles := []*altStruct.ConfigSofiaProfile{profile}
	eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{SofiaProfiles: &profiles}}
	eventChannel <- profile
}

func sofiaGatewayHandler(event string, id int, eventChannel chan interface{}) {
	if !altData.IsSofiaExists() {
		return
	}
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)
	gatewayName := eventMap[HeaderGateway]
	gateway := altData.GetSofiaProfileGateway(gatewayName)
	if gateway == nil {
		log.Println("no gateway")
		return
	}
	switch eventMap[EventSubclass] {
	case NameSubGatewayAdd:
		gateway.Started = true
		gateway.State = "NOREG()"
	case NameSubGatewayDelete:
		gateway.Started = false
		gateway.State = ""
	case NameSubGatewayState:
		gateway.State = eventMap[HeaderState]
	}

	eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{SofiaGateways: gateway}}
	eventChannel <- gateway
}

func vertoRegsHandler(event string, id int, eventChannel chan interface{}) {
	if !altData.IsVertoExists() {
		return
	}
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)
	if eventMap[HeaderVertoLogin] == "" {
		return
	}

	profile := altData.GetVertoProfileByName(eventMap[HeaderVertoProfileName])
	if profile == nil {
		return
	}
	var domain *altStruct.DirectoryDomain
	forceDomainParam := altData.GetVertoProfileParamByName(profile.Id, "force-register-domain")
	if forceDomainParam != nil && forceDomainParam.Enabled {
		domain = GetDomainByName(forceDomainParam.Value)
	}

	userName := eventMap[HeaderVertoLogin]

	r := regexp.MustCompile(`^(.+)@(.+)$`)
	res := r.FindStringSubmatch(eventMap[HeaderVertoLogin])

	if len(res) == 3 && res[1] != "" && res[2] != "" {
		if domain == nil {
			return
		}
		userName = res[1]
	}
	if domain == nil {
		domain = GetDomainByName(res[2])
		if domain == nil {
			log.Println("no domain")
			return
		}
	}
	user := GetDomainUserByName(userName, domain)
	if user == nil {
		log.Println("no user")
		return
	}
	directoryCache := cache.GetDirectoryCache()
	cUser := directoryCache.UserCache.GetById(user.Id)
	if cUser == nil {
		cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
	}
	if eventMap[HeaderVertoSuccess] == "1" {
		if cUser.VertoRegister {
			return
		}
		// domain.SipRegsCounter++
		cUser.VertoRegister = true
		cUser.UpdateUser(user)
		eventChannel <- user
	} else {
		if !cUser.VertoRegister {
			return
		}
		// domain.SipRegsCounter--
		cUser.VertoRegister = false
		cUser.UpdateUser(user)
		eventChannel <- user
	}
	// eventChannel <- &mainStruct.Dashboard{FSMetrics: &mainStruct.FSMetrics{DomainSipRegs: map[string]int{domain.Name: domain.SipRegsCounter}}}
}

func callcenterHandler(event string, id int, eventChannel chan interface{}) {
	if !altData.IsCallcenterEnabled() {
		return
	}
	var filterHeaders []string
	eventMap := fsock.FSEventStrToMap(event, filterHeaders)
	switch eventMap[NameCallcenterAction] {
	case "agent-status-change":
		agent := altData.GetCallcenterAgentByName(eventMap[NameCallcenterAgent])
		if agent == nil {
			return
		}
		eventStatus := eventMap[NameCallcenterAgentStatus]
		if agent.Status == eventStatus || eventStatus == "" {
			return
		}
		/*"Logged Out":"Available":"Available (On Demand)":	"On Break":*/
		eventChannel <- agent
	case "agent-state-change":
		agent := altData.GetCallcenterAgentByName(eventMap[NameCallcenterAgent])
		if agent == nil {
			return
		}
		eventState := eventMap[NameCallcenterAgentState]
		if agent.State == eventState || eventState == "" {
			return
		}
		/*"Idle":"Waiting":"Receiving":"In a queue call":*/
		eventChannel <- agent
		/*
			case "member-queue-start":
				epoch, err := strconv.ParseInt(eventMap[VariableCCQueueJoinedEpoch], 10, 64)
				if err != nil {
					epoch = 0
				}
				_, _ = pbxcache.SetConfCallcenterMemberCache(
					eventMap[NameCallcenterMemberUuid], CCMemberStateWaiting, eventMap[NameCallcenterQueue], "", 0, 0,
					0, eventMap[NameCallcenterMemberCIDName], eventMap[NameCallcenterMemberCIDNumber], epoch, 0, "", "",
					eventMap[NameCallcenterMemberSessionUuid], 0, epoch)

			case "member-queue-end":
				member := pbxcache.GetCallcenterMember(eventMap[NameCallcenterMemberUuid])
				if member == nil {
					return
				}
				pbxcache.DelCallcenterMemberCache(member)
			case "agent-offering":
				member := pbxcache.GetCallcenterMember(eventMap[NameCallcenterMemberUuid])
				if member == nil {
					return
				}
				member.ServingAgent = eventMap[NameCallcenterAgent]
				member.State = CCMemberStateTrying
			case "bridge-agent-fail":
				member := pbxcache.GetCallcenterMember(eventMap[NameCallcenterMemberUuid])
				if member == nil {
					return
				}
				member.ServingAgent = ""
				member.State = CCMemberStateWaiting
			case "bridge-agent-start":
				member := pbxcache.GetCallcenterMember(eventMap[NameCallcenterMemberUuid])
				if member == nil {
					return
				}
				member.ServingAgent = eventMap[NameCallcenterAgent]
				epoch, err := strconv.ParseInt(eventMap[NameCallcenterMemberAnsweredTime], 10, 64)
				if err != nil {
					epoch = 0
				}
				member.AbandonedEpoch = epoch
				member.State = CCMemberStateAnswered
			case "bridge-agent-end":*/
	}
}

func GetDirectorySipRegs() {
	if esl == nil || !esl.Connected() {
		return
	}
	if !altData.IsDirectoryEnabled() {
		return
	}
	asJson, err := esl.SendApiCmd("show registrations as json")
	if err != nil {
		return
	}
	regs := regsAsJson{}
	err = json.Unmarshal([]byte(asJson), &regs)
	if err != nil {
		return
	}
	domains := GetDomainsWithNames()
	directoryCache := cache.GetDirectoryCache()
	for _, reg := range regs.Registrations {
		domain := domains[reg.Realm]
		if domain == nil {
			log.Println("no domain")
			continue
		}
		user := GetDomainUserByName(reg.RegUser, domain)
		if user == nil {
			log.Println("no user")
			continue
		}
		cDomain := directoryCache.DomainCache.GetById(domain.Id)
		if cDomain == nil {
			cDomain = directoryCache.DomainCache.SetByData(domain.Id, domain.Name, 0)
		}
		cUser := directoryCache.UserCache.GetById(user.Id)
		if cUser == nil {
			cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
		}
		cDomain.SipRegsCounter++
		cUser.SipRegister = true
		//domain.SipRegsCounter++
		//user.SipRegister = true
	}
}

func GetChannels() {
	if esl == nil || !esl.Connected() {
		return
	}
	asJson, err := esl.SendApiCmd("show channels as json")
	if err != nil {
		log.Println("cant get channels", err)
		return
	}
	channels := channelsAsJson{}
	err = json.Unmarshal([]byte(asJson), &channels)
	if err != nil {
		log.Println("can't marshal channels", err)
		return
	}
	cCache := pbxcache.GetChannelsCache()
	for _, mChannel := range channels.Channels {
		channel := mChannel
		cCache.Set(&channel)

		possibleUserName := channel.PresenceId
		if possibleUserName == "" {
			possibleUserName = channel.Name
		}
		user := getUserByPresence(possibleUserName)
		if user == nil {
			continue
		}
		directoryCache := cache.GetDirectoryCache()
		cUser := directoryCache.UserCache.GetById(user.Id)
		if cUser == nil {
			cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
		}

		cUser.InCall = true
		intDate, err := strconv.ParseInt(channel.CreatedEpoch, 10, 64)
		if err == nil {
			cUser.CallDate = intDate
		}
		if channel.Callstate == ValueChannelStateActive {
			cUser.LastUuid = channel.Uuid
			cUser.Talking = true
		}
	}
	cCache.Total, cCache.Answered = cCache.GetLength()
}

func GetXMLSofia() *XmlSofia {
	asXML, err := esl.SendApiCmd("sofia xmlstatus")
	if err != nil {
		//log.Println(err.Error())
		return nil
	}
	statuses := XmlSofia{}
	reader := strings.NewReader(asXML)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&statuses)
	if err != nil {
		return nil
	}
	return &statuses
}

func GetSofiaProfilesStatuses() map[int64]*altStruct.ConfigSofiaProfile {
	profiles := map[int64]*altStruct.ConfigSofiaProfile{}
	if esl == nil || !esl.Connected() {
		return profiles
	}

	if !altData.IsSofiaExists() {
		return profiles
	}

	statuses := GetXMLSofia()
	if statuses == nil {
		return profiles
	}
	var newUris []string
	for _, xmlProfile := range statuses.Profiles {
		profile := altData.GetSofiaProfileByName(xmlProfile.Name)
		if profile == nil {
			continue
		}
		if xmlProfile.Data != "" && (strings.Contains(xmlProfile.Data, "ws") || strings.Contains(xmlProfile.Data, "WS")) {
			r := regexp.MustCompile(`^.*@(.+);.*=(.+)$`)
			res := r.FindStringSubmatch(xmlProfile.Data)
			if len(res) == 3 && res[1] != "" && res[2] != "" {
				newUris = append(newUris, res[2]+"://"+res[1])
			}
		}
		if profile.Started {
			continue
		}
		profile.Started = true
		profile.State = xmlProfile.State
		profile.Uri = xmlProfile.Data
		profiles[profile.Id] = profile
	}
	webcache.GetWebMetaData().SetWssUris(newUris)

	return profiles
}

func GetSofiaGatewaysStatuses() map[int64]*altStruct.ConfigSofiaProfileGateway {
	gateways := map[int64]*altStruct.ConfigSofiaProfileGateway{}
	if esl == nil || !esl.Connected() {
		return gateways
	}

	statuses := GetXMLSofia()
	if statuses == nil {
		return gateways
	}
	for _, xmlGateway := range statuses.Gateways {
		gateway := altData.GetSofiaProfileGateway(xmlGateway.Name)
		if gateway == nil {
			continue
		}
		gateway.Started = true
		gateway.State = xmlGateway.State

		gateways[gateway.Id] = gateway
	}

	return gateways
}

func updateStatusesForProfilesStruct() {

}

/*
func GetSofiaStatuses() {
	if esl == nil || !esl.Connected() {
		return
	}

	if !altData.IsSofiaExists() {
		return
	}
	asXML, err := esl.SendApiCmd("sofia xmlstatus")
	if err != nil {
		//log.Println(err.Error())
		return
	}
	statuses := XmlSofia{}
	reader := strings.NewReader(asXML)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&statuses)
	if err != nil {
		//log.Println(err.Error())
		return
	}
	var newUris []string
	for _, xmlProfile := range statuses.Profiles {
		profile := altData.GetSofiaProfileByName(xmlProfile.Name)
		if profile == nil {
			continue
		}
		if xmlProfile.Data != "" && (strings.Contains(xmlProfile.Data, "ws") || strings.Contains(xmlProfile.Data, "WS")) {
			r := regexp.MustCompile(`^.*@(.+);.*=(.+)$`)
			res := r.FindStringSubmatch(xmlProfile.Data)
			if len(res) == 3 && res[1] != "" && res[2] != "" {
				newUris = append(newUris, res[2]+"://"+res[1])
			}
		}
		if profile.Started {
			continue
		}
		profile.Started = true
		profile.State = xmlProfile.State
		profile.Uri = xmlProfile.Data
	}
	webcache.GetWebMetaData().SetWssUris(newUris)

	for _, xmlGateway := range statuses.Gateways {
		gateway := altData.GetSofiaProfileGateway(xmlGateway.Name)
		if gateway == nil {
			continue
		}
		gateway.Started = true
		gateway.State = xmlGateway.State
	}
}
*/

func GetVertoStatuses() {
	if esl == nil || !esl.Connected() {
		return
	}
	if !altData.IsVertoExists() {
		return
	}
	asXML, err := esl.SendApiCmd("verto xmlstatus")
	if err != nil {
		//log.Println(err.Error())
		return
	}
	statuses := XmlVerto{}
	reader := strings.NewReader(asXML)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&statuses)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var newUris []string
	for _, xmlProfile := range statuses.Profile {
		profile := altData.GetVertoProfileByName(xmlProfile.Name)
		if profile == nil {
			continue
		}
		if xmlProfile.Data != "" {
			newUris = append(newUris, strings.Replace(xmlProfile.Data, "s:", "s://", 1))
		}
	}

	webcache.GetWebMetaData().SetVertoWsUris(newUris)
	for _, xmlClient := range statuses.Clients {
		r := regexp.MustCompile(`^(.*)@(.+)$`)
		res := r.FindStringSubmatch(xmlClient.Name)
		if len(res) != 3 || res[1] == "" || res[2] == "" {
			continue
		}
		domain := GetDomainByName(res[2])
		if domain == nil {
			log.Println("no domain")
			continue
		}
		user := GetDomainUserByName(res[1], domain)
		if user == nil {
			log.Println("no user")
			continue
		}
		directoryCache := cache.GetDirectoryCache()
		cUser := directoryCache.UserCache.GetById(user.Id)
		if cUser == nil {
			cUser = directoryCache.UserCache.SetByData(user.Id, user.Name, false)
		}
		// domain.SipRegsCounter++
		cUser.VertoRegister = true
		//user.VertoRegister = true
	}
}

// TODO
func GetLoadedModules() {
	if esl == nil || !esl.Connected() {
		return
	}
	names := mainStruct.GetModulesNames()

	for _, name := range names {
		go SetModuleLoadedStatus(name)
	}
}

func SetModuleLoadedStatus(name string) {
	module, err := altData.GetModuleByName(name)
	if module == nil || err != nil {
		return
	}
	if name == mainStruct.ModAcl {
		module.Loaded = true
		return
	}

	out, err := esl.SendBgapiCmd(CommandModuleExists + " " + module.Module)
	if err != nil {
		return
	}

	module.Loaded = false

	val := <-out
	if strings.TrimSpace(val) == "true" {
		module.Loaded = true
	}
	pbxcache.ConfigurationCache.FillConfigurations(module)
}

func SendBgapiCmd(command string) (string, error) {
	if esl == nil || !esl.Connected() {
		return "", errors.New("no esl connection")
	}
	if command == "" {
		return "", errors.New("no command")
	}
	out, err := esl.SendBgapiCmd(command)
	if err != nil {
		return "", err
	}

	val := <-out
	if strings.Contains(val, "-ERR") {
		return val, errors.New(strings.TrimSpace(val))
	}

	return val, nil
}

func GetXMLDirectory() {
	if esl == nil || !esl.Connected() {
		return
	}
	scriptName, args, scriptPath := CheckLuaXMLHandlerDirectory()
	params := map[string]string{
		"Event-Calling-Function": "populate_database",
		"Event-Calling-File":     "mod_directory.c",
	}
	resp, ok := luaXMLRequest("directory", scriptName, args, scriptPath, params)
	if ok {
		luaXMLDirectoryImport(resp, scriptName, args, scriptPath)
		return
	}

	rawXML, err := esl.SendApiCmd("xml_locate directory")
	if err != nil {
		return
	}
	var directoryXML xmlStruct.Section
	err = xml.Unmarshal([]byte(rawXML), &directoryXML)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = initDirectoryDomain(directoryXML.Domain)
}

func CheckLuaXMLHandlerDirectory() (string, []string, string) {
	return CheckLuaXMLHandler("directory")
}

func CheckLuaXMLHandlerDialplan() (string, []string, string) {
	return CheckLuaXMLHandler("dialplan")
}

func CheckLuaXMLHandler(bind string) (string, []string, string) {
	configurationXML, err := GetXMLModuleConfigurationStruct(mainStruct.ConfLua)
	if err != nil || configurationXML == nil || configurationXML.Settings == nil {
		return "", []string{}, ""
	}

	var script []string
	var scriptPath string
	for _, param := range configurationXML.Settings.Param {
		switch param.Attrname {
		case "script-directory":
			scriptPath = ParseGlobalVars(param.Attrvalue)
		case "xml-handler-script":
			script = strings.Fields(param.Attrvalue)
		case "xml-handler-bindings":
			if !strings.Contains(strings.ToLower(param.Attrvalue), bind) {
				return "", []string{}, ""
			}
		}
	}

	if len(script) == 0 || script[0] == "" {
		return "", []string{}, ""
	}
	if scriptPath == "" {
		scriptPath = ParseGlobalVars("script_dir")
	}

	if scriptPath == "" {
		scriptPath = ParseGlobalVars("scripts")
	}

	return script[0], script[len(script)-1:], scriptPath
}

func generateLuaCommand(handlerCase, scriptName, scriptPath string, args []string, params map[string]string) string {
	command := "lua " +
		"~XML_REQUEST={" +
		"section=\"" + handlerCase + "\"," +
		"tag_name=\"\"," +
		"context=\"\"," +
		"key_name=\"\"," +
		"key_value=\"\"" +
		"};Params={};Params.__index=Params;function Params:getHeader (name) return self[name];end;function Params:serialize() return;end;function Params:setHeader(name,value) self[name]=value;end;params={};setmetatable(params,Params);"

	for name, value := range params {
		command += "params:setHeader(\"" + name + "\",\"" + value + "\");"
	}
	command += "argv={"
	for _, arg := range args {
		command += "\"" + arg + "\""
	}
	command += "};"
	command += "loadfile(\"" + scriptPath + "/" + scriptName + "\")();" +
		"stream:write(XML_STRING);"

	return command
}

func luaXMLRequest(handlerCase, scriptName string, args []string, scriptPath string, params map[string]string) (string, bool) {
	if scriptName == "" {
		return "", false
	}

	if scriptPath == "" {
		scriptPath = "/usr/share/freeswitch/scripts"
	} else {
		scriptPath = strings.TrimRight(scriptPath, "/?.lua")
	}

	command := generateLuaCommand(handlerCase, scriptName, scriptPath, args, params)

	rawXML, err := esl.SendApiCmd(command)
	if err != nil {
		return "", false
	}

	return rawXML, true
}

func luaXMLDirectoryImport(rawXML string, scriptName string, args []string, scriptPath string) {
	rawXML = fixBrokenXML(rawXML)
	var directoryXML xmlStruct.Document
	err := xml.Unmarshal([]byte(rawXML), &directoryXML)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for s := 0; s < len(directoryXML.Section); s++ {
		for d := 0; d < len(directoryXML.Section[s].Domain); d++ {
			firstUser := true
			if directoryXML.Section[s].Domain[d].Groups == nil {
				directoryXML.Section[s].Domain[d].Groups = &xmlStruct.Groups{}
			}
			for g := 0; g < len(directoryXML.Section[s].Domain[d].Groups.Group); g++ {
				for u := 0; u < len(directoryXML.Section[s].Domain[d].Groups.Group[g].Users.User); u++ {
					params := map[string]string{
						"domain": directoryXML.Section[s].Domain[d].Attrname,
						"user":   directoryXML.Section[s].Domain[d].Groups.Group[g].Users.User[u].Attrid,
					}
					res, ok := luaXMLRequest("directory", scriptName, args, scriptPath, params)
					if !ok {
						continue
					}
					res = fixBrokenXML(res)
					var directoryUserXML xmlStruct.Document
					err := xml.Unmarshal([]byte(res), &directoryUserXML)
					if err != nil {
						log.Print(err.Error())
						directoryXML.Section[s].Domain[d].Groups.Group[g].Users.User[u] = nil
						continue
					}
					usr := getFirstUser(directoryUserXML)
					directoryXML.Section[s].Domain[d].Groups.Group[g].Users.User[u] = usr
					if firstUser {
						if len(directoryUserXML.Section) == 0 {
							continue
						}
						if len(directoryUserXML.Section[0].Domain) == 0 {
							continue
						}
						userDomain := directoryUserXML.Section[0].Domain[0]
						directoryXML.Section[s].Domain[d].Params = userDomain.Params
						directoryXML.Section[s].Domain[d].Variables = userDomain.Variables
						firstUser = false
					}
				}
			}
			if directoryXML.Section[s].Domain[d].Users == nil {
				continue
			}
			for u := 0; u < len(directoryXML.Section[s].Domain[d].Users.User); u++ {
				params := map[string]string{
					"domain": directoryXML.Section[s].Domain[d].Attrname,
					"user":   directoryXML.Section[s].Domain[d].Users.User[u].Attrid,
				}
				res, ok := luaXMLRequest("directory", scriptName, args, scriptPath, params)
				if !ok {
					continue
				}
				res = fixBrokenXML(res)
				var directoryUserXML xmlStruct.Document
				err := xml.Unmarshal([]byte(res), &directoryUserXML)
				if err != nil {
					directoryXML.Section[s].Domain[d].Users.User[u] = nil
					log.Print(err.Error())
					continue
				}
				usr := getFirstUser(directoryUserXML)
				directoryXML.Section[s].Domain[d].Users.User[u] = usr
				if firstUser {
					if len(directoryUserXML.Section) == 0 {
						continue
					}
					if len(directoryUserXML.Section[0].Domain) == 0 {
						continue
					}
					userDomain := directoryUserXML.Section[0].Domain[0]
					directoryXML.Section[s].Domain[d].Params = userDomain.Params
					directoryXML.Section[s].Domain[d].Variables = userDomain.Variables
					firstUser = false
				}
			}
		}
		_ = initDirectoryDomain(directoryXML.Section[s].Domain)
	}

	//lua ~XML_REQUEST={section="directory",tag_name="",context="",key_name="",key_value=""};Params={};Params.__index=Params;function Params:getHeader (name) return self[name];end;function Params:serialize() return;end;function Params:setHeader(name,value) self[name]=value;end;params={};setmetatable(params,Params);params:setHeader("Event-Calling-Function","populate_database");params:setHeader("Event-Calling-File","mod_directory.c");argv={"xml_handler"};loadfile("/usr/share/freeswitch/scripts/app.lua")();stream:write(XML_STRING);
}

func fixBrokenXML(in string) string {
	//TODO need better regex
	re, _ := regexp.Compile("=\"[^\"]+\"")
	in = re.ReplaceAllStringFunc(in, func(s string) string {
		s = strings.Replace(s, "<", "&lt;", -1)
		return strings.Replace(s, ">", "&gt;", -1)
	})

	return strings.Replace(in, "type=>", "type=\"\">", -1)
}

func getFirstUser(directoryUserXML xmlStruct.Document) *xmlStruct.User {
	if len(directoryUserXML.Section) == 0 {
		return nil
	}
	if len(directoryUserXML.Section[0].Domain) == 0 {
		return nil
	}
	if directoryUserXML.Section[0].Domain[0].Users != nil && len(directoryUserXML.Section[0].Domain[0].Users.User) > 0 {
		return directoryUserXML.Section[0].Domain[0].Users.User[0]
	}
	if len(directoryUserXML.Section[0].Domain[0].Groups.Group) == 0 {
		return nil
	}

	if directoryUserXML.Section[0].Domain[0].Groups == nil ||
		directoryUserXML.Section[0].Domain[0].Groups.Group[0].Users == nil ||
		len(directoryUserXML.Section[0].Domain[0].Groups.Group[0].Users.User) == 0 {
		return nil
	}

	return directoryUserXML.Section[0].Domain[0].Groups.Group[0].Users.User[0]
}

func initDirectoryDomain(domains []*xmlStruct.Domain) error {
	//fmt.Println(rawXML)
	//fmt.Printf("%+v", directoryXML)

	var errs []string
	for _, eslDomain := range domains {
		domainId, err := altData.SetDirectoryDomain(eslDomain.Attrname)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		if domainId == 0 {
			errs = append(errs, "no domain")
			continue
		}
		if eslDomain.Params != nil {
			for _, eslDomainParam := range eslDomain.Params.Param {
				_, err := altData.SetDirectoryDomainParameter(domainId, eslDomainParam.Attrname, eslDomainParam.Attrvalue)
				if err != nil {
					errs = append(errs, err.Error())
					continue
				}
			}
		}
		if eslDomain.Variables != nil {
			for _, eslDomainVar := range eslDomain.Variables.Variable {
				_, err := altData.SetDirectoryDomainVariable(domainId, eslDomainVar.Attrname, eslDomainVar.Attrvalue)
				if err != nil {
					errs = append(errs, err.Error())
					continue
				}
			}
		}
		if eslDomain.Users != nil {
			for _, eslDomainUser := range eslDomain.Users.User {
				err = UserList(eslDomainUser, domainId)
				if err != nil {
					errs = append(errs, err.Error())
				}
			}
		}
		if eslDomain.Groups != nil {
			for _, eslDomainGroup := range eslDomain.Groups.Group {
				groupId, err := altData.SetDirectoryDomainGroup(domainId, eslDomainGroup.Attrname)
				if err != nil || groupId == 0 {
					continue
				}
				if eslDomainGroup.Users != nil {
					for _, groupUser := range eslDomainGroup.Users.User {
						if groupUser.Attrtype == "" {
							err = UserList(groupUser, domainId)
							if err != nil {
								errs = append(errs, err.Error())
							}
						}
						_, err := altData.SetDirectoryDomainGroupUser(groupId, domainId, groupUser.Attrid)
						if err != nil {
							errs = append(errs, err.Error())
							continue
						}
					}
				}
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}
	return nil
}

func UserList(eslDomainUser *xmlStruct.User, domainId int64) error {
	userId, err := altData.SetDirectoryDomainUser(domainId, eslDomainUser.Attrid, "", eslDomainUser.Attrcidr, eslDomainUser.Attrnumberalias)
	if err != nil || userId == 0 {
		return err
	}
	var errs []string
	if eslDomainUser.Params != nil {
		for _, eslDomainUserParams := range eslDomainUser.Params.Param {
			_, err := altData.SetDirectoryDomainUserParameter(userId, eslDomainUserParams.Attrname, eslDomainUserParams.Attrvalue)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
		}
	}
	if eslDomainUser.Variables != nil {
		for _, eslDomainUserVars := range eslDomainUser.Variables.Variable {
			_, err := altData.SetDirectoryDomainUserVariable(userId, eslDomainUserVars.Attrname, eslDomainUserVars.Attrvalue)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
		}
	}
	if eslDomainUser.Gateways != nil {
		for _, eslDomainUserGateway := range eslDomainUser.Gateways.Gateway {
			gateway, err := altData.SetDirectoryDomainUserGateway(userId, eslDomainUserGateway.Attrname)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			if gateway == 0 {
				continue
			}
			if eslDomainUserGateway != nil {
				for _, eslDomainUserGatewaysParam := range eslDomainUserGateway.Param {
					_, err := altData.SetDirectoryDomainUserGatewayParameter(gateway, eslDomainUserGatewaysParam.Attrname, eslDomainUserGatewaysParam.Attrvalue)
					if err != nil {
						errs = append(errs, err.Error())
						continue
					}
				}
				if eslDomainUserGateway.Variables != nil && eslDomainUserGateway.Variables.Variable != nil {
					for _, eslDomainUserGatewaysVar := range eslDomainUserGateway.Variables.Variable {
						_, err := altData.SetDirectoryDomainUserGatewayVariable(gateway, eslDomainUserGatewaysVar.Attrname, eslDomainUserGatewaysVar.Attrvalue, eslDomainUserGatewaysVar.Attrdirection)
						if err != nil {
							errs = append(errs, err.Error())
							continue
						}
					}
				}
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}
	return nil
}

func ParseDirectoryXML(rawXML string) error {
	var documentXML xmlStruct.Document
	err := xml.Unmarshal([]byte(rawXML), &documentXML)
	if err != nil {
		err = xml.Unmarshal([]byte(rawXML), &documentXML.Section)
		if err != nil {
			section := &xmlStruct.Section{Attrname: "directory"}
			documentXML.Section = append(documentXML.Section, section)
			err = xml.Unmarshal([]byte(rawXML), &section.Domain)
		}
	}
	if err != nil {
		return errors.New("cant parse")
	}

	for _, section := range documentXML.Section {
		if section.Attrname == "directory" {
			if section.Domain == nil {
				return errors.New("domain not recognized")
			}
			return initDirectoryDomain(section.Domain)
		}
	}
	return errors.New("domains not found")
}

func ParseDirectoryUsersXML(domainId int64, rawXML string) error {
	var usersDomainXML xmlStruct.Domain
	err := xml.Unmarshal([]byte(rawXML), &usersDomainXML)
	if err != nil {
		err = xml.Unmarshal([]byte(rawXML), &usersDomainXML.Users)
		if err != nil {
			err = xml.Unmarshal([]byte(rawXML), &usersDomainXML.Users.User)
		}
	}
	if err != nil {
		return errors.New("cant parse")
	}
	if usersDomainXML.Users.User == nil {
		return errors.New("users not recognized")
	}

	var errs []string
	for _, eslDomainUser := range usersDomainXML.Users.User {
		err = UserList(eslDomainUser, domainId)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}
	return nil
}

func GetXMLConfigurations() {
	if esl == nil || !esl.Connected() {
		return
	}
	rawXML, err := esl.SendApiCmd("xml_locate configuration")
	if err != nil {
		return
	}
	var configurationXML xmlStruct.Section
	err = xml.Unmarshal([]byte(rawXML), &configurationXML)

	if err != nil {
		fmt.Println(err)
		return
	}

	if configurationXML.Configuration != nil {
		for _, eslConfig := range configurationXML.Configuration {
			_ = InitConfigModule(eslConfig)
		}
	}
}

func GetXMLModuleConfiguration(confName string) error {
	configurationXML, err := GetXMLModuleConfigurationStruct(confName)
	if err != nil {
		return err
	}
	err = InitConfigModule(configurationXML)

	return err
}

func GetXMLModuleConfigurationStruct(confName string) (*xmlStruct.Configuration, error) {
	if esl == nil || confName == "" {
		return nil, errors.New("no esl connection")
	}
	if !mainStruct.IsConfName(confName) {
		return nil, errors.New("module name not found")
	}
	if confName == mainStruct.ConfPostLoadSwitch {
		confName = mainStruct.ConfSwitch
	}
	rawXML, err := esl.SendApiCmd("xml_locate configuration configuration name " + confName)
	if err != nil {
		return nil, errors.New("can't send command")
	}
	var configurationXML *xmlStruct.Configuration
	err = xml.Unmarshal([]byte(rawXML), &configurationXML)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can't read config")
	}

	if configurationXML.Attrname != confName {
		return nil, errors.New("can't read config")
	}

	return configurationXML, nil
}

func ParseConfigXML(rawXML string) error {
	var sectionXML xmlStruct.Section
	err := xml.Unmarshal([]byte(rawXML), &sectionXML)
	if err != nil {
		err = xml.Unmarshal([]byte(rawXML), &sectionXML.Configuration)
	}
	if err != nil {
		return errors.New("cant parse")
	}
	if sectionXML.Configuration == nil {
		return errors.New("config not recognized")
	}
	var errs []string
	for _, eslConfig := range sectionXML.Configuration {
		err = InitConfigModule(eslConfig)
		if err != nil {
			errs = append(errs, eslConfig.Attrname+": "+err.Error())
			err = nil
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}
	return nil
}

func InitConfigModule(conf *xmlStruct.Configuration) error {
	var err error
	switch conf.Attrname {
	case mainStruct.ConfSwitch:
		err = setConfigPostSwitch(conf)
	case mainStruct.ConfConferenceLayouts:
		err = setConfigConferenceLayouts(conf)
	case mainStruct.ConfAcl:
		err = setConfigAcl(conf)
		SetModuleLoadedStatus(mainStruct.ModAcl)
	case mainStruct.ConfCallcenter:
		err = setConfigCallcenter(conf)
		SetModuleLoadedStatus(mainStruct.ModCallcenter)
	case mainStruct.ConfCdrPgCsv:
		err = setConfigCdrPgCsv(conf)
		SetModuleLoadedStatus(mainStruct.ModCdrPgCsv)
	case mainStruct.ConfOdbcCdr:
		err = setConfigOdbcCdr(conf)
		SetModuleLoadedStatus(mainStruct.ModOdbcCdr)
	case mainStruct.ConfSofia:
		err = setConfigSofia(conf)
		SetModuleLoadedStatus(mainStruct.ModSofia)
		//GetSofiaStatuses()
	case mainStruct.ConfVerto:
		err = setConfigVerto(conf)
		SetModuleLoadedStatus(mainStruct.ModVerto)
	case mainStruct.ConfLcr:
		err = setConfigLcr(conf)
		SetModuleLoadedStatus(mainStruct.ModLcr)
	case mainStruct.ConfShout:
		err = setConfigShout(conf)
		SetModuleLoadedStatus(mainStruct.ModShout)
	case mainStruct.ConfRedis:
		err = setConfigRedis(conf)
		SetModuleLoadedStatus(mainStruct.ModRedis)
	case mainStruct.ConfNibblebill:
		err = setConfigNibblebill(conf)
		SetModuleLoadedStatus(mainStruct.ModNibblebill)
	case mainStruct.ConfDb:
		err = setConfigDb(conf)
		SetModuleLoadedStatus(mainStruct.ModDb)
	case mainStruct.ConfMemcache:
		err = setConfigMemcache(conf)
		SetModuleLoadedStatus(mainStruct.ModMemcache)
	case mainStruct.ConfAvmd:
		err = setConfigAvmd(conf)
		SetModuleLoadedStatus(mainStruct.ModAvmd)
	case mainStruct.ConfTtsCommandline:
		err = setConfigTtsCommandline(conf)
		SetModuleLoadedStatus(mainStruct.ModTtsCommandline)
	case mainStruct.ConfCdrMongodb:
		err = setConfigCdrMongodb(conf)
		SetModuleLoadedStatus(mainStruct.ModCdrMongodb)
	case mainStruct.ConfHttpCache:
		err = setConfigHttpCache(conf)
		SetModuleLoadedStatus(mainStruct.ModHttapiCache)
	case mainStruct.ConfOpus:
		err = setConfigOpus(conf)
		SetModuleLoadedStatus(mainStruct.ModOpus)
	case mainStruct.ConfPython:
		err = setConfigPython(conf)
		SetModuleLoadedStatus(mainStruct.ModPython)
	case mainStruct.ConfAlsa:
		err = setConfigAlsa(conf)
		SetModuleLoadedStatus(mainStruct.ModAlsa)
	case mainStruct.ConfAmr:
		err = setConfigAmr(conf)
		SetModuleLoadedStatus(mainStruct.ModAmr)
	case mainStruct.ConfAmrwb:
		err = setConfigAmrwb(conf)
		SetModuleLoadedStatus(mainStruct.ModAmrwb)
	case mainStruct.ConfCepstral:
		err = setConfigCepstral(conf)
		SetModuleLoadedStatus(mainStruct.ModCepstral)
	case mainStruct.ConfCidlookup:
		err = setConfigCidlookup(conf)
		SetModuleLoadedStatus(mainStruct.ModCidlookup)
	case mainStruct.ConfCurl:
		err = setConfigCurl(conf)
		SetModuleLoadedStatus(mainStruct.ModCurl)
	case mainStruct.ConfDialplanDirectory:
		err = setConfigDialplanDirectory(conf)
		SetModuleLoadedStatus(mainStruct.ModDialplanDirectory)
	case mainStruct.ConfEasyroute:
		err = setConfigEasyroute(conf)
		SetModuleLoadedStatus(mainStruct.ModEasyroute)
	case mainStruct.ConfErlangEvent:
		err = setConfigErlangEvent(conf)
		SetModuleLoadedStatus(mainStruct.ModErlangEvent)
	case mainStruct.ConfEventMulticast:
		err = setConfigEventMulticast(conf)
		SetModuleLoadedStatus(mainStruct.ModEventMulticast)
	case mainStruct.ConfFax:
		err = setConfigFax(conf)
		SetModuleLoadedStatus(mainStruct.ModFax)
	case mainStruct.ConfLua:
		err = setConfigLua(conf)
		SetModuleLoadedStatus(mainStruct.ModLua)
	case mainStruct.ConfMongo:
		err = setConfigMongo(conf)
		SetModuleLoadedStatus(mainStruct.ModMongo)
	case mainStruct.ConfMsrp:
		err = setConfigMsrp(conf)
		SetModuleLoadedStatus(mainStruct.ModMsrp)
	case mainStruct.ConfOreka:
		err = setConfigOreka(conf)
		SetModuleLoadedStatus(mainStruct.ModOreka)
	case mainStruct.ConfPerl:
		err = setConfigPerl(conf)
		SetModuleLoadedStatus(mainStruct.ModPerl)
	case mainStruct.ConfPocketsphinx:
		err = setConfigPocketsphinx(conf)
		SetModuleLoadedStatus(mainStruct.ModPocketsphinx)
	case mainStruct.ConfSangomaCodec:
		err = setConfigSangomaCodec(conf)
		SetModuleLoadedStatus(mainStruct.ModSangomaCodec)
	case mainStruct.ConfSndfile:
		err = setConfigSndfile(conf)
		SetModuleLoadedStatus(mainStruct.ModSndfile)
	case mainStruct.ConfXmlCdr:
		err = setConfigXmlCdr(conf)
		SetModuleLoadedStatus(mainStruct.ModXmlCdr)
	case mainStruct.ConfXmlRpc:
		err = setConfigXmlRpc(conf)
		SetModuleLoadedStatus(mainStruct.ModXmlRpc)
	case mainStruct.ConfZeroconf:
		err = setConfigZeroconf(conf)
		SetModuleLoadedStatus(mainStruct.ModZeroconf)
	case mainStruct.ConfDistributor:
		err = setConfigDistributor(conf)
		SetModuleLoadedStatus(mainStruct.ModDistributor)
	case mainStruct.ConfDirectory:
		err = setConfigDirectory(conf)
		SetModuleLoadedStatus(mainStruct.ModDirectory)
	case mainStruct.ConfFifo:
		err = setConfigFifo(conf)
		SetModuleLoadedStatus(mainStruct.ModFifo)
	case mainStruct.ConfOpal:
		err = setConfigOpal(conf)
		SetModuleLoadedStatus(mainStruct.ModOpal)
	case mainStruct.ConfOsp:
		err = setConfigOsp(conf)
		SetModuleLoadedStatus(mainStruct.ModOsp)
	case mainStruct.ConfUnicall:
		err = setConfigUnicall(conf)
		SetModuleLoadedStatus(mainStruct.ModUnicall)
	case mainStruct.ConfConference:
		err = setConfigConference(conf)
		SetModuleLoadedStatus(mainStruct.ModConference)
	case mainStruct.ConfPostLoadModules:
		err = setConfigPostLoadModules(conf)
	case mainStruct.ConfVoicemail:
		err = setConfigVoicemail(conf)
	default:
		return errors.New("not suitable module: " + conf.Attrname)
	}
	return err
}

func setConfigAcl(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}
	mod, err := altData.SetConfAcl()
	if err != nil {
		return err
	}
	log.Printf("%v", mod)
	log.Printf("%v", conf)
	if conf.NetworkLists == nil || conf.NetworkLists.List == nil {
		return errors.New("empty config object")
	}
	for _, eslList := range conf.NetworkLists.List {
		listId, err := altData.SetConfAclList(mod, eslList.Attrname, eslList.Attrdefault)
		if err != nil {
			continue
		}
		if eslList.Node == nil {
			continue
		}
		for _, node := range eslList.Node {
			if node.Attrcidr == "" && node.Attrhost != "" {
				suffix := "/32"
				switch node.Attrmask {
				case "255.0.0.0":
					suffix = "/8"
				case "255.128.0.0":
					suffix = "/9"
				case "255.192.0.0":
					suffix = "/10"
				case "255.224.0.0":
					suffix = "/11"
				case "255.240.0.0":
					suffix = "/12"
				case "255.248.0.0":
					suffix = "/13"
				case "255.252.0.0":
					suffix = "/14"
				case "255.254.0.0":
					suffix = "/15"
				case "255.255.0.0":
					suffix = "/16"
				case "255.255.128.0":
					suffix = "/17"
				case "255.255.192.0":
					suffix = "/18"
				case "255.255.224.0":
					suffix = "/19"
				case "255.255.240.0":
					suffix = "/20"
				case "255.255.248.0":
					suffix = "/21"
				case "255.255.252.0":
					suffix = "/22"
				case "255.255.254.0":
					suffix = "/23"
				case "255.255.255.0":
					suffix = "/24"
				case "255.255.255.128":
					suffix = "/25"
				case "255.255.255.192":
					suffix = "/26"
				case "255.255.255.224":
					suffix = "/27"
				case "255.255.255.240":
					suffix = "/28"
				case "255.255.255.248":
					suffix = "/29"
				case "255.255.255.252":
					suffix = "/30"
				case "255.255.255.254":
					suffix = "/31"
				}
				node.Attrcidr = node.Attrhost
				node.Attrcidr += suffix
			}
			_, err := altData.SetConfAclNode(listId, node.Attrtype, node.Attrcidr, node.Attrdomain)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigCallcenter(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}

	mod, err := altData.SetConfCallcenter()
	if err != nil {
		return err
	}

	if conf.Settings != nil {
		var dbSwitched bool
		ourDB := fmt.Sprintf("pgsql://host=%s port=%d dbname=%s user=%s password=%s application_name=%s",
			cfg.CustomPbx.Db.Host, cfg.CustomPbx.Db.Port, cfg.CustomPbx.Db.Name, cfg.CustomPbx.Db.User, cfg.CustomPbx.Db.Pass, cfg.AppName)

		for _, param := range conf.Settings.Param {
			if param.Attrname == "odbc-dsn" {
				param.Attrvalue = ourDB
				dbSwitched = true
			}

			_, err := altData.SetConfCallcenterSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
		if !dbSwitched {
			_, _ = altData.SetConfCallcenterSetting(mod, "odbc-dsn", ourDB)
		}
	}

	if conf.Queues != nil {
		for _, xmlQueue := range conf.Queues.Queue {
			if xmlQueue.Param == nil {
				continue
			}
			queue, err := altData.SetConfCallcenterQueue(mod, xmlQueue.Attrname)
			if err != nil {
				continue
			}
			for _, param := range xmlQueue.Param {
				_, err := altData.SetConfCallcenterQueueParam(queue, param.Attrname, param.Attrvalue)
				if err != nil {
					continue
				}
			}
		}
	}
	err = GetCallcenterAgents()
	if err != nil {
		return err
	}
	err = GetCallcenterTiers()
	if err != nil {
		return err
	}
	return GetCallcenterMembers()
}

func GetCallcenterAgents() error {
	if esl == nil || !esl.Connected() {
		log.Println("no ECL connection")
		return errors.New("no ECL connection")
	}
	agents, err := esl.SendApiCmd(`json {"command": "callcenter_config", "data": {"arguments":"agent list"}}`)
	if err != nil {
		log.Println(err)
		return err
	}

	var agentsResponse mainStruct.CallcenterAgentsJSONResponse
	err = json.Unmarshal([]byte(agents), &agentsResponse)
	if err != nil {
		log.Println(err)
		return err
	}
	if agentsResponse.Status != "success" {
		log.Println("response status: ", agentsResponse.Status)
		return errors.New("agents response status: " + agentsResponse.Status)
	}
	for _, agent := range agentsResponse.Response {
		_, err := altData.SetConfCallcenterAgent(
			agent.Name,
			agent.Type,
			agent.System,
			agent.InstanceId,
			agent.Uuid,
			agent.Contact,
			agent.Status,
			agent.State,
			localParseInt(agent.MaxNoAnswer),
			localParseInt(agent.WrapUpTime),
			localParseInt(agent.RejectDelayTime),
			localParseInt(agent.BusyDelayTime),
			localParseInt(agent.NoAnswerDelayTime),
			localParseInt(agent.LastBridgeStart),
			localParseInt(agent.LastBridgeEnd),
			localParseInt(agent.LastOfferedCall),
			localParseInt(agent.LastStatusChange),
			localParseInt(agent.NoAnswerCount),
			localParseInt(agent.CallsAnswered),
			localParseInt(agent.TalkTime),
			localParseInt(agent.ReadyTime))
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}

func GetCallcenterTiers() error {
	tiers, err := esl.SendApiCmd(`json {"command": "callcenter_config","data": {"arguments":"tier list"}}"`)
	if err != nil {
		log.Println(err)
		return err
	}
	var tiersResponse mainStruct.CallcenterTiersJSONResponse
	err = json.Unmarshal([]byte(tiers), &tiersResponse)
	if err != nil {
		log.Println(err)
		return err
	}
	if tiersResponse.Status != "success" {
		log.Println("tiers response status: ", tiersResponse.Status)
		return errors.New("tiers response status: " + tiersResponse.Status)
	}
	for _, tier := range tiersResponse.Response {
		_, err := altData.SetConfCallcenterTier(tier.Queue, tier.Agent, tier.State, localParseInt(tier.Position), localParseInt(tier.Level))
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}

func GetCallcenterMembers() error {
	members, err := esl.SendApiCmd(`json {"command": "callcenter_config","data": {"arguments":"member list"}}"`)
	if err != nil {
		log.Println(err)
		return err
	}
	var memberResponse mainStruct.CallcenterMembersJSONResponse
	err = json.Unmarshal([]byte(members), &memberResponse)
	if err != nil {
		log.Println(err)
		return err
	}
	if memberResponse.Status != "success" {
		return errors.New("members response status: " + memberResponse.Status)
	}
	for _, member := range memberResponse.Response {
		_, err := altData.SetConfCallcenterMember(
			member.Uuid, member.State, member.Queue, member.InstanceId, member.AbandonedEpoch, member.BaseScore,
			member.BridgeEpoch, member.CidName, member.CidNumber, member.JoinedEpoch, member.RejoinedEpoch,
			member.ServingAgent, member.ServingSystem, member.SessionUuid, member.SkillScore, member.SystemEpoch,
		)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}

func setConfigCdrPgCsv(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}
	mod, err := altData.SetConfCdrPgCsv()
	if err != nil {
		return err
	}

	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfCdrPgCsvSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}

	if conf.Schema != nil {
		for _, field := range conf.Schema.Field {
			_, err := altData.SetConfCdrPgCsvSchemaField(mod, field.Attrvar, field.Attrcolumn)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigOdbcCdr(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}
	mod, err := altData.SetConfOdbcCdr()
	if err != nil {
		return err
	}

	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfOdbcCdrSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}

	if conf.Tables != nil {
		for _, XMLTable := range conf.Tables.Table {
			tableId, err := altData.SetConfOdbcCdrTable(mod, XMLTable.Attrname, XMLTable.AttrlogLeg)
			if err != nil {
				continue
			}
			for _, field := range XMLTable.Fields {
				_, err := altData.SetConfOdbcCdrTableField(tableId, field.Attrname, field.AttrchanVarName)
				if err != nil {
					continue
				}

			}
		}
	}
	return nil
}

func setConfigDistributor(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}
	mod, err := altData.SetConfDistributor()
	if err != nil {
		return err
	}
	if conf.Lists == nil || conf.Lists.List == nil {
		return errors.New("empty config object")
	}
	for _, eslList := range conf.Lists.List {
		listId, err := altData.SetConfDistributorList(mod, eslList.Attrname)
		if err != nil {
			continue
		}
		if eslList.Node == nil {
			continue
		}
		for _, node := range eslList.Node {
			_, err := altData.SetConfDistributorNode(listId, node.Attrname, node.Attrweight)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigConference(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}
	mod, err := altData.SetConfConference()
	if err != nil {
		return err
	}
	if conf.Advertise != nil {
		for _, room := range conf.Advertise.Room {
			_, err := altData.SetConfConferenceAdvertise(mod, room.Attrname, room.Attrstatus)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}

	if conf.ChatPermissions != nil {
		for _, eslProfile := range conf.ChatPermissions.Profile {
			profileId, err := altData.SetConfConferenceChatPermissionsProfile(mod, eslProfile.Attrname)
			if err != nil {
				continue
			}
			for _, user := range eslProfile.User {
				_, err := altData.SetConfConferenceChatPermissionsUser(profileId, user.Attrname, user.Attrcommands)
				if err != nil {
					continue
				}
			}
		}
	}

	if conf.CallerControls != nil {
		for _, eslGroup := range conf.CallerControls.Group {
			groupId, err := altData.SetConfConferenceCallerControlsGroup(mod, eslGroup.Attrname)
			if err != nil {
				continue
			}
			for _, control := range eslGroup.Control {
				_, err := altData.SetConfConferenceCallerControlsGroupControl(groupId, control.Attraction, control.Attrdigits)
				if err != nil {
					continue
				}
			}
		}
	}

	if conf.Profiles != nil {
		for _, eslProfile := range conf.Profiles.Profile {
			profileId, err := altData.SetConfConferenceProfile(mod, eslProfile.Attrname)
			if err != nil {
				continue
			}
			for _, param := range eslProfile.Param {
				_, err := altData.SetConfConferenceProfileParam(profileId, param.Attrname, param.Attrvalue)
				if err != nil {
					continue
				}
			}
		}
	}
	return nil
}

func setConfigConferenceLayouts(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}
	mod, err := altData.SetConfConferenceLayouts()
	if err != nil {
		return err
	}
	if conf.LayoutSettings != nil {
		return nil
	}
	if conf.LayoutSettings.Groups != nil {
		for _, eslGroup := range conf.LayoutSettings.Groups.Group {
			group, err := altData.SetConfConferenceLayoutsGroups(mod, eslGroup.Attrname)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, eslLayout := range eslGroup.Layout {
				_, err := altData.SetConfConferenceLayoutsGroupLayout(group, eslLayout.Body)
				if err != nil {
					continue
				}
			}
		}
	}

	if conf.LayoutSettings.Layouts != nil {
		for _, eslLayout := range conf.LayoutSettings.Layouts.Layout {
			layout, err := altData.SetConfConferenceLayoutLayouts(mod, eslLayout.Attrname)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, eslImage := range eslLayout.Image {
				_, err := altData.SetConfConferenceLayoutLayoutsImage(
					layout,
					eslImage.Attrx,
					eslImage.Attry,
					eslImage.Attrhscale,
					eslImage.Attrfloor,
					eslImage.AttrfloorOnly,
					eslImage.Attrhscale,
					eslImage.Attroverlap,
					eslImage.Attrreservation_id,
					eslImage.Attrzoom,
				)
				if err != nil {
					continue
				}
			}
		}
	}

	return nil
}

/*
func setConfigEventSocket(conf *Configuration) {

		if conf.Settings != nil {
			return
		}
		for _, param := range conf.Settings.Param {
			_, err := pbxcash.SetConfigEvenSocketParam(param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}

		}
	}

func setConfigFormatCdr(conf *Configuration) {

		if conf.Profiles != nil {
			return
		}
		for _, profile := range conf.Profiles.Profile {
			_, err := pbxcash.SetConfigFormatCdrProfile(profile.Attrname)
			if err != nil {
				continue
			}
			for _, param := range profile.Settings.Param {
				_, err := pbxcash.SetConfigFormatCdrProfileParam(profile.Attrname, param.Attrname, param.Attrvalue)
				if err != nil {
					continue
				}
			}
		}
	}

func setConfigHttapi(conf *Configuration) {

		if conf.Settings != nil {
			for _, param := range conf.Settings.Param {
				_, err := pbxcash.SetHttapiParam(param.Attrname, param.Attrvalue)
				if err != nil {
					continue
				}
			}
		}

		if conf.Profiles != nil {
			for _, profile := range conf.Profiles.Profile {
				_, err := pbxcash.SetCHttapiProfile(profile.Attrname)
				if err != nil {
					continue
				}
				for _, confParam := range profile.Conference.Param {
					_, err := pbxcash.SetConfHttapiConferenceParam(profile.Attrname, confParam.Attrname, confParam.Attrname)
					if err != nil {
						continue
					}
				}
				for _, dialParam := range profile.Dial.Param {
					_, err := pbxcash.SetConfHttapiDialParam(profile.Attrname, dialParam.Attrname, dialParam.Attrname)
					if err != nil {
						continue
					}
				}
				for _, permission := range profile.Permissions.Permission {
					_, err := pbxcash.SetConfHttapiPermission(profile.Attrname, permission.Attrname, permission.Attrname)
					if err != nil {
						continue
					}
					if permission.ApplicationList != nil {
						_, err := pbxcash.SetConfHttapiPermissionApplicationList(permission.Attrname, permission.ApplicationList.Attrdefault)
						if err != nil {
							continue
						}
						for _, application := range permission.ApplicationList.Application {
							_, err := pbxcash.SetConfHttapiPermissionApplicationListApplication(application.Attrname)
							if err != nil {
								continue
							}
						}
					}
					if permission.VariableList != nil {
						_, err := pbxcash.SetConfHttapiPermissionVariableList(permission.Attrname, permission.VariableList.Attrdefault)
						if err != nil {
							continue
						}
						for _, variable := range permission.VariableList.Variable {
							_, err := pbxcash.SetConfHttapiPermissionVariableListVariable(variable.Attrname)
							if err != nil {
								continue
							}
						}
					}
					if permission.ApiList != nil {
						_, err := pbxcash.SetConfHttapiPermissionApiList(permission.Attrname, permission.ApiList.Attrdefault)
						if err != nil {
							continue
						}
						for _, api := range permission.VariableList.Variable {
							_, err := pbxcash.SetConfHttapiPermissionApiListApi(api.Attrname)
							if err != nil {
								continue
							}
						}
					}
				}
				for _, param := range profile.Params.Param{
					_, err := pbxcash.SetConfHttapiParam(param.Attrname, param.Attrname)
					if err != nil {
						continue
					}
				}
			}
		}
	}

	func setConfigHttpCache(conf *Configuration) {
		if conf.Settings != nil {
			return
		}
		for _, param := range conf.Settings.Param {
			_, err := pbxcash.SetConfigHttpCache(param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}

		}
	}

	func setConfigIvr(conf *Configuration) {
		if conf.Menus != nil {
			return
		}
		for _, menu := range conf.Menus.Menu {
			_, err := pbxcash.SetConfigIvrMenu(menu)
			if err != nil {
				continue
			}
			for _, entry := range menu.Entry {
				_, err := pbxcash.SetConfigIvrMenuEntry(menu.Attrname, entry.Attraction, entry.Attrdigits, entry.Attrparam)
				if err != nil {
					continue
				}
			}
		}
	}

	func setConfigLogfile(conf *Configuration) {
		if conf.Settings != nil {
			return
		}
		for _, param := range conf.Settings.Param {
			_, err := pbxcash.SetConfigLogfileSettings(param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
		if conf.Profiles != nil {
			return
		}
		for _, profile := range conf.Profiles.Profile {
			_, err := pbxcash.SetConfigLogfileProfile(profile.Attrname)
			if err != nil {
				continue
			}
			if profile.Settings == nil {
				continue
			}
			for _, param := range profile.Settings.Param {
				_, err := pbxcash.SetConfigLogfileProfileParam(profile.Attrname, param.Attrname, param.Attrvalue)
				if err != nil {
					continue
				}
			}
		}
	}

func setConfigModules(conf *Configuration) {

		if conf.Modules != nil {
			return
		}
		for _, module := range conf.Modules.Load {
			_, err := pbxcash.SetConfigModule(module.Attrmodule)
			if err != nil {
				continue
			}

		}
	}

	func setConfigNiblebill(conf *Configuration) {
		if conf.Settings != nil {
			return
		}
		for _, param := range conf.Settings.Param {
			_, err := pbxcash.SetConfigNiblebill(param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}

		}
	}

	func setConfigPocketsphinx(conf *Configuration) {
		if conf.Settings != nil {
			return
		}
		for _, param := range conf.Settings.Param {
			_, err := pbxcash.SetConfigPocketsphinx(param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
*/
func setConfigSofia(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfSofia()
	if err != nil {
		return err
	}

	if conf.GlobalSettings != nil {
		for _, param := range conf.GlobalSettings.Param {
			_, err := altData.SetConfigSofiaGlobalSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	if conf.Profiles == nil {
		return nil
	}
	for _, eslProfile := range conf.Profiles.Profile {
		profile, err := altData.SetConfigSofiaProfile(mod, eslProfile.Attrname)
		if err != nil {
			continue
		}
		if eslProfile.Aliases != nil {
			for _, alias := range eslProfile.Aliases.Alias {
				_, err := altData.SetConfigSofiaProfileAliases(profile, alias.Attrname)
				if err != nil {
					continue
				}
			}
		}
		if eslProfile.Gateways != nil {
			for _, eslGateway := range eslProfile.Gateways.Gateway {
				gateway, err := altData.SetConfigSofiaGateway(profile, eslGateway.Attrname)
				if err != nil {
					continue
				}
				for _, param := range eslGateway.Param {
					_, err := altData.SetConfigSofiaGatewayParam(gateway, param.Attrname, param.Attrvalue)
					if err != nil {
						continue
					}
				}
				if eslGateway.Variables != nil && eslGateway.Variables.Variable != nil {
					for _, variable := range eslGateway.Variables.Variable {
						_, err := altData.SetConfigSofiaGatewayVar(gateway, variable.Attrname, variable.Attrvalue, variable.Attrdirection)
						if err != nil {
							continue
						}
					}
				}
			}
		}
		if eslProfile.Domains != nil {
			for _, domain := range eslProfile.Domains.Domain {
				_, err := altData.SetConfigSofiaProfileDomain(profile, domain.Attrname, domain.Attralias, domain.Attrparse)
				if err != nil {
					continue
				}
			}
		}
		if eslProfile.Settings == nil {
			continue
		}
		for _, param := range eslProfile.Settings.Param {
			_, err := altData.SetConfigSofiaProfileParam(profile, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigVerto(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfVerto()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigVertoSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	if conf.Profiles == nil {
		return nil
	}
	for _, eslProfile := range conf.Profiles.Profile {
		profile, err := altData.SetConfigVertoProfile(mod, eslProfile.Attrname)
		if err != nil {
			continue
		}

		if eslProfile.Param == nil {
			continue
		}
		for _, param := range eslProfile.Param {
			_, err := altData.SetConfigVertoProfileParam(profile, param.Attrname, param.Attrvalue, param.Attrsecure)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigLcr(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfLcr()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigLcrSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	if conf.Profiles == nil {
		return nil
	}
	for _, eslProfile := range conf.Profiles.Profile {
		profile, err := altData.SetConfigLcrProfile(mod, eslProfile.Attrname)
		if err != nil {
			continue
		}

		if eslProfile.Param == nil {
			continue
		}
		for _, param := range eslProfile.Param {
			_, err := altData.SetConfigLcrProfileParam(profile, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigShout(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfShout()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigShoutSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigRedis(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfRedis()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigRedisSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigNibblebill(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfNibblebill()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigNibblebillSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigDb(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfDb()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigDbSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigMemcache(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfMemcache()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigMemcacheSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigAvmd(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfAvmd()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigAvmdSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigTtsCommandline(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfTtsCommandline()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigTtsCommandlineSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigCdrMongodb(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfCdrMongodb()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigCdrMongodbSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigHttpCache(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfHttpCache()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigHttpCacheSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigOpus(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfOpus()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigOpusSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigPython(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfPython()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigPythonSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigAlsa(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfAlsa()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigAlsaSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigAmr(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfAmr()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigAmrSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigAmrwb(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfAmrwb()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigAmrwbSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigCepstral(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfCepstral()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigCepstralSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigCidlookup(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfCidlookup()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigCidlookupSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigCurl(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfCurl()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigCurlSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigDialplanDirectory(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfDialplanDirectory()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigDialplanDirectorySetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigEasyroute(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfEasyroute()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigEasyrouteSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigErlangEvent(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfErlangEvent()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigErlangEventSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigEventMulticast(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfEventMulticast()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigEventMulticastSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigFax(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfFax()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigFaxSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigLua(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfLua()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigLuaSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigMongo(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfMongo()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigMongoSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigMsrp(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfMsrp()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigMsrpSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigOreka(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfOreka()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigOrekaSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigPerl(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfPerl()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigPerlSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigPocketsphinx(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfPocketsphinx()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigPocketsphinxSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigSangomaCodec(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfSangomaCodec()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigSangomaCodecSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigSndfile(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfSndfile()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigSndfileSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigXmlCdr(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfXmlCdr()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigXmlCdrSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigXmlRpc(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfXmlRpc()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigXmlRpcSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigZeroconf(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfZeroconf()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigZeroconfSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigPostSwitch(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfPostSwitch()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigPostSwitchSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	if conf.CliKeybindings != nil {
		for _, param := range conf.CliKeybindings.Key {
			_, err := altData.SetConfigPostSwitchCliKeybinding(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	if conf.DefaultPtimes != nil {
		for _, param := range conf.DefaultPtimes.Codec {
			_, err := altData.SetConfigPostSwitchDefaultPtime(mod, param.Attrname, param.Attrptime)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return nil
}

func setConfigDirectory(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfDirectory()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigDirectorySetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	log.Println(conf)
	log.Println(conf.Profiles)
	log.Println(conf.Profiles.Profile)
	if conf.Profiles == nil {
		return nil
	}
	for _, eslProfile := range conf.Profiles.Profile {
		profile, err := altData.SetConfigDirectoryProfile(mod, eslProfile.Attrname)
		if err != nil {
			continue
		}

		if eslProfile.Param == nil {
			continue
		}
		for _, param := range eslProfile.Param {
			_, err := altData.SetConfigDirectoryProfileParam(profile, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigFifo(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfFifo()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigFifoSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	log.Println(conf)
	log.Println(conf.Fifos)
	log.Println(conf.Fifos.Fifo)
	if conf.Fifos == nil {
		return nil
	}
	for _, eslFifo := range conf.Fifos.Fifo {
		profile, err := altData.SetConfigFifoFifo(mod, eslFifo.Attrname, eslFifo.Attrimportance)
		if err != nil {
			continue
		}

		if eslFifo.Member == nil {
			continue
		}
		for _, member := range eslFifo.Member {
			_, err := altData.SetConfigFifoFifoParam(profile, member.Attrtimeout, member.Attrsimo, member.Attlag, member.Body)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigOpal(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfOpal()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigOpalSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	log.Println(conf)
	log.Println(conf.Listeners)
	log.Println(conf.Listeners.Listener)
	if conf.Listeners == nil {
		return nil
	}
	for _, eslListener := range conf.Listeners.Listener {
		profile, err := altData.SetConfigOpalListener(mod, eslListener.Attrname)
		if err != nil {
			continue
		}

		if eslListener.Param == nil {
			continue
		}
		for _, param := range eslListener.Param {
			_, err := altData.SetConfigOpalListenerParam(profile, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigOsp(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfOsp()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigOspSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	log.Println(conf)
	log.Println(conf.Profiles)
	log.Println(conf.Profiles.Profile)
	if conf.Profiles == nil {
		return nil
	}
	for _, eslProfile := range conf.Profiles.Profile {
		profile, err := altData.SetConfigOspProfile(mod, eslProfile.Attrname)
		if err != nil {
			continue
		}

		if eslProfile.Param == nil {
			continue
		}
		for _, param := range eslProfile.Param {
			_, err := altData.SetConfigOspProfileParam(profile, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigUnicall(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no ECL connection")
	}

	mod, err := altData.SetConfUnicall()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, param := range conf.Settings.Param {
			_, err := altData.SetConfigUnicallSetting(mod, param.Attrname, param.Attrvalue)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	log.Println(conf)
	log.Println(conf.Spans)
	log.Println(conf.Spans.Span)
	if conf.Spans == nil {
		return nil
	}
	for _, eslSpan := range conf.Spans.Span {
		profile, err := altData.SetConfigUnicallSpan(mod, eslSpan.Attrid)
		if err != nil {
			continue
		}

		if eslSpan.Param == nil {
			continue
		}
		for _, param := range eslSpan.Param {
			_, err := altData.SetConfigUnicallSpanParam(profile, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func setConfigPostLoadModules(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}
	mod, err := altData.SetConfPostLoadModules()
	if err != nil {
		return err
	}
	if conf.Modules == nil {
		return errors.New("empty config object")
	}
	for _, eslList := range conf.Modules.Load {
		_, _ = altData.SetPostLoadModule(mod, eslList.Attrmodule)
	}
	return nil
}

func setConfigVoicemail(conf *xmlStruct.Configuration) error {
	if conf == nil {
		return errors.New("no config object")
	}
	mod, err := altData.SetConfVoicemail()
	if err != nil {
		return err
	}
	if conf.Settings != nil {
		for _, eslList := range conf.Settings.Param {
			_, _ = altData.SetVoicemailSetting(mod, eslList.Attrname, eslList.Attrvalue)
		}
	}
	if conf.Profiles != nil {
		for _, eslList := range conf.Profiles.Profile {
			profile, err := altData.SetVoicemailProfile(mod, eslList.Attrname)
			if err != nil {
				continue
			}
			for _, param := range eslList.Param {
				_, err := altData.SetVoicemailProfileParam(profile, param.Attrname, param.Attrvalue)
				if err != nil {
					continue
				}
			}
			if eslList.Email == nil {
				continue
			}
			for _, emailParam := range eslList.Email.Param {
				_, err := altData.SetConfigVoicemailProfileEmail(profile, emailParam.Attrname, emailParam.Attrvalue)
				if err != nil {
					continue
				}
			}
		}
	}
	return nil
}

/*
func setConfigXmlCurl(conf *Configuration) {
	if conf.Bindings != nil {
		return
	}
	for _, bind := range conf.Bindings.Binding {
		_, err := pbxcash.SetConfigXmlCurlBinding(bind.Attrname)
		if err != nil {
			continue
		}
		for _, param := range bind.Param {
			_, err := pbxcash.SetConfigXmlCurlBindingParam(bind.Attrname, param.Attrname, param.Attrvalue)
			if err != nil {
				continue
			}
		}
	}
}
*/

// Dialplan
func GetXMLDialplan() {
	if esl == nil || !esl.Connected() {
		return
	}
	dialplanXML, ok := getFPBXContexts()
	if !ok {
		rawXML, err := esl.SendApiCmd("xml_locate dialplan")
		if err != nil {
			return
		}
		err = xml.Unmarshal([]byte(rawXML), &dialplanXML)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//fmt.Println(rawXML)
	//fmt.Printf("%+v", directoryXML)

	for _, eslContex := range dialplanXML.Context {
		context, err := pbxcache.SetContext(eslContex.Attrname)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if context == nil {
			continue
		}

		temp := context.Dialplan.NoProceed
		context.Dialplan.NoProceed = false
		if eslContex.Extension != nil {
			getXMLDialplanContextData(eslContex, context)
		}
		context.Dialplan.NoProceed = temp
		context.CacheFullXML()
	}
	//lua ~XML_REQUEST={section="dialplan",tag_name="",context="",key_name="",key_value=""};Params={};Params.__index=Params;function Params:getHeader (name) return self[name];end;function Params:serialize() return;end;function Params:setHeader(name,value) self[name]=value;end;params={};setmetatable(params,Params);params:setHeader("","");argv={"xml_handler"};loadfile("/usr/share/freeswitch/scripts/app.lua")();stream:write(XML_STRING);
	//lua ~local Database=require "resources.functions.database";dbh=Database.new('system');assert(dbh:connected());hash={};local res="[";sql="select coalesce(dialplan_context, '') from v_dialplans;";dbh:query(sql,{},function(row) if (not hash[row.dialplan_context]) then res=res..'"'..row.dialplan_context..'",';hash[row.dialplan_context]=true;end;end);dbh:release();res=res.."]";stream:write(res);
}

func getXMLDialplanContextData(eslContex *xmlStruct.Context, context *mainStruct.Context) {
	for _, eslExtension := range eslContex.Extension {
		extension, err := pbxcache.SetContextExtension(context, eslExtension.Attrname, eslExtension.Attrcontinue)
		if err != nil || extension == nil {
			continue
		}
		if eslExtension.Condition != nil {
			for _, eslCondition := range eslExtension.Condition {
				/*							expression := ""
											if eslCondition.Expression != nil {
												expression = eslCondition.Expression.string
											}*/
				condition, err := pbxcache.SetExtensionCondition(
					extension,
					eslCondition.Attrbreak,
					eslCondition.Attrfield,
					eslCondition.Attrexpression,
					eslCondition.Attrhour,
					eslCondition.Attrmday,
					eslCondition.Attrmon,
					eslCondition.Attrmweek,
					eslCondition.Attrwday,
					eslCondition.DateTime,
					eslCondition.TimeOfDay,
					eslCondition.Year,
					eslCondition.Minute,
					eslCondition.Week,
					eslCondition.Yday,
					eslCondition.Minday,
					eslCondition.TzOffset,
					eslCondition.Dst,
					eslCondition.Attrregex,
				)

				if err != nil || condition == nil {
					continue
				}

				if eslCondition.Regex != nil {
					for _, eslRegex := range eslCondition.Regex {
						_, err := pbxcache.SetConditionRegex(condition, eslRegex.Attrfield, eslRegex.Attrexpression)
						if err != nil {
							continue
						}
					}
				}
				if eslCondition.Action != nil {
					for _, eslAction := range eslCondition.Action {
						var inline bool
						if eslAction.Attrinline == "true" {
							inline = true
						}
						_, err := pbxcache.SetConditionAction(condition, eslAction.Attrapplication, eslAction.Attrdata, inline)
						if err != nil {
							continue
						}
					}
				}
				if eslCondition.AntiAction != nil {
					for _, eslAntiAction := range eslCondition.AntiAction {
						_, err := pbxcache.SetConditionAntiAction(condition, eslAntiAction.Attrapplication, eslAntiAction.Attrdata)
						if err != nil {
							continue
						}
					}
				}
			}
		}
	}
}

func getFPBXContexts() (xmlStruct.Section, bool) {
	scriptName, args, scriptPath := CheckLuaXMLHandlerDialplan()
	if scriptName == "" {
		return xmlStruct.Section{}, false
	}
	command := `lua ~local Database=require "resources.functions.database";dbh=Database.new('system');assert(dbh:connected());hash={};local res="[";sql="select coalesce(dialplan_context, '') as dialplan_context from v_dialplans;";dbh:query(sql,{},function(row) if (not hash[row.dialplan_context]) then res=res..'"'..row.dialplan_context..'",';hash[row.dialplan_context]=true;end;end);dbh:release();res=res..'""]';stream:write(res);`
	res, err := esl.SendApiCmd(command)
	if err != nil {
		log.Println(err.Error())
		return xmlStruct.Section{}, false
	}
	var contexts []string
	err = json.Unmarshal([]byte(res), &contexts)
	if err != nil {
		log.Println(err.Error())
		return xmlStruct.Section{}, false
	}

	var dialplanXML xmlStruct.Section
	for _, context := range contexts {
		if context == "" || string(context[0]) == "$" {
			continue
		}
		resp, ok := luaXMLRequest("dialplan", scriptName, args, scriptPath, map[string]string{"Caller-Context": context})
		if !ok {
			log.Println("NOT OK!")
			continue
		}
		resp = fixBrokenXML(resp)
		var dialplanNewXML xmlStruct.Document
		err := xml.Unmarshal([]byte(resp), &dialplanNewXML)
		if err != nil {
			log.Println(err.Error(), resp)
			continue
		}
		if len(dialplanNewXML.Section) == 0 || len(dialplanNewXML.Section[0].Context) == 0 {
			continue
		}
		dialplanXML.Context = append(dialplanXML.Context, dialplanNewXML.Section[0].Context[0])
	}

	return dialplanXML, true
	//lua ~XML_REQUEST={section="dialplan",tag_name="",context="",key_name="",key_value=""};Params={};Params.__index=Params;function Params:getHeader (name) return self[name];end;function Params:serialize() return;end;function Params:setHeader(name,value) self[name]=value;end;params={};setmetatable(params,Params);params:setHeader("Caller-Context","public");argv={"xml_handler"};loadfile("/usr/share/freeswitch/scripts/app.lua")();stream:write(XML_STRING);

}

func localParseInt(str string) int64 {
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}
