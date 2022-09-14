package main

import (
	"fmt"
	"time"

)

type fork struct {
	id   int
	status chan bool
}
type guy struct {
	id                  int
	forkleft, forkright *fork
}

func (g guy) eat() {

	statusr := <- g.forkleft.status
	statusl := <- g.forkright.status

	if(!statusr && statusl){
		g.forkleft.status <- true
		g.forkright.status <- false	
	}
	if(!statusl && statusr){
		g.forkleft.status <- false
		g.forkright.status <- true	
	}
	if(!statusl && !statusr){
		g.forkleft.status <- false
		g.forkright.status <- false	
	}
	if(statusr && statusl){
		g.forkleft.status <- false
		g.forkright.status <- false
		fmt.Println(fmt.Sprint("guy number ", g.id, " is eating"))
		time.Sleep(time.Millisecond*5)
		
		fmt.Println(fmt.Sprint("guy number ", g.id, " is full"))	
		g.forkleft.status <- true
		g.forkright.status <- true
	} else {
		fmt.Println(fmt.Sprint("guy number ", g.id, " is thinking"))
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
			id: i, status: make(chan bool,2),
		} }
		for _, v := range forks{
			v.status<-true
		}
		
		
		//making 5 guys with with unique ID's and with their left and right forks
		guys := make([]*guy, 5)
		for i := 0; i < 5; i++ {
			guys[i] = &guy{
				id: i, forkleft: forks[i], forkright: forks[(i+1)%5]}
				
		}
		//makes sure every guy eats 3 times

		for i := 0; i<3; i++{
			for _, v := range guys {
				go v.eat()
				time.Sleep(time.Second)
				
			}
		}
	}
