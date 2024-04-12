package web

import (
	"crypto/rand"
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/apps"
	"custompbx/cache"
	"custompbx/cdrDb"
	"custompbx/cfg"
	"custompbx/daemonCache"
	"custompbx/db"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"custompbx/webcache"
	"encoding/json"
	"fmt"
	"github.com/custompbx/customorm"
	"github.com/custompbx/hepparser"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var eventChannel chan interface{}

func onlyAdminGroup() []int {
	return []int{mainStruct.GetAdminId()}
}

func onlyAdminAndManagerGroup() []int {
	return []int{mainStruct.GetAdminId(), mainStruct.GetManagerId()}
}

func onlyAdminManagerAndUserGroup() []int {
	return []int{mainStruct.GetAdminId(), mainStruct.GetManagerId(), mainStruct.GetUserId()}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SetBroadcastChannel(brChannel chan interface{}) {
	eventChannel = brChannel
}

var b = &webStruct.WsHub{}

func TimeEvents() {
	twoSecondsTick := time.Tick(2 * time.Second)

	for {
		select {
		case event := <-eventChannel:
			switch event.(type) {
			case *mainStruct.Dashboard:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetDashboard, Dashboard: event.(*mainStruct.Dashboard)})
			case *altStruct.ConfigSofiaProfile:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetSofiaProfiles, Data: map[int64]*altStruct.ConfigSofiaProfile{event.(*altStruct.ConfigSofiaProfile).Id: event.(*altStruct.ConfigSofiaProfile)}})
			case *altStruct.ConfigSofiaProfileGateway:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetSofiaProfileGateways, Data: map[int64]*altStruct.ConfigSofiaProfileGateway{event.(*altStruct.ConfigSofiaProfileGateway).Id: event.(*altStruct.ConfigSofiaProfileGateway)}})
			case *altStruct.Configurations:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetModules, Module: event.(*altStruct.Configurations)})
			case *mainStruct.DialplanDebug:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.DialplanDebug, DialplanDebug: event.(*mainStruct.DialplanDebug)})
			case *altStruct.DirectoryDomainUser:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetDirectoryUser, Data: struct {
					A interface{} `json:"directory_users"`
				}{A: event.(*altStruct.DirectoryDomainUser)}})
			case *map[int64]*mainStruct.Agent:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.SubscribeCallcenterAgents, CallcenterAgentsList: event.(*map[int64]*mainStruct.Agent)})
			case *mainStruct.DaemonState:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.BroadcastConnection, Daemon: event.(*mainStruct.DaemonState)})
			case *hepparser.HEP:
				// b.Broadcast(webStruct.UserResponse{MessageType: webStruct.SubscribeHepPackages, HEPs: event.(*hepparser.HEP)})
			default:
				log.Printf("Unknown event type: %T - %+v\n", event, event)
			}

		case <-twoSecondsTick:
			b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetDashboard, Dashboard: &mainStruct.Dashboard{DashboardData: webcache.GetDashboardData()}})
		}
	}
}

func StartWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("ERROR on StartWS: %+v", err.Error())
		return
	}
	fmt.Println("NEW WS CONNECTION")
	wsContext := webStruct.CreateWsContext(ws)

	fmt.Println("STARTING GOROUTINES")
	go wsContext.SendWaiter()
	go wsContext.ReadWaiter(messageHandler)
	b.Hub = append(b.Hub, wsContext)
}

func PostAPIRequest(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.Write([]byte("empty body\n"))
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var msg webStruct.Message
	err := decoder.Decode(&msg)
	if err != nil {
		w.Write([]byte("can't parse request\n"))
		return
	}

	var resp webStruct.UserResponse
	msg.Data.Trim()
	msg.Data.Event = msg.Event
	resp = messageMainHandler(msg.Data)
	res, err := json.Marshal(resp)
	if err != nil {
		res = []byte("can't marshal response\n")
		return
	}
	_, err = w.Write(res)
	if err != nil {
		log.Printf("%+v", err)
		res = []byte("can't send response\n")
	}
}

func tokenGenerator() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func messageHandler(msg *webStruct.Message, wsContext *webStruct.WsContext) {
	if !daemonCache.State.DatabaseConnection {
		wsContext.SendChannel <- &webStruct.UserResponse{Daemon: daemonCache.State, MessageType: webStruct.BroadcastConnection}
		return
	}

	msg.Data.Trim()
	msg.Data.Event = msg.Event
	msg.Data.Context = wsContext
	var resp webStruct.UserResponse
	switch msg.Event {
	case "login":
		resp = checkLogin(msg.Data)
	case "[Auth] Logout":
		wsContext.Subscriptions.Clear()
		resp = getUser(msg.Data, loginOut, onlyAdminGroup())
	case "relogin":
		resp = getUser(msg.Data, checkRelogin, onlyAdminGroup())
	case webStruct.DialplanDebug:
		resp = getUser(msg.Data, getDialplanDebug, onlyAdminGroup())
	case webStruct.SubscribeHepPackages:
		resp = getUser(msg.Data, getDialplanDebug, onlyAdminGroup())
	case "SubscriptionList":
		resp = getUser(
			msg.Data,
			func(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
				resp.MessageType = "SubscriptionList"
				wsContext.Subscriptions.Clear()
				if len(msg.Data.ArrVal) > 10 || len(msg.Data.ArrVal) == 0 {
					resp.Error = "can't subscribe!"
				} else {
					for _, name := range msg.Data.ArrVal {
						wsContext.Subscriptions.Set(name)
					}
				}
				return resp
			},
			onlyAdminGroup(),
		)
	case webStruct.Unsubscribe:
		resp = getUser(
			msg.Data,
			func(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
				if msg.Data.Name != "" {
					wsContext.Subscriptions.Del(msg.Data.Name)
				} else {
					wsContext.Subscriptions.Clear()
				}
				resp.MessageType = "OK"
				resp.MessageType = "SubscriptionList"
				return resp
			},
			onlyAdminGroup(),
		)
	case "[Dialplan][Switch] Debug":
		resp = getUser(msg.Data, switchDialplanDebug, onlyAdminGroup())
	case "AddUserToken":
		resp = getUser(msg.Data, createAPIToken, onlyAdminGroup())
	case "GetUserTokens":
		resp = getUser(msg.Data, GetUserTokens, onlyAdminGroup())
	case "UserGetOwnTokens":
		resp = getUser(msg.Data, UserGetOwnTokens, onlyAdminGroup())
	case "RemoveUserToken":
		resp = getUser(msg.Data, RemoveUserToken, onlyAdminGroup())
	default:
		resp = messageMainHandler(msg.Data)
	}

	wsContext.SendChannel <- &resp
}

func messageMainHandler(msg *webStruct.MessageData) webStruct.UserResponse {
	// if !daemonCache.State.DatabaseConnection || !daemonCache.State.ESLConnection {
	if !daemonCache.State.DatabaseConnection {
		return webStruct.UserResponse{Daemon: daemonCache.State, MessageType: webStruct.BroadcastConnection}
	}

	var resp webStruct.UserResponse

	switch msg.Event {
	case "get_settings":
		resp = getUser(msg, checkSettings, onlyAdminGroup())
	case "set_settings":
		resp = getUser(msg, setSettings, onlyAdminGroup())
	case webStruct.GetDashboard:
		resp = getUser(msg, getDashboardData, onlyAdminGroup())
	case "GetInstances":
		resp = getUser(msg, GetInstances, onlyAdminGroup())
	case "UpdateInstanceDescription":
		resp = getUser(msg, UpdateInstanceDescription, onlyAdminGroup())
	//Doc started ---- (//Request:.*parent":\{.*id":\d+)[^\}]+ //(Request:.*),"description":""(.+) //(Request:(?!.*Switch)+)(.+)(,|\{)"enabled":(?:false|true)(?:,|\{) //(Request:(.*Switch)+.+),"parent":\{"id":\d+\} //(Request:(?!.*Move)+.*),"position":\d+
	//## Directory
	//### Domains
	//Request:{"event":"GetDirectoryDomains","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDirectoryDomains","data":{"4":{"id":4,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0}}}
	//Errors:no id, DB error
	case "GetDirectoryDomains":
		msg.Id = cache.GetCurrentInstanceId()
		resp = getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomain{}, onlyAdminGroup())
	//Request:{"event":"ImportDirectory","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"ImportDirectory"}
	//Errors:empty data, can't parse
	case "ImportDirectory":
		resp = getUser(msg, importDirectory, onlyAdminGroup())
	//Request:{"event":"ImportXMLDomain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","file":"\r\n  <!--the domain or ip (the right hand side of the @ in the addr-->\r\n  <domain name=\"test_do\">\r\n    <params>\r\n      <param name=\"jsonrpc-allowed-methods\" value=\"verto\"/>>\r\n    </params>\r\n    <variables>\r\n      <variable name=\"record_stereo\" value=\"true\"/>\r\n      <variable name=\"default_gateway\" value=\"$${default_provider}\"/>\r\n    </variables>\r\n    <groups>\r\n      <group name=\"default\">\r\n\t<users>\r\n\t</users>\r\n      </group>\r\n    </groups>\r\n  </domain>"}}
	//Response:{"MessageType":"ImportXMLDomain","data":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0},"7":{"id":7,"position":2,"enabled":true,"name":"test_do","parent":{"id":1},"sip_regs_counter":0}}}
	//Errors:can't parse
	case "ImportXMLDomain":
		resp = getUser(msg, ImportXMLDomain, onlyAdminGroup())
		if resp.Error != "" {
			return resp
		}
		msg.Id = cache.GetCurrentInstanceId()
		resp = getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomain{}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryDomain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"test"}}
	//Response:{"MessageType":"AddDirectoryDomain","data":{"id":9,"position":4,"enabled":true,"name":"test","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	case "AddDirectoryDomain":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomain{Name: msg.Name, Enabled: true, Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}}, onlyAdminGroup())
	//Request:{"event":"RenameDirectoryDomain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":9,"name":"test2"}}
	//Response:{"MessageType":"RenameDirectoryDomain","data":{"id":9,"position":4,"enabled":true,"name":"test2","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	case "RenameDirectoryDomain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomain{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryDomain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":8,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryDomain","data":{"id":8,"position":3,"enabled":false,"name":"test_do2","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	case "SwitchDirectoryDomain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomain{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryDomain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":9}}
	//Response:{"MessageType":"DeleteDirectoryDomain","data":{"id":9,"position":4,"enabled":true,"name":"test2","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	case "DeleteDirectoryDomain":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomain{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetDirectoryDomainDetails","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6}}
	//Response:{"MessageType":"GetDirectoryDomainDetails","data":{"parameters":{"10":{"id":10,"position":2,"enabled":true,"name":"jsonrpc-allowed-methods","value":"verto","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"9":{"id":9,"position":1,"enabled":true,"name":"dial-string","value":"{^^:sip_invite_domain=${dialed_domain}:presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(*/${dialed_user}@${dialed_domain})},${verto_contact(${dialed_user}@${dialed_domain})}","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}},"variables":{"13":{"id":13,"position":1,"enabled":true,"name":"record_stereo","value":"true","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"14":{"id":14,"position":2,"enabled":true,"name":"default_gateway","value":"example.com","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"15":{"id":15,"position":3,"enabled":true,"name":"default_areacode","value":"918","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"16":{"id":16,"position":4,"enabled":true,"name":"transfer_fallback_extension","value":"operator","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}}}
	//Errors:DB error
	case "GetDirectoryDomainDetails":
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainParameter{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainVariable{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"parameters"`
			Sch interface{} `json:"variables"`
		}{S: resp1.Data, Sch: resp2.Data}}
	//Request:{"event":"AddDirectoryDomainParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6,"name":"paramn","value":"paramv"}}
	//Response:{"MessageType":"AddDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:DB error
	case "AddDirectoryDomainParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainParameter{Name: msg.Name, Value: msg.Value, Enabled: true, Parent: &altStruct.DirectoryDomain{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryDomainParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":15,"name":"paramn1","value":"paramv1"}}
	//Response:{"MessageType":"UpdateDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":true,"name":"paramn1","value":"paramv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "UpdateDirectoryDomainParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainParameter{Id: msg.Id, Name: msg.Name, Value: msg.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryDomainParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":15,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":false,"name":"paramn1","value":"paramv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "SwitchDirectoryDomainParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainParameter{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryDomainParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"DeleteDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":false,"name":"paramn1","value":"paramv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "DeleteDirectoryDomainParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainParameter{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryDomainVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6,"name":"varn","value":"varv"}}
	//Response:{"MessageType":"AddDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":true,"name":"varn","value":"varv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "AddDirectoryDomainVariable":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainVariable{Name: msg.Name, Value: msg.Value, Enabled: true, Parent: &altStruct.DirectoryDomain{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryDomainVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":21,"name":"varn1","value":"varv1"}}
	//Response:{"MessageType":"UpdateDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":true,"name":"varn1","value":"varv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "UpdateDirectoryDomainVariable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainVariable{Id: msg.Id, Name: msg.Name, Value: msg.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryDomainVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":21,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":false,"name":"varn1","value":"varv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "SwitchDirectoryDomainVariable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainVariable{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryDomainVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"DeleteDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":false,"name":"varn1","value":"varv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "DeleteDirectoryDomainVariable":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainVariable{Id: msg.Id}, onlyAdminGroup())
	//### Users
	//Request:{"event":"GetDirectoryUsers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDirectoryUsers","data":{"domains":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0},"7":{"id":7,"position":2,"enabled":true,"name":"test_do","parent":{"id":1},"sip_regs_counter":0},"8":{"id":8,"position":3,"enabled":true,"name":"test_do2","parent":{"id":1},"sip_regs_counter":0}},"directory_users":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}}
	//Errors:
	case webStruct.GetDirectoryUser:
		msg.Id = cache.GetCurrentInstanceId()
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomain{}, onlyAdminGroup())
		domains, ok := resp1.Data.(map[int64]interface{})
		if !ok {
			return webStruct.UserResponse{Error: "domains not found", MessageType: msg.Event}
		}
		var ids []int64
		for _, d := range domains {
			domain, ok := d.(altStruct.DirectoryDomain)
			if !ok || domain.Id == 0 {
				continue
			}
			ids = append(ids, domain.Id)
		}
		msg.IntSlice = ids
		resp2 := getUserForConfig(msg, getDirectoryByParents, &altStruct.DirectoryDomainUser{}, onlyAdminGroup())
		users, ok := resp2.Data.(map[int64]interface{})
		if !ok {
			return webStruct.UserResponse{Error: "users not found", MessageType: msg.Event}
		}
		directoryCache := cache.GetDirectoryCache()
		for k, u := range users {
			user, ok := u.(altStruct.DirectoryDomainUser)
			if !ok || user.Id == 0 {
				continue
			}
			cUser := directoryCache.UserCache.GetById(user.Id)
			if cUser == nil {
				continue
			}
			cUser.UpdateUser(&user)
			users[k] = user
		}

		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"domains"`
			Sch interface{} `json:"directory_users"`
		}{S: resp1.Data, Sch: users}}
	//Request:{"event":"GetDirectoryUserDetails","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":91}}
	//Response:{"MessageType":"GetDirectoryUserDetails","data":{"parameters":{"93":{"id":93,"position":1,"enabled":true,"name":"password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"94":{"id":94,"position":2,"enabled":true,"name":"vm-password","value":"1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}},"variables":{"346":{"id":346,"position":1,"enabled":true,"name":"toll_allow","value":"domestic,international,local","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"347":{"id":347,"position":2,"enabled":true,"name":"accountcode","value":"1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"348":{"id":348,"position":3,"enabled":true,"name":"user_context","value":"default","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"349":{"id":349,"position":4,"enabled":true,"name":"effective_caller_id_name","value":"Extension 1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"350":{"id":350,"position":5,"enabled":true,"name":"effective_caller_id_number","value":"1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"351":{"id":351,"position":6,"enabled":true,"name":"outbound_caller_id_name","value":"FreeSWITCH","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"352":{"id":352,"position":7,"enabled":true,"name":"outbound_caller_id_number","value":"0000000000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"353":{"id":353,"position":8,"enabled":true,"name":"callgroup","value":"techsupport","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}},"user":{"id":91,"position":1,"enabled":true,"name":"1000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "GetDirectoryUserDetails":
		userMsg := getUserForConfig(msg, getDirectoryById, &altStruct.DirectoryDomainUser{}, onlyAdminGroup())
		user, ok := userMsg.Data.(altStruct.DirectoryDomainUser)
		if !ok {
			return webStruct.UserResponse{Error: "user not found", MessageType: msg.Event}
		}
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUserParameter{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUserVariable{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			A interface{} `json:"parameters"`
			B interface{} `json:"variables"`
			C interface{} `json:"user"`
		}{A: resp1.Data, B: resp2.Data, C: user}}
	//Request:{"event":"AddDirectoryUserParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":91,"name":"paramn","value":"paramv"}}
	//Response:{"MessageType":"AddDirectoryUserParameter","data":{"id":137,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "AddDirectoryUserParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainUserParameter{Name: msg.Name, Value: msg.Value, Enabled: true, Parent: &altStruct.DirectoryDomainUser{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryUserVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":91,"name":"varn","value":"varv"}}
	//Response:{"MessageType":"AddDirectoryUserVariable","data":{"id":514,"position":9,"enabled":true,"name":"varn","value":"varv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "AddDirectoryUserVariable":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainUserVariable{Name: msg.Name, Value: msg.Value, Enabled: true, Parent: &altStruct.DirectoryDomainUser{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryUserParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":13}}
	//Response:{"MessageType":"DeleteDirectoryUserParameter","data":{"id":137,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "DeleteDirectoryUserParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainUserParameter{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryUserVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":51}}
	//Response:{"MessageType":"DeleteDirectoryUserVariable","data":{"id":514,"position":9,"enabled":true,"name":"varn","value":"varv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "DeleteDirectoryUserVariable":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainUserVariable{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":138,"name":"paramn1","value":"paramv1"}}
	//Response:{"MessageType":"UpdateDirectoryUserParameter","data":{"id":138,"position":3,"enabled":true,"name":"paramn1","value":"paramv1","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "UpdateDirectoryUserParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserParameter{Id: msg.Id, Name: msg.Name, Value: msg.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":515,"name":"varn1","value":"varv1"}}
	//Response:{"MessageType":"UpdateDirectoryUserVariable","data":{"id":515,"position":9,"enabled":true,"name":"varn1","value":"varv1","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "UpdateDirectoryUserVariable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserVariable{Id: msg.Id, Name: msg.Name, Value: msg.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserCache","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","value":"3000","id":91}}
	//Response:{"MessageType":"UpdateDirectoryUserCache","data":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	case "UpdateDirectoryUserCache":
		cacheValue, err := strconv.ParseUint(msg.Value, 10, 32)
		if err != nil {
			return webStruct.UserResponse{Error: "wrong data", MessageType: msg.Event}
		}
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUser{Id: msg.Id, Cache: uint(cacheValue)}, []string{"Cache"}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserCidr","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","value":"0.0.0.0","id":91}}
	//Response:{"MessageType":"UpdateDirectoryUserCidr","data":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"0.0.0.0","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	case "UpdateDirectoryUserCidr":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUser{Id: msg.Id, Cidr: msg.Value}, []string{"Cidr"}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserNumberAlias","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","value":"555","id":91}}
	//Response:{"MessageType":"UpdateDirectoryUserNumberAlias","data":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"0.0.0.0","number_alias":"555","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	case "UpdateDirectoryUserNumberAlias":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUser{Id: msg.Id, NumberAlias: msg.Value}, []string{"NumberAlias"}}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"5000","id":6}}
	//Response:{"MessageType":"AddDirectoryUser","data":{"115":{"id":115,"position":25,"enabled":true,"name":"5000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "AddDirectoryUser":
		resp = getUser(msg, addNewUser, onlyAdminGroup())
	//Request:{"event":"ImportXMLDomainUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","file":"<user id=\"1099\">\r\n    <params>\r\n      <param name=\"password\" value=\"$${default_password}\"/>\r\n      <param name=\"vm-password\" value=\"1099\"/>\r\n    </params>\r\n    <variables>\r\n      <variable name=\"toll_allow\" value=\"domestic,international,local\"/>\r\n      <variable name=\"accountcode\" value=\"1099\"/>\r\n      <variable name=\"user_context\" value=\"default\"/>\r\n      <variable name=\"effective_caller_id_name\" value=\"Extension 1990\"/>\r\n      <variable name=\"effective_caller_id_number\" value=\"1099\"/>\r\n      <variable name=\"outbound_caller_id_name\" value=\"$${outbound_caller_name}\"/>\r\n      <variable name=\"outbound_caller_id_number\" value=\"$${outbound_caller_id}\"/>\r\n      <variable name=\"callgroup\" value=\"techsupport\"/>\r\n    </variables>\r\n  </user>","id":6}}
	//Response:{"MessageType":"ImportXMLDomainUser","data":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"102":{"id":102,"position":12,"enabled":true,"name":"1011","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"103":{"id":103,"position":13,"enabled":true,"name":"1012","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"104":{"id":104,"position":14,"enabled":true,"name":"1013","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"105":{"id":105,"position":15,"enabled":true,"name":"1014","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"106":{"id":106,"position":16,"enabled":true,"name":"1015","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"107":{"id":107,"position":17,"enabled":true,"name":"1016","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"108":{"id":108,"position":18,"enabled":true,"name":"1017","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"109":{"id":109,"position":19,"enabled":true,"name":"1018","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"110":{"id":110,"position":20,"enabled":true,"name":"1019","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"111":{"id":111,"position":21,"enabled":true,"name":"brian","cache":1000,"cidr":"192.0.2.0/24","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"112":{"id":112,"position":22,"enabled":true,"name":"default","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"113":{"id":113,"position":23,"enabled":true,"name":"example.com","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"114":{"id":114,"position":24,"enabled":true,"name":"SEP001120AABBCC","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"115":{"id":115,"position":25,"enabled":true,"name":"5000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"116":{"id":116,"position":26,"enabled":true,"name":"1099","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"91":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"0.0.0.0","number_alias":"555","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"92":{"id":92,"position":2,"enabled":true,"name":"1001","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"93":{"id":93,"position":3,"enabled":true,"name":"1002","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"94":{"id":94,"position":4,"enabled":true,"name":"1003","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"95":{"id":95,"position":5,"enabled":true,"name":"1004","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"96":{"id":96,"position":6,"enabled":true,"name":"1005","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"97":{"id":97,"position":7,"enabled":true,"name":"1006","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"98":{"id":98,"position":8,"enabled":true,"name":"1007","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"99":{"id":99,"position":9,"enabled":true,"name":"1008","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "ImportXMLDomainUser":
		resp = getUser(msg, ImportXMLDomainUser, onlyAdminGroup())
		if resp.Error != "" {
			return resp
		}
		resp = getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUser{}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":11}}
	//Response:{"MessageType":"DeleteDirectoryUser","data":{"id":115,"position":25,"enabled":true,"name":"5000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	case "DeleteDirectoryUser":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainUser{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserName","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":116,"name":"1098"}}
	//Response:{"MessageType":"UpdateDirectoryUserName","data":{"id":116,"position":26,"enabled":true,"name":"1098","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	case "UpdateDirectoryUserName":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUser{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":116,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUser","data":{"id":116,"position":26,"enabled":false,"name":"1098","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	case "SwitchDirectoryUser":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUser{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryUserParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":140,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserParameter","data":{"id":140,"position":2,"enabled":false,"name":"vm-password","value":"1099","description":"","parent":{"id":116,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "SwitchDirectoryUserParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserParameter{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryUserVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":516,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserVariable","data":{"id":516,"position":1,"enabled":false,"name":"toll_allow","value":"domestic,international,local","description":"","parent":{"id":116,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "SwitchDirectoryUserVariable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserVariable{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//### Groups
	//Request:{"event":"GetDirectoryGroups","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDirectoryGroups","data":{"domains":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0},"7":{"id":7,"position":2,"enabled":true,"name":"test_do","parent":{"id":1},"sip_regs_counter":0},"8":{"id":8,"position":3,"enabled":true,"name":"test_do2","parent":{"id":1},"sip_regs_counter":0}},"list":{"10":{"id":10,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"11":{"id":11,"position":2,"enabled":true,"name":"sales","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"12":{"id":12,"position":3,"enabled":true,"name":"billing","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"13":{"id":13,"position":4,"enabled":true,"name":"support","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"14":{"id":14,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"15":{"id":15,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}}}
	//Errors:
	case "GetDirectoryGroups":
		msg.Id = cache.GetCurrentInstanceId()
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomain{}, onlyAdminGroup())
		domains, ok := resp1.Data.(map[int64]interface{})
		if !ok {
			return webStruct.UserResponse{Error: "domains not found", MessageType: msg.Event}
		}
		var ids []int64
		for _, d := range domains {
			domain, ok := d.(altStruct.DirectoryDomain)
			if !ok || domain.Id == 0 {
				continue
			}
			ids = append(ids, domain.Id)
		}
		msg.IntSlice = ids
		resp2 := getUserForConfig(msg, getDirectoryByParents, &altStruct.DirectoryDomainGroup{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"domains"`
			Sch interface{} `json:"list"`
		}{S: resp1.Data, Sch: resp2.Data}}
	//Request:{"event":"GetDirectoryGroupUsers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":12}}
	//Response:{"MessageType":"GetDirectoryGroupUsers","data":{"group_users":{"114":{"id":114,"position":2,"enabled":true,"type":"","description":"","parent":{"id":12,"position":0,"enabled":false,"name":"","description":"","parent":null},"user":{"id":100,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}},"users":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}}
	//Errors:
	case "GetDirectoryGroupUsers":
		resp0 := getUserForConfig(msg, getDirectoryById, &altStruct.DirectoryDomainGroup{}, onlyAdminGroup())
		group, ok := resp0.Data.(altStruct.DirectoryDomainGroup)
		if !ok || group.Id == 0 || group.Parent.Id == 0 {
			return webStruct.UserResponse{Error: "group not found", MessageType: msg.Event}
		}
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainGroupUser{}, onlyAdminGroup())
		msg.Id = group.Parent.Id
		resp2 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUser{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"group_users"`
			Sch interface{} `json:"users"`
		}{S: resp1.Data, Sch: resp2.Data}}
	//Request:{"event":"AddNewDirectoryGroup","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6,"name":"new_group"}}
	//Response:{"MessageType":"AddNewDirectoryGroup","data":{"id":16,"position":5,"enabled":true,"name":"new_group","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "AddNewDirectoryGroup":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainGroup{Name: msg.Name, Enabled: true, Parent: &altStruct.DirectoryDomain{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryGroup","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"DeleteDirectoryGroup","data":{"id":16,"position":5,"enabled":true,"name":"new_group","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "DeleteDirectoryGroup":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainGroup{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryGroupName","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":17,"name":"newnew"}}
	//Response:{"MessageType":"UpdateDirectoryGroupName","data":{"id":17,"position":5,"enabled":true,"name":"newnew","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	case "UpdateDirectoryGroupName":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainGroup{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryGroupUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id_int":91,"id":13}}
	//Response:{"MessageType":"AddDirectoryGroupUser","data":{"id":120,"position":3,"enabled":true,"type":"","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","description":"","parent":null},"user":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "AddDirectoryGroupUser":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainGroupUser{UserId: &altStruct.DirectoryDomainUser{Id: msg.IdInt}, Enabled: true, Parent: &altStruct.DirectoryDomainGroup{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryGroupUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":12}}
	//Response:{"MessageType":"DeleteDirectoryGroupUser","data":{"id":120,"position":3,"enabled":true,"type":"","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","description":"","parent":null},"user":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "DeleteDirectoryGroupUser":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainGroupUser{Id: msg.Id}, onlyAdminGroup())
	//### Gateways
	//Request:{"event":"GetDirectoryUserGateways","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDirectoryUserGateways","data":{"domains":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0}},"directory_users":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"102":{"id":102,"position":12,"enabled":true,"name":"1011","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"103":{"id":103,"position":13,"enabled":true,"name":"1012","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"104":{"id":104,"position":14,"enabled":true,"name":"1013","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"105":{"id":105,"position":15,"enabled":true,"name":"1014","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"user_gateways":{"5":{"id":5,"position":1,"enabled":true,"name":"example.com","description":"","parent":{"id":113,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}}}
	//Errors:
	case "GetDirectoryUserGateways":
		msg.Id = cache.GetCurrentInstanceId()
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomain{}, onlyAdminGroup())
		msg.Id = 0
		domains, ok := resp1.Data.(map[int64]interface{})
		if !ok {
			return webStruct.UserResponse{Error: "domains not found", MessageType: msg.Event}
		}
		for _, d := range domains {
			domain, ok := d.(altStruct.DirectoryDomain)
			if !ok || domain.Id == 0 {
				continue
			}
			msg.Id = domain.Id
			break
		}
		resp2 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUser{}, onlyAdminGroup())
		users, ok := resp2.Data.(map[int64]interface{})
		if !ok {
			return webStruct.UserResponse{Error: "users not found", MessageType: msg.Event}
		}
		var ids []int64
		for _, u := range users {
			user, ok := u.(altStruct.DirectoryDomainUser)
			if !ok || user.Id == 0 {
				continue
			}
			ids = append(ids, user.Id)
		}
		msg.IntSlice = ids
		resp3 := getUserForConfig(msg, getDirectoryByParents, &altStruct.DirectoryDomainUserGateway{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"domains"`
			Sch interface{} `json:"directory_users"`
			U   interface{} `json:"user_gateways"`
		}{S: resp1.Data, Sch: resp2.Data, U: resp3.Data}}
	//Request:{"event":"GetDirectoryUserGatewayDetails","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5}}
	//Response:{"MessageType":"GetDirectoryUserGatewayDetails","data":{"parameters":{"29":{"id":29,"position":1,"enabled":true,"name":"username","value":"joeuser","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"30":{"id":30,"position":2,"enabled":true,"name":"password","value":"password","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"31":{"id":31,"position":3,"enabled":true,"name":"from-user","value":"joeuser","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"32":{"id":32,"position":4,"enabled":true,"name":"from-domain","value":"example.com","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"33":{"id":33,"position":5,"enabled":true,"name":"expire-seconds","value":"600","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"34":{"id":34,"position":6,"enabled":true,"name":"register","value":"false","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"35":{"id":35,"position":7,"enabled":true,"name":"retry-seconds","value":"30","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"36":{"id":36,"position":8,"enabled":true,"name":"extension","value":"5000","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"37":{"id":37,"position":9,"enabled":true,"name":"context","value":"public","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}},"variables":{}}}
	//Errors:
	case "GetDirectoryUserGatewayDetails":
		resp1 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUserGatewayParameter{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getDirectoryByParent, &altStruct.DirectoryDomainUserGatewayVariable{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"parameters"`
			Sch interface{} `json:"variables"`
		}{S: resp1.Data, Sch: resp2.Data}}
	//Request:{"event":"AddDirectoryUserGatewayParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5,"name":"paramn","value":"paramv"}}
	//Response:{"MessageType":"AddDirectoryUserGatewayParameter","data":{"id":38,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddDirectoryUserGatewayParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainUserGatewayParameter{Name: msg.Name, Value: msg.Value, Enabled: true, Parent: &altStruct.DirectoryDomainUserGateway{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryUserGatewayParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":3}}
	//Response:{"MessageType":"DeleteDirectoryUserGatewayParameter","data":{"id":38,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DeleteDirectoryUserGatewayParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainUserGatewayParameter{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserGatewayParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":39,"name":"param","value":"param_new_val"}}
	//Response:{"MessageType":"UpdateDirectoryUserGatewayParameter","data":{"id":39,"position":10,"enabled":true,"name":"param","value":"param_new_val","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateDirectoryUserGatewayParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserGatewayParameter{Id: msg.Id, Name: msg.Name, Value: msg.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryUserGatewayVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5,"name":"varn","value":"varv","direction":"vard"}}
	//Response:{"MessageType":"AddDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":true,"name":"varn","value":"varv","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddDirectoryUserGatewayVariable":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainUserGatewayVariable{Name: msg.Name, Value: msg.Value, Direction: msg.Direction, Enabled: true, Parent: &altStruct.DirectoryDomainUserGateway{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserGatewayVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4,"name":"varn","value":"varv2222","direction":"vard"}}
	//Response:{"MessageType":"UpdateDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":true,"name":"varn","value":"varv2222","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateDirectoryUserGatewayVariable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserGatewayVariable{Id: msg.Id, Name: msg.Name, Value: msg.Value, Direction: msg.Direction}, []string{"Name", "Value", "Direction"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryUserGatewayVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":false,"name":"varn","value":"varv2222","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchDirectoryUserGatewayVariable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserGatewayVariable{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryUserGatewayVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4}}
	//Response:{"MessageType":"DeleteDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":false,"name":"varn","value":"varv2222","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DeleteDirectoryUserGatewayVariable":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainUserGatewayVariable{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryUserGateway","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_gw","id":93}}
	//Response:{"MessageType":"AddDirectoryUserGateway","data":{"id":6,"position":1,"enabled":true,"name":"new_gw","description":"","parent":{"id":93,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "AddDirectoryUserGateway":
		resp = getUserForConfig(msg, setConfig, &altStruct.DirectoryDomainUserGateway{Name: msg.Name, Enabled: true, Parent: &altStruct.DirectoryDomainUser{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DeleteDirectoryUserGateway","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6}}
	//Response:{"MessageType":"DeleteDirectoryUserGateway","data":{"id":6,"position":1,"enabled":true,"name":"new_gw2","description":"","parent":{"id":93,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "DeleteDirectoryUserGateway":
		resp = getUserForConfig(msg, delConfig, &altStruct.DirectoryDomainUserGateway{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryUserGatewayName","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6,"name":"new_gw2"}}
	//Response:{"MessageType":"UpdateDirectoryUserGatewayName","data":{"id":6,"position":1,"enabled":true,"name":"new_gw2","description":"","parent":{"id":93,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	case "UpdateDirectoryUserGatewayName":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserGateway{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryUserGatewayParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":39,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserGatewayParameter","data":{"id":39,"position":10,"enabled":false,"name":"param","value":"param_new_val","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchDirectoryUserGatewayParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.DirectoryDomainUserGatewayParameter{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//## Configuration
	//### Modules
	//Request:{"event":"[Config][Get] Modules","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config][Get] Modules","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":{"id":2,"position":2,"enabled":false,"name":"alsa.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"amr":{"id":52,"position":52,"enabled":true,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	case webStruct.GetModules:
		resp = getUser(msg, GetConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][Reload] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":52}}
	//Response:{"MessageType":"[Config][Reload] Module"}
	//Errors:
	case "[Config][Reload] Module":
		resp = getUser(msg, reloadConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][Unload] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":52}}
	//Response:{"MessageType":"[Config][Unload] Module"}
	//Errors:
	case "[Config][Unload] Module":
		resp = getUser(msg, unloadConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][Load] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":52}}
	//Response:{"MessageType":"[Config][Load] Module"}
	//Errors:
	case "[Config][Load] Module":
		resp = getUser(msg, loadConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][Switch] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":52,"enabled":false}}
	//Response:{"MessageType":"[Config][Switch] Module","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":{"id":2,"position":2,"enabled":false,"name":"alsa.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"amr":{"id":52,"position":52,"enabled":false,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	case "[Config][Switch] Module":
		resp = getUser(msg, switchConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][From scratch] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"alsa"}}
	//Response:{"MessageType":"[Config][From scratch] Module"}
	//Errors:
	case "[Config][From scratch] Module":
		resp = getUser(msg, fromScratchConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][Import] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"alsa"}}
	//Response:{"MessageType":"[Config][Import] Module"}
	//Errors:
	case "[Config][Import] Module":
		resp = getUser(msg, importConfModules, onlyAdminGroup())
	//Request:{"event":"[Config][Import] All Modules","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config][Import] All Modules"}
	//Errors:
	case "[Config][Import] All Modules":
		resp = getUser(msg, importConfAllModules, onlyAdminGroup())
	//Request:{"event":"TruncateModuleConfig","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":54}}
	//Response:{"MessageType":"TruncateModuleConfig","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":null,"amr":{"id":52,"position":52,"enabled":false,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	case "TruncateModuleConfig":
		resp = getUser(msg, TruncateModuleConfig, onlyAdminGroup())
	//Request:{"event":"ImportXMLModuleConfig","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","file":"<configuration name=\"alsa.conf\" description=\"Soundcard Endpoint\">\r\n  <settings>\r\n    <!--Default dialplan and caller-id info -->\r\n    <param name=\"dialplan\" value=\"XML\"/>\r\n    <param name=\"cid-name\" value=\"N800 Alsa\"/>\r\n    <param name=\"cid-num\" value=\"5555551212\"/>\r\n\r\n    <!--audio sample rate and interval -->\r\n    <param name=\"sample-rate\" value=\"8000\"/>\r\n    <param name=\"codec-ms\" value=\"20\"/>\r\n  </settings>\r\n</configuration>"}}
	//Response:{"MessageType":"ImportXMLModuleConfig","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":{"id":55,"position":53,"enabled":true,"name":"alsa.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"amr":{"id":52,"position":52,"enabled":false,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	case "ImportXMLModuleConfig":
		resp = getUser(msg, ImportXMLModuleConfig, onlyAdminGroup())
	//Request:{"event":"[Config][Autoload] Module","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":55}}
	//Response:{"MessageType":"[Config][Autoload] Module","data":{"id":15,"position":12,"enabled":true,"name":"mod_alsa","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Autoload] Module":
		res, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigurationsList{Id: msg.Id})
		if err != nil || res == nil {
			return webStruct.UserResponse{Error: "module not found", MessageType: msg.Event}
		}
		module, ok := res.(altStruct.ConfigurationsList)
		if !ok {
			return webStruct.UserResponse{Error: "module not found", MessageType: msg.Event}
		}
		module.Module = mainStruct.GetModuleNameByConfName(module.Name)
		result, err := intermediateDB.GetByValue(&altStruct.ConfigPostLoadModule{Name: module.Module, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadModule{}))}, map[string]bool{"Parent": true, "Name": true})
		if err != nil || module.Module == "" {
			return webStruct.UserResponse{Error: "module not found", MessageType: msg.Event}
		}
		if len(result) == 0 {
			resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPostLoadModule{Name: module.Module, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadModule{}))}, onlyAdminGroup())
		} else {
			postloadMod, ok := result[0].(altStruct.ConfigPostLoadModule)
			if !ok {
				return webStruct.UserResponse{Error: "module not found", MessageType: msg.Event}
			}
			resp = getUserForConfig(msg, updateConfig, struct {
				S interface{}
				A []string
			}{&altStruct.ConfigPostLoadModule{Id: postloadMod.Id, Enabled: !postloadMod.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
		}
	//### Acl
	//Request:{"event":"[Config] Get_acl_lists","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config] Get_acl_lists","data":{"1":{"id":1,"position":1,"enabled":true,"name":"lan","default":"allow","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"domains","default":"deny","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "[Config] Get_acl_lists":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigAclList{}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_acl_list","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"dddd","default":"deny"}}
	//Response:{"MessageType":"[Config] Add_acl_list","data":{"id":4,"position":3,"enabled":true,"name":"dddd","default":"deny","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Add_acl_list":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigAclList{Name: msg.Name, Default: msg.Default, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigAclList{}))}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_acl_list","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4,"name":"dddd2"}}
	//Response:{"MessageType":"[Config] Update_acl_list","data":{"id":4,"position":3,"enabled":true,"name":"dddd2","default":"deny","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Update_acl_list":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAclList{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_acl_list","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4}}
	//Response:{"MessageType":"[Config] Del_acl_list","data":{"id":4,"position":3,"enabled":true,"name":"dddd2","default":"deny","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Del_acl_list":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigAclList{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_acl_list_default","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","value":"allow","id":5}}
	//Response:{"MessageType":"[Config] Update_acl_list_default","data":{"id":5,"position":3,"enabled":true,"name":"ccc","default":"allow","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Update_acl_list_default":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAclList{Id: msg.Id, Default: msg.Value}, []string{"Default"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Get_acl_nodes","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"[Config] Get_acl_nodes","data":{"1":{"id":1,"position":1,"enabled":true,"type":"deny","cidr":"192.168.42.0/24","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}},"2":{"id":2,"position":4,"enabled":true,"type":"allow","cidr":"192.168.42.42/32","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}},"7":{"id":7,"position":5,"enabled":false,"type":"2","cidr":"2","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}}}}
	//Errors:
	case "[Config] Get_acl_nodes":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigAclNode{}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_acl_node","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"node":{"type":"allow","cidr":"0.0.0.0","domain":""}}}
	//Response:{"MessageType":"[Config] Add_acl_node","data":{"id":9,"position":6,"enabled":true,"type":"allow","cidr":"0.0.0.0","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}}}
	//Errors:
	case "[Config] Add_acl_node":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigAclNode{Type: msg.Node.Type, Cidr: msg.Node.Cidr, Domain: msg.Node.Domain, Enabled: true, Parent: &altStruct.ConfigAclList{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_acl_node","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":9}}
	//Response:{"MessageType":"[Config] Del_acl_node","data":{"id":9,"position":6,"enabled":true,"type":"allow","cidr":"0.0.0.0","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}}}
	//Errors:
	case "[Config] Del_acl_node":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigAclNode{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_acl_node","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","node":{"id":7,"type":"deny","cidr":"0.0.0.0","domain":""}}}
	//Response:{"MessageType":"[Config] Update_acl_node","data":{"id":7,"position":5,"enabled":true,"type":"deny","cidr":"0.0.0.0","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}}}
	//Errors:
	case "[Config] Update_acl_node":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAclNode{Id: msg.Node.Id, Type: msg.Node.Type, Cidr: msg.Node.Cidr, Domain: msg.Node.Domain}, []string{"Type", "Cidr", "Domain"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_acl_node","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","node":{"id":7,"enabled":false}}}
	//Response:{"MessageType":"[Config] Switch_acl_node","data":{"id":7,"position":5,"enabled":false,"type":"deny","cidr":"0.0.0.0","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}}}
	//Errors:
	case "[Config] Switch_acl_node":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAclNode{Id: msg.Node.Id, Enabled: msg.Node.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"MoveAclListNode","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","previous_index":1,"current_index":4,"id":1}}
	//Response:{"MessageType":"MoveAclListNode","data":{"1":{"id":1,"position":4,"enabled":true,"type":"deny","cidr":"192.168.42.0/24","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}},"2":{"id":2,"position":3,"enabled":true,"type":"allow","cidr":"192.168.42.42/32","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}},"7":{"id":7,"position":5,"enabled":false,"type":"deny","cidr":"0.0.0.0","domain":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","default":"","description":"","parent":null}}}}
	//Errors:
	case "MoveAclListNode":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAclNode{Id: msg.Id, Position: msg.CurrentIndex}, []string{"Position"}}, onlyAdminGroup())
	//### Sofia
	//Request:{"event":"[Config] Get_sofia_global_settings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config] Get_sofia_global_settings","data":{"1":{"id":1,"position":1,"enabled":true,"name":"log-level","value":"0","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"debug-presence","value":"0","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "[Config] Get_sofia_global_settings":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaGlobalSetting{}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_sofia_global_setting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"[Config] Update_sofia_global_setting","data":{"id":10,"position":3,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Update_sofia_global_setting":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaGlobalSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_sofia_global_setting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"[Config] Switch_sofia_global_setting","data":{"id":2,"position":2,"enabled":false,"name":"debug-presence","value":"0","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Switch_sofia_global_setting":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaGlobalSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_global_setting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"[Config] Add_sofia_global_setting","data":{"id":10,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Add_sofia_global_setting":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaGlobalSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigSofiaGlobalSetting{}))}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_global_setting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10}}}
	//Response:{"MessageType":"[Config] Del_sofia_global_setting","data":{"id":10,"position":3,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config] Del_sofia_global_setting":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaGlobalSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Get_sofia_profiles","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config] Get_sofia_profiles","data":{"1":{"id":1,"position":1,"enabled":true,"name":"external-ipv6","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":true,"state":"RUNNING (0)","uri":"sip:mod_sofia@[::1]:5080"},"2":{"id":2,"position":2,"enabled":true,"name":"external","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":true,"state":"RUNNING (0)","uri":"sip:mod_sofia@45.61.54.76:5080"},"3":{"id":3,"position":3,"enabled":true,"name":"internal-ipv6","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":true,"state":"RUNNING (0)","uri":"sip:mod_sofia@[::1]:5060"},"4":{"id":4,"position":4,"enabled":true,"name":"internal","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":true,"state":"RUNNING (21) (WSS)","uri":"sips:mod_sofia@45.61.54.76:7443;transport=wss"}}}
	//Errors:
	case webStruct.GetSofiaProfiles:
		//yeah getting profiles twice
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfile{}, onlyAdminGroup())
		profiles, ok := resp.Data.(map[int64]interface{})
		if ok {
			profilesX := fsesl.GetSofiaProfilesStatuses()
			for _, profileI := range profiles {
				profile, ok := profileI.(altStruct.ConfigSofiaProfile)
				if !ok {
					continue
				}
				profileX := profilesX[profile.Id]
				if profileX == nil {
					continue
				}
				profile.Started = profileX.Started
				profile.State = profileX.State
				profile.Uri = profileX.Uri
				profiles[profile.Id] = profile
			}
			resp.Data = profiles
		}
	//Request:{"event":"[Config] Get_sofia_profiles_params","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"[Config] Get_sofia_profiles_params","data":{"1":{"id":1,"position":1,"enabled":true,"name":"debug","value":"0","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"10":{"id":10,"position":10,"enabled":true,"name":"outbound-codec-prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"11":{"id":11,"position":11,"enabled":true,"name":"hold-music","value":"local_stream://moh","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"12":{"id":12,"position":12,"enabled":true,"name":"rtp-timer-name","value":"soft","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"13":{"id":13,"position":13,"enabled":true,"name":"local-network-acl","value":"localnet.auto","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"14":{"id":14,"position":14,"enabled":true,"name":"manage-presence","value":"false","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"15":{"id":15,"position":15,"enabled":true,"name":"inbound-codec-negotiation","value":"generous","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"16":{"id":16,"position":16,"enabled":true,"name":"nonce-ttl","value":"60","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"17":{"id":17,"position":17,"enabled":true,"name":"auth-calls","value":"false","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"18":{"id":18,"position":18,"enabled":true,"name":"inbound-late-negotiation","value":"true","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"19":{"id":19,"position":19,"enabled":true,"name":"inbound-zrtp-passthru","value":"true","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"2":{"id":2,"position":2,"enabled":true,"name":"sip-trace","value":"no","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"20":{"id":20,"position":20,"enabled":true,"name":"rtp-ip","value":"::1","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"21":{"id":21,"position":21,"enabled":true,"name":"sip-ip","value":"::1","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"22":{"id":22,"position":22,"enabled":true,"name":"rtp-timeout-sec","value":"300","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"23":{"id":23,"position":23,"enabled":true,"name":"rtp-hold-timeout-sec","value":"1800","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"24":{"id":24,"position":24,"enabled":true,"name":"tls","value":"false","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"25":{"id":25,"position":25,"enabled":true,"name":"tls-only","value":"false","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"26":{"id":26,"position":26,"enabled":true,"name":"tls-bind-params","value":"transport=tls","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"27":{"id":27,"position":27,"enabled":true,"name":"tls-sip-port","value":"5081","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"28":{"id":28,"position":28,"enabled":true,"name":"tls-passphrase","value":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"29":{"id":29,"position":29,"enabled":true,"name":"tls-verify-date","value":"true","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"3":{"id":3,"position":3,"enabled":true,"name":"sip-capture","value":"no","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"30":{"id":30,"position":30,"enabled":true,"name":"tls-verify-policy","value":"none","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"31":{"id":31,"position":31,"enabled":true,"name":"tls-verify-depth","value":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"32":{"id":32,"position":32,"enabled":true,"name":"tls-verify-in-subjects","value":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"33":{"id":33,"position":33,"enabled":true,"name":"tls-version","value":"tlsv1,tlsv1.1,tlsv1.2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"4":{"id":4,"position":4,"enabled":true,"name":"rfc2833-pt","value":"101","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"5":{"id":5,"position":5,"enabled":true,"name":"sip-port","value":"5080","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"6":{"id":6,"position":6,"enabled":true,"name":"dialplan","value":"XML","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"7":{"id":7,"position":7,"enabled":true,"name":"context","value":"public","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"8":{"id":8,"position":8,"enabled":true,"name":"dtmf-duration","value":"2000","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}},"9":{"id":9,"position":9,"enabled":true,"name":"inbound-codec-prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}}
	//Errors:
	case "[Config] Get_sofia_profiles_params":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_profile_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"[Config] Add_sofia_profile_param","data":{"id":180,"position":34,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Add_sofia_profile_param":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_profile_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":180}}}
	//Response:{"MessageType":"[Config] Del_sofia_profile_param","data":{"id":180,"position":34,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Del_sofia_profile_param":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_sofia_profile_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":180,"enabled":false}}}
	//Response:{"MessageType":"[Config] Switch_sofia_profile_param","data":{"id":180,"position":34,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Switch_sofia_profile_param":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_sofia_profile_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":180,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"[Config] Update_sofia_profile_param","data":{"id":180,"position":34,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Update_sofia_profile_param":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Get_sofia_profile_gateways","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"keep_subscription":true}}
	//Response:{"MessageType":"[Config] Get_sofia_profile_gateways","data":{"9":{"id":9,"position":1,"enabled":true,"name":"test","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""},"started":false,"state":""}}}
	//Errors:
	case "[Config] Get_sofia_profile_gateways":
		//yeah getting gateways twice
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfileGateway{}, onlyAdminGroup())
		gateways, ok := resp.Data.(map[int64]interface{})
		if ok {
			gatewaysX := fsesl.GetSofiaGatewaysStatuses()
			for _, gatewayI := range gateways {
				gateway, ok := gatewayI.(altStruct.ConfigSofiaProfileGateway)
				if !ok {
					continue
				}
				profileX := gatewaysX[gateway.Id]
				if profileX == nil {
					continue
				}
				gateway.Started = profileX.Started
				gateway.State = profileX.State
				gateways[gateway.Id] = gateway
			}
			resp.Data = gateways
		}
	//Request:
	//Response:
	//Errors:
	case "GetSofiaProfileGatewayVariables":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfileGatewayVariable{}, onlyAdminGroup())
	//Request:{"event":"GetSofiaProfileGatewayParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":9}}
	//Response:{"MessageType":"GetSofiaProfileGatewayParameters","data":{"20":{"id":20,"position":1,"enabled":true,"name":"test","value":"param","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":""}}}}
	//Errors:
	case "GetSofiaProfileGatewayParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfileGatewayParameter{}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_profile_gateway_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":9,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"[Config] Add_sofia_profile_gateway_param","data":{"id":21,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":""}}}
	//Errors:
	case "[Config] Add_sofia_profile_gateway_param":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfileGatewayParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigSofiaProfileGateway{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_sofia_profile_gateway_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":21,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"[Config] Update_sofia_profile_gateway_param","data":{"id":21,"position":2,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":""}}}
	//Errors:
	case "[Config] Update_sofia_profile_gateway_param":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileGatewayParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_sofia_profile_gateway_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":21,"enabled":false}}}
	//Response:{"MessageType":"[Config] Switch_sofia_profile_gateway_param","data":{"id":21,"position":2,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":""}}}
	//Errors:
	case "[Config] Switch_sofia_profile_gateway_param":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileGatewayParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_profile_gateway_param","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"[Config] Del_sofia_profile_gateway_param","data":{"id":21,"position":2,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":""}}}
	//Errors:
	case "[Config] Del_sofia_profile_gateway_param":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfileGatewayParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Config] Add_sofia_profile_gateway_var":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfileGatewayVariable{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigSofiaProfileGateway{Id: msg.Id}}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Config] Update_sofia_profile_gateway_var":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileGatewayVariable{Id: msg.Variable.Id, Name: msg.Variable.Name, Value: msg.Variable.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Config] Switch_sofia_profile_gateway_var":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileGatewayVariable{Id: msg.Variable.Id, Enabled: msg.Variable.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Config] Del_sofia_profile_gateway_var":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfileGatewayVariable{Id: msg.Variable.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_profile_gateway","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"test_gw","id":2}}
	//Response:{"MessageType":"[Config] Add_sofia_profile_gateway","data":{"id":10,"position":2,"enabled":true,"name":"test_gw","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""},"started":false,"state":""}}
	//Errors:
	case "[Config] Add_sofia_profile_gateway":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfileGateway{Name: msg.Name, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_profile_gateway","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"[Config] Del_sofia_profile_gateway","data":{"id":10,"position":2,"enabled":true,"name":"test_gw2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""},"started":false,"state":""}}
	//Errors:
	case "[Config] Del_sofia_profile_gateway":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfileGateway{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Rename_sofia_profile_gateway","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":10,"name":"test_gw2"}}
	//Response:{"MessageType":"[Config] Rename_sofia_profile_gateway","data":{"id":10,"position":2,"enabled":true,"name":"test_gw2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""},"started":false,"state":""}}
	//Errors:
	case "[Config] Rename_sofia_profile_gateway":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileGateway{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Get_sofia_profile_domains","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"[Config] Get_sofia_profile_domains","data":{"1":{"id":1,"position":1,"enabled":true,"name":"all","alias":false,"parse":true,"description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}}
	//Errors:
	case "[Config] Get_sofia_profile_domains":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfileDomain{}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_profile_domain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"sofia_domain":{"name":"domain2","alias":true,"parse":false}}}
	//Response:{"MessageType":"[Config] Add_sofia_profile_domain","data":{"id":12,"position":2,"enabled":true,"name":"domain2","alias":true,"parse":false,"description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Add_sofia_profile_domain":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfileDomain{Name: msg.SofiaDomain.Name, Alias: msg.SofiaDomain.Alias, Parse: msg.SofiaDomain.Parse, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_profile_domain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","sofia_domain":{"id":1}}}
	//Response:{"MessageType":"[Config] Del_sofia_profile_domain","data":{"id":12,"position":2,"enabled":true,"name":"domain2","alias":true,"parse":false,"description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Del_sofia_profile_domain":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfileDomain{Id: msg.SofiaDomain.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_sofia_profile_domain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","sofia_domain":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"[Config] Switch_sofia_profile_domain","data":{"id":1,"position":1,"enabled":false,"name":"all","alias":false,"parse":true,"description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Switch_sofia_profile_domain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileDomain{Id: msg.SofiaDomain.Id, Enabled: msg.SofiaDomain.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_sofia_profile_domain","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","sofia_domain":{"id":1,"name":"all","alias":false,"parse":true}}}
	//Response:{"MessageType":"[Config] Update_sofia_profile_domain","data":{"id":1,"position":1,"enabled":true,"name":"all","alias":false,"parse":true,"description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Update_sofia_profile_domain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileDomain{Id: msg.SofiaDomain.Id, Name: msg.SofiaDomain.Name, Alias: msg.SofiaDomain.Alias, Parse: msg.SofiaDomain.Parse}, []string{"Name", "Alias", "Parse"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Get_sofia_profile_aliases","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"[Config] Get_sofia_profile_aliases","data":{"4":{"id":4,"position":1,"enabled":true,"name":"domain_alias","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}}
	//Errors:
	case "[Config] Get_sofia_profile_aliases":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSofiaProfileAlias{}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_profile_alias","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"sofia_alias":{"name":"domain_alias2"}}}
	//Response:{"MessageType":"[Config] Add_sofia_profile_alias","data":{"id":5,"position":2,"enabled":true,"name":"domain_alias2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Add_sofia_profile_alias":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfileAlias{Name: msg.SofiaAlias.Name, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_profile_alias","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","sofia_alias":{"id":5}}}
	//Response:{"MessageType":"[Config] Del_sofia_profile_alias","data":{"id":5,"position":2,"enabled":true,"name":"domain_alias2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Del_sofia_profile_alias":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfileAlias{Id: msg.SofiaAlias.Id}, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_sofia_profile_alias","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","sofia_alias":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"[Config] Switch_sofia_profile_alias","data":{"id":4,"position":1,"enabled":false,"name":"domain_alias","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Switch_sofia_profile_alias":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileAlias{Id: msg.SofiaAlias.Id, Enabled: msg.SofiaAlias.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Update_sofia_profile_alias","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","sofia_alias":{"id":4,"name":"domain_alias3"}}}
	//Response:{"MessageType":"[Config] Update_sofia_profile_alias","data":{"id":4,"position":1,"enabled":true,"name":"domain_alias3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config] Update_sofia_profile_alias":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfileAlias{Id: msg.SofiaAlias.Id, Name: msg.SofiaAlias.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Add_sofia_profile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"[Config] Add_sofia_profile","data":{"id":19,"position":5,"enabled":true,"name":"new_profile","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":false,"state":"","uri":""}}
	//Errors:
	case "[Config] Add_sofia_profile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSofiaProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigSofiaProfile{}))}, onlyAdminGroup())
	//Request:{"event":"[Config] Rename_sofia_profile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":19,"name":"new_profile2"}}
	//Response:{"MessageType":"[Config] Rename_sofia_profile","data":{"id":19,"position":5,"enabled":true,"name":"new_profile2","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":false,"state":"","uri":""}}
	//Errors:
	case "[Config] Rename_sofia_profile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"[Config] Del_sofia_profile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"[Config] Del_sofia_profile","data":{"id":19,"position":5,"enabled":true,"name":"new_profile2","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":false,"state":"","uri":""}}
	//Errors:
	case "[Config] Del_sofia_profile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSofiaProfile{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"[API] Sofia profile command","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"stop","id":1,"id_int":0}}
	//Response:{"MessageType":"[API] Sofia profile command"}
	//Errors:
	case "[API] Sofia profile command":
		//TODO: replace
		resp = getUser(msg, runProfileCommand, onlyAdminGroup())
	//Request:{"event":"[Config] Switch_sofia_profile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"enabled":false}}
	//Response:{"MessageType":"[Config] Switch_sofia_profile","data":{"id":1,"position":1,"enabled":false,"name":"external-ipv6","description":"","parent":{"id":42,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null},"started":false,"state":"","uri":""}}
	//Errors:
	case "[Config] Switch_sofia_profile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSofiaProfile{Id: msg.Id, Enabled: *msg.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//### Cdr_Pg_Csv
	//Request:{"event":"[Config][Get] Cdr_Pg_Csv","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config][Get] Cdr_Pg_Csv","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"db-info","value":"host=localhost dbname=cdr connect_timeout=10","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"legs","value":"a","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"spool-format","value":"csv","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"rotate-on-hup","value":"true","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"schemas":{}}}
	//Errors:
	case "[Config][Get] Cdr_Pg_Csv":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigCdrPgCsvSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigCdrPgCsvSchema{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"settings"`
			Sch interface{} `json:"schemas"`
		}{S: resp1.Data, Sch: resp2.Data}}
	//Request:{"event":"[Config][Add] Cdr_Pg_Csv Parameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"[Config][Add] Cdr_Pg_Csv Parameter","data":{"id":12,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Add] Cdr_Pg_Csv Parameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCdrPgCsvSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCdrPgCsvSetting{}))}, onlyAdminGroup())
	//Request:{"event":"[Config][Update] Cdr_Pg_Csv Parameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":12,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"[Config][Update] Cdr_Pg_Csv Parameter","data":{"id":12,"position":5,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Update] Cdr_Pg_Csv Parameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCdrPgCsvSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Switch] Cdr_Pg_Csv Parameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":12,"enabled":false}}}
	//Response:{"MessageType":"[Config][Switch] Cdr_Pg_Csv Parameter","data":{"id":12,"position":5,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Switch] Cdr_Pg_Csv Parameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCdrPgCsvSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Delete] Cdr_Pg_Csv Parameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1}}}
	//Response:{"MessageType":"[Config][Delete] Cdr_Pg_Csv Parameter","data":{"id":12,"position":5,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Delete] Cdr_Pg_Csv Parameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCdrPgCsvSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config][Add] Cdr_Pg_Csv Field","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","field":{"var":"var","column":"column"}}}
	//Response:{"MessageType":"[Config][Add] Cdr_Pg_Csv Field","data":{"id":21,"position":1,"enabled":true,"var":"var","column":"column","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Add] Cdr_Pg_Csv Field":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCdrPgCsvSchema{Var: msg.Field.Var, Column: msg.Field.Column, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCdrPgCsvSchema{}))}, onlyAdminGroup())
	//Request:{"event":"[Config][Update] Cdr_Pg_Csv Field","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","field":{"id":21,"var":"var2","column":"column2"}}}
	//Response:{"MessageType":"[Config][Update] Cdr_Pg_Csv Field","data":{"id":21,"position":1,"enabled":true,"var":"var2","column":"column2","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Update] Cdr_Pg_Csv Field":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCdrPgCsvSchema{Id: msg.Field.Id, Var: msg.Field.Var, Column: msg.Field.Column}, []string{"Var", "Column"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Switch] Cdr_Pg_Csv Field","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","field":{"id":21,"enabled":false}}}
	//Response:{"MessageType":"[Config][Switch] Cdr_Pg_Csv Field","data":{"id":21,"position":1,"enabled":false,"var":"var2","column":"column2","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Switch] Cdr_Pg_Csv Field":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCdrPgCsvSchema{Id: msg.Field.Id, Enabled: msg.Field.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Delete] Cdr_Pg_Csv Field","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","field":{"id":2}}}
	//Response:{"MessageType":"[Config][Delete] Cdr_Pg_Csv Field","data":{"id":21,"position":1,"enabled":false,"var":"var2","column":"column2","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Delete] Cdr_Pg_Csv Field":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCdrPgCsvSchema{Id: msg.Field.Id}, onlyAdminGroup())
	//### GetOdbcCdr
	//Request:{"event":"GetOdbcCdr","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetOdbcCdr","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"safdfsadf","value":"dsafasdf","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"dfdsfd","value":"fdfd","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"tables":{"1":{"id":1,"position":1,"enabled":true,"name":"sfasf","log_leg":"asasasa2","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"cccc","log_leg":"","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetOdbcCdr":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigOdbcCdrSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigOdbcCdrTable{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S   interface{} `json:"settings"`
			Sch interface{} `json:"tables"`
		}{S: resp1.Data, Sch: resp2.Data}}
	//Request:{"event":"GetOdbcCdrField","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetOdbcCdrField","data":{"1":{"id":1,"position":1,"enabled":true,"name":"gg","chan_var_name":"gddd","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","log_leg":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"sss","chan_var_name":"ssss","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","log_leg":"","description":"","parent":null}}}}
	//Errors:
	case "GetOdbcCdrField":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigOdbcCdrTableField{Parent: &altStruct.ConfigOdbcCdrTable{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"AddOdbcCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOdbcCdrParameter","data":{"id":7,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOdbcCdrParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOdbcCdrSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOdbcCdrSetting{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateOdbcCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateOdbcCdrParameter","data":{"id":7,"position":3,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOdbcCdrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOdbcCdrSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOdbcCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7,"enabled":false}}}
	//Response:{"MessageType":"SwitchOdbcCdrParameter","data":{"id":7,"position":3,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchOdbcCdrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOdbcCdrSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteOdbcCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7}}}
	//Response:{"MessageType":"DeleteOdbcCdrParameter","data":{"id":7,"position":3,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DeleteOdbcCdrParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOdbcCdrSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"AddOdbcCdrTable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","table":{"name":"new_table"}}}
	//Response:{"MessageType":"AddOdbcCdrTable","data":{"id":9,"position":3,"enabled":true,"name":"new_table","log_leg":"","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOdbcCdrTable":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOdbcCdrTable{Name: msg.Table.Name, LogLeg: msg.Table.LogLeg, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOdbcCdrTable{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateOdbcCdrTable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","table":{"id":9,"name":"new_table2","log_leg":""}}}
	//Response:{"MessageType":"UpdateOdbcCdrTable","data":{"id":9,"position":3,"enabled":true,"name":"new_table2","log_leg":"","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOdbcCdrTable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOdbcCdrTable{Id: msg.Table.Id, Name: msg.Table.Name, LogLeg: msg.Table.LogLeg}, []string{"Name", "LogLeg"}}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "SwitchOdbcCdrTable":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOdbcCdrTable{Id: msg.Table.Id, Enabled: msg.Table.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteOdbcCdrTable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","table":{"id":9}}}
	//Response:{"MessageType":"DeleteOdbcCdrTable","data":{"id":9,"position":3,"enabled":true,"name":"new_table2","log_leg":"","description":"","parent":{"id":51,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DeleteOdbcCdrTable":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOdbcCdrTable{Id: msg.Table.Id}, onlyAdminGroup())
	//Request:{"event":"AddOdbcCdrField","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","odbc_cdr_field":{"name":"new_field","chan_var_name":"New_chan_var_name"},"id":1}}
	//Response:{"MessageType":"AddOdbcCdrField","data":{"id":4,"position":3,"enabled":true,"name":"new_field","chan_var_name":"New_chan_var_name","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","log_leg":"","description":"","parent":null}}}
	//Errors:
	case "AddOdbcCdrField":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOdbcCdrTableField{Name: msg.OdbcCdrField.Name, ChanVarName: msg.OdbcCdrField.ChanVarName, Enabled: true, Parent: &altStruct.ConfigOdbcCdrTable{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"UpdateOdbcCdrField","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","odbc_cdr_field":{"id":4,"name":"new_field2","chan_var_name":"New_chan_var_name2"}}}
	//Response:{"MessageType":"UpdateOdbcCdrField","data":{"id":4,"position":3,"enabled":true,"name":"new_field2","chan_var_name":"New_chan_var_name2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","log_leg":"","description":"","parent":null}}}
	//Errors:
	case "UpdateOdbcCdrField":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOdbcCdrTableField{Id: msg.OdbcCdrField.Id, Name: msg.OdbcCdrField.Name, ChanVarName: msg.OdbcCdrField.ChanVarName}, []string{"Name", "ChanVarName"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOdbcCdrField","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","odbc_cdr_field":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchOdbcCdrField","data":{"id":4,"position":3,"enabled":false,"name":"new_field2","chan_var_name":"New_chan_var_name2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","log_leg":"","description":"","parent":null}}}
	//Errors:
	case "SwitchOdbcCdrField":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOdbcCdrTableField{Id: msg.OdbcCdrField.Id, Enabled: msg.OdbcCdrField.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"DeleteOdbcCdrField","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","odbc_cdr_field":{"id":4}}}
	//Response:{"MessageType":"DeleteOdbcCdrField","data":{"id":4,"position":3,"enabled":false,"name":"new_field2","chan_var_name":"New_chan_var_name2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","log_leg":"","description":"","parent":null}}}
	//Errors:
	case "DeleteOdbcCdrField":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOdbcCdrTableField{Id: msg.OdbcCdrField.Id}, onlyAdminGroup())
	//### Verto
	//Request:{"event":"[Config][Verto][Get]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"[Config][Verto][Get]","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"debug","value":"0","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default-v4","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"default-v6","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "[Config][Verto][Get]":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigVertoSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigVertoProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"[Config][Verto][Profile][Parameters][Get]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"[Config][Verto][Profile][Parameters][Get]","data":{"1":{"id":1,"position":1,"enabled":true,"name":"bind-local","value":"45.61.54.76:8081","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":9,"enabled":true,"name":"rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":10,"enabled":true,"name":"ext-rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":11,"enabled":true,"name":"local-network","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"13":{"id":13,"position":12,"enabled":true,"name":"outbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"14":{"id":14,"position":13,"enabled":true,"name":"inbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"15":{"id":15,"position":14,"enabled":true,"name":"apply-candidate-acl","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"19":{"id":19,"position":15,"enabled":true,"name":"timer-name","value":"soft","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":2,"enabled":true,"name":"force-register-domain","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":7,"enabled":true,"name":"secure-combined","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":3,"enabled":true,"name":"secure-chain","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":4,"enabled":true,"name":"userauth","value":"true","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":5,"enabled":true,"name":"blind-reg","value":"false","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":6,"enabled":true,"name":"mcast-ip","value":"224.1.1.1","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":8,"enabled":true,"name":"mcast-port","value":"1337","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "[Config][Verto][Profile][Parameters][Get]":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVertoProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Settings][Update]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"debug","value":"1"}}}
	//Response:{"MessageType":"[Config][Verto][Settings][Update]","data":{"id":1,"position":1,"enabled":true,"name":"debug","value":"1","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Settings][Update]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Setting][Switch]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"[Config][Verto][Setting][Switch]","data":{"id":1,"position":1,"enabled":false,"name":"debug","value":"1","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Setting][Switch]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Setting][Add]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"param","value":"0"}}}
	//Response:{"MessageType":"[Config][Verto][Setting][Add]","data":{"id":5,"position":2,"enabled":true,"name":"param","value":"0","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Setting][Add]":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVertoSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVertoSetting{}))}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Setting][Del]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"[Config][Verto][Setting][Del]","data":{"id":5,"position":2,"enabled":true,"name":"param","value":"0","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Setting][Del]":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVertoSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Param][Add]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Param][Add]","data":{"id":39,"position":16,"enabled":true,"name":"paramn","value":"paramv","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Param][Add]":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVertoProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Secure: msg.Param.Secure, Enabled: true, Parent: &altStruct.ConfigVertoProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"MoveVertoProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","previous_index":16,"current_index":14,"id":39}}
	//Response:{"MessageType":"MoveVertoProfileParameter","data":{"1":{"id":1,"position":1,"enabled":true,"name":"bind-local","value":"45.61.54.76:8081","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":9,"enabled":true,"name":"rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":10,"enabled":true,"name":"ext-rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":11,"enabled":true,"name":"local-network","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"13":{"id":13,"position":12,"enabled":true,"name":"outbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"14":{"id":14,"position":13,"enabled":true,"name":"inbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"15":{"id":15,"position":15,"enabled":true,"name":"apply-candidate-acl","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"19":{"id":19,"position":16,"enabled":true,"name":"timer-name","value":"soft","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":2,"enabled":true,"name":"force-register-domain","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"39":{"id":39,"position":14,"enabled":true,"name":"paramn","value":"paramv","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":7,"enabled":true,"name":"secure-combined","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":3,"enabled":true,"name":"secure-chain","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":4,"enabled":true,"name":"userauth","value":"true","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":5,"enabled":true,"name":"blind-reg","value":"false","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":6,"enabled":true,"name":"mcast-ip","value":"224.1.1.1","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":8,"enabled":true,"name":"mcast-port","value":"1337","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "MoveVertoProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfileParameter{Id: msg.Id, Position: msg.CurrentIndex}, []string{"Position"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Param][Del]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":39}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Param][Del]","data":{"id":39,"position":14,"enabled":true,"name":"paramn","value":"paramv","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Param][Del]":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVertoProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Param][Switch]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":19,"enabled":false}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Param][Switch]","data":{"id":19,"position":16,"enabled":false,"name":"timer-name","value":"soft","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Param][Switch]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Update]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":19,"name":"timer-name","value":"hard","secure":""}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Update]","data":{"id":19,"position":19,"enabled":true,"name":"timer-name","value":"hard","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config][Verto][Profile][Update]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value, Secure: msg.Param.Secure}, []string{"Name", "Value", "secure"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Add]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"[Config][Verto][Profile][Add]","data":{"id":4,"position":3,"enabled":true,"name":"new_profile","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Add]":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVertoProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVertoProfile{}))}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Rename]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4,"name":"new_profile2"}}
	//Response:{"MessageType":"[Config][Verto][Profile][Rename]","data":{"id":4,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Rename]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Del]","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4}}
	//Response:{"MessageType":"[Config][Verto][Profile][Del]","data":{"id":4,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Del]":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVertoProfile{Id: msg.Id}, onlyAdminGroup())
	//### Callcenter
	//Request:{"event":"GetCallcenterQueues","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetCallcenterQueues","data":{"2":{"id":2,"position":2,"enabled":true,"name":"ddaaaaw","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ggdsf","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"a","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCallcenterQueues":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCallcenterQueue{}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterQueuesParams","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"GetCallcenterQueuesParams","data":{"1":{"id":1,"position":1,"enabled":true,"name":"ddd","value":"ddd","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetCallcenterQueuesParams":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCallcenterQueueParameter{}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterSettings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetCallcenterSettings","data":{"1":{"id":1,"position":1,"enabled":true,"name":"qqq","value":"qqq","description":"qqq","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCallcenterSettings":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCallcenterSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterSettings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"qqq2","value":"qqq2","description":"qqq"}}}
	//Response:{"MessageType":"UpdateCallcenterSettings","data":{"id":1,"position":1,"enabled":true,"name":"qqq2","value":"qqq2","description":"qqq","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateCallcenterSettings":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchCallcenterSettings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchCallcenterSettings","data":{"id":1,"position":1,"enabled":false,"name":"qqq2","value":"qqq2","description":"qqq","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchCallcenterSettings":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterSettings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCallcenterSettings","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCallcenterSettings":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCallcenterSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCallcenterSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterSettings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelCallcenterSettings","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCallcenterSettings":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCallcenterSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterQueueParam","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCallcenterQueueParam","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddCallcenterQueueParam":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCallcenterQueueParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigCallcenterQueue{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterQueueParam","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelCallcenterQueueParam","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelCallcenterQueueParam":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCallcenterQueueParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchCallcenterQueueParam","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchCallcenterQueueParam","data":{"id":1,"position":1,"enabled":false,"name":"ddd","value":"ddd","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchCallcenterQueueParam":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterQueueParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterQueueParam","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"new_param","value":"new_value"}}}
	//Response:{"MessageType":"UpdateCallcenterQueueParam","data":{"id":1,"position":1,"enabled":true,"name":"new_param","value":"new_value","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateCallcenterQueueParam":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterQueueParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterQueue","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_queue"}}
	//Response:{"MessageType":"AddCallcenterQueue","data":{"id":5,"position":5,"enabled":true,"name":"new_queue","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCallcenterQueue":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCallcenterQueue{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCallcenterQueue{}))}, onlyAdminGroup())
	//Request:{"event":"RenameCallcenterQueue","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5,"name":"new_queue2"}}
	//Response:{"MessageType":"RenameCallcenterQueue","data":{"id":5,"position":5,"enabled":true,"name":"new_queue2","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "RenameCallcenterQueue":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterQueue{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterQueue","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5}}
	//Response:{"MessageType":"DelCallcenterQueue","data":{"id":5,"position":5,"enabled":true,"name":"new_queue2","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCallcenterQueue":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCallcenterQueue{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"ImportCallcenterAgentsAndTiers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"ImportCallcenterAgentsAndTiers","data":{"callcenter_agents":{"items":null,"total":0},"callcenter_tiers":{"items":[{"id":3,"agent":"1007@45.61.54.76","queue":"n","level":4,"position":4,"state":"Ready"}],"total":1}}}
	//Errors:
	case "ImportCallcenterAgentsAndTiers":
		//TODO: replace
		getUser(msg, ImportCallcenterAgentsAdnTiers, onlyAdminGroup())
		msg.DBRequest = mainStruct.DBRequest{Limit: 25}
		resp1 := getUserForConfig(msg, getConfig, &altStruct.Agent{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.Tier{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event,
			Data: struct {
				S interface{} `json:"callcenter_agents"`
				P interface{} `json:"callcenter_tiers"`
			}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"GetCallcenterAgents","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","db_request":{"limit":25,"offset":0,"filters":[],"order":{"fields":[],"desc":false}}}}
	//Response:{"MessageType":"GetCallcenterAgents","data":{"items":[{"id":34,"name":"30","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":3,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10},{"id":32,"name":"28","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":7,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10},{"id":6,"name":"2","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":4,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10},{"id":7,"name":"1000@45.61.54.76","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":0,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}],"total":30}}
	//Errors:
	case "GetCallcenterAgents":
		resp = getUserForConfig(msg, getConfig, &altStruct.Agent{}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterAgent","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_agent"}}
	//Response:{"MessageType":"AddCallcenterAgent","data":{"id":35,"name":"new_agent","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":0,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}}
	//Errors:
	case "AddCallcenterAgent":
		resp = getUserForConfig(msg, setConfig, &altStruct.Agent{
			Name:              msg.Name,
			Type:              "callback",
			System:            "single_box",
			InstanceId:        "single_box",
			Status:            "On Break",
			State:             "Waiting",
			WrapUpTime:        10,
			ReadyTime:         10,
			BusyDelayTime:     10,
			NoAnswerDelayTime: 10,
		}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterAgent","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":37,"name":"max_no_answer","value":"5"}}}
	//Response:{"MessageType":"UpdateCallcenterAgent","data":{"id":37,"name":"agent@domain","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":5,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}}
	//Errors:
	case "UpdateCallcenterAgent":
		resp = getCallcenterAgents(msg)
		/*
					switch name {
					//Title:
			//Request:
			//Response:
			//Errors:
			case "state":
						eventChannel <- &map[int64]*mainStruct.Agent{agent.Id: agent}
					//Title:
			//Request:
			//Response:
			//Errors:
			case "status":
						agent.LastStatusChange = time.Now().Unix()
						eventChannel <- &map[int64]*mainStruct.Agent{agent.Id: agent}
					}
		*/
		// resp = getUser(msg, UpdateCallcenterAgent, onlyAdminGroup())
	//Request:{"event":"DelCallcenterAgent","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":3}}
	//Response:{"MessageType":"DelCallcenterAgent","data":{"id":37,"name":"agent@domain","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":5,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}}
	//Errors:
	case "DelCallcenterAgent":
		resp = getUserForConfig(msg, delConfig, &altStruct.Agent{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterTiers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","db_request":{"limit":25,"offset":0,"filters":[],"order":{"fields":[],"desc":false}}}}
	//Response:{"MessageType":"GetCallcenterTiers","data":{"items":[{"id":4,"agent":"agent","queue":"ddaaaaw","level":1,"position":1,"state":"Ready"}],"total":1}}
	//Errors:
	case "GetCallcenterTiers":
		resp = getUserForConfig(msg, getConfig, &altStruct.Tier{}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterTier","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"name":"new_agent"}}
	//Response:{"MessageType":"AddCallcenterTier","data":{"id":5,"agent":"new_agent","queue":"ddaaaaw","level":1,"position":1,"state":"Ready"}}
	//Errors:
	case "AddCallcenterTier":
		queueI, err := intermediateDB.GetByIdArg(&altStruct.ConfigCallcenterQueue{}, msg.Id)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: msg.Event}
		}
		queue, ok := queueI.(altStruct.ConfigCallcenterQueue)
		if !ok {
			return webStruct.UserResponse{Error: "queue not found", MessageType: msg.Event}
		}
		resp = getUserForConfig(msg, setConfig, &altStruct.Tier{Queue: queue.Name, Agent: msg.Name, State: "Ready", Position: 1, Level: 1}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterTier","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5,"name":"level","value":"2"}}}
	//Response:{"MessageType":"UpdateCallcenterTier","data":{"id":5,"agent":"new_agent","queue":"ddaaaaw","level":2,"position":7,"state":"Ready"}}
	//Errors:
	case "UpdateCallcenterTier":
		resp = getCallcenterTiers(msg)
	//Request:{"event":"DelCallcenterTier","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5}}
	//Response:{"MessageType":"DelCallcenterTier","data":{"id":5,"agent":"new_agent","queue":"ddaaaaw","level":2,"position":7,"state":"Ready"}}
	//Errors:
	case "DelCallcenterTier":
		resp = getUserForConfig(msg, delConfig, &altStruct.Tier{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterMembers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","db_request":{"limit":25,"offset":0,"filters":[],"order":{"fields":[],"desc":false}}}}
	//Response:{"MessageType":"GetCallcenterMembers","data":{"items":null,"total":0}}
	//Errors:
	case "GetCallcenterMembers":
		resp = getUserForConfig(msg, getConfig, &altStruct.Member{}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterMember","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5}}
	//Response:{"MessageType":"DelCallcenterMember","data":{"id":5}}
	//Errors:
	case "DelCallcenterMember":
		resp = getUserForConfig(msg, delConfig, &altStruct.Member{Uuid: msg.Uuid}, onlyAdminGroup())
	//Request:{"event":"SendCallcenterCommand","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"load","id":2}}
	//Response:{"MessageType":"SendCallcenterCommand"}
	//Errors:
	case "SendCallcenterCommand":
		resp = getUser(msg, runCallcenterQueueCommand, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case webStruct.SubscribeCallcenterAgents:
		//TODO: replace
		resp = getUserForConfig(msg, getConfig, &altStruct.Agent{}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case webStruct.SubscribeCallcenterTiers:
		//TODO: replace
		resp = getUserForConfig(msg, getConfig, &altStruct.Tier{}, onlyAdminGroup())
	//### LCR
	//Request:{"event":"GetLcr","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetLcr","data":{"settings":{},"profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"qual_rel","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"rel_qual","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetLcr":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigLcrSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigLcrProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"GetLcrProfileParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetLcrProfileParameters","data":{"1":{"id":1,"position":1,"enabled":true,"name":"id","value":"0","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"order_by","value":"rate,quality,reliability","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetLcrProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigLcrProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"UpdateLcrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateLcrParameter","data":{"id":7,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateLcrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLcrSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchLcrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7,"enabled":false}}}
	//Response:{"MessageType":"SwitchLcrParameter","data":{"id":7,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchLcrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLcrSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddLcrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddLcrParameter","data":{"id":7,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddLcrParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigLcrSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigLcrSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelLcrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7}}}
	//Response:{"MessageType":"DelLcrParameter","data":{"id":7,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelLcrParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigLcrSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"AddLcrProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddLcrProfileParameter","data":{"id":16,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddLcrProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigLcrProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigLcrProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelLcrProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":16}}}
	//Response:{"MessageType":"DelLcrProfileParameter","data":{"id":16,"position":3,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelLcrProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigLcrProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchLcrProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":16,"enabled":false}}}
	//Response:{"MessageType":"SwitchLcrProfileParameter","data":{"id":16,"position":3,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchLcrProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLcrProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateLcrProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":16,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateLcrProfileParameter","data":{"id":16,"position":3,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateLcrProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLcrProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddLcrProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"AddLcrProfile","data":{"id":10,"position":4,"enabled":true,"name":"new_profile","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddLcrProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigLcrProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigLcrProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateLcrProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":10,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateLcrProfile","data":{"id":10,"position":4,"enabled":true,"name":"new_profile2","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateLcrProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLcrProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelLcrProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"DelLcrProfile","data":{"id":10,"position":4,"enabled":true,"name":"new_profile2","description":"","parent":{"id":24,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelLcrProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigLcrProfile{Id: msg.Id}, onlyAdminGroup())
	//### Shout
	//Request:{"event":"GetShout","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetShout","data":{"2":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":40,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetShout":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigShoutSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateShoutParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateShoutParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":40,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateShoutParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigShoutSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchShoutParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchShoutParameter","data":{"id":2,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":40,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchShoutParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigShoutSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddShoutParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddShoutParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":40,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddShoutParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigShoutSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigShoutSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelShoutParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"DelShoutParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":40,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelShoutParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigShoutSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Redis
	//Request:{"event":"GetRedis","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetRedis","data":{"1":{"id":1,"position":1,"enabled":true,"name":"host","value":"localhost","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"port","value":"6379","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"timeout","value":"10000","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetRedis":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigRedisSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateRedisParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"name":"timeout","value":"100000"}}}
	//Response:{"MessageType":"UpdateRedisParameter","data":{"id":3,"position":3,"enabled":true,"name":"timeout","value":"100000","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateRedisParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigRedisSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchRedisParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchRedisParameter","data":{"id":3,"position":3,"enabled":false,"name":"timeout","value":"100000","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchRedisParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigRedisSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddRedisParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddRedisParameter","data":{"id":5,"position":4,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddRedisParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigRedisSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigRedisSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelRedisParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelRedisParameter","data":{"id":5,"position":4,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelRedisParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigRedisSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Nibblebill
	//Request:{"event":"GetNibblebill","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetNibblebill","data":{"1":{"id":1,"position":1,"enabled":true,"name":"odbc-dsn","value":"bandwidth.com","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"percall_max_amt","value":"100","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"percall_action","value":"hangup","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"db_table","value":"accounts","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"db_column_cash","value":"cash","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"db_column_account","value":"id","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"global_heartbeat","value":"60","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"lowbal_amt","value":"5","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"lowbal_action","value":"play ding","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"nobal_amt","value":"0","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"nobal_action","value":"hangup","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetNibblebill":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigNibblebillSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateNibblebillParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10,"name":"percall_max_amt","value":"1000"}}}
	//Response:{"MessageType":"UpdateNibblebillParameter","data":{"id":10,"position":10,"enabled":true,"name":"percall_max_amt","value":"1000","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateNibblebillParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigNibblebillSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchNibblebillParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":11,"enabled":false}}}
	//Response:{"MessageType":"SwitchNibblebillParameter","data":{"id":11,"position":11,"enabled":false,"name":"percall_action","value":"hangup","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchNibblebillParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigNibblebillSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddNibblebillParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddNibblebillParameter","data":{"id":14,"position":12,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddNibblebillParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigNibblebillSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigNibblebillSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelNibblebillParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14}}}
	//Response:{"MessageType":"DelNibblebillParameter","data":{"id":14,"position":12,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":29,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelNibblebillParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigNibblebillSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### DB
	//Request:{"event":"GetDb","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDb","data":{"3":{"id":3,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":14,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetDb":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDbSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateDbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateDbParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":14,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateDbParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDbSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchDbParameter","data":{"id":3,"position":1,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":14,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchDbParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDbSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddDbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddDbParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":14,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddDbParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDbSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigDbSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelDbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3}}}
	//Response:{"MessageType":"DelDbParameter","data":{"id":3,"position":1,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":14,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelDbParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDbSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Memcache
	//Request:{"event":"GetMemcache","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetMemcache","data":{"1":{"id":1,"position":1,"enabled":true,"name":"memcache-servers","value":"localhost","description":"","parent":{"id":26,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetMemcache":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigMemcacheSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateMemcacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"memcache-servers","value":"localhost"}}}
	//Response:{"MessageType":"UpdateMemcacheParameter","data":{"id":1,"position":1,"enabled":true,"name":"memcache-servers","value":"localhost","description":"","parent":{"id":26,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateMemcacheParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigMemcacheSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchMemcacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchMemcacheParameter","data":{"id":1,"position":1,"enabled":false,"name":"memcache-servers","value":"localhost","description":"","parent":{"id":26,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchMemcacheParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigMemcacheSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddMemcacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddMemcacheParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":26,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddMemcacheParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigMemcacheSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigMemcacheSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelMemcacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4}}}
	//Response:{"MessageType":"DelMemcacheParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":26,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelMemcacheParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigMemcacheSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Avmd
	//Request:{"event":"GetAvmd","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetAvmd","data":{"1":{"id":1,"position":1,"enabled":true,"name":"debug","value":"0","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"inbound_channel","value":"0","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"outbound_channel","value":"1","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"12":{"id":12,"position":12,"enabled":true,"name":"detection_mode","value":"2","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"13":{"id":13,"position":13,"enabled":true,"name":"detectors_n","value":"36","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"14":{"id":14,"position":14,"enabled":true,"name":"detectors_lagged_n","value":"1","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"report_status","value":"1","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"fast_math","value":"0","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"require_continuous_streak","value":"1","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"sample_n_continuous_streak","value":"3","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"sample_n_to_skip","value":"0","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"require_continuous_streak_amp","value":"1","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"sample_n_continuous_streak_amp","value":"3","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"simplified_estimation","value":"1","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetAvmd":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigAvmdSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateAvmdParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14,"name":"detectors_lagged_n","value":"2"}}}
	//Response:{"MessageType":"UpdateAvmdParameter","data":{"id":14,"position":14,"enabled":true,"name":"detectors_lagged_n","value":"2","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateAvmdParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAvmdSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchAvmdParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14,"enabled":false}}}
	//Response:{"MessageType":"SwitchAvmdParameter","data":{"id":14,"position":14,"enabled":false,"name":"detectors_lagged_n","value":"2","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchAvmdParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAvmdSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddAvmdParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddAvmdParameter","data":{"id":17,"position":15,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddAvmdParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigAvmdSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigAvmdSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelAvmdParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":17}}}
	//Response:{"MessageType":"DelAvmdParameter","data":{"id":17,"position":15,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelAvmdParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigAvmdSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Tts Commandline
	//Request:{"event":"GetTtsCommandline","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetTtsCommandline","data":{"1":{"id":1,"position":1,"enabled":true,"name":"command","value":"echo ${text} | text2wave -f ${rate} \u003e ${file}","description":"","parent":{"id":44,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetTtsCommandline":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigTtsCommandlineSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateTtsCommandlineParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"command","value":"echo ${text} | text2wave -f ${rate} > ${file}"}}}
	//Response:{"MessageType":"UpdateTtsCommandlineParameter","data":{"id":1,"position":1,"enabled":true,"name":"command","value":"echo ${text} | text2wave -f ${rate} \u003e ${file}","description":"","parent":{"id":44,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateTtsCommandlineParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigTtsCommandlineSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchTtsCommandlineParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchTtsCommandlineParameter","data":{"id":1,"position":1,"enabled":false,"name":"command","value":"echo ${text} | text2wave -f ${rate} \u003e ${file}","description":"","parent":{"id":44,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchTtsCommandlineParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigTtsCommandlineSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddTtsCommandlineParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddTtsCommandlineParameter","data":{"id":2,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":44,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddTtsCommandlineParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigTtsCommandlineSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigTtsCommandlineSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelTtsCommandlineParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"DelTtsCommandlineParameter","data":{"id":2,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":44,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelTtsCommandlineParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigTtsCommandlineSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### CDR Mongodb
	//Request:{"event":"GetCdrMongodb","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetCdrMongodb","data":{"1":{"id":1,"position":1,"enabled":true,"name":"host","value":"127.0.0.1","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"port","value":"27017","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"namespace","value":"test.cdr","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"log-b-leg","value":"false","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCdrMongodb":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCdrMongodbSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateCdrMongodbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"log-b-leg","value":"true"}}}
	//Response:{"MessageType":"UpdateCdrMongodbParameter","data":{"id":4,"position":4,"enabled":true,"name":"log-b-leg","value":"true","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateCdrMongodbParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCdrMongodbSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchCdrMongodbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchCdrMongodbParameter","data":{"id":4,"position":4,"enabled":false,"name":"log-b-leg","value":"true","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchCdrMongodbParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCdrMongodbSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddCdrMongodbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCdrMongodbParameter","data":{"id":7,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCdrMongodbParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCdrMongodbSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCdrMongodbSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelCdrMongodbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7}}}
	//Response:{"MessageType":"DelCdrMongodbParameter","data":{"id":7,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCdrMongodbParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCdrMongodbSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### HTTP Cache
	//Request:{"event":"GetHttpCache","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetHttpCache","data":{"1":{"id":1,"position":1,"enabled":true,"name":"enable-file-formats","value":"false","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"max-urls","value":"10000","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"location","value":"/var/cache/freeswitch","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"default-max-age","value":"86400","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"prefetch-thread-count","value":"8","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"prefetch-queue-size","value":"100","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"ssl-cacert","value":"/etc/freeswitch/tls/cacert.pem","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"ssl-verifypeer","value":"true","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"ssl-verifyhost","value":"true","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetHttpCache":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateHttpCacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":9,"name":"ssl-verifyhost","value":"false"}}}
	//Response:{"MessageType":"UpdateHttpCacheParameter","data":{"id":9,"position":9,"enabled":true,"name":"ssl-verifyhost","value":"false","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateHttpCacheParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchHttpCacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":9,"enabled":false}}}
	//Response:{"MessageType":"SwitchHttpCacheParameter","data":{"id":9,"position":9,"enabled":false,"name":"ssl-verifyhost","value":"false","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchHttpCacheParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddHttpCacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddHttpCacheParameter","data":{"id":12,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddHttpCacheParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigHttpCacheSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigHttpCacheSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelHttpCacheParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":12}}}
	//Response:{"MessageType":"DelHttpCacheParameter","data":{"id":12,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelHttpCacheParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigHttpCacheSetting{Id: msg.Param.Id}, onlyAdminGroup())
	case "GetHttpCacheProfile":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfile{}, onlyAdminGroup())
	case "AddHttpCacheProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigHttpCacheProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigHttpCacheProfile{}))}, onlyAdminGroup())
	case "RenameHttpCacheProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	case "DelHttpCacheProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigHttpCacheProfile{Id: msg.Id}, onlyAdminGroup())
	case "GetHttpCacheProfileParameters":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfileDomain{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfileAzureBlob{}, onlyAdminGroup())
		resp3 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfileAWSS3{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"domains"`
			P interface{} `json:"azure"`
			R interface{} `json:"aws_s3"`
		}{S: resp1.Data, P: resp2.Data, R: resp3.Data}}
	case "AddHttpCacheProfileDomain":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigHttpCacheProfileDomain{Name: msg.Param.Name, Enabled: true, Parent: &altStruct.ConfigHttpCacheProfile{Id: msg.Id}}, onlyAdminGroup())
	case "DelHttpCacheProfileDomain":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigHttpCacheProfileDomain{Id: msg.Param.Id}, onlyAdminGroup())
	case "SwitchHttpCacheProfileDomain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileDomain{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	case "UpdateHttpCacheProfileDomain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileDomain{Id: msg.Param.Id, Name: msg.Param.Name}, []string{"Name"}}, onlyAdminGroup())
	case "UpdateHttpCacheProfileAws":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileAWSS3{Id: msg.AwsS3.Id,
			AccessKeyId:     msg.AwsS3.AccessKeyId,
			SecretAccessKey: msg.AwsS3.SecretAccessKey,
			BaseDomain:      msg.AwsS3.BaseDomain,
			Region:          msg.AwsS3.Region,
			Expires:         msg.AwsS3.Expires,
		}, []string{"AccessKeyId", "SecretAccessKey", "BaseDomain", "Region", "Expires"}}, onlyAdminGroup())
	case "UpdateHttpCacheProfileAzure":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileAzureBlob{Id: msg.Azure.Id,
			SecretAccessKey: msg.Azure.SecretAccessKey,
		}, []string{"SecretAccessKey"}}, onlyAdminGroup())

	//### Opus
	//Request:{"event":"GetOpus","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetOpus","data":{"1":{"id":1,"position":1,"enabled":true,"name":"use-vbr","value":"1","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"complexity","value":"10","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"keep-fec-enabled","value":"1","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"maxaveragebitrate","value":"0","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"maxplaybackrate","value":"0","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetOpus":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigOpusSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateOpusParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5,"name":"maxplaybackrate","value":"512"}}}
	//Response:{"MessageType":"UpdateOpusParameter","data":{"id":5,"position":5,"enabled":true,"name":"maxplaybackrate","value":"512","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOpusParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpusSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOpusParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5,"enabled":false}}}
	//Response:{"MessageType":"SwitchOpusParameter","data":{"id":5,"position":5,"enabled":false,"name":"maxplaybackrate","value":"512","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchOpusParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpusSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddOpusParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOpusParameter","data":{"id":8,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOpusParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOpusSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOpusSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelOpusParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":8}}}
	//Response:{"MessageType":"DelOpusParameter","data":{"id":8,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":31,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelOpusParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOpusSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Python
	//Request:{"event":"GetPython","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetPython","data":{"2":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":37,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetPython":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigPythonSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdatePythonParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdatePythonParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":37,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePythonParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPythonSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPythonParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchPythonParameter","data":{"id":2,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":37,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchPythonParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPythonSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddPythonParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddPythonParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":37,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddPythonParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPythonSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPythonSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelPythonParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"DelPythonParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":37,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPythonParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPythonSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### ALsa
	//Request:{"event":"GetAlsa","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetAlsa","data":{"16":{"id":16,"position":1,"enabled":true,"name":"dialplan","value":"XML","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"17":{"id":17,"position":2,"enabled":true,"name":"cid-name","value":"N800 Alsa","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"18":{"id":18,"position":3,"enabled":true,"name":"cid-num","value":"5555551212","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"19":{"id":19,"position":4,"enabled":true,"name":"sample-rate","value":"8000","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"20":{"id":20,"position":5,"enabled":true,"name":"codec-ms","value":"20","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetAlsa":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigAlsaSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateAlsaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":17,"name":"cid-name","value":"N800 Alsa2"}}}
	//Response:{"MessageType":"UpdateAlsaParameter","data":{"id":17,"position":2,"enabled":true,"name":"cid-name","value":"N800 Alsa2","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateAlsaParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAlsaSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchAlsaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":18,"enabled":false}}}
	//Response:{"MessageType":"SwitchAlsaParameter","data":{"id":18,"position":3,"enabled":false,"name":"cid-num","value":"5555551212","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchAlsaParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAlsaSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddAlsaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddAlsaParameter","data":{"id":21,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddAlsaParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigAlsaSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigAlsaSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelAlsaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":21}}}
	//Response:{"MessageType":"DelAlsaParameter","data":{"id":21,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":55,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelAlsaParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigAlsaSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Amr
	//Request:{"event":"GetAmr","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetAmr","data":{"10":{"id":10,"position":3,"enabled":true,"name":"adjust-bitrate","value":"0","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":4,"enabled":true,"name":"force-oa","value":"0","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":1,"enabled":true,"name":"default-bitrate","value":"7","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":2,"enabled":true,"name":"volte","value":"0","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetAmr":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigAmrSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateAmrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":11,"name":"force-oa","value":"1"}}}
	//Response:{"MessageType":"UpdateAmrParameter","data":{"id":11,"position":4,"enabled":true,"name":"force-oa","value":"1","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateAmrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAmrSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchAmrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":11,"enabled":false}}}
	//Response:{"MessageType":"SwitchAmrParameter","data":{"id":11,"position":4,"enabled":false,"name":"force-oa","value":"1","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchAmrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAmrSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddAmrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddAmrParameter","data":{"id":12,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddAmrParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigAmrSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigAmrSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelAmrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":12}}}
	//Response:{"MessageType":"DelAmrParameter","data":{"id":12,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":52,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelAmrParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigAmrSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Amrwb
	//Request:{"event":"GetAmrwb","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetAmrwb","data":{"1":{"id":1,"position":1,"enabled":true,"name":"default-bitrate","value":"8","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"volte","value":"1","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"adjust-bitrate","value":"0","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"force-oa","value":"0","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetAmrwb":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigAmrwbSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateAmrwbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"force-oa","value":"1"}}}
	//Response:{"MessageType":"UpdateAmrwbParameter","data":{"id":4,"position":4,"enabled":true,"name":"force-oa","value":"1","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateAmrwbParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAmrwbSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchAmrwbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchAmrwbParameter","data":{"id":4,"position":4,"enabled":false,"name":"force-oa","value":"1","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchAmrwbParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigAmrwbSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddAmrwbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddAmrwbParameter","data":{"id":10,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddAmrwbParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigAmrwbSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigAmrwbSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelAmrwbParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10}}}
	//Response:{"MessageType":"DelAmrwbParameter","data":{"id":10,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":4,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelAmrwbParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigAmrwbSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Cepstral
	//Request:{"event":"GetCepstral","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetCepstral","data":{"1":{"id":1,"position":1,"enabled":true,"name":"encoding","value":"utf-8","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCepstral":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCespalSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateCepstralParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"encoding","value":"utf-8"}}}
	//Response:{"MessageType":"UpdateCepstralParameter","data":{"id":1,"position":1,"enabled":true,"name":"encoding","value":"utf-8","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateCepstralParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCespalSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchCepstralParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchCepstralParameter","data":{"id":1,"position":1,"enabled":false,"name":"encoding","value":"utf-8","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchCepstralParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCespalSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddCepstralParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCepstralParameter","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCepstralParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCespalSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCespalSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelCepstralParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelCepstralParameter","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":9,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCepstralParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCespalSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Cidlookup
	//Request:{"event":"GetCidlookup","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetCidlookup","data":{"1":{"id":1,"position":1,"enabled":true,"name":"url","value":"http://query.voipcnam.com/query.php?api_key=MYAPIKEY\u0026number=${caller_id_number}","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"whitepages-apikey","value":"MYAPIKEY","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"cache","value":"true","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"cache-expire","value":"86400","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"odbc-dsn","value":"phone:phone:phone","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"sql","value":"      SELECT name||' ('||type||')' AS name        FROM phonebook p JOIN numbers n ON p.id = n.phonebook_id       WHERE n.number='${caller_id_number}'        LIMIT 1       ","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"citystate-sql","value":"      SELECT ratecenter||' '||state as name       FROM npa_nxx_company_ocn       WHERE npa = ${caller_id_number:1:3} AND nxx = ${caller_id_number:4:3}       LIMIT 1       ","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCidlookup":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCidlookupSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateCidlookupParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"cache-expire","value":"86421"}}}
	//Response:{"MessageType":"UpdateCidlookupParameter","data":{"id":4,"position":4,"enabled":true,"name":"cache-expire","value":"86421","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateCidlookupParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCidlookupSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchCidlookupParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchCidlookupParameter","data":{"id":4,"position":4,"enabled":false,"name":"cache-expire","value":"86421","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchCidlookupParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCidlookupSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddCidlookupParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCidlookupParameter","data":{"id":12,"position":8,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCidlookupParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCidlookupSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCidlookupSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelCidlookupParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":12}}}
	//Response:{"MessageType":"DelCidlookupParameter","data":{"id":12,"position":8,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCidlookupParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCidlookupSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Curl
	//Request:{"event":"GetCurl","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetCurl","data":{"1":{"id":1,"position":1,"enabled":true,"name":"max-bytes","value":"64000","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCurl":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCurlSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateCurlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"max-bytes","value":"64001"}}}
	//Response:{"MessageType":"UpdateCurlParameter","data":{"id":1,"position":1,"enabled":true,"name":"max-bytes","value":"64001","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateCurlParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCurlSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchCurlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchCurlParameter","data":{"id":1,"position":1,"enabled":false,"name":"max-bytes","value":"64001","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchCurlParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCurlSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddCurlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCurlParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCurlParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCurlSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCurlSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelCurlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4}}}
	//Response:{"MessageType":"DelCurlParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCurlParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCurlSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### DDialplan Directory
	//Request:{"event":"GetDialplanDirectory","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDialplanDirectory","data":{"1":{"id":1,"position":1,"enabled":true,"name":"directory-name","value":"ldap","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"host","value":"ldap.freeswitch.org","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"dn","value":"cn=Manager,dc=freeswitch,dc=org","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"pass","value":"test","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"base","value":"dc=freeswitch,dc=org","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetDialplanDirectory":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDialplanDirectorySetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateDialplanDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"pass","value":"test2"}}}
	//Response:{"MessageType":"UpdateDialplanDirectoryParameter","data":{"id":4,"position":4,"enabled":true,"name":"pass","value":"test2","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateDialplanDirectoryParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDialplanDirectorySetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDialplanDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchDialplanDirectoryParameter","data":{"id":4,"position":4,"enabled":false,"name":"pass","value":"test2","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchDialplanDirectoryParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDialplanDirectorySetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddDialplanDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddDialplanDirectoryParameter","data":{"id":8,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddDialplanDirectoryParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDialplanDirectorySetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigDialplanDirectorySetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelDialplanDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":8}}}
	//Response:{"MessageType":"DelDialplanDirectoryParameter","data":{"id":8,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelDialplanDirectoryParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDialplanDirectorySetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Easyroute
	//Request:{"event":"GetEasyroute","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetEasyroute","data":{"1":{"id":1,"position":1,"enabled":true,"name":"db-username","value":"root","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"db-password","value":"password","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"db-dsn","value":"easyroute","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"default-techprofile","value":"sofia/default","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"default-gateway","value":"192.168.66.6","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"odbc-retries","value":"120","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetEasyroute":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigEasyrouteSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateEasyrouteParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":6,"name":"odbc-retries","value":"125"}}}
	//Response:{"MessageType":"UpdateEasyrouteParameter","data":{"id":6,"position":6,"enabled":true,"name":"odbc-retries","value":"125","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateEasyrouteParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigEasyrouteSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchEasyrouteParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":6,"enabled":false}}}
	//Response:{"MessageType":"SwitchEasyrouteParameter","data":{"id":6,"position":6,"enabled":false,"name":"odbc-retries","value":"125","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchEasyrouteParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigEasyrouteSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddEasyrouteParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddEasyrouteParameter","data":{"id":9,"position":7,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddEasyrouteParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigEasyrouteSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigEasyrouteSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelEasyrouteParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":9}}}
	//Response:{"MessageType":"DelEasyrouteParameter","data":{"id":9,"position":7,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":18,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelEasyrouteParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigEasyrouteSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Erlang Event
	//Request:{"event":"GetErlangEvent","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetErlangEvent","data":{"1":{"id":1,"position":1,"enabled":true,"name":"listen-ip","value":"0.0.0.0","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"listen-port","value":"8031","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"nodename","value":"freeswitch","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"cookie","value":"ClueCon","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"shortname","value":"true","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetErlangEvent":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigErlangEventSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateErlangEventParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5,"name":"shortname","value":"false"}}}
	//Response:{"MessageType":"UpdateErlangEventParameter","data":{"id":5,"position":5,"enabled":true,"name":"shortname","value":"false","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateErlangEventParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigErlangEventSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchErlangEventParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5,"enabled":false}}}
	//Response:{"MessageType":"SwitchErlangEventParameter","data":{"id":5,"position":5,"enabled":false,"name":"shortname","value":"false","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchErlangEventParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigErlangEventSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddErlangEventParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddErlangEventParameter","data":{"id":10,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddErlangEventParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigErlangEventSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigErlangEventSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelErlangEventParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10}}}
	//Response:{"MessageType":"DelErlangEventParameter","data":{"id":10,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":19,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelErlangEventParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigErlangEventSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Event Multicast
	//Request:{"event":"GetEventMulticast","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetEventMulticast","data":{"1":{"id":1,"position":1,"enabled":true,"name":"address","value":"225.1.1.1","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"port","value":"4242","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"bindings","value":"all","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"ttl","value":"1","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetEventMulticast":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigEventMulticastSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateEventMulticastParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"ttl","value":"2"}}}
	//Response:{"MessageType":"UpdateEventMulticastParameter","data":{"id":4,"position":4,"enabled":true,"name":"ttl","value":"2","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateEventMulticastParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigEventMulticastSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchEventMulticastParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchEventMulticastParameter","data":{"id":4,"position":4,"enabled":false,"name":"ttl","value":"2","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchEventMulticastParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigEventMulticastSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddEventMulticastParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddEventMulticastParameter","data":{"id":7,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddEventMulticastParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigEventMulticastSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigEventMulticastSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelEventMulticastParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7}}}
	//Response:{"MessageType":"DelEventMulticastParameter","data":{"id":7,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":20,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelEventMulticastParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigEventMulticastSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Fax
	//Request:{"event":"GetFax","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetFax","data":{"1":{"id":1,"position":1,"enabled":true,"name":"use-ecm","value":"true","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"verbose","value":"false","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"disable-v17","value":"false","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"ident","value":"SpanDSP Fax Ident","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"header","value":"SpanDSP Fax Header","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"spool-dir","value":"/tmp","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"file-prefix","value":"faxrx","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetFax":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigFaxSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateFaxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"verbose","value":"true"}}}
	//Response:{"MessageType":"UpdateFaxParameter","data":{"id":2,"position":2,"enabled":true,"name":"verbose","value":"true","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateFaxParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFaxSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchFaxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchFaxParameter","data":{"id":3,"position":3,"enabled":false,"name":"disable-v17","value":"false","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchFaxParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFaxSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddFaxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddFaxParameter","data":{"id":10,"position":8,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddFaxParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigFaxSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigFaxSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelFaxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10}}}
	//Response:{"MessageType":"DelFaxParameter","data":{"id":10,"position":8,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":21,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelFaxParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigFaxSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Lua
	//Request:{"event":"GetLua","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetLua","data":{"3":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":25,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetLua":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigLuaSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateLuaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateLuaParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":25,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateLuaParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLuaSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchLuaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchLuaParameter","data":{"id":3,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":25,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchLuaParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigLuaSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddLuaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddLuaParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":25,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddLuaParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigLuaSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigLuaSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelLuaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3}}}
	//Response:{"MessageType":"DelLuaParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":25,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelLuaParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigLuaSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Mongo
	//Request:{"event":"GetMongo","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetMongo","data":{"1":{"id":1,"position":1,"enabled":true,"name":"connection-string","value":"mongodb://127.0.0.1:27017/?connectTimeoutMS=10000","description":"","parent":{"id":27,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetMongo":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigMongoSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateMongoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"connection-string","value":"mongodb://127.0.0.1:27017/?connectTimeoutMS=10001"}}}
	//Response:{"MessageType":"UpdateMongoParameter","data":{"id":1,"position":1,"enabled":true,"name":"connection-string","value":"mongodb://127.0.0.1:27017/?connectTimeoutMS=10001","description":"","parent":{"id":27,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateMongoParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigMongoSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchMongoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchMongoParameter","data":{"id":1,"position":1,"enabled":false,"name":"connection-string","value":"mongodb://127.0.0.1:27017/?connectTimeoutMS=10001","description":"","parent":{"id":27,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchMongoParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigMongoSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddMongoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddMongoParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":27,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddMongoParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigMongoSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigMongoSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelMongoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4}}}
	//Response:{"MessageType":"DelMongoParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":27,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelMongoParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigMongoSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Msrp
	//Request:{"event":"GetMsrp","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetMsrp","data":{"1":{"id":1,"position":1,"enabled":true,"name":"listen-ip","value":"45.61.54.76","description":"","parent":{"id":28,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetMsrp":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigMsrpSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateMsrpParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"listen-ip","value":"127.0.0.1"}}}
	//Response:{"MessageType":"UpdateMsrpParameter","data":{"id":1,"position":1,"enabled":true,"name":"listen-ip","value":"127.0.0.1","description":"","parent":{"id":28,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateMsrpParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigMsrpSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchMsrpParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchMsrpParameter","data":{"id":1,"position":1,"enabled":false,"name":"listen-ip","value":"127.0.0.1","description":"","parent":{"id":28,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchMsrpParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigMsrpSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddMsrpParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddMsrpParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":28,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddMsrpParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigMsrpSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigMsrpSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelMsrpParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4}}}
	//Response:{"MessageType":"DelMsrpParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":28,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelMsrpParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigMsrpSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Oreka
	//Request:{"event":"GetOreka","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetOreka","data":{"3":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":32,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetOreka":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigOrekaSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateOrekaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateOrekaParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":32,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOrekaParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOrekaSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOrekaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchOrekaParameter","data":{"id":3,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":32,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchOrekaParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOrekaSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddOrekaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOrekaParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":32,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOrekaParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOrekaSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOrekaSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelOrekaParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3}}}
	//Response:{"MessageType":"DelOrekaParameter","data":{"id":3,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":32,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelOrekaParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOrekaSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Perl
	//Request:{"event":"GetPerl","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetPerl","data":{"2":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":34,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetPerl":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigPerlSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdatePerlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdatePerlParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":34,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePerlParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPerlSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPerlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchPerlParameter","data":{"id":2,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":34,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchPerlParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPerlSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddPerlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddPerlParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":34,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddPerlParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPerlSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPerlSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelPerlParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"DelPerlParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":34,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPerlParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPerlSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Pocketsphinx
	//Request:{"event":"GetPocketsphinx","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetPocketsphinx","data":{"1":{"id":1,"position":1,"enabled":true,"name":"threshold","value":"400","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"silence-hits","value":"25","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"listen-hits","value":"1","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"auto-reload","value":"true","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetPocketsphinx":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigPocketsphinxSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdatePocketsphinxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"name":"listen-hits","value":"2"}}}
	//Response:{"MessageType":"UpdatePocketsphinxParameter","data":{"id":3,"position":3,"enabled":true,"name":"listen-hits","value":"2","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePocketsphinxParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPocketsphinxSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPocketsphinxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchPocketsphinxParameter","data":{"id":4,"position":4,"enabled":false,"name":"auto-reload","value":"true","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchPocketsphinxParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPocketsphinxSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddPocketsphinxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddPocketsphinxParameter","data":{"id":6,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddPocketsphinxParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPocketsphinxSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPocketsphinxSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelPocketsphinxParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":6}}}
	//Response:{"MessageType":"DelPocketsphinxParameter","data":{"id":6,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":35,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPocketsphinxParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPocketsphinxSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Sangoma Codec
	//Request:{"event":"GetSangomaCodec","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetSangomaCodec","data":{"2":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetSangomaCodec":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSangomaCodecSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateSangomaCodecParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"UpdateSangomaCodecParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateSangomaCodecParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSangomaCodecSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchSangomaCodecParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchSangomaCodecParameter","data":{"id":2,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchSangomaCodecParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSangomaCodecSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddSangomaCodecParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddSangomaCodecParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddSangomaCodecParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSangomaCodecSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigSangomaCodecSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelSangomaCodecParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"DelSangomaCodecParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":38,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelSangomaCodecParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSangomaCodecSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Sndfile
	//Request:{"event":"GetSndfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetSndfile","data":{"2":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":41,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetSndfile":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigSndfileSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateSndfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateSndfileParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":41,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateSndfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSndfileSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchSndfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchSndfileParameter","data":{"id":2,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":41,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchSndfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigSndfileSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddSndfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddSndfileParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":41,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddSndfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigSndfileSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigSndfileSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelSndfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2}}}
	//Response:{"MessageType":"DelSndfileParameter","data":{"id":2,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":41,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelSndfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigSndfileSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### XML CDR
	//Request:{"event":"GetXmlCdr","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetXmlCdr","data":{"1":{"id":1,"position":1,"enabled":true,"name":"log-dir","value":"","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"log-b-leg","value":"false","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"prefix-a-leg","value":"true","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"encode","value":"true","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetXmlCdr":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigXmlCdrSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateXmlCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"encode","value":"false"}}}
	//Response:{"MessageType":"UpdateXmlCdrParameter","data":{"id":4,"position":4,"enabled":true,"name":"encode","value":"false","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateXmlCdrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigXmlCdrSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchXmlCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchXmlCdrParameter","data":{"id":4,"position":4,"enabled":false,"name":"encode","value":"false","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchXmlCdrParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigXmlCdrSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddXmlCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddXmlCdrParameter","data":{"id":5,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddXmlCdrParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigXmlCdrSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigXmlCdrSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelXmlCdrParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelXmlCdrParameter","data":{"id":5,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":48,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelXmlCdrParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigXmlCdrSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### XML RPC
	//Request:{"event":"GetXmlRpc","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetXmlRpc","data":{"1":{"id":1,"position":1,"enabled":true,"name":"http-port","value":"8080","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"auth-realm","value":"freeswitch","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"auth-user","value":"freeswitch","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"auth-pass","value":"works","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetXmlRpc":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigXmlRpcSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateXmlRpcParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":true,"name":"auth-user","value":"freeswitch2"}}}
	//Response:{"MessageType":"UpdateXmlRpcParameter","data":{"id":3,"position":3,"enabled":true,"name":"auth-user","value":"freeswitch2","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateXmlRpcParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigXmlRpcSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchXmlRpcParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchXmlRpcParameter","data":{"id":4,"position":4,"enabled":false,"name":"auth-pass","value":"works","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchXmlRpcParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigXmlRpcSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddXmlRpcParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddXmlRpcParameter","data":{"id":5,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddXmlRpcParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigXmlRpcSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigXmlRpcSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelXmlRpcParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelXmlRpcParameter","data":{"id":5,"position":5,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":49,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelXmlRpcParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigXmlRpcSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Zeroconf
	//Request:{"event":"GetZeroconf","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetZeroconf","data":{"1":{"id":1,"position":1,"enabled":true,"name":"publish","value":"yes","description":"","parent":{"id":50,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"browse","value":"_sip._udp","description":"","parent":{"id":50,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetZeroconf":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigZeroconfSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateZeroconfParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"publish","value":"no"}}}
	//Response:{"MessageType":"UpdateZeroconfParameter","data":{"id":1,"position":1,"enabled":true,"name":"publish","value":"no","description":"","parent":{"id":50,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateZeroconfParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigZeroconfSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchZeroconfParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchZeroconfParameter","data":{"id":1,"position":1,"enabled":false,"name":"publish","value":"no","description":"","parent":{"id":50,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchZeroconfParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigZeroconfSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddZeroconfParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddZeroconfParameter","data":{"id":3,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":50,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddZeroconfParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigZeroconfSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigZeroconfSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelZeroconfParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3}}}
	//Response:{"MessageType":"DelZeroconfParameter","data":{"id":3,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":50,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelZeroconfParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigZeroconfSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//### Post Load Switch
	//Request:{"event":"GetPostSwitch","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetPostSwitch","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"colorize-console","value":"true","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"dump-cores","value":"yes","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"rtp-enable-zrtp","value":"false","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"dialplan-timestamps","value":"false","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"max-db-handles","value":"50","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"db-handle-timeout","value":"10","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"max-sessions","value":"1000","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"sessions-per-second","value":"30","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"loglevel","value":"debug","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"mailer-app","value":"sendmail","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"mailer-app-args","value":"-t","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"cli_keybinding":{"1":{"id":1,"position":1,"enabled":true,"name":"1","value":"help","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"10","value":"sofia profile internal siptrace on","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"11","value":"sofia profile internal siptrace off","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"12":{"id":12,"position":12,"enabled":true,"name":"12","value":"version","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"2","value":"status","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"3","value":"show channels","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"4","value":"show calls","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"5","value":"sofia status","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"6","value":"reloadxml","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"7","value":"console loglevel 0","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"8","value":"console loglevel 7","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"9","value":"sofia status profile internal","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"default_ptime":{}}}
	//Errors:
	case "GetPostSwitch":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigPostLoadSwitchSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigPostLoadSwitchCliKeybinding{}, onlyAdminGroup())
		resp3 := getUserForConfig(msg, getConfig, &altStruct.ConfigPostLoadSwitchDefaultPtime{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			C interface{} `json:"cli_keybinding"`
			P interface{} `json:"default_ptime"`
		}{S: resp1.Data, C: resp2.Data, P: resp3.Data}}
	//Request:{"event":"UpdatePostSwitchParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":6,"enabled":true}}}
	//Response:{"MessageType":"UpdatePostSwitchParameter","data":{"id":6,"position":6,"enabled":true,"name":"sessions-per-second","value":"35","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePostSwitchParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadSwitchSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPostSwitchParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":6,"enabled":false}}}
	//Response:{"MessageType":"SwitchPostSwitchParameter","data":{"id":6,"position":6,"enabled":false,"name":"sessions-per-second","value":"35","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchPostSwitchParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadSwitchSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddPostSwitchParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"enabled":true,"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddPostSwitchParameter","data":{"id":13,"position":12,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddPostSwitchParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPostLoadSwitchSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadSwitchSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelPostSwitchParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":13}}}
	//Response:{"MessageType":"DelPostSwitchParameter","data":{"id":13,"position":12,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPostSwitchParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPostLoadSwitchSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"UpdatePostSwitchCliKeybinding","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7,"enabled":true}}}
	//Response:{"MessageType":"UpdatePostSwitchCliKeybinding","data":{"id":7,"position":7,"enabled":true,"name":"7","value":"console loglevel 1","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePostSwitchCliKeybinding":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadSwitchCliKeybinding{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPostSwitchCliKeybinding","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7,"enabled":false}}}
	//Response:{"MessageType":"SwitchPostSwitchCliKeybinding","data":{"id":7,"position":7,"enabled":false,"name":"7","value":"console loglevel 1","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchPostSwitchCliKeybinding":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadSwitchCliKeybinding{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddPostSwitchCliKeybinding","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"enabled":true,"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddPostSwitchCliKeybinding","data":{"id":15,"position":13,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddPostSwitchCliKeybinding":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPostLoadSwitchCliKeybinding{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadSwitchCliKeybinding{}))}, onlyAdminGroup())
	//Request:{"event":"DelPostSwitchCliKeybinding","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":15}}}
	//Response:{"MessageType":"DelPostSwitchCliKeybinding","data":{"id":15,"position":13,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPostSwitchCliKeybinding":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPostLoadSwitchCliKeybinding{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"UpdatePostSwitchDefaultPtime","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdatePostSwitchDefaultPtime","data":{"id":4,"position":1,"enabled":true,"codec_name":"paramn2","codec_ptime":"paramv2","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePostSwitchDefaultPtime":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadSwitchDefaultPtime{Id: msg.Param.Id, CodecName: msg.Param.Name, CodecPtime: msg.Param.Value}, []string{"Name", "CodecName", "CodecPtime"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPostSwitchDefaultPtime","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchPostSwitchDefaultPtime","data":{"id":4,"position":1,"enabled":false,"codec_name":"paramn","codec_ptime":"paramv","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchPostSwitchDefaultPtime":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadSwitchDefaultPtime{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddPostSwitchDefaultPtime","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"enabled":true,"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddPostSwitchDefaultPtime","data":{"id":4,"position":1,"enabled":true,"codec_name":"paramn","codec_ptime":"paramv","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddPostSwitchDefaultPtime":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPostLoadSwitchDefaultPtime{CodecName: msg.Param.Name, CodecPtime: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadSwitchDefaultPtime{}))}, onlyAdminGroup())
	//Request:{"event":"DelPostSwitchDefaultPtime","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":9}}}
	//Response:{"MessageType":"DelPostSwitchDefaultPtime","data":{"id":9,"position":2,"enabled":true,"codec_name":"dfsdfs","codec_ptime":"sfsfsd","description":"","parent":{"id":43,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPostSwitchDefaultPtime":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPostLoadSwitchDefaultPtime{Id: msg.Param.Id}, onlyAdminGroup())
	//### Distributor
	//Request:{"event":"GetDistributorConfig","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDistributorConfig","data":{"1":{"id":1,"position":1,"enabled":true,"name":"test","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetDistributorConfig":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDistributorList{}, onlyAdminGroup())
	//Request:{"event":"AddDistributorList","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_list"}}
	//Response:{"MessageType":"AddDistributorList","data":{"id":7,"position":2,"enabled":true,"name":"new_list","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddDistributorList":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDistributorList{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigDistributorList{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateDistributorList","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":7,"name":"new_list2"}}
	//Response:{"MessageType":"UpdateDistributorList","data":{"id":7,"position":2,"enabled":true,"name":"new_list2","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateDistributorList":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDistributorList{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelDistributorList","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":7}}
	//Response:{"MessageType":"DelDistributorList","data":{"id":7,"position":2,"enabled":true,"name":"new_list2","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelDistributorList":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDistributorList{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetDistributorNodes","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetDistributorNodes","data":{"1":{"id":1,"position":1,"enabled":true,"name":"foo1","weight":"1","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"foo2","weight":"9","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetDistributorNodes":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDistributorListNode{}, onlyAdminGroup())
	//Request:{"event":"AddDistributorNode","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"distributor_node":{"name":"paramn","weight":"paramv"}}}
	//Response:{"MessageType":"AddDistributorNode","data":{"id":15,"position":3,"enabled":true,"name":"paramn","weight":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddDistributorNode":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDistributorListNode{Name: msg.DistributorNode.Name, Weight: msg.DistributorNode.Weight, Enabled: true, Parent: &altStruct.ConfigDistributorList{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelDistributorNode","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","distributor_node":{"id":15}}}
	//Response:{"MessageType":"DelDistributorNode","data":{"id":15,"position":3,"enabled":true,"name":"paramn","weight":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelDistributorNode":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDistributorListNode{Id: msg.DistributorNode.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDistributorNode","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","distributor_node":{"id":2,"name":"foo2","weight":"2"}}}
	//Response:{"MessageType":"UpdateDistributorNode","data":{"id":2,"position":2,"enabled":true,"name":"foo2","weight":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateDistributorNode":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDistributorListNode{Id: msg.DistributorNode.Id, Name: msg.DistributorNode.Name, Weight: msg.DistributorNode.Weight}, []string{"Name", "Weight"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDistributorNode","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","distributor_node":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchDistributorNode","data":{"id":2,"position":2,"enabled":false,"name":"foo2","weight":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchDistributorNode":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDistributorListNode{Id: msg.DistributorNode.Id, Enabled: msg.DistributorNode.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//### Directory
	//Request:{"event":"GetDirectory","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetDirectory","data":{"settings":{},"profiles":{"2":{"id":2,"position":1,"enabled":true,"name":"ccvcx","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":2,"enabled":true,"name":"sefsef","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetDirectory":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigDirectorySetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigDirectoryProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":8,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateDirectoryParameter","data":{"id":8,"position":8,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateDirectoryParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDirectoryProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":8,"enabled":false}}}
	//Response:{"MessageType":"SwitchDirectoryParameter","data":{"id":8,"position":1,"enabled":false,"name":"paramn","value":"paramv","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchDirectoryParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDirectorySetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddDirectoryParameter","data":{"id":8,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddDirectoryParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDirectorySetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigDirectorySetting{}))}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "DelDirectoryParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDirectorySetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetDirectoryProfileParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"GetDirectoryProfileParameters","data":{"14":{"id":14,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetDirectoryProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDirectoryProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddDirectoryProfileParameter","data":{"id":14,"position":1,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddDirectoryProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDirectoryProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigDirectoryProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelDirectoryProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14}}}
	//Response:{"MessageType":"DelDirectoryProfileParameter","data":{"id":14,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelDirectoryProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDirectoryProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchDirectoryProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14,"enabled":false}}}
	//Response:{"MessageType":"SwitchDirectoryProfileParameter","data":{"id":14,"position":1,"enabled":false,"name":"paramn2","value":"paramv2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchDirectoryProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDirectoryProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14,"name":"paramn2","value":"paramv2"}}}
	//Response:{"MessageType":"UpdateDirectoryProfileParameter","data":{"id":14,"position":1,"enabled":true,"name":"paramn2","value":"paramv2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateDirectoryProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDirectoryProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddDirectoryProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"AddDirectoryProfile","data":{"id":5,"position":3,"enabled":true,"name":"new_profile","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddDirectoryProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDirectoryProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigDirectoryProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateDirectoryProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateDirectoryProfile","data":{"id":5,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateDirectoryProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDirectoryProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelDirectoryProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5}}
	//Response:{"MessageType":"DelDirectoryProfile","data":{"id":5,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":15,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelDirectoryProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDirectoryProfile{Id: msg.Id}, onlyAdminGroup())
	//### Fifo
	//Request:{"event":"GetFifo","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetFifo","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"delete-all-outbound-member-on-startup","value":"false","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"cool_fifo@45.61.54.76","importance":"444","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetFifo":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigFifoSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigFifoFifo{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateFifoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"delete-all-outbound-member-on-startup","value":"true"}}}
	//Response:{"MessageType":"UpdateFifoParameter","data":{"id":1,"position":1,"enabled":true,"name":"delete-all-outbound-member-on-startup","value":"true","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateFifoParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFifoSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchFifoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchFifoParameter","data":{"id":1,"position":1,"enabled":false,"name":"delete-all-outbound-member-on-startup","value":"true","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchFifoParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFifoSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddFifoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddFifoParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddFifoParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigFifoSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigFifoSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelFifoParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4}}}
	//Response:{"MessageType":"DelFifoParameter","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelFifoParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigFifoSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetFifoFifoMembers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetFifoFifoMembers","data":{"2":{"id":2,"position":1,"enabled":true,"timeout":"3","simo":"3","lag":"user","body":"user","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","importance":"","description":"","parent":null}}}}
	//Errors:
	case "GetFifoFifoMembers":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigFifoFifoMember{}, onlyAdminGroup())
	//Request:{"event":"AddFifoFifoMember","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"fifo_fifo_member":{"timeout":"3","simo":"3","lag":"3","body":"user"}}}
	//Response:{"MessageType":"AddFifoFifoMember","data":{"id":2,"position":1,"enabled":true,"timeout":"3","simo":"3","lag":"user","body":"3","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","importance":"","description":"","parent":null}}}
	//Errors:
	case "AddFifoFifoMember":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigFifoFifoMember{
			Id:      msg.FifoFifoMember.Id,
			Timeout: msg.FifoFifoMember.Timeout,
			Simo:    msg.FifoFifoMember.Simo,
			Lag:     msg.FifoFifoMember.Lag,
			Body:    msg.FifoFifoMember.Body,
			Enabled: true,
			Parent:  &altStruct.ConfigFifoFifo{Id: msg.Id},
		}, onlyAdminGroup())
	//Request:{"event":"DelFifoFifoMember","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","fifo_fifo_member":{"id":5}}}
	//Response:{"MessageType":"DelFifoFifoMember","data":{"id":5,"position":1,"enabled":true,"timeout":"3","simo":"3","lag":"3","body":"user","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","importance":"","description":"","parent":null}}}
	//Errors:
	case "DelFifoFifoMember":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigFifoFifoMember{Id: msg.FifoFifoMember.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchFifoFifoMember","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","fifo_fifo_member":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchFifoFifoMember","data":{"id":2,"position":1,"enabled":false,"timeout":"3","simo":"3","lag":"user","body":"user","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","importance":"","description":"","parent":null}}}
	//Errors:
	case "SwitchFifoFifoMember":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFifoFifoMember{Id: msg.FifoFifoMember.Id, Enabled: msg.FifoFifoMember.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateFifoFifoMember","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","fifo_fifo_member":{"id":2,"timeout":"5","simo":"5","lag":"4","body":"user"}}}
	//Response:{"MessageType":"UpdateFifoFifoMember","data":{"id":2,"position":1,"enabled":true,"timeout":"5","simo":"5","lag":"4","body":"user","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","importance":"","description":"","parent":null}}}
	//Errors:
	case "UpdateFifoFifoMember":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFifoFifoMember{Id: msg.FifoFifoMember.Id, Timeout: msg.FifoFifoMember.Timeout, Simo: msg.FifoFifoMember.Simo, Lag: msg.FifoFifoMember.Lag, Body: msg.FifoFifoMember.Body}, []string{"Timeout", "Simo", "Lag", "Body"}}, onlyAdminGroup())
	//Request:{"event":"AddFifoFifo","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_fifo"}}
	//Response:{"MessageType":"AddFifoFifo","data":{"id":3,"position":2,"enabled":true,"name":"new_fifo","importance":"0","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddFifoFifo":
		if msg.Importance == "" {
			msg.Importance = "0"
		}
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigFifoFifo{Name: msg.Name, Importance: msg.Importance, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigFifoFifo{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateFifoFifo","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":3,"name":"new_fifo2"}}
	//Response:{"MessageType":"UpdateFifoFifo","data":{"id":3,"position":2,"enabled":true,"name":"new_fifo2","importance":"0","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateFifoFifo":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFifoFifo{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelFifoFifo","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":3}}
	//Response:{"MessageType":"DelFifoFifo","data":{"id":3,"position":2,"enabled":true,"name":"new_fifo2","importance":"0","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelFifoFifo":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigFifoFifo{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateFifoFifoImportance","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","value":"5","id":1}}
	//Response:{"MessageType":"UpdateFifoFifoImportance","data":{"id":1,"position":1,"enabled":true,"name":"cool_fifo@45.61.54.76","importance":"5","description":"","parent":{"id":22,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateFifoFifoImportance":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigFifoFifo{Id: msg.Id, Importance: msg.Value}, []string{"Importance"}}, onlyAdminGroup())
	//### Opal
	//Request:{"event":"GetOpal","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetOpal","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"trace-level","value":"3","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"context","value":"default","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"dialplan","value":"XML","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"dtmf-type","value":"signal","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"jitter-size","value":"40,100","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"gk-address","value":"","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"gk-identifer","value":"","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"gk-interface","value":"45.61.54.76","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"listeners":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"fsfsdf","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetOpal":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigOpalSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigOpalListener{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"listeners"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateOpalParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5,"name":"jitter-size","value":"40,101"}}}
	//Response:{"MessageType":"UpdateOpalParameter","error":"can't get updated item"}
	//Errors:
	case "UpdateOpalParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpalListenerParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOpalParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"enabled":false}}}
	//Response:{"MessageType":"SwitchOpalParameter","data":{"id":4,"position":4,"enabled":false,"name":"dtmf-type","value":"signal","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchOpalParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpalSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddOpalParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOpalParameter","data":{"id":11,"position":9,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOpalParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOpalSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOpalSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelOpalParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":11}}}
	//Response:{"MessageType":"DelOpalParameter","data":{"id":11,"position":9,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelOpalParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOpalSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetOpalListenerParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetOpalListenerParameters","data":{"1":{"id":1,"position":1,"enabled":true,"name":"h323-ip","value":"45.61.54.76","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"h323-port","value":"1720","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetOpalListenerParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigOpalListenerParameter{}, onlyAdminGroup())
	//Request:{"event":"AddOpalListenerParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOpalListenerParameter","data":{"id":5,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddOpalListenerParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOpalListenerParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigOpalListener{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelOpalListenerParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelOpalListenerParameter","data":{"id":5,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelOpalListenerParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOpalListenerParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchOpalListenerParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchOpalListenerParameter","data":{"id":2,"position":2,"enabled":false,"name":"h323-port","value":"1720","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchOpalListenerParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpalListenerParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateOpalListenerParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"h323-port","value":"1721"}}}
	//Response:{"MessageType":"UpdateOpalListenerParameter","data":{"id":2,"position":2,"enabled":true,"name":"h323-port","value":"1721","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateOpalListenerParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpalListenerParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddOpalListener","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_listener"}}
	//Response:{"MessageType":"AddOpalListener","data":{"id":6,"position":3,"enabled":true,"name":"new_listener","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOpalListener":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOpalListener{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOpalListener{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateOpalListener","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6,"name":"new_listener2"}}
	//Response:{"MessageType":"UpdateOpalListener","data":{"id":6,"position":3,"enabled":true,"name":"new_listener2","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOpalListener":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOpalListener{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelOpalListener","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6}}
	//Response:{"MessageType":"DelOpalListener","data":{"id":6,"position":3,"enabled":true,"name":"new_listener2","description":"","parent":{"id":30,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelOpalListener":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOpalListener{Id: msg.Id}, onlyAdminGroup())
	//### Osp
	//Request:{"event":"GetOsp","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetOsp","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"debug-info","value":"disabled","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"log-level","value":"info","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"crypto-hardware","value":"disabled","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"sip","value":"","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"default-protocol","value":"sip","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetOsp":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigOspSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigOspProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateOspParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"log-level","value":"debug"}}}
	//Response:{"MessageType":"UpdateOspParameter","data":{"id":2,"position":2,"enabled":true,"name":"log-level","value":"debug","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOspParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOspSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOspParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchOspParameter","data":{"id":3,"position":3,"enabled":false,"name":"crypto-hardware","value":"disabled","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchOspParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOspSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddOspParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOspParameter","data":{"id":10,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOspParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOspSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOspSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelOspParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":10}}}
	//Response:{"MessageType":"DelOspParameter","data":{"id":10,"position":6,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelOspParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOspSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetOspProfileParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetOspProfileParameters","data":{"1":{"id":1,"position":1,"enabled":true,"name":"service-point-url","value":"http://127.0.0.1:5045/osp","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"service-type","value":"voice","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"max-destinations","value":"12","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"device-ip","value":"127.0.0.1:5080","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ssl-lifetime","value":"300","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"http-max-connections","value":"20","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"http-persistence","value":"60","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"http-retry-delay","value":"0","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"http-retry-limit","value":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"http-timeout","value":"10000","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"work-mode","value":"direct","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetOspProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigOspProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"AddOspProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddOspProfileParameter","data":{"id":19,"position":12,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddOspProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOspProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigOspProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelOspProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":19}}}
	//Response:{"MessageType":"DelOspProfileParameter","data":{"id":19,"position":12,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelOspProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOspProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "SwitchOspProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOspProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"SwitchOspProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":11,"enabled":false}}}
	//Response:{"MessageType":"SwitchOspProfileParameter","data":{"id":11,"position":11,"enabled":false,"name":"max-destinations","value":"12","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateOspProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOspProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddOspProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"AddOspProfile","data":{"id":6,"position":2,"enabled":true,"name":"new_profile","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddOspProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigOspProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOspProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateOspProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateOspProfile","data":{"id":6,"position":2,"enabled":true,"name":"new_profile2","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateOspProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigOspProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelOspProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":6}}
	//Response:{"MessageType":"DelOspProfile","data":{"id":6,"position":2,"enabled":true,"name":"new_profile2","description":"","parent":{"id":33,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelOspProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigOspProfile{Id: msg.Id}, onlyAdminGroup())
	//### Unicall
	//Request:{"event":"GetUnicall","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetUnicall","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"context","value":"default","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"dialplan","value":"XML","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"suppress-dtmf-tone","value":"true","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"profiles":{"1":{"id":1,"position":1,"enabled":true,"span_id":"1","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"span_id":"2","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetUnicall":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigUnicallSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigUnicallSpan{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateUnicallParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"name":"suppress-dtmf-tone","value":"truefalse"}}}
	//Response:{"MessageType":"UpdateUnicallParameter","data":{"id":3,"position":3,"enabled":true,"name":"suppress-dtmf-tone","value":"truefalse","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateUnicallParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigUnicallSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchUnicallParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":3,"enabled":false}}}
	//Response:{"MessageType":"SwitchUnicallParameter","data":{"id":3,"position":3,"enabled":false,"name":"suppress-dtmf-tone","value":"truefalse","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchUnicallParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigUnicallSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddUnicallParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddUnicallParameter","data":{"id":5,"position":4,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddUnicallParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigUnicallSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigUnicallSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelUnicallParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":5}}}
	//Response:{"MessageType":"DelUnicallParameter","data":{"id":5,"position":4,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelUnicallParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigUnicallSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetUnicallSpanParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetUnicallSpanParameters","data":{"1":{"id":1,"position":1,"enabled":true,"name":"protocol-class","value":"mfcr2","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"protocol-variant","value":"ar","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"protocol-end","value":"peer","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"outgoing-allowed","value":"true","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"dialplan","value":"XML","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"context","value":"default","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}}}}
	//Errors:
	case "GetUnicallSpanParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigUnicallSpanParameter{}, onlyAdminGroup())
	//Request:{"event":"AddUnicallSpanParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddUnicallSpanParameter","data":{"id":14,"position":7,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}}}
	//Errors:
	case "AddUnicallSpanParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigUnicallSpanParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigUnicallSpan{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelUnicallSpanParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":14}}}
	//Response:{"MessageType":"DelUnicallSpanParameter","data":{"id":14,"position":7,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}}}
	//Errors:
	case "DelUnicallSpanParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigUnicallSpanParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchUnicallSpanParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":6,"enabled":false}}}
	//Response:{"MessageType":"SwitchUnicallSpanParameter","data":{"id":6,"position":6,"enabled":false,"name":"context","value":"default","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}}}
	//Errors:
	case "SwitchUnicallSpanParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigUnicallSpanParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateUnicallSpanParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4,"name":"outgoing-allowed","value":"false"}}}
	//Response:{"MessageType":"UpdateUnicallSpanParameter","data":{"id":4,"position":4,"enabled":true,"name":"outgoing-allowed","value":"false","description":"","parent":{"id":1,"position":0,"enabled":false,"span_id":"","description":"","parent":null}}}
	//Errors:
	case "UpdateUnicallSpanParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigUnicallSpanParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddUnicallSpan","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_span"}}
	//Response:{"MessageType":"AddUnicallSpan","data":{"id":5,"position":3,"enabled":true,"span_id":"new_span","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddUnicallSpan":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigUnicallSpan{SpanId: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigUnicallSpan{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateUnicallSpan","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5,"name":"new_span3"}}
	//Response:{"MessageType":"UpdateUnicallSpan","data":{"id":5,"position":3,"enabled":true,"span_id":"new_span3","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateUnicallSpan":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigUnicallSpan{Id: msg.Id, SpanId: msg.Name}, []string{"SpanId"}}, onlyAdminGroup())
	//Request:{"event":"DelUnicallSpan","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":5}}
	//Response:{"MessageType":"DelUnicallSpan","data":{"id":5,"position":3,"enabled":true,"span_id":"new_span3","description":"","parent":{"id":45,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelUnicallSpan":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigUnicallSpan{Id: msg.Id}, onlyAdminGroup())
	//### Conference
	//Request:{"event":"GetConference","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetConference","data":{"conference_rooms":{"1":{"id":1,"position":1,"enabled":true,"name":"3001@45.61.54.76","status":"FreeSWITCH","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"conference_profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"a1","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"wideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ultrawideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"cdquality","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"video-mcu-stereo","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"video-mcu-stereo-720","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"video-mcu-stereo-480","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"video-mcu-stereo-320","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"sla","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"conference_caller_control_groups":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"s2","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"conference_chat_permissions_profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"a1","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"wideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ultrawideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"cdquality","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"video-mcu-stereo","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"video-mcu-stereo-720","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"video-mcu-stereo-480","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"video-mcu-stereo-320","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"sla","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetConference":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceAdvertiseRoom{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceProfile{}, onlyAdminGroup())
		resp3 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceCallerControlGroup{}, onlyAdminGroup())
		resp4 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceChatPermissionProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			A interface{} `json:"conference_rooms"`
			B interface{} `json:"conference_profiles"`
			C interface{} `json:"conference_caller_control_groups"`
			D interface{} `json:"conference_chat_permissions_profiles"`
		}{A: resp1.Data, B: resp2.Data, C: resp3.Data, D: resp4.Data}}
	//Request:{"event":"UpdateConferenceRoom","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":true,"name":"3001@45.61.54.76","status":"FreeSWITCH2"}}}
	//Response:{"MessageType":"UpdateConferenceRoom","data":{"id":1,"position":1,"enabled":true,"name":"3001@45.61.54.76","status":"","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceRoom":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceAdvertiseRoom{Id: msg.Param.Id, Name: msg.Param.Name, Status: msg.Param.Value}, []string{"Name", "Status"}}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceRoom","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceRoom","data":{"id":1,"position":1,"enabled":false,"name":"3001@45.61.54.76","status":"","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchConferenceRoom":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceAdvertiseRoom{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceRoom","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"room","value":"status"}}}
	//Response:{"MessageType":"AddConferenceRoom","data":{"id":7,"position":2,"enabled":true,"name":"room","status":"status","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceRoom":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceAdvertiseRoom{Name: msg.Param.Name, Status: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceAdvertiseRoom{}))}, onlyAdminGroup())
	//Request:{"event":"DelConferenceRoom","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":7}}}
	//Response:{"MessageType":"DelConferenceRoom","data":{"id":7,"position":2,"enabled":true,"name":"room","status":"status","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceRoom":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceAdvertiseRoom{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetConferenceCallerControls","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"GetConferenceCallerControls","data":{"1":{"id":1,"position":1,"enabled":true,"action":"mute","digits":"0","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":10,"enabled":true,"action":"vol listen zero","digits":"5","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":11,"enabled":true,"action":"vol listen dn","digits":"4","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":12,"enabled":true,"action":"hangup","digits":"#","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"action":"deaf mute","digits":"*","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":3,"enabled":true,"action":"energy up","digits":"9","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":4,"enabled":true,"action":"energy equ","digits":"8","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":5,"enabled":true,"action":"energy dn","digits":"7","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":6,"enabled":true,"action":"vol talk up","digits":"3","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":7,"enabled":true,"action":"vol talk zero","digits":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":8,"enabled":true,"action":"vol talk dn","digits":"1","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":9,"enabled":true,"action":"vol listen up","digits":"6","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetConferenceCallerControls":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceCallerControlGroupControl{}, onlyAdminGroup())
	//Request:{"event":"AddConferenceCallerControl","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1,"param":{"name":"action","value":"2"}}}
	//Response:{"MessageType":"AddConferenceCallerControl","data":{"id":19,"position":13,"enabled":true,"action":"action","digits":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddConferenceCallerControl":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceCallerControlGroupControl{Action: msg.Param.Name, Digits: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceCallerControlGroup{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceCallerControl","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":19}}}
	//Response:{"MessageType":"DelConferenceCallerControl","data":{"id":19,"position":13,"enabled":true,"action":"action","digits":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelConferenceCallerControl":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceCallerControlGroupControl{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceCallerControl","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":12,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceCallerControl","data":{"id":12,"position":12,"enabled":false,"action":"hangup","digits":"#","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchConferenceCallerControl":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceCallerControlGroupControl{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceCallerControl","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":11,"name":"vol listen dn","value":"4"}}}
	//Response:{"MessageType":"UpdateConferenceCallerControl","data":{"id":11,"position":11,"enabled":true,"action":"vol listen dn","digits":"4","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateConferenceCallerControl":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceCallerControlGroupControl{Id: msg.Param.Id, Action: msg.Param.Name, Digits: msg.Param.Value}, []string{"Action", "Digits"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceCallerControlGroup","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_group"}}
	//Response:{"MessageType":"AddConferenceCallerControlGroup","data":{"id":4,"position":3,"enabled":true,"name":"new_group","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceCallerControlGroup":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceCallerControlGroup{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceCallerControlGroup{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceCallerControlGroup","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4,"name":"new_group2"}}
	//Response:{"MessageType":"UpdateConferenceCallerControlGroup","data":{"id":4,"position":3,"enabled":true,"name":"new_group2","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceCallerControlGroup":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceCallerControlGroup{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceCallerControlGroup","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":4}}
	//Response:{"MessageType":"DelConferenceCallerControlGroup","data":{"id":4,"position":3,"enabled":true,"name":"new_group2","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceCallerControlGroup":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceCallerControlGroup{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetConferenceProfileParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":10}}
	//Response:{"MessageType":"GetConferenceProfileParameters","data":{"204":{"id":204,"position":1,"enabled":true,"name":"domain","value":"45.61.54.76","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"205":{"id":205,"position":2,"enabled":true,"name":"rate","value":"8000","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"206":{"id":206,"position":3,"enabled":true,"name":"interval","value":"20","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"207":{"id":207,"position":4,"enabled":true,"name":"energy-level","value":"100","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"208":{"id":208,"position":5,"enabled":true,"name":"muted-sound","value":"conference/conf-muted.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"209":{"id":209,"position":6,"enabled":true,"name":"unmuted-sound","value":"conference/conf-unmuted.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"210":{"id":210,"position":7,"enabled":true,"name":"alone-sound","value":"conference/conf-alone.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"211":{"id":211,"position":8,"enabled":true,"name":"moh-sound","value":"local_stream://moh","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"212":{"id":212,"position":9,"enabled":true,"name":"enter-sound","value":"tone_stream://%(200,0,500,600,700)","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"213":{"id":213,"position":10,"enabled":true,"name":"exit-sound","value":"tone_stream://%(500,0,300,200,100,50,25)","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"214":{"id":214,"position":11,"enabled":true,"name":"kicked-sound","value":"conference/conf-kicked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"215":{"id":215,"position":12,"enabled":true,"name":"locked-sound","value":"conference/conf-locked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"216":{"id":216,"position":13,"enabled":true,"name":"is-locked-sound","value":"conference/conf-is-locked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"217":{"id":217,"position":14,"enabled":true,"name":"is-unlocked-sound","value":"conference/conf-is-unlocked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"218":{"id":218,"position":15,"enabled":true,"name":"pin-sound","value":"conference/conf-pin.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"219":{"id":219,"position":16,"enabled":true,"name":"bad-pin-sound","value":"conference/conf-bad-pin.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"220":{"id":220,"position":17,"enabled":true,"name":"caller-id-name","value":"FreeSWITCH","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"221":{"id":221,"position":18,"enabled":true,"name":"caller-id-number","value":"0000000000","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"222":{"id":222,"position":19,"enabled":true,"name":"comfort-noise","value":"true","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetConferenceProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"AddConferenceProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":10,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddConferenceProfileParameter","data":{"id":407,"position":20,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddConferenceProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":407}}}
	//Response:{"MessageType":"DelConferenceProfileParameter","data":{"id":407,"position":20,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelConferenceProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":222,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceProfileParameter","data":{"id":222,"position":19,"enabled":false,"name":"comfort-noise","value":"true","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchConferenceProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":222,"name":"comfort-noise","value":"false"}}}
	//Response:{"MessageType":"UpdateConferenceProfileParameter","data":{"id":222,"position":19,"enabled":true,"name":"comfort-noise","value":"false","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateConferenceProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"AddConferenceProfile","data":{"id":19,"position":10,"enabled":true,"name":"new_profile","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":19,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateConferenceProfile","data":{"id":19,"position":10,"enabled":true,"name":"new_profile2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":1}}
	//Response:{"MessageType":"DelConferenceProfile","data":{"id":19,"position":10,"enabled":true,"name":"new_profile2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceProfile{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetConferenceChatPermissionUsers","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":10}}
	//Response:{"MessageType":"GetConferenceChatPermissionUsers","data":{"1":{"id":1,"position":1,"enabled":true,"name":"paramn","commands":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetConferenceChatPermissionUsers":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceChatPermissionProfileUser{}, onlyAdminGroup())
	//Request:{"event":"AddConferenceChatPermissionUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":10,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":true,"name":"paramn","commands":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddConferenceChatPermissionUser":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceChatPermissionProfileUser{Name: msg.Param.Name, Commands: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceChatPermissionProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceChatPermissionUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1}}}
	//Response:{"MessageType":"DelConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":true,"name":"paramn2","commands":"","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelConferenceChatPermissionUser":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceChatPermissionProfileUser{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceChatPermissionUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":false,"name":"paramn","commands":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchConferenceChatPermissionUser":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceChatPermissionProfileUser{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceChatPermissionUser","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":1,"name":"paramn2","commands":"paramv2"}}}
	//Response:{"MessageType":"UpdateConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":true,"name":"paramn2","commands":"","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateConferenceChatPermissionUser":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceChatPermissionProfileUser{Id: msg.Param.Id, Name: msg.Param.Name, Commands: msg.Param.Value}, []string{"Name", "Commands"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceChatPermission","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_permission"}}
	//Response:{"MessageType":"AddConferenceChatPermission","data":{"id":20,"position":10,"enabled":true,"name":"new_permission","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceChatPermission":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceChatPermissionProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceChatPermissionProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceChatPermission","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":20,"name":"new_permission2"}}
	//Response:{"MessageType":"UpdateConferenceChatPermission","data":{"id":20,"position":10,"enabled":true,"name":"new_permission2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceChatPermission":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceChatPermissionProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceChatPermission","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"DelConferenceChatPermission","data":{"id":20,"position":10,"enabled":true,"name":"new_permission2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceChatPermission":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceChatPermissionProfile{Id: msg.Id}, onlyAdminGroup())
	//### Post Load Modules
	//Request:{"event":"GetPostLoadModules","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetPostLoadModules","data":{"1":{"id":1,"position":1,"enabled":false,"name":"mod_sofia","description":" ","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":7,"enabled":false,"name":"mod_unicall","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":8,"enabled":false,"name":"mod_xml_cdr","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"12":{"id":12,"position":9,"enabled":false,"name":"mod_xml_rpc","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"13":{"id":13,"position":10,"enabled":true,"name":"mod_shout","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"14":{"id":14,"position":11,"enabled":true,"name":"mod_pocketsphinx","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"15":{"id":15,"position":12,"enabled":true,"name":"mod_alsa","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":false,"name":"mod_amr","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":3,"enabled":false,"name":"mod_db","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":4,"enabled":false,"name":"mod_verto","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":5,"enabled":true,"name":"mod_voicemail","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":6,"enabled":false,"name":"mod_zeroconf","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetPostLoadModules":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigPostLoadModule{}, onlyAdminGroup())
	//Request:{"event":"UpdatePostLoadModule","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":13,"name":"mod_shout"}}}
	//Response:{"MessageType":"UpdatePostLoadModule","data":{"id":13,"position":10,"enabled":true,"name":"mod_shout","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePostLoadModule":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadModule{Id: msg.Param.Id, Name: msg.Param.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPostLoadModule","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":13,"enabled":false}}}
	//Response:{"MessageType":"SwitchPostLoadModule","data":{"id":13,"position":10,"enabled":false,"name":"mod_shout","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case webStruct.SwitchPostLoadModule:
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadModule{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"GetVoicemailSettings","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetVoicemailSettings","data":{"2":{"id":2,"position":1,"enabled":true,"name":"dsfsdf2","value":"sdfsfs2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case webStruct.AddPostLoadModule:
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPostLoadModule{Name: msg.Param.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadModule{}))}, onlyAdminGroup())
	//Request:{"event":"DelPostLoadModule","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":16}}}
	//Response:{"MessageType":"DelPostLoadModule","data":{"id":16,"position":13,"enabled":true,"name":"mod_fifo","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPostLoadModule":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPostLoadModule{Id: msg.Param.Id}, onlyAdminGroup())
	//### Voicemail
	//Request:{"event":"GetVoicemailProfiles","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetVoicemailProfiles","data":{"2":{"id":2,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":2,"enabled":true,"name":"ccc","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetVoicemailSettings":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateVoicemailSetting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"name":"dsfsdf2","value":"sdfsfs2"}}}
	//Response:{"MessageType":"UpdateVoicemailSetting","data":{"id":2,"position":1,"enabled":true,"name":"dsfsdf2","value":"sdfsfs2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateVoicemailSetting":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchVoicemailSetting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchVoicemailSetting","data":{"id":2,"position":1,"enabled":false,"name":"dsfsdf2","value":"sdfsfs2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchVoicemailSetting":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddVoicemailSetting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddVoicemailSetting","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddVoicemailSetting":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVoicemailSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVoicemailSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelVoicemailSetting","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":4}}}
	//Response:{"MessageType":"DelVoicemailSetting","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelVoicemailSetting":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVoicemailSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetVoicemailProfiles","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetVoicemailProfiles","data":{"2":{"id":2,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":2,"enabled":true,"name":"ccc","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetVoicemailProfiles":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailProfile{}, onlyAdminGroup())
	//Request:{"event":"AddVoicemailProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","name":"new_profile"}}
	//Response:{"MessageType":"AddVoicemailProfile","data":{"id":8,"position":3,"enabled":true,"name":"new_profile","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddVoicemailProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVoicemailProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVoicemailProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateVoicemailProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":8,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateVoicemailProfile","data":{"id":8,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateVoicemailProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelVoicemailProfile","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":8}}
	//Response:{"MessageType":"DelVoicemailProfile","data":{"id":8,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelVoicemailProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVoicemailProfile{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetVoicemailProfileParameters","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2}}
	//Response:{"MessageType":"GetVoicemailProfileParameters","data":{"1":{"id":1,"position":1,"enabled":true,"name":"file-extension","value":"wav","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"callback-context","value":"default","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"play-new-messages-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":12,"enabled":true,"name":"play-saved-messages-key","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"13":{"id":13,"position":13,"enabled":true,"name":"login-keys","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"14":{"id":14,"position":14,"enabled":true,"name":"main-menu-key","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"15":{"id":15,"position":15,"enabled":true,"name":"config-menu-key","value":"5","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"16":{"id":16,"position":16,"enabled":true,"name":"record-greeting-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"17":{"id":17,"position":17,"enabled":true,"name":"choose-greeting-key","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"18":{"id":18,"position":18,"enabled":true,"name":"change-pass-key","value":"6","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"19":{"id":19,"position":19,"enabled":true,"name":"record-name-key","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"terminator-key","value":"#","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"20":{"id":20,"position":20,"enabled":true,"name":"record-file-key","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"21":{"id":21,"position":21,"enabled":true,"name":"listen-file-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"22":{"id":22,"position":22,"enabled":true,"name":"save-file-key","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"23":{"id":23,"position":23,"enabled":true,"name":"delete-file-key","value":"7","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"24":{"id":24,"position":24,"enabled":true,"name":"undelete-file-key","value":"8","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"25":{"id":25,"position":25,"enabled":true,"name":"email-key","value":"4","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"26":{"id":26,"position":26,"enabled":true,"name":"pause-key","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"27":{"id":27,"position":27,"enabled":true,"name":"restart-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"28":{"id":28,"position":28,"enabled":true,"name":"ff-key","value":"6","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"29":{"id":29,"position":29,"enabled":true,"name":"rew-key","value":"4","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"max-login-attempts","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"30":{"id":30,"position":30,"enabled":true,"name":"skip-greet-key","value":"#","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"31":{"id":31,"position":31,"enabled":true,"name":"previous-message-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"32":{"id":32,"position":32,"enabled":true,"name":"next-message-key","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"33":{"id":33,"position":33,"enabled":true,"name":"skip-info-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"34":{"id":34,"position":34,"enabled":true,"name":"repeat-message-key","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"35":{"id":35,"position":35,"enabled":true,"name":"record-silence-threshold","value":"200","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"36":{"id":36,"position":36,"enabled":true,"name":"record-silence-hits","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"37":{"id":37,"position":37,"enabled":true,"name":"web-template-file","value":"web-vm.tpl","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"38":{"id":38,"position":38,"enabled":true,"name":"db-password-override","value":"false","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"39":{"id":39,"position":39,"enabled":true,"name":"allow-empty-password-auth","value":"true","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"digit-timeout","value":"10000","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"40":{"id":40,"position":40,"enabled":true,"name":"operator-extension","value":"operator XML default","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"41":{"id":41,"position":41,"enabled":true,"name":"operator-key","value":"9","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"42":{"id":42,"position":42,"enabled":true,"name":"vmain-extension","value":"vmain XML default","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"43":{"id":43,"position":43,"enabled":true,"name":"vmain-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"min-record-len","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"max-record-len","value":"300","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"max-retries","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"tone-spec","value":"%(1000, 0, 640)","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"callback-dialplan","value":"XML","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetVoicemailProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"AddVoicemailProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","id":2,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddVoicemailProfileParameter","data":{"id":48,"position":44,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddVoicemailProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVoicemailProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigVoicemailProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelVoicemailProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":48}}}
	//Response:{"MessageType":"DelVoicemailProfileParameter","data":{"id":48,"position":44,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelVoicemailProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVoicemailProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchVoicemailProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":43,"enabled":false}}}
	//Response:{"MessageType":"SwitchVoicemailProfileParameter","data":{"id":43,"position":43,"enabled":false,"name":"vmain-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchVoicemailProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateVoicemailProfileParameter","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","param":{"id":43,"name":"vmain-key","value":"*"}}}
	//Response:{"MessageType":"UpdateVoicemailProfileParameter","data":{"id":43,"position":43,"enabled":true,"name":"vmain-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateVoicemailProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
		/*
			case "GetVoicemailProfileParameters":
				resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailProfileEmailParameter{}, onlyAdminGroup())
			case "AddVoicemailProfileParameter":
				resp = getUser(msg, AddVoicemailProfileParameter, onlyAdminGroup())
			case "DelVoicemailProfileParameter":
				resp = getUser(msg, DelVoicemailProfileParameter, onlyAdminGroup())
			case "SwitchVoicemailProfileParameter":
				resp = getUser(msg, SwitchVoicemailProfileParameter, onlyAdminGroup())
			case "UpdateVoicemailProfileParameter":
				resp = getUser(msg, UpdateVoicemailProfileParameter, onlyAdminGroup())
		*/
	//## Global Variables
	//Request:{"event":"GetGlobalVariables","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"GetGlobalVariables","global_variables":{"1":{"id":1,"enabled":true,"dynamic":true,"name":"hostname","value":"debian-05","type":"set","position":1},"10":{"id":10,"enabled":true,"dynamic":true,"name":"log_dir","value":"/var/log/freeswitch","type":"set","position":10},"100":{"id":100,"enabled":true,"dynamic":false,"name":"video_mute_png","value":"/var/lib/freeswitch/images/default-mute.png","type":"set","position":100},"101":{"id":101,"enabled":true,"dynamic":false,"name":"video_no_avatar_png","value":"/var/lib/freeswitch/images/default-avatar.png","type":"set","position":101},"102":{"id":102,"enabled":true,"dynamic":false,"name":"rtp_liberal_dtmf","value":"true","type":"set","position":102},"103":{"id":103,"enabled":true,"dynamic":false,"name":"AT_EPENT1","value":"0 0 0 -1 -1 0 -1 0 -1 -1 0 -1","type":"set","position":103},"104":{"id":104,"enabled":true,"dynamic":false,"name":"AT_EPENT2","value":"1 1 1 -1 -1 1 -1 1 -1 -1 1 -1","type":"set","position":104},"105":{"id":105,"enabled":true,"dynamic":false,"name":"AT_CPENT1","value":"0 -1 -1 0 -1 0 0 0 -1 -1 0 -1","type":"set","position":105},"106":{"id":106,"enabled":true,"dynamic":false,"name":"AT_CPENT2","value":"1 -1 -1 1 -1 1 1 1 -1 -1 1 -1","type":"set","position":106},"107":{"id":107,"enabled":true,"dynamic":false,"name":"AT_CMAJ1","value":"0 -1 0 0 -1 0 -1 0 0 -1 0 -1","type":"set","position":107},"108":{"id":108,"enabled":true,"dynamic":false,"name":"AT_CMAJ2","value":"1 -1 1 1 -1 1 -1 1 1 -1 1 -1","type":"set","position":109},"109":{"id":109,"enabled":true,"dynamic":false,"name":"AT_BBLUES","value":"1 -1 1 -1 -1 1 -1 1 1 1 -1 -1","type":"set","position":110},"11":{"id":11,"enabled":true,"dynamic":true,"name":"run_dir","value":"/var/run/freeswitch","type":"set","position":11},"110":{"id":110,"enabled":true,"dynamic":false,"name":"ATGPENT2","value":"-1 1 -1 1 -1 1 -1 -1 1 -1 1 -1","type":"set","position":111},"111":{"id":111,"enabled":true,"dynamic":true,"name":"zrtp_enabled","value":"false","type":"set","position":112},"112":{"id":112,"enabled":true,"dynamic":true,"name":"core_uuid","value":"set","type":"set","position":113},"113":{"id":113,"enabled":true,"dynamic":false,"name":"sfsdfsdf","value":"dsfcsdfsfsdfsd","type":"set","position":108},"12":{"id":12,"enabled":true,"dynamic":true,"name":"db_dir","value":"/var/lib/freeswitch/db","type":"set","position":12},"13":{"id":13,"enabled":true,"dynamic":true,"name":"mod_dir","value":"/usr/lib/freeswitch/mod","type":"set","position":13},"14":{"id":14,"enabled":true,"dynamic":true,"name":"htdocs_dir","value":"/usr/share/freeswitch/htdocs","type":"set","position":14},"15":{"id":15,"enabled":true,"dynamic":true,"name":"script_dir","value":"/usr/share/freeswitch/scripts","type":"set","position":15},"16":{"id":16,"enabled":true,"dynamic":true,"name":"temp_dir","value":"/tmp","type":"set","position":16},"17":{"id":17,"enabled":true,"dynamic":true,"name":"grammar_dir","value":"/usr/share/freeswitch/grammar","type":"set","position":17},"18":{"id":18,"enabled":true,"dynamic":true,"name":"certs_dir","value":"/etc/freeswitch/tls","type":"set","position":18},"19":{"id":19,"enabled":true,"dynamic":true,"name":"storage_dir","value":"/var/lib/freeswitch/storage","type":"set","position":19},"2":{"id":2,"enabled":true,"dynamic":true,"name":"local_ip_v4","value":"45.61.54.76","type":"set","position":2},"20":{"id":20,"enabled":true,"dynamic":true,"name":"cache_dir","value":"/var/cache/freeswitch","type":"set","position":20},"21":{"id":21,"enabled":true,"dynamic":true,"name":"switch_serial","value":"2d3d364cd6cc","type":"set","position":21},"22":{"id":22,"enabled":true,"dynamic":false,"name":"fonts_dir","value":"/usr/share/freeswitch/fonts","type":"set","position":22},"23":{"id":23,"enabled":true,"dynamic":false,"name":"images_dir","value":"/var/lib/freeswitch/images","type":"set","position":23},"24":{"id":24,"enabled":true,"dynamic":false,"name":"data_dir","value":"/usr/share/freeswitch","type":"set","position":24},"25":{"id":25,"enabled":true,"dynamic":false,"name":"localstate_dir","value":"/var/lib/freeswitch","type":"set","position":25},"26":{"id":26,"enabled":true,"dynamic":false,"name":"default_password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","type":"set","position":26},"27":{"id":27,"enabled":true,"dynamic":false,"name":"domain","value":"45.61.54.76","type":"set","position":27},"28":{"id":28,"enabled":true,"dynamic":false,"name":"domain_name","value":"45.61.54.76","type":"set","position":28},"29":{"id":29,"enabled":true,"dynamic":false,"name":"hold_music","value":"local_stream://moh","type":"set","position":29},"3":{"id":3,"enabled":true,"dynamic":true,"name":"local_mask_v4","value":"255.255.255.0","type":"set","position":3},"30":{"id":30,"enabled":true,"dynamic":false,"name":"use_profile","value":"external","type":"set","position":30},"31":{"id":31,"enabled":true,"dynamic":false,"name":"rtp_sdes_suites","value":"AEAD_AES_256_GCM_8|AEAD_AES_128_GCM_8|AES_CM_256_HMAC_SHA1_80|AES_CM_192_HMAC_SHA1_80|AES_CM_128_HMAC_SHA1_80|AES_CM_256_HMAC_SHA1_32|AES_CM_192_HMAC_SHA1_32|AES_CM_128_HMAC_SHA1_32|AES_CM_128_NULL_AUTH","type":"set","position":31},"32":{"id":32,"enabled":true,"dynamic":false,"name":"zrtp_secure_media","value":"true","type":"set","position":32},"33":{"id":33,"enabled":true,"dynamic":false,"name":"global_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":33},"34":{"id":34,"enabled":true,"dynamic":false,"name":"outbound_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":34},"35":{"id":35,"enabled":true,"dynamic":false,"name":"xmpp_client_profile","value":"xmppc","type":"set","position":35},"36":{"id":36,"enabled":true,"dynamic":false,"name":"xmpp_server_profile","value":"xmpps","type":"set","position":36},"37":{"id":37,"enabled":true,"dynamic":false,"name":"bind_server_ip","value":"auto","type":"set","position":37},"38":{"id":38,"enabled":true,"dynamic":false,"name":"external_rtp_ip","value":"45.61.54.76","type":"set","position":38},"39":{"id":39,"enabled":true,"dynamic":false,"name":"external_sip_ip","value":"45.61.54.76","type":"set","position":39},"4":{"id":4,"enabled":true,"dynamic":true,"name":"local_ip_v6","value":"::1","type":"set","position":4},"40":{"id":40,"enabled":true,"dynamic":false,"name":"unroll_loops","value":"true","type":"set","position":40},"41":{"id":41,"enabled":true,"dynamic":false,"name":"outbound_caller_name","value":"FreeSWITCH","type":"set","position":41},"42":{"id":42,"enabled":true,"dynamic":false,"name":"outbound_caller_id","value":"0000000000","type":"set","position":42},"43":{"id":43,"enabled":true,"dynamic":false,"name":"call_debug","value":"false","type":"set","position":43},"44":{"id":44,"enabled":true,"dynamic":false,"name":"console_loglevel","value":"info","type":"set","position":44},"45":{"id":45,"enabled":true,"dynamic":false,"name":"default_areacode","value":"918","type":"set","position":45},"46":{"id":46,"enabled":true,"dynamic":false,"name":"default_country","value":"US","type":"set","position":46},"47":{"id":47,"enabled":true,"dynamic":false,"name":"presence_privacy","value":"false","type":"set","position":47},"48":{"id":48,"enabled":true,"dynamic":false,"name":"au-ring","value":"%(400,200,383,417);%(400,2000,383,417)","type":"set","position":48},"49":{"id":49,"enabled":true,"dynamic":false,"name":"be-ring","value":"%(1000,3000,425)","type":"set","position":49},"5":{"id":5,"enabled":true,"dynamic":true,"name":"base_dir","value":"/usr","type":"set","position":5},"50":{"id":50,"enabled":true,"dynamic":false,"name":"ca-ring","value":"%(2000,4000,440,480)","type":"set","position":50},"51":{"id":51,"enabled":true,"dynamic":false,"name":"cn-ring","value":"%(1000,4000,450)","type":"set","position":51},"52":{"id":52,"enabled":true,"dynamic":false,"name":"cy-ring","value":"%(1500,3000,425)","type":"set","position":52},"53":{"id":53,"enabled":true,"dynamic":false,"name":"cz-ring","value":"%(1000,4000,425)","type":"set","position":53},"54":{"id":54,"enabled":true,"dynamic":false,"name":"de-ring","value":"%(1000,4000,425)","type":"set","position":54},"55":{"id":55,"enabled":true,"dynamic":false,"name":"dk-ring","value":"%(1000,4000,425)","type":"set","position":55},"56":{"id":56,"enabled":true,"dynamic":false,"name":"dz-ring","value":"%(1500,3500,425)","type":"set","position":56},"57":{"id":57,"enabled":true,"dynamic":false,"name":"eg-ring","value":"%(2000,1000,475,375)","type":"set","position":57},"58":{"id":58,"enabled":true,"dynamic":false,"name":"es-ring","value":"%(1500,3000,425)","type":"set","position":58},"59":{"id":59,"enabled":true,"dynamic":false,"name":"fi-ring","value":"%(1000,4000,425)","type":"set","position":59},"6":{"id":6,"enabled":true,"dynamic":true,"name":"recordings_dir","value":"/var/lib/freeswitch/recordings","type":"set","position":6},"60":{"id":60,"enabled":true,"dynamic":false,"name":"fr-ring","value":"%(1500,3500,440)","type":"set","position":60},"61":{"id":61,"enabled":true,"dynamic":false,"name":"hk-ring","value":"%(400,200,440,480);%(400,3000,440,480)","type":"set","position":61},"62":{"id":62,"enabled":true,"dynamic":false,"name":"hu-ring","value":"%(1250,3750,425)","type":"set","position":62},"63":{"id":63,"enabled":true,"dynamic":false,"name":"il-ring","value":"%(1000,3000,400)","type":"set","position":63},"64":{"id":64,"enabled":true,"dynamic":false,"name":"in-ring","value":"%(400,200,425,375);%(400,2000,425,375)","type":"set","position":64},"65":{"id":65,"enabled":true,"dynamic":false,"name":"jp-ring","value":"%(1000,2000,420,380)","type":"set","position":65},"66":{"id":66,"enabled":true,"dynamic":false,"name":"ko-ring","value":"%(1000,2000,440,480)","type":"set","position":66},"67":{"id":67,"enabled":true,"dynamic":false,"name":"pk-ring","value":"%(1000,2000,400)","type":"set","position":67},"68":{"id":68,"enabled":true,"dynamic":false,"name":"pl-ring","value":"%(1000,4000,425)","type":"set","position":68},"69":{"id":69,"enabled":true,"dynamic":false,"name":"ro-ring","value":"%(1850,4150,475,425)","type":"set","position":69},"7":{"id":7,"enabled":true,"dynamic":true,"name":"sound_prefix","value":"/usr/share/freeswitch/sounds","type":"set","position":7},"70":{"id":70,"enabled":true,"dynamic":false,"name":"rs-ring","value":"%(1000,4000,425)","type":"set","position":70},"71":{"id":71,"enabled":true,"dynamic":false,"name":"ru-ring","value":"%(800,3200,425)","type":"set","position":71},"72":{"id":72,"enabled":true,"dynamic":false,"name":"sa-ring","value":"%(1200,4600,425)","type":"set","position":72},"73":{"id":73,"enabled":true,"dynamic":false,"name":"tr-ring","value":"%(2000,4000,450)","type":"set","position":73},"74":{"id":74,"enabled":true,"dynamic":false,"name":"uk-ring","value":"%(400,200,400,450);%(400,2000,400,450)","type":"set","position":74},"75":{"id":75,"enabled":true,"dynamic":false,"name":"us-ring","value":"%(2000,4000,440,480)","type":"set","position":75},"76":{"id":76,"enabled":true,"dynamic":false,"name":"bong-ring","value":"v","type":"set","position":76},"77":{"id":77,"enabled":true,"dynamic":false,"name":"beep","value":"%(1000,0,640)","type":"set","position":77},"78":{"id":78,"enabled":true,"dynamic":false,"name":"sit","value":"%(274,0,913.8);%(274,0,1370.6);%(380,0,1776.7)","type":"set","position":78},"79":{"id":79,"enabled":true,"dynamic":false,"name":"df_us_ssn","value":"(?!219099999|078051120)(?!666|000|9\\d{2})\\d{3}(?!00)\\d{2}(?!0{4})\\d{4}","type":"set","position":79},"8":{"id":8,"enabled":true,"dynamic":true,"name":"sounds_dir","value":"/usr/share/freeswitch/sounds","type":"set","position":8},"80":{"id":80,"enabled":true,"dynamic":false,"name":"df_luhn","value":"?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\\d{3})\\d{11}","type":"set","position":80},"81":{"id":81,"enabled":true,"dynamic":false,"name":"default_provider","value":"example.com","type":"set","position":81},"82":{"id":82,"enabled":true,"dynamic":false,"name":"default_provider_username","value":"joeuser","type":"set","position":82},"83":{"id":83,"enabled":true,"dynamic":false,"name":"default_provider_password","value":"password","type":"set","position":83},"84":{"id":84,"enabled":true,"dynamic":false,"name":"default_provider_from_domain","value":"example.com","type":"set","position":84},"85":{"id":85,"enabled":true,"dynamic":false,"name":"default_provider_register","value":"false","type":"set","position":85},"86":{"id":86,"enabled":true,"dynamic":false,"name":"default_provider_contact","value":"5000","type":"set","position":86},"87":{"id":87,"enabled":true,"dynamic":false,"name":"sip_tls_version","value":"tlsv1,tlsv1.1,tlsv1.2","type":"set","position":87},"88":{"id":88,"enabled":true,"dynamic":false,"name":"sip_tls_ciphers","value":"ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH","type":"set","position":88},"89":{"id":89,"enabled":true,"dynamic":false,"name":"internal_auth_calls","value":"true","type":"set","position":89},"9":{"id":9,"enabled":true,"dynamic":true,"name":"conf_dir","value":"/etc/freeswitch","type":"set","position":9},"90":{"id":90,"enabled":true,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":90},"91":{"id":91,"enabled":true,"dynamic":false,"name":"internal_tls_port","value":"5061","type":"set","position":91},"92":{"id":92,"enabled":true,"dynamic":false,"name":"internal_ssl_enable","value":"false","type":"set","position":92},"93":{"id":93,"enabled":true,"dynamic":false,"name":"external_auth_calls","value":"false","type":"set","position":93},"94":{"id":94,"enabled":true,"dynamic":false,"name":"external_sip_port","value":"5080","type":"set","position":94},"95":{"id":95,"enabled":true,"dynamic":false,"name":"external_tls_port","value":"5081","type":"set","position":95},"96":{"id":96,"enabled":true,"dynamic":false,"name":"external_ssl_enable","value":"false","type":"set","position":96},"97":{"id":97,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_in","value":"3mb","type":"set","position":97},"98":{"id":98,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_out","value":"3mb","type":"set","position":98},"99":{"id":99,"enabled":true,"dynamic":false,"name":"suppress_cng","value":"true","type":"set","position":99}}}
	//Errors:
	case "GetGlobalVariables":
		resp = getUser(msg, GetGlobalVariables, onlyAdminGroup())
	//Request:{"event":"UpdateGlobalVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","variable":{"id":117,"dynamic":false,"name":"new_var2","value":"new_val2","type":"set"}}}
	//Response:{"MessageType":"UpdateGlobalVariable","global_variables":{"117":{"id":117,"enabled":true,"dynamic":false,"name":"new_var2","value":"new_val2","type":"set","position":114}}}
	//Errors:
	case "UpdateGlobalVariable":
		resp = getUser(msg, UpdateGlobalVariable, onlyAdminGroup())
	//Request:{"event":"SwitchGlobalVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","variable":{"id":90,"enabled":false}}}
	//Response:{"MessageType":"SwitchGlobalVariable","global_variables":{"90":{"id":90,"enabled":false,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":90}}}
	//Errors:
	case "SwitchGlobalVariable":
		resp = getUser(msg, SwitchGlobalVariable, onlyAdminGroup())
	//Request:{"event":"AddGlobalVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","variable":{"name":"new_var","value":"new_val","type":"set"}}}
	//Response:{"MessageType":"AddGlobalVariable","global_variables":{"117":{"id":117,"enabled":true,"dynamic":false,"name":"new_var","value":"new_val","type":"set","position":114}}}
	//Errors:
	case "AddGlobalVariable":
		resp = getUser(msg, AddGlobalVariable, onlyAdminGroup())
	//Request:{"event":"DelGlobalVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","variable":{"id":117}}}
	//Response:{"MessageType":"DelGlobalVariable","id":117}
	//Errors:
	case "DelGlobalVariable":
		resp = getUser(msg, DelGlobalVariable, onlyAdminGroup())
	//Request:{"event":"MoveGlobalVariable","data":{"token":"3c2f3200f73699a28c96783a15dff1d7","previous_index":111,"current_index":108,"id":110}}
	//Response:{"MessageType":"MoveGlobalVariable","global_variables":{"1":{"id":1,"enabled":true,"dynamic":true,"name":"hostname","value":"debian-05","type":"set","position":1},"10":{"id":10,"enabled":true,"dynamic":true,"name":"log_dir","value":"/var/log/freeswitch","type":"set","position":10},"100":{"id":100,"enabled":true,"dynamic":false,"name":"video_mute_png","value":"/var/lib/freeswitch/images/default-mute.png","type":"set","position":100},"101":{"id":101,"enabled":true,"dynamic":false,"name":"video_no_avatar_png","value":"/var/lib/freeswitch/images/default-avatar.png","type":"set","position":101},"102":{"id":102,"enabled":true,"dynamic":false,"name":"rtp_liberal_dtmf","value":"true","type":"set","position":102},"103":{"id":103,"enabled":true,"dynamic":false,"name":"AT_EPENT1","value":"0 0 0 -1 -1 0 -1 0 -1 -1 0 -1","type":"set","position":103},"104":{"id":104,"enabled":true,"dynamic":false,"name":"AT_EPENT2","value":"1 1 1 -1 -1 1 -1 1 -1 -1 1 -1","type":"set","position":104},"105":{"id":105,"enabled":true,"dynamic":false,"name":"AT_CPENT1","value":"0 -1 -1 0 -1 0 0 0 -1 -1 0 -1","type":"set","position":105},"106":{"id":106,"enabled":true,"dynamic":false,"name":"AT_CPENT2","value":"1 -1 -1 1 -1 1 1 1 -1 -1 1 -1","type":"set","position":106},"107":{"id":107,"enabled":true,"dynamic":false,"name":"AT_CMAJ1","value":"0 -1 0 0 -1 0 -1 0 0 -1 0 -1","type":"set","position":107},"108":{"id":108,"enabled":true,"dynamic":false,"name":"AT_CMAJ2","value":"1 -1 1 1 -1 1 -1 1 1 -1 1 -1","type":"set","position":110},"109":{"id":109,"enabled":true,"dynamic":false,"name":"AT_BBLUES","value":"1 -1 1 -1 -1 1 -1 1 1 1 -1 -1","type":"set","position":111},"11":{"id":11,"enabled":true,"dynamic":true,"name":"run_dir","value":"/var/run/freeswitch","type":"set","position":11},"110":{"id":110,"enabled":true,"dynamic":false,"name":"ATGPENT2","value":"-1 1 -1 1 -1 1 -1 -1 1 -1 1 -1","type":"set","position":108},"111":{"id":111,"enabled":true,"dynamic":true,"name":"zrtp_enabled","value":"false","type":"set","position":112},"112":{"id":112,"enabled":true,"dynamic":true,"name":"core_uuid","value":"set","type":"set","position":113},"113":{"id":113,"enabled":true,"dynamic":false,"name":"sfsdfsdf","value":"dsfcsdfsfsdfsd","type":"set","position":109},"12":{"id":12,"enabled":true,"dynamic":true,"name":"db_dir","value":"/var/lib/freeswitch/db","type":"set","position":12},"13":{"id":13,"enabled":true,"dynamic":true,"name":"mod_dir","value":"/usr/lib/freeswitch/mod","type":"set","position":13},"14":{"id":14,"enabled":true,"dynamic":true,"name":"htdocs_dir","value":"/usr/share/freeswitch/htdocs","type":"set","position":14},"15":{"id":15,"enabled":true,"dynamic":true,"name":"script_dir","value":"/usr/share/freeswitch/scripts","type":"set","position":15},"16":{"id":16,"enabled":true,"dynamic":true,"name":"temp_dir","value":"/tmp","type":"set","position":16},"17":{"id":17,"enabled":true,"dynamic":true,"name":"grammar_dir","value":"/usr/share/freeswitch/grammar","type":"set","position":17},"18":{"id":18,"enabled":true,"dynamic":true,"name":"certs_dir","value":"/etc/freeswitch/tls","type":"set","position":18},"19":{"id":19,"enabled":true,"dynamic":true,"name":"storage_dir","value":"/var/lib/freeswitch/storage","type":"set","position":19},"2":{"id":2,"enabled":true,"dynamic":true,"name":"local_ip_v4","value":"45.61.54.76","type":"set","position":2},"20":{"id":20,"enabled":true,"dynamic":true,"name":"cache_dir","value":"/var/cache/freeswitch","type":"set","position":20},"21":{"id":21,"enabled":true,"dynamic":true,"name":"switch_serial","value":"2d3d364cd6cc","type":"set","position":21},"22":{"id":22,"enabled":true,"dynamic":false,"name":"fonts_dir","value":"/usr/share/freeswitch/fonts","type":"set","position":22},"23":{"id":23,"enabled":true,"dynamic":false,"name":"images_dir","value":"/var/lib/freeswitch/images","type":"set","position":23},"24":{"id":24,"enabled":true,"dynamic":false,"name":"data_dir","value":"/usr/share/freeswitch","type":"set","position":24},"25":{"id":25,"enabled":true,"dynamic":false,"name":"localstate_dir","value":"/var/lib/freeswitch","type":"set","position":25},"26":{"id":26,"enabled":true,"dynamic":false,"name":"default_password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","type":"set","position":26},"27":{"id":27,"enabled":true,"dynamic":false,"name":"domain","value":"45.61.54.76","type":"set","position":27},"28":{"id":28,"enabled":true,"dynamic":false,"name":"domain_name","value":"45.61.54.76","type":"set","position":28},"29":{"id":29,"enabled":true,"dynamic":false,"name":"hold_music","value":"local_stream://moh","type":"set","position":29},"3":{"id":3,"enabled":true,"dynamic":true,"name":"local_mask_v4","value":"255.255.255.0","type":"set","position":3},"30":{"id":30,"enabled":true,"dynamic":false,"name":"use_profile","value":"external","type":"set","position":30},"31":{"id":31,"enabled":true,"dynamic":false,"name":"rtp_sdes_suites","value":"AEAD_AES_256_GCM_8|AEAD_AES_128_GCM_8|AES_CM_256_HMAC_SHA1_80|AES_CM_192_HMAC_SHA1_80|AES_CM_128_HMAC_SHA1_80|AES_CM_256_HMAC_SHA1_32|AES_CM_192_HMAC_SHA1_32|AES_CM_128_HMAC_SHA1_32|AES_CM_128_NULL_AUTH","type":"set","position":31},"32":{"id":32,"enabled":true,"dynamic":false,"name":"zrtp_secure_media","value":"true","type":"set","position":32},"33":{"id":33,"enabled":true,"dynamic":false,"name":"global_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":33},"34":{"id":34,"enabled":true,"dynamic":false,"name":"outbound_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":34},"35":{"id":35,"enabled":true,"dynamic":false,"name":"xmpp_client_profile","value":"xmppc","type":"set","position":35},"36":{"id":36,"enabled":true,"dynamic":false,"name":"xmpp_server_profile","value":"xmpps","type":"set","position":36},"37":{"id":37,"enabled":true,"dynamic":false,"name":"bind_server_ip","value":"auto","type":"set","position":37},"38":{"id":38,"enabled":true,"dynamic":false,"name":"external_rtp_ip","value":"45.61.54.76","type":"set","position":38},"39":{"id":39,"enabled":true,"dynamic":false,"name":"external_sip_ip","value":"45.61.54.76","type":"set","position":39},"4":{"id":4,"enabled":true,"dynamic":true,"name":"local_ip_v6","value":"::1","type":"set","position":4},"40":{"id":40,"enabled":true,"dynamic":false,"name":"unroll_loops","value":"true","type":"set","position":40},"41":{"id":41,"enabled":true,"dynamic":false,"name":"outbound_caller_name","value":"FreeSWITCH","type":"set","position":41},"42":{"id":42,"enabled":true,"dynamic":false,"name":"outbound_caller_id","value":"0000000000","type":"set","position":42},"43":{"id":43,"enabled":true,"dynamic":false,"name":"call_debug","value":"false","type":"set","position":43},"44":{"id":44,"enabled":true,"dynamic":false,"name":"console_loglevel","value":"info","type":"set","position":44},"45":{"id":45,"enabled":true,"dynamic":false,"name":"default_areacode","value":"918","type":"set","position":45},"46":{"id":46,"enabled":true,"dynamic":false,"name":"default_country","value":"US","type":"set","position":46},"47":{"id":47,"enabled":true,"dynamic":false,"name":"presence_privacy","value":"false","type":"set","position":47},"48":{"id":48,"enabled":true,"dynamic":false,"name":"au-ring","value":"%(400,200,383,417);%(400,2000,383,417)","type":"set","position":48},"49":{"id":49,"enabled":true,"dynamic":false,"name":"be-ring","value":"%(1000,3000,425)","type":"set","position":49},"5":{"id":5,"enabled":true,"dynamic":true,"name":"base_dir","value":"/usr","type":"set","position":5},"50":{"id":50,"enabled":true,"dynamic":false,"name":"ca-ring","value":"%(2000,4000,440,480)","type":"set","position":50},"51":{"id":51,"enabled":true,"dynamic":false,"name":"cn-ring","value":"%(1000,4000,450)","type":"set","position":51},"52":{"id":52,"enabled":true,"dynamic":false,"name":"cy-ring","value":"%(1500,3000,425)","type":"set","position":52},"53":{"id":53,"enabled":true,"dynamic":false,"name":"cz-ring","value":"%(1000,4000,425)","type":"set","position":53},"54":{"id":54,"enabled":true,"dynamic":false,"name":"de-ring","value":"%(1000,4000,425)","type":"set","position":54},"55":{"id":55,"enabled":true,"dynamic":false,"name":"dk-ring","value":"%(1000,4000,425)","type":"set","position":55},"56":{"id":56,"enabled":true,"dynamic":false,"name":"dz-ring","value":"%(1500,3500,425)","type":"set","position":56},"57":{"id":57,"enabled":true,"dynamic":false,"name":"eg-ring","value":"%(2000,1000,475,375)","type":"set","position":57},"58":{"id":58,"enabled":true,"dynamic":false,"name":"es-ring","value":"%(1500,3000,425)","type":"set","position":58},"59":{"id":59,"enabled":true,"dynamic":false,"name":"fi-ring","value":"%(1000,4000,425)","type":"set","position":59},"6":{"id":6,"enabled":true,"dynamic":true,"name":"recordings_dir","value":"/var/lib/freeswitch/recordings","type":"set","position":6},"60":{"id":60,"enabled":true,"dynamic":false,"name":"fr-ring","value":"%(1500,3500,440)","type":"set","position":60},"61":{"id":61,"enabled":true,"dynamic":false,"name":"hk-ring","value":"%(400,200,440,480);%(400,3000,440,480)","type":"set","position":61},"62":{"id":62,"enabled":true,"dynamic":false,"name":"hu-ring","value":"%(1250,3750,425)","type":"set","position":62},"63":{"id":63,"enabled":true,"dynamic":false,"name":"il-ring","value":"%(1000,3000,400)","type":"set","position":63},"64":{"id":64,"enabled":true,"dynamic":false,"name":"in-ring","value":"%(400,200,425,375);%(400,2000,425,375)","type":"set","position":64},"65":{"id":65,"enabled":true,"dynamic":false,"name":"jp-ring","value":"%(1000,2000,420,380)","type":"set","position":65},"66":{"id":66,"enabled":true,"dynamic":false,"name":"ko-ring","value":"%(1000,2000,440,480)","type":"set","position":66},"67":{"id":67,"enabled":true,"dynamic":false,"name":"pk-ring","value":"%(1000,2000,400)","type":"set","position":67},"68":{"id":68,"enabled":true,"dynamic":false,"name":"pl-ring","value":"%(1000,4000,425)","type":"set","position":68},"69":{"id":69,"enabled":true,"dynamic":false,"name":"ro-ring","value":"%(1850,4150,475,425)","type":"set","position":69},"7":{"id":7,"enabled":true,"dynamic":true,"name":"sound_prefix","value":"/usr/share/freeswitch/sounds","type":"set","position":7},"70":{"id":70,"enabled":true,"dynamic":false,"name":"rs-ring","value":"%(1000,4000,425)","type":"set","position":70},"71":{"id":71,"enabled":true,"dynamic":false,"name":"ru-ring","value":"%(800,3200,425)","type":"set","position":71},"72":{"id":72,"enabled":true,"dynamic":false,"name":"sa-ring","value":"%(1200,4600,425)","type":"set","position":72},"73":{"id":73,"enabled":true,"dynamic":false,"name":"tr-ring","value":"%(2000,4000,450)","type":"set","position":73},"74":{"id":74,"enabled":true,"dynamic":false,"name":"uk-ring","value":"%(400,200,400,450);%(400,2000,400,450)","type":"set","position":74},"75":{"id":75,"enabled":true,"dynamic":false,"name":"us-ring","value":"%(2000,4000,440,480)","type":"set","position":75},"76":{"id":76,"enabled":true,"dynamic":false,"name":"bong-ring","value":"v","type":"set","position":76},"77":{"id":77,"enabled":true,"dynamic":false,"name":"beep","value":"%(1000,0,640)","type":"set","position":77},"78":{"id":78,"enabled":true,"dynamic":false,"name":"sit","value":"%(274,0,913.8);%(274,0,1370.6);%(380,0,1776.7)","type":"set","position":78},"79":{"id":79,"enabled":true,"dynamic":false,"name":"df_us_ssn","value":"(?!219099999|078051120)(?!666|000|9\\d{2})\\d{3}(?!00)\\d{2}(?!0{4})\\d{4}","type":"set","position":79},"8":{"id":8,"enabled":true,"dynamic":true,"name":"sounds_dir","value":"/usr/share/freeswitch/sounds","type":"set","position":8},"80":{"id":80,"enabled":true,"dynamic":false,"name":"df_luhn","value":"?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\\d{3})\\d{11}","type":"set","position":80},"81":{"id":81,"enabled":true,"dynamic":false,"name":"default_provider","value":"example.com","type":"set","position":81},"82":{"id":82,"enabled":true,"dynamic":false,"name":"default_provider_username","value":"joeuser","type":"set","position":82},"83":{"id":83,"enabled":true,"dynamic":false,"name":"default_provider_password","value":"password","type":"set","position":83},"84":{"id":84,"enabled":true,"dynamic":false,"name":"default_provider_from_domain","value":"example.com","type":"set","position":84},"85":{"id":85,"enabled":true,"dynamic":false,"name":"default_provider_register","value":"false","type":"set","position":85},"86":{"id":86,"enabled":true,"dynamic":false,"name":"default_provider_contact","value":"5000","type":"set","position":86},"87":{"id":87,"enabled":true,"dynamic":false,"name":"sip_tls_version","value":"tlsv1,tlsv1.1,tlsv1.2","type":"set","position":87},"88":{"id":88,"enabled":true,"dynamic":false,"name":"sip_tls_ciphers","value":"ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH","type":"set","position":88},"89":{"id":89,"enabled":true,"dynamic":false,"name":"internal_auth_calls","value":"true","type":"set","position":89},"9":{"id":9,"enabled":true,"dynamic":true,"name":"conf_dir","value":"/etc/freeswitch","type":"set","position":9},"90":{"id":90,"enabled":true,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":90},"91":{"id":91,"enabled":true,"dynamic":false,"name":"internal_tls_port","value":"5061","type":"set","position":91},"92":{"id":92,"enabled":true,"dynamic":false,"name":"internal_ssl_enable","value":"false","type":"set","position":92},"93":{"id":93,"enabled":true,"dynamic":false,"name":"external_auth_calls","value":"false","type":"set","position":93},"94":{"id":94,"enabled":true,"dynamic":false,"name":"external_sip_port","value":"5080","type":"set","position":94},"95":{"id":95,"enabled":true,"dynamic":false,"name":"external_tls_port","value":"5081","type":"set","position":95},"96":{"id":96,"enabled":true,"dynamic":false,"name":"external_ssl_enable","value":"false","type":"set","position":96},"97":{"id":97,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_in","value":"3mb","type":"set","position":97},"98":{"id":98,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_out","value":"3mb","type":"set","position":98},"99":{"id":99,"enabled":true,"dynamic":false,"name":"suppress_cng","value":"true","type":"set","position":99}}}
	//Errors:
	case "MoveGlobalVariable":
		resp = getUser(msg, MoveGlobalVariable, onlyAdminGroup())
	//Request:{"event":"ImportGlobalVariables","data":{"token":"3c2f3200f73699a28c96783a15dff1d7"}}
	//Response:{"MessageType":"ImportGlobalVariables","global_variables":{"1":{"id":1,"enabled":true,"dynamic":true,"name":"hostname","value":"debian-05","type":"set","position":1},"10":{"id":10,"enabled":true,"dynamic":true,"name":"log_dir","value":"/var/log/freeswitch","type":"set","position":10},"100":{"id":100,"enabled":true,"dynamic":false,"name":"video_no_avatar_png","value":"/var/lib/freeswitch/images/default-avatar.png","type":"set","position":100},"101":{"id":101,"enabled":true,"dynamic":false,"name":"rtp_liberal_dtmf","value":"true","type":"set","position":101},"102":{"id":102,"enabled":true,"dynamic":false,"name":"sfsdfsdf","value":"dsfcsdfsfsdfsd","type":"set","position":102},"103":{"id":103,"enabled":true,"dynamic":false,"name":"AT_EPENT1","value":"0 0 0 -1 -1 0 -1 0 -1 -1 0 -1","type":"set","position":103},"104":{"id":104,"enabled":true,"dynamic":false,"name":"AT_EPENT2","value":"1 1 1 -1 -1 1 -1 1 -1 -1 1 -1","type":"set","position":104},"105":{"id":105,"enabled":true,"dynamic":false,"name":"AT_CPENT1","value":"0 -1 -1 0 -1 0 0 0 -1 -1 0 -1","type":"set","position":105},"106":{"id":106,"enabled":true,"dynamic":false,"name":"AT_CPENT2","value":"1 -1 -1 1 -1 1 1 1 -1 -1 1 -1","type":"set","position":106},"107":{"id":107,"enabled":true,"dynamic":false,"name":"AT_CMAJ1","value":"0 -1 0 0 -1 0 -1 0 0 -1 0 -1","type":"set","position":107},"108":{"id":108,"enabled":true,"dynamic":false,"name":"AT_CMAJ2","value":"1 -1 1 1 -1 1 -1 1 1 -1 1 -1","type":"set","position":108},"109":{"id":109,"enabled":true,"dynamic":false,"name":"AT_BBLUES","value":"1 -1 1 -1 -1 1 -1 1 1 1 -1 -1","type":"set","position":109},"11":{"id":11,"enabled":true,"dynamic":true,"name":"run_dir","value":"/var/run/freeswitch","type":"set","position":11},"110":{"id":110,"enabled":true,"dynamic":false,"name":"ATGPENT2","value":"-1 1 -1 1 -1 1 -1 -1 1 -1 1 -1","type":"set","position":110},"111":{"id":111,"enabled":true,"dynamic":true,"name":"core_uuid","value":"4ee847e9-b9fb-49a8-99be-11e42a8cfdd4","type":"set","position":111},"112":{"id":112,"enabled":true,"dynamic":true,"name":"zrtp_enabled","value":"false","type":"set","position":112},"113":{"id":113,"enabled":true,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":113},"12":{"id":12,"enabled":true,"dynamic":true,"name":"db_dir","value":"/var/lib/freeswitch/db","type":"set","position":12},"13":{"id":13,"enabled":true,"dynamic":true,"name":"mod_dir","value":"/usr/lib/freeswitch/mod","type":"set","position":13},"14":{"id":14,"enabled":true,"dynamic":true,"name":"htdocs_dir","value":"/usr/share/freeswitch/htdocs","type":"set","position":14},"15":{"id":15,"enabled":true,"dynamic":true,"name":"script_dir","value":"/usr/share/freeswitch/scripts","type":"set","position":15},"16":{"id":16,"enabled":true,"dynamic":true,"name":"temp_dir","value":"/tmp","type":"set","position":16},"17":{"id":17,"enabled":true,"dynamic":true,"name":"grammar_dir","value":"/usr/share/freeswitch/grammar","type":"set","position":17},"18":{"id":18,"enabled":true,"dynamic":true,"name":"certs_dir","value":"/etc/freeswitch/tls","type":"set","position":18},"19":{"id":19,"enabled":true,"dynamic":true,"name":"storage_dir","value":"/var/lib/freeswitch/storage","type":"set","position":19},"2":{"id":2,"enabled":true,"dynamic":true,"name":"local_ip_v4","value":"45.61.54.76","type":"set","position":2},"20":{"id":20,"enabled":true,"dynamic":true,"name":"cache_dir","value":"/var/cache/freeswitch","type":"set","position":20},"21":{"id":21,"enabled":true,"dynamic":true,"name":"switch_serial","value":"2d3d364cd6cc","type":"set","position":21},"22":{"id":22,"enabled":true,"dynamic":false,"name":"fonts_dir","value":"/usr/share/freeswitch/fonts","type":"set","position":22},"23":{"id":23,"enabled":true,"dynamic":false,"name":"images_dir","value":"/var/lib/freeswitch/images","type":"set","position":23},"24":{"id":24,"enabled":true,"dynamic":false,"name":"data_dir","value":"/usr/share/freeswitch","type":"set","position":24},"25":{"id":25,"enabled":true,"dynamic":false,"name":"localstate_dir","value":"/var/lib/freeswitch","type":"set","position":25},"26":{"id":26,"enabled":true,"dynamic":false,"name":"default_password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","type":"set","position":26},"27":{"id":27,"enabled":true,"dynamic":false,"name":"domain","value":"45.61.54.76","type":"set","position":27},"28":{"id":28,"enabled":true,"dynamic":false,"name":"domain_name","value":"45.61.54.76","type":"set","position":28},"29":{"id":29,"enabled":true,"dynamic":false,"name":"hold_music","value":"local_stream://moh","type":"set","position":29},"3":{"id":3,"enabled":true,"dynamic":true,"name":"local_mask_v4","value":"255.255.255.0","type":"set","position":3},"30":{"id":30,"enabled":true,"dynamic":false,"name":"use_profile","value":"external","type":"set","position":30},"31":{"id":31,"enabled":true,"dynamic":false,"name":"rtp_sdes_suites","value":"AEAD_AES_256_GCM_8|AEAD_AES_128_GCM_8|AES_CM_256_HMAC_SHA1_80|AES_CM_192_HMAC_SHA1_80|AES_CM_128_HMAC_SHA1_80|AES_CM_256_HMAC_SHA1_32|AES_CM_192_HMAC_SHA1_32|AES_CM_128_HMAC_SHA1_32|AES_CM_128_NULL_AUTH","type":"set","position":31},"32":{"id":32,"enabled":true,"dynamic":false,"name":"zrtp_secure_media","value":"true","type":"set","position":32},"33":{"id":33,"enabled":true,"dynamic":false,"name":"global_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":33},"34":{"id":34,"enabled":true,"dynamic":false,"name":"outbound_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":34},"35":{"id":35,"enabled":true,"dynamic":false,"name":"xmpp_client_profile","value":"xmppc","type":"set","position":35},"36":{"id":36,"enabled":true,"dynamic":false,"name":"xmpp_server_profile","value":"xmpps","type":"set","position":36},"37":{"id":37,"enabled":true,"dynamic":false,"name":"bind_server_ip","value":"auto","type":"set","position":37},"38":{"id":38,"enabled":true,"dynamic":false,"name":"external_rtp_ip","value":"45.61.54.76","type":"set","position":38},"39":{"id":39,"enabled":true,"dynamic":false,"name":"external_sip_ip","value":"45.61.54.76","type":"set","position":39},"4":{"id":4,"enabled":true,"dynamic":true,"name":"local_ip_v6","value":"::1","type":"set","position":4},"40":{"id":40,"enabled":true,"dynamic":false,"name":"unroll_loops","value":"true","type":"set","position":40},"41":{"id":41,"enabled":true,"dynamic":false,"name":"outbound_caller_name","value":"FreeSWITCH","type":"set","position":41},"42":{"id":42,"enabled":true,"dynamic":false,"name":"outbound_caller_id","value":"0000000000","type":"set","position":42},"43":{"id":43,"enabled":true,"dynamic":false,"name":"call_debug","value":"false","type":"set","position":43},"44":{"id":44,"enabled":true,"dynamic":false,"name":"console_loglevel","value":"info","type":"set","position":44},"45":{"id":45,"enabled":true,"dynamic":false,"name":"default_areacode","value":"918","type":"set","position":45},"46":{"id":46,"enabled":true,"dynamic":false,"name":"default_country","value":"US","type":"set","position":46},"47":{"id":47,"enabled":true,"dynamic":false,"name":"presence_privacy","value":"false","type":"set","position":47},"48":{"id":48,"enabled":true,"dynamic":false,"name":"au-ring","value":"%(400,200,383,417);%(400,2000,383,417)","type":"set","position":48},"49":{"id":49,"enabled":true,"dynamic":false,"name":"be-ring","value":"%(1000,3000,425)","type":"set","position":49},"5":{"id":5,"enabled":true,"dynamic":true,"name":"base_dir","value":"/usr","type":"set","position":5},"50":{"id":50,"enabled":true,"dynamic":false,"name":"ca-ring","value":"%(2000,4000,440,480)","type":"set","position":50},"51":{"id":51,"enabled":true,"dynamic":false,"name":"cn-ring","value":"%(1000,4000,450)","type":"set","position":51},"52":{"id":52,"enabled":true,"dynamic":false,"name":"cy-ring","value":"%(1500,3000,425)","type":"set","position":52},"53":{"id":53,"enabled":true,"dynamic":false,"name":"cz-ring","value":"%(1000,4000,425)","type":"set","position":53},"54":{"id":54,"enabled":true,"dynamic":false,"name":"de-ring","value":"%(1000,4000,425)","type":"set","position":54},"55":{"id":55,"enabled":true,"dynamic":false,"name":"dk-ring","value":"%(1000,4000,425)","type":"set","position":55},"56":{"id":56,"enabled":true,"dynamic":false,"name":"dz-ring","value":"%(1500,3500,425)","type":"set","position":56},"57":{"id":57,"enabled":true,"dynamic":false,"name":"eg-ring","value":"%(2000,1000,475,375)","type":"set","position":57},"58":{"id":58,"enabled":true,"dynamic":false,"name":"es-ring","value":"%(1500,3000,425)","type":"set","position":58},"59":{"id":59,"enabled":true,"dynamic":false,"name":"fi-ring","value":"%(1000,4000,425)","type":"set","position":59},"6":{"id":6,"enabled":true,"dynamic":true,"name":"recordings_dir","value":"/var/lib/freeswitch/recordings","type":"set","position":6},"60":{"id":60,"enabled":true,"dynamic":false,"name":"fr-ring","value":"%(1500,3500,440)","type":"set","position":60},"61":{"id":61,"enabled":true,"dynamic":false,"name":"hk-ring","value":"%(400,200,440,480);%(400,3000,440,480)","type":"set","position":61},"62":{"id":62,"enabled":true,"dynamic":false,"name":"hu-ring","value":"%(1250,3750,425)","type":"set","position":62},"63":{"id":63,"enabled":true,"dynamic":false,"name":"il-ring","value":"%(1000,3000,400)","type":"set","position":63},"64":{"id":64,"enabled":true,"dynamic":false,"name":"in-ring","value":"%(400,200,425,375);%(400,2000,425,375)","type":"set","position":64},"65":{"id":65,"enabled":true,"dynamic":false,"name":"jp-ring","value":"%(1000,2000,420,380)","type":"set","position":65},"66":{"id":66,"enabled":true,"dynamic":false,"name":"ko-ring","value":"%(1000,2000,440,480)","type":"set","position":66},"67":{"id":67,"enabled":true,"dynamic":false,"name":"pk-ring","value":"%(1000,2000,400)","type":"set","position":67},"68":{"id":68,"enabled":true,"dynamic":false,"name":"pl-ring","value":"%(1000,4000,425)","type":"set","position":68},"69":{"id":69,"enabled":true,"dynamic":false,"name":"ro-ring","value":"%(1850,4150,475,425)","type":"set","position":69},"7":{"id":7,"enabled":true,"dynamic":true,"name":"sound_prefix","value":"/usr/share/freeswitch/sounds","type":"set","position":7},"70":{"id":70,"enabled":true,"dynamic":false,"name":"rs-ring","value":"%(1000,4000,425)","type":"set","position":70},"71":{"id":71,"enabled":true,"dynamic":false,"name":"ru-ring","value":"%(800,3200,425)","type":"set","position":71},"72":{"id":72,"enabled":true,"dynamic":false,"name":"sa-ring","value":"%(1200,4600,425)","type":"set","position":72},"73":{"id":73,"enabled":true,"dynamic":false,"name":"tr-ring","value":"%(2000,4000,450)","type":"set","position":73},"74":{"id":74,"enabled":true,"dynamic":false,"name":"uk-ring","value":"%(400,200,400,450);%(400,2000,400,450)","type":"set","position":74},"75":{"id":75,"enabled":true,"dynamic":false,"name":"us-ring","value":"%(2000,4000,440,480)","type":"set","position":75},"76":{"id":76,"enabled":true,"dynamic":false,"name":"bong-ring","value":"v","type":"set","position":76},"77":{"id":77,"enabled":true,"dynamic":false,"name":"beep","value":"%(1000,0,640)","type":"set","position":77},"78":{"id":78,"enabled":true,"dynamic":false,"name":"sit","value":"%(274,0,913.8);%(274,0,1370.6);%(380,0,1776.7)","type":"set","position":78},"79":{"id":79,"enabled":true,"dynamic":false,"name":"df_us_ssn","value":"(?!219099999|078051120)(?!666|000|9\\d{2})\\d{3}(?!00)\\d{2}(?!0{4})\\d{4}","type":"set","position":79},"8":{"id":8,"enabled":true,"dynamic":true,"name":"sounds_dir","value":"/usr/share/freeswitch/sounds","type":"set","position":8},"80":{"id":80,"enabled":true,"dynamic":false,"name":"df_luhn","value":"?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\\d{3})\\d{11}","type":"set","position":80},"81":{"id":81,"enabled":true,"dynamic":false,"name":"default_provider","value":"example.com","type":"set","position":81},"82":{"id":82,"enabled":true,"dynamic":false,"name":"default_provider_username","value":"joeuser","type":"set","position":82},"83":{"id":83,"enabled":true,"dynamic":false,"name":"default_provider_password","value":"password","type":"set","position":83},"84":{"id":84,"enabled":true,"dynamic":false,"name":"default_provider_from_domain","value":"example.com","type":"set","position":84},"85":{"id":85,"enabled":true,"dynamic":false,"name":"default_provider_register","value":"false","type":"set","position":85},"86":{"id":86,"enabled":true,"dynamic":false,"name":"default_provider_contact","value":"5000","type":"set","position":86},"87":{"id":87,"enabled":true,"dynamic":false,"name":"sip_tls_version","value":"tlsv1,tlsv1.1,tlsv1.2","type":"set","position":87},"88":{"id":88,"enabled":true,"dynamic":false,"name":"sip_tls_ciphers","value":"ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH","type":"set","position":88},"89":{"id":89,"enabled":true,"dynamic":false,"name":"internal_auth_calls","value":"true","type":"set","position":89},"9":{"id":9,"enabled":true,"dynamic":true,"name":"conf_dir","value":"/etc/freeswitch","type":"set","position":9},"90":{"id":90,"enabled":true,"dynamic":false,"name":"internal_tls_port","value":"5061","type":"set","position":90},"91":{"id":91,"enabled":true,"dynamic":false,"name":"internal_ssl_enable","value":"false","type":"set","position":91},"92":{"id":92,"enabled":true,"dynamic":false,"name":"external_auth_calls","value":"false","type":"set","position":92},"93":{"id":93,"enabled":true,"dynamic":false,"name":"external_sip_port","value":"5080","type":"set","position":93},"94":{"id":94,"enabled":true,"dynamic":false,"name":"external_tls_port","value":"5081","type":"set","position":94},"95":{"id":95,"enabled":true,"dynamic":false,"name":"external_ssl_enable","value":"false","type":"set","position":95},"96":{"id":96,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_in","value":"3mb","type":"set","position":96},"97":{"id":97,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_out","value":"3mb","type":"set","position":97},"98":{"id":98,"enabled":true,"dynamic":false,"name":"suppress_cng","value":"true","type":"set","position":98},"99":{"id":99,"enabled":true,"dynamic":false,"name":"video_mute_png","value":"/var/lib/freeswitch/images/default-mute.png","type":"set","position":99}}}
	//Errors:
	case "ImportGlobalVariables":
		resp = getUser(msg, ImportGlobalVariables, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	case "DialplanChangeNotProceed":
		resp = getUser(msg, SwitchDialplanNoProceed, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "DialplanGetSettings":
		resp = getUser(msg, DialplanGetSettings, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Get] Contexts":
		resp = getUser(msg, getDialplanContexts, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Import]":
		resp = getUser(msg, importDialplan, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Get] Extensions":
		resp = getUser(msg, getDialplanExtensions, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Get] Conditions":
		resp = getUser(msg, getDialplanConditions, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Get] Extension details":
		resp = getUser(msg, getDialplanExtenDetails, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Move] Extension":
		resp = getUser(msg, moveDialplanExtension, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Move] Condition":
		resp = getUser(msg, moveDialplanCondition, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Move] Action":
		resp = getUser(msg, moveDialplanAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Move] Antiaction":
		resp = getUser(msg, moveDialplanAntiAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Add] Regex":
		resp = getUser(msg, addRegex, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Add] Action":
		resp = getUser(msg, addAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Add] Antiaction":
		resp = getUser(msg, addAntiAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Delete] Regex":
		resp = getUser(msg, delRegex, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Delete] Action":
		resp = getUser(msg, delAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Delete] Antiaction":
		resp = getUser(msg, delAntiAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Update] Regex":
		resp = getUser(msg, updateRegex, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Update] Action":
		resp = getUser(msg, updateAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Update] Antiaction":
		resp = getUser(msg, updateAntiAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Switch] Regex":
		resp = getUser(msg, switchRegex, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Switch] Action":
		resp = getUser(msg, switchAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Switch] Antiaction":
		resp = getUser(msg, switchAntiAction, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Add] Context":
		resp = getUser(msg, addContext, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Add] Extension":
		resp = getUser(msg, addExtension, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Add] Condition":
		resp = getUser(msg, addCondition, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Rename] Context":
		resp = getUser(msg, renameContext, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Rename] Extension":
		resp = getUser(msg, renameExtension, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Delete] Context":
		resp = getUser(msg, deleteContext, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Delete] Extension":
		resp = getUser(msg, deleteExtension, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Switch] Extension Continue":
		resp = getUser(msg, switchExtensionContinue, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Update] Condition":
		resp = getUser(msg, updateCondition, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Switch] Condition":
		resp = getUser(msg, switchCondition, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Dialplan][Delete] Condition":
		resp = getUser(msg, deleteCondition, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users] Get":
		resp = getUser(msg, getWebUsers, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "GetWebUsersByDirectory":
		resp = getUser(msg, GetWebUsersByDirectory, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users] Add":
		resp = getUser(msg, addWebUsers, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users] Rename":
		resp = getUser(msg, renameWebUsers, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users] Delete":
		resp = getUser(msg, deleteWebUsers, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Switch] Web user":
		resp = getUser(msg, switchWebUser, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Password":
		resp = getUser(msg, updateWebUsersPassword, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Lang":
		resp = getUser(msg, updateWebUsersLang, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Sip user":
		resp = getUser(msg, updateWebUsersSipUser, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Ws":
		resp = getUser(msg, updateWebUsersWs, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Verto Ws":
		resp = getUser(msg, updateWebUsersVertoWs, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] WebRTC Lib":
		resp = getUser(msg, UpdateWebUserWebRTCLib, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Stun":
		resp = getUser(msg, updateWebUsersStun, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Update] Avatar":
		resp = getUser(msg, updateWebUsersAvatar, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Settings][Users][Clear] Avatar":
		resp = getUser(msg, clearWebUsersAvatar, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[CDR] Get":
		resp = getUser(msg, getCDR, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "GetHEP":
		resp = getUser(msg, getHEP, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "GetHEPDetails":
		resp = getUser(msg, GetHEPDetails, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "GetLogs":
		resp = getUser(msg, GetLogs, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "GetWebSettings":
		resp = getUser(msg, GetWebSettings, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "SaveWebSettings":
		resp = getUser(msg, SaveWebSettings, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "[Phone][Get] Creds":
		resp = getUser(msg, getPhoneCreds, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "get_status":
		resp = webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "connection"}
	//Request:
	//Response:
	//Errors:
	case "SendFSCLICommand":
		resp = getUser(msg, runCLICommand, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "RealFSCLIConnect":
		resp = getUser(msg, RealFSCLIConnect, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "RealFSCLICommand":
		resp = getUser(msg, RealFSCLICommand, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	case "UpdateWebUserGroup":
		resp = getUser(msg, UpdateWebUserGroup, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	case "GetWebDirectoryUsersTemplates":
		resp = getUser(msg, GetWebDirectoryUsersTemplates, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "AddWebDirectoryUsersTemplate":
		resp = getUser(msg, AddWebDirectoryUsersTemplate, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "DelWebDirectoryUsersTemplate":
		resp = getUser(msg, DelWebDirectoryUsersTemplate, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "SwitchWebDirectoryUsersTemplate":
		resp = getUser(msg, SwitchWebDirectoryUsersTemplate, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "UpdateWebDirectoryUsersTemplate":
		resp = getUser(msg, UpdateWebDirectoryUsersTemplate, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	case "GetWebDirectoryUsersTemplateParameters":
		resp = getUser(msg, GetWebDirectoryUsersTemplateParameters, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "AddWebDirectoryUsersTemplateParameter":
		resp = getUser(msg, AddWebDirectoryUsersTemplateParameter, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "DelWebDirectoryUsersTemplateParameter":
		resp = getUser(msg, DelWebDirectoryUsersTemplateParameter, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "SwitchWebDirectoryUsersTemplateParameter":
		resp = getUser(msg, SwitchWebDirectoryUsersTemplateParameter, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "UpdateWebDirectoryUsersTemplateParameter":
		resp = getUser(msg, UpdateWebDirectoryUsersTemplateParameter, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	case "GetWebDirectoryUsersTemplateVariables":
		resp = getUser(msg, GetWebDirectoryUsersTemplateVariables, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "AddWebDirectoryUsersTemplateVariable":
		resp = getUser(msg, AddWebDirectoryUsersTemplateVariable, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "DelWebDirectoryUsersTemplateVariable":
		resp = getUser(msg, DelWebDirectoryUsersTemplateVariable, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "SwitchWebDirectoryUsersTemplateVariable":
		resp = getUser(msg, SwitchWebDirectoryUsersTemplateVariable, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "UpdateWebDirectoryUsersTemplateVariable":
		resp = getUser(msg, UpdateWebDirectoryUsersTemplateVariable, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	case "GetWebDirectoryUsersTemplatesList":
		resp = getUser(msg, GetWebDirectoryUsersTemplatesList, onlyAdminAndManagerGroup())
	//Request:
	//Response:
	//Errors:
	case "GetWebDirectoryUsersTemplateForm":
		resp = getUser(msg, GetWebDirectoryUsersTemplateForm, onlyAdminAndManagerGroup())
	//Request:
	//Response:
	//Errors:
	case "CreateWebDirectoryUsersByTemplate":
		resp = getUser(msg, CreateWebDirectoryUsersByTemplate, onlyAdminAndManagerGroup())

	//Request:
	//Response:
	//Errors:
	case "UpdateAutoDialerListMember":
		resp = UpdateAutoDialerListMember(msg)
	//Request:
	//Response:
	//Errors:
	case "AddAutoDialerListMembers":
		resp = getUser(msg, AddAutoDialerListMembers, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case "GetAutoDialerListMembers":
		resp = getUserForConfig(msg, getByStruct, &apps.AutoDialerListMember{}, onlyAdminGroup())
	default:
		foo, exists := apps.WebCases[msg.Event]
		if exists {
			return getUser(msg, foo, onlyAdminAndManagerGroup())
		}
		resp = webStruct.UserResponse{Error: "Wrong event", MessageType: "none"}
	}

	return resp
}

func checkRelogin(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	return webStruct.UserResponse{User: user, Token: data.Token, MessageType: "relogin"}
}

func checkSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	return webStruct.UserResponse{Settings: &cfg.CustomPbx, MessageType: "settings"}
}

func setSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	log.Println(data.Payload)
	if data.Payload.Fs.Esl.Pass == "" || data.Payload.Fs.Esl.Port == 0 || data.Payload.Fs.Esl.Host == "" ||
		data.Payload.Db.Host == "" || data.Payload.Db.Port == 0 || data.Payload.Db.Name == "" ||
		data.Payload.Db.User == "" || data.Payload.Db.Pass == "" ||
		data.Payload.Web.Host == "" || data.Payload.Web.Port == 0 ||
		data.Payload.Web.Route == "" || data.Payload.Web.CertPath == "" || data.Payload.Web.KeyPath == "" ||
		data.Payload.XMLCurl.Host == "" || data.Payload.XMLCurl.Port == 0 ||
		data.Payload.XMLCurl.Route == "" || data.Payload.XMLCurl.CertPath == "" || data.Payload.XMLCurl.KeyPath == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: "settings"}
	}
	cfg.CustomPbx.Fs.Esl = data.Payload.Fs.Esl
	cfg.CustomPbx.Db = data.Payload.Db
	conf, err := cfg.WD(cfg.CustomPbx)
	if err != nil {
		cfg.RD()
		return webStruct.UserResponse{Error: "can't save", MessageType: "settings"}
	}

	return webStruct.UserResponse{Settings: &conf, MessageType: "settings"}
}

func getCDR(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 250 {
		limit = 25
	}
	offset := data.DBRequest.Offset * limit
	cdr, err := cdrDb.GetList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	/*
		if cdr == nil {
			return webStruct.UserResponse{Error: "nothing", MessageType: data.Event}
		}*/

	return webStruct.UserResponse{CDR: &cdr, MessageType: data.Event}
}

func getPhoneCreds(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if !user.SipId.Valid {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}

	userI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: user.SipId.Int64})
	if err != nil || userI == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	directoryUser, ok := userI.(altStruct.DirectoryDomainUser)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	param, err := intermediateDB.GetByValue(
		&altStruct.DirectoryDomainUserParameter{Name: "password", Parent: &directoryUser},
		map[string]bool{"Parent": true, "Name": true},
	)
	if err != nil || len(param) == 0 {
		return webStruct.UserResponse{Error: "user password not found", MessageType: data.Event}
	}
	directoryUserParam, ok := param[0].(altStruct.DirectoryDomainUserParameter)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	domainI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomain{Id: directoryUser.Parent.Id})
	if err != nil || domainI == nil {
		return webStruct.UserResponse{Error: "user domain not found", MessageType: data.Event}
	}
	domain, ok := domainI.(altStruct.DirectoryDomain)
	if !ok {
		return webStruct.UserResponse{Error: "user domain not found", MessageType: data.Event}
	}

	password := directoryUserParam.Value

	if password == "" {
		paramI, err := intermediateDB.GetByValue(
			&altStruct.DirectoryDomainParameter{Name: "password", Parent: &domain},
			map[string]bool{"Parent": true, "Name": true},
		)
		if err != nil || len(paramI) == 0 {
			return webStruct.UserResponse{Error: "domain directory password not found", MessageType: data.Event}
		}
		domainParam, ok := paramI[0].(altStruct.DirectoryDomainParameter)
		if !ok {
			return webStruct.UserResponse{Error: "domain directory not found", MessageType: data.Event}
		}

		password = domainParam.Value
	}
	if password == "" || (user.Ws == "" && user.VertoWs == "") /*|| user.Stun == ""*/ {
		return webStruct.UserResponse{Error: "no enough params params", MessageType: data.Event}
	}

	creds := webStruct.PhoneCreds{}
	creds.UserName = directoryUser.Name
	creds.Password = password
	creds.Domain = domain.Name
	creds.WebRTCLib = user.WebRTCLib
	creds.Ws = user.Ws
	creds.VertoWs = user.VertoWs
	creds.Stun = user.Stun
	if creds.Stun == "" && daemonCache.State.StunServerStatus {
		creds.Stun = "stun:" + cfg.CustomPbx.Web.Host + ":" + strconv.Itoa(cfg.CustomPbx.Web.StunPort)
	}

	return webStruct.UserResponse{PhoneCreds: &creds, MessageType: data.Event}
}

func runCLICommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func RealFSCLIConnect(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func RealFSCLICommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func GetLogs(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 5000 {
		limit = 250
	}
	offset := data.DBRequest.Offset * limit
	logs, err := db.GetList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order, cache.GetCurrentInstanceId())
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Logs: &logs, MessageType: data.Event}
}

func getHEP(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 5000 {
		limit = 250
	}
	offset := data.DBRequest.Offset * limit
	heps, err := db.GetHEPList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{HEPs: &heps, MessageType: data.Event}
}

func GetHEPDetails(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if len(data.ArrVal) == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	heps, err := db.GetHEPDetailsList(data.ArrVal, cache.GetCurrentInstanceId())
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{HEPsDetails: &heps, MessageType: data.Event}
}

func GetInstances(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	cache.UpdateCacheInstances()
	var res = cache.GetFSInstances().GetList()
	var currentId = cache.GetCurrentInstanceId()
	return webStruct.UserResponse{FSInstances: &res, MessageType: "GetInstances", Id: &currentId}
}

func UpdateInstanceDescription(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	instance := cache.GetFSInstances().GetById(data.Id)
	if instance == nil {
		return webStruct.UserResponse{Error: "instance not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateFSInstanceDescription(instance, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.FsInstance{instance.Id: instance}

	return webStruct.UserResponse{MessageType: data.Event, FSInstances: &item}

}

func getCallcenterAgents(msg *webStruct.MessageData) webStruct.UserResponse {
	if msg.Param.Name == "" || msg.Param.Value == "" {
		return webStruct.UserResponse{Error: "wrong params", MessageType: msg.Event}
	}
	item := &altStruct.Agent{Id: msg.Param.Id}
	fieldName := mainStruct.GetItemNameByTag(item, msg.Param.Name)
	if fieldName == "id" {
		return webStruct.UserResponse{Error: "please dont", MessageType: msg.Event}
	}
	f := reflect.ValueOf(item).Elem().FieldByName(fieldName)
	switch f.Type().Name() {
	case "string":
		f.SetString(msg.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(msg.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(msg.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}
	// and update with updateConfig func
	return getUserForConfig(msg, updateConfig, struct {
		S interface{}
		A []string
	}{item, []string{fieldName}}, onlyAdminGroup())
}

func getCallcenterTiers(msg *webStruct.MessageData) webStruct.UserResponse {
	if msg.Param.Name == "" || msg.Param.Value == "" {
		return webStruct.UserResponse{Error: "wrong params", MessageType: msg.Event}
	}
	item := &altStruct.Tier{Id: msg.Param.Id}
	fieldName := mainStruct.GetItemNameByTag(item, msg.Param.Name)
	if fieldName == "id" {
		return webStruct.UserResponse{Error: "please dont", MessageType: msg.Event}
	}
	f := reflect.ValueOf(item).Elem().FieldByName(fieldName)
	switch f.Type().Name() {
	case "string":
		f.SetString(msg.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(msg.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(msg.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}

	return getUserForConfig(msg, updateConfig, struct {
		S interface{}
		A []string
	}{item, []string{fieldName}}, onlyAdminGroup())
}

// like getConfig
func getByStruct(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no parent id", MessageType: data.Event}
	}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true, UseValue: true, Value: data.Id}}

	var res interface{}
	var err error
	if data.DBRequest.Limit != 0 || data.DBRequest.Filters != nil {
		limit := data.DBRequest.Limit
		if limit < 25 || limit > 250 {
			limit = 25
		}
		offset := 0
		if data.DBRequest.Offset > 0 {
			offset = data.DBRequest.Offset * limit
		}
		for _, v := range data.DBRequest.Filters {
			filter[v.Field] = customorm.FilterFields{Flag: true, UseValue: true, Value: v.FieldValue, Operand: v.Operand}
		}
		filterStr := customorm.Filters{
			Fields: filter,
			Limit:  limit,
			Offset: offset,
			Order:  customorm.Order{Desc: data.DBRequest.Order.Desc, Fields: data.DBRequest.Order.Fields}}
		res, err = intermediateDB.GetByFilteredValues(
			item,
			filterStr,
		)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
		}
		//TODO: with total all the time
		filterStr.Count = true
		resCount, err := intermediateDB.GetByFilteredValues(
			item,
			filterStr,
		)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
		}
		if len(resCount) == 0 {
			return webStruct.UserResponse{Error: "can't count total", MessageType: data.Event}
		}
		total, ok := resCount[0].(int64)
		if !ok {
			return webStruct.UserResponse{Error: "can't get total", MessageType: data.Event}
		}
		res = struct {
			Items interface{} `json:"items"`
			Total int64       `json:"total"`
		}{Items: res, Total: total}
	} else {
		res, err = intermediateDB.GetByValuesAsMap(
			item,
			filter,
		)
	}

	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: res, MessageType: data.Event}
}

func AddAutoDialerListMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}
	var items []struct {
		Name       string `json:"name"`
		ToNumber   string `json:"to_number"`
		FromNumber string `json:"from_number"`
		Retries    string `json:"retries"`
		CustomVars string `json:"custom_vars"`
	}
	err := json.Unmarshal(data.Data, &items)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	counter := 0
	for _, sub := range items {
		retry, _ := strconv.ParseInt(sub.Retries, 10, 64)
		item := &apps.AutoDialerListMember{
			Name:       sub.Name,
			ToNumber:   sub.ToNumber,
			FromNumber: sub.FromNumber,
			Retries:    retry,
			CustomVars: sub.CustomVars,
			Enabled:    true,
			Parent:     &apps.AutoDialerList{Id: data.Id},
		}
		_, err := intermediateDB.InsertItem(item)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		counter++
	}
	return webStruct.UserResponse{MessageType: data.Event, Total: &counter}
}

func UpdateAutoDialerListMember(msg *webStruct.MessageData) webStruct.UserResponse {
	if msg.Param.Name == "" || msg.Param.Value == "" {
		return webStruct.UserResponse{Error: "wrong params", MessageType: msg.Event}
	}
	item := &apps.AutoDialerListMember{Id: msg.Param.Id}
	fieldName := mainStruct.GetItemNameByTag(item, msg.Param.Name)
	if fieldName == "id" {
		return webStruct.UserResponse{Error: "please dont", MessageType: msg.Event}
	}
	f := reflect.ValueOf(item).Elem().FieldByName(fieldName)
	switch f.Type().Name() {
	case "string":
		f.SetString(msg.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(msg.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(msg.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}

	return getUserForConfig(msg, updateConfig, struct {
		S interface{}
		A []string
	}{item, []string{fieldName}}, onlyAdminGroup())
}
