package tui
type HTTPField  int 

const (
    MethodField HTTPField = iota
    URLField
    HeadersField
    BodyField
    TimeoutField
    SendField
)

var httpclientitems = []string{
	"method",
	"url",
	"header",
    "body",
    "timeout",
    "send",
}

func (m Model) renderHTTPClient() string {
	var output string
	output = "HTTP Client\n\n"
	
	for i, item := range httpclientitems {
		if i == int(m.http.selectedField) {
			output += "> " + item + "\n"
			
			// ← NEW: When URL field is selected, show the textinput below it
			if HTTPField(i) == URLField {
				output += "  " + m.http.urlInput.View() + "\n"
			}
		} else {
			output += "  " + item + "\n"
		}
	}
	
	return output
}