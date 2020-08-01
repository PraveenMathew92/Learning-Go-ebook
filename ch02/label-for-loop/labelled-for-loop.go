package main

import "fmt"

func main() {
	I: for i := 5; i <= 10; i++ {
		for j := 10; j>=0; j-- {
			if i >= j {
				break I
			}
			fmt.Println(i)
		}
	}
	fmt.Println("END OF EXECUTION")
}
