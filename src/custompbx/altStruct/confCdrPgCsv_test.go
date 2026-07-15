package altStruct

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestConfigCdrPgCsvSchemaOmitsEmptyOptionalAttributes(t *testing.T) {
	encoded, err := xml.Marshal(ConfigCdrPgCsvSchema{Var: "uuid"})
	if err != nil {
		t.Fatalf("marshal schema field: %v", err)
	}

	got := string(encoded)
	if !strings.Contains(got, `var="uuid"`) {
		t.Fatalf("required var attribute missing: %s", got)
	}
	if strings.Contains(got, `column=`) {
		t.Fatalf("empty column attribute must be omitted: %s", got)
	}
	if strings.Contains(got, `quote=`) {
		t.Fatalf("empty quote attribute must be omitted: %s", got)
	}
}

func TestConfigCdrPgCsvSchemaKeepsConfiguredOptionalAttributes(t *testing.T) {
	encoded, err := xml.Marshal(ConfigCdrPgCsvSchema{
		Var:    "duration",
		Column: "duration_seconds",
		Quote:  "false",
	})
	if err != nil {
		t.Fatalf("marshal schema field: %v", err)
	}

	got := string(encoded)
	for _, attribute := range []string{`var="duration"`, `column="duration_seconds"`, `quote="false"`} {
		if !strings.Contains(got, attribute) {
			t.Fatalf("configured attribute %s missing: %s", attribute, got)
		}
	}
}
