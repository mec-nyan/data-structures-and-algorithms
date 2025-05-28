package main

import (
	"fmt"

	"github.com/mec-nyan/data-structures-and-algorithms/go/pkg/stack"
)

type Hooman struct {
	First, Last string
	Age uint
}

func (h *Hooman) String() string {
	return fmt.Sprintf("%s, %s (%d)", h.Last, h.First, h.Age)
}

func main() {
	s := stack.New[string]()

	fmt.Printf("%v\n", s)

	s.Push("Megan")
	s.Push("7")
	s.Push("Neko")
	s.Push("true")

	fmt.Printf("%v\n", s)

	t := stack.New[int]()

	fmt.Printf("%v\n", t)

	t.Push(42)
	t.Push(7)
	t.Push(123)
	t.Push(314)

	fmt.Printf("%v\n", t)

	hoomans := stack.New[Hooman]()

	fmt.Printf("%v\n", hoomans)

	hoomans.Push(Hooman{
		First: "SpongeBob",
		Last: "SquarePants",
		Age: 19,
	})

	hoomans.Push(Hooman{
		First: "Patrick",
		Last: "Star",
		Age: 18,
	})

	hoomans.Push(Hooman{
		First: "Squidward",
		Last: "Tentacles",
		Age: 38,
	})

	hoomans.Push(Hooman{
		First: "Sandy",
		Last: "Cheeks",
		Age: 18,
	})

	hoomans.Push(Hooman{
		First: "Eugene H.",
		Last: "Crabs",
		Age: 18,
	})

	fmt.Printf("%v\n", hoomans)
}
