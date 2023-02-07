package pbxcache

import (
	"custompbx/cache"
	"custompbx/db"
	"custompbx/mainStruct"
	"errors"
)

func SetContext(contextName string) (*mainStruct.Context, error) {
	if contextName == "" {
		return nil, errors.New("empty domain name")
	}
	if dialplan.Contexts.HasName(contextName) {
		return nil, errors.New("context name already exists")
	}
	res, err := db.SetContext(contextName, cache.GetCurrentInstanceId())
	if err != nil {
		return nil, err
	}
	context := &mainStruct.Context{Id: res, Name: contextName, Extensions: mainStruct.NewExtensions(), Enabled: true, Dialplan: dialplan}
	context.CacheFullXML()
	dialplan.Contexts.Set(context)

	return context, err
}

func GetContext(id int64) *mainStruct.Context {
	context, _ := dialplan.Contexts.GetById(id)
	return context
}

func SwitchDialplanNoProceed(switcher bool) bool {
	if dialplan == nil {
		return false
	}
	value := "false"
	if switcher {
		value = "true"
	}
	err := db.SetDialplanSettings(dialplan, mainStruct.NoProceedName, value)
	if err != nil {
		return false
	}

	dialplan.NoProceed = switcher

	if !switcher {
		return true
	}
	for _, context := range dialplan.Contexts.Props() {
		context.CacheFullXML()
	}

	return true
}

func GetExtension(id int64) *mainStruct.Extension {
	extension, _ := dialplan.Extensions.GetById(id)
	return extension
}

func GetCondition(id int64) *mainStruct.Condition {
	condition, _ := dialplan.Conditions.GetById(id)
	return condition
}

func GetRegex(id int64) *mainStruct.Regex {
	action, _ := dialplan.Regexes.GetById(id)
	return action
}

func GetAction(id int64) *mainStruct.Action {
	action, _ := dialplan.Actions.GetById(id)
	return action
}

func GetAntiAction(id int64) *mainStruct.AntiAction {
	antiAction, _ := dialplan.AntiActions.GetById(id)
	return antiAction
}

func SetContextExtension(context *mainStruct.Context, extensionName, extensionContinue string) (*mainStruct.Extension, error) {
	if context == nil {
		return nil, errors.New("context doesn't exists")
	}
	if context.Extensions.GetByName(extensionName) != nil {
		return nil, errors.New("extensions name already exists")
	}

	id, position, err := db.SetContextExtension(context.Id, extensionName, extensionContinue)
	if err != nil {
		return nil, err
	}
	extension := &mainStruct.Extension{Id: id, Name: extensionName, Continue: extensionContinue, Context: context, Conditions: mainStruct.NewConditions(), Enabled: true, Position: position}
	context.Extensions.Set(extension)
	dialplan.Extensions.Set(extension)
	context.CacheFullXML()

	return extension, err
}

func SetExtensionCondition(
	extension *mainStruct.Extension,
	conditionBreak,
	field,
	expression,
	hour,
	mday,
	mon,
	mweek,
	wday,
	dateTime,
	timeOfDay,
	year,
	minute,
	week,
	yday,
	minday,
	tzOffset,
	dst,
	regex string,
) (*mainStruct.Condition, error) {
	if extension == nil {
		return nil, errors.New("extension doesn't exists")
	}

	id, position, err := db.SetExtensionCondition(
		extension.Id,
		conditionBreak,
		field,
		expression,
		hour,
		mday,
		mon,
		mweek,
		wday,
		dateTime,
		timeOfDay,
		year,
		minute,
		week,
		yday,
		minday,
		tzOffset,
		dst,
		regex)
	if err != nil {
		return nil, err
	}
	condition := &mainStruct.Condition{
		Id:         id,
		Break:      conditionBreak,
		Field:      field,
		Expression: expression,
		Hour:       hour,
		Mday:       mday,
		Mon:        mon,
		Mweek:      mweek,
		Wday:       wday,

		DateTime:  dateTime,
		TimeOfDay: timeOfDay,
		Year:      year,
		Minute:    minute,
		Week:      week,
		Yday:      yday,
		Minday:    minday,
		TzOffset:  tzOffset,
		Dst:       dst,
		Regex:     regex,

		Regexes:     mainStruct.NewRegexes(),
		Actions:     mainStruct.NewActions(),
		AntiActions: mainStruct.NewAntiActions(),
		Extension:   extension,
		Enabled:     true,
		Position:    position,
	}
	extension.Conditions.Set(condition)
	dialplan.Conditions.Set(condition)
	extension.Context.CacheFullXML()

	return condition, err
}

func SetConditionRegex(condition *mainStruct.Condition, field, expr string) (*mainStruct.Regex, error) {
	if condition == nil {
		return nil, errors.New("condition doesn't exists")
	}

	res, err := db.SetConditionRegex(condition.Id, field, expr)
	if err != nil {
		return nil, err
	}
	regex := &mainStruct.Regex{Id: res, Condition: condition, Field: field, Expression: expr, Enabled: true}
	condition.Regexes.Set(regex)
	dialplan.Regexes.Set(regex)
	condition.Extension.Context.CacheFullXML()

	return regex, err
}

func SetConditionAction(condition *mainStruct.Condition, application, data string, inline bool) (*mainStruct.Action, error) {
	if condition == nil {
		return nil, errors.New("condition doesn't exists")
	}

	id, position, err := db.SetConditionAction(condition.Id, application, data, inline)
	if err != nil {
		return nil, err
	}
	action := &mainStruct.Action{Id: id, Condition: condition, Application: application, Data: data, Inline: inline, Enabled: true, Position: position}
	condition.Actions.Set(action)
	dialplan.Actions.Set(action)
	condition.Extension.Context.CacheFullXML()

	return action, err
}

func SetConditionAntiAction(condition *mainStruct.Condition, application, data string) (*mainStruct.AntiAction, error) {
	if condition == nil {
		return nil, errors.New("condition doesn't exists")
	}

	id, position, err := db.SetConditionAntiAction(condition.Id, application, data)
	if err != nil {
		return nil, err
	}
	antiAction := &mainStruct.AntiAction{Id: id, Condition: condition, Application: application, Data: data, Enabled: true, Position: position}
	condition.AntiActions.Set(antiAction)
	dialplan.AntiActions.Set(antiAction)
	condition.Extension.Context.CacheFullXML()

	return antiAction, err
}

func GetContextByName(name string) *mainStruct.Context {
	context, _ := dialplan.Contexts.GetByName(name)
	return context
}

func GetContexts() (map[int64]*mainStruct.Context, bool) {
	if dialplan.Contexts == nil {
		return map[int64]*mainStruct.Context{}, false
	}

	item := dialplan.Contexts.GetList()
	return item, true
}

func GetDialplan() *mainStruct.Dialplans {
	return dialplan
}

func MoveContextExtension(extension *mainStruct.Extension, newPosition int64) error {
	if extension == nil || newPosition == 0 {
		return errors.New("extension doesn't exists")
	}

	err := db.MoveContextExtension(extension, newPosition)
	if err != nil {
		return err
	}
	extension.Context.CacheFullXML()
	return err
}

func MoveContextCondition(condition *mainStruct.Condition, newPosition int64) error {
	if condition == nil || newPosition == 0 {
		return errors.New("extension doesn't exists")
	}

	err := db.MoveContextCondition(condition, newPosition)
	if err != nil {
		return err
	}

	condition.Extension.Context.CacheFullXML()
	return err
}

func MoveContextAction(action *mainStruct.Action, newPosition int64) error {
	if action == nil || newPosition == 0 {
		return errors.New("action doesn't exists")
	}

	err := db.MoveContextAction(action, newPosition)
	if err != nil {
		return err
	}
	action.Condition.Extension.Context.CacheFullXML()
	return err
}

func MoveContextAntiAction(antiAction *mainStruct.AntiAction, newPosition int64) error {
	if antiAction == nil || newPosition == 0 {
		return errors.New("antiAction doesn't exists")
	}

	err := db.MoveContextAntiAction(antiAction, newPosition)
	if err != nil {
		return err
	}
	antiAction.Condition.Extension.Context.CacheFullXML()

	return err
}

func DelContextRegex(regex *mainStruct.Regex) int64 {
	if regex == nil {
		return 0
	}
	ok := db.DelContextRegex(regex.Id)
	if !ok {
		return 0
	}

	regex.Condition.Regexes.Remove(regex)
	dialplan.Regexes.Remove(regex)
	regex.Condition.Extension.Context.CacheFullXML()
	return regex.Id
}

func DelContextAction(action *mainStruct.Action) int64 {
	if action == nil {
		return 0
	}
	ok := db.DelContextAction(action.Id)
	if !ok {
		return 0
	}

	action.Condition.Actions.Remove(action)
	dialplan.Actions.Remove(action)
	action.Condition.Extension.Context.CacheFullXML()
	return action.Id
}

func DelContextAntiAction(antiAction *mainStruct.AntiAction) int64 {
	if antiAction == nil {
		return 0
	}
	ok := db.DelContextAntiAction(antiAction.Id)
	if !ok {
		return 0
	}

	antiAction.Condition.AntiActions.Remove(antiAction)
	dialplan.AntiActions.Remove(antiAction)
	antiAction.Condition.Extension.Context.CacheFullXML()
	return antiAction.Id
}

func UpdateContextRegex(regex *mainStruct.Regex, field, expression string) int64 {
	if regex == nil {
		return 0
	}
	ok := db.UpdateContextRegex(regex.Id, field, expression)
	if !ok {
		return 0
	}

	regex.Field = field
	regex.Expression = expression
	regex.Condition.Extension.Context.CacheFullXML()
	return regex.Id
}

func UpdateContextAction(action *mainStruct.Action, application, data string, inline bool) int64 {
	if action == nil {
		return 0
	}
	ok := db.UpdateContextAction(action.Id, application, data, inline)
	if !ok {
		return 0
	}

	action.Application = application
	action.Data = data
	action.Inline = inline
	action.Condition.Extension.Context.CacheFullXML()
	return action.Id
}

func UpdateContextAntiAction(antiAction *mainStruct.AntiAction, application, data string) int64 {
	if antiAction == nil {
		return 0
	}
	ok := db.UpdateContextAntiAction(antiAction.Id, application, data)
	if !ok {
		return 0
	}

	antiAction.Application = application
	antiAction.Data = data
	antiAction.Condition.Extension.Context.CacheFullXML()
	return antiAction.Id
}

func SwitchContextRegex(regex *mainStruct.Regex, switcher bool) int64 {
	if regex == nil {
		return 0
	}
	ok := db.SwitchContextRegex(regex.Id, switcher)
	if !ok {
		return 0
	}

	regex.Enabled = switcher
	regex.Condition.Extension.Context.CacheFullXML()
	return regex.Id
}

func SwitchContextAction(action *mainStruct.Action, switcher bool) int64 {
	if action == nil {
		return 0
	}
	ok := db.SwitchContextAction(action.Id, switcher)
	if !ok {
		return 0
	}

	action.Enabled = switcher
	action.Condition.Extension.Context.CacheFullXML()
	return action.Id
}

func SwitchContextAntiAction(antiAction *mainStruct.AntiAction, switcher bool) int64 {
	if antiAction == nil {
		return 0
	}
	ok := db.SwitchContextAntiAction(antiAction.Id, switcher)
	if !ok {
		return 0
	}

	antiAction.Enabled = switcher
	antiAction.Condition.Extension.Context.CacheFullXML()
	return antiAction.Id
}

func RenameContext(context *mainStruct.Context, newName string) (*mainStruct.Context, error) {
	if context == nil {
		return nil, errors.New("no context")
	}
	if newName == "" {
		return nil, errors.New("empty context name")
	}
	if dialplan.Contexts.HasName(newName) {
		return nil, errors.New("domain name already exists")
	}
	err := db.RenameContext(context.Id, newName)
	if err != nil {
		return nil, err
	}
	dialplan.Contexts.Rename(context.Name, newName)
	context.CacheFullXML()

	return context, err
}

func RenameExtension(extension *mainStruct.Extension, newName string) (*mainStruct.Extension, error) {
	if extension == nil {
		return nil, errors.New("no extension")
	}
	if newName == "" {
		return nil, errors.New("empty extension name")
	}
	if dialplan.Extensions.HasName(newName) {
		return nil, errors.New("domain name already exists")
	}
	err := db.RenameExtension(extension.Id, newName)
	if err != nil {
		return nil, err
	}
	dialplan.Extensions.Rename(extension.Name, newName)
	extension.Context.Extensions.Rename(extension.Name, newName)
	extension.Context.CacheFullXML()

	return extension, err
}

func DeleteContext(context *mainStruct.Context) error {
	if context == nil {
		return errors.New("no context")
	}
	err := db.DeleteContext(context.Id)
	if err != nil {
		return err
	}
	dialplan.Contexts.Remove(context)
	dialplan.ClearUp()

	return err
}

func DeleteExtension(extension *mainStruct.Extension) error {
	if extension == nil {
		return errors.New("no extension")
	}
	err := db.DeleteExtension(extension.Id)
	if err != nil {
		return err
	}
	dialplan.Extensions.Remove(extension)
	extension.Context.Extensions.Remove(extension)
	dialplan.ClearExtensions()
	extension.Context.CacheFullXML()

	return err
}

func SwitchExtensionContinue(extension *mainStruct.Extension, extensionContinue string) error {
	if extension == nil {
		return errors.New("no extension")
	}
	err := db.SwitchExtensionContinue(extension.Id, extensionContinue)
	if err != nil {
		return err
	}
	extension.Continue = extensionContinue
	extension.Context.CacheFullXML()

	return err
}

func UpdateExtensionCondition(
	condition *mainStruct.Condition, conditionBreak, field, expression, hour, mday, mon, mweek, wday, dateTime, timeOfDay,
	year, minute, week, yday, minday, tzOffset, dst, regex string) error {
	if condition == nil {
		return errors.New("condition doesn't exists")
	}

	if field != "" {
		regex = ""
		hour = ""
		mday = ""
		mon = ""
		mweek = ""
		wday = ""
		dateTime = ""
		timeOfDay = ""
		year = ""
		minute = ""
		week = ""
		yday = ""
		minday = ""
		tzOffset = ""
		dst = ""
	} else if regex != "" {
		field = ""
		expression = ""
		hour = ""
		mday = ""
		mon = ""
		mweek = ""
		wday = ""
		dateTime = ""
		timeOfDay = ""
		year = ""
		minute = ""
		week = ""
		yday = ""
		minday = ""
		tzOffset = ""
		dst = ""
	} else {
		field = ""
		expression = ""
		regex = ""
	}

	err := db.UpdateExtensionCondition(
		condition.Id, conditionBreak, field, expression, hour, mday, mon, mweek, wday, dateTime, timeOfDay,
		year, minute, week, yday, minday, tzOffset, dst, regex)

	if err != nil {
		return err
	}
	condition.Break = conditionBreak
	condition.Field = field
	condition.Expression = expression
	condition.Hour = hour
	condition.Mday = mday
	condition.Mon = mon
	condition.Mweek = mweek
	condition.Wday = wday
	condition.DateTime = dateTime
	condition.TimeOfDay = timeOfDay
	condition.Year = year
	condition.Minute = minute
	condition.Week = week
	condition.Yday = yday
	condition.Minday = minday
	condition.TzOffset = tzOffset
	condition.Dst = dst
	condition.Regex = regex

	condition.Extension.Context.CacheFullXML()

	return err
}

func SwitchContextCondition(condition *mainStruct.Condition, switcher bool) int64 {
	if condition == nil {
		return 0
	}
	err := db.SwitchContextConditon(condition.Id, switcher)
	if err != nil {
		return 0
	}

	condition.Enabled = switcher
	condition.Extension.Context.CacheFullXML()

	return condition.Id
}

func DeleteContextCondition(condition *mainStruct.Condition) int64 {
	if condition == nil {
		return 0
	}
	err := db.DeleteContextConditon(condition.Id)
	if err != nil {
		return 0
	}

	condition.Extension.Conditions.Remove(condition)
	dialplan.Conditions.Remove(condition)
	dialplan.ClearConditions()
	condition.Extension.Context.CacheFullXML()

	return condition.Id
}

func GetDialplanDebug() bool {
	return dialplan.EnableDebug
}

func SwitchDialplanDebug(switcher bool) bool {
	dialplan.EnableDebug = switcher
	return dialplan.EnableDebug
}
