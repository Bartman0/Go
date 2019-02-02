package main

/*
Starting two goroutines which use the same variable, thus share data, will result in a race condition. 
The retrieved value may be based on the original value or that value as already increased in the 'other' thread. 

Experiments show that the race conditions occurs: 
$ while true; do ./race >> race.log; done 
$ wc -l race.log 6704 race.log 
$ grep 1 race.log | wc -l 2442 
$ grep 2 race.log | wc -l 4262 

More '2's than '1's are being generated but not in the ratio one might expect when taking the race condition into account.
*/

import (
	"fmt"
	"sync"
)

func inc(wg *sync.WaitGroup, i *int) {
	*i = *i + 1
	fmt.Println(*i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var i = 0

	wg.Add(2)
	go inc(&wg, &i)
	go inc(&wg, &i)
	wg.Wait()
}
