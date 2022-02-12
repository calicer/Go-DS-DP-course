package BinaryTree

type Tree struct {
	Node *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.Node == nil {
		t.Node = &Node{Value: value}
	} else {
		t.Node.insert(value)
	}
	return t
}

func (n *Node) insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.insert(value)
		}

	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.insert(value)
		}
	}
}
