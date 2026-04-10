package main

import "fmt"

// node is one element on the list
type Node struct {
	value int // data of the node
	next *Node // pointer to the next node
}

// linkedlist, the whole list
type LinkedList struct {
	head *Node // pointer to the head node where the list starts
	tail *Node // pointer to the tail node where the list ends
}

// create linkedlist & add value to the list
func (l *LinkedList) Add(value int) {

	// create a new node
	newNode := &Node{
		value: value,
		next: nil, // nil for now
	}

	// if the list is empty
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	// connect the new node to the last node 
	l.tail.next = newNode

	// update the tail to the latest added node
	l.tail = newNode

}

func (l *LinkedList) RemoveFirst() {
	// if the linklist is empty
	if l.head == nil {
		return
	}
	
	// move the head pointer to the next node
	l.head = l.head.next

	// if there are no further nodes
	if l.head == nil {
		l.tail = nil
	}
}

// To print the linked list
func (l *LinkedList) Print() {
	current := l.head

	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}

	// for i := range l {
	// 	fmt.Print(l.current.value, " -> ")
	// 	current = l.current.next
	// }

	fmt.Println("nil")
}

func (l *LinkedList) Removal() {
	for l.head != nil {
		l.RemoveFirst()
		l.Print()
	}
	fmt.Print("Link list now empty")
}

func main() {

	// instance of the linked list itself
	// list := LinkedList{}
	var list LinkedList

	list.Add(3)
	list.Add(8)
	list.Add(4)
	list.Add(2)
	list.Add(6)
	list.Add(1)

	// list.Print()
	list.Removal()
}

