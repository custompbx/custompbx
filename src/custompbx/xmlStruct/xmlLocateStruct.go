package xmlStruct

import "encoding/xml"

type X_NO_PRE_PROCESS struct {
	XMLName  xml.Name `xml:"X-NO-PRE-PROCESS,omitempty" json:"X-NO-PRE-PROCESS,omitempty"`
	Attrcmd  string   `xml:"cmd,attr"  json:",omitempty"`
	Attrdata string   `xml:"data,attr"  json:",omitempty"`
}

type Action struct {
	XMLName         xml.Name `xml:"action,omitempty" json:"action,omitempty"`
	Attrapplication string   `xml:"application,attr"  json:",omitempty"`
	Attrdata        string   `xml:"data,attr"  json:",omitempty"`
	Attrfunction    string   `xml:"function,attr,omitempty"  json:",omitempty"`
	Attrmethod      string   `xml:"method,attr,omitempty"  json:",omitempty"`
	Attrphrase      string   `xml:"phrase,attr,omitempty"  json:",omitempty"`
	Attrtype        string   `xml:"type,attr,omitempty"  json:",omitempty"`
	Attrinline      string   `xml:"inline,attr,omitempty"  json:",omitempty"`
	string          string   `xml:",chardata" json:",omitempty"`
}

type Regex struct {
	XMLName        xml.Name `xml:"regex,omitempty" json:"action,omitempty"`
	Attrexpression string   `xml:"expression,attr"  json:",omitempty"`
	Attrfield      string   `xml:"field,attr"  json:",omitempty"`
	string         string   `xml:",chardata" json:",omitempty"`
}

type Advertise struct {
	XMLName xml.Name `xml:"advertise,omitempty" json:"advertise,omitempty"`
	Room    []*Room  `xml:"room,omitempty" json:"room,omitempty"`
}

type Agent struct {
	XMLName             xml.Name `xml:"agent,omitempty" json:"agent,omitempty"`
	AttrbusyDelayTime   string   `xml:"busy-delay-time,attr"  json:",omitempty"`
	Attrcontact         string   `xml:"contact,attr"  json:",omitempty"`
	AttrmaxNoAnswer     string   `xml:"max-no-answer,attr"  json:",omitempty"`
	Attrname            string   `xml:"name,attr"  json:",omitempty"`
	AttrrejectDelayTime string   `xml:"reject-delay-time,attr"  json:",omitempty"`
	Attrstatus          string   `xml:"status,attr"  json:",omitempty"`
	Attrtype            string   `xml:"type,attr"  json:",omitempty"`
	AttrwrapUpTime      string   `xml:"wrap-up-time,attr"  json:",omitempty"`
}

type Agents struct {
	XMLName xml.Name `xml:"agents,omitempty" json:"agents,omitempty"`
	Agent   []*Agent `xml:"agent,omitempty" json:"agent,omitempty"`
}

type Aliases struct {
	XMLName xml.Name `xml:"aliases,omitempty" json:"aliases,omitempty"`
	Alias   []*Alias `xml:"alias,omitempty" json:"alias,omitempty"`
}

type Alias struct {
	XMLName  xml.Name `xml:"alias,omitempty" json:"alias,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
}

type AntiAction struct {
	XMLName         xml.Name `xml:"anti-action,omitempty" json:"anti-action,omitempty"`
	Attrapplication string   `xml:"application,attr"  json:",omitempty"`
	Attrdata        string   `xml:"data,attr"  json:",omitempty"`
}

type Api struct {
	XMLName         xml.Name `xml:"api,omitempty" json:"api,omitempty"`
	Attrargument    string   `xml:"argument,attr"  json:",omitempty"`
	Attrdescription string   `xml:"description,attr"  json:",omitempty"`
	Attrdestination string   `xml:"destination,attr"  json:",omitempty"`
	Attrname        string   `xml:"name,attr"  json:",omitempty"`
	Attrparse       string   `xml:"parse,attr"  json:",omitempty"`
	Attrsyntax      string   `xml:"syntax,attr"  json:",omitempty"`
	Attrvalue       string   `xml:"value,attr"  json:",omitempty"`
}

type Apis struct {
	XMLName xml.Name `xml:"apis,omitempty" json:"apis,omitempty"`
	Api     []*Api   `xml:"api,omitempty" json:"api,omitempty"`
}

type Application struct {
	XMLName  xml.Name `xml:"application,omitempty" json:"application,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
}

type ApplicationList struct {
	XMLName     xml.Name       `xml:"application-list,omitempty" json:"application-list,omitempty"`
	Attrdefault string         `xml:"default,attr"  json:",omitempty"`
	Application []*Application `xml:"application,omitempty" json:"application,omitempty"`
}

type VariableList struct {
	XMLName     xml.Name    `xml:"variable-list,omitempty" json:"variable-list,omitempty"`
	Attrdefault string      `xml:"default,attr"  json:",omitempty"`
	Variable    []*Variable `xml:"variable,omitempty" json:"variable,omitempty"`
}

type ApiList struct {
	XMLName     xml.Name    `xml:"api-list,omitempty" json:"api-list,omitempty"`
	Attrdefault string      `xml:"default,attr"  json:",omitempty"`
	Variable    []*Variable `xml:"variable,omitempty" json:"api,omitempty"`
}

type Binding struct {
	XMLName  xml.Name `xml:"binding,omitempty" json:"binding,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Bindings struct {
	XMLName xml.Name   `xml:"bindings,omitempty" json:"bindings,omitempty"`
	Binding []*Binding `xml:"binding,omitempty" json:"binding,omitempty"`
}

type Button struct {
	XMLName        xml.Name `xml:"button,omitempty" json:"button,omitempty"`
	AttrcallerName string   `xml:"caller-name,attr"  json:",omitempty"`
	Attrlabel      string   `xml:"label,attr"  json:",omitempty"`
	Attrposition   string   `xml:"position,attr"  json:",omitempty"`
	Attrtype       string   `xml:"type,attr"  json:",omitempty"`
	Attrvalue      string   `xml:"value,attr"  json:",omitempty"`
}

type Buttons struct {
	XMLName xml.Name  `xml:"buttons,omitempty" json:"buttons,omitempty"`
	Button  []*Button `xml:"button,omitempty" json:"button,omitempty"`
}

type CallerControls struct {
	XMLName xml.Name `xml:"caller-controls,omitempty" json:"caller-controls,omitempty"`
	Group   []*Group `xml:"group,omitempty" json:"group,omitempty"`
}

type ChatPermissions struct {
	XMLName xml.Name   `xml:"chat-permissions,omitempty" json:"chat-permissions,omitempty"`
	Profile []*Profile `xml:"group,omitempty" json:"group,omitempty"`
}

type CliKeybindings struct {
	XMLName xml.Name `xml:"cli-keybindings,omitempty" json:"cli-keybindings,omitempty"`
	Key     []*Key   `xml:"key,omitempty" json:"key,omitempty"`
}

type Commands struct {
	XMLName xml.Name   `xml:"commands,omitempty" json:"commands,omitempty"`
	Profile []*Profile `xml:"profile,omitempty" json:"profile,omitempty"`
}

type Condition struct {
	XMLName        xml.Name      `xml:"condition,omitempty" json:"condition,omitempty"`
	Attrbreak      string        `xml:"break,attr,omitempty"  json:",omitempty"`
	Attrexpression string        `xml:"expression,attr"  json:",omitempty"`
	Attrfield      string        `xml:"field,attr"  json:",omitempty"`
	Attrhour       string        `xml:"hour,attr,omitempty"  json:",omitempty"`
	Attrmday       string        `xml:"mday,attr,omitempty"  json:",omitempty"`
	Attrmon        string        `xml:"mon,attr,omitempty"  json:",omitempty"`
	Attrmweek      string        `xml:"mweek,attr,omitempty"  json:",omitempty"`
	Attrwday       string        `xml:"wday,attr,omitempty"  json:",omitempty"`
	Action         []*Action     `xml:"action,omitempty" json:"action,omitempty"`
	AntiAction     []*AntiAction `xml:"anti-action,omitempty" json:"anti-action,omitempty"`
	//Expression     string   `xml:"expression,omitempty" json:"expression,omitempty"`
	DateTime  string   `xml:"date-time,attr,omitempty" json:"date-time"`
	TimeOfDay string   `xml:"time-of-day,attr,omitempty" json:"time-of-day"`
	Year      string   `xml:"year,attr,omitempty" json:"year"`
	Minute    string   `xml:"minute,attr,omitempty" json:"minute"`
	Week      string   `xml:"week,attr,omitempty" json:"week"`
	Yday      string   `xml:"yday,attr,omitempty" json:"yday"`
	Minday    string   `xml:"minday,attr,omitempty" json:"minday"`
	TzOffset  string   `xml:"tz-offset,attr,omitempty" json:"tz-offset"`
	Dst       string   `xml:"dst,attr,omitempty" json:"dst"`
	Attrregex string   `xml:"regex,attr,omitempty" json:"regex,omitempty"`
	Regex     []*Regex `xml:"regex,omitempty" json:"regex,omitempty"`
}

type Conference struct {
	XMLName xml.Name `xml:"conference,omitempty" json:"conference,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Configuration struct {
	XMLName         xml.Name         `xml:"configuration,omitempty" json:"configuration,omitempty"`
	Attrdescription string           `xml:"description,attr"  json:",omitempty"`
	Attrname        string           `xml:"name,attr"  json:",omitempty"`
	Advertise       *Advertise       `xml:"advertise,omitempty" json:"advertise,omitempty"`
	Agents          *Agents          `xml:"agents,omitempty" json:"agents,omitempty"`
	Apis            *Apis            `xml:"apis,omitempty" json:"apis,omitempty"`
	Bindings        *Bindings        `xml:"bindings,omitempty" json:"bindings,omitempty"`
	CallerControls  *CallerControls  `xml:"caller-controls,omitempty" json:"caller-controls,omitempty"`
	ChatPermissions *ChatPermissions `xml:"chat-permissions,omitempty" json:"chat-permissions,omitempty"`
	CliKeybindings  *CliKeybindings  `xml:"cli-keybindings,omitempty" json:"cli-keybindings,omitempty"`
	Commands        *Commands        `xml:"commands,omitempty" json:"commands,omitempty"`
	DefaultPtimes   *DefaultPtimes   `xml:"default-ptimes,omitempty" json:"default-ptimes,omitempty"`
	Descriptors     *Descriptors     `xml:"descriptors,omitempty" json:"descriptors,omitempty"`
	Directory       []*Directory     `xml:"directory,omitempty" json:"directory,omitempty"`
	Domains         *Domains         `xml:"domains,omitempty" json:"domains,omitempty"`
	Endpoints       *Endpoints       `xml:"endpoints,omitempty" json:"endpoints,omitempty"`
	EventFilter     *EventFilter     `xml:"event-filter,omitempty" json:"event-filter,omitempty"`
	FaxSettings     *FaxSettings     `xml:"fax-settings,omitempty" json:"fax-settings,omitempty"`
	Feeds           *Feeds           `xml:"feeds,omitempty" json:"feeds,omitempty"`
	Fifos           *Fifos           `xml:"fifos,omitempty" json:"fifos,omitempty"`
	Gateways        *Gateways        `xml:"gateways,omitempty" json:"gateways,omitempty"`
	GlobalSettings  *GlobalSettings  `xml:"global_settings,omitempty" json:"global_settings,omitempty"`
	Javavm          *Javavm          `xml:"javavm,omitempty" json:"javavm,omitempty"`
	LayoutSettings  *LayoutSettings  `xml:"layout-settings,omitempty" json:"layout-settings,omitempty"`
	Listeners       *Listeners       `xml:"listeners,omitempty" json:"listeners,omitempty"`
	Lists           *Lists           `xml:"lists,omitempty" json:"lists,omitempty"`
	Logging         *Logging         `xml:"logging,omitempty" json:"logging,omitempty"`
	Mappings        *Mappings        `xml:"mappings,omitempty" json:"mappings,omitempty"`
	Menus           *Menus           `xml:"menus,omitempty" json:"menus,omitempty"`
	ModemSettings   *ModemSettings   `xml:"modem-settings,omitempty" json:"modem-settings,omitempty"`
	Modules         *Modules         `xml:"modules,omitempty" json:"modules,omitempty"`
	NetworkLists    *NetworkLists    `xml:"network-lists,omitempty" json:"network-lists,omitempty"`
	Options         *Options         `xml:"options,omitempty" json:"options,omitempty"`
	Producers       *Producers       `xml:"producers,omitempty" json:"producers,omitempty"`
	Profiles        *Profiles        `xml:"profiles,omitempty" json:"profiles,omitempty"`
	Queues          *Queues          `xml:"queues,omitempty" json:"queues,omitempty"`
	Remotes         *Remotes         `xml:"remotes,omitempty" json:"remotes,omitempty"`
	Routes          *Routes          `xml:"routes,omitempty" json:"routes,omitempty"`
	Schema          *Schema          `xml:"schema,omitempty" json:"schema,omitempty"`
	Settings        *Settings        `xml:"settings,omitempty" json:"settings,omitempty"`
	Spans           *Spans           `xml:"spans,omitempty" json:"spans,omitempty"`
	Startup         *Startup         `xml:"startup,omitempty" json:"startup,omitempty"`
	Streams         *Streams         `xml:"streams,omitempty" json:"streams,omitempty"`
	Templates       *Templates       `xml:"templates,omitempty" json:"templates,omitempty"`
	Tiers           *Tiers           `xml:"tiers,omitempty" json:"tiers,omitempty"`
	Timezones       *Timezones       `xml:"timezones,omitempty" json:"timezones,omitempty"`
	XProfile        []*XProfile      `xml:"x-profile,omitempty" json:"x-profile,omitempty"`
	Tables          *Tables          `xml:"tables,omitempty" json:"tables,omitempty"`
}

type Connection struct {
	XMLName  xml.Name `xml:"connection,omitempty" json:"connection,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Connections struct {
	XMLName    xml.Name      `xml:"connections,omitempty" json:"connections,omitempty"`
	Connection []*Connection `xml:"connection,omitempty" json:"connection,omitempty"`
}

type Context struct {
	XMLName   xml.Name     `xml:"context,omitempty" json:"context,omitempty"`
	Attrname  string       `xml:"name,attr"  json:",omitempty"`
	Extension []*Extension `xml:"extension,omitempty" json:"extension,omitempty"`
}

type Control struct {
	XMLName    xml.Name `xml:"control,omitempty" json:"control,omitempty"`
	Attraction string   `xml:"action,attr"  json:",omitempty"`
	Attrdigits string   `xml:"digits,attr"  json:",omitempty"`
}

type DefaultPtimes struct {
	XMLName xml.Name `xml:"default-ptimes,omitempty" json:"default-ptimes,omitempty"`
	Codec   []*Codec `xml:"codec,omitempty" json:"codec,omitempty"`
}

type Codec struct {
	XMLName   xml.Name `xml:"codec,omitempty" json:"control,omitempty"`
	Attrname  string   `xml:"name,attr"  json:",omitempty"`
	Attrptime string   `xml:"ptime,attr"  json:"ptime,omitempty"`
}

type Descriptor struct {
	XMLName  xml.Name `xml:"descriptor,omitempty" json:"descriptor,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Tone     []*Tone  `xml:"tone,omitempty" json:"tone,omitempty"`
}

type Descriptors struct {
	XMLName    xml.Name      `xml:"descriptors,omitempty" json:"descriptors,omitempty"`
	Descriptor []*Descriptor `xml:"descriptor,omitempty" json:"descriptor,omitempty"`
}

type DeviceType struct {
	XMLName xml.Name `xml:"device-type,omitempty" json:"device-type,omitempty"`
	Attrid  string   `xml:"id,attr"  json:",omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type DeviceTypes struct {
	XMLName    xml.Name    `xml:"device-types,omitempty" json:"device-types,omitempty"`
	DeviceType *DeviceType `xml:"device-type,omitempty" json:"device-type,omitempty"`
}

type Dial struct {
	XMLName xml.Name `xml:"dial,omitempty" json:"dial,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Directory struct {
	XMLName  xml.Name `xml:"directory,omitempty" json:"directory,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Attrpath string   `xml:"path,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Document struct {
	XMLName  xml.Name   `xml:"document,omitempty" json:"document,omitempty"`
	Attrtype string     `xml:"type,attr"  json:",omitempty"`
	Section  []*Section `xml:"section,omitempty" json:"section,omitempty"`
}

type Domain struct {
	XMLName   xml.Name   `xml:"domain,omitempty" json:"domain,omitempty"`
	Attralias string     `xml:"alias,attr"  json:",omitempty"`
	Attrname  string     `xml:"name,attr"  json:",omitempty"`
	Attrparse string     `xml:"parse,attr"  json:",omitempty"`
	Exten     *Exten     `xml:"exten,omitempty" json:"exten,omitempty"`
	Groups    *Groups    `xml:"groups,omitempty" json:"groups,omitempty"`
	Params    *Params    `xml:"params,omitempty" json:"params,omitempty"`
	Variables *Variables `xml:"variables,omitempty" json:"variables,omitempty"`
	Users     *Users     `xml:"users,omitempty" json:"users,omitempty"`
}

type Domains struct {
	XMLName xml.Name  `xml:"domains,omitempty" json:"domains,omitempty"`
	Domain  []*Domain `xml:"domain,omitempty" json:"domain,omitempty"`
}

type Element struct {
	XMLName   xml.Name `xml:"element,omitempty" json:"element,omitempty"`
	Attrfreq1 string   `xml:"freq1,attr"  json:",omitempty"`
	Attrfreq2 string   `xml:"freq2,attr"  json:",omitempty"`
	Attrmax   string   `xml:"max,attr"  json:",omitempty"`
	Attrmin   string   `xml:"min,attr"  json:",omitempty"`
}

type Email struct {
	XMLName xml.Name `xml:"email,omitempty" json:"email,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Endpoint struct {
	XMLName  xml.Name `xml:"endpoint,omitempty" json:"endpoint,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Endpoints struct {
	XMLName  xml.Name    `xml:"endpoints,omitempty" json:"endpoints,omitempty"`
	Endpoint []*Endpoint `xml:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type Entry struct {
	XMLName    xml.Name `xml:"entry,omitempty" json:"entry,omitempty"`
	Attraction string   `xml:"action,attr"  json:",omitempty"`
	Attrdigits string   `xml:"digits,attr"  json:",omitempty"`
	Attrparam  string   `xml:"param,attr"  json:",omitempty"`
}

type EventFilter struct {
	XMLName  xml.Name  `xml:"event-filter,omitempty" json:"event-filter,omitempty"`
	Attrtype string    `xml:"type,attr"  json:",omitempty"`
	Header   []*Header `xml:"header,omitempty" json:"header,omitempty"`
}

type Expression struct {
	XMLName xml.Name `xml:"expression,omitempty" json:"expression,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type Exten struct {
	XMLName   xml.Name `xml:"exten,omitempty" json:"exten,omitempty"`
	Attrproto string   `xml:"proto,attr"  json:",omitempty"`
	Attrregex string   `xml:"regex,attr"  json:",omitempty"`
}

type Extension struct {
	XMLName      xml.Name     `xml:"extension,omitempty" json:"extension,omitempty"`
	Attrcontinue string       `xml:"continue,attr,omitempty"  json:",omitempty"`
	Attrname     string       `xml:"name,attr"  json:",omitempty"`
	Condition    []*Condition `xml:"condition,omitempty" json:"condition,omitempty"`
}

type FaxSettings struct {
	XMLName xml.Name `xml:"fax-settings,omitempty" json:"fax-settings,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Feeds struct {
	XMLName xml.Name `xml:"feeds,omitempty" json:"feeds,omitempty"`
}

type Field struct {
	XMLName         xml.Name `xml:"field,omitempty" json:"field,omitempty"`
	Attrquote       string   `xml:"quote,attr"  json:",omitempty"`
	Attrvar         string   `xml:"var,attr"  json:",omitempty"`
	Attrcolumn      string   `xml:"column,attr"  json:",omitempty"`
	Attrname        string   `xml:"name,attr"  json:",omitempty"`
	AttrchanVarName string   `xml:"chan-var-name,attr"  json:",omitempty"`
}

type Fields struct {
	XMLName xml.Name `xml:"fields,omitempty" json:"fields,omitempty"`
}

type Member struct {
	XMLName     xml.Name `xml:"member,omitempty" json:"member,omitempty"`
	Attrtimeout string   `xml:"timeout,attr"  json:",omitempty"`
	Attrsimo    string   `xml:"simo,attr"  json:",omitempty"`
	Attlag      string   `xml:"lag,attr"  json:",omitempty"`
	Body        string   `xml:",chardata" json:",omitempty"`
}

type Fifo struct {
	XMLName        xml.Name  `xml:"fifo,omitempty" json:"fifo,omitempty"`
	Attrimportance string    `xml:"importance,attr"  json:",omitempty"`
	Attrname       string    `xml:"name,attr"  json:",omitempty"`
	Member         []*Member `xml:"member,omitempty"  json:"member,omitempty"`
}

type Fifos struct {
	XMLName xml.Name `xml:"fifos,omitempty" json:"fifos,omitempty"`
	Fifo    []*Fifo  `xml:"fifo,omitempty" json:"fifo,omitempty"`
}

type Gateway struct {
	XMLName   xml.Name   `xml:"gateway,omitempty" json:"gateway,omitempty"`
	Attrname  string     `xml:"name,attr"  json:",omitempty"`
	Param     []*Param   `xml:"param,omitempty" json:"param,omitempty"`
	Params    *Params    `xml:"params,omitempty" json:"params,omitempty"`
	Variables *Variables `xml:"variables,omitempty" json:"variables,omitempty"`
}

type Gateways struct {
	XMLName xml.Name   `xml:"gateways,omitempty" json:"gateways,omitempty"`
	Gateway []*Gateway `xml:"gateway,omitempty" json:"gateway,omitempty"`
}

type GlobalSettings struct {
	XMLName xml.Name `xml:"global_settings,omitempty" json:"global_settings,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Group struct {
	XMLName  xml.Name   `xml:"group,omitempty" json:"group,omitempty"`
	Attrname string     `xml:"name,attr"  json:",omitempty"`
	Control  []*Control `xml:"control,omitempty" json:"control,omitempty"`
	Layout   []*Layout  `xml:"layout,omitempty" json:"layout,omitempty"`
	Users    *Users     `xml:"users,omitempty" json:"users,omitempty"`
}

type Groups struct {
	XMLName xml.Name `xml:"groups,omitempty" json:"groups,omitempty"`
	Group   []*Group `xml:"group,omitempty" json:"group,omitempty"`
}

type Header struct {
	XMLName  xml.Name `xml:"header,omitempty" json:"header,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
}

type Image struct {
	XMLName            xml.Name `xml:"image,omitempty" json:"image,omitempty"`
	Attrfloor          string   `xml:"floor,attr"  json:",omitempty"`
	AttrfloorOnly      string   `xml:"floor-only,attr"  json:",omitempty"`
	Attrhscale         string   `xml:"hscale,attr"  json:",omitempty"`
	Attroverlap        string   `xml:"overlap,attr"  json:",omitempty"`
	Attrreservation_id string   `xml:"reservation_id,attr"  json:",omitempty"`
	Attrscale          string   `xml:"scale,attr"  json:",omitempty"`
	Attrx              string   `xml:"x,attr"  json:",omitempty"`
	Attry              string   `xml:"y,attr"  json:",omitempty"`
	Attrzoom           string   `xml:"zoom,attr"  json:",omitempty"`
}

type Input struct {
	XMLName            xml.Name `xml:"input,omitempty" json:"input,omitempty"`
	Attrbreak_on_match string   `xml:"break_on_match,attr"  json:",omitempty"`
	Attrfield          string   `xml:"field,attr"  json:",omitempty"`
	Attrpattern        string   `xml:"pattern,attr"  json:",omitempty"`
	Match              *Match   `xml:"match,omitempty" json:"match,omitempty"`
	Nomatch            *Nomatch `xml:"nomatch,omitempty" json:"nomatch,omitempty"`
}

type Javavm struct {
	XMLName  xml.Name `xml:"javavm,omitempty" json:"javavm,omitempty"`
	Attrpath string   `xml:"path,attr"  json:",omitempty"`
}

type Key struct {
	XMLName      xml.Name `xml:"key,omitempty" json:"key,omitempty"`
	Attraction   string   `xml:"action,attr"  json:",omitempty"`
	Attrdtmf     string   `xml:"dtmf,attr"  json:",omitempty"`
	Attrname     string   `xml:"name,attr"  json:",omitempty"`
	Attrvalue    string   `xml:"value,attr"  json:",omitempty"`
	Attrvariable string   `xml:"variable,attr"  json:",omitempty"`
}

type Keys struct {
	XMLName xml.Name `xml:"keys,omitempty" json:"keys,omitempty"`
	Key     []*Key   `xml:"key,omitempty" json:"key,omitempty"`
}

type Language struct {
	XMLName         xml.Name `xml:"language,omitempty" json:"language,omitempty"`
	Attrname        string   `xml:"name,attr"  json:",omitempty"`
	AttrsayModule   string   `xml:"say-module,attr"  json:",omitempty"`
	AttrsoundPath   string   `xml:"sound-path,attr"  json:",omitempty"`
	AttrsoundPrefix string   `xml:"sound-prefix,attr"  json:",omitempty"`
	AttrttsEngine   string   `xml:"tts-engine,attr"  json:",omitempty"`
	AttrttsVoice    string   `xml:"tts-voice,attr"  json:",omitempty"`
	Phrases         *Phrases `xml:"phrases,omitempty" json:"phrases,omitempty"`
}

type Layout struct {
	XMLName            xml.Name `xml:"layout,omitempty" json:"layout,omitempty"`
	Attrauto3dPosition string   `xml:"auto-3d-position,attr"  json:",omitempty"`
	Attrname           string   `xml:"name,attr"  json:",omitempty"`
	Image              []*Image `xml:"image,omitempty" json:"image,omitempty"`
	Body               string   `xml:",chardata" json:",omitempty"`
}

type LayoutSettings struct {
	XMLName xml.Name `xml:"layout-settings,omitempty" json:"layout-settings,omitempty"`
	Groups  *Groups  `xml:"groups,omitempty" json:"groups,omitempty"`
	Layouts *Layouts `xml:"layouts,omitempty" json:"layouts,omitempty"`
}

type Layouts struct {
	XMLName xml.Name  `xml:"layouts,omitempty" json:"layouts,omitempty"`
	Layout  []*Layout `xml:"layout,omitempty" json:"layout,omitempty"`
}

type List struct {
	XMLName     xml.Name `xml:"list,omitempty" json:"list,omitempty"`
	Attrdefault string   `xml:"default,attr"  json:",omitempty"`
	Attrname    string   `xml:"name,attr"  json:",omitempty"`
	Node        []*Node  `xml:"node,omitempty" json:"node,omitempty"`
}

type Listener struct {
	XMLName  xml.Name `xml:"listener,omitempty" json:"listener,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Listeners struct {
	XMLName  xml.Name    `xml:"listeners,omitempty" json:"listeners,omitempty"`
	Listener []*Listener `xml:"listener,omitempty" json:"listener,omitempty"`
}

type Lists struct {
	XMLName xml.Name `xml:"lists,omitempty" json:"lists,omitempty"`
	List    []*List  `xml:"list,omitempty" json:"list,omitempty"`
}

type Load struct {
	XMLName    xml.Name `xml:"load,omitempty" json:"load,omitempty"`
	Attrmodule string   `xml:"module,attr"  json:",omitempty"`
}

type Logging struct {
	XMLName xml.Name   `xml:"logging,omitempty" json:"logging,omitempty"`
	Profile []*Profile `xml:"profile,omitempty" json:"profile,omitempty"`
}

type Macro struct {
	XMLName   xml.Name `xml:"macro,omitempty" json:"macro,omitempty"`
	Attrname  string   `xml:"name,attr"  json:",omitempty"`
	Attrpause string   `xml:"pause,attr"  json:",omitempty"`
	Input     []*Input `xml:"input,omitempty" json:"input,omitempty"`
}

type Macros struct {
	XMLName  xml.Name `xml:"macros,omitempty" json:"macros,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Macro    []*Macro `xml:"macro,omitempty" json:"macro,omitempty"`
}

type Map struct {
	XMLName   xml.Name `xml:"map,omitempty" json:"map,omitempty"`
	Attrname  string   `xml:"name,attr"  json:",omitempty"`
	Attrvalue string   `xml:"value,attr"  json:",omitempty"`
}

type Mappings struct {
	XMLName xml.Name `xml:"mappings,omitempty" json:"mappings,omitempty"`
	Map     *Map     `xml:"map,omitempty" json:"map,omitempty"`
}

type Match struct {
	XMLName xml.Name  `xml:"match,omitempty" json:"match,omitempty"`
	Action  []*Action `xml:"action,omitempty" json:"action,omitempty"`
}

type Menu struct {
	XMLName               xml.Name  `xml:"menu,omitempty" json:"menu,omitempty"`
	AttrconfirmAttempts   string    `xml:"confirm-attempts,attr"  json:",omitempty"`
	AttrconfirmKey        string    `xml:"confirm-key,attr"  json:",omitempty"`
	AttrconfirmMacro      string    `xml:"confirm-macro,attr"  json:",omitempty"`
	AttrdigitLen          string    `xml:"digit-len,attr"  json:",omitempty"`
	AttrexitSound         string    `xml:"exit-sound,attr"  json:",omitempty"`
	AttrgreetLong         string    `xml:"greet-long,attr"  json:",omitempty"`
	AttrgreetShort        string    `xml:"greet-short,attr"  json:",omitempty"`
	AttrinterDigitTimeout string    `xml:"inter-digit-timeout,attr"  json:",omitempty"`
	AttrinvalidSound      string    `xml:"invalid-sound,attr"  json:",omitempty"`
	AttrmaxFailures       string    `xml:"max-failures,attr"  json:",omitempty"`
	AttrmaxTimeouts       string    `xml:"max-timeouts,attr"  json:",omitempty"`
	Attrname              string    `xml:"name,attr"  json:",omitempty"`
	Attrtimeout           string    `xml:"timeout,attr"  json:",omitempty"`
	AttrttsEngine         string    `xml:"tts-engine,attr"  json:",omitempty"`
	AttrttsVoice          string    `xml:"tts-voice,attr"  json:",omitempty"`
	Entry                 []*Entry  `xml:"entry,omitempty" json:"entry,omitempty"`
	Keys                  *Keys     `xml:"keys,omitempty" json:"keys,omitempty"`
	Phrases               *Phrases  `xml:"phrases,omitempty" json:"phrases,omitempty"`
	Settings              *Settings `xml:"settings,omitempty" json:"settings,omitempty"`
}

type Menus struct {
	XMLName xml.Name `xml:"menus,omitempty" json:"menus,omitempty"`
	Menu    []*Menu  `xml:"menu,omitempty" json:"menu,omitempty"`
}

type ModemSettings struct {
	XMLName xml.Name `xml:"modem-settings,omitempty" json:"modem-settings,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Modules struct {
	XMLName xml.Name `xml:"modules,omitempty" json:"modules,omitempty"`
	Load    []*Load  `xml:"load,omitempty" json:"load,omitempty"`
}

type NetworkLists struct {
	XMLName xml.Name `xml:"network-lists,omitempty" json:"network-lists,omitempty"`
	List    []*List  `xml:"list,omitempty" json:"list,omitempty"`
}

type Node struct {
	XMLName    xml.Name `xml:"node,omitempty" json:"node,omitempty"`
	Attrcidr   string   `xml:"cidr,attr"  json:",omitempty"`
	Attrdomain string   `xml:"domain,attr"  json:",omitempty"`
	Attrname   string   `xml:"name,attr"  json:",omitempty"`
	Attrtype   string   `xml:"type,attr"  json:",omitempty"`
	Attrhost   string   `xml:"host,attr"  json:",omitempty"`
	Attrmask   string   `xml:"mask,attr"  json:",omitempty"`
	Attrweight string   `xml:"weight,attr"  json:",omitempty"`
}

type Nomatch struct {
	XMLName xml.Name  `xml:"nomatch,omitempty" json:"nomatch,omitempty"`
	Action  []*Action `xml:"action,omitempty" json:"action,omitempty"`
}

type Option struct {
	XMLName   xml.Name `xml:"option,omitempty" json:"option,omitempty"`
	Attrvalue string   `xml:"value,attr"  json:",omitempty"`
}

type Options struct {
	XMLName xml.Name  `xml:"options,omitempty" json:"options,omitempty"`
	Option  []*Option `xml:"option,omitempty" json:"option,omitempty"`
}

type Param struct {
	XMLName      xml.Name `xml:"param,omitempty" json:"param,omitempty"`
	Attrbindings string   `xml:"bindings,attr"  json:",omitempty"`
	Attrmodule   string   `xml:"module,attr"  json:",omitempty"`
	Attrname     string   `xml:"name,attr"  json:",omitempty"`
	Attrprofile  string   `xml:"profile,attr"  json:",omitempty"`
	Attrsecure   string   `xml:"secure,attr"  json:",omitempty"`
	Attrvalue    string   `xml:"value,attr"  json:",omitempty"`
}

type Params struct {
	XMLName xml.Name `xml:"params,omitempty" json:"params,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Permission struct {
	XMLName         xml.Name         `xml:"permission,omitempty" json:"permission,omitempty"`
	Attrname        string           `xml:"name,attr"  json:",omitempty"`
	Attrvalue       string           `xml:"value,attr"  json:",omitempty"`
	ApplicationList *ApplicationList `xml:"application-list,omitempty" json:"application-list,omitempty"`
	VariableList    *VariableList    `xml:"variable-list,omitempty" json:"variable-list,omitempty"`
	ApiList         *ApiList         `xml:"api-list,omitempty" json:"api-list,omitempty"`
}

type Permissions struct {
	XMLName    xml.Name      `xml:"permissions,omitempty" json:"permissions,omitempty"`
	Permission []*Permission `xml:"permission,omitempty" json:"permission,omitempty"`
}

type Phrase struct {
	XMLName   xml.Name `xml:"phrase,omitempty" json:"phrase,omitempty"`
	Attrname  string   `xml:"name,attr"  json:",omitempty"`
	Attrvalue string   `xml:"value,attr"  json:",omitempty"`
}

type Phrases struct {
	XMLName xml.Name  `xml:"phrases,omitempty" json:"phrases,omitempty"`
	Macros  []*Macros `xml:"macros,omitempty" json:"macros,omitempty"`
	Phrase  []*Phrase `xml:"phrase,omitempty" json:"phrase,omitempty"`
}

type Producers struct {
	XMLName xml.Name   `xml:"producers,omitempty" json:"producers,omitempty"`
	Profile []*Profile `xml:"profile,omitempty" json:"profile,omitempty"`
}

type Profile struct {
	XMLName        xml.Name        `xml:"profile,omitempty" json:"profile,omitempty"`
	Attrname       string          `xml:"name,attr"  json:",omitempty"`
	Attrversion    string          `xml:"version,attr"  json:",omitempty"`
	Aliases        *Aliases        `xml:"aliases,omitempty" json:"aliases,omitempty"`
	Apis           *Apis           `xml:"apis,omitempty" json:"apis,omitempty"`
	Conference     *Conference     `xml:"conference,omitempty" json:"conference,omitempty"`
	Connections    *Connections    `xml:"connections,omitempty" json:"connections,omitempty"`
	DeviceTypes    *DeviceTypes    `xml:"device-types,omitempty" json:"device-types,omitempty"`
	Dial           *Dial           `xml:"dial,omitempty" json:"dial,omitempty"`
	Domains        *Domains        `xml:"domains,omitempty" json:"domains,omitempty"`
	Email          *Email          `xml:"email,omitempty" json:"email,omitempty"`
	Gateways       *Gateways       `xml:"gateways,omitempty" json:"gateways,omitempty"`
	Mappings       *Mappings       `xml:"mappings,omitempty" json:"mappings,omitempty"`
	Menus          *Menus          `xml:"menus,omitempty" json:"menus,omitempty"`
	Param          []*Param        `xml:"param,omitempty" json:"param,omitempty"`
	Params         *Params         `xml:"params,omitempty" json:"params,omitempty"`
	Permissions    *Permissions    `xml:"permissions,omitempty" json:"permissions,omitempty"`
	Recogparams    *Recogparams    `xml:"recogparams,omitempty" json:"recogparams,omitempty"`
	Rule           []*Rule         `xml:"rule,omitempty" json:"rule,omitempty"`
	Settings       *Settings       `xml:"settings,omitempty" json:"settings,omitempty"`
	SoftKeySetSets *SoftKeySetSets `xml:"soft-key-set-sets,omitempty" json:"soft-key-set-sets,omitempty"`
	Synthparams    *Synthparams    `xml:"synthparams,omitempty" json:"synthparams,omitempty"`
	User           []*User         `xml:"user,omitempty" json:"user,omitempty"`
}

type Profiles struct {
	XMLName xml.Name   `xml:"profiles,omitempty" json:"profiles,omitempty"`
	Profile []*Profile `xml:"profile,omitempty" json:"profile,omitempty"`
}

type Queue struct {
	XMLName  xml.Name `xml:"queue,omitempty" json:"queue,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Queues struct {
	XMLName xml.Name `xml:"queues,omitempty" json:"queues,omitempty"`
	Queue   []*Queue `xml:"queue,omitempty" json:"queue,omitempty"`
}

type Recogparams struct {
	XMLName xml.Name `xml:"recogparams,omitempty" json:"recogparams,omitempty"`
}

type Remotes struct {
	XMLName xml.Name `xml:"remotes,omitempty" json:"remotes,omitempty"`
}

type Room struct {
	XMLName    xml.Name `xml:"room,omitempty" json:"room,omitempty"`
	Attrname   string   `xml:"name,attr"  json:",omitempty"`
	Attrstatus string   `xml:"status,attr"  json:",omitempty"`
}

type Route struct {
	XMLName     xml.Name `xml:"route,omitempty" json:"route,omitempty"`
	Attrregex   string   `xml:"regex,attr"  json:",omitempty"`
	Attrreplace string   `xml:"replace,attr"  json:",omitempty"`
	Attrservice string   `xml:"service,attr"  json:",omitempty"`
}

type Routes struct {
	XMLName xml.Name `xml:"routes,omitempty" json:"routes,omitempty"`
	Route   []*Route `xml:"route,omitempty" json:"route,omitempty"`
}

type Rule struct {
	XMLName     xml.Name `xml:"rule,omitempty" json:"rule,omitempty"`
	Attrregex   string   `xml:"regex,attr"  json:",omitempty"`
	Attrreplace string   `xml:"replace,attr"  json:",omitempty"`
}

type Schema struct {
	XMLName xml.Name `xml:"schema,omitempty" json:"schema,omitempty"`
	Field   []*Field `xml:"field,omitempty" json:"field,omitempty"`
}

type Section struct {
	XMLName         xml.Name            `xml:"section,omitempty" json:"section,omitempty"`
	Attrdescription string              `xml:"description,attr"  json:",omitempty"`
	Attrname        string              `xml:"name,attr"  json:",omitempty"`
	X_NOPRE_PROCESS []*X_NO_PRE_PROCESS `xml:"X-NO-PRE-PROCESS,omitempty" json:"X-NO-PRE-PROCESS,omitempty"`
	Configuration   []*Configuration    `xml:"configuration,omitempty" json:"configuration,omitempty"`
	Context         []*Context          `xml:"context,omitempty" json:"context,omitempty"`
	Domain          []*Domain           `xml:"domain,omitempty" json:"domain,omitempty"`
	Language        []*Language         `xml:"language,omitempty" json:"language,omitempty"`
}

type Settings struct {
	XMLName xml.Name `xml:"settings,omitempty" json:"settings,omitempty"`
	Fields  *Fields  `xml:"fields,omitempty" json:"fields,omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Tables struct {
	XMLName xml.Name `xml:"tables,omitempty" json:"tables,omitempty"`
	Table   []*Table `xml:"table,omitempty" json:"table,omitempty"`
}

type Table struct {
	XMLName    xml.Name `xml:"table,omitempty" json:"table,omitempty"`
	Attrname   string   `xml:"name,attr"  json:",omitempty"`
	AttrlogLeg string   `xml:"log-leg,attr"  json:",omitempty"`
	Fields     []*Field `xml:"field,omitempty" json:"field,omitempty"`
}

type Skinny struct {
	XMLName xml.Name `xml:"skinny,omitempty" json:"skinny,omitempty"`
	Buttons *Buttons `xml:"buttons,omitempty" json:"buttons,omitempty"`
}

type SoftKeySet struct {
	XMLName   xml.Name `xml:"soft-key-set,omitempty" json:"soft-key-set,omitempty"`
	Attrname  string   `xml:"name,attr"  json:",omitempty"`
	Attrvalue string   `xml:"value,attr"  json:",omitempty"`
}

type SoftKeySetSet struct {
	XMLName    xml.Name      `xml:"soft-key-set-set,omitempty" json:"soft-key-set-set,omitempty"`
	Attrname   string        `xml:"name,attr"  json:",omitempty"`
	SoftKeySet []*SoftKeySet `xml:"soft-key-set,omitempty" json:"soft-key-set,omitempty"`
}

type SoftKeySetSets struct {
	XMLName       xml.Name       `xml:"soft-key-set-sets,omitempty" json:"soft-key-set-sets,omitempty"`
	SoftKeySetSet *SoftKeySetSet `xml:"soft-key-set-set,omitempty" json:"soft-key-set-set,omitempty"`
}

type Span struct {
	XMLName xml.Name `xml:"span,omitempty" json:"span,omitempty"`
	Attrid  string   `xml:"id,attr"  json:",omitempty"`
	Param   []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Spans struct {
	XMLName xml.Name `xml:"spans,omitempty" json:"spans,omitempty"`
	Span    []*Span  `xml:"span,omitempty" json:"span,omitempty"`
}

type Startup struct {
	XMLName    xml.Name `xml:"startup,omitempty" json:"startup,omitempty"`
	Attrclass  string   `xml:"class,attr"  json:",omitempty"`
	Attrmethod string   `xml:"method,attr"  json:",omitempty"`
}

type Stream struct {
	XMLName  xml.Name `xml:"stream,omitempty" json:"stream,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Streams struct {
	XMLName xml.Name  `xml:"streams,omitempty" json:"streams,omitempty"`
	Stream  []*Stream `xml:"stream,omitempty" json:"stream,omitempty"`
}

type Synthparams struct {
	XMLName xml.Name `xml:"synthparams,omitempty" json:"synthparams,omitempty"`
}

type Template struct {
	XMLName  xml.Name `xml:"template,omitempty" json:"template,omitempty"`
	Attrname string   `xml:"name,attr"  json:",omitempty"`
	string   string   `xml:",chardata" json:",omitempty"`
}

type Templates struct {
	XMLName  xml.Name    `xml:"templates,omitempty" json:"templates,omitempty"`
	Template []*Template `xml:"template,omitempty" json:"template,omitempty"`
}

type Tier struct {
	XMLName      xml.Name `xml:"tier,omitempty" json:"tier,omitempty"`
	Attragent    string   `xml:"agent,attr"  json:",omitempty"`
	Attrlevel    string   `xml:"level,attr"  json:",omitempty"`
	Attrposition string   `xml:"position,attr"  json:",omitempty"`
	Attrqueue    string   `xml:"queue,attr"  json:",omitempty"`
}

type Tiers struct {
	XMLName xml.Name `xml:"tiers,omitempty" json:"tiers,omitempty"`
	Tier    []*Tier  `xml:"tier,omitempty" json:"tier,omitempty"`
}

type Timezones struct {
	XMLName xml.Name `xml:"timezones,omitempty" json:"timezones,omitempty"`
	Zone    []*Zone  `xml:"zone,omitempty" json:"zone,omitempty"`
}

type Tone struct {
	XMLName         xml.Name   `xml:"tone,omitempty" json:"tone,omitempty"`
	Attrdescription string     `xml:"description,attr"  json:",omitempty"`
	Attrname        string     `xml:"name,attr"  json:",omitempty"`
	Element         []*Element `xml:"element,omitempty" json:"element,omitempty"`
}

type User struct {
	XMLName         xml.Name   `xml:"user,omitempty" json:"user,omitempty"`
	Attrcidr        string     `xml:"cidr,attr"  json:",omitempty"`
	Attrnumberalias string     `xml:"number-alias,attr"  json:",omitempty"`
	Attrid          string     `xml:"id,attr"  json:",omitempty"`
	Gateways        *Gateways  `xml:"gateways,omitempty" json:"gateways,omitempty"`
	Params          *Params    `xml:"params,omitempty" json:"params,omitempty"`
	Skinny          *Skinny    `xml:"skinny,omitempty" json:"skinny,omitempty"`
	Variables       *Variables `xml:"variables,omitempty" json:"variables,omitempty"`
	Vcard           *Vcard     `xml:"vcard,omitempty" json:"vcard,omitempty"`
	Attrtype        string     `xml:"type,attr" json:"type,omitempty"`
	Attrcommands    string     `xml:"commands,attr" json:"commands,omitempty"`
	Attrname        string     `xml:"name,attr" json:"name,omitempty"`
}

type Users struct {
	XMLName xml.Name `xml:"users,omitempty" json:"users,omitempty"`
	User    []*User  `xml:"user,omitempty" json:"user,omitempty"`
}

type Variable struct {
	XMLName       xml.Name `xml:"variable,omitempty" json:"variable,omitempty"`
	Attrname      string   `xml:"name,attr"  json:",omitempty"`
	Attrvalue     string   `xml:"value,attr"  json:",omitempty"`
	Attrdirection string   `xml:"direction,attr"  json:",omitempty"`
}

type Variables struct {
	XMLName  xml.Name    `xml:"variables,omitempty" json:"variables,omitempty"`
	Variable []*Variable `xml:"variable,omitempty" json:"variable,omitempty"`
}

type Vcard struct {
	XMLName xml.Name `xml:"vcard,omitempty" json:"vcard,omitempty"`
}

type XProfile struct {
	XMLName  xml.Name `xml:"x-profile,omitempty" json:"x-profile,omitempty"`
	Attrtype string   `xml:"type,attr"  json:",omitempty"`
	Param    []*Param `xml:"param,omitempty" json:"param,omitempty"`
}

type Zone struct {
	XMLName   xml.Name `xml:"zone,omitempty" json:"zone,omitempty"`
	Attrname  string   `xml:"name,attr"  json:",omitempty"`
	Attrvalue string   `xml:"value,attr"  json:",omitempty"`
}
