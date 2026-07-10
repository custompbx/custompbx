package web

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestMigratedAggregateResponseShapes(t *testing.T) {
	tests := []struct {
		name  string
		event string
		keys  []string
	}{
		{
			name:  "conference config",
			event: "GetConference",
			keys: []string{
				"conference_rooms",
				"conference_profiles",
				"conference_caller_control_groups",
				"conference_chat_permissions_profiles",
			},
		},
		{
			name:  "conference layouts",
			event: "GetConferenceLayouts",
			keys: []string{
				"conference_layouts",
				"conference_layouts_groups",
			},
		},
		{
			name:  "verto config",
			event: "[Config][Verto][Get]",
			keys: []string{
				"settings",
				"profiles",
			},
		},
		{
			name:  "callcenter import",
			event: "ImportCallcenterAgentsAndTiers",
			keys: []string{
				"callcenter_agents",
				"callcenter_tiers",
			},
		},
		{
			name:  "http cache config",
			event: "GetHttpCache",
			keys: []string{
				"settings",
				"profiles",
			},
		},
		{
			name:  "http cache profile parameters",
			event: "GetHttpCacheProfileParameters",
			keys: []string{
				"domains",
				"azure",
				"aws_s3",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs := make([]responseDataPair, 0, len(tt.keys))
			for _, key := range tt.keys {
				pairs = append(pairs, responseDataPair{name: key, data: key + "-data"})
			}
			resp := combinedDataResponse(tt.event, pairs...)

			if resp.MessageType != tt.event {
				t.Fatalf("message type = %q, want %q", resp.MessageType, tt.event)
			}
			body, err := json.Marshal(resp.Data)
			if err != nil {
				t.Fatal(err)
			}
			var got map[string]string
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatal(err)
			}
			if len(got) != len(tt.keys) {
				t.Fatalf("keys = %v, want %v", sortedMapKeys(got), tt.keys)
			}
			for _, key := range tt.keys {
				if got[key] != key+"-data" {
					t.Fatalf("key %q = %q, want %q; all keys=%v", key, got[key], key+"-data", sortedMapKeys(got))
				}
			}
		})
	}
}

func sortedMapKeys[V any](items map[string]V) []string {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
