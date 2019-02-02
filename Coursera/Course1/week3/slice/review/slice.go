package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	number := make([]int, 3)
	fmt.Printf("%v", number)
	var input string
	for true {
		fmt.Printf("Enter a number to add into the slice, or press x to exit the loop: ")
		fmt.Scan(&input)
		if strings.ContainsAny(input, "xX") {
			break
		}
		if strings.ContainsAny(input, "abcdefghjiklmnopqrstuvwyzABCDEFGHIJKLMNOPQRSTUVWYZ") {
			fmt.Printf("Invalid, please type a number!\n")
			continue
		} else {
			input2, _ := strconv.Atoi(input)
			number = append(number, input2)
			sort.Ints(number)
			fmt.Println(number)
		}
	}
}
