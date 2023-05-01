package main

import (
	"bytes"
	"custompbx/apps"
	"custompbx/cache"
	"custompbx/cfg"
	"custompbx/cweb"
	"custompbx/daemonCache"
	"custompbx/db"
	"custompbx/fsesl"
	"custompbx/hepHandler"
	"custompbx/mainStruct"
	"custompbx/metrics"
	"custompbx/nocache"
	"custompbx/pbxcache"
	"custompbx/web"
	"custompbx/webcache"
	"custompbx/xmlcurl"
	"encoding/base64"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pion/turn/v2"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	log.Println("CustomPBX development version: " + mainStruct.Version)
	log.Println("Starting...")
	daemonCache.InitDaemonState()
	log.Println("DB")
	db.StartDB()
	pbxcache.InitCacheObjects()
	webcache.InitCacheObjects()

	log.Println("Events Handler")
	eventChannel := make(chan interface{}, 42)
	logsChannel := make(chan mainStruct.LogType, 420)

	web.SetBroadcastChannel(eventChannel)
	go web.TimeEvents()

	fsesl.FirstConnectData()
	switchName := cache.GetCurrentInstanceName()
	if switchName != "" {
		if switchName != cfg.CustomPbx.Fs.Switchname {
			log.Println("FS actual switchname is different from config file!!!")
			return
		}
	} else {
		if cfg.CustomPbx.Fs.Switchname == "" {
			log.Println("FS actual switchname is different from config file!!!")
			return
		}
		switchName = cfg.CustomPbx.Fs.Switchname
		cache.SetCurrentInstanceName(cfg.CustomPbx.Fs.Switchname)
	}
	if switchName == "" {
		log.Println("FS switchname is empty!")
		return
	}
	log.Println("FS SWITHNAME is " + switchName)

	log.Println("Cache")
	if daemonCache.State.DatabaseConnection {
		pbxcache.InitRootDB()
		db.InitCustomDB()
		log.Println("Checking DB migration")
		cache.InitCache()
		ok, err := db.Migrate(switchName)
		if err != nil {
			log.Println("[ERROR] Database schema update met error " + err.Error() + "!")
		}
		if ok {
			log.Println("Database schema updated for version " + mainStruct.Version)
		}
		nocache.InitDB()
		apps.InitApps()
		db.InitLogDB()
		web.InitDB(cache.GetCurrentInstanceId())
		db.InitGlobalVariablesDB()

		pbxcache.InitPBXCache()
		webcache.InitUsersCache(cache.GetCurrentInstanceId())
		webcache.InitWebSettings(cache.GetCurrentInstanceId())
		webcache.InitWebData()

		log.Println("ESL Connection and handlers")
		go fsesl.ESLConnectKeeper(eventChannel, logsChannel)
		go metrics.UpdateMetrics()
		log.Println("FS logs collecting")
		go freeswitchLogHandler(logsChannel)

		log.Println("HEP collecting")
		db.InitHEPDb()
		go func() { hepHandler.StartHepListener(db.SaveHEPPackets, cache.GetCurrentInstanceId()) }()
	} else {
		log.Printf("no service connections DB: %t ESL: %t\n", daemonCache.State.DatabaseConnection, daemonCache.State.ESLConnection)
	}

	log.Println("Web Handlers")
	rCurl := chi.NewRouter()
	rCurl.Use(middleware.RequestID)
	rCurl.Use(middleware.RealIP)
	rCurl.Use(middleware.Logger)
	rCurl.Use(middleware.Recoverer)
	rCurl.Use(render.SetContentType(render.ContentTypeJSON))
	rCurl.Use(middleware.Timeout(60 * time.Second))

	if cfg.CustomPbx.XMLCurl.Route[:1] != "/" {
		cfg.CustomPbx.XMLCurl.Route = "/" + cfg.CustomPbx.XMLCurl.Route
	}
	rCurl.Post(cfg.CustomPbx.XMLCurl.Route, func(w http.ResponseWriter, r *http.Request) { dispatcher(w, r, eventChannel) })

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	if cfg.CustomPbx.Web.Route[:1] != "/" {
		cfg.CustomPbx.Web.Route = "/" + cfg.CustomPbx.Web.Route
	}
	r.Get(cfg.CustomPbx.Web.Route, web.StartWS)
	r.Post("/api/v1", web.PostAPIRequest)

	r.Route("/cweb", func(r chi.Router) {
		r.Get("/{any1}/{any2}/{any3}/{any4}/{any5}/{any6}/{any7}/{any8}/{any9}", Web)
		r.Get("/{any1}/{any2}/{any3}/{any4}/{any5}/{any6}/{any7}/{any8}", Web)
		r.Get("/{any1}/{any2}/{any3}/{any4}/{any5}/{any6}/{any7}", Web)
		r.Get("/{any1}/{any2}/{any3}/{any4}/{any5}/{any6}", Web)
		r.Get("/{any1}/{any2}/{any3}/{any4}/{any5}", Web)
		r.Get("/{any1}/{any2}/{any3}/{any4}", Web)
		r.Get("/{any1}/{any2}/{any3}", Web)
		r.Get("/{any1}/{any2}", Web)
		r.Get("/{any1}", Web)
		r.Get("/", Web)
		//r.Get("/reload/callcenter", ReloadCallcenter)

		/*	r.Get("/assets/sounds/{sounds}", Web)
			r.Get("/assets/img/{img}", Web)
			r.Get("/{file}", Web)*/
	})

	go turnServer()

	go func() {
		if cfg.CustomPbx.XMLCurl.CertPath != "" {
			log.Println("Secure XMLCurl Server")
			err := http.ListenAndServeTLS(cfg.CustomPbx.XMLCurl.Host+":"+strconv.Itoa(cfg.CustomPbx.XMLCurl.Port), cfg.CustomPbx.XMLCurl.CertPath, cfg.CustomPbx.XMLCurl.KeyPath, rCurl)
			if err != nil {
				log.Println(err)
				log.Println("Insecure XMLCurl Server")
				log.Fatal(http.ListenAndServe(cfg.CustomPbx.XMLCurl.Host+":"+strconv.Itoa(cfg.CustomPbx.XMLCurl.Port), rCurl))
			}
		} else {
			log.Println("Insecure XMLCurl Server")
			log.Fatal(http.ListenAndServe(cfg.CustomPbx.XMLCurl.Host+":"+strconv.Itoa(cfg.CustomPbx.XMLCurl.Port), rCurl))
		}
	}()

	if cfg.CustomPbx.Web.CertPath != "" {
		log.Println("Secure Web Server")
		err := http.ListenAndServeTLS(cfg.CustomPbx.Web.Host+":"+strconv.Itoa(cfg.CustomPbx.Web.Port), cfg.CustomPbx.Web.CertPath, cfg.CustomPbx.Web.KeyPath, r)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Insecure Web Server")
	log.Fatal(http.ListenAndServe(cfg.CustomPbx.Web.Host+":"+strconv.Itoa(cfg.CustomPbx.Web.Port), r))
}

func dispatcher(w http.ResponseWriter, r *http.Request, events chan interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in dispatcher", r)
		}
	}()
	var response []byte

	var Header = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>` + "\n"
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		response = xmlcurl.NotFound(make(map[string]string))
		w.Header().Set("Content-Type", "application/xml")
		w.Write(append([]byte(Header), response...))
		return
	}

	form := make(map[string]string)
	for key := range r.PostForm {
		form[key] = r.PostFormValue(key)
		//log.Println(key, r.PostFormValue(key))
	}
	//log.Println()
	switch r.FormValue("section") {
	case "directory":
		response = xmlcurl.Directory(form)
	case "configuration":
		response = xmlcurl.Configuration(form)
	case "dialplan":
		response = xmlcurl.Dialplan(form, events)
	default:
		response = xmlcurl.NotFound(form)
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(append([]byte(Header), response...))
}

func Web(rw http.ResponseWriter, r *http.Request) {
	var err error
	defer r.Body.Close()

	route1 := chi.URLParam(r, "any1")
	route2 := chi.URLParam(r, "any2")
	route3 := chi.URLParam(r, "any3")
	route4 := chi.URLParam(r, "any4")
	route5 := chi.URLParam(r, "any5")
	route6 := chi.URLParam(r, "any6")
	route7 := chi.URLParam(r, "any7")
	route8 := chi.URLParam(r, "any8")
	route9 := chi.URLParam(r, "any9")
	filePath := "index.html"
	if route1 != "" {
		filePath = route1
	}
	if route2 != "" {
		filePath = filePath + "/" + route2
	}
	if route3 != "" {
		filePath = filePath + "/" + route3
	}
	if route4 != "" {
		filePath = filePath + "/" + route4
	}

	var file []byte
	if route1 == "assets" && route2 == "img" && route3 == "avatar" {
		c, err := r.Cookie("token")
		if err != nil {
			log.Println(err.Error())
			_, _ = rw.Write([]byte("access denied"))
			return
		}
		requester, err := webcache.GetWebUserByToken(c.Value)
		if err != nil || requester == nil || requester.Login == "" {
			_, _ = rw.Write([]byte("access denied"))
			return
		}

		extension := filepath.Ext(route4)
		userId := route4[0 : len(route4)-len(extension)]
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			return
		}
		user := webcache.GetWebUserById(id)
		if user == nil {
			return
		}
		unbased, _ := base64.StdEncoding.DecodeString(user.Avatar)
		var img image.Image
		buf := new(bytes.Buffer)
		switch user.AvatarFormat {
		case "jpeg":
			img, err = jpeg.Decode(bytes.NewReader(unbased))
			if err != nil {
				return
			}
			err = jpeg.Encode(buf, img, nil)
			if err != nil {
				return
			}
			file = buf.Bytes()
		case "png":
			img, err = png.Decode(bytes.NewReader(unbased))
			if err != nil {
				return
			}
			err = png.Encode(buf, img)
			if err != nil {
				return
			}
			file = buf.Bytes()
		case "gif":
			img, err = gif.Decode(bytes.NewReader(unbased))
			if err != nil {
				return
			}
			err = gif.Encode(buf, img, nil)
			if err != nil {
				return
			}
			file = buf.Bytes()
		}
	} else if route1 == "cdr" && route2 == "records" {
		c, err := r.Cookie("token")
		if err != nil {
			log.Println(err.Error())
			_, _ = rw.Write([]byte("access denied"))
			return
		}
		requester, err := webcache.GetWebUserByToken(c.Value)
		if err != nil || requester == nil || requester.Login == "" {
			_, _ = rw.Write([]byte("access denied"))
			return
		}

		if route3 != "" {
			servePath := webcache.GetWebSetting(webcache.CdrFileServerPath)
			if servePath != "" && servePath != "/" {
				servePath += "/" + route3
				if route4 != "" {
					servePath += "/" + route4
				}
				if route5 != "" {
					servePath += "/" + route5
				}
				if route6 != "" {
					servePath += "/" + route6
				}
				if route7 != "" {
					servePath += "/" + route7
				}
				if route8 != "" {
					servePath += "/" + route8
				}
				if route9 != "" {
					servePath += "/" + route9
				}
				log.Println(servePath)

				http.ServeFile(rw, r, servePath)

				return
				/*				file, err = ioutil.ReadFile(servePath)
								if err != nil {
									file = []byte("not found")
								}
								rw.Header().Set("accept-ranges", "bytes")*/
			} else {
				file = []byte("cdr path is filesystem root")
			}
		} else {
			file = []byte("no route")
		}
	}

	path := r.URL.Path[1:]
	// set content-type and content-encoding
	rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path))+"; charset=utf-8")
	//rw.Header().Set("Content-Encoding", "gzip")

	// security headers
	// SecurityHeaders(rw, r)
	// gzip compression
	//w, _ := gzip.NewWriterLevel(rw, 2)
	//defer w.Close()

	if file == nil {
		file, err = cweb.Asset(filePath)
		if err != nil {
			log.Println(err)
			file = cweb.MustAsset("index.html")
		}
	}

	// write file
	_, err = rw.Write(file)
	if err != nil {
		// print unknown error
		log.Println(err)
	}
}

func SecurityHeaders(w http.ResponseWriter, r *http.Request) {
	// content-security-policy
	w.Header().Set("Content-Security-Policy",
		"default-src 'self'; img-src 'self' data:; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';")

	// access-control-allow-origin
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-ConfigConferenceCallerControlsControls-Allow-Origin", origin)
	}

	// access-control-allow-headers
	w.Header().Set("Access-ConfigConferenceCallerControlsControls-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// access-control-allow-methods
	w.Header().Set("Access-ConfigConferenceCallerControlsControls-Allow-Methods", "POST")
}

func freeswitchLogHandler(logsChannel chan mainStruct.LogType) {
	chuck := 300
	collector := make([]mainStruct.LogType, 0, chuck)
	tick := time.Tick(300 * time.Millisecond)

	for {
		select {
		case logLine := <-logsChannel:
			collector = append(collector, logLine)
			if len(collector) < chuck {
				continue
			}
			go db.SetLogLines(collector, cache.GetCurrentInstanceId())
			collector = []mainStruct.LogType{}
		case <-tick:
			if len(collector) == 0 {
				continue
			}
			go db.SetLogLines(collector, cache.GetCurrentInstanceId())
			collector = []mainStruct.LogType{}
		}
	}
}

func turnServer() {
	if cfg.CustomPbx.Web.StunPort < 3000 || cfg.CustomPbx.Web.StunPort > 65765 {
		log.Println("STUN/TURN: stun port not in range 3000-65535 ")
		daemonCache.State.StunServerStatus = false
		return
	}
	udpListener, err := net.ListenPacket("udp4", "0.0.0.0:"+strconv.Itoa(cfg.CustomPbx.Web.StunPort))
	if err != nil {
		log.Println("STUN/TURN: Failed to create TURN server listener: ", err)
		return
	}

	_, err = turn.NewServer(turn.ServerConfig{
		ChannelBindTimeout: time.Second * 2,
		Realm:              "custom-pbx.com",
		AuthHandler: func(username string, realm string, srcAddr net.Addr) ([]byte, bool) {
			return nil, false
		},
		PacketConnConfigs: []turn.PacketConnConfig{
			{
				PacketConn: udpListener,
				RelayAddressGenerator: &turn.RelayAddressGeneratorStatic{
					RelayAddress: net.ParseIP(cfg.CustomPbx.Web.Host),
					Address:      "0.0.0.0",
				},
			},
		},
	})
	if err != nil {
		log.Println("STUN: ", err)
		daemonCache.State.StunServerStatus = false
		return
	}

	log.Println("STUN Server")
	daemonCache.State.StunServerStatus = true
	sigs := make(chan interface{}, 1)
	<-sigs

	daemonCache.State.StunServerStatus = false
}

func ReloadCallcenter(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	pbxcache.ReloadCallcenter()
	_, _ = rw.Write([]byte("Done"))
}
