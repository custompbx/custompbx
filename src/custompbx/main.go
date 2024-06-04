package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
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
	"encoding/pem"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/pion/turn/v2"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"math/big"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	eventChannel := make(chan interface{}, 42)
	logsChannel := make(chan mainStruct.LogType, 420)

	log.Println("CustomPBX development version: " + mainStruct.Version)
	log.Println("Starting...")
	daemonCache.InitDaemonState()
	log.Println("DB")
	db.StartDB()
	log.Println("Database connected.")

	pbxcache.InitCacheObjects()
	webcache.InitCacheObjects()

	log.Println("Events Handler")
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
		db.InitWebDB(cache.GetCurrentInstanceId())
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

		db.InitHEPDb()
		go func() { hepHandler.StartHepListener(db.SaveHEPPackets, cache.GetCurrentInstanceId()) }()
	} else {
		log.Printf("no service connections DB: %t ESL: %t\n", daemonCache.State.DatabaseConnection, daemonCache.State.ESLConnection)
	}

	log.Println("Web Handlers")
	rCurl := chi.NewRouter()
	configureMiddleware(rCurl)
	rCurl.Post(cfg.CustomPbx.XMLCurl.Route, func(w http.ResponseWriter, r *http.Request) { dispatcher(w, r, eventChannel) })

	rWeb := chi.NewRouter()
	configureMiddleware(rWeb)
	rWeb.Get(cfg.CustomPbx.Web.Route, web.StartWS)
	rWeb.Post("/api/v1", web.PostAPIRequest)
	configureStaticRoutes(rWeb)

	go turnServer()
	startServers(rWeb, rCurl)
}

func startServers(r chi.Router, cr chi.Router) {
	curlCert, curlKey, webCert, webKey := checkAndCreateCerts()

	if cfg.CustomPbx.XMLCurl.Secure {
		log.Println("Starting XMLCurl Secure Server...")
		go func() {
			log.Fatal(http.ListenAndServeTLS(cfg.CustomPbx.XMLCurl.Host+":"+strconv.Itoa(cfg.CustomPbx.XMLCurl.Port), curlCert, curlKey, cr))
		}()
	} else {
		log.Println("Starting XMLCurl Insecure Server...")
		go func() {
			log.Fatal(http.ListenAndServe(cfg.CustomPbx.XMLCurl.Host+":"+strconv.Itoa(cfg.CustomPbx.XMLCurl.Port), cr))
		}()
	}

	if cfg.CustomPbx.Web.Secure {
		log.Println("Starting Web Secure Server...")
		log.Fatal(http.ListenAndServeTLS(cfg.CustomPbx.Web.Host+":"+strconv.Itoa(cfg.CustomPbx.Web.Port), webCert, webKey, r))
	} else {
		log.Println("Starting Web Insecure Server...")
		log.Fatal(http.ListenAndServe(cfg.CustomPbx.Web.Host+":"+strconv.Itoa(cfg.CustomPbx.Web.Port), r))
	}
}

func configureMiddleware(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))
}

func configureStaticRoutes(r chi.Router) {
	// Define other routes here
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

func checkAndCreateCerts() (string, string, string, string) {
	var curlCert = cfg.CustomPbx.XMLCurl.CertPath
	var curlKey = cfg.CustomPbx.XMLCurl.KeyPath
	var webCert = cfg.CustomPbx.Web.CertPath
	var webKey = cfg.CustomPbx.Web.KeyPath

	var err error
	log.Println("Checking certs")
	curlCert, curlKey, err = loadCertificateAndKey(curlCert, curlKey)
	if err != nil {
		log.Println("Certs Error: ", err.Error())
		curlCert = "./cert.pem"
		curlKey = curlCert
		_, err := os.ReadFile(curlCert)
		if err != nil {
			createCert(curlCert)
		}
	}

	if cfg.CustomPbx.XMLCurl.CertPath == cfg.CustomPbx.Web.CertPath {
		return curlCert, curlKey, curlCert, curlKey
	}

	log.Println("Checking Web cert", webCert, webKey)
	webCert, webKey, err = loadCertificateAndKey(webCert, webKey)
	if err != nil {
		log.Println("Certs Error: ", err.Error())
		webCert = "./cert.pem"
		webKey = webCert
		_, err := os.ReadFile(webCert)
		if err != nil {
			createCert(webCert)
		}
	}

	return curlCert, curlKey, webCert, webKey
}

func createCert(filePath string) {
	// Generate a new private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return
	}

	// Create a template for the certificate
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Self-Signed Certificate"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // Valid for 1 year
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// Generate the certificate using the template and the private key
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		fmt.Println("Failed to create certificate:", err)
		return
	}

	// Create a buffer to hold the PEM data
	pemBuffer := []byte{}

	// Append the private key PEM block to the buffer
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	pemBuffer = append(pemBuffer, pem.EncodeToMemory(privateKeyPEM)...)

	// Append the certificate PEM block to the buffer
	certPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	pemBuffer = append(pemBuffer, pem.EncodeToMemory(certPEM)...)

	// Save the combined PEM data to a file
	combinedFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create", filePath, "PEM file:", err)
		return
	}
	defer combinedFile.Close()

	if _, err := combinedFile.Write(pemBuffer); err != nil {
		fmt.Println("Failed to create", filePath, "PEM file:", err)
		return
	}

	fmt.Println("Self-signed certificate and private key have been generated and saved in", filePath)
}

func loadCertificateAndKey(certFilePath, keyFilePath string) (string, string, error) {
	certBytes, err := os.ReadFile(certFilePath)
	if err != nil {
		return "", "", err
	}

	// Decode the PEM data and get the certificate and private key
	var cert *x509.Certificate
	//var privateKey interface{}
	var certContainsKey bool

	for {
		block, rest := pem.Decode(certBytes)
		if block == nil {
			break
		}

		certBytes = rest

		switch block.Type {
		case "CERTIFICATE":
			cert, err = x509.ParseCertificate(block.Bytes)
			if err != nil {
				return "", "", err
			}
		case "RSA PRIVATE KEY", "PRIVATE KEY":
			_, err = x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				// If PKCS#1 parsing fails, try PKCS#8 parsing
				_, err = x509.ParsePKCS8PrivateKey(block.Bytes)
				if err != nil {
					return "", "", err
				}
			}
			certContainsKey = true
		default:
			// Handle other PEM blocks if needed
		}
	}

	keyFile := certFilePath
	// If the certificate doesn't contain the key, load the key from the key file
	if !certContainsKey {
		keyFile = keyFilePath
		privateKeyBytes, err := os.ReadFile(keyFilePath)
		if err != nil {
			return "", "", err
		}

		_, err = x509.ParsePKCS8PrivateKey(privateKeyBytes)
		if err != nil {
			return "", "", err
		}
	}

	fmt.Println(
		"The cert", certFilePath, "is valid.",
		"Subject:", cert.Subject.CommonName,
		"Issuer:", cert.Issuer.CommonName,
		"Not Before:", cert.NotBefore,
		"Not After:", cert.NotAfter,
		"The key ", keyFile, "is valid.",
	)

	return certFilePath, keyFile, nil
}
