package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func printHeaders(r *http.Request) {
	for key, values := range r.Header {
		value := strings.Join(values, " ; ")
		fmt.Println(key, value)
	}
}

func handleFileUpload(r *http.Request) error {

	file, header, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("could not get file from request: %w", err)
	}

	fmt.Printf("filename: %s\nsize: %d\n", header.Filename, header.Size)

	writeFile, err := os.Create("received")
	if err != nil {
		return fmt.Errorf("could not create received file: %w", err)
	}

	n, err := io.Copy(writeFile, file)
	if err != nil {
		return fmt.Errorf("could not copy file to destination: %w", err)
	}

	fmt.Println("wrote ", n, "bytes")

	return nil
}

func uploadHander(w http.ResponseWriter, r *http.Request) {

	printHeaders(r)
	err := handleFileUpload(r)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println("could not write to response")
		}
	}

	_, err = w.Write([]byte("status: ok"))
	if err != nil {
		fmt.Println("could not write to response")
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHander)

	server := http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
