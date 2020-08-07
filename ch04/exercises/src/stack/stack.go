package stack

type Stack struct {
	count int
	elements [] byte
}

func Push(stack *Stack, element byte) {
	stack.count++
	stack.elements = append(stack.elements, element)
}

func Pop(stack *Stack) {
	if stack.count > 0 {
		stack.count--
		stack.elements = stack.elements[:(stack.count)]
	}
}