package util

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReplaceInFiles(dir, old, new string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = ReplaceInFile(path, old, new)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func ReplaceInFile(filePath, old, new string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var content strings.Builder
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		content.WriteString(string(buf[:n]))
	}

	newContent := strings.ReplaceAll(content.String(), old, new)

	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
