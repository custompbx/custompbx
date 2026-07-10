package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/webStruct"
	"fmt"
	"reflect"
)

type responseDataPair struct {
	name string
	data interface{}
}

type simpleParamConfigEvents struct {
	Get    string
	Update string
	Switch string
	Add    string
	Delete string
}

type namedConfigEvents struct {
	Add    string
	Update string
	Delete string
}

type parentedParamConfigEvents struct {
	Add    string
	Delete string
	Switch string
	Update string
}

func configGet(sample interface{}) eventHandler {
	return func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUserForConfig(data, getConfig, sample, adminOnly())
	}
}

func configSet(factory func(*webStruct.MessageData) interface{}) eventHandler {
	return func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUserForConfig(data, setConfig, factory(data), adminOnly())
	}
}

func configUpdate(factory func(*webStruct.MessageData) interface{}, fields ...string) eventHandler {
	return func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUserForConfig(data, updateConfig, configUpdatePayload(factory(data), fields...), adminOnly())
	}
}

func configDelete(factory func(*webStruct.MessageData) interface{}) eventHandler {
	return func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUserForConfig(data, delConfig, factory(data), adminOnly())
	}
}

func configParentFor(sample interface{}) *altStruct.ConfigurationsList {
	return getConfParent(altData.GetConfNameByStruct(sample))
}

func configUpdatePayload(item interface{}, fields ...string) interface{} {
	return struct {
		S interface{}
		A []string
	}{S: item, A: fields}
}

func configWithFields(sample interface{}, fields map[string]interface{}) interface{} {
	return configValue(sample, func(v reflect.Value) {
		for name, value := range fields {
			setConfigField(v, name, value)
		}
	})
}

func configSetWithFields(sample interface{}, fields func(*webStruct.MessageData) map[string]interface{}) eventHandler {
	return configSet(func(data *webStruct.MessageData) interface{} {
		return configWithFields(sample, fields(data))
	})
}

func configUpdateWithFields(sample interface{}, updateFields []string, fields func(*webStruct.MessageData) map[string]interface{}) eventHandler {
	return configUpdate(func(data *webStruct.MessageData) interface{} {
		return configWithFields(sample, fields(data))
	}, updateFields...)
}

func configDeleteWithFields(sample interface{}, fields func(*webStruct.MessageData) map[string]interface{}) eventHandler {
	return configDelete(func(data *webStruct.MessageData) interface{} {
		return configWithFields(sample, fields(data))
	})
}

func configGetNamedID(data *webStruct.MessageData) map[string]interface{} {
	return map[string]interface{}{"Id": data.Id}
}

func configGetParamID(data *webStruct.MessageData) map[string]interface{} {
	return map[string]interface{}{"Id": data.Param.Id}
}

func configDataIDName(data *webStruct.MessageData) map[string]interface{} {
	return map[string]interface{}{"Id": data.Id, "Name": data.Name}
}

func configParamIDName(data *webStruct.MessageData) map[string]interface{} {
	return map[string]interface{}{"Id": data.Param.Id, "Name": data.Param.Name}
}

func configSetTopLevelName(sample interface{}, name string) map[string]interface{} {
	return map[string]interface{}{"Name": name, "Enabled": true, "Parent": configParentFor(sample)}
}

func configSetChildParamNameValue(parent interface{}) func(*webStruct.MessageData) map[string]interface{} {
	return func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Name": data.Param.Name, "Value": data.Param.Value, "Enabled": true, "Parent": parent}
	}
}

func configUpdateParamNameValue(data *webStruct.MessageData) map[string]interface{} {
	return map[string]interface{}{"Id": data.Param.Id, "Name": data.Param.Name, "Value": data.Param.Value}
}

func configSwitchParamEnabled(data *webStruct.MessageData) map[string]interface{} {
	return map[string]interface{}{"Id": data.Param.Id, "Enabled": data.Param.Enabled}
}

func combinedDataResponse(event string, pairs ...responseDataPair) webStruct.UserResponse {
	data := make(map[string]interface{}, len(pairs))
	for _, pair := range pairs {
		data[pair.name] = pair.data
	}
	return webStruct.UserResponse{MessageType: event, Data: data}
}

func registerNamedConfigMutationsForSample(
	r *handlerRegistry,
	overrides map[string]eventHandler,
	events namedConfigEvents,
	sample interface{},
	name func(*webStruct.MessageData) string,
	parent func(*webStruct.MessageData) interface{},
) {
	if events.Add != "" {
		mustRegisterAdmin(r, events.Add, configSetWithFields(sample, func(data *webStruct.MessageData) map[string]interface{} {
			return map[string]interface{}{"Name": name(data), "Enabled": true, "Parent": parent(data)}
		}), overrides)
	}
	if events.Update != "" {
		mustRegisterAdmin(r, events.Update, configUpdateWithFields(sample, []string{"Name"}, configDataIDName), overrides)
	}
	if events.Delete != "" {
		mustRegisterAdmin(r, events.Delete, configDeleteWithFields(sample, configGetNamedID), overrides)
	}
}

func registerParentedParamConfigMutationsForSample(
	r *handlerRegistry,
	overrides map[string]eventHandler,
	events parentedParamConfigEvents,
	sample interface{},
	parent func(*webStruct.MessageData) interface{},
) {
	if events.Add != "" {
		mustRegisterAdmin(r, events.Add, configSetWithFields(sample, func(data *webStruct.MessageData) map[string]interface{} {
			return map[string]interface{}{"Name": data.Param.Name, "Value": data.Param.Value, "Enabled": true, "Parent": parent(data)}
		}), overrides)
	}
	if events.Delete != "" {
		mustRegisterAdmin(r, events.Delete, configDeleteWithFields(sample, configGetParamID), overrides)
	}
	if events.Switch != "" {
		mustRegisterAdmin(r, events.Switch, configUpdateWithFields(sample, []string{"Enabled"}, configSwitchParamEnabled), overrides)
	}
	if events.Update != "" {
		mustRegisterAdmin(r, events.Update, configUpdateWithFields(sample, []string{"Name", "Value"}, configUpdateParamNameValue), overrides)
	}
}

func registerSimpleParamConfig(
	r *handlerRegistry,
	overrides map[string]eventHandler,
	events simpleParamConfigEvents,
	empty interface{},
	add func(*webStruct.MessageData) interface{},
	update func(*webStruct.MessageData) interface{},
	switchEnabled func(*webStruct.MessageData) interface{},
	delete func(*webStruct.MessageData) interface{},
) {
	mustRegisterAdmin(r, events.Get, configGet(empty), overrides)
	mustRegisterAdmin(r, events.Update, configUpdate(update, "Name", "Value"), overrides)
	mustRegisterAdmin(r, events.Switch, configUpdate(switchEnabled, "Enabled"), overrides)
	mustRegisterAdmin(r, events.Add, configSet(add), overrides)
	mustRegisterAdmin(r, events.Delete, configDelete(delete), overrides)
}

func registerSimpleParamConfigForSample(r *handlerRegistry, overrides map[string]eventHandler, events simpleParamConfigEvents, sample interface{}) {
	validateSimpleParamSample(sample)
	registerSimpleParamConfig(
		r,
		overrides,
		events,
		sample,
		func(data *webStruct.MessageData) interface{} {
			return simpleParamConfigValue(sample, func(v reflect.Value) {
				setSimpleParamField(v, "Name", data.Param.Name)
				setSimpleParamField(v, "Value", data.Param.Value)
				setSimpleParamField(v, "Enabled", true)
				setSimpleParamField(v, "Parent", configParentFor(sample))
			})
		},
		func(data *webStruct.MessageData) interface{} {
			return simpleParamConfigValue(sample, func(v reflect.Value) {
				setSimpleParamField(v, "Id", data.Param.Id)
				setSimpleParamField(v, "Name", data.Param.Name)
				setSimpleParamField(v, "Value", data.Param.Value)
			})
		},
		func(data *webStruct.MessageData) interface{} {
			return simpleParamConfigValue(sample, func(v reflect.Value) {
				setSimpleParamField(v, "Id", data.Param.Id)
				setSimpleParamField(v, "Enabled", data.Param.Enabled)
			})
		},
		func(data *webStruct.MessageData) interface{} {
			return simpleParamConfigValue(sample, func(v reflect.Value) {
				setSimpleParamField(v, "Id", data.Param.Id)
			})
		},
	)
}

func registerSimpleParamConfigMutationsForSample(r *handlerRegistry, overrides map[string]eventHandler, events simpleParamConfigEvents, sample interface{}) {
	validateSimpleParamSample(sample)
	mustRegisterAdmin(r, events.Update, configUpdate(func(data *webStruct.MessageData) interface{} {
		return simpleParamConfigValue(sample, func(v reflect.Value) {
			setSimpleParamField(v, "Id", data.Param.Id)
			setSimpleParamField(v, "Name", data.Param.Name)
			setSimpleParamField(v, "Value", data.Param.Value)
		})
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, events.Switch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return simpleParamConfigValue(sample, func(v reflect.Value) {
			setSimpleParamField(v, "Id", data.Param.Id)
			setSimpleParamField(v, "Enabled", data.Param.Enabled)
		})
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, events.Add, configSet(func(data *webStruct.MessageData) interface{} {
		return simpleParamConfigValue(sample, func(v reflect.Value) {
			setSimpleParamField(v, "Name", data.Param.Name)
			setSimpleParamField(v, "Value", data.Param.Value)
			setSimpleParamField(v, "Enabled", true)
			setSimpleParamField(v, "Parent", configParentFor(sample))
		})
	}), overrides)
	mustRegisterAdmin(r, events.Delete, configDelete(func(data *webStruct.MessageData) interface{} {
		return simpleParamConfigValue(sample, func(v reflect.Value) {
			setSimpleParamField(v, "Id", data.Param.Id)
		})
	}), overrides)
}

func validateSimpleParamSample(sample interface{}) {
	v := newSimpleParamConfigValue(sample)
	for _, field := range []string{"Id", "Name", "Value", "Enabled", "Parent"} {
		f := v.FieldByName(field)
		if !f.IsValid() || !f.CanSet() {
			panic(fmt.Sprintf("simple config sample %T is missing settable %s field", sample, field))
		}
	}
}

func simpleParamConfigValue(sample interface{}, fill func(reflect.Value)) interface{} {
	v := newConfigValue(sample)
	fill(v)
	return v.Addr().Interface()
}

func newSimpleParamConfigValue(sample interface{}) reflect.Value {
	return newConfigValue(sample)
}

func configValue(sample interface{}, fill func(reflect.Value)) interface{} {
	v := newConfigValue(sample)
	fill(v)
	return v.Addr().Interface()
}

func newConfigValue(sample interface{}) reflect.Value {
	t := reflect.TypeOf(sample)
	if t == nil {
		panic("config sample is nil")
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("config sample %T must be a struct or pointer to struct", sample))
	}
	return reflect.New(t).Elem()
}

func setSimpleParamField(v reflect.Value, name string, value interface{}) {
	field := v.FieldByName(name)
	if !field.IsValid() || !field.CanSet() {
		panic(fmt.Sprintf("simple config value %s is missing settable %s field", v.Type(), name))
	}
	rv := reflect.ValueOf(value)
	if !rv.Type().AssignableTo(field.Type()) {
		panic(fmt.Sprintf("simple config field %s.%s expects %s, got %s", v.Type(), name, field.Type(), rv.Type()))
	}
	field.Set(rv)
}

func setConfigField(v reflect.Value, name string, value interface{}) {
	field := v.FieldByName(name)
	if !field.IsValid() || !field.CanSet() {
		panic(fmt.Sprintf("config value %s is missing settable %s field", v.Type(), name))
	}
	rv := reflect.ValueOf(value)
	if !rv.Type().AssignableTo(field.Type()) {
		panic(fmt.Sprintf("config field %s.%s expects %s, got %s", v.Type(), name, field.Type(), rv.Type()))
	}
	field.Set(rv)
}
