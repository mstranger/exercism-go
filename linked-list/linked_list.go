package linkedlist

// package main

import "fmt"

var ErrEmptyList error = fmt.Errorf("empty list")

// Element represents a Node for the linked list.
type Element struct {
	Val  interface{}
	prev *Element
	next *Element
}

type List struct {
	first *Element
	last  *Element
}

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

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (l *List) First() *Element {
	return l.first
}

func (l *List) Last() *Element {
	return l.last
}

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

func (l *List) Empty() bool {
	return l.first == nil && l.last == nil
}

// Print displays a list.
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

// func main() {
// 	list := NewList(1, 2, 3, 4, 5)
// 	list.Reverse()
// 	list.Print()
// 	// 	// 	list.PushFront(0)
// 	// 	// 	list.PushBack(5)
// 	// 	// 	fmt.Printf("first: %+v\n", list.first)
// 	// 	// 	fmt.Printf("last: %+v\n", list.last)
// 	// 	// 	fmt.Printf("after first: %+v\n", list.first.Next())
// 	// 	// 	fmt.Printf("before last: %+v\n", list.last.Prev())

// 	// 	// 	list.Print()
// 	// 	// 	list.Reverse().Print()

// 	// 	list := NewList(3, 4)
// 	// 	v, _ := list.PopBack()
// 	// 	fmt.Println(v)
// 	// 	v, _ = list.PopBack()
// 	// 	fmt.Println(v)
// 	// 	v, err := list.PopBack()
// 	// 	fmt.Println(v, err)
// 	// 	list.Print()
// }
