package tui

import (
	"goproxy/internal/httpclient"

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
	urlInput      textinput.Model
	response      *httpclient.Summary  // ← NEW: store the response
	responseBody  []byte               // ← NEW: raw response bytes
}
type Screen int
const (
    MainMenu Screen = iota
    HTTPClientScreen
    ProxyScreen
		HTTPResponseScreen  
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
	urlInput := textinput.New()
	urlInput.Placeholder = "https://example.com"
	urlInput.CharLimit = 2048
	urlInput.Width = 50
	urlInput.Focus()

	return Model{
		selected: 0,
		screen:   MainMenu,
		http: HTTPClientState{
			selectedField: MethodField,
			method:        "GET",
			url:           "",
			headers:       []string{},
			body:          "",
			timeout:       30,
			urlInput:      urlInput,
			response:      nil,  // ← start with no response
			responseBody:  nil,
		},
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
