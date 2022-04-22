/*
 * @brief 文件操作
 */

package encapsutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// PathExists probe path is exist.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

// IsDir probe path is dir.
func IsDir(path string) bool {
	if PathExists(path) {
		s, err := os.Stat(path)
		if err != nil {
			return false
		}

		return s.IsDir()
	}
	return false
}

// IsFile probe path is file.
func IsFile(path string) bool {
	if PathExists(path) {
		s, err := os.Stat(path)
		if err != nil {
			return false
		}

		return !s.IsDir()
	}
	return false
}

// MustSaveToFile if path file not exist, first create it and write binary to file.
func MustSaveToFile(binary []byte, path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if IsDir(abs) {
		return fmt.Errorf("path [%s] already exists and it's Dir", abs)
	}

	if err := MustCreateDir(filepath.Dir(abs), os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(abs, binary, 0666)
}

// MustCreateDir create path is not exist,
func MustCreateDir(path string, perm os.FileMode) error {
	if !PathExists(path) {
		return os.MkdirAll(path, perm)
	}
	if !IsDir(path) {
		return fmt.Errorf("path is already occupied by a file")
	}
	return nil
}

// MustCreateFile create a file if it not exist.
func MustCreateFile(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if IsFile(abs) {
		return nil
	}

	if IsDir(abs) {
		return fmt.Errorf("%s is dir", abs)
	}

	if err = MustCreateDir(filepath.Dir(abs), os.ModePerm); err != nil {
		return err
	}

	var f *os.File
	f, err = os.Create(abs)
	if err != nil {
		return err
	}
	return f.Close()
}

// MustAppendToFile if path file not exist, first create it and append write binary to file.
func MustAppendToFile(binary []byte, path string) error {
	if !IsFile(path) {
		return MustSaveToFile(binary, path)
	}

	fd, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()

	n, err := fd.Write(binary)
	if err == nil && n != len(binary) {
		return fmt.Errorf("not write full")
	}
	return nil
}
