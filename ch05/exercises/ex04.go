package main

import (
	"fmt"
	"os"
	"bufio"
)

func main () {
	path := os.Args[len(os.Args) - 1]
	shouldPrintLines := os.Args[1] == "-n"
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if shouldPrintLines {
		lineNumber := 1
		for scanner.Scan() {
			lineNumber+=1
			fmt.Printf("%d %s\n", lineNumber, scanner.Text())
		}
	} else {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}