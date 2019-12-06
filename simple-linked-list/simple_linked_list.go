package linkedlist

import "errors"

// Element represents a Node of the linked list.
type Element struct {
	data int
	next *Element
}

// List represents a linked list.
type List struct {
	head *Element
	size int
}

// New creates a new linked list.
func New(input []int) *List {
	l := &List{}
	var curr, prev *Element

	if len(input) == 0 {
		return l
	}

	prev = &Element{data: input[0]}
	l.head = prev
	l.size = len(input)

	for i := range input {
		if i == 0 {
			continue
		}

		curr = &Element{data: input[i]}
		prev.next = curr

		prev = curr
	}

	return l
}

// Size returns the length of the linked list.
func (l *List) Size() int {
	return l.size
}

// Push adds new Element at the end of the linked list.
func (l *List) Push(elem int) {
	newElem := &Element{data: elem}

	if l.size == 0 {
		l.head = newElem
		l.size = 1
		return
	}

	e := l.head
	for e.next != nil {
		e = e.next
	}

	e.next = newElem
	l.size++
}

// Pop removes Element from the end of the linked list.
func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("empty list")
	}

	if l.size == 1 {
		t := l.head
		l.head = nil
		l.size = 0
		return t.data, nil
	}

	e := l.head
	for i := 0; i < l.size-2; i++ {
		e = e.next
	}

	t := e.next
	e.next = nil
	l.size--

	return t.data, nil
}

// Array returns an array from the given linked list.
func (l *List) Array() []int {
	arr := make([]int, 0)
	if l.size == 0 {
		return arr
	}

	e := l.head
	arr = append(arr, e.data)
	for e.next != nil {
		e = e.next
		arr = append(arr, e.data)
	}

	return arr
}

// Reverse returns a new reversed linked list.
func (l *List) Reverse() *List {
	arr := l.Array()

	// revert array
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	rl := New(arr)
	return rl
}
