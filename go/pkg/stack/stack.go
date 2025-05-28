package stack

import "fmt"

type Stack struct {
	items []any
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() any {
	if s.Size() > 0 {
		last := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return last
	}
	return nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Peek() any {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func PrintStack(s *Stack) {
	if s.IsEmpty() {
		println("Stack: empty.")
	} else {
		outputStr := "--------------------------------------------------------------------------------\nStack:"
		for i, v := range s.items {
			outputStr += fmt.Sprintf("\n\t%2d: %v", i+1, v)
		}
		outputStr += fmt.Sprintf("\nSize: %d", s.Size())
		println(outputStr)
	}
}
