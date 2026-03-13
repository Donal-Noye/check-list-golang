package todo

import (
	"sync"

	"github.com/google/uuid"
)

type List struct {
	tasks map[uuid.UUID]Task
	mtx   sync.RWMutex
}

func GetList() *List {
	return &List{
		tasks: make(map[uuid.UUID]Task),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[task.ID]; ok {
		return ErrTaskAlreadyExist
	}

	l.tasks[task.ID] = task

	return nil
}

func (l *List) GetTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task)
	for k, v := range l.tasks {
		tmp[k.String()] = v
	}

	return tmp
}

func (l *List) DoneTask(id uuid.UUID) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Done()

	l.tasks[id] = task

	return task, nil
}

func (l *List) DeleteTask(id uuid.UUID) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, id)

	return nil
}
