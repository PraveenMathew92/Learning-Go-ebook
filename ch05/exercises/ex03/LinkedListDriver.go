package main

import (
	"fmt"
)

func main() {
	list := new(List)
	list.PushBack(1)
	list.PushBack(4)
	list.PushBack(5)
	len := list.Len()
	var element = list.Front()
	for i := 0; i < len; i++ {
		fmt.Printf("%v\n", element.Value)
		element = element.Next()
	}
}