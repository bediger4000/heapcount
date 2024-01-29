package main

/*
 * Permutations of N integers via "Heap's Algorithm"
 * https://en.wikipedia.org/wiki/Heap's_algorithm
 */

import (
	"fmt"
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

	generate(len(ary), ary)
}

// generate comprises a Go transliteration of whatever
// pseudocode in which Wikipedia shows the algorithm.
func generate(k int, a []int) {
	if k == 1 {
		fmt.Printf("%v\n", a)
		return
	}

	generate(k-1, a)

	for i := 0; i < k-1; i++ {
		if (k % 2) == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		generate(k-1, a)
	}
}
