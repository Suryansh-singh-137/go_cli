package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)
func  saveResponse(dataToSave [] byte,filename string){
		err := os.WriteFile(filename, dataToSave, 0644)

	if err != nil {
		fmt.Printf("Error saving response to file: %v\n", err)
	} else {
		fmt.Printf("Response saved to %s\n", filename)
	}
}
func prettyPrintJSON(body []byte) ([]byte, error){
	var data interface{}

err := json.Unmarshal(body, &data)
if err != nil {
	return nil, err
}

prettyJSON, err := json.MarshalIndent(data, "", "  ")
if err != nil {
	return nil, err
}

return prettyJSON, nil
}
func printSummary(  method string ,  endTime time.Duration,
    status string,
    responseLength int,
    contentType string,
    prettyJSON []byte,
    previewData string ){

fmt.Println("==================================== REQUEST SUMMARY ====================================")
fmt.Printf("Method: %s\n", method)
fmt.Printf("Response Time: %v\n", endTime)
fmt.Printf("Status Code: %s\n", status)
fmt.Printf("Length: %d\n", responseLength)

if strings.Contains( contentType, "application/json") {
	fmt.Printf("Response: %s\n", string(prettyJSON))
} else {
	fmt.Printf("Response: %s\n",previewData)
}

fmt.Printf("Response Type: %s\n", contentType)
		}
func BuildURL(userURL string,queries []string) (string, error) {

	parsedURL, err := url.Parse(userURL)

	if err != nil {

		return "", err

	}

	params := parsedURL.Query()

	for _, query := range queries {

		parts := strings.SplitN(query, "=", 2)

		if len(parts) != 2 {

			fmt.Printf("Invalid query format: %sn", query)

			continue

		}

		key := strings.TrimSpace(parts[0])

		value := strings.TrimSpace(parts[1])

		params.Set(key, value)

	}

	parsedURL.RawQuery = params.Encode()

	return parsedURL.String(), nil

}
func ApplyHeaders(req *http.Request,  headers []string) {
	// parse header
	// add header
	for _, header := range headers {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid header format: %s\n", header)
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		req.Header.Set(key, value)
	}
}
// internal/httpclient/client.go

func SendRequest(req *http.Request, timeout int) (*http.Response, error) {
    client := http.Client{
        Timeout: time.Duration(timeout) * time.Second,
    }

    return client.Do(req)
}
func ReadResponseBody(
    response *http.Response,
) ([]byte, error) {

    return io.ReadAll(response.Body)
}