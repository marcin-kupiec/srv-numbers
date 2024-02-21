package file

import "context"

type Storage struct {
	Numbers []int64
}

func NewStorage() *Storage {
	return &Storage{Numbers: []int64{0, 1, 5, 18, 24, 56, 78, 129, 329, 4359, 438593, 567893, 654837, 768492, 849576, 928372, 989311}}
}

func (s Storage) GetAll(ctx context.Context) ([]int64, error) {
	return s.Numbers, nil
}
