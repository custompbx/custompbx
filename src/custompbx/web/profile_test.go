package web

import (
	"testing"

	"custompbx/mainStruct"
	"custompbx/webStruct"
)

func TestRequireOwnWebUserOrAdmin(t *testing.T) {
	tests := []struct {
		name      string
		user      *mainStruct.WebUser
		targetID  int64
		wantError string
	}{
		{
			name:     "user may update self",
			user:     &mainStruct.WebUser{Id: 7, GroupId: mainStruct.GetUserId()},
			targetID: 7,
		},
		{
			name:      "user may not update another user",
			user:      &mainStruct.WebUser{Id: 7, GroupId: mainStruct.GetUserId()},
			targetID:  8,
			wantError: "access denied",
		},
		{
			name:     "administrator may update another user",
			user:     &mainStruct.WebUser{Id: 1, GroupId: mainStruct.GetAdminId()},
			targetID: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			context := webStruct.CreateWsContext(nil)
			context.SetUser(test.user)
			response := requireOwnWebUserOrAdmin(&webStruct.MessageData{
				Event:   "profile update",
				Id:      test.targetID,
				Context: context,
			})
			if test.wantError == "" {
				if response != nil {
					t.Fatalf("unexpected error: %s", response.Error)
				}
				return
			}
			if response == nil || response.Error != test.wantError {
				t.Fatalf("error = %#v, want %q", response, test.wantError)
			}
		})
	}
}
