package tg

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInputFile(t *testing.T) {
	file := NewInputFile("test.txt", strings.NewReader("test"))
	if assert.NotNil(t, file) {
		assert.Equal(t, "test.txt", file.Name)
		assert.NotNil(t, file.Body)
		err := file.Close()
		assert.NoError(t, err)
	}
}

func TestNewInputFileBytes(t *testing.T) {
	file := NewInputFileBytes("test.txt", []byte("test"))

	if assert.NotNil(t, file) {
		assert.Equal(t, "test.txt", file.Name)
		data, err := ioutil.ReadAll(file.Body)
		assert.NoError(t, err, "read file")
		assert.Equal(t, []byte("test"), data)
	}
}

func TestNewInputFileLocal(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		file, err := NewInputFileLocal("./README.md")
		if assert.NoError(t, err) {
			defer file.Close()

			assert.Equal(t, "README.md", file.Name)
		}
	})

	t.Run("FAIL", func(t *testing.T) {
		file, err := NewInputFileLocal("./NOT_FOUND.txt")
		assert.Error(t, err)
		assert.Nil(t, file)
	})
}

func TestNewInputFileLocalBuffer(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		file, err := NewInputFileLocalBuffer("./README.md")
		if assert.NoError(t, err) {
			defer file.Close()

			assert.Equal(t, "README.md", file.Name)
		}
	})

	t.Run("FAIL", func(t *testing.T) {
		file, err := NewInputFileLocalBuffer("./NOT_FOUND.txt")
		assert.Error(t, err)
		assert.Nil(t, file)
	})
}