// Package hamming calculates the hamming distance between DNA strands
package hamming

import "errors"

// Distance returns the hamming distance between DMA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Cannot compute hamming distance from strands of differing length")
	}

	var hammingDistance int

	for i := range a {
		if a[i] == b[i] {
			continue
		}

		hammingDistance++
	}

	return hammingDistance, nil
}
