package apps

import (
	"custompbx/db"
	"custompbx/webStruct"
	"github.com/custompbx/customorm"
)

var WebCases map[string]func(*webStruct.MessageData) webStruct.UserResponse

func InitApps() {
	InitTables()
	InitCases()
}

func InitTables() {
	corm := customorm.Init(db.GetDB())

	corm.CreateTable(&AutoDialerReducer{})
	corm.CreateTable(&AutoDialerTeam{})
	corm.CreateTable(&AutoDialerList{})
	corm.CreateTable(&AutoDialerCompany{})
	corm.CreateTable(&AutoDialerTeamMember{})
	corm.CreateTable(&AutoDialerListMember{})
	corm.CreateTable(&AutoDialerReducerMember{})
}

func InitCases() {
	WebCases = make(map[string]func(*webStruct.MessageData) webStruct.UserResponse)
	WebCases["GetAutoDialerCompanies"] = GetAutoDialerCompanies
	WebCases["AddAutoDialerCompany"] = AddAutoDialerCompany
	WebCases["DelAutoDialerCompany"] = DelAutoDialerCompany
	WebCases["UpdateAutoDialerCompany"] = UpdateAutoDialerCompany
	WebCases["GetAutoDialerTeams"] = GetAutoDialerTeams
	WebCases["AddAutoDialerTeam"] = AddAutoDialerTeam
	WebCases["DelAutoDialerTeam"] = DelAutoDialerTeam
	WebCases["UpdateAutoDialerTeam"] = UpdateAutoDialerTeam
	WebCases["GetAutoDialerTeamMembers"] = GetAutoDialerTeamMembers
	WebCases["AddAutoDialerTeamMember"] = AddAutoDialerTeamMember
	WebCases["AddAutoDialerTeamMembers"] = AddAutoDialerTeamMembers
	WebCases["DelAutoDialerTeamMember"] = DelAutoDialerTeamMember
	WebCases["UpdateAutoDialerTeamMember"] = UpdateAutoDialerTeamMember
	WebCases["GetAutoDialerLists"] = GetAutoDialerLists
	WebCases["AddAutoDialerList"] = AddAutoDialerList
	WebCases["DelAutoDialerList"] = DelAutoDialerList
	WebCases["UpdateAutoDialerList"] = UpdateAutoDialerList
	//WebCases["GetAutoDialerListMembers"] = GetAutoDialerListMembers
	WebCases["AddAutoDialerListMember"] = AddAutoDialerListMember
	//WebCases["AddAutoDialerListMembers"] = AddAutoDialerListMembers
	WebCases["DelAutoDialerListMember"] = DelAutoDialerListMember
	//WebCases["UpdateAutoDialerListMember"] = UpdateAutoDialerListMember
	WebCases["GetAutoDialerReducers"] = GetAutoDialerReducers
	WebCases["AddAutoDialerReducer"] = AddAutoDialerReducer
	WebCases["DelAutoDialerReducer"] = DelAutoDialerReducer
	WebCases["UpdateAutoDialerReducer"] = UpdateAutoDialerReducer
	WebCases["GetAutoDialerReducerMembers"] = GetAutoDialerReducerMembers
	WebCases["AddAutoDialerReducerMember"] = AddAutoDialerReducerMember
	WebCases["DelAutoDialerReducerMember"] = DelAutoDialerReducerMember
	WebCases["UpdateAutoDialerReducerMember"] = UpdateAutoDialerReducerMember
}
