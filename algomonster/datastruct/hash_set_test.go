package datastruct_test

import (
	"facebook/algomonster/datastruct"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashSet(t *testing.T) {
	set := datastruct.HashSet{}

	r := require.New(t)
	r.False(set.Contains(1))

	set.Remove(1) // does not panic
	set.Add(1)
	r.True(set.Contains(1))
	set.Remove(1)
	r.False(set.Contains(1))
}
