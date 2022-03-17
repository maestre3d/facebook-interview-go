package leetcode_test

import (
	"facebook/leetcode"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	exp := leetcode.RemoveDuplicateLetters("bcabc")
	t.Log(exp)
}
