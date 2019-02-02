package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {}
type Bird struct {}
type Snake struct {}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var neworq, name, actionstr string
	m := make(map[string]string)

	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s %s", &neworq, &name, &actionstr)
		switch neworq {
		case "newanimal":
			if actionstr == "cow" || actionstr == "snake" || actionstr == "bird" {
				m[name] = actionstr
			} else {
				fmt.Printf("Not supported animal %s\n", actionstr)
				continue
			}
		case "query":
			na, found := m[name]
			if found {
			        var animal Animal
				switch na {
				case "cow":
					animal = Cow{}
				case "snake":
					animal = Snake{}
				case "bird":
					animal = Bird{}
				default:
					fmt.Printf("%s is not supported, retry\n", na)
					continue
				}
				switch actionstr {
				case "eat":
					animal.Eat()
				case "speak":
					animal.Speak()
				case "move":
					animal.Move()
				default:
					fmt.Printf("%s is not supported, retry\n", actionstr)
					continue
				}
			} else {
				fmt.Printf("%s not found, retry\n", name)
				continue
			}
		default:
			fmt.Printf("%s is not supported, retry\n", neworq)
			continue
		}
	}
}
