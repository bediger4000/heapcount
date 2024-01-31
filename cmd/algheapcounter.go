package main

// algorithmic determination of number of heaps of N
// distinct integers.
// https://mathworld.wolfram.com/Heap.html

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	count := a(n)

	fmt.Printf("%d heaps of %d distinct integers\n", count, n)
}

func a(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	h := int(math.Log2(float64(n+1))) - 1

	h2 := 1 // 2^h
	for i := 0; i < h; i++ {
		h2 *= 2
	}

	b := h2 - 1

	r := n - 1 - 2*b

	r1 := r - int(float64(r)/float64(h2))*(r-h2)
	r2 := r - r1

	return combinations(n-1, b+r1) * a(b+r1) * a(b+r2)
}

func combinations(n, r int) int {
	c := math.Gamma(float64(n + 1))
	c /= math.Gamma(float64(r + 1))
	c /= math.Gamma(float64(n - r + 1))
	return int(c)
}
