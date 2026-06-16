/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goproxy/internal/httpclient"
	"net/http"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var  headers []string 
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
        []string{
            "q=tom",
            "page=1",
        },
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
    nil,
)
if err != nil {
    fmt.Println("Failed to create request:", err)
    return
}
// apply heaader 
httpclient.ApplyHeaders(req,  headers,)
		fmt.Println(req.Header)
		// send request 
		response, err := httpclient.SendRequest(req, 10)
		if  err!=nil{
			fmt.Println("error while sending the request",err)
			return 
		}
		fmt.Println(response)
		body, err := httpclient.ReadResponseBody(response)
if err != nil {
    fmt.Println(err)
    return
}
defer response.Body.Close()
fmt.Println("Body Size:", len(body))
  prettyJSON, err := httpclient.PrettyPrintJSON(body)
if err != nil {
    fmt.Println("unable to pretty print json")
    return
}
fmt.Println(string(prettyJSON))
fmt.Println("Headers:", headers)
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
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
