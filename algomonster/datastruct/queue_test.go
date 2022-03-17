package datastruct_test

import (
	"facebook/algomonster/datastruct"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	queue := datastruct.Queue{}

	r := require.New(t)
	r.Equal(0, queue.Len())
	r.True(queue.IsEmpty())

	out := queue.Pop()
	r.Nil(out)

	v := 90
	queue.Push(v)
	queue.Push(v + 1)
	r.False(queue.IsEmpty())
	r.Equal(2, queue.Len())

	out = queue.Pop()
	r.NotNil(out)
	r.Equal(v, out.(int))
	r.Equal(1, queue.Len())
	r.False(queue.IsEmpty())

	queue.Pop()
	r.Equal(0, queue.Len())
	r.True(queue.IsEmpty())
}
