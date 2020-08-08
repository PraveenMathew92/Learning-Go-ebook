package calculator

import (
	"stack"
	"fmt"
)

func main() {
	stack := stack.Stack{count: 3, elements:([] byte{'3', '5', '+'})}
}

func calculate(stack *stack.Stack) {
	operator := stack.Pop(stack)
	operand2 := int(stack.Pop(stack))
	operand1 := int(stack.Pop(stack))

	switch operator {
	case '+':
		fmt.Println(operand1 + operand2)
	}
}
