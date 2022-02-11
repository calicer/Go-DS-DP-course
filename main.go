package main

import (
	"DP-DS/LinkedList"
	"fmt"
)

func main() {
	callLinkedList()
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
