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

func main() {
	
	if len(os.Args) <2{
		fmt.Println("Usage: go run main.go <URL> <filename>")
		return
	}
userURL := os.Args[1]
fmt.Printf("FETCHING: %s\n", userURL)
start_time:=time.Now()
req ,err:= http.NewRequest("GET" , userURL , nil)
if err != nil {
	fmt.Println("unable to create  a  new request")
	return ;
}
	 client :=  http.Client{}
	req.Header.Set(
    "User-Agent",
    "goproxy/1.0",
)
fmt.Println("Request Headers:")
fmt.Println(req.Header)
	response ,err:= client.Do(req)
	if err!=nil {
		fmt.Println("unbale  to  get a reponse  back from   the    server")
		return ;
	} 
	end_time:=time.Since(start_time)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err !=nil{
		fmt.Printf("Error Getting  the Response  Body  : %v\n", err)
		
	}
	len_of_res := len(body)
	previewLen := 300

if len(body) < previewLen {
    previewLen = len(body)
}

res_data := string(body[:previewLen])
	res_type:= response.Header.Get("Content-Type")
var prettyJSON []byte
	var data  interface{}
	if strings.Contains(res_type, "application/json"){
err = json.Unmarshal(body, &data)
if err != nil {
    fmt.Println("unable to unmarshal")
}
	prettyJSON, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Print("preetier  unaplicable")
	}
	}

	var dataToSave []byte
if len(os.Args) > 2 {
    if strings.Contains(res_type, "application/json") {
    dataToSave = prettyJSON
} else {
    dataToSave = body
}

err = os.WriteFile(os.Args[2], dataToSave, 0644)
if err != nil {
    fmt.Printf("Error saving response to file: %v\n", err)
}else{
	
fmt.Printf("Response saved to %s\n", os.Args[2])
}
}

	res_header := response.Header
fmt.Println("==================================== REQUEST SUMMARY ====================================")   
	fmt.Printf("Response Time: %v\n", end_time)
	fmt.Printf("Status Code: %s\n", response.Status)
		fmt.Printf("Length: %d\n", len_of_res)
	if strings.Contains(res_type, "application/json") {
		fmt.Printf("Response: %s\n",string(prettyJSON))
	}else{
			fmt.Printf("Response: %s\n",res_data)
	}
				fmt.Printf("Response Type: %s\n",res_type)
		fmt.Println("==================================== RESPONSE HEADER ====================================")   
				for key,value:= range res_header{
					
					fmt.Println(key,":",value)
					
				}
}  