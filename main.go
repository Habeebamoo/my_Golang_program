package main

//welcome to my world

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//an addition function
func add(x, y int) int {
	return x + y
}

//a swap function
func swap(a, b string) (string, string) {
	return b, a
}

func split(num int) (x, y int) {
	x = num - 10
	y = num - x
	return
}

//ssa := "ho"
// variables: doesnt work here

//but works here :) 😁
var (
	toBe bool = false
	myAge int = 17

	//false values by default
	val1 bool
	val2 string
	val3 int
)

//iota for eums
const (
	sunday = iota
	monday
	tuesday
)

func main() {
	//Hello world
	fmt.Println("Hello world")

	//Time
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	//Next week
	nextWeek := currentTime.Add(8 * 24 * time.Hour)
	nextWeekFormatted := nextWeek.Format("Monday")
	fmt.Println("Next week is", nextWeekFormatted)

	//Next month
	nextMonth := currentTime.Add(30 * 24 * time.Hour)
	nextMonthFormatted := nextMonth.Format("3, Jan 2006")
	fmt.Println("Next month is", nextMonthFormatted)

	//Clock time
	clockTime := currentTime.Format("15-04-05 PM")
	fmt.Println("Clock Time is", clockTime)

	//todays date
	todaysDate := currentTime.Format("2, Jan 2006")
	fmt.Println("Todays date is", todaysDate)

	//math function
	//numbers between 0 and 9
	fmt.Println("My favorite number is", rand.Intn(10))

	//my own rand range logic :)
	min := 10
	max := 50
	randInt := rand.Intn(max - min + 1) + min //(50 - 40 + 1) = Intn(40) + 10 😁
	fmt.Println("Logic number from 10 to 50 is:", randInt)

	//format verbs
	result := math.Sqrt(48)
	result2 := "habeeb"
	fmt.Printf("type %T, Value: %g problems. \n", result2, result)
	fmt.Println(math.Pi)


	///////////////////////////////////////////////////////////
	//functions

	fmt.Println(add(10, 20))

	firstName := "Habeeb"
	lastName := "Amoo"
	fmt.Println(swap(firstName, lastName))

	fmt.Println(split(18))

	////////////////////////////////////////////////////////////
	//variables

	var c, python, rust bool //false by default
	fmt.Println(c, python, rust)

	var java, Go bool = true, true // now it takes its assigned value
	fmt.Println(java, Go)

	var i, j int = 1, 2
	fmt.Println(i, j)

	k := "habeeb" //dosent need var, assign itself (and adds type) :)
	fmt.Println(k)
	fmt.Println(toBe)

	//varaibles with verbs
	fmt.Printf("Type of toBe: %T actual value: %v \n", toBe, toBe)
	fmt.Printf("Type of myAge: %T , myAge is: %d \n", myAge, myAge)
	fmt.Printf("%v %v %v \n", val1, val2, val3) // second parameter dosent show :)
	//val2 = string ""

	var opp, adj int = 3, 4
	var hyp float64 = math.Sqrt(float64(opp*opp + adj*adj))
	fmt.Printf("opposite: %v , adjacent: %v , hypotenus: %v \n", opp, adj, hyp)

	//print safety (convert to positive integers) :(
	var safeHyp uint = uint(hyp)
	fmt.Printf("opposite: %v , adjacent: %v , safe hypotenus: %v \n", opp, adj, safeHyp)	

	//type inference
	text := 43 //change me :)
	fmt.Printf("text is a type of %T\n", text)

	//constant variables
	const myConstantvariable = "olakunle"
	fmt.Println(myConstantvariable)

	//enums
	fmt.Println(sunday, monday, tuesday)
}