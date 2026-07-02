package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type HTTPClientState struct {
	selectedField HTTPField
	method        string
	url           string
	headers       []string
	body          string
	timeout       int
	urlInput      textinput.Model  // ← add this
}

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
	http  HTTPClientState
}
func New() Model {
	// Create a fresh textinput
	urlInput := textinput.New()
	urlInput.Placeholder = "https://example.com"
	urlInput.CharLimit = 2048
	urlInput.Width = 50
	urlInput.Focus()  // ← ADD THIS LINE — enables the textinput to accept input

	return Model{
		selected: 0,
		screen:   MainMenu,
		http: HTTPClientState{
			selectedField: MethodField,
			urlInput:      urlInput,
		},
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
