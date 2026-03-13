package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func GetList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExist
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *List) GetTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task)
	for k, v := range tmp {
		tmp[k] = v
	}

	return tmp
}

func (l *List) DoneTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Done()

	l.tasks[title] = task

	return task, nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[title]; !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
