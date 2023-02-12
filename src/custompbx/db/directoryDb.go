package db

import (
	"custompbx/mainStruct"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDirectoryDB() {
	createDirectoryDomainsTable(db)
	createDirectoryUsersTable(db)
	createDirectoryUsersDomainParamsTable(db)
	createDirectoryUsersParamsTable(db)
	createDirectoryUsersDomainVarsTable(db)
	createDirectoryUsersVarsTable(db)
	createDirectoryGroupsTable(db)
	createDirectoryGroupUsersTable(db)
	createDirectoryUsersGateways(db)
	createDirectoryUsersGatewaysParams(db)
	createDirectoryUsersGatewaysVars(db)
}

func createDirectoryDomainsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_domains(
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

func createDirectoryUsersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_users(
		id serial NOT NULL PRIMARY KEY,
		domain_id bigint NOT NULL REFERENCES directory_domains (id) ON DELETE CASCADE,
		name VARCHAR,
		cache INTEGER DEFAULT 5000,
		cidr VARCHAR DEFAULT NULL,
		number_alias VARCHAR DEFAULT NULL,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (domain_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryUsersDomainParamsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_domain_parameters(
		id serial NOT NULL PRIMARY KEY,
		domain_id bigint NOT NULL REFERENCES directory_domains (id) ON DELETE CASCADE,
		name VARCHAR,
		value VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (domain_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}
func createDirectoryUsersParamsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_parameters(
		id serial NOT NULL PRIMARY KEY,
		user_id bigint NOT NULL REFERENCES directory_users (id) ON DELETE CASCADE,
		name VARCHAR,
		value VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (user_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryUsersDomainVarsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_domain_variables(
		id serial NOT NULL PRIMARY KEY,
		domain_id bigint NOT NULL REFERENCES directory_domains (id) ON DELETE CASCADE,
		name VARCHAR,
		value VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (domain_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryUsersVarsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_variables(
		id serial NOT NULL PRIMARY KEY,
		user_id bigint NOT NULL REFERENCES directory_users (id) ON DELETE CASCADE,
		name VARCHAR,
		value VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (user_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryGroupsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_groups(
		id serial NOT NULL PRIMARY KEY,
		domain_id bigint NOT NULL REFERENCES directory_domains (id) ON DELETE CASCADE,
		name VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (domain_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryGroupUsersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_group_users(
		id serial NOT NULL PRIMARY KEY,
		group_id bigint NOT NULL REFERENCES directory_groups (id) ON DELETE CASCADE,
		user_id bigint NOT NULL REFERENCES directory_users (id) ON DELETE CASCADE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (group_id, user_id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryUsersGateways(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_users_gateways(
		id serial NOT NULL PRIMARY KEY,
		user_id bigint NOT NULL REFERENCES directory_users (id) ON DELETE CASCADE,
		name VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (user_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryUsersGatewaysParams(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_users_gateways_params(
		id serial NOT NULL PRIMARY KEY,
		gateway_id bigint NOT NULL REFERENCES directory_users_gateways (id) ON DELETE CASCADE,
		name VARCHAR,
		value VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (gateway_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDirectoryUsersGatewaysVars(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS directory_users_gateways_vars(
		id serial NOT NULL PRIMARY KEY,
		gateway_id bigint NOT NULL REFERENCES directory_users_gateways (id) ON DELETE CASCADE,
		name VARCHAR,
		value VARCHAR,
		direction VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (gateway_id, name)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func SetDomain(name string, instanceId int64) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO directory_domains(name, instance_id) values($1, $2) returning id;", name, instanceId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetDomainUser(domainId int64, name, cidr, numberAlias string) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO directory_users(domain_id, name, cidr, number_alias) values($1, $2, $3, $4) returning id;", domainId, name, cidr, numberAlias).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetDomainParameter(domainId int64, name, value string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_domain_parameters(domain_id, name, value)
							values($1, $2, $3) returning id;`, domainId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetDomainVariable(domainId int64, name, value string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_domain_variables(domain_id, name, value) 
							values($1, $2, $3) returning id;`, domainId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetUserParameter(directoryId int64, name, value string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_parameters(user_id, name, value)
							values($1, $2, $3) returning id;`, directoryId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetUserVariable(directoryId int64, name, value string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_variables(user_id, name, value) 
							values($1, $2, $3) returning id;`, directoryId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetDomainGroup(domainId int64, name string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_groups(domain_id, name)
							values($1, $2) returning id;`, domainId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetGroupMember(groupId, userId int64) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_group_users(group_id, user_id)
							values($1, $2) returning id;`, groupId, userId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetUserGateway(userId int64, name string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_users_gateways(user_id, name)
							values($1, $2) returning id;`, userId, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetUserGatewayParam(gatewayId int64, name, value string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_users_gateways_params(gateway_id, name, value)
							values($1, $2, $3) returning id;`, gatewayId, name, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetUserGatewayVar(gatewayId int64, name, value, direction string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO directory_users_gateways_vars(gateway_id, name, value, direction)
							values($1, $2, $3, $4) returning id;`, gatewayId, name, value, direction).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

/*
func SetUser(name string) (int64, error) {
		sqlReq := fmt.Sprintf(
		`INSERT INTO directory_users(domain_id, name) values(%d,'%s') returning id;

	`, , name)
	var id int64
	err := db.QueryRow(sqlReq).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}*/
//GET
func GetDomains(directory *mainStruct.DirectoryItems, instanceId int64) {
	domains, err := db.Query(`SELECT dd.id AS id, dd.name AS name, dd.enabled FROM directory_domains dd WHERE instance_id = $1;`, instanceId)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer domains.Close()
	for domains.Next() {
		var domain mainStruct.Domain
		err := domains.Scan(&domain.Id, &domain.Name, &domain.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		directory.Domains.Set(&domain)
	}
}

func GetDomainParams(domain *mainStruct.Domain, directory *mainStruct.DirectoryItems) {
	params, err := db.Query(`SELECT ddp.id AS id, ddp.name AS name, ddp.value AS value, ddp.enabled FROM directory_domain_parameters ddp WHERE domain_id = $1;`, domain.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.DomainParam
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		param.Domain = domain
		domain.Params.Set(&param)
		directory.DomainParams.Set(&param)
	}
}

func GetDomainVars(domain *mainStruct.Domain, directory *mainStruct.DirectoryItems) {
	vars, err := db.Query(`SELECT ddv.id AS id, ddv.name AS name, ddv.value AS value, ddv.enabled FROM directory_domain_variables ddv WHERE domain_id = $1;`, domain.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer vars.Close()
	for vars.Next() {
		var variable mainStruct.DomainVariable
		err := vars.Scan(&variable.Id, &variable.Name, &variable.Value, &variable.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		variable.Domain = domain
		domain.Vars.Set(&variable)
		directory.DomainVars.Set(&variable)
	}
}

func GetUser(domain *mainStruct.Domain, directory *mainStruct.DirectoryItems) {
	directories, err := db.Query(`SELECT dn.id AS id, dn.name AS name, dn.cache AS cache, COALESCE(dn.cidr, '') AS cidr, COALESCE(dn.number_alias, '') AS number_alias, dn.enabled FROM directory_users dn WHERE domain_id = $1;`, domain.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer directories.Close()
	for directories.Next() {
		var user mainStruct.User
		err := directories.Scan(&user.Id, &user.Name, &user.Cache, &user.Cidr, &user.NumberAlias, &user.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		user.Domain = domain
		domain.Users.Set(&user)
		directory.Users.Set(&user)
	}
}

func GetUserParams(user *mainStruct.User, directory *mainStruct.DirectoryItems) {
	params, err := db.Query(`SELECT dp.id AS id, dp.name AS name, dp.value AS value, dp.enabled FROM directory_parameters dp WHERE user_id = $1;`, user.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var param mainStruct.UserParam
		err := params.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		param.User = user
		user.Params.Set(&param)
		directory.UserParams.Set(&param)
	}
}

func GetUserVars(user *mainStruct.User, directory *mainStruct.DirectoryItems) {
	vars, err := db.Query(`SELECT dv.id AS id, dv.name AS name, dv.value AS value, dv.enabled FROM directory_variables dv WHERE user_id = $1;`, user.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer vars.Close()
	for vars.Next() {
		var variable mainStruct.UserVariable
		err := vars.Scan(&variable.Id, &variable.Name, &variable.Value, &variable.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		variable.User = user
		user.Vars.Set(&variable)
		directory.UserVars.Set(&variable)
	}
}

func GetDomainGroups(domain *mainStruct.Domain, directory *mainStruct.DirectoryItems) {
	directories, err := db.Query(`SELECT dg.id AS id, dg.name AS name, dg.enabled FROM directory_groups dg WHERE domain_id = $1;`, domain.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer directories.Close()
	for directories.Next() {
		var group mainStruct.Group
		err := directories.Scan(&group.Id, &group.Name, &group.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		group.Domain = domain
		domain.Groups.Set(&group)
		directory.Groups.Set(&group)
	}
}

func GetGroupMembers(group *mainStruct.Group, directory *mainStruct.DirectoryItems) {
	directories, err := db.Query(`SELECT dgu.id AS id, du.name AS name, dgu.user_id AS user_id, du.enabled FROM directory_group_users dgu LEFT JOIN directory_users du ON dgu.user_id = du.id WHERE dgu.group_id = $1;`, group.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer directories.Close()
	for directories.Next() {
		var user mainStruct.GroupUser
		err := directories.Scan(&user.Id, &user.Name, &user.UserId, &user.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		user.Type = "pointer"
		user.Group = group
		group.Users.Set(&user)
		directory.GroupUsers.Set(&user)
	}
}

func GetUserGateways(user *mainStruct.User, directory *mainStruct.DirectoryItems) {
	directories, err := db.Query(`SELECT dug.id AS id, dug.name AS name, dug.enabled FROM directory_users_gateways dug WHERE user_id = $1;`, user.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer directories.Close()
	for directories.Next() {
		var gateway mainStruct.UserGateway
		err := directories.Scan(&gateway.Id, &gateway.Name, &gateway.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		gateway.User = user
		user.Gateways.Set(&gateway)
		directory.UserGateways.Set(&gateway)
	}
}

func GetUserGatewaysParams(gateway *mainStruct.UserGateway, directory *mainStruct.DirectoryItems) {
	directories, err := db.Query(`SELECT dugp.id AS id, dugp.name AS name, dugp.value AS value, dugp.enabled FROM directory_users_gateways_params dugp WHERE dugp.gateway_id = $1;`, gateway.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer directories.Close()
	for directories.Next() {
		var param mainStruct.GatewayParam
		err := directories.Scan(&param.Id, &param.Name, &param.Value, &param.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		param.Gateway = gateway
		gateway.Params.Set(&param)
		directory.GatewayParams.Set(&param)
	}
}

func GetUserGatewaysVars(gateway *mainStruct.UserGateway, directory *mainStruct.DirectoryItems) {
	directories, err := db.Query(`SELECT id, name, value, COALESCE(direction, ''), enabled FROM directory_users_gateways_vars WHERE gateway_id = $1;`, gateway.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer directories.Close()
	for directories.Next() {
		var variable mainStruct.GatewayVariable
		err := directories.Scan(&variable.Id, &variable.Name, &variable.Value, &variable.Direction, &variable.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		variable.Gateway = gateway
		gateway.Vars.Set(&variable)
		directory.GatewayVars.Set(&variable)
	}
}

// DELETE
func DelUser(userId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_users WHERE id = $1;`, userId)
	if err != nil {
		return false
	}
	return true
}

func DelUserParameter(userId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_parameters WHERE id = $1;`, userId)
	if err != nil {
		return false
	}
	return true
}

func DelUserVariable(varId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_variables WHERE id = $1;`, varId)
	if err != nil {
		return false
	}
	return true
}

func DelDomain(userId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_domains WHERE id = $1;`, userId)
	if err != nil {
		return false
	}
	return true
}

func DelDomainParameter(paramId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_domain_parameters WHERE id = $1;`, paramId)
	if err != nil {
		return false
	}
	return true
}

func DelDomainVariable(varId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_domain_variables WHERE id = $1;`, varId)
	if err != nil {
		return false
	}
	return true
}

func DelUserGateway(gatewayId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_users_gateways WHERE id = $1;`, gatewayId)
	if err != nil {
		return false
	}
	return true
}

func DelUserGatewayParam(paramId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_users_gateways_params WHERE id = $1;`, paramId)
	if err != nil {
		return false
	}
	return true
}

func DelGroup(groupId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_groups WHERE id = $1;`, groupId)
	if err != nil {
		return false
	}
	return true
}

func DelGroupMember(userId int64) bool {
	_, err := db.Exec(`DELETE FROM directory_group_users WHERE  id = $1;`, userId)
	if err != nil {
		return false
	}
	return true
}

// UPDATE
func UpdateDomain(domainId int64, newName string) error {
	_, err := db.Exec("UPDATE directory_domains SET name = $1 WHERE id = $2;", newName, domainId)
	if err != nil {
		return err
	}
	return err
}

func UpdateDomainUser(userId int64, newName string) error {
	_, err := db.Exec("UPDATE directory_users SET name = $1 WHERE id = $2;", newName, userId)
	if err != nil {
		return err
	}
	return err
}

func UpdateUserParameterValue(paramId int64, newValue string) error {
	_, err := db.Exec("UPDATE directory_parameters SET value = $1 WHERE id = $2;", newValue, paramId)
	if err != nil {
		return err
	}
	return err
}

func UpdateUserVariableValue(varId int64, newValue string) error {
	_, err := db.Exec("UPDATE directory_variables SET value = $1 WHERE id = $2;", newValue, varId)
	if err != nil {
		return err
	}
	return err
}

func UpdateUserParameter(paramId int64, newName string, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE directory_parameters SET name = $1, value = $2 WHERE id = $3;", newName, newValue, paramId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateUserVariable(varId int64, newName string, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE directory_variables SET name = $1, value = $2 WHERE id = $3;", newName, newValue, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateDomainParameterValue(paramId int64, newValue string) error {
	_, err := db.Exec("UPDATE directory_domain_parameters SET value = $1 WHERE id = $2;", newValue, paramId)
	if err != nil {
		return err
	}
	return err
}

func UpdateDomainParameter(paramId int64, newName string, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE directory_domain_parameters SET name = $1, value = $2 WHERE id = $3;", newName, newValue, paramId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateUserCache(userId int64, newValue uint) (int64, error) {
	res, err := db.Exec("UPDATE directory_users SET cache = $1 WHERE id = $2;", newValue, userId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateUserCidr(userId int64, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE directory_users SET cidr = $1 WHERE id = $2;", newValue, userId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateDomainVariableValue(varId int64, newValue string) error {
	_, err := db.Exec("UPDATE directory_domain_variables SET value = $1 WHERE id = $2;", newValue, varId)
	if err != nil {
		return err
	}
	return err
}

func UpdateDomainVariable(varId int64, newName string, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE directory_domain_variables SET name = $1, value = $2 WHERE id = $3;", newName, newValue, varId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateGroupName(groupId int64, newName string) error {
	_, err := db.Exec("UPDATE directory_groups SET name = $1 WHERE id = $2;", newName, groupId)
	if err != nil {
		return err
	}
	return err
}

func UpdateUserGatewayParameter(paramId int64, newName string, newValue string) (int64, error) {
	res, err := db.Exec("UPDATE directory_users_gateways_params SET name = $1, value = $2 WHERE id = $3;", newName, newValue, paramId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateUserGatewayVariable(paramId int64, newName string, newValue, direction string) (int64, error) {
	res, err := db.Exec("UPDATE directory_users_gateways_vars SET name = $1, value = $2, direction = $3 WHERE id = $4;", newName, newValue, direction, paramId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SwitchUserGatewayVariable(paramId int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_users_gateways_vars SET enabled = $1 WHERE id = $2;", enabled, paramId)
	return err
}

func DelUserGatewayVariable(id int64) bool {
	_, err := db.Exec(`DELETE FROM directory_users_gateways_vars WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateDomainUserGateway(userId int64, newName string) error {
	_, err := db.Exec("UPDATE directory_users_gateways SET name = $1 WHERE id = $2;", newName, userId)
	if err != nil {
		return err
	}
	return err
}

func SwitchDomain(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_domains SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchDomainParameter(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_domain_parameters SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchDomainVariable(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_domain_variables SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchUsers(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_users SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchUserParameter(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_parameters SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchUserVariable(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_variables SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchUserGateway(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_users_gateways SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchUserGatewayParameter(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE directory_users_gateways_params SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}

func SwitchGroup(id int64, enabled bool) error {
	_, err := db.Exec("UPDATE custompbx.public.directory_groups SET enabled = $1 WHERE id = $2;", enabled, id)
	return err
}
