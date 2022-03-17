package datastruct

// LinkedList
//
// - append to end is O(1)
// - finding an element is O(N)
type LinkedList struct {
	Head *LinkedListNode
	size int
}

func (l LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) Insert(v interface{}) {
	n := LinkedListNode{Value: v}
	if l.size == 0 {
		l.Head = &n
		l.size++
		return
	}

	nodeBuffer := l.Head
	for i := 0; i < l.size; i++ {
		if !nodeBuffer.HasNext() {
			nodeBuffer.Next = &n
			l.size++
			return
		}
		nodeBuffer = nodeBuffer.Next
	}
}

func (l LinkedList) GetAt(pos int) *LinkedListNode {
	if pos < 0 || pos > l.size {
		return nil
	} else if pos == 0 {
		return l.Head
	}

	nodeBuffer := l.Head
	for i := 0; i < l.size; i++ {
		if i == pos {
			return nodeBuffer
		}
		nodeBuffer = nodeBuffer.Next
	}
	return nil
}

func (l *LinkedList) InsertAt(pos int, v interface{}) {
	if pos < 0 || pos > l.size {
		return
	}

	nodeBuffer := LinkedListNode{
		Value: v,
	}
	currentNode := l.GetAt(pos)
	if pos == 0 {
		nodeBuffer.Next = l.Head // append head to second position
		l.Head = &nodeBuffer
		l.size++
		return
	}
	nodeBuffer.Next = currentNode
	previousNode := l.GetAt(pos - 1)
	previousNode.Next = &nodeBuffer
	l.size++
}

func (l LinkedList) GetPos(v interface{}) int {
	nodeBuffer := l.Head
	for i := 0; i < l.size; i++ {
		if nodeBuffer == nil {
			return -1
		} else if nodeBuffer.Value == v {
			return i
		}
		nodeBuffer = nodeBuffer.Next
	}
	return -1
}

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

func (n LinkedListNode) HasNext() bool {
	return n.Next != nil
}
