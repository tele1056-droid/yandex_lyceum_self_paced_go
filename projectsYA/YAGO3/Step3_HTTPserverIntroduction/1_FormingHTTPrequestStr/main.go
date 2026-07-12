package main

import (
	"fmt"
)

func BuildHTTPRequest(method, url, host, headers, body string) string {
	http1 := "HTTP/1.1"
	
	res := fmt.Sprintf("%s %s %s\r\nHost: %s\r\n%s\r\n%s", method, url,http1, host, headers, body)
	return res
}

func main() {
	result := BuildHTTPRequest(
	"POST",
	"/api/users",
	"example.com",
	"",
	"",
	)
	fmt.Println(result)
	// url := "/api/data"

	// met := "GET"

	// res := fmt.Sprintf("%s %s", met, url)
	// fmt.Println(res)
	
}