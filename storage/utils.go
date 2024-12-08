package storage

import (
	"errors"
	"fmt"
	"os"
)

func createDir(name string) error {
	info, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("creating asset dir")
		err = os.Mkdir("assets", 0660)
		if err != nil {
			return fmt.Errorf("could not create %s dir: %w", name, err)
		}

		fmt.Println(name, "dir created")
		return nil
	} else if err != nil {
		return fmt.Errorf("could not create %s dir: %w", name, err)
	}

	fmt.Println(name, " : ", info.Size(), "bytes")
	return nil
}

func CreateDirs() error {

	createDir("assets")

	return nil

}
