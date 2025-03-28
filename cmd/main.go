package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"task-tui/internal/storage"
	"task-tui/internal/task"
	"task-tui/internal/ui"
)

func main() {
	// Determine storage path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not find home directory:", err)
		os.Exit(1)
	}

	storagePath := filepath.Join(homeDir, ".task-manager", "tasks.json")

	// Initialize storage and task manager
	jsonStorage := storage.NewJSONStorage(storagePath)
	taskManager := task.NewManager(jsonStorage)

	// Initialize bubbletea program
	p := tea.NewProgram(ui.InitialModel(taskManager))
	if err := p.Start(); err != nil {
		fmt.Println("Error rinning program:", err)
		os.Exit(1)
	}
}
