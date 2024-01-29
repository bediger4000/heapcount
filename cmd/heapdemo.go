package main

// heapdemo - read string representations of integers from command line,
// convert to ints, insert into a heap.
// Read maximum values from heap one at a time, print them out.

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

	for len(h) > 0 {
		var max int
		h, max = h.Delete()
		fmt.Printf("%d ", max)
	}
	fmt.Println()
}
