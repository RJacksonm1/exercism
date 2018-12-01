package isogram

import "strings"

// IsIsogram test whether a word is an isogram!
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, l := range word {
		if l == '-' || l == ' ' {
			continue
		}

		if strings.ContainsRune(word[:i], l) {
			return false
		}
	}

	return true
}
