package testhelper

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func RequireFileContent(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func RequireFileReader(filename string) io.Reader {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(content)
}

func RequireTempFile() (f *os.File, cleanup func()) {
	destinationFile, err := ioutil.TempFile("", "test_*.txt")
	if err != nil {
		panic(fmt.Errorf("failed to create tmp file: %v", err))
	}
	return destinationFile, func() {
		destinationFile.Close()
		os.Remove(destinationFile.Name())
	}
}
