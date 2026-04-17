package service

import (
	"context"
	"errors"
	"strings"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/domain"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/repository"
	"gorm.io/gorm"
)

var ErrInvalidTitle = errors.New("title must not be empty")
var ErrTodoNotFound = errors.New("todo not found")

type CreateTodoInput struct {
	Title       string
	Description string
}

type UpdateTodoInput struct {
	Title       *string
	Description *string
	Completed   *bool
}

type TodoService interface {
	ListTodos(ctx context.Context, completed *bool) ([]domain.Todo, error)
	CreateTodo(ctx context.Context, input CreateTodoInput) (*domain.Todo, error)
	UpdateTodo(ctx context.Context, id uint, input UpdateTodoInput) (*domain.Todo, error)
	DeleteTodo(ctx context.Context, id uint) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) ListTodos(ctx context.Context, completed *bool) ([]domain.Todo, error) {
	return s.repo.List(ctx, completed)
}

func (s *todoService) CreateTodo(ctx context.Context, input CreateTodoInput) (*domain.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return nil, ErrInvalidTitle
	}

	todo := &domain.Todo{
		Title:       title,
		Description: strings.TrimSpace(input.Description),
		Completed:   false,
	}
	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id uint, input UpdateTodoInput) (*domain.Todo, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}

	if input.Title != nil {
		trimmed := strings.TrimSpace(*input.Title)
		if trimmed == "" {
			return nil, ErrInvalidTitle
		}
		todo.Title = trimmed
	}
	if input.Description != nil {
		todo.Description = strings.TrimSpace(*input.Description)
	}
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id uint) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTodoNotFound
		}
		return err
	}
	return s.repo.Delete(ctx, id)
}
