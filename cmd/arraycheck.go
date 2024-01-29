package main

// Read a sequence of integers from the command line,
// creating a slice with the same order in the process.
// Cast that slice to a heap.Heap instance, then check
// if the slice, considered as an array-backed data structure,
// has the heap property or not.

import (
	"fmt"
	"heapcounter/heap"
	"os"
	"strconv"
)

func main() {
	var ary []int
	for _, str := range os.Args[1:] {
		if n, err := strconv.Atoi(str); err == nil {
			ary = append(ary, n)
		}
	}
	if heap.Heap(ary).IsHeap() {
		fmt.Printf("It's a heap\n")
	} else {
		fmt.Printf("It's NOT a heap\n")
	}
}
