package db

import (
	"errors"
	"strings"
	"testing"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"gorm.io/gorm"
)

func TestConnectFromConfigSelectsPostgresInProduction(t *testing.T) {
	originalPostgres := openPostgres
	originalSQLite := openSQLite
	t.Cleanup(func() {
		openPostgres = originalPostgres
		openSQLite = originalSQLite
	})

	postgresCalled := false
	sqliteCalled := false
	openPostgres = func(dsn string) (*gorm.DB, error) {
		postgresCalled = true
		return &gorm.DB{}, nil
	}
	openSQLite = func(path string) (*gorm.DB, error) {
		sqliteCalled = true
		return &gorm.DB{}, nil
	}

	_, err := ConnectFromConfig(config.Config{
		Environment: "production",
		DatabaseDSN: "postgres://todo:todo@localhost:5432/todos?sslmode=disable",
		SQLitePath:  "ignored.db",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !postgresCalled {
		t.Fatalf("expected postgres connector to be used")
	}
	if sqliteCalled {
		t.Fatalf("did not expect sqlite connector to be used")
	}
}

func TestConnectFromConfigSelectsSQLiteForNonProduction(t *testing.T) {
	cfg := config.Config{
		Environment: "dev",
		SQLitePath:  t.TempDir() + "/todos.db",
	}

	database, err := ConnectFromConfig(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		t.Fatalf("unexpected db conversion error: %v", err)
	}
	t.Cleanup(func() {
		_ = sqlDB.Close()
	})
}

func TestConnectFromConfigProductionDoesNotFallbackToSQLite(t *testing.T) {
	originalPostgres := openPostgres
	originalSQLite := openSQLite
	t.Cleanup(func() {
		openPostgres = originalPostgres
		openSQLite = originalSQLite
	})

	sqliteCalled := false
	openPostgres = func(dsn string) (*gorm.DB, error) {
		return nil, gorm.ErrInvalidDB
	}
	openSQLite = func(path string) (*gorm.DB, error) {
		sqliteCalled = true
		return &gorm.DB{}, nil
	}

	_, err := ConnectFromConfig(config.Config{
		Environment: "production",
		DatabaseDSN: "postgres://todo:todo@localhost:5432/todos?sslmode=disable",
		SQLitePath:  t.TempDir() + "/todos.db",
	})
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "postgres connector failed") {
		t.Fatalf("expected wrapped postgres error, got %v", err)
	}
	if sqliteCalled {
		t.Fatalf("expected no sqlite fallback in production")
	}
}

func TestConnectFromConfigRequiresProductionDSN(t *testing.T) {
	_, err := ConnectFromConfig(config.Config{
		Environment: "production",
		DatabaseDSN: "  ",
	})
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "database dsn is required") {
		t.Fatalf("expected dsn-required error, got %v", err)
	}
}

func TestConnectFromConfigWrapsSQLiteOpenFailures(t *testing.T) {
	originalSQLite := openSQLite
	t.Cleanup(func() {
		openSQLite = originalSQLite
	})
	openSQLite = func(path string) (*gorm.DB, error) {
		return nil, errors.New("boom")
	}

	_, err := ConnectFromConfig(config.Config{
		Environment: "local",
		SQLitePath:  t.TempDir() + "/todos.db",
	})
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "sqlite connector failed") {
		t.Fatalf("expected wrapped sqlite error, got %v", err)
	}
}
