package main

import (
	"fmt"
	"time"

)

type fork struct {
	id   int
	free  chan bool
	status bool
}
type guy struct {
	id                  int
	forkleft, forkright *fork
}

func (g guy) eat() {
	number := g.id

	if(checkFork(g.forkright) && checkFork(g.forkleft)){
		fmt.Println("its working")
	} else {fmt.Println("its not working")}

	//checkFork(chright, chleft, g.forkright,g.forkleft)

	//if checkbooking(g.forkleft.free) {
		if(1==1){
		fmt.Println(fmt.Sprint("guy number ", number, " is eating"))
		
		time.Sleep(time.Second)
	} else {
		fmt.Println(fmt.Sprint("guy number ", number, " is thinking"))
		time.Sleep(time.Second)
		return
	}
	//fmt.Println(fmt.Sprint("guy number ", number, " is done eating"))
}


func  checkFork(forky *fork) bool{
	forky.free <- forky.status
	availabiliy := <- forky.free
	return availabiliy
}

func main() {
	begin()
	time.Sleep(time.Second)

}

func begin() {
	//making 5 forks with unique ID's and a boolean if they're taking or not
	forks := make([]*fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &fork{
			id: i, free: make(chan bool,1), status: true,
		} }
		
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
