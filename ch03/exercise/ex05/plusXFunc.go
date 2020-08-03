package main

import "fmt"

func main() {
	plusx := func(n int, x int) int {
		return n + x
	}
	fmt.Println(plusx(5, 7))
}