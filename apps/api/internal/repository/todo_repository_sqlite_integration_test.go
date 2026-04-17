package repository

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/db"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/domain"
	"gorm.io/gorm"
)

func TestTodoRepositoryCRUDWithSQLiteConnector(t *testing.T) {
	cfg := config.Config{
		Environment: "local",
		SQLitePath:  filepath.Join(t.TempDir(), "todos.db"),
	}

	database, err := db.ConnectFromConfig(cfg)
	if err != nil {
		t.Fatalf("unexpected connect error: %v", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		t.Fatalf("unexpected db conversion error: %v", err)
	}
	t.Cleanup(func() {
		_ = sqlDB.Close()
	})

	repo := NewTodoRepository(database)
	ctx := context.Background()

	todo := &domain.Todo{
		Title:       "write integration test",
		Description: "verify sqlite connector and repository",
	}
	if err := repo.Create(ctx, todo); err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}
	if todo.ID == 0 {
		t.Fatalf("expected id to be assigned")
	}

	fetched, err := repo.GetByID(ctx, todo.ID)
	if err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}
	if fetched.Title != todo.Title {
		t.Fatalf("expected title %q, got %q", todo.Title, fetched.Title)
	}

	fetched.Completed = true
	if err := repo.Update(ctx, fetched); err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}

	completed := true
	listed, err := repo.List(ctx, &completed)
	if err != nil {
		t.Fatalf("unexpected list error: %v", err)
	}
	if len(listed) != 1 {
		t.Fatalf("expected one completed todo, got %d", len(listed))
	}

	if err := repo.Delete(ctx, todo.ID); err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}

	_, err = repo.GetByID(ctx, todo.ID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record-not-found after delete, got %v", err)
	}
}
