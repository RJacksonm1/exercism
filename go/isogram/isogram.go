package isogram

import (
	"strings"
	"unicode"
)

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

// IsIsogramWithSeenMap - see if a map of seen runes is faster
func IsIsogramWithSeenMap(word string) bool {
	word = strings.ToLower(word)
	seenRunes := make(map[rune]bool, len(word))

	for _, l := range word {
		if l == '-' || l == ' ' {
			continue
		}

		if seenRunes[l] {
			return false
		}

		seenRunes[l] = true
	}

	return true
}

// IsIsogramWithSeenMapWithUnicodeLower - and if we use unicode.ToLower instead of looping the string twice?
func IsIsogramWithSeenMapWithUnicodeLower(word string) bool {
	seenRunes := make(map[rune]bool, len(word))

	for _, l := range word {
		if l == '-' || l == ' ' {
			continue
		}

		l = unicode.ToLower(l)
		if seenRunes[l] {
			return false
		}

		seenRunes[l] = true
	}

	return true
}
