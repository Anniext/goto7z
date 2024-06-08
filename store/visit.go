package store

import (
	"fmt"
	"os"
)

func Visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	return nil
}
