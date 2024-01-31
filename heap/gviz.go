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
	PrintNodePrefixed(out, h, max, idx, "n")
}

func PrintNodePrefixed(out *os.File, h Heap, max, idx int, prefix string) {
	if idx > max {
		return
	}

	if idx != 0 {
		parentIndex := (idx - 1) / 2
		fmt.Fprintf(out, "%s%d -> %s%d\n", prefix, parentIndex, prefix, idx)
	}

	fmt.Fprintf(out, "%s%d [label=\"%d\"];\n", prefix, idx, h[idx])

	leftIndex := 2*idx + 1
	PrintNodePrefixed(out, h, max, leftIndex, prefix)

	rightIndex := 2*idx + 2
	PrintNodePrefixed(out, h, max, rightIndex, prefix)
}
