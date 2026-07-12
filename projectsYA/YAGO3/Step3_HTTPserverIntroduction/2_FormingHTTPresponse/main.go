package main

import (
	"fmt"
)

func BuildHTTPResponse(statusLine, headers, body string) string {
	res := fmt.Sprintf("%s\r\n%s\r\n%s", statusLine, headers,body)
	return res
}