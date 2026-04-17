package api

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/domain"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/repository"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/service"
)

func TestTodoCRUDFlow(t *testing.T) {
	repo := &inMemoryRepo{items: map[uint]domain.Todo{}, next: 1}
	svc := service.NewTodoService(repo)
	h := NewRouter(svc, config.Config{})

	createBody := []byte(`{"title":"demo todo","description":"from integration test"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRes := httptest.NewRecorder()
	h.ServeHTTP(createRes, createReq)
	if createRes.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", createRes.Code)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	listRes := httptest.NewRecorder()
	h.ServeHTTP(listRes, listReq)
	if listRes.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", listRes.Code)
	}

	patchBody := []byte(`{"completed":true}`)
	patchReq := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/1", bytes.NewReader(patchBody))
	patchReq.Header.Set("Content-Type", "application/json")
	patchRes := httptest.NewRecorder()
	h.ServeHTTP(patchRes, patchReq)
	if patchRes.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", patchRes.Code)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/v1/todos/1", nil)
	deleteRes := httptest.NewRecorder()
	h.ServeHTTP(deleteRes, deleteReq)
	if deleteRes.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", deleteRes.Code)
	}
}

func TestHealthAndMetrics(t *testing.T) {
	repo := &inMemoryRepo{items: map[uint]domain.Todo{}, next: 1}
	svc := service.NewTodoService(repo)
	h := NewRouter(svc, config.Config{})

	for _, path := range []string{"/healthz", "/readyz", "/metrics"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		res := httptest.NewRecorder()
		h.ServeHTTP(res, req)
		if res.Code != http.StatusOK {
			t.Fatalf("path %s expected 200 got %d", path, res.Code)
		}
	}
}

type inMemoryRepo struct {
	items map[uint]domain.Todo
	next  uint
}

func (r *inMemoryRepo) List(_ context.Context, completed *bool) ([]domain.Todo, error) {
	out := []domain.Todo{}
	for _, t := range r.items {
		if completed != nil && t.Completed != *completed {
			continue
		}
		out = append(out, t)
	}
	return out, nil
}

func (r *inMemoryRepo) Create(_ context.Context, todo *domain.Todo) error {
	todo.ID = r.next
	r.next++
	r.items[todo.ID] = *todo
	return nil
}

func (r *inMemoryRepo) Update(_ context.Context, todo *domain.Todo) error {
	r.items[todo.ID] = *todo
	return nil
}

func (r *inMemoryRepo) GetByID(_ context.Context, id uint) (*domain.Todo, error) {
	t, ok := r.items[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &t, nil
}

func (r *inMemoryRepo) Delete(_ context.Context, id uint) error {
	delete(r.items, id)
	return nil
}

var _ repository.TodoRepository = (*inMemoryRepo)(nil)
