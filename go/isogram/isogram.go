package isogram

import (
	"strings"
)

// IsIsogram test whether a word is an isogram!
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, l := range word {
		s := string(l)

		if s == "-" || s == " " {
			continue
		}

		if strings.Contains(word[:i], s) {
			return false
		}
	}

	return true
}
