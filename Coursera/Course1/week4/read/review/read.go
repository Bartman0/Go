package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	type name struct {
		fname, lname string
	}
	names := []name{}

	var fileName string
	fmt.Println("Enter the file name?")
	fmt.Scan(&fileName)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err.Error())
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if len(line) > 0 {
			fields := strings.Fields(line)
			names = append(names, name{fields[0], fields[1]})
		}
	}

	for _, outName := range names {
		fmt.Printf("First Name is %s, Last Name is %s \n", outName.fname, outName.lname)
	}

}
