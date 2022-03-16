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
			t.Fatalf("%s is not equal to expected value: %s", exp, tt.Exp)
		}
	}
}

func BenchmarkRotationalCipher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		_ = stringsquest.RotationalCipher("abcdefghijklmNOPQRSTUVWXYZ0123456789", 39)
	}
}
