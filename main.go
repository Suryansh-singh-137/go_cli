package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	userURL := os.Args[1]
	fmt.Printf("FETCHING: %s\n", userURL)
start_time:=time.Now()
	response, err := http.Get(userURL)
	end_time:=time.Since(start_time)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	body, err := io.ReadAll(response.Body)
	if err !=nil{
		fmt.Printf("Error Getting  the Response  Body  : %v\n", err)
		
	}
	len_of_res := len(body)
	res_data:= string(body[:300])
	res_type:= response.Header.Get("Content-Type")
	res_header:= response.Header
	defer response.Body.Close()
	fmt.Printf("Response Time: %v\n", end_time)
	fmt.Printf("Status Code: %s\n", response.Status)
		fmt.Printf("Length: %d\n", len_of_res)
			fmt.Printf("Response: %s\n",res_data)
				fmt.Printf("Response Type: %s\n",res_type)
				fmt.Printf("Response Header: %s\n",res_header)
}   