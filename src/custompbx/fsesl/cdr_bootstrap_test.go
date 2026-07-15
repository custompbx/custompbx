package fsesl

import (
	"custompbx/cfg"
	"testing"
)

func TestDemoCDRBootstrapValue(t *testing.T) {
	oldDB := cfg.CustomPbx.Db
	t.Cleanup(func() { cfg.CustomPbx.Db = oldDB })
	cfg.CustomPbx.Db.Host = "postgres-host"
	cfg.CustomPbx.Db.Port = 5432
	cfg.CustomPbx.Db.Name = "custompbx"
	cfg.CustomPbx.Db.User = "custompbx"
	cfg.CustomPbx.Db.Pass = "demo-password"

	t.Run("upgrades legacy sample in demo mode", func(t *testing.T) {
		t.Setenv(demoCDRBootstrapEnv, "true")
		got := demoCDRBootstrapValue("db-info", legacyDemoCDRDBInfo)
		want := "host=postgres-host port=5432 dbname=custompbx user=custompbx password=demo-password connect_timeout=10"
		if got != want {
			t.Fatalf("unexpected bootstrap value: %q", got)
		}
	})

	t.Run("preserves operator configuration", func(t *testing.T) {
		t.Setenv(demoCDRBootstrapEnv, "true")
		const custom = "host=cdr.example.test dbname=customer_cdr"
		if got := demoCDRBootstrapValue("db-info", custom); got != custom {
			t.Fatalf("operator value changed: %q", got)
		}
	})

	t.Run("does nothing outside demo mode", func(t *testing.T) {
		t.Setenv(demoCDRBootstrapEnv, "false")
		if got := demoCDRBootstrapValue("db-info", legacyDemoCDRDBInfo); got != legacyDemoCDRDBInfo {
			t.Fatalf("legacy value changed outside demo mode: %q", got)
		}
	})
}
