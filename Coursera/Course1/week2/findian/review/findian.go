package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Printf("Enter a string:")
	fmt.Scan(&str)

	if strings.ContainsAny(str, "ianIAN") {
		str = strings.Replace(str, "I", "i", -1)
		str = strings.Replace(str, "A", "a", -1)
		str = strings.Replace(str, "N", "n", -1)
	}
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
