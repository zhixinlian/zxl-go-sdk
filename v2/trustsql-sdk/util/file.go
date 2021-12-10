package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	// ErrFileExists file exists
	ErrFileExists = errors.New("file exists")
)

// CreateDirIfNotExist create dir
func CreateDirIfNotExist(dir string) error {
	if exist, err := FileExists(dir); !exist || err != nil {
		if err != nil {
			return err
		}
		err = os.MkdirAll(dir, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// FileExists check file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// FileWrite write file to path
func FileWrite(file string, content []byte, overwrite bool) error {
	// Create the keystore directory with appropriate permissions
	if err := CreateDirIfNotExist(filepath.Dir(file)); err != nil {
		return err
	}
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()

	if overwrite {
		if exist, _ := FileExists(file); exist {
			if err := os.Remove(file); err != nil {
				os.Remove(f.Name())
				return err
			}
		}
	}

	return os.Rename(f.Name(), file)
}
