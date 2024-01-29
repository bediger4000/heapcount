package heap

// Delete removes the maximum-value item from the heap
// and returns the remaining heap and that max value
// item.
func (h Heap) Delete() (Heap, int) {
	n := h[0]

	h[0] = h[len(h)-1]
	h = h[:len(h)-1]

	h.siftDown(0)

	return h, n
}

// siftDown checks for appropriate position of item
// in partially-ordered array at index idx.
// It moves items not in the right place.
func (h Heap) siftDown(idx int) {
	if idx > len(h)-1 {
		return
	}

	left := 2*idx + 1
	if left > len(h)-1 {
		return
	}
	if h[idx] < h[left] {
		h[idx], h[left] = h[left], h[idx]
		h.siftDown(left)
	}

	right := 2*idx + 2
	if right > len(h)-1 {
		return
	}
	if h[idx] < h[right] {
		h[idx], h[right] = h[right], h[idx]
		h.siftDown(right)
	}
}
