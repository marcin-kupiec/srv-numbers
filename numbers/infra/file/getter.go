package file

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

const (
	fileNameInput = "input.txt"
)

type fileReader func(string) ([]byte, error)

type Storage struct {
	fileReader fileReader
	numbers    []int64
}

func NewStorage(fReader fileReader) *Storage {
	return &Storage{fileReader: fReader}
}

func (s *Storage) Init() error {
	b, err := s.fileReader(fileNameInput)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	lines := strings.Split(string(b), "\n")
	if len(lines) == 1 {
		return fmt.Errorf("no lines read from input file")
	}

	numbersStr := strings.Split(lines[0], " ")
	if len(numbersStr) <= 1 {
		return fmt.Errorf("no numbers in input file")
	}

	numbers := make([]int64, len(numbersStr)-1)
	for i := 1; i < len(numbersStr); i++ {
		n, err := strconv.ParseInt(numbersStr[i], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse string number=%s index=%d: %w", numbersStr[i], i, err)
		}
		numbers[i-1] = n
	}

	s.numbers = numbers
	return nil
}

func (s *Storage) GetAll(ctx context.Context) ([]int64, error) {
	return s.numbers, nil
}
