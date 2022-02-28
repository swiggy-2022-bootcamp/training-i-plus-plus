/*
 Recursive Structs := When a struct of same type difined as a property
 of the same struct.
*/

package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func (head *Node) addNodeAtStart(val int) {
	// Check for empty head
	if *head == (Node{}) {
		head.next = &Node{val: val, next: new(Node)}
		return
	} else {
		temp := &Node{val: val, next: head.next}
		head.next = temp
	}
}

func (head *Node) printLinkedList() {
	if *head == (Node{}) {
		fmt.Println("No element.")
		return
	}
	var curr *Node = head.next
	for *curr != (Node{}) {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

func (head *Node) getTail() (tail *Node) {
	tail = head.next
	for *tail.next != (Node{}) {
		tail = tail.next
	}
	return tail
}

func (head *Node) addNodeAtEnd(val int) {
	if *head == (Node{}) {
		head.next = &Node{val: val, next: new(Node)}
		return
	}
	tail := head.getTail()
	temp := &Node{val: val, next: tail.next}
	tail.next = temp
}

func main() {
	head := new(Node)
	head.printLinkedList()
	head.addNodeAtStart(8)
	head.addNodeAtStart(8)
	head.addNodeAtStart(5)
	head.addNodeAtEnd(10)
	head.printLinkedList()
}
