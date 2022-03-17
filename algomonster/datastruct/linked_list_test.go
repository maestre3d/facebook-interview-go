package datastruct_test

import (
	"facebook/algomonster/datastruct"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedListNode(t *testing.T) {
	l := datastruct.LinkedList{}
	nodeA := "foo"
	nodeB := "bar"
	l.Insert(nodeA)
	l.Insert(nodeB)

	r := require.New(t)
	r.Equal(2, l.Len())

	nodeC := "baz"
	l.Insert(nodeC)
	r.Equal(3, l.Len())

	n := l.GetAt(2)
	r.NotNil(n)
	r.Equal(nodeC, n.Value)

	n = l.GetAt(1)
	r.NotNil(n)
	r.Equal(nodeB, n.Value)

	n = l.GetAt(0)
	r.NotNil(n)
	r.Equal(nodeA, n.Value)

	l.InsertAt(1, "foobar")
	n = l.GetAt(1)
	r.NotNil(n)
	r.Equal("foobar", n.Value)

	nodeD := "aruiz"
	l.InsertAt(0, nodeD)
	n = l.GetAt(0)
	r.NotNil(n)
	r.Equal(nodeD, n.Value)
	r.Equal(nodeA, n.Next.Value)

	i := l.GetPos("aruiz")
	r.Equal(0, i)
	i = l.GetPos(nodeA)
	r.Equal(1, i)
	i = l.GetPos(nodeC)
	r.Equal(4, i)
	i = l.GetPos(nodeB)
	r.Equal(3, i)
}
