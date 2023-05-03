package domain

import (
	"time"
)

type Task struct {
	ID        string
	Title     string
	Content   string
	Done      bool
	CreatedBy string
	CreatedAt time.Time
}
