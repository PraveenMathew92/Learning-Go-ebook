package main

import "fmt"

func main() {
	n := 10
	fiboNumbers := fibonacci(n)
	fmt.Println(fiboNumbers)
}

func fibonacci(n int) (terms [] int) {
	a := 0
	b := 1
	terms = make([] int, n)
	for i := range terms {
		terms[i] = b
		b = a + b
		a = b - a
	}
	return terms
}