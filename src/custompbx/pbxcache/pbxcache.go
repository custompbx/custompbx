package pbxcache

import (
	"custompbx/cache"
	"custompbx/db"
	"custompbx/mainStruct"
	"errors"
)

var dialplan *mainStruct.Dialplans
var channels *mainStruct.Channels
var globalVariables *mainStruct.GlobalVariables

func InitCacheObjects() {
	dialplan = mainStruct.NewDialplanItems()
	channels = mainStruct.NewChannelsCache()
	globalVariables = mainStruct.NewGlobalVariables()
}

func InitRootDB() {
	db.InitRootDB()
}

func InitPBXCache() {
	InitDialplanCache()
	InitGlobalVariablesCache()
}

func GetGlobalVariables() *mainStruct.GlobalVariables {
	return globalVariables
}

func GetGlobalVariableByName(name string) *mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetByName(name)
}

func GetGlobalVariableById(id int64) *mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetById(id)
}

func GetGlobalVariableNamedList() map[string]*mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetNamedList()
}

func GetGlobalVariableList() map[int64]*mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.GetList()
}

func GetGlobalVariableNotDynamicsProps() []*mainStruct.GlobalVariable {
	if globalVariables == nil {
		return nil
	}
	return globalVariables.NotDynamicsProps()
}

func UpdateFSInstanceDescription(instance *mainStruct.FsInstance, description string) error {
	if instance == nil {
		return errors.New("no id")
	}

	err := db.UpdateFSInstanceDescription(instance.Id, description)
	if err != nil {
		return err
	}
	instance.Description = description
	return nil
}

func InitGlobalVariablesCache() {
	db.GetGlobalVariables(globalVariables, cache.GetCurrentInstanceId())

}

func InitDialplanCache() {
	db.GetDialplanSettings(dialplan, cache.GetCurrentInstanceId())
	db.GetContexts(dialplan, cache.GetCurrentInstanceId())
	for _, context := range dialplan.Contexts.Props() {
		context.Extensions = mainStruct.NewExtensions()
		db.GetContextExtensions(context, dialplan)
		for _, extension := range context.Extensions.Props() {
			extension.Conditions = mainStruct.NewConditions()
			db.GetExtensionConditions(extension, dialplan)
			for _, condition := range extension.Conditions.Props() {
				condition.Regexes = mainStruct.NewRegexes()
				condition.Actions = mainStruct.NewActions()
				condition.AntiActions = mainStruct.NewAntiActions()
				db.GetConditionRegexes(condition, dialplan)
				db.GetConditionActions(condition, dialplan)
				db.GetConditionAntiActions(condition, dialplan)
			}
		}
		context.CacheFullXML()
	}
}

func GetChannelsCache() *mainStruct.Channels {
	return channels
}

func GetChannelsCounter() (int, int) {
	return channels.Total, channels.Answered
}
