package logsafe

import (
	"strings"
	"testing"
)

func TestRedactMasksKnownSecretFields(t *testing.T) {
	input := map[string]interface{}{
		"token":           "abc123",
		"password":        "pw",
		"SecretAccessKey": "secret",
		"Authorization":   "Digest username=\"1000\"",
		"safe":            "visible",
	}

	got := Redact(input)
	for _, secret := range []string{"abc123", "pw", "secret", "Digest username"} {
		if strings.Contains(got, secret) {
			t.Fatalf("secret %q leaked in %s", secret, got)
		}
	}
	if !strings.Contains(got, "visible") {
		t.Fatalf("safe value was unexpectedly removed: %s", got)
	}
}

func TestRedactMasksConnectionStrings(t *testing.T) {
	got := RedactString("pgsql://host=db user=custom password=s3cr3t application_name=CustomPBX postgres://user:pw@example/db")
	for _, secret := range []string{"s3cr3t", ":pw@"} {
		if strings.Contains(got, secret) {
			t.Fatalf("secret %q leaked in %s", secret, got)
		}
	}
}
