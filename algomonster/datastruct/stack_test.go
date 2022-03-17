package datastruct_test

import (
	"facebook/algomonster/datastruct"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	r := require.New(t)
	stack := datastruct.Stack{}
	r.True(stack.IsEmpty())

	i := 90
	stack.Push(i)
	r.False(stack.IsEmpty())
	r.Equal(1, stack.Len())

	out := stack.Pop()
	r.NotNil(out)
	r.Equal(i, out.(int))
	r.True(stack.IsEmpty())
	r.Equal(0, stack.Len())

	stack.Push(i)
	stack.Push(i + 1)
	out = stack.Pop()
	r.Equal(i+1, out.(int))
}
