package main

import "fmt"

func main() {
	slice := []int {1, 2, 3, 4, 5, 6, 7, 7, 8, 9}
	doubler := func(x int) int {
		return 2*x
	}
	fmt.Println(mapFunc(doubler, slice))
}

func mapFunc(f (func(int) int), list [] int) []int {
	outputList := make([] int, len(list))
	for i := range list {
		outputList [i] = f(list[i])
	}
	return outputList
}