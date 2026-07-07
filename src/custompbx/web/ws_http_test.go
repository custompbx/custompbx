package web

import (
	"custompbx/cfg"
	"custompbx/mainStruct"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestBearerToken(t *testing.T) {
	tests := []struct {
		name string
		auth string
		want string
	}{
		{name: "missing"},
		{name: "basic rejected", auth: "Basic abc"},
		{name: "bearer", auth: "Bearer abc", want: "abc"},
		{name: "lowercase bearer", auth: "bearer abc", want: "abc"},
		{name: "trim token", auth: "Bearer   abc  ", want: "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/ws/metrics", nil)
			if tt.auth != "" {
				req.Header.Set("Authorization", tt.auth)
			}
			if got := requestBearerToken(req); got != tt.want {
				t.Fatalf("token = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHubMetricsRequiresAdminBearerToken(t *testing.T) {
	oldLookup := metricsTokenLookup
	defer func() { metricsTokenLookup = oldLookup }()
	metricsTokenLookup = func(token string) (*mainStruct.WebUser, error) {
		switch token {
		case "admin":
			return &mainStruct.WebUser{Id: 1, GroupId: mainStruct.GetAdminId()}, nil
		case "user":
			return &mainStruct.WebUser{Id: 2, GroupId: mainStruct.GetUserId()}, nil
		default:
			return nil, nil
		}
	}

	tests := []struct {
		name   string
		token  string
		status int
	}{
		{name: "missing", status: http.StatusUnauthorized},
		{name: "unknown", token: "bad", status: http.StatusUnauthorized},
		{name: "non admin", token: "user", status: http.StatusForbidden},
		{name: "admin", token: "admin", status: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/ws/metrics", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			rr := httptest.NewRecorder()

			HubMetrics(rr, req)

			if rr.Code != tt.status {
				t.Fatalf("status = %d, want %d, body=%q", rr.Code, tt.status, rr.Body.String())
			}
			if tt.status == http.StatusOK && !strings.Contains(rr.Body.String(), "active") {
				t.Fatalf("metrics response does not look like JSON metrics: %q", rr.Body.String())
			}
		})
	}
}

func TestPostAPIRequestRejectsMalformedMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1", strings.NewReader(`{"event":"x","data":null}`))
	rr := httptest.NewRecorder()

	PostAPIRequest(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusBadRequest)
	}
}

func TestStartWSAppliesOriginPolicy(t *testing.T) {
	oldPolicy := cfg.CustomPbx.Web.OriginPolicy
	oldAllowed := cfg.CustomPbx.Web.AllowedOrigins
	defer func() {
		cfg.CustomPbx.Web.OriginPolicy = oldPolicy
		cfg.CustomPbx.Web.AllowedOrigins = oldAllowed
	}()

	t.Run("rejects cross origin by default", func(t *testing.T) {
		cfg.CustomPbx.Web.OriginPolicy = cfg.OriginPolicySameOrigin
		cfg.CustomPbx.Web.AllowedOrigins = nil

		server := httptest.NewServer(http.HandlerFunc(StartWS))
		defer server.Close()

		header := http.Header{"Origin": []string{"https://evil.example"}}
		conn, resp, err := websocket.DefaultDialer.Dial(wsURL(server.URL), header)
		if conn != nil {
			_ = conn.Close()
		}
		if err == nil {
			t.Fatal("cross-origin websocket upgrade succeeded")
		}
		if resp == nil || resp.StatusCode != http.StatusForbidden {
			t.Fatalf("status = %v, want 403", responseStatus(resp))
		}
	})

	t.Run("allows explicit development allow all", func(t *testing.T) {
		cfg.CustomPbx.Web.OriginPolicy = cfg.OriginPolicyAllowAll
		cfg.CustomPbx.Web.AllowedOrigins = nil

		server := httptest.NewServer(http.HandlerFunc(StartWS))
		defer server.Close()

		header := http.Header{"Origin": []string{"https://evil.example"}}
		conn, resp, err := websocket.DefaultDialer.Dial(wsURL(server.URL), header)
		if err != nil {
			t.Fatalf("websocket upgrade failed: status=%v err=%v", responseStatus(resp), err)
		}
		_ = conn.Close()
	})
}

func wsURL(httpURL string) string {
	return "ws" + strings.TrimPrefix(httpURL, "http")
}

func responseStatus(resp *http.Response) interface{} {
	if resp == nil {
		return nil
	}
	return resp.StatusCode
}
