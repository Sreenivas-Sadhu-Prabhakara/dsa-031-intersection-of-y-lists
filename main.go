package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// Intersection returns the first node shared by both lists (by identity, not
// value), or nil. Two pointers walk both lists; each one switches to the
// other list's head when it runs out, equalizing the path lengths.
func Intersection(a, b *Node) *Node {
	p, q := a, b
	for p != q {
		if p == nil {
			p = b
		} else {
			p = p.Next
		}
		if q == nil {
			q = a
		} else {
			q = q.Next
		}
	}
	return p
}

func main() {
	shared := &Node{Val: 8, Next: &Node{Val: 10}}
	a := &Node{Val: 3, Next: shared}
	b := &Node{Val: 99, Next: &Node{Val: 1, Next: shared}}
	fmt.Println(Intersection(a, b).Val)
}
