package ui

import (
	"task-tui/internal/model"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Base Styles
	baseStyle = lipgloss.NewStyle().Padding(1, 2)

	// Task list styles
	taskListStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63"))

	// Priority Styles
	lowPriorityStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("green"))

	mediumPriorityStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("yellow"))

	highPriorityStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("red")).
				Bold(true)

	// Status Styles
	pendingStatusStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240"))

	inProgressStatusStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("39"))

	completedStatusStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("42")).
				Strikethrough(true)

	// Input Styles
	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Background(lipgloss.Color("237")).
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	inputFocusedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Background(lipgloss.Color("237")).
				Padding(0, 1).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("51"))

	// Help Styles
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)
)

func StyleForPriority(priority model.Priority) lipgloss.Style {
	switch priority {
	case model.PriorityLow:
		return lowPriorityStyle
	case model.PriorityMedium:
		return mediumPriorityStyle
	case model.PriorityHigh:
		return highPriorityStyle
	default:
		return lipgloss.NewStyle()
	}
}

func StyleForStatus(status model.Status) lipgloss.Style {
	switch status {
	case model.StatusPending:
		return pendingStatusStyle
	case model.StatusInProgress:
		return inProgressStatusStyle
	case model.StatusCompleted:
		return completedStatusStyle
	default:
		return lipgloss.NewStyle()
	}
}
