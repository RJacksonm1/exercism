// Package twofer facilitates sharing 😉
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith makes you share!
func ShareWith(name string) string {
	name = strings.TrimSpace(name)

	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
