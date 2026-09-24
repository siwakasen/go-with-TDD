// Package file is example of using file in go
package file

import (
	"errors"
	"io"
	"os"
	"strings"
)

func createFile(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		return errors.New("file already exist")
	}
	if !os.IsNotExist(err) {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	return file.Close()
}

func writeFile(filePath string, text string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func readFile(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0o644)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1024)
	var result strings.Builder

	for {
		n, err := file.Read(buffer)
		result.Write(buffer[:n])

		if err == io.EOF {
			break
		}

		if err != nil {
			return "", err
		}
	}

	return result.String(), file.Close()
}

func simpleReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
