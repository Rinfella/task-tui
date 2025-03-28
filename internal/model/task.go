package model

import (
	"time"

	"github.com/google/uuid"
)

type Priority int
type Status int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

const (
	StatusPending Status = iota
	StatusInProgress
	StatusCompleted
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    Priority   `json:"priority"`
	Status      Status     `json:"status"`
	Tags        []string   `json:"tags"`
}

func NewTask(title string) *Task {
	return &Task{
		ID:        uuid.New(),
		Title:     title,
		CreatedAt: time.Now(),
		Priority:  PriorityMedium,
		Status:    StatusPending,
	}
}

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	default:
		return "Unknown"
	}
}

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "Pending"
	case StatusInProgress:
		return "In Progress"
	case StatusCompleted:
		return "Completed"
	default:
		return "Unknown"
	}
}
