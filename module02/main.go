package main

// welcome to my world :) let's ~GO

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	//if else
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// v declared
	if v := math.Pow(x, n); v >= lim {
		return v
	} else {
		fmt.Printf("%g <= %g\n", v, lim)
	}
	//cant use v here
	return lim

	//v := math.Pow(x, n)
	//if v >= lim {
	//return v
	//}
	//return lim
}


func main() {
	//basic for loop
	sum := 1
	for i := 0; i < 3; i++ {
		sum += i
	}
	fmt.Println(sum) 

	//a while loop (but for acts as while) :(
	age := 1
	for age < 2 {
		age += age
	}
	fmt.Println(age)

	//infinite loop
	//for {
	// fmt.Println("loop forever :)")
	//}

	//basic if else if else statements
	userAge := 51
	if userAge < 18 {
		println("Programmer is young")
	} else if userAge < 50 {
		println("Programmer is a youth")
	} else {
		println("Programmer is a senior citizen")
	}

	//declare a local variable in an if else statement
	if myAge := 25; myAge < 18 {
		fmt.Printf("%d is a teen age.\n", myAge)
	} else {
		fmt.Printf("%d is above Teenagehood.\n", myAge)
	}

	//calling the if else function
	fmt.Println(sqrt(-5))
	fmt.Println(
		pow(2, 3, 9),
		pow(5, 2, 26),
	)

	//switch statement
	fmt.Println("Go runs on")
	//you can declare it here ---- os := runtime.GOOS
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		//freebsd, openbsd, plan9, windows
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When is saturday")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Next Tomorrow")
	default:
		fmt.Println("Too Far Away")
	}

	//using multiple values in case
	devAge := 16
	switch devAge {
	case 2,3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Dev is a kid")
	case 13,14,15,16,17,18,19:
		fmt.Println("Dev is a teen")
	case 20:
		fmt.Println("Dev is an adult")
	}

	//using switch instead of if else (no condition acts like if else)
	t := time.Now()
	switch { // no condition for switch means (true)
	case t.Hour() < 12:
		fmt.Println("Good Morining")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	//defering a function
	fmt.Println("First")
	//defer fmt.Println("Second") comes last :)
	fmt.Println("Third")

	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //prints backward
	}
	fmt.Println("done")
}