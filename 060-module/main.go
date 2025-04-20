package main

// welcome to my world :) let's ~GO

import (
	"fmt"
	"sync"
	"time"
)

func main() {
// 	//goroutines
// 	var wg sync.WaitGroup //wait group is 0

// 	wg.Add(1) // wait group becomes 1
// 	go func ()  {
// 		defer wg.Done() //wait group waits for program to finish
// 		for i := 0; i < 5; i++ {
// 			time.Sleep(5 * time.Second)
// 			fmt.Println("first is processing")
// 		}
// 	}()

// 	wg.Add(1) // wait group becomes 2
// 	go func() {
// 		defer wg.Done() // wait  group waits for program to finish
// 		for i := 0; i < 5; i++ {
// 			time.Sleep(5 * time.Second)
// 			fmt.Println("second is processing")
// 		}
// 	}()

// 	wg.Wait() //wait group waits for all to finish and becomes 0 back
// 	fmt.Println("All execution are done") // final message after the programs ran in concurrency
// 

	// c := make(chan int) //created a channel

	// go func() {
	// 	sum := 0
	// 	for i := 1; i < 101; i++ {
	// 		sum += i
	// 	}
	// 	c <- sum // channel enters this routine and takes the data
	// }()

	// output := <-c // channel comes out back and return to data to the main program and also waits for the goroutine to finish

	// fmt.Printf("output: %d\n", output)


	//test own goroutine :)
	var wg sync.WaitGroup

	wg.Add(1)
	go cook("rice", &wg)

	wg.Add(1)
	go cook("beans", &wg)

	wg.Add(1)
	go cook("spag", &wg)

	wg.Wait()
	fmt.Println("Finished cooking all foods")
}

func cook(food string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("starting %s\n", food)
	time.Sleep(5 * time.Second)
	fmt.Printf("Finished %s\n", food)
}