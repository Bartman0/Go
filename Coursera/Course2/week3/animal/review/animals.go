package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Printf("%s", a.food)
}

func (a Animal) Move() {
	fmt.Printf("%s", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("%s", a.noise)
}

func main() {
	animalMap := make(map[string]Animal)
	animalMap["cow"] = Animal{"grass", "walk", "moo"}
	animalMap["bird"] = Animal{"worms", "fly", "peep"}
	animalMap["snake"] = Animal{"mice", "slither", "hsss"}

	functionMap := make(map[string]func (a Animal))
	functionMap["eat"] = Animal.Eat
	functionMap["move"] = Animal.Move
	functionMap["speak"] = Animal.Speak

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		line = strings.TrimSuffix(line, "\n")
		lineSplice := strings.Fields(line)
		if len(lineSplice) != 2 {
			fmt.Printf("%s%d.\n", "Expected 2 strings per line, got: ", len(lineSplice))
			continue
		}
		creature, found := animalMap[lineSplice[0]]
		if !found {
			fmt.Printf("%s%s\n", "Received non-supported animal: ", lineSplice[0])
			continue
		}

		function, found := functionMap[lineSplice[1]]
		if !found {
			fmt.Printf("%s%s\n", "Received non-supported function: ", lineSplice[1])
                        continue
                }

		function(creature)
		fmt.Printf("\n")
	}
}
