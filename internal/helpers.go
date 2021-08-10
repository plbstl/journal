package internal

import (
	"errors"
	"os"
)

// FileExists checks if a file already exists in the dir.
func FileExists(item, filename, dir string) (bool, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false, errors.New("cannot read " + item + "s directory: " + dir)
	}
	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == filename {
				return true, nil
			}
		}
	}
	return false, nil
}
