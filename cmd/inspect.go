/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	  "goproxy/internal/httpclient"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
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
},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
