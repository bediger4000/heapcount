package main

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

	c := make(chan []int)

	go outerGenerate(c, ary)

	// track unique heaps
	uniqueHeaps := make(map[string]bool)

	count := 0
	for h := range c {
		hp := heap.Heap(h)
		if hp.IsHeap() {
			stringRep := fmt.Sprintf("%v", h)
			if _, ok := uniqueHeaps[stringRep]; !ok {
				count++
				fmt.Printf("%s\t%v\n", hp.ToString(), h)
				uniqueHeaps[stringRep] = true
			}
		}
	}
	fmt.Printf("%d heaps possible\n", count)
}

func outerGenerate(c chan []int, ary []int) {
	generate(c, len(ary), ary)
	close(c)
}

func generate(c chan []int, k int, a []int) {
	if k == 1 {
		cpy := make([]int, len(a))
		copy(cpy, a)
		c <- cpy
		return
	}

	generate(c, k-1, a)

	for i := 0; i < k-1; i++ {
		if (k % 2) == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		generate(c, k-1, a)
	}
}
