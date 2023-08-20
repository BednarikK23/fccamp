package main

import "fmt"

// structs are pretty same as in C
type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel // We can nest structs into each oother...
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material int
}

type messageToSend struct {
	message   string
	sender    user // NESTED
	recipient user
}

type user struct {
	name   string
	number int
}

// EMBEDDED STRUCTS
// not same as nested, we take field from one struct and shave them into another struct
type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all the fields, of the car struct
	// even if go can do this, it is just syntactic sugar, and go is not object-oriented pl
	car
	bedSize int
}

type rectangle struct {
	width  int
	height int
}

// METHODS ON STRUCTS (even tho Go is not O-O programming language! - just syntactic sugar)
// done using this special parameter before name of function
func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	myCar := car{} // simple declaration all intialized with 0
	myCar.FrontWheel.Radius = 0

	sender := user{"me", 770}
	receiver := user{"you", 152}
	mess := messageToSend{"bla", sender, receiver}
	if canSend(mess) {
		fmt.Println("send")
	}

	// ANONYMOUS STRUCT
	// type of struct does not have a name
	// prevent you from re-using a struct definition you never intended to re-use :))
	currCar := struct {
		Make  string // no ,
		Model string
	}{
		Make:  "tesla", // ,
		Model: "model 3",
	}

	println(currCar.Make, currCar.Model)

	// EMBEDDED
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			Make:  "toyota",
			Model: "camry",
		},
	}
	// it's not lanesTruck.car.Model it shortens it to just this, and it is really good...
	fmt.Println(lanesTruck.Model, lanesTruck.Make) // HERE!!

	// METHODS ON STRUCT
	r := rectangle{40, 50}
	fmt.Println(r.area())
	if mess.betterMTS() {
		fmt.Println("send")
	}
}

func (u user) authentication() bool {
	return u.name != "" && u.number != 0
}

func (m messageToSend) betterMTS() bool {
	return m.sender.authentication() && m.recipient.authentication() && m.message != ""
}

func canSend(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.recipient.name == "" {
		return false
	}

	if mToSend.sender.number != 0 || mToSend.recipient.number != 0 {
		return false
	}
	return true
}
