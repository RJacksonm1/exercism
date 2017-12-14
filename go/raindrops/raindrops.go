// Package raindrops is onomatopoeiamous
package raindrops

import (
	"fmt"
)

// Convert onomatopoeiazes an amount of raindrops
func Convert(n int) string {
	var onomatopoeia string

	if n%3 == 0 {
		onomatopoeia += "Pling"
	}

	if n%5 == 0 {
		onomatopoeia += "Plang"
	}

	if n%7 == 0 {
		onomatopoeia += "Plong"
	}

	if onomatopoeia == "" {
		return fmt.Sprintf("%d", n)
	}

	return onomatopoeia
}
