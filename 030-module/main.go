package main

//welcome to my world :) lets ~GO

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X, Y int
}

type Vertex2 struct {
	Lat, Long float64
}

type Nest struct {
	Math int
	Eng int
}

var (
	v1 = Vertex{4,5}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

//a pointer function
func incrementValue(x *int) {
	*x = *x + 1 //modifies the value of point
}

//A function takes a fuction as params which takes 2 float64 type as params and returns a float64 type which then returns another float64 type
// :) :(
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	//pointers
	i, j := 23, 1400
	fmt.Println(i, j)

	point := 50
	fmt.Println("before point was", point)

	incrementValue(&point)
	fmt.Println("now point is", point)

	p := &i //points to i
	fmt.Println(*p)

	*p = 30 //changes i
	fmt.Println(i)

	p = &j //points to j
	*p = *p / 700 //changes j
	fmt.Println(j)

	k := 43
	l := new(int)
	*l = k
	fmt.Println(*l)

	//struct-s
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // v = {4, 2}

	a := &v
	a.Y = 4
	fmt.Println(a)

	fmt.Println(v1, v2, v3, v4)

	////////////////////////////////////////
	//Arrays
	var array1 [4]int
	array1[0] = 20
	array1[1] = 30
	fmt.Println(array1[0], array1[1])

	primes := [6]int {5,10,15,20,25,30}
	fmt.Println(primes)

	var splitPrims []int = primes[3:6] //slices the array
	fmt.Println((splitPrims))

	names := [4]string {
		"habeeb",
		"amoo",
		"olakunle",
		"whatever",
	}

	//takes their own portions
	formalNames := names[0:2]
	otherNames := names[1:3]
	whateverNames := names[3:4]

	otherNames[0] = "the-amoo" //changes the parents and reflects on children
	fmt.Println(formalNames, otherNames, whateverNames)

	//strict literals
	nums := []int {1,23,4,5,6,78}
	fmt.Println(nums)

	booleans := []bool {true, false, false, true, false}
	fmt.Println(booleans)

	evenNums := []struct {
		p int
		v bool
	} {
		{1, false},
		{2, true},
		{3, false},
		{4, true},
	}

	fmt.Println(evenNums)

	type PeopleAgeTypes struct { //works like typescript interfaces/types :)
		p string
		v int
	}

	myAge := PeopleAgeTypes{"habeeb", 17} //myAge is a type of PeopleAgeTypes :)
	fmt.Println(myAge)

	months := [12]string {"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

	firstQuater := months[0:4]
	secondQuater := months[4:8]
	thirdQuater := months[8:12]
	monthsRangeInMyHouse := months[2:9]
	fmt.Println(firstQuater, secondQuater, thirdQuater, monthsRangeInMyHouse)

	//loops through the arrays
	for  i := 0; i < len(months); i++ {
		fmt.Println(months[i])
	}

	firtSix := months[:6] //to 6th
	lastSix := months[6:] //from 6th
	allTwelve := months[:]
	fmt.Printf("First 6 %v\n", firtSix)
	fmt.Printf("Last 6 %v\n", lastSix)
	fmt.Printf("All 12 %v\n", allTwelve)

	//slice
	integers := []int {1,2,3,4,5,6,7,8,9,10}
	printSlice(integers)

	//now it has zero lenght but still obtains its capacity
	fewIntegers := integers[:0]
	printSlice(fewIntegers)

	//increased the length
	someIntegers := integers[:4]
	printSlice(someIntegers)

	//dropped its first 2 values
	manyIntegers := integers[2:] //now becomes 8 :(
	printSlice(manyIntegers)

	var sector []int
	fmt.Println(sector, len(sector), cap(sector))
	if sector == nil {
		fmt.Println("sector is nill")
	}

	//a multi dimensioanal array :)
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println(matrix)

	//make function
	langs := make([]int, 5) //creates an slice of 5-len
	frams := make([]int, 0, 5) //creates an slics 0-len 5-cap

	frams = frams[:cap(frams)] //prepopulates the slice to 5
	frams = frams[1:] //slice it to 4

	fmt.Println(langs)
	fmt.Println(frams)

	test := make([]int, 5)
	secondPrintSlice("test", test)

	test2 := make([]int, 0, 5)
	secondPrintSlice("test2", test2)

	replit_test2 := test2[:2]
	secondPrintSlice("replit_test2", replit_test2)


	//just decided to over complicate things here :)
	replit_test2_test2 := replit_test2[2:5]
	secondPrintSlice("replit_test2_test2", replit_test2_test2)


	//a tic tac toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//make turns
	board[0][0] = "X"
	board[0][2] = "O"
	board[1][1] = "X"
	board[1][2] = "O"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//append function
	var schools []int
	thirdPrintSlice(schools)

	//append works on nil slices
	schools = append(schools, 0)
	thirdPrintSlice(schools)

	//the slice grow as needed
	schools = append(schools, 1)
	thirdPrintSlice(schools)

	//i can add more than one elements at a time
	schools = append(schools, 2,3,4,5,6)
	thirdPrintSlice(schools)

	//looped through this array to prepopulate the schools array
	wannaBeSchools := []int{7,8,9,10,11,12}
	for i := 0; i < len(wannaBeSchools); i++ {
		schools = append(schools, wannaBeSchools[i])
	}
	thirdPrintSlice(schools)

	//range
	//range returns to values
	//1: the index
	//2: the value
	var pow = []int{1,2,4,8,16,32,64,128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d.\n", i, v)
	}

	for i, v := range pow {
		fmt.Printf("pow index %d = %d.\n", i, v)
	}

	service := make([]int, 10) //make the arry
	for i := range service { //loops through and populates the values
		service[i] = 1 << uint(i)
	}
	fmt.Println(service)

	for _, value := range service { //now it has access to the values
		fmt.Printf("%d.\n", value)
	}



	////////////////////////////////
	//maps

	var m = make(map[string]Vertex2)
	m["Bella Labs"] = Vertex2{
		40.64, 54.23,
	}
	fmt.Println(m["Bella Labs"])

	var n = map[string]Vertex2 {
		"Bella labs": Vertex2 {
			32.89, 90.34,
		},
		"Google": Vertex2{
			64.11, 35.14,
		},
	}

	fmt.Println(n)

	//i can omit here
	var mapped = map[string]Vertex2{ //type a string array of type Vertex2
		"First": {65.45, 89.0}, //omited vertex
	}
	fmt.Println(mapped)

	mapTest := make(map[string]int) //an array of prop-string : value-int
	mapTest["Answer"] = 10 //assigning the values
	mapTest["Answer"] = 15 //i can reassign :)

	delete(mapTest, "Answer") //delete elem : key
	fmt.Println(mapTest)
	mapTest["Answer"] = 20
	
	mapV, ok := mapTest["Answer"] //checks the array
	fmt.Println("The value: ", mapV, "Present? ",ok)


	//declaring an array of prop-string : value-string
	myMapValue := map[string]string{
		"Monday": "first",
		"Tuesday": "second",
	}

	//programs checks if wednesday exists on the array
	value, ok := myMapValue["Monday"]
	if ok {
		fmt.Println("wednesday exists")
	} else {
		fmt.Println("wednesday does not exists")
	}
	fmt.Println(value)

	//nested maps
	students := map[string]map[string]int {
		"Alice": {
			"maths": 78,
			"eng": 76,
		},
	}
	fmt.Printf("Alice maths score is %d\n", students["Alice"]["maths"])

	//nested maps with structs
	pupils := map[string]Nest {
		"Bob": {Math: 67, Eng: 87},
	}
	fmt.Printf("Bob's english score is %d\n", pupils["Bob"].Eng) //logic!!! :)


	////////////////////////////////////////
	//functions as a value
	test50 := returnFunc() //basic 😎
	fmt.Println(test50)

	//now passing the function as a value
	hypot := func (x, y float64) float64  {  //like js arrow func
		return  math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	
	fmt.Println(compute(hypot))

	//passing funcs as aguments
	fmt.Println(perform(addition, 10, 100))
}

func perform(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}

func addition(x, y int) int {
	return x + y
}

func returnFunc() string {
	return "yes"
}

func thirdPrintSlice(arr []int) {
	fmt.Printf("len=%d , cap=%d , %v.\n", len(arr), cap(arr), arr)
}

func secondPrintSlice (s string, x []int) {
	fmt.Printf("%s len=%d, cap=%d, %v.\n", s, len(x), cap(x), x)
}

func printSlice(x []int) {
	//len() length
	//cap() capacity
	fmt.Printf("len=%d , cap=%d , %v\n", len(x), cap(x), x)
}
