/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
    fmt.Println(r.Method, r.URL.Path)

    w.Write([]byte("Hello from goproxy"))
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
