package isogram

import "strings"

// IsIsogram test whether a word is an isogram!
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)

	for i, l := range word {
		if strings.Contains(word[:i], string(l)) {
			return false
		}
	}

	return true
}
