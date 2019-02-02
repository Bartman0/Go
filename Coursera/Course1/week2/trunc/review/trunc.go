package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var input_str string
var input_float float64
var truncated string
var e error

func main() {
	for {
		fmt.Printf("Please, enter a valid floating number ot hit `Ctrl+C` to quit.\r\n")
		fmt.Scan(&input_str)
		input_float, e = strconv.ParseFloat(input_str, 32)
		reflect.TypeOf(e)
		if e != nil {
			fmt.Printf("%s seems not to be a valid float.\r\n", input_str)
			continue
		}
		truncated = strings.Split(input_str, ".")[0]
		if truncated == "" {
			truncated = "0" // if there is no zero prefix, presume it
		}
		fmt.Printf("%s\r\n", truncated)
	}
}
