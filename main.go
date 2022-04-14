package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := NewModel("Hello nerds")
	p := tea.NewProgram(m)
	p.EnterAltScreen()
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
