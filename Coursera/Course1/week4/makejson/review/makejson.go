package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Printf("Enter the name: ")
	fmt.Scan(&name)
	fmt.Printf("Enter the address: ")
	fmt.Scan(&address)

	person := make(map[string]string)

	person["name"] = name
	person["address"] = address

	j, err := json.Marshal(person)

	if err != nil {
		fmt.Printf("ERROR")
	} else {
		fmt.Println(j)
	}

}
