package web

import (
	"custompbx/fsesl"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
)

func importDialplan(data *webStruct.MessageData) webStruct.UserResponse {
	fsesl.GetXMLDialplan()

	return webStruct.UserResponse{MessageType: data.Event}
}

func SwitchDialplanNoProceed(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Enabled == nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if !pbxcache.SwitchDialplanNoProceed(*data.Enabled) {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, DialplanSettings: pbxcache.GetDialplan()}
}

func DialplanGetSettings(data *webStruct.MessageData) webStruct.UserResponse {
	items := pbxcache.GetDialplan()
	if items == nil {
		return webStruct.UserResponse{MessageType: data.Event}
	}

	return webStruct.UserResponse{DialplanSettings: items, MessageType: data.Event}
}

func getDialplanContexts(data *webStruct.MessageData) webStruct.UserResponse {
	items, exists := pbxcache.GetContexts()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{DialplanContexts: &items, MessageType: data.Event}
}

func getDialplanExtensions(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	context := pbxcache.GetContext(data.Id)
	if context == nil {
		return webStruct.UserResponse{Error: "context not found", MessageType: data.Event}
	}

	extensions := context.Extensions.Props()

	return webStruct.UserResponse{MessageType: data.Event, DialplanExtensions: &extensions, Id: &context.Id}
}

func getDialplanConditions(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	extension := pbxcache.GetExtension(data.Id)
	if extension == nil {
		return webStruct.UserResponse{Error: "extension not found", MessageType: data.Event}
	}

	conditions := extension.Conditions.Props()
	item := map[int64]*[]*mainStruct.Condition{extension.Id: &conditions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanConditions: &item, Id: &extension.Context.Id}
}

func getDialplanExtenDetails(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	actions := condition.Actions.Props()
	antiActions := condition.AntiActions.Props()
	regexes := condition.Regexes.Props()
	dialplanActions := webStruct.DialplanDetails{DialplanActions: &actions, DialplanAntiActions: &antiActions, DialplanRegexes: &regexes}
	item := map[int64]*webStruct.DialplanDetails{condition.Id: &dialplanActions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &condition.Extension.Context.Id, AffectedId: &condition.Extension.Id}
}

func moveDialplanExtension(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	extension := pbxcache.GetExtension(data.Id)
	if extension == nil || extension.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "extension not found", MessageType: data.Event}
	}

	err := pbxcache.MoveContextExtension(extension, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move extension", MessageType: data.Event}
	}

	extensions := extension.Context.Extensions.Props()

	return webStruct.UserResponse{MessageType: data.Event, DialplanExtensions: &extensions, Id: &extension.Context.Id}
}

func moveDialplanCondition(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Id)
	if condition == nil || condition.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	err := pbxcache.MoveContextCondition(condition, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move condition", MessageType: data.Event}
	}

	conditions := condition.Extension.Conditions.Props()
	item := map[int64]*[]*mainStruct.Condition{condition.Extension.Id: &conditions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanConditions: &item, Id: &condition.Extension.Context.Id}
}

func moveDialplanAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	action := pbxcache.GetAction(data.Id)
	if action == nil || action.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	err := pbxcache.MoveContextAction(action, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move action", MessageType: data.Event}
	}

	actions := action.Condition.Actions.Props()
	item := map[int64]*webStruct.DialplanDetails{action.Condition.Id: {DialplanActions: &actions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &action.Condition.Extension.Context.Id, AffectedId: &action.Condition.Extension.Id}
}

func moveDialplanAntiAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	antiAction := pbxcache.GetAntiAction(data.Id)
	if antiAction == nil || antiAction.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "antiAction not found", MessageType: data.Event}
	}

	err := pbxcache.MoveContextAntiAction(antiAction, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move antiAction", MessageType: data.Event}
	}

	antiActions := antiAction.Condition.AntiActions.Props()
	item := map[int64]*webStruct.DialplanDetails{antiAction.Condition.Id: {DialplanAntiActions: &antiActions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &antiAction.Condition.Extension.Context.Id, AffectedId: &antiAction.Condition.Extension.Id}
}

func addRegex(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Regex.Id != 0 || data.Regex.Field == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	regex, err := pbxcache.SetConditionRegex(condition, data.Regex.Field, data.Regex.Expression)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	regexes := regex.Condition.Regexes.Props()
	item := map[int64]*webStruct.DialplanDetails{regex.Condition.Id: {DialplanRegexes: &regexes}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &regex.Condition.Extension.Context.Id, AffectedId: &regex.Condition.Extension.Id}
}

func addAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Action.Id != 0 || data.Action.Application == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	action, err := pbxcache.SetConditionAction(condition, data.Action.Application, data.Action.Data, data.Action.Inline)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	actions := action.Condition.Actions.Props()
	item := map[int64]*webStruct.DialplanDetails{action.Condition.Id: {DialplanActions: &actions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &action.Condition.Extension.Context.Id, AffectedId: &action.Condition.Extension.Id}
}

func addAntiAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.AntiAction.Id != 0 || data.AntiAction.Application == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	antiAction, err := pbxcache.SetConditionAntiAction(condition, data.AntiAction.Application, data.AntiAction.Data)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	antiActions := antiAction.Condition.AntiActions.Props()
	item := map[int64]*webStruct.DialplanDetails{antiAction.Condition.Id: {DialplanAntiActions: &antiActions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &antiAction.Condition.Extension.Context.Id, AffectedId: &antiAction.Condition.Extension.Id}
}

func delRegex(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Regex.Id == 0 {
		return webStruct.UserResponse{Error: "regex not found", MessageType: data.Event}
	}

	regex := pbxcache.GetRegex(data.Regex.Id)
	if regex == nil {
		return webStruct.UserResponse{Error: "regex not found", MessageType: data.Event}
	}

	res := pbxcache.DelContextRegex(regex)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	regexes := regex.Condition.Regexes.Props()
	item := map[int64]*webStruct.DialplanDetails{regex.Condition.Id: {DialplanRegexes: &regexes}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &regex.Condition.Extension.Context.Id, AffectedId: &regex.Condition.Extension.Id}
}

func delAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Action.Id == 0 {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	action := pbxcache.GetAction(data.Action.Id)
	if action == nil {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	res := pbxcache.DelContextAction(action)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	actions := action.Condition.Actions.Props()
	item := map[int64]*webStruct.DialplanDetails{action.Condition.Id: {DialplanActions: &actions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &action.Condition.Extension.Context.Id, AffectedId: &action.Condition.Extension.Id}
}

func delAntiAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.AntiAction.Id == 0 {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	antiAction := pbxcache.GetAntiAction(data.AntiAction.Id)
	if antiAction == nil {
		return webStruct.UserResponse{Error: "antiAction not found", MessageType: data.Event}
	}

	res := pbxcache.DelContextAntiAction(antiAction)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	antiActions := antiAction.Condition.AntiActions.Props()
	item := map[int64]*webStruct.DialplanDetails{antiAction.Condition.Id: {DialplanAntiActions: &antiActions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &antiAction.Condition.Extension.Context.Id, AffectedId: &antiAction.Condition.Extension.Id}
}

func updateRegex(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Regex.Id == 0 {
		return webStruct.UserResponse{Error: "regex not found", MessageType: data.Event}
	}

	regex := pbxcache.GetRegex(data.Regex.Id)
	if regex == nil {
		return webStruct.UserResponse{Error: "regex not found", MessageType: data.Event}
	}

	res := pbxcache.UpdateContextRegex(regex, data.Regex.Field, data.Regex.Expression)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	regexes := regex.Condition.Regexes.Props()
	item := map[int64]*webStruct.DialplanDetails{regex.Condition.Id: {DialplanRegexes: &regexes}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &regex.Condition.Extension.Context.Id, AffectedId: &regex.Condition.Extension.Id}
}

func updateAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Action.Id == 0 {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	action := pbxcache.GetAction(data.Action.Id)
	if action == nil {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	res := pbxcache.UpdateContextAction(action, data.Action.Application, data.Action.Data, data.Action.Inline)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	actions := action.Condition.Actions.Props()
	item := map[int64]*webStruct.DialplanDetails{action.Condition.Id: {DialplanActions: &actions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &action.Condition.Extension.Context.Id, AffectedId: &action.Condition.Extension.Id}
}

func updateAntiAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.AntiAction.Id == 0 {
		return webStruct.UserResponse{Error: "antiAction not found", MessageType: data.Event}
	}

	antiAction := pbxcache.GetAntiAction(data.AntiAction.Id)
	if antiAction == nil {
		return webStruct.UserResponse{Error: "antiAction not found", MessageType: data.Event}
	}

	res := pbxcache.UpdateContextAntiAction(antiAction, data.AntiAction.Application, data.AntiAction.Data)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	antiActions := antiAction.Condition.AntiActions.Props()
	item := map[int64]*webStruct.DialplanDetails{antiAction.Condition.Id: {DialplanAntiActions: &antiActions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &antiAction.Condition.Extension.Context.Id, AffectedId: &antiAction.Condition.Extension.Id}
}

func switchRegex(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Regex.Id == 0 {
		return webStruct.UserResponse{Error: "regex not found", MessageType: data.Event}
	}

	regex := pbxcache.GetRegex(data.Regex.Id)
	if regex == nil {
		return webStruct.UserResponse{Error: "regex not found", MessageType: data.Event}
	}

	res := pbxcache.SwitchContextRegex(regex, data.Regex.Enabled)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	regexes := regex.Condition.Regexes.Props()
	item := map[int64]*webStruct.DialplanDetails{regex.Condition.Id: {DialplanRegexes: &regexes}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &regex.Condition.Extension.Context.Id, AffectedId: &regex.Condition.Extension.Id}
}

func switchAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Action.Id == 0 {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	action := pbxcache.GetAction(data.Action.Id)
	if action == nil {
		return webStruct.UserResponse{Error: "action not found", MessageType: data.Event}
	}

	res := pbxcache.SwitchContextAction(action, data.Action.Enabled)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	actions := action.Condition.Actions.Props()
	item := map[int64]*webStruct.DialplanDetails{action.Condition.Id: {DialplanActions: &actions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &action.Condition.Extension.Context.Id, AffectedId: &action.Condition.Extension.Id}
}

func switchAntiAction(data *webStruct.MessageData) webStruct.UserResponse {
	if data.AntiAction.Id == 0 {
		return webStruct.UserResponse{Error: "antiAction not found", MessageType: data.Event}
	}

	antiAction := pbxcache.GetAntiAction(data.AntiAction.Id)
	if antiAction == nil {
		return webStruct.UserResponse{Error: "antiAction not found", MessageType: data.Event}
	}

	res := pbxcache.SwitchContextAntiAction(antiAction, data.AntiAction.Enabled)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	antiActions := antiAction.Condition.AntiActions.Props()
	item := map[int64]*webStruct.DialplanDetails{antiAction.Condition.Id: {DialplanAntiActions: &antiActions}}

	return webStruct.UserResponse{MessageType: data.Event, DialplanDetails: &item, Id: &antiAction.Condition.Extension.Context.Id, AffectedId: &antiAction.Condition.Extension.Id}
}

func addContext(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	res, err := pbxcache.SetContext(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.Context)
	items[res.Id] = res

	return webStruct.UserResponse{DialplanContexts: &items, MessageType: data.Event}
}

func addExtension(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	context := pbxcache.GetContext(data.Id)
	if context == nil {
		return webStruct.UserResponse{Error: "context not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetContextExtension(context, data.Name, "")
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := []*mainStruct.Extension{res}

	return webStruct.UserResponse{DialplanExtensions: &items, Id: &context.Id, MessageType: data.Event}
}

func addCondition(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	extension := pbxcache.GetExtension(data.Id)
	if extension == nil {
		return webStruct.UserResponse{Error: "extension not found", MessageType: data.Event}
	}

	condition, err := pbxcache.SetExtensionCondition(extension, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	conditions := []*mainStruct.Condition{condition}
	item := map[int64]*[]*mainStruct.Condition{condition.Extension.Id: &conditions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanConditions: &item, Id: &condition.Extension.Context.Id}
}

func renameContext(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	context := pbxcache.GetContext(data.Id)
	if context == nil {
		return webStruct.UserResponse{Error: "context not found", MessageType: data.Event}
	}

	res, err := pbxcache.RenameContext(context, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.Context)
	items[res.Id] = res

	return webStruct.UserResponse{DialplanContexts: &items, MessageType: data.Event}
}

func renameExtension(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	extension := pbxcache.GetExtension(data.Id)
	if extension == nil {
		return webStruct.UserResponse{Error: "extension not found", MessageType: data.Event}
	}

	res, err := pbxcache.RenameExtension(extension, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := []*mainStruct.Extension{res}

	return webStruct.UserResponse{DialplanExtensions: &items, Id: &extension.Context.Id, MessageType: data.Event}
}

func deleteContext(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	context := pbxcache.GetContext(data.Id)
	if context == nil {
		return webStruct.UserResponse{Error: "context not found", MessageType: data.Event}
	}

	err := pbxcache.DeleteContext(context)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Id: &context.Id, MessageType: data.Event}
}

func deleteExtension(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	extension := pbxcache.GetExtension(data.Id)
	if extension == nil {
		return webStruct.UserResponse{Error: "extension not found", MessageType: data.Event}
	}

	err := pbxcache.DeleteExtension(extension)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{AffectedId: &extension.Id, Id: &extension.Context.Id, MessageType: data.Event}
}

func switchExtensionContinue(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	extension := pbxcache.GetExtension(data.Id)
	if extension == nil {
		return webStruct.UserResponse{Error: "extension not found", MessageType: data.Event}
	}

	if data.Value != "" && data.Value != "true" {
		return webStruct.UserResponse{Error: "wrong value", MessageType: data.Event}
	}

	err := pbxcache.SwitchExtensionContinue(extension, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	items := []*mainStruct.Extension{extension}

	return webStruct.UserResponse{DialplanExtensions: &items, Id: &extension.Context.Id, MessageType: data.Event}
}

func updateCondition(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Condition.Id == 0 {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Condition.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateExtensionCondition(
		condition,
		data.Condition.Break,
		data.Condition.Field,
		data.Condition.Expression,
		data.Condition.Hour,
		data.Condition.Mday,
		data.Condition.Mon,
		data.Condition.Mweek,
		data.Condition.Wday,
		data.Condition.DateTime,
		data.Condition.TimeOfDay,
		data.Condition.Year,
		data.Condition.Minute,
		data.Condition.Week,
		data.Condition.Yday,
		data.Condition.Minday,
		data.Condition.TzOffset,
		data.Condition.Dst,
		data.Condition.Regex,
	)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	conditions := []*mainStruct.Condition{condition}
	item := map[int64]*[]*mainStruct.Condition{condition.Extension.Id: &conditions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanConditions: &item, Id: &condition.Extension.Context.Id}
}

func switchCondition(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Condition.Id == 0 {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Condition.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	res := pbxcache.SwitchContextCondition(condition, data.Condition.Enabled)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	conditions := []*mainStruct.Condition{condition}
	item := map[int64]*[]*mainStruct.Condition{condition.Extension.Id: &conditions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanConditions: &item, Id: &condition.Extension.Context.Id}
}

func deleteCondition(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Condition.Id == 0 {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	condition := pbxcache.GetCondition(data.Condition.Id)
	if condition == nil {
		return webStruct.UserResponse{Error: "condition not found", MessageType: data.Event}
	}

	res := pbxcache.DeleteContextCondition(condition)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	conditions := []*mainStruct.Condition{condition}
	item := map[int64]*[]*mainStruct.Condition{condition.Extension.Id: &conditions}

	return webStruct.UserResponse{MessageType: data.Event, DialplanConditions: &item, Id: &condition.Extension.Context.Id}
}

func getDialplanDebug(data *webStruct.MessageData) webStruct.UserResponse {
	enabled := pbxcache.GetDialplanDebug()
	return webStruct.UserResponse{MessageType: data.Event, Enabled: &enabled}
}

func switchDialplanDebug(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Enabled == nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	enabled := pbxcache.SwitchDialplanDebug(*data.Enabled)
	return webStruct.UserResponse{MessageType: data.Event, Enabled: &enabled}
}
