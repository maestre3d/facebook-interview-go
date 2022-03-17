package datastruct_test

import (
	"facebook/algomonster/datastruct"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircularQueue(t *testing.T) {
	queue := datastruct.NewCircularQueue(2)

	r := require.New(t)
	i := queue.Pop()
	r.Nil(i)

	v := 90
	queue.Push(v)
	queue.Push(v + 1)
	queue.Push(v + 2)

	i = queue.Pop()
	r.NotNil(i)
	r.Equal(v+1, i.(int))

	i = queue.Pop()
	r.NotNil(i)
	r.Equal(v+2, i.(int))

	i = queue.Pop()
	r.Nil(i)

	queue.Push(v)
	i = queue.Pop()
	r.NotNil(i)
	r.Equal(v, i.(int))
}
