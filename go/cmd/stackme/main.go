package main

import "github.com/mec-nyan/data-structures-and-algorithms/go/pkg/stack"

func main() {
	s := stack.New()

	stack.PrintStack(s)

	s.Push("Megan")
	s.Push(7)
	s.Push("Neko")
	s.Push(true)

	stack.PrintStack(s)
}
