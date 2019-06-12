package linkedlist

import "fmt"

// ErrEmptyList contains an error for an empty list.
var ErrEmptyList error = fmt.Errorf("empty list")

// Element represents a Node for the linked list.
// It holds a value and pointers to the next and previous elements.
type Element struct {
	Val  interface{}
	prev *Element
	next *Element
}

// List represents a doubly linked list structure.
// It holds references to the first and last node.
type List struct {
	first *Element
	last  *Element
}

// NewList creates a new list from the given arguments.
func NewList(args ...interface{}) *List {
	list := List{}
	var prev *Element

	for i, v := range args {
		e := Element{Val: v}

		if list.first == nil {
			list.first = &e
		}
		if i == len(args)-1 {
			list.last = &e
		}
		if prev != nil {
			prev.next = &e
		}

		e.prev = prev
		prev = &e
	}

	return &list
}

// Next returns a pointer to the next Node (Element).
func (e *Element) Next() *Element {
	return e.next
}

// Prev returns a pointer to the previous Node (Element).
func (e *Element) Prev() *Element {
	return e.prev
}

// First returns a pointer to the first element in the list.
func (l *List) First() *Element {
	return l.first
}

// Last returns a pointer to the last element in the list.
func (l *List) Last() *Element {
	return l.last
}

// PushFront adds a new node to the top (from left) of the list.
func (l *List) PushFront(v interface{}) {
	e := Element{Val: v}
	if l.Empty() {
		l.last, l.first = &e, &e
		return
	}

	l.first.prev = &e
	e.next = l.first
	l.first = &e
}

// PushBack adds a new node to the end (from right) of the list.
func (l *List) PushBack(v interface{}) {
	e := Element{Val: v}
	if l.Empty() {
		l.last, l.first = &e, &e
		return
	}

	l.last.next = &e
	e.prev = l.last
	l.last = &e
}

// PopFront removes a node from the top of the list.
// Returns this node and error.
func (l *List) PopFront() (interface{}, error) {
	if l.Empty() {
		return 0, ErrEmptyList
	}

	v := l.first
	if v == l.last {
		l.first, l.last = nil, nil
		return v.Val, nil
	}
	l.first.next.prev = nil
	l.first = l.first.next
	return v.Val, nil
}

// PopBack removes a node from the end of the list.
// Returns this node and error.
func (l *List) PopBack() (interface{}, error) {
	if l.Empty() {
		return 0, ErrEmptyList
	}

	v := l.last
	if v == l.first {
		l.first, l.last = nil, nil
		return v.Val, nil
	}

	l.last.prev.next = nil
	l.last = l.last.prev
	return v.Val, nil
}

// Reverse flips a list in place.
func (l *List) Reverse() *List {
	if l.Empty() {
		*l = *NewList()
		return l
	}

	reversed := []interface{}{}

	e := l.last
	for e.prev != nil {
		reversed = append(reversed, e.Val)
		e = e.prev
	}
	reversed = append(reversed, e.Val)

	*l = *NewList(reversed...)
	return l
}

// Empty checks if the list is empty.
func (l *List) Empty() bool {
	return l.first == nil && l.last == nil
}

// Print displays the list.
func (l *List) Print() {
	if l.Empty() {
		fmt.Println("{ }")
		return
	}

	e := l.first
	fmt.Print("{ ")
	for e.next != nil {
		fmt.Print(e.Val, " ")
		e = e.next
	}
	fmt.Print(e.Val, " }\n")
}
