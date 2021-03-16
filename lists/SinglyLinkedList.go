package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head   *Node
	length int
}

func NewList() *List {
	return &List{}
}

/**
* Prints Head
 */
func (l *List) printHead() {
	node := l.head
	if node != nil {
		fmt.Printf("Head Node - %+v ", node.val)
	} else {
		fmt.Println("No head")
	}
}

func (l *List) print() {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}

/**
* Returns the last Node
 */
func (l *List) peekLast() *Node {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		if node.next == nil {
			break
		}
	}
	return node
}

/**
* Replaces the Head of a List
 */
func (l *List) addHead(val int) {
	node := &Node{val: val}
	if l.head != nil {
		node.next = l.head
	}
	l.head = node
	l.length++
}

/*
* Pushes Node into end of list.
 */
func (l *List) pushNode(val int) {
	node := l.peekLast()
	node.next = &Node{val: val}
	l.length++
}

/**
* Pops Node from end of list
 */
func (l *List) popNode() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	var node *Node
	for node = l.head; node != nil; node = node.next {
		if node.next.next == nil {
			node.next = nil
			l.length--
		}
	}
}

func main() {
	list := NewList()
	list.print()
	list.addHead(0)
	list.pushNode(2)
	fmt.Printf("List Length %d\n", list.length)
	list.print()
	list.addHead(1)
	list.printHead()
	list.peekLast()
	list.popNode()
	list.peekLast()
}
