package mainStruct

import "testing"

func TestChannelsOperationsAreIdempotent(t *testing.T) {
	channels := NewChannelsCache()
	channels.Set(&Channel{Uuid: "one"})
	channels.Set(&Channel{Uuid: "one"})
	if total, answered := channels.Counters(); total != 1 || answered != 0 {
		t.Fatalf("duplicate create changed counters: total=%d answered=%d", total, answered)
	}

	channels.MarkAnswered("one")
	channels.MarkAnswered("one")
	if total, answered := channels.Counters(); total != 1 || answered != 1 {
		t.Fatalf("duplicate answer changed counters: total=%d answered=%d", total, answered)
	}

	channels.RemoveByUUID("one", "")
	channels.RemoveByUUID("one", "")
	if total, answered := channels.Counters(); total != 0 || answered != 0 {
		t.Fatalf("duplicate destroy changed counters: total=%d answered=%d", total, answered)
	}
}

func TestChannelsReplaceRemovesMissingEntries(t *testing.T) {
	channels := NewChannelsCache()
	channels.Set(&Channel{Uuid: "stale", Callstate: "ACTIVE"})
	channels.Replace([]Channel{{Uuid: "current"}})

	if channels.GetByUuid("stale") != nil || channels.GetByUuid("current") == nil {
		t.Fatal("replace did not install an exact snapshot")
	}
	if total, answered := channels.Counters(); total != 1 || answered != 0 {
		t.Fatalf("unexpected counters after replace: total=%d answered=%d", total, answered)
	}
}
