package db

import "fmt"

func allSql() {
	fmt.Sprintf(`CREATE DATABASE customconf;`)
	fmt.Sprintf(`CREATE ROLE custompbx LOGIN PASSWORD 'custompbx';`)
	fmt.Sprintf(` GRANT ALL ON DATABASE customconf TO custompbx;`)
	fmt.Sprintf(` GRANT ALL ON DATABASE customconf TO custompbx;`)
	fmt.Sprintf(`
hostname=testpopov
section=directory
tag_name=domain
key_name=name
key_value=[46.229.223.143:5080]
key=id
user=3890
domain=%5B46.229.223.143%3A5080%5D
action=sip_auth
Event-Name=REQUEST_PARAMS
Core-UUID=83ce4dc7-bc82-46eb-b17e-a33e859af9ed
FreeSWITCH-Hostname=testpopov
FreeSWITCH-Switchname=testpopov
FreeSWITCH-IPv4=46.229.223.143
FreeSWITCH-IPv6=%3A%3A1
Event-Date-Local=2018-06-07%2022%3A23%3A59
Event-Date-GMT=Thu,%2007%20Jun%202018%2019%3A23%3A59%20GMT
Event-Date-Timestamp=1528399439482986
Event-Calling-File=sofia_reg.c
Event-Calling-Function=sofia_reg_parse_auth
Event-Calling-Line-Number=2820
Event-Sequence=1009546
sip_profile=external
sip_user_agent=VaxSIPUserAgent/3.5
sip_auth_username=3890
sip_auth_realm=%5B46.229.223.143%3A5080%5D
sip_auth_nonce=71ec2b31-7e52-484d-8f22-f9d002674966
sip_auth_uri=sip%3A46.229.223.143%3A5080
sip_contact_user=3890
sip_contact_host=94.23.35.226
sip_to_user=3890
sip_to_host=%5B46.229.223.143%3A5080%5D
sip_via_protocol=udp
sip_from_user=3890
sip_from_host=%5B46.229.223.143%3A5080%5D
sip_call_id=5c4a2273-b925d3422ef71c22-e7add524%4046.229.223.143%3A5080
sip_request_host=46.229.223.143
sip_request_port=5080
sip_auth_qop=auth
sip_auth_cnonce=5c4a22d0
sip_auth_nc=00000002
sip_auth_response=88fe378d971b0d56906a1f2c3ddafb35
sip_auth_method=REGISTER
client_port=5099
ip=94.23.35.226
`)
	fmt.Sprintf(`
<section name="directory" description="User Directory">
<domain name="46.229.223.143">
<params>
<param name="dial-string" value="{^^:sip_invite_domain=${dialed_domain}:presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(*/${dialed_user}@${dialed_domain})},${verto_contact(${dialed_user}@${dialed_domain})}">
</param>
<param name="jsonrpc-allowed-methods" value="verto">
</param>
</params>
<variables>
<variable name="record_stereo" value="true">
</variable>
<variable name="default_gateway" value="example.com">
</variable>
<variable name="default_areacode" value="918">
</variable>
<variable name="transfer_fallback_extension" value="operator">
</variable>
</variables>
<groups>
<group name="default">
<users>
<user id="1000">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1000">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1000">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1000">
</variable>
<variable name="effective_caller_id_number" value="1000">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1001">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1001">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1001">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1001">
</variable>
<variable name="effective_caller_id_number" value="1001">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1002">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1002">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1002">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1002">
</variable>
<variable name="effective_caller_id_number" value="1002">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1003">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1003">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1003">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1003">
</variable>
<variable name="effective_caller_id_number" value="1003">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1004">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1004">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1004">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1004">
</variable>
<variable name="effective_caller_id_number" value="1004">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1005">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1005">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1005">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1005">
</variable>
<variable name="effective_caller_id_number" value="1005">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1006">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1006">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1006">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1006">
</variable>
<variable name="effective_caller_id_number" value="1006">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1007">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1007">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1007">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1007">
</variable>
<variable name="effective_caller_id_number" value="1007">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1008">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1008">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1008">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1008">
</variable>
<variable name="effective_caller_id_number" value="1008">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1009">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1009">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1009">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1009">
</variable>
<variable name="effective_caller_id_number" value="1009">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1010">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1010">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1010">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1010">
</variable>
<variable name="effective_caller_id_number" value="1010">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1011">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1011">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1011">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1011">
</variable>
<variable name="effective_caller_id_number" value="1011">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1012">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1012">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1012">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1012">
</variable>
<variable name="effective_caller_id_number" value="1012">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1013">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1013">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1013">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1013">
</variable>
<variable name="effective_caller_id_number" value="1013">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1014">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1014">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1014">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1014">
</variable>
<variable name="effective_caller_id_number" value="1014">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1015">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1015">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1015">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1015">
</variable>
<variable name="effective_caller_id_number" value="1015">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1016">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1016">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1016">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1016">
</variable>
<variable name="effective_caller_id_number" value="1016">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1017">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1017">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1017">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1017">
</variable>
<variable name="effective_caller_id_number" value="1017">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1018">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1018">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1018">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1018">
</variable>
<variable name="effective_caller_id_number" value="1018">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="1019">
<params>
<param name="password" value="1234">
</param>
<param name="vm-password" value="1019">
</param>
</params>
<variables>
<variable name="toll_allow" value="domestic,international,local">
</variable>
<variable name="accountcode" value="1019">
</variable>
<variable name="user_context" value="default">
</variable>
<variable name="effective_caller_id_name" value="Extension 1019">
</variable>
<variable name="effective_caller_id_number" value="1019">
</variable>
<variable name="outbound_caller_id_name" value="FreeSWITCH">
</variable>
<variable name="outbound_caller_id_number" value="0000000000">
</variable>
<variable name="callgroup" value="techsupport">
</variable>
</variables>
</user>
<user id="brian" cidr="192.0.2.0/24">
	<gateways>
	</gateways>
	<params>
		<param name="password" value="1234">
		</param>
		<param name="vm-password" value="9999">
		</param>
	</params>
	<variables>
		<variable name="user_context" value="default">
		</variable>
		<variable name="effective_caller_id_name" value="Brian West">
		</variable>
		<variable name="effective_caller_id_number" value="1000">
		</variable>
		<variable name="process_cdr" value="true">
		</variable>
		<variable name="rtp_secure_media" value="true">
		</variable>
	</variables>
<vcard>
</vcard>
</user>
<user id="default">
	<variables>
		<variable name="numbering_plan" value="US">
		</variable>
		<variable name="default_areacode" value="918">
		</variable>
		<variable name="default_gateway" value="example.com">
		</variable>
	</variables>
</user>
<user id="example.com">
	<gateways>
		<gateway name="example.com">
			<param name="username" value="joeuser">
			</param>
			<param name="password" value="password">
			</param>
			<param name="from-user" value="joeuser">
			</param>
			<param name="from-domain" value="example.com">
			</param>
			<param name="expire-seconds" value="600">
			</param>
			<param name="register" value="false">
			</param>
			<param name="retry-seconds" value="30">
			</param>
			<param name="extension" value="5000">
			</param>
			<param name="context" value="public">
			</param>
		</gateway>
	</gateways>
<params>
<param name="password" value="password">
</param>
</params>
</user>
		<user id="SEP001120AABBCC">
		<params>
		<param name="foo" value="bar">
		</param>
		</params>
		<skinny>
		<buttons>
		<button position="1" type="Line" label="Line 1" value="1101" caller-name="Calling as 1101">
		</button>
		<button position="3" type="Line" label="Shared Line 10" value="1110" caller-name="Calling as 1110">
		</button>
		<button position="5" type="SpeedDial" label="Call 1001" value="1001">
		</button>
		<button position="6" type="ServiceUrl" label="Some URL" value="http://phone-xml.berbee.com/menu.xml">
		</button>
		</buttons>
		</skinny>
		</user>
</users>
</group>
<group name="sales">
<users>
<user id="1000" type="pointer">
</user>
<user id="1001" type="pointer">
</user>
<user id="1002" type="pointer">
</user>
<user id="1003" type="pointer">
</user>
<user id="1004" type="pointer">
</user>
</users>
</group>
<group name="billing">
<users>
<user id="1005" type="pointer">
</user>
<user id="1006" type="pointer">
</user>
<user id="1007" type="pointer">
</user>
<user id="1008" type="pointer">
</user>
<user id="1009" type="pointer">
</user>
</users>
</group>
<group name="support">
<users>
<user id="1010" type="pointer">
</user>
<user id="1011" type="pointer">
</user>
<user id="1012" type="pointer">
</user>
<user id="1013" type="pointer">
</user>
<user id="1014" type="pointer">
</user>
</users>
</group>
</groups>
</domain>
</section> 
`)
	fmt.Sprintf(`
hostname=testpopov
section=directory
tag_name=domain
key_name=name
key_value=%{domain_name}
Event-Name=REQUEST_PARAMS
Core-UUID=83ce4dc7-bc82-46eb-b17e-a33e859af9ed
FreeSWITCH-Hostname=testpopov
FreeSWITCH-Switchname=testpopov
FreeSWITCH-IPv4=46.229.223.143
FreeSWITCH-IPv6=%3A%3A1
Event-Date-Local=2018-06-10%2016%3A37%3A25
Event-Date-GMT=Sun,%2010%20Jun%202018%2013%3A37%3A25%20GMT
Event-Date-Timestamp=1528637845563102
Event-Calling-File=mod_commands.c
Event-Calling-Function=group_call_function
Event-Calling-Line-Number=1026
Event-Sequence=1152018
group=support
action=group_call
group_name=support
domain=%25%7Bdomain_name%7D]
`)
	fmt.Sprintf(`
<section name="configuration" description="Various Configuration">
  <configuration name="abstraction.conf" description="Abstraction">
    <apis>
	      <api name="user_name" description="Return Name for extension" syntax="&lt;exten&gt;" parse="(.*)" destination="user_data" argument="$1@default var effective_caller_id_name"></api>
    </apis>
  </configuration>
  <configuration name="acl.conf" description="Network Lists">
      <network-lists>
          <list name="lan" default="allow">
              <node type="deny" cidr="192.168.42.0/24"></node>
              <node type="allow" cidr="192.168.42.42/32"></node>
    </list>
          <list name="domains" default="deny">
              <node type="allow" domain="46.229.223.143"></node>
    </list>
  </network-lists>
  </configuration>
  <configuration name="alsa.conf" description="Soundcard Endpoint">
      <settings>
          <param name="dialplan" value="XML"></param>
          <param name="cid-name" value="N800 Alsa"></param>
          <param name="cid-num" value="5555551212"></param>
          <param name="sample-rate" value="8000"></param>
          <param name="codec-ms" value="20"></param>
  </settings>
  </configuration>
  <configuration name="amqp.conf" description="mod_amqp">
      <producers>
          <profile name="default">
              <connections>
	          <connection name="primary">
	              <param name="hostname" value="localhost"></param>
	              <param name="virtualhost" value="/"></param>
	              <param name="username" value="guest"></param>
	              <param name="password" value="guest"></param>
	              <param name="port" value="5673"></param>
	              <param name="heartbeat" value="0"></param>
	</connection>
	          <connection name="secondary">
	              <param name="hostname" value="localhost"></param>
	              <param name="virtualhost" value="/"></param>
	              <param name="username" value="guest"></param>
	              <param name="password" value="guest"></param>
	              <param name="port" value="5672"></param>
	              <param name="heartbeat" value="0"></param>
	</connection>
      </connections>
              <params>
	          <param name="exchange-name" value="TAP.Events"></param>
	          <param name="exchange-type" value="topic"></param>
	          <param name="circuit_breaker_ms" value="10000"></param>
	          <param name="reconnect_interval_ms" value="1000"></param>
	          <param name="send_queue_size" value="5000"></param>
	          <param name="enable_fallback_format_fields" value="1"></param>
	          <param name="format_fields" value="#FreeSWITCH,FreeSWITCH-Hostname,Event-Name,Event-Subclass,Unique-ID"></param>
	          <param name="event_filter" value="SWITCH_EVENT_CHANNEL_CREATE,SWITCH_EVENT_CHANNEL_DESTROY,SWITCH_EVENT_HEARTBEAT,SWITCH_EVENT_DTMF"></param>
      </params>
    </profile>
  </producers>
      <commands>
          <profile name="default">
              <connections>
	          <connection name="primary">
	              <param name="hostname" value="localhost"></param>
	              <param name="virtualhost" value="/"></param>
	              <param name="username" value="guest"></param>
	              <param name="password" value="guest"></param>
	              <param name="port" value="5672"></param>
	              <param name="heartbeat" value="0"></param>
	</connection>
      </connections>
              <params>
	          <param name="exchange-name" value="TAP.Commands"></param>
	          <param name="binding_key" value="commandBindingKey"></param>
	          <param name="reconnect_interval_ms" value="1000"></param>
      </params>
    </profile>
  </commands>
      <logging>
          <profile name="default">
              <connections>
	          <connection name="primary">
	              <param name="hostname" value="localhost"></param>
	              <param name="virtualhost" value="/"></param>
	              <param name="username" value="guest"></param>
	              <param name="password" value="guest"></param>
	              <param name="port" value="5672"></param>
	              <param name="heartbeat" value="0"></param>
	</connection>
      </connections>
              <params>
	          <param name="exchange-name" value="TAP.Logging"></param>
	          <param name="send_queue_size" value="5000"></param>
	          <param name="reconnect_interval_ms" value="1000"></param>
	          <param name="log-levels" value="debug,info,notice,warning,err,crit,alert"></param>
      </params>
    </profile>
  </logging>
  </configuration>
  <configuration name="avmd.conf" description="AVMD config">
        <settings>
            
                  <param name="debug" value="0"></param>
            
                  <param name="report_status" value="1"></param>
            
                  <param name="fast_math" value="0"></param>
            
                  <param name="require_continuous_streak" value="1"></param>
            
                  <param name="sample_n_continuous_streak" value="15"></param>
            
                  <param name="sample_n_to_skip" value="15"></param>
            
                  <param name="simplified_estimation" value="1"></param>
            
                  <param name="inbound_channel" value="0"></param>
            
                  <param name="outbound_channel" value="1"></param>
    </settings>
  </configuration>
  <configuration name="mod_blacklist.conf" description="Blacklist module">
    <lists>
    </lists>
  </configuration>
  <configuration name="callcenter.conf" description="CallCenter">
      <settings>
  </settings>
      <queues>
          <queue name="support@default">
              <param name="strategy" value="longest-idle-agent"></param>
              <param name="moh-sound" value="local_stream://moh"></param>
              <param name="time-base-score" value="system"></param>
              <param name="max-wait-time" value="0"></param>
              <param name="max-wait-time-with-no-agent" value="0"></param>
              <param name="max-wait-time-with-no-agent-time-reached" value="5"></param>
              <param name="tier-rules-apply" value="false"></param>
              <param name="tier-rule-wait-second" value="300"></param>
              <param name="tier-rule-wait-multiply-level" value="true"></param>
              <param name="tier-rule-no-agent-no-wait" value="false"></param>
              <param name="discard-abandoned-after" value="60"></param>
              <param name="abandoned-resume-allowed" value="false"></param>
    </queue>
  </queues>
      <agents>
  </agents>
      <tiers>
  </tiers>
  </configuration>
  <configuration name="cdr_csv.conf" description="CDR CSV Format">
      <settings>
          <param name="default-template" value="example"></param>
          <param name="rotate-on-hup" value="true"></param>
          <param name="legs" value="a"></param>
  </settings>
      <templates>
          <template name="sql">INSERT INTO cdr VALUES ("${caller_id_name}","${caller_id_number}","${destination_number}","${context}","${start_stamp}","${answer_stamp}","${end_stamp}","${duration}","${billsec}","${hangup_cause}","${uuid}","${bleg_uuid}", "${accountcode}");</template>
          <template name="example">"${caller_id_name}","${caller_id_number}","${destination_number}","${context}","${start_stamp}","${answer_stamp}","${end_stamp}","${duration}","${billsec}","${hangup_cause}","${uuid}","${bleg_uuid}","${accountcode}","${read_codec}","${write_codec}"</template>
          <template name="snom">"${caller_id_name}","${caller_id_number}","${destination_number}","${context}","${start_stamp}","${answer_stamp}","${end_stamp}","${duration}","${billsec}","${hangup_cause}","${uuid}","${bleg_uuid}", "${accountcode}","${read_codec}","${write_codec}","${sip_user_agent}","${call_clientcode}","${sip_rtp_rxstat}","${sip_rtp_txstat}","${sofia_record_file}"</template>
          <template name="linksys">"${caller_id_name}","${caller_id_number}","${destination_number}","${context}","${start_stamp}","${answer_stamp}","${end_stamp}","${duration}","${billsec}","${hangup_cause}","${uuid}","${bleg_uuid}","${accountcode}","${read_codec}","${write_codec}","${sip_user_agent}","${sip_p_rtp_stat}"</template>
          <template name="asterisk">"${accountcode}","${caller_id_number}","${destination_number}","${context}","${caller_id}","${channel_name}","${bridge_channel}","${last_app}","${last_arg}","${start_stamp}","${answer_stamp}","${end_stamp}","${duration}","${billsec}","${hangup_cause}","${amaflags}","${uuid}","${userfield}"</template>
          <template name="opencdrrate">"${uuid}","${signal_bond}","${direction}","${ani}","${destination_number}","${answer_stamp}","${end_stamp}","${billsec}","${accountcode}","${userfield}","${network_addr}","${regex('${original_caller_id_name}'|^.)}","${sip_gateway_name}"</template>
  </templates>
  </configuration>
  <configuration name="cdr_mongodb.conf" description="MongoDB CDR logger">
      <settings>
          <param name="host" value="127.0.0.1"></param>
          <param name="port" value="27017"></param>
          <param name="namespace" value="test.cdr"></param>
          <param name="log-b-leg" value="false"></param>
  </settings>
  </configuration>
  <configuration name="cdr_pg_csv.conf" description="CDR PG CSV Format">
      <settings>
          <param name="db-info" value="host=localhost dbname=cdr connect_timeout=10"></param>
          <param name="legs" value="a"></param>
          <param name="spool-format" value="csv"></param>
          <param name="rotate-on-hup" value="true"></param>
  </settings>
      <schema>
          <field var="local_ip_v4"></field>
          <field var="caller_id_name"></field>
          <field var="caller_id_number"></field>
          <field var="destination_number"></field>
          <field var="context"></field>
          <field var="start_stamp"></field>
          <field var="answer_stamp"></field>
          <field var="end_stamp"></field>
          <field var="duration" quote="false"></field>
          <field var="billsec" quote="false"></field>
          <field var="hangup_cause"></field>
          <field var="uuid"></field>
          <field var="bleg_uuid"></field>
          <field var="accountcode"></field>
          <field var="read_codec"></field>
          <field var="write_codec"></field>
  </schema>
  </configuration>
  <configuration name="cdr_sqlite.conf" description="SQLite CDR">
      <settings>
          <param name="legs" value="a"></param>
          <param name="default-template" value="example"></param>
  </settings>
      <templates>
          <template name="example">"${caller_id_name}","${caller_id_number}","${destination_number}","${context}","${start_stamp}","${answer_stamp}","${end_stamp}",${duration},${billsec},"${hangup_cause}","${uuid}","${bleg_uuid}","${accountcode}"</template>
  </templates>
  </configuration>
  <configuration name="cepstral.conf" description="Cepstral TTS configuration">
      <settings>
          <param name="encoding" value="utf-8"></param>
  </settings>
  </configuration>
    <configuration name="cidlookup.conf" description="cidlookup Configuration">
      <settings>
          <param name="url" value="http://query.voipcnam.com/query.php?api_key=MYAPIKEY&amp;number=${caller_id_number}"></param>
          <param name="whitepages-apikey" value="MYAPIKEY"></param>
          <param name="cache" value="true"></param>
          <param name="cache-expire" value="86400"></param>
          <param name="odbc-dsn" value="phone:phone:phone"></param>
          <param name="sql" value="      SELECT name||' ('||type||')' AS name        FROM phonebook p JOIN numbers n ON p.id = n.phonebook_id       WHERE n.number='${caller_id_number}'        LIMIT 1       "></param>
          <param name="citystate-sql" value="      SELECT ratecenter||' '||state as name       FROM npa_nxx_company_ocn       WHERE npa = ${caller_id_number:1:3} AND nxx = ${caller_id_number:4:3}       LIMIT 1       "></param>
  </settings>
  </configuration>
  <configuration name="conference.conf" description="Audio Conference">
      <advertise>
          <room name="3001@46.229.223.143" status="FreeSWITCH"></room>
  </advertise>
  	
  	
      <caller-controls>
          <group name="default">
              <control action="mute" digits="0"></control>
              <control action="deaf mute" digits="*"></control>
              <control action="energy up" digits="9"></control>
              <control action="energy equ" digits="8"></control>
              <control action="energy dn" digits="7"></control>
              <control action="vol talk up" digits="3"></control>
              <control action="vol talk zero" digits="2"></control>
              <control action="vol talk dn" digits="1"></control>
              <control action="vol listen up" digits="6"></control>
              <control action="vol listen zero" digits="5"></control>
              <control action="vol listen dn" digits="4"></control>
              <control action="hangup" digits="#"></control>
    </group>
  </caller-controls>
      <profiles>
          <profile name="default">
              <param name="domain" value="46.229.223.143"></param>
              <param name="rate" value="8000"></param>
              <param name="interval" value="20"></param>
              <param name="energy-level" value="100"></param>
              <param name="muted-sound" value="conference/conf-muted.wav"></param>
              <param name="unmuted-sound" value="conference/conf-unmuted.wav"></param>
              <param name="alone-sound" value="conference/conf-alone.wav"></param>
              <param name="moh-sound" value="local_stream://moh"></param>
              <param name="enter-sound" value="tone_stream://%(200,0,500,600,700)"></param>
              <param name="exit-sound" value="tone_stream://%(500,0,300,200,100,50,25)"></param>
              <param name="kicked-sound" value="conference/conf-kicked.wav"></param>
              <param name="locked-sound" value="conference/conf-locked.wav"></param>
              <param name="is-locked-sound" value="conference/conf-is-locked.wav"></param>
              <param name="is-unlocked-sound" value="conference/conf-is-unlocked.wav"></param>
              <param name="pin-sound" value="conference/conf-pin.wav"></param>
              <param name="bad-pin-sound" value="conference/conf-bad-pin.wav"></param>
              <param name="caller-id-name" value="FreeSWITCH"></param>
              <param name="caller-id-number" value="0000000000"></param>
              <param name="comfort-noise" value="true"></param>
    </profile>
          <profile name="wideband">
              <param name="domain" value="46.229.223.143"></param>
              <param name="rate" value="16000"></param>
              <param name="interval" value="20"></param>
              <param name="energy-level" value="100"></param>
              <param name="muted-sound" value="conference/conf-muted.wav"></param>
              <param name="unmuted-sound" value="conference/conf-unmuted.wav"></param>
              <param name="alone-sound" value="conference/conf-alone.wav"></param>
              <param name="moh-sound" value="local_stream://moh"></param>
              <param name="enter-sound" value="tone_stream://%(200,0,500,600,700)"></param>
              <param name="exit-sound" value="tone_stream://%(500,0,300,200,100,50,25)"></param>
              <param name="kicked-sound" value="conference/conf-kicked.wav"></param>
              <param name="locked-sound" value="conference/conf-locked.wav"></param>
              <param name="is-locked-sound" value="conference/conf-is-locked.wav"></param>
              <param name="is-unlocked-sound" value="conference/conf-is-unlocked.wav"></param>
              <param name="pin-sound" value="conference/conf-pin.wav"></param>
              <param name="bad-pin-sound" value="conference/conf-bad-pin.wav"></param>
              <param name="caller-id-name" value="FreeSWITCH"></param>
              <param name="caller-id-number" value="0000000000"></param>
              <param name="comfort-noise" value="true"></param>
    </profile>
          <profile name="ultrawideband">
              <param name="domain" value="46.229.223.143"></param>
              <param name="rate" value="32000"></param>
              <param name="interval" value="20"></param>
              <param name="energy-level" value="100"></param>
              <param name="muted-sound" value="conference/conf-muted.wav"></param>
              <param name="unmuted-sound" value="conference/conf-unmuted.wav"></param>
              <param name="alone-sound" value="conference/conf-alone.wav"></param>
              <param name="moh-sound" value="local_stream://moh"></param>
              <param name="enter-sound" value="tone_stream://%(200,0,500,600,700)"></param>
              <param name="exit-sound" value="tone_stream://%(500,0,300,200,100,50,25)"></param>
              <param name="kicked-sound" value="conference/conf-kicked.wav"></param>
              <param name="locked-sound" value="conference/conf-locked.wav"></param>
              <param name="is-locked-sound" value="conference/conf-is-locked.wav"></param>
              <param name="is-unlocked-sound" value="conference/conf-is-unlocked.wav"></param>
              <param name="pin-sound" value="conference/conf-pin.wav"></param>
              <param name="bad-pin-sound" value="conference/conf-bad-pin.wav"></param>
              <param name="caller-id-name" value="FreeSWITCH"></param>
              <param name="caller-id-number" value="0000000000"></param>
              <param name="comfort-noise" value="true"></param>
    </profile>
          <profile name="cdquality">
              <param name="domain" value="46.229.223.143"></param>
              <param name="rate" value="48000"></param>
              <param name="interval" value="20"></param>
              <param name="energy-level" value="100"></param>
              <param name="muted-sound" value="conference/conf-muted.wav"></param>
              <param name="unmuted-sound" value="conference/conf-unmuted.wav"></param>
              <param name="alone-sound" value="conference/conf-alone.wav"></param>
              <param name="moh-sound" value="local_stream://moh"></param>
              <param name="enter-sound" value="tone_stream://%(200,0,500,600,700)"></param>
              <param name="exit-sound" value="tone_stream://%(500,0,300,200,100,50,25)"></param>
              <param name="kicked-sound" value="conference/conf-kicked.wav"></param>
              <param name="locked-sound" value="conference/conf-locked.wav"></param>
              <param name="is-locked-sound" value="conference/conf-is-locked.wav"></param>
              <param name="is-unlocked-sound" value="conference/conf-is-unlocked.wav"></param>
              <param name="pin-sound" value="conference/conf-pin.wav"></param>
              <param name="bad-pin-sound" value="conference/conf-bad-pin.wav"></param>
              <param name="caller-id-name" value="FreeSWITCH"></param>
              <param name="caller-id-number" value="0000000000"></param>
              <param name="comfort-noise" value="true"></param>
    </profile>
          <profile name="video-mcu-stereo">
              <param name="domain" value="46.229.223.143"></param>
              <param name="rate" value="48000"></param>
              <param name="channels" value="2"></param>
              <param name="interval" value="20"></param>
              <param name="energy-level" value="200"></param>
              <param name="muted-sound" value="conference/conf-muted.wav"></param>
              <param name="unmuted-sound" value="conference/conf-unmuted.wav"></param>
              <param name="alone-sound" value="conference/conf-alone.wav"></param>
              <param name="moh-sound" value="local_stream://moh"></param>
              <param name="enter-sound" value="tone_stream://%(200,0,500,600,700)"></param>
              <param name="exit-sound" value="tone_stream://%(500,0,300,200,100,50,25)"></param>
              <param name="kicked-sound" value="conference/conf-kicked.wav"></param>
              <param name="locked-sound" value="conference/conf-locked.wav"></param>
              <param name="is-locked-sound" value="conference/conf-is-locked.wav"></param>
              <param name="is-unlocked-sound" value="conference/conf-is-unlocked.wav"></param>
              <param name="pin-sound" value="conference/conf-pin.wav"></param>
              <param name="bad-pin-sound" value="conference/conf-bad-pin.wav"></param>
              <param name="caller-id-name" value="FreeSWITCH"></param>
              <param name="caller-id-number" value="0000000000"></param>
              <param name="comfort-noise" value="false"></param>
              <param name="conference-flags" value="video-floor-only|rfc-4579|livearray-sync|minimize-video-encoding"></param>
              <param name="video-mode" value="mux"></param>
              <param name="video-layout-name" value="3x3"></param>
              <param name="video-layout-name" value="group:grid"></param>
              <param name="video-canvas-size" value="1920x1080"></param>
              <param name="video-canvas-bgcolor" value="#333333"></param>
              <param name="video-layout-bgcolor" value="#000000"></param>
              <param name="video-codec-bandwidth" value="1mb"></param>
              <param name="video-fps" value="15"></param>
    </profile>
          <profile name="sla">
              <param name="domain" value="46.229.223.143"></param>
              <param name="rate" value="16000"></param>
              <param name="interval" value="20"></param>
              <param name="caller-controls" value="none"></param>
              <param name="energy-level" value="200"></param>
              <param name="moh-sound" value="silence"></param>
              <param name="comfort-noise" value="true"></param>
    </profile>
  </profiles>
  </configuration>
  <configuration name="conference_layouts.conf" description="Audio Conference">
      <layout-settings>
          <layouts>
              <layout name="1x1">
	          <image x="0" y="0" scale="360" floor="true"></image>
      </layout>
              <layout name="1x2" auto-3d-position="true">
	          <image x="90" y="0" scale="180"></image>
	          <image x="90" y="180" scale="180"></image>
      </layout>
              <layout name="2x1" auto-3d-position="true">
	          <image x="0" y="90" scale="180"></image>
	          <image x="180" y="90" scale="180"></image>
      </layout>
              <layout name="2x1-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="180" hscale="360" zoom="true"></image>
	          <image x="180" y="0" scale="180" hscale="360" zoom="true"></image>
      </layout>
              <layout name="3x1-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="120" hscale="360" zoom="true"></image>
	          <image x="120" y="0" scale="120" hscale="360" zoom="true"></image>
	          <image x="240" y="0" scale="120" hscale="360" zoom="true"></image>
      </layout>
              <layout name="5-grid-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="180"></image>
	          <image x="180" y="0" scale="180"></image>
	          <image x="0" y="180" scale="120" hscale="180" zoom="true"></image>
	          <image x="120" y="180" scale="120" hscale="180" zoom="true"></image>
	          <image x="240" y="180" scale="120" hscale="180" zoom="true"></image>
      </layout>
              <layout name="3x2-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="120" hscale="180" zoom="true"></image>
	          <image x="120" y="0" scale="120" hscale="180" zoom="true"></image>
	          <image x="240" y="0" scale="120" hscale="180" zoom="true"></image>
	          <image x="0" y="180" scale="120" hscale="180" zoom="true"></image>
	          <image x="120" y="180" scale="120" hscale="180" zoom="true"></image>
	          <image x="240" y="180" scale="120" hscale="180" zoom="true"></image>
      </layout>
              <layout name="7-grid-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="120" hscale="180" zoom="true"></image>
	          <image x="120" y="0" scale="120" hscale="180" zoom="true"></image>
	          <image x="240" y="0" scale="120" hscale="180" zoom="true"></image>
	          <image x="0" y="180" scale="90" hscale="180" zoom="true"></image>
	          <image x="90" y="180" scale="90" hscale="180" zoom="true"></image>
	          <image x="180" y="180" scale="90" hscale="180" zoom="true"></image>
	          <image x="270" y="180" scale="90" hscale="180" zoom="true"></image>
      </layout>
              <layout name="4x2-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="90" hscale="180" zoom="true"></image>
	          <image x="90" y="0" scale="90" hscale="180" zoom="true"></image>
	          <image x="180" y="0" scale="90" hscale="180" zoom="true"></image>
	          <image x="270" y="0" scale="90" hscale="180" zoom="true"></image>
	          <image x="0" y="180" scale="90" hscale="180" zoom="true"></image>
	          <image x="90" y="180" scale="90" hscale="180" zoom="true"></image>
	          <image x="180" y="180" scale="90" hscale="180" zoom="true"></image>
	          <image x="270" y="180" scale="90" hscale="180" zoom="true"></image>
      </layout>
              <layout name="1x1+2x1" auto-3d-position="true">
	          <image x="90" y="0" scale="180"></image>
	          <image x="0" y="180" scale="180"></image>
	          <image x="180" y="180" scale="180"></image>
      </layout>
              <layout name="2x2" auto-3d-position="true">
	          <image x="0" y="0" scale="180"></image>
	          <image x="180" y="0" scale="180"></image>
	          <image x="0" y="180" scale="180"></image>
	          <image x="180" y="180" scale="180"></image>
      </layout>
              <layout name="3x3" auto-3d-position="true">
	          <image x="0" y="0" scale="120"></image>
	          <image x="120" y="0" scale="120"></image>
	          <image x="240" y="0" scale="120"></image>
	          <image x="0" y="120" scale="120"></image>
	          <image x="120" y="120" scale="120"></image>
	          <image x="240" y="120" scale="120"></image>
	          <image x="0" y="240" scale="120"></image>
	          <image x="120" y="240" scale="120"></image>
	          <image x="240" y="240" scale="120"></image>
      </layout>
              <layout name="4x4" auto-3d-position="true">
	          <image x="0" y="0" scale="90"></image>
	          <image x="90" y="0" scale="90"></image>
	          <image x="180" y="0" scale="90"></image>
	          <image x="270" y="0" scale="90"></image>
	          <image x="0" y="90" scale="90"></image>
	          <image x="90" y="90" scale="90"></image>
	          <image x="180" y="90" scale="90"></image>
	          <image x="270" y="90" scale="90"></image>
	          <image x="0" y="180" scale="90"></image>
	          <image x="90" y="180" scale="90"></image>
	          <image x="180" y="180" scale="90"></image>
	          <image x="270" y="180" scale="90"></image>
	          <image x="0" y="270" scale="90"></image>
	          <image x="90" y="270" scale="90"></image>
	          <image x="180" y="270" scale="90"></image>
	          <image x="270" y="270" scale="90"></image>
      </layout>
              <layout name="5x5" auto-3d-position="true">
	          <image x="0" y="0" scale="72"></image>
	          <image x="72" y="0" scale="72"></image>
	          <image x="144" y="0" scale="72"></image>
	          <image x="216" y="0" scale="72"></image>
	          <image x="288" y="0" scale="72"></image>
	          <image x="0" y="72" scale="72"></image>
	          <image x="72" y="72" scale="72"></image>
	          <image x="144" y="72" scale="72"></image>
	          <image x="216" y="72" scale="72"></image>
	          <image x="288" y="72" scale="72"></image>
	          <image x="0" y="144" scale="72"></image>
	          <image x="72" y="144" scale="72"></image>
	          <image x="144" y="144" scale="72"></image>
	          <image x="216" y="144" scale="72"></image>
	          <image x="288" y="144" scale="72"></image>
	          <image x="0" y="216" scale="72"></image>
	          <image x="72" y="216" scale="72"></image>
	          <image x="144" y="216" scale="72"></image>
	          <image x="216" y="216" scale="72"></image>
	          <image x="288" y="216" scale="72"></image>
	          <image x="0" y="288" scale="72"></image>
	          <image x="72" y="288" scale="72"></image>
	          <image x="144" y="288" scale="72"></image>
	          <image x="216" y="288" scale="72"></image>
	          <image x="288" y="288" scale="72"></image>
      </layout>
              <layout name="6x6" auto-3d-position="true">
	          <image x="0" y="0" scale="60"></image>
	          <image x="60" y="0" scale="60"></image>
	          <image x="120" y="0" scale="60"></image>
	          <image x="180" y="0" scale="60"></image>
	          <image x="240" y="0" scale="60"></image>
	          <image x="300" y="0" scale="60"></image>
	          <image x="0" y="60" scale="60"></image>
	          <image x="60" y="60" scale="60"></image>
	          <image x="120" y="60" scale="60"></image>
	          <image x="180" y="60" scale="60"></image>
	          <image x="240" y="60" scale="60"></image>
	          <image x="300" y="60" scale="60"></image>
	          <image x="0" y="120" scale="60"></image>
	          <image x="60" y="120" scale="60"></image>
	          <image x="120" y="120" scale="60"></image>
	          <image x="180" y="120" scale="60"></image>
	          <image x="240" y="120" scale="60"></image>
	          <image x="300" y="120" scale="60"></image>
	          <image x="0" y="180" scale="60"></image>
	          <image x="60" y="180" scale="60"></image>
	          <image x="120" y="180" scale="60"></image>
	          <image x="180" y="180" scale="60"></image>
	          <image x="240" y="180" scale="60"></image>
	          <image x="300" y="180" scale="60"></image>
	          <image x="0" y="240" scale="60"></image>
	          <image x="60" y="240" scale="60"></image>
	          <image x="120" y="240" scale="60"></image>
	          <image x="180" y="240" scale="60"></image>
	          <image x="240" y="240" scale="60"></image>
	          <image x="300" y="240" scale="60"></image>
	          <image x="0" y="300" scale="60"></image>
	          <image x="60" y="300" scale="60"></image>
	          <image x="120" y="300" scale="60"></image>
	          <image x="180" y="300" scale="60"></image>
	          <image x="240" y="300" scale="60"></image>
	          <image x="300" y="300" scale="60"></image>
      </layout>
              <layout name="8x8" auto-3d-position="true">
	          <image x="0" y="0" scale="45"></image>
	          <image x="45" y="0" scale="45"></image>
	          <image x="90" y="0" scale="45"></image>
	          <image x="135" y="0" scale="45"></image>
	          <image x="180" y="0" scale="45"></image>
	          <image x="225" y="0" scale="45"></image>
	          <image x="270" y="0" scale="45"></image>
	          <image x="315" y="0" scale="45"></image>
	          <image x="0" y="45" scale="45"></image>
	          <image x="45" y="45" scale="45"></image>
	          <image x="90" y="45" scale="45"></image>
	          <image x="135" y="45" scale="45"></image>
	          <image x="180" y="45" scale="45"></image>
	          <image x="225" y="45" scale="45"></image>
	          <image x="270" y="45" scale="45"></image>
	          <image x="315" y="45" scale="45"></image>
	          <image x="0" y="90" scale="45"></image>
	          <image x="45" y="90" scale="45"></image>
	          <image x="90" y="90" scale="45"></image>
	          <image x="135" y="90" scale="45"></image>
	          <image x="180" y="90" scale="45"></image>
	          <image x="225" y="90" scale="45"></image>
	          <image x="270" y="90" scale="45"></image>
	          <image x="315" y="90" scale="45"></image>
	          <image x="0" y="135" scale="45"></image>
	          <image x="45" y="135" scale="45"></image>
	          <image x="90" y="135" scale="45"></image>
	          <image x="135" y="135" scale="45"></image>
	          <image x="180" y="135" scale="45"></image>
	          <image x="225" y="135" scale="45"></image>
	          <image x="270" y="135" scale="45"></image>
	          <image x="315" y="135" scale="45"></image>
	          <image x="0" y="180" scale="45"></image>
	          <image x="45" y="180" scale="45"></image>
	          <image x="90" y="180" scale="45"></image>
	          <image x="135" y="180" scale="45"></image>
	          <image x="180" y="180" scale="45"></image>
	          <image x="225" y="180" scale="45"></image>
	          <image x="270" y="180" scale="45"></image>
	          <image x="315" y="180" scale="45"></image>
	          <image x="0" y="225" scale="45"></image>
	          <image x="45" y="225" scale="45"></image>
	          <image x="90" y="225" scale="45"></image>
	          <image x="135" y="225" scale="45"></image>
	          <image x="180" y="225" scale="45"></image>
	          <image x="225" y="225" scale="45"></image>
	          <image x="270" y="225" scale="45"></image>
	          <image x="315" y="225" scale="45"></image>
	          <image x="0" y="270" scale="45"></image>
	          <image x="45" y="270" scale="45"></image>
	          <image x="90" y="270" scale="45"></image>
	          <image x="135" y="270" scale="45"></image>
	          <image x="180" y="270" scale="45"></image>
	          <image x="225" y="270" scale="45"></image>
	          <image x="270" y="270" scale="45"></image>
	          <image x="315" y="270" scale="45"></image>
	          <image x="0" y="315" scale="45"></image>
	          <image x="45" y="315" scale="45"></image>
	          <image x="90" y="315" scale="45"></image>
	          <image x="135" y="315" scale="45"></image>
	          <image x="180" y="315" scale="45"></image>
	          <image x="225" y="315" scale="45"></image>
	          <image x="270" y="315" scale="45"></image>
	          <image x="315" y="315" scale="45"></image>
      </layout>
              <layout name="1up_top_left+5" auto-3d-position="true">
	          <image x="0" y="0" scale="240" floor="true"></image>
	          <image x="240" y="0" scale="120"></image>
	          <image x="240" y="120" scale="120"></image>
	          <image x="0" y="240" scale="120"></image>
	          <image x="120" y="240" scale="120"></image>
	          <image x="240" y="240" scale="120"></image>
      </layout>
              <layout name="1up_top_left+7" auto-3d-position="true">
	          <image x="0" y="0" scale="270" floor="true"></image>
	          <image x="270" y="0" scale="90"></image>
	          <image x="270" y="90" scale="90"></image>
	          <image x="270" y="180" scale="90"></image>
	          <image x="0" y="270" scale="90"></image>
	          <image x="90" y="270" scale="90"></image>
	          <image x="180" y="270" scale="90"></image>
	          <image x="270" y="270" scale="90"></image>
      </layout>
              <layout name="1up_top_left+9" auto-3d-position="true">
	          <image x="0" y="0" scale="288" floor="true"></image>
	          <image x="288" y="0" scale="72"></image>
	          <image x="288" y="72" scale="72"></image>
	          <image x="288" y="144" scale="72"></image>
	          <image x="288" y="216" scale="72"></image>
	          <image x="0" y="288" scale="72"></image>
	          <image x="72" y="288" scale="72"></image>
	          <image x="144" y="288" scale="72"></image>
	          <image x="216" y="288" scale="72"></image>
	          <image x="288" y="288" scale="72"></image>
      </layout>
              <layout name="2up_top+8" auto-3d-position="true">
	          <image x="0" y="0" scale="180" floor="true"></image>
	          <image x="180" y="0" scale="180" reservation_id="secondary"></image>
	          <image x="0" y="180" scale="90"></image>
	          <image x="90" y="180" scale="90"></image>
	          <image x="180" y="180" scale="90"></image>
	          <image x="270" y="180" scale="90"></image>
	          <image x="0" y="270" scale="90"></image>
	          <image x="90" y="270" scale="90"></image>
	          <image x="180" y="270" scale="90"></image>
	          <image x="270" y="270" scale="90"></image>
      </layout>
              <layout name="2up_middle+8" auto-3d-position="true">
	          <image x="0" y="90" scale="180" floor="true"></image>
	          <image x="180" y="90" scale="180" reservation_id="secondary"></image>
	          <image x="0" y="0" scale="90"></image>
	          <image x="90" y="0" scale="90"></image>
	          <image x="180" y="0" scale="90"></image>
	          <image x="270" y="0" scale="90"></image>
	          <image x="0" y="270" scale="90"></image>
	          <image x="90" y="270" scale="90"></image>
	          <image x="180" y="270" scale="90"></image>
	          <image x="270" y="270" scale="90"></image>
      </layout>
              <layout name="2up_bottom+8" auto-3d-position="true">
	          <image x="0" y="180" scale="180" floor="true"></image>
	          <image x="180" y="180" scale="180" reservation_id="secondary"></image>
	          <image x="0" y="0" scale="90"></image>
	          <image x="90" y="0" scale="90"></image>
	          <image x="180" y="0" scale="90"></image>
	          <image x="270" y="0" scale="90"></image>
	          <image x="0" y="90" scale="90"></image>
	          <image x="90" y="90" scale="90"></image>
	          <image x="180" y="90" scale="90"></image>
	          <image x="270" y="90" scale="90"></image>
      </layout>
              <layout name="3up+4" auto-3d-position="true">
	          <image x="0" y="0" scale="180" floor="true"></image>
	          <image x="180" y="0" scale="180" reservation_id="secondary"></image>
	          <image x="0" y="180" scale="180" reservation_id="third"></image>
	          <image x="180" y="180" scale="90"></image>
	          <image x="270" y="180" scale="90"></image>
	          <image x="180" y="270" scale="90"></image>
	          <image x="270" y="270" scale="90"></image>
      </layout>
              <layout name="3up+9" auto-3d-position="true">
	          <image x="0" y="0" scale="180" floor="true"></image>
	          <image x="180" y="0" scale="180" reservation_id="secondary"></image>
	          <image x="0" y="180" scale="180" reservation_id="third"></image>
	          <image x="180" y="180" scale="60"></image>
	          <image x="240" y="180" scale="60"></image>
	          <image x="300" y="180" scale="60"></image>
	          <image x="180" y="240" scale="60"></image>
	          <image x="240" y="240" scale="60"></image>
	          <image x="300" y="240" scale="60"></image>
	          <image x="180" y="300" scale="60"></image>
	          <image x="240" y="300" scale="60"></image>
	          <image x="300" y="300" scale="60"></image>
      </layout>
              <layout name="2x1-presenter-zoom" auto-3d-position="true">
	          <image x="0" y="0" scale="180" hscale="360" zoom="true" floor="true"></image>
	          <image x="180" y="0" scale="180" hscale="360" zoom="true" reservation_id="presenter"></image>
      </layout>
              <layout name="presenter-dual-vertical">
	          <image x="90" y="0" scale="180" floor-only="true"></image>
	          <image x="90" y="180" scale="180" reservation_id="presenter"></image>
      </layout>
              <layout name="presenter-dual-horizontal">
	          <image x="0" y="90" scale="180" floor-only="true"></image>
	          <image x="180" y="90" scale="180" reservation_id="presenter"></image>
      </layout>
              <layout name="presenter-overlap-small-top-right">
	          <image x="0" y="0" scale="360" floor-only="true"></image>
	          <image x="300" y="0" scale="60" overlap="true" reservation_id="presenter"></image>
      </layout>
              <layout name="presenter-overlap-small-bot-right">
	          <image x="0" y="0" scale="360" floor-only="true"></image>
	          <image x="300" y="300" scale="60" overlap="true" reservation_id="presenter"></image>
      </layout>
              <layout name="presenter-overlap-large-top-right">
	          <image x="0" y="0" scale="360" floor-only="true"></image>
	          <image x="180" y="0" scale="180" overlap="true" reservation_id="presenter"></image>
      </layout>
              <layout name="presenter-overlap-large-bot-right">
	          <image x="0" y="0" scale="360" floor-only="true"></image>
	          <image x="180" y="180" scale="180" overlap="true" reservation_id="presenter"></image>
      </layout>
              <layout name="overlaps" auto-3d-position="true">
	          <image x="0" y="0" scale="360" floor-only="true"></image>
	          <image x="300" y="300" scale="60" overlap="true"></image>
	          <image x="240" y="300" scale="60" overlap="true"></image>
	          <image x="180" y="300" scale="60" overlap="true"></image>
	          <image x="120" y="300" scale="60" overlap="true"></image>
	          <image x="60" y="300" scale="60" overlap="true"></image>
	          <image x="0" y="300" scale="60" overlap="true"></image>
      </layout>
    </layouts>
          <groups>
              <group name="grid">
	          <layout>1x1</layout>
	          <layout>2x1</layout>
	          <layout>1x1+2x1</layout>
	          <layout>2x2</layout>
	          <layout>3x3</layout>
	          <layout>4x4</layout>
	          <layout>5x5</layout>
	          <layout>6x6</layout>
	          <layout>8x8</layout>
      </group>
              <group name="grid-zoom">
	          <layout>1x1</layout>
	          <layout>2x1-zoom</layout>
	          <layout>3x1-zoom</layout>
	          <layout>2x2</layout>
	          <layout>5-grid-zoom</layout>
	          <layout>3x2-zoom</layout>
	          <layout>7-grid-zoom</layout>
	          <layout>4x2-zoom</layout>
	          <layout>3x3</layout>
      </group>
              <group name="1up_top_left_plus">
	          <layout>1up_top_left+5</layout>
	          <layout>1up_top_left+7</layout>
	          <layout>1up_top_left+9</layout>
      </group>
              <group name="3up_plus">
	          <layout>3up+4</layout>
	          <layout>3up+9</layout>
      </group>
    </groups>
  </layout-settings>
  </configuration>
  <configuration name="console.conf" description="Console Logger">
      <mappings>
          <map name="all" value="console,debug,info,notice,warning,err,crit,alert"></map>
  </mappings>
      <settings>
          <param name="colorize" value="true"></param>
          <param name="loglevel" value="info"></param>
  </settings>
  </configuration>
  <configuration name="db.conf" description="LIMIT DB Configuration">
      <settings>
  </settings>
  </configuration>
  <configuration name="dialplan_directory.conf" description="Dialplan Directory">
      <settings>
          <param name="directory-name" value="ldap"></param>
          <param name="host" value="ldap.freeswitch.org"></param>
          <param name="dn" value="cn=Manager,dc=freeswitch,dc=org"></param>
          <param name="pass" value="test"></param>
          <param name="base" value="dc=freeswitch,dc=org"></param>
  </settings>
  </configuration>
  <configuration name="dingaling.conf" description="XMPP Jingle Endpoint">
      <settings>
          <param name="debug" value="0"></param>
          <param name="codec-prefs" value="H264,PCMU"></param>
  </settings>
      <x-profile type="client">
          <param name="name" value="xmppc"></param>
          <param name="login" value="myjid@myserver.com/talk"></param>
          <param name="password" value="mypass"></param>
          <param name="dialplan" value="XML"></param>
          <param name="context" value="public"></param>
          <param name="message" value="Jingle all the way"></param>
          <param name="rtp-ip" value="auto"></param>
          <param name="auto-login" value="true"></param>
          <param name="sasl" value="plain"></param>
          <param name="tls" value="true"></param>
          <param name="use-rtp-timer" value="true"></param>
          <param name="exten" value="888"></param>
          <param name="local-network-acl" value="localnet.auto"></param>
  </x-profile>
      <x-profile type="component">
          <param name="name" value="xmpps"></param>
          <param name="password" value="secret"></param>
          <param name="dialplan" value="XML"></param>
          <param name="context" value="public"></param>
          <param name="rtp-ip" value="auto"></param>
          <param name="server" value="jabber.server.org:5347"></param>
          <param name="use-rtp-timer" value="true"></param>
          <param name="exten" value="_auto_"></param>
  </x-profile>
  </configuration>
  <configuration name="directory.conf" description="Directory">
      <settings>
  </settings>
      <profiles>
          <profile name="default">
              <param name="max-menu-attempts" value="3"></param>
              <param name="min-search-digits" value="3"></param>
              <param name="terminator-key" value="#"></param>
              <param name="digit-timeout" value="3000"></param>
              <param name="max-result" value="5"></param>
              <param name="next-key" value="6"></param>
              <param name="prev-key" value="4"></param>
              <param name="switch-order-key" value="*"></param>
              <param name="select-name-key" value="1"></param>
              <param name="new-search-key" value="3"></param>
              <param name="search-order" value="last_name"></param>
    </profile>
  </profiles>
  </configuration>
  <configuration name="distributor.conf" description="Distributor Configuration">
      <lists>
          <list name="test">
              <node name="foo1" weight="1"></node>
              <node name="foo2" weight="9"></node>
    </list>
  </lists>
  </configuration>
  <configuration name="easyroute.conf" description="EasyRoute Module">
      <settings>
          <param name="db-username" value="root"></param>
          <param name="db-password" value="password"></param>
          <param name="db-dsn" value="easyroute"></param>
          <param name="default-techprofile" value="sofia/default"></param>
          <param name="default-gateway" value="192.168.66.6"></param>
          <param name="odbc-retries" value="120"></param>
  </settings>
  </configuration>
  <configuration name="enum.conf" description="ENUM Module">
      <settings>
          <param name="default-root" value="e164.org"></param>
          <param name="default-isn-root" value="freenum.org"></param>
          <param name="auto-reload" value="true"></param>
          <param name="query-timeout-ms" value="200"></param>
          <param name="query-timeout-retry" value="2"></param>
          <param name="random-nameserver" value="false"></param>
  </settings>
      <routes>
          <route service="E2U+SIP" regex="sip:(.*)" replace="sofia/${use_profile}-ipv6/$1;transport=udp|sofia/${use_profile}/$1;transport=udp"></route>
          <route service="E2T+SIP" regex="sip:(.*)" replace="sofia/${use_profile}-ipv6/$1;transport=tcp|sofia/${use_profile}/$1;transport=tcp"></route>
          <route service="E2T+SIPS" regex="sip:(.*)" replace="sofia/${use_profile}-ipv6/$1;transport=tls|sofia/${use_profile}/$1;transport=tls"></route>
  </routes>
  </configuration>
  <configuration name="erlang_event.conf" description="Erlang Socket Client">
      <settings>
          <param name="listen-ip" value="0.0.0.0"></param>
          <param name="listen-port" value="8031"></param>
          <param name="nodename" value="freeswitch"></param>
          <param name="cookie" value="ClueCon"></param>
          <param name="shortname" value="true"></param>
  </settings>
  </configuration>
  <configuration name="event_multicast.conf" description="Multicast Event">
      <settings>
          <param name="address" value="225.1.1.1"></param>
          <param name="port" value="4242"></param>
          <param name="bindings" value="all"></param>
          <param name="ttl" value="1"></param>
  </settings>
  </configuration>
  <configuration name="event_socket.conf" description="Socket Client">
      <settings>
          <param name="nat-map" value="false"></param>
          <param name="listen-ip" value="::"></param>
          <param name="listen-port" value="8021"></param>
          <param name="password" value="ClueCon"></param>
  </settings>
  </configuration>
  <configuration name="fax.conf" description="FAX application configuration">
        <settings>
	      <param name="use-ecm" value="true"></param>
	      <param name="verbose" value="false"></param>
	      <param name="disable-v17" value="false"></param>
	      <param name="ident" value="SpanDSP Fax Ident"></param>
	      <param name="header" value="SpanDSP Fax Header"></param>
	      <param name="spool-dir" value="/tmp"></param>
	      <param name="file-prefix" value="faxrx"></param>
    </settings>
  </configuration>
  <configuration name="fifo.conf" description="FIFO Configuration">
      <settings>
          <param name="delete-all-outbound-member-on-startup" value="false"></param>
  </settings>
      <fifos>
          <fifo name="cool_fifo@46.229.223.143" importance="0">
    </fifo>
  </fifos>
  </configuration>
  <configuration name="format_cdr.conf" description="Multi Format CDR CURL logger">
     <profiles>
       <profile name="default">
          <settings>
              <param name="format" value="xml"></param>
              <param name="log-dir" value=""></param>
              <param name="log-b-leg" value="false"></param>
              <param name="prefix-a-leg" value="true"></param>
              <param name="encode" value="true"></param>
              <param name="encode-values" value="true"></param>
  </settings>
 </profile>
 </profiles>
  </configuration>
  <configuration name="graylog2.conf" description="Graylog2 Logger">
      <settings>
          <param name="server-host" value="192.168.0.69"></param>
          <param name="server-port" value="12201"></param>
          <param name="loglevel" value="warning"></param>
          <fields>
    </fields>
  </settings>
  </configuration>
  <configuration name="hash.conf" description="Hash Configuration">
      <remotes>
  </remotes>
  </configuration>
  <configuration name="hiredis.conf" description="mod_hiredis">
      <profiles>
          <profile name="default">
              <connections>
	          <connection name="primary">
	              <param name="hostname" value="172.18.101.101"></param>
	              <param name="password" value="redis"></param>
	              <param name="port" value="6379"></param>
	              <param name="timeout_ms" value="500"></param>
	</connection>
	          <connection name="secondary">
	              <param name="hostname" value="localhost"></param>
	              <param name="password" value="redis"></param>
	              <param name="port" value="6380"></param>
	              <param name="timeout_ms" value="500"></param>
	</connection>
      </connections>
              <params>
	          <param name="ignore-connect-fail" value="true"></param>
      </params>
    </profile>
  </profiles>
  </configuration>
  <configuration name="httapi.conf" description="HT-TAPI Hypertext Telephony API">
      <settings>
          <param name="debug" value="true"></param>
          <param name="file-not-found-expires" value="300"></param>
          <param name="file-cache-ttl" value="300"></param>
  </settings>
      <profiles>
          <profile name="default">
              <conference>
	          <param name="default-profile" value="default"></param>
      </conference>
              <dial>
	          <param name="context" value="default"></param>
	          <param name="dialplan" value="XML"></param>
      </dial>
              <permissions>
	          <permission name="set-params" value="true"></permission>
	          <permission name="set-vars" value="false">
	</permission>
	          <permission name="get-vars" value="false">
	</permission>
	          <permission name="extended-data" value="false"></permission>
	          <permission name="execute-apps" value="true">
	              <application-list default="deny">
	                  <application name="info"></application>
	                  <application name="hangup"></application>
	  </application-list>
	</permission>
	          <permission name="expand-vars-in-tag-body" value="false">
	</permission>
	          <permission name="dial" value="true"></permission>
	          <permission name="dial-set-context" value="false"></permission>
	          <permission name="dial-set-dialplan" value="false"></permission>
	          <permission name="dial-set-cid-name" value="false"></permission>
	          <permission name="dial-set-cid-number" value="false"></permission>
	          <permission name="dial-full-originate" value="false"></permission>
	          <permission name="conference" value="true"></permission>
	          <permission name="conference-set-profile" value="false"></permission>
      </permissions>
              <params>
	          <param name="gateway-url" value="http://www.freeswitch.org/api/index.cgi"></param>
      </params>
    </profile>
  </profiles>
  </configuration>
  <configuration name="http_cache.conf" description="HTTP GET cache">
      <settings>
          <param name="enable-file-formats" value="false"></param>
          <param name="max-urls" value="10000"></param>
          <param name="location" value="/var/cache/freeswitch"></param>
          <param name="default-max-age" value="86400"></param>
          <param name="prefetch-thread-count" value="8"></param>
          <param name="prefetch-queue-size" value="100"></param>
          <param name="ssl-cacert" value="/etc/freeswitch/tls/cacert.pem"></param>
          <param name="ssl-verifypeer" value="true"></param>
          <param name="ssl-verifyhost" value="true"></param>
  </settings>
  </configuration>
  <configuration name="ivr.conf" description="IVR menus">
      <menus>
        <menu name="demo_ivr" greet-long="phrase:demo_ivr_main_menu" greet-short="phrase:demo_ivr_main_menu_short" invalid-sound="ivr/ivr-that_was_an_invalid_entry.wav" exit-sound="voicemail/vm-goodbye.wav" confirm-macro="" confirm-key="" tts-engine="flite" tts-voice="rms" confirm-attempts="3" timeout="10000" inter-digit-timeout="2000" max-failures="3" max-timeouts="3" digit-len="4">
            <entry action="menu-exec-app" digits="1" param="bridge sofia/46.229.223.143/888@conference.freeswitch.org"></entry>
            <entry action="menu-exec-app" digits="2" param="transfer 9196 XML default"></entry>
            <entry action="menu-exec-app" digits="3" param="transfer 9664 XML default"></entry>
            <entry action="menu-exec-app" digits="4" param="transfer 9191 XML default"></entry>
            <entry action="menu-exec-app" digits="5" param="transfer 1234*256 enum"></entry>
            <entry action="menu-sub" digits="6" param="demo_ivr_submenu"></entry>
                  
            <entry action="menu-exec-app" digits="/^(10[01][0-9])$/" param="transfer $1 XML features"></entry>
            <entry action="menu-top" digits="9"></entry>
          
  </menu>
        <menu name="demo_ivr_submenu" greet-long="phrase:demo_ivr_sub_menu" greet-short="phrase:demo_ivr_sub_menu_short" invalid-sound="ivr/ivr-that_was_an_invalid_entry.wav" exit-sound="voicemail/vm-goodbye.wav" timeout="15000" max-failures="3" max-timeouts="3">
            <entry action="menu-top" digits="*"></entry>
   </menu>
        <menu name="new_demo_ivr" greet-long="phrase:new_demo_ivr_main_menu" greet-short="phrase:new_demo_ivr_main_menu_short" invalid-sound="ivr/ivr-that_was_an_invalid_entry.wav" exit-sound="voicemail/vm-goodbye.wav" confirm-macro="" confirm-key="" tts-engine="flite" tts-voice="rms" confirm-attempts="3" timeout="10000" inter-digit-timeout="2000" max-failures="3" max-timeouts="3" digit-len="4">
            <entry action="menu-sub" digits="1" param="freeswitch_ivr_submenu"></entry>
            
            <entry action="menu-sub" digits="2" param="freeswitch_solutions_ivr_submenu"></entry>
            <entry action="menu-sub" digits="3" param="cluecon_ivr_submenu"></entry>
               
            <entry action="menu-exec-app" digits="4" param="5000 XML default"></entry>
                  
            <entry action="menu-top" digits="9"></entry>
                                           
  </menu>
        <menu name="freeswitch_ivr_submenu" greet-long="phrase:learn_about_freeswitch_sub_menu" greet-short="phrase:learn_about_freeswitch_sub_menu" invalid-sound="ivr/ivr-that_was_an_invalid_entry.wav" exit-sound="voicemail/vm-goodbye.wav" timeout="15000" max-failures="3" max-timeouts="3">
            <entry action="menu-sub" digits="9" param="freeswitch_ivr_submenu"></entry>
            <entry action="menu-top" digits="*"></entry>
  </menu>
        <menu name="freeswitch_solutions_ivr_submenu" greet-long="phrase:learn_about_freeswitch_solutions_sub_menu" greet-short="phrase:learn_about_freeswitch_solutions_sub_menu" invalid-sound="ivr/ivr-that_was_an_invalid_entry.wav" exit-sound="voicemail/vm-goodbye.wav" timeout="15000" max-failures="3" max-timeouts="3">
            <entry action="menu-sub" digits="9" param="freeswitch_solutions_ivr_submenu"></entry>
            <entry action="menu-top" digits="*"></entry>
  </menu>
        <menu name="cluecon_ivr_submenu" greet-long="phrase:learn_about_cluecon_sub_menu" greet-short="phrase:learn_about_cluecon_sub_menu" invalid-sound="ivr/ivr-that_was_an_invalid_entry.wav" exit-sound="voicemail/vm-goodbye.wav" timeout="15000" max-failures="3" max-timeouts="3">
            <entry action="menu-sub" digits="9" param="cluecon_ivr_submenu"></entry>
            <entry action="menu-top" digits="*"></entry>
  </menu>
  </menus>
  </configuration>
  <configuration name="java.conf" description="Java Plug-Ins">
      <javavm path="/opt/jdk1.6.0_04/jre/lib/amd64/server/libjvm.so"></javavm>
      <options>
          <option value="-Djava.class.path=/usr/share/freeswitch/scripts/freeswitch.jar:/usr/share/freeswitch/scripts/example.jar"></option>
          <option value="-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=0.0.0.0:8000"></option>
  </options>
      <startup class="org/freeswitch/example/ApplicationLauncher" method="startup"></startup>
  </configuration>
  <configuration name="kazoo.conf" description="General purpose Erlang c-node produced to better fit the Kazoo project">
      <settings>
          <param name="listen-ip" value="0.0.0.0"></param>
          <param name="listen-port" value="8031"></param>
          <param name="cookie" value="change_me"></param>
          <param name="shortname" value="false"></param>
          <param name="nodename" value="freeswitch"></param>
          <param name="send-msg-batch-size" value="10"></param>
          <param name="receive-timeout" value="1"></param>
  </settings>
      <event-filter type="whitelist">
          <header name="Acquired-UUID"></header>
          <header name="action"></header>
          <header name="Action"></header>
          <header name="alt_event_type"></header>
          <header name="Answer-State"></header>
          <header name="Application"></header>
          <header name="Application-Data"></header>
          <header name="Application-Name"></header>
          <header name="Application-Response"></header>
          <header name="att_xfer_replaced_by"></header>
          <header name="Auth-Method"></header>
          <header name="Auth-Realm"></header>
          <header name="Auth-User"></header>
          <header name="Bridge-A-Unique-ID"></header>
          <header name="Bridge-B-Unique-ID"></header>
          <header name="Call-Direction"></header>
          <header name="Caller-Callee-ID-Name"></header>
          <header name="Caller-Callee-ID-Number"></header>
          <header name="Caller-Caller-ID-Name"></header>
          <header name="Caller-Caller-ID-Number"></header>
          <header name="Caller-Context"></header>
          <header name="Caller-Controls"></header>
          <header name="Caller-Destination-Number"></header>
          <header name="Caller-Dialplan"></header>
          <header name="Caller-Network-Addr"></header>
          <header name="Caller-Unique-ID"></header>
          <header name="Call-ID"></header>
          <header name="Channel-Call-State"></header>
          <header name="Channel-Call-UUID"></header>
          <header name="Channel-Presence-ID"></header>
          <header name="Channel-State"></header>
          <header name="Chat-Permissions"></header>
          <header name="Conference-Name"></header>
          <header name="Conference-Profile-Name"></header>
          <header name="Conference-Unique-ID"></header>
          <header name="Conference-Size"></header>
          <header name="New-ID"></header>
          <header name="Old-ID"></header>
          <header name="Detected-Tone"></header>
          <header name="dialog_state"></header>
          <header name="direction"></header>
          <header name="Distributed-From"></header>
          <header name="DTMF-Digit"></header>
          <header name="DTMF-Duration"></header>
          <header name="Event-Date-Timestamp"></header>
          <header name="Event-Name"></header>
          <header name="Event-Subclass"></header>
          <header name="Expires"></header>
          <header name="Ext-SIP-IP"></header>
          <header name="File"></header>
          <header name="FreeSWITCH-Hostname"></header>
          <header name="from"></header>
          <header name="Hunt-Destination-Number"></header>
          <header name="ip"></header>
          <header name="Message-Account"></header>
          <header name="metadata"></header>
          <header name="old_node_channel_uuid"></header>
          <header name="Other-Leg-Callee-ID-Name"></header>
          <header name="Other-Leg-Callee-ID-Number"></header>
          <header name="Other-Leg-Caller-ID-Name"></header>
          <header name="Other-Leg-Caller-ID-Number"></header>
          <header name="Other-Leg-Destination-Number"></header>
          <header name="Other-Leg-Direction"></header>
          <header name="Other-Leg-Unique-ID"></header>
          <header name="Participant-Type"></header>
          <header name="Path"></header>
          <header name="profile_name"></header>
          <header name="Profiles"></header>
          <header name="proto-specific-event-name"></header>
          <header name="Raw-Application-Data"></header>
          <header name="Resigning-UUID"></header>
          <header name="set"></header>
          <header name="sip_auto_answer"></header>
          <header name="sip_auth_method"></header>
          <header name="sip_from_host"></header>
          <header name="sip_from_user"></header>
          <header name="sip_to_host"></header>
          <header name="sip_to_user"></header>
          <header name="sub-call-id"></header>
          <header name="technology"></header>
          <header name="to"></header>
          <header name="Unique-ID"></header>
          <header name="URL"></header>
          <header name="variable_channel_is_moving"></header>
          <header name="variable_collected_digits"></header>
          <header name="variable_current_application"></header>
          <header name="variable_current_application_data"></header>
          <header name="variable_domain_name"></header>
          <header name="variable_effective_caller_id_name"></header>
          <header name="variable_effective_caller_id_number"></header>
          <header name="variable_fax_bad_rows"></header>
          <header name="variable_fax_document_total_pages"></header>
          <header name="variable_fax_document_transferred_pages"></header>
          <header name="variable_fax_ecm_used"></header>
          <header name="variable_fax_result_code"></header>
          <header name="variable_fax_result_text"></header>
          <header name="variable_fax_success"></header>
          <header name="variable_fax_transfer_rate"></header>
          <header name="variable_holding_uuid"></header>
          <header name="variable_hold_music"></header>
          <header name="variable_media_group_id"></header>
          <header name="variable_originate_disposition"></header>
          <header name="variable_playback_terminator_used"></header>
          <header name="variable_presence_id"></header>
          <header name="variable_record_ms"></header>
          <header name="variable_recovered"></header>
          <header name="variable_silence_hits_exhausted"></header>
          <header name="variable_sip_auth_realm"></header>
          <header name="variable_sip_from_host"></header>
          <header name="variable_sip_from_user"></header>
          <header name="variable_sip_h_X-AUTH-IP"></header>
          <header name="variable_sip_received_ip"></header>
          <header name="variable_sip_to_host"></header>
          <header name="variable_sip_to_user"></header>
          <header name="variable_sofia_profile_name"></header>
          <header name="variable_transfer_history"></header>
          <header name="variable_user_name"></header>
          <header name="variable_endpoint_disposition"></header>
          <header name="variable_originate_disposition"></header>
          <header name="variable_bridge_hangup_cause"></header>
          <header name="variable_hangup_cause"></header>
          <header name="variable_last_bridge_proto_specific_hangup_cause"></header>
          <header name="variable_proto_specific_hangup_cause"></header>
          <header name="VM-Call-ID"></header>
          <header name="VM-sub-call-id"></header>
          <header name="whistle_application_name"></header>
          <header name="whistle_application_response"></header>
          <header name="whistle_event_name"></header>
          <header name="sip_auto_answer_notify"></header>
          <header name="eavesdrop_group"></header>
          <header name="origination_caller_id_name"></header>
          <header name="origination_caller_id_number"></header>
          <header name="origination_callee_id_name"></header>
          <header name="origination_callee_id_number"></header>
          <header name="sip_auth_username"></header>
          <header name="sip_auth_password"></header>
          <header name="effective_caller_id_name"></header>
          <header name="effective_caller_id_number"></header>
          <header name="effective_callee_id_name"></header>
          <header name="effective_callee_id_number"></header>
          <header name="call-id"></header>
          <header name="profile-name"></header>
          <header name="from-user"></header>
          <header name="from-host"></header>
          <header name="presence-hosts"></header>
          <header name="contact"></header>
          <header name="rpid"></header>
          <header name="status"></header>
          <header name="expires"></header>
          <header name="to-user"></header>
          <header name="to-host"></header>
          <header name="network-ip"></header>
          <header name="network-port"></header>
          <header name="username"></header>
          <header name="realm"></header>
          <header name="user-agent"></header>
          <header name="Hangup-Cause"></header>
          <header name="Unique-ID"></header>
          <header name="variable_switch_r_sdp"></header>
          <header name="variable_sip_local_sdp_str"></header>
          <header name="variable_sip_to_uri"></header>
          <header name="variable_sip_from_uri"></header>
          <header name="variable_effective_caller_id_number"></header>
          <header name="Caller-Caller-ID-Number"></header>
          <header name="variable_effective_caller_id_name"></header>
          <header name="Caller-Caller-ID-Name"></header>
          <header name="Caller-Callee-ID-Name"></header>
          <header name="Caller-Callee-ID-Number"></header>
          <header name="Other-Leg-Unique-ID"></header>
          <header name="variable_sip_user_agent"></header>
          <header name="variable_duration"></header>
          <header name="variable_billsec"></header>
          <header name="variable_progresssec"></header>
          <header name="variable_progress_uepoch"></header>
          <header name="variable_progress_media_uepoch"></header>
          <header name="variable_start_uepoch"></header>
          <header name="variable_digits_dialed"></header>
          <header name="variable_sip_cid_type"></header>
          <header name="Hear"></header>
          <header name="Speak"></header>
          <header name="Video"></header>
          <header name="Talking"></header>
          <header name="Mute-Detect"></header>
          <header name="Member-ID"></header>
          <header name="Member-Type"></header>
          <header name="Energy-Level"></header>
          <header name="Current-Energy"></header>
          <header name="Floor"></header>
  </event-filter>
  </configuration>
  <configuration name="lcr.conf" description="LCR Configuration">
      <settings>
          <param name="odbc-dsn" value="freeswitch-mysql:freeswitch:Fr33Sw1tch"></param>
  </settings>
      <profiles>
          <profile name="default">
              <param name="id" value="0"></param>
              <param name="order_by" value="rate,quality,reliability"></param>
    </profile>
          <profile name="qual_rel">
              <param name="id" value="1"></param>
              <param name="order_by" value="quality,reliability"></param>
    </profile>
          <profile name="rel_qual">
              <param name="id" value="2"></param>
              <param name="order_by" value="reliability,quality"></param>
    </profile>
  </profiles>
  </configuration>
  <configuration name="local_stream.conf" description="stream files from local dir">
      <directory name="default" path="/usr/share/freeswitch/sounds/music/8000">
          <param name="rate" value="8000"></param>
          <param name="shuffle" value="true"></param>
          <param name="channels" value="1"></param>
          <param name="interval" value="20"></param>
          <param name="timer-name" value="soft"></param>
  </directory>
      <directory name="moh/8000" path="/usr/share/freeswitch/sounds/music/8000">
          <param name="rate" value="8000"></param>
          <param name="shuffle" value="true"></param>
          <param name="channels" value="1"></param>
          <param name="interval" value="20"></param>
          <param name="timer-name" value="soft"></param>
  </directory>
      <directory name="moh/16000" path="/usr/share/freeswitch/sounds/music/16000">
          <param name="rate" value="16000"></param>
          <param name="shuffle" value="true"></param>
          <param name="channels" value="1"></param>
          <param name="interval" value="20"></param>
          <param name="timer-name" value="soft"></param>
  </directory>
      <directory name="moh/32000" path="/usr/share/freeswitch/sounds/music/32000">
          <param name="rate" value="32000"></param>
          <param name="shuffle" value="true"></param>
          <param name="channels" value="1"></param>
          <param name="interval" value="20"></param>
          <param name="timer-name" value="soft"></param>
  </directory>
      <directory name="moh/48000" path="/usr/share/freeswitch/sounds/music/48000">
          <param name="rate" value="48000"></param>
          <param name="shuffle" value="true"></param>
          <param name="channels" value="1"></param>
          <param name="interval" value="10"></param>
          <param name="timer-name" value="soft"></param>
  </directory>
  </configuration>
  <configuration name="logfile.conf" description="File Logging">
      <settings>
         <param name="rotate-on-hup" value="true"></param>
  </settings>
      <profiles>
          <profile name="default">
              <settings>
                  <param name="rollover" value="1048576000"></param>
		          <param name="maximum-rotate" value="32"></param>
                  <param name="uuid" value="true"></param>
      </settings>
              <mappings>
	          <map name="all" value="console,debug,info,notice,warning,err,crit,alert"></map>
      </mappings>
    </profile>
  </profiles>
  </configuration>
  <configuration name="lua.conf" description="LUA Configuration">
      <settings>
  </settings>
  </configuration>
  <configuration name="memcache.conf" description="memcache Configuration">
      <settings>
          <param name="memcache-servers" value="localhost"></param>
  </settings>
  </configuration>
  <configuration name="modules.conf" description="Modules">
      <modules>
          <load module="mod_console"></load>
          <load module="mod_logfile"></load>
          <load module="mod_enum"></load>
          <load module="mod_xml_curl"></load>
          <load module="mod_cdr_csv"></load>
          <load module="mod_event_socket"></load>
          <load module="mod_sofia"></load>
          <load module="mod_loopback"></load>
          <load module="mod_rtc"></load>
          <load module="mod_verto"></load>
          <load module="mod_commands"></load>
          <load module="mod_conference"></load>
          <load module="mod_db"></load>
          <load module="mod_dptools"></load>
          <load module="mod_expr"></load>
          <load module="mod_fifo"></load>
          <load module="mod_hash"></load>
          <load module="mod_voicemail"></load>
          <load module="mod_esf"></load>
          <load module="mod_fsv"></load>
          <load module="mod_valet_parking"></load>
          <load module="mod_httapi"></load>
          <load module="mod_dialplan_xml"></load>
          <load module="mod_dialplan_asterisk"></load>
          <load module="mod_spandsp"></load>
          <load module="mod_g723_1"></load>
          <load module="mod_g729"></load>
          <load module="mod_amr"></load>
          <load module="mod_b64"></load>
          <load module="mod_opus"></load>
          <load module="mod_sndfile"></load>
          <load module="mod_native_file"></load>
          <load module="mod_png"></load>
          <load module="mod_local_stream"></load>
          <load module="mod_tone_stream"></load>
          <load module="mod_lua"></load>
          <load module="mod_say_en"></load>
  </modules>
  </configuration>
  <configuration name="mongo.conf">
      <settings>
          <param name="connection-string" value="mongodb://127.0.0.1:27017/?connectTimeoutMS=10000"></param>
  </settings>
  </configuration>
  <configuration name="nibblebill.conf" description="Nibble Billing">
      <settings>
          <param name="odbc-dsn" value="bandwidth.com"></param>
          <param name="db_table" value="accounts"></param>
          <param name="db_column_cash" value="cash"></param>
          <param name="db_column_account" value="id"></param>
          <param name="global_heartbeat" value="60"></param>
          <param name="lowbal_amt" value="5"></param>
          <param name="lowbal_action" value="play ding"></param>
          <param name="nobal_amt" value="0"></param>
          <param name="nobal_action" value="hangup"></param>
          <param name="percall_max_amt" value="100"></param>
          <param name="percall_action" value="hangup"></param>
  </settings>
  </configuration>
  <configuration name="opal.conf" description="Opal Endpoints">
       <settings>
            <param name="trace-level" value="3"></param>
            <param name="context" value="default"></param>
            <param name="dialplan" value="XML"></param>
            <param name="dtmf-type" value="signal"></param>
                   
            <param name="jitter-size" value="40,100"></param>
                 
           
            <param name="gk-address" value=""></param>
                        
            <param name="gk-identifer" value=""></param>
                      
            <param name="gk-interface" value="46.229.223.143"></param>
   </settings>
       <listeners>
            <listener name="default">
                 <param name="h323-ip" value="46.229.223.143"></param>
                 <param name="h323-port" value="1720"></param>
      </listener>
   </listeners>
  </configuration>
  <configuration name="opus.conf">
        <settings>
              <param name="use-vbr" value="1"></param>
              <param name="complexity" value="10"></param>
              <param name="keep-fec-enabled" value="1"></param>
              <param name="maxaveragebitrate" value="0"></param>
              <param name="maxplaybackrate" value="0"></param>
    </settings>
  </configuration>
  <configuration name="oreka.conf" description="Oreka Recorder configuration">
      <settings>
  </settings>
  </configuration>
  <configuration name="osp.conf" description="OSP Module Configuration">
	    <settings>
		      <param name="debug-info" value="disabled"></param>
		      <param name="log-level" value="info"></param>
		      <param name="crypto-hardware" value="disabled"></param>
		      <param name="sip" module="sofia" profile="external"></param>
		      <param name="default-protocol" value="sip"></param>
	</settings>
	    <profiles>
		      <profile name="default">
			        <param name="service-point-url" value="http://127.0.0.1:5045/osp"></param>
			        <param name="device-ip" value="127.0.0.1:5080"></param>
			        <param name="ssl-lifetime" value="300"></param>
			        <param name="http-max-connections" value="20"></param>
			        <param name="http-persistence" value="60"></param>
			        <param name="http-retry-delay" value="0"></param>
			        <param name="http-retry-limit" value="2"></param>
			        <param name="http-timeout" value="10000"></param>
			        <param name="work-mode" value="direct"></param>
			        <param name="service-type" value="voice"></param>
			        <param name="max-destinations" value="12"></param>
		</profile>
	</profiles>
  </configuration>
  <configuration name="perl.conf" description="PERL Configuration">
      <settings>
  </settings>
  </configuration>
  <configuration name="pocketsphinx.conf" description="PocketSphinx ASR Configuration">
      <settings>
          <param name="threshold" value="400"></param>
          <param name="silence-hits" value="25"></param>
          <param name="listen-hits" value="1"></param>
          <param name="auto-reload" value="true"></param>
  </settings>
  </configuration>
  <configuration name="portaudio.conf" description="Soundcard Endpoint">
      <settings>
          <param name="indev" value=""></param>
          <param name="outdev" value=""></param>
          <param name="hold-file" value="local_stream://moh"></param>
          <param name="dialplan" value="XML"></param>
          <param name="cid-name" value="FreeSWITCH"></param>
          <param name="cid-num" value="0000000000"></param>
          <param name="sample-rate" value="48000"></param>
          <param name="codec-ms" value="20"></param>
  </settings>
      <streams>
  	
  	      <stream name="usb1">
		        <param name="indev" value="#2"></param>
		        <param name="outdev" value="#2"></param>
		        <param name="sample-rate" value="48000"></param>
		        <param name="codec-ms" value="10"></param>
		        <param name="channels" value="2"></param>
  	</stream>
  	      <stream name="default">
		        <param name="indev" value="#0"></param>
		        <param name="outdev" value="#1"></param>
		        <param name="sample-rate" value="48000"></param>
		        <param name="codec-ms" value="10"></param>
		        <param name="channels" value="1"></param>
  	</stream>
  </streams>
      <endpoints>
  	      <endpoint name="default">
		        <param name="instream" value="default:0"></param>
		        <param name="outstream" value="default:0"></param>
	</endpoint>
  	      <endpoint name="usb1out-left">
		        <param name="outstream" value="usb1:0"></param>
	</endpoint>
  	      <endpoint name="usb1out-right">
		        <param name="outstream" value="usb1:1"></param>
	</endpoint>
  	      <endpoint name="usb1in-left">
		        <param name="instream" value="usb1:0"></param>
	</endpoint>
  	      <endpoint name="usb1in-right">
		        <param name="instream" value="usb1:1"></param>
	</endpoint>
  	      <endpoint name="usb1-left">
		        <param name="instream" value="usb1:0"></param>
		        <param name="outstream" value="usb1:0"></param>
	</endpoint>
  	      <endpoint name="usb1-right">
		        <param name="instream" value="usb1:1"></param>
		        <param name="outstream" value="usb1:1"></param>
	</endpoint>
  </endpoints>
  </configuration>
  <configuration name="post_load_modules.conf" description="Modules">
      <modules>
  </modules>
  </configuration>
  <configuration name="presence_map.conf" description="PRESENCE MAP">
      <domains>
          <domain name="46.229.223.143">
              <exten regex="3\d+" proto="conf"></exten>
    </domain>
  </domains>
  </configuration>
  <configuration name="python.conf" description="PYTHON Configuration">
      <settings>
  </settings>
  </configuration>
  <configuration name="redis.conf" description="mod_redis Configuration">
      <settings>
          <param name="host" value="localhost"></param>
          <param name="port" value="6379"></param>
          <param name="timeout" value="10000"></param>
  </settings>
  </configuration>
  <configuration name="rss.conf" description="RSS Parser">
      <feeds>
  </feeds>
  </configuration>
  <configuration name="rtmp.conf" description="RTMP Endpoint">
      <profiles>
	      <profile name="default">
		        <settings>
			          <param name="bind-address" value="0.0.0.0:1935"></param>
			          <param name="context" value="public"></param>
			          <param name="dialplan" value="XML"></param>
			          <param name="auth-calls" value="true"></param>
			          <param name="buffer-len" value="50"></param>
			          <param name="chunksize" value="512"></param>
		</settings>
	</profile>
  </profiles>
  </configuration>
  <configuration name="sangoma_codec.conf" description="Sangoma Codec Configuration">
	    <settings>
	</settings>
  </configuration>
  <configuration name="shout.conf" description="mod shout config">
      <settings>
  </settings>
  </configuration>
  <configuration name="skinny.conf" description="Skinny Endpoints">
      <profiles>
      <profile name="internal">
          <settings>
              <param name="domain" value="46.229.223.143"></param>
              <param name="ip" value="46.229.223.143"></param>
              <param name="port" value="2000"></param>
              <param name="patterns-dialplan" value="XML"></param>
              <param name="patterns-context" value="skinny-patterns"></param>
              <param name="dialplan" value="XML"></param>
              <param name="context" value="default"></param>
              <param name="keep-alive" value="60"></param>
              <param name="date-format" value="D/M/Y"></param>
              <param name="odbc-dsn" value=""></param>
              <param name="debug" value="4"></param>
              <param name="auto-restart" value="true"></param>
              <param name="digit-timeout" value="10000"></param>
  </settings>
          <soft-key-set-sets>
              <soft-key-set-set name="default">
                  <soft-key-set name="KeySetOnHook" value="SoftkeyNewcall,SoftkeyRedial"></soft-key-set>
                  <soft-key-set name="KeySetConnected" value="SoftkeyEndcall,SoftkeyHold,SoftkeyNewcall,SoftkeyTransfer"></soft-key-set>
                  <soft-key-set name="KeySetOnHold" value="SoftkeyNewcall,SoftkeyResume,SoftkeyEndcall"></soft-key-set>
                  <soft-key-set name="KeySetRingIn" value="SoftkeyAnswer,SoftkeyEndcall,SoftkeyNewcall"></soft-key-set>
                  <soft-key-set name="KeySetOffHook" value=",SoftkeyRedial,SoftkeyEndcall"></soft-key-set>
                  <soft-key-set name="KeySetConnectedWithTransfer" value="SoftkeyEndcall,SoftkeyHold,SoftkeyNewcall,SoftkeyTransfer"></soft-key-set>
                  <soft-key-set name="KeySetDigitsAfterDialingFirstDigit" value="SoftkeyBackspace,,SoftkeyEndcall"></soft-key-set>
                  <soft-key-set name="KeySetRingOut" value=",,SoftkeyEndcall,SoftkeyTransfer"></soft-key-set>
                  <soft-key-set name="KeySetOffHookWithFeatures" value=",SoftkeyRedial,SoftkeyEndcall"></soft-key-set>
                  <soft-key-set name="KeySetInUseHint" value="SoftkeyNewcall,SoftkeyRedial"></soft-key-set>
    </soft-key-set-set>
  </soft-key-set-sets>
          <device-types>
              <device-type id="Cisco ATA 186">
                    <param name="firmware-version" value="ATA030101SCCP04"></param>
    </device-type>
  </device-types>
      </profile>
  </profiles>
  </configuration>
  <configuration name="smpp.conf" description="SMPP client and server Gateway">
      <gateways>
          <gateway name="example.com">
              <params>
	          <param name="host" value="example.com"></param>
	          <param name="port" value="2775"></param>
	          <param name="debug" value="1"></param>
	          <param name="profile" value="default"></param>
	          <param name="system_id" value="username"></param>
	          <param name="password" value="password"></param>
	          <param name="system_type" value="remote_smpp"></param>
      </params>
    </gateway>
  </gateways>
  </configuration>
  <configuration name="sms_flowroute.conf" description="SMS_FLOWROUTE send configs">
      <profiles>
          <profile name="default">
              <params>
	          <param name="host" value="https://api.flowroute.com/v2/messages"></param>
	          <param name="debug" value="1"></param>
	          <param name="port" value="8090"></param>
	          <param name="access-key" value="ACCESS-KEY"></param>
	          <param name="secret-key" value="SECRET-KEY"></param>
      </params>
    </profile>
  </profiles>
  </configuration>
  <configuration name="sofia.conf" description="sofia Endpoint">
      <global_settings>
          <param name="log-level" value="0"></param>
          <param name="debug-presence" value="0"></param>
  </global_settings>
      <profiles>
      <profile name="external-ipv6">
          <gateways>
  </gateways>
          <aliases>
  </aliases>
          <domains>
  </domains>
          <settings>
              <param name="debug" value="0"></param>
              <param name="sip-trace" value="no"></param>
              <param name="sip-capture" value="no"></param>
              <param name="rfc2833-pt" value="101"></param>
              <param name="sip-port" value="5080"></param>
              <param name="dialplan" value="XML"></param>
              <param name="context" value="public"></param>
              <param name="dtmf-duration" value="2000"></param>
              <param name="inbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="outbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="hold-music" value="local_stream://moh"></param>
              <param name="rtp-timer-name" value="soft"></param>
              <param name="local-network-acl" value="localnet.auto"></param>
              <param name="manage-presence" value="false"></param>
              <param name="inbound-codec-negotiation" value="generous"></param>
              <param name="nonce-ttl" value="60"></param>
              <param name="auth-calls" value="false"></param>
              <param name="inbound-late-negotiation" value="true"></param>
              <param name="inbound-zrtp-passthru" value="true"></param>
              <param name="rtp-ip" value="::1"></param>
              <param name="sip-ip" value="::1"></param>
              <param name="rtp-timeout-sec" value="300"></param>
              <param name="rtp-hold-timeout-sec" value="1800"></param>
              <param name="tls" value="false"></param>
              <param name="tls-only" value="false"></param>
              <param name="tls-bind-params" value="transport=tls"></param>
              <param name="tls-sip-port" value="5081"></param>
              <param name="tls-passphrase" value=""></param>
              <param name="tls-verify-date" value="true"></param>
              <param name="tls-verify-policy" value="none"></param>
              <param name="tls-verify-depth" value="2"></param>
              <param name="tls-verify-in-subjects" value=""></param>
              <param name="tls-version" value="tlsv1,tlsv1.1,tlsv1.2"></param>
  </settings>
      </profile>
      <profile name="external">
          <gateways>
  </gateways>
          <aliases>
  </aliases>
          <domains>
              <domain name="all" alias="false" parse="true"></domain>
  </domains>
          <settings>
              <param name="debug" value="0"></param>
              <param name="sip-trace" value="no"></param>
              <param name="sip-capture" value="no"></param>
              <param name="rfc2833-pt" value="101"></param>
              <param name="sip-port" value="5080"></param>
              <param name="dialplan" value="XML"></param>
              <param name="context" value="public"></param>
              <param name="dtmf-duration" value="2000"></param>
              <param name="inbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="outbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="hold-music" value="local_stream://moh"></param>
              <param name="rtp-timer-name" value="soft"></param>
              <param name="local-network-acl" value="localnet.auto"></param>
              <param name="manage-presence" value="false"></param>
              <param name="inbound-codec-negotiation" value="generous"></param>
              <param name="nonce-ttl" value="60"></param>
              <param name="auth-calls" value="false"></param>
              <param name="inbound-late-negotiation" value="true"></param>
              <param name="inbound-zrtp-passthru" value="true"></param>
              <param name="rtp-ip" value="46.229.223.143"></param>
              <param name="sip-ip" value="46.229.223.143"></param>
              <param name="ext-rtp-ip" value="auto-nat"></param>
              <param name="ext-sip-ip" value="auto-nat"></param>
              <param name="rtp-timeout-sec" value="300"></param>
              <param name="rtp-hold-timeout-sec" value="1800"></param>
              <param name="tls" value="false"></param>
              <param name="tls-only" value="false"></param>
              <param name="tls-bind-params" value="transport=tls"></param>
              <param name="tls-sip-port" value="5081"></param>
              <param name="tls-passphrase" value=""></param>
              <param name="tls-verify-date" value="true"></param>
              <param name="tls-verify-policy" value="none"></param>
              <param name="tls-verify-depth" value="2"></param>
              <param name="tls-verify-in-subjects" value=""></param>
              <param name="tls-version" value="tlsv1,tlsv1.1,tlsv1.2"></param>
  </settings>
      </profile>
      <profile name="internal-ipv6">
          <settings>
              <param name="debug" value="0"></param>
              <param name="sip-trace" value="no"></param>
              <param name="context" value="public"></param>
              <param name="rfc2833-pt" value="101"></param>
              <param name="sip-port" value="5060"></param>
              <param name="dialplan" value="XML"></param>
              <param name="dtmf-duration" value="2000"></param>
              <param name="inbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="outbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="use-rtp-timer" value="true"></param>
              <param name="rtp-timer-name" value="soft"></param>
              <param name="rtp-ip" value="::1"></param>
              <param name="sip-ip" value="::1"></param>
              <param name="hold-music" value="local_stream://moh"></param>
              <param name="apply-inbound-acl" value="domains"></param>
              <param name="record-template" value="/var/lib/freeswitch/recordings/${caller_id_number}.${strftime(%Y-%m-%d-%H-%M-%S)}.wav"></param>
              <param name="manage-presence" value="true"></param>
              <param name="inbound-codec-negotiation" value="generous"></param>
              <param name="tls" value="false"></param>
              <param name="tls-bind-params" value="transport=tls"></param>
              <param name="tls-sip-port" value="5061"></param>
              <param name="tls-cert-dir" value=""></param>
              <param name="tls-version" value="tlsv1,tlsv1.1,tlsv1.2"></param>
              <param name="inbound-late-negotiation" value="true"></param>
              <param name="inbound-zrtp-passthru" value="true"></param>
              <param name="nonce-ttl" value="60"></param>
              <param name="auth-calls" value="true"></param>
              <param name="auth-all-packets" value="false"></param>
              <param name="rtp-timeout-sec" value="300"></param>
              <param name="rtp-hold-timeout-sec" value="1800"></param>
              <param name="force-register-domain" value="46.229.223.143"></param>
              <param name="force-register-db-domain" value="46.229.223.143"></param>
  </settings>
      </profile>
      <profile name="internal">
          <aliases>
  </aliases>
          <gateways>
  </gateways>
          <domains>
              <domain name="all" alias="true" parse="false"></domain>
  </domains>
          <settings>
              <param name="debug" value="0"></param>
              <param name="sip-trace" value="no"></param>
              <param name="sip-capture" value="no"></param>
              <param name="watchdog-enabled" value="no"></param>
              <param name="watchdog-step-timeout" value="30000"></param>
              <param name="watchdog-event-timeout" value="30000"></param>
              <param name="log-auth-failures" value="false"></param>
              <param name="forward-unsolicited-mwi-notify" value="false"></param>
              <param name="context" value="public"></param>
              <param name="rfc2833-pt" value="101"></param>
              <param name="sip-port" value="5060"></param>
              <param name="dialplan" value="XML"></param>
              <param name="dtmf-duration" value="2000"></param>
              <param name="inbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="outbound-codec-prefs" value="OPUS,G722,PCMU,PCMA,VP8"></param>
              <param name="rtp-timer-name" value="soft"></param>
              <param name="rtp-ip" value="46.229.223.143"></param>
              <param name="sip-ip" value="46.229.223.143"></param>
              <param name="hold-music" value="local_stream://moh"></param>
              <param name="apply-nat-acl" value="nat.auto"></param>
              <param name="apply-inbound-acl" value="domains"></param>
              <param name="local-network-acl" value="localnet.auto"></param>
              <param name="record-path" value="/var/lib/freeswitch/recordings"></param>
              <param name="record-template" value="${caller_id_number}.${target_domain}.${strftime(%Y-%m-%d-%H-%M-%S)}.wav"></param>
              <param name="manage-presence" value="true"></param>
              <param name="presence-hosts" value="46.229.223.143,46.229.223.143"></param>
              <param name="presence-privacy" value="false"></param>
              <param name="inbound-codec-negotiation" value="generous"></param>
              <param name="tls" value="false"></param>
              <param name="tls-only" value="false"></param>
              <param name="tls-bind-params" value="transport=tls"></param>
              <param name="tls-sip-port" value="5061"></param>
              <param name="tls-passphrase" value=""></param>
              <param name="tls-verify-date" value="true"></param>
              <param name="tls-verify-policy" value="none"></param>
              <param name="tls-verify-depth" value="2"></param>
              <param name="tls-verify-in-subjects" value=""></param>
              <param name="tls-version" value="tlsv1,tlsv1.1,tlsv1.2"></param>
              <param name="tls-ciphers" value="ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH"></param>
              <param name="inbound-late-negotiation" value="true"></param>
              <param name="inbound-zrtp-passthru" value="true"></param>
              <param name="nonce-ttl" value="60"></param>
              <param name="auth-calls" value="true"></param>
              <param name="inbound-reg-force-matching-username" value="true"></param>
              <param name="auth-all-packets" value="false"></param>
              <param name="ext-rtp-ip" value="auto-nat"></param>
              <param name="ext-sip-ip" value="auto-nat"></param>
              <param name="rtp-timeout-sec" value="300"></param>
              <param name="rtp-hold-timeout-sec" value="1800"></param>
              <param name="force-register-domain" value="46.229.223.143"></param>
              <param name="force-subscription-domain" value="46.229.223.143"></param>
              <param name="force-register-db-domain" value="46.229.223.143"></param>
              <param name="ws-binding" value=":5066"></param>
              <param name="wss-binding" value=":7443"></param>
              <param name="challenge-realm" value="auto_from"></param>
  </settings>
      </profile>
  </profiles>
  </configuration>
  <configuration name="spandsp.conf" description="SpanDSP config">
        <modem-settings>
            <param name="total-modems" value="0"></param>
            <param name="context" value="default"></param>
            <param name="dialplan" value="XML"></param>
            <param name="verbose" value="false"></param>
    </modem-settings>
        <fax-settings>
	      <param name="use-ecm" value="true"></param>
	      <param name="verbose" value="false"></param>
	      <param name="disable-v17" value="false"></param>
	      <param name="ident" value="SpanDSP Fax Ident"></param>
	      <param name="header" value="SpanDSP Fax Header"></param>
	      <param name="spool-dir" value="/tmp"></param>
	      <param name="file-prefix" value="faxrx"></param>
    </fax-settings>
        <descriptors>
           <descriptor name="1">
               <tone name="CED_TONE">
                   <element freq1="2100" freq2="0" min="700" max="0"></element>
       </tone>
               <tone name="SIT">
                   <element freq1="950" freq2="0" min="256" max="400"></element>
                   <element freq1="1400" freq2="0" min="256" max="400"></element>
                   <element freq1="1800" freq2="0" min="256" max="400"></element>
       </tone>
               <tone name="RING_TONE" description="North America ring">
                   <element freq1="440" freq2="480" min="1200" max="0"></element>
       </tone>
               <tone name="REORDER_TONE">
                   <element freq1="480" freq2="620" min="224" max="316"></element>
                   <element freq1="0" freq2="0" min="168" max="352"></element>
                   <element freq1="480" freq2="620" min="224" max="316"></element>
       </tone>
               <tone name="BUSY_TONE">
                   <element freq1="480" freq2="620" min="464" max="536"></element>
                   <element freq1="0" freq2="0" min="464" max="572"></element>
                   <element freq1="480" freq2="620" min="464" max="536"></element>
       </tone>
     </descriptor>
           <descriptor name="44">
               <tone name="CED_TONE">
                   <element freq1="2100" freq2="0" min="500" max="0"></element>
       </tone>
               <tone name="SIT">
                   <element freq1="950" freq2="0" min="256" max="400"></element>
                   <element freq1="1400" freq2="0" min="256" max="400"></element>
                   <element freq1="1800" freq2="0" min="256" max="400"></element>
       </tone>
               <tone name="REORDER_TONE">
                   <element freq1="400" freq2="0" min="368" max="416"></element>
                   <element freq1="0" freq2="0" min="336" max="368"></element>
                   <element freq1="400" freq2="0" min="256" max="288"></element>
                   <element freq1="0" freq2="0" min="512" max="544"></element>
       </tone>
               <tone name="BUSY_TONE">
                   <element freq1="400" freq2="0" min="352" max="384"></element>
                   <element freq1="0" freq2="0" min="352" max="384"></element>
                   <element freq1="400" freq2="0" min="352" max="384"></element>
                   <element freq1="0" freq2="0" min="352" max="384"></element>
       </tone>
     </descriptor>
           <descriptor name="49">
               <tone name="CED_TONE">
                   <element freq1="2100" freq2="0" min="500" max="0"></element>
       </tone>
               <tone name="SIT">
                   <element freq1="900" freq2="0" min="256" max="400"></element>
                   <element freq1="1400" freq2="0" min="256" max="400"></element>
                   <element freq1="1800" freq2="0" min="256" max="400"></element>
       </tone>
               <tone name="REORDER_TONE">
                   <element freq1="425" freq2="0" min="224" max="272"></element>
                   <element freq1="0" freq2="0" min="224" max="272"></element>
       </tone>
               <tone name="BUSY_TONE">
                   <element freq1="425" freq2="0" min="464" max="516"></element>
                   <element freq1="0" freq2="0" min="464" max="516"></element>
       </tone>
     </descriptor>
   </descriptors>
  </configuration>
  <configuration name="switch.conf" description="Core Configuration">
      <cli-keybindings>
          <key name="1" value="help"></key>
          <key name="2" value="status"></key>
          <key name="3" value="show channels"></key>
          <key name="4" value="show calls"></key>
          <key name="5" value="sofia status"></key>
          <key name="6" value="reloadxml"></key>
          <key name="7" value="console loglevel 0"></key>
          <key name="8" value="console loglevel 7"></key>
          <key name="9" value="sofia status profile internal"></key>
          <key name="10" value="sofia profile internal siptrace on"></key>
          <key name="11" value="sofia profile internal siptrace off"></key>
          <key name="12" value="version"></key>
  </cli-keybindings>
      <default-ptimes>
  </default-ptimes>
      <settings>
          <param name="colorize-console" value="true"></param>
          <param name="dialplan-timestamps" value="false"></param>
          <param name="max-db-handles" value="50"></param>
          <param name="db-handle-timeout" value="10"></param>
          <param name="max-sessions" value="1000"></param>
          <param name="sessions-per-second" value="30"></param>
          <param name="loglevel" value="debug"></param>
          <param name="mailer-app" value="sendmail"></param>
          <param name="mailer-app-args" value="-t"></param>
          <param name="dump-cores" value="yes"></param>
          <param name="rtp-enable-zrtp" value="false"></param>
  </settings>
  </configuration>
  <configuration name="syslog.conf" description="Syslog Logger">
      <settings>
          <param name="facility" value="user"></param>
          <param name="ident" value="freeswitch"></param>
          <param name="loglevel" value="warning"></param>
          <param name="uuid" value="true"></param>
  </settings>
  </configuration>
  <configuration name="timezones.conf" description="Timezones">
        <timezones>
	      <zone name="Africa/Abidjan" value="GMT0"></zone>
	      <zone name="Africa/Accra" value="GMT0"></zone>
	      <zone name="Africa/Addis_Ababa" value="EAT-3"></zone>
	      <zone name="Africa/Algiers" value="CET-1"></zone>
	      <zone name="Africa/Asmara" value="EAT-3"></zone>
	      <zone name="Africa/Asmera" value="EAT-3"></zone>
	      <zone name="Africa/Bamako" value="GMT0"></zone>
	      <zone name="Africa/Bangui" value="WAT-1"></zone>
	      <zone name="Africa/Banjul" value="GMT0"></zone>
	      <zone name="Africa/Bissau" value="GMT0"></zone>
	      <zone name="Africa/Blantyre" value="CAT-2"></zone>
	      <zone name="Africa/Brazzaville" value="WAT-1"></zone>
	      <zone name="Africa/Bujumbura" value="CAT-2"></zone>
	      <zone name="Africa/Cairo" value="EET-2"></zone>
	      <zone name="Africa/Casablanca" value="WET0WEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Africa/Ceuta" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Africa/Conakry" value="GMT0"></zone>
	      <zone name="Africa/Dakar" value="GMT0"></zone>
	      <zone name="Africa/Dar_es_Salaam" value="EAT-3"></zone>
	      <zone name="Africa/Djibouti" value="EAT-3"></zone>
	      <zone name="Africa/Douala" value="WAT-1"></zone>
	      <zone name="Africa/El_Aaiun" value="WET0WEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Africa/Freetown" value="GMT0"></zone>
	      <zone name="Africa/Gaborone" value="CAT-2"></zone>
	      <zone name="Africa/Harare" value="CAT-2"></zone>
	      <zone name="Africa/Johannesburg" value="SAST-2"></zone>
	      <zone name="Africa/Juba" value="EAT-3"></zone>
	      <zone name="Africa/Kampala" value="EAT-3"></zone>
	      <zone name="Africa/Khartoum" value="EAT-3"></zone>
	      <zone name="Africa/Kigali" value="CAT-2"></zone>
	      <zone name="Africa/Kinshasa" value="WAT-1"></zone>
	      <zone name="Africa/Lagos" value="WAT-1"></zone>
	      <zone name="Africa/Libreville" value="WAT-1"></zone>
	      <zone name="Africa/Lome" value="GMT0"></zone>
	      <zone name="Africa/Luanda" value="WAT-1"></zone>
	      <zone name="Africa/Lubumbashi" value="CAT-2"></zone>
	      <zone name="Africa/Lusaka" value="CAT-2"></zone>
	      <zone name="Africa/Malabo" value="WAT-1"></zone>
	      <zone name="Africa/Maputo" value="CAT-2"></zone>
	      <zone name="Africa/Maseru" value="SAST-2"></zone>
	      <zone name="Africa/Mbabane" value="SAST-2"></zone>
	      <zone name="Africa/Mogadishu" value="EAT-3"></zone>
	      <zone name="Africa/Monrovia" value="GMT0"></zone>
	      <zone name="Africa/Nairobi" value="EAT-3"></zone>
	      <zone name="Africa/Ndjamena" value="WAT-1"></zone>
	      <zone name="Africa/Niamey" value="WAT-1"></zone>
	      <zone name="Africa/Nouakchott" value="GMT0"></zone>
	      <zone name="Africa/Ouagadougou" value="GMT0"></zone>
	      <zone name="Africa/Porto-Novo" value="WAT-1"></zone>
	      <zone name="Africa/Sao_Tome" value="GMT0"></zone>
	      <zone name="Africa/Timbuktu" value="GMT0"></zone>
	      <zone name="Africa/Tripoli" value="EET-2"></zone>
	      <zone name="Africa/Tunis" value="CET-1"></zone>
	      <zone name="Africa/Windhoek" value="WAT-1WAST,M9.1.0,M4.1.0"></zone>
	      <zone name="America/Adak" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Anchorage" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Anguilla" value="AST4"></zone>
	      <zone name="America/Antigua" value="AST4"></zone>
	      <zone name="America/Araguaina" value="BRT3"></zone>
	      <zone name="America/Argentina/Buenos_Aires" value="ART3"></zone>
	      <zone name="America/Argentina/Catamarca" value="ART3"></zone>
	      <zone name="America/Argentina/ComodRivadavia" value="ART3"></zone>
	      <zone name="America/Argentina/Cordoba" value="ART3"></zone>
	      <zone name="America/Argentina/Jujuy" value="ART3"></zone>
	      <zone name="America/Argentina/La_Rioja" value="ART3"></zone>
	      <zone name="America/Argentina/Mendoza" value="ART3"></zone>
	      <zone name="America/Argentina/Rio_Gallegos" value="ART3"></zone>
	      <zone name="America/Argentina/Salta" value="ART3"></zone>
	      <zone name="America/Argentina/San_Juan" value="ART3"></zone>
	      <zone name="America/Argentina/San_Luis" value="ART3"></zone>
	      <zone name="America/Argentina/Tucuman" value="ART3"></zone>
	      <zone name="America/Argentina/Ushuaia" value="ART3"></zone>
	      <zone name="America/Aruba" value="AST4"></zone>
	      <zone name="America/Asuncion" value="PYT4PYST,M10.1.0/0,M3.4.0/0"></zone>
	      <zone name="America/Atikokan" value="EST5"></zone>
	      <zone name="America/Atka" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Bahia" value="BRT3"></zone>
	      <zone name="America/Bahia_Banderas" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Barbados" value="AST4"></zone>
	      <zone name="America/Belem" value="BRT3"></zone>
	      <zone name="America/Belize" value="CST6"></zone>
	      <zone name="America/Blanc-Sablon" value="AST4"></zone>
	      <zone name="America/Boa_Vista" value="AMT4"></zone>
	      <zone name="America/Bogota" value="COT5"></zone>
	      <zone name="America/Boise" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Buenos_Aires" value="ART3"></zone>
	      <zone name="America/Cambridge_Bay" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Campo_Grande" value="AMT4AMST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="America/Cancun" value="EST5"></zone>
	      <zone name="America/Caracas" value="VET4:30"></zone>
	      <zone name="America/Catamarca" value="ART3"></zone>
	      <zone name="America/Cayenne" value="GFT3"></zone>
	      <zone name="America/Cayman" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Chicago" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Chihuahua" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Coral_Harbour" value="EST5"></zone>
	      <zone name="America/Cordoba" value="ART3"></zone>
	      <zone name="America/Costa_Rica" value="CST6"></zone>
	      <zone name="America/Creston" value="MST7"></zone>
	      <zone name="America/Cuiaba" value="AMT4AMST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="America/Curacao" value="AST4"></zone>
	      <zone name="America/Danmarkshavn" value="GMT0"></zone>
	      <zone name="America/Dawson" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Dawson_Creek" value="MST7"></zone>
	      <zone name="America/Denver" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Detroit" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Dominica" value="AST4"></zone>
	      <zone name="America/Edmonton" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Eirunepe" value="ACT5"></zone>
	      <zone name="America/El_Salvador" value="CST6"></zone>
	      <zone name="America/Ensenada" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Fort_Nelson" value="MST7"></zone>
	      <zone name="America/Fort_Wayne" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Fortaleza" value="BRT3"></zone>
	      <zone name="America/Glace_Bay" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Godthab" value="WGST"></zone>
	      <zone name="America/Goose_Bay" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Grand_Turk" value="AST4"></zone>
	      <zone name="America/Grenada" value="AST4"></zone>
	      <zone name="America/Guadeloupe" value="AST4"></zone>
	      <zone name="America/Guatemala" value="CST6"></zone>
	      <zone name="America/Guayaquil" value="ECT5"></zone>
	      <zone name="America/Guyana" value="GYT4"></zone>
	      <zone name="America/Halifax" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Havana" value="CST5CDT,M3.2.0/0,M11.1.0/1"></zone>
	      <zone name="America/Hermosillo" value="MST7"></zone>
	      <zone name="America/Indiana/Indianapolis" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Knox" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Marengo" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Petersburg" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Tell_City" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Vevay" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Vincennes" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indiana/Winamac" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Indianapolis" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Inuvik" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Iqaluit" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Jamaica" value="EST5"></zone>
	      <zone name="America/Jujuy" value="ART3"></zone>
	      <zone name="America/Juneau" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Kentucky/Louisville" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Kentucky/Monticello" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Knox_IN" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Kralendijk" value="AST4"></zone>
	      <zone name="America/La_Paz" value="BOT4"></zone>
	      <zone name="America/Lima" value="PET5"></zone>
	      <zone name="America/Los_Angeles" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Louisville" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Lower_Princes" value="AST4"></zone>
	      <zone name="America/Maceio" value="BRT3"></zone>
	      <zone name="America/Managua" value="CST6"></zone>
	      <zone name="America/Manaus" value="AMT4"></zone>
	      <zone name="America/Marigot" value="AST4"></zone>
	      <zone name="America/Martinique" value="AST4"></zone>
	      <zone name="America/Matamoros" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Mazatlan" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Mendoza" value="ART3"></zone>
	      <zone name="America/Menominee" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Merida" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Metlakatla" value="PST8"></zone>
	      <zone name="America/Mexico_City" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Miquelon" value="PMST3PMDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Moncton" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Monterrey" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Montevideo" value="UYT3"></zone>
	      <zone name="America/Montreal" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Montserrat" value="AST4"></zone>
	      <zone name="America/Nassau" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/New_York" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Nipigon" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Nome" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Noronha" value="FNT2"></zone>
	      <zone name="America/North_Dakota/Beulah" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/North_Dakota/Center" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/North_Dakota/New_Salem" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Ojinaga" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Panama" value="EST5"></zone>
	      <zone name="America/Pangnirtung" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Paramaribo" value="SRT3"></zone>
	      <zone name="America/Phoenix" value="MST7"></zone>
	      <zone name="America/Port-au-Prince" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Port_of_Spain" value="AST4"></zone>
	      <zone name="America/Porto_Acre" value="ACT5"></zone>
	      <zone name="America/Porto_Velho" value="AMT4"></zone>
	      <zone name="America/Puerto_Rico" value="AST4"></zone>
	      <zone name="America/Rainy_River" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Rankin_Inlet" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Recife" value="BRT3"></zone>
	      <zone name="America/Regina" value="CST6"></zone>
	      <zone name="America/Resolute" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Rio_Branco" value="ACT5"></zone>
	      <zone name="America/Rosario" value="ART3"></zone>
	      <zone name="America/Santa_Isabel" value="PST8PDT,M4.1.0,M10.5.0"></zone>
	      <zone name="America/Santarem" value="BRT3"></zone>
	      <zone name="America/Santiago" value="CLT3"></zone>
	      <zone name="America/Santo_Domingo" value="AST4"></zone>
	      <zone name="America/Sao_Paulo" value="BRT3BRST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="America/Scoresbysund" value="EGT1EGST,M3.5.0/0,M10.5.0/1"></zone>
	      <zone name="America/Shiprock" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Sitka" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/St_Barthelemy" value="AST4"></zone>
	      <zone name="America/St_Johns" value="NST3:30NDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/St_Kitts" value="AST4"></zone>
	      <zone name="America/St_Lucia" value="AST4"></zone>
	      <zone name="America/St_Thomas" value="AST4"></zone>
	      <zone name="America/St_Vincent" value="AST4"></zone>
	      <zone name="America/Swift_Current" value="CST6"></zone>
	      <zone name="America/Tegucigalpa" value="CST6"></zone>
	      <zone name="America/Thule" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Thunder_Bay" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Tijuana" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Toronto" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Tortola" value="AST4"></zone>
	      <zone name="America/Vancouver" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Virgin" value="AST4"></zone>
	      <zone name="America/Whitehorse" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Winnipeg" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Yakutat" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="America/Yellowknife" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Antarctica/Casey" value="AWST-8"></zone>
	      <zone name="Antarctica/Davis" value="DAVT-7"></zone>
	      <zone name="Antarctica/DumontDUrville" value="DDUT-10"></zone>
	      <zone name="Antarctica/Macquarie" value="MIST-11"></zone>
	      <zone name="Antarctica/Mawson" value="MAWT-5"></zone>
	      <zone name="Antarctica/McMurdo" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="Antarctica/Palmer" value="CLT3"></zone>
	      <zone name="Antarctica/Rothera" value="ROTT3"></zone>
	      <zone name="Antarctica/South_Pole" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="Antarctica/Syowa" value="SYOT-3"></zone>
	      <zone name="Antarctica/Troll" value="UTC0CEST-2,M3.5.0/1,M10.5.0/3"></zone>
	      <zone name="Antarctica/Vostok" value="VOST-6"></zone>
	      <zone name="Arctic/Longyearbyen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Asia/Aden" value="AST-3"></zone>
	      <zone name="Asia/Almaty" value="ALMT-6"></zone>
	      <zone name="Asia/Amman" value="EEST"></zone>
	      <zone name="Asia/Anadyr" value="ANAT-12"></zone>
	      <zone name="Asia/Aqtau" value="AQTT-5"></zone>
	      <zone name="Asia/Aqtobe" value="AQTT-5"></zone>
	      <zone name="Asia/Ashgabat" value="TMT-5"></zone>
	      <zone name="Asia/Ashkhabad" value="TMT-5"></zone>
	      <zone name="Asia/Baghdad" value="AST-3"></zone>
	      <zone name="Asia/Bahrain" value="AST-3"></zone>
	      <zone name="Asia/Baku" value="AZT-4AZST,M3.5.0/4,M10.5.0/5"></zone>
	      <zone name="Asia/Bangkok" value="ICT-7"></zone>
	      <zone name="Asia/Beirut" value="EET-2EEST,M3.5.0/0,M10.5.0/0"></zone>
	      <zone name="Asia/Bishkek" value="KGT-6"></zone>
	      <zone name="Asia/Brunei" value="BNT-8"></zone>
	      <zone name="Asia/Calcutta" value="IST-5:30"></zone>
	      <zone name="Asia/Chita" value="IRKT-8"></zone>
	      <zone name="Asia/Choibalsan" value="CHOT-8CHOST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="Asia/Chongqing" value="CST-8"></zone>
	      <zone name="Asia/Chungking" value="CST-8"></zone>
	      <zone name="Asia/Colombo" value="IST-5:30"></zone>
	      <zone name="Asia/Dacca" value="BDT-6"></zone>
	      <zone name="Asia/Damascus" value="EET-2EEST,M3.5.5/0,M10.5.5/0"></zone>
	      <zone name="Asia/Dhaka" value="BDT-6"></zone>
	      <zone name="Asia/Dili" value="TLT-9"></zone>
	      <zone name="Asia/Dubai" value="GST-4"></zone>
	      <zone name="Asia/Dushanbe" value="TJT-5"></zone>
	      <zone name="Asia/Gaza" value="EEST"></zone>
	      <zone name="Asia/Harbin" value="CST-8"></zone>
	      <zone name="Asia/Hebron" value="EEST"></zone>
	      <zone name="Asia/Ho_Chi_Minh" value="ICT-7"></zone>
	      <zone name="Asia/Hong_Kong" value="HKT-8"></zone>
	      <zone name="Asia/Hovd" value="HOVT-7HOVST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="Asia/Irkutsk" value="IRKT-8"></zone>
	      <zone name="Asia/Istanbul" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Asia/Jakarta" value="WIB-7"></zone>
	      <zone name="Asia/Jayapura" value="WIT-9"></zone>
	      <zone name="Asia/Jerusalem" value="IDDT"></zone>
	      <zone name="Asia/Kabul" value="AFT-4:30"></zone>
	      <zone name="Asia/Kamchatka" value="PETT-12"></zone>
	      <zone name="Asia/Karachi" value="PKT-5"></zone>
	      <zone name="Asia/Kashgar" value="XJT-6"></zone>
	      <zone name="Asia/Kathmandu" value="NPT-5:45"></zone>
	      <zone name="Asia/Katmandu" value="NPT-5:45"></zone>
	      <zone name="Asia/Khandyga" value="YAKT-9"></zone>
	      <zone name="Asia/Kolkata" value="IST-5:30"></zone>
	      <zone name="Asia/Krasnoyarsk" value="KRAT-7"></zone>
	      <zone name="Asia/Kuala_Lumpur" value="MYT-8"></zone>
	      <zone name="Asia/Kuching" value="MYT-8"></zone>
	      <zone name="Asia/Kuwait" value="AST-3"></zone>
	      <zone name="Asia/Macao" value="CST-8"></zone>
	      <zone name="Asia/Macau" value="CST-8"></zone>
	      <zone name="Asia/Magadan" value="MAGT-10"></zone>
	      <zone name="Asia/Makassar" value="WITA-8"></zone>
	      <zone name="Asia/Manila" value="PHT-8"></zone>
	      <zone name="Asia/Muscat" value="GST-4"></zone>
	      <zone name="Asia/Nicosia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Asia/Novokuznetsk" value="KRAT-7"></zone>
	      <zone name="Asia/Novosibirsk" value="NOVT-6"></zone>
	      <zone name="Asia/Omsk" value="OMST-6"></zone>
	      <zone name="Asia/Oral" value="ORAT-5"></zone>
	      <zone name="Asia/Phnom_Penh" value="ICT-7"></zone>
	      <zone name="Asia/Pontianak" value="WIB-7"></zone>
	      <zone name="Asia/Pyongyang" value="KST-8:30"></zone>
	      <zone name="Asia/Qatar" value="AST-3"></zone>
	      <zone name="Asia/Qyzylorda" value="QYZT-6"></zone>
	      <zone name="Asia/Rangoon" value="MMT-6:30"></zone>
	      <zone name="Asia/Riyadh" value="AST-3"></zone>
	      <zone name="Asia/Saigon" value="ICT-7"></zone>
	      <zone name="Asia/Sakhalin" value="SAKT-10"></zone>
	      <zone name="Asia/Samarkand" value="UZT-5"></zone>
	      <zone name="Asia/Seoul" value="KST-9"></zone>
	      <zone name="Asia/Shanghai" value="CST-8"></zone>
	      <zone name="Asia/Singapore" value="SGT-8"></zone>
	      <zone name="Asia/Srednekolymsk" value="SRET-11"></zone>
	      <zone name="Asia/Taipei" value="CST-8"></zone>
	      <zone name="Asia/Tashkent" value="UZT-5"></zone>
	      <zone name="Asia/Tbilisi" value="GET-4"></zone>
	      <zone name="Asia/Tehran" value="IRDT"></zone>
	      <zone name="Asia/Tel_Aviv" value="IDDT"></zone>
	      <zone name="Asia/Thimbu" value="BTT-6"></zone>
	      <zone name="Asia/Thimphu" value="BTT-6"></zone>
	      <zone name="Asia/Tokyo" value="JST-9"></zone>
	      <zone name="Asia/Ujung_Pandang" value="WITA-8"></zone>
	      <zone name="Asia/Ulaanbaatar" value="ULAT-8ULAST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="Asia/Ulan_Bator" value="ULAT-8ULAST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="Asia/Urumqi" value="XJT-6"></zone>
	      <zone name="Asia/Ust-Nera" value="VLAT-10"></zone>
	      <zone name="Asia/Vientiane" value="ICT-7"></zone>
	      <zone name="Asia/Vladivostok" value="VLAT-10"></zone>
	      <zone name="Asia/Yakutsk" value="YAKT-9"></zone>
	      <zone name="Asia/Yekaterinburg" value="YEKT-5"></zone>
	      <zone name="Asia/Yerevan" value="AMT-4"></zone>
	      <zone name="Atlantic/Azores" value="AZOT1AZOST,M3.5.0/0,M10.5.0/1"></zone>
	      <zone name="Atlantic/Bermuda" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="Atlantic/Canary" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Atlantic/Cape_Verde" value="CVT1"></zone>
	      <zone name="Atlantic/Faeroe" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Atlantic/Faroe" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Atlantic/Jan_Mayen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Atlantic/Madeira" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Atlantic/Reykjavik" value="GMT0"></zone>
	      <zone name="Atlantic/South_Georgia" value="GST2"></zone>
	      <zone name="Atlantic/St_Helena" value="GMT0"></zone>
	      <zone name="Atlantic/Stanley" value="FKST3"></zone>
	      <zone name="Australia/ACT" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Adelaide" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Brisbane" value="AEST-10"></zone>
	      <zone name="Australia/Broken_Hill" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Canberra" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Currie" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Darwin" value="ACST-9:30"></zone>
	      <zone name="Australia/Eucla" value="ACWST-8:45"></zone>
	      <zone name="Australia/Hobart" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/LHI" value="LHST-10:30LHDT-11,M10.1.0,M4.1.0"></zone>
	      <zone name="Australia/Lindeman" value="AEST-10"></zone>
	      <zone name="Australia/Lord_Howe" value="LHST-10:30LHDT-11,M10.1.0,M4.1.0"></zone>
	      <zone name="Australia/Melbourne" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/NSW" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/North" value="ACST-9:30"></zone>
	      <zone name="Australia/Perth" value="AWST-8"></zone>
	      <zone name="Australia/Queensland" value="AEST-10"></zone>
	      <zone name="Australia/South" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Sydney" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Tasmania" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/Victoria" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Australia/West" value="AWST-8"></zone>
	      <zone name="Australia/Yancowinna" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="Brazil/Acre" value="ACT5"></zone>
	      <zone name="Brazil/DeNoronha" value="FNT2"></zone>
	      <zone name="Brazil/East" value="BRT3BRST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="Brazil/West" value="AMT4"></zone>
	      <zone name="CET" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="CST6CDT" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/Atlantic" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/Central" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/East-Saskatchewan" value="CST6"></zone>
	      <zone name="Canada/Eastern" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/Mountain" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/Newfoundland" value="NST3:30NDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/Pacific" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Canada/Saskatchewan" value="CST6"></zone>
	      <zone name="Canada/Yukon" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Chile/Continental" value="CLT3"></zone>
	      <zone name="Chile/EasterIsland" value="EAST5"></zone>
	      <zone name="Cuba" value="CST5CDT,M3.2.0/0,M11.1.0/1"></zone>
	      <zone name="EET" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="EST" value="EST5"></zone>
	      <zone name="EST5EDT" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Egypt" value="EET-2"></zone>
	      <zone name="Eire" value="GMT0IST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Etc/GMT" value="GMT0"></zone>
	      <zone name="Etc/GMT+0" value="GMT0"></zone>
	      <zone name="Etc/GMT+1" value="&lt;GMT+1&gt;1"></zone>
	      <zone name="Etc/GMT+10" value="&lt;GMT+10&gt;10"></zone>
	      <zone name="Etc/GMT+11" value="&lt;GMT+11&gt;11"></zone>
	      <zone name="Etc/GMT+12" value="&lt;GMT+12&gt;12"></zone>
	      <zone name="Etc/GMT+2" value="&lt;GMT+2&gt;2"></zone>
	      <zone name="Etc/GMT+3" value="&lt;GMT+3&gt;3"></zone>
	      <zone name="Etc/GMT+4" value="&lt;GMT+4&gt;4"></zone>
	      <zone name="Etc/GMT+5" value="&lt;GMT+5&gt;5"></zone>
	      <zone name="Etc/GMT+6" value="&lt;GMT+6&gt;6"></zone>
	      <zone name="Etc/GMT+7" value="&lt;GMT+7&gt;7"></zone>
	      <zone name="Etc/GMT+8" value="&lt;GMT+8&gt;8"></zone>
	      <zone name="Etc/GMT+9" value="&lt;GMT+9&gt;9"></zone>
	      <zone name="Etc/GMT-0" value="GMT0"></zone>
	      <zone name="Etc/GMT-1" value="&lt;GMT-1&gt;-1"></zone>
	      <zone name="Etc/GMT-10" value="&lt;GMT-10&gt;-10"></zone>
	      <zone name="Etc/GMT-11" value="&lt;GMT-11&gt;-11"></zone>
	      <zone name="Etc/GMT-12" value="&lt;GMT-12&gt;-12"></zone>
	      <zone name="Etc/GMT-13" value="&lt;GMT-13&gt;-13"></zone>
	      <zone name="Etc/GMT-14" value="&lt;GMT-14&gt;-14"></zone>
	      <zone name="Etc/GMT-2" value="&lt;GMT-2&gt;-2"></zone>
	      <zone name="Etc/GMT-3" value="&lt;GMT-3&gt;-3"></zone>
	      <zone name="Etc/GMT-4" value="&lt;GMT-4&gt;-4"></zone>
	      <zone name="Etc/GMT-5" value="&lt;GMT-5&gt;-5"></zone>
	      <zone name="Etc/GMT-6" value="&lt;GMT-6&gt;-6"></zone>
	      <zone name="Etc/GMT-7" value="&lt;GMT-7&gt;-7"></zone>
	      <zone name="Etc/GMT-8" value="&lt;GMT-8&gt;-8"></zone>
	      <zone name="Etc/GMT-9" value="&lt;GMT-9&gt;-9"></zone>
	      <zone name="Etc/GMT0" value="GMT0"></zone>
	      <zone name="Etc/Greenwich" value="GMT0"></zone>
	      <zone name="Etc/UCT" value="UCT0"></zone>
	      <zone name="Etc/UTC" value="UTC0"></zone>
	      <zone name="Etc/Universal" value="UTC0"></zone>
	      <zone name="Etc/Zulu" value="UTC0"></zone>
	      <zone name="Europe/Amsterdam" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Andorra" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Athens" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Belfast" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Belgrade" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Berlin" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Bratislava" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Brussels" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Bucharest" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Budapest" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Busingen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Chisinau" value="EET-2EEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Copenhagen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Dublin" value="GMT0IST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Gibraltar" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Guernsey" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Helsinki" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Isle_of_Man" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Istanbul" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Jersey" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Kaliningrad" value="EET-2"></zone>
	      <zone name="Europe/Kiev" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Lisbon" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Ljubljana" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/London" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Europe/Luxembourg" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Madrid" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Malta" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Mariehamn" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Minsk" value="MSK-3"></zone>
	      <zone name="Europe/Monaco" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Moscow" value="MSK-3"></zone>
	      <zone name="Europe/Nicosia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Oslo" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Paris" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Podgorica" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Prague" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Riga" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Rome" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Samara" value="SAMT-4"></zone>
	      <zone name="Europe/San_Marino" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Sarajevo" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Simferopol" value="MSK-3"></zone>
	      <zone name="Europe/Skopje" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Sofia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Stockholm" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Tallinn" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Tirane" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Tiraspol" value="EET-2EEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Uzhgorod" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Vaduz" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Vatican" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Vienna" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Vilnius" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Volgograd" value="MSK-3"></zone>
	      <zone name="Europe/Warsaw" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Zagreb" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Europe/Zaporozhye" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="Europe/Zurich" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="GB" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="GB-Eire" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="GMT" value="GMT0"></zone>
	      <zone name="GMT+0" value="GMT0"></zone>
	      <zone name="GMT-0" value="GMT0"></zone>
	      <zone name="GMT0" value="GMT0"></zone>
	      <zone name="Greenwich" value="GMT0"></zone>
	      <zone name="HST" value="HST10"></zone>
	      <zone name="Hongkong" value="HKT-8"></zone>
	      <zone name="Iceland" value="GMT0"></zone>
	      <zone name="Indian/Antananarivo" value="EAT-3"></zone>
	      <zone name="Indian/Chagos" value="IOT-6"></zone>
	      <zone name="Indian/Christmas" value="CXT-7"></zone>
	      <zone name="Indian/Cocos" value="CCT-6:30"></zone>
	      <zone name="Indian/Comoro" value="EAT-3"></zone>
	      <zone name="Indian/Kerguelen" value="TFT-5"></zone>
	      <zone name="Indian/Mahe" value="SCT-4"></zone>
	      <zone name="Indian/Maldives" value="MVT-5"></zone>
	      <zone name="Indian/Mauritius" value="MUT-4"></zone>
	      <zone name="Indian/Mayotte" value="EAT-3"></zone>
	      <zone name="Indian/Reunion" value="RET-4"></zone>
	      <zone name="Iran" value="IRDT"></zone>
	      <zone name="Israel" value="IDDT"></zone>
	      <zone name="Jamaica" value="EST5"></zone>
	      <zone name="Japan" value="JST-9"></zone>
	      <zone name="Kwajalein" value="MHT-12"></zone>
	      <zone name="Libya" value="EET-2"></zone>
	      <zone name="MET" value="MET-1MEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="MST" value="MST7"></zone>
	      <zone name="MST7MDT" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Mexico/BajaNorte" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Mexico/BajaSur" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="Mexico/General" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="NZ" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="NZ-CHAT" value="CHAST-12:45CHADT,M9.5.0/2:45,M4.1.0/3:45"></zone>
	      <zone name="Navajo" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="PRC" value="CST-8"></zone>
	      <zone name="PST8PDT" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="Pacific/Apia" value="WSST-13WSDT,M9.5.0/3,M4.1.0/4"></zone>
	      <zone name="Pacific/Auckland" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="Pacific/Bougainville" value="BST-11"></zone>
	      <zone name="Pacific/Chatham" value="CHAST-12:45CHADT,M9.5.0/2:45,M4.1.0/3:45"></zone>
	      <zone name="Pacific/Chuuk" value="CHUT-10"></zone>
	      <zone name="Pacific/Easter" value="EAST5"></zone>
	      <zone name="Pacific/Efate" value="VUT-11"></zone>
	      <zone name="Pacific/Enderbury" value="PHOT-13"></zone>
	      <zone name="Pacific/Fakaofo" value="TKT-13"></zone>
	      <zone name="Pacific/Fiji" value="FJT-12FJST,M11.1.0,M1.3.0/3"></zone>
	      <zone name="Pacific/Funafuti" value="TVT-12"></zone>
	      <zone name="Pacific/Galapagos" value="GALT6"></zone>
	      <zone name="Pacific/Gambier" value="GAMT9"></zone>
	      <zone name="Pacific/Guadalcanal" value="SBT-11"></zone>
	      <zone name="Pacific/Guam" value="ChST-10"></zone>
	      <zone name="Pacific/Honolulu" value="HST10"></zone>
	      <zone name="Pacific/Johnston" value="HST10"></zone>
	      <zone name="Pacific/Kiritimati" value="LINT-14"></zone>
	      <zone name="Pacific/Kosrae" value="KOST-11"></zone>
	      <zone name="Pacific/Kwajalein" value="MHT-12"></zone>
	      <zone name="Pacific/Majuro" value="MHT-12"></zone>
	      <zone name="Pacific/Marquesas" value="MART9:30"></zone>
	      <zone name="Pacific/Midway" value="SST11"></zone>
	      <zone name="Pacific/Nauru" value="NRT-12"></zone>
	      <zone name="Pacific/Niue" value="NUT11"></zone>
	      <zone name="Pacific/Norfolk" value="NFT-11"></zone>
	      <zone name="Pacific/Noumea" value="NCT-11"></zone>
	      <zone name="Pacific/Pago_Pago" value="SST11"></zone>
	      <zone name="Pacific/Palau" value="PWT-9"></zone>
	      <zone name="Pacific/Pitcairn" value="PST8"></zone>
	      <zone name="Pacific/Pohnpei" value="PONT-11"></zone>
	      <zone name="Pacific/Ponape" value="PONT-11"></zone>
	      <zone name="Pacific/Port_Moresby" value="PGT-10"></zone>
	      <zone name="Pacific/Rarotonga" value="CKT10"></zone>
	      <zone name="Pacific/Saipan" value="ChST-10"></zone>
	      <zone name="Pacific/Samoa" value="SST11"></zone>
	      <zone name="Pacific/Tahiti" value="TAHT10"></zone>
	      <zone name="Pacific/Tarawa" value="GILT-12"></zone>
	      <zone name="Pacific/Tongatapu" value="TOT-13"></zone>
	      <zone name="Pacific/Truk" value="CHUT-10"></zone>
	      <zone name="Pacific/Wake" value="WAKT-12"></zone>
	      <zone name="Pacific/Wallis" value="WFT-12"></zone>
	      <zone name="Pacific/Yap" value="CHUT-10"></zone>
	      <zone name="Poland" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="Portugal" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="ROC" value="CST-8"></zone>
	      <zone name="ROK" value="KST-9"></zone>
	      <zone name="Singapore" value="SGT-8"></zone>
	      <zone name="Turkey" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="UCT" value="UCT0"></zone>
	      <zone name="US/Alaska" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Aleutian" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Arizona" value="MST7"></zone>
	      <zone name="US/Central" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/East-Indiana" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Eastern" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Hawaii" value="HST10"></zone>
	      <zone name="US/Indiana-Starke" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Michigan" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Mountain" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Pacific" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Pacific-New" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="US/Samoa" value="SST11"></zone>
	      <zone name="UTC" value="UTC0"></zone>
	      <zone name="Universal" value="UTC0"></zone>
	      <zone name="W-SU" value="MSK-3"></zone>
	      <zone name="WET" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="Zulu" value="UTC0"></zone>
	      <zone name="posix/Africa/Abidjan" value="GMT0"></zone>
	      <zone name="posix/Africa/Accra" value="GMT0"></zone>
	      <zone name="posix/Africa/Addis_Ababa" value="EAT-3"></zone>
	      <zone name="posix/Africa/Algiers" value="CET-1"></zone>
	      <zone name="posix/Africa/Asmara" value="EAT-3"></zone>
	      <zone name="posix/Africa/Asmera" value="EAT-3"></zone>
	      <zone name="posix/Africa/Bamako" value="GMT0"></zone>
	      <zone name="posix/Africa/Bangui" value="WAT-1"></zone>
	      <zone name="posix/Africa/Banjul" value="GMT0"></zone>
	      <zone name="posix/Africa/Bissau" value="GMT0"></zone>
	      <zone name="posix/Africa/Blantyre" value="CAT-2"></zone>
	      <zone name="posix/Africa/Brazzaville" value="WAT-1"></zone>
	      <zone name="posix/Africa/Bujumbura" value="CAT-2"></zone>
	      <zone name="posix/Africa/Cairo" value="EET-2"></zone>
	      <zone name="posix/Africa/Casablanca" value="WET0WEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Africa/Ceuta" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Africa/Conakry" value="GMT0"></zone>
	      <zone name="posix/Africa/Dakar" value="GMT0"></zone>
	      <zone name="posix/Africa/Dar_es_Salaam" value="EAT-3"></zone>
	      <zone name="posix/Africa/Djibouti" value="EAT-3"></zone>
	      <zone name="posix/Africa/Douala" value="WAT-1"></zone>
	      <zone name="posix/Africa/El_Aaiun" value="WET0WEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Africa/Freetown" value="GMT0"></zone>
	      <zone name="posix/Africa/Gaborone" value="CAT-2"></zone>
	      <zone name="posix/Africa/Harare" value="CAT-2"></zone>
	      <zone name="posix/Africa/Johannesburg" value="SAST-2"></zone>
	      <zone name="posix/Africa/Juba" value="EAT-3"></zone>
	      <zone name="posix/Africa/Kampala" value="EAT-3"></zone>
	      <zone name="posix/Africa/Khartoum" value="EAT-3"></zone>
	      <zone name="posix/Africa/Kigali" value="CAT-2"></zone>
	      <zone name="posix/Africa/Kinshasa" value="WAT-1"></zone>
	      <zone name="posix/Africa/Lagos" value="WAT-1"></zone>
	      <zone name="posix/Africa/Libreville" value="WAT-1"></zone>
	      <zone name="posix/Africa/Lome" value="GMT0"></zone>
	      <zone name="posix/Africa/Luanda" value="WAT-1"></zone>
	      <zone name="posix/Africa/Lubumbashi" value="CAT-2"></zone>
	      <zone name="posix/Africa/Lusaka" value="CAT-2"></zone>
	      <zone name="posix/Africa/Malabo" value="WAT-1"></zone>
	      <zone name="posix/Africa/Maputo" value="CAT-2"></zone>
	      <zone name="posix/Africa/Maseru" value="SAST-2"></zone>
	      <zone name="posix/Africa/Mbabane" value="SAST-2"></zone>
	      <zone name="posix/Africa/Mogadishu" value="EAT-3"></zone>
	      <zone name="posix/Africa/Monrovia" value="GMT0"></zone>
	      <zone name="posix/Africa/Nairobi" value="EAT-3"></zone>
	      <zone name="posix/Africa/Ndjamena" value="WAT-1"></zone>
	      <zone name="posix/Africa/Niamey" value="WAT-1"></zone>
	      <zone name="posix/Africa/Nouakchott" value="GMT0"></zone>
	      <zone name="posix/Africa/Ouagadougou" value="GMT0"></zone>
	      <zone name="posix/Africa/Porto-Novo" value="WAT-1"></zone>
	      <zone name="posix/Africa/Sao_Tome" value="GMT0"></zone>
	      <zone name="posix/Africa/Timbuktu" value="GMT0"></zone>
	      <zone name="posix/Africa/Tripoli" value="EET-2"></zone>
	      <zone name="posix/Africa/Tunis" value="CET-1"></zone>
	      <zone name="posix/Africa/Windhoek" value="WAT-1WAST,M9.1.0,M4.1.0"></zone>
	      <zone name="posix/America/Adak" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Anchorage" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Anguilla" value="AST4"></zone>
	      <zone name="posix/America/Antigua" value="AST4"></zone>
	      <zone name="posix/America/Araguaina" value="BRT3"></zone>
	      <zone name="posix/America/Argentina/Buenos_Aires" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Catamarca" value="ART3"></zone>
	      <zone name="posix/America/Argentina/ComodRivadavia" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Cordoba" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Jujuy" value="ART3"></zone>
	      <zone name="posix/America/Argentina/La_Rioja" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Mendoza" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Rio_Gallegos" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Salta" value="ART3"></zone>
	      <zone name="posix/America/Argentina/San_Juan" value="ART3"></zone>
	      <zone name="posix/America/Argentina/San_Luis" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Tucuman" value="ART3"></zone>
	      <zone name="posix/America/Argentina/Ushuaia" value="ART3"></zone>
	      <zone name="posix/America/Aruba" value="AST4"></zone>
	      <zone name="posix/America/Asuncion" value="PYT4PYST,M10.1.0/0,M3.4.0/0"></zone>
	      <zone name="posix/America/Atikokan" value="EST5"></zone>
	      <zone name="posix/America/Atka" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Bahia" value="BRT3"></zone>
	      <zone name="posix/America/Bahia_Banderas" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Barbados" value="AST4"></zone>
	      <zone name="posix/America/Belem" value="BRT3"></zone>
	      <zone name="posix/America/Belize" value="CST6"></zone>
	      <zone name="posix/America/Blanc-Sablon" value="AST4"></zone>
	      <zone name="posix/America/Boa_Vista" value="AMT4"></zone>
	      <zone name="posix/America/Bogota" value="COT5"></zone>
	      <zone name="posix/America/Boise" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Buenos_Aires" value="ART3"></zone>
	      <zone name="posix/America/Cambridge_Bay" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Campo_Grande" value="AMT4AMST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="posix/America/Cancun" value="EST5"></zone>
	      <zone name="posix/America/Caracas" value="VET4:30"></zone>
	      <zone name="posix/America/Catamarca" value="ART3"></zone>
	      <zone name="posix/America/Cayenne" value="GFT3"></zone>
	      <zone name="posix/America/Cayman" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Chicago" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Chihuahua" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Coral_Harbour" value="EST5"></zone>
	      <zone name="posix/America/Cordoba" value="ART3"></zone>
	      <zone name="posix/America/Costa_Rica" value="CST6"></zone>
	      <zone name="posix/America/Creston" value="MST7"></zone>
	      <zone name="posix/America/Cuiaba" value="AMT4AMST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="posix/America/Curacao" value="AST4"></zone>
	      <zone name="posix/America/Danmarkshavn" value="GMT0"></zone>
	      <zone name="posix/America/Dawson" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Dawson_Creek" value="MST7"></zone>
	      <zone name="posix/America/Denver" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Detroit" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Dominica" value="AST4"></zone>
	      <zone name="posix/America/Edmonton" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Eirunepe" value="ACT5"></zone>
	      <zone name="posix/America/El_Salvador" value="CST6"></zone>
	      <zone name="posix/America/Ensenada" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Fort_Nelson" value="MST7"></zone>
	      <zone name="posix/America/Fort_Wayne" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Fortaleza" value="BRT3"></zone>
	      <zone name="posix/America/Glace_Bay" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Godthab" value="WGST"></zone>
	      <zone name="posix/America/Goose_Bay" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Grand_Turk" value="AST4"></zone>
	      <zone name="posix/America/Grenada" value="AST4"></zone>
	      <zone name="posix/America/Guadeloupe" value="AST4"></zone>
	      <zone name="posix/America/Guatemala" value="CST6"></zone>
	      <zone name="posix/America/Guayaquil" value="ECT5"></zone>
	      <zone name="posix/America/Guyana" value="GYT4"></zone>
	      <zone name="posix/America/Halifax" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Havana" value="CST5CDT,M3.2.0/0,M11.1.0/1"></zone>
	      <zone name="posix/America/Hermosillo" value="MST7"></zone>
	      <zone name="posix/America/Indiana/Indianapolis" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Knox" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Marengo" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Petersburg" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Tell_City" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Vevay" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Vincennes" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indiana/Winamac" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Indianapolis" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Inuvik" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Iqaluit" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Jamaica" value="EST5"></zone>
	      <zone name="posix/America/Jujuy" value="ART3"></zone>
	      <zone name="posix/America/Juneau" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Kentucky/Louisville" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Kentucky/Monticello" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Knox_IN" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Kralendijk" value="AST4"></zone>
	      <zone name="posix/America/La_Paz" value="BOT4"></zone>
	      <zone name="posix/America/Lima" value="PET5"></zone>
	      <zone name="posix/America/Los_Angeles" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Louisville" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Lower_Princes" value="AST4"></zone>
	      <zone name="posix/America/Maceio" value="BRT3"></zone>
	      <zone name="posix/America/Managua" value="CST6"></zone>
	      <zone name="posix/America/Manaus" value="AMT4"></zone>
	      <zone name="posix/America/Marigot" value="AST4"></zone>
	      <zone name="posix/America/Martinique" value="AST4"></zone>
	      <zone name="posix/America/Matamoros" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Mazatlan" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Mendoza" value="ART3"></zone>
	      <zone name="posix/America/Menominee" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Merida" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Metlakatla" value="PST8"></zone>
	      <zone name="posix/America/Mexico_City" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Miquelon" value="PMST3PMDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Moncton" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Monterrey" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Montevideo" value="UYT3"></zone>
	      <zone name="posix/America/Montreal" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Montserrat" value="AST4"></zone>
	      <zone name="posix/America/Nassau" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/New_York" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Nipigon" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Nome" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Noronha" value="FNT2"></zone>
	      <zone name="posix/America/North_Dakota/Beulah" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/North_Dakota/Center" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/North_Dakota/New_Salem" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Ojinaga" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Panama" value="EST5"></zone>
	      <zone name="posix/America/Pangnirtung" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Paramaribo" value="SRT3"></zone>
	      <zone name="posix/America/Phoenix" value="MST7"></zone>
	      <zone name="posix/America/Port-au-Prince" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Port_of_Spain" value="AST4"></zone>
	      <zone name="posix/America/Porto_Acre" value="ACT5"></zone>
	      <zone name="posix/America/Porto_Velho" value="AMT4"></zone>
	      <zone name="posix/America/Puerto_Rico" value="AST4"></zone>
	      <zone name="posix/America/Rainy_River" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Rankin_Inlet" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Recife" value="BRT3"></zone>
	      <zone name="posix/America/Regina" value="CST6"></zone>
	      <zone name="posix/America/Resolute" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Rio_Branco" value="ACT5"></zone>
	      <zone name="posix/America/Rosario" value="ART3"></zone>
	      <zone name="posix/America/Santa_Isabel" value="PST8PDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/America/Santarem" value="BRT3"></zone>
	      <zone name="posix/America/Santiago" value="CLT3"></zone>
	      <zone name="posix/America/Santo_Domingo" value="AST4"></zone>
	      <zone name="posix/America/Sao_Paulo" value="BRT3BRST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="posix/America/Scoresbysund" value="EGT1EGST,M3.5.0/0,M10.5.0/1"></zone>
	      <zone name="posix/America/Shiprock" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Sitka" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/St_Barthelemy" value="AST4"></zone>
	      <zone name="posix/America/St_Johns" value="NST3:30NDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/St_Kitts" value="AST4"></zone>
	      <zone name="posix/America/St_Lucia" value="AST4"></zone>
	      <zone name="posix/America/St_Thomas" value="AST4"></zone>
	      <zone name="posix/America/St_Vincent" value="AST4"></zone>
	      <zone name="posix/America/Swift_Current" value="CST6"></zone>
	      <zone name="posix/America/Tegucigalpa" value="CST6"></zone>
	      <zone name="posix/America/Thule" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Thunder_Bay" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Tijuana" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Toronto" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Tortola" value="AST4"></zone>
	      <zone name="posix/America/Vancouver" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Virgin" value="AST4"></zone>
	      <zone name="posix/America/Whitehorse" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Winnipeg" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Yakutat" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/America/Yellowknife" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Antarctica/Casey" value="AWST-8"></zone>
	      <zone name="posix/Antarctica/Davis" value="DAVT-7"></zone>
	      <zone name="posix/Antarctica/DumontDUrville" value="DDUT-10"></zone>
	      <zone name="posix/Antarctica/Macquarie" value="MIST-11"></zone>
	      <zone name="posix/Antarctica/Mawson" value="MAWT-5"></zone>
	      <zone name="posix/Antarctica/McMurdo" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="posix/Antarctica/Palmer" value="CLT3"></zone>
	      <zone name="posix/Antarctica/Rothera" value="ROTT3"></zone>
	      <zone name="posix/Antarctica/South_Pole" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="posix/Antarctica/Syowa" value="SYOT-3"></zone>
	      <zone name="posix/Antarctica/Troll" value="UTC0CEST-2,M3.5.0/1,M10.5.0/3"></zone>
	      <zone name="posix/Antarctica/Vostok" value="VOST-6"></zone>
	      <zone name="posix/Arctic/Longyearbyen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Asia/Aden" value="AST-3"></zone>
	      <zone name="posix/Asia/Almaty" value="ALMT-6"></zone>
	      <zone name="posix/Asia/Amman" value="EEST"></zone>
	      <zone name="posix/Asia/Anadyr" value="ANAT-12"></zone>
	      <zone name="posix/Asia/Aqtau" value="AQTT-5"></zone>
	      <zone name="posix/Asia/Aqtobe" value="AQTT-5"></zone>
	      <zone name="posix/Asia/Ashgabat" value="TMT-5"></zone>
	      <zone name="posix/Asia/Ashkhabad" value="TMT-5"></zone>
	      <zone name="posix/Asia/Baghdad" value="AST-3"></zone>
	      <zone name="posix/Asia/Bahrain" value="AST-3"></zone>
	      <zone name="posix/Asia/Baku" value="AZT-4AZST,M3.5.0/4,M10.5.0/5"></zone>
	      <zone name="posix/Asia/Bangkok" value="ICT-7"></zone>
	      <zone name="posix/Asia/Beirut" value="EET-2EEST,M3.5.0/0,M10.5.0/0"></zone>
	      <zone name="posix/Asia/Bishkek" value="KGT-6"></zone>
	      <zone name="posix/Asia/Brunei" value="BNT-8"></zone>
	      <zone name="posix/Asia/Calcutta" value="IST-5:30"></zone>
	      <zone name="posix/Asia/Chita" value="IRKT-8"></zone>
	      <zone name="posix/Asia/Choibalsan" value="CHOT-8CHOST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="posix/Asia/Chongqing" value="CST-8"></zone>
	      <zone name="posix/Asia/Chungking" value="CST-8"></zone>
	      <zone name="posix/Asia/Colombo" value="IST-5:30"></zone>
	      <zone name="posix/Asia/Dacca" value="BDT-6"></zone>
	      <zone name="posix/Asia/Damascus" value="EET-2EEST,M3.5.5/0,M10.5.5/0"></zone>
	      <zone name="posix/Asia/Dhaka" value="BDT-6"></zone>
	      <zone name="posix/Asia/Dili" value="TLT-9"></zone>
	      <zone name="posix/Asia/Dubai" value="GST-4"></zone>
	      <zone name="posix/Asia/Dushanbe" value="TJT-5"></zone>
	      <zone name="posix/Asia/Gaza" value="EEST"></zone>
	      <zone name="posix/Asia/Harbin" value="CST-8"></zone>
	      <zone name="posix/Asia/Hebron" value="EEST"></zone>
	      <zone name="posix/Asia/Ho_Chi_Minh" value="ICT-7"></zone>
	      <zone name="posix/Asia/Hong_Kong" value="HKT-8"></zone>
	      <zone name="posix/Asia/Hovd" value="HOVT-7HOVST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="posix/Asia/Irkutsk" value="IRKT-8"></zone>
	      <zone name="posix/Asia/Istanbul" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Asia/Jakarta" value="WIB-7"></zone>
	      <zone name="posix/Asia/Jayapura" value="WIT-9"></zone>
	      <zone name="posix/Asia/Jerusalem" value="IDDT"></zone>
	      <zone name="posix/Asia/Kabul" value="AFT-4:30"></zone>
	      <zone name="posix/Asia/Kamchatka" value="PETT-12"></zone>
	      <zone name="posix/Asia/Karachi" value="PKT-5"></zone>
	      <zone name="posix/Asia/Kashgar" value="XJT-6"></zone>
	      <zone name="posix/Asia/Kathmandu" value="NPT-5:45"></zone>
	      <zone name="posix/Asia/Katmandu" value="NPT-5:45"></zone>
	      <zone name="posix/Asia/Khandyga" value="YAKT-9"></zone>
	      <zone name="posix/Asia/Kolkata" value="IST-5:30"></zone>
	      <zone name="posix/Asia/Krasnoyarsk" value="KRAT-7"></zone>
	      <zone name="posix/Asia/Kuala_Lumpur" value="MYT-8"></zone>
	      <zone name="posix/Asia/Kuching" value="MYT-8"></zone>
	      <zone name="posix/Asia/Kuwait" value="AST-3"></zone>
	      <zone name="posix/Asia/Macao" value="CST-8"></zone>
	      <zone name="posix/Asia/Macau" value="CST-8"></zone>
	      <zone name="posix/Asia/Magadan" value="MAGT-10"></zone>
	      <zone name="posix/Asia/Makassar" value="WITA-8"></zone>
	      <zone name="posix/Asia/Manila" value="PHT-8"></zone>
	      <zone name="posix/Asia/Muscat" value="GST-4"></zone>
	      <zone name="posix/Asia/Nicosia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Asia/Novokuznetsk" value="KRAT-7"></zone>
	      <zone name="posix/Asia/Novosibirsk" value="NOVT-6"></zone>
	      <zone name="posix/Asia/Omsk" value="OMST-6"></zone>
	      <zone name="posix/Asia/Oral" value="ORAT-5"></zone>
	      <zone name="posix/Asia/Phnom_Penh" value="ICT-7"></zone>
	      <zone name="posix/Asia/Pontianak" value="WIB-7"></zone>
	      <zone name="posix/Asia/Pyongyang" value="KST-8:30"></zone>
	      <zone name="posix/Asia/Qatar" value="AST-3"></zone>
	      <zone name="posix/Asia/Qyzylorda" value="QYZT-6"></zone>
	      <zone name="posix/Asia/Rangoon" value="MMT-6:30"></zone>
	      <zone name="posix/Asia/Riyadh" value="AST-3"></zone>
	      <zone name="posix/Asia/Saigon" value="ICT-7"></zone>
	      <zone name="posix/Asia/Sakhalin" value="SAKT-10"></zone>
	      <zone name="posix/Asia/Samarkand" value="UZT-5"></zone>
	      <zone name="posix/Asia/Seoul" value="KST-9"></zone>
	      <zone name="posix/Asia/Shanghai" value="CST-8"></zone>
	      <zone name="posix/Asia/Singapore" value="SGT-8"></zone>
	      <zone name="posix/Asia/Srednekolymsk" value="SRET-11"></zone>
	      <zone name="posix/Asia/Taipei" value="CST-8"></zone>
	      <zone name="posix/Asia/Tashkent" value="UZT-5"></zone>
	      <zone name="posix/Asia/Tbilisi" value="GET-4"></zone>
	      <zone name="posix/Asia/Tehran" value="IRDT"></zone>
	      <zone name="posix/Asia/Tel_Aviv" value="IDDT"></zone>
	      <zone name="posix/Asia/Thimbu" value="BTT-6"></zone>
	      <zone name="posix/Asia/Thimphu" value="BTT-6"></zone>
	      <zone name="posix/Asia/Tokyo" value="JST-9"></zone>
	      <zone name="posix/Asia/Ujung_Pandang" value="WITA-8"></zone>
	      <zone name="posix/Asia/Ulaanbaatar" value="ULAT-8ULAST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="posix/Asia/Ulan_Bator" value="ULAT-8ULAST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="posix/Asia/Urumqi" value="XJT-6"></zone>
	      <zone name="posix/Asia/Ust-Nera" value="VLAT-10"></zone>
	      <zone name="posix/Asia/Vientiane" value="ICT-7"></zone>
	      <zone name="posix/Asia/Vladivostok" value="VLAT-10"></zone>
	      <zone name="posix/Asia/Yakutsk" value="YAKT-9"></zone>
	      <zone name="posix/Asia/Yekaterinburg" value="YEKT-5"></zone>
	      <zone name="posix/Asia/Yerevan" value="AMT-4"></zone>
	      <zone name="posix/Atlantic/Azores" value="AZOT1AZOST,M3.5.0/0,M10.5.0/1"></zone>
	      <zone name="posix/Atlantic/Bermuda" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Atlantic/Canary" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Atlantic/Cape_Verde" value="CVT1"></zone>
	      <zone name="posix/Atlantic/Faeroe" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Atlantic/Faroe" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Atlantic/Jan_Mayen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Atlantic/Madeira" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Atlantic/Reykjavik" value="GMT0"></zone>
	      <zone name="posix/Atlantic/South_Georgia" value="GST2"></zone>
	      <zone name="posix/Atlantic/St_Helena" value="GMT0"></zone>
	      <zone name="posix/Atlantic/Stanley" value="FKST3"></zone>
	      <zone name="posix/Australia/ACT" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Adelaide" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Brisbane" value="AEST-10"></zone>
	      <zone name="posix/Australia/Broken_Hill" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Canberra" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Currie" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Darwin" value="ACST-9:30"></zone>
	      <zone name="posix/Australia/Eucla" value="ACWST-8:45"></zone>
	      <zone name="posix/Australia/Hobart" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/LHI" value="LHST-10:30LHDT-11,M10.1.0,M4.1.0"></zone>
	      <zone name="posix/Australia/Lindeman" value="AEST-10"></zone>
	      <zone name="posix/Australia/Lord_Howe" value="LHST-10:30LHDT-11,M10.1.0,M4.1.0"></zone>
	      <zone name="posix/Australia/Melbourne" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/NSW" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/North" value="ACST-9:30"></zone>
	      <zone name="posix/Australia/Perth" value="AWST-8"></zone>
	      <zone name="posix/Australia/Queensland" value="AEST-10"></zone>
	      <zone name="posix/Australia/South" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Sydney" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Tasmania" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/Victoria" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Australia/West" value="AWST-8"></zone>
	      <zone name="posix/Australia/Yancowinna" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="posix/Brazil/Acre" value="ACT5"></zone>
	      <zone name="posix/Brazil/DeNoronha" value="FNT2"></zone>
	      <zone name="posix/Brazil/East" value="BRT3BRST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="posix/Brazil/West" value="AMT4"></zone>
	      <zone name="posix/CET" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/CST6CDT" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/Atlantic" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/Central" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/East-Saskatchewan" value="CST6"></zone>
	      <zone name="posix/Canada/Eastern" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/Mountain" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/Newfoundland" value="NST3:30NDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/Pacific" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Canada/Saskatchewan" value="CST6"></zone>
	      <zone name="posix/Canada/Yukon" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Chile/Continental" value="CLT3"></zone>
	      <zone name="posix/Chile/EasterIsland" value="EAST5"></zone>
	      <zone name="posix/Cuba" value="CST5CDT,M3.2.0/0,M11.1.0/1"></zone>
	      <zone name="posix/EET" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/EST" value="EST5"></zone>
	      <zone name="posix/EST5EDT" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Egypt" value="EET-2"></zone>
	      <zone name="posix/Eire" value="GMT0IST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Etc/GMT" value="GMT0"></zone>
	      <zone name="posix/Etc/GMT+0" value="GMT0"></zone>
	      <zone name="posix/Etc/GMT+1" value="&lt;GMT+1&gt;1"></zone>
	      <zone name="posix/Etc/GMT+10" value="&lt;GMT+10&gt;10"></zone>
	      <zone name="posix/Etc/GMT+11" value="&lt;GMT+11&gt;11"></zone>
	      <zone name="posix/Etc/GMT+12" value="&lt;GMT+12&gt;12"></zone>
	      <zone name="posix/Etc/GMT+2" value="&lt;GMT+2&gt;2"></zone>
	      <zone name="posix/Etc/GMT+3" value="&lt;GMT+3&gt;3"></zone>
	      <zone name="posix/Etc/GMT+4" value="&lt;GMT+4&gt;4"></zone>
	      <zone name="posix/Etc/GMT+5" value="&lt;GMT+5&gt;5"></zone>
	      <zone name="posix/Etc/GMT+6" value="&lt;GMT+6&gt;6"></zone>
	      <zone name="posix/Etc/GMT+7" value="&lt;GMT+7&gt;7"></zone>
	      <zone name="posix/Etc/GMT+8" value="&lt;GMT+8&gt;8"></zone>
	      <zone name="posix/Etc/GMT+9" value="&lt;GMT+9&gt;9"></zone>
	      <zone name="posix/Etc/GMT-0" value="GMT0"></zone>
	      <zone name="posix/Etc/GMT-1" value="&lt;GMT-1&gt;-1"></zone>
	      <zone name="posix/Etc/GMT-10" value="&lt;GMT-10&gt;-10"></zone>
	      <zone name="posix/Etc/GMT-11" value="&lt;GMT-11&gt;-11"></zone>
	      <zone name="posix/Etc/GMT-12" value="&lt;GMT-12&gt;-12"></zone>
	      <zone name="posix/Etc/GMT-13" value="&lt;GMT-13&gt;-13"></zone>
	      <zone name="posix/Etc/GMT-14" value="&lt;GMT-14&gt;-14"></zone>
	      <zone name="posix/Etc/GMT-2" value="&lt;GMT-2&gt;-2"></zone>
	      <zone name="posix/Etc/GMT-3" value="&lt;GMT-3&gt;-3"></zone>
	      <zone name="posix/Etc/GMT-4" value="&lt;GMT-4&gt;-4"></zone>
	      <zone name="posix/Etc/GMT-5" value="&lt;GMT-5&gt;-5"></zone>
	      <zone name="posix/Etc/GMT-6" value="&lt;GMT-6&gt;-6"></zone>
	      <zone name="posix/Etc/GMT-7" value="&lt;GMT-7&gt;-7"></zone>
	      <zone name="posix/Etc/GMT-8" value="&lt;GMT-8&gt;-8"></zone>
	      <zone name="posix/Etc/GMT-9" value="&lt;GMT-9&gt;-9"></zone>
	      <zone name="posix/Etc/GMT0" value="GMT0"></zone>
	      <zone name="posix/Etc/Greenwich" value="GMT0"></zone>
	      <zone name="posix/Etc/UCT" value="UCT0"></zone>
	      <zone name="posix/Etc/UTC" value="UTC0"></zone>
	      <zone name="posix/Etc/Universal" value="UTC0"></zone>
	      <zone name="posix/Etc/Zulu" value="UTC0"></zone>
	      <zone name="posix/Europe/Amsterdam" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Andorra" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Athens" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Belfast" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Belgrade" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Berlin" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Bratislava" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Brussels" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Bucharest" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Budapest" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Busingen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Chisinau" value="EET-2EEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Copenhagen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Dublin" value="GMT0IST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Gibraltar" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Guernsey" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Helsinki" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Isle_of_Man" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Istanbul" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Jersey" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Kaliningrad" value="EET-2"></zone>
	      <zone name="posix/Europe/Kiev" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Lisbon" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Ljubljana" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/London" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Europe/Luxembourg" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Madrid" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Malta" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Mariehamn" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Minsk" value="MSK-3"></zone>
	      <zone name="posix/Europe/Monaco" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Moscow" value="MSK-3"></zone>
	      <zone name="posix/Europe/Nicosia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Oslo" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Paris" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Podgorica" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Prague" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Riga" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Rome" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Samara" value="SAMT-4"></zone>
	      <zone name="posix/Europe/San_Marino" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Sarajevo" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Simferopol" value="MSK-3"></zone>
	      <zone name="posix/Europe/Skopje" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Sofia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Stockholm" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Tallinn" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Tirane" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Tiraspol" value="EET-2EEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Uzhgorod" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Vaduz" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Vatican" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Vienna" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Vilnius" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Volgograd" value="MSK-3"></zone>
	      <zone name="posix/Europe/Warsaw" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Zagreb" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Europe/Zaporozhye" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/Europe/Zurich" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/GB" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/GB-Eire" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/GMT" value="GMT0"></zone>
	      <zone name="posix/GMT+0" value="GMT0"></zone>
	      <zone name="posix/GMT-0" value="GMT0"></zone>
	      <zone name="posix/GMT0" value="GMT0"></zone>
	      <zone name="posix/Greenwich" value="GMT0"></zone>
	      <zone name="posix/HST" value="HST10"></zone>
	      <zone name="posix/Hongkong" value="HKT-8"></zone>
	      <zone name="posix/Iceland" value="GMT0"></zone>
	      <zone name="posix/Indian/Antananarivo" value="EAT-3"></zone>
	      <zone name="posix/Indian/Chagos" value="IOT-6"></zone>
	      <zone name="posix/Indian/Christmas" value="CXT-7"></zone>
	      <zone name="posix/Indian/Cocos" value="CCT-6:30"></zone>
	      <zone name="posix/Indian/Comoro" value="EAT-3"></zone>
	      <zone name="posix/Indian/Kerguelen" value="TFT-5"></zone>
	      <zone name="posix/Indian/Mahe" value="SCT-4"></zone>
	      <zone name="posix/Indian/Maldives" value="MVT-5"></zone>
	      <zone name="posix/Indian/Mauritius" value="MUT-4"></zone>
	      <zone name="posix/Indian/Mayotte" value="EAT-3"></zone>
	      <zone name="posix/Indian/Reunion" value="RET-4"></zone>
	      <zone name="posix/Iran" value="IRDT"></zone>
	      <zone name="posix/Israel" value="IDDT"></zone>
	      <zone name="posix/Jamaica" value="EST5"></zone>
	      <zone name="posix/Japan" value="JST-9"></zone>
	      <zone name="posix/Kwajalein" value="MHT-12"></zone>
	      <zone name="posix/Libya" value="EET-2"></zone>
	      <zone name="posix/MET" value="MET-1MEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/MST" value="MST7"></zone>
	      <zone name="posix/MST7MDT" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Mexico/BajaNorte" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Mexico/BajaSur" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/Mexico/General" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="posix/NZ" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="posix/NZ-CHAT" value="CHAST-12:45CHADT,M9.5.0/2:45,M4.1.0/3:45"></zone>
	      <zone name="posix/Navajo" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/PRC" value="CST-8"></zone>
	      <zone name="posix/PST8PDT" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/Pacific/Apia" value="WSST-13WSDT,M9.5.0/3,M4.1.0/4"></zone>
	      <zone name="posix/Pacific/Auckland" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="posix/Pacific/Bougainville" value="BST-11"></zone>
	      <zone name="posix/Pacific/Chatham" value="CHAST-12:45CHADT,M9.5.0/2:45,M4.1.0/3:45"></zone>
	      <zone name="posix/Pacific/Chuuk" value="CHUT-10"></zone>
	      <zone name="posix/Pacific/Easter" value="EAST5"></zone>
	      <zone name="posix/Pacific/Efate" value="VUT-11"></zone>
	      <zone name="posix/Pacific/Enderbury" value="PHOT-13"></zone>
	      <zone name="posix/Pacific/Fakaofo" value="TKT-13"></zone>
	      <zone name="posix/Pacific/Fiji" value="FJT-12FJST,M11.1.0,M1.3.0/3"></zone>
	      <zone name="posix/Pacific/Funafuti" value="TVT-12"></zone>
	      <zone name="posix/Pacific/Galapagos" value="GALT6"></zone>
	      <zone name="posix/Pacific/Gambier" value="GAMT9"></zone>
	      <zone name="posix/Pacific/Guadalcanal" value="SBT-11"></zone>
	      <zone name="posix/Pacific/Guam" value="ChST-10"></zone>
	      <zone name="posix/Pacific/Honolulu" value="HST10"></zone>
	      <zone name="posix/Pacific/Johnston" value="HST10"></zone>
	      <zone name="posix/Pacific/Kiritimati" value="LINT-14"></zone>
	      <zone name="posix/Pacific/Kosrae" value="KOST-11"></zone>
	      <zone name="posix/Pacific/Kwajalein" value="MHT-12"></zone>
	      <zone name="posix/Pacific/Majuro" value="MHT-12"></zone>
	      <zone name="posix/Pacific/Marquesas" value="MART9:30"></zone>
	      <zone name="posix/Pacific/Midway" value="SST11"></zone>
	      <zone name="posix/Pacific/Nauru" value="NRT-12"></zone>
	      <zone name="posix/Pacific/Niue" value="NUT11"></zone>
	      <zone name="posix/Pacific/Norfolk" value="NFT-11"></zone>
	      <zone name="posix/Pacific/Noumea" value="NCT-11"></zone>
	      <zone name="posix/Pacific/Pago_Pago" value="SST11"></zone>
	      <zone name="posix/Pacific/Palau" value="PWT-9"></zone>
	      <zone name="posix/Pacific/Pitcairn" value="PST8"></zone>
	      <zone name="posix/Pacific/Pohnpei" value="PONT-11"></zone>
	      <zone name="posix/Pacific/Ponape" value="PONT-11"></zone>
	      <zone name="posix/Pacific/Port_Moresby" value="PGT-10"></zone>
	      <zone name="posix/Pacific/Rarotonga" value="CKT10"></zone>
	      <zone name="posix/Pacific/Saipan" value="ChST-10"></zone>
	      <zone name="posix/Pacific/Samoa" value="SST11"></zone>
	      <zone name="posix/Pacific/Tahiti" value="TAHT10"></zone>
	      <zone name="posix/Pacific/Tarawa" value="GILT-12"></zone>
	      <zone name="posix/Pacific/Tongatapu" value="TOT-13"></zone>
	      <zone name="posix/Pacific/Truk" value="CHUT-10"></zone>
	      <zone name="posix/Pacific/Wake" value="WAKT-12"></zone>
	      <zone name="posix/Pacific/Wallis" value="WFT-12"></zone>
	      <zone name="posix/Pacific/Yap" value="CHUT-10"></zone>
	      <zone name="posix/Poland" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="posix/Portugal" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/ROC" value="CST-8"></zone>
	      <zone name="posix/ROK" value="KST-9"></zone>
	      <zone name="posix/Singapore" value="SGT-8"></zone>
	      <zone name="posix/Turkey" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="posix/UCT" value="UCT0"></zone>
	      <zone name="posix/US/Alaska" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Aleutian" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Arizona" value="MST7"></zone>
	      <zone name="posix/US/Central" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/East-Indiana" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Eastern" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Hawaii" value="HST10"></zone>
	      <zone name="posix/US/Indiana-Starke" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Michigan" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Mountain" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Pacific" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Pacific-New" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="posix/US/Samoa" value="SST11"></zone>
	      <zone name="posix/UTC" value="UTC0"></zone>
	      <zone name="posix/Universal" value="UTC0"></zone>
	      <zone name="posix/W-SU" value="MSK-3"></zone>
	      <zone name="posix/WET" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="posix/Zulu" value="UTC0"></zone>
	      <zone name="posixrules" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Africa/Abidjan" value="GMT0"></zone>
	      <zone name="right/Africa/Accra" value="GMT0"></zone>
	      <zone name="right/Africa/Addis_Ababa" value="EAT-3"></zone>
	      <zone name="right/Africa/Algiers" value="CET-1"></zone>
	      <zone name="right/Africa/Asmara" value="EAT-3"></zone>
	      <zone name="right/Africa/Asmera" value="EAT-3"></zone>
	      <zone name="right/Africa/Bamako" value="GMT0"></zone>
	      <zone name="right/Africa/Bangui" value="WAT-1"></zone>
	      <zone name="right/Africa/Banjul" value="GMT0"></zone>
	      <zone name="right/Africa/Bissau" value="GMT0"></zone>
	      <zone name="right/Africa/Blantyre" value="CAT-2"></zone>
	      <zone name="right/Africa/Brazzaville" value="WAT-1"></zone>
	      <zone name="right/Africa/Bujumbura" value="CAT-2"></zone>
	      <zone name="right/Africa/Cairo" value="EET-2"></zone>
	      <zone name="right/Africa/Casablanca" value="WET0WEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Africa/Ceuta" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Africa/Conakry" value="GMT0"></zone>
	      <zone name="right/Africa/Dakar" value="GMT0"></zone>
	      <zone name="right/Africa/Dar_es_Salaam" value="EAT-3"></zone>
	      <zone name="right/Africa/Djibouti" value="EAT-3"></zone>
	      <zone name="right/Africa/Douala" value="WAT-1"></zone>
	      <zone name="right/Africa/El_Aaiun" value="WET0WEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Africa/Freetown" value="GMT0"></zone>
	      <zone name="right/Africa/Gaborone" value="CAT-2"></zone>
	      <zone name="right/Africa/Harare" value="CAT-2"></zone>
	      <zone name="right/Africa/Johannesburg" value="SAST-2"></zone>
	      <zone name="right/Africa/Juba" value="EAT-3"></zone>
	      <zone name="right/Africa/Kampala" value="EAT-3"></zone>
	      <zone name="right/Africa/Khartoum" value="EAT-3"></zone>
	      <zone name="right/Africa/Kigali" value="CAT-2"></zone>
	      <zone name="right/Africa/Kinshasa" value="WAT-1"></zone>
	      <zone name="right/Africa/Lagos" value="WAT-1"></zone>
	      <zone name="right/Africa/Libreville" value="WAT-1"></zone>
	      <zone name="right/Africa/Lome" value="GMT0"></zone>
	      <zone name="right/Africa/Luanda" value="WAT-1"></zone>
	      <zone name="right/Africa/Lubumbashi" value="CAT-2"></zone>
	      <zone name="right/Africa/Lusaka" value="CAT-2"></zone>
	      <zone name="right/Africa/Malabo" value="WAT-1"></zone>
	      <zone name="right/Africa/Maputo" value="CAT-2"></zone>
	      <zone name="right/Africa/Maseru" value="SAST-2"></zone>
	      <zone name="right/Africa/Mbabane" value="SAST-2"></zone>
	      <zone name="right/Africa/Mogadishu" value="EAT-3"></zone>
	      <zone name="right/Africa/Monrovia" value="GMT0"></zone>
	      <zone name="right/Africa/Nairobi" value="EAT-3"></zone>
	      <zone name="right/Africa/Ndjamena" value="WAT-1"></zone>
	      <zone name="right/Africa/Niamey" value="WAT-1"></zone>
	      <zone name="right/Africa/Nouakchott" value="GMT0"></zone>
	      <zone name="right/Africa/Ouagadougou" value="GMT0"></zone>
	      <zone name="right/Africa/Porto-Novo" value="WAT-1"></zone>
	      <zone name="right/Africa/Sao_Tome" value="GMT0"></zone>
	      <zone name="right/Africa/Timbuktu" value="GMT0"></zone>
	      <zone name="right/Africa/Tripoli" value="EET-2"></zone>
	      <zone name="right/Africa/Tunis" value="CET-1"></zone>
	      <zone name="right/Africa/Windhoek" value="WAT-1WAST,M9.1.0,M4.1.0"></zone>
	      <zone name="right/America/Adak" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Anchorage" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Anguilla" value="AST4"></zone>
	      <zone name="right/America/Antigua" value="AST4"></zone>
	      <zone name="right/America/Araguaina" value="BRT3"></zone>
	      <zone name="right/America/Argentina/Buenos_Aires" value="ART3"></zone>
	      <zone name="right/America/Argentina/Catamarca" value="ART3"></zone>
	      <zone name="right/America/Argentina/ComodRivadavia" value="ART3"></zone>
	      <zone name="right/America/Argentina/Cordoba" value="ART3"></zone>
	      <zone name="right/America/Argentina/Jujuy" value="ART3"></zone>
	      <zone name="right/America/Argentina/La_Rioja" value="ART3"></zone>
	      <zone name="right/America/Argentina/Mendoza" value="ART3"></zone>
	      <zone name="right/America/Argentina/Rio_Gallegos" value="ART3"></zone>
	      <zone name="right/America/Argentina/Salta" value="ART3"></zone>
	      <zone name="right/America/Argentina/San_Juan" value="ART3"></zone>
	      <zone name="right/America/Argentina/San_Luis" value="ART3"></zone>
	      <zone name="right/America/Argentina/Tucuman" value="ART3"></zone>
	      <zone name="right/America/Argentina/Ushuaia" value="ART3"></zone>
	      <zone name="right/America/Aruba" value="AST4"></zone>
	      <zone name="right/America/Asuncion" value="PYT4PYST,M10.1.0/0,M3.4.0/0"></zone>
	      <zone name="right/America/Atikokan" value="EST5"></zone>
	      <zone name="right/America/Atka" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Bahia" value="BRT3"></zone>
	      <zone name="right/America/Bahia_Banderas" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Barbados" value="AST4"></zone>
	      <zone name="right/America/Belem" value="BRT3"></zone>
	      <zone name="right/America/Belize" value="CST6"></zone>
	      <zone name="right/America/Blanc-Sablon" value="AST4"></zone>
	      <zone name="right/America/Boa_Vista" value="AMT4"></zone>
	      <zone name="right/America/Bogota" value="COT5"></zone>
	      <zone name="right/America/Boise" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Buenos_Aires" value="ART3"></zone>
	      <zone name="right/America/Cambridge_Bay" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Campo_Grande" value="AMT4AMST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="right/America/Cancun" value="EST5"></zone>
	      <zone name="right/America/Caracas" value="VET4:30"></zone>
	      <zone name="right/America/Catamarca" value="ART3"></zone>
	      <zone name="right/America/Cayenne" value="GFT3"></zone>
	      <zone name="right/America/Cayman" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Chicago" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Chihuahua" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Coral_Harbour" value="EST5"></zone>
	      <zone name="right/America/Cordoba" value="ART3"></zone>
	      <zone name="right/America/Costa_Rica" value="CST6"></zone>
	      <zone name="right/America/Creston" value="MST7"></zone>
	      <zone name="right/America/Cuiaba" value="AMT4AMST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="right/America/Curacao" value="AST4"></zone>
	      <zone name="right/America/Danmarkshavn" value="GMT0"></zone>
	      <zone name="right/America/Dawson" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Dawson_Creek" value="MST7"></zone>
	      <zone name="right/America/Denver" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Detroit" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Dominica" value="AST4"></zone>
	      <zone name="right/America/Edmonton" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Eirunepe" value="ACT5"></zone>
	      <zone name="right/America/El_Salvador" value="CST6"></zone>
	      <zone name="right/America/Ensenada" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Fort_Nelson" value="MST7"></zone>
	      <zone name="right/America/Fort_Wayne" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Fortaleza" value="BRT3"></zone>
	      <zone name="right/America/Glace_Bay" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Godthab" value="WGST"></zone>
	      <zone name="right/America/Goose_Bay" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Grand_Turk" value="AST4"></zone>
	      <zone name="right/America/Grenada" value="AST4"></zone>
	      <zone name="right/America/Guadeloupe" value="AST4"></zone>
	      <zone name="right/America/Guatemala" value="CST6"></zone>
	      <zone name="right/America/Guayaquil" value="ECT5"></zone>
	      <zone name="right/America/Guyana" value="GYT4"></zone>
	      <zone name="right/America/Halifax" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Havana" value="CST5CDT,M3.2.0/0,M11.1.0/1"></zone>
	      <zone name="right/America/Hermosillo" value="MST7"></zone>
	      <zone name="right/America/Indiana/Indianapolis" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Knox" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Marengo" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Petersburg" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Tell_City" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Vevay" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Vincennes" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indiana/Winamac" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Indianapolis" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Inuvik" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Iqaluit" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Jamaica" value="EST5"></zone>
	      <zone name="right/America/Jujuy" value="ART3"></zone>
	      <zone name="right/America/Juneau" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Kentucky/Louisville" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Kentucky/Monticello" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Knox_IN" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Kralendijk" value="AST4"></zone>
	      <zone name="right/America/La_Paz" value="BOT4"></zone>
	      <zone name="right/America/Lima" value="PET5"></zone>
	      <zone name="right/America/Los_Angeles" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Louisville" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Lower_Princes" value="AST4"></zone>
	      <zone name="right/America/Maceio" value="BRT3"></zone>
	      <zone name="right/America/Managua" value="CST6"></zone>
	      <zone name="right/America/Manaus" value="AMT4"></zone>
	      <zone name="right/America/Marigot" value="AST4"></zone>
	      <zone name="right/America/Martinique" value="AST4"></zone>
	      <zone name="right/America/Matamoros" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Mazatlan" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Mendoza" value="ART3"></zone>
	      <zone name="right/America/Menominee" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Merida" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Metlakatla" value="PST8"></zone>
	      <zone name="right/America/Mexico_City" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Miquelon" value="PMST3PMDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Moncton" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Monterrey" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Montevideo" value="UYT3"></zone>
	      <zone name="right/America/Montreal" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Montserrat" value="AST4"></zone>
	      <zone name="right/America/Nassau" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/New_York" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Nipigon" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Nome" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Noronha" value="FNT2"></zone>
	      <zone name="right/America/North_Dakota/Beulah" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/North_Dakota/Center" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/North_Dakota/New_Salem" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Ojinaga" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Panama" value="EST5"></zone>
	      <zone name="right/America/Pangnirtung" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Paramaribo" value="SRT3"></zone>
	      <zone name="right/America/Phoenix" value="MST7"></zone>
	      <zone name="right/America/Port-au-Prince" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Port_of_Spain" value="AST4"></zone>
	      <zone name="right/America/Porto_Acre" value="ACT5"></zone>
	      <zone name="right/America/Porto_Velho" value="AMT4"></zone>
	      <zone name="right/America/Puerto_Rico" value="AST4"></zone>
	      <zone name="right/America/Rainy_River" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Rankin_Inlet" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Recife" value="BRT3"></zone>
	      <zone name="right/America/Regina" value="CST6"></zone>
	      <zone name="right/America/Resolute" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Rio_Branco" value="ACT5"></zone>
	      <zone name="right/America/Rosario" value="ART3"></zone>
	      <zone name="right/America/Santa_Isabel" value="PST8PDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/America/Santarem" value="BRT3"></zone>
	      <zone name="right/America/Santiago" value="CLT3"></zone>
	      <zone name="right/America/Santo_Domingo" value="AST4"></zone>
	      <zone name="right/America/Sao_Paulo" value="BRT3BRST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="right/America/Scoresbysund" value="EGT1EGST,M3.5.0/0,M10.5.0/1"></zone>
	      <zone name="right/America/Shiprock" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Sitka" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/St_Barthelemy" value="AST4"></zone>
	      <zone name="right/America/St_Johns" value="NST3:30NDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/St_Kitts" value="AST4"></zone>
	      <zone name="right/America/St_Lucia" value="AST4"></zone>
	      <zone name="right/America/St_Thomas" value="AST4"></zone>
	      <zone name="right/America/St_Vincent" value="AST4"></zone>
	      <zone name="right/America/Swift_Current" value="CST6"></zone>
	      <zone name="right/America/Tegucigalpa" value="CST6"></zone>
	      <zone name="right/America/Thule" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Thunder_Bay" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Tijuana" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Toronto" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Tortola" value="AST4"></zone>
	      <zone name="right/America/Vancouver" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Virgin" value="AST4"></zone>
	      <zone name="right/America/Whitehorse" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Winnipeg" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Yakutat" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/America/Yellowknife" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Antarctica/Casey" value="AWST-8"></zone>
	      <zone name="right/Antarctica/Davis" value="DAVT-7"></zone>
	      <zone name="right/Antarctica/DumontDUrville" value="DDUT-10"></zone>
	      <zone name="right/Antarctica/Macquarie" value="MIST-11"></zone>
	      <zone name="right/Antarctica/Mawson" value="MAWT-5"></zone>
	      <zone name="right/Antarctica/McMurdo" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="right/Antarctica/Palmer" value="CLT3"></zone>
	      <zone name="right/Antarctica/Rothera" value="ROTT3"></zone>
	      <zone name="right/Antarctica/South_Pole" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="right/Antarctica/Syowa" value="SYOT-3"></zone>
	      <zone name="right/Antarctica/Troll" value="UTC0CEST-2,M3.5.0/1,M10.5.0/3"></zone>
	      <zone name="right/Antarctica/Vostok" value="VOST-6"></zone>
	      <zone name="right/Arctic/Longyearbyen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Asia/Aden" value="AST-3"></zone>
	      <zone name="right/Asia/Almaty" value="ALMT-6"></zone>
	      <zone name="right/Asia/Amman" value="EEST"></zone>
	      <zone name="right/Asia/Anadyr" value="ANAT-12"></zone>
	      <zone name="right/Asia/Aqtau" value="AQTT-5"></zone>
	      <zone name="right/Asia/Aqtobe" value="AQTT-5"></zone>
	      <zone name="right/Asia/Ashgabat" value="TMT-5"></zone>
	      <zone name="right/Asia/Ashkhabad" value="TMT-5"></zone>
	      <zone name="right/Asia/Baghdad" value="AST-3"></zone>
	      <zone name="right/Asia/Bahrain" value="AST-3"></zone>
	      <zone name="right/Asia/Baku" value="AZT-4AZST,M3.5.0/4,M10.5.0/5"></zone>
	      <zone name="right/Asia/Bangkok" value="ICT-7"></zone>
	      <zone name="right/Asia/Beirut" value="EET-2EEST,M3.5.0/0,M10.5.0/0"></zone>
	      <zone name="right/Asia/Bishkek" value="KGT-6"></zone>
	      <zone name="right/Asia/Brunei" value="BNT-8"></zone>
	      <zone name="right/Asia/Calcutta" value="IST-5:30"></zone>
	      <zone name="right/Asia/Chita" value="IRKT-8"></zone>
	      <zone name="right/Asia/Choibalsan" value="CHOT-8CHOST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="right/Asia/Chongqing" value="CST-8"></zone>
	      <zone name="right/Asia/Chungking" value="CST-8"></zone>
	      <zone name="right/Asia/Colombo" value="IST-5:30"></zone>
	      <zone name="right/Asia/Dacca" value="BDT-6"></zone>
	      <zone name="right/Asia/Damascus" value="EET-2EEST,M3.5.5/0,M10.5.5/0"></zone>
	      <zone name="right/Asia/Dhaka" value="BDT-6"></zone>
	      <zone name="right/Asia/Dili" value="TLT-9"></zone>
	      <zone name="right/Asia/Dubai" value="GST-4"></zone>
	      <zone name="right/Asia/Dushanbe" value="TJT-5"></zone>
	      <zone name="right/Asia/Gaza" value="EEST"></zone>
	      <zone name="right/Asia/Harbin" value="CST-8"></zone>
	      <zone name="right/Asia/Hebron" value="EEST"></zone>
	      <zone name="right/Asia/Ho_Chi_Minh" value="ICT-7"></zone>
	      <zone name="right/Asia/Hong_Kong" value="HKT-8"></zone>
	      <zone name="right/Asia/Hovd" value="HOVT-7HOVST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="right/Asia/Irkutsk" value="IRKT-8"></zone>
	      <zone name="right/Asia/Istanbul" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Asia/Jakarta" value="WIB-7"></zone>
	      <zone name="right/Asia/Jayapura" value="WIT-9"></zone>
	      <zone name="right/Asia/Jerusalem" value="IDDT"></zone>
	      <zone name="right/Asia/Kabul" value="AFT-4:30"></zone>
	      <zone name="right/Asia/Kamchatka" value="PETT-12"></zone>
	      <zone name="right/Asia/Karachi" value="PKT-5"></zone>
	      <zone name="right/Asia/Kashgar" value="XJT-6"></zone>
	      <zone name="right/Asia/Kathmandu" value="NPT-5:45"></zone>
	      <zone name="right/Asia/Katmandu" value="NPT-5:45"></zone>
	      <zone name="right/Asia/Khandyga" value="YAKT-9"></zone>
	      <zone name="right/Asia/Kolkata" value="IST-5:30"></zone>
	      <zone name="right/Asia/Krasnoyarsk" value="KRAT-7"></zone>
	      <zone name="right/Asia/Kuala_Lumpur" value="MYT-8"></zone>
	      <zone name="right/Asia/Kuching" value="MYT-8"></zone>
	      <zone name="right/Asia/Kuwait" value="AST-3"></zone>
	      <zone name="right/Asia/Macao" value="CST-8"></zone>
	      <zone name="right/Asia/Macau" value="CST-8"></zone>
	      <zone name="right/Asia/Magadan" value="MAGT-10"></zone>
	      <zone name="right/Asia/Makassar" value="WITA-8"></zone>
	      <zone name="right/Asia/Manila" value="PHT-8"></zone>
	      <zone name="right/Asia/Muscat" value="GST-4"></zone>
	      <zone name="right/Asia/Nicosia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Asia/Novokuznetsk" value="KRAT-7"></zone>
	      <zone name="right/Asia/Novosibirsk" value="NOVT-6"></zone>
	      <zone name="right/Asia/Omsk" value="OMST-6"></zone>
	      <zone name="right/Asia/Oral" value="ORAT-5"></zone>
	      <zone name="right/Asia/Phnom_Penh" value="ICT-7"></zone>
	      <zone name="right/Asia/Pontianak" value="WIB-7"></zone>
	      <zone name="right/Asia/Pyongyang" value="KST-8:30"></zone>
	      <zone name="right/Asia/Qatar" value="AST-3"></zone>
	      <zone name="right/Asia/Qyzylorda" value="QYZT-6"></zone>
	      <zone name="right/Asia/Rangoon" value="MMT-6:30"></zone>
	      <zone name="right/Asia/Riyadh" value="AST-3"></zone>
	      <zone name="right/Asia/Saigon" value="ICT-7"></zone>
	      <zone name="right/Asia/Sakhalin" value="SAKT-10"></zone>
	      <zone name="right/Asia/Samarkand" value="UZT-5"></zone>
	      <zone name="right/Asia/Seoul" value="KST-9"></zone>
	      <zone name="right/Asia/Shanghai" value="CST-8"></zone>
	      <zone name="right/Asia/Singapore" value="SGT-8"></zone>
	      <zone name="right/Asia/Srednekolymsk" value="SRET-11"></zone>
	      <zone name="right/Asia/Taipei" value="CST-8"></zone>
	      <zone name="right/Asia/Tashkent" value="UZT-5"></zone>
	      <zone name="right/Asia/Tbilisi" value="GET-4"></zone>
	      <zone name="right/Asia/Tehran" value="IRDT"></zone>
	      <zone name="right/Asia/Tel_Aviv" value="IDDT"></zone>
	      <zone name="right/Asia/Thimbu" value="BTT-6"></zone>
	      <zone name="right/Asia/Thimphu" value="BTT-6"></zone>
	      <zone name="right/Asia/Tokyo" value="JST-9"></zone>
	      <zone name="right/Asia/Ujung_Pandang" value="WITA-8"></zone>
	      <zone name="right/Asia/Ulaanbaatar" value="ULAT-8ULAST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="right/Asia/Ulan_Bator" value="ULAT-8ULAST,M3.5.6,M9.5.6/0"></zone>
	      <zone name="right/Asia/Urumqi" value="XJT-6"></zone>
	      <zone name="right/Asia/Ust-Nera" value="VLAT-10"></zone>
	      <zone name="right/Asia/Vientiane" value="ICT-7"></zone>
	      <zone name="right/Asia/Vladivostok" value="VLAT-10"></zone>
	      <zone name="right/Asia/Yakutsk" value="YAKT-9"></zone>
	      <zone name="right/Asia/Yekaterinburg" value="YEKT-5"></zone>
	      <zone name="right/Asia/Yerevan" value="AMT-4"></zone>
	      <zone name="right/Atlantic/Azores" value="AZOT1AZOST,M3.5.0/0,M10.5.0/1"></zone>
	      <zone name="right/Atlantic/Bermuda" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Atlantic/Canary" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Atlantic/Cape_Verde" value="CVT1"></zone>
	      <zone name="right/Atlantic/Faeroe" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Atlantic/Faroe" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Atlantic/Jan_Mayen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Atlantic/Madeira" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Atlantic/Reykjavik" value="GMT0"></zone>
	      <zone name="right/Atlantic/South_Georgia" value="GST2"></zone>
	      <zone name="right/Atlantic/St_Helena" value="GMT0"></zone>
	      <zone name="right/Atlantic/Stanley" value="FKST3"></zone>
	      <zone name="right/Australia/ACT" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Adelaide" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Brisbane" value="AEST-10"></zone>
	      <zone name="right/Australia/Broken_Hill" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Canberra" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Currie" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Darwin" value="ACST-9:30"></zone>
	      <zone name="right/Australia/Eucla" value="ACWST-8:45"></zone>
	      <zone name="right/Australia/Hobart" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/LHI" value="LHST-10:30LHDT-11,M10.1.0,M4.1.0"></zone>
	      <zone name="right/Australia/Lindeman" value="AEST-10"></zone>
	      <zone name="right/Australia/Lord_Howe" value="LHST-10:30LHDT-11,M10.1.0,M4.1.0"></zone>
	      <zone name="right/Australia/Melbourne" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/NSW" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/North" value="ACST-9:30"></zone>
	      <zone name="right/Australia/Perth" value="AWST-8"></zone>
	      <zone name="right/Australia/Queensland" value="AEST-10"></zone>
	      <zone name="right/Australia/South" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Sydney" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Tasmania" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/Victoria" value="AEST-10AEDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Australia/West" value="AWST-8"></zone>
	      <zone name="right/Australia/Yancowinna" value="ACST-9:30ACDT,M10.1.0,M4.1.0/3"></zone>
	      <zone name="right/Brazil/Acre" value="ACT5"></zone>
	      <zone name="right/Brazil/DeNoronha" value="FNT2"></zone>
	      <zone name="right/Brazil/East" value="BRT3BRST,M10.3.0/0,M2.3.0/0"></zone>
	      <zone name="right/Brazil/West" value="AMT4"></zone>
	      <zone name="right/CET" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/CST6CDT" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/Atlantic" value="AST4ADT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/Central" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/East-Saskatchewan" value="CST6"></zone>
	      <zone name="right/Canada/Eastern" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/Mountain" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/Newfoundland" value="NST3:30NDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/Pacific" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Canada/Saskatchewan" value="CST6"></zone>
	      <zone name="right/Canada/Yukon" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Chile/Continental" value="CLT3"></zone>
	      <zone name="right/Chile/EasterIsland" value="EAST5"></zone>
	      <zone name="right/Cuba" value="CST5CDT,M3.2.0/0,M11.1.0/1"></zone>
	      <zone name="right/EET" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/EST" value="EST5"></zone>
	      <zone name="right/EST5EDT" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Egypt" value="EET-2"></zone>
	      <zone name="right/Eire" value="GMT0IST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Etc/GMT" value="GMT0"></zone>
	      <zone name="right/Etc/GMT+0" value="GMT0"></zone>
	      <zone name="right/Etc/GMT+1" value="&lt;GMT+1&gt;1"></zone>
	      <zone name="right/Etc/GMT+10" value="&lt;GMT+10&gt;10"></zone>
	      <zone name="right/Etc/GMT+11" value="&lt;GMT+11&gt;11"></zone>
	      <zone name="right/Etc/GMT+12" value="&lt;GMT+12&gt;12"></zone>
	      <zone name="right/Etc/GMT+2" value="&lt;GMT+2&gt;2"></zone>
	      <zone name="right/Etc/GMT+3" value="&lt;GMT+3&gt;3"></zone>
	      <zone name="right/Etc/GMT+4" value="&lt;GMT+4&gt;4"></zone>
	      <zone name="right/Etc/GMT+5" value="&lt;GMT+5&gt;5"></zone>
	      <zone name="right/Etc/GMT+6" value="&lt;GMT+6&gt;6"></zone>
	      <zone name="right/Etc/GMT+7" value="&lt;GMT+7&gt;7"></zone>
	      <zone name="right/Etc/GMT+8" value="&lt;GMT+8&gt;8"></zone>
	      <zone name="right/Etc/GMT+9" value="&lt;GMT+9&gt;9"></zone>
	      <zone name="right/Etc/GMT-0" value="GMT0"></zone>
	      <zone name="right/Etc/GMT-1" value="&lt;GMT-1&gt;-1"></zone>
	      <zone name="right/Etc/GMT-10" value="&lt;GMT-10&gt;-10"></zone>
	      <zone name="right/Etc/GMT-11" value="&lt;GMT-11&gt;-11"></zone>
	      <zone name="right/Etc/GMT-12" value="&lt;GMT-12&gt;-12"></zone>
	      <zone name="right/Etc/GMT-13" value="&lt;GMT-13&gt;-13"></zone>
	      <zone name="right/Etc/GMT-14" value="&lt;GMT-14&gt;-14"></zone>
	      <zone name="right/Etc/GMT-2" value="&lt;GMT-2&gt;-2"></zone>
	      <zone name="right/Etc/GMT-3" value="&lt;GMT-3&gt;-3"></zone>
	      <zone name="right/Etc/GMT-4" value="&lt;GMT-4&gt;-4"></zone>
	      <zone name="right/Etc/GMT-5" value="&lt;GMT-5&gt;-5"></zone>
	      <zone name="right/Etc/GMT-6" value="&lt;GMT-6&gt;-6"></zone>
	      <zone name="right/Etc/GMT-7" value="&lt;GMT-7&gt;-7"></zone>
	      <zone name="right/Etc/GMT-8" value="&lt;GMT-8&gt;-8"></zone>
	      <zone name="right/Etc/GMT-9" value="&lt;GMT-9&gt;-9"></zone>
	      <zone name="right/Etc/GMT0" value="GMT0"></zone>
	      <zone name="right/Etc/Greenwich" value="GMT0"></zone>
	      <zone name="right/Etc/UCT" value="UCT0"></zone>
	      <zone name="right/Etc/UTC" value="UTC0"></zone>
	      <zone name="right/Etc/Universal" value="UTC0"></zone>
	      <zone name="right/Etc/Zulu" value="UTC0"></zone>
	      <zone name="right/Europe/Amsterdam" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Andorra" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Athens" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Belfast" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Belgrade" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Berlin" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Bratislava" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Brussels" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Bucharest" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Budapest" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Busingen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Chisinau" value="EET-2EEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Copenhagen" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Dublin" value="GMT0IST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Gibraltar" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Guernsey" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Helsinki" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Isle_of_Man" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Istanbul" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Jersey" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Kaliningrad" value="EET-2"></zone>
	      <zone name="right/Europe/Kiev" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Lisbon" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Ljubljana" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/London" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Europe/Luxembourg" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Madrid" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Malta" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Mariehamn" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Minsk" value="MSK-3"></zone>
	      <zone name="right/Europe/Monaco" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Moscow" value="MSK-3"></zone>
	      <zone name="right/Europe/Nicosia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Oslo" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Paris" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Podgorica" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Prague" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Riga" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Rome" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Samara" value="SAMT-4"></zone>
	      <zone name="right/Europe/San_Marino" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Sarajevo" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Simferopol" value="MSK-3"></zone>
	      <zone name="right/Europe/Skopje" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Sofia" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Stockholm" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Tallinn" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Tirane" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Tiraspol" value="EET-2EEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Uzhgorod" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Vaduz" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Vatican" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Vienna" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Vilnius" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Volgograd" value="MSK-3"></zone>
	      <zone name="right/Europe/Warsaw" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Zagreb" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Europe/Zaporozhye" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/Europe/Zurich" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/GB" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/GB-Eire" value="GMT0BST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/GMT" value="GMT0"></zone>
	      <zone name="right/GMT+0" value="GMT0"></zone>
	      <zone name="right/GMT-0" value="GMT0"></zone>
	      <zone name="right/GMT0" value="GMT0"></zone>
	      <zone name="right/Greenwich" value="GMT0"></zone>
	      <zone name="right/HST" value="HST10"></zone>
	      <zone name="right/Hongkong" value="HKT-8"></zone>
	      <zone name="right/Iceland" value="GMT0"></zone>
	      <zone name="right/Indian/Antananarivo" value="EAT-3"></zone>
	      <zone name="right/Indian/Chagos" value="IOT-6"></zone>
	      <zone name="right/Indian/Christmas" value="CXT-7"></zone>
	      <zone name="right/Indian/Cocos" value="CCT-6:30"></zone>
	      <zone name="right/Indian/Comoro" value="EAT-3"></zone>
	      <zone name="right/Indian/Kerguelen" value="TFT-5"></zone>
	      <zone name="right/Indian/Mahe" value="SCT-4"></zone>
	      <zone name="right/Indian/Maldives" value="MVT-5"></zone>
	      <zone name="right/Indian/Mauritius" value="MUT-4"></zone>
	      <zone name="right/Indian/Mayotte" value="EAT-3"></zone>
	      <zone name="right/Indian/Reunion" value="RET-4"></zone>
	      <zone name="right/Iran" value="IRDT"></zone>
	      <zone name="right/Israel" value="IDDT"></zone>
	      <zone name="right/Jamaica" value="EST5"></zone>
	      <zone name="right/Japan" value="JST-9"></zone>
	      <zone name="right/Kwajalein" value="MHT-12"></zone>
	      <zone name="right/Libya" value="EET-2"></zone>
	      <zone name="right/MET" value="MET-1MEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/MST" value="MST7"></zone>
	      <zone name="right/MST7MDT" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Mexico/BajaNorte" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Mexico/BajaSur" value="MST7MDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/Mexico/General" value="CST6CDT,M4.1.0,M10.5.0"></zone>
	      <zone name="right/NZ" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="right/NZ-CHAT" value="CHAST-12:45CHADT,M9.5.0/2:45,M4.1.0/3:45"></zone>
	      <zone name="right/Navajo" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/PRC" value="CST-8"></zone>
	      <zone name="right/PST8PDT" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/Pacific/Apia" value="WSST-13WSDT,M9.5.0/3,M4.1.0/4"></zone>
	      <zone name="right/Pacific/Auckland" value="NZST-12NZDT,M9.5.0,M4.1.0/3"></zone>
	      <zone name="right/Pacific/Bougainville" value="BST-11"></zone>
	      <zone name="right/Pacific/Chatham" value="CHAST-12:45CHADT,M9.5.0/2:45,M4.1.0/3:45"></zone>
	      <zone name="right/Pacific/Chuuk" value="CHUT-10"></zone>
	      <zone name="right/Pacific/Easter" value="EAST5"></zone>
	      <zone name="right/Pacific/Efate" value="VUT-11"></zone>
	      <zone name="right/Pacific/Enderbury" value="PHOT-13"></zone>
	      <zone name="right/Pacific/Fakaofo" value="TKT-13"></zone>
	      <zone name="right/Pacific/Fiji" value="FJT-12FJST,M11.1.0,M1.3.0/3"></zone>
	      <zone name="right/Pacific/Funafuti" value="TVT-12"></zone>
	      <zone name="right/Pacific/Galapagos" value="GALT6"></zone>
	      <zone name="right/Pacific/Gambier" value="GAMT9"></zone>
	      <zone name="right/Pacific/Guadalcanal" value="SBT-11"></zone>
	      <zone name="right/Pacific/Guam" value="ChST-10"></zone>
	      <zone name="right/Pacific/Honolulu" value="HST10"></zone>
	      <zone name="right/Pacific/Johnston" value="HST10"></zone>
	      <zone name="right/Pacific/Kiritimati" value="LINT-14"></zone>
	      <zone name="right/Pacific/Kosrae" value="KOST-11"></zone>
	      <zone name="right/Pacific/Kwajalein" value="MHT-12"></zone>
	      <zone name="right/Pacific/Majuro" value="MHT-12"></zone>
	      <zone name="right/Pacific/Marquesas" value="MART9:30"></zone>
	      <zone name="right/Pacific/Midway" value="SST11"></zone>
	      <zone name="right/Pacific/Nauru" value="NRT-12"></zone>
	      <zone name="right/Pacific/Niue" value="NUT11"></zone>
	      <zone name="right/Pacific/Norfolk" value="NFT-11"></zone>
	      <zone name="right/Pacific/Noumea" value="NCT-11"></zone>
	      <zone name="right/Pacific/Pago_Pago" value="SST11"></zone>
	      <zone name="right/Pacific/Palau" value="PWT-9"></zone>
	      <zone name="right/Pacific/Pitcairn" value="PST8"></zone>
	      <zone name="right/Pacific/Pohnpei" value="PONT-11"></zone>
	      <zone name="right/Pacific/Ponape" value="PONT-11"></zone>
	      <zone name="right/Pacific/Port_Moresby" value="PGT-10"></zone>
	      <zone name="right/Pacific/Rarotonga" value="CKT10"></zone>
	      <zone name="right/Pacific/Saipan" value="ChST-10"></zone>
	      <zone name="right/Pacific/Samoa" value="SST11"></zone>
	      <zone name="right/Pacific/Tahiti" value="TAHT10"></zone>
	      <zone name="right/Pacific/Tarawa" value="GILT-12"></zone>
	      <zone name="right/Pacific/Tongatapu" value="TOT-13"></zone>
	      <zone name="right/Pacific/Truk" value="CHUT-10"></zone>
	      <zone name="right/Pacific/Wake" value="WAKT-12"></zone>
	      <zone name="right/Pacific/Wallis" value="WFT-12"></zone>
	      <zone name="right/Pacific/Yap" value="CHUT-10"></zone>
	      <zone name="right/Poland" value="CET-1CEST,M3.5.0,M10.5.0/3"></zone>
	      <zone name="right/Portugal" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/ROC" value="CST-8"></zone>
	      <zone name="right/ROK" value="KST-9"></zone>
	      <zone name="right/Singapore" value="SGT-8"></zone>
	      <zone name="right/Turkey" value="EET-2EEST,M3.5.0/3,M10.5.0/4"></zone>
	      <zone name="right/UCT" value="UCT0"></zone>
	      <zone name="right/US/Alaska" value="AKST9AKDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Aleutian" value="HST10HDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Arizona" value="MST7"></zone>
	      <zone name="right/US/Central" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/East-Indiana" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Eastern" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Hawaii" value="HST10"></zone>
	      <zone name="right/US/Indiana-Starke" value="CST6CDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Michigan" value="EST5EDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Mountain" value="MST7MDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Pacific" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Pacific-New" value="PST8PDT,M3.2.0,M11.1.0"></zone>
	      <zone name="right/US/Samoa" value="SST11"></zone>
	      <zone name="right/UTC" value="UTC0"></zone>
	      <zone name="right/Universal" value="UTC0"></zone>
	      <zone name="right/W-SU" value="MSK-3"></zone>
	      <zone name="right/WET" value="WET0WEST,M3.5.0/1,M10.5.0"></zone>
	      <zone name="right/Zulu" value="UTC0"></zone>
    </timezones>
  </configuration>
    <configuration name="translate.conf" description="Number Translation Rules">
        <profiles>
            <profile name="US">
	        <rule regex="^\+(\d+)$" replace="$1"></rule>
	        <rule regex="^(1[2-9]\d{2}[2-9]\d{6})$" replace="$1"></rule>
	        <rule regex="^([2-9]\d{2}[2-9]\d{6})$" replace="1$1"></rule>
	        <rule regex="^([2-9]\d{6})$" replace="1${areacode}$1"></rule>
	        <rule regex="^011(\d+)$" replace="$1"></rule>
      </profile>
            <profile name="GB">
	        <rule regex="^\+(\d+)$" replace="$1"></rule>
	        <rule regex="^$" replace="$1"></rule>
      </profile>
            <profile name="HK">
	        <rule regex="\+(\d+)$" replace="$1"></rule>
	        <rule regex="^(852\d{8})$" replace="$1"></rule>
	        <rule regex="^(\d{8})$" replace="852$1"></rule>
      </profile>
    </profiles>
  </configuration>
  <configuration name="tts_commandline.conf" description="TextToSpeech Commandline configuration">
        <settings>
	      <param name="command" value="echo ${text} | text2wave -f ${rate} &gt; ${file}"></param>
    </settings>
  </configuration>
  <configuration name="unicall.conf" description="Unicall Configuration">
      <settings>
          <param name="context" value="default"></param>
          <param name="dialplan" value="XML"></param>
          <param name="suppress-dtmf-tone" value="true"></param>
  </settings>
      <spans>
          <span id="1">
              <param name="protocol-class" value="mfcr2"></param>
              <param name="protocol-variant" value="ar"></param>
              <param name="protocol-end" value="peer"></param>
              <param name="outgoing-allowed" value="true"></param>
              <param name="dialplan" value="XML"></param>
              <param name="context" value="default"></param>
    </span>
          <span id="2">
              <param name="protocol-class" value="mfcr2"></param>
              <param name="protocol-variant" value="ar"></param>
              <param name="protocol-end" value="peer"></param>
              <param name="outgoing-allowed" value="true"></param>
              <param name="dialplan" value="XML"></param>
              <param name="context" value="default"></param>
    </span>
  </spans>
  </configuration>
  <configuration name="unimrcp.conf" description="UniMRCP Client">
      <settings>
          <param name="default-tts-profile" value="voxeo-prophecy8.0-mrcp1"></param>
          <param name="default-asr-profile" value="voxeo-prophecy8.0-mrcp1"></param>
          <param name="log-level" value="DEBUG"></param>
          <param name="enable-profile-events" value="false"></param>
          <param name="max-connection-count" value="100"></param>
          <param name="offer-new-connection" value="1"></param>
          <param name="request-timeout" value="3000"></param>
  </settings>
      <profiles>
        <profile name="loquendo7-mrcp2" version="2">
            <param name="client-ip" value="auto"></param>
            <param name="client-port" value="5090"></param>
            <param name="server-ip" value="10.5.5.152"></param>
            <param name="server-port" value="5060"></param>
            <param name="sip-transport" value="udp"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="4000"></param>
            <param name="rtp-port-max" value="5000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <param name="jsgf-mime-type" value="application/jsgf"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
        <profile name="nuance-mrcp1" version="1">
            <param name="server-ip" value="10.5.5.152"></param>
            <param name="server-port" value="554"></param>
            <param name="resource-location" value=""></param>
            <param name="speechsynth" value="synthesizer"></param>
            <param name="speechrecog" value="recognizer"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="4000"></param>
            <param name="rtp-port-max" value="5000"></param>
            <param name="rtcp" value="1"></param>
            <param name="rtcp-bye" value="2"></param>
            <param name="rtcp-tx-interval" value="5000"></param>
            <param name="rtcp-rx-resolution" value="1000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
        <profile name="nuance5-mrcp1" version="1">
            <param name="server-ip" value="10.5.5.152"></param>
            <param name="server-port" value="4900"></param>
            <param name="resource-location" value="media"></param>
            <param name="speechsynth" value="speechsynthesizer"></param>
            <param name="speechrecog" value="speechrecognizer"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="4000"></param>
            <param name="rtp-port-max" value="5000"></param>
            <param name="rtcp" value="1"></param>
            <param name="rtcp-bye" value="2"></param>
            <param name="rtcp-tx-interval" value="5000"></param>
            <param name="rtcp-rx-resolution" value="1000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
        <profile name="nuance5-mrcp2" version="2">
            <param name="client-ip" value="auto"></param>
            <param name="client-port" value="5090"></param>
            <param name="server-ip" value="10.5.5.152"></param>
            <param name="server-port" value="5060"></param>
            <param name="sip-transport" value="udp"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="4000"></param>
            <param name="rtp-port-max" value="5000"></param>
            <param name="rtcp" value="1"></param>
            <param name="rtcp-bye" value="2"></param>
            <param name="rtcp-tx-interval" value="5000"></param>
            <param name="rtcp-rx-resolution" value="1000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
        <profile name="unimrcpserver-mrcp1" version="1">
            <param name="server-ip" value="10.5.5.152"></param>
            <param name="server-port" value="1554"></param>
            <param name="resource-location" value=""></param>
            <param name="speechsynth" value="speechsynthesizer"></param>
            <param name="speechrecog" value="speechrecognizer"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="4000"></param>
            <param name="rtp-port-max" value="5000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
        <profile name="vestec-mrcp-v1" version="1">
            <param name="server-ip" value="127.0.0.1"></param>
            <param name="server-port" value="1554"></param>
            <param name="resource-location" value=""></param>
            <param name="speechsynth" value="speechsynthesizer"></param>
            <param name="speechrecog" value="speechrecognizer"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="14000"></param>
            <param name="rtp-port-max" value="15000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
        <profile name="voxeo-prophecy8.0-mrcp1" version="1">
            <param name="server-ip" value="99.185.85.31"></param>
            <param name="server-port" value="554"></param>
            <param name="resource-location" value=""></param>
            <param name="speechsynth" value="synthesizer"></param>
            <param name="speechrecog" value="recognizer"></param>
            <param name="rtp-ip" value="auto"></param>
            <param name="rtp-port-min" value="4000"></param>
            <param name="rtp-port-max" value="5000"></param>
            <param name="codecs" value="PCMU PCMA L16/96/8000"></param>
            <synthparams>
    </synthparams>
            <recogparams>
    </recogparams>
  </profile>
  </profiles>
  </configuration>
  <configuration name="v8.conf" description="Google V8 JavaScript Plug-Ins">
      <modules>
  </modules>
  </configuration>
  <configuration name="verto.conf" description="HTML5 Verto Endpoint">
      <settings>
          <param name="debug" value="0"></param>
  </settings>
      <profiles>
          <profile name="default-v4">
              <param name="bind-local" value="46.229.223.143:8081"></param>
              <param name="bind-local" value="46.229.223.143:8082" secure="true"></param>
              <param name="force-register-domain" value="46.229.223.143"></param>
              <param name="secure-combined" value="/etc/freeswitch/tls/wss.pem"></param>
              <param name="secure-chain" value="/etc/freeswitch/tls/wss.pem"></param>
              <param name="userauth" value="true"></param>
              <param name="blind-reg" value="false"></param>
              <param name="mcast-ip" value="224.1.1.1"></param>
              <param name="mcast-port" value="1337"></param>
              <param name="rtp-ip" value="46.229.223.143"></param>
              <param name="local-network" value="localnet.auto"></param>
              <param name="outbound-codec-string" value="opus,vp8"></param>
              <param name="inbound-codec-string" value="opus,vp8"></param>
              <param name="apply-candidate-acl" value="localnet.auto"></param>
              <param name="apply-candidate-acl" value="wan_v4.auto"></param>
              <param name="apply-candidate-acl" value="rfc1918.auto"></param>
              <param name="apply-candidate-acl" value="any_v4.auto"></param>
              <param name="timer-name" value="soft"></param>
    </profile>
          <profile name="default-v6">
              <param name="bind-local" value="[::1]:8081"></param>
              <param name="bind-local" value="[::1]:8082" secure="true"></param>
              <param name="force-register-domain" value="46.229.223.143"></param>
              <param name="secure-combined" value="/etc/freeswitch/tls/wss.pem"></param>
              <param name="secure-chain" value="/etc/freeswitch/tls/wss.pem"></param>
              <param name="userauth" value="true"></param>
              <param name="blind-reg" value="false"></param>
              <param name="rtp-ip" value="::1"></param>
              <param name="outbound-codec-string" value="opus,vp8"></param>
              <param name="inbound-codec-string" value="opus,vp8"></param>
              <param name="apply-candidate-acl" value="wan_v6.auto"></param>
              <param name="apply-candidate-acl" value="rfc1918.auto"></param>
              <param name="apply-candidate-acl" value="any_v6.auto"></param>
              <param name="apply-candidate-acl" value="wan_v4.auto"></param>
              <param name="apply-candidate-acl" value="any_v4.auto"></param>
              <param name="timer-name" value="soft"></param>
    </profile>
  </profiles>
  </configuration>
  <configuration name="voicemail.conf" description="Voicemail">
      <settings>
  </settings>
      <profiles>
          <profile name="default">
              <param name="file-extension" value="wav"></param>
              <param name="terminator-key" value="#"></param>
              <param name="max-login-attempts" value="3"></param>
              <param name="digit-timeout" value="10000"></param>
              <param name="min-record-len" value="3"></param>
              <param name="max-record-len" value="300"></param>
              <param name="max-retries" value="3"></param>
              <param name="tone-spec" value="%(1000, 0, 640)"></param>
              <param name="callback-dialplan" value="XML"></param>
              <param name="callback-context" value="default"></param>
              <param name="play-new-messages-key" value="1"></param>
              <param name="play-saved-messages-key" value="2"></param>
              <param name="login-keys" value="0"></param>
              <param name="main-menu-key" value="0"></param>
              <param name="config-menu-key" value="5"></param>
              <param name="record-greeting-key" value="1"></param>
              <param name="choose-greeting-key" value="2"></param>
              <param name="change-pass-key" value="6"></param>
              <param name="record-name-key" value="3"></param>
              <param name="record-file-key" value="3"></param>
              <param name="listen-file-key" value="1"></param>
              <param name="save-file-key" value="2"></param>
              <param name="delete-file-key" value="7"></param>
              <param name="undelete-file-key" value="8"></param>
              <param name="email-key" value="4"></param>
              <param name="pause-key" value="0"></param>
              <param name="restart-key" value="1"></param>
              <param name="ff-key" value="6"></param>
              <param name="rew-key" value="4"></param>
              <param name="skip-greet-key" value="#"></param>
              <param name="previous-message-key" value="1"></param>
              <param name="next-message-key" value="3"></param>
              <param name="skip-info-key" value="*"></param>
              <param name="repeat-message-key" value="0"></param>
              <param name="record-silence-threshold" value="200"></param>
              <param name="record-silence-hits" value="2"></param>
              <param name="web-template-file" value="web-vm.tpl"></param>
              <param name="db-password-override" value="false"></param>
              <param name="allow-empty-password-auth" value="true"></param>
              <param name="operator-extension" value="operator XML default"></param>
              <param name="operator-key" value="9"></param>
              <param name="vmain-extension" value="vmain XML default"></param>
              <param name="vmain-key" value="*"></param>
              <email>
	          <param name="template-file" value="voicemail.tpl"></param>
	          <param name="notify-template-file" value="notify-voicemail.tpl"></param>
                  <param name="date-fmt" value="%A, %B %d %Y, %I %M %p"></param>
                  <param name="email-from" value="${voicemail_account}@${voicemail_domain}"></param>
      </email>
    </profile>
  </profiles>
  </configuration>
  <configuration name="voicemail_ivr.conf" description="Voicemail IVR">
    <profiles>
	      <profile name="default">
		        <settings>
			          <param name="IVR-Maximum-Attempts" value="3"></param>
			          <param name="IVR-Entry-Timeout" value="3000"></param>
			          <param name="Record-Format" value="wav"></param>
			          <param name="Record-Silence-Hits" value="4"></param>
			          <param name="Record-Silence-Threshold" value="200"></param>
			          <param name="Record-Maximum-Length" value="30"></param>
			          <param name="Exit-Purge" value="true"></param>
			          <param name="Password-Mask" value="XXX."></param>
			          <param name="User-Mask" value="X."></param>
		</settings>
		        <apis>
			          <api name="auth_login" value="vm_fsdb_auth_login"></api>
			          <api name="msg_list" value="vm_fsdb_msg_list"></api>
			          <api name="msg_count" value="vm_fsdb_msg_count"></api>
			          <api name="msg_delete" value="vm_fsdb_msg_delete"></api>
			          <api name="msg_undelete" value="vm_fsdb_msg_undelete"></api>
			          <api name="msg_save" value="vm_fsdb_msg_save"></api>
			          <api name="msg_purge" value="vm_fsdb_msg_purge"></api>
			          <api name="msg_get" value="vm_fsdb_msg_get"></api>
			          <api name="msg_forward" value="vm_fsdb_msg_forward"></api>
			          <api name="pref_greeting_set" value="vm_fsdb_pref_greeting_set"></api>
			          <api name="pref_greeting_get" value="vm_fsdb_pref_greeting_get"></api>
			          <api name="pref_recname_set" value="vm_fsdb_pref_recname_set"></api>
			          <api name="pref_password_set" value="vm_fsdb_pref_password_set"></api>
		</apis>
		        <menus>
			          <menu name="std_authenticate">
			            <phrases>
				              <phrase name="fail_auth" value="fail_auth@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
			</keys>
			</menu>
			          <menu name="std_authenticate_ask_user">
			            <phrases>
				              <phrase name="instructions" value="enter_id@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="#" action="ivrengine:terminate_entry" variable="VM-Key-Terminator"></key>
			</keys>
			</menu>
			          <menu name="std_authenticate_ask_password">
			            <phrases>
				              <phrase name="instructions" value="enter_pass@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="#" action="ivrengine:terminate_entry" variable="VM-Key-Terminator"></key>
			</keys>
			</menu>
			          <menu name="std_main_menu">
			            <settings>
				              <param name="Action-On-New-Message" value="new_msg:std_navigator"></param>
			</settings>
			            <phrases>
				              <phrase name="msg_count" value="message_count@voicemail_ivr"></phrase>
				              <phrase name="say_date" value="say_date_event@voicemail_ivr"></phrase>
				              <phrase name="say_msg_number" value="say_message_number@voicemail_ivr"></phrase>
				              <phrase name="menu_options" value="menu@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="new_msg:std_navigator" variable="VM-Key-Play-New-Messages"></key>
				              <key dtmf="2" action="saved_msg:std_navigator" variable="VM-Key-Play-Saved-Messages"></key>
				              <key dtmf="5" action="menu:std_preference" variable="VM-Key-Config-Menu"></key>
				              <key dtmf="#" action="return" variable="VM-Key-Terminator"></key>
			</keys>
			</menu>
			          <menu name="std_navigator">
			            <settings>
			</settings>
			            <phrases>
				              <phrase name="msg_count" value="message_count@voicemail_ivr"></phrase>
				              <phrase name="say_date" value="say_date_event@voicemail_ivr"></phrase>
				              <phrase name="say_msg_number" value="say_message_number@voicemail_ivr"></phrase>
				              <phrase name="menu_options" value="listen_file_check@voicemail_ivr"></phrase>
				              <phrase name="ack" value="ack@voicemail_ivr"></phrase>
				              <phrase name="play_message" value="play_message@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="skip_intro" variable="VM-Key-Main-Listen-File"></key>
				              <key dtmf="6" action="next_msg" variable="VM-Key-Main-Next-Msg"></key>
				              <key dtmf="4" action="prev_msg"></key>
				              <key dtmf="7" action="delete_msg" variable="VM-Key-Main-Delete-File"></key>
				              <key dtmf="8" action="menu:std_forward" variable="VM-Key-Main-Forward"></key>
				              <key dtmf="2" action="save_msg" variable="VM-Key-Main-Save-File"></key>
				              <key dtmf="5" action="callback" variable="VM-Key-Main-Callback"></key>
				              <key dtmf="#" action="return"></key>
			</keys>
			</menu>
			          <menu name="std_preference">
			            <phrases>
				              <phrase name="menu_options" value="config_menu@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="menu:std_record_greeting_with_slot" variable="VM-Key-Record-Greeting"></key>
				              <key dtmf="2" action="menu:std_select_greeting_slot" variable="VM-Key-Choose-Greeting"></key>
				              <key dtmf="3" action="menu:std_record_name" variable="VM-Key-Record-Name"></key>
				              <key dtmf="6" action="menu:std_set_password" variable="VM-Key-Change-Password"></key>
				              <key dtmf="0" action="return" variable="VM-Key-Main-Menu"></key>
			</keys>
			</menu>
			          <menu name="std_record_greeting">
			            <phrases>
				              <phrase name="instructions" value="record_greeting@voicemail_ivr"></phrase>
				              <phrase name="play_recording" value="play_recording@voicemail_ivr"></phrase>
				              <phrase name="menu_options" value="record_file_check@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="listen" variable="VM-Key-Listen-File"></key>
				              <key dtmf="2" action="save" variable="VM-Key-Save-File"></key>
				              <key dtmf="4" action="rerecord" variable="VM-Key-ReRecord-File"></key>
				              <key dtmf="#" action="skip_instruction"></key>
			</keys>
			</menu>
			          <menu name="std_record_name">
			            <phrases>
				              <phrase name="instructions" value="record_name@voicemail_ivr"></phrase>
				              <phrase name="play_recording" value="play_recording@voicemail_ivr"></phrase>
				              <phrase name="menu_options" value="record_file_check@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="listen" variable="VM-Key-Listen-File"></key>
				              <key dtmf="2" action="save" variable="VM-Key-Save-File"></key>
				              <key dtmf="4" action="rerecord" variable="VM-Key-ReRecord-File"></key>
				              <key dtmf="#" action="skip_instruction"></key>
			</keys>
			</menu>
			          <menu name="std_record_message">
			            <phrases>
				              <phrase name="instructions" value="record_message@voicemail_ivr"></phrase>
				              <phrase name="play_recording" value="play_recording@voicemail_ivr"></phrase>
				              <phrase name="menu_options" value="record_file_check@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="listen" variable="VM-Key-Listen-File"></key>
				              <key dtmf="2" action="save" variable="VM-Key-Save-File"></key>
				              <key dtmf="4" action="rerecord" variable="VM-Key-ReRecord-File"></key>
				              <key dtmf="#" action="skip_instruction"></key>
			</keys>
			</menu>
			          <menu name="std_forward_ask_prepend">
			            <phrases>
				              <phrase name="menu_options" value="forward_ask_prepend@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="1" action="prepend" variable="VM-Key-Prepend"></key>
				              <key dtmf="8" action="forward" variable="VM-Key-Forward"></key>
				              <key dtmf="#" action="return" variable="VM-Key-Return"></key>
			</keys>
			</menu>
			          <menu name="std_forward_ask_extension">
			            <phrases>
				              <phrase name="instructions" value="forward_ask_extension@voicemail_ivr"></phrase>
				              <phrase name="ack" value="ack@voicemail_ivr"></phrase>
				              <phrase name="invalid_extension" value="invalid_extension@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="#" action="ivrengine:terminate_entry" variable="VM-Key-Terminator"></key>
			</keys>
			</menu>
			          <menu name="std_select_greeting_slot">
			            <phrases>
				              <phrase name="instructions" value="choose_greeting@voicemail_ivr"></phrase>
				              <phrase name="invalid_slot" value="choose_greeting_fail@voicemail_ivr"></phrase>
				              <phrase name="selected_slot" value="greeting_selected@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
			</keys>
			</menu>
			          <menu name="std_record_greeting_with_slot">
			            <phrases>
				              <phrase name="instructions" value="choose_greeting@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
			</keys>
			</menu>
			          <menu name="std_set_password">
			            <phrases>
				              <phrase name="instructions" value="enter_pass@voicemail_ivr"></phrase>
			</phrases>
			            <keys>
				              <key dtmf="#" action="ivrengine:terminate_entry" variable="VM-Key-Terminator"></key>
			</keys>
			</menu>
		</menus>
	</profile>
    </profiles>
  </configuration>
  <configuration name="xml_cdr.conf" description="XML CDR CURL logger">
      <settings>
          <param name="log-dir" value=""></param>
          <param name="log-b-leg" value="false"></param>
          <param name="prefix-a-leg" value="true"></param>
          <param name="encode" value="true"></param>
  </settings>
  </configuration>
  <configuration name="xml_curl.conf" description="cURL XML Gateway">
      <bindings>
          <binding name="example">
              <param name="gateway-url" value="http://127.0.0.1:8080/conf/directory" bindings="directory"></param>
              <param name="timeout" value="10"></param>
    </binding>
  </bindings>
  </configuration>
  <configuration name="xml_rpc.conf" description="XML RPC">
      <settings>
          <param name="http-port" value="8080"></param>
          <param name="auth-realm" value="freeswitch"></param>
          <param name="auth-user" value="freeswitch"></param>
          <param name="auth-pass" value="works"></param>
  </settings>
  </configuration>
  <configuration name="xml_scgi.conf" description="SCGI XML Gateway">
      <bindings>
          <binding name="example">
    </binding>
  </bindings>
  </configuration>
  <configuration name="zeroconf.conf" description="Zeroconf Event Handler">
      <settings>
          <param name="publish" value="yes"></param>
          <param name="browse" value="_sip._udp"></param>
  </settings>
  </configuration>
  </section>
`)
}
