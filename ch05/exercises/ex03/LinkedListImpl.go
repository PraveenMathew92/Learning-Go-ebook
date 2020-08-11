package main

type Element struct {
	Value int
	next *Element
}

type List struct {
	front *Element
	tail *Element
	len int
}

func (element *Element) Next() *Element {
	return element.next
}

func (list *List) New() *List {
	return new(List)
}

func (list *List) PushBack(value int) *Element {
	element := Element{value, list.front}
	list.tail.next = &element
	list.tail = &element
	list.len++
	return &element
}

func (list *List) Len() int {
	return list.len
}

func (list *List) Front() *Element {
	return list.front
}