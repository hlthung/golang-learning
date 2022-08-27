package filehelper

import (
	"context"
	"fmt"
	"github.com/hlthung/golang-learning/pkg/utils/testhelper"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	fp := "/tmp/test/writetest.txt"
	_, err := Create(fp)
	assert.NoError(t, err)
	_, err = os.Stat(fp) // check if file exists
	assert.NoError(t, err)
	_ = os.RemoveAll("/tmp/test") // cleanup
}

func TestRemoveExtension(t *testing.T) {
	var removeExtensionTest = []struct {
		filePath     string
		restrictions []string
		want         string
	}{
		{"folder1/file.jpg.gpg", []string{"gpg"}, "folder1/file.jpg"},
		{"file2.pdf.pgp", []string{"pgp"}, "file2.pdf"},
		{"folder2/file.jpg", []string{"pgp"}, "folder2/file.jpg"},
		{"folder2/file.jpg.csv", []string{}, "folder2/file.jpg"},
	}

	for _, tt := range removeExtensionTest {
		filePath := RemoveExtension(tt.filePath, tt.restrictions...)
		if filePath != tt.want {
			t.Errorf(`removeExtension(%q) = (%v) != (%v)`, tt.filePath, tt.want, filePath)
		}
	}

}

func TestSafeRemoveFileIfExist(t *testing.T) {
	const inputFileName = "in.txt"
	t.Run("it should move file to .deleted folder", func(t *testing.T) {
		expectedContent := uuid.New().String()
		file, err := os.OpenFile(inputFileName, os.O_CREATE|os.O_WRONLY, 0644)
		if !assert.NoError(t, err) {
			return
		}

		fmt.Fprint(file, expectedContent)
		file.Close()

		assert.FileExists(t, file.Name())

		SafeRemoveFileIfExist(context.Background(), file.Name())
		defer os.RemoveAll(".deleted")

		assert.NoFileExists(t, file.Name())

		dir, err := ioutil.ReadDir(path.Dir(file.Name()) + "/.deleted")
		if !assert.NoError(t, err) {
			return
		}
		if !assert.Len(t, dir, 1) {
			return
		}

		name := ".deleted/" + dir[0].Name()
		actualContent := testhelper.RequireFileContent(name)

		assert.Equal(t, expectedContent, actualContent)
	})

	t.Run("it not return error file do not exist", func(t *testing.T) {
		err := SafeRemoveFileIfExist(context.Background(), "not-exist")
		assert.NoError(t, err)
	})
}

func TestFileExists(t *testing.T) {
	fp := "/tmp/test/testFileExists.txt"
	_, err := Create(fp)
	assert.NoError(t, err)

	res := FileExists(fp)
	assert.Equal(t, true, res)
	res = FileExists("/tmp/test/fileNotExists.txt")
	assert.Equal(t, false, res)
	_ = os.RemoveAll("/tmp/test") // cleanup
}

func TestCopyFile(t *testing.T) {
	fp := "/tmp/test/original.txt"
	_, err := Create(fp)
	assert.NoError(t, err)

	err = CopyFile(fp, "/tmp/test/copy.txt")
	assert.NoError(t, err)
	_, err = os.Stat("/tmp/test/copy.txt")
	assert.NoError(t, err)

	_ = os.RemoveAll("/tmp/test") // cleanup
}
