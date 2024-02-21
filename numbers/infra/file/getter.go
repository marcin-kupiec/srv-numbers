package file

import "context"

type Storage struct {
	Numbers []int64
}

func NewStorage() *Storage {
	return &Storage{Numbers: []int64{5, 3, 6, 8, 9, 4, 6, 7, 8, 2, 7, 5, 8, 4}}
}

func (s Storage) GetAll(ctx context.Context) ([]int64, error) {
	return s.Numbers, nil
}
