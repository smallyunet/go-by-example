package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

func sum(l1 *Node, l2 *Node) int {
	a := ""
	for l1 != nil {
		a += strconv.Itoa(l1.Value)
		l1 = l1.Next
	}
	b := ""
	for l2 != nil {
		b += strconv.Itoa(l2.Value)
		l2 = l2.Next
	}
	fmt.Println(a, "-", b)
	return 0

	//strconv.Atoi(a) + strconv.Atoi(b)

	//46
	//"46"
	//"4" "6"
	//4 6
	//l3root := 4 -> 6 -> nil
}

func main() {
	L1 := Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next:  nil,
		},
	}
	L2 := Node{
		Value: 3,
		Next: &Node{
			Value: 4,
			Next:  nil,
		},
	}
	sum(&L1, &L2)
}
