package main

type Heap struct {
	data       []int
	heapSize   int
	betterThan func(a, b int) bool
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *Heap) heapify(i int) {
	l := left(i)
	r := right(i)

	best := i

	if l < h.heapSize && h.betterThan(h.data[l], h.data[best]) {
		best = l
	}
	if r < h.heapSize && h.betterThan(h.data[r], h.data[best]) {
		best = r
	}

	if best != i {
		h.data[i], h.data[best] = h.data[best], h.data[i]
		h.heapify(best)
	}
}

func BuildMaxHeap(arr []int) *Heap {
	h := &Heap{
		data:     arr,
		heapSize: len(arr),
		betterThan: func(a, b int) bool {
			return a > b
		},
	}

	for i := h.heapSize/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}

	return h
}

func BuildMinHeap(arr []int) *Heap {
	h := &Heap{
		data:     arr,
		heapSize: len(arr),
		betterThan: func(a, b int) bool {
			return a < b
		},
	}

	for i := h.heapSize/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}
