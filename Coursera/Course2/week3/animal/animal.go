package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	name                    string
	food, locomotion, noise string
}

func main() {
	var animals = [...]animal{
		animal{"cow", "grass", "walk", "moo"},
		animal{"bird", "worms", "fly", "peep"},
		animal{"snake", "mice", "slither", "hsss"},
	}

	fmt.Println("type a request with <animal>(cow, bird, snake) <request>(eat, move, speak):")

	for {
		fmt.Print("> ")
		anim, request := readStrings()
		for _, a := range animals {
			if a.name == anim {
				switch request {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			}
		}
	}
}

func readStrings() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimSpace(line)

	strings := strings.Split(line, " ")

	return strings[0], strings[1]
}

func (a animal) Eat() {
	fmt.Println(a.name, "eats:", a.food)
}

func (a animal) Move() {
	fmt.Println(a.name, "moves by:", a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.name, "speaks with:", a.noise)
}
