package filehelper

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/hlthung/golang-learning/pkg/utils/slicehelper"
)

func Create(fp string) (*os.File, error) {
	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create dir: %w", err)
	}
	return os.Create(fp)
}

// SafeRemoveFileIfExist move a the file to the directory ./deleted/ and remove it
// from the original directory.
func SafeRemoveFileIfExist(ctx context.Context, fileName string) error {
	src, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open file (%s): %v", fileName, err)
	}

	destinationDir := path.Join(path.Dir(fileName), ".deleted")
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return err
	}

	dst, err := os.Create(path.Join(destinationDir, getDeletedFileName(fileName)))
	if err != nil {
		return err
	}
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	if err := dst.Close(); err != nil {
		return err
	}
	if err := os.Remove(fileName); err != nil {
		return err
	}

	return nil
}

func getDeletedFileName(fileName string) string {
	base := RemoveExtension(path.Base(fileName))
	ext := path.Ext(fileName)
	return fmt.Sprintf("_%s_%d.%s", base, time.Now().Unix(), ext)
}

// RemoveExtension allows us to remove an extension from a file path. If you pass in
// to restrictedExtensions, only those extension will be removed. Otherwise, this function
// will remove any extensions available.
// Eg. RemoveExtension("test.csv", "pgp", "gpg") will return test.csv since csv is not within
// the restrictedExtensions.
func RemoveExtension(pathToFile string, restrictedExtensions ...string) string {
	fileExt := filepath.Ext(pathToFile)

	// Do not remove if not within restrictedExtensions
	if len(restrictedExtensions) > 0 && !slicehelper.ContainsString(restrictedExtensions, strings.TrimPrefix(fileExt, ".")) {
		return pathToFile
	}

	return strings.TrimSuffix(pathToFile, fileExt)
}

func FileExists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}

func CopyFile(srcFilePath, dstFilePath string) error {
	src, err := os.Open(srcFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file (%s): %v", srcFilePath, err)
	}

	destinationDir := path.Dir(dstFilePath)
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return err
	}

	dst, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	if err := dst.Close(); err != nil {
		return err
	}

	return nil
}
