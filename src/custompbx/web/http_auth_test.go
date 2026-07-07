package web

import (
	"custompbx/mainStruct"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBearerToken(t *testing.T) {
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
			if got := BearerToken(req); got != tt.want {
				t.Fatalf("token = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHTTPUserAuthentication(t *testing.T) {
	oldLookup := HTTPTokenLookup
	defer func() { HTTPTokenLookup = oldLookup }()
	HTTPTokenLookup = func(token string) (*mainStruct.WebUser, error) {
		switch token {
		case "admin":
			return &mainStruct.WebUser{Id: 1, Login: "admin", GroupId: mainStruct.GetAdminId()}, nil
		case "user":
			return &mainStruct.WebUser{Id: 2, Login: "user", GroupId: mainStruct.GetUserId()}, nil
		case "error":
			return nil, errors.New("lookup failed")
		default:
			return nil, nil
		}
	}

	t.Run("bearer", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer admin")
		user, status := UserFromBearer(req)
		if status != http.StatusOK || user.Login != "admin" {
			t.Fatalf("user=%+v status=%d", user, status)
		}
	})

	t.Run("cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "token", Value: "user"})
		user, status := UserFromCookie(req, "token")
		if status != http.StatusOK || user.Login != "user" {
			t.Fatalf("user=%+v status=%d", user, status)
		}
	})

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		if _, status := UserFromBearer(req); status != http.StatusUnauthorized {
			t.Fatalf("status=%d, want %d", status, http.StatusUnauthorized)
		}
	})

	t.Run("lookup error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer error")
		if _, status := UserFromBearer(req); status != http.StatusUnauthorized {
			t.Fatalf("status=%d, want %d", status, http.StatusUnauthorized)
		}
	})
}

func TestRequireGroups(t *testing.T) {
	if got := RequireGroups(nil, mainStruct.GetAdminId()); got != http.StatusUnauthorized {
		t.Fatalf("nil user status=%d", got)
	}
	user := &mainStruct.WebUser{Id: 1, GroupId: mainStruct.GetUserId()}
	if got := RequireGroups(user, mainStruct.GetAdminId()); got != http.StatusForbidden {
		t.Fatalf("wrong group status=%d", got)
	}
	if got := RequireGroups(user, mainStruct.GetAdminId(), mainStruct.GetUserId()); got != http.StatusOK {
		t.Fatalf("matching group status=%d", got)
	}
}
