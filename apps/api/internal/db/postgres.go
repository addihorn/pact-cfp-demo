package db

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/domain"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector interface {
	Connect(cfg config.Config) (*gorm.DB, error)
	BackendName() string
}

type PostgresConnector struct{}

func (PostgresConnector) BackendName() string {
	return "postgres"
}

type SQLiteConnector struct{}

func (SQLiteConnector) BackendName() string {
	return "sqlite"
}

var openPostgres = func(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

var openSQLite = func(path string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), &gorm.Config{})
}

func ConnectFromConfig(cfg config.Config) (*gorm.DB, error) {
	connector := selectConnector(cfg)
	database, err := connector.Connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s connector failed: %w", connector.BackendName(), err)
	}
	return database, nil
}

func selectConnector(cfg config.Config) Connector {
	if strings.EqualFold(strings.TrimSpace(cfg.Environment), "production") {
		return PostgresConnector{}
	}
	return SQLiteConnector{}
}

func (PostgresConnector) Connect(cfg config.Config) (*gorm.DB, error) {
	dsn := strings.TrimSpace(cfg.DatabaseDSN)
	if dsn == "" {
		return nil, fmt.Errorf("database dsn is required in production")
	}
	return openPostgres(dsn)
}

func (SQLiteConnector) Connect(cfg config.Config) (*gorm.DB, error) {
	path := strings.TrimSpace(cfg.SQLitePath)
	if path == "" {
		return nil, fmt.Errorf("sqlite path is required")
	}

	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("ensure sqlite directory %q: %w", dir, err)
		}
	}

	database, err := openSQLite(path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database %q: %w", path, err)
	}

	if err := database.AutoMigrate(&domain.Todo{}); err != nil {
		return nil, fmt.Errorf("auto-migrate sqlite schema: %w", err)
	}

	return database, nil
}
