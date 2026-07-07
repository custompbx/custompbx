package web

import (
	"custompbx/mainStruct"
	"custompbx/webcache"
	"net/http"
	"strings"
)

type TokenLookupFunc func(token string) (*mainStruct.WebUser, error)

var HTTPTokenLookup TokenLookupFunc = webcache.GetWebUserByToken

func UserFromBearer(r *http.Request) (*mainStruct.WebUser, int) {
	return userFromToken(BearerToken(r))
}

func UserFromCookie(r *http.Request, name string) (*mainStruct.WebUser, int) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, http.StatusUnauthorized
	}
	return userFromToken(cookie.Value)
}

func RequireGroups(user *mainStruct.WebUser, groups ...int) int {
	if user == nil || user.Id == 0 {
		return http.StatusUnauthorized
	}
	for _, group := range groups {
		if user.GroupId == group {
			return http.StatusOK
		}
	}
	return http.StatusForbidden
}

func BearerToken(r *http.Request) string {
	const prefix = "bearer "
	auth := strings.TrimSpace(r.Header.Get("Authorization"))
	if len(auth) <= len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return ""
	}
	return strings.TrimSpace(auth[len(prefix):])
}

func userFromToken(token string) (*mainStruct.WebUser, int) {
	token = strings.TrimSpace(token)
	if token == "" {
		return nil, http.StatusUnauthorized
	}
	user, err := HTTPTokenLookup(token)
	if err != nil || user == nil || user.Id == 0 {
		return nil, http.StatusUnauthorized
	}
	return user, http.StatusOK
}
