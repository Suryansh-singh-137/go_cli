package tui

import (
	"fmt"
	"strings"
)

func (m Model) renderHTTPResponse() string {
	if m.http.response == nil {
		return "No response yet"
	}

	resp := m.http.response

	var output string
	output += "HTTP Response\n"
	output += "==================================\n\n"
	output += fmt.Sprintf("Status: %s\n", resp.Status)
	output += fmt.Sprintf("Response Time: %v\n", resp.ResponseTime)
	output += fmt.Sprintf("Content-Type: %s\n", resp.ContentType)
	output += fmt.Sprintf("Length: %d bytes\n", resp.ResponseLength)
	output += "\n---\n\n"

	// Show body (pretty-printed if JSON, otherwise preview)
	if strings.Contains(resp.ContentType, "application/json") {
		output += "Response Body (JSON):\n"
		output += string(resp.PrettyJSON)
	} else {
		output += "Response Body (Prev`iew):\n"
		output += resp.PreviewData
		if resp.ResponseLength > 300 {
			output += "\n... (truncated, total " + fmt.Sprintf("%d", resp.ResponseLength) + " bytes)"
		}
	}

	output += "\n\n---\nPress 'esc' to go back"

	return output
}