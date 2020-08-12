package main

import (
	"fmt"
	"os"
	"bufio"
)

func main () {
	path := os.Args[len(os.Args) - 1]
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}