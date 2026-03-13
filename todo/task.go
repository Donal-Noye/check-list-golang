package todo

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewTask(Title string, Description string) Task {
	return Task{
		ID:          uuid.New(),
		Title:       Title,
		Description: Description,
	}
}

func (t *Task) Done() {
	updatedAt := time.Now()

	t.IsDone = true
	t.UpdatedAt = &updatedAt
}
