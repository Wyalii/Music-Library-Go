package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	message string
	count   int
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) View() string {
	return fmt.Sprint(m.message, m.count)
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
func main() {
	p := tea.NewProgram(model{message: "this is the internal state: ", count: 0})
	if _, teaErr := p.Run(); teaErr != nil {
		log.Panic(teaErr)
	}
}
