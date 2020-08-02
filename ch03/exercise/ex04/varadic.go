package main

import "fmt"

func main() {
	printNumbers(4, 5, 6, 2, 3, 5)
}

func printNumbers (args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}