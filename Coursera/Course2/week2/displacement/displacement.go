package main

import "fmt"

func main() {
	var n int

	var acceleration float32
	n = 0
	for n == 0 {
		fmt.Println("Please input an acceleration: ")
		n, _ = fmt.Scanln(&acceleration)
	}

	var initialVelocity float32
	n = 0
	for n == 0 {
		fmt.Println("Please input an initial velocity: ")
		n, _ = fmt.Scanln(&initialVelocity)
	}

	var initialDisplacement float32
	n = 0
	for n == 0 {
		fmt.Println("Please input an initial displacement: ")
		n, _ = fmt.Scanln(&initialDisplacement)
	}

	var time float32
	n = 0
	for n == 0 {
		fmt.Println("Please input a time: ")
		n, _ = fmt.Scanln(&time)
	}

	fn := genDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println(fn(time))
}

func genDisplaceFn(acceleration, initialVelocity, initialDisplacement float32) func(time float32) float32 {
	return func(time float32) float32 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}
