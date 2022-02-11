package main

import (
	"DP-DS/LinkedList"
	"DP-DS/Stack"
	"fmt"
)

func main() {
	//callLinkedList()
	callStack()
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
