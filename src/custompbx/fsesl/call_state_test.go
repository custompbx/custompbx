package fsesl

import "testing"

func TestParseFreeSWITCHEpoch(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  int64
	}{
		{name: "event microseconds", value: "1783737004000000", want: 1783737004},
		{name: "snapshot seconds", value: "1783737004", want: 1783737004},
		{name: "invalid", value: "", want: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := parseFreeSWITCHEpoch(test.value); got != test.want {
				t.Fatalf("parseFreeSWITCHEpoch(%q)=%d, want %d", test.value, got, test.want)
			}
		})
	}
}

func TestNormalizedFreeSWITCHEpochAppliesObservedClockOffset(t *testing.T) {
	previous := freeSwitchClockOffsetSeconds
	defer func() { freeSwitchClockOffsetSeconds = previous }()
	freeSwitchClockOffsetSeconds = 209934

	if got := normalizedFreeSWITCHEpoch("1783737004000000"); got != 1783946938 {
		t.Fatalf("normalized epoch=%d, want 1783946938", got)
	}
}
