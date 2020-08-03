package main

import "fmt"

func main () {
	slice := [] int {4, 5, 6, 9, 8, 3, 2, 1, 7}
	maximum := max(slice)
	fmt.Println(maximum)
}

func max (slice [] int) int {
	max := slice[0]
	for _, v := range slice[1:] {
		if (v > max) {
			max = v
		}
	}
	return max
}