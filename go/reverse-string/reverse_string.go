package reverse

import (
	"bytes"
)

// String reverses a string
func String(s string) string {
	var reversed bytes.Buffer

	for i := len(s) - 1; i >= 0; i-- {
		reversed.WriteByte(s[i])
	}

	return reversed.String()
}
