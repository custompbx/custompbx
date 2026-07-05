package db

import "testing"

func TestHashToken(t *testing.T) {
	a := HashToken("one")
	b := HashToken("two")
	if len(a) != 64 || a == b || a != HashToken("one") {
		t.Fatalf("unexpected hashes: %q %q", a, b)
	}
	if a == "one" {
		t.Fatal("token stored without hashing")
	}
}
