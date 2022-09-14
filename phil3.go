package main

import (
	"fmt"
	"math/rand"
	"time"
)

type signal struct{}

type communication struct {
	available chan signal
	pickup    chan signal
	putdown   chan signal
}

func randomTime(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func fork(c communication) {
	c.available <- signal{}
	for {
		<-c.pickup
		<-c.putdown
		c.available <- signal{}
	}
}

func philosopher(right_com communication, left_com communication, id int) {
	var times_eaten int
	for times_eaten < 3 {
		// think
		fmt.Printf("Philosopher %d thinks\n", id)
		randomTime(2)

		// wait for forks
		/*
			This if statement is to prevent a possible deadlock.
			Thinking of the forks as being in a partial order,
			each philosopher should always pick up the highest order first
			- which for all philosophers is the left fork, except for the last (4th) one.
			Here the last philosopher acts as a "tie-breaker", in the case of every other philosopher
			picks up the left fork and then waits for the right fork. Should the last philosopher
			in this case also pick up the left and wait for the right, this would result in a deadlock.
		*/
		if id == 4 {
			<-right_com.available
			right_com.pickup <- signal{}
			<-left_com.available
			left_com.pickup <- signal{}
		} else {
			<-left_com.available
			left_com.pickup <- signal{}
			<-right_com.available
			right_com.pickup <- signal{}
		}

		// eat
		times_eaten++
		fmt.Printf("Philosopher %d eats\n", id)
		randomTime(3)

		// put back forks
		left_com.putdown <- signal{}
		right_com.putdown <- signal{}
	}
	fmt.Printf("Philosopher %d is full\n", id)
	done <- signal{}
}

func count_done() {
	var count int
	for count < 5 {
		<-done
		count++
	}
	all_done <- signal{}
}

var done chan signal
var all_done chan signal

func main() {
	done = make(chan signal)
	all_done = make(chan signal)

	var communications [5]communication
	for i := 0; i < 5; i++ {
		communications[i] = communication{
			available: make(chan signal),
			pickup:    make(chan signal),
			putdown:   make(chan signal),
		}
	}

	go count_done()

	for i := 0; i < 5; i++ {
		go fork(communications[i])
	}

	for i := 0; i < 5; i++ {
		go philosopher(communications[i], communications[(i+1)%5], i)
	}

	<-all_done
	fmt.Println("All done.")
}
