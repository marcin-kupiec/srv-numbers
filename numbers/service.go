package numbers

import (
	"context"
	"fmt"
)

var (
	ErrNumberNotFound = fmt.Errorf("number not found")
)

type numbersRepo interface {
	GetAll(ctx context.Context) ([]int64, error)
}

type Getter struct {
	NumbersRepo numbersRepo
}

func NewGetter(repo numbersRepo) *Getter {
	return &Getter{NumbersRepo: repo}
}

func (g Getter) Get(ctx context.Context, index int) (int64, int64, error) {
	numbers, err := g.NumbersRepo.GetAll(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get all numbers: %w", err)
	}

	if index > len(numbers) {
		return 0, 0, ErrNumberNotFound
	}

	return int64(index), numbers[index], nil
}
