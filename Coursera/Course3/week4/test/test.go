package main

import (
	"fmt"
	"sync"
	"time"
)

type host struct {
	tickets chan int
}

func main() {
	h := new(host)
	h.tickets = make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go h.put(&wg)
	go h.get(&wg)
	wg.Wait()
}

func (h *host) put(wg *sync.WaitGroup) {
	fmt.Println("starting put")
	h.tickets <- 1
	time.Sleep(1000 * time.Millisecond)
	wg.Done()
	fmt.Println("end of put")
}

func (h *host) get(wg *sync.WaitGroup) {
	fmt.Println("starting get")
	<-h.tickets
	time.Sleep(1000 * time.Millisecond)
	wg.Done()
	fmt.Println("end of get")
}
