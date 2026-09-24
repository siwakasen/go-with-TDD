// Package file is example of using file in go
package file

import (
	"fmt"
	"os"
)

func createFile(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, _ := os.Create(filePath)

		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println(err)
			}
		}()
	}
}
