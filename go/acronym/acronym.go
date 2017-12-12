// Package acronym generates acronyms!
package acronym

import (
	"bytes"
	"strings"
)

// Abbreviate a string!
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	words := strings.Fields(s)

	if len(words) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	for _, word := range words {
		buffer.WriteString(strings.ToUpper(string(word[0])))
	}

	return buffer.String()
}
