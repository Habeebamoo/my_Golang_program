package main

// welcome to my world :) let's ~GO

import (
	"errors"
	"fmt"
)

//generics
//checks for the index
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

//works on type int and float64
func sumAny[T int | float64](a, b T) T {
	return a + b
}

//generic structs
type Pair[T any] struct {
	First T
	Second T
}

var ErrInvalid = errors.New("cannot divide by zero") //defined my own errror

//error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrInvalid // manually throw an error
	}

	return a / b, nil
}

func main() {
	//Index works with slice if int
	si := []int{1,4,7,21,89,3}
	fmt.Println(Index(si, 21))

	//Index works with slice of strings also
	words := []string{"foo", "bar", "wolf", "cup"}
	fmt.Println(Index(words, "foo"))

	//sumAny func works cuz of the generics
	fmt.Println(sumAny(2, 3))
	fmt.Println(sumAny(4.4, 5.7))
	fmt.Println(sumAny(5, 6.8))

	//generics structs
	p := Pair[int]{First: 1, Second: 5}
	s := Pair[string]{First: "a", Second: "b"}
	fmt.Println(p, s)

	//error handling
	result, err := divide(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result: ", result)
}