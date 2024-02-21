package numbers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	ctx := context.Background()

	useCases := []struct {
		name          string
		numbers       []int64
		number        int64
		expectedID    int64
		expectedValue int64
		expectedErr   error
	}{
		{
			name:          "should return index and value of existing number",
			numbers:       []int64{4, 6, 8, 9},
			number:        6,
			expectedID:    1,
			expectedValue: 6,
			expectedErr:   nil,
		},
		{
			name:          "should return index and value of first number",
			numbers:       []int64{4, 6, 8, 9},
			number:        4,
			expectedID:    0,
			expectedValue: 4,
			expectedErr:   nil,
		},
		{
			name:          "should return index and value of last number",
			numbers:       []int64{4, 6, 8, 9},
			number:        9,
			expectedID:    3,
			expectedValue: 9,
			expectedErr:   nil,
		},
		{
			name:          "should return index and value of number withing range when no match - min",
			numbers:       []int64{4, 6, 38, 42, 56},
			number:        41,
			expectedID:    2,
			expectedValue: 38,
			expectedErr:   nil,
		},
		{
			name:          "should return index and value of number withing range when no match - max",
			numbers:       []int64{4, 6, 38, 45, 56},
			number:        51,
			expectedID:    4,
			expectedValue: 56,
			expectedErr:   nil,
		},
		{
			name:          "should return error when no match",
			numbers:       []int64{4, 36, 58, 79},
			number:        45,
			expectedID:    0,
			expectedValue: 0,
			expectedErr:   ErrNumberNotFound,
		},
	}

	for _, uc := range useCases {
		t.Run(uc.name, func(t *testing.T) {
			repo := &numbersRepoMock{
				GetAllFunc: func(ctx context.Context) ([]int64, error) {
					return uc.numbers, nil
				},
			}

			svc := NewService(repo)

			index, value, err := svc.Get(ctx, uc.number)

			assert.Equal(t, uc.expectedID, index)
			assert.Equal(t, uc.expectedValue, value)
			assert.ErrorIs(t, err, uc.expectedErr)
		})
	}

	t.Run("should return error on getting all numbers error", func(t *testing.T) {
		repo := &numbersRepoMock{
			GetAllFunc: func(ctx context.Context) ([]int64, error) {
				return nil, assert.AnError
			},
		}

		svc := NewService(repo)

		_, _, err := svc.Get(ctx, 6)

		assert.ErrorIs(t, err, assert.AnError)
	})
}
