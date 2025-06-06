package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mec-nyan/data-structures-and-algorithms/go/pkg/stack"
)

// binaryme converts a decimal number to its binary representation.


func main() {
	for i := 0; i < 16; i++ {
		bin, _ := toBin(i)
		fmt.Printf("dec: %2d, bin: %04s\n", i, bin)
	}
}

func toBin(dec int) (string, error) {
	if dec < 0 {
		return "", errors.New("only positives for now")
	}

	if dec == 0 {
		return "0", nil
	}

	stk := stack.New[int]()

	for dec > 0 {
		stk.Push(dec % 2)
		dec /= 2
	}

	var out string

	for !stk.IsEmpty() {
		next := stk.Pop().(int)
		out += strconv.Itoa(next)
	}

	return out, nil
}
