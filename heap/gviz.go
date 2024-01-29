package heap

import (
	"fmt"
	"os"
)

// PrintDot writes a GraphViz dot-format text representation
// of an instance of type Heap on the file descriptor.
func PrintDot(out *os.File, h Heap) {
	fmt.Fprintf(out, "digraph g {\n")
	printNode(out, h, len(h)-1, 0)
	fmt.Fprintf(out, "}\n")
}

func printNode(out *os.File, h Heap, max, idx int) {
	if idx > max {
		return
	}

	if idx != 0 {
		parentIndex := (idx - 1) / 2
		fmt.Fprintf(out, "n%d -> n%d\n", parentIndex, idx)
	}

	fmt.Fprintf(out, "n%d [label=\"%d\"];\n", idx, h[idx])

	leftIndex := 2*idx + 1
	printNode(out, h, max, leftIndex)

	rightIndex := 2*idx + 2
	printNode(out, h, max, rightIndex)
}
