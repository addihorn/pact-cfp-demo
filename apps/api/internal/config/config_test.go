package config

import "testing"

func TestLoadDefaultsToSQLiteInLocal(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("SQLITE_PATH", "")

	cfg := Load()

	if cfg.Environment != "local" {
		t.Fatalf("expected default environment local, got %q", cfg.Environment)
	}
	if cfg.SQLitePath != "apps/api/data/todos.db" {
		t.Fatalf("expected default sqlite path, got %q", cfg.SQLitePath)
	}
}

func TestLoadProductionEnvironment(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	cfg := Load()

	if cfg.Environment != "production" {
		t.Fatalf("expected production environment, got %q", cfg.Environment)
	}
}

func TestLoadNonProductionEnvironmentDefaultsToSQLiteSelection(t *testing.T) {
	t.Setenv("APP_ENV", "dev")

	cfg := Load()

	if cfg.Environment != "dev" {
		t.Fatalf("expected dev environment, got %q", cfg.Environment)
	}
	if cfg.SQLitePath == "" {
		t.Fatalf("expected sqlite path to be set for non-production environment")
	}
}

func TestLoadSQLitePathOverride(t *testing.T) {
	t.Setenv("SQLITE_PATH", "tmp/custom.db")

	cfg := Load()

	if cfg.SQLitePath != "tmp/custom.db" {
		t.Fatalf("expected overridden sqlite path, got %q", cfg.SQLitePath)
	}
}
