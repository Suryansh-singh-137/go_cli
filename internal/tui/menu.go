package tui

func (m Model) renderMainMenu() string {
	var output string

	output += "GoProxy\n"
	output += "HTTP Client & Forward Proxy\n\n"

	for i, item := range menuItems {
		if i == m.selected {
			output += "> " + item + "\n"
		} else {
			output += "  " + item + "\n"
		}
	}

	return output
}