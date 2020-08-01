package main

import "fmt"

func main() {
	var i = 1
	BEGIN:
		fmt.Println(i)
		i++
		if i <= 10 {
			goto BEGIN
		}
}