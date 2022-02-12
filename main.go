package main

import (
	"DP-DS/LinkedList"
	"DP-DS/Stack"
	"DP-DS/Que"
	"DP-DS/BinaryTree"
	"fmt"
)

func main() {
	//callLinkedList()
	//callStack()
	//callQue()
	callTree()
}


func callTree(){
	t := &BinaryTree.Tree{}
	v := t.Insert(10).Insert(20).Insert(3).Insert(12).Insert(15).Insert(25).Insert(23)
	fmt.Printf("%+v",v.Node)

}
func callQue(){
	q := &Que.Que{ Item: make(chan int, 16),}
	q.EnQue(1)
	q.EnQue(2)
	fmt.Printf("%+v",q.Item)
//	fmt.Printf("%+v",q.DeQue)
	println(q.DeQue())
}

func callLinkedList() {
	l := &LinkedList.List{}

	l.Push(1)
	l.Push(81)
	l.Push(21)
	l.Push(31)

	//n := l.First()
	//fmt.Println(n.value)
	//n := l.First()
	// for {
	// 	println(n.value)
	// 	n = n.Next()
	// 	if n == nil {
	// 		break
	// 	}
	// }
	v, err := l.GetValue(21)
	if err != nil {
		println("--")
		fmt.Println(err)
	}
	println(v)
}

func callStack() {
	s := Stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	i, err := s.Pop()
	i, err = s.Pop()
	i, err = s.Pop()
	i, err = s.Pop()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
