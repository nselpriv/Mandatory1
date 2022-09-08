package main

import (
	"fmt"
	"time"

)

type fork struct {
	id   int
	free chan bool
}
type guy struct {
	id                  int
	forkleft, forkright *fork
}

func (g guy) eat() {
	number := g.id


	//if checkbooking(g.forkleft.free) {
		if(1==1){
		fmt.Println(fmt.Sprint("guy number ", number, " is eating"))
		
		time.Sleep(time.Second)
	} else {
		fmt.Println(fmt.Sprint("guy number ", number, " is thinking"))
		time.Sleep(time.Second)
		return
	}
	fmt.Println(fmt.Sprint("guy number ", number, " is done eating"))
}

func main() {
	//ch1 := make(chan int, 2)

	makeStructures()
	time.Sleep(time.Second)

}
func checkstatus(c chan bool){
	booking := <-c
	fmt.Println(booking)
}

func checkbooking(c chan bool) bool {
	booking := <-c
	return booking
}

func makeStructures() {
	//making 5 forks with unique ID's and a boolean if they're taking or not
	forks := make([]*fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &fork{
			id: i, free: make(chan bool),
		}
		//c := forks[1].free

		//c <- true
		//forks[1].free <- true
		/*
		for _, v := range forks {
			v.free <- true
		}
		*/
		//making 5 guys with with unique ID's and with their left and right forks
		guys := make([]*guy, 5)
		for i := 0; i < 5; i++ {
			guys[i] = &guy{
				id: i, forkleft: forks[i], forkright: forks[(i+1)%5]}
		}
		//makes sure every guy eats 3 times

		for _, v := range guys {
			go v.eat()
		}
	}
	//go checkstatus(forks[1].free)
}
