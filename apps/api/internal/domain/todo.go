package domain

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt,format:'2006-01-02T15:04:05Z'"`
	UpdatedAt   time.Time `json:"updatedAt,format:'2006-01-02T15:04:05Z'"`
}
