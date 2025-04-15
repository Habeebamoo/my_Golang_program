package main

//welcome to my world :) lets ~GO

import (
	"fmt"
)

type Vertex struct {
	X, Y int
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
}

func printSlice(x []int) {
	//len() length
	//cap() capacity
	fmt.Printf("len=%d , cap=%d , %v\n", len(x), cap(x), x)
}