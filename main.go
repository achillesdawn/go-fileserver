package main

import (
	"fmt"

	"github.com/achillesdawn/go-fileserver/storage"
)

func main() {
	err := storage.CreateDirs()
	if err != nil {
		panic(err)
	}

	server := newServer(":5000")

	fmt.Println("Listening on port :5000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
