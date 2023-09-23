package circularlinkedlist

import "fmt"

type Node struct {
	Data int
	next *Node
}

type CircularLinkedList struct {
	tail *Node
	len  int
}

func NewNode(data int) *Node {
	return &Node{Data: data, next: nil}
}

func (c *CircularLinkedList) Init() *CircularLinkedList {
	c.tail = nil
	c.len = 0
	return c
}

func New() *CircularLinkedList {
	return new(CircularLinkedList).Init()
}

func (c *CircularLinkedList) Prepend(data int) {
	newNode := NewNode(data)
	if c.tail == nil {
		c.tail = newNode
		newNode.next = newNode
		c.len++
		return
	}
	newNode.next = c.tail.next
	c.tail.next = newNode
	c.len++
}

func (c *CircularLinkedList) Append(data int) {
	newNode := NewNode(data)
	if c.tail == nil {
		c.tail = newNode
		newNode.next = newNode
		c.len++
		return
	}
	newNode.next = c.tail.next
	c.tail.next = newNode
	c.tail = newNode
	c.len++
}

func (c *CircularLinkedList) InsertAt(data, pos int) {
	if pos <= 0 || pos > c.len+1 {
		return
	}

	if pos == 1 {
		c.Prepend(data)
		return
	}

	if pos == c.len+1 {
		c.Append(data)
		return
	}

	current := c.tail.next
	for i := 1; i < pos-1; i++ {
		current = current.next
	}
	newNode := NewNode(data)
	newNode.next = current.next
	current.next = newNode
	c.len++
}

func (c *CircularLinkedList) DeleteHead() {
	if c.len == 0 {
		return
	}

	if c.len == 1 {
		c.tail = nil
		c.len--
		return
	}

	c.tail.next = c.tail.next.next
	c.len--
	return
}

func (c *CircularLinkedList) DeleteTail() {
	if c.len == 0 {
		return
	}

	if c.len == 1 {
		c.tail = nil
		c.len--
		return
	}

	current := c.tail.next
	for current.next != c.tail {
		current = current.next
	}
	current.next = c.tail.next
	c.tail = current
	c.len--
}

func (c *CircularLinkedList) DeleteAt(pos int) {
	if c.tail == nil || pos > c.len || pos < 1 {
		return
	}

	if pos == 1 {
		c.DeleteHead()
		return
	}

	if pos == c.len {
		c.DeleteTail()
		return
	}

	current := c.tail.next
	for i := 1; i < pos-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	c.len--
}

func (c *CircularLinkedList) Print() string {
	var result string
	if c.tail == nil {
		result = "Circular linked list is empty"
		return result
	}
	head := c.tail.next
	for head != c.tail {
		result += fmt.Sprintf("%d -> ", head.Data)
		head = head.next
	}
	result += fmt.Sprintf("%d", head.Data)
	return result
}

func (c *CircularLinkedList) Len() int { return c.len }

func (c *CircularLinkedList) Head() *Node {
	if c.len == 0 {
		return nil
	}
	return c.tail.next
}

func (c *CircularLinkedList) Tail() *Node {
	if c.len == 0 {
		return nil
	}
	return c.tail
}
