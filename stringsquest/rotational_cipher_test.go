package stringsquest_test

import (
	"facebook/stringsquest"
	"testing"

	"github.com/stretchr/testify/require"
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
		r := require.New(t)
		exp := stringsquest.RotationalCipher(tt.In, tt.InFactor)
		r.Equal(tt.Exp, exp)
	}
}

func BenchmarkRotationalCipher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		_ = stringsquest.RotationalCipher("abcdefghijklmNOPQRSTUVWXYZ0123456789", 39)
	}
}
