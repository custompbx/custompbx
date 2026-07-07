package hepHandler

import (
	"custompbx/cfg"
	"fmt"
	"github.com/custompbx/hepparser"
	"log"
	"net"
	"time"
)

const maxPktLen = 8192

func StartHepListener(DBSaveHandler func(hep []*hepparser.HEP, instanceId int64) error, instanceId int64) {
	if cfg.CustomPbx.Fs.HEPCollector.Host == "" || cfg.CustomPbx.Fs.HEPCollector.Port < 1024 || cfg.CustomPbx.Fs.HEPCollector.Port > 65535 {
		log.Println("HEP collecting disabled")
		return
	}

	log.Println("HEP collecting")
	ua, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", cfg.CustomPbx.Fs.HEPCollector.Host, cfg.CustomPbx.Fs.HEPCollector.Port))
	if err != nil {
		log.Printf("%v", err)
		return
	}

	uc, err := net.ListenUDP("udp", ua)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in StartHepListener", r)
		}
		log.Printf("stopping UDP listener on %s", uc.LocalAddr())
		uc.Close()
	}()

	HEPChannel := make(chan *hepparser.HEP, 40000)
	go HEPSaver(DBSaveHandler, HEPChannel, instanceId)

	for {
		//uc.SetReadDeadline(time.Now().Add(1e9))
		buf := make([]byte, maxPktLen)
		n, err := uc.Read(buf)
		if err != nil {
			//if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
			//	log.Printf("%v\n continue", err)
			//	continue
			//} else {
			log.Printf("%v\n return", err)
			return
			//}
		} else if n > maxPktLen {
			log.Printf("received too big packet with %d bytes", n)
			continue
		}
		go ParseAndSaveHEP(buf[:n], HEPChannel)
	}
}

func HEPSaver(SaveHEPPackets func(hep []*hepparser.HEP, instanceId int64) error, HEPChannel chan *hepparser.HEP, instanceId int64) {
	chuck := 50
	collector := make([]*hepparser.HEP, 0, chuck)
	tick := time.Tick(300 * time.Millisecond)

	for {
		select {
		case hep := <-HEPChannel:
			collector = append(collector, hep)
			if len(collector) < chuck {
				continue
			}
			go SaveHEPPackets(collector, instanceId)
			collector = []*hepparser.HEP{}
		case <-tick:
			if len(collector) == 0 {
				continue
			}
			go SaveHEPPackets(collector, instanceId)
			collector = []*hepparser.HEP{}
		}
	}
}

func ParseAndSaveHEP(packet []byte, HEPChannel chan *hepparser.HEP) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in ParseAndSaveHEP", r)
		}
	}()

	hepPacket, err := hepparser.DecodeHEP(packet)
	if err != nil {
		log.Println(err)
		return
	}
	HEPChannel <- hepPacket
}
