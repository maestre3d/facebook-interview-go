package stringsquest_test

import (
	"facebook/stringsquest"
	"testing"
)

var rotationalCipherTests = []struct {
	In       string
	InFactor int
	Exp      string
}{
	{
		In:       "Zebra-493?",
		InFactor: 3,
		Exp:      "Cheud-726?",
	},
	{
		In:       "abcdefghijklmNOPQRSTUVWXYZ0123456789",
		InFactor: 39,
		Exp:      "nopqrstuvwxyzABCDEFGHIJKLM9012345678",
	},
}

func TestRotationalCipher(t *testing.T) {
	for _, tt := range rotationalCipherTests {
		exp := stringsquest.RotationalCipher(tt.In, tt.InFactor)
		if exp != tt.Exp {
			t.Fail()
		}
	}
}
