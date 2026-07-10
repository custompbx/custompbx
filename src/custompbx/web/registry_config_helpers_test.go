package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
	"strings"
	"testing"
)

func TestCallcenterDynamicUpdatePayloadMapsAgentFields(t *testing.T) {
	data := &webStruct.MessageData{Event: eventCallcenterAgentUpdate}
	data.Param = webStruct.Param{Id: 7, Name: "max_no_answer", Value: "5"}

	payload, errText := callcenterDynamicUpdatePayload(data, &altStruct.Agent{Id: data.Param.Id})

	if errText != "" {
		t.Fatalf("unexpected error: %s", errText)
	}
	update, ok := payload.(struct {
		S interface{}
		A []string
	})
	if !ok {
		t.Fatalf("payload type = %T", payload)
	}
	if len(update.A) != 1 || update.A[0] != "MaxNoAnswer" {
		t.Fatalf("updated fields = %#v", update.A)
	}
	agent, ok := update.S.(*altStruct.Agent)
	if !ok {
		t.Fatalf("updated item type = %T", update.S)
	}
	if agent.Id != 7 || agent.MaxNoAnswer != 5 {
		t.Fatalf("agent update = %+v", agent)
	}
}

func TestCallcenterDynamicUpdatePayloadRejectsBadFieldsAndValues(t *testing.T) {
	tests := []struct {
		name  string
		param webStruct.Param
		want  string
	}{
		{name: "missing name", param: webStruct.Param{Value: "5"}, want: "wrong params"},
		{name: "id field", param: webStruct.Param{Name: "id", Value: "5"}, want: "please dont"},
		{name: "unknown field", param: webStruct.Param{Name: "missing", Value: "5"}, want: "unknown field"},
		{name: "bad int", param: webStruct.Param{Name: "level", Value: "not-int"}, want: "invalid syntax"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &webStruct.MessageData{Event: eventCallcenterTierUpdate, Param: tt.param}

			_, errText := callcenterDynamicUpdatePayload(data, &altStruct.Tier{Id: data.Param.Id})

			if !strings.Contains(errText, tt.want) {
				t.Fatalf("error = %q, want containing %q", errText, tt.want)
			}
		})
	}
}

func TestConfigUpdatePayloadUsesUpdateConfigShape(t *testing.T) {
	item := &altStruct.ConfigVertoSetting{Id: 7, Name: "debug"}

	payload := configUpdatePayload(item, "Name", "Value")

	update, ok := payload.(struct {
		S interface{}
		A []string
	})
	if !ok {
		t.Fatalf("payload type = %T", payload)
	}
	if update.S != item {
		t.Fatalf("payload item = %#v, want original item", update.S)
	}
	if strings.Join(update.A, ",") != "Name,Value" {
		t.Fatalf("payload fields = %#v", update.A)
	}
}

func TestConfigWithFieldsBuildsTopLevelParentedItem(t *testing.T) {
	item := configWithFields(&altStruct.ConfigVertoProfile{}, map[string]interface{}{
		"Name":    "profile-a",
		"Enabled": true,
		"Parent":  &altStruct.ConfigurationsList{Id: 46},
	})

	profile, ok := item.(*altStruct.ConfigVertoProfile)
	if !ok {
		t.Fatalf("item type = %T", item)
	}
	if profile.Name != "profile-a" || !profile.Enabled || profile.Parent == nil || profile.Parent.Id != 46 {
		t.Fatalf("profile = %+v", profile)
	}
}

func TestConfigWithFieldsBuildsChildParentedItem(t *testing.T) {
	item := configWithFields(&altStruct.ConfigVertoProfileParameter{}, map[string]interface{}{
		"Name":    "bind-local",
		"Value":   "127.0.0.1:8081",
		"Secure":  "true",
		"Enabled": true,
		"Parent":  &altStruct.ConfigVertoProfile{Id: 3},
	})

	param, ok := item.(*altStruct.ConfigVertoProfileParameter)
	if !ok {
		t.Fatalf("item type = %T", item)
	}
	if param.Name != "bind-local" || param.Value != "127.0.0.1:8081" || param.Secure != "true" || !param.Enabled || param.Parent == nil || param.Parent.Id != 3 {
		t.Fatalf("param = %+v", param)
	}
}

func TestConfigIDNameFieldHelpers(t *testing.T) {
	data := &webStruct.MessageData{Id: 11, Name: "profile-a"}
	got := configDataIDName(data)
	if got["Id"] != int64(11) || got["Name"] != "profile-a" {
		t.Fatalf("configDataIDName = %#v", got)
	}

	data.Param.Id = 22
	data.Param.Name = "param-a"
	got = configParamIDName(data)
	if got["Id"] != int64(22) || got["Name"] != "param-a" {
		t.Fatalf("configParamIDName = %#v", got)
	}
}

func TestNamedConfigMutationRegistrationUsesAccessChecksAndOverrides(t *testing.T) {
	events := namedConfigEvents{Add: "named-add", Update: "named-update", Delete: "named-delete"}
	calls := map[string]int{}
	r := newHandlerRegistry()
	registerNamedConfigMutationsForSample(
		r,
		map[string]eventHandler{
			events.Add: func(data *webStruct.MessageData) webStruct.UserResponse {
				calls[data.Event]++
				return webStruct.UserResponse{MessageType: data.Event}
			},
			events.Update: func(data *webStruct.MessageData) webStruct.UserResponse {
				calls[data.Event]++
				return webStruct.UserResponse{MessageType: data.Event}
			},
			events.Delete: func(data *webStruct.MessageData) webStruct.UserResponse {
				calls[data.Event]++
				return webStruct.UserResponse{MessageType: data.Event}
			},
		},
		events,
		&altStruct.ConfigVoicemailProfile{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} { return &altStruct.ConfigurationsList{Id: 1} },
	)

	for _, event := range []string{events.Add, events.Update, events.Delete} {
		ctx := adminContext()
		resp, ok := r.Dispatch(messageData(ctx, event), ctx)
		if !ok || resp.MessageType != event {
			t.Fatalf("%s dispatch resp=%+v ok=%t", event, resp, ok)
		}
		if calls[event] != 1 {
			t.Fatalf("%s calls = %d, want 1", event, calls[event])
		}

		userCtx := userContext()
		resp, ok = r.Dispatch(messageData(userCtx, event), userCtx)
		if !ok || resp.MessageType != "no_access" {
			t.Fatalf("%s unauthorized resp=%+v ok=%t", event, resp, ok)
		}
	}
}

func TestParentedParamConfigMutationRegistrationUsesAccessChecksAndOverrides(t *testing.T) {
	events := parentedParamConfigEvents{Add: "param-add", Delete: "param-delete", Switch: "param-switch", Update: "param-update"}
	calls := map[string]int{}
	overrides := map[string]eventHandler{}
	for _, event := range []string{events.Add, events.Delete, events.Switch, events.Update} {
		event := event
		overrides[event] = func(data *webStruct.MessageData) webStruct.UserResponse {
			calls[data.Event]++
			return webStruct.UserResponse{MessageType: data.Event}
		}
	}

	r := newHandlerRegistry()
	registerParentedParamConfigMutationsForSample(
		r,
		overrides,
		events,
		&altStruct.ConfigVoicemailProfileParameter{},
		func(data *webStruct.MessageData) interface{} { return &altStruct.ConfigVoicemailProfile{Id: data.Id} },
	)

	for _, event := range []string{events.Add, events.Delete, events.Switch, events.Update} {
		ctx := adminContext()
		resp, ok := r.Dispatch(messageData(ctx, event), ctx)
		if !ok || resp.MessageType != event {
			t.Fatalf("%s dispatch resp=%+v ok=%t", event, resp, ok)
		}
		if calls[event] != 1 {
			t.Fatalf("%s calls = %d, want 1", event, calls[event])
		}

		userCtx := userContext()
		resp, ok = r.Dispatch(messageData(userCtx, event), userCtx)
		if !ok || resp.MessageType != "no_access" {
			t.Fatalf("%s unauthorized resp=%+v ok=%t", event, resp, ok)
		}
	}
}

func TestCombinedDataResponse(t *testing.T) {
	resp := combinedDataResponse("event",
		responseDataPair{name: "settings", data: "settings-data"},
		responseDataPair{name: "profiles", data: "profiles-data"},
	)
	if resp.MessageType != "event" {
		t.Fatalf("message type = %q, want event", resp.MessageType)
	}
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("data type = %T, want map[string]interface{}", resp.Data)
	}
	if data["settings"] != "settings-data" {
		t.Fatalf("settings data = %v", data["settings"])
	}
	if data["profiles"] != "profiles-data" {
		t.Fatalf("profiles data = %v", data["profiles"])
	}
}
