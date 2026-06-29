package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:
switch m.screen {// check which screen i hree  , then  handle the key press accoring to the rendered  screen  

case MainMenu:
	switch msg.String(){

	
  case "up":
		m.selected = (m.selected - 1 + len(menuItems)) % len(menuItems)

	case "down":
		m.selected = (m.selected + 1) % len(menuItems)

	case "q", "ctrl+c":
		return m, tea.Quit

	case "enter":
	switch m.selected {
case 0:
   m.screen =  HTTPClientScreen
case 1:
    m.screen =  ProxyScreen
case 2:
    return  m , tea.Quit
}
	}

case HTTPClientScreen:
   switch msg.String() {

case "esc":
    m.screen = MainMenu
}

case ProxyScreen:
   switch msg.String() {

case "esc":
    m.screen = MainMenu
}
}
	}
	   return m, nil
}
