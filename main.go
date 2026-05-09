package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

// Defining applications state as a struct
type model struct {
	counter int
}

// Defining our applications initial state
func initialModel() model {
	return model{counter: 0}
}

// Any inital command i/o, network etc. i want to run before the update method
func (m model) Init() tea.Cmd {
	return nil
}

// The Update method checks the type of the message and updates the model accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			m.counter++
			return m, nil

		case "down", "j":
			for m.counter > 0 {
				m.counter--
				return m, nil
			}
		}
	}

	return m, nil
}

// Rerenders the UI every time the model changes
func (m model) View() tea.View {
	p := fmt.Sprintf("Your count value is: %d\n", m.counter)
	p += "Press up or down to change the counter, press q/ctrl+c to quit"

	return tea.NewView(p) // Wrap the string in tea.NewView
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
