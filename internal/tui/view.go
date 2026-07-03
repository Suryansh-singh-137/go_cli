package tui

func (m Model) View() string {
	switch m.screen {
	case MainMenu:
		return m.renderMainMenu()
	case HTTPClientScreen:
		return m.renderHTTPClient()
	case HTTPResponseScreen:
		return m.renderHTTPResponse() // ← NEW
	case ProxyScreen:
		return m.renderProxy()
	}

	return "Unknown Screen"
}