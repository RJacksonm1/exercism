// Package raindrops is onomatopoeiamous
package raindrops

import (
	"fmt"
	"math"
	"sort"
)

// Factors returns all the prime factors of a given integer
func Factors(n int) []int {
	// Add factors to a map for automatic deduplication
	factorMap := make(map[int]struct{})

	// How many 2s are in n
	for n%2 == 0 {
		factorMap[2] = struct{}{}
		n /= 2
	}

	// Iterate through the odd nums and find which ones
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			factorMap[i] = struct{}{}
			n /= i
		}
	}

	// If the final prime is greater than 2
	if n > 2 {
		factorMap[n] = struct{}{}
	}

	var factors []int
	for factor := range factorMap {
		factors = append(factors, factor)
	}

	// Maps are iterated in random order 😱
	sort.Ints(factors)

	return factors
}

// Convert onomatopoeiazes an amount of raindrops
func Convert(n int) string {
	var onomatopoeia string

	factors := Factors(n)

	fmt.Printf("%d: %d", n, factors)

	for _, factor := range factors {
		switch factor {
		case 3:
			onomatopoeia += "Pling"
		case 5:
			onomatopoeia += "Plang"
		case 7:
			onomatopoeia += "Plong"
		}
	}

	if onomatopoeia == "" {
		return fmt.Sprintf("%d", n)
	}

	return onomatopoeia
}
