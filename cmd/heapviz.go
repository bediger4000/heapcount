package main

// heap visualization via GraphViz
// Reads integers from command line, inserts them into an
// instance of the quick-and-dirty heap implemenatation.
// Checks that instance has the correct partial ordering,
// then puts a Dot-format GraphViz input on stdout.

import (
	"fmt"
	"heapcounter/heap"
	"os"
	"strconv"
)

func main() {
	var h heap.Heap
	for _, str := range os.Args[1:] {
		if n, err := strconv.Atoi(str); err == nil {
			h = h.Insert(n)
		}
	}
	if h.IsHeap() {
		fmt.Printf("/* It's a heap */\n")
	} else {
		fmt.Printf("/* It's NOT a heap */\n")
	}
	heap.PrintDot(os.Stdout, h)
}
