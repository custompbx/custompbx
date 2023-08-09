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
	/*
		err = DBSaveHandler(hepPacket)
		if err != nil {
			log.Println(err)
			return
		}
	*/
	// brChannel <- hepPacket
	/*
		log.Printf("Version %v\n", hepPacket.Version)
		log.Printf("Protocol %v\n", hepPacket.Protocol)
		log.Printf("SrcIP %v\n", hepPacket.SrcIP)
		log.Printf("DstIP %v\n", hepPacket.DstIP)
		log.Printf("SrcPort %v\n", hepPacket.SrcPort)
		log.Printf("DstPort %v\n", hepPacket.DstPort)
		log.Printf("Tsec %v\n", hepPacket.Tsec)
		log.Printf("Tmsec %v\n", hepPacket.Tmsec)
		log.Printf("ProtoType %v\n", hepPacket.ProtoType)
		log.Printf("NodeID %v\n", hepPacket.NodeID)
		log.Printf("NodePW %v\n", hepPacket.NodePW)
		log.Printf("Payload %v\n", hepPacket.Payload)
		log.Printf("CID %v\n", hepPacket.CID)
		log.Printf("Vlan %v\n", hepPacket.Vlan)
		log.Printf("ProtoString %v\n", hepPacket.ProtoString)
		log.Printf("Timestamp %v\n", hepPacket.Timestamp)
		log.Printf("NodeName %v\n", hepPacket.NodeName)
		log.Printf("SID %v\n", hepPacket.SID)

		log.Printf("State %+v\n", hepPacket.SIP.State)
		log.Printf("Error %+v\n", hepPacket.SIP.Error)
		log.Printf("Msg %+v\n", hepPacket.SIP.Msg)
		log.Printf("CallingParty %+v\n", hepPacket.SIP.CallingParty)
		log.Printf("Body %+v\n", hepPacket.SIP.Body)
		log.Printf("Authorization %+v\n", hepPacket.SIP.Authorization)
		log.Printf("AuthVal %+v\n", hepPacket.SIP.AuthVal)
		log.Printf("AuthUser %+v\n", hepPacket.SIP.AuthUser)
		log.Printf("ContentLength %+v\n", hepPacket.SIP.ContentLength)
		log.Printf("ContentType %+v\n", hepPacket.SIP.ContentType)
		log.Printf("From %+v\n", hepPacket.SIP.From)
		log.Printf("FromUser %+v\n", hepPacket.SIP.FromUser)
		log.Printf("FromHost %+v\n", hepPacket.SIP.FromHost)
		log.Printf("FromTag %+v\n", hepPacket.SIP.FromTag)
		log.Printf("MaxForwards %+v\n", hepPacket.SIP.MaxForwards)
		log.Printf("Organization %+v\n", hepPacket.SIP.Organization)
		log.Printf("To %+v\n", hepPacket.SIP.To)
		log.Printf("ToUser %+v\n", hepPacket.SIP.ToUser)
		log.Printf("ToHost %+v\n", hepPacket.SIP.ToHost)
		log.Printf("ToTag %+v\n", hepPacket.SIP.ToTag)
		log.Printf("Contact %+v\n", hepPacket.SIP.Contact)
		log.Printf("ContactVal %+v\n", hepPacket.SIP.ContactVal)
		log.Printf("ContactUser %+v\n", hepPacket.SIP.ContactUser)
		log.Printf("ContactHost %+v\n", hepPacket.SIP.ContactHost)
		log.Printf("ContactPort %+v\n", hepPacket.SIP.ContactPort)
		log.Printf("CallID %+v\n", hepPacket.SIP.CallID)
		log.Printf("XCallID %+v\n", hepPacket.SIP.XCallID)
		log.Printf("XHeader %+v\n", hepPacket.SIP.XHeader)
		log.Printf("Cseq %+v\n", hepPacket.SIP.Cseq)
		log.Printf("CseqMethod %+v\n", hepPacket.SIP.CseqMethod)
		log.Printf("CseqVal %+v\n", hepPacket.SIP.CseqVal)
		log.Printf("Reason %+v\n", hepPacket.SIP.Reason)
		log.Printf("ReasonVal %+v\n", hepPacket.SIP.ReasonVal)
		log.Printf("RTPStatVal %+v\n", hepPacket.SIP.RTPStatVal)
		log.Printf("ViaOne %+v\n", hepPacket.SIP.ViaOne)
		log.Printf("ViaOneBranch %+v\n", hepPacket.SIP.ViaOneBranch)
		log.Printf("Privacy %+v\n", hepPacket.SIP.Privacy)
		log.Printf("RemotePartyIdVal %+v\n", hepPacket.SIP.RemotePartyIdVal)
		log.Printf("DiversionVal %+v\n", hepPacket.SIP.DiversionVal)
		log.Printf("RemotePartyId %+v\n", hepPacket.SIP.RemotePartyId)
		log.Printf("PAssertedIdVal %+v\n", hepPacket.SIP.PAssertedIdVal)
		log.Printf("PaiUser %+v\n", hepPacket.SIP.PaiUser)
		log.Printf("PaiHost %+v\n", hepPacket.SIP.PaiHost)
		log.Printf("PAssertedId %+v\n", hepPacket.SIP.PAssertedId)
		log.Printf("UserAgent %+v\n", hepPacket.SIP.UserAgent)
		log.Printf("Server %+v\n", hepPacket.SIP.Server)
		log.Printf("URIHost %+v\n", hepPacket.SIP.URIHost)
		log.Printf("URIRaw %+v\n", hepPacket.SIP.URIRaw)
		log.Printf("URIUser %+v\n", hepPacket.SIP.URIUser)
		log.Printf("FirstMethod %+v\n", hepPacket.SIP.FirstMethod)
		log.Printf("FirstResp %+v\n", hepPacket.SIP.FirstResp)
		log.Printf("FirstRespText %+v\n", hepPacket.SIP.FirstRespText)*/
}

/*
2019/12/27 22:51:53 Version 2
2019/12/27 22:51:53 Protocol 17
2019/12/27 22:51:53 SrcIP 185.247.118.201
2019/12/27 22:51:53 DstIP 208.64.201.20
2019/12/27 22:51:53 SrcPort 5080
2019/12/27 22:51:53 DstPort 5060
2019/12/27 22:51:53 Tsec 1577487113
2019/12/27 22:51:53 Tmsec 269927
2019/12/27 22:51:53 ProtoType 1
2019/12/27 22:51:53 NodeID 200
2019/12/27 22:51:53 NodePW
2019/12/27 22:51:53 Payload REGISTER sip:asterlink.com;transport=udp SIP/2.0
Via: SIP/2.0/UDP 185.247.118.201:5080;rport;branch=z9hG4bK31mQK3m5mKvra
Max-Forwards: 70
From: <sip:cluecon@asterlink.com>;tag=cacNSm5r9727e
To: <sip:cluecon@asterlink.com>
Call-ID: e2a8c4ab-edb0-40bb-b704-454e5a673462
CSeq: 14189759 REGISTER
Contact: <sip:gw+asterlink.com@185.247.118.201:5080;transport=udp;gw=asterlink.com>
Expires: 60
User-Agent: FreeSWITCH-mod_sofia/1.10.1-release-12-f9990221e6~64bit
Allow: INVITE, ACK, BYE, CANCEL, OPTIONS, MESSAGE, INFO, UPDATE, REGISTER, REFER, NOTIFY
Supported: timer, path, replaces
Content-Length: 0


2019/12/27 22:51:53 CID e2a8c4ab-edb0-40bb-b704-454e5a673462
2019/12/27 22:51:53 Vlan 0
2019/12/27 22:51:53 ProtoString sip
2019/12/27 22:51:53 Timestamp 2019-12-27 22:51:53.269927 +0000 UTC
2019/12/27 22:51:53 NodeName 200
2019/12/27 22:51:53 SID e2a8c4ab-edb0-40bb-b704-454e5a673462
2019/12/27 22:51:53 State SipParseStateStartLine
2019/12/27 22:51:53 Error <nil>
2019/12/27 22:51:53 Msg REGISTER sip:asterlink.com;transport=udp SIP/2.0
Via: SIP/2.0/UDP 185.247.118.201:5080;rport;branch=z9hG4bK31mQK3m5mKvra
Max-Forwards: 70
From: <sip:cluecon@asterlink.com>;tag=cacNSm5r9727e
To: <sip:cluecon@asterlink.com>
Call-ID: e2a8c4ab-edb0-40bb-b704-454e5a673462
CSeq: 14189759 REGISTER
Contact: <sip:gw+asterlink.com@185.247.118.201:5080;transport=udp;gw=asterlink.com>
Expires: 60
User-Agent: FreeSWITCH-mod_sofia/1.10.1-release-12-f9990221e6~64bit
Allow: INVITE, ACK, BYE, CANCEL, OPTIONS, MESSAGE, INFO, UPDATE, REGISTER, REFER, NOTIFY
Supported: timer, path, replaces
Content-Length: 0


2019/12/27 22:51:53 CallingParty <nil>
2019/12/27 22:51:53 Body
2019/12/27 22:51:53 Authorization <nil>
2019/12/27 22:51:53 AuthVal
2019/12/27 22:51:53 AuthUser
2019/12/27 22:51:53 ContentLength 0
2019/12/27 22:51:53 ContentType
2019/12/27 22:51:53 From &{Error:<nil> Val:<sip:cluecon@asterlink.com>;tag=cacNSm5r9727e Name: Tag:cacNSm5r9727e URI:0xc000091550 endName:0 rightBrack:26 leftBrack:0 brackChk:true}
2019/12/27 22:51:53 FromUser cluecon
2019/12/27 22:51:53 FromHost asterlink.com
2019/12/27 22:51:53 FromTag cacNSm5r9727e
2019/12/27 22:51:53 MaxForwards 70
2019/12/27 22:51:53 Organization
2019/12/27 22:51:53 To &{Error:<nil> Val:<sip:cluecon@asterlink.com> Name: Tag: URI:0xc000091600 endName:0 rightBrack:26 leftBrack:0 brackChk:true}
2019/12/27 22:51:53 ToUser cluecon
2019/12/27 22:51:53 ToHost asterlink.com
2019/12/27 22:51:53 ToTag
2019/12/27 22:51:53 Contact &{Error:<nil> Val:Contact: <sip:gw+asterlink.com@185.247.118.201:5080;transport=udp;gw=asterlink.com> Name:Contact: Tag: URI:0xc0000916b0 endName:9 rightBrack:82 leftBrack:9 brackChk:true}
2019/12/27 22:51:53 ContactVal <sip:gw+asterlink.com@185.247.118.201:5080;transport=udp;gw=asterlink.com>
2019/12/27 22:51:53 ContactUser gw+asterlink.com
2019/12/27 22:51:53 ContactHost 185.247.118.201
2019/12/27 22:51:53 ContactPort 5080
2019/12/27 22:51:53 CallID e2a8c4ab-edb0-40bb-b704-454e5a673462
2019/12/27 22:51:53 XCallID
2019/12/27 22:51:53 XHeader []
2019/12/27 22:51:53 Cseq &{Val:14189759 REGISTER Method:REGISTER Digit:14189759}
2019/12/27 22:51:53 CseqMethod REGISTER
2019/12/27 22:51:53 CseqVal 14189759 REGISTER
2019/12/27 22:51:53 Reason <nil>
2019/12/27 22:51:53 ReasonVal
2019/12/27 22:51:53 RTPStatVal
2019/12/27 22:51:53 ViaOne SIP/2.0/UDP 185.247.118.201:5080;rport;branch=z9hG4bK31mQK3m5mKvra
2019/12/27 22:51:53 ViaOneBranch z9hG4bK31mQK3m5mKvra
2019/12/27 22:51:53 Privacy
2019/12/27 22:51:53 RemotePartyIdVal
2019/12/27 22:51:53 DiversionVal
2019/12/27 22:51:53 RemotePartyId <nil>
2019/12/27 22:51:53 PAssertedIdVal
2019/12/27 22:51:53 PaiUser
2019/12/27 22:51:53 PaiHost
2019/12/27 22:51:53 PAssertedId <nil>
2019/12/27 22:51:53 UserAgent FreeSWITCH-mod_sofia/1.10.1-release-12-f9990221e6~64bit
2019/12/27 22:51:53 Server
2019/12/27 22:51:53 URIHost asterlink.com
2019/12/27 22:51:53 URIRaw asterlink.com;transport=udp
2019/12/27 22:51:53 URIUser
2019/12/27 22:51:53 FirstMethod REGISTER
2019/12/27 22:51:53 FirstResp
2019/12/27 22:51:53 FirstRespText
*/
