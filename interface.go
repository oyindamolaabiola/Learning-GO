package main

import "fmt"

// type Greetings interface {
// 	Time() string
// }

// type Morning struct {}

// func (m Morning) Time() string {
// 	return "Good morning, Oyindamola"
// }

// type Afternoon struct {}

// func (a Afternoon) Time() string {
// 	return "Good afternoon, Oyindamola"
// }

// func main() {
// 	// define a struct type variable because the function Time is 
// 	// a receiver function and can take only a struct type variable
// 	// so to execute the function we need to create an instance of the struct
// 	// type (Morning & Afternoon)
// 	m := Morning{}
// 	a := Afternoon{}

// 	fmt.Println(m.Time())
// 	fmt.Println(a.Time())
// }



// EXAMPLE 2

type Measurements interface {
	Area()
	Circumference()
}

type Rectangle struct {
	length float32
	breath float32
}

type Circle struct {
	pi float32
	radius float32
}

func (r Rectangle) Area() float32 {
	return r.breath * r.length
}

func (c Circle) Circumference() float32 {
	r := 2 * c.radius 
	return r * c.pi
}

func main() {
	r := Rectangle{length: 3.1, breath: 2.5}
	c := Circle{pi: 2.24, radius: 0.9}

	fmt.Println(r.Area())
	fmt.Println(c.Circumference())
}
