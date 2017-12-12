// Package bob provide everything needed to handle a lackadaisical teenager named Bob
package bob

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"unicode"
)

// reduceToAlphaRegexp filters a string down to only alphanumeric chars
func reduceToAlphaRegexp(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, "")
}

// reduceToAlpha filters a string down to only alphanumeric chars
func reduceToAlpha(s string) string {
	var buffer bytes.Buffer

	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}

		buffer.WriteRune(c)
	}

	return buffer.String()
}

// Hey handles bob's 💩
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	empty := remark == ""
	question := strings.HasSuffix(remark, "?")

	// "yelling" if all alpha are caps
	alphasFromRemark := reduceToAlpha(remark)
	yelling := alphasFromRemark != "" && strings.ToUpper(remark) == remark

	switch {
	case empty:
		return "Fine. Be that way!"

	case yelling && question:
		return "Calm down, I know what I'm doing!"

	case question:
		return "Sure."

	case yelling:
		return "Whoa, chill out!"

	default:
		return "Whatever."
	}
}
