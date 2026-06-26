package tui

import tea "github.com/charmbracelet/bubbletea"
type Screen int
const (
    MainMenu Screen = iota
    HTTPClientScreen
    ProxyScreen
)
var menuItems = []string{
	"HTTP Client",
	"Proxy Server",
	"Exit",
}

type Model struct {
	selected int
	screen Screen
}

func New() Model {
	return Model{
		selected: 0,
		screen:  MainMenu,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
switch m.screen {

case MainMenu:
    return m.renderMainMenu()

case HTTPClientScreen:
    return m.renderHTTPClient()

case ProxyScreen:
    return m.renderProxy()
}
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