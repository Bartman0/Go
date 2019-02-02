package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list := []int{}
	for {
		var input string
		fmt.Println("Please input a new integer value")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		if strings.Trim(input, "+-0123456789") != "" {
			fmt.Println("Please input digits only")
			continue
		}
		value, _ := strconv.Atoi(input)
		list = append(list, value)
		sort.Ints(list)
		printSlice(list)
	}
	fmt.Println("done")
}

func printSlice(s []int) {
	fmt.Printf("sorted slice: %v\n", s)
}
