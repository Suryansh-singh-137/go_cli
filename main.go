package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)
// printing header  
func printHeaders(headers http.Header) {
	for key, value := range headers {
		fmt.Println(key, ":", value)
	}
}
// save  responses  
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
func printSummary(   endTime time.Duration,
    status string,
    responseLength int,
    contentType string,
    prettyJSON []byte,
    previewData string ){

fmt.Println("==================================== REQUEST SUMMARY ====================================")

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
func main() {
	// ====================
// Argument Validation
// ====================
// Expected usage:
// go run main.go <URL> [filename]
if len(os.Args) < 2 {
	fmt.Println("Usage: go run main.go <URL> <filename>")
	return
}

userURL := os.Args[1]
fmt.Printf("FETCHING: %s\n", userURL)

// ====================
// Build HTTP Request
// ====================
// Create a custom HTTP request instead of using http.Get()
// so we can modify headers, methods, body, etc.
start_time := time.Now()

req, err := http.NewRequest("GET", userURL, nil)
if err != nil {
	fmt.Println("unable to create a new request")
	return
}

// Add custom headers to the request
req.Header.Set(
	"User-Agent",
	"goproxy/1.0",
)

fmt.Println("Request Headers:")
fmt.Println(req.Header)

// ====================
// Send Request
// ====================
// The client is responsible for sending requests
// and receiving responses.
client := http.Client{}

response, err := client.Do(req)
if err != nil {
	fmt.Println("unable to get a response back from the server")
	return
}

end_time := time.Since(start_time)

defer response.Body.Close()

// ====================
// Read Response Body
// ====================
// Convert the response stream into bytes.
body, err := io.ReadAll(response.Body)
if err != nil {
	fmt.Printf("Error Getting the Response Body: %v\n", err)
}

len_of_res := len(body)

// ====================
// Create Response Preview
// ====================
// Only show first 300 characters for non-JSON responses.
previewLen := 300

if len(body) < previewLen {
	previewLen = len(body)
}

res_data := string(body[:previewLen])

// ====================
// Response Metadata
// ====================
res_type := response.Header.Get("Content-Type")

// ====================
// Pretty Print JSON
// ====================
// If response is JSON:
// 1. Parse JSON
// 2. Reformat with indentation
var prettyJSON []byte


if strings.Contains(res_type, "application/json") {
prettyJSON, err = prettyPrintJSON(body)

	if err != nil {
		fmt.Println("unable to pretty print json")
	}
}

// ====================
// ] File Saving
// ====================
// If user supplied a filename,
// save the response to disk.
var dataToSave []byte

if len(os.Args) > 2 {
filename:= os.Args[2]
if strings.Contains(res_type, "application/json") {
		dataToSave = prettyJSON
} else {
		dataToSave = body
	}
saveResponse(dataToSave,filename)
}
// ====================
// Display Results
// ====================
res_header := response.Header

printSummary(
    end_time,
    response.Status,
    len_of_res,
    res_type,
    prettyJSON,
    res_data,
)
// ====================
// Display Response Headers
// ====================
fmt.Println("==================================== RESPONSE HEADER ====================================")
printHeaders(res_header);
}