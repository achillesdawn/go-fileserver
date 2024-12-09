package storage

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Path         string    `json:"path,omitempty"`
	Name         string    `json:"name,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	Size         int64     `json:"size,omitempty"`
}

func Walk(path string) []FileInfo {

	fileList := make([]FileInfo, 0, 100)
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("%s", err.Error())
			return err
		}

		if !info.IsDir() {

			var file = FileInfo{
				Path:         path,
				Name:         info.Name(),
				ModifiedTime: info.ModTime(),
				Size:         info.Size(),
			}
			fileList = append(fileList, file)
		}
		return nil
	})

	return fileList
}
