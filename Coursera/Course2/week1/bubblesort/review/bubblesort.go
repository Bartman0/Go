package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ints := readInts()
	BubbleSort(ints)
	fmt.Println(ints)
}

func readInts() []int {
	fmt.Print("Please write up to 10 integers separated by a space: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimSpace(line)

	intsAsStrings := strings.Split(line, " ")
	ints := make([]int, 0)

	for _, v := range intsAsStrings {
		value, _ := strconv.Atoi(v)
		ints = append(ints, value)
	}

	return ints
}

//BubbleSort ...
func BubbleSort(original []int) {
	swapped := true
	for swapped {
		swapped = false
		for k := range original {
			if k+1 < len(original) && original[k+1] < original[k] {
				Swap(original, k)
				swapped = true
			}
		}
	}
}

//Swap ...
func Swap(original []int, index int) {
	tmp := original[index]
	original[index] = original[index+1]
	original[index+1] = tmp
}
