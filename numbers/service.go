package numbers

import (
	"context"
	"fmt"
	"math"
)

const (
	conformationCoeff = 0.1
)

var (
	ErrNumberNotFound = fmt.Errorf("number not found")
)

//go:generate moq -out numbersRepoMock_test.go . numbersRepo
type numbersRepo interface {
	GetAll(ctx context.Context) ([]int64, error)
}

type Service struct {
	NumbersRepo numbersRepo
}

func NewService(repo numbersRepo) *Service {
	return &Service{NumbersRepo: repo}
}

func (s Service) Get(ctx context.Context, number int64) (int64, int64, error) {
	numbers, err := s.NumbersRepo.GetAll(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get all numbers: %w", err)
	}

	delta := int64(math.Floor(conformationCoeff * float64(number)))
	min := number - delta
	max := number + delta

	notExactID := -1
	notExactValue := int64(-1)

	for id, n := range numbers {
		if n > max {
			break
		}
		if n == number {
			return int64(id), n, nil
		}
		if n > number && notExactID != -1 {
			return int64(notExactID), notExactValue, nil
		}
		if n >= min && n <= max {
			notExactID = id
			notExactValue = n
		}
	}

	if notExactID != -1 {
		return int64(notExactID), notExactValue, nil
	}
	return int64(0), int64(0), ErrNumberNotFound
}
