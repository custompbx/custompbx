package db

import (
	"custompbx/mainStruct"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func InitConfDB() {
	createConfigsTable(db)
	createConfigAclListsTable(db)
	createConfigAclNodesTable(db)
	createConfigCallcenterSettingsTable(db)
	createConfigCallcenterQueuesTable(db)
	createConfigCallcenterQueuesParamsTable(db)
	createConfigCallcenterAgentsTable(db)
	createConfigCallcenterTiersTable(db)
	createConfigCallcenterMembersTable(db)
	createConfigSofiaGlobalParamsTable(db)
	createConfigSofiaProfilesTable(db)
	createConfigSofiaProfileAliasesTable(db)
	createConfigSofiaProfileGatewaysTable(db)
	createConfigSofiaProfileGatewayParamsTable(db)
	createConfigSofiaProfileGatewayVareablesTable(db)
	createConfigSofiaProfileDoaminsTable(db)
	createConfigSofiaProfileSettingsTable(db)
	createConfigCdrPgCsvSettingsTable(db)
	createConfigCdrPgCsvSchemaTable(db)
	createConfigVertoSettingsTable(db)
	createConfigVertoProfilesTable(db)
	createConfigVertoProfileSettingsTable(db)
	createConfigOdbcCdrSettingsTable(db)
	createConfigOdbcCdrTablesTable(db)
	createConfigOdbcCdrTablesFieldsTable(db)
	createConfigLcrSettingsTable(db)
	createConfigLcrProfilesTable(db)
	createConfigLcrProfileSettingsTable(db)
	createConfigShoutSettingsTable(db)
	createConfigRedisSettingsTable(db)
	createConfigNibblebillSettingsTable(db)
	createConfigDbSettingsTable(db)
	createConfigMemcacheSettingsTable(db)
	createConfigAvmdSettingsTable(db)
	createConfigTtsCommandlineSettingsTable(db)
	createConfigCdrMongodbSettingsTable(db)
	createConfigHttpCacheSettingsTable(db)
	createConfigOpusSettingsTable(db)
	createConfigPythonSettingsTable(db)
	createConfigAlsaSettingsTable(db)
	createConfigAmrSettingsTable(db)
	createConfigAmrwbSettingsTable(db)
	createConfigCepstralSettingsTable(db)
	createConfigCidlookupSettingsTable(db)
	createConfigCurlSettingsTable(db)
	createConfigDialplanDirectorySettingsTable(db)
	createConfigEasyrouteSettingsTable(db)
	createConfigErlangEventSettingsTable(db)
	createConfigEventMulticastSettingsTable(db)
	createConfigFaxSettingsTable(db)
	createConfigLuaSettingsTable(db)
	createConfigMongoSettingsTable(db)
	createConfigMsrpSettingsTable(db)
	createConfigOrekaSettingsTable(db)
	createConfigPerlSettingsTable(db)
	createConfigPocketsphinxSettingsTable(db)
	createConfigSangomaCodecSettingsTable(db)
	createConfigSndfileSettingsTable(db)
	createConfigXmlCdrSettingsTable(db)
	createConfigXmlRpcSettingsTable(db)
	createConfigZeroconfSettingsTable(db)
	createConfigPostSwitchSettingsTable(db)
	createConfigPostSwitchCliKeybindingsTable(db)
	createConfigPostSwitchDefaultPtimesTable(db)
	createConfigDistributorListsTable(db)
	createConfigDistributorNodesTable(db)
	createConfigDirectorySettingsTable(db)
	createConfigDirectoryProfilesTable(db)
	createConfigDirectoryProfileSettingsTable(db)
	createConfigFifoSettingsTable(db)
	createConfigFifoFifosTable(db)
	createConfigFifoFifoMembersTable(db)
	createConfigOpalSettingsTable(db)
	createConfigOpalListenersTable(db)
	createConfigOpalListenerSettingsTable(db)
	createConfigOspSettingsTable(db)
	createConfigOspProfilesTable(db)
	createConfigOspProfileSettingsTable(db)
	createConfigUnicallSettingsTable(db)
	createConfigUnicallSpansTable(db)
	createConfigUnicallSpanSettingsTable(db)

	CreateTableByStruct(&mainStruct.ConfigConferenceAdvertiseRooms{})
	CreateTableByStruct(&mainStruct.ConfigConferenceProfiles{})
	CreateTableByStruct(&mainStruct.ConfigConferenceProfilesParams{})
	CreateTableByStruct(&mainStruct.ConfigConferenceCallerControlsGroups{})
	CreateTableByStruct(&mainStruct.ConfigConferenceCallerControlsControls{})
	CreateTableByStruct(&mainStruct.ConfigConferenceChatPermissions{})
	CreateTableByStruct(&mainStruct.ConfigConferenceChatPermissionUsers{})
	CreateTableByStruct(&mainStruct.ConfigConferenceLayoutsGroups{})
	CreateTableByStruct(&mainStruct.ConfigConferenceLayoutsGroupLayouts{})
	CreateTableByStruct(&mainStruct.ConfigConferenceLayouts{})
	CreateTableByStruct(&mainStruct.ConfigConferenceLayoutsImages{})

	CreateTableByStruct(&mainStruct.ModuleTag{})
	CreateTableByStruct(&mainStruct.VoicemailSettingsParameter{})
	CreateTableByStruct(&mainStruct.VoicemailProfile{})
	CreateTableByStruct(&mainStruct.VoicemailProfilesParameter{})
}

func createConfigsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_list(
		id serial NOT NULL PRIMARY KEY,
		name VARCHAR,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE(name, instance_id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigAclListsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_acl_lists(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		name VARCHAR NOT NULL,
		list_default VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigAclNodesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_acl_nodes(
		id serial NOT NULL PRIMARY KEY,
		list_id bigint NOT NULL REFERENCES config_acl_lists (id) ON DELETE CASCADE,
		node_type VARCHAR,
		cidr VARCHAR,
		domain VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		position integer NOT NULL,
		UNIQUE (list_id, cidr, domain),
		CONSTRAINT constraint_list_position UNIQUE (list_id, position)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCallcenterSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_callcenter_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCallcenterQueuesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_callcenter_queues(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		queue_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, queue_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCallcenterQueuesParamsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_callcenter_queues_params(
		id serial NOT NULL PRIMARY KEY,
		queue_id bigint NOT NULL REFERENCES config_callcenter_queues (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (queue_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCallcenterAgentsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS agents (
		name character varying(255) UNIQUE,
		system character varying(255),
		instance_id character varying(255),
		uuid character varying(255),
		type character varying(255),
		contact character varying(255),
		status character varying(255),
		state character varying(255),
		max_no_answer integer DEFAULT 0 NOT NULL,
		wrap_up_time integer DEFAULT 0 NOT NULL,
		reject_delay_time integer DEFAULT 0 NOT NULL,
		busy_delay_time integer DEFAULT 0 NOT NULL,
		no_answer_delay_time integer DEFAULT 0 NOT NULL,
		last_bridge_start integer DEFAULT 0 NOT NULL,
		last_bridge_end integer DEFAULT 0 NOT NULL,
		last_offered_call integer DEFAULT 0 NOT NULL,
		last_status_change integer DEFAULT 0 NOT NULL,
		no_answer_count integer DEFAULT 0 NOT NULL,
		calls_answered integer DEFAULT 0 NOT NULL,
		talk_time integer DEFAULT 0 NOT NULL,
		ready_time integer DEFAULT 0 NOT NULL,
		external_calls_count integer DEFAULT 0 NOT NULL,
		id serial PRIMARY KEY,
		UNIQUE (name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCallcenterTiersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS tiers (
		queue character varying(255),
		agent character varying(255),
		state character varying(255),
		level integer DEFAULT 1 NOT NULL,
		"position" integer DEFAULT 1 NOT NULL,
		id serial PRIMARY KEY,
		UNIQUE (queue, agent)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCallcenterMembersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS members (
	   queue	     VARCHAR(255),
	   instance_id	     VARCHAR(255),
	   uuid	     VARCHAR(255) NOT NULL DEFAULT '',
	   session_uuid     VARCHAR(255) NOT NULL DEFAULT '',
	   cid_number	     VARCHAR(255),
	   cid_name	     VARCHAR(255),
	   system_epoch     INTEGER NOT NULL DEFAULT 0,
	   joined_epoch     INTEGER NOT NULL DEFAULT 0,
	   rejoined_epoch   INTEGER NOT NULL DEFAULT 0,
	   bridge_epoch     INTEGER NOT NULL DEFAULT 0,
	   abandoned_epoch  INTEGER NOT NULL DEFAULT 0,
	   base_score       INTEGER NOT NULL DEFAULT 0,
	   skill_score      INTEGER NOT NULL DEFAULT 0,
	   serving_agent    VARCHAR(255),
	   serving_system   VARCHAR(255),
	   state	     VARCHAR(255) 
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaGlobalParamsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_global_params(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfilesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profiles(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		profile_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, profile_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfileAliasesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profile_aliases(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_sofia_profiles (id) ON DELETE CASCADE,
		alias_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, alias_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfileGatewaysTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profile_gateways(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_sofia_profiles (id) ON DELETE CASCADE,
		gateway_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, gateway_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfileGatewayParamsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profile_gateway_params(
		id serial NOT NULL PRIMARY KEY,
		gateway_id bigint NOT NULL REFERENCES config_sofia_profile_gateways (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (gateway_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfileGatewayVareablesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profile_gateway_variables(
		id serial NOT NULL PRIMARY KEY,
		gateway_id bigint NOT NULL REFERENCES config_sofia_profile_gateways (id) ON DELETE CASCADE,
		var_name VARCHAR NOT NULL,
		var_value VARCHAR DEFAULT NULL,
		var_direction VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (gateway_id, var_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfileDoaminsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profile_domains(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_sofia_profiles (id) ON DELETE CASCADE,
		domain_name VARCHAR NOT NULL,
		domain_alias BOOLEAN NOT NULL DEFAULT FALSE,
		parse BOOLEAN NOT NULL DEFAULT FALSE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, domain_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSofiaProfileSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sofia_profile_settings(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_sofia_profiles (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCdrPgCsvSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_cdr_pg_csv_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCdrPgCsvSchemaTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_cdr_pg_csv_schema(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		var VARCHAR NOT NULL,
		column_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, var)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigVertoSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_verto_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigVertoProfilesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_verto_profiles(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		profile_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, profile_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigVertoProfileSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_verto_profile_params(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_verto_profiles (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		secure VARCHAR NOT NULL DEFAULT '',
		position integer NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		CONSTRAINT constraint_verto_profile_param_position UNIQUE (profile_id, position)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)

	_, err = db.Exec(`
	CREATE UNIQUE INDEX IF NOT EXISTS x_config_verto_profile_params_uniq ON config_verto_profile_params (profile_id, param_name, secure)
    WHERE param_name <> 'apply-candidate-acl';
	`)
	panicErr(err)

	// UNIQUE (profile_id, param_name, secure)
}

func createConfigOdbcCdrSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_odbc_cdr_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOdbcCdrTablesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_odbc_cdr_tables(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		name VARCHAR NOT NULL,
		log_leg VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOdbcCdrTablesFieldsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_odbc_cdr_tables_fields(
		id serial NOT NULL PRIMARY KEY,
		table_id bigint NOT NULL REFERENCES config_odbc_cdr_tables (id) ON DELETE CASCADE,
		name VARCHAR NOT NULL,
		chan_var_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (table_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigLcrSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_lcr_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigLcrProfilesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_lcr_profiles(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		profile_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, profile_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigLcrProfileSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_lcr_profile_params(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_lcr_profiles (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigShoutSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_shout_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigRedisSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_redis_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigNibblebillSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_nibblebill_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDbSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_db_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigMemcacheSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_memcache_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigAvmdSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_avmd_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigTtsCommandlineSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_tts_commandline_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCdrMongodbSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_cdr_mongodb_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigHttpCacheSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_http_cache_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOpusSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_opus_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigPythonSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_python_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigAlsaSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_alsa_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigAmrSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_amr_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigAmrwbSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_amrwb_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCepstralSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_cepstral_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCidlookupSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_cidlookup_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigCurlSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_curl_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDialplanDirectorySettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_dialplan_directory_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigEasyrouteSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_easyroute_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigErlangEventSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_erlang_event_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigEventMulticastSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_event_multicast_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigFaxSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_fax_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigLuaSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_lua_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigMongoSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_mongo_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigMsrpSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_msrp_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOrekaSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_oreka_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigPerlSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_perl_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigPocketsphinxSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_pocketsphinx_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSangomaCodecSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sangoma_codec_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigSndfileSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_sndfile_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigXmlCdrSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_xml_cdr_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigXmlRpcSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_xml_rpc_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigZeroconfSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_zeroconf_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigPostSwitchSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_post_switch_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigPostSwitchCliKeybindingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_post_switch_cli_keybindings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		key_name VARCHAR NOT NULL,
		key_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, key_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigPostSwitchDefaultPtimesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_post_switch_default_ptimes(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		codec_name VARCHAR NOT NULL,
		codec_ptime VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, codec_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDistributorListsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_distributor_lists(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDistributorNodesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_distributor_nodes(
		id serial NOT NULL PRIMARY KEY,
		list_id bigint NOT NULL REFERENCES config_distributor_lists (id) ON DELETE CASCADE,
		name VARCHAR,
		weight VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (list_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDirectorySettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_directory_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDirectoryProfilesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_directory_profiles(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		profile_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, profile_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigDirectoryProfileSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_directory_profile_params(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_directory_profiles (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigFifoSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_fifo_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigFifoFifosTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_fifo_fifos(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		fifo_name VARCHAR NOT NULL,
		importance VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, fifo_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigFifoFifoMembersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_fifo_fifo_members(
		id serial NOT NULL PRIMARY KEY,
		fifo_id bigint NOT NULL REFERENCES config_fifo_fifos (id) ON DELETE CASCADE,
		timeout VARCHAR NOT NULL,
		simo VARCHAR NOT NULL,
		lag VARCHAR DEFAULT NULL,
		body VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE
		
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOpalSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_opal_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOpalListenersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_opal_listeners(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		listener_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, listener_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOpalListenerSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_opal_listener_params(
		id serial NOT NULL PRIMARY KEY,
		listener_id bigint NOT NULL REFERENCES config_opal_listeners (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (listener_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOspSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_osp_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOspProfilesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_osp_profiles(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		profile_name VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, profile_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigUnicallSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_unicall_settings(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigUnicallSpansTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_unicall_spans(
		id serial NOT NULL PRIMARY KEY,
		conf_id bigint NOT NULL REFERENCES config_list (id) ON DELETE CASCADE,
		span_id VARCHAR NOT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (conf_id, span_id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigUnicallSpanSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_unicall_span_params(
		id serial NOT NULL PRIMARY KEY,
		span_id bigint NOT NULL REFERENCES config_unicall_spans (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (span_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createConfigOspProfileSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS config_osp_profile_params(
		id serial NOT NULL PRIMARY KEY,
		profile_id bigint NOT NULL REFERENCES config_osp_profiles (id) ON DELETE CASCADE,
		param_name VARCHAR NOT NULL,
		param_value VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (profile_id, param_name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func SetConf(name string, instanceId int64) (int64, error) {
	sqlReq := "INSERT INTO config_list(name, instance_id) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, name, instanceId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfAclList(confId int64, name, listDefault string) (int64, error) {
	sqlReq := "INSERT INTO config_acl_lists(conf_id, name, list_default) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, listDefault).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfAclListNode(listId int64, nodeType, cidr, domain string) (int64, int64, error) {
	sqlReq := `INSERT INTO config_acl_nodes(list_id, node_type, cidr, domain, position)
							values($1, $2, $3, $4, (SELECT COALESCE((SELECT position + 1 FROM config_acl_nodes WHERE list_id = $1 ORDER BY position DESC LIMIT 1), 1))) returning id, position;`
	var id int64
	var position int64
	err := db.QueryRow(sqlReq, listId, nodeType, cidr, domain).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

func SetConfCallcenterSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_callcenter_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfCallcenterQueue(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_callcenter_queues(conf_id, queue_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfCallcenterQueueParam(queueId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_callcenter_queues_params(queue_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, queueId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfCallcenterAgent(
	name,
	agentType,
	system,
	instanceId,
	uuid,
	contact,
	status,
	state string,
	maxNoAnswer,
	wrapUpTime,
	rejectDelayTime,
	busyDelayTime,
	noAnswerDelayTime,
	lastBridgeStart,
	lastBridgeEnd,
	lastOfferedCall,
	lastStatusChange,
	noAnswerCount,
	callsAnswered,
	talkTime,
	readyTime int64) (int64, error) {
	sqlReq :=
		`INSERT INTO agents(name, type, system, uuid, contact, status, state, max_no_answer, wrap_up_time,
					reject_delay_time, busy_delay_time, no_answer_delay_time, last_bridge_start, last_bridge_end, 
					last_offered_call, last_status_change, no_answer_count, calls_answered, talk_time, ready_time, instance_id)
			values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) returning id;`
	var id int64
	err := db.QueryRow(sqlReq,
		name, agentType, system, uuid, contact, status, state, maxNoAnswer, wrapUpTime, rejectDelayTime, busyDelayTime,
		noAnswerDelayTime, lastBridgeStart, lastBridgeEnd, lastOfferedCall, lastStatusChange, noAnswerCount, callsAnswered,
		talkTime, readyTime, instanceId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfCallcenterTier(queue, agent, state string, position, level int64) (int64, error) {
	sqlReq :=
		`INSERT INTO tiers(queue, agent, position, level, state)
			values($1, $2, $3, $4, $5) returning id;`
	var id int64
	err := db.QueryRow(sqlReq,
		queue, agent, position, level, state).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfCallcenterMember(
	uuid, state, queue, instanceId string, abandonedEpoch, baseScore,
	bridgeEpoch int64, cidName, cidNumber string, joinedEpoch, rejoinedEpoch int64,
	servingAgent, servingSystem, sessionUuid string, skillScore, systemEpoch int64,
) error {
	sqlReq :=
		`INSERT INTO agents(
                   uuid, state, queue, instance_id, abandoned_epoch, base_score,
		bridge_epoch, cid_name, cid_number, joined_epoch, rejoined_epoch,
		serving_agent, serving_system, session_uuid, skill_score, system_epoch
)
			values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) returning id;`
	var id int64
	err := db.QueryRow(sqlReq,
		uuid, state, queue, instanceId, abandonedEpoch, baseScore,
		bridgeEpoch, cidName, cidNumber, joinedEpoch, rejoinedEpoch,
		servingAgent, servingSystem, sessionUuid, skillScore, systemEpoch).Scan(&id)
	return err
}

func SetConfSofiaGlobalSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_global_params(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfSofiaProfile(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_profiles(conf_id, profile_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfSofiaProfileAlias(profileId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_profile_aliases(profile_id, alias_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, profileId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfSofiaProfileGateway(profileId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_profile_gateways(profile_id, gateway_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, profileId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfSofiaProfileGatewayParam(gatewayId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_profile_gateway_params(gateway_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, gatewayId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfSofiaProfileGatewayVar(gatewayId int64, name, value, diretion string) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_profile_gateway_variables(gateway_id, var_name, var_value, var_direction) values($1, $2, $3, $4) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, gatewayId, name, value, diretion).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfSofiaProfileDomain(profileId int64, domainName string, alias, parse bool) (int64, error) {
	sqlReq := "INSERT INTO config_sofia_profile_domains(profile_id, domain_name, domain_alias, parse) values($1, $2, $3, $4) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, profileId, domainName, alias, parse).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

// GET
func GetConfig(name string, instanceId int64) *mainStruct.Configurations {
	var id int64
	var enabled bool
	var configs = &mainStruct.Configurations{}
	sqlReq := `SELECT cl.id AS id, cl.name AS name, cl.enabled FROM config_list cl WHERE instance_id = $1 and name = $2;`

	err := db.QueryRow(sqlReq, instanceId, name).Scan(&id, &enabled)
	if err == nil {
		log.Printf("%+v", err)
		return configs
	}

	switch name {
	case mainStruct.ConfPostLoadSwitch:
		configs.NewPostSwitch(id, enabled)
	case mainStruct.ConfAcl:
		configs.NewAcl(id, enabled)
	case mainStruct.ConfCallcenter:
		configs.NewCallcenter(id, enabled)
	case mainStruct.ConfCdrPgCsv:
		configs.NewCdrPgCsv(id, enabled)
	case mainStruct.ConfOdbcCdr:
		configs.NewOdbcCdr(id, enabled)
	case mainStruct.ConfSofia:
		configs.NewSofia(id, enabled)
	case mainStruct.ConfVerto:
		configs.NewVerto(id, enabled)
	case mainStruct.ConfLcr:
		configs.NewLcr(id, enabled)
	case mainStruct.ConfShout:
		configs.NewShout(id, enabled)
	case mainStruct.ConfRedis:
		configs.NewRedis(id, enabled)
	case mainStruct.ConfNibblebill:
		configs.NewNibblebill(id, enabled)
	case mainStruct.ConfDb:
		configs.NewDb(id, enabled)
	case mainStruct.ConfMemcache:
		configs.NewMemcache(id, enabled)
	case mainStruct.ConfAvmd:
		configs.NewAvmd(id, enabled)
	case mainStruct.ConfTtsCommandline:
		configs.NewTtsCommandline(id, enabled)
	case mainStruct.ConfCdrMongodb:
		configs.NewCdrMongodb(id, enabled)
	case mainStruct.ConfHttpCache:
		configs.NewHttpCache(id, enabled)
	case mainStruct.ConfOpus:
		configs.NewOpus(id, enabled)
	case mainStruct.ConfPython:
		configs.NewPython(id, enabled)
	case mainStruct.ConfAlsa:
		configs.NewAlsa(id, enabled)
	case mainStruct.ConfAmr:
		configs.NewAmr(id, enabled)
	case mainStruct.ConfAmrwb:
		configs.NewAmrwb(id, enabled)
	case mainStruct.ConfCepstral:
		configs.NewCepstral(id, enabled)
	case mainStruct.ConfCidlookup:
		configs.NewCidlookup(id, enabled)
	case mainStruct.ConfCurl:
		configs.NewCurl(id, enabled)
	case mainStruct.ConfDialplanDirectory:
		configs.NewDialplanDirectory(id, enabled)
	case mainStruct.ConfEasyroute:
		configs.NewEasyroute(id, enabled)
	case mainStruct.ConfErlangEvent:
		configs.NewErlangEvent(id, enabled)
	case mainStruct.ConfEventMulticast:
		configs.NewEventMulticast(id, enabled)
	case mainStruct.ConfFax:
		configs.NewFax(id, enabled)
	case mainStruct.ConfLua:
		configs.NewLua(id, enabled)
	case mainStruct.ConfMongo:
		configs.NewMongo(id, enabled)
	case mainStruct.ConfMsrp:
		configs.NewMsrp(id, enabled)
	case mainStruct.ConfOreka:
		configs.NewOreka(id, enabled)
	case mainStruct.ConfPerl:
		configs.NewPerl(id, enabled)
	case mainStruct.ConfPocketsphinx:
		configs.NewPocketsphinx(id, enabled)
	case mainStruct.ConfSangomaCodec:
		configs.NewSangomaCodec(id, enabled)
	case mainStruct.ConfSndfile:
		configs.NewSndfile(id, enabled)
	case mainStruct.ConfXmlCdr:
		configs.NewXmlCdr(id, enabled)
	case mainStruct.ConfXmlRpc:
		configs.NewXmlRpc(id, enabled)
	case mainStruct.ConfZeroconf:
		configs.NewZeroconf(id, enabled)
	case mainStruct.ConfDistributor:
		configs.NewDistributor(id, enabled)
	case mainStruct.ConfDirectory:
		configs.NewDirectory(id, enabled)
	case mainStruct.ConfFifo:
		configs.NewFifo(id, enabled)
	case mainStruct.ConfOpal:
		configs.NewOpal(id, enabled)
	case mainStruct.ConfOsp:
		configs.NewOsp(id, enabled)
	case mainStruct.ConfUnicall:
		configs.NewUnicall(id, enabled)
	case mainStruct.ConfConference:
		configs.NewConference(id, enabled)
	case mainStruct.ConfPostLoadModules:
		configs.NewPostLoadModules(id, enabled)
	case mainStruct.ConfVoicemail:
		configs.NewVoicemail(id, enabled)
	}
	return configs
}

func GetConfigs(configs *mainStruct.Configurations, instanceId int64) {
	sqlReq := `SELECT cl.id AS id, cl.name AS name, cl.enabled FROM config_list cl WHERE instance_id = $1;`
	configList, err := db.Query(sqlReq, instanceId)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer configList.Close()
	for configList.Next() {
		var id int64
		var name string
		var enabled bool

		err := configList.Scan(&id, &name, &enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		switch name {
		case mainStruct.ConfPostLoadSwitch:
			configs.NewPostSwitch(id, enabled)
		case mainStruct.ConfAcl:
			configs.NewAcl(id, enabled)
		case mainStruct.ConfCallcenter:
			configs.NewCallcenter(id, enabled)
		case mainStruct.ConfCdrPgCsv:
			configs.NewCdrPgCsv(id, enabled)
		case mainStruct.ConfOdbcCdr:
			configs.NewOdbcCdr(id, enabled)
		case mainStruct.ConfSofia:
			configs.NewSofia(id, enabled)
		case mainStruct.ConfVerto:
			configs.NewVerto(id, enabled)
		case mainStruct.ConfLcr:
			configs.NewLcr(id, enabled)
		case mainStruct.ConfShout:
			configs.NewShout(id, enabled)
		case mainStruct.ConfRedis:
			configs.NewRedis(id, enabled)
		case mainStruct.ConfNibblebill:
			configs.NewNibblebill(id, enabled)
		case mainStruct.ConfDb:
			configs.NewDb(id, enabled)
		case mainStruct.ConfMemcache:
			configs.NewMemcache(id, enabled)
		case mainStruct.ConfAvmd:
			configs.NewAvmd(id, enabled)
		case mainStruct.ConfTtsCommandline:
			configs.NewTtsCommandline(id, enabled)
		case mainStruct.ConfCdrMongodb:
			configs.NewCdrMongodb(id, enabled)
		case mainStruct.ConfHttpCache:
			configs.NewHttpCache(id, enabled)
		case mainStruct.ConfOpus:
			configs.NewOpus(id, enabled)
		case mainStruct.ConfPython:
			configs.NewPython(id, enabled)
		case mainStruct.ConfAlsa:
			configs.NewAlsa(id, enabled)
		case mainStruct.ConfAmr:
			configs.NewAmr(id, enabled)
		case mainStruct.ConfAmrwb:
			configs.NewAmrwb(id, enabled)
		case mainStruct.ConfCepstral:
			configs.NewCepstral(id, enabled)
		case mainStruct.ConfCidlookup:
			configs.NewCidlookup(id, enabled)
		case mainStruct.ConfCurl:
			configs.NewCurl(id, enabled)
		case mainStruct.ConfDialplanDirectory:
			configs.NewDialplanDirectory(id, enabled)
		case mainStruct.ConfEasyroute:
			configs.NewEasyroute(id, enabled)
		case mainStruct.ConfErlangEvent:
			configs.NewErlangEvent(id, enabled)
		case mainStruct.ConfEventMulticast:
			configs.NewEventMulticast(id, enabled)
		case mainStruct.ConfFax:
			configs.NewFax(id, enabled)
		case mainStruct.ConfLua:
			configs.NewLua(id, enabled)
		case mainStruct.ConfMongo:
			configs.NewMongo(id, enabled)
		case mainStruct.ConfMsrp:
			configs.NewMsrp(id, enabled)
		case mainStruct.ConfOreka:
			configs.NewOreka(id, enabled)
		case mainStruct.ConfPerl:
			configs.NewPerl(id, enabled)
		case mainStruct.ConfPocketsphinx:
			configs.NewPocketsphinx(id, enabled)
		case mainStruct.ConfSangomaCodec:
			configs.NewSangomaCodec(id, enabled)
		case mainStruct.ConfSndfile:
			configs.NewSndfile(id, enabled)
		case mainStruct.ConfXmlCdr:
			configs.NewXmlCdr(id, enabled)
		case mainStruct.ConfXmlRpc:
			configs.NewXmlRpc(id, enabled)
		case mainStruct.ConfZeroconf:
			configs.NewZeroconf(id, enabled)
		case mainStruct.ConfDistributor:
			configs.NewDistributor(id, enabled)
		case mainStruct.ConfDirectory:
			configs.NewDirectory(id, enabled)
		case mainStruct.ConfFifo:
			configs.NewFifo(id, enabled)
		case mainStruct.ConfOpal:
			configs.NewOpal(id, enabled)
		case mainStruct.ConfOsp:
			configs.NewOsp(id, enabled)
		case mainStruct.ConfUnicall:
			configs.NewUnicall(id, enabled)
		case mainStruct.ConfConference:
			configs.NewConference(id, enabled)
		case mainStruct.ConfPostLoadModules:
			configs.NewPostLoadModules(id, enabled)
		case mainStruct.ConfVoicemail:
			configs.NewVoicemail(id, enabled)
		}
	}
}

func GetConfigAclLists(conf *mainStruct.Configurations) {
	if conf.Acl == nil {
		return
	}
	sqlReq := `SELECT cal.id AS id, cal.name AS name, cal.list_default AS default, enabled FROM config_acl_lists cal WHERE conf_id = $1;`
	lists, err := db.Query(sqlReq, conf.Acl.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer lists.Close()
	for lists.Next() {
		var list mainStruct.List
		err := lists.Scan(&list.Id, &list.Name, &list.Default, &list.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		list.Nodes = mainStruct.NewNodes()
		conf.Acl.Lists.Set(&list)
	}
}

func GetConfigAclList(conf *mainStruct.Configurations) {
	if conf.Acl == nil {
		return
	}
	var list mainStruct.List
	sqlReq := `SELECT cal.id AS id, cal.name AS name, cal.list_default AS default, enabled FROM config_acl_lists cal WHERE conf_id = $1;`
	err := db.QueryRow(sqlReq, conf.Acl.Id).Scan(&list.Id, &list.Name, &list.Default, &list.Enabled)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	conf.Acl.Lists.Set(&list)
}

func GetConfigAclListNodes(aclList *mainStruct.List, conf *mainStruct.Configurations) {
	if aclList == nil {
		return
	}
	sqlReq := `SELECT can.id AS id, can.node_type AS nodeType, can.cidr AS cidr, domain AS domain, enabled, position FROM config_acl_nodes can WHERE list_id = $1;`
	nodes, err := db.Query(sqlReq, aclList.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer nodes.Close()
	for nodes.Next() {
		var node mainStruct.Node
		err := nodes.Scan(&node.Id, &node.Type, &node.Cidr, &node.Domain, &node.Enabled, &node.Position)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		node.List = aclList
		aclList.Nodes.Set(&node)
		conf.Acl.Nodes.Set(&node)
	}
}

func MoveAclListNode(extension *mainStruct.Node, newPosition int64) error {
	pos1 := newPosition
	pos2 := extension.Position
	pos3 := newPosition + 1

	if extension.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE config_acl_nodes SET position = (position + 1)*-1 WHERE list_id = $1 AND position > $2`,
		extension.List.Id, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_acl_nodes SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_acl_nodes SET position = $2 WHERE id = $1`,
		extension.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_acl_nodes SET position = (position - 1)*-1 WHERE list_id = $1 AND position > $2`,
		extension.List.Id, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_acl_nodes SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	node := extension.List.Nodes.Props()
	switch extension.Position > newPosition {
	case true:
		for _, v := range node {
			if v.Position >= newPosition && v.Position < extension.Position {
				v.Position = v.Position + 1
			}
		}
	case false:
		for _, v := range node {
			if v.Position > extension.Position && v.Position <= newPosition {
				v.Position = v.Position - 1
			}
		}
	}
	extension.Position = newPosition

	return err
}

func GetConfigCallcenterSettings(conf *mainStruct.Configurations) {
	if conf.Callcenter == nil {
		return
	}
	sqlReq := `SELECT ccs.id AS id, ccs.param_name AS name, ccs.param_value AS value, enabled FROM config_callcenter_settings ccs WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Callcenter.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Callcenter.Settings.Set(&param)
	}
}

func GetConfigCallcenterQueues(conf *mainStruct.Configurations) {
	if conf.Callcenter == nil {
		return
	}
	sqlReq := `SELECT ccq.id AS id, ccq.queue_name AS name, enabled FROM config_callcenter_queues ccq WHERE conf_id = $1;`
	queues, err := db.Query(sqlReq, conf.Callcenter.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	for queues.Next() {
		var queue mainStruct.Queue
		err := queues.Scan(&queue.Id, &queue.Name, &queue.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		queue.Params = mainStruct.NewQueueParams()
		conf.Callcenter.Queues.Set(&queue)
	}
}

func GetConfigCallcenterQueuesParams(queue *mainStruct.Queue, conf *mainStruct.Configurations) {
	if queue == nil || conf.Callcenter == nil {
		return
	}
	sqlReq := `SELECT ccqp.id AS id, ccqp.param_name AS name, ccqp.param_value AS value, enabled FROM config_callcenter_queues_params ccqp WHERE queue_id = $1;`
	params, err := db.Query(sqlReq, queue.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.QueueParam
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		param.Queue = queue
		queue.Params.Set(&param)
		conf.Callcenter.QueueParams.Set(&param)
	}
}

func GetConfigCallcenterAgents(conf *mainStruct.Configurations, directory *mainStruct.DirectoryItems) {
	if conf.Callcenter == nil {
		return
	}
	sqlReq := `SELECT name, type, system, uuid, contact, status, state, max_no_answer, wrap_up_time,
					reject_delay_time, busy_delay_time, no_answer_delay_time, last_bridge_start, last_bridge_end, 
					last_offered_call, last_status_change, no_answer_count, calls_answered, talk_time, ready_time, instance_id, id FROM agents;`
	items, err := db.Query(sqlReq)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	for items.Next() {
		var item mainStruct.Agent
		err := items.Scan(&item.Name, &item.Type, &item.System, &item.Uuid, &item.Contact, &item.Status, &item.State, &item.MaxNoAnswer, &item.WrapUpTime, &item.RejectDelayTime, &item.BusyDelayTime,
			&item.NoAnswerDelayTime, &item.LastBridgeStart, &item.LastBridgeEnd, &item.LastOfferedCall, &item.LastStatusChange, &item.NoAnswerCount, &item.CallsAnswered,
			&item.TalkTime, &item.ReadyTime, &item.InstanceId, &item.Id)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Callcenter.Agents.Set(&item)

		if directory != nil {
			r := regexp.MustCompile(`^(.+)@(.+)$`)
			res := r.FindStringSubmatch(item.Name)

			if len(res) != 3 || res[1] == "" || res[2] == "" {
				continue
			}

			domain := directory.Domains.GetByName(res[2])
			if domain == nil {
				continue
			}
			user := domain.Users.GetByName(res[1])
			if user == nil {
				continue
			}
			user.CCAgent = item.Id
		}
	}
}

func GetConfigCallcenterTiers(conf *mainStruct.Configurations) {
	if conf.Callcenter == nil {
		return
	}
	sqlReq := `SELECT queue, agent, state, level, position, id FROM tiers;`
	items, err := db.Query(sqlReq)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	for items.Next() {
		var item mainStruct.Tier
		err := items.Scan(&item.Queue, &item.Agent, &item.State, &item.Level, &item.Position, &item.Id)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Callcenter.Tiers.Set(&item)
	}
}

func GetConfigSofiaSettings(conf *mainStruct.Configurations) {
	if conf.Sofia == nil {
		return
	}
	sqlReq := `SELECT csgp.id AS id, csgp.param_name AS name, csgp.param_value AS value, enabled FROM config_sofia_global_params csgp WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Sofia.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Sofia.GlobalSettings.Set(&param)
	}
}

func GetConfigSofiaProfiles(conf *mainStruct.Configurations) {
	if conf.Sofia == nil {
		return
	}
	sqlReq := `SELECT id, profile_name, enabled FROM config_sofia_profiles WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Sofia.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.SofiaProfile
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewSofiaProfileParams()
		item.Aliases = mainStruct.NewAliases()
		item.Domains = mainStruct.NewSofiaDomains()
		item.Gateways = mainStruct.NewSofiaGateways()
		conf.Sofia.Profiles.Set(&item)
	}
}

func GetConfigSofiaProfileAliases(profile *mainStruct.SofiaProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Sofia == nil {
		return
	}

	sqlReq := `SELECT id, alias_name, enabled FROM config_sofia_profile_aliases  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.Alias
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Aliases.Set(&item)
		conf.Sofia.ProfileAliases.Set(&item)
	}
}

func GetConfigSofiaProfileDomains(profile *mainStruct.SofiaProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Sofia == nil {
		return
	}

	sqlReq := `SELECT id, domain_name, domain_alias, parse, enabled FROM config_sofia_profile_domains  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.SofiaDomain
		err := items.Scan(&item.Id, &item.Name, &item.Alias, &item.Parse, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Domains.Set(&item)
		conf.Sofia.ProfileDomains.Set(&item)
	}
}

func GetConfigSofiaProfileParams(profile *mainStruct.SofiaProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Sofia == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_sofia_profile_settings  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.SofiaProfileParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Params.Set(&item)
		conf.Sofia.ProfileParams.Set(&item)
	}
}

func GetConfigSofiaProfileGateways(profile *mainStruct.SofiaProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Sofia == nil {
		return
	}

	sqlReq := `SELECT id, gateway_name, enabled FROM config_sofia_profile_gateways  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.SofiaGateway
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Params = mainStruct.NewSofiaGatewayParams()
		item.Vars = mainStruct.NewSofiaGatewayVars()
		item.Profile = profile
		profile.Gateways.Set(&item)
		conf.Sofia.ProfileGateways.Set(&item)
	}
}

func GetConfigSofiaProfileGatewayParams(gateway *mainStruct.SofiaGateway, conf *mainStruct.Configurations) {
	if gateway == nil || conf.Sofia == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_sofia_profile_gateway_params WHERE gateway_id = $1;`
	items, err := db.Query(sqlReq, gateway.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.SofiaGatewayParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Gateway = gateway
		gateway.Params.Set(&item)
		conf.Sofia.GatewayParams.Set(&item)
	}
}

func GetConfigSofiaProfileGatewayVars(gateway *mainStruct.SofiaGateway, conf *mainStruct.Configurations) {
	if gateway == nil || conf.Sofia == nil {
		return
	}

	sqlReq := `SELECT id, var_name, var_value, COALESCE(var_direction, ''), enabled FROM config_sofia_profile_gateway_variables WHERE gateway_id = $1;`
	items, err := db.Query(sqlReq, gateway.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.SofiaGatewayVariable
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Direction, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Gateway = gateway
		gateway.Vars.Set(&item)
		conf.Sofia.GatewayVars.Set(&item)
	}
}

// DELETE
func DelConfig(configId int64) bool {
	sqlReq := `DELETE FROM config_list WHERE id = $1;`

	_, err := db.Exec(sqlReq, configId)
	if err != nil {
		return false
	}
	return true
}

func DelConfigAclList(listId int64) bool {
	sqlReq := `DELETE FROM config_acl_lists WHERE id = $1;`

	_, err := db.Exec(sqlReq, listId)
	if err != nil {
		return false
	}
	return true
}

func DelConfigAclListNode(nodeId int64) bool {
	sqlReq := `DELETE FROM config_acl_nodes WHERE id = $1;`

	_, err := db.Exec(sqlReq, nodeId)
	if err != nil {
		return false
	}
	return true
}

// UPDATE
func UpdateConfig(configId int64, enabled bool) error {
	sqlReq := "UPDATE config_list SET enabled = $1 WHERE id = $2;"
	_, err := db.Exec(sqlReq, enabled, configId)
	if err != nil {
		return err
	}
	return err
}

func UpdateConfigAclLictName(listId int64, newName string) error {
	sqlReq := "UPDATE config_acl_lists SET name = $1 WHERE id = $2;"
	_, err := db.Exec(sqlReq, newName, listId)
	if err != nil {
		return err
	}
	return err
}

func UpdateConfigAclLictDefault(listId int64, allow bool) error {
	defaultValue := "deny"
	if allow {
		defaultValue = "allow"
	}
	sqlReq := "UPDATE config_acl_lists SET list_default = $1 WHERE id = $2;"
	_, err := db.Exec(sqlReq, defaultValue, listId)
	if err != nil {
		return err
	}
	return err
}

func UpdateConfigAclNodeType(nodeId int64, newValue string) error {
	sqlReq := "UPDATE config_acl_nodes SET node_type = $1 WHERE id = $2;"
	_, err := db.Exec(sqlReq, newValue, nodeId)
	if err != nil {
		return err
	}
	return err
}

func UpdateConfigAclNodeCidr(nodeId int64, newValue string) error {
	sqlReq := "UPDATE config_acl_nodes SET cidr = $1 WHERE id = $2;"
	_, err := db.Exec(sqlReq, newValue, nodeId)
	if err != nil {
		return err
	}
	return err
}

func UpdateConfigAclNodeDomain(nodeId int64, newValue string) error {
	sqlReq := "UPDATE config_acl_nodes SET domain = $1 WHERE id = $2;"
	_, err := db.Exec(sqlReq, newValue, nodeId)
	if err != nil {
		return err
	}
	return err
}

func UpdateAclListDefault(userId int64, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE config_acl_lists SET list_default = $1 WHERE id = $2;", newValue, userId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateAclNode(varId int64, nodeType, cidr, domain string) (int64, error) {
	if cidr != "" {
		domain = ""
	}
	res, err := db.Exec("UPDATE config_acl_nodes SET node_type = $1, cidr = $2, domain = $3 WHERE id = $4;", nodeType, cidr, domain, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchAclNode(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_acl_nodes SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelAclNode(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_acl_nodes WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func DelAclList(userId int64) bool {
	_, err := db.Exec(`DELETE FROM config_acl_lists WHERE id = $1;`, userId)
	if err != nil {
		return false
	}
	return true
}

func UpdateAclList(id int64, newName string) error {
	_, err := db.Exec("UPDATE config_acl_lists SET name = $1 WHERE id = $2;", newName, id)
	if err != nil {
		return err
	}
	return err
}

func UpdateSofiaGlobalSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_global_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchSofiaGlobalSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_global_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelSofiaGlobalSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_global_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfSofiaProfileParam(profileId int64, name, value string) (int64, error) {
	sqlReq := `INSERT INTO config_sofia_profile_settings(profile_id, param_name, param_value, enabled)
							values($1, $2, $3, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, profileId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelProfileParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profile_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchProfileParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateProfileParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateProfileGatewayParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_gateway_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchProfileGatewayParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_gateway_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelProfileGatewayParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profile_gateway_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateProfileGatewayVariable(id int64, name, value, direction string) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_gateway_variables SET var_name = $1, var_value = $2, var_direction = $3 WHERE id = $4;", name, value, direction, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchProfileGatewayVariable(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_gateway_variables SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelProfileGatewayVariable(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profile_gateway_variables WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func DelProfileDomain(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profile_domains WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchProfileDomain(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_domains SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateProfileDomain(id int64, name string, alias, parse bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_domains SET domain_name = $1, domain_alias = $2, parse = $3 WHERE id = $4;", name, alias, parse, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelProfileAlias(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profile_aliases WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchProfileAlias(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_aliases SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchProfile(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profiles SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateProfileAlias(id int64, name string) (int64, error) {
	res, err := db.Exec("UPDATE config_sofia_profile_aliases SET alias_name = $1 WHERE id = $2;", name, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateSofiaProfile(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_sofia_profiles SET profile_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelProfile(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profiles WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateSofiaProfileGateway(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_sofia_profile_gateways SET gateway_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelProfileGateway(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sofia_profile_gateways WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchModule(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_list SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func TruncateModuleConfig(id int64) error {
	_, err := db.Exec("DELETE FROM config_list WHERE id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}

func GetConfigCdrPgCsvSettings(conf *mainStruct.Configurations) {
	if conf.CdrPgCsv == nil {
		return
	}
	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_cdr_pg_csv_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.CdrPgCsv.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.CdrPgCsv.Settings.Set(&param)
	}
}

func GetConfigCdrPgCsvSchema(conf *mainStruct.Configurations) {
	if conf.CdrPgCsv == nil {
		return
	}
	sqlReq := `SELECT id, var, column_name, enabled FROM config_cdr_pg_csv_schema WHERE conf_id = $1;`
	fields, err := db.Query(sqlReq, conf.CdrPgCsv.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer fields.Close()
	for fields.Next() {
		var field mainStruct.Field
		err := fields.Scan(&field.Id, &field.Var, &field.Column, &field.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.CdrPgCsv.Schema.Set(&field)
	}
}

func SetConfCdrPgCsvSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_cdr_pg_csv_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfCdrPgCsvSchema(confId int64, variable, column string) (int64, error) {
	sqlReq := "INSERT INTO config_cdr_pg_csv_schema(conf_id, var, column_name) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, variable, column).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SwitchCdrPgCsvParam(id int64, switcher bool) error {
	_, err := db.Exec("UPDATE config_cdr_pg_csv_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	return err
}

func SwitchCdrPgCsvField(id int64, switcher bool) error {
	_, err := db.Exec("UPDATE config_cdr_pg_csv_schema SET enabled = $1 WHERE id = $2;", switcher, id)
	return err
}

func UpdateCdrPgCsvParam(id int64, name, value string) error {
	_, err := db.Exec("UPDATE config_cdr_pg_csv_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	return err
}

func UpdateCdrPgCsvField(id int64, variable, column string) error {
	_, err := db.Exec("UPDATE config_cdr_pg_csv_schema SET var = $1, column_name = $2 WHERE id = $3;", variable, column, id)
	return err
}

func DelCdrPgCsvParam(id int64) error {
	_, err := db.Exec(`DELETE FROM config_cdr_pg_csv_settings WHERE id = $1;`, id)
	return err
}

func DelCdrPgCsvField(id int64) error {
	_, err := db.Exec(`DELETE FROM config_cdr_pg_csv_schema WHERE id = $1;`, id)
	return err
}

func SetConfVertoSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_verto_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfVertoProfile(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_verto_profiles(conf_id, profile_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigVertoSettings(conf *mainStruct.Configurations) {
	if conf.Verto == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_verto_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Verto.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Verto.Settings.Set(&param)
	}
}

func GetConfigVertoProfiles(conf *mainStruct.Configurations) {
	if conf.Verto == nil {
		return
	}
	sqlReq := `SELECT id, profile_name, enabled FROM config_verto_profiles WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Verto.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.VertoProfile
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewVertoProfileParams()
		conf.Verto.Profiles.Set(&item)
	}
}

func GetConfigVertoProfileParams(profile *mainStruct.VertoProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Verto == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, secure, enabled, position FROM config_verto_profile_params WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.VertoProfileParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Secure, &item.Enabled, &item.Position)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Params.Set(&item)
		conf.Verto.ProfileParams.Set(&item)
	}
}

func MoveVertoProfileParam(extension *mainStruct.VertoProfileParam, newPosition int64) error {
	pos1 := newPosition
	pos2 := extension.Position
	pos3 := newPosition + 1

	if extension.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE config_verto_profile_params SET position = (position + 1)*-1 WHERE profile_id = $1 AND position > $2`,
		extension.Profile.Id, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_verto_profile_params SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_verto_profile_params SET position = $2 WHERE id = $1`,
		extension.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_verto_profile_params SET position = (position - 1)*-1 WHERE profile_id = $1 AND position > $2`,
		extension.Profile.Id, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE config_verto_profile_params SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	node := extension.Profile.Params.Props()
	switch extension.Position > newPosition {
	case true:
		for _, v := range node {
			if v.Position >= newPosition && v.Position < extension.Position {
				v.Position = v.Position + 1
			}
		}
	case false:
		for _, v := range node {
			if v.Position > extension.Position && v.Position <= newPosition {
				v.Position = v.Position - 1
			}
		}
	}
	extension.Position = newPosition

	return err
}

func UpdateVertoSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_verto_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchVertoSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_verto_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelVertoSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_verto_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfVertoProfileParam(profileId int64, name, value, secure string) (int64, int64, error) {
	sqlReq := `INSERT INTO config_verto_profile_params(profile_id, param_name, param_value, secure, enabled, position)
							values($1, $2, $3, $4, TRUE, (SELECT COALESCE((SELECT position + 1 FROM config_verto_profile_params WHERE profile_id = $1 ORDER BY position DESC LIMIT 1), 1))) returning id, position;`
	var id int64
	var position int64
	err := db.QueryRow(sqlReq, profileId, name, value, secure).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

func DelVertoProfileParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_verto_profile_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchVertoProfileParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_verto_profile_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateVertoProfileParam(id int64, name, value, secure string) (int64, error) {
	res, err := db.Exec("UPDATE config_verto_profile_params SET param_name = $1, param_value = $2, secure = $3 WHERE id = $4;", name, value, secure, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchVertoProfile(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_verto_profiles SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateVertoProfile(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_verto_profiles SET profile_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelVertoProfile(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_verto_profiles WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateCallcenterSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_callcenter_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchCallcenterSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_callcenter_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCallcenterSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_callcenter_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchCallcenterQueueParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_callcenter_queues_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateCallcenterQueueParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_callcenter_queues_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCallcenterQueueParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_callcenter_queues_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateCallcenterQueue(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_callcenter_queues SET queue_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCallcenterQueue(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_callcenter_queues WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateCallcenterTableColumn(table string, id int64, name, value string) (string, error) {
	var column string
	var dataType string
	err := db.QueryRow(`SELECT column_name, data_type FROM information_schema.columns WHERE table_name = $1 and column_name = $2;`, table, name).Scan(&column, &dataType)
	if err != nil {
		return "", err
	}
	if column == "" {
		return "", errors.New("column not found")
	}
	switch dataType {
	case "integer":
		if value == "" {
			value = "0"
		} else {
			re := regexp.MustCompile("^[0-9]+$")
			res := re.FindAllString(value, 1)
			if len(res) != 1 {
				return "", errors.New("not a number")
			}
			value = res[0]
		}
	case "boolean":
		if value != "true" {
			value = "false"
		}

	}
	query := fmt.Sprintf("UPDATE %s SET %s = $1 WHERE id = $2;", table, column)
	_, err = db.Exec(query, value, id) //TODO
	if err != nil {
		return "", err
	}
	return value, nil
}

func DelCallcenterAgent(id int64) bool {
	_, err := db.Exec(`DELETE FROM agents WHERE id = $1;`, id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func DelCallcenterTier(id int64) bool {
	_, err := db.Exec(`DELETE FROM tiers WHERE id = $1;`, id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func DelCallcenterMember(uuid string) bool {
	_, err := db.Exec(`DELETE FROM agents WHERE uuid = $1;`, uuid)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func SetConfOdbcCdrSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_odbc_cdr_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfOdbcCdrTable(confId int64, name, logLeg string) (int64, error) {
	sqlReq := "INSERT INTO config_odbc_cdr_tables(conf_id, name, log_leg) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, logLeg).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfOdbcCdrTableField(tableId int64, chanVarName, logLeg string) (int64, error) {
	sqlReq := "INSERT INTO config_odbc_cdr_tables_fields(table_id, name, chan_var_name) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, tableId, chanVarName, logLeg).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigOdbcCdrSettings(conf *mainStruct.Configurations) {
	if conf.OdbcCdr == nil {
		return
	}
	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_odbc_cdr_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.OdbcCdr.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.OdbcCdr.Settings.Set(&param)
	}
}

func GetConfigOdbcCdrTables(conf *mainStruct.Configurations) {
	if conf.OdbcCdr == nil {
		return
	}
	sqlReq := `SELECT id, name, log_leg, enabled FROM config_odbc_cdr_tables WHERE conf_id = $1;`
	tables, err := db.Query(sqlReq, conf.OdbcCdr.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer tables.Close()
	for tables.Next() {
		var table mainStruct.Table
		err := tables.Scan(&table.Id, &table.Name, &table.LogLeg, &table.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		table.Fields = mainStruct.NewOdbcFields()
		conf.OdbcCdr.Tables.Set(&table)
	}
}

func GetConfigOdbcCdrTableFields(table *mainStruct.Table, conf *mainStruct.Configurations) {
	if table == nil || conf.OdbcCdr == nil {
		return
	}

	sqlReq := `SELECT id, name, chan_var_name, enabled FROM config_odbc_cdr_tables_fields  WHERE table_id = $1;`
	items, err := db.Query(sqlReq, table.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.ODBCField
		err := items.Scan(&item.Id, &item.Name, &item.ChanVarName, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Table = table
		table.Fields.Set(&item)
		conf.OdbcCdr.TableFields.Set(&item)
	}
}

func UpdateOdbcCdrParam(id int64, name, value string) error {
	_, err := db.Exec("UPDATE config_odbc_cdr_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	return err
}

func SwitchOdbcCdrParam(id int64, switcher bool) error {
	_, err := db.Exec("UPDATE config_odbc_cdr_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	return err
}

func DelOdbcCdrParam(id int64) error {
	_, err := db.Exec(`DELETE FROM config_odbc_cdr_settings WHERE id = $1;`, id)
	return err
}

func UpdateOdbcCdrTable(id int64, name, logLeg string) error {
	_, err := db.Exec("UPDATE config_odbc_cdr_tables SET name = $1, log_leg = $2 WHERE id = $3;", name, logLeg, id)
	return err
}

func SwitchOdbcCdrTable(id int64, switcher bool) error {
	_, err := db.Exec("UPDATE config_odbc_cdr_tables SET enabled = $1 WHERE id = $2;", switcher, id)
	return err
}

func DelOdbcCdrTable(id int64) error {
	_, err := db.Exec(`DELETE FROM config_odbc_cdr_tables WHERE id = $1;`, id)
	return err
}

func SwitchOdbcCdrField(id int64, switcher bool) error {
	_, err := db.Exec("UPDATE config_odbc_cdr_tables_fields SET enabled = $1 WHERE id = $2;", switcher, id)
	return err
}

func UpdateOdbcCdrField(id int64, name, chanVarName string) error {
	_, err := db.Exec("UPDATE config_odbc_cdr_tables_fields SET name = $1, chan_var_name = $2 WHERE id = $3;", name, chanVarName, id)
	return err
}

func DelOdbcCdrField(id int64) error {
	_, err := db.Exec(`DELETE FROM config_odbc_cdr_tables_fields WHERE id = $1;`, id)
	return err
}

func SetConfLcrSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_lcr_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfLcrProfile(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_lcr_profiles(conf_id, profile_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigLcrSettings(conf *mainStruct.Configurations) {
	if conf.Lcr == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_lcr_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Lcr.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Lcr.Settings.Set(&param)
	}
}

func GetConfigLcrProfiles(conf *mainStruct.Configurations) {
	if conf.Lcr == nil {
		return
	}
	sqlReq := `SELECT id, profile_name, enabled FROM config_lcr_profiles WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Lcr.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.LcrProfile
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewLcrProfileParams()
		conf.Lcr.Profiles.Set(&item)
	}
}

func GetConfigLcrProfileParams(profile *mainStruct.LcrProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Lcr == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_lcr_profile_params  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.LcrProfileParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Params.Set(&item)
		conf.Lcr.ProfileParams.Set(&item)
	}
}

func UpdateLcrSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_lcr_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchLcrSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_lcr_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelLcrSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_lcr_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfLcrProfileParam(profileId int64, name, value string) (int64, error) {
	sqlReq := `INSERT INTO config_lcr_profile_params(profile_id, param_name, param_value, enabled)
							values($1, $2, $3, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, profileId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelLcrProfileParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_lcr_profile_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchLcrProfileParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_lcr_profile_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateLcrProfileParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_lcr_profile_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchLcrProfile(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_lcr_profiles SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateLcrProfile(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_lcr_profiles SET profile_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelLcrProfile(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_lcr_profiles WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfShoutSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_shout_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigShoutSettings(conf *mainStruct.Configurations) {
	if conf.Shout == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_shout_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Shout.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Shout.Settings.Set(&param)
	}
}

func UpdateShoutSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_shout_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchShoutSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_shout_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelShoutSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_shout_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfRedisSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_redis_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigRedisSettings(conf *mainStruct.Configurations) {
	if conf.Redis == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_redis_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Redis.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Redis.Settings.Set(&param)
	}
}

func UpdateRedisSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_redis_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchRedisSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_redis_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelRedisSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_redis_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfNibblebillSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_nibblebill_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigNibblebillSettings(conf *mainStruct.Configurations) {
	if conf.Nibblebill == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_nibblebill_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Nibblebill.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Nibblebill.Settings.Set(&param)
	}
}

func UpdateNibblebillSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_nibblebill_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchNibblebillSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_nibblebill_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelNibblebillSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_nibblebill_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfDbSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_db_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigDbSettings(conf *mainStruct.Configurations) {
	if conf.Db == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_db_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Db.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Db.Settings.Set(&param)
	}
}

func UpdateDbSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_db_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchDbSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_db_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelDbSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_db_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfMemcacheSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_memcache_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigMemcacheSettings(conf *mainStruct.Configurations) {
	if conf.Memcache == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_memcache_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Memcache.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Memcache.Settings.Set(&param)
	}
}

func UpdateMemcacheSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_memcache_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchMemcacheSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_memcache_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelMemcacheSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_memcache_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfAvmdSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_avmd_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigAvmdSettings(conf *mainStruct.Configurations) {
	if conf.Avmd == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_avmd_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Avmd.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Avmd.Settings.Set(&param)
	}
}

func UpdateAvmdSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_avmd_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchAvmdSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_avmd_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelAvmdSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_avmd_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfTtsCommandlineSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_tts_commandline_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigTtsCommandlineSettings(conf *mainStruct.Configurations) {
	if conf.TtsCommandline == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_tts_commandline_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.TtsCommandline.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.TtsCommandline.Settings.Set(&param)
	}
}

func UpdateTtsCommandlineSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_tts_commandline_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchTtsCommandlineSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_tts_commandline_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelTtsCommandlineSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_tts_commandline_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfCdrMongodbSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_cdr_mongodb_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigCdrMongodbSettings(conf *mainStruct.Configurations) {
	if conf.CdrMongodb == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_cdr_mongodb_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.CdrMongodb.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.CdrMongodb.Settings.Set(&param)
	}
}

func UpdateCdrMongodbSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_cdr_mongodb_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchCdrMongodbSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_cdr_mongodb_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCdrMongodbSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_cdr_mongodb_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfHttpCacheSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_http_cache_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigHttpCacheSettings(conf *mainStruct.Configurations) {
	if conf.HttpCache == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_http_cache_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.HttpCache.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.HttpCache.Settings.Set(&param)
	}
}

func UpdateHttpCacheSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_http_cache_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchHttpCacheSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_http_cache_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelHttpCacheSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_http_cache_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfOpusSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_opus_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigOpusSettings(conf *mainStruct.Configurations) {
	if conf.Opus == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_opus_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Opus.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Opus.Settings.Set(&param)
	}
}

func UpdateOpusSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_opus_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchOpusSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_opus_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelOpusSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_opus_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPythonSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_python_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigPythonSettings(conf *mainStruct.Configurations) {
	if conf.Python == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_python_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Python.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Python.Settings.Set(&param)
	}
}

func UpdatePythonSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_python_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchPythonSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_python_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelPythonSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_python_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfAlsaSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_alsa_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigAlsaSettings(conf *mainStruct.Configurations) {
	if conf.Alsa == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_alsa_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Alsa.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Alsa.Settings.Set(&param)
	}
}

func UpdateAlsaSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_alsa_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchAlsaSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_alsa_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelAlsaSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_alsa_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfAmrSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_amr_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigAmrSettings(conf *mainStruct.Configurations) {
	if conf.Amr == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_amr_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Amr.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Amr.Settings.Set(&param)
	}
}

func UpdateAmrSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_amr_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchAmrSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_amr_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelAmrSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_amr_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfAmrwbSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_amrwb_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigAmrwbSettings(conf *mainStruct.Configurations) {
	if conf.Amrwb == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_amrwb_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Amrwb.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Amrwb.Settings.Set(&param)
	}
}

func UpdateAmrwbSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_amrwb_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchAmrwbSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_amrwb_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelAmrwbSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_amrwb_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfCepstralSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_cepstral_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigCepstralSettings(conf *mainStruct.Configurations) {
	if conf.Cepstral == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_cepstral_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Cepstral.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Cepstral.Settings.Set(&param)
	}
}

func UpdateCepstralSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_cepstral_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchCepstralSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_cepstral_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCepstralSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_cepstral_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfCidlookupSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_cidlookup_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigCidlookupSettings(conf *mainStruct.Configurations) {
	if conf.Cidlookup == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_cidlookup_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Cidlookup.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Cidlookup.Settings.Set(&param)
	}
}

func UpdateCidlookupSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_cidlookup_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchCidlookupSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_cidlookup_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCidlookupSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_cidlookup_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfCurlSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_curl_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigCurlSettings(conf *mainStruct.Configurations) {
	if conf.Curl == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_curl_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Curl.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Curl.Settings.Set(&param)
	}
}

func UpdateCurlSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_curl_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchCurlSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_curl_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelCurlSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_curl_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfDialplanDirectorySetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_dialplan_directory_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigDialplanDirectorySettings(conf *mainStruct.Configurations) {
	if conf.DialplanDirectory == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_dialplan_directory_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.DialplanDirectory.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.DialplanDirectory.Settings.Set(&param)
	}
}

func UpdateDialplanDirectorySetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_dialplan_directory_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchDialplanDirectorySetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_dialplan_directory_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelDialplanDirectorySetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_dialplan_directory_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfEasyrouteSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_easyroute_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigEasyrouteSettings(conf *mainStruct.Configurations) {
	if conf.Easyroute == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_easyroute_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Easyroute.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Easyroute.Settings.Set(&param)
	}
}

func UpdateEasyrouteSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_easyroute_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchEasyrouteSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_easyroute_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelEasyrouteSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_easyroute_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfErlangEventSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_erlang_event_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigErlangEventSettings(conf *mainStruct.Configurations) {
	if conf.ErlangEvent == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_erlang_event_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.ErlangEvent.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.ErlangEvent.Settings.Set(&param)
	}
}

func UpdateErlangEventSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_erlang_event_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchErlangEventSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_erlang_event_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelErlangEventSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_erlang_event_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfEventMulticastSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_event_multicast_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigEventMulticastSettings(conf *mainStruct.Configurations) {
	if conf.EventMulticast == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_event_multicast_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.EventMulticast.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.EventMulticast.Settings.Set(&param)
	}
}

func UpdateEventMulticastSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_event_multicast_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchEventMulticastSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_event_multicast_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelEventMulticastSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_event_multicast_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfFaxSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_fax_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigFaxSettings(conf *mainStruct.Configurations) {
	if conf.Fax == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_fax_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Fax.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Fax.Settings.Set(&param)
	}
}

func UpdateFaxSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_fax_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchFaxSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_fax_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelFaxSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_fax_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfLuaSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_lua_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigLuaSettings(conf *mainStruct.Configurations) {
	if conf.Lua == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_lua_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Lua.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Lua.Settings.Set(&param)
	}
}

func UpdateLuaSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_lua_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchLuaSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_lua_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelLuaSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_lua_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfMongoSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_mongo_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigMongoSettings(conf *mainStruct.Configurations) {
	if conf.Mongo == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_mongo_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Mongo.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Mongo.Settings.Set(&param)
	}
}

func UpdateMongoSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_mongo_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchMongoSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_mongo_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelMongoSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_mongo_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfMsrpSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_msrp_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigMsrpSettings(conf *mainStruct.Configurations) {
	if conf.Msrp == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_msrp_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Msrp.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Msrp.Settings.Set(&param)
	}
}

func UpdateMsrpSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_msrp_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchMsrpSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_msrp_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelMsrpSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_msrp_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfOrekaSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_oreka_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigOrekaSettings(conf *mainStruct.Configurations) {
	if conf.Oreka == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_oreka_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Oreka.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Oreka.Settings.Set(&param)
	}
}

func UpdateOrekaSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_oreka_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchOrekaSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_oreka_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelOrekaSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_oreka_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPerlSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_perl_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigPerlSettings(conf *mainStruct.Configurations) {
	if conf.Perl == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_perl_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Perl.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Perl.Settings.Set(&param)
	}
}

func UpdatePerlSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_perl_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchPerlSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_perl_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelPerlSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_perl_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPocketsphinxSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_pocketsphinx_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigPocketsphinxSettings(conf *mainStruct.Configurations) {
	if conf.Pocketsphinx == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_pocketsphinx_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Pocketsphinx.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Pocketsphinx.Settings.Set(&param)
	}
}

func UpdatePocketsphinxSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_pocketsphinx_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchPocketsphinxSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_pocketsphinx_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelPocketsphinxSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_pocketsphinx_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfSangomaCodecSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_sangoma_codec_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigSangomaCodecSettings(conf *mainStruct.Configurations) {
	if conf.SangomaCodec == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_sangoma_codec_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.SangomaCodec.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.SangomaCodec.Settings.Set(&param)
	}
}

func UpdateSangomaCodecSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_sangoma_codec_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchSangomaCodecSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sangoma_codec_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelSangomaCodecSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sangoma_codec_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfSndfileSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_sndfile_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigSndfileSettings(conf *mainStruct.Configurations) {
	if conf.Sndfile == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_sndfile_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Sndfile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Sndfile.Settings.Set(&param)
	}
}

func UpdateSndfileSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_sndfile_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchSndfileSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_sndfile_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelSndfileSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_sndfile_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfXmlCdrSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_xml_cdr_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigXmlCdrSettings(conf *mainStruct.Configurations) {
	if conf.XmlCdr == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_xml_cdr_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.XmlCdr.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.XmlCdr.Settings.Set(&param)
	}
}

func UpdateXmlCdrSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_xml_cdr_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchXmlCdrSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_xml_cdr_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelXmlCdrSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_xml_cdr_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfXmlRpcSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_xml_rpc_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigXmlRpcSettings(conf *mainStruct.Configurations) {
	if conf.XmlRpc == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_xml_rpc_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.XmlRpc.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.XmlRpc.Settings.Set(&param)
	}
}

func UpdateXmlRpcSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_xml_rpc_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchXmlRpcSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_xml_rpc_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelXmlRpcSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_xml_rpc_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfZeroconfSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_zeroconf_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigZeroconfSettings(conf *mainStruct.Configurations) {
	if conf.Zeroconf == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_zeroconf_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Zeroconf.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Zeroconf.Settings.Set(&param)
	}
}

func UpdateZeroconfSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_zeroconf_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchZeroconfSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_zeroconf_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelZeroconfSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_zeroconf_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPostSwitchSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_post_switch_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigPostSwitchSettings(conf *mainStruct.Configurations) {
	if conf.PostSwitch == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_post_switch_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.PostSwitch.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.PostSwitch.Settings.Set(&param)
	}
}

func UpdatePostSwitchSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_post_switch_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchPostSwitchSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_post_switch_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelPostSwitchSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_post_switch_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPostSwitchCliKeybinding(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_post_switch_cli_keybindings(conf_id, key_name, key_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigPostSwitchCliKeybindings(conf *mainStruct.Configurations) {
	if conf.PostSwitch == nil {
		return
	}
	sqlReq := `SELECT id AS id, key_name AS name, key_value AS value, enabled FROM config_post_switch_cli_keybindings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.PostSwitch.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.PostSwitch.CliKeybindings.Set(&param)
	}
}

func UpdatePostSwitchCliKeybinding(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_post_switch_cli_keybindings SET key_name = $1, key_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchPostSwitchCliKeybinding(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_post_switch_cli_keybindings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelPostSwitchCliKeybinding(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_post_switch_cli_keybindings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPostSwitchDefaultPtime(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_post_switch_default_ptimes(conf_id, codec_name, codec_ptime) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigPostSwitchDefaultPtimes(conf *mainStruct.Configurations) {
	if conf.PostSwitch == nil {
		return
	}
	sqlReq := `SELECT id AS id, codec_name AS name, codec_ptime AS ptime, enabled FROM config_post_switch_default_ptimes WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.PostSwitch.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.DefaultPtime
		err := params.Scan(&param.Id, &param.Name, &param.Ptime, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.PostSwitch.DefaultPtimes.Set(&param)
	}
}

func UpdatePostSwitchDefaultPtime(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_post_switch_default_ptimes SET codec_name = $1, codec_ptime = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchPostSwitchDefaultPtime(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_post_switch_default_ptimes SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelPostSwitchDefaultPtime(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_post_switch_default_ptimes WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfPostLoadModule(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_post_load_modules(conf_id, param_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

/*
func GetConfigPostLoadModules(conf *mainStruct.Configurations) {
	//if conf.PostLoadSwitch == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, enabled FROM config_post_load_modules WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.PostLoadSwitch.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.PostLoadSwitch.Settings.Set(&param)
	}
}
*/

func SetConfDistributorList(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_distributor_lists(conf_id, name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfDistributorListNode(listId int64, name, weight string) (int64, error) {
	sqlReq := `INSERT INTO config_distributor_nodes(list_id, name, weight)
							values($1, $2, $3) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, listId, name, weight).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigDistributorLists(conf *mainStruct.Configurations) {
	if conf.Distributor == nil {
		return
	}
	sqlReq := `SELECT id AS id, name AS name, enabled FROM config_distributor_lists WHERE conf_id = $1;`
	lists, err := db.Query(sqlReq, conf.Distributor.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer lists.Close()
	for lists.Next() {
		var list mainStruct.DistributorList
		err := lists.Scan(&list.Id, &list.Name, &list.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		list.Nodes = mainStruct.NewDistributorNodes()
		conf.Distributor.Lists.Set(&list)
	}
}

func GetConfigDistributorListNodes(distributorList *mainStruct.DistributorList, conf *mainStruct.Configurations) {
	if distributorList == nil {
		return
	}
	sqlReq := `SELECT id AS id, name AS name, weight AS weight, enabled FROM config_distributor_nodes WHERE list_id = $1;`
	nodes, err := db.Query(sqlReq, distributorList.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer nodes.Close()
	for nodes.Next() {
		var node mainStruct.DistributorNode
		err := nodes.Scan(&node.Id, &node.Name, &node.Weight, &node.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		node.List = distributorList
		distributorList.Nodes.Set(&node)
		conf.Distributor.Nodes.Set(&node)
	}
}

func UpdateDistributorListDefault(userId int64, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE config_distributor_lists SET list_default = $1 WHERE id = $2;", newValue, userId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateDistributorNode(varId int64, name, weight string) (int64, error) {
	res, err := db.Exec("UPDATE config_distributor_nodes SET name = $1, weight = $2 WHERE id = $3;", name, weight, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchDistributorNode(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_distributor_nodes SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelDistributorNode(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_distributor_nodes WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func DelDistributorList(userId int64) bool {
	_, err := db.Exec(`DELETE FROM config_distributor_lists WHERE id = $1;`, userId)
	if err != nil {
		return false
	}
	return true
}

func UpdateDistributorList(id int64, newName string) error {
	_, err := db.Exec("UPDATE config_distributor_lists SET name = $1 WHERE id = $2;", newName, id)
	if err != nil {
		return err
	}
	return err
}

func SetConfDirectorySetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_directory_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfDirectoryProfile(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_directory_profiles(conf_id, profile_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigDirectorySettings(conf *mainStruct.Configurations) {
	if conf.Directory == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_directory_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Directory.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Directory.Settings.Set(&param)
	}
}

func GetConfigDirectoryProfiles(conf *mainStruct.Configurations) {
	if conf.Directory == nil {
		return
	}
	sqlReq := `SELECT id, profile_name, enabled FROM config_directory_profiles WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Directory.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.DirectoryProfile
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewDirectoryProfileParams()
		conf.Directory.Profiles.Set(&item)
	}
}

func GetConfigDirectoryProfileParams(profile *mainStruct.DirectoryProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Directory == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_directory_profile_params  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.DirectoryProfileParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Params.Set(&item)
		conf.Directory.ProfileParams.Set(&item)
	}
}

func UpdateDirectorySetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_directory_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchDirectorySetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_directory_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelDirectorySetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_directory_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfDirectoryProfileParam(profileId int64, name, value string) (int64, error) {
	sqlReq := `INSERT INTO config_directory_profile_params(profile_id, param_name, param_value, enabled)
							values($1, $2, $3, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, profileId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelDirectoryProfileParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_directory_profile_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchDirectoryProfileParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_directory_profile_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateDirectoryProfileParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_directory_profile_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchDirectoryProfile(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_directory_profiles SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateDirectoryProfile(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_directory_profiles SET profile_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelDirectoryProfile(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_directory_profiles WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfFifoSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_fifo_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfFifoFifo(confId int64, name, importance string) (int64, error) {
	sqlReq := "INSERT INTO config_fifo_fifos(conf_id, fifo_name, importance) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, importance).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func UpdateFifoFifoImportance(profileId int64, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE config_fifo_fifos SET importance = $1 WHERE id = $2;", newValue, profileId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func GetConfigFifoSettings(conf *mainStruct.Configurations) {
	if conf.Fifo == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_fifo_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Fifo.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Fifo.Settings.Set(&param)
	}
}

func GetConfigFifoFifos(conf *mainStruct.Configurations) {
	if conf.Fifo == nil {
		return
	}
	sqlReq := `SELECT id, fifo_name, importance, enabled FROM config_fifo_fifos WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Fifo.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.FifoFifo
		err := items.Scan(&item.Id, &item.Name, &item.Importance, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewFifoFifoParams()
		conf.Fifo.Fifos.Set(&item)
	}
}

func GetConfigFifoFifoParams(fifo *mainStruct.FifoFifo, conf *mainStruct.Configurations) {
	if fifo == nil || conf.Fifo == nil {
		return
	}

	sqlReq := `SELECT id, timeout, simo, lag, body, enabled FROM config_fifo_fifo_members  WHERE fifo_id = $1;`
	items, err := db.Query(sqlReq, fifo.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.FifoFifoMember
		err := items.Scan(&item.Id, &item.Timeout, &item.Simo, &item.Lag, &item.Body, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Fifo = fifo
		fifo.Params.Set(&item)
		conf.Fifo.FifoParams.Set(&item)
	}
}

func UpdateFifoSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_fifo_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchFifoSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_fifo_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelFifoSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_fifo_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfFifoFifoParam(fifoId int64, timeout, simo, lag, body string) (int64, error) {
	sqlReq := `INSERT INTO config_fifo_fifo_members(fifo_id, timeout, simo, lag, body, enabled)
							values($1, $2, $3, $4, $5, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, fifoId, timeout, simo, lag, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelFifoFifoParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_fifo_fifo_members WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchFifoFifoParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_fifo_fifo_members SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateFifoFifoParam(id int64, timeout, simo, lag, body string) (int64, error) {
	res, err := db.Exec("UPDATE config_fifo_fifo_members SET timeout = $1, simo = $2, lag = $3, body = $4 WHERE id = $5;", timeout, simo, lag, body, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchFifoFifo(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_fifo_fifos SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateFifoFifo(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_fifo_fifos SET fifo_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelFifoFifo(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_fifo_fifos WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfOpalSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_opal_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfOpalListener(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_opal_listeners(conf_id, listener_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigOpalSettings(conf *mainStruct.Configurations) {
	if conf.Opal == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_opal_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Opal.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Opal.Settings.Set(&param)
	}
}

func GetConfigOpalListeners(conf *mainStruct.Configurations) {
	if conf.Opal == nil {
		return
	}
	sqlReq := `SELECT id, listener_name, enabled FROM config_opal_listeners WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Opal.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.OpalListener
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewOpalListenerParams()
		conf.Opal.Listeners.Set(&item)
	}
}

func GetConfigOpalListenerParams(listener *mainStruct.OpalListener, conf *mainStruct.Configurations) {
	if listener == nil || conf.Opal == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_opal_listener_params  WHERE listener_id = $1;`
	items, err := db.Query(sqlReq, listener.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.OpalListenerParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Listener = listener
		listener.Params.Set(&item)
		conf.Opal.ListenerParams.Set(&item)
	}
}

func UpdateOpalSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_opal_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchOpalSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_opal_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelOpalSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_opal_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfOpalListenerParam(listenerId int64, name, value string) (int64, error) {
	sqlReq := `INSERT INTO config_opal_listener_params(listener_id, param_name, param_value, enabled)
							values($1, $2, $3, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, listenerId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelOpalListenerParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_opal_listener_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchOpalListenerParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_opal_listener_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateOpalListenerParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_opal_listener_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchOpalListener(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_opal_listeners SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateOpalListener(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_opal_listeners SET listener_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelOpalListener(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_opal_listeners WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfOspSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_osp_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfOspProfile(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_osp_profiles(conf_id, profile_name) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigOspSettings(conf *mainStruct.Configurations) {
	if conf.Osp == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_osp_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Osp.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Osp.Settings.Set(&param)
	}
}

func GetConfigOspProfiles(conf *mainStruct.Configurations) {
	if conf.Osp == nil {
		return
	}
	sqlReq := `SELECT id, profile_name, enabled FROM config_osp_profiles WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Osp.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.OspProfile
		err := items.Scan(&item.Id, &item.Name, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewOspProfileParams()
		conf.Osp.Profiles.Set(&item)
	}
}

func GetConfigOspProfileParams(profile *mainStruct.OspProfile, conf *mainStruct.Configurations) {
	if profile == nil || conf.Osp == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_osp_profile_params  WHERE profile_id = $1;`
	items, err := db.Query(sqlReq, profile.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.OspProfileParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Profile = profile
		profile.Params.Set(&item)
		conf.Osp.ProfileParams.Set(&item)
	}
}

func UpdateOspSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_osp_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchOspSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_osp_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelOspSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_osp_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfOspProfileParam(profileId int64, name, value string) (int64, error) {
	sqlReq := `INSERT INTO config_osp_profile_params(profile_id, param_name, param_value, enabled)
							values($1, $2, $3, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, profileId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelOspProfileParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_osp_profile_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchOspProfileParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_osp_profile_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateOspProfileParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_osp_profile_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchOspProfile(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_osp_profiles SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateOspProfile(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_osp_profiles SET profile_name = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelOspProfile(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_osp_profiles WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfUnicallSetting(confId int64, name, value string) (int64, error) {
	sqlReq := "INSERT INTO config_unicall_settings(conf_id, param_name, param_value) values($1, $2, $3) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConfUnicallSpan(confId int64, name string) (int64, error) {
	sqlReq := "INSERT INTO config_unicall_spans(conf_id, span_id) values($1, $2) returning id;"
	var id int64
	err := db.QueryRow(sqlReq, confId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetConfigUnicallSettings(conf *mainStruct.Configurations) {
	if conf.Unicall == nil {
		return
	}
	sqlReq := `SELECT id AS id, param_name AS name, param_value AS value, enabled FROM config_unicall_settings WHERE conf_id = $1;`
	params, err := db.Query(sqlReq, conf.Unicall.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.Param
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		conf.Unicall.Settings.Set(&param)
	}
}

func GetConfigUnicallSpans(conf *mainStruct.Configurations) {
	if conf.Unicall == nil {
		return
	}
	sqlReq := `SELECT id, span_id, enabled FROM config_unicall_spans WHERE conf_id = $1;`
	items, err := db.Query(sqlReq, conf.Unicall.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.UnicallSpan
		err := items.Scan(&item.Id, &item.SpanId, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		item.Params = mainStruct.NewUnicallSpanParams()
		conf.Unicall.Spans.Set(&item)
	}
}

func GetConfigUnicallSpanParams(span *mainStruct.UnicallSpan, conf *mainStruct.Configurations) {
	if span == nil || conf.Unicall == nil {
		return
	}

	sqlReq := `SELECT id, param_name, param_value, enabled FROM config_unicall_span_params  WHERE span_id = $1;`
	items, err := db.Query(sqlReq, span.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.UnicallSpanParam
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		item.Span = span
		span.Params.Set(&item)
		conf.Unicall.SpanParams.Set(&item)
	}
}

func UpdateUnicallSetting(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_unicall_settings SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchUnicallSetting(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_unicall_settings SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelUnicallSetting(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_unicall_settings WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SetConfUnicallSpanParam(spanId int64, name, value string) (int64, error) {
	sqlReq := `INSERT INTO config_unicall_span_params(span_id, param_name, param_value, enabled)
							values($1, $2, $3, TRUE) returning id;`
	var id int64
	err := db.QueryRow(sqlReq, spanId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DelUnicallSpanParam(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_unicall_span_params WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func SwitchUnicallSpanParam(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_unicall_span_params SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateUnicallSpanParam(id int64, name, value string) (int64, error) {
	res, err := db.Exec("UPDATE config_unicall_span_params SET param_name = $1, param_value = $2 WHERE id = $3;", name, value, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchUnicallSpan(id int64, switcher bool) (int64, error) {
	res, err := db.Exec("UPDATE config_unicall_spans SET enabled = $1 WHERE id = $2;", switcher, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateUnicallSpan(varId int64, name string) (int64, error) {
	if name == "" {
		return 0, errors.New("no new name")
	}
	res, err := db.Exec("UPDATE config_unicall_spans SET span_id = $1 WHERE id = $2;", name, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DelUnicallSpan(id int64) bool {
	_, err := db.Exec(`DELETE FROM config_unicall_spans WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func InsertFields(confId int64, tableName string, item interface{}) (int64, error) {
	if tableName == "" {
		return 0, errors.New("no table name")
	}
	var sqlReq string
	var id int64
	var err error
	names, values, fkey := StrutToSliceOfSqlNamesAndValues(item)
	names, i := FilterSlice("id", names)
	if i >= 0 {
		copy(values[i:], values[i+1:])
		values = values[:len(values)-1]
	}
	names = append([]string{fkey}, names...)
	values = append([]interface{}{confId}, values...)
	sqlReq = fmt.Sprintf("INSERT INTO %s(%s) values(%s) returning id;", tableName, strings.Join(names, ", "), ValuesPlaceholders(names))

	err = db.QueryRow(sqlReq, values...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetDataById(rowId int64, tableName string, item interface{}) []interface{} {
	if tableName == "" {
		return []interface{}{}
	}
	sqlFields, fields, fkey := StrutToSliceOfFieldAddressAndNames(item)
	if len(sqlFields) == 0 || fkey == "" || rowId == 0 {
		return []interface{}{}
	}
	sqlReq := fmt.Sprintf(`SELECT %s FROM %s WHERE %s = $1;`, strings.Join(sqlFields[:], ", "), tableName, fkey)
	params, err := db.Query(sqlReq, rowId)
	if err != nil {
		log.Printf("%+v", err)
		return nil
	}
	defer params.Close()
	var res []interface{}
	for params.Next() {
		err := params.Scan(fields...)
		if err != nil {
			log.Printf("%+v", err)
			return nil
		}
		indirect := reflect.Indirect(reflect.ValueOf(item))
		newIndirect := reflect.New(indirect.Type())
		newIndirect.Elem().Set(reflect.ValueOf(indirect.Interface()))
		res = append(res, newIndirect.Interface())
	}
	return res
}

func GetData(tableName string, item interface{}) []interface{} {
	if tableName == "" {
		return []interface{}{}
	}
	sqlFields, fields, _ := StrutToSliceOfFieldAddressAndNames(item)
	if len(sqlFields) == 0 {
		return []interface{}{}
	}
	sqlReq := fmt.Sprintf(`SELECT %s FROM %s;`, strings.Join(sqlFields[:], ", "), tableName)
	params, err := db.Query(sqlReq)
	if err != nil {
		log.Printf("%+v", err)
		return nil
	}
	defer params.Close()
	var res []interface{}
	for params.Next() {
		err := params.Scan(fields...)
		if err != nil {
			log.Printf("%+v", err)
			return nil
		}
		indirect := reflect.Indirect(reflect.ValueOf(item))
		newIndirect := reflect.New(indirect.Type())
		newIndirect.Elem().Set(reflect.ValueOf(indirect.Interface()))
		res = append(res, newIndirect.Interface())
	}
	return res
}

func UpdateFields(tableName string, item interface{}) error {
	if tableName == "" {
		return errors.New("no table name")
	}
	var sqlReq string
	var err error
	names, values, _ := StrutToSliceOfSqlNamesAndValues(item)
	names, i := FilterSlice("id", names)
	if i < 0 {
		return errors.New("no id field")
	}
	itemId := values[i]
	copy(values[i:], values[i+1:])
	values = values[:len(values)-1]

	names, i = FilterSlice("enabled", names)
	if i >= 0 {
		copy(values[i:], values[i+1:])
		values = values[:len(values)-1]
	}
	sqlReq = fmt.Sprintf("UPDATE %s SET %s WHERE id = %d;", tableName, ValuesEqualPlaceholders(names), itemId)

	_, err = db.Exec(sqlReq, values...)
	if err != nil {
		return err
	}
	return err
}

func SwitchFields(tableName string, itemId int64, itemEnabled bool) error {
	if tableName == "" {
		return errors.New("no table name")
	}
	var sqlReq string
	var err error
	sqlReq = fmt.Sprintf("UPDATE %s SET enabled = $1 WHERE id = $2;", tableName)

	_, err = db.Exec(sqlReq, itemEnabled, itemId)
	if err != nil {
		return err
	}
	return err
}

func DeleteRow(tableName string, itemId int64) bool {
	if tableName == "" {
		return false
	}
	var sqlReq string
	var err error

	sqlReq = fmt.Sprintf("DELETE FROM %s WHERE id = $1;", tableName)

	_, err = db.Exec(sqlReq, itemId)
	if err != nil {
		return false
	}
	return true
}

func StrutSqlValue(s interface{}, name string) interface{} {
	t := reflect.TypeOf(s).Elem()
	fieldArr := reflect.ValueOf(s).Elem()

	for i := 0; i < t.NumField(); i++ {
		tag, ok := t.Field(i).Tag.Lookup("customsql")
		if !ok || tag == "" || tag != name {
			continue
		}
		return fieldArr.Field(i).Interface()
	}

	return nil
}

func CreateTableByStruct(row mainStruct.RowItem) bool {
	tableName := row.GetTableName()
	fkTable := row.GetFKTableName()
	if tableName == "" || fkTable == "" {
		panicErr(errors.New("no enough data for create table. tableName:  " + tableName + "; fkTable: " + fkTable))
		return false
	}
	prepared, fkey, uniq := StrutToSliceOfFieldTypesAndNames(row)
	uniqLine := strings.Join(uniq, ", ")
	if len(uniqLine) > 0 {
		uniqLine = ",\n UNIQUE (" + uniqLine + ")"
	}
	if len(prepared) == 0 {
		panicErr(errors.New("no table struct provided to create table"))
		return false
	}

	sqlReq := fmt.Sprintf(`	CREATE TABLE IF NOT EXISTS %s(
		%s bigint NOT NULL REFERENCES %s (id) ON DELETE CASCADE,
		%s%s
	)
	WITH (OIDS=FALSE);`,
		tableName, fkey, fkTable, strings.Join(prepared, ",\n"), uniqLine)

	_, err := db.Exec(sqlReq)
	panicErr(err)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}

	return true
}

func StrutToSliceOfFieldAddressAndNames(s interface{}) ([]string, []interface{}, string) {
	parentKey := ""
	t := reflect.TypeOf(s).Elem()
	fieldRowNameArr := make([]string, 0, t.NumField())
	fieldArr := reflect.ValueOf(s).Elem()
	fieldAddrArr := make([]interface{}, 0, fieldArr.NumField())

	for i := 0; i < t.NumField(); i++ {
		tag, ok := t.Field(i).Tag.Lookup("customsql")
		if !ok || tag == "" {
			continue
		}
		subConstrain := strings.Split(tag, ";")
		if len(subConstrain) > 1 {
			switch subConstrain[1] {
			case "unique":
				tag = subConstrain[0]
			}
		}
		subOption := strings.Split(tag, ":")
		if len(subOption) > 1 {
			switch subOption[0] {
			case "fkey":
				parentKey = subOption[1]
				continue
			case "pkey":
				tag = subOption[1]
			}
		}
		f := fieldArr.Field(i)
		fieldAddrArr = append(fieldAddrArr, f.Addr().Interface())
		fieldRowNameArr = append(fieldRowNameArr, tag)
	}

	return fieldRowNameArr, fieldAddrArr, parentKey
}

func StrutToSliceOfFieldTypesAndNames(s interface{}) ([]string, string, []string) {
	parentKey := ""
	var uniq []string
	t := reflect.TypeOf(s).Elem()
	fieldRowNameType := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		notNeed := false
		isUniq := false
		isSerial := false
		ending := " NOT NULL"
		tag, ok := t.Field(i).Tag.Lookup("customsql")
		if !ok || tag == "" {
			continue
		}
		subConstrain := strings.Split(tag, ";")
		if len(subConstrain) > 1 {
			switch subConstrain[1] {
			case "unique":
				tag = subConstrain[0]
				isUniq = true
			}
		}
		subOption := strings.Split(tag, ":")
		if len(subOption) > 1 {
			tag = subOption[1]
			switch subOption[0] {
			case "fkey":
				parentKey = subOption[1]
				notNeed = true
			case "pkey":
				ending += " PRIMARY KEY"
				isSerial = true
			}
		}
		if isUniq {
			uniq = append(uniq, tag)
		}
		if notNeed {
			continue
		}
		switch t.Field(i).Type.Name() {
		case "bool":
			fieldRowNameType = append(fieldRowNameType, tag+" BOOLEAN"+ending)
		case "int64":
			if isSerial {
				fieldRowNameType = append(fieldRowNameType, tag+" SERIAL"+ending)
			} else {
				fieldRowNameType = append(fieldRowNameType, tag+" BIGINT"+ending)
			}
		case "string":
			fieldRowNameType = append(fieldRowNameType, tag+" VARCHAR"+ending)
		default:
			log.Println(t.Field(i).Type.Name())
			log.Println(t.Field(i).Type.Elem().Name())
		}
	}

	return fieldRowNameType, parentKey, uniq
}

func StrutToSliceOfSqlNamesAndValues(s interface{}) ([]string, []interface{}, string) {
	parentKey := ""
	t := reflect.TypeOf(s).Elem()
	fieldRowNameArr := make([]string, 0, t.NumField())
	fieldArr := reflect.ValueOf(s).Elem()
	fieldAddrArr := make([]interface{}, 0, fieldArr.NumField())

	for i := 0; i < t.NumField(); i++ {
		tag, ok := t.Field(i).Tag.Lookup("customsql")
		if !ok || tag == "" {
			continue
		}
		subConstrain := strings.Split(tag, ";")
		if len(subConstrain) > 1 {
			switch subConstrain[1] {
			case "unique":
				tag = subConstrain[0]
			}
		}
		subOption := strings.Split(tag, ":")
		if len(subOption) > 1 {
			switch subOption[0] {
			case "fkey":
				parentKey = subOption[1]
				continue
			case "pkey":
				tag = subOption[1]
			}
		}
		fieldAddrArr = append(fieldAddrArr, fieldArr.Field(i).Interface())
		fieldRowNameArr = append(fieldRowNameArr, tag)
	}

	return fieldRowNameArr, fieldAddrArr, parentKey
}

func StrutToSliceOfValues(s interface{}) []interface{} {
	fieldArr := reflect.ValueOf(s).Elem()
	fieldAddrArr := make([]interface{}, fieldArr.NumField())

	for i := 0; i < fieldArr.NumField(); i++ {
		f := fieldArr.Field(i)
		fieldAddrArr[i] = reflect.ValueOf(f)
	}

	return fieldAddrArr
}

func StrutToSliceOfFieldAddress(s interface{}) []interface{} {
	fieldArr := reflect.ValueOf(s).Elem()
	fieldAddrArr := make([]interface{}, fieldArr.NumField())

	for i := 0; i < fieldArr.NumField(); i++ {
		f := fieldArr.Field(i)
		fieldAddrArr[i] = f.Addr().Interface()
	}

	return fieldAddrArr
}

func StrutToSliceOfSqlNames(s interface{}) []string {
	t := reflect.TypeOf(s).Elem()
	fieldRowNameArr := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		tag, ok := t.Field(i).Tag.Lookup("customsql")
		if !ok {
			continue
		}
		fieldRowNameArr[i] = tag
	}

	return fieldRowNameArr
}

func FilterSlice(toRemove string, sl []string) ([]string, int) {
	j := 0
	index := -1
	for _, n := range sl {
		if n == toRemove {
			index = j
			continue
		}
		sl[j] = n
		j++
	}
	sl = sl[:j]
	return sl, index
}

func ValuesPlaceholders(sl []string) string {
	var res string
	for i := 1; i <= len(sl); i++ {
		res += "$" + strconv.Itoa(i)
		if i != len(sl) {
			res += ", "
		}
	}
	return res
}

func ValuesEqualPlaceholders(sl []string) string {
	var res string
	for i := 1; i <= len(sl); i++ {
		res += sl[i-1] + "=$" + strconv.Itoa(i)
		if i != len(sl) {
			res += ", "
		}
	}
	return res
}
