package main

import (
	"fmt"
)

// s =Â½ a t2 + vot + so

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

	return func(t float64) float64 {
		return ((0.5 * a * (t * t)) + (vo * t) + so)
	}
}

func main() {

	param := [4]float64{}

	fmt.Printf(" %s\n", "Please enter acceleration ...  ")
	fmt.Scan(&param[0])
	fmt.Printf(" %s\n", "Please enter initial velocity ...  ")
	fmt.Scan(&param[1])
	fmt.Printf(" %s\n", "Please enter initial displacement ...  ")
	fmt.Scan(&param[2])
	fmt.Printf(" %s\n", "Please enter time (seconds) ...  ")
	fmt.Scan(&param[3])

	fn := GenDisplaceFn(param[0], param[1], param[2])

	fmt.Printf("\n%s%v%s  ...   %v\n\n", "Displacement after ", param[3], " seconds", fn(param[3]))

}

/*
Examine the code. If the code contains a function called GenDisplaceFn()
which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so and
returns a function, then give another 2 points.
If the function returned by GenDisplaceFn() is used to compute the time, give another 2 points.


Let us assume the following formula for displacement s as a function of time t,
acceleration a, initial velocity vo, and initial displacement so.

s =Â½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration,
initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program
should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn()
which takes three float64 arguments, acceleration a, initial velocity vo, and
initial displacement so. GenDisplaceFn() should return a function
which computes displacement as a function of time, assuming the given values acceleration,
initial velocity, and initial displacement.
The function returned by GenDisplaceFn() should take one float64 argument t,
representing time, and return one float64 argument which is the displacement travelled after time t.

For example, letâ€™s say that I want to assume the following values for acceleration,
initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function fn
which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

*/
