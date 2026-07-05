package web

import (
	"custompbx/cfg"
	"net/http/httptest"
	"testing"
)

func TestCheckWebSocketOrigin(t *testing.T) {
	req := httptest.NewRequest("GET", "https://pbx.example/ws", nil)
	req.Header.Set("Origin", "https://pbx.example")
	if !CheckWebSocketOrigin(cfg.OriginPolicySameOrigin, nil, req) {
		t.Fatal("same origin rejected")
	}
	req.Header.Set("Origin", "https://evil.example")
	if CheckWebSocketOrigin(cfg.OriginPolicySameOrigin, nil, req) {
		t.Fatal("cross origin accepted")
	}
	if !CheckWebSocketOrigin(cfg.OriginPolicyAllowAll, nil, req) {
		t.Fatal("allow_all rejected")
	}
	if !CheckWebSocketOrigin(cfg.OriginPolicyAllowList, []string{"https://evil.example"}, req) {
		t.Fatal("allow-list origin rejected")
	}
}
