package main

import "fmt"

type stack struct {
	count int
	elements [] int
}

func main() {
	var stack = stack { count:0, elements: make([] int,0)}
	fmt.Println(stack.count)
	push(&stack, 5)
	push(&stack, 3)
	push(&stack, 4)
	pop(&stack)
	fmt.Println(stack.elements)

}

func push (stack *stack, value int) {
	stack.count++
	stack.elements = append(stack.elements, value)
}

func pop (stack *stack) {
	stack.count--
	stack.elements = stack.elements[:(stack.count)]
}