package main

import "fmt"

func main() {
	slice := [] int {4, 5, 6, 7, 3, 2, 1, 8, 9}
	sortedSlice := bubbleSort(slice)
	fmt.Println(sortedSlice)
}

func bubbleSort(s [] int) (sortedSlice [] int)  {
	length  := len(s)
	for i := length - 1; i >=0; i-- {
		for k, _ := range s[:i] {
			if (s[k] > s[k + 1]) {
				s[k], s[k + 1] = s[k + 1], s[k]
			}
		}
	}
	return s
}