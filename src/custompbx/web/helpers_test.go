package web

import (
	"custompbx/mainStruct"
	"testing"

	"github.com/custompbx/customorm"
)

func TestNormalizeConfigPagination(t *testing.T) {
	tests := []struct {
		name       string
		limit      int
		page       int
		wantLimit  int
		wantOffset int
	}{
		{name: "default when empty", limit: 0, page: 0, wantLimit: 25, wantOffset: 0},
		{name: "keeps small positive existing behavior", limit: 10, page: 2, wantLimit: 10, wantOffset: 20},
		{name: "clamps above maximum", limit: 500, page: 1, wantLimit: 25, wantOffset: 25},
		{name: "negative page becomes zero", limit: 50, page: -1, wantLimit: 50, wantOffset: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit, offset := normalizeConfigPagination(tt.limit, tt.page)
			if limit != tt.wantLimit || offset != tt.wantOffset {
				t.Fatalf("got limit=%d offset=%d", limit, offset)
			}
		})
	}
}

func TestBuildFilteredConfigRequestCopiesBaseAndOrder(t *testing.T) {
	base := map[string]customorm.FilterFields{
		"Parent": {Flag: true, UseValue: true, Value: int64(42)},
	}
	request := mainStruct.DBRequest{
		Limit:  50,
		Offset: 2,
		Filters: []mainStruct.Filter{
			{Field: "name", Operand: "LIKE", FieldValue: "%abc%"},
		},
		Order: mainStruct.Order{Desc: true, Fields: []string{"name"}},
	}

	got := buildFilteredConfigRequest(base, request)
	if got.Limit != 50 || got.Offset != 100 {
		t.Fatalf("unexpected pagination: limit=%d offset=%d", got.Limit, got.Offset)
	}
	if got.Fields["Parent"].Value != int64(42) {
		t.Fatalf("parent filter changed: %+v", got.Fields["Parent"])
	}
	if got.Fields["name"].Operand != "LIKE" || got.Fields["name"].Value != "%abc%" {
		t.Fatalf("request filter missing: %+v", got.Fields["name"])
	}
	if _, exists := base["name"]; exists {
		t.Fatal("base filter map was mutated")
	}
	got.Order.Fields[0] = "changed"
	if request.Order.Fields[0] != "name" {
		t.Fatal("order fields were aliased")
	}
}

func TestResponseHelpersKeepMessageType(t *testing.T) {
	errResp := errorResponse("event", "bad")
	if errResp.MessageType != "event" || errResp.Error != "bad" {
		t.Fatalf("unexpected error response: %+v", errResp)
	}
	dataResp := dataResponse("event", 123)
	if dataResp.MessageType != "event" || dataResp.Data.(int) != 123 {
		t.Fatalf("unexpected data response: %+v", dataResp)
	}
}
