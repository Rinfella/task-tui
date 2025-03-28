package task

import (
	"errors"

	"github.com/google/uuid"
	"task-tui/internal/model"
	"task-tui/internal/storage"
)

type Manager struct {
	tasks   []*model.Task
	storage *storage.JSONStorage
}

func NewManager(storage *storage.JSONStorage) *Manager {
	tasks, _ := storage.Load()
	return &Manager{
		tasks:   tasks,
		storage: storage,
	}
}

func (m *Manager) AddTask(title string) *model.Task {
	task := model.NewTask(title)
	m.tasks = append(m.tasks, task)
	m.save()
	return task
}

func (m *Manager) UpdateTask(id uuid.UUID, updateFn func(*model.Task)) error {
	for _, task := range m.tasks {
		if task.ID == id {
			updateFn(task)
			m.save()
			return nil
		}
	}
	return errors.New("Task not found")
}

func (m *Manager) DeleteTask(id uuid.UUID) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			m.save()
			return nil
		}
	}
	return errors.New("Task not found")
}

func (m *Manager) GetTasks() []*model.Task {
	return m.tasks
}

func (m *Manager) FilterTask(
	status *model.Status,
	priority *model.Priority,
	tag string,
) []*model.Task {
	var filtered []*model.Task

	for _, task := range m.tasks {
		if status != nil && task.Status != *status {
			continue
		}

		if priority != nil && task.Priority != *priority {
			continue
		}

		if tag != "" {
			hasTag := false
			for _, t := range task.Tags {
				if t == tag {
					hasTag = true
					break
				}
			}
			if !hasTag {
				continue
			}
		}

		filtered = append(filtered, task)
	}

	return filtered
}

func (m *Manager) save() {
	m.storage.Save(m.tasks)
}
