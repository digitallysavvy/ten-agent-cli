package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListCreatedFiles(projectName string) error {
	return filepath.Walk(projectName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
}
