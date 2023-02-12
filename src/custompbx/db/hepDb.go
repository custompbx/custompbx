package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/custompbx/hepparser"
	"log"
	"strconv"
	"strings"
)

func InitHEPDb() {
	createHepTable(db)
	createHEPTableHepTimestampIndex(db)
	createHEPTableSipCallIdIndex(db)
	createHEPTableSipFirstMethodIndex(db)
	createHEPTableSipFromUserIndex(db)
	createHEPTableSipFromHostIndex(db)
	createHEPTableSipUriUserIndex(db)
}

func createHepTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS hep_packets(
		hep_timestamp TIMESTAMP,
		hep_sid VARCHAR,
		hep_node_name VARCHAR,
		hep_proto_string VARCHAR,
		hep_cid VARCHAR,
		hep_payload VARCHAR,
		hep_node_pw VARCHAR,
		hep_dst_ip VARCHAR,
		hep_src_ip VARCHAR,
		hep_vlan INTEGER,
		hep_node_id INTEGER,
		hep_proto_type INTEGER,
		hep_t_msec INTEGER,
		hep_t_sec INTEGER,
		hep_dst_port INTEGER,
		hep_src_port INTEGER,
		hep_protocol INTEGER,
		hep_version INTEGER,
		sip_state VARCHAR,
		sip_msg VARCHAR,
		sip_body VARCHAR,
		sip_auth_val VARCHAR,
		sip_auth_user VARCHAR,
		sip_content_type VARCHAR,
		sip_from_user VARCHAR,
		sip_from_host VARCHAR,
		sip_from_tag VARCHAR,
		sip_organization VARCHAR,
		sip_max_forwards VARCHAR,
		sip_to_user VARCHAR,
		sip_to_host VARCHAR,
		sip_Content_Length VARCHAR,
		sip_to_tag VARCHAR,
		sip_contact_val VARCHAR,
		sip_contact_user VARCHAR,
		sip_contact_host VARCHAR,
		sip_call_id VARCHAR,
		sip_xcall_id VARCHAR,
		sip_cseq_method VARCHAR,
		sip_cseq_val VARCHAR,
		sip_reason_val VARCHAR,
		sip_rtpstat_val VARCHAR,
		sip_via_one VARCHAR,
		sip_viaone_branch VARCHAR,
		sip_privacy VARCHAR,
		sip_remote_partyid_val VARCHAR,
		sip_diversion_val VARCHAR,
		sip_passertedid_val VARCHAR,
		sip_pai_user VARCHAR,
		sip_pai_host VARCHAR,
		sip_user_agent VARCHAR,
		sip_server VARCHAR,
		sip_uri_host VARCHAR,
		sip_uri_raw VARCHAR,
		sip_uri_user VARCHAR,
		sip_first_method VARCHAR,
		sip_first_resp VARCHAR,
		sip_first_resp_text VARCHAR,
		sip_hdr VARCHAR,
		sip_hdrv VARCHAR,
		sip_contact_port integer,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createHEPTableHepTimestampIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_hep_timestamp ON hep_packets(
		hep_timestamp
	);`,
	)
	panicErr(err)
}

func createHEPTableSipCallIdIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_hep_sip_call_id ON hep_packets(sip_call_id);`,
	)
	panicErr(err)
}

func createHEPTableSipFromUserIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_hep_sip_from_user ON hep_packets(sip_from_user);`,
	)
	panicErr(err)
}

func createHEPTableSipFromHostIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_hep_sip_from_host ON hep_packets(sip_from_host);`,
	)
	panicErr(err)
}

func createHEPTableSipUriUserIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_hep_sip_uri_user ON hep_packets(sip_uri_user);`,
	)
	panicErr(err)
}

func createHEPTableSipFirstMethodIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_hep_sip_first_method ON hep_packets(sip_first_method);`,
	)
	panicErr(err)
}

/*eof              int
x_header          []string
calling_party     *CallingPartyInfo
authorization    *Authorization
from             *From
to               *From
contact          *From
cseq             *Cseq
reason           *Reason
remote_party_id    *RemotePartyId
passerted_id      *PAssertedId*/

func SaveHEPPacket(packet *hepparser.HEP) error {
	if packet == nil || packet.SIP == nil {
		return errors.New("empty sip msg")
	}
	sqlReq := `INSERT INTO hep_packets(        
		hep_timestamp,
		hep_sid,
		hep_node_name,
		hep_proto_string,
		hep_cid,
		hep_payload,
		hep_node_pw,
		hep_dst_ip,
		hep_src_ip,
		hep_vlan,
		hep_node_id,
		hep_proto_type,
		hep_t_msec,
		hep_t_sec,
		hep_dst_port,
		hep_src_port,
		hep_protocol,
		hep_version,       
		sip_state,
		sip_msg,
		sip_body,
		sip_auth_val,
		sip_auth_user,
		sip_content_type,
		sip_from_user,
		sip_from_host,
		sip_from_tag,
		sip_organization,
		sip_max_forwards,
		sip_to_user,
		sip_to_host,
		sip_Content_Length,
		sip_to_tag,
		sip_contact_val,
		sip_contact_user,
		sip_contact_host,
		sip_call_id,
		sip_xcall_id,
		sip_cseq_method,
		sip_cseq_val,
		sip_reason_val,
		sip_rtpstat_val,
		sip_via_one,
		sip_viaone_branch,
		sip_privacy,
		sip_remote_partyid_val,
		sip_diversion_val,
		sip_passertedid_val,
		sip_pai_user,
		sip_pai_host,
		sip_user_agent,
		sip_server,
		sip_uri_host,
		sip_uri_raw,
		sip_uri_user,
		sip_first_method,
		sip_first_resp,
		sip_first_resp_text,
		sip_contact_port
) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29,
         $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59
         )`
	_, err := db.Exec(sqlReq, packet.Timestamp, packet.SID, packet.NodeName, packet.ProtoString, packet.CID, packet.Payload, packet.NodePW, packet.DstIP,
		packet.SrcIP, packet.Vlan, packet.NodeID, packet.ProtoType, packet.Tmsec, packet.Tsec, packet.DstPort, packet.SrcPort, packet.Protocol, packet.Version,
		packet.SIP.State, packet.SIP.Msg, packet.SIP.Body, packet.SIP.AuthVal, packet.SIP.AuthUser, packet.SIP.ContentType, packet.SIP.FromUser, packet.SIP.FromHost,
		packet.SIP.FromTag, packet.SIP.Organization, packet.SIP.MaxForwards, packet.SIP.ToUser, packet.SIP.ToHost, packet.SIP.ContentLength, packet.SIP.ToTag, packet.SIP.ContactVal,
		packet.SIP.ContactUser, packet.SIP.ContactHost, packet.SIP.CallID, packet.SIP.XCallID, packet.SIP.CseqMethod, packet.SIP.CseqVal, packet.SIP.ReasonVal, packet.SIP.RTPStatVal,
		packet.SIP.ViaOne, packet.SIP.ViaOneBranch, packet.SIP.Privacy, packet.SIP.RemotePartyIdVal, packet.SIP.DiversionVal, packet.SIP.PAssertedIdVal, packet.SIP.PaiUser, packet.SIP.PaiHost,
		packet.SIP.UserAgent, packet.SIP.Server, packet.SIP.URIHost, packet.SIP.URIRaw, packet.SIP.URIUser, packet.SIP.FirstMethod, packet.SIP.FirstResp, packet.SIP.FirstRespText,
		packet.SIP.ContactPort)
	if err != nil {
		return err
	}

	return err
}

func clearOldHEPs() {
	_, err := db.Exec(
		`DELETE from hep_packets
               WHERE hep_timestamp < now() - INTERVAL '1 day';`)
	if err != nil {
		log.Printf("clearOldHEPLogs ERROR: %+v", err.Error())
	}
}

func SaveHEPPackets(packets []*hepparser.HEP, instanceId int64) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in SaveHEPPackets", r)
		}
	}()

	query := `INSERT INTO hep_packets(        
		hep_timestamp,
		hep_sid,
		hep_node_name,
		hep_proto_string,
		hep_cid,
		hep_payload,
		hep_node_pw,
		hep_dst_ip,
		hep_src_ip,
		hep_vlan,
		hep_node_id,
		hep_proto_type,
		hep_t_msec,
		hep_t_sec,
		hep_dst_port,
		hep_src_port,
		hep_protocol,
		hep_version,       
		sip_state,
		sip_msg,
		sip_body,
		sip_auth_val,
		sip_auth_user,
		sip_content_type,
		sip_from_user,
		sip_from_host,
		sip_from_tag,
		sip_organization,
		sip_max_forwards,
		sip_to_user,
		sip_to_host,
		sip_Content_Length,
		sip_to_tag,
		sip_contact_val,
		sip_contact_user,
		sip_contact_host,
		sip_call_id,
		sip_xcall_id,
		sip_cseq_method,
		sip_cseq_val,
		sip_reason_val,
		sip_rtpstat_val,
		sip_via_one,
		sip_viaone_branch,
		sip_privacy,
		sip_remote_partyid_val,
		sip_diversion_val,
		sip_passertedid_val,
		sip_pai_user,
		sip_pai_host,
		sip_user_agent,
		sip_server,
		sip_uri_host,
		sip_uri_raw,
		sip_uri_user,
		sip_first_method,
		sip_first_resp,
		sip_first_resp_text,
		sip_contact_port,
        instance_id
) VALUES`
	var values []interface{}

	counterMultipiler := 0
	for i, packet := range packets {
		if packet == nil || packet.SIP == nil || packet.Timestamp.IsZero() {
			log.Println("empty sip msg")
			counterMultipiler++
			continue
		}
		values = append(values, packet.Timestamp, packet.SID, packet.NodeName, packet.ProtoString, packet.CID, packet.Payload, packet.NodePW, packet.DstIP,
			packet.SrcIP, packet.Vlan, packet.NodeID, packet.ProtoType, packet.Tmsec, packet.Tsec, packet.DstPort, packet.SrcPort, packet.Protocol, packet.Version,
			packet.SIP.State, packet.SIP.Msg, packet.SIP.Body, packet.SIP.AuthVal, packet.SIP.AuthUser, packet.SIP.ContentType, packet.SIP.FromUser, packet.SIP.FromHost,
			packet.SIP.FromTag, packet.SIP.Organization, packet.SIP.MaxForwards, packet.SIP.ToUser, packet.SIP.ToHost, packet.SIP.ContentLength, packet.SIP.ToTag, packet.SIP.ContactVal,
			packet.SIP.ContactUser, packet.SIP.ContactHost, packet.SIP.CallID, packet.SIP.XCallID, packet.SIP.CseqMethod, packet.SIP.CseqVal, packet.SIP.ReasonVal, packet.SIP.RTPStatVal,
			packet.SIP.ViaOne, packet.SIP.ViaOneBranch, packet.SIP.Privacy, packet.SIP.RemotePartyIdVal, packet.SIP.DiversionVal, packet.SIP.PAssertedIdVal, packet.SIP.PaiUser, packet.SIP.PaiHost,
			packet.SIP.UserAgent, packet.SIP.Server, packet.SIP.URIHost, packet.SIP.URIRaw, packet.SIP.URIUser, packet.SIP.FirstMethod, packet.SIP.FirstResp, packet.SIP.FirstRespText,
			packet.SIP.ContactPort, instanceId)

		numFields := 60

		n := (i - counterMultipiler) * numFields

		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1]
	_, err := db.Exec(query, values...)
	if err != nil {
		log.Printf("%+v", err.Error())
		return err
	}
	return nil
}

func GetHEPDetailsList(callIds []string, instanceId int64) ([]map[string]*interface{}, error) {
	table := "hep_packets"

	var fieldsArr = []string{
		"hep_timestamp",
		"hep_dst_ip",
		"hep_src_ip",
		"hep_dst_port",
		"hep_src_port",
		"hep_payload",
		"sip_first_method",
		"sip_cseq_method",
	}

	fieldsArrModified := make([]string, len(fieldsArr))
	copy(fieldsArrModified, fieldsArr)
	fieldsArrModified[0] = fmt.Sprintf("to_char(%s, 'YYYY-MM-DD HH24:MI:SS.MS')", fieldsArrModified[0])

	queryBuilder := squirrel.Select(fieldsArrModified...).From(table).PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{"instance_id": instanceId})
	queryBuilder = queryBuilder.Where(fmt.Sprintf("sip_call_id IN ('%s')", strings.Join(callIds[:], "','")))
	queryBuilder = queryBuilder.OrderBy("hep_timestamp ASC")

	query, args, _ := queryBuilder.ToSql()
	log.Println(query)
	log.Println(args)
	logs, err := db.Query(query, args...)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error on select")
	}
	defer logs.Close()

	var logLines []map[string]*interface{}

	for logs.Next() {
		var rows []interface{}
		result := make(map[string]*interface{})
		for _, name := range fieldsArr {
			var str interface{}
			rows = append(rows, &str)
			result[name] = &str
		}
		err := logs.Scan(rows...)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		logLines = append(logLines, result)
	}
	return logLines, nil
}
