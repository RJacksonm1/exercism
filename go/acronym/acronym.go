// Package acronym generates acronyms!
package acronym

import "strings"

// Abbreviate a string!
func Abbreviate(s string) string {
	var acronym string

	s = strings.Replace(s, "-", " ", -1)
	words := strings.Fields(s)

	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}

	return acronym
}
