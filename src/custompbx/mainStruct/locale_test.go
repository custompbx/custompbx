package mainStruct

import "testing"

func TestNormalizeLocale(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"en", "en"}, {"pt-BR", "pt-BR"}, {"zh-Hans", "zh-Hans"},
		{" ar ", "ar"}, {"", "en"}, {"../../secret", "en"}, {"EN", "en"},
	}
	for _, test := range tests {
		if got := NormalizeLocale(test.input); got != test.want {
			t.Fatalf("NormalizeLocale(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestLegacyLocale(t *testing.T) {
	if LegacyLocale(0) != "en" || LegacyLocale(1) != "ru" || LegacyLocale(99) != "en" {
		t.Fatal("legacy locale mapping changed")
	}
}
