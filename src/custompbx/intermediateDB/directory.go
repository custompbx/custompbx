package intermediateDB

import "custompbx/altStruct"

func InitDirectoryDB() {

	corm := GetCORM()

	corm.CreateTable(&altStruct.DirectoryDomain{})
	corm.CreateTable(&altStruct.DirectoryDomainParameter{})
	corm.CreateTable(&altStruct.DirectoryDomainVariable{})
	corm.CreateTable(&altStruct.DirectoryDomainUser{})
	corm.CreateTable(&altStruct.DirectoryDomainUserParameter{})
	corm.CreateTable(&altStruct.DirectoryDomainUserVariable{})
	corm.CreateTable(&altStruct.DirectoryDomainUserGateway{})
	corm.CreateTable(&altStruct.DirectoryDomainUserGatewayParameter{})
	corm.CreateTable(&altStruct.DirectoryDomainUserGatewayVariable{})
	corm.CreateTable(&altStruct.DirectoryDomainGroup{})
	corm.CreateTable(&altStruct.DirectoryDomainGroupUser{})
}
