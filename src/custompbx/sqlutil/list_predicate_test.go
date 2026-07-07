package sqlutil

import (
	"strings"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestStringListEqualParameterized(t *testing.T) {
	predicate, err := StringListEqual("sip_call_id", []string{"a", `b'); SELECT pg_sleep(10); --`})
	if err != nil {
		t.Fatal(err)
	}
	query, args, err := squirrel.Select("*").
		From("hep_packets").
		Where(predicate).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(query, "pg_sleep") || strings.Contains(query, "SELECT pg_sleep") {
		t.Fatalf("payload was interpolated into query: %s", query)
	}
	if len(args) != 2 || args[1] != `b'); SELECT pg_sleep(10); --` {
		t.Fatalf("unexpected args: %#v", args)
	}
}

func TestStringListEqualRejectsMalformedColumn(t *testing.T) {
	if _, err := StringListEqual("sip_call_id); DROP TABLE hep_packets; --", []string{"a"}); err == nil {
		t.Fatal("malformed identifier accepted")
	}
}

func TestStringListEqualEmptyListMatchesNothing(t *testing.T) {
	predicate, err := StringListEqual("sip_call_id", nil)
	if err != nil {
		t.Fatal(err)
	}
	query, args, err := squirrel.Select("*").
		From("hep_packets").
		Where(predicate).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(query, "1 = 0") {
		t.Fatalf("empty list should match nothing: %s", query)
	}
	if len(args) != 0 {
		t.Fatalf("unexpected args: %#v", args)
	}
}
