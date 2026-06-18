/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}
var startCmd = &cobra.Command{
    Use: "start",
		Run : func(cmd *cobra.Command, args []string) {
			  fmt.Println("Proxy listening on :8080")
http.HandleFunc("/",func(w http.ResponseWriter ,  r *http.Request){
 fmt.Println("Received request:")
fmt.Println("================================")
fmt.Println("Incoming Request")
fmt.Println("Method:", r.Method)
fmt.Println("Path:", r.URL.Path)
fmt.Println("Host:", r.Host)
fmt.Println("Remote Addr:", r.RemoteAddr)
fmt.Println("Query Params:")
// printing req params  
for key, values := range r.URL.Query() {
    fmt.Printf("%s = %v\n", key, values)
}
// printing req body 
body,err:= io.ReadAll(r.Body)
if err!=nil{
	fmt.Println("unable to read request bosy",err)
}
fmt.Println(string(body))

for key,values := range r.Header{
	fmt.Printf("%s: %v\n", key, values)
}
// creaeting a reuwst  to  fowraed it to my serve using  my sereer as proxy
user_requested_url:= "https://httpbingo.org"+   r.URL.RequestURI()
req, err := http.NewRequest(
    r.Method,
   user_requested_url,
     bytes.NewReader(body),
)
if err != nil {
    fmt.Println(err)
    return
}
// setiing headers  fo rour req  
for key,values := range  r.Header{
	for _,value := range values{
		req.Header.Set(key,value)
	}
}
client := &http.Client{}
response, err := client.Do(req)

if err != nil {
    fmt.Println(err)
    return
}
defer response.Body.Close()
responseBody, err := io.ReadAll(response.Body)
if err != nil {
    fmt.Println(err)
    return
}
_, err = w.Write(responseBody)

if err != nil {
    fmt.Println(err)
}
})
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println(err)
    }
		
	
		},
}
func init() {
	rootCmd.AddCommand(proxyCmd)
 proxyCmd.AddCommand(startCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proxyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proxyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
