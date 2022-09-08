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

	if(checkForkBoth(g.forkright, g.forkleft)){
		fmt.Println(fmt.Sprint("guy number ", number, " is eating"))
		
		resetforks(g.forkright, g.forkleft)

				fmt.Println(fmt.Sprint("guy number ", number, " is full"))
		g.forkleft.free <- true
		g.forkright.free <- true		

	} else {fmt.Println(fmt.Sprint("guy number ", number, " is thinking"))}
}

func  checkForkBoth(forkyright *fork, forkyleft *fork) bool{
	forkyright.free <- forkyright.status
	availabiliyright := <- forkyright.free

	forkyleft.free <- forkyleft.status
	availabiliyleft := <- forkyleft.free
	
	if(availabiliyright && !availabiliyleft){
		forkyright.free <- true
	}
	if(availabiliyleft && !availabiliyright){
		forkyright.free <- true
	}
	if(!availabiliyleft && !availabiliyright){
		forkyright.free <- false
		forkyleft.free <- false
	}
	if(availabiliyleft && availabiliyright){
		forkyright.free <- false
		forkyleft.free <- false
		return true
	}
	return false
}

func resetforks(forkyright *fork, forkyleft *fork){
	x1 := <- forkyright.free 
	x2 := <- forkyleft.free 

	fmt.Println(fmt.Sprint(x1, " and ", x2, " are now reset" ))

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

		for i := 0; i<3; i++{
			for _, v := range guys {
				go v.eat()
			}
		}
	}
