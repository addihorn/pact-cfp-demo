package repository

import (
	"context"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/domain"
	"gorm.io/gorm"
)

type TodoRepository interface {
	List(ctx context.Context, completed *bool) ([]domain.Todo, error)
	Create(ctx context.Context, todo *domain.Todo) error
	Update(ctx context.Context, todo *domain.Todo) error
	GetByID(ctx context.Context, id uint) (*domain.Todo, error)
	Delete(ctx context.Context, id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) List(ctx context.Context, completed *bool) ([]domain.Todo, error) {
	var todos []domain.Todo
	q := r.db.WithContext(ctx).Order("created_at desc")
	if completed != nil {
		q = q.Where("completed = ?", *completed)
	}
	err := q.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *todoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *todoRepository) GetByID(ctx context.Context, id uint) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Todo{}, id).Error
}
