package LinkedList

import (
	"errors"
)

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *Node) Next() *Node {
	return l.next
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node

}

func (l *List) GetValue(i int) (int, error) {
	k := l.First()
	for {
		g := k.value
		// printlng)
		if g == i {
			return g, nil
			break
		}
		k = k.Next()
		if k.next == nil {
			return 0, errors.New("end of LL reached 404")

		}
	}

	return 0, errors.New("end of LL reached 404")
}
