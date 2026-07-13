package cache

import (
	"custompbx/altStruct"
	"testing"
)

func TestUserCallStateTracksMultipleChannels(t *testing.T) {
	directory := NewFsDirectoryCache()
	user := &altStruct.DirectoryDomainUser{Id: 5, Name: "1004"}

	directory.UserCache.SetCall(user, UserCallState{UUID: "ringing", CreatedAt: 100, Direction: "inbound"}, true)
	directory.UserCache.SetCall(user, UserCallState{UUID: "answered", CreatedAt: 200, Direction: "outbound", Talking: true}, true)

	if !user.InCall || !user.Talking {
		t.Fatalf("expected active talking user, got in_call=%v talking=%v", user.InCall, user.Talking)
	}
	if user.CallDate != 100 {
		t.Fatalf("call date must remain the oldest active creation time, got %d", user.CallDate)
	}
	if user.LastUuid != "answered" {
		t.Fatalf("talking channel should be selected, got %q", user.LastUuid)
	}

	directory.UserCache.SetCall(user, UserCallState{UUID: "answered"}, false)
	if !user.InCall || user.Talking || user.CallDate != 100 || user.LastUuid != "ringing" {
		t.Fatalf("removing one channel must preserve the other: %+v", user)
	}

	directory.UserCache.SetCall(user, UserCallState{UUID: "ringing"}, false)
	if user.InCall || user.Talking || user.CallDate != 0 || user.LastUuid != "" || user.CallDirection != "" {
		t.Fatalf("last channel removal must clear call state: %+v", user)
	}
}

func TestReplaceCallsClearsStaleState(t *testing.T) {
	directory := NewFsDirectoryCache()
	stale := &altStruct.DirectoryDomainUser{Id: 5, Name: "1004"}
	directory.UserCache.SetCall(stale, UserCallState{UUID: "old", CreatedAt: 100, Talking: true}, true)

	directory.UserCache.ReplaceCalls(nil, map[int64]*altStruct.DirectoryDomainUser{stale.Id: stale})

	if stale.InCall || stale.Talking || stale.CallDate != 0 || stale.LastUuid != "" {
		t.Fatalf("snapshot replacement must clear absent channels: %+v", stale)
	}
}
