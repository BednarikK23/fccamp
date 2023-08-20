package main

import "math"

// are collections of method signatures.
// A type "implements" an interface if it has all the methods of the given interface defined on it.

// any type that implements the these two methods and matches both of the methods signatures
// will implement the shape interface like so:
type shape interface {
	area() float64
	perimeter() float64
}

// types can fulfill more interfaces - circle and rect fulfill this one and shape also...
type yourMom interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (f fullTime) getName() string {
	return f.name
}
func (f fullTime) getSalary() int {
	return f.salary
}

func test(e employee) {
	println(e.getSalary(), e.getName())
	println("==========================================")
}

func shapeInfo(s shape) (float64, float64) {
	// TYPE ASSERTION
	// sometimes you need to know what type is the object of, and you have only its interface, so:
	// "r" is new "rect" cast from "s" - instance of "shape"
	// "ok" is a bool that is true if "r" was "rect" or false otherwise
	r, ok := s.(rect)
	if ok {
		return r.height, r.width
	}
	c, ok := s.(circle)
	if ok {
		return c.radius, -1
	}
	return -1, -1
}

func shapeSwitchInfo(s shape) (float64, float64) {
	// SWITCH + type
	// same as up only using switch:
	switch v := s.(type) {
	case rect:
		return v.width, v.height
	case circle:
		return v.radius, -1
	default:
		return -1, -1
	}

}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 500_000, // this is allowed same as python, happy
	})
	test(contractor{
		name:         "Daniels",
		hourlyPay:    20,
		hoursPerYear: 9000,
	})

	// TYPE ASSERTION
	rt := rect{width: 12.5, height: 85.44}
	cic := circle{4566}
	shapeInfo(rt)
	shapeSwitchInfo(cic)

}
