package web

import (
	"crypto/rand"
	"custompbx/altStruct"
	"custompbx/cfg"
	"custompbx/daemonCache"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"custompbx/webcache"
	"encoding/json"
	"fmt"
	"github.com/custompbx/hepparser"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

var eventChannel chan interface{}
var messageUserLookup = findUser

func onlyAdminGroup() []int {
	return []int{mainStruct.GetAdminId()}
}

func onlyAdminAndManagerGroup() []int {
	return []int{mainStruct.GetAdminId(), mainStruct.GetManagerId()}
}

func adminOnly() []int {
	return onlyAdminGroup()
}

func adminOrManager() []int {
	return onlyAdminAndManagerGroup()
}

func onlyAdminManagerAndUserGroup() []int {
	return []int{mainStruct.GetAdminId(), mainStruct.GetManagerId(), mainStruct.GetUserId()}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return CheckWebSocketOrigin(cfg.CustomPbx.Web.OriginPolicy, cfg.CustomPbx.Web.AllowedOrigins, r)
	},
}

func SetBroadcastChannel(brChannel chan interface{}) {
	eventChannel = brChannel
}

var b = webStruct.NewWsHub()

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
	b.Register(wsContext)

	fmt.Println("STARTING GOROUTINES")
	go wsContext.SendWaiter()
	go wsContext.ReadWaiter(messageHandler)
}

func HubMetrics(w http.ResponseWriter, r *http.Request) {
	user, status := UserFromBearer(r)
	if status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}
	if status := RequireGroups(user, mainStruct.GetAdminId()); status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(b.Metrics()); err != nil {
		log.Printf("component=websocket operation=write_metrics error=%q", err)
	}
}

func tokenGenerator() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func messageHandler(msg *webStruct.Message, wsContext *webStruct.WsContext) {
	defer func() {
		if r := recover(); r != nil {
			wsContext.RecordHandlerFailure()
			log.Printf("component=websocket connection_id=%d user_id=%d operation=message_handler error=%v", wsContext.ID, wsContext.UserID(), r)
			log.Printf("component=websocket connection_id=%d user_id=%d operation=message_handler stacktrace=%q", wsContext.ID, wsContext.UserID(), string(debug.Stack()))
		}
	}()
	if err := msg.Validate(); err != nil {
		wsContext.RecordHandlerFailure()
		sendWSResponse(wsContext, &webStruct.UserResponse{Error: "invalid message", MessageType: "none"})
		return
	}
	if !daemonCache.State.DatabaseConnection {
		sendWSResponse(wsContext, &webStruct.UserResponse{Daemon: daemonCache.State, MessageType: webStruct.BroadcastConnection})
		return
	}

	msg.Data.Trim()
	msg.Data.Event = msg.Event
	msg.Data.Context = wsContext

	// first check if it login request
	if msg.Event == eventLogin {
		resp := checkLogin(msg.Data)
		sendWSResponse(wsContext, &resp)
		return
	}
	// allow without token
	if msg.Event == "get_status" {
		resp := webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "connection"}
		sendWSResponse(wsContext, &resp)
		return
	}
	log.Println("EVENT: ", msg.Event)

	// find user by token
	user, response := messageUserLookup(msg.Data)
	if user == nil {
		log.Println("EVENT: ", msg.Event, "NO USER")
		sendWSResponse(wsContext, &response)
		return
	}
	subsResp := subscribeUser(msg.Data)
	if subsResp != nil {
		log.Println("EVENT: ", msg.Event, "NO SUBS")
		sendWSResponse(wsContext, subsResp)
		return
	}

	resp := dispatchMessage(msg.Data, wsContext)
	sendWSResponse(wsContext, &resp)
}

func dispatchMessage(data *webStruct.MessageData, wsContext *webStruct.WsContext) webStruct.UserResponse {
	if registeredResponse, ok := coreEvents.Dispatch(data, wsContext); ok {
		return registeredResponse
	}
	return messageMainHandler(data)
}

func sendWSResponse(wsContext *webStruct.WsContext, resp *webStruct.UserResponse) {
	if !wsContext.Enqueue(resp) {
		_ = wsContext.CloseWithReason("outbound queue full")
	}
}
