package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	numParts = 4
)

func partitionSort(wg *sync.WaitGroup, p []int, partition int) {
	fmt.Printf("partition %d, len %d, contents %v\n", partition, len(p), p)
	sort.Ints(p)
	wg.Done()
}

func mergeSort(p1 []int, p2 []int) []int {
	var result []int
	l1 := len(p1)
	l2 := len(p2)
	var i, j int
	for i < l1 || j < l2 {
		switch {
		case i == l1:
			result = append(result, p2[j])
			j++
		case j == l2:
			result = append(result, p1[i])
			i++
		case p1[i] < p2[j]:
			result = append(result, p1[i])
			i++
		case p1[i] >= p2[j]:
			result = append(result, p2[j])
			j++
		default:
			panic("should not be reached")
		}
	}
	return result
}

func main() {
	fmt.Println("input a list of integers separated by space; minimum number of ints = 4")
	ints, err := readInts()
	switch err {
	case -1:
		fmt.Println("error in input")
		return
	case -4:
		fmt.Println("less than 4 integers in input")
		return
	}
	l := len(ints) / numParts
	var wg sync.WaitGroup
	wg.Add(numParts)
	for i := 0; i < numParts-1; i++ {
		go partitionSort(&wg, ints[i*l:(i+1)*l], i+1)
	}
	go partitionSort(&wg, ints[(numParts-1)*l:], numParts)
	wg.Wait()
	result1 := mergeSort(ints[0:l], ints[l:2*l])
	result2 := mergeSort(ints[2*l:3*l], ints[3*l:])
	complete := mergeSort(result1, result2)
	fmt.Println("sorted result:", complete)
}

func readInts() ([]int, int) {
	var err int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimSpace(line)

	strings := strings.Split(line, " ")
	var ints []int
	for _, s := range strings {
		i, e := strconv.Atoi(s)
		if e != nil {
			return ints, -1
		}
		ints = append(ints, i)
	}

	if len(ints) < 4 {
		return ints, -4
	}

	return ints, err
}
