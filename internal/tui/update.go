package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:
	switch msg.String() {

	case "up":
		m.selected = (m.selected - 1 + len(menuItems)) % len(menuItems)

	case "down":
		m.selected = (m.selected + 1) % len(menuItems)

	case "q", "ctrl+c":
		return m, tea.Quit

	case "enter":
		m.screen = Screen(m.selected)
		 
	}
	}
	   return m, nil
}