package ui

import (
	"fmt"

	"task-tui/internal/model"
	"task-tui/internal/task"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tasks         []*model.Task
	taskManager   *task.Manager
	selectedIndex int
	view          string
	textInput     textinput.Model
	err           error
}

func InitialModel(taskManager *task.Manager) Model {
	ti := textinput.New()
	ti.Placeholder = "Enter a new task: "
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	return Model{
		tasks:       taskManager.GetTasks(),
		taskManager: taskManager,
		view:        "list",
		textInput:   ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case "j":
			if m.selectedIndex < len(m.tasks)-1 {
				m.selectedIndex++
			}
		case "a":
			if m.view != "add" {
				m.view = "add"
				m.textInput.Reset()
				m.textInput.Focus()
				return m, nil
			}
		case "esc":
			m.view = "list"
			m.textInput.Blur()
			m.textInput.Reset()
		case "enter":
			if m.view == "add" {
				if m.textInput.Value() != "" {
					m.taskManager.AddTask(m.textInput.Value())
					m.tasks = m.taskManager.GetTasks()
					m.textInput.Reset()
					m.view = "list"
				}
			}
		case "p":
			if m.view == "list" && len(m.tasks) > 0 {
				m.togglePriority()
			}
		case "s":
			if m.view == "list" && len(m.tasks) > 0 {
				m.toggleStatus()
			}
		}

	// hendle window size if needed
	case tea.WindowSizeMsg:
		m.textInput.Width = 50
	}

	if m.view == "add" {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m *Model) togglePriority() {
	if m.selectedIndex < 0 || m.selectedIndex >= len(m.tasks) {
		return
	}

	task := m.tasks[m.selectedIndex]

	// Cycle through priority states
	switch task.Priority {
	case model.PriorityLow:
		task.Priority = model.PriorityMedium
	case model.PriorityMedium:
		task.Priority = model.PriorityHigh
	case model.PriorityHigh:
		task.Priority = model.PriorityLow
	}

	// Update the task in the task manager
	m.taskManager.UpdateTask(task.ID, func(t *model.Task) {
		t.Priority = task.Priority
	})

	// Refresh tasks
	m.tasks = m.taskManager.GetTasks()
}

func (m *Model) toggleStatus() {
	if m.selectedIndex < 0 || m.selectedIndex >= len(m.tasks) {
		return
	}

	task := m.tasks[m.selectedIndex]

	// Cycle through status states
	switch task.Status {
	case model.StatusPending:
		task.Status = model.StatusInProgress
	case model.StatusInProgress:
		task.Status = model.StatusCompleted
	case model.StatusCompleted:
		task.Status = model.StatusPending
	}

	// Update the task in the task manager
	m.taskManager.UpdateTask(task.ID, func(t *model.Task) {
		t.Status = task.Status
	})

	// Refresh Task
	m.tasks = m.taskManager.GetTasks()

}

func (m Model) View() string {
	switch m.view {
	case "list":
		return m.renderTaskList()
	case "add":
		return m.renderAddTask()
	default:
		return m.renderTaskList()
	}
}

func (m Model) renderTaskList() string {
	s := "Task Manager \n\n"

	for i, task := range m.tasks {
		// Styling based on selection and task properties
		cursor := " "
		if m.selectedIndex == i {
			cursor = "→"
		}

		priorityStyle := StyleForPriority(task.Priority)
		statusStyle := StyleForStatus(task.Status)

		s += fmt.Sprintf("%s%-25s %s %s\n",
			cursor,
			priorityStyle.Render(task.Title),
			statusStyle.Render(fmt.Sprintf("%-12s", task.Status.String())),
			helpStyle.Render(fmt.Sprintf("[Priority: %s]", task.Priority.String())),
		)
	}

	s += "\n" + helpStyle.Render("Press 'a' to add task, 'q' to quit")
	return s
}

func (m Model) renderAddTask() string {
	// Use different style when focused
	inputStyleToUse := inputStyle
	if m.textInput.Focused() {
		inputStyleToUse = inputFocusedStyle
	}

	return fmt.Sprintf(
		"Add a new task:\n\n %s \n\n %s",
		inputStyleToUse.Render(m.textInput.View()),
		helpStyle.Render("Press Enter to save, 'q' to cancel"),
	)
}
