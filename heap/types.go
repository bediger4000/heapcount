// Package heap implements a max heap data structure as a slice of int.
// The two children of the node at index n are at indexes
// 2n + 1 and 2n + 2 in the slice.
// The index of the parent node of element n
// is (n-1)/2, which takes advantage of ones-complement integer division.
package heap

import (
	"fmt"
	"strings"
)

// Heap as a slice: simplifies maintaining the shape of the heap.
type Heap []int

// IsHeap decides whether an instance of type Heap actually
// has the heap property (returning true) or not (returning false).
// This allows double checking where methods Insert and Delete maintain
// the heap properties.
func (h Heap) IsHeap() bool {
	return isHeap(h, len(h)-1, 0)
}

// ToString creates a lisp-like, parenthesized, string representation
// of an instance of type Heap
func (h Heap) ToString() string {
	s := &strings.Builder{}

	s.WriteRune('(')
	addElement(h, 0, s)
	s.WriteRune(')')

	return s.String()
}

// addElement recursively adds node values to the strings.Builder.
// Because of heap "shape", it does not have to account for a nil
// (non-existent) left child node.
func addElement(h Heap, idx int, s *strings.Builder) {
	s.WriteString(fmt.Sprintf("%d", h[idx]))
	if 2*idx+1 <= len(h)-1 {
		s.WriteString(" (")
		addElement(h, 2*idx+1, s)
		s.WriteRune(')')
	}
	if 2*idx+2 <= len(h)-1 {
		s.WriteString(" (")
		addElement(h, 2*idx+2, s)
		s.WriteRune(')')
	}
}

func isHeap(h Heap, max, idx int) bool {
	leftChild := 2*idx + 1
	if leftChild > max {
		// h[idx] is a leaf node, because of the "shape" of a heap.
		// If node at h[idx] doesn't have a left child, it also doesn't
		// have a right child.
		return true
	}
	if h[idx] < h[leftChild] {
		return false
	}
	if !isHeap(h, max, leftChild) {
		return false
	}
	rightChild := 2*idx + 2
	if rightChild > max {
		// h[idx] might have a left child, but it doesn't have
		// a right child. We've already check heap property with
		// respect to left child.
		return true
	}
	if h[idx] < h[rightChild] {
		return false
	}
	return isHeap(h, max, rightChild)
}
