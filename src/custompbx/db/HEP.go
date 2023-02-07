package db

import "database/sql"

func createHEPSIPTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS sip_package(
state            VARCHAR DEFAULT '',
auth_val          VARCHAR DEFAULT '',
auth_user         VARCHAR DEFAULT '',
content_length    VARCHAR DEFAULT '',
content_type      VARCHAR DEFAULT '',
from_user         VARCHAR DEFAULT '',
from_host         VARCHAR DEFAULT '',
from_tag          VARCHAR DEFAULT '',
max_forwards      VARCHAR DEFAULT '',
organization     VARCHAR DEFAULT '',
msg              VARCHAR DEFAULT '',
body             VARCHAR DEFAULT '',
to_user           VARCHAR DEFAULT '',
to_host           VARCHAR DEFAULT '',
to_tag            VARCHAR DEFAULT '',
contact_val       VARCHAR DEFAULT '',
contact_user      VARCHAR DEFAULT '',
contact_host      VARCHAR DEFAULT '',
call_id           VARCHAR DEFAULT '',
xCall_id          VARCHAR DEFAULT '',
cseq_method       VARCHAR DEFAULT '',
cseq_val          VARCHAR DEFAULT '',
reason_val        VARCHAR DEFAULT '',
rTPStat_val       VARCHAR DEFAULT '',
via_one           VARCHAR DEFAULT '',
via_one_branch     VARCHAR DEFAULT '',
privacy          VARCHAR DEFAULT '',
remote_partyId_val VARCHAR DEFAULT '',
diversion_val     VARCHAR DEFAULT '',
pAssertedId_val   VARCHAR DEFAULT '',
pai_user          VARCHAR DEFAULT '',
pai_host          VARCHAR DEFAULT '',
user_agent        VARCHAR DEFAULT '',
server           VARCHAR DEFAULT '',
uri_host          VARCHAR DEFAULT '',
uri_raw           VARCHAR DEFAULT '',
uri_user          VARCHAR DEFAULT '',
first_method      VARCHAR DEFAULT '',
first_resp        VARCHAR DEFAULT '',
first_resp_text    VARCHAR DEFAULT '',
contact_port    INTEGER DEFAULT 0,
cseq_digit INTEGER DEFAULT 0,
reason_proto VARCHAR DEFAULT '',
reason_cause VARCHAR DEFAULT '',
reason_text VARCHAR DEFAULT '',

from_val VARCHAR DEFAULT '',
from_name VARCHAR DEFAULT '',
from_party VARCHAR DEFAULT '',
from_screen VARCHAR DEFAULT '',
from_privacy VARCHAR DEFAULT '',
from_uri_scheme VARCHAR DEFAULT '',
from_uri_raw VARCHAR DEFAULT '',
from_uri_user_info VARCHAR DEFAULT '',
from_uri_user VARCHAR DEFAULT '',
from_uri_user_password VARCHAR DEFAULT '',
from_uri_host_info VARCHAR DEFAULT '',
from_uri_host VARCHAR DEFAULT '',
from_uri_port INTEGER DEFAULT 0,
from_uri_secure BOOLEAN DEFAULT FALSE,

to_val VARCHAR DEFAULT '',
to_name VARCHAR DEFAULT '',
to_party VARCHAR DEFAULT '',
to_screen VARCHAR DEFAULT '',
to_privacy VARCHAR DEFAULT '',
to_uri_scheme VARCHAR DEFAULT '',
to_uri_raw VARCHAR DEFAULT '',
to_uri_user_info VARCHAR DEFAULT '',
to_uri_user VARCHAR DEFAULT '',
to_uri_user_password VARCHAR DEFAULT '',
to_uri_host_info VARCHAR DEFAULT '',
to_uri_host VARCHAR DEFAULT '',
to_uri_port INTEGER DEFAULT 0,
to_uri_secure BOOLEAN DEFAULT FALSE,

contact_val VARCHAR DEFAULT '',
contact_name VARCHAR DEFAULT '',
contact_party VARCHAR DEFAULT '',
contact_screen VARCHAR DEFAULT '',
contact_privacy VARCHAR DEFAULT '',
contact_uri_scheme VARCHAR DEFAULT '',
contact_uri_raw VARCHAR DEFAULT '',
contact_uri_user_info VARCHAR DEFAULT '',
contact_uri_user VARCHAR DEFAULT '',
contact_uri_user_password VARCHAR DEFAULT '',
contact_uri_host_info VARCHAR DEFAULT '',
contact_uri_host VARCHAR DEFAULT '',
contact_uri_port INTEGER DEFAULT 0,
contact_uri_secure BOOLEAN DEFAULT FALSE

	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

// Error            error
/*XHeader          []string
RemotePartyId    *RemotePartyId
PAssertedId      *PAssertedId
CallingParty     *CallingPartyInfo
Authorization    *Authorization
Contact          *From

Val     string
Name    string
URI     *URI
Party   string
Screen  string
Privacy string
Params  []*Param
*/
