package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/achillesdawn/go-fileserver/storage"
)

func handleFileUpload(r *http.Request) error {

	file, header, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("could not get file from request: %w", err)
	}

	fmt.Printf("filename: %s\nsize: %d\n", header.Filename, header.Size)

	path := filepath.Join("assets", header.Filename)

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

func errorHandle(err error, w http.ResponseWriter) {

	fmt.Println("Error:", err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		fmt.Println("could not write to response")
	}
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
		errorHandle(err, w)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("status: ok"))
	if err != nil {
		fmt.Println("could not write to response")
		return
	}
}

func fileListHandler(w http.ResponseWriter, r *http.Request) {

	serveDir, err := filepath.Abs("./assets")
	if err != nil {
		panic(err)
	}

	files := storage.Walk(serveDir)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

func newServer(port uint16) *http.Server {
	mux := http.NewServeMux()

	serveDir, err := filepath.Abs("./assets")
	if err != nil {
		panic(err)
	}

	mux.Handle(
		"/files/",
		http.StripPrefix(
			"/files/",
			http.FileServer(http.Dir(serveDir)),
		),
	)

	mux.HandleFunc("/upload", uploadHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	mux.HandleFunc("/filelist", fileListHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	PrintLocalAddress(port)

	return &server
}
