package doublylinkedlist

import "fmt"

// Node represents a node in the doubly linked list.
type Node struct {
	Data       int
	next, prev *Node
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList struct {
	head, tail *Node
	len        int
}

// NewNode creates a new node with the given 
// data and a nil next and prev pointer.
func NewNode(data int) *Node {
	return &Node{Data: data, next: nil, prev: nil}
}

// NewDoublyLinkedList creates a new empty doubly linked list 
// with a nil head and tail pointer and a length of 0.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil, len: 0}
}

// Prepend adds a new node to the beginning of the doubly linked list
func (l *DoublyLinkedList) Prepend(data int) *Node {
	newNode := NewNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.len++
		return newNode
	}
	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
	l.len++
	return newNode
}

// Append adds a new node to the end of the doubly linked list
func (l *DoublyLinkedList) Append(data int) *Node {
	newNode := NewNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.len++
		return newNode
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.len++
	return newNode
}

// DeleteAt deletes the node at the specified 
// position in the doubly linked list.
func (l *DoublyLinkedList) DeleteAt(pos int) {
	if l.head == nil || pos > l.len || pos < 1 {
		return
	}
	if pos == 1 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		l.len--
		return
	}
	current := l.head
	for i := 1; i < pos-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	if current.next != nil {
		current.next.prev = current
	} else {
		l.tail = current
	}
	l.len--
}

// Display returns a string representation of the doubly linked list.
func (l *DoublyLinkedList) Display() string {
	var result string
	if l.head == nil {
		result = "Doubly linked list is empty"
		return result
	}
	result += "nil <-> "
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%d <-> ", current.Data)
		current = current.next
	}
	result += "nil"
	return result
}

// Head returns the head node of the doubly linked list.
func (l *DoublyLinkedList) Head() *Node { return l.head }

// Tail returns the tail node of the doubly linked list.
func (l *DoublyLinkedList) Tail() *Node { return l.tail }

// Len returns the length of the doubly linked list.
func (l *DoublyLinkedList) Len() int { return l.len }
