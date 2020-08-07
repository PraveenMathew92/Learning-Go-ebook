package stack

import "testing"

func TestShouldPopAfterPush (t *testing.T) {
	stack := Stack{elements: make([] byte, 0), count: 0}
	Push(&stack, 5)
	t.Log("Pushed 5 to stack")
	Pop(&stack)
	t.Log("Poped from stack")
}