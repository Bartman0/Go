package main

import "fmt"

func main() {
	input := 0.0
	fmt.Printf("input float to be truncated: ")
	fmt.Scan(&input)
	fmt.Printf("original value: %f, truncated value: %d\n", input, int64(input))
}
