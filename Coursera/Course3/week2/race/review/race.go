/*

This program has two goroutines, the main routine and the anonymous function.
The main routine triggers the another routine and starts the loop. Inside the loop of
the main routine, the routine prints the value of x. However, as the other routine also
references the same variable x (increments it), there is no deterministic way of knowing 
what value the x will contain when the main routine prints it. It depends on the order 
how they are executed -> race condition. Run it a few times and see. 

*/

package main

import (
	"fmt"
	"time"
)

func main() {

	var x int = 0

	counter := 10

	go func(x *int, counter int) {
		for i:= 0; i < counter; i++ {
			*x = *x + 1
			time.Sleep(1 * time.Second)
		}
	}(&x, counter)

	for i:= 0; i < counter; i++ {
            fmt.Println(x)
	    time.Sleep(1 * time.Second)
        }
}
