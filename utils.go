package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func printHeaders(r *http.Request) {
	for key, values := range r.Header {
		value := strings.Join(values, " ; ")
		fmt.Println(key, value)
	}
}

func PrintLocalAddress() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("could not determine local interface address: %s\n", err.Error())
		return
	}

	for _, addr := range addrs {
		if strings.HasPrefix(addr.String(), "192.") {
			fmt.Println("listening on: ", addr)
			return
		}
	}
}
