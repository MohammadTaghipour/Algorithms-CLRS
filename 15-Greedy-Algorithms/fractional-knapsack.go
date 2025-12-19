package main

// FractionalKnapsack computes the maximum total value that fits in a knapsack of capacity W.
//
// It assumes that the arrays w (weights) and v (values) are already sorted in descending order
// of value per unit weight (v[i]/w[i]). The algorithm iterates through the items, taking each
// item entirely if it fits, or a fraction if the remaining capacity is smaller than the item's weight.
//
// Parameters:
//   W - maximum capacity of the knapsack
//   w - slice of item weights
//   v - slice of item values
//   n - number of items
//
// Returns:
//   The maximum total value that can be obtained in the knapsack.
//
// Time complexity: O(nlgn)   (assuming the items are pre-sorted with O(lgn))
// Space complexity: O(1)
func FractionalKnapsack(W float64, w, v []float64, n int) float64 {
	var totalValue float64 = 0.0
	remainingCapacity := W

	for i := 0; i < n && remainingCapacity > 0; i++ {
		if w[i] <= remainingCapacity {
			// Take the whole item
			totalValue += v[i]
			remainingCapacity -= w[i]
		} else {
			// Take a fraction of the item
			fraction := remainingCapacity / w[i]
			totalValue += v[i] * fraction
			remainingCapacity = 0
			break
		}
	}

	return totalValue
}
