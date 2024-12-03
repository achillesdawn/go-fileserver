package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	path := filepath.Join("received", header.Filename)
	writeFile, err := os.Create(path)
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	printHeaders(r)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	err := handleFileUpload(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println("could not write to response")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("status: ok"))
	if err != nil {
		fmt.Println("could not write to response")
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHandler)

	server := http.Server{
		Addr:    ":5000",
		Handler: mux,
	}
	fmt.Println("Listening on port :5000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
