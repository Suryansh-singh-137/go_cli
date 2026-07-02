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
type HTTPClientState struct {
    selectedField HTTPField

    method  string
    url     string
    headers []string
    body    string
    timeout int
}
var httpclientitems = []string{
	"method",
	"url",
	"header",
    "body",
    "timeout",
    "send",
}

func (m Model) renderHTTPClient() string {
	output := "HTTP Client\n\n"
	for i, item := range httpclientitems {
		if i == int(m.http.selectedField) {
			output += "> " + item + "\n"
		} else {
			output += "  " + item + "\n"
		}
	}
	return output
}