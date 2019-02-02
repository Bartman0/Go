package main

import (
	"fmt"
	"sync"
	"time"
)

type host struct {
	tickets chan int
}

type chopStick struct {
	sema chan int
}

type philosopher struct {
	num         int
	host        *host
	left, right *chopStick
}

func main() {
	host := new(host)

	host.serve() // start serving diner

	chopSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(chopStick)
		chopSticks[i].init()
	}
	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{i + 1, host, chopSticks[i], chopSticks[(i+1)%5]}
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg)
	}
	wg.Wait()
}

func (p *philosopher) eat(wg *sync.WaitGroup) {
	fmt.Printf("getting ready for diner %d\n", p.num)
	for i := 0; i < 3; i++ {
		// get a host ticket first
		p.host.ask()
		select { // try to get chopstick tokens
		case <-p.left.sema:
			p.right.pickUp()
		case <-p.right.sema:
			p.left.pickUp()
		default:
			panic("starving here...")
		}

		fmt.Printf("starting to eat %d\n", p.num)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("finished eating %d\n", p.num)

		// give the tokens back for the chopsticks
		p.right.putDown()
		p.left.putDown()
		// give back the host ticket
		p.host.give()
	}
	wg.Done()
}

func (h *host) serve() {
	h.tickets = make(chan int, 2)
	// put two meal tickets on the host
	h.give()
	h.give()
	// for {
	// 	time.Sleep(10 * time.Millisecond)
	// }
}

// ask permission to eat
func (h *host) ask() {
	<-h.tickets
}

func (h *host) give() {
	h.tickets <- 1
}

func (c *chopStick) init() {
	c.sema = make(chan int, 1)
	c.putDown()
	// for {
	// 	// time.Sleep(10 * time.Millisecond)
	// }
}

func (c *chopStick) pickUp() {
	<-c.sema
}

func (c *chopStick) putDown() {
	c.sema <- 1
}
