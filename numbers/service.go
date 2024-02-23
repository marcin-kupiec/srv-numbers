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

	return s.search(number, numbers)
}

func (s Service) search(number int64, numbers []int64) (int64, int64, error) {
	delta := int64(math.Floor(conformationCoeff * float64(number)))
	min := number - delta
	max := number + delta

	notExactID := int64(-1)
	notExactValue := int64(-1)

	start := int64(0)
	end := int64(len(numbers) - 1)

	for start <= end {
		middleID := int64(math.Floor(float64(start+end) / 2))
		middleValue := numbers[middleID]

		// return id and value if found
		if middleValue == number {
			return middleID, middleValue, nil
		}

		// save value within the range in case no exact matching number in the slice
		if middleValue >= min && middleValue <= max {
			notExactID = middleID
			notExactValue = middleValue
		}

		// continue in left or right half
		if middleValue < number {
			start = middleID + 1
		} else {
			end = middleID - 1
		}
	}

	// no exact match, so check if there is anything within the range
	if notExactID != -1 {
		return notExactID, notExactValue, nil
	}

	return 0, 0, ErrNumberNotFound
}
