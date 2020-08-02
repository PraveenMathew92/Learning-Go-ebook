package main

import "fmt"

func main() {
	slice := []float64 {1, 2, 3, 4, 5, 6, 7, 8, 9}
	avg := average(slice)
	fmt.Print(avg)
}

func average(slice []float64) float64 {
	length := len(slice)
	var sum float64 = 0
	for _, value := range slice {
		sum += value
	}
	return sum/float64(length);
}