package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"task-tui/internal/model"
)

type JSONStorage struct {
	filepath string
}

func NewJSONStorage(filepath string) *JSONStorage {
	return &JSONStorage{filepath: filepath}
}

func (s *JSONStorage) Save(tasks []*model.Task) error {
	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(s.filepath), 0755); err != nil {
		return err
	}

	// Marshall tasks to JSON
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(s.filepath, data, 0644)
}

func (s *JSONStorage) Load() ([]*model.Task, error) {
	// Check if file exists
	if _, err := os.Stat(s.filepath); os.IsNotExist(err) {
		return []*model.Task{}, nil
	}

	// Read file
	data, err := os.ReadFile(s.filepath)
	if err != nil {
		return nil, err
	}

	// If file is empty, return empty slice
	if len(data) == 0 {
		return []*model.Task{}, nil
	}

	// Unmarshal JSON
	var tasks []*model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
