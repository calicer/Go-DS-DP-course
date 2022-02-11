package Stack

import (
	"errors"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("Nothing in stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item, nil
}
