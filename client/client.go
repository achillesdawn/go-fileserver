package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func createFileForm(path string) (*bytes.Buffer, *multipart.Writer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file: %w", err)
	}

	var body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return nil, nil, fmt.Errorf("could not create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {

		return nil, nil, fmt.Errorf("copy error creating multipart form: %w", err)
	}
	writer.Close()

	return body, writer, nil
}

func main() {
	url := "http://192.168.2.100:5000/upload"

	body, writer, err := createFileForm("client/hello.txt")
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "go client v0.0.1")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", res.Status)

	defer res.Body.Close()
	bites, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bites))
}
