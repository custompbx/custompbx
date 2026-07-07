package main

import (
	"custompbx/mainStruct"
	"custompbx/web"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveFileUnderRoot(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "2026")
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "call.wav")
	if err := os.WriteFile(file, []byte("test"), 0600); err != nil {
		t.Fatal(err)
	}
	got, err := resolveFileUnderRoot(root, "2026", "call.wav")
	if err != nil || got != file {
		t.Fatalf("got %q, %v", got, err)
	}
	for _, parts := range [][]string{{"..", "secret"}, {`..\\secret`}, {"2026", "missing.wav"}} {
		if got, err := resolveFileUnderRoot(root, parts...); err == nil {
			t.Fatalf("unsafe path accepted: %q", got)
		}
	}
}

func TestWebProtectedStaticRoutesRequireValidCookie(t *testing.T) {
	oldLookup := web.HTTPTokenLookup
	defer func() { web.HTTPTokenLookup = oldLookup }()
	web.HTTPTokenLookup = func(token string) (*mainStruct.WebUser, error) {
		if token == "valid" {
			return &mainStruct.WebUser{Id: 1, Login: "admin", GroupId: mainStruct.GetAdminId()}, nil
		}
		return nil, nil
	}

	router := chi.NewRouter()
	configureStaticRoutes(router)

	tests := []struct {
		name   string
		path   string
		token  string
		status int
	}{
		{name: "cdr missing cookie", path: "/cweb/cdr/records/call.wav", status: http.StatusUnauthorized},
		{name: "cdr invalid cookie", path: "/cweb/cdr/records/call.wav", token: "bad", status: http.StatusUnauthorized},
		{name: "avatar missing cookie", path: "/cweb/assets/img/avatar/1.png", status: http.StatusUnauthorized},
		{name: "avatar invalid cookie", path: "/cweb/assets/img/avatar/1.png", token: "bad", status: http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			if tt.token != "" {
				req.AddCookie(&http.Cookie{Name: "token", Value: tt.token})
			}
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if rr.Code != tt.status {
				t.Fatalf("status=%d, want %d, body=%q", rr.Code, tt.status, rr.Body.String())
			}
		})
	}
}
