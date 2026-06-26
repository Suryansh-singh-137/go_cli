package tui

import tea "github.com/charmbracelet/bubbletea"

var menuItems = []string{
	"HTTP Client",
	"Proxy Server",
	"Exit",
}

type Model struct {
	selected int
}

func New() Model {
	return Model{
		selected: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
   var output string

output += "GoProxy\n"
output += "HTTP Client & Forward Proxy\n\n"
for i, item := range menuItems {
	if i  == m.selected{
		output+= "> "+item+"\n"
	}else {
		output+="  "+item+"\n"
	}
}
return output
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up":
			m.selected = (m.selected - 1 + len(menuItems)) % len(menuItems)

		case "down":
			m.selected = (m.selected + 1) % len(menuItems)

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			// We'll implement this next.
		}
		
	}

	return m, nil
}