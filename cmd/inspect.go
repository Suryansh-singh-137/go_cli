/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goproxy/internal/httpclient"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
// defining flags  for  inspect cmnd
var  headers []string 
var  timeout  int 
var query []string  
var  body string 
var  save string 
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
Run: func(cmd *cobra.Command, args []string) {

    if len(args) < 2 {
        fmt.Println("Usage: inspect <METHOD> <URL>")
        return
    }

    method := args[0]
    userURL := args[1]

    fmt.Println("Method:", method)
    fmt.Println("URL:", userURL)

    finalURL, err := httpclient.BuildURL(
        userURL,
      query,
    )

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Final URL:", finalURL)
// bilding request 
req, err := http.NewRequest(
    method,
    finalURL,
 strings.NewReader(body),
)
req.Header.Set(
    "User-Agent",
    "goproxy/1.0",
)
if err != nil {
    fmt.Println("Failed to create request:", err)
    return
}
// apply heaader 
httpclient.ApplyHeaders(req,  headers,)
		fmt.Println(req.Header)
		// send request 
		startTime := time.Now()
		response, err := httpclient.SendRequest(req, timeout)
		endTime := time.Since(startTime)
		if  err!=nil{
			fmt.Println("error while sending the request",err)
			return 
		}
		defer response.Body.Close()
		fmt.Println(response)
		responseBody, err := httpclient.ReadResponseBody(response)
		contentType := response.Header.Get("Content-Type")
if err != nil {
    fmt.Println(err)
    return
}
previewLen := 300
fmt.Println("Body Size:", len(responseBody))

if len(responseBody) < previewLen {
    previewLen = len(responseBody)
}

previewData := string(responseBody[:previewLen])
  prettyJSON, err := httpclient.PrettyPrintJSON(responseBody)
if err != nil {
    fmt.Println("unable to pretty print json")
    return
}
// building struct 
summary := httpclient.Summary{
    Method:         method,
    ResponseTime:   endTime,
    Status:         response.Status,
    ResponseLength: len(responseBody),
    ContentType:    contentType,
    PrettyJSON:     prettyJSON,
    PreviewData:    previewData,
}
//save resposne 
if save != "" {

    if strings.Contains(contentType, "application/json") {
        httpclient.SaveResponse(prettyJSON, save)
    } else {
        httpclient.SaveResponse(responseBody, save)
    }

}
httpclient.PrintSummary(summary)

},

}
func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.
  inspectCmd.Flags().StringSliceVar(
        &headers,
        "header",
        []string{},
        "custom headers",
    )
		inspectCmd.Flags().IntVar(
    &timeout,
    "timeout",
    30,
    "request timeout in seconds",
)
  inspectCmd.Flags().StringSliceVar(
        &query,
        "query",
        []string{},
        "query parameters",
    )
		  inspectCmd.Flags().StringVar(
        &body,
        "body",
       "",
        "request  body",
    )
		inspectCmd.Flags().StringVar(
        &save,
        "save",
       "",
        "save reponse from server in provided file",
    )
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
