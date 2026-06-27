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
