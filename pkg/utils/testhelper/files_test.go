package testhelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequireTempFile(t *testing.T) {
	t.Run("it should create and provide hook for delete a temporary file", func(t *testing.T) {
		file, cleanup := RequireTempFile()
		assert.FileExists(t, file.Name())
		cleanup()
		assert.NoFileExists(t, file.Name())
	})
}

func TestRequireFileReader(t *testing.T) {
	t.Run("it should get a reader for a given file", func(t *testing.T) {
		file, cleanup := RequireTempFile()
		defer cleanup()
		fmt.Fprint(file, "[test]")
		_ = file.Close()

		var got string
		fmt.Fscanf(RequireFileReader(file.Name()), "%s", &got)
		assert.Equal(t, "[test]", got)
	})

	t.Run("it should panic if file do not exist", func(t *testing.T) {
		assert.Panics(t, func() {
			RequireFileReader("not-exist")
		})
	})
}

func TestRequireFileContent(t *testing.T) {
	t.Run("it should get content of a given file", func(t *testing.T) {
		file, cleanup := RequireTempFile()
		defer cleanup()
		fmt.Fprint(file, "[test]")
		_ = file.Close()

		assert.Equal(t, "[test]", RequireFileContent(file.Name()))
	})

	t.Run("it should panic if file do not exist", func(t *testing.T) {
		assert.Panics(t, func() {
			RequireFileContent("not-exist")
		})
	})
}
