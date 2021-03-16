package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type List struct {
	length int
	head   *Node
}

func NewDoubleList() *List {
	return &List{}
}

func (l *List) print() {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		fmt.Printf("Node: %d\n", node.value)
	}
}
func (l *List) reversePrint() {
	var node *Node
	for node = l.peekLast(); node != nil; node = node.prev {
		fmt.Printf("Node: %d\n", node.value)
	}
}
func (l *List) peekLast() *Node {
	var node *Node
	if l.length == 0 {
		fmt.Println("No Nodes in List")
		return nil
	}
	if l.length == 1 {
		return l.head
	}
	for node = l.head; node != nil; node = node.next {
		if node.next == nil {
			break
		}
	}
	return node

}
func (l *List) addToHead(value int) {
	node := &Node{value: value}
	if l.head != nil {
		node.next = l.head
		l.head.prev = node
	}
	l.head = node
	l.length++
}
func (l *List) push(value int) {
	node := &Node{value: value}
	last := l.peekLast()
	last.next = node
	node.prev = last
}

func main() {
	list := NewDoubleList()

	list.addToHead(2)
	list.addToHead(10)
	fmt.Printf("Head Value %d \n", list.head.value)
	list.push(1)
	list.peekLast()
	list.print()
	list.reversePrint()
}
