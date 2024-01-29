package main

// Fill in a heap data structure with integers from command line
// using heap.Heap.Insert method calls.
// Check if the data structure actually has the partial ordering
// required of a heap data structure.
// Outputs two views of the filled-in heap structure.

import (
	"fmt"
	"heapcounter/heap"
	"log"
	"os"
	"strconv"
)

func main() {
	var h heap.Heap
	for _, str := range os.Args[1:] {
		if n, err := strconv.Atoi(str); err == nil {
			h = h.Insert(n)
		} else {
			log.Fatal(err)
		}
	}

	phrase := " not"
	if h.IsHeap() {
		phrase = ""
	}

	fmt.Printf("actually is%s a heap of %d items\n", phrase, len(h))

	fmt.Printf("heap %v\n", h)
	fmt.Printf("lispy: %v\n", h.ToString())
}
