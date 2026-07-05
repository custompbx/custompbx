package web

import (
	"custompbx/webStruct"
	"testing"
)

func TestHandlerRegistryRejectsDuplicates(t *testing.T) {
	r := newHandlerRegistry()
	h := func(data *webStruct.MessageData) webStruct.UserResponse {
		return webStruct.UserResponse{MessageType: data.Event}
	}
	groups := func() []int { return []int{1} }
	if err := r.Register("event", h, groups); err != nil {
		t.Fatal(err)
	}
	if err := r.Register("event", h, groups); err == nil {
		t.Fatal("duplicate registration accepted")
	}
}

func TestNormalizePagination(t *testing.T) {
	limit, offset := normalizePagination(0, -1)
	if limit != 250 || offset != 0 {
		t.Fatalf("got %d, %d", limit, offset)
	}
	limit, offset = normalizePagination(100, 3)
	if limit != 100 || offset != 300 {
		t.Fatalf("got %d, %d", limit, offset)
	}
}
