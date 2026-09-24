package file

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	path := "/Users/siwakasen/repos/go-with-TDD/file/"

	t.Run("create new file", func(t *testing.T) {
		fileName := "example.txt"
		createFile(path + fileName)

		_, err := os.Stat(path + fileName)
		if err != nil {
			t.Fatalf("file not found: %v", err)
		}
	})
}
