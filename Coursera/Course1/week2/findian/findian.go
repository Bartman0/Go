package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ""
	fmt.Printf("input string to be searched: ")
	fmt.Scan(&input)
	if strings.HasPrefix(strings.ToLower(input), "i") &&
		strings.HasSuffix(strings.ToLower(input), "n") &&
		strings.Contains(strings.ToLower(input), "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not found!\n")
	}
}
