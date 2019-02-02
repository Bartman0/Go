package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

type cow struct {
	name                    string
	food, locomotion, noise string
}

type bird struct {
	name                    string
	food, locomotion, noise string
}

type snake struct {
	name                    string
	food, locomotion, noise string
}

func main() {
	animals := make(map[string]animal, 0)

	fmt.Println("type a request with newanimal/query name type:")

	for {
		fmt.Print("> ")
		cmd, name, req := readStrings()
		switch cmd {
		case "newanimal":
			switch req {
			case "cow":
				animals[name] = cow{name, "grass", "walk", "moo"}
			case "bird":
				animals[name] = bird{name, "worms", "fly", "peep"}
			case "snake":
				animals[name] = snake{name, "mice", "slither", "hsss"}
			}
			fmt.Println("Created it!")
		case "query":
			a, err := animals[name]
			if a == nil {
				fmt.Println(err)
				continue
			}
			switch req {
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

func readStrings() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimSpace(line)

	strings := strings.Split(line, " ")

	return strings[0], strings[1], strings[2]
}

func (a cow) Name() string {
	return a.name
}

func (a bird) Name() string {
	return a.name
}

func (a snake) Name() string {
	return a.name
}

func (a cow) Eat() {
	fmt.Println(a.name, "eats:", a.food)
}

func (a bird) Eat() {
	fmt.Println(a.name, "eats:", a.food)
}

func (a snake) Eat() {
	fmt.Println(a.name, "eats:", a.food)
}

func (a cow) Move() {
	fmt.Println(a.name, "moves by:", a.locomotion)
}

func (a bird) Move() {
	fmt.Println(a.name, "moves by:", a.locomotion)
}

func (a snake) Move() {
	fmt.Println(a.name, "moves by:", a.locomotion)
}

func (a cow) Speak() {
	fmt.Println(a.name, "speaks with:", a.noise)
}

func (a bird) Speak() {
	fmt.Println(a.name, "speaks with:", a.noise)
}

func (a snake) Speak() {
	fmt.Println(a.name, "speaks with:", a.noise)
}
