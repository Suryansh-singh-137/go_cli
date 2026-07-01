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
    "send"
}

func (m Model) renderHTTPClient() string {
    var  output string;
    output = "HTTP Client"
    for i,item := range  httpclientitems{
        	if i == m.HTTPField {
			output += "> " + httpclientitems + "\n"
		} else {
			output += "  " + httpclientitems + "\n"
		}

    }

}