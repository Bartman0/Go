package main

import (
	"fmt"
	"os"
)

func main() {
	type Fullname struct {
		fname string
		lname string
	}

	var name string
	fmt.Println("Please input a filename")
	fmt.Scanln(&name)
	file, err := os.Open(name)
	defer file.Close()
	var full Fullname
	fullnames := make([]Fullname, 0)
	itemCount, err := 1, nil
	for {
		itemCount, err = fmt.Fscanf(file, "%s %s\n", &full.fname, &full.lname)
		if itemCount > 0 && err == nil {
			fullnames = append(fullnames, full)
		} else {
			break
		}
	}
	for _, full = range fullnames {
		fmt.Println(full.fname, full.lname)
	}
}
