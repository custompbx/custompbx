package cfg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	Route    string `json:"route"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	StunPort int    `json:"stun_port"`
	CertPath string `json:"cert_path"`
	KeyPath  string `json:"key_path"`
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
	workDir, err := os.Getwd()
	if err != nil || workDir == "" {
		err = fmt.Errorf("Couldn't not detect pwd directory: " + err.Error())
		os.Exit(1)
	}

	file, err := os.ReadFile(workDir + "/" + configFile)
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
	workDir, err := os.Getwd()
	if err != nil || workDir == "" {
		err = fmt.Errorf("Couldn't not detect pwd directory: " + err.Error())
		os.Exit(1)
	}

	file, _ := json.MarshalIndent(conf, "", "    ")
	err = os.WriteFile(workDir+"/"+configFile, file, 0644)
	if err != nil {
		err = fmt.Errorf("Couldn't write to file: " + err.Error())
		return config, err
	}
	return conf, err
}

func init() {
	var err error
	workDir, err := os.Getwd()
	if err != nil || workDir == "" {
		err = fmt.Errorf("Couldn't not detect pwd directory: " + err.Error())
		os.Exit(1)
	}
	log.Println("Config file at " + workDir + "/" + configFile)
	CustomPbx, err = RD()
	if err != nil {
		fmt.Printf("ERROR: Could not read configuration file: "+configFile+", cause: %v\n", err.Error())
		os.Exit(1)
	}

	if CustomPbx.XMLCurl.Route == "" || CustomPbx.XMLCurl.Route[:1] != "/" {
		CustomPbx.XMLCurl.Route = "/" + CustomPbx.XMLCurl.Route
	}

	if CustomPbx.Web.Route == "" || CustomPbx.Web.Route[:1] != "/" {
		CustomPbx.Web.Route = "/" + CustomPbx.Web.Route
	}
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
	item.Fs.Esl.Pass = "ClueCon"
	item.Fs.Esl.Host = "127.0.0.1"
	item.Fs.Esl.Timeout = 10
	item.Fs.Esl.CollectLogs = 7
	item.Web.Route = "/ws"
	item.Web.Host = "127.0.0.1"
	item.Web.Port = 8080
	item.Web.StunPort = 3478
	item.Web.CertPath = ""
	item.Web.KeyPath = ""
	item.XMLCurl.Route = "/conf/config"
	item.XMLCurl.Host = "127.0.0.1"
	item.XMLCurl.Port = 8081
	item.XMLCurl.CertPath = ""
	item.XMLCurl.KeyPath = ""

	return item
}
