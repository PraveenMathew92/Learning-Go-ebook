package main

import (
	"fmt"
	"container/list"
)

func main() {
	list := list.New()
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