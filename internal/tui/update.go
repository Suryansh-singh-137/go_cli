package tui

import (
	"goproxy/internal/httpclient"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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

	case "up":
		m.http.selectedField = HTTPField((int(m.http.selectedField) - 1 + len(httpclientitems)) % len(httpclientitems))
		return m, nil

	case "down":
		m.http.selectedField = HTTPField((int(m.http.selectedField) + 1) % len(httpclientitems))
		return m, nil

	case "enter":  // ← NEW: handle Enter on Send button
		if m.http.selectedField == SendField {
			// Extract URL from textinput
			finalURL := m.http.urlInput.Value()

			// Build the request (for now, just GET with the URL)
			req, err := http.NewRequest(m.http.method, finalURL, nil)
			if err != nil {
				// TODO: show error message in UI
				return m, nil
			}

			// Send it
			startTime := time.Now()
			response, err := httpclient.SendRequest(req, m.http.timeout)
			responseTime := time.Since(startTime)
			if err != nil {
				// TODO: show error message in UI
				return m, nil
			}
			defer response.Body.Close()

			// Read response body
			responseBody, err := httpclient.ReadResponseBody(response)
			if err != nil {
				return m, nil
			}

			// Try to pretty-print as JSON
			prettyJSON, err := httpclient.PrettyPrintJSON(responseBody)
			if err != nil {
				prettyJSON = responseBody
			}

			// Build summary
			summary := httpclient.Summary{
				Method:         m.http.method,
				ResponseTime:   responseTime,
				Status:         response.Status,
				ResponseLength: len(responseBody),
				ContentType:    response.Header.Get("Content-Type"),
				PrettyJSON:     prettyJSON,
				PreviewData:    string(responseBody[:min(len(responseBody), 300)]),
			}

			// Store in model
			m.http.response = &summary
			m.http.responseBody = responseBody

			// Show response screen (we'll create this next)
			m.screen = HTTPResponseScreen

			return m, nil
		}

	case "esc":
		m.screen = MainMenu
		return m, nil
	}

	// If URL field is selected, let textinput handle keystrokes
	if m.http.selectedField == URLField {
		updated, cmd := m.http.urlInput.Update(msg)
		m.http.urlInput = updated
		return m, cmd
	}
case ProxyScreen:
   switch msg.String() {

case "esc":
    m.screen = MainMenu
}
case HTTPResponseScreen:
	switch msg.String() {
	case "esc":
		m.screen = HTTPClientScreen
		return m, nil
	}
}
	}
	   return m, nil
}
