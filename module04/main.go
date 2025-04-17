package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Person struct {
	Age int
}

type MyFloat float64

type Dev struct {
	Job string
	Lang string
}

//methods in functions
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (p Person) GetAge() float64 {
	return float64(p.Age + 1)
}

func (f MyFloat) GetAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) { //pointer address must be included
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Object() Vertex {
	return v
}

//tested B & A
func (d Dev) Prop() Dev {
	return Dev{d.Job, d.Lang}
}

func (d *Dev) UpdateLang(lang string) {
	d.Lang = lang
}

func funcAbs(v Vertex) Vertex {
	return Vertex{v.X, v.Y}
}

func funcScale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) PointAbs() Vertex {
	return Vertex{v.X, v.Y}
}

func (v *Vertex) PointScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Abser interface {
	AbserAbs() Vertex // Abser can hold an value that has this method
}

//types for methods
type I interface {
	M() //the method
}

type T struct {
	S string
}

type SchoolMethod interface {
	GetSchool()
}

type School struct {
	Name string
}

func (s School) GetSchool() {
	fmt.Println(s)
}

func (t T) M() {
	fmt.Println(t.S)
}

type HomeMethod interface {
	GetHome() string
}

type Home struct{
	Country string
}

func (h Home) GetHome() string {
	return h.Country
}

type Inter interface {
	Meth()
}

type Text struct {
	s string
}

type Inter2 interface {
	CallMeth()
}

type String struct {
	s string
}

func (s *String) CallMeth() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(s.s)
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.Abs())

	habeeb := Person{17}
	fmt.Println(habeeb.GetAge())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.GetAbs())

	num := Vertex{5,3}
	num.Scale(2)
	fmt.Println(num.Object())

	//before
	mike := Dev{"Web Developer", "JavaScript"}
	fmt.Println(mike.Prop())
	
	//after
	mike.UpdateLang("Typescript")
	fmt.Println(mike.Prop())

	//methods replaced with functions
	graph := Vertex{10,20}
	funcScale(&graph, 2)
	fmt.Println(funcAbs(graph))

	vector := &Vertex{50, 100} //points here thats why pointScale takes a pointer
	vector.PointScale(2)
	fmt.Println(vector.PointAbs())


	//defining the type of method
	var i I = T{"hello"}
	i.M()
	
	var defaultSchool SchoolMethod = School{"j"}
	fmt.Println(defaultSchool)

	var secondSchool SchoolMethod = School{"NYC"}
	fmt.Println(secondSchool)

	var myHome HomeMethod = Home{"Niger"}
	fmt.Println(myHome.GetHome())

	var variab Inter
	variab = &Text{"This is a server"}
	describe(variab) //return the (value, type) of ~variab~
	variab.Meth()

	var try Inter2
	var say *String
	try = say
	try.CallMeth() //no value so my func returns nil
	describeAgain(try)

	try = &String{"Present"} //now has a value cuz it has been assigned
	try.CallMeth()
	describeAgain(try)

	var testInter interface{}
	test(testInter) //nil, nil

	testInter = 42
	test(testInter) //42, int

	testInter = "hello"
	test(testInter)//hello, string


	//////////////////////////////////////////////////
	//type assertions

	var typeAssert interface{} = "hello"

	check1, ok := typeAssert.(string) //checks the type
	fmt.Println(check1, ok)

	do(12)
	do("hello")
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("value is an %T.\n", v)
	case string:
		fmt.Printf("values is a %T.\n", v)
	default:
		fmt.Println("I dont know the type")
	}
}

func test(i interface{}) { //return the value and type
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeAgain(t Inter2) {
	fmt.Printf("(%v, %T)\n", t, t)
}

func describe(i Inter) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (t *Text) Meth() {
	fmt.Println(t)
}
