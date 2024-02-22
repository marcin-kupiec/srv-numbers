package file

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	ctx := context.Background()

	t.Run("should read and parse input file", func(t *testing.T) {
		fReader := func(fName string) ([]byte, error) {
			assert.Equal(t, fileNameInput, fName)
			text := bytes.NewBufferString("Value: 0 1 5 18 24 56 78 129 329 4359 438593 567893 654837 768492 849576 928372 989311\nIndex: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17").Bytes()
			return text, nil
		}
		expectedNumbers := []int64{0, 1, 5, 18, 24, 56, 78, 129, 329, 4359, 438593, 567893, 654837, 768492, 849576, 928372, 989311}

		storage := NewStorage(fReader)

		err := storage.Init()
		assert.NoError(t, err)

		numbers, err := storage.GetAll(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedNumbers, numbers)
	})

	t.Run("should read and parse input file with only one number", func(t *testing.T) {
		fReader := func(string) ([]byte, error) {
			return bytes.NewBufferString("Value: 0\nIndex: 0").Bytes(), nil
		}
		expectedNumbers := []int64{0}

		storage := NewStorage(fReader)

		err := storage.Init()
		assert.NoError(t, err)

		numbers, err := storage.GetAll(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedNumbers, numbers)
	})

	t.Run("should return error on reading file error", func(t *testing.T) {
		fReader := func(string) ([]byte, error) {
			return nil, assert.AnError
		}

		storage := NewStorage(fReader)

		err := storage.Init()
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("should return error when file is empty", func(t *testing.T) {
		fReader := func(string) ([]byte, error) {
			return bytes.NewBufferString("").Bytes(), nil
		}

		storage := NewStorage(fReader)

		err := storage.Init()
		assert.ErrorContains(t, err, "no lines read from input file")
	})

	t.Run("should return error when file has no numbers", func(t *testing.T) {
		fReader := func(string) ([]byte, error) {
			return bytes.NewBufferString("Value:\nIndex:").Bytes(), nil
		}

		storage := NewStorage(fReader)

		err := storage.Init()
		assert.ErrorContains(t, err, "no numbers in input file")
	})

	t.Run("should return parsing error when file has malformed number", func(t *testing.T) {
		fReader := func(string) ([]byte, error) {
			return bytes.NewBufferString("Value: 1 5 a 5\nIndex:").Bytes(), nil
		}

		storage := NewStorage(fReader)

		err := storage.Init()
		assert.ErrorContains(t, err, "failed to parse string number")
	})
}
