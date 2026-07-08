package cfg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	configFile = "config.json"
	AppName    = "CustomPBX"
)

var CustomPbx GeneralCfg

type FreeSWITCH struct {
	Switchname   string       `json:"switchname"`
	Esl          Esl          `json:"esl"`
	HEPCollector HEPCollector `json:"hep_collector"`
}

type Esl struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Pass        string `json:"pass"`
	Timeout     int    `json:"timeout"`
	CollectLogs int    `json:"collect_logs"`
}

type HEPCollector struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type WebServer struct {
	Route               string   `json:"route"`
	Host                string   `json:"host"`
	Port                int      `json:"port"`
	StunPort            int      `json:"stun_port"`
	CertPath            string   `json:"cert_path"`
	KeyPath             string   `json:"key_path"`
	Secure              bool     `json:"secure"`
	OriginPolicy        string   `json:"origin_policy,omitempty"`
	AllowedOrigins      []string `json:"allowed_origins,omitempty"`
	WriteTimeoutSeconds int      `json:"ws_write_timeout_seconds,omitempty"`
	ReadTimeoutSeconds  int      `json:"ws_read_timeout_seconds,omitempty"`
	PingIntervalSeconds int      `json:"ws_ping_interval_seconds,omitempty"`
	WebSocketQueueSize  int      `json:"websocket_queue_size,omitempty"`
}

const (
	OriginPolicySameOrigin = "same_origin"
	OriginPolicyAllowList  = "allow_list"
	OriginPolicyAllowAll   = "allow_all"

	DefaultWSWriteTimeoutSeconds = 10
	DefaultWSReadTimeoutSeconds  = 60
	DefaultWSPingIntervalSeconds = 30
	DefaultWebSocketQueueSize    = 64
	MaxWebSocketQueueSize        = 1024
)

func (w *WebServer) NormalizeAndValidateOrigins() error {
	if w.WriteTimeoutSeconds <= 0 {
		w.WriteTimeoutSeconds = DefaultWSWriteTimeoutSeconds
	}
	if w.ReadTimeoutSeconds <= 1 {
		w.ReadTimeoutSeconds = DefaultWSReadTimeoutSeconds
	}
	if w.PingIntervalSeconds <= 0 {
		w.PingIntervalSeconds = DefaultWSPingIntervalSeconds
	}
	if w.PingIntervalSeconds >= w.ReadTimeoutSeconds {
		w.PingIntervalSeconds = w.ReadTimeoutSeconds / 2
		if w.PingIntervalSeconds < 1 {
			w.PingIntervalSeconds = 1
		}
	}
	w.WebSocketQueueSize = NormalizeWebSocketQueueSize(w.WebSocketQueueSize)
	if w.OriginPolicy == "" {
		w.OriginPolicy = OriginPolicySameOrigin
	}
	if w.OriginPolicy != OriginPolicySameOrigin && w.OriginPolicy != OriginPolicyAllowList && w.OriginPolicy != OriginPolicyAllowAll {
		return fmt.Errorf("invalid webserver origin_policy %q", w.OriginPolicy)
	}
	if w.OriginPolicy == OriginPolicyAllowList && len(w.AllowedOrigins) == 0 {
		return fmt.Errorf("webserver allowed_origins must not be empty when origin_policy is allow_list")
	}
	for i, origin := range w.AllowedOrigins {
		normalized, err := NormalizeOrigin(origin)
		if err != nil {
			return fmt.Errorf("invalid webserver allowed_origins[%d]: %w", i, err)
		}
		w.AllowedOrigins[i] = normalized
	}
	return nil
}

func NormalizeWebSocketQueueSize(size int) int {
	if size <= 0 {
		return DefaultWebSocketQueueSize
	}
	if size > MaxWebSocketQueueSize {
		return MaxWebSocketQueueSize
	}
	return size
}

func NormalizeOrigin(origin string) (string, error) {
	origin = strings.TrimSpace(origin)
	u, err := url.Parse(origin)
	if err != nil || u.Scheme == "" || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
		return "", fmt.Errorf("origin must contain only http(s) scheme and host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("origin scheme must be http or https")
	}
	return strings.ToLower(u.Scheme) + "://" + strings.ToLower(u.Host), nil
}

type Database struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type GeneralCfg struct {
	Fs      FreeSWITCH `json:"freeswitch"`
	Web     WebServer  `json:"webserver"`
	XMLCurl WebServer  `json:"xml_curl_server"`
	Db      Database   `json:"database"`
}

func RD() (config GeneralCfg, err error) {
	path, err := ConfigPath()
	if err != nil {
		return config, err
	}
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Config file not found. Creating...\n")
		config = createConfig()
		config, err = WD(createConfig())
		return config, err
	}
	err = json.Unmarshal(file, &config)
	return config, err
}

func WD(conf GeneralCfg) (config GeneralCfg, err error) {
	path, err := ConfigPath()
	if err != nil {
		return config, err
	}
	file, _ := json.MarshalIndent(conf, "", "    ")
	err = os.WriteFile(path, file, 0600)
	if err != nil {
		err = fmt.Errorf("could not write config file: %w", err)
		return config, err
	}
	return conf, err
}

func init() {
	var err error
	path, err := ConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config file at " + path)
	CustomPbx, err = RD()
	if err != nil {
		fmt.Printf("ERROR: Could not read configuration file: "+configFile+", cause: %v\n", err.Error())
		os.Exit(1)
	}
	if err = CustomPbx.Web.NormalizeAndValidateOrigins(); err != nil {
		log.Fatal(err)
	}

	if CustomPbx.XMLCurl.Route == "" || CustomPbx.XMLCurl.Route[:1] != "/" {
		CustomPbx.XMLCurl.Route = "/" + CustomPbx.XMLCurl.Route
	}

	if CustomPbx.Web.Route == "" || CustomPbx.Web.Route[:1] != "/" {
		CustomPbx.Web.Route = "/" + CustomPbx.Web.Route
	}
}

func ConfigPath() (string, error) {
	if configured := strings.TrimSpace(os.Getenv("CUSTOMPBX_CONFIG")); configured != "" {
		return filepath.Abs(configured)
	}
	workDir, err := os.Getwd()
	if err != nil || workDir == "" {
		return "", fmt.Errorf("could not detect working directory: %w", err)
	}
	return filepath.Join(workDir, configFile), nil
}

func createConfig() GeneralCfg {
	hostname, _ := os.Hostname()
	var item GeneralCfg
	item.Fs.Switchname = hostname
	item.Db.Host = "127.0.0.1"
	item.Db.Name = "custompbx"
	item.Db.User = "custompbx"
	item.Db.Pass = "custompbx"
	item.Db.Port = 5432
	item.Fs.HEPCollector.Host = "127.0.0.1"
	item.Fs.HEPCollector.Port = 9060
	item.Fs.Esl.Port = 8021
	item.Fs.Esl.Pass = "change-me"
	item.Fs.Esl.Host = "127.0.0.1"
	item.Fs.Esl.Timeout = 10
	item.Fs.Esl.CollectLogs = 7
	item.Web.Route = "/ws"
	item.Web.Host = "127.0.0.1"
	item.Web.Port = 8080
	item.Web.StunPort = 3478
	item.Web.CertPath = ""
	item.Web.KeyPath = ""
	item.Web.Secure = true
	item.Web.OriginPolicy = OriginPolicySameOrigin
	item.Web.WriteTimeoutSeconds = DefaultWSWriteTimeoutSeconds
	item.Web.ReadTimeoutSeconds = DefaultWSReadTimeoutSeconds
	item.Web.PingIntervalSeconds = DefaultWSPingIntervalSeconds
	item.Web.WebSocketQueueSize = DefaultWebSocketQueueSize
	item.XMLCurl.Route = "/conf/config"
	item.XMLCurl.Host = "127.0.0.1"
	item.XMLCurl.Port = 8081
	item.XMLCurl.CertPath = ""
	item.XMLCurl.KeyPath = ""
	item.XMLCurl.Secure = true

	return item
}
