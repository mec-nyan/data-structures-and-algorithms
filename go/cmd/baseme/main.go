package main

import (
	"errors"
	"fmt"

	"github.com/mec-nyan/data-structures-and-algorithms/go/pkg/stack"
)

// baseme converts a decimal number to another base.


func main() {
	fmt.Printf("%4s . %4s . %4s . %4s\n", "hex", "dec", "oct", "bin")
	for i := 0; i < 16; i++ {
		dec := i
		bin, err := decToBase(i, 2)
		if err != nil {
			panic(err)
		}
		oct, err := decToBase(i, 8)
		if err != nil {
			panic(err)
		}
		hex, err := decToBase(i, 16)
		if err != nil {
			panic(err)
		}
		fmt.Printf("0x%02s . %4d . 0o%02s . 0b%04s\n", hex, dec, oct, bin)
	}
}

func decToBase(dec, base int) (string, error) {
	if dec < 0 {
		return "", errors.New("only positives for now")
	}

	if dec == 0 {
		return "0", nil
	}

	if base > 16 {
		return "", errors.New("base too f**king big")
	}

	digits := "0123456789ABCDEF"
	stk := stack.New[int]()

	for dec > 0 {
		stk.Push(dec % base)
		dec /= base
	}

	var out string

	for !stk.IsEmpty() {
		next := stk.Pop().(int)
		out += string(digits[next])
	}

	return out, nil
}
