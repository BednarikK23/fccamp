package main

// Declare a main package (a package is a way to group functions,
// and it's made up of all the files in the same directory).

import (
	"errors"
	"fmt"
)

// contains functions for formatting text, including printing to the console.
// one of the standard library packages you got when you installed Go

/* POSSIBLE TYPES:
- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- byte  // alias for uint8
- rune  // alias for int32
		// represents unicode code point
- float32, float64
- complex64, complex128
*/

func main() {
	// declaration, two types:
	// short - go will automatically add type to the variable
	// for short variable declaration is used := operator, automatically inferring the type based on the assigned value
	// := It is only used within functions or local scopes, not at the package level
	congrats := "happy birthday"
	someInt := 45
	fmt.Println(congrats, someInt) // example println: fmt.Println(name, "is", age, "years old.")

	// long - u can specify type of variable even the size - like in C fe
	var smsLimit int // beautifully readable variable smsLimit is type int...
	var text string
	var days uint8

	// single line, multi init possible:
	smsLimit, text, days = 15, "Sms limit is: ", 30
	days++ // same as in c...

	fmt.Printf("%s%d in %d days.\n", text, smsLimit, days)
	/*
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
		default variable:		 %v - when you are unsure of the type u can use this, but it is better to use specific
	*/

	// const are immutable, and their value have to be known in compile time
	// so, we can compute value of fullName, but it has to be competed from values
	// that are known before we compile program...
	const firstName = "Krystof"
	const lastName = "Bednarik"
	const fullName = firstName + " " + lastName

	// ifs don't require () around conditions
	// you can assign (or do some operation) before condition then add; and after that put condition:
	if length := len(firstName); length < 0 {
		fmt.Println("first name is not valid")
	} else if length > 10 {
		fmt.Println("first name too long change your fuckin name lol")
	} else {
		fmt.Println("first name valid", firstName)
	}
	// it also removes length from parent's scope

	// FUNCTIONS ________________________________________________
	decrement(sub, 50)
	increment(sub, 50)
}

// now type is: f func(func(func(int, int) int, int) int
func decrement(sub func(int, int) int, num int) int {
	return sub(num, 1)
}

// so could do ALIAS:
type callBack func(int, int) int

func increment(sub callBack, num int) int {
	return sub(num, -1)
}

// functions are really easy as all other things in go:
// types are after names for better reading
func sub(x int, y int) int {
	return x - y
}

// we can return tuple
func getCoords() (int, int) {
	return 10, 20
}

// ERRORS
// first of all we can type like this and both have type of int...
func divide(dividend, divisor int) (int, error) { // error is special interface
	if divisor == 0 {
		return 0, errors.New("Cannot devide by zero")
	}
	return dividend / divisor, nil // -> nil means no error
}

func concat(s1 string, s2 string) string {
	return s1 + s2 // the strings in go can be normally concat like so
}

func strCmp(s1 string, s2 string) bool {
	return s1 == s2 // the strings in go can be normally checked on equality like that
}
