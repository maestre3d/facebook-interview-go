package stringsquest

import (
	"strings"
)

func getRotatedChar(c byte, rotationFactor uint8) byte {
	// Time: O(1) lookup with lang's ASCII character internal hashtable
	// Space: O(m) where m = number of alphanumeric items
	switch {
	case c >= 'A' && c <= 'Z':
		return ((c - 'A' + rotationFactor) % 26) + 'A'
	case c >= 'a' && c <= 'z':
		return ((c - 'a' + rotationFactor) % 26) + 'a'
	case c >= '0' && c <= '9':
		return ((c - '0' + rotationFactor) % 10) + '0'
	}
	return 0
}

func isValidAlphanumeric(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func RotationalCipher(input string, rotationFactor int) string {
	// Time: O(n)
	// Space: O(n)
	wordBuffer := strings.Builder{}
	for i := 0; i < len(input); i++ {
		c := input[i]
		if !isValidAlphanumeric(c) {
			wordBuffer.WriteByte(c)
			continue
		}

		wordBuffer.WriteByte(getRotatedChar(c, uint8(rotationFactor)))
	}

	return wordBuffer.String()
}
