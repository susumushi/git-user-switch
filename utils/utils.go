package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathParser(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	} else {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("path parse error :%s", err)
		}
		config := homedir + string(os.PathSeparator) + path
		return config, nil
	}
}
