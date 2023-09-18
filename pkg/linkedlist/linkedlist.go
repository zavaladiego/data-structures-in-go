package linkedlist

import "fmt"

// Node represents a node in the linked list.
type Node struct {
	Data int   // Data stored in the node.
	next *Node // Pointer to the next node.
}

// LinkedList represents a singly linked list.
type LinkedList struct {
	head *Node // Pointer to the head node.
	len  int   // Length of the linked list.
}

// NewNode creates a new node with the given data and a nil next pointer.
func NewNode(data int) *Node {
	return &Node{Data: data, next: nil}
}

// NewLinkedList creates a new empty linked list with a nil head pointer and a length of 0.
func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, len: 0}
}

// Prepend adds a new node to the beginning of the linked list and returns the new node.
func (l *LinkedList) Prepend(data int) *Node {
	newNode := NewNode(data)
	if l.head == nil {
		l.head = newNode
		l.len++
		return newNode
	}
	newNode.next = l.head
	l.head = newNode
	l.len++
	return newNode
}

// Append adds a new node to the end of the linked list and returns the new node.
func (l *LinkedList) Append(data int) *Node {
	newNode := NewNode(data)
	if l.head == nil {
		l.head = newNode
		l.len++
		return newNode
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	l.len++
	return newNode
}

// DeleteAt deletes the node at the specified position in the linked list.
func (l *LinkedList) DeleteAt(pos int) {
	if l.head == nil || pos > l.len || pos < 1 {
		return
	}
	if pos == 1 {
		l.head = l.head.next
		l.len--
		return
	}
	current := l.head
	for i := 1; i < pos-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	l.len--
}

// Display returns a string representation of the linked list.
func (l *LinkedList) Display() string {
	var result string
	if l.head == nil {
		result = "Linked list is empty"
		return result
	}
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%d -> ", current.Data)
		current = current.next
	}
	result += "nil"
	return result
}

// Len returns the length of the linked list.
func (l *LinkedList) Len() int { return l.len }

// Head returns the head node of the linked list.
func (l *LinkedList) Head() *Node { return l.head }
