package main

import (
	"fmt"

	"github.com/mec-nyan/data-structures-and-algorithms/go/pkg/sort"
)


func main() {
	arr := []int{88, 42, 7, 101, 999, 23, 65, 7}

	fmt.Println("Before:", arr)

	sort.BubbleSort(arr)

	fmt.Println("After:", arr)
}
