// package main

import (
	"fmt"
	"time"
)

type fork struct {
	id     int
	status chan bool
}
type guy struct {
	id                  int
	forkleft, forkright *fork
}

func (g guy) eat() {
	go checkfork(g.forkright, g.forkleft, g.id)
}


func checkfork(forkyright *fork, forkyleft *fork, id int){
	statusr := <- forkyright.status
	statusl := <- forkyleft.status
	if(!statusr && statusl){
		forkyright.status <- true
		forkyleft.status <- false	
	}
	if(!statusl && statusr){
		forkyleft.status <- false
		forkyright.status <- true	
	}
	if(!statusl && !statusr){
		forkyleft.status <- false
		forkyright.status <- false	
	}
	if(statusr && statusl){
		fmt.Println(fmt.Sprint("guy number ", id, " is eating"))
		
		fmt.Println(fmt.Sprint("guy number ", id, " is full"))	
		forkyleft.status <- true
		forkyright.status <- true
	} else {
		fmt.Println(fmt.Sprint("guy number ", id, " is thinking"))
	}
}

func main() {
	begin()
	
}

func begin() {
	//making 5 forks with unique ID's and a boolean if they're taking or not
	forks := make([]*fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &fork{
			id: i, status: make(chan bool, 2),
		}
	}
	for _, v := range forks {
		v.status <- true
	}

	//making 5 guys with with unique ID's and with their left and right forks
	guys := make([]*guy, 5)
	for i := 0; i < 5; i++ {
		guys[i] = &guy{
			id: i, forkleft: forks[i], forkright: forks[(i+1)%5]}

	}
	//makes sure every guy eats 3 times

	for i := 0; i < 3; i++ {
		for _, v := range guys {
			go v.eat()
			// time.Sleep(time.Second)

		for i := 0; i<3; i++{
			for _, v := range guys {
				go v.eat()
				time.Sleep(time.Millisecond)	
				
			}
		}
	}
	time.Sleep(time.Second)
}
