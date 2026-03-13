package todo

import "time"

type Task struct {
	Title       string
	Description string
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewTask(Title string, Description string) Task {
	return Task{
		Title:       Title,
		Description: Description,
	}
}

func (t *Task) Done() {
	updatedAt := time.Now()

	t.IsDone = true
	t.UpdatedAt = &updatedAt
}
