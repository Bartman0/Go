package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := []int{}
	for {
		var input string
		fmt.Println("Please input a new integer value, X to start sorting")
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
	}
	bubbleSort(list)
	printSlice(list)
}

func bubbleSort(list []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := range list {
			if i == len(list)-1 {
				break
			}
			if list[i] > list[i+1] {
				swap(list, i)
				swapped = true
			}
		}
	}
}

func swap(list []int, i int) {
	tmp := list[i]
	list[i] = list[i+1]
	list[i+1] = tmp
}

func printSlice(s []int) {
	fmt.Printf("sorted slice: %v\n", s)
}
