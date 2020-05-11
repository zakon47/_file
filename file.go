package _file


import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrNoNameFile = errors.New("File's name is empty")
)
const (
	mode = 0744
)

func CreateFile(name string, dir string, data []byte) error {
	if dir == ""{
		dir = "."
	}
	if name == ""{
		return ErrNoNameFile
	}
	if err := os.MkdirAll(dir, mode); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, name), data, 0744); err != nil {
		return err
	}
	return nil
}
func ReadFile(name string, dir string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(dir, name))
}