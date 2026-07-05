package web

import (
	"custompbx/cfg"
	"net/http"
	"strings"
)

func CheckWebSocketOrigin(policy string, allowed []string, r *http.Request) bool {
	if policy == cfg.OriginPolicyAllowAll {
		return true
	}
	origin, err := cfg.NormalizeOrigin(r.Header.Get("Origin"))
	if err != nil {
		return false
	}
	if policy == cfg.OriginPolicyAllowList {
		for _, item := range allowed {
			if origin == item {
				return true
			}
		}
		return false
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	if forwarded := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-Proto"), ",")[0]); forwarded == "http" || forwarded == "https" {
		scheme = forwarded
	}
	expected, err := cfg.NormalizeOrigin(scheme + "://" + r.Host)
	return err == nil && origin == expected
}
