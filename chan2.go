package main

import (
	"fmt"
)

func main(){
	jobs := make(chan int, 66)
	results := make(chan int, 50)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	
	
	for i:=0; i<100; i++ {
		jobs <-i
	}

	fibList := [] int{}

	close(jobs)

	for j:=0; j<100; j++ {
		fibList = append(fibList, <-results)

		fmt.Printf("%v - %v\n", j, fibList[j] )
		//fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs{
		//fmt.Printf("This is n in worker func%v\n", n)
		results <- fib(n)
	}
}

func fib(n int) int{
	
	if n <=1 {
		return n
	}
	//fmt.Printf("N inside worker %v\n", n)
	return fib(n - 1) + fib(n - 2)
}