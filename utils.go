package main

import (
	"fmt"
	"net/http"
	"strings"
)

func printHeaders(r *http.Request) {
	for key, values := range r.Header {
		value := strings.Join(values, " ; ")
		fmt.Println(key, value)
	}
}
