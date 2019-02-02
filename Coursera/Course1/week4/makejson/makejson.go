package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	fmt.Println("Please input a name")
	fmt.Scanln(&name)
	var address string
	fmt.Println("Please input an address")
	fmt.Scanln(&address)
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address
	var out []byte
	out, _ = json.Marshal(m)
	fmt.Println(string(out))
}
