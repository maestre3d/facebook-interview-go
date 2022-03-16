package stringsquest

import (
	"regexp"
	"strings"
)

const alphanumericPattern = "^[A-z0-9]$"

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

func RotationalCipher(input string, rotationFactor int) string {
	// Time: O(n)
	// Space: O(n)
	wordBuffer := strings.Builder{}
	for i := 0; i < len(input); i++ {
		c := input[i]
		if ok, _ := regexp.Match(alphanumericPattern, []byte{c}); !ok {
			wordBuffer.WriteByte(c)
			continue
		}

		wordBuffer.WriteByte(getRotatedChar(c, uint8(rotationFactor)))
	}

	return wordBuffer.String()
}
