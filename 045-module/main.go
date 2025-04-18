package main

import (
	"fmt"
	"image"
	"io"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	//Stringer
	return fmt.Sprintf("%v (%v years).\n", p.Name, p.Age)
}

type IPAddr [4]byte

type MyError struct {
	When  time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s.\n", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	man := Person{"John", 34}
	fmt.Println(man)

	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 0},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		a := ip[0]
		b := ip[1]
		c := ip[2]
		d := ip[3]

		fmt.Printf("%v : {%v.%v.%v.%v}.\n", name, a, b, c, d)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Hello, World")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v.\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	pic := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(pic.Bounds())
	fmt.Println(pic.At(0, 0).RGBA())
}