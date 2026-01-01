package main

// Huffman constructs a Huffman tree from a set of characters and
// their frequencies using a greedy algorithm.
//
// The algorithm repeatedly combines the two nodes with the smallest
// frequencies into a new internal node whose frequency is the sum
// of the two. This process continues until only one node remains,
// which becomes the root of the Huffman tree.
//
// Time complexity: O(n log n)
//   - Building the initial min-heap takes O(n)
//   - Each of the n−1 iterations performs two Extract-Min and one Insert,
//     each costing O(log n)
//
// Space complexity: O(n)
//   - n nodes are stored in the heap and the resulting tree
func Huffman(C map[rune]int) *HuffmanNode {
	var nodes []*HuffmanNode
	for ch, freq := range C {
		nodes = append(nodes, &HuffmanNode{
			char: ch,
			freq: freq,
		})
	}

	Q := BuildHuffmanMinHeap(nodes)
	n := len(nodes)

	for i := 0; i < n-1; i++ {
		x := Q.ExtractMin()
		y := Q.ExtractMin()

		z := &HuffmanNode{
			freq:  x.freq + y.freq,
			left:  x,
			right: y,
		}
		Q.Insert(z)
	}
	return Q.ExtractMin()
}

type HuffmanNode struct {
	char  rune
	freq  int
	left  *HuffmanNode
	right *HuffmanNode
}

type HuffmanHeap struct {
	data       []*HuffmanNode
	heapSize   int
	betterThan func(a, b *HuffmanNode) bool
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

func (h *HuffmanHeap) heapify(i int) {
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

func BuildHuffmanMinHeap(nodes []*HuffmanNode) *HuffmanHeap {
	h := &HuffmanHeap{
		data:     nodes,
		heapSize: len(nodes),
		betterThan: func(a, b *HuffmanNode) bool {
			return a.freq < b.freq
		},
	}

	for i := h.heapSize/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *HuffmanHeap) ExtractMin() *HuffmanNode {
	if h.heapSize == 0 {
		return nil
	}

	root := h.data[0]

	h.data[0] = h.data[h.heapSize-1]
	h.heapSize--
	h.data = h.data[:h.heapSize]

	if h.heapSize > 0 {
		h.heapify(0)
	}

	return root
}

func (h *HuffmanHeap) Insert(node *HuffmanNode) {
	h.data = append(h.data, node)
	h.heapSize++

	i := h.heapSize - 1
	for i > 0 && h.betterThan(h.data[i], h.data[parent(i)]) {
		h.data[i], h.data[parent(i)] =
			h.data[parent(i)], h.data[i]
		i = parent(i)
	}
}
