package service

import (
	"context"
	"testing"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/domain"
)

type fakeRepo struct {
	items map[uint]domain.Todo
	next  uint
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{items: map[uint]domain.Todo{}, next: 1}
}

func (f *fakeRepo) List(_ context.Context, completed *bool) ([]domain.Todo, error) {
	result := []domain.Todo{}
	for _, t := range f.items {
		if completed != nil && t.Completed != *completed {
			continue
		}
		result = append(result, t)
	}
	return result, nil
}

func (f *fakeRepo) Create(_ context.Context, todo *domain.Todo) error {
	todo.ID = f.next
	f.next++
	f.items[todo.ID] = *todo
	return nil
}

func (f *fakeRepo) Update(_ context.Context, todo *domain.Todo) error {
	f.items[todo.ID] = *todo
	return nil
}

func (f *fakeRepo) GetByID(_ context.Context, id uint) (*domain.Todo, error) {
	t, ok := f.items[id]
	if !ok {
		return nil, ErrTodoNotFound
	}
	return &t, nil
}

func (f *fakeRepo) Delete(_ context.Context, id uint) error {
	delete(f.items, id)
	return nil
}

func TestCreateTodoRejectsBlankTitle(t *testing.T) {
	svc := NewTodoService(newFakeRepo())
	_, err := svc.CreateTodo(context.Background(), CreateTodoInput{Title: "   "})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestCreateAndUpdateTodo(t *testing.T) {
	svc := NewTodoService(newFakeRepo())
	created, err := svc.CreateTodo(context.Background(), CreateTodoInput{Title: "write docs"})
	if err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}

	done := true
	title := "write better docs"
	updated, err := svc.UpdateTodo(context.Background(), created.ID, UpdateTodoInput{Title: &title, Completed: &done})
	if err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}
	if !updated.Completed || updated.Title != title {
		t.Fatalf("update did not apply")
	}
}
