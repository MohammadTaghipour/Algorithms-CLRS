package main

// OfflineCaching computes the number of cache hits for a given sequence of requests
// using the offline (furthest-in-future) caching strategy.
//
// Important points:
//   - The algorithm knows the entire sequence of requests in advance (offline).
//   - The cache can hold at most k blocks at any time.
//   - On a cache miss when the cache is full, the block whose next use is farthest
//     in the future is evicted.
//
// Parameters:
//   requests []string : slice of block identifiers representing the request sequence
//   k        int      : cache size (maximum number of blocks it can hold)
//
// Returns:
//   int : total number of cache hits
//
// Algorithm:
//   1. Initialize an empty cache (map of block -> next use index).
//   2. For each request b in requests:
//        - If b is already in the cache, increment hits.
//        - Else (cache miss):
//            - If cache has space, insert b.
//            - If cache is full, evict the block whose next use is farthest in the future,
//              then insert b.
//   3. Return the total number of hits.
//
// Time complexity: O(n * k)
// Space complexity: O(k)
func OfflineCaching(requests []string, k int) int {
	cache := make(map[string]int) // store block -> next use index
	hits := 0

	for i := 0; i < len(requests); i++ {
		b := requests[i]

		if _, exists := cache[b]; exists {
			hits++ // Cache hit
		} else {
			if len(cache) < k {
				// Cache has space. only insert
				cache[b] = nextUseIndex(requests, b, i+1)
			} else {
				// Cache full, evict block used farthest in future
				furthestBlock := ""
				furthestIndex := -1
				for block := range cache {
					next := nextUseIndex(requests, block, i+1)
					if next > furthestIndex {
						furthestIndex = next
						furthestBlock = block
					}
				}
				delete(cache, furthestBlock)
				cache[b] = nextUseIndex(requests, b, i+1)
			}
		}
	}

	return hits
}

// nextUseIndex returns the index of the next use of block b after position start.
// If b is never used again, it returns len(requests).
func nextUseIndex(requests []string, b string, start int) int {
	for j := start; j < len(requests); j++ {
		if requests[j] == b {
			return j
		}
	}
	return len(requests) // never used again
}
