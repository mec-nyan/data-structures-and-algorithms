package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return "Stack: empty."
	}
	outputStr := strings.Repeat("-", 80) 
	outputStr += fmt.Sprintf("\nStack (%T):", s.items[0])
	for i, v := range s.items {
		outputStr += fmt.Sprintf("\n\t%2d: %v", i+1, v)
	}
	outputStr += fmt.Sprintf("\nSize: %d", s.Size())
	return outputStr
}
