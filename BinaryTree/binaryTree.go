package BinaryTree

import "fmt"
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
func PrintNode(n *Node){
	if n == nil{
		return
	}
	fmt.Printf("%+v, ",n.Value)
	// for {
	// 	if n.Left != nil{
	// 		PrintNode(n.Left)
	// 		n.Left = n.Left.Left
	// 	} else{
	// 		break
	// 	}
	// }
	// for {
	// 	if n.Right != nil{
	// 		PrintNode(n.Right)
	// 		n.Right = n.Right.Right
	// 	} else{
	// 		break
	// 	}
	// }
	PrintNode(n.Left)
	PrintNode(n.Right)
	
}

func (n *Node) Exists(value int) bool{
	if n == nil{
		println("Here")
		return false
	} 
	if n.Value == value{
		return true
	} 
	if value <= n.Value{
		return n.Left.Exists(value)
	} else{
		return n.Right.Exists(value)
	}

}