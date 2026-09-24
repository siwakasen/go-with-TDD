package file

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	path := "/Users/siwakasen/repos/go-with-TDD/file/"

	fileName := "example.txt"
	t.Run("create new file", func(t *testing.T) {
		err := createFile(path + fileName)
		if err != nil {
			t.Errorf("file failed to create: %v", err)
		}

		_, err = os.Stat(path + fileName)
		if err != nil {
			t.Errorf("file not found: %v", err)
		}
	})

	t.Run("write a file", func(t *testing.T) {
		want := "hello world!"

		err := writeFile(path+fileName, want)
		if err != nil {
			t.Errorf("failed to write within file %v", err)
		}

		got, err := readFile(path + fileName)
		if err != nil {
			t.Errorf("failed to read the file %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("delete a file", func(t *testing.T) {
		err := deleteFile(path + fileName)
		if err != nil {
			t.Errorf("file failed to delete: %v", err)
		}
		_, err = os.Stat(path + fileName)
		if err == nil {
			t.Errorf("file is still exist %v", err)
		}
	})
}
