package main

import "fmt"

func main() {
	p := func(n int) int {
		return n + 2
	}
	fmt.Printf("%v\n", p(2))
}