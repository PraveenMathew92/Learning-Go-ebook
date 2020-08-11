package main

import "fmt"

type Element interface {
	mapper() Element
}

type intSlice [] int
type stringSlice [] string

func (slice *intSlice) mapper (transformer func (int) int) intSlice {
	var result intSlice
	for _, v := range *slice {
		result = append(result, transformer(v))
	}
	return result
}

func (slice *stringSlice) mapper (transformer func (string) string) stringSlice {
	var result stringSlice
	for _, v := range *slice {
		result = append(result, transformer(v))
	}
	return result
}

func main() {
	var a intSlice = []int{3, 4, 5, 6, 5, 7, 3}
	b := a.mapper(func (x int) int { return x*2 })
	fmt.Println(b)
	var c stringSlice = []string{"3", "4", "5", "6", "5", "7", "3"}
	d := c.mapper(func (x string) string { return x + "0" })
	fmt.Println(d)
}