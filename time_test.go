package _file_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zakon47/_file"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFile(t *testing.T) {
	data := []struct {
		name     string
		filename string
		dirname  string
		data     []byte
	}{
		{name: "test1", filename: "test1", dirname: ".", data: []byte("HI!")},
		{name: "test2", filename: "", dirname: ".", data: []byte("HI!")},
		{name: "test3", filename: "test2", dirname: "test_dir", data: []byte("HI!")},
		{name: "test4", filename: "test3", dirname: "test_dir/dasda", data: []byte("HI!")},
		{name: "test5", filename: "test4", dirname: "test_dir\\dasda", data: []byte("HI!")},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			//создание файла
			err := _file.CreateFile(test.name, test.dirname, test.data)
			assert.NoError(t, err)

			//чтение файла - проверка на содержимое
			file := filepath.Join(test.dirname,test.name)
			b, err := ioutil.ReadFile(file)
			assert.NoError(t, err)
			log.Println(111,file)
			assert.Equal(t, test.data, b)

			//удаление файла
			f, err := os.Stat(file)
			assert.NoError(t, err)
			if f != nil {
				err := os.RemoveAll(file)
				assert.NoError(t, err)
			}
		})
	}

	//удаление тестовую папку
	f, err := os.Stat("test_dir")
	assert.NoError(t, err)
	if f != nil {
		err := os.RemoveAll(f.Name())
		assert.NoError(t, err)
	}
}
