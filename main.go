package main

import ("fmt"
"sync")

var wg *sync.WaitGroup
var limit = 100

func main(){

	wg = new(sync.WaitGroup)
	even := make(chan int,1)
	odd := make(chan int,1)

	wg.Add(2)

	go oddFunc(odd,even)
	go evenFunc(odd,even)
	
	even <- 1

	wg.Wait()
}

func oddFunc(odd chan int,even chan int){

	for i:=1;i<=limit;i++ {
		<- even
		if i%2!=0{
			fmt.Println("odd_go_routine:",i)	
		}
		odd <- 1
	}

	wg.Done()
}

func evenFunc(odd chan int, even chan int){

	for i:=1;i<=limit;i++ {
		<- odd
		if i%2==0{
			fmt.Println("even_go_routine:",i)	
		}
		even <- 1
	}
	
	wg.Done()
}
